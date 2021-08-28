package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/youpps/response"
)

func main() {
	http.HandleFunc("/api/user", func(rw http.ResponseWriter, r *http.Request) {
		response.Status(rw, 200).JSON(response.J{
			"id":       7,
			"name":     "youpps",
			"password": "1234567890",
		})
	})

	http.HandleFunc("/api/username", func(rw http.ResponseWriter, r *http.Request) {
		response.Status(rw, 200).String("youpps")
	})
	http.HandleFunc("/api/bytes/username", func(rw http.ResponseWriter, r *http.Request) {
		response.Status(rw, 200).Bytes([]byte("youpps"))
	})
	http.HandleFunc("/api/user/set-in-cookie", func(rw http.ResponseWriter, r *http.Request) {
		bytes, err := json.Marshal(response.J{
			"id":       7,
			"name":     "youpps",
			"password": "1234567890",
		})

		if err != nil {
			log.Fatal(err)
		}

		response.Cookie(rw, &http.Cookie{
			Name:  "user",
			Value: string(bytes),
		})
	})
	http.ListenAndServe(":5000", nil)
}
