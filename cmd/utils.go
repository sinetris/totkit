// Copyright © 2023 Duilio Ruggiero <duilio@sinetris.info>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	addTemplateFuncs()
	RootCmd.SetUsageTemplate(heredoc.Doc(`
		{{"Usage:" | green}}{{if .Runnable}}
		  {{.UseLine | bold}}{{end}}{{if .HasAvailableSubCommands}}
		  {{.CommandPath | bold}} {{"[command]" | bold}}{{end}}{{if gt (len .Aliases) 0}}

		{{"Aliases:" | yellow}}
		  {{.NameAndAliases}}{{end}}{{if .HasExample}}

		{{"Examples:" | yellow}}
		  {{.Example}}{{end}}{{if .HasAvailableSubCommands}}{{$cmds := .Commands}}{{if eq (len .Groups) 0}}

		{{"Available Commands:" | yellow}}{{range $cmds}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
		  {{rpad .Name .NamePadding | bold}} {{.Short | hiBlue}}{{end}}{{end}}{{else}}{{range $group := .Groups}}
		  {{.Title}}{{range $cmds}}{{if (and (eq .GroupID $group.ID) (or .IsAvailableCommand (eq .Name "help")))}}
		  {{rpad .Name .NamePadding | bold}} {{.Short | hiBlue}}{{end}}{{end}}{{end}}{{if not .AllChildCommandsHaveGroup}}

		Additional Commands:{{range $cmds}}{{if (and (eq .GroupID "") (or .IsAvailableCommand (eq .Name "help")))}}
		  {{rpad .Name .NamePadding | bold}} {{.Short | hiBlue}}{{end}}{{end}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

		{{"Flags:" | yellow}}
		{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

		{{"Global Flags:" | yellow}}
		{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

		{{"Additional help topics:" | yellow}}{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
		  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

		Use "{{.CommandPath | bold}} {{"[command]" | yellow}} {{"--help" | bold}}" for more information about a command.{{end}}
	`))
	RootCmd.SetHelpTemplate(heredoc.Doc(`
		{{with (or .Long .Short)}}{{.| trimTrailingWhitespaces | yellow | bold }}
		{{end}}{{if or .Runnable .HasSubCommands}}
		{{.UsageString}}{{end}}`))
}

func addTemplateFuncs() {
	cobra.AddTemplateFunc("cyan", color.CyanString)
	cobra.AddTemplateFunc("green", color.GreenString)
	cobra.AddTemplateFunc("yellow", color.YellowString)
	cobra.AddTemplateFunc("red", color.RedString)
	cobra.AddTemplateFunc("blue", color.BlueString)
	cobra.AddTemplateFunc("magenta", color.MagentaString)
	cobra.AddTemplateFunc("hiCyan", color.HiCyanString)
	cobra.AddTemplateFunc("hiGreen", color.HiGreenString)
	cobra.AddTemplateFunc("hiYellow", color.HiYellowString)
	cobra.AddTemplateFunc("hiRed", color.HiRedString)
	cobra.AddTemplateFunc("hiBlue", color.HiBlueString)
	cobra.AddTemplateFunc("hiMagenta", color.HiMagentaString)
	cobra.AddTemplateFunc("bold", func(format string) string {
		return color.New(color.Bold).Sprint(format)
	})
	cobra.AddTemplateFunc("grey", func(format string) string {
		return color.New(color.FgWhite).Sprint(format)
	})
}
