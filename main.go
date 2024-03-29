package main

import (
	"crypto/rand"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"
)

type jokes []joke

type joke struct {
	Q string `json:"q"`
	A string `json:"a"`
}

func randnum(length int) int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(27))
	if err != nil {
		panic(err)
	}
	n := nBig.Int64()
	return n
}

func getJoke(lastnum int64) ([]byte, int64) {
	file, _ := ioutil.ReadFile("jokelist.json")
	data := jokes{}
	_ = json.Unmarshal([]byte(file), &data)
	// for i := 0; i < len(data); i++ {
	// 	fmt.Println("q: ", data[i].Q)
	// 	fmt.Println("a: ", data[i].A)
	// }
	jokenum := randnum(len(data))
	answer := data[jokenum]
	result, _ := json.Marshal(answer)
	// print(result)
	return result, jokenum
}

func getOneJokeRand() {
	lastJoke := int64(1)
	http.HandleFunc("/getonejokerand", func(w http.ResponseWriter, r *http.Request) {
		joke, jokenum := getJoke(lastJoke)
		w.Write(joke)
		lastJoke = jokenum
	})
}

// LegacyEndPoint to keep compatiblity with old clients
func legacyEndPoint() {
	lastJoke := int64(1)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		joke, jokenum := getJoke(lastJoke)
		w.Write(joke)
		lastJoke = jokenum
	})
}

func main() {

	getOneJokeRand()
	http.ListenAndServe(":8080", nil)

}
