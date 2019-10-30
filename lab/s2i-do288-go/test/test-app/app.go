package main

import 	"net/http"
import "flag"
import	"fmt"


var lang = flag.String("lang", "en", "run app with language support - dafault is english")

func main() {
	var port = "8080"
	http.HandleFunc("/", rootHandler)
	fmt.Printf("Starting server on port %v...\n", port)
	http.ListenAndServe(":"+port, nil)
}

func rootHandler(response http.ResponseWriter, request *http.Request) {
	
	flag.Parse()
	
	switch *lang {
		case "en":
			fmt.Fprintf(response, "Hello %s!. Welcome!\n", request.URL.Path[1:])
		case "es":
			fmt.Fprintf(response, "Hola %s!. Bienvenido!\n", request.URL.Path[1:])
		default:
			fmt.Fprintf(response, "Error!. Unknown lang option -> %s\n", *lang)
	}
}
