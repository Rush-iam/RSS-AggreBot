package callbacks

import (
	"fmt"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/command"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/errors"
	"github.com/Rush-iam/RSS-AggreBot.git/internal/bot_ui/markup"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func cbSourceRemoveReplyText(sourceName, sourceUrl string) string {
	return fmt.Sprintf("Removed:\n🗑 %s\n%s", sourceName, sourceUrl)
}

func (m *Manager) cbSourceRemove(c *command.Command) (string, *tgbotapi.InlineKeyboardMarkup) {
	source, errReply := m.getSourceIdFromArg(c.UserId, c.Text)
	if errReply != "" {
		keyboard := markup.KeyboardBackToMenu()
		return errReply, &keyboard
	}
	err := m.backend.DeleteSource(source.Id)
	if err != nil {
		keyboard := markup.KeyboardBackToMenu()
		return errors.ErrInternalError, &keyboard
	}

	reply := cbSourceRemoveReplyText(source.Name, source.Url)
	keyboard := tgbotapi.NewInlineKeyboardMarkup(markup.ButtonBackToList())
	return reply, &keyboard
}
