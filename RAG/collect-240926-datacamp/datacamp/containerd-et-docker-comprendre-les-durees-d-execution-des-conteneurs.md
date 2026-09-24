---
id: collect-240926-datacamp/datacamp/containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs
title: "containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs"
domain: datacamp
role: reference
task: reference
actors: ["AWS", "Google"]
dates: []
keywords: ["agent", "attention", "aws", "latency", "memory", "training"]
source: docs/RAG/clean_en/datacamp/containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs.md
source_anchor: ""
source_lines: [1, 307]
sha256: 2210cac76205ca74b0e1c39e4daabd67cc41e4a680a743cade2e15987392f2fa
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

Containerd, by contrast, focuses exclusively on running containers. It does not include image building, orchestration, or graphical interface features. For networking, Docker integrates a container network interface (`Libnetwork`) into its daemon, while containerd relies on external Container Network Interface (CNI) plugins that can be swapped out as needed.

This architectural difference has performance implications. In high-churn environments where containers start and stop frequently, such as autoscaling Kubernetes clusters, containerd's streamlined design can result in faster startup times and reduced resource consumption. Lower overhead means containerd uses less memory and CPU, which becomes significant at scale.

### Kubernetes integration

When it comes to Kubernetes, the relationship between container runtimes and orchestration platforms has evolved considerably, particularly regarding how Kubernetes connects to containerd.

The relationship between Kubernetes and container runtimes underwent a major change with Kubernetes version 1.24, released in 2022. This version removed "Dockershim," a compatibility layer that allowed Kubernetes to use Docker as a container runtime.

Dockershim was always intended as a temporary solution. When Kubernetes introduced the Container Runtime Interface (CRI) to standardize communication with runtimes, Docker could not implement it directly because Docker predated CRI's design. Dockershim provided translation between Kubernetes and Docker, adding a translation layer that could prove superfluous.

Modern versions of Kubernetes communicate directly with containerd via CRI, completely eliminating the Docker translation layer. This simplification brings concrete benefits:

- **Reduced latency:** Direct communication eliminates translation overhead in container operations.
- **Improved stability:** Fewer moving parts means fewer potential points of failure.
- **Better performance:** The streamlined execution stack is particularly advantageous for large-scale deployments.
- **Simplified debugging:** Direct CRI integration simplifies troubleshooting.

For Kubernetes users, this change is largely transparent. Docker images remain fully compatible because they comply with OCI standards. In practice, this means production Kubernetes clusters now run more efficiently by using containerd directly, while developers can continue using Docker locally for building and testing.

If you're not sure about the pros and cons of using Kubernetes, check out this comparison between Docker Compose and Kubernetes.

### Image building and management

While runtime integration is essential for orchestration, developer workflows depend heavily on how each tool handles image creation and storage. Image building represents a significant capability gap between Docker and containerd.

Docker provides a built-in build system via Dockerfiles and BuildKit, allowing developers to create complex multi-stage builds with caching, parallelization, and advanced features such as build secrets and SSH agent forwarding.

True to its design, Containerd includes no native image-building workflow. To build images with containerd, developers must use external tools. Options include running BuildKit as a separate daemon and using `buildctl` for command-line builds, or adopting `nerdctl`, a Docker-compatible CLI that integrates BuildKit.

Storage mechanisms also differ in their approach. Docker's volume management provides an abstraction that feels natural to developers, with named volumes that retain data independently of the container lifecycle. Containerd uses a lower-level snapshot system, in which different snapshot drivers can be plugged in to handle layered file systems differently depending on the underlying storage requirements.

This difference reflects the intended audiences of these tools:

- Docker optimizes developer productivity through convenient built-in features.
- Containerd provides flexible primitives that platform developers can assemble according to their specific needs.

### CLI, developer experience, and nerdctl

Beyond architecture and features, developers' day-to-day experience is directly influenced by the command-line interface provided by each tool. The command-line experience clearly highlights the different design philosophies.

Docker's CLI is renowned for its user-friendliness. Commands such as `docker run`, `docker build`, and `docker logs` are intuitive, well documented, and designed for users. The CLI includes useful defaults, clear error messages, and extensive options that cover most use cases.

