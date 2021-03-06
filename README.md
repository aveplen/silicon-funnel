# Imap Concentrator Service

```
  ┌────────┐                                                 ┌──────────────┐
 ┌┴───────┐│    ┌──────────────┐                            ┌┴─────────────┐│
┌┴───────┐││    │     imap     │  chat  ┌──────────────┐   ┌┴─────────────┐├┘
│        │││<───│ concentrator │<─ id ──│    gateway   │<──│    client    ├┘
│  IMAP  │├┘    │    service   │  auth  └──────────────┘   └──────────────┘
│        ├┘     └──┬────────┬──┘
└────────┘         ^        ^ white
                   │        │ list
            ╭──────┴─╮      │ auth
            ╰────────╯  ┌───┴──────────┐
            │   db   │  │ telegram bot │
            ╰────────╯  └──────────────┘

```

Сервис создан для отслеживания новых сообщений на imap-ящиках. Основной сервер
предоставляет grpc-api с возможностью подписки для получения уведомлений по
принципу server-side-streaming'а, также есть возможность подписаться на обновления
по websocket'у, который реализован через grpc-gateway. 

Клиенты могут подписываться на оповещения о всех ящиках, для этого они должны
предоставить ключ аутентификации из whitelist'а, или на сообщения, которые бы
отправлялись в тг-чат, для этого они должны предоставить идентификатор чата.
Идея в том, чтобы использовать бота для "регистрации", и websocket для
оповещений.

В БД хранятся идентификаторы чатов, которые инициировали диалог с ботом, а 
также imap-ящики, которые были добавлены через эти чаты, вместе с оффсетами.
Сервис узнает о появлении новых сообщений при помощи периодического поллинга,
решение о том, является ли сообщение новым принимается на основании оффсета из
БД (с хранением оффсетов связана проблема отсутствия оповещений о новых
сообщениях после удаления старых с ящика, пока количество не дойдет до
хранимого оффсета).

## TODO:
1. Отрефакторить `internal/bot/telegram`
2. Форматирование оповещений
3. Возможность добавлять базовые ссылки на веб-интерфейс, чтобы получать 
ссылки на письма в оповещениях
4. Добавить возможность ручного рефетчинга оффсетов
5. Докер
6. Миграции
7. Тесты)

## Запускать:
```
docker compose -f deployment/docker-compose.yaml up
# тут надо руками создать таблицы (код в ./migration/)

go run cmd/server/server.go -c config/imapc.yaml
go run cmd/bot/bot.go -c config/bot.yaml
```