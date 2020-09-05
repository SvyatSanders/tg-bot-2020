package main

// Config - конфигурация бота
type Config struct {
	tgToken       string
	botName       string
	username      string
	herokuAppName string
	vkAppName     string
	vkServiceKey  string
	vkAppID       string
}

var botConfig = Config{
	tgToken:       "1003625678:AAFWvu23-E_5TcmeBXo0QeBTUNLMdZg28As",
	botName:       "SandersBot",
	username:      "SvyatSandersBot",
	herokuAppName: "sandersbot",
	vkAppName:     "sandersbot",
	vkServiceKey:  "69cde26069cde26069cde2603869be209a669cd69cde260369c3ee8230f5bcdf28661eb",
	vkAppID:       "7586554",
}
