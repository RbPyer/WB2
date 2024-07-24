package wget

import (
	"io"
	"net/http"
	"os"
	"strings"
)

type IWget interface {
	Wget(url string) error
}

func NewWget() IWget {
	return &ImplWget{}
}

type ImplWget struct{}

func (w *ImplWget) Wget(url string) error {
	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create("index.html")
	defer file.Close()
	if err != nil {
		return err
	}

	_, err = io.Copy(file, response.Body)
	return err
}
