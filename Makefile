build:
	@go run . build -file main.lumi -out out.lbc

run: build
	@go run . run out.lbc
