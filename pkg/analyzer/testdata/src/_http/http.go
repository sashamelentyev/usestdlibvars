package _http

import "net/http"

func _200() {
	_ = 200 // want `can use http.StatusOK instead "200"`
}

func _200_1() {
	var w http.ResponseWriter
	w.WriteHeader(200) // want `can use http.StatusOK instead "200"`
}
