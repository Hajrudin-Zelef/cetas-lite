# INDEX — Step 7 — Phase B: Containers, Orchestration & Sandbox Runtimes

Corpus `etape7-phaseb-containers` · **12 fichiers** · 796 lignes source · ~5065 mots · partition exacte de `docs/RAG/etape7_phaseB_containers.md`.

## Mode d'emploi

1. Filtrer dans `manifest.json` (ou les tableaux ci-dessous) sur `domain`, `task`, `actors`, `dates` ou `keywords`.
2. Ouvrir 1 à 3 fichiers ciblés ; chaque fichier est une unité thématique auto-suffisante avec un en-tête YAML.
3. Pour un événement répété dans plusieurs sections, préférer le fichier marqué `canonical_for` (voir la table Événements canoniques).

## Domaines (dossiers → fichiers)

### `00-containers/`

| # | fichier | lignes source | rôle | tâche |
|---|---|---|---|---|
| 01 | [Step 7 — Phase B: Containers, Orchestration & Sandbox Runtimes](00-containers/overview.md) | 1–87 | deep-dive | reference |
| 02 | [2. Desktop and local development tooling](00-containers/2-desktop-and-local-development-tooling.md) | 88–164 | deep-dive | reference |
| 03 | [4. Compose and build tooling](00-containers/4-compose-and-build-tooling.md) | 165–227 | deep-dive | reference |
| 04 | [6. Kubernetes distributions](00-containers/6-kubernetes-distributions.md) | 228–316 | deep-dive | reference |
| 05 | [7. Packaging and deployment: Helm, Kustomize, operators](00-containers/7-packaging-and-deployment-helm-kustomize-operators.md) | 317–361 | deep-dive | reference |
| 06 | [8. System containers: LXC, LXD, Incus](00-containers/8-system-containers-lxc-lxd-incus.md) | 362–406 | deep-dive | reference |
| 07 | [9. Sandbox and microVM runtimes](00-containers/9-sandbox-and-microvm-runtimes.md) | 407–464 | deep-dive | reference |
| 08 | [10. Container registries](00-containers/10-container-registries.md) | 465–517 | deep-dive | reference |
| 09 | [11. Kubernetes networking: CNI, load balancing](00-containers/11-kubernetes-networking-cni-load-balancing.md) | 518–565 | deep-dive | reference |
| 10 | [12. Storage: CSI drivers](00-containers/12-storage-csi-drivers.md) | 566–610 | deep-dive | reference |
| 11 | [14. GPUs in containers](00-containers/14-gpus-in-containers.md) | 611–726 | deep-dive | hardware |
| 12 | [18. Source index (verbatim URLs)](00-containers/18-source-index-verbatim-urls.md) | 727–796 | deep-dive | reference |

## Par tâche

- **hardware** — [14. GPUs in containers](00-containers/14-gpus-in-containers.md)
- **reference** — [Step 7 — Phase B: Containers, Orchestration & Sandbox Runtimes](00-containers/overview.md), [2. Desktop and local development tooling](00-containers/2-desktop-and-local-development-tooling.md), [4. Compose and build tooling](00-containers/4-compose-and-build-tooling.md), [6. Kubernetes distributions](00-containers/6-kubernetes-distributions.md), [7. Packaging and deployment: Helm, Kustomize, operators](00-containers/7-packaging-and-deployment-helm-kustomize-operators.md), [8. System containers: LXC, LXD, Incus](00-containers/8-system-containers-lxc-lxd-incus.md), [9. Sandbox and microVM runtimes](00-containers/9-sandbox-and-microvm-runtimes.md), [10. Container registries](00-containers/10-container-registries.md), [11. Kubernetes networking: CNI, load balancing](00-containers/11-kubernetes-networking-cni-load-balancing.md), [12. Storage: CSI drivers](00-containers/12-storage-csi-drivers.md), [18. Source index (verbatim URLs)](00-containers/18-source-index-verbatim-urls.md)

## Par acteur

