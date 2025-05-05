# 🕹️ Go Telegram PC Controller Bot

A lightweight Go bot that lets you remotely control your computer via **Telegram**.  
Perfect for keeping your game in focus or shutting down your system remotely.

## ✨ Features

- 📩 Receive commands through a Telegram bot
- 🪟 `/focus`: Bring your game window to the foreground
- ⏱️ `/shutdown`: Schedule a computer shutdown (60-second delay)
- ❌ `/abort`: Cancel a scheduled shutdown
- 🆘 `/help`: Get a list of available commands

## 🧠 How It Works

This bot uses the Telegram Bot API to receive messages, and runs local OS-level commands on your machine accordingly.  
Ideal for remote control of a gaming PC or any unattended system running Windows.

## 📦 Requirements

- Built with Go 1.24.2
- A Telegram bot token (get it from [BotFather](https://t.me/BotFather))
- Windows OS (currently supports Windows-specific commands)

## 🚀 Getting Started

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/go-telegram-pc-controller.git
   cd go-telegram-pc-controller

2. Set the .env properties (Bot Token, and allowed user id)
3. 
3. go run main.go (you'll probably have to run as admin for some commands, for example window focus)