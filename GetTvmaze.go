package tvmaze

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetTvmaze(query string) []byte {
	res, err := http.Get("http://api.tvmaze.com/search/shows?q=" + query)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return data
}