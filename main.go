package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type configJSON struct {
	APIInfo apiInfo `json:"api_info"`
}

type apiInfo struct {
	URL            string                 `json:"url"`
	Method         string                 `json:"method"`
	Headers        map[string]string      `json:"headers"`
	RequestBody    map[string]interface{} `json:"request_body"`
	SampleResponse map[string]interface{} `json:"sample_response"`
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("ERROR: Please provide file path.")
	}

	filePath := os.Args[1]

	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	cfg := configJSON{}

	if err := json.Unmarshal(content, &cfg); err != nil {
		log.Fatal(err)
	}

	cfgRespByte, _ := json.Marshal(cfg.APIInfo.SampleResponse)

	fmt.Println("Config File Content: ")
	fmt.Printf("URL: %s\nMethod: %s\nHeaders: %v\nSample Response: %s\n",
		cfg.APIInfo.URL, cfg.APIInfo.Method, cfg.APIInfo.Headers, string(cfgRespByte))

	var body io.Reader

	if len(cfg.APIInfo.RequestBody) > 0 {
		reqBodyBytes, _ := json.Marshal(cfg.APIInfo.RequestBody)
		body = bytes.NewBuffer(reqBodyBytes)

	}

	req, err := http.NewRequest(cfg.APIInfo.Method, cfg.APIInfo.URL, body)
	if err != nil {
		log.Fatal(err)
	}

	for key, value := range cfg.APIInfo.Headers {
		req.Header.Add(key, value)
	}

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// TO-DO: Making this portion dynamic
	fmt.Println("Status Code:", resp.StatusCode)

	if resp.StatusCode != http.StatusOK {
		bdy, _ := ioutil.ReadAll(resp.Body)
		log.Fatal(string(bdy))
	}

	respSample := map[string]interface{}{}

	if err := json.NewDecoder(resp.Body).Decode(&respSample); err != nil {
		log.Fatal(err)
	}

	respByte, _ := json.Marshal(respSample)

	fmt.Println("Actual Response:", string(respByte))

	if !bytes.Equal(cfgRespByte, respByte) {
		fmt.Println("Test Failed!")
	} else {
		fmt.Println("Test Passed!")
	}

}
