---
id: collect-240926-datacamp/datacamp/docker-vs-podman-quel-outil-de-conteneurisation-vous-convient-le-mieux-1
title: "docker-vs-podman-quel-outil-de-conteneurisation-vous-convient-le-mieux"
domain: datacamp
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws"]
source: docs/RAG/clean_en/datacamp/docker-vs-podman-quel-outil-de-conteneurisation-vous-convient-le-mieux.md
source_anchor: ""
source_lines: [1, 88]
sha256: 743504f07062d7426713e0b9772757106803c86a7d0a82599ee70b751af1d39a
---

# docker-vs-podman-quel-outil-de-conteneurisation-vous-convient-le-mieux

<!-- source: https://www.datacamp.com/fr/blog/docker-vs-podman -->

Course

Containers run applications and data workloads all over the world. First designed in the 1970s, containers allow you to bundle everything needed to run an application or workload into a single object. Containers solve the "it runs on my machine" problem by providing an isolated and portable solution for developing, testing, and shipping code. Tools like Kubernetes rely heavily on containers as a central element of their architecture. For now, containers aren't going anywhere.

To run these containers, you'll need a container management solution. Enter Docker and Podman.

Docker and Podman are used to build, manage, and deploy containers. Together, we'll analyze the similarities and differences between Docker and Podman, as well as the unique features of each. We'll explore topics such as daemon and daemonless architecture, multi-container management, and cross-platform integration. By the end, you'll be armed with the information you need to choose the perfect container solution for your needs.

**If you're not yet familiar with these tools, you can also check out our Introduction to Docker and our Introduction to Podman tutorial for machine learning.**

## Become a Data Engineer

## What are Podman and Docker?

Let's start with an overview of these tools to begin our comparison:

### Docker Overview

Docker is the de facto standard for building, running, and shipping containers. Containers are objects that combine operating system-level dependencies and some kind of application code to package and run things like complete applications or ETL pipelines in their own isolated environment. Containers are like small computers that only have the essentials to run some kind of code.

Docker is quite young and was first released as an open-source project in 2013. Since then, the project has exploded.

When it comes to running containers in an enterprise, almost every software and data team uses Docker.

Developers can use Docker on all three major operating systems, and it integrates seamlessly with almost all modern technologies. This means a data engineer can write and package a data pipeline using a Docker container on their local Mac, and ship that container to run on AWS ECS.

Tools like Docker CLI, Docker Desktop, and Docker Hub allow developers of all skill levels to get started easily.

If you're looking for a more convenient way to learn Docker, we have several Docker projects and information on Docker certifications to help you improve your Docker skills!

### Podman Overview

Like Docker, Podman is an open-source tool for developing and managing containers. Podman was originally developed by Red Hat as a Linux-native alternative to Docker and was released in 2019.

Notably, the underlying architecture of the two container runtimes differs; while Docker uses daemons, Podman operates daemonless (more on this later).

Unlike Docker, Podman does not require root access to the machine on which the pods it manages run, making it a more security-conscious option for teams using containers to run their applications and workloads.

Podman users get a user experience similar to Docker's; developers can use a CLI or a graphical interface (Podman Desktop) to interact with Podman in their local environment.

Users on Linux, Mac, and Windows can use Podman to build and test their containers locally before deploying them to some kind of remote environment, such as Kubernetes.

## Key Differences Between Podman and Docker

### Daemon vs. Daemonless Architecture

The biggest difference between Docker and Podman is the underlying architecture they are built on. Docker relies heavily on a daemon, while Podman is daemonless.

You can think of a daemon as a process that runs in the background on the host operating system. In Docker's case, its daemon is responsible for managing Docker objects (images and containers) and communicating with other systems. To run its daemon, Docker uses a package called dockerd.

Why does this matter? First, daemons generally require root-level access to the machine they run on. This situation is conducive to security vulnerabilities: if a malicious actor manages to access a daemon, they now have access to the entire machine.

Podman's daemonless architecture has a few advantages. Since running daemons almost always requires administrator privileges, a daemonless architecture can be considered "rootless." This means users who don't have system-level access to the machine their containers run on can still use Podman, which isn't always the case with Docker.

Instead of a daemon, Podman uses a Linux package known as systemd. Because systemd is an integral part of the Linux operating system, Podman is often considered more "lightweight" than Docker; Podman users typically experience faster container startup times than when using Docker.

### Building images and containers

Despite their fundamentally different architectures, Docker and Podman share the same primary goal: creating and running images and containers. However, their approaches to this process differ slightly.

With Docker, an image is built by first adding commands to a Dockerfile. Then, a command such as docker build is executed. This operation calls each of the instructions in the Docker file, creating an image. An image can then be "run" as a container. As you may have guessed, this is done using the docker run command, specifying an image ID or tag. To build and run multiple containers, we will use a special tool called docker-compose, which we will explore a bit further on.

The process of building images and running them as containers is almost identical in Podman. Rather than a Docker file (although this filename still works), Podman users will create a Container file. The image composition syntax is the same. Once the appropriate commands are added to the container file, the image can be built and run using the Podman API.

For the most part, Podman is compatible with most Docker elements. You will find differences here and there, but for the most part, the Docker API can be replaced by the Podman API without issue.

## Podman vs Docker Desktop

### Docker Desktop for simplified cross-platform access

There are several ways to work with Docker. Experienced software and data practitioners generally rely on the Docker CLI (AKA the Docker client) to interact with their Docker images and containers.

However, there is an even simpler way to get started, and that is Docker Desktop.

Docker Desktop is a free GUI-based tool that allows users to create and manage the images and containers that run their applications or workloads. A data engineer can use Docker Desktop to view the images available on their machine and turn that image into a container. Similarly, a software developer can download an image from Docker Hub to use as part of their next project.

The user interface is simple and intuitive, while maintaining complete visibility and control over your Docker environment.

However, Docker Desktop is not limited to viewing and managing Docker objects.

Users can notably manage (down to the byte) the resources available to their Docker objects, attach to a running container, or launch a Kubernetes cluster on their local machine. Docker Desktop users can choose from hundreds of extensions, or get started with Docker through helpful tutorials and sample environments. Fortunately for you, Docker Desktop is widely accessible and works on Mac, Windows, or Linux.

### Features and limitations of Podman Desktop

Podman Desktop looks very much like its Docker counterpart. From the Podman Desktop user interface, users can view and manage containers, images, pods, and volumes. As with Docker, Podman supports plugins and integrations that allow you to run a Red Hat OpenShift cluster locally or work with LLMs using the Podman AI Lab.

