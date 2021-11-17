package main

import (
	"fmt"
	"log"
	"net/http"
)

func form(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":		
		 http.ServeFile(w, r, "index.html");
	case "POST":
		firstName := r.FormValue("fname");
		lastName  := r.FormValue("lname");
		fmt.Printf("firstName = %s\n", firstName);
		fmt.Printf("lastName = %s\n", lastName);
	}
}

func main() {
	http.HandleFunc("/", form);

	log.Fatal(http.ListenAndServe(":9000", nil));
}
