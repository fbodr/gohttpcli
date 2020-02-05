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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// getTokenCmd represents the getToken command
var getTokenCmd = &cobra.Command{
	Use:   "getToken",
	Short: "get an authentication token from oauth provider",
	Long:  `get an authentication token from oauth provider`,
	Run: func(cmd *cobra.Command, args []string) {
		accessTokenUrl := viper.GetString("access_token_url")
		clientId := viper.GetString("client_id")
		clientSecret := viper.GetString("client_secret")
		audience := viper.GetString("audience")
		grantType := viper.GetString("grant_type")
		printAsJson := false

		if isVerbose {
			fmt.Printf("access_token_url = %s\n", accessTokenUrl)
			fmt.Printf("client_id = %s\n", clientId)
			fmt.Printf("client_secret = %s\n", clientSecret)
			fmt.Printf("audience = %s\n", audience)
			fmt.Printf("grant_type = %s\n", grantType)
			printAsJson = true
		}

		accessToken := lib.GetToken(accessTokenUrl, clientId, clientSecret, audience, grantType, printAsJson)
		fmt.Println(accessToken)

	},
}

func init() {
	rootCmd.AddCommand(getTokenCmd)
}
