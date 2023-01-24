run-db:
	sudo docker run --name=todo-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres
create-migrate:
	migrate create -ext sql -dir ./schema -seq init
up-migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' up
down-migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/postgres?sslmode=disable' down
open-db:
	sudo docker exec -it todo-db bash
run-server:
	go run cmd/main.go