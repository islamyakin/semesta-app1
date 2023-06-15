package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, `<body style="background-color:#000000">`)
		fmt.Fprint(w, `<center><h1 style="color:#FFFFFF">Selamat datang di Semesta System Administrator</center></h1>`)
		fmt.Fprint(w, `<center><img src="https://maukuliah.id/assets/img/semesta/logo-semesta-light.png" alt="Gambar"></center>`)
		fmt.Fprint(w, `</body>`)
	} else if r.URL.Path == "/aboutus" {
		err := godotenv.Load(".env")
		if err != nil {
			http.Error(w, "Gagal memuat file .env", http.StatusInternalServerError)
			fmt.Println("Gagal memuat file .env:", err)
			return
		}
		targetURL := os.Getenv("APP2_URL")
		if targetURL == "" {
			http.Error(w, "URL tujuan tidak ditentukan", http.StatusInternalServerError)
			fmt.Println("URL tujuan tidak ditentukan")
			return
		}
		resp, err := http.Get(targetURL)
		if err != nil {
			http.Error(w, fmt.Sprintf("Gagal memuat konten dari %s: %s", targetURL, err.Error()), http.StatusInternalServerError)
			fmt.Println("Gagal memuat konten:", err)
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			http.Error(w, "Gagal membaca respons", http.StatusInternalServerError)
			fmt.Println("Gagal membaca respons:", err)
			return
		}
		fmt.Fprintf(w, "%s", body)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>Halaman yang dicari tidak ditemukan</h1>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
