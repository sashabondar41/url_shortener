# Тестовое задание Ozon Fintech. Сервис для сокращения ссылок
## Установка и запуск
### Для запуска с выбором типа хранилища
`docker-compose run -p 8000:8000 urlshortener go run main.go -type repo ` - с БД в качестве хранилища

#### или

`docker-compose run -p 8000:8000 urlshortener go run main.go -type mem ` - с in-memory хранилищем
### Для быстрого запуска только с БД
`docker-compose up `

## Скриншоты api-тестов
### Успешное создание короткой ссылки
![valid_post_long](https://github.com/sashabondar41/url_shortener/assets/75033340/c9c3b641-31d8-487f-a3aa-19b86faf8ded)
### Некорректный URL при создании короткой ссылки
![invalid_url_post_long](https://github.com/sashabondar41/url_shortener/assets/75033340/85ed95ce-c867-432d-a5ef-13076091445e)
### Успешное получение оригинальной ссылки по короткой
![valid_get_long](https://github.com/sashabondar41/url_shortener/assets/75033340/6ea5d63d-6253-42f5-ae02-8b9e61db66cf)
### Отстутсвующая в хранилище короткая ссылка при поиске оригинальной
![not_found_get_long](https://github.com/sashabondar41/url_shortener/assets/75033340/07748357-bd28-4e23-b657-e5fcb95eb24b)
### Некорректный URL при поиске оригинальной ссылки
![invalid_url_get_long](https://github.com/sashabondar41/url_shortener/assets/75033340/85e462b7-46ee-4966-8dc9-9d9aae9e183c)






