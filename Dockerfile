FROM golang:alpine as builder
 
ADD . /go/src/
WORKDIR /go/src/
RUN go mod download
 
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /cmd/webserver .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
 
COPY --from=builder / ./
RUN mkdir ./configs
COPY ./configs/config.json ./configs
RUN mkdir ./resources 
COPY ./resources/ ./resources
 
EXPOSE 8080
 
ENTRYPOINT ["./"]