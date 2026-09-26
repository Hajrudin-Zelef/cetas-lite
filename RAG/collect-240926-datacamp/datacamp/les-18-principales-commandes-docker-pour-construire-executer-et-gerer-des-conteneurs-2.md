---
id: collect-240926-datacamp/datacamp/les-18-principales-commandes-docker-pour-construire-executer-et-gerer-des-conteneurs-2
title: "syntax=docker/dockerfile:1"
domain: datacamp
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/clean_en/datacamp/les-18-principales-commandes-docker-pour-construire-executer-et-gerer-des-conteneurs.md
source_anchor: ""
source_lines: [127, 288]
sha256: d0b8bc1e60d1de6b38aa2e02e8ddff53e161c62bf82906ea39642fd4a1ba37a3
---

# syntax=docker/dockerfile:1

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

