FROM golang:1.21 as builder
ARG CGO_ENABLED=0
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN mkdir dist
RUN go build -o dist/main
RUN cp -r html/ dist/html/ && cp -r static/ dist/static/

FROM scratch
COPY --from=builder /app/dist/ /app/
WORKDIR /app
ENTRYPOINT ["./main"]
