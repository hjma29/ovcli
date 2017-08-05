// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"fmt"
	"log"
	"os"

	//"github.com/docker/machine/libmachine/log"
	"github.com/hashicorp/logutils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "ovcli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.AddCommand(showCmd)
	RootCmd.AddCommand(deleteCmd)
	RootCmd.AddCommand(createCmd)

	RootCmd.PersistentFlags().BoolVarP(&Debugmode, "debug", "d", false, "Debug:true,false")

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.ovcli.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".ovcli") // name of config file (without extension)
	viper.AddConfigPath("$HOME")  // adding home directory as first search path
	viper.AutomaticEnv()          // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	//viper.BindPFlag("debugbit", RootCmd.PersistentFlags().Lookup("Debugmode"))

	//fmt.Println(viper.GetBool("debugbit"))

	//fmt.Println(Debugmode)

	if Debugmode {
		//log.SetDebug(true)
		filter := &logutils.LevelFilter{
			Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
			MinLevel: logutils.LogLevel("DEBUG"),
			Writer:   os.Stderr,
		}
		log.SetOutput(filter)

	} else {
		filter := &logutils.LevelFilter{
			Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
			MinLevel: logutils.LogLevel("WARN"),
			Writer:   os.Stderr,
		}
		log.SetOutput(filter)

	}

	// if !Debugmode {
	// 	//fmt.Println("setting output to discard")
	// 	//if viper.GetBool("debugbit") {
	// 	// log.SetOutput(ioutil.Discard)
	// } else {
	// 	//fmt.Println("setting output to stdout")

	// 	// log.SetOutput(os.Stdout)
	// }

}
