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
						Uid:     uuid.NewString(),
						Message: "Sign In",
						Size:    20,
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
	compos := make([]components.Component, 0)
	if !isAuth {
		return ui.UserNotFoundInterface()
	}
	compos = []components.Component{}
	return compos
}

func (ui *UserInterface) Header() []components.Component {
	return make([]components.Component, 0)
}
