//go:build linux
// +build linux

package cli

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/essentialkaos/ek/v12/env"
	"github.com/essentialkaos/ek/v12/fmtc"
	"github.com/essentialkaos/ek/v12/fmtutil"
	"github.com/essentialkaos/ek/v12/fsutil"
	"github.com/essentialkaos/ek/v12/mathutil"
	"github.com/essentialkaos/ek/v12/options"
	"github.com/essentialkaos/ek/v12/strutil"
	"github.com/essentialkaos/ek/v12/system"
	"github.com/essentialkaos/ek/v12/system/process"
	"github.com/essentialkaos/ek/v12/terminal/window"
	"github.com/essentialkaos/ek/v12/usage"
	"github.com/essentialkaos/ek/v12/usage/completion/bash"
	"github.com/essentialkaos/ek/v12/usage/completion/fish"
	"github.com/essentialkaos/ek/v12/usage/completion/zsh"
	"github.com/essentialkaos/ek/v12/usage/man"
	"github.com/essentialkaos/ek/v12/usage/update"

	"github.com/essentialkaos/swptop/cli/support"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	APP  = "swptop"
	VER  = "0.6.4"
	DESC = "Utility for viewing swap consumption of processes"
)

const (
	OPT_USER     = "u:user"
	OPT_FILTER   = "f:filter"
	OPT_NO_COLOR = "nc:no-color"
	OPT_HELP     = "h:help"
	OPT_VER      = "v:version"

	OPT_VERB_VER     = "vv:verbose-version"
	OPT_COMPLETION   = "completion"
	OPT_GENERATE_MAN = "generate-man"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// ProcessInfo contains basic info about process
type ProcessInfo struct {
	PID     int
	VmSwap  uint64
	User    string
	Command string
}

// ProcessInfoSlice is ProcessInfo slice
type ProcessInfoSlice []ProcessInfo

// ////////////////////////////////////////////////////////////////////////////////// //

func (s ProcessInfoSlice) Len() int           { return len(s) }
func (s ProcessInfoSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ProcessInfoSlice) Less(i, j int) bool { return s[i].VmSwap < s[j].VmSwap }

// ////////////////////////////////////////////////////////////////////////////////// //

// optMap is map with options
var optMap = options.Map{
	OPT_USER:     {},
	OPT_FILTER:   {},
	OPT_NO_COLOR: {Type: options.BOOL},
	OPT_HELP:     {Type: options.BOOL},
	OPT_VER:      {Type: options.BOOL, Alias: "ver"},

	OPT_VERB_VER:     {Type: options.BOOL},
	OPT_COMPLETION:   {},
	OPT_GENERATE_MAN: {Type: options.BOOL},
}

// useRawOutput is raw output flag
var useRawOutput bool

// winWidth is current window width
var winWidth int

// cmdEllipsis command ellipsis size
var cmdEllipsis = 64

// ////////////////////////////////////////////////////////////////////////////////// //

// Init is main function
func Init(gitRev string, gomod []byte) {
	_, errs := options.Parse(optMap)

	if len(errs) != 0 {
		for _, err := range errs {
			printError(err.Error())
		}

		os.Exit(1)
	}

	if options.GetB(OPT_NO_COLOR) {
		fmtc.DisableColors = true
	}

	switch {
	case options.Has(OPT_COMPLETION):
		os.Exit(genCompletion())
	case options.Has(OPT_GENERATE_MAN):
		os.Exit(genMan())
	case options.GetB(OPT_VER):
		showAbout(gitRev)
		return
	case options.GetB(OPT_VERB_VER):
		support.ShowSupportInfo(APP, VER, gitRev, gomod)
		return
	case options.GetB(OPT_HELP):
		showUsage()
		return
	}

	if useRawOutput {
		printRawTop()
	} else {
		printPrettyTop()
	}
}

// configureUI configures user interface
func configureUI() {
	envVars := env.Get()
	term := envVars.GetS("TERM")

	fmtc.DisableColors = true

	if term != "" {
		switch {
		case strings.Contains(term, "xterm"),
			strings.Contains(term, "color"),
			term == "screen":
			fmtc.DisableColors = false
		}
	}

	if options.GetB(OPT_NO_COLOR) {
		fmtc.DisableColors = true
	}

	if !fsutil.IsCharacterDevice("/dev/stdout") && envVars.GetS("FAKETTY") == "" {
		fmtc.DisableColors = true
		useRawOutput = true
	}

	if !useRawOutput {
		fmtutil.SeparatorFullscreen = true
		winWidth = window.GetWidth()
	}

	if winWidth > 110 {
		cmdEllipsis = winWidth - 40
	}
}

// printPrettyTop prints info with separators and headers
func printPrettyTop() {
	procInfo, memUsage, err := collectInfo()

	if err != nil {
		printErrorAndExit(err.Error())
	}

	if len(procInfo) == 0 && memUsage.SwapUsed == 0 {
		fmtc.Println("{g}Can't find any process with swap usage{!}")
		return
	}

	fmtc.NewLine()

	if len(procInfo) != 0 {
		printPrettyProcessList(procInfo)
	}

	fmtutil.Separator(true)
	fmtc.NewLine()

	printOverallInfo(procInfo, memUsage)

	fmtc.NewLine()
	fmtutil.Separator(true)

	fmtc.NewLine()
}

// printPrettyProcessList prints info about swap usage by processes
func printPrettyProcessList(procInfo ProcessInfoSlice) {
	fmtutil.Separator(true)

	fmtc.Printf(
		" {*}%5s{!} {s}|{!} {*}%16s{!} {s}|{!} {*}%8s{!} {s}|{!} {*}%-s{!}\n",
		"PID", "USERNAME", "SWAP", "COMMAND",
	)

	fmtutil.Separator(true)

	for _, pi := range procInfo {
		fmtc.Printf(
			" %5d {s}|{!} %16s {s}|{!} %8s {s}|{!} %-s\n",
			pi.PID, pi.User, fmtutil.PrettySize(pi.VmSwap),
			strutil.Ellipsis(pi.Command, cmdEllipsis),
		)
	}
}

// printOverallInfo prints overall swap usage info
func printOverallInfo(procInfo ProcessInfoSlice, memUsage *system.MemUsage) {
	var procUsed uint64
	var procUsedPerc float64

	if len(procInfo) != 0 {
		procUsed = calculateUsage(procInfo)
		procUsedPerc = (float64(procUsed) / float64(memUsage.SwapTotal)) * 100.0
		procUsedPerc = mathutil.BetweenF(procUsedPerc, 0.0001, 100.0)
	}

	overallUsed := memUsage.SwapUsed

	// Procfs cannot show values less than 1kb, so we have use calculated processes usage
	if procUsed > memUsage.SwapUsed {
		overallUsed = procUsed
	}

	overallUsedPerc := (float64(overallUsed) / float64(memUsage.SwapTotal)) * 100.0
	overallUsedPerc = mathutil.BetweenF(overallUsedPerc, 0.0001, 100.0)

	if len(procInfo) == 0 || math.IsNaN(procUsedPerc) {
		fmtc.Println("  {*}Processes:{!} n/a")
	} else {
		fmtc.Printf(
			"  {*}Processes:{!} %s {s-}(%s){!}\n",
			fmtutil.PrettySize(procUsed),
			fmtutil.PrettyPerc(procUsedPerc),
		)
	}

	if math.IsNaN(overallUsedPerc) {
		fmtc.Println("  {*}Overall:{!} n/a")
	} else {
		fmtc.Printf(
			"  {*}Overall:{!}   %s {s-}(%s){!}\n",
			fmtutil.PrettySize(overallUsed),
			fmtutil.PrettyPerc(overallUsedPerc),
		)
	}

	fmtc.Printf("  {*}Total:{!}     %s\n", fmtutil.PrettySize(memUsage.SwapTotal))
}

// printRawTop just prints raw info
func printRawTop() {
	procInfo, _, err := collectInfo()

	if err != nil {
		printErrorAndExit(err.Error())
	}

	if len(procInfo) == 0 {
		return
	}

	for _, pi := range procInfo {
		fmt.Printf("%d %s %d %s\n", pi.PID, pi.User, pi.VmSwap, pi.Command)
	}
}

// collectInfo collects info about processes and sort result slice
func collectInfo() (ProcessInfoSlice, *system.MemUsage, error) {
	memInfo, err := system.GetMemUsage()

	if err != nil {
		return nil, nil, err
	}

	procInfo, err := getProcessesSwapUsage()

	if err != nil {
		return nil, nil, err
	}

	return procInfo, memInfo, err
}

// ignoreInfo returns true if we must ignore this info
func ignoreInfo(info ProcessInfo) bool {
	if options.Has(OPT_USER) {
		if info.User != options.GetS(OPT_USER) {
			return true
		}
	}

	if options.Has(OPT_FILTER) {
		if !strings.Contains(info.Command, options.GetS(OPT_FILTER)) {
			return true
		}
	}

	return false
}

// getProcessesSwapUsage returns slice with info about swap usage by processes
func getProcessesSwapUsage() (ProcessInfoSlice, error) {
	processes, err := process.GetList()

	if err != nil {
		fmt.Errorf("Can't collect info about processes")
	}

	var result ProcessInfoSlice

	for _, pi := range processes {
		if pi.IsThread {
			continue
		}

		memInfo, err := process.GetMemInfo(pi.PID)

		if err != nil || memInfo.VmSwap == 0 {
			continue
		}

		info := ProcessInfo{
			PID:     pi.PID,
			User:    pi.User,
			Command: pi.Command,
			VmSwap:  memInfo.VmSwap,
		}

		if !ignoreInfo(info) {
			result = append(result, info)
		}
	}

	sort.Sort(sort.Reverse(result))

	return result, nil
}

// calculateUsage calculates total swap usage
func calculateUsage(info ProcessInfoSlice) uint64 {
	var result uint64

	for _, processInfo := range info {
		result += processInfo.VmSwap
	}

	return result
}

// printError prints error message to console
func printError(f string, a ...interface{}) {
	fmtc.Fprintf(os.Stderr, "{r}"+f+"{!}\n", a...)
}

// printError prints warning message to console
func printWarn(f string, a ...interface{}) {
	fmtc.Fprintf(os.Stderr, "{y}"+f+"{!}\n", a...)
}

// printErrorAndExit prints error mesage and exit with exit code 1
func printErrorAndExit(f string, a ...interface{}) {
	printError(f, a...)
	os.Exit(1)
}

// ////////////////////////////////////////////////////////////////////////////////// //

// showUsage prints usage info
func showUsage() {
	genUsage().Render()
}

// showAbout prints info about version
func showAbout(gitRev string) {
	genAbout(gitRev).Render()
}

// genCompletion generates completion for different shells
func genCompletion() int {
	info := genUsage()

	switch options.GetS(OPT_COMPLETION) {
	case "bash":
		fmt.Printf(bash.Generate(info, "swptop"))
	case "fish":
		fmt.Printf(fish.Generate(info, "swptop"))
	case "zsh":
		fmt.Printf(zsh.Generate(info, optMap, "swptop"))
	default:
		return 1
	}

	return 0
}

// genMan generates man page
func genMan() int {
	fmt.Println(
		man.Generate(
			genUsage(),
			genAbout(""),
		),
	)

	return 0
}

// genUsage generates usage info
func genUsage() *usage.Info {
	info := usage.NewInfo()

	info.AddOption(OPT_USER, "Filter output by user")
	info.AddOption(OPT_FILTER, "Filter output by part of command")
	info.AddOption(OPT_NO_COLOR, "Disable colors in output")
	info.AddOption(OPT_HELP, "Show this help message")
	info.AddOption(OPT_VER, "Show version")

	info.AddExample("", "Show current swap consumption of all processes")
	info.AddExample("-u redis", "Show current swap consumption by webserver user processes")
	info.AddExample("-f redis-server", "Show current swap consumption by processes with 'redis-server' in command")
	info.AddExample("| wc -l", "Count number of processes which use swap")

	return info
}

// genAbout generates info about version
func genAbout(gitRev string) *usage.About {
	about := &usage.About{
		App:           APP,
		Version:       VER,
		Desc:          DESC,
		Year:          2006,
		Owner:         "ESSENTIAL KAOS",
		License:       "Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>",
		UpdateChecker: usage.UpdateChecker{"essentialkaos/swptop", update.GitHubChecker},
	}

	if gitRev != "" {
		about.Build = "git:" + gitRev
	}

	return about
}
