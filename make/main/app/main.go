package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// PhraseHash : Phrase - comming string, Hash - hash of Phrase
type PhraseHash struct {
	Phrase string
	Hash   int
}

// Converting hexdigest string to int64, then int
func hex2int(hexStr string) int {
	val, _ := strconv.ParseInt(hexStr, 16, 64)
	return int(val)
}

// Checking that hex-string between min and max int64 value
func checkHashSize(hexStr string) int {
	maxint64 := 9223372036854775807
	minint64 := -9223372036854775807
	int64val := hex2int(hexStr)

	for len(hexStr) > 0 && int64val >= maxint64 || int64val <= minint64 {
		hexStr = hexStr[:len(hexStr)-1]
		int64val = hex2int(hexStr)
	}
	return int64val
}

// Starts here.
func main() {
	http.HandleFunc("/get-phrase-hash", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

// Handling /get-phrase-hash request with Phrase and returns json of PhraseHash
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var p PhraseHash

		json.NewDecoder(r.Body).Decode(&p)

		sha := sha256.New()
		sha.Write([]byte(p.Phrase))

		p.Hash = checkHashSize(fmt.Sprintf("%x", sha.Sum(nil)))
		jsn, err := json.Marshal(p)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsn)

	} else {
		fmt.Fprintf(w, "%s method not allowed", r.Method)
	}
}
