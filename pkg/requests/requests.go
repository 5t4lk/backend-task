package requests

import (
	"io/ioutil"
	"net/http"
)

func reqLatestArticles(l string) ([]byte, error) {
	r, err := http.Get(sourceLatestN + l)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	byteData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return byteData, nil
}

func reqArticleByID(ID string) ([]byte, error) {
	r, err := http.Get(sourceArticleID + ID)
	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	byteData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	return byteData, nil
}
