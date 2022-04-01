/*
Copyright © 2022 Farye Nwede <farye@aeekay.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type scrapemsg struct {
	URL string `json:"url", yaml:"url"`
}

// scheduleScrapeCmd represents the getBuckets command
var scheduleScrapeCmd = &cobra.Command{
	Use:   "schedule-scrape",
	Short: "Schedule the scrape of a website",
	Long: `Schedule the scrape of a website. The scrape will include
a pdf, an image, and a HTML document of the site depending on the 
features available from the scrape service.`,
	Run: func(cmd *cobra.Command, args []string) {
		var sw string
		sw = strings.TrimSpace(args[0])

		if sw == "" {
			fmt.Println("[error] empty argument")
			return
		}

		fmt.Printf("scheduling scrape of %s\n", sw)
		o := scrapemsg{
			URL: sw,
		}

		msg, err := json.Marshal(o)
		if err != nil {
			fmt.Println("[error] couldn't marshal object")
			return
		}

		// this should really be a web hook but for now, we'll just use mqtt
		// for quick and dirty
		mh := viper.GetString("mqtt.host")
		mp := viper.GetInt("mqtt.port")
		mu := viper.GetString("mqtt.username")
		mpwd := viper.GetString("mqtt.password")
		mt := viper.GetString("mqtt.topics.scrape")
		mc := fmt.Sprintf("cli-%s", uuid.New().String())

		opts := mqtt.NewClientOptions()
		opts.AddBroker(fmt.Sprintf("tls://%s:%d", mh, mp))
		opts.SetClientID(mc) // set a name as you desire
		opts.SetUsername(mu) // these are the credentials that you declare for your cluster
		opts.SetPassword(mpwd)

		client := mqtt.NewClient(opts)
		// throw an error if the connection isn't successfull
		if token := client.Connect(); token.Wait() && token.Error() != nil {
			fmt.Printf("[error]: couldn't connect to client %v\n", token.Error())
			return
		}

		err = publishMqtt(client, mt, string(msg))

		if err != nil {
			fmt.Printf("[error]: couldn't publish message to client %v\n", err)
			return
		}

		fmt.Printf("published message\n")

		client.Disconnect(250)
	},
}

func init() {
	mlCmd.AddCommand(scheduleScrapeCmd)
}

func publishMqtt(client mqtt.Client, topic string, msg string) error {
	token := client.Publish(topic, 0, false, msg)
	token.Wait()
	// Check for errors during publishing (More on error reporting https://pkg.go.dev/github.com/eclipse/paho.mqtt.golang#readme-error-handling)
	if token.Error() != nil {
		return fmt.Errorf("failed to publish to topic %s", topic)
	}

	return nil
}
