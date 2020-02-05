package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetToken(
	accessTokenUrl string,
	clientId string,
	clientSecret string,
	audience string,
	grantType string,
	fullJson bool) string {

	var payloadMap map[string]string
	payloadMap = make(map[string]string)
	payloadMap["client_id"] = clientId
	payloadMap["client_secret"] = clientSecret
	payloadMap["audience"] = audience
	payloadMap["grant_type"] = grantType
	payload, _ := json.Marshal(payloadMap)

	auth0Req, err := http.NewRequest("POST", accessTokenUrl, strings.NewReader(string(payload)))
	if err != nil {
		panic(err)
	}
	auth0Req.Header.Add("content-type", "application/json")
	auth0Res, err := http.DefaultClient.Do(auth0Req)
	if err != nil {
		panic(err)
	}
	defer auth0Res.Body.Close()
	auth0Body, _ := ioutil.ReadAll(auth0Res.Body)
	var auth0Fields map[string]interface{}
	err = json.Unmarshal(auth0Body, &auth0Fields)
	if err != nil {
		fmt.Println(auth0Body)
		panic(err)
	}

	accessToken := fmt.Sprintf("%s %s", auth0Fields["token_type"], auth0Fields["access_token"])
	ContextSet("access_token", accessToken)

	if fullJson {
		return string(auth0Body)
	} else {
		return accessToken
	}

}
