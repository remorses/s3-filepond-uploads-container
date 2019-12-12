

start: server client

.PHONY: server
server:
	docker-compose up --build

.PHONY: client
client:
	cd client && yarn && yarn start
