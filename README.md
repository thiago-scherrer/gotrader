# GoTrader bot

Project to create a trade bot for bitmex trade platform.

[![Go Report Card](https://goreportcard.com/badge/github.com/thiago-scherrer/gotrader)](https://goreportcard.com/report/github.com/thiago-scherrer/gotrader) [![Build Status](https://travis-ci.org/thiago-scherrer/gotrader.svg?branch=master)](https://travis-ci.org/thiago-scherrer/gotrader)
[![GoDoc](https://godoc.org/github.com/thiago-scherrer/gotrader?status.svg)](https://godoc.org/github.com/thiago-scherrer/gotrader)

![gopher](assets/gopher.png)

## Requirements

- docker => 18.09.2
- docker-compose => 1.23.2
- bitmex account

## How it works

Its purpose is to automate a rule created by the trader.

## Caution

This bot still under construction and does not guarantee anything, it may not even work properly. You can lose money with it! **Test the gotrader bot in the test network first!**
But if you have good results, PR the logic =)

Don't panic!

## Enabled

- New and close order, buy or sell
- leverage
- Send messages to matrix.org
- Able to use custom logic

## Supported Contracts

- XBTUSD
- ETHUSD

## Runing with docker

Go to the config folder and then, copy **config-example.yml** to the **config.yml**. Add your settings to the file and save.

Go to the **example/** folder and then, choose or create a strategy and modify with your trade strategies and after this, copy the file like **logic-short-trade_go** to the folder **internal/logic** with the name **logic.go**. Go back to the root dir and follow the steps below:

```bash
docker-compose build
```

After *test* and *build*, run the bot (background):

```bash
docker-compose up -d
```

You can see the logs with *docker logs* command, like:

```bash
docker logs -f gotrader
```

To stop the bot, run:

```bash
docker-compose down
```

## Logic

The acual logic can be changed on *internal/logic/*. A simple example can be found on *examples/*.

## TO-DO

- best log control
- more documentation
- more logic
- more redis 

## References

- [bitmex api](https://www.bitmex.com/api/explorer/)
- [docker-compose install](https://docs.docker.com/compose/install/)
- [matrix doc](https://matrix.org/docs/spec/client_server/latest#sending-events-to-a-room)
- [goreportcard](https://goreportcard.com/)
- [gopherize](https://gopherize.me)
- [go-yaml](https://github.com/go-yaml/yaml)
- [project-layout](https://github.com/golang-standards/project-layout)
