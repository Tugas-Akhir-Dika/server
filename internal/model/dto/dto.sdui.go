package dto

import "SDUI_Server/internal/model/components"

type SDUIResponseDTO struct {
	Title  string                 `json:"title"`
	Header []components.Component `json:"header"`
	Body   []components.Component `json:"body"`
}
