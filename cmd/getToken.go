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
)

// getTokenCmd represents the getToken command
var getTokenCmd = &cobra.Command{
	Use:   "getToken",
	Short: "get an authentication token from oauth provider",
	Long:  `get an authentication token from oauth provider`,
	Run: func(cmd *cobra.Command, args []string) {
		access_token_url, _ := cmd.Flags().GetString("access_token_url")
		client_id, _ := cmd.Flags().GetString("client_id")
		client_secret, _ := cmd.Flags().GetString("client_secret")
		audience, _ := cmd.Flags().GetString("audience")
		grant_type, _ := cmd.Flags().GetString("grant_type")
		verbose, _ := cmd.Flags().GetBool("verbose")
		if verbose {
			fmt.Printf("access_token_url = %s\n", access_token_url)
			fmt.Printf("client_id = %s\n", client_id)
			fmt.Printf("client_secret = %s\n", client_secret)
			fmt.Printf("audience = %s\n", audience)
			fmt.Printf("grant_type = %s\n", grant_type)
			fmt.Println(oauth.GetToken(access_token_url, client_id, client_secret, audience, grant_type, true))
		} else {
			fmt.Println(oauth.GetToken(access_token_url, client_id, client_secret, audience, grant_type, false))
		}
	},
}

func init() {
	rootCmd.AddCommand(getTokenCmd)

	getTokenCmd.Flags().String("access_token_url", "", "access_token_url")
	getTokenCmd.Flags().String("client_id", "", "client_id")
	getTokenCmd.Flags().String("client_secret", "", "client_secret")
	getTokenCmd.Flags().String("audience", "", "audience")
	getTokenCmd.Flags().String("grant_type", "client_credentials", "grant_type")
	getTokenCmd.Flags().BoolP("verbose", "v", false, "verbose output")

	getTokenCmd.MarkFlagRequired("access_token_url")
	getTokenCmd.MarkFlagRequired("client_id")
	getTokenCmd.MarkFlagRequired("client_secret")
	getTokenCmd.MarkFlagRequired("audience")
}
