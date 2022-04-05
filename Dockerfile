FROM golang:1.18.0-alpine3.14 as builder

WORKDIR /go/app

ADD go.mod go.sum ./
RUN go mod download
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main source/cmd/main.go

# copy artifacts to a clean image
FROM alpine
COPY --from=builder /go/app/main .
ADD source/configuration/enviroments source/configuration/enviroments 
ENTRYPOINT [ "./main" ]