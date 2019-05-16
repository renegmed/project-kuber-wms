.PHONY: build
build:
	go build -o inven-wms .

.PHONY: docker
docker:
	docker build -t inven-wms .

.PHONY: run  
run:
	docker run --name inven-wms --net host -d -p 8484:8484 inven-wms

.PHONY: tag
tag:
	docker tag inven-wms:latest renegmedal/inven-wms:1.0.1

.PHONY: push
push:
	docker push renegmedal/inven-wms:1.0.1

.PHONY: up
up:
	docker-compose up --build -d 
