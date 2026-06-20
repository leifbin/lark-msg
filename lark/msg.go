package lark

type Message interface {
	MsgType() string
	Build() any
}
