package main

import (
	"encoding/json"
	"regexp"
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
