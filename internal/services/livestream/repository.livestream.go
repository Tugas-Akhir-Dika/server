package livestream

import (
	"SDUI_Server/internal/model/entity"
	"fmt"
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
	return roomUrl, nil
}
