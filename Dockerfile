FROM golang:1.12

RUN apt-get update

# install and configure app
WORKDIR /go/src/github.com/213-team/recbot_backend
COPY . .
RUN go get ./...
WORKDIR /go/bin
RUN go build -o backend /go/src/github.com/213-team/recbot_backend/apps/backend/main.go

# install and configure supervisor
RUN apt-get install -y supervisor
ADD ./supervisord.conf /etc/supervisord.conf

CMD ["supervisord", "-n"]