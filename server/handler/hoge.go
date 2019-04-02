package handler

import (
	"fmt"
	"net/http"
)

func Hoge(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hoge")
}
