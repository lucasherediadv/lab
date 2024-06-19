# The base image
FROM ubuntu:latest

# Install python and pip
RUN apt-get update && apt-get install -y \
  python3 \
  python3-pip \
  python3-dev \
  build-essential

# Install Python modules needed by the Python app
COPY requirements.txt /usr/src/app/
RUN pip3 install --break-system-packages --no-cache-dir -r /usr/src/app/requirements.txt

# Copy files required for the app to run
COPY app.py /usr/src/app/

# Declare the port number the container should expose.
# The EXPOSE instruction does not actually publish the port. It functions as a
# type of documentation between the person who build the image and the person
# who runs the container, about which ports are intended to be published. You
# need the -p/-P command to actually opne the host ports.
EXPOSE 5000

# CMD Tell the container which command it should run by default when it is started.
CMD ["python3", "/usr/src/app/app.py"]
