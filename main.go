package gva

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func IsErr(e error) {
	if e != nil {
		log.Fatal(e)
	}
}


func ApiLogin(client *http.Client, profile Profile, address string, username string) TokenModel {

	// client := Client(secure)

	password := os.Getenv("VEEAM_API_PASSWORD")

	if password == "" {
		log.Fatal("VEEAM_API_PASSWORD not set in ENV")
	}

	data := url.Values{}
	data.Add("grant_type", "password")
	data.Add("username", username)
	data.Add("password", password)

	urlString := fmt.Sprintf("https://%s%s", address, profile.URL)

	r, err := http.NewRequest("POST", urlString, strings.NewReader(data.Encode()))
	IsErr(err)
	r.Header.Add("accept", profile.Headers.Accept)
	r.Header.Add("x-api-version", profile.Headers.XAPIVersion)
	r.Header.Add("Content-Type", profile.Headers.ContentType)

	res, err := client.Do(r)
	IsErr(err)

	if res.StatusCode == 401 {
		log.Fatalf("Not Authorized: %v", res.StatusCode)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	IsErr(err)

	var token TokenModel

	if err := json.Unmarshal(body, &token); err != nil {
		log.Fatalf("Could not unmarshal token - %v", err)
	}
	
	return token
}

func BuildRequestUrl(address string, endpoint string, profile Profile) string {

	apibit := "/api/"

	if profile.Name == "vb365" {
		apibit = "/"
	}

	cs := fmt.Sprintf("https://%v:%v%v%v/%v", address, profile.Port, apibit, profile.APIVersion, endpoint)
	return cs
}

func AddHeaders(req *http.Request, profile Profile, token TokenModel) {
	req.Header.Add("accept", profile.Headers.Accept)
	req.Header.Add("x-api-version", profile.Headers.XAPIVersion)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer " + token.AccessToken)
}