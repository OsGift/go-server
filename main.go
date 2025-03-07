package main

import (
	"fmt"
	"log"
	"net/http"
)

// func formHandler(w http.ResponseWriter, r *http.Request) {
// 	if err := r.ParseForm(); err != nil {
// 		fmt.Fprintf(w, "ParseForm() err: %v", err)
// 		return
// 	}
// 	fmt.Fprintf(w, "POST request successful")
// 	name := r.FormValue("name")
// 	address := r.FormValue("address")
// 	fmt.Fprintf(w, "Name = %s\n", name)
// 	fmt.Fprintf(w, "Address = %s\n", address)
// }

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "ParseForm() error", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	address := r.FormValue("address")

	// Define a simple success page
	successPage := fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Form Submitted</title>
		<style>
			body { font-family: Arial, sans-serif; text-align: center; padding: 50px; background: #f4f4f4; }
			.container { background: white; padding: 20px; border-radius: 10px; display: inline-block; box-shadow: 0 0 10px rgba(0, 0, 0, 0.1); }
			h2 { color: #333; }
			p { color: #555; }
			a { display: inline-block; margin-top: 15px; padding: 10px 20px; background: #667eea; color: white; text-decoration: none; border-radius: 5px; }
			a:hover { background: #5643a8; }
		</style>
	</head>
	<body>
		<div class="container">
			<h2>Form Submitted Successfully!</h2>
			<p><strong>Name:</strong> %s</p>
			<p><strong>Address:</strong> %s</p>
			<a href="/">Go Back</a>
		</div>
	</body>
	</html>`, name, address)

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, successPage)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello!")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
