---
id: collect-240926-datacamp/datacamp/les-18-principales-commandes-docker-pour-construire-executer-et-gerer-des-conteneurs
title: "syntax=docker/dockerfile:1"
domain: datacamp
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/clean_en/datacamp/les-18-principales-commandes-docker-pour-construire-executer-et-gerer-des-conteneurs.md
source_anchor: ""
source_lines: [1, 429]
sha256: 891c1a5a2306bca339aac6a302021695ce8cbe48aebd3da6ec9dc84e2476599a
---

# syntax=docker/dockerfile:1

<!-- source: https://www.datacamp.com/fr/blog/docker-commands -->

Course

Docker has become an essential tool for developers and data professionals who need to create, share, and run applications consistently across different environments. Whether you're building containers for local development or deploying microservices in production, mastering Docker commands is essential.

In this guide, I'll introduce you to 18 essential Docker commands, covering images, containers, networking, volumes, and Compose, that will improve your workflow and make your work with Docker smoother and more efficient.

## What is Docker?

Docker is a platform for developing, shipping, and running applications. It allows you to separate your application from the underlying infrastructure, which speeds up software delivery and lets you manage your configuration the same way you manage your applications.

Docker runs applications in isolated environments using lightweight packages called containers, which include everything an application needs to run, such as dependencies and installations, helping you save system resources. We can easily share containers with teammates, run multiple at the same time, and manage them all using Docker's tools and platform.

We can use Docker for many things, including:

- Responsive deployment and scaling.
- Running more workloads on the same hardware.
- Fast and consistent application delivery.

When you use Docker, you work with different Docker objects such as images, containers, networks, plugins, and volumes. These are the building blocks of the Docker installation. Under the hood, Docker uses Linux kernel features to make all of this work. We interact with it using simple commands in the terminal, and every Docker command starts with `docker`.

**> If you're just getting started, the Introduction to Docker will help you get familiar with Docker.***This Introduction to Docker course provides a practical foundation for learning the basics of containerization.***This course provides a practical foundation for learning the basics of containerization.**

## Basic Docker Commands

Now that we've explained what Docker is and how it works, let's look at some of the most common commands. They will help you build, run, and manage containers in your daily work.

### docker --version and docker info

In Docker commands, anything that starts with `--` is called a flag.

For example, `--version` is a flag that indicates the version of the Docker CLI you're using. You can also use `docker version` (without the flag) to get detailed information about the version of all Docker components.

The output is divided into two parts:

- Client presents information about the Docker CLI and associated tools.
- Server displays details about the Docker engine and what it's running on.

You can also format this output using the `--format` option with a custom template.

The `docker info` command gives you a complete overview of your Docker configuration. It's the same as going to the `docker system info` site, but with a shorter name. You'll see details like your kernel version, the number of containers and images, and other system details.

Depending on your storage driver, you may also see information such as pool names and data files. As with `docker version`, you can format the output using `--format` or `-f`.

### docker pull <image>

The pull command downloads a Docker image from a registry, usually Docker Hub, a public library of pre-built images that you can use without configuring anything yourself. You can run it as `docker pull` or `docker pull`, and both will do the same thing.

- The full syntax is as follows: `docker image pull [OPTIONS] NAME[:TAG|@DIGEST]`

If you don't specify a tag, Docker will use `:latest` by default. For example: `docker image pull debian` pulls the `debian:latest` image.

You can also add options after the command to customize how the image is pulled, for example by limiting bandwidth or skipping image verification. The image below shows all available options and their functions.

Options for the docker pull command. Sour*ce: Docker Documentation*

### docker run <image>

The `docker run` command creates and starts a new container from a specified image, meaning it runs the image in a new container. It's a shortcut for `docker container run`, and both work the same way.

- Here's the basic syntax: `docker container run [OPTIONS] IMAGE [COMMAND] [ARG...]`

If you've already launched a container and want to restart it with all previous changes, use: `docker start`.

The run command offers many options for customizing how your container runs. We'll go over a few of the most common ones with examples:

| Flag | Example command | Description |
| `-- name` | `docker run --name test -d nginx:alpine` | The custom identifier specified for the container named test using the `nginx:alpine` image. |
| `-w` ,`--workdir` | `docker run -w /path/to/dir/ -i -t ubuntu pwd` | Runs the command in the specified directory, in this example, `/path/to/dir/` |
| `--pid` | `docker run --rm -it --pid=host alpine` | By default, the PID namespace is enabled in all containers, which allows processes to be separated. The example uses an alpine container with the `--pid = host` option. |
| `--cidfile` | `docker run --cidfile /tmp/docker_test.cid ubuntu echo "test"` | This creates a container and prints a test to the console. The `cidfile` option allows Docker to create a new file and write the container ID to it. |

