i am following through this youtube video and its corresponding written blog.
https://www.youtube.com/watch?v=RqTEHSBrYFw&t=35s
https://courses.devopsdirective.com/docker-beginner-to-pro/lessons/04-using-3rd-party-containers/01-data-persistence
- first understand virtualization and containers under the hood
- [virtualization](../os/virtualization.md)
- [containers](../os/containers.md)

## Concepts:
### Docker Container Lifecycle
A Docker container has multiple states:
- Created: The container is defined (from an image), but not running yet.
- Running: Actively executing the container’s main process.
- Exited (aka Stopped): The main process finished or crashed, and the container is no longer running.
- Removed: The container is deleted from the system — it no longer exists.
### Persisting Data Produced by the Application:
[source](https://courses.devopsdirective.com/docker-beginner-to-pro/lessons/04-using-3rd-party-containers/01-data-persistence)
![alt text](images/2.png)
- allow us to specify a location where data should persist beyond the lifecycle of a single container. 
 - The data can live in a location managed by Docker (volume mount), a location in your host filesystem (bind mount), or in memory (tmpfs mount, not pictured).
- by default docker is volume mount
 - tmpfs mount does not persist the data after the container exits, and is instead used as a temporary store for data you specifically DON'T want to persist

### objects
[source](https://courses.devopsdirective.com/docker-beginner-to-pro/lessons/10-interacting-with-docker-objects/01-images)
- docker mostly consists of 4 objects or components which we need to set up
  - image
  - containers
  - volume
  - network
### build context
- To resolve relative paths inside the Dockerfile

- To determine what files get sent to Docker during docker build

- The build context is the directory you provide to Docker when running the docker build command. This directory and its contents are available to the Dockerfile during the build process, meaning that when you use COPY, ADD, or similar commands in the Dockerfile, Docker looks for the files relative to the build context.

eg: When You Run docker build -t godocker:v1 -f ./docker_images/Dockerfile.golang .
Here’s what’s happening step by step:

Context (.): The . at the end of the command means "use the current directory as the build context." So when you run docker build from a directory (say project/), the contents of project/ are available to Docker as the build context.

Dockerfile Location (-f ./docker_images/Dockerfile.golang): This specifies the path to the Dockerfile. In this case, it's located in the docker_images/ folder. Docker will use this specific Dockerfile to build the image.
- so make sure the build context covers everyfile and advisable to be in root so that the image file understands it
### combining docker compose
docker compose -f $(DEV_COMPOSE_FILE) -f $(DEBUG_COMPOSE_FILE) up --build
- You are combining multiple Docker Compose files (DEV_COMPOSE_FILE and DEBUG_COMPOSE_FILE) into a single configuration using override behavior.
- If both files define a service with the same name (e.g., golang), the configuration from the second file (DEBUG_COMPOSE_FILE) will override if the same command exists or merge into the first.
- second compose will override which is already there and the not common ones are retained from first file
services:
  db:
    container_name: brainwars_pgsql
  golang:
    dsfhsdkfhksdjh
db = service name (this becomes the Docker DNS hostname that other services in the network to use)
brainwars_pgsql = actual container name (visible via docker ps, for humans)
- When your golang service starts, Docker Compose:
  - Connects both containers (golang and db) to the same default network.
  -Sets up internal DNS so that db resolves to the IP 
# your conatiners are in a private neighborhood while using docker compose
- When you're using Docker Compose:
  - Each service (like db or golang) is like a house in a private neighborhood (Docker network).
  - Inside that neighborhood, you don't refer to houses by their fancy house names (container names) — you refer to them by their service names.
### tips and tricks
- for dev environment bind mount the source file into the container so that any change in source code will immideately be displayed (hot reloading)
- have 2 docker compose file , one for production and one for dev. build the image such a way that the dockerFile should solve both dev and prod
- All relative paths in docker-compose.yml are relative to the location of the docker-compose.yml file itself.
- named volume are not deleted when you do docker compose command. we need to explicitly remove the volume.
- 
##  Commands:
1. run an container (which includes pulling the image if not exists)
```bash
docker run --interactive --tty --rm ubuntu:22.04
```
- ubuntu:22.04 is name:tags
- If ubuntu:22.04 is not available locally, Docker pulls it from Docker Hub 
```bash
docker pull ubuntu:22.04
```
This image consists of several read-only layers that make up the Ubuntu root filesystem.

**--interactive (-i):**
- Purpose: Keeps the container’s stdin (standard input) open even if you're not attached to it directly.
- Without -i The container runs, but you can't provide input.
- Without -i is Useful for containers that run background jobs or services.
- With -i: The container accepts input from your terminal.Required if you want to type inside the container (e.g., a shell session).
**--tty (-t):**
- gives you a proper look for interactive terminal. with -t or --tty the interactive standard input looks like a proper formatted terminal
**--rm:**
- When you type exit:
Docker stops the container
- Because of --rm, Docker automatically deletes the container and its writable layer
---
2. list all containers that are running
```bash
docker ps
```
---
3. list all containers taht are running as well as stopped or exited
```bash
docker ps -a
```
![alt text](images/1.png)
---
4. giving a name to the created container
- all containers has an unique id apart from that we can name the container
- if we are not giving our name docker gives random name to our container
```bash
docker run -it --name void-ubuntu-container ubuntu:22.04
```
---
5.  Build a container image  as base and softwares installed while building the container 
- we usually dont create the container and install the required softwares upon logging into the container. we preinstall and build
- for example we need to ping after running the ubuntu container. so we need to install ping inside ubuntu container
```bash
# Build a container image with ubuntu image as base and ping installed
docker build --tag my-ubuntu-image -<<EOF
FROM ubuntu:22.04
RUN apt update && apt install iputils-ping --yes
EOF

# Run a container based on that image
docker run -it --rm my-ubuntu-image

# Confirm that ping was pre-installed
ping google.com -c 1 # Success! 
```
- this is the right way to do
---
6. DATA MOUNT:
 a) volume mount:
- We can use volumes and mounts to safely persist the data.
mounting volume
  ![alt text](images/4.png)
 b) bind mount:
