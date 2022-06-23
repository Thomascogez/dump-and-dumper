![Dump & dumper](./assets/dump-and-dumper.png)

<center>
    <h1>Dump & dumper</h1>
</center

Dump & dumper is a simple cli tool written in go to ease the process of containerized database backup and upload result into a s3 bucket

> Dump & dumper 1.0.0 currently support only _postgress_ database

# Install

Using *Homebrew*

```sh
brew tap Thomascogez/thomascogez

```
Then
```sh
brew install dump-and-dumper
```

# Usage


Add labels to the containers you want to dump

The above labels are currently available
| label             | description                                        |
| ----------------- | -------------------------------------------------- |
| go-dumper.enabled | Use to determined if the container should be dump  |
| go-dumper.user    | The database user to use during the dump process   |
| go-dumper.type    | The database type (for now only "pg" is available) |

> docker-compose example
```yml
version: '3.0'
    services:
        pg_dump_1:
        image: postgres:latest
        environment:
        POSTGRES_PASSWORD: test
        POSTGRES_DB: test
        labels:
        - 'go-dumper.enabled=true'
        - 'go-dumper.user=postgres'
        - 'go-dumper.type=pg'
```

Then start or restart your docker containers


After that you can run the dump command using desired s3 options

The above flags are available

| flag              | description                        |
| ----------------- | ---------------------------------- |
| --s3-endpoint     | Set the endpoint of your s3 bucket |
| --s3-bucket       | Set the targeted bucket name       |
| --s3-secretKeyId  | Set the secret key id              |
| --s3-secretKey    | Set the secret key                 |

```sh
dump-and-dumper dump  --s3-endpoint=S3_ENDPOINT --s3-bucket=BUCKET_NAME --s3-secretKeyId=SECRET_KEY_ID --s3-secretKey=SECRET_KEY --s3-region=REGION

```
