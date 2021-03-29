# Docker
### What is Docker?
### What is Container?
 - Container is a way to **package the application** with all the neccessary dependencies and configuration.
 - Run on own isolated environment
 - **Portable Artifact** easily shared and moved around between various team
 - Make development and deployment process more effcient.
 - Container images available in the container repository.
 - Run the same application with different version dependencies. 
### How is Containers work?
 - Developer and operation team work together and package the application.
 - **No server configuration** needed, except docker runtime.
### How is Containers build?
 - Layers of images
 - Mostly **Linux Base image**, becasuse small size
 -  Application image on top of base image.

### Docker vs Virtual Machine
  - Docker virtualize OS Application layer alone.
  - Docker images small size, in terms of few MB.
  - Docker images run faster than VM.
  - VM virtualize OS Kernal & application.
### Difference between Container and image
- **Container** is a running Environment of **Image**
### Container port vs Host port
- **Multiple containers** can run on your host machine(your laptop).
- Your laptop has only certain ports available.
- Port mapping helps to resolve the conflict.
### Main Docker commands
```cmd
docker images # list all the images
docker ps # list all the running containers
docker ps -a # list all the running/not running containers.
docker run -d redis # run the container in detach mode
docker stop <container_id> #  stops the container
docker start <container_id> # start the container
docker run redis:4.5 # Pull the images and start the container

docker run -p<host_port>:<container_port> -d redis  # run the image and localhost port.
```
### Debug Docker commands
```cmd
docker logs <container_id or container name>
docker exec -it <container_id or container name> /bin/bash/ # look inside the container cmd.
```
### Docker compose?
### Deploy containerized Applications
### Volumes - Presisting Data
