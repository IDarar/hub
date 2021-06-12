run:
	go run ./cmd/hub/main.go --env
container-run:
	sudo docker run --network hub_main_app --env-file .env  hub
startcompose:
	sudo docker-compose up -d
stopcompose:
	sudo docker-compose down
enterpostgres:
	sudo docker exec -it postgres12 psql -U root
dockerbuild:
	sudo docker build -t hub .
tag:
	sudo docker tag hub:latest aince/hub
push:
	sudo docker push aince/hub	
cert:
	cd cert; bash gen.sh; cd ..
swag:
	swag init -g internal/app/app.go
genmock:
	mockgen -source=internal/service/services.go -destination internal/service/mocks/mock.go
.PHONY: gen clean cert

