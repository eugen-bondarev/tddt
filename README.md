# tddt - tiny database dump tool

## Motivation

I wanted a simple Docker-first tool that works anywhere, follows cloud native principles and is easy to deploy.

`tddt` gives you **one** endpoint to create a dump of your database (mysql, postgres, more in the future) and upload it to the cloud storage of your choice (for now only GCP, more in the future).

## Deployment

This section describes how to start the service in Docker. You can derive other deployment methods (e.g. kubernetes) from this.

1. Create a `.env` file based on the `.env.example` file.

```bash
cp .env.example .env
```

2. Run the container

```bash
docker run -p 8080:8080 --env-file .env eugbondarev/tddt:1.0.0
```

## Usage

Create a dump

```bash
curl -X POST http://admin:password@localhost:8080/v1/dump \
    -H "Content-Type: application/json" \
    -d '{
        "dump": {
            "database": "test",
            "type": "pg"
        },
        "output": {
            "bucket": "my-nice-gcp-bucket-123",
            "path": "pg/1.sql"
        }
    }'
```

## Requirements

- mysqldump
- pg_dump

## Roadmap

- [ ] Refactoring for better code structure
- [ ] Swagger documentation
- [ ] Encryption
- [ ] Upload to S3, Azure blob, SFTP and more
- [ ] Support more databases
- [ ] Better logging
- [ ] More flexibility through env variables
- [ ] Add tests
- [ ] Many Docker images, only with required stuff (e.g. no mysql if you only need postgres)

> [!NOTE]  
> It's only v1, the API might change in the future
