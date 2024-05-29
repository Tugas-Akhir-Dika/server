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

func (ui *UserInterface) CreateItemListChildInterface(items []entity.ItemEntity) []components.Component {
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
	spl := 3.0
	var gray = "#808080"
	splTwo := 10.0
	compo := []components.Component{
		{
			Type: components.SPACER_TYPE,
			Information: components.SpacerComponentInfo{
				Uid:    uuid.NewString(),
				Length: &spl,
			},
		},
		{
			Type: components.REGULAR_IMG_TYPE,
			Information: components.RegularImageComponentInfo{
				Uid:      uuid.NewString(),
				ImageURL: item.Image,
				Height:   300,
			},
		},
		{
			Type: components.STACK_TYPE,
			Information: components.StackComponentInfo{
				Uid:  uuid.NewString(),
				Type: components.STACK_VERTICAL,
				Children: []components.Component{
					{
						Type: components.TEXT_TYPE,
						Information: components.TextComponentInfo{
							Uid:      uuid.NewString(),
							Message:  item.Title,
							Size:     30,
							Alpha:    nil,
							ColorHex: nil,
						},
					},
					{
						Type: components.TEXT_TYPE,
						Information: components.TextComponentInfo{
							Uid:      uuid.NewString(),
							Message:  item.Category,
							Size:     10,
							Alpha:    nil,
							ColorHex: &gray,
						},
					},
					{
						Type: components.TEXT_TYPE,
						Information: components.TextComponentInfo{
							Uid:     uuid.NewString(),
							Message: fmt.Sprintf("$%.2f", item.Price),
							Size:    20,
						},
					},
					{
						Type: components.SPACER_TYPE,
						Information: components.SpacerComponentInfo{
							Uid:    uuid.NewString(),
							Length: &splTwo,
						},
					},
					{
						Type: components.TEXT_TYPE,
						Information: components.TextComponentInfo{
							Uid:      uuid.NewString(),
							Message:  item.Description,
							Size:     12,
							Alpha:    nil,
							ColorHex: nil,
						},
					},
				},
			},
		},
	}
	return compo
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
	if isAuth {
		return []components.Component{
			{
				Type: components.BUTTON_TYPE,
				Information: components.ButtonComponentInfo{
					Uid: components.OPEN_LIVE_BTN,
					Child: components.Component{
						Type: components.TEXT_TYPE,
						Information: components.TextComponentInfo{
							Uid:     uuid.NewString(),
							Message: "Live",
							Size:    15,
						},
					},
					IsClear: true,
				},
			},
			{
				Type: components.BUTTON_TYPE,
				Information: components.ButtonComponentInfo{
					Uid:     components.OPEN_PROFILE_BTN,
					IsClear: true,
					Child: components.Component{
						Type: components.CIRCULAR_IMG_TYPE,
						Information: components.CircularImageComponentInfo{
							Uid:      uuid.NewString(),
							ImageURL: "https://pbs.twimg.com/media/GGyPlsyW0AAhnbH?format=jpg",
							Size:     34,
						},
					},
				},
			},
		}
	}
	return []components.Component{
		{
			Type: components.BUTTON_TYPE,
			Information: components.ButtonComponentInfo{
				Uid: components.OPEN_LIVE_BTN,
				Child: components.Component{
					Type: components.TEXT_TYPE,
					Information: components.TextComponentInfo{
						Uid:     uuid.NewString(),
						Message: "Live",
						Size:    15,
					},
				},
				IsClear: true,
			},
		},
		{
			Type: components.BUTTON_TYPE,
			Information: components.ButtonComponentInfo{
				Uid: components.OPEN_PROFILE_BTN,
				Child: components.Component{
					Type: components.TEXT_TYPE,
					Information: components.TextComponentInfo{
						Uid:     uuid.NewString(),
						Message: "Login",
						Size:    15,
					},
				},
				IsClear: true,
			},
		},
		//{
		//	Type: components.BUTTON_TYPE,
		//	Information: components.ButtonComponentInfo{
		//		Uid:     components.OPEN_PROFILE_BTN,
		//		IsClear: true,
		//		Child: components.Component{
		//			Type: components.CIRCULAR_IMG_TYPE,
		//			Information: components.CircularImageComponentInfo{
		//				Uid:      uuid.NewString(),
		//				ImageURL: "https://t4.ftcdn.net/jpg/04/00/24/31/360_F_400243185_BOxON3h9avMUX10RsDkt3pJ8iQx72kS3.jpg",
		//				Size:     34,
		//			},
		//		},
		//	},
		//},
	}
}

func (ui *UserInterface) CreateSuggestionCarouselInterface(items []entity.ItemEntity) []components.Component {
	compos := make([]components.Component, 0)
	props := make([]components.CarouselProperties, 0)
	for _, item := range items {
		props = append(props, components.CarouselProperties{
			Title:    item.Title,
			Image:    item.Image,
			SubTitle: item.Category,
		})
	}
	carousel := components.Component{
		Type: components.CAROUSEL_TYPE,
		Information: components.CarouselComponentInfo{
			Uid:        "",
			Properties: props,
		},
	}
	compos = append(compos, carousel)
	return compos
}
