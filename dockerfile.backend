FROM golang:1.21.4

WORKDIR /app

COPY MongoAPI .
RUN go mod download && go mod verify


RUN go build -v -o bin .

ENTRYPOINT [ "/app/bin" ]