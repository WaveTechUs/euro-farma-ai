# Project farmaIA

## Getting Started

## Install
```
scoop install migrate
```
```
go install github.com/swaggo/swag/cmd/swag@latest
```

## Docker
```
docker compose up
```

## MakeFile
run the application
```bash
make run
```

run the migrations
```bash
make migrate
```

create migrations
```bash
migrate create -ext=sql -dir=cmd/migrate/ -seq init
```
