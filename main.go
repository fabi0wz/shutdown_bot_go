package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

var shutdownTimer *time.Timer

// Command handler type
type commandHandler struct {
	handler     func(bot *tgbotapi.BotAPI, chatID int64)
	description string
}

// Define command handlers globally so they can be accessed from any function
var commandHandlers map[string]commandHandler

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	// Get bot token and allowed chat ID from environment variables
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is required in .env file")
	}

	allowedChatIDStr := os.Getenv("ALLOWED_CHAT_ID")
	if allowedChatIDStr == "" {
		log.Fatal("ALLOWED_CHAT_ID is required in .env file")
	}

	allowedChatID, err := strconv.ParseInt(allowedChatIDStr, 10, 64)
	if err != nil {
		log.Fatal("ALLOWED_CHAT_ID must be a valid integer: ", err)
	}

	// Optional environment variable for the game window title
	gameWindowTitle := os.Getenv("GAME_WINDOW_TITLE")
	if gameWindowTitle == "" {
		gameWindowTitle = "TwelveSky2" // Default value
	}

	// Initialize command handlers
	commandHandlers = map[string]commandHandler{
		"/shutdown": {
			handler:     handleShutdown,
			description: "Schedule a PC shutdown in 60 seconds",
		},
		"/abort": {
			handler:     handleAbort,
			description: "Cancel scheduled shutdown",
		},
		"/focus": {
			handler: func(bot *tgbotapi.BotAPI, chatID int64) {
				focusWindow(bot, chatID, gameWindowTitle)
			},
			description: "Focus the game window if minimized",
		},
		"/help": {
			handler:     handleHelp,
			description: "Display available commands",
		},
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic("Failed to connect to Telegram Bot:", err)
	}

	bot.Debug = false // Set to true for debugging
	log.Printf("Bot authorized on account @%s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		chatID := update.Message.Chat.ID
		userText := update.Message.Text

		// Print incoming messages for debugging
		log.Printf("Message from chatID %d: %s", chatID, userText)

		if chatID != allowedChatID {
			log.Println("Unauthorized user tried to send:", userText)
			continue
		}

		// Extract command (first word) from the message
		command := strings.Split(userText, " ")[0]
		command = strings.ToLower(command)

		// Check if the command exists in our handlers map
		if cmdHandler, exists := commandHandlers[command]; exists {
			cmdHandler.handler(bot, chatID) // Call the corresponding handler function
		} else {
			// Send help message if command doesn't exist
			sendHelpMessage(bot, chatID)
		}
	}
}

// handleShutdown schedules the shutdown command
func handleShutdown(bot *tgbotapi.BotAPI, chatID int64) {
	log.Println("Shutdown command received! Shutting down...")

	// Send confirmation message
	msg := tgbotapi.NewMessage(chatID, "Shutdown has been scheduled. You can abort in 60 seconds. 💀")
	bot.Send(msg)

	// Schedule shutdown in 60 seconds
	shutdownTimer = time.AfterFunc(60*time.Second, func() {
		cmd := exec.Command("shutdown", "/s", "/t", "0")
		err := cmd.Run()
		if err != nil {
			log.Println("Failed to shutdown:", err)
			errorMsg := tgbotapi.NewMessage(chatID, "Failed to shutdown: "+err.Error())
			bot.Send(errorMsg)
		}
	})
}

// handleAbort cancels the scheduled shutdown
func handleAbort(bot *tgbotapi.BotAPI, chatID int64) {
	if shutdownTimer != nil && shutdownTimer.Stop() {
		log.Println("Shutdown aborted!")

		// Send confirmation message
		msg := tgbotapi.NewMessage(chatID, "Shutdown aborted. Your PC is safe! 😅")
		bot.Send(msg)
	} else {
		log.Println("No shutdown scheduled to abort.")
		msg := tgbotapi.NewMessage(chatID, "No shutdown scheduled, nothing to abort.")
		bot.Send(msg)
	}
}

// focusWindow focuses the game window
func focusWindow(bot *tgbotapi.BotAPI, chatID int64, windowTitle string) {
	// Find the window by its title
	hwnd, err := FindWindow(windowTitle)
	if err != nil {
		log.Printf("Failed to find the window: %v", err)
		msg := tgbotapi.NewMessage(chatID, "Failed to find the game window.")
		bot.Send(msg)
		return
	}

	// Attempt to restore and focus the window
	err = restoreAndFocusWindow(hwnd)
	if err != nil {
		log.Printf("Failed to focus the window: %v", err)
		msg := tgbotapi.NewMessage(chatID, "Failed to focus the game window.")
		bot.Send(msg)
		return
	}

	// Send confirmation message
	log.Println("Game window focused successfully!")
	msg := tgbotapi.NewMessage(chatID, "The game window has been focused and unminimized.")
	bot.Send(msg)
}

// handleHelp displays all available commands
func handleHelp(bot *tgbotapi.BotAPI, chatID int64) {
	sendHelpMessage(bot, chatID)
}

// sendHelpMessage creates and sends the help message
func sendHelpMessage(bot *tgbotapi.BotAPI, chatID int64) {
	var helpText strings.Builder
	helpText.WriteString("Available commands:\n\n")

	for cmd, handler := range commandHandlers {
		helpText.WriteString(cmd + " - " + handler.description + "\n")
	}

	msg := tgbotapi.NewMessage(chatID, helpText.String())
	bot.Send(msg)
}
