from golang:latest

COPY . .

# build
RUN go build main.go
err
EXPOSE 8080
ENTRYPOINT [ "./main" ]