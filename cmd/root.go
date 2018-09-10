// Copyright © 2018 Martin Lebeda <martin.lebeda@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"os"

	"github.com/martinlebeda/taskmaster/termout"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"path/filepath"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "taskmaster",
	Short: "manage your tasks and worktime",
	Long: `Tool for manage your tasts and worktime on computer in simple way.

Base usage is on command line as CLI application.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.taskmaster.yaml)")

	rootCmd.PersistentFlags().String("dbfile", filepath.Join(os.Getenv("HOME"), ".taskmaster.db"), "database file")
	viper.BindPFlag("dbfile", rootCmd.PersistentFlags().Lookup("dbfile"))

	rootCmd.PersistentFlags().String("notifycmd", "notify-send", "command for system notification")
	viper.BindPFlag("notifycmd", rootCmd.PersistentFlags().Lookup("notifycmd"))

	rootCmd.PersistentFlags().String("afterchange", "", "command for run after change (ie. some refresh in external app)")
	viper.BindPFlag("afterchange", rootCmd.PersistentFlags().Lookup("afterchange"))

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Verbose output")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	rootCmd.PersistentFlags().BoolP("color", "C", false, "Force color output output")
	viper.BindPFlag("color", rootCmd.PersistentFlags().Lookup("color"))

	rootCmd.PersistentFlags().BoolP("list-after-change", "L", false, "list all modules after any change")
	viper.BindPFlag("listafterchange", rootCmd.PersistentFlags().Lookup("list-after-change"))
	// TODO Lebeda - list-after-change
}

// TODO Lebeda - subcommand txt - import/export tasks in todo.txt for sync with mobile simpletask application
// TODO Lebeda - subcommand edit/ed ID - run editor for yaml generated and reimport id back (old edit command rename to update/upd)
// TODO Lebeda - subcommand termui/tui - run fzf
// TODO Lebeda - subcommand webserver/web - run server if need on port
// TODO Lebeda - subcommand gui - run server if need on port and open browser gui

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".taskmaster" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".taskmaster")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		termout.Verbose("Using config file:", viper.ConfigFileUsed())
	}
}
