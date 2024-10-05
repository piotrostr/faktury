/*
Copyright © 2024 essa.dev
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/piotrostr/essadev/faktury/pkg/config"
	"github.com/piotrostr/essadev/faktury/pkg/invoice"
	"github.com/piotrostr/essadev/faktury/pkg/pdf"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var projectFile string

var rootCmd = &cobra.Command{
	Use:   "faktury",
	Short: "Simple CLI for generating invoices for working on open-source projects",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.LoadConfig(cfgFile)
		if projectFile == "" {
			fmt.Println("Please provide a project file")
			os.Exit(1)
		}
		project := invoice.LoadProject(projectFile)
		pdfGen := pdf.NewGenerator(cfg)
		pdfGen.GenerateInvoice(project)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.json)")
	rootCmd.PersistentFlags().StringVar(&projectFile, "project", "", "project file (JSON)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
