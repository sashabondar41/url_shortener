# Сервис для сокращения ссылок
## Тестовое задание Ozon Fintech
### Для запуска с выбором типа хранилища
####Установка
`docker-compose build `
#### Запуск
`docker-compose run go-app go run main.go -type repo ` - с БД в качестве хранилища

или

`docker-compose run go-app go run main.go -type mem ` - с in-memory хранилищем
### Для быстрого запуска только с БД
`docker-compose up `
