// +build linux

package main

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                     Copyright (c) 2009-2017 ESSENTIAL KAOS                         //
//        Essential Kaos Open Source License <https://essentialkaos.com/ekol>         //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"pkg.re/essentialkaos/ek.v9/env"
	"pkg.re/essentialkaos/ek.v9/fmtc"
	"pkg.re/essentialkaos/ek.v9/fmtutil"
	"pkg.re/essentialkaos/ek.v9/fsutil"
	"pkg.re/essentialkaos/ek.v9/options"
	"pkg.re/essentialkaos/ek.v9/strutil"
	"pkg.re/essentialkaos/ek.v9/system/process"
	"pkg.re/essentialkaos/ek.v9/terminal/window"
	"pkg.re/essentialkaos/ek.v9/usage"
	"pkg.re/essentialkaos/ek.v9/usage/update"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	APP  = "swptop"
	VER  = "0.2.0"
	DESC = "Utility for viewing swap consumption of processes"
)

const (
	OPT_NO_COLOR = "nc:no-color"
	OPT_HELP     = "h:help"
	OPT_VER      = "v:version"
)

// ////////////////////////////////////////////////////////////////////////////////// //

type ProcessInfo struct {
	PID     int
	VmSwap  uint64
	User    string
	Command string
}

type ProcessInfoSlice []ProcessInfo

// ////////////////////////////////////////////////////////////////////////////////// //

func (s ProcessInfoSlice) Len() int           { return len(s) }
func (s ProcessInfoSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ProcessInfoSlice) Less(i, j int) bool { return s[i].VmSwap < s[j].VmSwap }

// ////////////////////////////////////////////////////////////////////////////////// //

// optMap is map with options
var optMap = options.Map{
	OPT_NO_COLOR: {Type: options.BOOL},
	OPT_HELP:     {Type: options.BOOL, Alias: "u:usage"},
	OPT_VER:      {Type: options.BOOL, Alias: "ver"},
}

// Raw output flag
var useRawOutput bool

// Window width
var winWidth int

// ////////////////////////////////////////////////////////////////////////////////// //

func main() {
	_, errs := options.Parse(optMap)

	if len(errs) != 0 {
		for _, err := range errs {
			printError(err.Error())
		}

		os.Exit(1)
	}

	configureUI()

	if options.GetB(OPT_VER) {
		showAbout()
		return
	}

	if options.GetB(OPT_HELP) {
		showUsage()
		return
	}

	if useRawOutput {
		printRawTop()
	} else {
		printPrettyTop()
	}
}

// configureUI configure user interface
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
}

// printPrettyTop print info with separators and headers
func printPrettyTop() {
	info, err := collectInfo()

	if err != nil {
		printErrorAndExit(err.Error())
	}

	if len(info) == 0 {
		fmtc.Println("{g}System doesn't have any process with swap usage{!}")
		return
	}

	cmdEllipsis := 64

	if winWidth > 110 {
		cmdEllipsis = winWidth - 40
	}

	fmtutil.Separator(true)
	fmtc.Printf(
		" {*}%5s{!} {s}|{!} {*}%16s{!} {s}|{!} {*}%8s{!} {s}|{!} {*}%-s{!}\n",
		"PID", "USERNAME", "SWAP", "COMMAND",
	)
	fmtutil.Separator(true)

	for _, pi := range info {
		fmtc.Printf(
			" %5d {s}|{!} %16s {s}|{!} %8s {s}|{!} %-s\n",
			pi.PID, pi.User, fmtutil.PrettySize(pi.VmSwap),
			strutil.Ellipsis(pi.Command, cmdEllipsis),
		)
	}

	fmtutil.Separator(true)
}

// printRawTop just print raw info
func printRawTop() {
	info, err := collectInfo()

	if err != nil {
		printErrorAndExit(err.Error())
	}

	if len(info) == 0 {
		return
	}

	for _, pi := range info {
		fmt.Printf("%d %s %d %s\n", pi.PID, pi.User, pi.VmSwap, pi.Command)
	}
}

// collectInfo collect info about processes and sort result slice
func collectInfo() (ProcessInfoSlice, error) {
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

		if err != nil {
			continue
		}

		if memInfo.VmSwap == 0 {
			continue
		}

		result = append(
			result,
			ProcessInfo{
				PID:     pi.PID,
				User:    pi.User,
				Command: pi.Command,
				VmSwap:  memInfo.VmSwap,
			},
		)
	}

	sort.Sort(sort.Reverse(result))

	return result, nil
}

// printError prints error message to console
func printError(f string, a ...interface{}) {
	fmtc.Fprintf(os.Stderr, "{r}"+f+"{!}\n", a...)
}

// printError prints warning message to console
func printWarn(f string, a ...interface{}) {
	fmtc.Fprintf(os.Stderr, "{y}"+f+"{!}\n", a...)
}

// printErrorAndExit print error mesage and exit with exit code 1
func printErrorAndExit(f string, a ...interface{}) {
	printError(f, a...)
	os.Exit(1)
}

// ////////////////////////////////////////////////////////////////////////////////// //

func showUsage() {
	info := usage.NewInfo("")

	info.AddOption(OPT_NO_COLOR, "Disable colors in output")
	info.AddOption(OPT_HELP, "Show this help message")
	info.AddOption(OPT_VER, "Show version")

	info.Render()
}

func showAbout() {
	about := &usage.About{
		App:           APP,
		Version:       VER,
		Desc:          DESC,
		Year:          2006,
		Owner:         "ESSENTIAL KAOS",
		License:       "Essential Kaos Open Source License <https://essentialkaos.com/ekol>",
		UpdateChecker: usage.UpdateChecker{"essentialkaos/swptop", update.GitHubChecker},
	}

	about.Render()
}
