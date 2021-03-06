package callbacks

import (
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/command"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/user_state"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func cbFilterReplyText() string {
	return "👇 Type your Feed Filter (it can be Regular Expression):"
}

func (m *Manager) cbFilter(c *command.Command) (string, *tgbotapi.InlineKeyboardMarkup) {
	c.UserState.State = user_state.FilterSet
	return cbFilterReplyText(), nil
}
