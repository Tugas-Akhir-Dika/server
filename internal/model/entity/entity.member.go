package entity

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2/log"
	"io"
	"os"
)

type MemberEntity struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	SubTitle string `json:"subTitle"`
	PhotoURL string `json:"photoUrl"`
}

func LoadMembers() ([]MemberEntity, error) {
	filePath := "./files/members.json"
	members := make([]MemberEntity, 0)

	// Open the JSON file
	file, err := os.Open(filePath)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer file.Close()

	// Read the file contents
	fileData, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Parse the JSON data into the Config struct
	err = json.Unmarshal(fileData, &members)
	if err != nil {
		log.Error(err)
	}
	return members, err
}

type StreamingEntity struct {
	StreamingUrlList []StreamingListEntity `json:"streaming_url_list"`
}

type StreamingListEntity struct {
	IsDefault bool   `json:"is_default"`
	Url       string `json:"url"`
	Type      string `json:"type"`
	Id        int    `json:"id"`
	Label     string `json:"label"`
	Quality   int    `json:"quality"`
}
