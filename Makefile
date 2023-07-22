swagger-generate:
	swagger generate spec -o ./swagger-ui/swagger.json --scan-models

swagger-validate:swagger-generate
	swagger validate ./swagger-ui/swagger.json

run:
	go run main.go


