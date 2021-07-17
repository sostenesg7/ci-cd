from golang:latest

COPY . .

# build
RUN go build main.go

EXPOSE 8080
ENTRYPOINT [ "./main" ]