Containerd ships with `ctr`, a minimal CLI intended exclusively for debugging and testing containerd's low-level features. The `ctr` tool is intentionally not developer-friendly. It lacks common features such as port mapping shortcuts, automatic restart policies, and integration with credential helpers. It is designed for container developers, not application developers.

#### Bridging the gap with nerdctl

This usability gap led to the creation of `nerdctl`, a Docker-compatible CLI for containerd. Using `nerdctl` is similar to using Docker: same command syntax, same flags, same workflow, but with containerd as the underlying runtime. This makes `nerdctl` an excellent transitional tool for teams moving from Docker to containerd in their development environments.

In practice, developers rarely interact directly with containerd. When necessary, `nerdctl` provides the familiar interface they expect, while platform operators and administrators use containerd's APIs programmatically through orchestration systems such as Kubernetes.

To illustrate the differences in practice, here is a comparison of common container operations across the three CLI tools:

| **Task** | **Docker** | **nerdctl** | **ctr** | 
| Run container | `docker run -d -p 8080:80 nginx` | `nerdctl run -d -p 8080:80 nginx` | `ctr run --net-host -d docker.io/library/nginx:latest nginx_id` | 
| List containers | `docker ps` | `nerdctl ps` | `ctr tasks list` | 
| Build an image | `docker build -t myapp .` | `nerdctl build -t myapp .` | *Not supported* | 
| View logs | `docker logs`  | `nerdctl logs`  | *Not supported* | 
| Please inspect the container. | `docker inspect`  | `nerdctl inspect`  | `ctr containers info`  | 
| Please click on the image. | `docker pull nginx` | `nerdctl pull nginx` | `ctr images pull docker.io/library/nginx:latest` | 
| Compose support | `docker compose up` | `nerdctl compose up` | *Not supported* |

Note: ctr does not have port mapping (`-p`) and requires host networking (`--net-host`) to expose services. It does not automatically download images.

### Overview of key differences

Before we get into recommendations for specific use cases, let's recap the differences we've covered so far. The following table summarizes the key functional differences between Docker and containerd, highlighting their distinct capabilities and target audiences:

| **Feature** | **Docker** | **Containerd** |
| Image build | Built-in (Dockerfiles, BuildKit) | Requires external tools (buildctl, nerdctl) |
| Orchestration | Docker Swarm / Kubernetes | None (used by Kubernetes) |
| Storage management | Volume management | Snapshotter system |
| Graphical interface | Docker Desktop | None |
| Container lifecycle | Full management (via containerd) | Primary purpose (CRI-compatible) |
| Primary users | Application developers | Cluster operators, platform developers |

## Why choose Docker?

Now that we've covered the technical differences, let's look at practical scenarios where each tool excels. Despite the rise of containers in production environments, Docker remains the preferred choice for specific scenarios where developer experience and the availability of comprehensive tools are paramount.

### For local development and prototyping

Docker is particularly effective when you need an all-in-one solution for writing and testing code. Thanks to the integrated toolchain, developers can go from zero to running containers in minutes, without having to assemble multiple components or configure a complex network.

The Docker ecosystem offers significant productivity benefits:

- **Docker Hub:** Millions of ready-to-use images for databases, message queues, and web servers.
- **Docker Compose:** Define multi-container applications in a single YAML file and launch complete development environments using a single command.
- **Docker Desktop GUI:** Visual container management, volume exploration, resource limit adjustments, and built-in Kubernetes support.
- **Lower barriers to entry:** Graphical tools and intuitive commands make containers accessible to developers new to the technology.

For teams using Docker Desktop, the GUI offers additional features that significantly accelerate onboarding and daily workflows.

### For complex build pipelines

Beyond development, Docker's build capabilities make it a natural choice for sophisticated continuous integration and deployment workflows.

BuildKit, integrated into Docker, offers advanced features essential for modern CI/CD pipelines. Multi-stage builds reduce image size while maintaining the readability of Dockerfiles. BuildKit's caching mechanisms intelligently reuse layers from one build to the next, significantly reducing build times in continuous integration environments.

