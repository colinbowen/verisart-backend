# VERISART API

Build & run main.go

```
docker build -t verisart .
docker run -it --rm -p 8000:8000 verisart
```

## Certificates

### Get a list of all certificates in DB

Send a GET request to

``` http://localhost:8000/api/certificates ```

#### Get a list of all certificates owned by a user

Send a GET request to

```
http://localhost:8000/users/1/certificates
```

Change ownerID to required user.

```
http://localhost:8000/users/{ownerID}/certificates
```

#### Create a certificate

Send a POST request to

```
http://localhost:8000/api/certificates/{ownerID}
```

- Set content/type to application/json
- In body of request input data e.g.

```
{
	"id":"150",
	"title":"New Certificate",
	"createdAt":"0000",
	"ownerID":"Owner 1",
	"year":2019,
	"note":"My newly added note"
}
```

#### Update a certificate

Submit PUT request to

```
http://localhost:8000/api/certificates/1
```

- Make sure body contains data in format

```
{
	"id":"150",
	"title":"New Certificate",
	"createdAt":"0000",
	"ownerID":"Owner 1",
	"year":2019,
	"note":"My newly added note"
}
```

#### Delete a certificate

Submit DELETE request to

```
http://localhost:8000/api/certificates/{id}
```

- Where id is certificate id

## Notes

@TODO Implement transfer of certificates

@TODO Structure project to GO specification

@TODO Implement unit tests for all functions.

@TODO Detailed documentation

@TODO Add support for CORS (Cross Origin Resource Sharing)

@TODO Containerise project using Docker - DONE