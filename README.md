# go-tuner-api
## 1. How to run 🚀

Running the API using docker-compose
```sh
docker-compose up
```


Running the API without Docker
```sh
$ go run cmd/main.go 
```


## 2. Using the tune endpoint 🎵
Easily use the `/tune/:frequency` endpoint
```sh
# example for retrieving a Note from 110.1Hz
curl http://localhost:3000/tune/110.1
```
🎶 Sample response:

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

------

### TODO:
Status | Feature
:---:| ---
⬜️| gRPC Adapter (NTH)
⬜️| WebSocket Adapter (NTH)
✅| Dockerfile
✅| Docker-compose
⬜️| Unit and Integration tests (MUST)

