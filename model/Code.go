package model

type Message struct {
	Code int
	Info string
}

func GetMessage(code int, info string) Message {
	message := Message{
		Code: code,
		Info: info,
	}
	return message
}
