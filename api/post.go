package api

import (
	"log"
	"net/http"
)

func postImage(w http.ResponseWriter, req *http.Request) {
	//err := req.ParseMultipartForm(32 << 20)
	//if err != nil {
	//	_, err := w.Write([]byte("no"))
	//}
	//for _, item := range req.MultipartForm.File {
	//
	//}
	//req.MultipartForm.Value[0]
}

func serve() {
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
