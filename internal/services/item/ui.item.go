package item

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

func (ui *UserInterface) CreateItemCardInterface(items []entity.ItemEntity) []components.Component {
	compos := make([]components.Component, 0)
	for _, item := range items {
		compos = append(compos, components.Component{
			Type: components.NAVIGATION_TYPE,
			Information: components.NavigationComponentInfo{
				Uid:       fmt.Sprintf("/item/%d", item.Id),
				TargetURL: fmt.Sprintf("/item/%d", item.Id),
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
									ImageURL: item.Image,
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
												Message: item.Title,
												Size:    18,
											},
										},
										{
											Type: components.TEXT_TYPE,
											Information: components.TextComponentInfo{
												Uid:     uuid.NewString(),
												Message: item.Category,
												Size:    8,
											},
										},
										{
											Type: components.TEXT_TYPE,
											Information: components.TextComponentInfo{
												Uid:     uuid.NewString(),
												Message: fmt.Sprintf("$%.2f", item.Price),
												Size:    10,
											},
										}, {
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
			},
		})
	}
	return compos
}

func (ui *UserInterface) CreateItemDetailInterface(item entity.ItemEntity) []components.Component {
	return []components.Component{
		{
			Type: components.TEXT_TYPE,
			Information: components.TextComponentInfo{
				Uid:     uuid.NewString(),
				Message: item.Title,
				Size:    10,
			},
		},
		{
			Type: components.NAVIGATION_TYPE,
			Information: components.NavigationComponentInfo{
				Uid:       uuid.NewString(),
				TargetURL: "/item/2",
				Child: components.Component{
					Type: components.TEXT_TYPE,
					Information: components.TextComponentInfo{
						Uid:     uuid.NewString(),
						Message: item.Category,
						Size:    10,
					},
				},
			},
		},
	}
}

func (ui *UserInterface) CreateNotFoundInterface() []components.Component {
	return []components.Component{
		{
			Type: components.TEXT_TYPE,
			Information: components.TextComponentInfo{
				Uid:     uuid.NewString(),
				Message: "Cannot Find Item",
				Size:    10,
			},
		},
	}
}

func (ui *UserInterface) CreateHeader(isAuth bool) []components.Component {
	if !isAuth {
		return make([]components.Component, 0)
	}
	return []components.Component{
		{
			Type: components.CIRCULAR_IMG_TYPE,
			Information: components.CircularImageComponentInfo{
				Uid:      uuid.NewString(),
				ImageURL: "https://pbs.twimg.com/media/GGyPlsyW0AAhnbH?format=jpg",
				Size:     34,
			},
		},
	}
}
