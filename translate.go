package translate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const translateApi = "https://translate.googleapis.com/translate_a/single?client=gtx&sl=%v&tl=%v&dt=t&q=%v"

type translation struct {
	translated string
	source     string
	code       float64
}

func toTranslation(element []interface{}) *translation {
	return &translation{
		translated: element[0].(string),
		source:     element[1].(string),
		code:       element[4].(float64)}
}

//Translate translates a source language 'sl' to target language 'tl' given a query 'q'
//If httpClient is nil http.DefaultClient will be used for requests
func Translate(sl, tl, q string, httpClient *http.Client) (string, error) {
	client := http.DefaultClient
	if httpClient != nil {
		client = httpClient
	}

	q = strings.TrimSpace(q)
	query := strings.Replace(url.QueryEscape(q), ".", "%2E", -1)

	u := fmt.Sprintf(translateApi, sl, tl, query)
	req, _ := http.NewRequest("POST", u, nil)
	req.Header.Set("Content-Length", strconv.Itoa(len(query)))
	r, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("could not translate text: %v", err)
	}

	if r.StatusCode != 200 {
		return "", fmt.Errorf("could read request response code: %v", r.StatusCode)
	}

	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", fmt.Errorf("could not read response file: %v", err)
	}
	var element []interface{}
	if err := json.Unmarshal(b, &element); err != nil {
		return "", fmt.Errorf("could not unmarshal str %v: %v", string(b), err)
	}
	subEl := element[0].([]interface{})
	var translated string
	for _, t := range subEl {
		subSubEl := t.([]interface{})
		translated += toTranslation(subSubEl).translated
	}

	return translated, nil
}
