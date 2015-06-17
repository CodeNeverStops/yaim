package models

type GlobalEnv struct {
	sessionServer *SessionPool
	chatServer    *ChatServer
}

var App GlobalEnv

func RunApp() {
	App = &GlobalEnv{
		sessionServer: NewSessionPool(),
		chatServer: NewChatServer(),
	}
}