- we can mount a directory from the host system using a bind mount
```bash
# Create a container that mounts a directory from the host filesystem into the container
docker run  -it --rm --mount type=bind,source="${PWD}"/my-data,destination=/my-data ubuntu:22.04
# Again, there is a similar (but shorter) syntax using -v which accomplishes the same
docker run  -it --rm -v ${PWD}/my-data:/my-data ubuntu:22.04

echo "Hello from the container!" > /my-data/hello.txt

# You should also be able to see the hello.txt file on your host system
cat my-data/hello.txt
exit
```
- volume mount is preferred and it is the default
---
7. to list all the volumes
```bash
docker volume ls
```
---
8. go to an exited continer
```bash
docker start void-ubuntu-container
```
- Starts a container that was previously created or stopped, identified by name (void-ubuntu-container) or ID.
- It does not attach your terminal to it — it starts the container in the background.

```bash
 docker attach void-ubuntu-container
```
- Attaches your terminal to a running container’s STDIN/STDOUT/STDERR.
- You'll see output from the container and can interact if the container has a foreground process (like a shell or sleep, etc.).
- To detach safely, use Ctrl+P then Ctrl+Q.
---
9. common docker run options

- -d
  - -d (Detach): Run a container in the background.
  ```bash
  docker run -d ubuntu sleep 5
  ```
---
10. docker compose
```bash
docker compose up -d
docker-compose -f docker-compose-dev.yml up -d # if we use some other name for docker-compose
docker-compose -f ./docker/docker-compose-dev.yaml up --build # to explicitly build the image not take from cache

```
using -v flag will remove everything including volume
```bash
docker compose down -v
```
---
11. enter into running container 
```bash
docker exec -it <container name or id> bash
```
- multistage docker build
https://docs.docker.com/build/building/multi-stage/

---
12. docker images:
- list all images
```bash
docker image ls
```
- remove image
```bash
docker rmi <image name:tag or image id>
```
---
13. docker container logs
```bash
docker logs --details <container id or container name:tags>
```
---
14. export and save docker image/ load
```bash
docker save  -o image.tar <image id or name:tag>
 docker load -i image.tar
```

---
15. volume
Docker volumes are stored on your host machine
```bash
docker volume ls
docker volume rm <volume-name>
docker compose down -v # add -v to remove the volume corresponds to docker compose
docker volume inspect pg_data
```
- cache mount 
- size and speed
- buildx
