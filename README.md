# Тестовое задание Ozon Fintech. Сервис для сокращения ссылок
## Установка и запуск
### Для запуска с выбором типа хранилища
#### Установка
`docker-compose build `
#### Запуск
`docker-compose run -p 8000:8000 urlshortener go run main.go -type repo ` - с БД в качестве хранилища

или

`docker-compose run -p 8000:8000 urlshortener go run main.go -type mem ` - с in-memory хранилищем
### Для быстрого запуска только с БД
`docker-compose up `

## Скриншоты api-тестов
### Успешное создание короткой ссылки
![valid_post_long](https://github.com/sashabondar41/url_shortener/assets/75033340/915629d7-09ca-4b76-9634-2929611fd65e)
### Некорректный URL при создании короткой ссылки
![invalid_url_post_long](https://github.com/sashabondar41/url_shortener/assets/75033340/85bee76c-5705-4016-a6ef-29877c4cd39d)
### Успешное получение оригинальной ссылки по короткой
![valid_get_long](https://github.com/sashabondar41/url_shortener/assets/75033340/f361bb41-8eeb-4f58-b956-e004a6a74385)
### Отстутсвующая в хранилище короткая ссылка при поиске оригинальной
![not_found_get_long](https://github.com/sashabondar41/url_shortener/assets/75033340/d967144b-cc75-4af3-a9f2-0969a9ed75f1)
### Некорректный URL при поиске оригинальной ссылки
![invalid_url_get_long](https://github.com/sashabondar41/url_shortener/assets/75033340/c019fbb2-8a6f-4fce-b5c9-e6c350923656)
