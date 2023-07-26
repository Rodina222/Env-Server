FROM golang:1.20 AS builder

WORKDIR /go/src/github.com/codescalersinternships/EnvServer-Rodina/

COPY . ./

RUN CGO_ENABLED=0  go build -o app cmd/main.go


FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/github.com/codescalersinternships/EnvServer-Rodina/app ./

EXPOSE 8080

CMD [ "./app","-p","8080" ]
