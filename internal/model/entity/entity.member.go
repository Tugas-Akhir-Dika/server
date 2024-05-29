package entity

import "database/sql"

type MemberEntity struct {
	Id        int64        `db:"id"`
	Name      string       `db:"name"`
	SubTitle  string       `db:"sub_title"`
	PhotoURL  string       `db:"photo_url"`
	BornDate  sql.NullTime `db:"born_date"`
	Horoscope string       `db:"horoscope"`
	Height    int64        `db:"height"`
	Jiko      string       `db:"jiko"`
}

type StreamingEntity struct {
	StreamingUrlList []StreamingListEntity `db:"streaming_url_list"`
}

type StreamingListEntity struct {
	IsDefault bool   `db:"is_default"`
	Url       string `db:"url"`
	Type      string `db:"type"`
	Id        int    `db:"id"`
	Label     string `db:"label"`
	Quality   int    `db:"quality"`
}
