# go-tuner-api
### How to run

```sh
# Run the API without docker-compose
$ go run cmd/main.go 

# Success output:
> [GIN-debug] Listening and serving HTTP on :8080
```

### Using `/tune/:frequency` endpoint
```sh
curl http://localhost:8080/tune/440
```
Sample response:

```json
{
  "id": "1",
  "name": "A4",
  "keyNumber": 49,
  "frequency": {
    "min": 408.87616512680097,
    "max": 456.5655947149062
  },
  "pitchPerfect": 432
}
```

`TODO:`
  - gRPC Adapter (NTH)
  - WebSocket Adapter (NTH)
  - Dockerfile (MUST)
  - Docker-compose (MUST)
  - Unit and Integration tests (MUST)
  
