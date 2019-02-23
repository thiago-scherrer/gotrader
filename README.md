# GoTrader bot

Project to create a trade bot for a bitmex

[![Go Report Card](https://goreportcard.com/badge/github.com/thiago-scherrer/gotrader)](https://goreportcard.com/report/github.com/thiago-scherrer/gotrader)

![gopher](assets/gopher.png)

## requirements

- bitmex api
- gopkg.in/yaml.v2 (go get gopkg.in/yaml.v2)

## how it works

This robot is still under construction. Its purpose is to automate a rule created by the trader. It's not a money machine ...

## test

Copy the sample configuration file, which is inside configs. Change the required data.
Enter to the ./internal/gotrader folder and run the test:

```bash
go test -args config ../../config.yml
```

## build

```bash
go build -o gotrader 
```

## runing

```bash
./gotrader config ../../config.yml
```


## references

- [golang-standards](https://github.com/golang-standards/project-layout)
- [bitmex api](https://www.bitmex.com/api/explorer/)
- [goreportcard](https://goreportcard.com/)
- [gopherize](https://gopherize.me)
