package main

import (
	"encoding/json"
	"fmt"
	"bytes"
	"net/http"
)

func elasticDocument(document interface{}, elasticHost string, docIndex string, docType string, docId string) {
	data := &document

	payloadBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
	}

	body := bytes.NewReader(payloadBytes)
	loadUrl := fmt.Sprintf("%v/%v/%v/%v", elasticHost, docIndex, docType, docId)
	req, err := http.NewRequest("POST", loadUrl, body)
	if err != nil {
		fmt.Println(err)
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
}