### docker stop <container> and docker start <container> :

The `docker start` command starts one or more stopped containers. For example, in `docker start my_container`, `my_container` is the name of the container we want to start. We can also use its alias: `docker container start`.

- Its full syntax is as follows: `docker container start [OPTIONS] CONTAINER [CONTAINER...]`

Similarly, `docker stop` stops one or more running containers. For example, in `docker stop my_container`, `my_container` is the name of a running container.

It also has an alias: `docker container stop`. 

- Its full syntax is as follows: `docker container stop [OPTIONS] CONTAINER [CONTAINER...]`

Like `start`, the `stop` command has options that allow you to customize how containers are stopped. We will now present a few of them to you:

Options for `docker stop`. Source*: Docker Documentation*

## Mastering Docker and Kubernetes

## Working with Docker images

Images are the foundation of every container. In this section, we will see how to create, manage, and inspect Docker images using its common commands.

### docker build

The `docker build` command is one of Docker's most widely used features. Although it is part of a larger ecosystem that supports advanced use cases, we will focus on how to use it to build an image from a simple Dockerfile. 

A `Dockerfile` is a plain text file (without an extension) that contains step-by-step instructions that Docker uses to build an image. Here is how to create one:

1. In the root directory of your application, create a file named `Dockerfile` with the following content:

```
# syntax=docker/dockerfile:1
FROM node:lts-alpine
WORKDIR /app
COPY . .
RUN yarn install --production
CMD ["node", "src/index.js"]
EXPOSE 3000
```
This Dockerfile starts with a lightweight base image that includes Node.js and Yarn. It copies your application's source code into the image, installs the dependencies, and defines how to start the application.

1. Now, build the image using the following command:

`docker build -t getting-started.`
The `-t` flag allows you to tag the image. In this case, we named it `getting-started.` The `.` at the end tells Docker to look for the Dockerfile in the current directory.

### docker images

The `docker images` command lists all your top-level images, their repository, their tags, and their size. You can also use its aliases:

- `docker image list`
- `docker image ls`

Its syntax is as follows: `docker image ls [OPTIONS] [REPOSITORY[:TAG]]`

This command has several options. For example, to display all images, including intermediate images, you can add the `-a` or `--all` flag as follows: `docker images -a`.

### docker rmi <image>

The `docker rmi` command removes one or more images from your system. If an image has multiple tags, running this command with a specific tag will only remove that tag. But if the tag is the only one linked to the image, both the tag and the image will be removed.

You can also use one of these aliases:

- `docker image remove`
- `docker image rm`

Its syntax is as follows: `docker image rm [OPTIONS] IMAGE [IMAGE...]`

If you need to remove an image still being used by a running container, you will need to force it by adding the `-f` or `--force` option.

## Managing Docker containers

We often need to manage containers - start them, stop them, inspect them, or remove them as our application evolves. So I will introduce you to the most useful Docker commands for managing containers in your daily work.

### docker exec <container> <command>

The `docker exec` command allows you to run a command inside a running container without restarting it. This is particularly useful for debugging or manually checking something inside the container. You can also use its alias: `docker container exec`.

The command only works if the container's main process (PID 1) is running. It does not run automatically if the container restarts.

Its syntax is as follows: `docker exec [OPTIONS] CONTAINER COMMAND [ARG...]`

You can use several optional flags. For example, `--privileged` gives the command extended permissions inside the container. For the full list of options, see the official documentation.

Here is an example of running a command on the container:

`docker exec -d mycontainer touch /tmp/execWorks`
The `touch` command creates a new file `/tmp/execWorks` in the running container `mycontainer`, in the background.

### docker logs <container>

The `docker logs` command allows you to view the logs of a specific container and display everything printed to standard output and errors at runtime. You can also use its alias: `docker container logs`

Its syntax is as follows: `docker container logs [OPTIONS] CONTAINER`

You can add a number of useful options. For example:

- `--details` displays additional attributes such as environment variables and labels.
- `--until` allows you to retrieve logs up to a specific time.
- `docker logs -f --until=2s test` follows the log output of the `test` container and stops after displaying the last two seconds of logs.

Here are all the options we can use with `docker logs`:

Options for `docker log`. Source: Docker Documentation

### docker rm <container>

The `docker rm` command removes one or more containers from your system. You can also use its aliases:

- `docker container remove`
- `docker container rm`

Its syntax is as follows: `docker container rm [OPTIONS] CONTAINER [CONTAINER...]`

This command has a few options you can use. Here are some of them:

Options for `docker rm`. Source: Docker Documentation

