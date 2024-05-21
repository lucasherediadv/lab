# Constructing a docker image

A Dockerfile is a text file containing a list of commands that the Docker daemon calls while creating an image. The Dockerfile contains all the information that Docker need to know to run the app:

```
FROM ubuntu:latest
RUN apt-get update -y
RUN apt-get install -y python3 python3-pip python3-dev build-essential
COPY requirements.txt /usr/src/app/
RUN pip3 install --no-cache-dir -r /usr/src/app/requirements.txt
COPY app.py /usr/src/app/
EXPOSE 5000
CMD ["python3", "/usr/src/app/app.py"]
```

### Build the image

The docker build command takes and optional tag name with the `-t` flag, and the location of the directory containing the Dockerfile. The `.` indicates the current directory:

```
docker build -t myfirstapp .
```

### Run your image

Run the image and see if it actually works:

```
docker run -p 8888:5000 --name myfirstapp myfirstapp
```

Expected output:

```
 * Running on http://0.0.0.0:5000/ (Press CTRL+C to quit)
```

