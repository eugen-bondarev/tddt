# tddt - tiny database dump tool

## Motivation

I wanted a simple Docker-first tool that works anywhere, follows cloud native principles and is easy to deploy.

`tddt` gives you **one** endpoint to create a dump of your database (mysql, postgres, more in the future) and upload it to the cloud storage of your choice (for now only GCP, more in the future).

## What `tddt` is not

- `tddt` is not a CLI tool
- `tddt` doesn't take care of any scheduling. It's up to you to schedule backups and call the endpoint of `tddt`, e.g. using cron

## TL;DR

1. Call the endpoint with a JSON body to create a dump

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

2. Enjoy your dump

![image](./docs/images/1.png)

## Deployment

This section describes how to start the service in Docker. You can derive other deployment methods (e.g. kubernetes) from this.

1. Create a `.env` file based on the `.env.example` file.

```bash
cp .env.example .env
```

2. Fill in the variables. Make sure your `GOOGLE_CLOUD_CREDENTIALS` is set properly. It should be a Base64 encoded string of your JSON credentials.

3. Run the container

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

### Other request examples

```json
{
  "dump": {
    "database": "test",
    "type": "mysql"
  },
  "output": {
    "bucket": "my-nice-gcp-bucket-123",
    "path": "1.sql"
  },
  "async": true // optional, default is false. This will immediately return and create a dump in the background
}
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
- [ ] Many Docker images, only with required stuff (e.g. no mysql if you only need postgres, etc.)

> [!NOTE]  
> It's only v1, the API might change in the future
