---
id: collect-240926-datacamp/datacamp/containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs-2
title: "containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs"
domain: datacamp
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "latency", "memory"]
source: docs/RAG/clean_en/datacamp/containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs.md
source_anchor: ""
source_lines: [96, 178]
sha256: 432c78874c12b8ba1a51ff33c6ca9891318eebde02f021f217ed8586ca11a26c
---

# containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs

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

