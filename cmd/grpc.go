// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"cobra/config"
	"cobra/rpcs/gateway"
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// grpcCmd represents the grpc command
var grpcCmd = &cobra.Command{
	Use:   "grpc",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		GrpcRun()
	},
}

func init() {
	rootCmd.AddCommand(grpcCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grpcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

//GrpcRun 服务启动
func GrpcRun() {
	var grpcServer *grpc.Server
	defer func() {
		config.Close()
	}()

	config.MonitorConfig()

	if err := config.LoadManages(); err != nil {
		log.Printf("load manages fail :%v", err)
		return
	}

	log.Println("Grpc server starting...")
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", viper.GetString("rpc.host"), viper.GetInt("rpc.port")))
	if err != nil {
		log.Fatalf("Grpc initialize failed. %s\n", err.Error())
	}

	grpcServer = grpc.NewServer()
	gateway.RegisterPayCenterSrvServer(grpcServer, new(gateway.Service))
	grpcServer.Serve(lis)
}