For example, you can use `docker rm /redis` to remove the container identified by the link `/redis`. Note, however, that this command only removes running containers.

If you want to remove stopped containers, you need to use `docker container prune`. To keep your environment clean, learn how to safely remove unused Docker resources with this Docker prune tutorial.

### docker restart <container>

The `docker restart` command stops and then restarts one or more containers. You can also use its alias: `docker container restart`.

Its syntax is as follows: `docker restart [OPTIONS] CONTAINER [CONTAINER...]`

The restart command has several useful options for customizing container restarts. Below are the two options, along with examples.

| Option | Description | Example |
| -s, --signal | Signal to send to the container. | `docker restart -s SIGTERM mycontainer` |
| -t, --timeout | Number of seconds to wait before killing the container. | `docker restart -t 10 mycontainer` |

> For hands-on practice, explore these Docker project ideas that range from beginner to advanced level.

## Docker Networking

Container networking enables communication between containers and external workloads. By default, containers have networking enabled and can establish outbound connections, but they don't automatically know what type of network they're on or what other workloads they're connected to.

Unless you use the `none` network driver (which disables networking), containers can interact with network elements such as IP addresses, gateways, and DNS.

Let's explore some common Docker networking commands you can use.

### docker network ls

The `docker network ls` command lists all networks known to the Docker engine, including those on multiple hosts in a cluster. Its alias is: `docker network list`.

- Here is its syntax: `docker network ls [OPTIONS]`

You can use several options, as shown in the image below:

By default, this command displays the name of each network:

- ID
- Name
- Driver
- Scope

You can use flags to customize the output. For example, the `--no-trunc` flag displays full network IDs instead of shortened IDs. Here's how to use this command:

`docker network ls --no-trunc`
### docker network create <network name>

The `docker network create` command creates a new Docker network. By default, it uses the `bridge` driver, unless you specify another one using the `--driver` (or `-d`) flag.

Docker supports built-in network drivers such as:

- `bridge` for single-host networks.
- `overlay` for multi-host networks in swarm mode.

You can also use third-party or custom drivers if needed.

- Here is the basic syntax: `docker network create [OPTIONS] NETWORK`

The command offers many options for different purposes. Consult the official documentation to see the complete list of options.

Here is an example of creating a bridge network:

`docker network create -d bridge my-bridge-network`
Bridge networks are limited to a single Docker engine, so they do not connect containers on different hosts.

Once Swarm mode is enabled, you can create a network that spans multiple Docker hosts:

`docker network create --scope=swarm --attachable -d overlay my-multihost-network`
> Curious about how Docker compares to Kubernetes? This Kubernetes vs Docker analysis covers the key differences and use cases.

## Docker Volumes

Docker volumes are used to store data that must persist, even when a container stops or is deleted. You can create them explicitly or let Docker create them automatically when starting a container.

Volumes are stored on the host system but are isolated from the host's main files. They are mounted into containers in a manner similar to bind mounts, but with better portability and security.

Let's explore some relevant Docker volume commands:

### docker volume ls

The `docker volume ls` command lists all volumes known to Docker. You can also use its alias: `docker volume list`.

- Its syntax is as follows: `docker volume ls [OPTIONS]`

This command supports a few optional options to help you filter or format the output - here's what they are:

Options for the `docker volume ls` command. Source: Docker Documentation

### docker volume create <volume name>

The `docker volume create` command creates a new volume for storing persistent data. If you don't provide a name, Docker generates one for you automatically.

Creating volumes is a common step when you want data to persist beyond the lifetime of a single container.

- Here's the syntax: `docker volume create [OPTIONS] [VOLUME_NAME]`

Let's look at an example of creating a volume and configuring a container to use it:

```
docker volume create hello
docker run -d -v hello:/world busybox ls /world
```
In this example, a volume named `hello` is created. It is then mounted in a container at `/world`. This allows the container to write or read data on this volume.

Multiple containers can use the same volume, which is useful if one container needs to write data while another reads it.

Note: Volume names must be unique across drivers. You cannot use the same volume name in two different storage drivers.

## Docker Compose Commands

Docker Compose makes it easy to manage multi-container applications using a simple YAML file. It supports different environments such as development, testing, staging, production, and critical analysis. With a single command, you can control services, configure networks, and manage volumes, all in one place.

Let's look at some common compose commands:

### docker-compose up

The `docker compose up` command builds, (re)creates, starts, and attaches containers for a service. If the containers aren't running yet, it also starts all related services automatically.

- Here's the syntax: `docker compose up [OPTIONS] [SERVICE...]`

By default, this command combines the output of all containers. If you want to focus on specific services, you can do the following:

