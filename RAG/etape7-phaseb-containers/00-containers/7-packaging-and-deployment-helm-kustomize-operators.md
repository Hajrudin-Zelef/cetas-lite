---
id: etape7-phaseb-containers/00-containers/7-packaging-and-deployment-helm-kustomize-operators
title: "7. Packaging and deployment: Helm, Kustomize, operators"
domain: step-7-phase-b-containers-orchestration-sandbox-runtimes
role: deep-dive
task: reference
actors: ["EU", "Nvidia"]
dates: ["2025-11-12", "2026-07", "2026-07-08", "2026-08-05", "2026-08-13", "2026-09-09", "2026-11-11"]
keywords: ["packaging", "gpu", "nvidia"]
source: docs/RAG/etape7_phaseB_containers.md
source_anchor: ""
source_lines: [317, 361]
section: "Step 7 — Phase B: Containers, Orchestration & Sandbox Runtimes"
sha256: 38ead8b89c637b96d0a9be37e009a623bb75169185d273bb00e515adeb8d07a9
---

# 7. Packaging and deployment: Helm, Kustomize, operators

## 7. Packaging and deployment: Helm, Kustomize, operators

### 7.1 Helm

- **Latest stable:** **v4.2.4** (2026-08-13); Helm 4 is the current stable line, developed on
  `main`; **Helm 4.0.0** released **November 12, 2025** at KubeCon NA — first major in six years [official].
  - Sources: http://en.wikipedia.org/wiki/Helm_(package_manager),
    https://github.com/helm/helm/blob/HEAD/README.md
- **Helm 4 key changes:** secrets-only release storage; built-in OCI push (no plugin);
  Lua + Sprig templating; server-side apply default for new installs; kstatus-based
  `--wait` readiness; WebAssembly plugins [independent].
  - Source: https://medium.com/@rameshavutu/helm-4-vs-helm-3-should-you-upgrade-abb62e66f174
- **Helm 3 EOL:** bug fixes until **2026-07-08**, security fixes until **2026-11-11**;
  final feature release 2026-09-09; Helm 3 charts (Chart API v2) run unmodified on Helm 4 [official].
- **Migration gotchas** (independent, July 2026): post-renderers must be registered plugins;
  `helm registry login` takes domain only; `--atomic` → `--rollback-on-failure`,
  `--force` → `--force-replace` (renamed, deprecation warnings); `--wait` needs `watch`
  RBAC verb for kstatus [independent].
  - Source: https://medium.com/@amareswer/helm-4-migration-guide-what-breaks-and-how-to-fix-it-before-eol-7779b340e3ee
- NVIDIA Cloud Native Stack pins **Helm 4.0.4** [vendor-reported].
- Production practice (2026): Helm + Kustomize post-renderer, OCI chart registries,
  Artifact Hub as chart discovery [independent].

### 7.2 Kustomize

- Kustomize remains the built-in `kubectl -k` overlay engine; 2026 usage centers on
  environment overlays, Flux/Kustomize controllers, and Helm post-rendering [independent].
- Flux Kustomizations commonly vendor operator manifests (e.g., KubeVirt operator +
  CR split) with `dependsOn` ordering [secondary].
  - Source: https://github.com/alinanova21/home-ops/blob/HEAD/docs/superpowers/plans/2026-08-05-kubevirt-deployment.md

### 7.3 Operators and ecosystem

- Operator pattern dominant for stateful workloads: KubeVirt, CDI (v1.65.0/1.66.0),
  NVIDIA GPU Operator, cert-manager, DocumentDB operator (requires K8s 1.35+ ImageVolume GA),
  CNPG Postgres, Karpenter [secondary].
- **CNCF landscape 2026:** Kubernetes, Prometheus, Envoy, Helm, containerd, Cilium,
  Argo, Backstage graduated; KubeVirt incubating→stable trajectory since CNCF adoption
  (KubeCon EU 2026 marked its "post-KVM" 1.8/1.9 era) [independent/secondary].
- **Artifact Hub** is the CNCF chart/operator discovery index (Helm Hub retired 2020) [official].
- **Karpenter** vs cluster-autoscaler: Karpenter is the 2026 default for EKS workload
  autoscaling in ML stacks [secondary].

---

