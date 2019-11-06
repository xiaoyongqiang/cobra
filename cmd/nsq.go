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
	"cobra/nsqpubsub/testsub"
	"cobra/pkg/nsq"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// nsqCmd represents the nsq command
var nsqCmd = &cobra.Command{
	Use:   "nsq",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		NsqRun()
	},
}

func init() {
	rootCmd.AddCommand(nsqCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nsqCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nsqCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

//NsqRun 运行Nsq服务
func NsqRun() {

	defer func() {
		config.Close()
	}()

	config.MonitorConfig()

	if err := config.LoadManages(); err != nil {
		log.Printf("load manages fail :%v", err)
		return
	}

	consumer := nsq.NsqConsumer("topic", "channel")
	consumer.AddHandler(&testsub.TestHandler{})

	if err := consumer.ConnectToNSQLookupd(fmt.Sprintf("%s:%d", viper.GetString("nsqlookupd.host"), viper.GetInt("nsqlookupd.port"))); err != nil {
		log.Fatalf("consumer.ConnectToNSQD() fail. Error info: %s\n", err.Error())
	}
	shutdown := make(chan os.Signal, 2)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-consumer.StopChan:
			return
		case <-shutdown:
			consumer.Stop()
		}
	}
}
