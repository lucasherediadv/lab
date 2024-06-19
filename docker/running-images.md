# Running your first container from image

Pull the Alpine Linux (a lightweight linux distribution) image:

```
docker pull alpine
```

The pull command fetches the alpine image from the Docker registry and saves it in your system. You can use `docker image ls` command to see a list of all images on your system.

### docker run

```
docker run alpine ls -l

docker run echo "Hello from alpine"

docker run alpine /bin/sh
```

These interactive shell will exit after running any scripted commands, unless they are run in an intereactive terminal. To not exit, you need to add the parameters `i` and `t`.
The flag `-i` is the short form of `--interactive` (Keep STDIN open)  and `t` is the short form of `--tty` (Allocate a terminal).

```
docker run -it alpine /bin/sh
```

Now you are inside the container.

### docker ps

The `docker ps` command shows you all containers that are currently running.

The `docker ps -a` list all containers, both stopped and started.

### Naming containers

All containers have and ID and name, both are generated every time a new container spins up with a random seed for uniqueness.

If you want to assign a specific name to a container then you can use the `--name` option.

