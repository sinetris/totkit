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
	"log"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var debug bool = false

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "totkit",
	Short: "TechOps ToolKit",
	Long: heredoc.Doc(`
		TechOps ToolKit
		A CLI to simplify TechOps tasks.
	`),
	Version: "0.0.1",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Persistent flags global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "$HOME/.totkit.yaml", "config file")
	RootCmd.PersistentFlags().BoolVar(&color.NoColor, "no-color", false, "disable colors in command output [$NO_COLOR=1]")
	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Print Debug info")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if debug {
		log.Println("initConfig")
	}

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".totkit" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".totkit")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			if debug {
				log.Printf("Config file not found:\n%v\n", err)
			}
		} else {
			// Config file was found but another error was produced
			if debug {
				log.Printf("Config error:\n%v\n", err)
			}
		}
	} else {
		if debug {
			log.Println("Config is OK")
		}
	}
	if debug {
		log.Printf("Config:\n%v\n", viper.AllSettings())
	}
}
