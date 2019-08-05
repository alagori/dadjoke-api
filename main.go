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

func getjoke() []byte {
	file, _ := ioutil.ReadFile("jokelist.json")
	data := jokes{}
	_ = json.Unmarshal([]byte(file), &data)
	// for i := 0; i < len(data); i++ {
	// 	fmt.Println("q: ", data[i].Q)
	// 	fmt.Println("a: ", data[i].A)
	// }

	answer := data[randnum(len(data))]
	result, _ := json.Marshal(answer)
	// print(result)
	return result
}

// func webserver() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		getjoke()
// 	})

// 	http.ListenAndServe(":80", nil)

// }

func main() {
	// webserver()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(getjoke())

	})

	http.ListenAndServe(":8080", nil)

}
