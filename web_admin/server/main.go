package main

import (
    "github.com/spf13/cobra"
    "fmt"
    "os"
    "server"
)

var (
    bindAddr string
    bindPort int64
)

var rootCmd = &cobra.Command {
    Use: "./web --bind_addr addr --bind_port port",
    Short: "Statistic Web show Stock data automatic",
    Run: func (cmd *cobra.Command, args []string) {
        server.RunServer(bindAddr, bindPort)
    },
}

func init() {
    rootCmd.PersistentFlags().StringVarP(&bindAddr, "bind_addr", "", "0.0.0.0", "server bind interface")
    rootCmd.PersistentFlags().Int64VarP(&bindPort, "bind_port", "", 8080, "server bind port")
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
