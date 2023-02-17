run:
	go run ./cmd/main.go

build:
	docker image build -t forum .

d-run:
	docker run -d -p 8080:8080 --name forumapp forum 