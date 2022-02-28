FROM golang:latest
RUN mkdir src/app
WORKDIR /src/app
COPY . .
RUN go build cmd/app/main.go
CMD ./main
EXPOSE 9876
