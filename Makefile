run:
	go run ./cmd/hub/main.go --env
startdbs:
	sudo docker-compose up -d
stopdbs:
	sudo docker-compose stop
enterpostgres:
	sudo docker exec -it postgres12 psql -U root
 
cert:
	cd cert; bash gen.sh; cd ..
genmock:
	mockgen -source=internal/service/services.go -destination internal/service/mocks/mock.go
.PHONY: gen clean cert
