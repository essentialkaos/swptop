//go:build linux
// +build linux

package cli

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2024 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strings"

	"github.com/essentialkaos/ek/v13/fmtc"
	"github.com/essentialkaos/ek/v13/fmtutil"
	"github.com/essentialkaos/ek/v13/mathutil"
	"github.com/essentialkaos/ek/v13/options"
	"github.com/essentialkaos/ek/v13/strutil"
	"github.com/essentialkaos/ek/v13/support"
	"github.com/essentialkaos/ek/v13/support/deps"
	"github.com/essentialkaos/ek/v13/support/kernel"
	"github.com/essentialkaos/ek/v13/support/resources"
	"github.com/essentialkaos/ek/v13/system"
	"github.com/essentialkaos/ek/v13/system/process"
	"github.com/essentialkaos/ek/v13/terminal"
	"github.com/essentialkaos/ek/v13/terminal/tty"
	"github.com/essentialkaos/ek/v13/usage"
	"github.com/essentialkaos/ek/v13/usage/completion/bash"
	"github.com/essentialkaos/ek/v13/usage/completion/fish"
	"github.com/essentialkaos/ek/v13/usage/completion/zsh"
	"github.com/essentialkaos/ek/v13/usage/man"
	"github.com/essentialkaos/ek/v13/usage/update"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	APP  = "swptop"
	VER  = "1.1.1"
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

// optMap is map with options
var optMap = options.Map{
	OPT_USER:     {},
	OPT_FILTER:   {},
	OPT_NO_COLOR: {Type: options.BOOL},
	OPT_HELP:     {Type: options.BOOL},
	OPT_VER:      {Type: options.MIXED},

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
	preConfigureUI()

	_, errs := options.Parse(optMap)

	if !errs.IsEmpty() {
		terminal.Error("Options parsing errors:")
		terminal.Error(errs.Error(" - "))
		os.Exit(1)
	}

	configureUI()

	switch {
	case options.Has(OPT_COMPLETION):
		os.Exit(printCompletion())
	case options.Has(OPT_GENERATE_MAN):
		printMan()
		os.Exit(0)
	case options.GetB(OPT_VER):
		genAbout(gitRev).Print(options.GetS(OPT_VER))
		os.Exit(0)
	case options.GetB(OPT_VERB_VER):
		support.Collect(APP, VER).
			WithRevision(gitRev).
			WithDeps(deps.Extract(gomod)).
			WithResources(resources.Collect()).
			WithKernel(kernel.Collect("vm.swappiness")).
			Print()
		os.Exit(0)
	case options.GetB(OPT_HELP):
		genUsage().Print()
		os.Exit(0)
	}

	if useRawOutput {
		printRawTop()
	} else {
		printPrettyTop()
	}
}

// preConfigureUI preconfigures UI based on information about user terminal
func preConfigureUI() {
	if !tty.IsTTY() {
		fmtc.DisableColors = true
		useRawOutput = true
	}
}

// configureUI configures user interface
func configureUI() {
	if options.GetB(OPT_NO_COLOR) {
		fmtc.DisableColors = true
	}

	if !useRawOutput {
		fmtutil.SeparatorFullscreen = true
		winWidth = tty.GetWidth()
	}

	if winWidth > 110 {
		cmdEllipsis = winWidth - 40
	}
}

// printPrettyTop prints info with separators and headers
func printPrettyTop() error {
	procInfo, memUsage, err := collectInfo()

	if err != nil {
		return err
	}

	if len(procInfo) == 0 && memUsage.SwapUsed == 0 {
		fmtc.Println("{g}Can't find any process with swap usage{!}")
		return nil
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

	return nil
}

// printPrettyProcessList prints info about swap usage by processes
func printPrettyProcessList(procInfo ProcessInfoSlice) {
	fmtutil.Separator(true)

	fmtc.Printfn(
		" {*}%5s{!} {s}|{!} {*}%16s{!} {s}|{!} {*}%8s{!} {s}|{!} {*}%-s{!}",
		"PID", "USERNAME", "SWAP", "COMMAND",
	)

	fmtutil.Separator(true)

	for _, pi := range procInfo {
		fmtc.Printfn(
			" %5d {s}|{!} %16s {s}|{!} %8s {s}|{!} %-s",
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
		procUsedPerc = mathutil.Perc(procUsed, memUsage.SwapTotal)
		procUsedPerc = mathutil.Between(procUsedPerc, 0.0001, 100.0)
	}

	overallUsed := memUsage.SwapUsed

	// Procfs cannot show values less than 1kb, so we have use calculated processes usage
	if procUsed > memUsage.SwapUsed {
		overallUsed = procUsed
	}

	overallUsedPerc := mathutil.Perc(overallUsed, memUsage.SwapTotal)
	overallUsedPerc = mathutil.Between(overallUsedPerc, 0.0001, 100.0)

	if len(procInfo) == 0 || math.IsNaN(procUsedPerc) {
		fmtc.Println("  {*}Processes:{!} n/a")
	} else {
		fmtc.Printfn(
			"  {*}Processes:{!} %s {s-}(%s){!}",
			fmtutil.PrettySize(procUsed),
			fmtutil.PrettyPerc(procUsedPerc),
		)
	}

	if math.IsNaN(overallUsedPerc) {
		fmtc.Println("  {*}Overall:{!} n/a")
	} else {
		fmtc.Printfn(
			"  {*}Overall:{!}   %s {s-}(%s){!}",
			fmtutil.PrettySize(overallUsed),
			fmtutil.PrettyPerc(overallUsedPerc),
		)
	}

	fmtc.Printfn("  {*}Total:{!}     %s", fmtutil.PrettySize(memUsage.SwapTotal))
}

// printRawTop just prints raw info
func printRawTop() error {
	procInfo, _, err := collectInfo()

	if err != nil {
		return err
	}

	if len(procInfo) == 0 {
		return nil
	}

	for _, pi := range procInfo {
		fmt.Printf("%d %s %d %s\n", pi.PID, pi.User, pi.VmSwap, pi.Command)
	}

	return nil
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
		return nil, fmt.Errorf("Can't collect info about processes")
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

	sort.Slice(result, func(i, j int) bool {
		return result[i].VmSwap > result[j].VmSwap
	})

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

// ////////////////////////////////////////////////////////////////////////////////// //

// printCompletion prints completion for given shell
func printCompletion() int {
	switch options.GetS(OPT_COMPLETION) {
	case "bash":
		fmt.Print(bash.Generate(genUsage(), APP))
	case "fish":
		fmt.Print(fish.Generate(genUsage(), APP))
	case "zsh":
		fmt.Print(zsh.Generate(genUsage(), optMap, APP))
	default:
		return 1
	}

	return 0
}

// printMan prints man page
func printMan() {
	fmt.Println(man.Generate(genUsage(), genAbout("")))
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
	info.AddExample(
		"-f redis-server",
		`Show current swap consumption by processes with "redis-server" in command`,
	)
	info.AddExample("| wc -l", "Count number of processes which use swap")

	return info
}

// genAbout generates info about version
func genAbout(gitRev string) *usage.About {
	about := &usage.About{
		App:     APP,
		Version: VER,
		Desc:    DESC,
		Year:    2006,
		Owner:   "ESSENTIAL KAOS",
		License: "Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>",

		DescSeparator: "{s}—{!}",
	}

	if gitRev != "" {
		about.Build = "git:" + gitRev
		about.UpdateChecker = usage.UpdateChecker{"essentialkaos/swptop", update.GitHubChecker}
	}

	return about
}
