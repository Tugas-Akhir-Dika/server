package dto

type MemberResponseDTO struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	SubTitle string `json:"subTitle"`
	PhotoURL string `json:"photoUrl"`
}

type MemberLiveStreamURLResponseDTO struct {
	Id            int64  `json:"id"`
	LiveStreamURL string `json:"liveStreamUrl"`
}
