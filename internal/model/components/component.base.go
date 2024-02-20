package components

const (
	STACK_TYPE        = "stack"
	TEXT_TYPE         = "text"
	TEXT_FIELD_TYPE   = "text_field"
	SPACER_TYPE       = "spacer"
	BUTTON_TYPE       = "button"
	NAVIGATION_TYPE   = "navigation"
	CIRCULAR_IMG_TYPE = "circular_img"
	NORMAL_IMG_TYPE   = "normal_img"
	STACK_VERTICAL    = "vertical"
	STACK_HORIZONTAL  = "horizontal"
)

type Component struct {
	Type        string      `json:"type"`
	Information interface{} `json:"info"`
}

type StackComponentInfo struct {
	Uid      string      `json:"uid"`
	Type     string      `json:"type"`
	Children []Component `json:"children"`
}

type TextComponentInfo struct {
	Uid     string  `json:"uid"`
	Message string  `json:"message"`
	Size    float64 `json:"size"`
}

type SpacerComponentInfo struct {
	Uid    string   `json:"uid"`
	Length *float64 `json:"length"`
}

const SIGN_IN_BTN = "sign_in_btn"

type ButtonComponentInfo struct {
	Uid   string    `json:"uid"`
	Child Component `json:"child"`
}

type NavigationComponentInfo struct {
	Uid       string    `json:"uid"`
	TargetURL string    `json:"to"` // example, open item
	Child     Component `json:"child"`
}

type ImageComponentInfo struct {
	Uid      string `json:"uid"`
	ImageURL string `json:"image_url"`
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
