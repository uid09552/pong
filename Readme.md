# This small service will pong the headers provided in the request as json response
* server will run on port 5000
* content will be served at path /request


# Build & Run
```
docker build . -t myimage
docker run -it -p 5000:5000 myimage
```

## Prebuild image
```
docker run -it -p 5000:5000 uid09552/pong
```

## As k8s deployment
```
kubectl apply -f k8s/*.yml
```