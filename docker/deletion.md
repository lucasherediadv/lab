# Throw your container away

As containers are just a thin base layer on top of the host kernel, it is really fast to spin up a new instance if you crashed your old one.

Try to run an alpine container with `docker run -ti alpine` and delete the file system:

```
rm -rf /bin
```

This will remove almost all binaries from your container

Exit out by pressing `Ctrl+d` and create a new instance of the Alpine image. Now you have a fresh new instance ready to go.

### Auto-remove a container after use

If you are creating a lot of new containers, you can tell Docker daemon to remove the container once stopped with the `--rm` option:

```
docker run --rm -it alpine
```

This will remove the container immediately after it is stopped.

### Cleaning up containers you do not use anymore

Containers are still persisted, even though they are stopped. If you want to delete them from your server you can use the `docker rm` command.

```
docker container rm ecstatic_cray
```

`docker rm` can take either the `CONTAINER ID` or `NAME`.

### Deleting images

List images you have downloaded:

```
docker image ls
```

Use the `docker image rm` command together with the id of the docker image:

```
docker image rm 48b5124b2768
```

What docker did here was to `untag` the image removing the references to the SHA of the image. After the image has no references, removes the two layers that make up the image itself.

### Cleaning up

Docker provides a `prune` command, taking al dangling container/images/networks/volumes:

- `docker container prune`
- `docker image prune`
- `docker network prune`
- `docker volume prune`

If you want a general cleanup:

```
docker system prune
```