Most automation platforms, such as GitHub Actions, GitLab CI, Jenkins, and others, have proven, mature Docker integrations. These integrations handle authentication, caching, and image publishing without any difficulty.

While other tools may offer similar results, Docker's ubiquity means that solutions and troubleshooting assistance are easily accessible, which is another considerable advantage.

## Why choose Containerd?

Containerd stands out in production scenarios where minimalism, performance, and stability take precedence over the convenience of built-in tools.

### For production Kubernetes clusters

Using containerd as the runtime for Kubernetes nodes brings significant benefits:

- **Reduced overhead:** Removing the Docker daemon reduces resource consumption per node and achieves significant savings at scale.
- **Improved stability:** Fewer moving components means fewer potential points of failure in your infrastructure.
- **Reduced attack surface:** Less code to audit and fewer potential security vulnerabilities
- **Simplified debugging:** Direct CRI integration eliminates the dockershim translation layer, simplifying troubleshooting.
- **Better performance:** The optimized runtime stack improves container startup times and responsiveness.

Containerd's CNCF graduation status is a testament to its maturity and reliability. Major cloud service providers, including AWS, Google Cloud, and Azure, have adopted containerd as the standard for their managed Kubernetes offerings, demonstrating their confidence in its ability to be used in production for critical infrastructure.

### For specialized and minimalist environments

While Kubernetes is the most common production use case, the lightweight nature of containerd paves the way for deployment scenarios where Docker would be impractical.

Edge computing and IoT devices often operate under severely resource-constrained conditions. Containerd's lightweight design makes it viable in these environments where the full Docker stack would be prohibitive. Every megabyte of memory and every CPU cycle matters when running on embedded hardware.

Advanced security scenarios benefit from containerd's modular runtime architecture. Organizations can integrate sandboxed runtime environments for workloads requiring additional security boundaries. Here are a few examples:

- **gVisor:** provides efficient kernel isolation
- **Kata Containers:** runs containers in lightweight virtual machines.

These integrations fit into containerd without requiring major modifications.

## Transitioning from Docker to Containerd

It is important to understand when to use each tool, but it is just as crucial to know how to switch from one to the other. Transitioning from Docker to containerd in existing environments requires careful planning, but the process is well documented and straightforward.

### Migrating Kubernetes Nodes

The operational steps to migrate Kubernetes nodes from Docker to containerd follow a standard pattern:

1. 
**Isolate the node:** Prevent new pods from being scheduled (`kubectl cordon` )
2. 
**Drain existing pods:** Move workloads to other nodes (`kubectl drain` )
3. 
**Update the Kubelet configuration:** Point to the containerd CRI socket at the following address:`/run/containerd/containerd.sock`
4. 
**Check the CNI plugins:** Make sure the necessary network plugins are installed on containerd.
5. 
**Restart Kubelet:** Register with the new runtime and rejoin the cluster.

Testing the migration on non-production nodes helps identify environment-specific issues before rolling out the changes cluster-wide. To avoid common mistakes, follow these best practices:

- **Log path changes:** Docker and containerd use different default log locations. Update your logging infrastructure accordingly.
- **Install missing CNI plugins:** Containerd requires the CNI plugin binaries for networking; these are not always installed by default.
- **Pay attention to image pull differences:** Authentication and registry settings may need adjustments.
- **Watch out for incompatibilities between storage drivers:** Make sure your persistent volumes are compatible with containerd's snapshotter.

### Understanding the Differences Between CLI Interfaces

Once your infrastructure is migrated, developers will need to adapt their daily workflows to work with the new runtime.

For developers used to Docker commands, nerdctl offers a similar experience. Commands such as ` `nerdctl run``, ` `nerdctl build`` and ` `nerdctl compose up` ` work exactly like their Docker equivalents, making the transition seamless.

For debugging, it is essential to understand how to map Docker debugging workflows to containerd. If you use `docker inspect` to examine a container, `ctr containers info` provides similar information, but in a different format. Likewise, `ctr tasks list` shows running containers.

Most developers find that `nerdctl` eliminates the need to learn `ctr` syntax for everyday tasks. The low-level `ctr` remains useful for troubleshooting runtime-specific issues or when working directly with containerd APIs.

