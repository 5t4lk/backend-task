package requests

import (
	"io/ioutil"
	"net/http"
)

const (
	sourceLatestN   = "https://www.brentfordfc.com/api/incrowd/getnewlistinformation?count="
	sourceArticleID = "https://www.brentfordfc.com/api/incrowd/getnewsarticleinformation?id="
)

func ReqLatestArticles(l string) ([]byte, error) {
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

func ReqArticleByID(ID string) ([]byte, error) {
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
