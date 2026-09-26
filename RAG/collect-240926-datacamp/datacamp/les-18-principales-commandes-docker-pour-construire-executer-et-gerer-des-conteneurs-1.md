---
id: collect-240926-datacamp/datacamp/les-18-principales-commandes-docker-pour-construire-executer-et-gerer-des-conteneurs-1
title: "syntax=docker/dockerfile:1"
domain: datacamp
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/clean_en/datacamp/les-18-principales-commandes-docker-pour-construire-executer-et-gerer-des-conteneurs.md
source_anchor: ""
source_lines: [1, 126]
sha256: a9c15f9bdd80cae6a4e817f76173e98922ea49824b99af1c01713ffd891743c4
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

