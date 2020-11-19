package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var port int
var address string
var httpAnonymousPort int
var httpAuthorizedPort int

var rootCmd = &cobra.Command{
	Use:   "xmas-socks",
	Short: "xmas-socks is a VERY SECURE SOCKS server written in Go.",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	go serveHTTP("./anonymous", httpAnonymousPort)
	go serveHTTP("./permitted", httpAuthorizedPort)
	startSOCKS()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&address, "address", "a", "0.0.0.0", "socks listen address")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 9000, "socks port")
	rootCmd.PersistentFlags().IntVarP(&httpAnonymousPort, "http1", "y", 8081, "http port for anonymous server")
	rootCmd.PersistentFlags().IntVarP(&httpAuthorizedPort, "http2", "z", 8082, "http port for authorized server")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
