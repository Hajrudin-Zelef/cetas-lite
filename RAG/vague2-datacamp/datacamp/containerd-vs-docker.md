---
id: vague2-datacamp/datacamp/containerd-vs-docker
title: "Containerd et Docker : Comprendre les durées d'exécution des conteneurs"
domain: datacamp
role: reference
task: article
actors: ["AWS"]
dates: ["2026-09-23"]
keywords: ["agent", "aws", "latency"]
source: docs/RAG/Collect RAG Vague 2/02_datacamp/containerd-vs-docker.md
source_anchor: ""
source_lines: [1, 76]
sha256: 2f2f64157d2ae03bc5a49b4cbd58e3bd710ae44c5027a28ce8d90b3c49d116a4
---

# Containerd et Docker : Comprendre les durées d'exécution des conteneurs

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/containerd-vs-docker
- **Site** : DataCamp
- **Type** : Article (comparatif)
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This DataCamp guide clarifies the relationship between **Docker** (a complete container platform) and **containerd** (a specialized runtime component). The analogy: comparing a full car to its engine — both essential but at different abstraction levels.

**Docker** is an all-in-one container platform with an integrated toolset covering most of the container lifecycle. It packages applications and dependencies into portable, self-contained images. Key benefits: extensive ecosystem (Docker Hub with millions of prebuilt images), universal portability, developer-first design, and fast onboarding. Architecture: the CLI communicates with the Docker daemon (`dockerd`), which manages containers, images, networks, and volumes; Docker Desktop adds a GUI for Windows/macOS including Kubernetes support. Crucially, Docker does not run containers directly — `dockerd` delegates actual container execution to containerd running as a separate background process. Typical workflow: build images with Dockerfiles, push to registries, run as containers.

**Containerd** is a lightweight, industry-standard container runtime certified by the CNCF. Initially part of Docker, it was extracted in 2017 and donated to the CNCF. Its scope is deliberately limited: it manages essential lifecycle operations (start, stop, pause, delete) and handles image transfer/storage. This focused approach makes it exceptionally stable and efficient. It acts as "plumbing" embedded in larger systems — Kubernetes, AWS Fargate, and GKE all rely on containerd. Architecture (OCI-compliant): a **gRPC API layer** (client-server model); the **containerd daemon** (manages running container state, image storage via snapshotters, network coordination via plugins); the **runc runtime** (lightweight OCI runtime interfacing with Linux kernel features — namespaces, cgroups, process launch); and **modular plugins** (custom snapshotters, alternative runtimes like gVisor and Kata Containers, network plugins).

**Key differences:**
- **Scope:** Docker is a platform (image building, container execution, networking, volume management); containerd focuses exclusively on execution (no image build, orchestration, or GUI). Docker uses built-in Libnetwork; containerd uses external CNI plugins.
- **Performance:** containerd's streamlined design can yield faster startup and lower resource use, significant at scale (e.g., autoscaling Kubernetes clusters).
- **Kubernetes integration:** Kubernetes 1.24 (2022) removed **Dockershim**, the compatibility layer that let Kubernetes use Docker as a runtime. Modern Kubernetes talks directly to containerd via the **CRI (Container Runtime Interface)**, eliminating translation. Benefits: reduced latency, improved stability, better performance, simplified debugging. Docker images remain fully compatible (OCI standards).
- **Image building:** Docker has integrated BuildKit/Dockerfiles (multi-stage builds, caching, parallelism, build secrets, SSH agent forwarding); containerd has no native build workflow — use external tools like BuildKit/`buildctl` or `nerdctl`.
- **Storage:** Docker volumes (developer-friendly named volumes); containerd uses lower-level snapshotters.
- **CLI/UX:** Docker's CLI is friendly (`docker run`, `docker build`, `docker logs`); containerd ships `ctr`, a minimal CLI for debugging only (no port mapping, no auto-restart, no auth helpers). **nerdctl** bridges the gap — a Docker-compatible CLI for containerd with identical syntax.

