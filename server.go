package main

import (
	"fmt"
	"net/http"
)

func rootHander(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello.")
}

func main() {
	http.HandleFunc("/", rootHander)
	err := http.ListenAndServeTLS(":9444", "./public_key", "private_key", nil)
	if err != nil {
		fmt.Println(err)
	}
}
