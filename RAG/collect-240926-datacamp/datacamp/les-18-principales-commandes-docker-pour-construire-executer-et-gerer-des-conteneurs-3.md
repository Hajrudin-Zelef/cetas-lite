---
id: collect-240926-datacamp/datacamp/les-18-principales-commandes-docker-pour-construire-executer-et-gerer-des-conteneurs-3
title: "syntax=docker/dockerfile:1"
domain: datacamp
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/clean_en/datacamp/les-18-principales-commandes-docker-pour-construire-executer-et-gerer-des-conteneurs.md
source_anchor: ""
source_lines: [289, 429]
sha256: 6112f67a650e2508e63adc1b4ecffb82d02133e3d2e28d169f77f7e994ab4b46
---

# syntax=docker/dockerfile:1

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
