# GoTrader bot

Project to create a trade bot for a bitmex

[![Go Report Card](https://goreportcard.com/badge/github.com/thiago-scherrer/gotrader)](https://goreportcard.com/report/github.com/thiago-scherrer/gotrader)

![gopher](assets/gopher.png)

## requirements

- docker => 18.09.2
- docker-compose => 1.23.2
- bitmex api
- gopkg.in/yaml.v2 (go get gopkg.in/yaml.v2)

## how it works

This robot is still under construction. Its purpose is to automate a rule created by the trader. It's not a money machine ...

## runing

Go to the config folder and then, copy **config-example.yml** to the **config.yml**. Add your settings to the file and then, go back to the root dir and run *docker-compose*:

```bash
docker-compose build
```

After *build*, run the docker:

```bash
docker-compose up -d
```

You can see the logs with *docker logs* command, like:

```bash
docker logs -f gotrader_gotrader_1
```

## logic

The acual logic can be changed on *internal/gotrader/logic.go*.

## test

Copy the sample configuration file, which is inside configs. Change the required data.
Enter to the *./internal/gotrader* folder and run the test:

```bash
go test -args config ../../configs/config-test.yml
```

## manual run

Build the bin:

```bash
cd internal/gotrader/
go build -o gotrader 
```

And then, run:

```bash
./gotrader config ../../configs/config.yml
```

## references

- [golang-standards](https://github.com/golang-standards/project-layout)
- [bitmex api](https://www.bitmex.com/api/explorer/)
- [goreportcard](https://goreportcard.com/)
- [gopherize](https://gopherize.me)
- [docker-compose install](https://docs.docker.com/compose/install/)

## irc

- freenode - Channel #gotrader