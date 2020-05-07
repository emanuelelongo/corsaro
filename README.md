# Corsaro 🏴‍☠️
A simple reverse proxy to allow cors request.

Useful for local development or testing. Don't use it in production.


## Run with Docker

To proxy a remote service, simply run:
``` bash
docker run -p 1337:1337 emanuelelongo/corsaro http://a-service-not-allowing-cors.com
```
Now the service is reachable to http://localhost:1337 allowing cors.


If the service you need to proxy runs on `localhost`, for example on port `8080`, then use `host.docker.internal` to refer it inside Docker:

``` bash
docker run -p 1337:1337 emanuelelongo/corsaro http://host.docker.internal:8080
```

## Build your own executalbe

If you need an executable that can be run without Docker or any dependencies you can build it within a Go environment.

``` bash
git clone https://github.com/emanuelelongo/corsaro.git
cd corsaro
go mod download
go build -o ./build/corsaro ./src
```

## Change default port
By default `Corsaro` listen on 1337 port but you can override this setting using an additional argument:

``` bash
corsaro http://a-service-not-allowing-cors.com 8080
```

If running with docker remember to map the port accordingly:
``` bash
docker run -p 8080:8080 emanuelelongo/corsaro http://a-service-not-allowing-cors.com 8080
```