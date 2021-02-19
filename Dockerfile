FROM golang:latest AS builder
ADD . /app
WORKDIR /app
#RUN go get -d -v
#RUN go mod download
#RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /app/main .
#RUN go mod download
RUN go build -o /main .

FROM scratch
#ENV MONGODB_USERNAME=MONGODB_USERNAME MONGODB_PASSWORD=MONGODB_PASSWORD MONGODB_ENDPOINT=MONGODB_ENDPOINT
COPY --from=builder /main ./
ENTRYPOINT ["./main"]
EXPOSE 8099