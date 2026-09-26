---
id: collect-240926-datacamp/datacamp/openshift-vs-kubernetes-explication-des-principales-differences-2
title: "openshift-vs-kubernetes-explication-des-principales-differences"
domain: datacamp
role: reference
task: reference
actors: []
dates: []
keywords: ["governance"]
source: docs/RAG/clean_en/datacamp/openshift-vs-kubernetes-explication-des-principales-differences.md
source_anchor: ""
source_lines: [105, 189]
sha256: 1df16d62158fa634ae32b66fb88a5b8a06270d609d3c646a7564bbf5ed33db9c
---

# openshift-vs-kubernetes-explication-des-principales-differences

At their core, both platforms run Kubernetes. OpenShift is built directly on top of Kubernetes and adheres to its upstream APIs and components. A standard OpenShift cluster includes all the essential Kubernetes elements, such as the API server, etcd, the scheduler, and kubelets. This means it is fully compatible with Kubernetes workloads and tools.

What sets OpenShift apart is the ecosystem that Red Hat builds around that foundation. It introduces features such as Red Hat Enterprise Linux CoreOS (RHCOS) as the node operating system, a built-in container registry, and additional platform services tightly coupled with Kubernetes.

These integrations reduce the amount of work teams need to do to build and secure a production environment.

Overview of the OpenShift Container Platform architecture. Image provided by Red Hat documentation.

### Installation and Configuration

One of the most important architectural differences is how each platform is installed.

- Kubernetes offers a great deal of deployment flexibility. You can set up a local cluster using tools like minikube, managed services in the cloud like GKE or EKS, or install it manually using tools like kubeadm. While this gives you full control, it also means you are responsible for everything, including setting up networking, configuring storage, and installing tools such as monitoring and logging.
- OpenShift, on the other hand, offers a more guided and automated installation process. OpenShift provides several installation methods to suit different infrastructure environments and user preferences.
- The first is installer-provisioned infrastructure, which automates cluster creation on supported platforms.
- The other is user-provisioned infrastructure, which gives infrastructure teams more control.
- You can learn more about installing OpenShift in the official user documentation.

### Deployment and Management Tools

Kubernetes and OpenShift both provide command-line tools for interacting with the cluster, but they differ in scope and ease of use:

- Kubernetes relies on `kubectl`, a powerful CLI tool that interacts directly with the Kubernetes API. It is flexible and widely used, but assumes a certain level of familiarity with YAML and the underlying architecture.
- OpenShift includes `oc`, a CLI that extends `kubectl` and has additional capabilities. It supports the same basic commands and adds features such as project-based access, image stream management, and streamlined login flows.

OpenShift also includes a comprehensive web console that makes it easy to manage applications, monitor workloads, and control access, without the need to write YAML by hand.

Kubernetes also offers a dashboard, but configuring and securing it requires more work.

Both platforms support Helm for managing application packages. OpenShift also integrates closely with operators, available through the built-in OperatorHub.

## Security and governance

Security is often a key differentiator between Kubernetes and OpenShift, especially for teams operating in regulated environments or enterprises with strict compliance requirements.

Understanding how both platforms handle security, access control, and compliance can help you decide which one is best suited to your organization's risk profile and governance needs.

### Default security policies

While Kubernetes provides a flexible security framework, OpenShift builds on it with a stricter and more secure approach by default.

- Kubernetes offers a lot of flexibility when it comes to security, but above all you have to build everything from scratch. Kubernetes does not enforce strict runtime security rules. It provides features such as PodSecurity Admission (PSA) and Security Contexts, which must be explicitly configured and managed.

- OpenShift is more proactive about security. It comes with Security Context Constraints (SCCs) enabled by default. They define what a pod can and cannot do, for example whether it can run as root or use the host network. By enforcing these constraints, OpenShift helps reduce the risk of misconfiguration and privilege escalation, which is especially important in multi-tenant clusters.

In practice, workloads that run well on Kubernetes might fail on OpenShift until they meet its stricter security standards. This may seem like an obstacle at first glance, but it forces teams to adopt best practices early in the development cycle.

### Role-based access control (RBAC)

Kubernetes and OpenShift both support RBAC, which allows you to control who has access to what within the cluster.

- Kubernetes provides a highly configurable RBAC system, allowing you to define roles and assign them to users or service accounts. However, it is up to the cluster administrator to define the roles correctly, and there is no imposed structure on how access should be defined.
- OpenShift builds on Kubernetes' RBAC model, but includes predefined roles and a project-based access model that makes consistent access management easier. Each OpenShift project (roughly equivalent to a Kubernetes namespace) comes with its own set of permissions, which helps teams isolate workloads and manage access more intuitively.

OpenShift is therefore better suited to teams that want to set up access control without starting from scratch.

### Built-in compliance and monitoring

- In Kubernetes, monitoring and compliance are generally add-ons. You can install tools such as Prometheus, Grafana, Falco, or audit logging systems. However, these require manual configuration and maintenance. This gives you flexibility but increases overhead.

- OpenShift includes a built-in monitoring stack (Prometheus, Alertmanager, and Grafana), centralized logging, and tools such as the Compliance Operator. They are integrated into the platform and supported from the start, which allows teams to know what is running and whether it meets security standards.

For teams working in regulated sectors, such as finance, healthcare, and government, this saves a great deal of time and effort and reduces the risk of missing something crucial.

## Developer and user experience

A significant part of choosing between Kubernetes and OpenShift comes down to how they feel to use.

While Kubernetes provides powerful building blocks, OpenShift adds a polished layer that simplifies workflows and improves productivity.

This section examines the process of developing, deploying, and managing applications on each platform.

### Interface and ease of use

- Kubernetes is primarily command-line driven, with `kubectl` as the main interface. It also offers an online dashboard, which is relatively basic and often disabled by default for security reasons. Most users interact with Kubernetes through configuration files and CLI commands. This gives you a lot of control, but also comes with a steep learning curve, especially for newcomers.
- OpenShift, on the other hand, includes a full web console designed for both developers and cluster administrators. You can create and manage projects, deploy applications from Git repositories or container images, view logs and metrics, and even trigger builds without touching the command line. The user interface is well integrated and user-friendly, which lowers the barrier to entry for teams that do not have deep Kubernetes expertise.

I like working with the OpenShift user interface because it allows you to manage your applications easily and quickly. In particular, debugging has become much faster thanks to using the UI instead of relying entirely on API commands.

I recommend reading What's New in the OpenShift 4.4 Web Console Developer Experience to learn more about the features and design of the OpenShift web console.

OpenShift also supports its own CLI tool (`oc`), which extends `kubectl` and provides additional commands tailored to OpenShift features. More experienced users thus have the same level of control they expect from Kubernetes.

### DevOps tool integration

