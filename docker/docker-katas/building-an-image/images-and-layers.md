# Images and layers

A layer, or image layer, is a change on an image. Every time you run one of the commands RUN, COPY or ADD in your Dockerfile it add a new layer, causes the image to change to the new layer.

Layers are intermediate images, if you make a change to your Dockerfile, docker will build only the layer that was changed and the ones after that. This is called layer caching.

Each layer is build on top of it's parent layer, meaning if the parent layer changes, the next layer does as well.

If you want to concatenate two layer (e.g. the update and install) then:

```
FROM ubuntu:latest
RUN apt-get update && apt-get install -y \
 python3 \
 python3-pip \
 python3-dev \
 build-essential
COPY requirements.txt /usr/src/app/
RUN pip3 install --no-cache-dir -r /usr/src/app/requirements.txt
COPY app.py /usr/src/app/
EXPOSE 5000
CMD ["python3", "/usr/src/app/app.py"]
```

If you want to be able to use any cached layers from the last time, they need to be run before the update command.

> NOTE:
> Once we build the layers, Docker will reuse them for new builds. This make the builds much faster.

### Every layer can be a container

Docker command will create a new layer in your image, and therefore also be an image of their own:

```
 ---> c1f2dc732c7c
Removing intermediate container f92f9c719287
Step 6/8 : COPY app.py /usr/src/app/
 ---> 6ed47d3c544a
Removing intermediate container 61a68a949d68
Step 7/8 : EXPOSE 5000
 ---> Running in 1f939928b7d5
 ---> 6c14a93b72f2
```

What docker does is:

- Taking the layer created just before
- Make a container based of it
- Run the command given
- Save the layer

In a loop until all the commands have been made.

