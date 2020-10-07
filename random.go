package secret

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strings"
)

func Random16() string {
	const length = 16
	const charSet = "0123456789abcdefghijk-mnopqrstvwxyz" // not have lower case L to minimise confusion
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "0000000000000000"
	}

	var output strings.Builder
	output.Grow(length)
	for i := 0; i < length; i++ {
		randomChar := charSet[int(math.Abs(float64(int(b[i]))))%35]
		output.WriteString(string(randomChar))
	}
	return output.String()

}

func KeyOutputPair(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	type submitType struct {
		Text string `json:"text"`
	}

	decoder := json.NewDecoder(req.Body)

	var t submitType
	err := decoder.Decode(&t)
	if err != nil {
		fmt.Fprint(w, `{"key":"Problem encountered. See log."}`)
		return
	}
	key := Random16()
	output := Encrypt(key, t.Text)

	fmt.Fprint(w, `{"key":"`+key+`","output":"`+output+`"}`)

}
