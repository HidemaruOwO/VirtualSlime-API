FROM golang:1.21-alpine

WORKDIR /app

RUN apk update && apk add git bash
RUN git clone https://github.com/HidemaruOwO/VirtualSlime-API.git --depth 1 /app
RUN go build -ldflags="-s -w" -trimpath src/main.go
RUN mv main virtualslime

EXPOSE 3000

CMD ["./virtualslime"]
