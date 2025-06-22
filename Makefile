.PHONY: db

run: build
	@./bin/app

build:
	@templ generate
	@go build -o bin/app .

css:
	@tailwindcss -i views/css/app.css -o public/tailwind.css --watch

templ: 
	@templ generate --watch --proxy=http://localhost:8080

air:
	@air

db:
	docker compose up db

dev:
	make -j3 css templ air
