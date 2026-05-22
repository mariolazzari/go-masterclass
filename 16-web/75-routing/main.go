package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerHome(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	homePage := `<html>
		<header>
			<title>Home</title>
		</header>
		<body>
			<h1>Home page</h1>
		</body>
		</html>`

	_, err := fmt.Fprint(w, homePage)
	if err != nil {
		w.Write([]byte("error loading home page"))
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlerHome)

	fmt.Println("Server started on port 8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}

}
