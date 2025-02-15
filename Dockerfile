FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /app-exec ./main.go

RUN apt-get update
RUN apt-get install -y lsb-release

# Prepare MySQL 8
RUN wget https://dev.mysql.com/get/mysql-apt-config_0.8.33-1_all.deb
RUN DEBIAN_FRONTEND=noninteractive dpkg -i mysql-apt-config_0.8.33-1_all.deb

# Prepare PostgreSQL 17
RUN sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
RUN curl -fsSL https://www.postgresql.org/media/keys/ACCC4CF8.asc | gpg --dearmor -o /etc/apt/trusted.gpg.d/postgresql.gpg

RUN apt-get update

# Install PostgreSQL
RUN apt-get install -y postgresql-client-17

# Install MySQL
RUN apt-get install -y mysql-client

CMD ["/app-exec"]

EXPOSE 8080