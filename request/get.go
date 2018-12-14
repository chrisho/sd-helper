package request

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)


func Get(url string) (r string, err error) {
	request := &http.Client{}
	response, err := request.Get(url)
	if err != nil {
		return r, err
	}
	if response.StatusCode != 200 {
		return r, errors.New("request fail code: " + fmt.Sprintf("%d", response.StatusCode))
	}
	byteSlice, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return r, err
	}
	return string(byteSlice), nil
}