- **AMD** (4) — [00-containers/overview.md](00-containers/overview.md), [00-containers/2-desktop-and-local-development-tooling.md](00-containers/2-desktop-and-local-development-tooling.md), [00-containers/9-sandbox-and-microvm-runtimes.md](00-containers/9-sandbox-and-microvm-runtimes.md), [00-containers/14-gpus-in-containers.md](00-containers/14-gpus-in-containers.md)
- **AWS** (6) — [00-containers/4-compose-and-build-tooling.md](00-containers/4-compose-and-build-tooling.md), [00-containers/6-kubernetes-distributions.md](00-containers/6-kubernetes-distributions.md), [00-containers/9-sandbox-and-microvm-runtimes.md](00-containers/9-sandbox-and-microvm-runtimes.md), [00-containers/10-container-registries.md](00-containers/10-container-registries.md), [00-containers/11-kubernetes-networking-cni-load-balancing.md](00-containers/11-kubernetes-networking-cni-load-balancing.md), [00-containers/12-storage-csi-drivers.md](00-containers/12-storage-csi-drivers.md)
- **EU** (3) — [00-containers/7-packaging-and-deployment-helm-kustomize-operators.md](00-containers/7-packaging-and-deployment-helm-kustomize-operators.md), [00-containers/10-container-registries.md](00-containers/10-container-registries.md), [00-containers/14-gpus-in-containers.md](00-containers/14-gpus-in-containers.md)
- **Google** (2) — [00-containers/6-kubernetes-distributions.md](00-containers/6-kubernetes-distributions.md), [00-containers/9-sandbox-and-microvm-runtimes.md](00-containers/9-sandbox-and-microvm-runtimes.md)
- **Intel** (2) — [00-containers/9-sandbox-and-microvm-runtimes.md](00-containers/9-sandbox-and-microvm-runtimes.md), [00-containers/14-gpus-in-containers.md](00-containers/14-gpus-in-containers.md)
- **Lambda** (1) — [00-containers/9-sandbox-and-microvm-runtimes.md](00-containers/9-sandbox-and-microvm-runtimes.md)
- **Microsoft** (1) — [00-containers/10-container-registries.md](00-containers/10-container-registries.md)
- **Nvidia** (7) — [00-containers/overview.md](00-containers/overview.md), [00-containers/6-kubernetes-distributions.md](00-containers/6-kubernetes-distributions.md), [00-containers/7-packaging-and-deployment-helm-kustomize-operators.md](00-containers/7-packaging-and-deployment-helm-kustomize-operators.md), [00-containers/10-container-registries.md](00-containers/10-container-registries.md), [00-containers/11-kubernetes-networking-cni-load-balancing.md](00-containers/11-kubernetes-networking-cni-load-balancing.md), [00-containers/14-gpus-in-containers.md](00-containers/14-gpus-in-containers.md), [00-containers/18-source-index-verbatim-urls.md](00-containers/18-source-index-verbatim-urls.md)

## Par date

