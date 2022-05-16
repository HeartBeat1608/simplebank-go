POSTGRES := "postgres14"
DATABASE := "simple_bank"
POSTGRES_USER := root
POSTGRES_PASSWORD := secret

postgres_new:
	docker run --name $(POSTGRES) -e POSTGRES_USER=$(POSTGRES_USER)-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -p 5432:5432 postgres:14-alpine 

postgres:
	docker start $(POSTGRES)

createdb: 
	docker exec -it $(POSTGRES) createdb --username=$(POSTGRES_USER) --owner=$(POSTGRES_USER) $(DATABASE)

dropdb:
	docker exec -it $(POSTGRES) dropdb $(DATABASE)
	
cleanup:
	docker stop $(POSTGRES)

migrateup:
	migrate -path db/migration -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(DATABASE)?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(DATABASE)?sslmode=disable" -verbose down

sqlc:
	../sqlc generate

test:
	go test -v -cover ./...

clean_setup: postgres dropdb createdb migrateup sqlc cleanup