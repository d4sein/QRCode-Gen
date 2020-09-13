package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	qrcode "github.com/skip2/go-qrcode"
)

type qrcodeStruct struct {
	Content string
	Size    int
}

func newQRCode() *qrcodeStruct {
	var qrcodeSpecs qrcodeStruct
	return &qrcodeSpecs
}

func setupResponse(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "data, X-Requested-With, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	(*w).Header().Set("Content-Type", "application/json")
}

func main() {
	port := "localhost:3000"
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Options("/", func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w)
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		setupResponse(&w)

		body, err := ioutil.ReadAll(r.Body)

		if err != nil || len(body) == 0 {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		qrcodeSpecs := newQRCode()
		err = json.Unmarshal(body, &qrcodeSpecs)

		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		png, err := qrcode.Encode(qrcodeSpecs.Content, qrcode.Medium, qrcodeSpecs.Size)

		if err != nil {
			http.Error(w, http.StatusText(204), 204)
			return
		}

		res := map[string][]byte{}
		res["content"] = png

		json.NewEncoder(w).Encode(res)
	})

	fmt.Println("Serving on http://" + port)
	http.ListenAndServe(port, r)
}
