package lark

import "fmt"

type TextMessage struct {
	content string
}

func NewText(format string, args ...any) *TextMessage {
	return &TextMessage{content: fmt.Sprintf(format, args...)}
}

func (t *TextMessage) MsgType() string { return "text" }

func (t *TextMessage) Build() any {
	return map[string]any{
		"msg_type": "text",
		"content": map[string]string{
			"text": t.content,
		},
	}
}