A CLI comparison table shows `docker`, `nerdctl`, and `ctr` equivalents; `ctr` does not support build, logs, or compose, requires `--net-host`, and doesn't auto-pull images.

**Why choose Docker:** local development/prototyping (all-in-one, Docker Hub, Docker Compose, Docker Desktop GUI) and complex build pipelines (BuildKit multi-stage builds, CI/CD integrations with GitHub Actions, GitLab CI, Jenkins).

**Why choose containerd:** production Kubernetes clusters (reduced overhead, improved stability, smaller attack surface, direct CRI integration, better performance; CNCF graduation, adopted by AWS/GCP/Azure managed Kubernetes); specialized/minimalist environments (edge computing, IoT; sandboxed runtimes gVisor for kernel isolation, Kata Containers for VM-based containers).

**Migration (Docker → containerd):** cordon the node (`kubectl cordon`), drain pods (`kubectl drain`), update kubelet config to point to the containerd CRI socket `/run/containerd/containerd.sock`, verify CNI plugins, restart kubelet. Pitfalls: log path differences, missing CNI plugin binaries, image pull/registry auth differences, snapshotter compatibility. For daily work, `nerdctl` mirrors Docker commands.

**Conclusion:** Docker is best for developers writing code; containerd is ideal for machines running code. The fact that Docker Engine uses containerd proves they complement rather than compete. Recommended approach: keep Docker on developer laptops, migrate production Kubernetes clusters to containerd.

## Key points

- Docker is a full platform; containerd is a specialized OCI-compliant runtime.
- Docker delegates actual container execution to containerd (since Docker Engine 1.11).
- containerd is CNCF-graduated and powers Kubernetes, AWS Fargate, and GKE.
- Kubernetes 1.24 removed Dockershim; modern Kubernetes uses CRI directly with containerd.
- Docker images remain compatible because they follow OCI standards.
- containerd has no native image build; use BuildKit/`buildctl` or `nerdctl`.
- `ctr` is a minimal debugging CLI; `nerdctl` provides a Docker-compatible experience.
- Recommendation: Docker for local dev, containerd for production Kubernetes.

## Technical data / figures

| Feature | Docker | Containerd |
|---|---|---|
| Image building | Integrated (Dockerfiles, BuildKit) | External tools (buildctl, nerdctl) |
| Orchestration | Docker Swarm / Kubernetes | None (used by Kubernetes) |
| Storage | Volume management | Snapshotters |
| GUI | Docker Desktop | None |
| Container lifecycle | Full management (via containerd) | Core focus (CRI-compatible) |
| Primary users | Application developers | Cluster operators, platform developers |

CLI comparison examples:

| Task | Docker | nerdctl | ctr |
|---|---|---|---|
| Run container | `docker run -d -p 8080:80 nginx` | `nerdctl run -d -p 8080:80 nginx` | `ctr run --net-host -d docker.io/library/nginx:latest nginx_id` |
| List containers | `docker ps` | `nerdctl ps` | `ctr tasks list` |
| Build image | `docker build -t myapp .` | `nerdctl build -t myapp .` | Not supported |
| View logs | `docker logs` | `nerdctl logs` | Not supported |
| Inspect | `docker inspect` | `nerdctl inspect` | `ctr containers info` |
| Pull image | `docker pull nginx` | `nerdctl pull nginx` | `ctr images pull docker.io/library/nginx:latest` |
| Compose | `docker compose up` | `nerdctl compose up` | Not supported |

Key terms: OCI, CRI, CNCF, runc, gRPC, snapshotter, Libnetwork, CNI, dockershim, gVisor, Kata Containers.

## Why this source matters for the RAG

It clearly explains the Docker/containerd relationship, OCI and CRI standards, and migration paths, resolving a common conceptual confusion. It is valuable for RAG queries on container runtimes, Kubernetes integration, and infrastructure decisions.
