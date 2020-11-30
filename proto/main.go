package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var listenAddress string
var tlsCommonName string
var flagPath string

var rootCmd = &cobra.Command{
	Use:   "xmas-proto",
	Short: "xmas-proto server.",
	Run:   startServer,
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&tlsCommonName, "tlsCommonName", "c", "xmas.rip", "TLS common name")
	rootCmd.PersistentFlags().StringVarP(&listenAddress, "address", "a", "0.0.0.0:8443", "listen address")
	rootCmd.PersistentFlags().StringVarP(&flagPath, "flag", "f", "flag.txt", "path to flag file")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
