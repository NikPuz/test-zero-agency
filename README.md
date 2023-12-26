Сборка прокета:
1. ``docker compose up``
2. дождаться запуска сервиса
3. ``goose -dir migrations postgres "user=root password=rpass dbname=ps_db host=localhost port=5432 sslmode=disable" up``