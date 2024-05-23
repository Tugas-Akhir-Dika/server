package livestream

import (
	"SDUI_Server/internal/model/components"
	"SDUI_Server/internal/model/entity"
	"fmt"
	"github.com/google/uuid"
)

type UserInterface struct {
}

func NewUserInterface() *UserInterface {
	return &UserInterface{}
}

func (ui *UserInterface) CreateMemberListChildInterface(members []entity.MemberEntity) []components.Component {
	compos := make([]components.Component, 0)
	for _, member := range members {
		compos = append(compos, components.Component{
			Type: components.BUTTON_TYPE,
			Information: components.ButtonComponentInfo{
				Uid: fmt.Sprintf("livestream/%d", member.Id),
				Child: components.Component{
					Type: components.STACK_TYPE,
					Information: components.StackComponentInfo{
						Uid:  uuid.NewString(),
						Type: components.STACK_HORIZONTAL,
						Children: []components.Component{
							{
								Type: components.CIRCULAR_IMG_TYPE,
								Information: components.CircularImageComponentInfo{
									Uid:      uuid.NewString(),
									ImageURL: member.PhotoURL,
									Size:     80,
								},
							}, {
								Type: components.STACK_TYPE,
								Information: components.StackComponentInfo{
									Uid:  uuid.NewString(),
									Type: components.STACK_VERTICAL,
									Children: []components.Component{
										{
											Type: components.TEXT_TYPE,
											Information: components.TextComponentInfo{
												Uid:     uuid.NewString(),
												Message: member.Name,
												Size:    18,
											},
										},
										{
											Type: components.TEXT_TYPE,
											Information: components.TextComponentInfo{
												Uid:     uuid.NewString(),
												Message: member.SubTitle,
												Size:    8,
											},
										},
										{
											Type: components.SPACER_TYPE,
											Information: components.SpacerComponentInfo{
												Uid:    uuid.NewString(),
												Length: nil,
											},
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
						},
					},
				},
				IsClear: true,
			},
		})
	}
	return compos
}
