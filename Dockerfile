FROM golang:1.14-alpine as BUILDER
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build .

FROM alpine:latest
COPY --from=BUILDER /app/graphql-gin-websocket-example .
COPY --from=BUILDER /app/server/config.yaml .
RUN chmod +x graphql-gin-websocket-example
RUN chmod +r config.yaml
ENTRYPOINT ["./graphql-gin-websocket-example","-c", "config.yaml"]

