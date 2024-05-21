# Multi container

### Making a container network

In order to connect multiple docker containers without binding them to the hosts network interface we need to create a docker network.

Make a new network for the containers to communicate through:

```
docker network create if_wordpress
```

You can take a deeper look into the container network:

```
docker network inspect if_wordpress
```

### Connect containers to the network

Now you need to connect containers to the network, by adding the `--network` option:

```
docker run --name mysql-container --rm --network if_wordpress -e MYSQL_ROOT_PASSWORD=wordpress -e MYSQL_DATABASE=wordpressdb -d mysql:latest

docker run --name wordpress-container --rm --network if_wordpress -e WORDPRESS_DB_HOST=mysql-container -e WORDPRESS_DB_PASSWORD=wordpress -e WORDPRESS_DB_USER=root -e WORDPRESS_DB_NAME=wordpressdb -p 8080:80 -d wordpress:latest
```

As, we have linked both the container now wordpress container can be accessed from browser using the address http://localhost:8080. MYSQL is not accessible from the outside.
