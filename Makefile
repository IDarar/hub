run:
	go run ./cmd/hub/main.go
startdbs:
	sudo docker-compose up -d
stopdbs:
	sudo docker-compose stop
enterpostgres:
	sudo docker exec -it postgres12 psql -U root
