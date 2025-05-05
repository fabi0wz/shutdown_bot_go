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
3. go run main.go (you'll probably have to run as admin for some commands, for example window focus)

## 💡 Future Ideas
This section will be updated with upcoming features and experiments that I plan to build on top of the bot:

🎮 Simulate key presses or mouse movements (e.g. auto-reconnect in games)

📸 Capture and send screenshots on request

🔊 Control volume or media playback remotely

🗂️ File access: List, send, or delete specific files via Telegram

🧠 Basic automation scripting via chat

🌐 Wake-on-LAN support or remote start from mobile

🪄 Custom game macros triggered by Telegram commands

Feel free to fork and suggest your own features via pull requests or issues!
