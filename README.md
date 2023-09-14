# ms-mongodb-api


## Generic Mongodb Restfull Api

### CRUD

```bash
GET    "/api/:documents"
GET    "/api/:documents/:id"
POST   "/api/:documents"
PUT    "/api/:documents/:id"
DELETE "/api/:documents/:id"
```

By default the absent collection will be created.

### Ops Health

```bash
GET "/ops/ping"
```

### Environment Parameters

```env
PORT=8080
ENVIRONMENT=development
DATABASE_URL=mongodb://localhost:27017
DATABASE_NAME=test
CONNECTION_TIMEOUT=3
```



References :

* https://blog.devops.dev/make-swagger-from-your-http-rest-full-in-golang-4dc9088f5292