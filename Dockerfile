FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /app-exec ./main.go

RUN apt-get update && apt-get install -y lsb-release curl gnupg

# Prepare PostgreSQL 17 (official PGDG repo)
RUN sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" > /etc/apt/sources.list.d/pgdg.list'
RUN curl -fsSL https://www.postgresql.org/media/keys/ACCC4CF8.asc | gpg --dearmor -o /etc/apt/trusted.gpg.d/postgresql.gpg

RUN apt-get update

# Install PostgreSQL client
RUN apt-get install -y postgresql-client-17

# Install MySQL-compatible client from Debian (provides mysqldump)
RUN apt-get install -y default-mysql-client

RUN rm -rf /var/lib/apt/lists/*

CMD ["/app-exec"]

EXPOSE 8080
