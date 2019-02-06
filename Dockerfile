FROM golang:1.11.5 as builder

RUN go get gopkg.in/yaml.v2