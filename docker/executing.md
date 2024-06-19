# Executing processes in your container

If you want to examine a running container, but do not want to disturb the running process you can execute another process inside the container with `exec`.

This could be a shell, or a script of some sort. In that way you can debug an existing environment befre starting a new up.

```
docker run -d -p 8000:80 nginx
```

Step into a new container by executing a bash shell inside the container:

```
docker exec CONTAINERNAME bash
```

