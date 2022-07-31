package chatbot

import "errors"

// Talk 定义了聊天的接口类型
type Talk interface {
	Hello(userName string) string
	Talk(heard string) (saying string, end bool, err error)
}

// Chatbot 定义了聊天机器人的接口类型
type Chatbot interface {
	Name() string
	Begin() (string, error)
	Talk
	ReportError(err error) string
	End() error
}

var (
	ErrInvalidaChatbotName = errors.New("Invalid chatbot name")
	ErrInvalidChatbot      = errors.New("Invalid chatbot")
	ErrExistingChatbot     = errors.New("Existing chatbot")
)

var chatbotMap = map[string]Chatbot{}

// 注册聊天机器人
func Register(chatbot Chatbot) error {
	if chatbot == nil {
		return ErrInvalidChatbot
	}
	name := chatbot.Name()
	if name == "" {
		return ErrInvalidaChatbotName
	}
	if _, ok := chatbotMap[name]; ok {
		return ErrExistingChatbot
	}
	chatbotMap[name] = chatbot
	return nil
}

func Get(name string) Chatbot {
	return chatbotMap[name]
}
