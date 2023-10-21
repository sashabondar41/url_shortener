FROM golang:latest
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /urlshortener
EXPOSE 8000
CMD ["/urlshortener"]