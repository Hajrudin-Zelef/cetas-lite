---
id: vague2-datacamp/datacamp/openshift-vs-kubernetes
title: "OpenShift vs. Kubernetes : Explication des principales différences"
domain: datacamp
role: reference
task: article
actors: ["AWS", "Google"]
dates: ["2026-09-23"]
keywords: ["aws", "cost", "distribution", "governance", "open source"]
source: docs/RAG/Collect RAG Vague 2/02_datacamp/openshift-vs-kubernetes.md
source_anchor: ""
source_lines: [1, 54]
sha256: 4b3c22cc651470d75655fbbebc7fa851915d50836614de40b87724b78b178370
---

# OpenShift vs. Kubernetes : Explication des principales différences

## Metadata

- **Source** : https://www.datacamp.com/fr/blog/openshift-vs-kubernetes
- **Site** : DataCamp
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This DataCamp article by Patrick Brus compares OpenShift and Kubernetes based on the author's experience with both (bare Kubernetes for personal projects/CKAD preparation, OpenShift for enterprise production). Kubernetes (K8s) is an open-source platform, originally developed by Google (based on its internal Borg system) and maintained by the CNCF, that automates deployment, scaling, and management of containerized applications. Its key features include container orchestration, declarative configuration (YAML/JSON), self-healing, service discovery and load balancing, horizontal scaling, rolling updates and rollbacks, and secrets/config management. Use cases include microservices, cloud-native apps, multi/hybrid-cloud, DevOps/CI-CD, and ML/data science (Kubeflow, MLflow).

OpenShift is Red Hat's Kubernetes distribution — "a fully equipped vehicle built around the raw engine" of Kubernetes. It adds an integrated container registry, a web-based admin console, developer tools, and stricter security defaults. Key components include built-in CI/CD (Tekton, GitOps), a web console, security enhancements (Security Context Constraints/SCC, image scanning, RBAC), Red Hat Enterprise Linux CoreOS (RHCOS), and OperatorHub. OpenShift comes in flavors: OpenShift Container Platform (self-managed), OpenShift Online (fully hosted), and OpenShift Dedicated (managed on AWS/Google Cloud).

The comparison covers architecture (Kubernetes modular/DIY vs OpenShift opinionated/integrated), installation (flexible minikube/kubeadm vs guided automated installers), and tooling (`kubectl` vs `oc`, which extends kubectl, plus a full web console). Security: Kubernetes requires explicit PodSecurity Admission/Security Contexts; OpenShift enables SCC by default, offers predefined RBAC roles and project-based access, and includes integrated monitoring (Prometheus/Alertmanager/Grafana), centralized logging, and the Compliance Operator. Developer experience: Kubernetes is CLI-first with a basic dashboard; OpenShift has a comprehensive web console, OpenShift Pipelines (Tekton) and OpenShift GitOps (Argo CD). Support: Kubernetes is CNCF-governed with community/vendor support (GKE, EKS, AKS); OpenShift has commercial Red Hat support with SLAs (ROSA, ARO, IBM Cloud, OpenShift Dedicated). Cost: Kubernetes is free but has infrastructure/operational costs; OpenShift is subscription-based.

The article concludes with guidance on when to choose each and key decision factors (budget, team expertise, time-to-production, security/compliance), plus an FAQ.

## Key points

- Kubernetes is an open-source orchestration platform (Google/CNCF, based on Borg); OpenShift is Red Hat's enterprise Kubernetes distribution.
- OpenShift builds on upstream Kubernetes APIs, ensuring workload/tool compatibility.
- OpenShift adds built-in registry, web console, RHCOS, OperatorHub, and stricter security defaults.
- Security: OpenShift enables Security Context Constraints (SCC) by default; Kubernetes requires explicit PodSecurity Admission/Security Contexts.
- CLI: Kubernetes uses `kubectl`; OpenShift uses `oc` (extends kubectl) plus a comprehensive web console.
- OpenShift bundles OpenShift Pipelines (Tekton) and OpenShift GitOps (Argo CD); Kubernetes requires manual DevOps tool setup.
- Kubernetes is free (open source) but has higher operational costs; OpenShift is subscription-based with Red Hat SLAs.
- Kubernetes suits experienced DevOps teams wanting flexibility and avoiding vendor lock-in; OpenShift suits regulated enterprises needing compliance, support, and fast time-to-production.
- Managed offerings: GKE, EKS, AKS (Kubernetes); ROSA, ARO, IBM Cloud, OpenShift Dedicated (OpenShift).

## Technical data / figures

| Aspect | Kubernetes | OpenShift |
|--------|-----------|-----------|
| Nature | Open-source orchestration platform | Red Hat Kubernetes distribution |
| Governance | CNCF (originally Google/Borg) | Red Hat (commercial support) |
| Installation | minikube, kubeadm, managed (GKE/EKS/AKS) | Installer-provisioned or user-provisioned infrastructure |
| CLI | `kubectl` | `oc` (extends kubectl) |
| Web console | Basic dashboard (often disabled) | Comprehensive developer/admin console |
| Node OS | Any | RHCOS (Red Hat Enterprise Linux CoreOS) |
| Security defaults | PodSecurity Admission, Security Contexts (manual) | SCC enabled by default, image scanning, RBAC |
| CI/CD | Jenkins, Argo CD, Tekton, Flux (manual setup) | OpenShift Pipelines (Tekton), OpenShift GitOps (Argo CD) |
| Monitoring | Prometheus, Grafana, Falco (add-ons) | Integrated Prometheus, Alertmanager, Grafana, Compliance Operator |
| Operator marketplace | Community/Helm | OperatorHub |
| Managed services | GKE, EKS, AKS | ROSA, ARO, IBM Cloud, OpenShift Dedicated |
| Cost model | Free (infra + operational costs) | Subscription (nodes, CPUs, deployment model) |
| Flavors | — | Container Platform, Online, Dedicated |

## Why this source matters for the RAG

This article delivers a detailed, multi-dimensional comparison of Kubernetes and its enterprise distribution OpenShift, covering architecture, security, tooling, support, and cost, which is essential for platform-selection and enterprise-deployment queries. Its concrete distinctions (SCC vs PSA, kubectl vs oc, managed offerings) provide high-value technical facts for the knowledge base.