## Conclusion

The relationship between Docker and containerd perfectly illustrates the success of modular design in software infrastructure. Docker remains the optimal tool for developers who write code, because it offers an integrated experience, a complete ecosystem, and user-friendly interfaces that make building containerized applications productive and enjoyable.

Containerd, meanwhile, stands out as the ideal runtime for machines running code, offering the stability, performance, and minimalism required for production orchestration platforms. The fact that Docker Engine uses containerd in the background demonstrates that these two tools complement each other rather than compete.

I recommend a pragmatic approach for most organizations: continue using Docker on developers' laptops, where its tools accelerate development workflows, but consider migrating production Kubernetes clusters to containerd directly to benefit from the operational advantages of reduced overhead and simplified runtime stacks.

Whether you choose Docker's full platform or containerd's specialized runtime, both solutions remain essential components of the modern container ecosystem, each optimized for different stages of the application lifecycle.

To continue your training, we invite you to sign up for our skills course "Containerization and virtualization with Docker and Kubernetes."

## FAQ on Containerd and Docker

### Is it possible to use Docker images with containerd?

**Yes, absolutely. Containerd supports all container images that comply with the OCI standard, including those created with Docker. Since Docker creates OCI-compliant images, they work perfectly with containerd and all other OCI-compatible runtime environments. You can use `docker build` locally and run those images with containerd in production without any compatibility issues.**

### Does Docker use containerd in the background?

**Yes, Docker Engine uses containerd as its primary container runtime. Starting from version 1.11, Docker integrated containerd to manage container lifecycle operations, such as creation, execution, and management. When you run `docker run`, the Docker daemon (`dockerd`) delegates the actual execution of the container to containerd, which then uses runc to interact with the Linux kernel.**

### Why did Kubernetes remove Docker support?

**In 2022, Kubernetes removed dockershim (the Docker compatibility layer) in order to eliminate an unnecessary translation layer. Docker predates the Container Runtime Interface (CRI), so Kubernetes needed dockershim to handle the conversion between its APIs and Docker. By communicating directly with containerd via CRI, Kubernetes achieves better performance, greater stability, and a simpler runtime stack. Docker images continue to work perfectly in Kubernetes.**

### Should I switch from Docker to containerd for local development?

**No, Docker remains the most appropriate choice for local development. Docker provides an integrated toolchain with Docker Compose, the Docker Desktop graphical interface, and extensive ecosystem support that speeds up development workflows. Please use containerd for production Kubernetes clusters, where its low overhead and direct integration with CRI offer clear advantages, but keep Docker on developers' laptops to benefit from its superior user experience.**

### What is nerdctl and do I need it?

**Nerdctl is a Docker-compatible CLI for containerd that offers the same user experience as Docker (supports most common commands and flags) but uses containerd as the runtime. You need it if you want to interact directly with containerd using the usual Docker commands. It is particularly useful for development environments that use containerd or when teams are transitioning from Docker to containerd-based workflows.**

As the founder of Martin Data Solutions and a freelance Data Scientist, ML and AI engineer, I bring a diverse portfolio in regression, classification, NLP, LLM, RAG, neural networks, ensemble methods, and computer vision.

- Successfully developed several end-to-end ML projects, including data cleaning, analysis, modeling, and deployment on AWS and GCP, delivering impactful and scalable solutions.
- Created interactive and scalable web applications using Streamlit and Gradio for various use cases in the industry.
- Teaches and mentors students in data science and analytics, fostering their professional development through personalized learning approaches.
- Designed course content for retrieval-augmented generation (RAG) applications tailored to enterprise requirements.
- Wrote high-impact technical blogs on AI and ML, covering topics such as MLOps, vector databases, and LLMs, with significant engagement.

In every project I take on, I make sure to apply up-to-date software engineering and DevOps practices, such as CI/CD, code linting, formatting, model monitoring, experiment tracking, and robust error handling. I am committed to delivering comprehensive solutions, transforming data insights into practical strategies that help businesses grow and make the most of data science, machine learning, and AI.
