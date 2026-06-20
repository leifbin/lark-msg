package lark

import "fmt"

type CardMessage struct {
	header   *cardHeader
	elements []interface{}
}

type cardHeader struct {
	Title string
	Color Color
}

func NewCard() *CardMessage {
	return &CardMessage{elements: make([]interface{}, 0)}
}

func (c *CardMessage) Title(title string, color Color) *CardMessage {
	c.header = &cardHeader{Title: title, Color: color}
	return c
}

func (c *CardMessage) Text(format string, args ...interface{}) *CardMessage {
	content := format
	if len(args) > 0 {
		content = fmt.Sprintf(format, args...)
	}
	c.elements = append(c.elements, map[string]string{
		"tag": "markdown", "content": content,
	})
	return c
}

func (c *CardMessage) Table(header1, header2 string, rows ...[2]string) *CardMessage {
	content := "**" + header1 + "** | **" + header2 + "**\n--- | ---\n"
	for _, row := range rows {
		content += row[0] + " | " + row[1] + "\n"
	}
	c.elements = append(c.elements, map[string]string{
		"tag": "markdown", "content": content,
	})
	return c
}

func (c *CardMessage) HR() *CardMessage {
	c.elements = append(c.elements, map[string]string{"tag": "hr"})
	return c
}

func (c *CardMessage) Button(text, url string) *CardMessage {
	c.elements = append(c.elements, map[string]interface{}{
		"tag": "action",
		"actions": []map[string]interface{}{{
			"tag":  "button",
			"text": map[string]string{"tag": "plain_text", "content": text},
			"url":  url, "type": "default",
		}},
	})
	return c
}

func (c *CardMessage) Note(format string, args ...interface{}) *CardMessage {
	content := format
	if len(args) > 0 {
		content = fmt.Sprintf(format, args...)
	}
	c.elements = append(c.elements, map[string]interface{}{
		"tag":      "note",
		"elements": []map[string]string{{"tag": "plain_text", "content": content}},
	})
	return c
}

func (c *CardMessage) MsgType() string { return "interactive" }

func (c *CardMessage) Build() interface{} {
	card := map[string]interface{}{"elements": c.elements}
	if c.header != nil {
		card["header"] = map[string]interface{}{
			"title":    map[string]string{"tag": "plain_text", "content": c.header.Title},
			"template": string(c.header.Color),
		}
	}
	return map[string]interface{}{"msg_type": "interactive", "card": card}
}
