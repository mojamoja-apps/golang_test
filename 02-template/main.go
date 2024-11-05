// templateに変数を渡してブラウザアクセスするサンプル
// http://localhost:8080/
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func handler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("layout.html", "page1.html"))
	err := t.Execute(w, map[string]interface{}{
		"Slice": []int{1, 2, 3, 4444, 5, 6, 9999},
		"Map":   map[string]string{"k1": "v1", "k2": "v2"},
	})
	if err != nil {
		fmt.Fprintln(w, err)
	}
}
