/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/ccheeliang/nats-example/nats-websocket/pkg/server"
	"github.com/spf13/cobra"
)

// server2Cmd represents the server2 command
var server2Cmd = &cobra.Command{
	Use:   "server2",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("server2 called")
		sv, err := server.StartServer("ws://localhost:9090")
		if err != nil {
			log.Fatal(err)
		}

		sv.ListenAndServe("9001")
	},
}

func init() {
	rootCmd.AddCommand(server2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// server2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// server2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
