---
id: collect-240926-datacamp/datacamp/kubernetes-vs-docker-differences-que-tout-developpeur-doit-connaitre-1
title: "Use the official Python base image with version 3.9"
domain: datacamp
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/clean_en/datacamp/kubernetes-vs-docker-differences-que-tout-developpeur-doit-connaitre.md
source_anchor: ""
source_lines: [1, 105]
sha256: 1c6035db9b3f549a5db453529ec881d942684849aafe7e818c5362d31dfcb669
---

# Use the official Python base image with version 3.9

<!-- source: https://www.datacamp.com/fr/blog/kubernetes-vs-docker -->

Course

Containerization has become the ideal solution for efficiently building, deploying, and scaling modern applications. The two big names in this space are Kubernetes and Docker, and although they are often mentioned together, they actually have different goals. Both are essential, but they have distinct functions.

In this guide, I will help you understand what differentiates Kubernetes and Docker, their features, and when to use each of them.

## What is containerization?

Before diving into Docker and Kubernetes, let's start by understanding what it is, namely containerization.

Containerization is a lightweight form of virtualization that packages an application and its dependencies into a single unit called a container.

Unlike traditional virtual machines, containers share the host system's operating system but maintain isolation between applications. They are therefore more efficient, lighter, and faster to start!

Containerization helps developers create consistent, portable, and easy-to-manage environments, regardless of where they run, whether on a developer's laptop, in a data center, or in the cloud.

### Containerization or virtualization

It is helpful to compare containerization to traditional virtualization to better understand it. Virtual machines (VMs) virtualize entire hardware systems, meaning that each VM includes a complete operating system as well as the necessary binaries and libraries. This approach provides isolation but comes with significant resource overhead - each VM requires its own operating system, making it resource-intensive and slower to start.

Containers, on the other hand, share the host operating system's kernel, making them much lighter and faster to start. Instead of virtualizing hardware, containers virtualize the operating system. This allows containers to run isolated processes without having to manage a complete operating system for each instance, resulting in better resource utilization and greater efficiency.

While virtual machines are ideal for complete isolation and running multiple different operating systems on the same hardware, containers are better suited for efficient, scalable, and consistent application deployment.

Virtual Machines vs. Containers. Image source: contentstack.io

If you want to learn more about the essentials of VMs, containers, Docker, and Kubernetes, check out the free course on containerization and virtualization concepts on DataCamp.

Now let's dive into the details of Docker and Kubernetes!

## Become a data engineer

## What is Docker?

Docker is an open-source platform that provides a lightweight and portable way to create, deploy, and manage containers. Unlike traditional virtual machines, Docker containers package everything, including application code, the runtime engine, system tools, and libraries, allowing applications to run consistently across different environments.

### How Docker works

Docker works by creating containers that, as we saw earlier, are lightweight packages that encapsulate all the components needed to run an application.

Containers are built from Docker images, which act as a blueprint defining what is inside each container. A Docker image can include an operating system, application binaries, and configuration files, making it easy to replicate environments.

Once the image is created, developers can use Docker to run containers based on that image. One of Docker's main strengths is its simplicity and consistency: regardless of where a container is run (on a developer's local machine, in an on-premises data center, or in the cloud), the behavior remains the same.

Overview of Docker architecture. Image source: Docker Documentation

The following example provides an overview of how Docker images are implemented. Take a look at the Docker file below:

```
# Use the official Python base image with version 3.9
FROM python:3.9
# Set the working directory within the container
WORKDIR /app
# Copy the requirements file to the container
COPY requirements.txt .
# Install the dependencies
RUN pip install -r requirements.txt
# Copy the application code to the container
COPY . .
# Set the command to run the application
CMD ["python", "app.py"]
```
A `Dockerfile` is a script that contains a series of instructions allowing Docker to build an image, which can then be used to create a container.

After creating a Docker file in your project, the next step is to build the Docker image. This operation is performed using the `docker build` command, which reads the instructions contained in the `Dockerfile` document to assemble the image. 

For example, running `docker build -t my-app .` in the terminal tells Docker to build an image with the tag `my-app` from the current directory (indicated by `.`).

During the build process, Docker executes each step in the Docker file, such as pulling the base image, installing dependencies, and copying the application code into the image. Once the image is built, it serves as a template that can be reused to create multiple containers.

Once the image is successfully built, you can create and run containers from it using the `docker run` command. For example, `docker run my-app` starts a new container based on the `my-app` image, thus launching your application in the isolated environment provided by Docker.

If you want to learn more about common Docker commands and industry best practices, check out the blog Docker for Data Science: An Introduction.

### Docker Features

- Portability: Docker containers can run consistently across different systems, providing a seamless experience in development, testing, and production environments.
- Ease of use: Docker's command-line interface and comprehensive set of tools make it accessible to developers, even those unfamiliar with containerization.
- Lightweight: Docker containers share the same operating system kernel, which reduces resource overhead compared to full virtual machines.
- Fast startup time: Docker containers can be started in seconds, making them highly efficient for applications that require rapid startup and teardown.

Check out DataCamp's Docker cheat sheet, which provides an overview of all available Docker commands.

## What is Kubernetes?

Kubernetes is a powerful open-source container orchestration platform designed to manage containerized applications across clusters of machines.

Initially developed by Google, Kubernetes, commonly referred to as K8s, manages the deployment, scaling, and operations of application containers, making it an essential tool for container management at scale.

*Evolution of deployment strategies over time. Image source: Kubernetes.io*

### How Kubernetes Works

Kubernetes relies on the concept of clusters, nodes, and pods, forming a layered architecture that provides flexibility and scalability. A cluster represents the entire infrastructure, composed of multiple nodes (virtual or physical machines).

These nodes work together to host and manage containerized applications. Nodes can be either master nodes, which control and manage the cluster, or worker nodes, which run the application workloads. The master node is responsible for managing the state of the cluster, making scheduling decisions, and monitoring its health.

Each worker node runs one or more pods, which are the smallest deployable units in Kubernetes and consist of one or more containers.

Pods act as logical hosts for containers and share the same network and storage, making it easier for containers within a pod to communicate. Pods are ephemeral by nature, meaning they can be created, destroyed, or replicated dynamically based on the application's needs.

Overview of Kubernetes architecture. Image source: Kubernetes.io

