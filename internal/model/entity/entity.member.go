package entity

type MemberEntity struct {
	Id       int64  `db:"id"`
	Name     string `db:"name"`
	SubTitle string `db:"sub_title"`
	PhotoURL string `db:"photo_url"`
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
