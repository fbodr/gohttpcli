package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetToken(
	access_token_url string,
	client_id string,
	client_secret string,
	audience string,
	grant_type string,
	fullJson bool) string {

	var payload_map map[string]string
	payload_map = make(map[string]string)
	payload_map["client_id"] = client_id
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
	err = json.Unmarshal(auth0_body, &auth0_fields)
	if err != nil {
		fmt.Println(auth0_body)
		panic(err)
	}

	access_token := fmt.Sprintf("%s %s", auth0_fields["token_type"], auth0_fields["access_token"])
	ContextSet("access_token", access_token)

	if fullJson {
		return string(auth0_body)
	} else {
		return access_token
	}

}
