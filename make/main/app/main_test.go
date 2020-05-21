package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

type Hash struct {
	Hash int `json:"hash"`
}

func TestPhraseHash(t *testing.T) {
	var p Hash
	requestBody, err := json.Marshal(map[string]string{"Phrase": "testing_string"})

	if err != nil {
		t.Error("Cant form a request body")
	}
	resp, err := http.Post("http://localhost/get-phrase-hash", "application/json", bytes.NewBuffer(requestBody))

	if err != nil {
		t.Error("Server error from /get-phrase-hash")
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(respBody), &p)

	if p.Hash != 914709792559998841 {
		t.Error("Values are not equal at from /get-phrase-hash")
	}
}
