package signin

import (
	"SDUI_Server/internal/model/components"
	"github.com/google/uuid"
)

type UserInterface struct {
}

func NewUserInterface() *UserInterface {
	return &UserInterface{}
}

func (ui *UserInterface) UserNotFoundInterface() []components.Component {
	l := 200.0
	l2 := 25.0
	return []components.Component{
		{
			Type: components.SPACER_TYPE,
			Information: components.SpacerComponentInfo{
				Uid:    "sdkdhfgdjkhg",
				Length: &l,
			},
		},
		{
			Type: components.TEXT_FIELD_TYPE,
			Information: components.TextFieldComponentInfo{
				Uid:      components.USERNAME_TXTFLD,
				IsSecure: false,
				Title:    "username",
			},
		}, {
			Type: components.TEXT_FIELD_TYPE,
			Information: components.TextFieldComponentInfo{
				Uid:      components.PASSSWORD_TXTFLD,
				IsSecure: true,
				Title:    "password",
			},
		},
		{
			Type: components.SPACER_TYPE,
			Information: components.SpacerComponentInfo{
				Uid:    "97aosdjsugsg",
				Length: &l2,
			},
		},
		{
			Type: components.BUTTON_TYPE,
			Information: components.ButtonComponentInfo{
				Uid: components.SIGN_IN_BTN,
				Child: components.Component{
					Type: components.TEXT_TYPE,
					Information: components.TextComponentInfo{
						Uid:      uuid.NewString(),
						Message:  "Sign In",
						Size:     20,
						ColorHex: &components.WHITE_COLOR,
					},
				},
			},
		},
		{
			Type: components.SPACER_TYPE,
			Information: components.SpacerComponentInfo{
				Uid:    uuid.NewString(),
				Length: nil,
			},
		},
	}
}

func (ui *UserInterface) UserFoundInterface(isAuth bool) []components.Component {
	l := 200.0
	compos := make([]components.Component, 0)
	if !isAuth {
		return ui.UserNotFoundInterface()
	}
	compos = append(compos, components.Component{
		Type: components.STACK_TYPE,
		Information: components.StackComponentInfo{
			Uid:  "d89yf389",
			Type: components.STACK_VERTICAL,
			Children: []components.Component{
				{
					Type: components.SPACER_TYPE,
					Information: components.SpacerComponentInfo{
						Uid:    "309dedc9ehdjewhc",
						Length: &l,
					},
				},
				{
					Type: components.BUTTON_TYPE,
					Information: components.ButtonComponentInfo{
						Uid: components.SIGN_OUT_BTN,
						Child: components.Component{
							Type: components.TEXT_TYPE,
							Information: components.TextComponentInfo{
								Uid:      uuid.NewString(),
								Message:  "Sign Out",
								Size:     20,
								ColorHex: &components.WHITE_COLOR,
							},
						},
					},
				},
			},
		},
	})
	return compos
}

func (ui *UserInterface) Header(isAuth bool) []components.Component {
	if !isAuth {
		return make([]components.Component, 0)
	}
	return []components.Component{
		{
			Type: components.BUTTON_TYPE,
			Information: components.ButtonComponentInfo{
				Uid: components.EXIT_BTN,
				Child: components.Component{
					Type: components.TEXT_TYPE,
					Information: components.TextComponentInfo{
						Uid:     uuid.NewString(),
						Message: "X",
						Size:    25,
					},
				},
			},
		},
	}
}
