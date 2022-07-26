package http

import "net/http"

const foo = 404

func _200() {
	_ = 200
}

func _200_1() {
	var w http.ResponseWriter
	w.WriteHeader(200) // want `"200" can be replaced by http\.StatusOK`
}

func _GET() {
	_ = "GET"
}

func _GET_1() {
	_, _ = http.NewRequest("GET", "", nil) // want `"GET" can be replaced by http\.MethodGet`
}

func _GET_2() {
	_, _ = http.NewRequestWithContext(nil, "GET", "", nil) // want `"GET" can be replaced by http\.MethodGet`
}

func _GET_3() {
	_, _ = http.NewRequestWithContext(nil, "GET", "", nil) // want `"GET" can be replaced by http\.MethodGet`
}
