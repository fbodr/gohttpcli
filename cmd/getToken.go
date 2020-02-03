/*
Copyright © 2020 Fabio Rafael da Rosa <fdr@fabiodarosa.org>

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

	"github.com/spf13/cobra"
	"io/ioutil"
	"net/http"
	"strings"
)

// getTokenCmd represents the getToken command
var getTokenCmd = &cobra.Command{
	Use:   "getToken",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		access_token_url, _ :=  cmd.Flags().GetString("access_token_url")
		client_id, _ := cmd.Flags().GetString("client_id")
		client_secret, _ := cmd.Flags().GetString("client_secret")
		audience, _ := cmd.Flags().GetString("audience")
		grant_type, _ := cmd.Flags().GetString("grant_type")

		/*fmt.Printf("access_token_url = %s\n", access_token_url)
		fmt.Printf("client_id = %s\n", client_id)
		fmt.Printf("client_secret = %s\n", client_secret)
		fmt.Printf("audience = %s\n", audience)
		fmt.Printf("grant_type = %s\n", grant_type)*/

		make_request(access_token_url, client_id, client_secret, audience, grant_type)
	},
}

func init() {
	rootCmd.AddCommand(getTokenCmd)

	getTokenCmd.Flags().String("access_token_url", "", "access_token_url")
	getTokenCmd.Flags().String("client_id", "", "client_id")
	getTokenCmd.Flags().String("client_secret", "", "client_secret")
	getTokenCmd.Flags().String("audience", "", "audience")
	getTokenCmd.Flags().String("grant_type", "client_credentials", "grant_type")

	getTokenCmd.MarkFlagRequired("access_token_url")
	getTokenCmd.MarkFlagRequired("client_id")
	getTokenCmd.MarkFlagRequired("client_secret")
	getTokenCmd.MarkFlagRequired("audience")
}

func make_request(access_token_url string, client_id string, client_secret string, audience string, grant_type string) {

	var payload_map map[string]string
	payload_map = make(map[string]string)
	payload_map["client_id"] =  client_id
	payload_map["client_secret"] = client_secret
	payload_map["audience"] = audience
	payload_map["grant_type"] = grant_type
	payload, _ := json.Marshal(payload_map)

	auth0_req, err := http.NewRequest("POST", access_token_url, strings.NewReader(string(payload)))
	if err != nil {
		panic(err)
	}
	auth0_req.Header.Add("content-type", "application/json")
	auth0_res, err := http.DefaultClient.Do(auth0_req)
	if err != nil {
		panic(err)
	}
	defer auth0_res.Body.Close()
	auth0_body, _ := ioutil.ReadAll(auth0_res.Body)
	var auth0_fields map[string]interface{}
	err = json.Unmarshal([]byte(auth0_body), &auth0_fields)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(auth0_body))

}
