# syntax=docker/dockerfile:1

FROM golang:1.16-alpine
EXPOSE 8080

WORKDIR /app

COPY . ./

RUN go get
RUN go build -o /camera-emulator

CMD [ "/camera-emulator" ]