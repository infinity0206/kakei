migrate-up:
	docker-compose exec api migrate -source file://migration/ -database 'mysql://admin:admin@tcp(mysql)/root' up

migrate-down:
	docker-compose exec api migrate -source file://migration/ -database 'mysql://admin:admin@tcp(mysql)/root' down -all