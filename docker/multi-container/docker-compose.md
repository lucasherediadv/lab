# Docker compose

Docker compose is a tool for defining and running your multi-container Docker applications.

Your applications can be defined in a YAML file where all the options that you used in `docker run` are defined.

Compose also allows you to manage your applications as a single entity rather than dealing with individual containers.

The example of `multi-container.md` is made in `docker-compose.yaml` with Compose.

Create and start the services stated in the compose file:

```
docker compose up
```

Stop and remove containers, networks, images and volumes:

```
docker compose down
```

