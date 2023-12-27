Сборка прокета:
1. ``docker compose up``
2. дождаться запуска сервиса
3. ``docker exec news_server goose -dir migrations postgres "user=root password=rpass dbname=news_db host=news_pg port=5432 sslmode=disable" up``

Сервис поддерживает запросы:
- POST /edit/:Id - изменение новости по Id
- GET /list - список новостей; Пагинация через Query Param /list?list=1&limit=100