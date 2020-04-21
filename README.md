# GoTrader bot

Project to create a trade bot for bitmex trade platform.

[![Go Report Card](https://goreportcard.com/badge/github.com/thiago-scherrer/gotrader)](https://goreportcard.com/report/github.com/thiago-scherrer/gotrader) [![Build Status](https://travis-ci.org/thiago-scherrer/gotrader.svg?branch=master)](https://travis-ci.org/thiago-scherrer/gotrader)
[![GoDoc](https://godoc.org/github.com/thiago-scherrer/gotrader?status.svg)](https://godoc.org/github.com/thiago-scherrer/gotrader)

![gopher](assets/gopher.png)

## Requirements

- docker => 19.03.8
- docker-compose => 1.25.4
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
- Able to use custom logic

## Supported Contracts

- XBTUSD
- ETHUSD

## Runing with docker

Add your settings to the file **configs/config.yml**

Go back to the root dir and follow the steps below:

```sh
docker-compose build
```

After *test* and *build*, run the bot (background):

```sh
docker-compose up -d
```

You can see the logs with *docker logs* command, like:

```sh
docker logs -f gotrader
```

To stop the bot, run:

```sh
docker-compose down
```

## Using other logic

The acual logic can be changed on *internal/logic/*. Some examples can be found on *examples/*.

Go to the **example/** folder and then, choose or create a strategy. Copy the file like **martingale_go** to the **internal/logic/logic.go**.

## References

- [bitmex api](https://www.bitmex.com/api/explorer/)
- [docker-compose install](https://docs.docker.com/compose/install/)
- [gopherize](https://gopherize.me)
- [project-layout](https://github.com/golang-standards/project-layout)
