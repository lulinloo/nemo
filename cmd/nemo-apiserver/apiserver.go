// apiserver is the api server for nemo-apiserver service.
// it is responsible for serving the platform RESTful resource management.
package main

import (
	"fmt"

	"github.com/lulinloo/nemo/pkg/config"
	"github.com/lulinloo/nemo/pkg/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var cfg *config.Config
var name string
var name1 string

// var age string
var logger *zap.Logger

var rootCmd = &cobra.Command{
	Use:   "nemo-apiserver",
	Short: "Command line tool for Database",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// name, err := cmd.Flags().GetString("name")
		// if err != nil {
		// 	fmt.Printf("name error", zap.Error(err))
		// }
		// fmt.Printf(name)
		cfgFilepath, err := cmd.Flags().GetString("config")
		if err != nil {
			logger.Fatal("config flag error", zap.Error(err))
		}

		cfg, err = config.LoadConfig(cfgFilepath)
		if err != nil {
			logger.Fatal("LoadConfig failed", zap.Error((err)))
		}

		fmt.Println(cfg.Dbtypes)

		logger, _ = log.NewLogger()
		// logger.Info("qq")
	},
	Run: func(cmd *cobra.Command, args []string) {
		// _ = logger.Sync()
		// fmt.Println("well")
		// logger, _ := log.NewLogger()
		logger.Info(name)

		// name, err := cmd.Flags().GetString("name")
		// if err != nil {
		// 	fmt.Printf("name error", zap.Error(err))
		// }
		// fmt.Printf(name)
	},
}

func main() {
	// MarkFlagRequired:必须填写的commands;PersistentFlags既可以设置给该 Command，又可以设置给该 Command 的子 Command;
	// Flags只能设置给指定的 Command
	rootCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "name to operate.")
	// rootCmd.MarkFlagRequired("name")
	viper.BindPFlag("name1", rootCmd.PersistentFlags().Lookup("name"))
	// fmt.Printf(name1)

	rootCmd.PersistentFlags().StringP("config", "f", "", "Config file path.")

	rootCmd.AddCommand(configCmd)
	configCmd.AddCommand(configCatCmd)

	rootCmd.AddCommand(runnowCmd)

	// viper.BindPFlag("age", rootCmd.PersistentFlags().Lookup("age"))

	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("command failed")
	}

	

}
