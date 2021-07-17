from golang:latest

COPY . .

RUN go build main.go


EXPOSE 8080
ENTRYPOINT [ "./main" ]