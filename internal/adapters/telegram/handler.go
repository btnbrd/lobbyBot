// telegram/handler.go
package telegram

import (
	"strings"
	"tgBot/internal/adapters/telegram/commands"
	"tgBot/internal/app"
	"tgBot/internal/domain/entities"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	bot        *tgbotapi.BotAPI
	service    *app.LobbyService
	userStates map[string]UserState
}

func NewHandler(bot *tgbotapi.BotAPI, service *app.LobbyService) *Handler {
	return &Handler{
		bot:        bot,
		service:    service,
		userStates: make(map[string]UserState),
	}
}

func (h *Handler) Handle(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	user := entities.NewUser(update.Message.From.UserName).WithChatId(update.Message.Chat.ID)
	text := update.Message.Text
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	if update.Message.ForwardFrom != nil {
		forwardedUser := entities.NewUser(update.Message.ForwardFrom.UserName)
		commands.Connect(h.service, user, forwardedUser.Name(), &msg)
	} else {
		if update.Message.IsCommand() {
			h.clearState(user.Name())
			h.handleCommand(update, &msg, user)
		} else {
			switch h.getState(user.Name()) {
			case StateJoinWait:
				h.clearState(user.Name())

				code := strings.ToUpper(strings.TrimSpace(text))
				commands.Join(h.service, user, code, &msg)

			case StateAddWait:
				h.clearState(user.Name())

				subName := strings.ToLower(strings.TrimSpace(text))
				commands.Connect(h.service, user, subName, &msg)

			case StateShuffleWait:
				h.clearState(user.Name())

				commands.Shuffle(h.service, user, text, &msg, h.bot)

			default:
				msg.Text = "Я не понимаю это сообщение. Попробуй /help"

			}
		}
	}
	h.bot.Send(msg)
}

func (h *Handler) handleCommand(update tgbotapi.Update, msg *tgbotapi.MessageConfig, user *entities.User) {
	command := update.Message.Command()
	args := update.Message.CommandArguments()

	switch command {
	case "help":
		commands.Help(user, msg)
	case "start":
		commands.Start(h.service, user)
	case "create":
		commands.Create(h.service, user, msg)
	case "join":
		if args == "" {
			msg.Text = "Введите код лобби:"
			h.setState(user.Name(), StateJoinWait)
		} else {
			commands.Join(h.service, user, args, msg)
		}
	case "random":
		commands.Random(h.service, user, msg)
	case "leave":
		commands.Leave(h.service, user, msg)
	case "whoami":
		commands.Whoami(h.service, user, msg)
	case "shuffle":
		if args == "" {
			msg.Text = "Введите количество перестановок:"
			h.setState(user.Name(), StateShuffleWait)
		} else {
			commands.Shuffle(h.service, user, args, msg, h.bot)
		}
	case "add":
		if args == "" {
			msg.Text = "Введите ник пользователя для добавления:"
			h.setState(user.Name(), StateAddWait)
		} else {
			commands.Connect(h.service, user, args, msg)
		}
	default:
		msg.Text = "Неизвестная команда"
	}
	msg.ReplyMarkup = h.getMainKeyboard()
}

func (h *Handler) getMainKeyboard() tgbotapi.ReplyKeyboardMarkup {
	buttons := [][]tgbotapi.KeyboardButton{
		{
			tgbotapi.NewKeyboardButton("/create"),
			tgbotapi.NewKeyboardButton("/join"),
		},
		{
			tgbotapi.NewKeyboardButton("/random"),
			tgbotapi.NewKeyboardButton("/leave"),
		},
		{
			tgbotapi.NewKeyboardButton("/whoami"),
			tgbotapi.NewKeyboardButton("/shuffle 5"),
		},
		{
			tgbotapi.NewKeyboardButton("/add"),
		},
	}
	return tgbotapi.NewReplyKeyboard(buttons...)
}
