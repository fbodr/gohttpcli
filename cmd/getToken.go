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
		access_token_url := viper.GetString("access_token_url")
		client_id := viper.GetString("client_id")
		client_secret := viper.GetString("client_secret")
		audience := viper.GetString("audience")
		grant_type := viper.GetString("grant_type")
		printAsJson := false

		if isVerbose {
			fmt.Printf("access_token_url = %s\n", access_token_url)
			fmt.Printf("client_id = %s\n", client_id)
			fmt.Printf("client_secret = %s\n", client_secret)
			fmt.Printf("audience = %s\n", audience)
			fmt.Printf("grant_type = %s\n", grant_type)
			printAsJson = true
		}

		accessToken := lib.GetToken(access_token_url, client_id, client_secret, audience, grant_type, printAsJson)
		fmt.Println(accessToken)

	},
}

func init() {
	rootCmd.AddCommand(getTokenCmd)
}
