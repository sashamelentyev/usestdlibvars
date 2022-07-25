package _http

import "net/http"

func _200() {
	_ = 200
}

func _200_1() {
	var w http.ResponseWriter
	w.WriteHeader(200) // want `can use http.StatusOK instead "200"`
}

func _GET() {
	_ = "GET"
}

func _GET_1() {
	_, _ = http.NewRequest("GET", "", nil) // want `can use http.MethodGet instead "GET"`
}

func _GET_2() {
	_, _ = http.NewRequestWithContext(nil, "GET", "", nil) // want `can use http.MethodGet instead "GET"`
}

func _GET_3() {
	_, _ = http.NewRequestWithContext(nil, "GET", "", nil) // want `can use http.MethodGet instead "GET"`
}
