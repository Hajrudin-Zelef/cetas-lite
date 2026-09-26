---
id: collect-240926-datacamp/datacamp/principales-alternatives-a-docker-en-2026-un-guide-complet-3
title: "Buildah scripting approach with CI integration"
domain: datacamp
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["gpus", "latency"]
source: docs/RAG/clean_en/datacamp/principales-alternatives-a-docker-en-2026-un-guide-complet.md
source_anchor: ""
source_lines: [167, 230]
sha256: 8a5847c7d89f00d8e8efcc5615b21077e795a998d2291dd6bbccda7d33b122d3
---

# Buildah scripting approach with CI integration

This approach is particularly effective for existing workloads that were not designed for containerization. Applications that expect to write to the home directory (`/etc`), run system services, or interact with the entire filesystem hierarchy work seamlessly in LXC containers. It is possible to migrate entire server configurations to LXC without having to refactor applications for microservices architectures.

LXC containers share the host kernel but offer stronger isolation than application containers. Each container has its own network stack, its own process tree, and its own filesystem namespace, creating isolation similar to that of a virtual machine without the overhead of hardware virtualization.

LXD, now available under Canonical, adds a powerful management layer on top of LXC. LXD provides REST APIs, image management, and advanced features such as live migration between hosts. You can move running containers from one physical machine to another without service interruption, similar to VMware vMotion, but with containers.

Hardware passthrough capabilities allow LXD containers to directly access GPUs, USB devices, and other hardware. This makes it suitable for workloads that require specialized hardware access while retaining the benefits of containers, such as density and rapid provisioning.

LXD also supports clustering, allowing you to manage multiple hosts as a single logical unit with automated placement and failover capabilities.

### Comparative Analysis

Here is a comparison between system containers and application containers based on their main characteristics:

Image 6 - Overview of the main characteristics of Docker, LXC, and virtual machines

System containers bridge the gap between lightweight application containers and heavy virtual machines. They are particularly suitable when you need virtual machine-like capabilities with container efficiency, or when you are migrating existing applications that cannot be easily decomposed into microservices.

The choice between system containers and application containers depends more on your workload characteristics and operational requirements than on technical superiority. They address different needs in the field of containerization.

## Security Architectures in Modern Runtime Environments

Security has moved from a secondary consideration to a fundamental design principle in modern container runtime environments. Current platforms implement defense-in-depth strategies that assume containers will be compromised and focus on limiting the blast radius.

### Rootless Operation

Rootless containers eliminate the primary security risk of traditional Docker deployments: the root daemon. Podman was the first to adopt this approach by running containers entirely under user privileges, using Linux user namespaces to map the container's root user to an unprivileged user ID on the host system.

The implementation relies on subordinate user and group ID ranges (`/etc/subuid` and `/etc/subgid`) that allow unprivileged users to create isolated namespaces. When a container process thinks it is running as root (UID 0), the kernel maps it to your actual user ID (for example, UID 1000) on the host. Even if an attacker manages to escape the container, they cannot exceed your user's permissions.

Containerd has implemented similar rootless capabilities through its rootless mode, which uses the same user namespace mapping techniques. The runtime can start containers, manage images, and handle networking without requiring root privileges on the host system.

Kubernetes now natively supports rootless operation through the Kubernetes-in-Rootless-Docker (KIND) project and rootless containerd integration. This means entire Kubernetes clusters can run without root privileges, significantly reducing the attack surface for multi-tenant environments and edge deployments where traditional security models do not apply.

The security impact goes beyond preventing privilege escalation. Rootless containers cannot bind to privileged ports (below 1024), do not have access to most `/proc` and `/sys` filesystems, and cannot perform operations requiring kernel capabilities. This creates natural boundaries that help contain potential security breaches.

### Seccomp/BPF Integration

Modern runtime environments integrate eBPF (extended Berkeley Packet Filter) for real-time enforcement of security policies that go beyond traditional access controls. eBPF programs run in kernel space and can monitor, filter, or modify system calls as they occur, providing unprecedented visibility and control over container behavior.

Seccomp (Secure Computing) profiles use BPF to filter system calls at the kernel level. Instead of allowing containers access to all of the more than 300 Linux system calls, seccomp profiles precisely define which calls are permitted. Docker's default seccomp profile blocks potentially dangerous system calls, while custom profiles can be even more restrictive depending on application requirements.

Advanced eBPF integration enables real-time behavioral monitoring. Tools such as Falco use eBPF programs to detect anomalous container behavior, such as unusual network connections, unexpected file access patterns, or attempts to use blocked system calls. These detections occur in real time with minimal performance overhead, since monitoring takes place in kernel space.

Network policy enforcement via eBPF enables granular traffic control at the packet level. Cilium, a highly regarded Kubernetes CNI, uses eBPF to implement network policies capable of filtering traffic based on application-layer protocols, not just IP addresses and ports. This means you can create policies such as "allow HTTP GET requests to `/api/v1/users` but block POST requests" directly in the kernel.

eBPF-based security also enables container-aware monitoring, which understands the relationship between processes, containers, and Kubernetes pods. Traditional monitoring tools examine individual processes, while eBPF programs can correlate system calls with container metadata to provide contextual security information.

These capabilities transform security from a reactive remediation approach to a proactive policy enforcement approach, where suspicious behaviors are automatically blocked before they can cause harm.

## Performance Optimization Strategies

Container performance is especially important when you run hundreds or even thousands of containers on your infrastructure. Let's explore strategies that minimize resource overhead and startup latency.

### Reducing Cold Starts

Cold start time, that is, the delay between requesting a container and its availability to handle traffic, has a direct impact on user experience and resource efficiency. Techniques have been developed to minimize these delays across different runtime architectures.

Pre-pulled images eliminate download time by keeping frequently used container images cached on nodes. Kubernetes DaemonSets can pre-pull critical images, while registries such as Harbor support image replication to edge locations. This technique can reduce cold start time from several seconds to a few milliseconds for cached images.

Image layer optimization reduces the amount of data to transfer and extract. Multi-stage builds produce smaller final images, while tools like dive help identify unnecessary layers. Google's Distroless images eliminate package managers and shells, which often significantly reduces image size.

Lazy loading with projects like Stargz allows containers to start before the entire image is downloaded. The runtime fetches only the files needed for initial startup, downloading additional layers on demand. This can reduce cold start time from several seconds to under a second for large images.

