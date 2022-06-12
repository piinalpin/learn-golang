# Run app
FROM alpine:3.16.0

WORKDIR /app
ARG APP_NAME

COPY bin/${APP_NAME} goapp

EXPOSE 8080
ENTRYPOINT ./goapp