- **2024-12-10** — [00-containers/2-desktop-and-local-development-tooling.md](00-containers/2-desktop-and-local-development-tooling.md)
- **2025-08** — [00-containers/4-compose-and-build-tooling.md](00-containers/4-compose-and-build-tooling.md)
- **2025-08-27** — [00-containers/6-kubernetes-distributions.md](00-containers/6-kubernetes-distributions.md)
- **2025-10-02** — [00-containers/6-kubernetes-distributions.md](00-containers/6-kubernetes-distributions.md)
- **2025-11-10** — [00-containers/overview.md](00-containers/overview.md)
- **2025-11-12** — [00-containers/7-packaging-and-deployment-helm-kustomize-operators.md](00-containers/7-packaging-and-deployment-helm-kustomize-operators.md)
- **2025-12-17** — [00-containers/4-compose-and-build-tooling.md](00-containers/4-compose-and-build-tooling.md), [00-containers/6-kubernetes-distributions.md](00-containers/6-kubernetes-distributions.md)
- **2026-01-02** — [00-containers/overview.md](00-containers/overview.md)
- **2026-01-27** — [00-containers/4-compose-and-build-tooling.md](00-containers/4-compose-and-build-tooling.md), [00-containers/6-kubernetes-distributions.md](00-containers/6-kubernetes-distributions.md)
- **2026-02-16** — [00-containers/12-storage-csi-drivers.md](00-containers/12-storage-csi-drivers.md)
- **2026-04** — [00-containers/11-kubernetes-networking-cni-load-balancing.md](00-containers/11-kubernetes-networking-cni-load-balancing.md)
- **2026-04-22** — [00-containers/4-compose-and-build-tooling.md](00-containers/4-compose-and-build-tooling.md), [00-containers/6-kubernetes-distributions.md](00-containers/6-kubernetes-distributions.md)
- **2026-04-27** — [00-containers/6-kubernetes-distributions.md](00-containers/6-kubernetes-distributions.md)
- **2026-05-01** — [00-containers/8-system-containers-lxc-lxd-incus.md](00-containers/8-system-containers-lxc-lxd-incus.md)
- **2026-05-14** — [00-containers/12-storage-csi-drivers.md](00-containers/12-storage-csi-drivers.md)
- **2026-05-27** — [00-containers/6-kubernetes-distributions.md](00-containers/6-kubernetes-distributions.md), [00-containers/18-source-index-verbatim-urls.md](00-containers/18-source-index-verbatim-urls.md)
- **2026-06-02** — [00-containers/4-compose-and-build-tooling.md](00-containers/4-compose-and-build-tooling.md), [00-containers/6-kubernetes-distributions.md](00-containers/6-kubernetes-distributions.md)
- **2026-06-03** — [00-containers/9-sandbox-and-microvm-runtimes.md](00-containers/9-sandbox-and-microvm-runtimes.md)
- **2026-06-23** — [00-containers/4-compose-and-build-tooling.md](00-containers/4-compose-and-build-tooling.md)
- **2026-07** — [00-containers/7-packaging-and-deployment-helm-kustomize-operators.md](00-containers/7-packaging-and-deployment-helm-kustomize-operators.md)
- **2026-07-08** — [00-containers/7-packaging-and-deployment-helm-kustomize-operators.md](00-containers/7-packaging-and-deployment-helm-kustomize-operators.md)
- **2026-07-26** — [00-containers/18-source-index-verbatim-urls.md](00-containers/18-source-index-verbatim-urls.md)
- **2026-07-30** — [00-containers/4-compose-and-build-tooling.md](00-containers/4-compose-and-build-tooling.md)
- **2026-08** — [00-containers/2-desktop-and-local-development-tooling.md](00-containers/2-desktop-and-local-development-tooling.md)
- **2026-08-05** — [00-containers/7-packaging-and-deployment-helm-kustomize-operators.md](00-containers/7-packaging-and-deployment-helm-kustomize-operators.md), [00-containers/18-source-index-verbatim-urls.md](00-containers/18-source-index-verbatim-urls.md)
- **2026-08-13** — [00-containers/7-packaging-and-deployment-helm-kustomize-operators.md](00-containers/7-packaging-and-deployment-helm-kustomize-operators.md)
- **2026-08-21** — [00-containers/9-sandbox-and-microvm-runtimes.md](00-containers/9-sandbox-and-microvm-runtimes.md)
- **2026-08-27** — [00-containers/12-storage-csi-drivers.md](00-containers/12-storage-csi-drivers.md)
- **2026-09** — [00-containers/overview.md](00-containers/overview.md), [00-containers/2-desktop-and-local-development-tooling.md](00-containers/2-desktop-and-local-development-tooling.md), [00-containers/4-compose-and-build-tooling.md](00-containers/4-compose-and-build-tooling.md), [00-containers/6-kubernetes-distributions.md](00-containers/6-kubernetes-distributions.md), [00-containers/9-sandbox-and-microvm-runtimes.md](00-containers/9-sandbox-and-microvm-runtimes.md), [00-containers/10-container-registries.md](00-containers/10-container-registries.md), [00-containers/11-kubernetes-networking-cni-load-balancing.md](00-containers/11-kubernetes-networking-cni-load-balancing.md), [00-containers/12-storage-csi-drivers.md](00-containers/12-storage-csi-drivers.md), [00-containers/14-gpus-in-containers.md](00-containers/14-gpus-in-containers.md)
- **2026-09-03** — [00-containers/overview.md](00-containers/overview.md)
- **2026-09-07** — [00-containers/9-sandbox-and-microvm-runtimes.md](00-containers/9-sandbox-and-microvm-runtimes.md)
- **2026-09-09** — [00-containers/7-packaging-and-deployment-helm-kustomize-operators.md](00-containers/7-packaging-and-deployment-helm-kustomize-operators.md), [00-containers/11-kubernetes-networking-cni-load-balancing.md](00-containers/11-kubernetes-networking-cni-load-balancing.md), [00-containers/12-storage-csi-drivers.md](00-containers/12-storage-csi-drivers.md)
- **2026-09-10** — [00-containers/9-sandbox-and-microvm-runtimes.md](00-containers/9-sandbox-and-microvm-runtimes.md)
- **2026-09-11** — [00-containers/11-kubernetes-networking-cni-load-balancing.md](00-containers/11-kubernetes-networking-cni-load-balancing.md)
- **2026-09-14** — [00-containers/2-desktop-and-local-development-tooling.md](00-containers/2-desktop-and-local-development-tooling.md)
- **2026-09-15** — [00-containers/overview.md](00-containers/overview.md), [00-containers/11-kubernetes-networking-cni-load-balancing.md](00-containers/11-kubernetes-networking-cni-load-balancing.md)
- **2026-09-22** — [00-containers/overview.md](00-containers/overview.md), [00-containers/4-compose-and-build-tooling.md](00-containers/4-compose-and-build-tooling.md), [00-containers/18-source-index-verbatim-urls.md](00-containers/18-source-index-verbatim-urls.md)
- **2026-10-12** — [00-containers/12-storage-csi-drivers.md](00-containers/12-storage-csi-drivers.md)
- **2026-10-27** — [00-containers/4-compose-and-build-tooling.md](00-containers/4-compose-and-build-tooling.md)
- **2026-11-11** — [00-containers/7-packaging-and-deployment-helm-kustomize-operators.md](00-containers/7-packaging-and-deployment-helm-kustomize-operators.md)
- **2026-11-26** — [00-containers/6-kubernetes-distributions.md](00-containers/6-kubernetes-distributions.md)
- **2026-12** — [00-containers/6-kubernetes-distributions.md](00-containers/6-kubernetes-distributions.md)
- **2026-12-02** — [00-containers/6-kubernetes-distributions.md](00-containers/6-kubernetes-distributions.md)
- **2026-12-03** — [00-containers/9-sandbox-and-microvm-runtimes.md](00-containers/9-sandbox-and-microvm-runtimes.md)
- **2027-02-28** — [00-containers/4-compose-and-build-tooling.md](00-containers/4-compose-and-build-tooling.md)
- **2027-03-10** — [00-containers/9-sandbox-and-microvm-runtimes.md](00-containers/9-sandbox-and-microvm-runtimes.md)
- **2027-03-27** — [00-containers/6-kubernetes-distributions.md](00-containers/6-kubernetes-distributions.md)
- **2027-06-28** — [00-containers/4-compose-and-build-tooling.md](00-containers/4-compose-and-build-tooling.md)
- **2027-07-29** — [00-containers/6-kubernetes-distributions.md](00-containers/6-kubernetes-distributions.md)
- **2027-08-02** — [00-containers/4-compose-and-build-tooling.md](00-containers/4-compose-and-build-tooling.md), [00-containers/6-kubernetes-distributions.md](00-containers/6-kubernetes-distributions.md)
- **2028-08-02** — [00-containers/4-compose-and-build-tooling.md](00-containers/4-compose-and-build-tooling.md)

