package main

import (
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/cookie-set", func(w http.ResponseWriter, r *http.Request) {

		expiration := time.Now().Add(24 * time.Hour)
		cookie := http.Cookie{Name: "d", Value: "i am cookie", Expires: expiration}
		http.SetCookie(w, &cookie)
		w.Write([]byte("设置cookie ok ==> d:i am cookie"))
	})

	http.ListenAndServe(":8009", nil)
}
