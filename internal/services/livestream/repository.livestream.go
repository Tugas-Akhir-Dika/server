package livestream

import (
	"SDUI_Server/internal/model/entity"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"io"
	"net/http"
)

type Repository struct {
	members []entity.MemberEntity
}

const liveURL = "https://www.showroom-live.com/api/live/streaming_url?room_id=%d"

func NewRepository(members []entity.MemberEntity) *Repository {
	return &Repository{members: members}
}

func (r *Repository) GetMembers() []entity.MemberEntity {
	return r.members
}

func (r *Repository) GetMemberLiveStreamURL(roomId int64) (string, error) {
	roomUrl := fmt.Sprintf(liveURL, roomId)
	// Create a new HTTP client
	client := &http.Client{}

	// Create a new GET request
	req, err := http.NewRequest(http.MethodGet, roomUrl, nil)
	if err != nil {
		log.Error(err)
		return "", err
	}

	// Send the request and get the response
	resp, err := client.Do(req)
	if err != nil {
		log.Error(err)
		return "", err
	}

	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return "", err
	}

	var streamingEntity entity.StreamingEntity

	err = json.Unmarshal(body, &streamingEntity)
	if err != nil {
		log.Error(err)
		return "", err
	}
	if len(streamingEntity.StreamingUrlList) < 1 {
		return "", errors.New("cannot find live stream url")
	}
	return streamingEntity.StreamingUrlList[0].Url, nil
}
