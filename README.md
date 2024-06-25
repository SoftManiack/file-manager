# File Manager

__Тесты__

_https://ieftimov.com/posts/testing-in-go-test-doubles-by-example/_
_https://habr.com/ru/articles/649363/_
_https://github.com/golang/mock_
___

__Миграции__

https://github.com/pressly/goose?ysclid=lr58j7plma383897397

"postgres://YourUserName:YourPassword@YourHostname:5432/YourDatabaseName";
goose -dir db/migrations postgres "postgresql://goose:password@127.0.0.1:8092/go_migrations?sslmode=disable" up

goose -dir db/migrations postgres "postgres://postgres:123@localhost:5432/fm" up