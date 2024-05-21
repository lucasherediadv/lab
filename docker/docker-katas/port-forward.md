# Basic webserver

Pull down the `nginx` Docker image from the Docker Hub:

```
docker pull nginx
```

Start a new container from the nginx image that exposes port 80 from the container to port 8080 on your host:

```
docker run -p 8080:80 nginx
```

### Running the webserver container in the background

Docker enables this with the `-d` parameter for the `run` command. `d` is the short from of `detach`.

```
docker run -p 8080:80 -d nginx
```

Docker prints out the container ID and returns to the terminal.

### Clean up

Stop the container you just started:

```
docker stop <container-ID>
```

