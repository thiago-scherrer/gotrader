# GoTrader bot

Project to create a trade bot for bitmex trade platform.

[![Go Report Card](https://goreportcard.com/badge/github.com/thiago-scherrer/gotrader)](https://goreportcard.com/report/github.com/thiago-scherrer/gotrader) [![Build Status](https://travis-ci.org/thiago-scherrer/gotrader.svg?branch=master)](https://travis-ci.org/thiago-scherrer/gotrader)

![gopher](assets/gopher.png)

## Requirements

- docker => 18.09.2
- docker-compose => 1.23.2
- bitmex account

## How it works

This robot still under construction. Its purpose is to automate a rule created by the trader. It's not a money machine ...

## Caution

This robo does not guarantee anything, it may not even work properly. You can lose money with it! Test the robo in the test network first.

You have been warned.

But if you have good results, share your extracts =)

Don't panic.

## Supported Contracts

- XBTUSD
- ETHUSD

## Runing

Go to the config folder and then, copy **config-example.yml** to the **config.yml**. Add your settings to the file and then, go back to the root dir and run *docker-compose*:

```bash
docker-compose build
```

After *test* and *build*, run the bot (background):

```bash
docker-compose up -d
```

You can see the logs with *docker logs* command, like:

```bash
docker logs -f gotrader_runner_1
```

To stop the robot, run:

```bash
docker-compose down
```

## Logic

The acual logic can be changed on *internal/logic/*.
Anothers logics or examples can be found on *configs/logics/*.

## References

- [golang-standards](https://github.com/golang-standards/project-layout)
- [bitmex api](https://www.bitmex.com/api/explorer/)
- [goreportcard](https://goreportcard.com/)
- [gopherize](https://gopherize.me)
- [docker-compose install](https://docs.docker.com/compose/install/)
- [go-yaml](https://github.com/go-yaml/yaml)

## IRC

- freenode - Channel #gotrader
