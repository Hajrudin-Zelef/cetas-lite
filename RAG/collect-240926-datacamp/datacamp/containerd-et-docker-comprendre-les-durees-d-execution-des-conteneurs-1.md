---
id: collect-240926-datacamp/datacamp/containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs-1
title: "containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs"
domain: datacamp
role: reference
task: reference
actors: ["AWS", "Google"]
dates: []
keywords: ["aws"]
source: docs/RAG/clean_en/datacamp/containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs.md
source_anchor: ""
source_lines: [1, 95]
sha256: 83ff41cc171b24d6b5c4bda670bb6e61154682147c85d9708bb40171d69501dc
---

# containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs

<!-- source: https://www.datacamp.com/fr/blog/containerd-vs-docker -->

Course

In the container ecosystem, a frequent source of confusion arises when developers compare Docker (a complete platform) to containerd, which is actually a specialized runtime component. This comparison is like comparing a complete automobile to its engine: both are essential, but they serve different functions at different levels of abstraction.

In this guide, I will clarify the relationship between these two technologies, exploring their architectures, Kubernetes integration, and the specific scenarios in which each tool excels. Whether you are building containers locally or managing production Kubernetes clusters, understanding when to use Docker rather than containerd can have a significant impact on your infrastructure decisions.

By the end of this article, you will have a clear understanding of how Docker and containerd complement each other, and you will know how to choose the right tool for your specific needs, whether it's rapid local development or large-scale production deployments.

If you are new to containerization, I highly recommend taking our course on containerization and virtualization concepts.

## What is Docker?

Docker revolutionized the software development landscape by making containerization accessible and practical for everyday developers. As a complete platform, Docker provides all the elements needed to build, deliver, and run containerized applications.

For beginners, this practical guide to containers is an excellent introduction to Docker and containers.

### A Complete Container Platform

Docker works as an all-in-one container platform, offering an integrated tool stack that covers most of the container lifecycle. At its core, Docker packages applications and all their dependencies into portable, self-contained images that can run consistently in any environment, from a developer's laptop to production servers.

This comprehensive approach has made Docker the industry standard for local development and developer experience. Key benefits include:

- **Extensive ecosystem:** Docker Hub hosts millions of pre-built images, from databases to web servers, eliminating complex installation procedures.
- **Universal portability:** A container created on macOS runs identically on Linux or Windows, as long as Docker is installed.
- **Developer-first:** Abstracts infrastructure differences so developers can focus on their applications rather than deployment complexities.
- **Rapid onboarding:** New team members can set up development environments in minutes rather than hours.

### Key Components and Workflow

Docker's architecture consists of several interconnected components that work together seamlessly.

The Docker CLI provides the user interface from which developers run commands. When you run a command such as ` `docker run``, the CLI communicates with the Docker daemon (``dockerd``), which serves as the central service for managing containers, images, networks, and volumes.

For Windows and macOS users, Docker Desktop offers an additional layer of convenience through a graphical interface, making container management accessible even to those who are less comfortable with command-line tools. This integrated environment includes everything needed for local development, from Kubernetes support to volume management.

The typical workflow follows a clear pattern:

1. Developers create images using Dockerfiles.
2. They ship them to registries such as Docker Hub.
3. Finally, the images are run as containers on any Docker-compatible host.

What many developers don't realize is that Docker doesn't directly run containers. `dockerd` delegates the actual container execution to containerd, which runs in the background as a separate process.

This modular design, in which Docker uses containerd for low-level operations, allows each component to focus on what it does best: Docker provides a developer-friendly interface and ecosystem, while containerd handles the technical details of container execution.

Whether you are new to using Docker or want to take the next step, we invite you to check out our 10 Docker project ideas suitable for all levels.

## What is Containerd?

Having examined Docker's comprehensive platform approach, we will now look at containerd, the specialized runtime that enables container execution in both Docker and the broader ecosystem.

### An Industry-Standard Container Runtime

Containerd is a lightweight, industry-standard container runtime that has been certified by the Cloud Native Computing Foundation (CNCF), which attests to its maturity, reliability, and widespread adoption. Originally integrated into Docker, containerd was extracted in 2017 and transferred to the CNCF to enable broader adoption within the ecosystem.

Unlike Docker's full feature set, containerd's scope is deliberately limited: it handles essential container lifecycle operations such as starting, stopping, pausing, and deleting, and supports image transfer and storage. This focused approach makes containerd exceptionally stable and efficient—qualities that are essential for production environments.

Think of containerd as "plumbing": infrastructure designed to be embedded in larger systems rather than used directly by users. Major platforms such as Kubernetes, AWS Fargate, and Google Kubernetes Engine all rely on containerd to run containers, even though users interact with these platforms through their own interfaces.

### Architecture and design

To understand why containerd is so widely adopted in production systems, it is necessary to examine its architectural principles.

Containerd's architecture perfectly illustrates the principles of modular design. Built according to Open Container Initiative (OCI) standards, containerd ensures compatibility across the entire container ecosystem. This means that any OCI-compliant image will work with containerd, regardless of the tool used to create it.

The architecture consists of several main layers:

- **gRPC API layer:** Provides a client-server model in which multiple clients (Docker, Kubernetes, custom tools) communicate with a single containerd instance.
- **Containerd daemon:** Manages the state of running containers, administers image storage via snapshotters, and coordinates networking via plugins.
- **runc runtime:** A lightweight, OCI-compliant runtime that interfaces directly with Linux kernel features, creating namespaces, configuring cgroups, and launching container processes.
- **Modular plugins:** Custom snapshotters for specialized storage, alternative runtimes (gVisor, Kata Containers), and network plugins can be integrated without modifying containerd.

When containerd needs to start a container, it launches runc, which does the actual work of creating isolated namespaces, configuring cgroups for resource limits, and launching the container process.

This separation of concerns makes containerd highly extensible. Organizations can customize almost every aspect without modifying the core container code.

## Containerd and Docker: key differences

Now that we have a good understanding of how each tool works, let's look at their differences in practice, from architecture to integration models.

### Container runtime vs. platform

The fundamental difference lies in scope and purpose.

Docker uses a centralized daemon architecture in which `dockerd` coordinates many different aspects, such as:

- Image building
- Container execution
- Networking
- Volume management

This all-in-one design simplifies the developer experience but introduces overhead due to the many layers of abstraction.