## Carte de couverture (lignes source)

| plage | fichier |
|---|---|
| 1–87 | etape7-phaseb-containers/00-containers/overview.md |
| 88–164 | etape7-phaseb-containers/00-containers/2-desktop-and-local-development-tooling.md |
| 165–227 | etape7-phaseb-containers/00-containers/4-compose-and-build-tooling.md |
| 228–316 | etape7-phaseb-containers/00-containers/6-kubernetes-distributions.md |
| 317–361 | etape7-phaseb-containers/00-containers/7-packaging-and-deployment-helm-kustomize-operators.md |
| 362–406 | etape7-phaseb-containers/00-containers/8-system-containers-lxc-lxd-incus.md |
| 407–464 | etape7-phaseb-containers/00-containers/9-sandbox-and-microvm-runtimes.md |
| 465–517 | etape7-phaseb-containers/00-containers/10-container-registries.md |
| 518–565 | etape7-phaseb-containers/00-containers/11-kubernetes-networking-cni-load-balancing.md |
| 566–610 | etape7-phaseb-containers/00-containers/12-storage-csi-drivers.md |
| 611–726 | etape7-phaseb-containers/00-containers/14-gpus-in-containers.md |
| 727–796 | etape7-phaseb-containers/00-containers/18-source-index-verbatim-urls.md |

