FROM golang

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ENV ENV=docker

EXPOSE 5000
EXPOSE 5432

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /server

ENTRYPOINT [ "/server" ]