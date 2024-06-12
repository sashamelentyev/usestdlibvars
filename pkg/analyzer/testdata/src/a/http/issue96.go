package http_test

import (
	"net/http"
)

func _() error {
	resp, err := http.DefaultClient.Do(&http.Request{})
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode/100 != 2 {
		return nil
	}
	if resp.StatusCode+100 != 2 {
		return nil
	}
	if resp.StatusCode-100 != 2 {
		return nil
	}
	if resp.StatusCode*100 != 2 {
		return nil
	}

	return nil
}
