/*
Copyright © 2020 Fabio Rafael da Rosa <fdr@fabiodarosa.org>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"github.com/fbodr/gohttpcli/lib"
	"github.com/spf13/cobra"
	"os"

	"github.com/spf13/viper"
)

var cfgFile string
var isVerbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gohttpcli",
	Short: "simple command line http client",
	Long:  `simple command line http client`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	isVerbose, _ = rootCmd.Flags().GetBool("verbose")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gohttpcli/config.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")

	rootCmd.PersistentFlags().String("access_token_url", "", "access_token_url")
	rootCmd.PersistentFlags().String("client_id", "", "client_id")
	rootCmd.PersistentFlags().String("client_secret", "", "client_secret")
	rootCmd.PersistentFlags().String("audience", "", "audience")
	rootCmd.PersistentFlags().String("grant_type", "client_credentials", "grant_type")

	rootCmd.MarkFlagRequired("access_token_url")
	rootCmd.MarkFlagRequired("client_id")
	rootCmd.MarkFlagRequired("client_secret")
	rootCmd.MarkFlagRequired("audience")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	appCfgPath := os.Getenv("HOME") + "/.gohttpcli"
	_, err := os.Stat(appCfgPath)
	if err != nil {
		/* dir does not exists */
		os.Mkdir(os.Getenv("HOME")+"/.gohttpcli", 0777)
	}

	viper.SetConfigType("yaml")
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(appCfgPath)
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	viper.BindPFlags(rootCmd.Flags())
	lib.ContextInit(appCfgPath)
}
