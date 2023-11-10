build:
	env GOOS=linux CGO_ENABLED=0  GOARCH=amd64 -ldflags="-s -w" go build -o bin/main main.go
deploy: build
	serverless deploy --stage prod --aws-profile default

	