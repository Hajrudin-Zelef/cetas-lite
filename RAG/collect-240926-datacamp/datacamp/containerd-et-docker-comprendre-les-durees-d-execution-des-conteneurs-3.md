---
id: collect-240926-datacamp/datacamp/containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs-3
title: "containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs"
domain: datacamp
role: reference
task: reference
actors: ["AWS", "Google"]
dates: []
keywords: ["attention", "aws", "memory"]
source: docs/RAG/clean_en/datacamp/containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs.md
source_anchor: ""
source_lines: [179, 274]
sha256: a988548c563b27dd839eb3636620a76fd07913f27737129c462b72254bb1c3e4
---

# containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs

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

