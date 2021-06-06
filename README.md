Starting with docker-compose needs app image, which can be built with ```make dockerbuild```

And also in the directory you need a ```.env``` file

Sample:

```
POSTGRES_HOST=postgres12
POSTGRES_PORT=5432
POSTGRES_USER=root
POSTGRES_DBNAME=root
POSTGRES_PASSWORD=secret
POSTGRES_PORT=5432

REDIS_ADDR=redis-hub:6379
REDIS_DB=0

ELASTICSEARCH_URL=http://elasticsearch:9200

JWT_SIGNINGKEY=signing_key

PASSWORD_SALT=1234
```