- Use the `--attach` flag to attach to certain services.
- Use the `--no-attach` option to exclude others

For example, `docker compose up --no-attach` starts all services except the one you excluded from the logs.

When the command finishes, the containers also stop. To keep them running in the background, use the `--detach` option:

`docker compose up --detach`
Before running this command, make sure you've navigated (`cd`) to the directory where your `docker-compose.yml` file is located.

### docker-compose down

The `docker compose down` command stops the containers and removes the containers, images, networks, and volumes created by `docker compose up`.

- Its syntax is as follows: `docker compose down [OPTIONS] [SERVICES]`

You can use various options, including the following:

Options for the `docker compose down` command. Source: Docker docs.

By default, the command removes the following:

- Containers for the services defined in the Compose file.
- Networks defined in the networks section of the Compose file.
- The default network, if there is one.

The following are not removed by default:

- Networks and volumes defined as external.
- Anonymous volumes, i.e., volumes that don't have a name.

Anonymous volumes are not automatically mounted when you run `docker compose up` again, because they don't have a name. If you need persistent data storage, it's better to use named volumes or bind mounts.

## Best Practices for Using Docker Commands

Docker is a powerful way to build, ship, and run applications in containers. But to use it properly and make your setup efficient and easily scalable, you should follow some best practices.

### Use Docker Volumes for Persistent Data

By default, container files are stored in a writable layer that is lost when the container is deleted! This layer is specific to each container and is not easily accessible. Docker uses different types of mounts to persist data, one of which is volumes, managed by the Docker daemon and stored on the host.

Volumes allow us to:

- Keep data even after a container is deleted.
- Store performance-critical data with host-level speed.
- Easily manage storage through Docker.

They are ideal for long-term data or when multiple containers need to share access. Keep in mind that if you need to access files directly from the host, bind mounts may be better suited, since volumes are entirely managed by Docker.

### Automate with Docker Compose

Manually managing multiple containers can be tedious. Docker Compose simplifies things by letting you define everything in a single YAML file so you can focus on building.

Here's why this is a good practice to follow:

- A single command: Start, stop, adapt, or rebuild your services all at once.
- Consistency for everyone: Whether you work in development, testing, or production, Compose aligns environments so you don't run into "it works on my machine" issues.
- Built-in networking: Compose creates a shared network so your services can easily communicate using service names instead of IP addresses.
- Easy scaling: You can quickly increase or decrease the size of services with the `--scale` option, which is ideal for testing how your application handles different loads.
- Clear and collaborative configuration: Your entire configuration, including containers, networks, and volumes, is version-controlled and readable.

## Conclusion

Docker can seem like a lot at first, but once you master the basic commands, it offers many possibilities. From running your first container to managing networks, volumes, and services, you are now better equipped to build and run containerized applications with confidence.

And if you want to keep learning, here are some interesting resources to explore:

- Containerization and Virtualization Concepts - a perfect course to build your conceptual foundation.
- Introduction to Docker - a beginner-friendly course to help you get started.
- Intermediate Docker - a course for those who want to go further.
- Containerization and Virtualization with Docker and Kubernetes - a skill path to develop on Kubernetes and real-world orchestration.

## Mastering Docker and Kubernetes

## FAQ

### What are the most commonly used Docker commands?

**Among the most common Docker commands are `docker run`, `docker ps`, `docker build`, `docker pull`, and `docker-compose up`.**

### How do Docker volumes differ from bind mounts?

**Docker volumes are managed by Docker and are ideal for portability and data persistence, while bind mounts are tied to a specific file path on the host system.**

### Can I launch multiple containers using a single command?

**Yes, using `docker-compose up` with a `docker-compose.yml` file allows you to run multiple services simultaneously.**

### What is the difference between docker start and docker run?

`docker run` creates and starts a new container, while `docker start` restarts an existing stopped container.

### How can I list all stopped containers?

**Use `docker ps -a` to see all containers, including those that are stopped.**

### What is the purpose of the docker exec command?

**It allows you to run commands inside a running container, often used for debugging or manual checks.**

### Is Docker only for Linux-based systems?

**No, Docker also supports macOS and Windows, using lightweight virtual machines to enable containerization.**

### What does the docker-compose down command do?

**It stops and removes the containers, networks, and volumes created by `docker-compose up`.**

### How do I remove dangling Docker images?

**Run `docker image prune` to clean up unused images and free up disk space.**

### Can Docker be used in production environments?

**Absolutely. Docker is widely used in production to deploy scalable and reproducible applications across cloud and on-prem configurations.**

I am a content strategist who enjoys simplifying complex topics. I have helped companies like Splunk, Hackernoon, and Tiiny Host create engaging and informative content for their audiences.
