package telegram

const helpText = "/start - начало работы с ботом\n" +
	"/insert - добавление нового imap-ящика для отслеживания\n" +
	"/list - просмотр списка отслеживаемых ящиков\n" +
	"/poll - проверка наличия новых сообщений прямо сейчас\n" +
	"/help - вывод информационного сообщения\n" +
	"\n" +
	"Бот предназначен для отслеживания новых сообщений на imap-ящиках. Создан " +
	"в качестве проекта для портфолио с использованием языка программирования " +
	"Go, системы удаленного вызова процедур gRPC, базы данных PostgreSQL и др." +
	"Репозиторий проекта: https://github.com/aveplen/silicon-funnel"

func (t *TelegramService) SendHelpText(chatID int64) error {
	return t.SendText(chatID, helpText)
}
