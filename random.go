package secret

import (
	"math/rand"
	"strings"
	"time"
)

func Random16() string {
	const length = 16
	const charSet = "0123456789abcdefghijk-mnopqrstvwxyz" // not have lower case L to minimise confusion
	rand.Seed(time.Now().Unix())

	var output strings.Builder

	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	return output.String()

}
