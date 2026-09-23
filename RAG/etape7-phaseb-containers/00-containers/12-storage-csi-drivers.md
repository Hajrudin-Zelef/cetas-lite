---
id: etape7-phaseb-containers/00-containers/12-storage-csi-drivers
title: "12. Storage: CSI drivers"
domain: step-7-phase-b-containers-orchestration-sandbox-runtimes
role: deep-dive
task: reference
actors: ["AWS"]
dates: ["2026-02-16", "2026-05-14", "2026-08-27", "2026-09", "2026-09-09", "2026-10-12"]
keywords: ["aws"]
source: docs/RAG/etape7_phaseB_containers.md
source_anchor: ""
source_lines: [566, 610]
section: "Step 7 — Phase B: Containers, Orchestration & Sandbox Runtimes"
sha256: ee087dfe7c9a830b3c333c94377d52d6ce193e04bce354987ba7d254494dbfbe
---

# 12. Storage: CSI drivers

## 12. Storage: CSI drivers

- **CSI** is the standard: cloud disks (EBS, Azure Disk, GCE PD), Ceph (cephcsi 3.13.0),
  Longhorn, OpenEBS, Rook-Ceph, Portworx (Pure Storage), local-path-provisioner [secondary].
- EBS CSI + VolumeAttributesClass GA (1.34) enables volume modification rollback on EKS;
  sidecars patched through EKS 1.33 EOL [official].
- **TopoLVM**, **LINSTOR/DRBD** (now an Incus 7.0 remote storage option), TrueNAS iSCSI
  driver for Incus — convergence between K8s CSI and Incus storage pools [secondary].
- Volume populators / DataVolumes (CDI v1.65.0–1.66.0) for VM disk import in KubeVirt [secondary].
- **[gap]** Exact latest CSI sidecar/driver versions for September 2026 not exhaustively
  enumerated here; covered in depth by the storage step (Step 9).

---

## 13. Service mesh

### 13.1 Istio

- **Supported releases (September 2026):** 1.31 (released 2026-08-27, EOL ~Feb 2027),
  1.30 (2026-05-14, EOL ~Dec 2026), 1.29 (2026-02-16, EOL 2026-10-12); 1.28 and older
  unsupported [official].
  - Source: http://istio.io/latest/docs/releases/supported-releases/
- Supported K8s: 1.32–1.36 for 1.30/1.31 [official].
- Architecture: istiod control plane + Envoy sidecars (ambient mesh sidecarless mode
  available since 1.22, maturing through 2026); distroless proxy images in Rancher
  charts [official/secondary].
- Solo Gloo Mesh Gateway 2.14 (2026-09-09) supports Istio 1.27–1.31 on K8s 1.30–1.36 [vendor-reported].
  - Source: https://docs.solo.io/gloo-mesh-gateway/main/reference/version/istio/versions/

### 13.2 Linkerd

- Buoyant's Linkerd: Rust-based micro-proxy (linkerd2-proxy), service mesh focused on
  simplicity and low overhead; 2.x line current in 2026 [secondary].
- **[gap]** Exact Linkerd 2026 stable version number not captured verbatim this pass.

### 13.3 Alternatives and trends

- **Cilium Service Mesh:** sidecar-free mesh via eBPF (Envoy DaemonSet option) —
  main challenger to Istio on operational simplicity [independent].
- **Consul Connect**, **AWS App Mesh** (EOL'd — migration to Istio/Envoy recommended),
  **NGINX Service Mesh** (discontinued) — consolidation around Istio/Cilium/Linkerd [secondary].
- Gateway API + mesh (GAMMA initiative) unifies ingress and east-west policy [independent].

---

