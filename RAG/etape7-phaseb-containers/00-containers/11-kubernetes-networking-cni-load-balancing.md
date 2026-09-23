---
id: etape7-phaseb-containers/00-containers/11-kubernetes-networking-cni-load-balancing
title: "11. Kubernetes networking: CNI, load balancing"
domain: step-7-phase-b-containers-orchestration-sandbox-runtimes
role: deep-dive
task: reference
actors: ["AWS", "Nvidia"]
dates: ["2026-04", "2026-09", "2026-09-09", "2026-09-11", "2026-09-15"]
keywords: ["aws", "nvidia"]
source: docs/RAG/etape7_phaseB_containers.md
source_anchor: ""
source_lines: [518, 565]
section: "Step 7 — Phase B: Containers, Orchestration & Sandbox Runtimes"
sha256: 2cded6600031858611ae1a51a7ddeef642e5b114b8d63be66afe80c89b1e0eb1
---

# 11. Kubernetes networking: CNI, load balancing

## 11. Kubernetes networking: CNI, load balancing

### 11.1 Cilium (eBPF)

- **Actively maintained (2026-09-15):** v1.20.2 (latest), v1.19.8, v1.18.14; images
  `quay.io/cilium/cilium:v1.20.x`, amd64 + arm64 [official].
  - Source: https://github.com/cilium/cilium
- v1.21.0-pre.2 published 2026-09-09 (dev only) [secondary].
- Modes: VXLAN/Geneve overlay, native routing, L2 ND, BGP route advertisement;
  kube-proxy replacement; Hubble observability; EKS chaining mode with AWS VPC CNI [official].
- Production migration pattern (2026): Flannel + kube-proxy → Cilium via Helm, with
  eBPF kube-proxy replacement and Hubble enabled [secondary].
  - Source: https://github.com/derio-net/frank/blob/HEAD/patches/phase02-cilium/README.md
- Verified April 2026 on K8s 1.35.3: Cilium 1.19.2 + Hubble 1.18.6 on Ubuntu 24.04
  (kernel 6.8) [independent].
  - Source: https://computingforgeeks.com/cilium-ebpf-cni-kubernetes-production/
- **[gap]** No riscv64 CI or release artifacts upstream (issue #39977, open as of
  2026-09-11) [secondary].

### 11.2 Calico (Tigera)

- **Calico 3.31.3** pinned in NVIDIA Cloud Native Stack 25.12.0; charmed-K8s pins
  calico 3.29.3 [vendor-reported/secondary].
- eBPF dataplane option; WireGuard encryption; NetworkPolicy + GlobalNetworkPolicy;
  Calico Cloud (managed) offering [secondary].
- **[gap]** Exact latest Calico release number for September 2026 not captured verbatim.

### 11.3 Others

- **Flannel:** still the K3s/Talos default simple overlay (VXLAN/host-gw); commonly
  replaced by Cilium in production [secondary].
- **Canal** (Flannel + Calico policy), **Weave** (discontinued — Weaveworks shut down
  2024; community forks only) [secondary].
- **kube-ovn 1.12.30** used in Charmed Kubernetes for OVN-based networking [secondary].
- **Multus:** multi-NIC / SR-IOV attachment for telco and AI workloads; NetworkBindingPlugins
  in KubeVirt 1.9 for non-default bindings [secondary].

### 11.4 Ingress and Gateway API

- **ingress-nginx 1.11.5 / 1.14.4** (CVE-2026-3288 fixed in 1.14.4) [secondary].
- **Traefik 3.7.8** bundled in K3s (chart v40.1.4) [vendor-reported].
- **Gateway API** is the 2026 successor to Ingress: GatewayClass, HTTPRoute, GRPCRoute;
  Istio, Cilium, Traefik, NGINX all ship Gateway API providers [independent].
- **MetalLB** remains the bare-metal LoadBalancer standard; **Cilium L2/BGP** announcements
  increasingly replace it [secondary].

---

