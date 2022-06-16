package main

import (
	"fmt"
	"os/exec"
	"net/http"
	"log"
)
func toggle(status string) {

	cmd := exec.Command("gpio", "mode", "7", status)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
	}
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	http.HandleFunc("/on", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
		toggle("up")
	})
	http.HandleFunc("/off", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
		toggle("down")
	})

	log.Fatal(http.ListenAndServe(":5000", nil))
}










