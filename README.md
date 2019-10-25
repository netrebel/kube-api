# kube-api


# Build docker image

```
docker build -t kube-api .
```

# Run docker image

```
docker run -d -p 8080:8080 kube-api
```

# Run docker image with logs volume mounted

```
mkdir -p ~/logs/kube-api
docker run -d -p 8080:8080 -v ~/logs/kube-api:/app/logs kube-api
```

# docker-compose

* `docker-compose build api` builds the image
* `docker-compose run api bash` connects you to the image
* `docker-compose up -d` starts the service in detached mode

# Logs

`tail -f ~/logs/kube-api.log`