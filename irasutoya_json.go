package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
)

// IrasutoyaJSON はいらすとやから帰ってくるJSONを格納する
type IrasutoyaJSON struct {
	Feed Feed `json:"feed"`
}

// Feed は IrasutoyaJSON を構成する
type Feed struct {
	OpenSearchTotalResults OpenSearchTotalResults `json:"openSearch$totalResults"`
	Entry                  []Entry                `json:"entry"`
}

// OpenSearchTotalResults は IrasutoyaJSON を構成する
type OpenSearchTotalResults struct {
	T string `json:"$t"`
}

// Entry は IrasutoyaJSON を構成する
type Entry struct {
	Title          Title          `json:"title"`
	Summary        Summary        `json:"summary"`
	MediaThumbnail MediaThumbnail `json:"media$thumbnail"`
}

// Title は IrasutoyaJSON を構成する
type Title struct {
	T string `json:"$t"`
}

// Summary は IrasutoyaJSON を構成する
type Summary struct {
	T string `json:"$t"`
}

// MediaThumbnail は IrasutoyaJSON を構成する
type MediaThumbnail struct {
	URL string `json:"url"`
}

// NewIrasutoyaJSON はいらすとやから返ってきた string を IrasutoyaJSON にする
func NewIrasutoyaJSON(s string) (*IrasutoyaJSON, error) {
	// JSON部分を取り出す
	regJSON := regexp.MustCompile(`{.*}`)
	JSON := regJSON.FindString(s)

	var irasutoyaJSON IrasutoyaJSON
	if err := json.Unmarshal([]byte(JSON), &irasutoyaJSON); err != nil {
		return nil, err
	}
	return &irasutoyaJSON, nil
}

// IrasutoyaURL はランダムなインデックスを含んだいらすとやのURLを返す
func IrasutoyaURL(idx int) string {
	return "http://www.irasutoya.com/feeds/posts/summary?start-index=" +
		strconv.Itoa(idx) +
		"&max-results=1&alt=json-in-script"
}

// FetchIrasutoyaJSON はidxに対応したIrasutoyaJSONを返す
func FetchIrasutoyaJSON(idx int) (*IrasutoyaJSON, error) {
	url := IrasutoyaURL(idx)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	irasutoyaJSON, err := NewIrasutoyaJSON(string(bytes))
	if err != nil {
		return nil, err
	}

	return irasutoyaJSON, nil
}

// FetchIrasutoyaMaxIndex はいらすとやの現在の最大インデックスを調べて返す
func FetchIrasutoyaMaxIndex() (idx int, err error) {
	// 現状でインデックスは20000より大きいので1-20000のインデックスのどれかにする
	seed := rand.Intn(20000)
	seed++

	irasutoyaJSON, err := FetchIrasutoyaJSON(seed)
	if err != nil {
		return 0, err
	}

	idx, err = strconv.Atoi(irasutoyaJSON.Feed.OpenSearchTotalResults.T)
	if err != nil {
		return 0, err
	}

	return idx, nil
}

// FetchRandomIrasutoyaJSON はランダムなIrasutoyaJSONを返す
func FetchRandomIrasutoyaJSON() (*IrasutoyaJSON, error) {
	maxIdx, err := FetchIrasutoyaMaxIndex()
	if err != nil {
		return nil, err
	}

	randIdx := rand.Intn(maxIdx)
	randIdx++

	irasutoyaJSON, err := FetchIrasutoyaJSON(randIdx)
	if err != nil {
		return nil, err
	}

	return irasutoyaJSON, nil
}
