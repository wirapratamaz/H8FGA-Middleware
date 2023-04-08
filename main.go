package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//sebuah router untuk menangani request http
	mux := http.NewServeMux()

	//menyimpan handler http greet
	endpoint := http.HandlerFunc(greet)

	//menangani request -> endpoint parameter yg diproses oleh middleware 1 & 2
	mux.Handle("/", middleware1(middleware2(endpoint)))

	fmt.Println("Listening to port 8080")

	//mux -> mengeksekusi middleware secara berurutan
	err := http.ListenAndServe(":3000", mux)

	log.Fatal(err)
}

// handler endpoint "/"
func greet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

// next -> menentukan handler berikutnya
func middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware pertama")
		next.ServeHTTP(w, r)
	})
}

func middleware2(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware kedua")
		next.ServeHTTP(w, r)
	})
}
