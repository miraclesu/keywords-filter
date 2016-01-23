# http service

# Run

1. install [MongoDB](http://docs.mongodb.org/manual/installation) and [Redis](http://redis.io/download#installation), MongoDB for storage keywords and symbols, Redis for keywords and symbols update notify.
2. config
	```
	cp load.json.sample load.json
	cp listen.json.sample listen.json
	```
3. go run main.go

## Usage

```
// add keywords
// echo 'PUBLISH kws \'{"Action":1,"Data":[{"Word":"xxoo","Kind":"porn", "Rate":100}]}\''|redis-cli
// add symbols
// echo 'PUBLISH sbs \'{"Action":1,"Data":["*"]}\''|redis-cli
curl -X POST -H "Content-Type: application/json" -d '{"Content":"test *xx**oo something"}' http://127.0.0.1:7520/filter
// result
{"Success":true,"Msg":"","Result":{"Threshold":100,"Rate":100,"Keywords":[{"Rate":100,"Index":6,"Kind":"porn","Word":"xx**oo"}]}}
```

## Note

you can start muitl serves and use nginx upstream to make the program stronger

