FROM golang:1.14.2 AS build
WORKDIR /app
COPY . .
RUN go mod download && CGO_ENABLED=0 go build -o ./build/corsaro ./src

FROM scratch
EXPOSE 1337
WORKDIR /
COPY --from=build /app/build/corsaro .
ENTRYPOINT [ "./corsaro"]