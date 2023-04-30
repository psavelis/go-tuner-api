# build
FROM --platform=linux/amd64 golang:1.19 AS build
WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 GOOS=linux go build -o api ./cmd/main.go

# runtime
FROM --platform=linux/amd64 scratch
FROM scratch 
WORKDIR /app
COPY --from=build /app/api ./
EXPOSE 3000
CMD [ "./api" ]