package miniq

import (
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func fetch(url string) (io.Reader, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return strings.NewReader(string(body)), nil
}
