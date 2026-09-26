---
id: collect-240926-datacamp/datacamp/les-26-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-concernant-d-1
title: "Step 1: Choose a base image"
domain: datacamp
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/clean_en/datacamp/les-26-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-concernant-docker-pour.md
source_anchor: ""
source_lines: [1, 141]
sha256: 9fdcfbb5c157091bef836c778214d7a04d71234c6107274405a66f42288b32b1
---

# Step 1: Choose a base image

<!-- source: https://www.datacamp.com/fr/blog/docker-interview-questions -->

Course

Docker has become the most popular containerization tool in modern software development, particularly in DevOps and CI/CD workflows. It simplifies application deployment and management through containers, enabling software delivery quickly and consistently.

Its scalability and flexibility make Docker an indispensable tool for data-related roles, such as data engineering, MLOps, and even data science. That's why I compiled the frequently asked questions in Docker-related job interviews, covering fundamental concepts and concrete scenarios.

## Become a data engineer

## What is Docker?

Docker is a container platform that developers use to package applications with all their dependencies so they can run properly in different environments.

Although containers share the same operating system kernel, each runs in its own isolated environment. This setup minimizes compatibility issues, reduces delays, and improves communication between development, testing, and operations teams.

*Docker logo. Image source*

In 2023, Docker was the leader in the containerization market with over 32% market share. This highlights its importance in modern software development. That's why you can expect recruiters to assess your Docker expertise in data-related job interviews.

## Basic Docker interview questions

First, familiarize yourself with some fundamental Docker concepts. These foundational questions will help you deepen your understanding and prepare for the initial phase of the interview.

### 1. What is a Docker image?

A Docker image is comparable to a blueprint that allows you to create containers. It contains all the elements a developer needs to run an application, including:

- Code
- Libraries
- Settings

When you use a Docker image, Docker turns it into a container, which is a fully isolated environment. This is where the application runs autonomously.

### 2. What is a Docker host?

A Docker host is the system on which we install Docker. It serves as the main environment for running and managing Docker containers. We can set up a Docker host on a local device or in a virtual or cloud environment.

### 3. How does a Docker client differ from a Docker daemon? Could you give us an example?

The Docker client and the Docker daemon work in parallel but have distinct roles. The Docker client is the tool that sends commands, and the Docker daemon is the engine that executes those commands.

For example, if we enter the `docker run` command to start a container, the client will take the request and send it to the Docker daemon. The Docker daemon will then handle the actual work by starting the container.

### 4. Could you explain what Docker networking is and which commands are used to create a bridge and an overlay network?

Docker networking allows containers to connect and communicate with other containers and hosts. The ` `docker network create` ` command allows us to configure user-defined networks.

- **Bridge network**: Creates a local network for communication between containers on the same Docker host.
- Command: `docker network create -d bridge my-bridge-network`
- This sets up a bridge network called `my-bridge-network` for containers on the same host.
- **Overlay network**: Enables communication between containers across multiple Docker hosts, often used in a Swarm setup.
- Command: `docker network create --scope=swarm --attachable -d overlay my-multihost-network`
- This creates an attachable overlay network called "`my-multihost-network`" for containers running on different hosts in a Docker Swarm.

### 5. Please explain how Docker bridge networking works.

The bridge network is the default configuration used by Docker to connect containers. If we do not specify a network, Docker connects it to the bridge network. This bridge connects all containers on the same Docker host. Each container has a unique IP address, which allows containers to communicate directly with each other.

## Intermediate Docker interview questions

These questions are asked to assess your knowledge of intermediate-level Docker concepts.

### 6. What is a Dockerfile? Please explain how you would write it.

A Dockerfile is a script that defines the instructions for creating a Docker image. Each command in the Dockerfile configures a specific part of the environment. When we run these commands, Docker builds an image layer by layer. Here is how we can write it:

1. First, please select a base image. It contains the essential tools for the application.
2. Next, please define a working directory inside the container. This is where the application files will be stored and executed.
3. In the third step, please use the `COPY . .` command to copy all project files into the container's working directory.
4. Please use the "`RUN`" command to install dependencies.
5. Please use the ` `EXPOSE` ` command to indicate the port on which your application runs.
6. Please now define the command that Docker should run when it starts the container.

Here is a simple example of a Dockerfile for a Python web application:

```
# Step 1: Choose a base image
FROM python:3.9-slim
# Step 2: Specify the working directory
WORKDIR /app
# Step 3: Copy project files into the container
COPY . .
# Step 4: Install dependencies
RUN pip install -r requirements.txt
# Step 5: Expose the port the app runs on
EXPOSE 5000
# Step 6: Define the default command
CMD ["python", "app.py"]
```
Using the Dockerfile above, you can create an image with the command ` `docker build -t my-python-app .` ` and run a container with the command ` `docker run -p 5000:5000 my-python-app``.

### 7. What is Docker Compose and how does it differ from Dockerfile?

Docker Compose is a tool for defining and managing multi-container Docker applications using a YAML file (`docker-compose.yml`). It allows us to configure services, networks, and volumes in a single file, which makes managing complex applications easier.

Differences from the Dockerfile:

- A Dockerfile is used to create a single Docker image by defining its layers and dependencies.
- Docker Compose is used to run and orchestrate multiple containers that may depend on each other (for example, a web application container and a database container).

For example, a `docker-compose.yml` file might look like this:

```
version: '3.9'
services:
  web:
    build: .
    ports:
      - "5000:5000"
    depends_on:
      - db
  db:
    image: postgres
    volumes:
      - db-data:/var/lib/postgresql/data
volumes:
  db-data:
```
This file defines two services, `web` and `db`, with network and volume configurations.

### 8. Why do we use volumes in Docker?

We use Docker volumes to ensure data safety outside Docker containers. They provide a separate location on the hosts where data is preserved even if the container is deleted. In addition, it is easier to manage, back up, and share volumes between containers.

### 9. What are Docker bind mounts and why do we prefer volumes over bind mounts?

With Docker bind mounts, it is possible to share files between the host machine and a container. They link a specific file on the host system to a location in the container. If we make changes to the files, they will appear immediately in the container.

Docker mounts are suitable for real-time file sharing, but they depend on the host operating system, which raises security concerns.

On the contrary, because Docker volumes work independently, they are more secure than mounts.

*Diagram of Docker Bind mounts and volumes. Image source: Docker*

### 10. What is Docker Swarm?

Docker Swarm is a container orchestration tool that manages and deploys services across a cluster of Docker nodes. It offers high availability, scalability, and load balancing, allowing multiple hosts to act as a single virtual Docker engine.

### 11. Is it possible to set up automatic scaling of Docker Swarm?

