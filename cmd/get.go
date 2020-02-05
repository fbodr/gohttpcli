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
	"fmt"
	"github.com/fbodr/gohttpcli/lib"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "make a http get request to a url",
	Long: `make a http get request to a url

    If informed, a auth token can be used from the environment
`,
	Run: func(cmd *cobra.Command, args []string) {
		var token string

		accessTokenUrl, _ := rootCmd.Flags().GetString("access_token_url")
		clientId, _ := rootCmd.Flags().GetString("client_id")
		clientSecret, _ := rootCmd.Flags().GetString("client_secret")
		audience, _ := rootCmd.Flags().GetString("audience")
		grantType, _ := rootCmd.Flags().GetString("grant_type")

		if lib.ContextHasKey("access_token") {
			token = lib.ContextGetValue("access_token")
		} else {
			token = lib.GetToken(accessTokenUrl, clientId, clientSecret, audience, grantType, false)
		}

		makeRequest(token, args[0])
	},
}

func makeRequest(token string, url string) {
	client := &http.Client{}

	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("Authorization", token)

	fmt.Println("")

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	responseBody, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(responseBody))
}

func init() {
	rootCmd.AddCommand(getCmd)
}
