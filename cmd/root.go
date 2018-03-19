package cmd

import (
  "github.com/spf13/cobra"
  "fmt"
  "net/http"
  "io/ioutil"
  "net/url"
  "strconv"
  "encoding/json"
  "os"
)

const translateApi = "https://translate.googleapis.com/translate_a/single?client=gtx&sl=%v&tl=%v&dt=t&q=%v"

var (
  sl, tl string
)

func init() {
  rootCmd.PersistentFlags().StringVar(&sl, "sl", "en", "source language")
  rootCmd.PersistentFlags().StringVar(&tl, "tl", "it", "target language")
}

var rootCmd = &cobra.Command{
  Use:   "",
  Short: "a simple cli for translation from google.",
  RunE:   translate}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Print(err)
    os.Exit(1)
  }
}

func translate(_ *cobra.Command, args []string) error {

  if len(args) == 0 {
    return fmt.Errorf("please provide at least 1 arg to translate")
  }
  var escapedQuery string
  var queryLen = 0
  for _, q := range args {
    escaped := url.PathEscape(q + " ")
    escapedQuery += escaped
    queryLen += len(escaped)
  }
  u := fmt.Sprintf(translateApi, sl, tl, escapedQuery)
  req, _ := http.NewRequest("POST", u, nil)
  req.Header.Set("Content-Length", strconv.Itoa(queryLen))
  r, err := http.DefaultClient.Do(req)
  if err != nil {
    return fmt.Errorf("could not translate text: %v", err)
  }

  if r.StatusCode != 200 {
    return fmt.Errorf("could read request response code: %v", r.StatusCode)
  }

  defer r.Body.Close()
  b, err := ioutil.ReadAll(r.Body)
  if err != nil {
    return fmt.Errorf("could not read response file: %v", err)
  }
  var element []interface{}
  if err := json.Unmarshal(b, &element); err != nil {
    return fmt.Errorf("could not unmarshal str %v: %v", string(b), err)
  }
  subEl := element[0].([]interface{})[0]
  subSubEl := subEl.([]interface{})
  fmt.Printf("%v\n", toTranslation(subSubEl).translated)
  return nil
}

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
