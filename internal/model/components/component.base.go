package components

const (
	STACK_TYPE        = "stack"
	TEXT_TYPE         = "text"
	TEXT_FIELD_TYPE   = "text_field"
	SPACER_TYPE       = "spacer"
	BUTTON_TYPE       = "button"
	NAVIGATION_TYPE   = "navigation"
	CIRCULAR_IMG_TYPE = "circular_img"
	REGULAR_IMG_TYPE  = "reg_img"
	STACK_VERTICAL    = "vertical"
	STACK_HORIZONTAL  = "horizontal"
	CAROUSEL_TYPE     = "carousel"
)

var WHITE_COLOR = "#FFFFFF"
var BLUE_COLOR = "#0000FF"

type Component struct {
	Type        string      `json:"type"`
	Information interface{} `json:"info"`
}

type CarouselComponentInfo struct {
	Uid        string               `json:"uid"`
	Properties []CarouselProperties `json:"properties"`
}

type CarouselProperties struct {
	Title    string `json:"title"`
	Image    string `json:"image"`
	SubTitle string `json:"sub_title"`
}

type StackComponentInfo struct {
	Uid      string      `json:"uid"`
	Type     string      `json:"type"`
	Children []Component `json:"children"`
}

type TextComponentInfo struct {
	Uid      string   `json:"uid"`
	Message  string   `json:"message"`
	Size     float64  `json:"size"`
	Alpha    *float32 `json:"alpha"`
	ColorHex *string  `json:"color_hex"`
}

type SpacerComponentInfo struct {
	Uid    string   `json:"uid"`
	Length *float64 `json:"length"`
}

const SIGN_IN_BTN = "sign_in_btn"
const SIGN_OUT_BTN = "sign_out_btn"
const EXIT_BTN = "exit_btn"
const OPEN_PROFILE_BTN = "open_profile_btn"
const OPEN_LIVE_BTN = "open_live_btn"

type ButtonComponentInfo struct {
	Uid     string    `json:"uid"`
	Child   Component `json:"child"`
	IsClear bool      `json:"is_clear"`
}

type NavigationComponentInfo struct {
	Uid       string    `json:"uid"`
	TargetURL string    `json:"to"` // example, open item
	Child     Component `json:"child"`
}

type RegularImageComponentInfo struct {
	Uid      string  `json:"uid"`
	ImageURL string  `json:"image_url"`
	Height   float64 `json:"height"`
}

type CircularImageComponentInfo struct {
	Uid      string  `json:"uid"`
	ImageURL string  `json:"image_url"`
	Size     float64 `json:"size"`
}

const (
	USERNAME_TXTFLD  = "username"
	PASSSWORD_TXTFLD = "password"
)

type TextFieldComponentInfo struct {
	Uid      string `json:"uid"`
	IsSecure bool   `json:"is_secure"`
	Title    string `json:"title"`
}

type HorizontalScrollComponentInfo struct {
	Uid      string      `json:"uid"`
	Children []Component `json:"children"`
}
