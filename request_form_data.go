package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Print("hello")

	http.HandleFunc("/images", func(w http.ResponseWriter, r *http.Request) {
		//r.FormFile("")
		fk := r.FormValue("k")
		fmt.Printf("[FK]:=>%s\n", fk)

		file, sdz, err := r.FormFile("img")
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}
		fmt.Printf("sdz.Filename=>%s\nsdz.Size==>%v\nmedia=>%s\n", sdz.Filename, sdz.Size, sdz.Header.Get("Content-Type"))

		// copy example
		f, err := os.OpenFile(sdz.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		defer f.Close()
		io.Copy(f, file)
		w.Write([]byte("ok"))
	})

	http.ListenAndServe(":8080", nil)

}
