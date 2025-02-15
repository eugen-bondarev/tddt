# tdbt - tiny database dump tool

## Motivation

I wanted a simple Docker-first tool that works anywhere, follows cloud native principles and is easy to deploy.

`tdbt` gives you **one** endpoint to create a dump of your database (mysql, postgres, more in the future) and upload it to the cloud storage of your choice (for now only GCP, more in the future).

## Deployment

This section describes how to start the service in Docker. You can derive other deployment methods (e.g. kubernetes) from this.

1. Create a `.env` file based on the `.env.example` file.

```bash
cp .env.example .env
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
