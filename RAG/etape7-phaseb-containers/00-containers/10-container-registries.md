---
id: etape7-phaseb-containers/00-containers/10-container-registries
title: "10. Container registries"
domain: step-7-phase-b-containers-orchestration-sandbox-runtimes
role: deep-dive
task: reference
actors: ["AWS", "EU", "Microsoft", "Nvidia"]
dates: ["2026-09"]
keywords: ["attribution", "aws", "nvidia", "pricing", "scout"]
source: docs/RAG/etape7_phaseB_containers.md
source_anchor: ""
source_lines: [465, 517]
section: "Step 7 — Phase B: Containers, Orchestration & Sandbox Runtimes"
sha256: 3303ab3ad6d831f50d8bee9e4bf45cd4dec692f843a08be37924db273e8654a3
---

# 10. Container registries

## 10. Container registries

### 10.1 Docker Hub

- **Pull rate limits** (official, current September 2026) [official]:
  - Source: https://docs.docker.com/docker-hub/usage/pulls/
  - Business/Team/Pro (authenticated): **unlimited** (fair use).
  - Personal (authenticated): **200 pulls / 6h**.
  - Unauthenticated: **100 pulls / 6h** per IPv4 or IPv6 /64.
- A "pull" = one manifest fetch (version checks excluded); multi-arch images count one
  pull per architecture [official].
- Pulls are attributed by requester's account, not image owner's; anonymous attribution
  by IP [official].
- Hub pricing ties to Docker plans: Personal free (1 private repo), Pro **$9/mo** unlimited
  private repos, Team/Business unlimited pulls [independent].
- Monthly pull caps (25k/100k/1M) once planned for paid tiers were dropped; paid
  subscribers keep unlimited pulls; storage billing "delayed indefinitely" [independent].
  - Source: https://kloudle.com/academy/are-you-missing-patching-your-clusters-because-of-the-new-docker-hub-limits/
- Operational impact: unauthenticated K8s nodes behind NAT share the 100-pull budget —
  imagePullSecrets or registry mirrors are standard mitigation [independent].

### 10.2 Alternatives

- **GitHub Container Registry (GHCR):** free for public packages; private packages draw
  from GitHub plan quotas (Free: 500MB storage/1GB transfer; Team: 2GB/10GB); overage
  $0.008/GB storage, $0.50/GB transfer; `GITHUB_TOKEN` auth in Actions [independent].
  - Source: https://devopsboys.com/blog/aws-ecr-vs-docker-hub-vs-ghcr-container-registry-2026
- **Amazon ECR:** $0.10/GB-mo storage; $0.09/GB egress to internet; **free** transfer to
  EC2/EKS in same region; IAM/IRSA auth (no stored credentials); Snyk/Clair scanning,
  lifecycle policies, immutable tags [independent].
- **Harbor (CNCF):** self-hosted registry with RBAC, replication, vulnerability scanning
  (Trivy), image signing (cosign/Notary), retention policies; standard enterprise
  air-gapped choice [secondary].
- **Quay.io:** Red Hat's registry; hosts Cilium (`quay.io/cilium/cilium:v1.20.2`),
  KubeVirt demo images; Clair scanning lineage [official/secondary].
- **nvcr.io:** NVIDIA NGC catalog hosts container-toolkit images
  (`nvcr.io/nvidia/k8s/container-toolkit:v1.18.1`) [official].
- **[gap]** Exact Harbor 2.x latest release number for September 2026 not captured
  verbatim this pass.

### 10.3 Image signing, SBOM, and policy

- **cosign (Sigstore):** keyless signing with OIDC identity; standard for CI-attested
  images; `cosign verify` in admission pipelines [secondary].
- **Notary v2 / notation:** OCI-native signing spec; Harbor and ECR integrations [secondary].
- **SBOM:** Cilium ships SPDX SBOMs in all images since 1.13.0; Docker Scout (3 repos on
  Personal, local analysis on Pro+) surfaces CVEs; EU CRA (in force, see Phase F3)
  pushes SBOM/provenance requirements [official/independent].
- **Admission control:** Kyverno / OPA Gatekeeper policies enforcing signed images,
  non-root, read-only filesystems are 2026 baseline practice [independent].

---

