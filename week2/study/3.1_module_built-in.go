// <3.1 빌트인 패키지 사용하기>
// Go 언어는 기본적으로 제공하는 빌트인 패키지를 사용하여 다양한 기능을 구현할 수 있습니다.
package study

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// 코드 3.1: OMDb API를 사용한 영화 정보 검색
/*
Go에서 빌트인 패키지만 사용하여
rest API로 영화 세부 정보를 검색하는 예시
*/

// omdbapi.com API key
const APIKEY = "193ef3a"

// omdbapi.com에서 반환된 JSON 구조체
// 예시의 간략화를 위해 일부 값은 구조체에 매핑하지 않음
type MovieInfo struct {
	Title      string `json:"Title"`
	Year       string `json:"Year"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Language   string `json:"Language"`
	Country    string `json:"Country"`
	Awards     string `json:"Awards"`
	Poster     string `json:"Poster"`
	ImdbRating string `json:"imdbRating"`
	ImdbID     string `json:"imdbID"`
}

func Code31() {
	body, _ := SearchById("tt3896198")
	fmt.Println(body.Title)
	body, _ = SearchByName("Game of")
	fmt.Println(body.Title)
}

func SearchByName(name string) (*MovieInfo, error) {
	params := url.Values{}
	params.Set("apikey", APIKEY)
	params.Set("t", name)
	siteURL := "http://www.omdbapi.com/?" + params.Encode()
	body, err := sendGetRequest(siteURL)
	if err != nil {
		return nil, errors.New(err.Error() + "\nBody: " + string(body))
	}
	mi := &MovieInfo{}
	return mi, json.Unmarshal([]byte(body), mi)
}

func SearchById(id string) (*MovieInfo, error) {
	params := url.Values{}
	params.Set("apikey", APIKEY)
	params.Set("i", id)
	siteURL := "http://www.omdbapi.com/?" + params.Encode()
	body, err := sendGetRequest(siteURL)
	if err != nil {
		return nil, errors.New(err.Error() + "\nBody: " + string(body))
	}
	mi := &MovieInfo{}
	return mi, json.Unmarshal([]byte(body), mi)
}

func sendGetRequest(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return string(body), errors.New(resp.Status)
	}
	return string(body), nil
}
