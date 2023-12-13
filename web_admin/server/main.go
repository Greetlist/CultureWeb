package main

import (
    "github.com/spf13/cobra"
    "fmt"
    "os"
    "github.com/Greetlist/CultureWeb/web_admin/server/server"
    _ "github.com/shopspring/decimal"
)

var (
    config_file string
    run_tls bool
)

var rootCmd = &cobra.Command {
    Use: "./web --config_file config_path --run_tls",
    Short: "Culture web admin",
    Run: func (cmd *cobra.Command, args []string) {
        if run_tls {
            server.RunTLSServer(config_file)
        } else {
            server.RunServer(config_file)
        }
    },
}

func init() {
    rootCmd.PersistentFlags().StringVarP(&config_file, "config_file", "", "config.yaml", "config file path")
    rootCmd.PersistentFlags().BoolVarP(&run_tls, "run_tls", "", false, "run tls or not")
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Printf("root cmd execute error.\n")
        os.Exit(1)
    }
}

// @title Stock-Show Web API
// @version 1.0
// @description server API for stock web
// @termsOfService http://github.com/greetlist/

// @contact.name Greetlist
// @contact.url http://github.com/greetlist
// @contact.email lrt12250914@outlook.com

// @BasePath /

func main() {
    Execute()
}
