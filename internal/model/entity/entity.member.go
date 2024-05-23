package entity

type MemberEntity struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	SubTitle string `json:"subTitle"`
	PhotoURL string `json:"photoUrl"`
}
