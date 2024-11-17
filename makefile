
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/tangapp_beta?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/tangapp_beta?sslmode=disable" -verbose down


.PHONY: migrateup migratedown