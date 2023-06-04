# WebRTC сервер для обмена потоком между клиентами

## О проекте

Учебный проект, был создан с желанием познать WebRTC и понять как устроны стриминговые сервисы и ресурсы вроде Yandex.Telemost, meet.google и др.

## Как работает

- Клиент может создать комнату с помощью `/create`, который, в свою очередь, вернёт ему уникальный хэш этой комнаты
- Сразу после создания он может присоединиться по WebSocket к комнате посредством `/join?roomID={hash_room_id}`, с этим же методом могут присоединиться и другие клиенты

Сервер абсолютно ни к чему не привязан и может обмениваться не только видеопотоком, но и сообщениями, например, для чата.

## Как запустить

Запускаем посредством `go run ./cmd/server/main.go` или делаем `go build` того же файла и запускаем бинарник

## Чтобы я доработал

У сервера есть парочка неисправленных ошибок, которые я на досуге, возможно, поправлю. Но у проекта была цель просто поиграться с WebRTC.

Точки роста:

- подключить хэш и бд, чтобы хранить сессии долговечно и не терять комнаты
- нет нормального отключения от комнаты в данной реализации
- нет контекста, нет контроля состояния сервера

## Автор данной идеи

Оставляю ссылку на источник, который помог мне разобраться в данной концепции: https://www.youtube.com/watch?v=JTIm3ChI-6w