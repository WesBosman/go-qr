package main

import (
	"fmt"
	"net/http"
	"text/template"

	qrcode "github.com/skip2/go-qrcode"
)

type PageData struct {
	Url string
	Src string
	Err string
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	hasUrl := r.URL.Query().Has("url")
	url := r.URL.Query().Get("url")
	templ := template.Must(template.ParseFiles("layout.html"))
	data := PageData{
		Url: url,
		Src: "/assets/qr.png",
		Err: "",
	}

	if hasUrl {
		err := qrcode.WriteFile(url, qrcode.Medium, 256, "./assets/qr.png")

		if err != nil {
			data.Err = fmt.Sprintf("Error: %v", err)
		}
	} else {
		data.Err = "Enter a URL query parameter"
	}

	templ.Execute(w, data)
}

func main() {
	port := 3333
	portStr := fmt.Sprintf(":%v", port)
	url := fmt.Sprintf("http://localhost%v", portStr)
	fmt.Println("Running server on port ", url)

	http.HandleFunc("/", getRoot)
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))
	err := http.ListenAndServe(portStr, nil)

	if err != nil {
		fmt.Println("Error: ", err)
	}

}
