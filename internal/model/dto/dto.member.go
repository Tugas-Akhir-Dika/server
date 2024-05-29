package dto

type MemberResponseDTO struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	SubTitle  string `json:"subTitle"`
	PhotoURL  string `json:"photoUrl"`
	BornDate  string `json:"born_date"`
	Horoscope string `json:"horoscope"`
	Height    int64  `json:"height"`
	Jiko      string `json:"jiko"`
}

type MemberLiveStreamURLResponseDTO struct {
	Id            int64  `json:"id"`
	LiveStreamURL string `json:"liveStreamUrl"`
}
