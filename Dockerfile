# build linux:
# CGO_ENABLE=0 GOOS=linux go build -a -installsuffix cgo -o app .

# FROM alpine

# WORKDIR /app1
# # ADD ./app .

# COPY ./app .

# ENTRYPOINT [ "./app" ]


# # docker build -t myapp:1.0 .
# # docker run -d --name myapp1 -p 3005:3005 myapp:1.0



FROM golang:1.22.6-alpine
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /out/main ./
ENTRYPOINT ["/out/main"]