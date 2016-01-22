## http service

## Usage

1. Start service

	```
	go run main.go
	```
2. Add keywords to service

	```
	curl -X POST -H "Content-Type: application/json" -d '[{"Word":"xxoo","Kind":"porn","Rate":100}]' http://127.0.0.1:7520/addkws
	```
3. Add ignore symbols to service

	```
	curl -X POST -H "Content-Type: application/json" -d '["*","%"]' http://127.0.0.1:7520/addsbs
	```
4. Fitler content

	```
	curl -X POST -H "Content-Type: application/json" -d '{"Content":"test *xx**oo something"}' http://127.0.0.1:7520/filter
	// result
	{"Success":true,"Msg":"","Result":{"Threshold":100,"Rate":100,"Keywords":[{"Rate":100,"Index":6,"Kind":"porn","Word":"xx**oo"}]}}
	```
