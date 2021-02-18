package tvmaze

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Results []map[string]string
}

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

	var response Response
	json.Unmarshal(data, &response)

	encjson, _ := json.Marshal(response.Results)

	return encjson
}