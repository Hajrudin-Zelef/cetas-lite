---
id: collect-240926-datacamp/datacamp/docker-vs-podman-quel-outil-de-conteneurisation-vous-convient-le-mieux-2
title: "docker-vs-podman-quel-outil-de-conteneurisation-vous-convient-le-mieux"
domain: datacamp
role: reference
task: reference
actors: ["AWS", "Google"]
dates: []
keywords: ["aws"]
source: docs/RAG/clean_en/datacamp/docker-vs-podman-quel-outil-de-conteneurisation-vous-convient-le-mieux.md
source_anchor: ""
source_lines: [89, 187]
sha256: 80b2936328d31e139b79024b71e1fdaf09211cbb53692971b6bcc551438d54de
---

# docker-vs-podman-quel-outil-de-conteneurisation-vous-convient-le-mieux

If you want to use a custom plugin, you can install it from Podman Desktop.

If you are both a Docker and Podman user, you may be surprised to see Docker and Podman objects in the Podman Desktop user interface. This is not a coincidence! We will examine what this implies in more detail. In the meantime, it means that users can interact with both their Podman AND Docker objects, all through a single pane.

The most common use case for containers is running them via Kubernetes. Despite Docker's title as the industry standard for containerization, Podman offers a more robust Kubernetes experience on Podman Desktop.

The ability to view and manage Kubernetes resources such as nodes, pods, deployments (and much more) makes Kubernetes administration and development a first-class citizen on Podman Desktop. These tools, along with plugins like the Red Hat OpenShift integration mentioned above, set Podman apart as a tool oriented toward Kubernetes workshops.

## Podman Compose vs Docker Compose

### Defining and managing multi-container applications with Docker

Some applications and workloads can be grouped into a single container. Some cannot. To facilitate managing multiple containers, Docker offers a tool called Docker Compose. Docker Compose uses a single YAML file to define the components of your application.

Then, using the docker-compose CLI, these containers and services can be started, stopped, or rebuilt. A Docker Compose YAML file might look like this:

```yaml
version: '3'
services:
	app:
		image: python:3.10
		container_name: app
		command: run app --host=0.0.0.0
	database:
		image: postgres:13
		container_name: database
		ports: 5432
		volumes:
			- postgres_data:/var/lib/postgresql/data
volumes:
	postgres_data
```
A lot is happening, but what Docker Compose allows us to do is define a YAML file with two services and a volume. Then, the docker-compose up command will launch these containers and we will have a running application.

For software and data teams managing large applications and workloads, Docker Compose makes local development easier, as well as shipping and running code in a production environment.

### Podman's approach to multi-container applications

Running multi-container applications with Podman looks almost identical to Docker. Podman does this using Podman Compose. As with Docker Compose, Podman Compose uses YAML files to define the components of an application declaratively.

Podman-compose can then be used to start, stop, or restart the services defined in the YAML file.

In most cases, podman-compose can be used in place of docker-compose (there are a few incompatibilities here and there). As with Docker, using Podman Compose allows you to manage multi-container applications independently and flexibly.

Below is a table comparing Docker and Podman.

| **Feature/aspect** | **Docker** | **Podman** | 
| **Architecture** | Docker relies on a daemon as a core architectural component. | Daemonless architecture. | 
| **Security** | Requires root privileges to build, run, and manage containers. | The daemonless nature of Podman's architecture makes it a more security-conscious container management tool. | 
| **User tool** | Docker Desktop, docker CLI | Podman Desktop, podman CLI | 
| **Compatibility** | Windows, Mac, Linux | Native to Linux, available for Windows and Mac. | 
| **Adoption** | Industry standard for container orchestration with a massive community and near-universal compatibility. | Alternative to Docker with a smaller but growing community. | 

## Use cases and best scenarios for Podman vs Docker

Let's now look at the key question you may be asking yourself: when to use Docker and when to use Podman? Let's take a closer look.

### When to use Docker?

Docker is the de facto standard for building, running, and shipping containers. If you are new to containerization (especially on your personal machine), try using Docker.

It is easy to set up your first (or fiftieth) container and get it running using tools such as Docker Desktop or Docker CLI. Docker has a massive community, and there is a good chance that what you are trying to do has already been done. This makes troubleshooting easier.

Docker offers greater cross-platform consistency than Podman. More importantly, Docker integrates with almost every container-based service, including AWS ECS, Azure AKS, and Google Cloud Run.

This means that when the time comes to run your containers in production, you are able to easily integrate the service of your choice. The ability to move from local development to production is one of the most powerful aspects of containerizing your code with Docker.

Software and data engineering teams are not the only ones using Docker. AI and ML engineers, data scientists, and even data analysts use Docker to improve their work!

### When to use Podman

For developers working in a security-sensitive or highly regulated environment, Podman may be your container manager of the day. Remember that Podman is rootless, which means that a user running Podman locally does not need root access to their machine to build and manage containers locally. Below are a few other reasons why it may be wise to choose Podman over Docker.

- You are developing locally on a Linux machine.
- The use of underlying resources and container startup time are important to you.
- You plan to ship your containers to a Kubernetes cluster or you want to mimic a Kubernetes environment on your local machine.

Here is the other thing to keep in mind: for the most part, Podman and Docker are interchangeable. This means that if you start using Docker and realize that Podman is the tool you need, it is easy to switch from one offering to the other.

## Conclusion

Operating containers requires a tool to manage these objects. Together, we explored two of the most popular containerization tools: Docker and Podman.

The industry standard for containerization, Docker is used by millions of people to run applications and data workloads around the world. Docker's architecture relies on the Docker daemon, which requires root access to the system on which the container runs.

To interact with Docker, developers can use the Docker CLI or Docker Desktop, both of which allow you to manage items such as images, containers, and volumes. Docker's widespread adoption results in a large and vibrant community, as well as support for the three main operating systems and almost all container-based services.

Podman offers an alternative container management solution. Podman is daemonless and rootless, meaning a user doesn't need root access to the machine they're using to run Podman. This is interesting for teams looking for a more security-conscious container management tool. Like Docker, Podman offers both a CLI and a user interface for building and managing containers. Although native to Linux, Podman can run on Windows and Mac and integrates quite well with tools like AWS ECS and Azure AKS.

Whichever tool you choose, learning to "containerize" the code you write is one of the fastest ways to develop your development skills. If you'd like to learn more about Docker and Podman, feel free to get hands-on with courses like Introduction to Docker or Containerization and Virtualization Concepts. Good luck and happy coding!

## Podman vs Docker FAQ

### What is a container?

**A container is an object that contains everything needed to run an application or workload. You can think of containers as small computers that only have the essentials to run some kind of code. Thankfully, we can run these containers on our local machines as well as on servers that make a solution accessible to the whole world.**

### Why use a container in my project?

