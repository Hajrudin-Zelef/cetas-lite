---
id: collect-261001-rattrapage/rattrapage/datacenter-cpu-ram-guide-20
title: "CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION"
domain: rattrapage
role: reference
task: reference
actors: ["AMD"]
dates: ["2026-09-27"]
keywords: ["amd", "benchmark", "benchmarks", "datacenter", "gpu", "intel", "memory"]
source: docs/RAG/collect-261001-rattrapage/datacenter_cpu_ram_guide.md
source_anchor: ""
source_lines: [3166, 3348]
sha256: 2c9f08ef4d14c538ebfb52a3d85adf08e04809c14f60ad28b180b998a04e546e
---

# CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION

**Quelle garantie ?** 5 ans sur site J+1 pour un cycle de 5 ans ; 3 ans
minimum. Négociée à l'achat (section 139).

**Tour ou rack ?** Tour si pas de local technique ; rack dès 3-4 serveurs
ou un local dédié (section 76).

## 177. FAQ acheteur (2/2)

**Combien de temps garder un serveur ?** 5 ans en production, +2 ans en
non-critique si l'état le permet (section 159). Au-delà, le risque panne
et la conso (vs générations récentes) plaident pour le renouvellement.

**Faut-il du 4S ?** Presque toujours non (section 27). Le scale-out gagne.

**Cloud ou on-premise ?** Comparez le TCO 5 ans on-premise (section 171 :
~250 €/VM/an) au prix cloud pour **votre** usage réel (les VM cloud
allumées 24/7 coûtent cher ; l'on-premise gagne sur les charges stables).

**Quel PUE viser ?** ≤ 1,5 en rénovation, ≤ 1,3 en neuf bien conçu
(sections 81, 92).

**Par où commencer ?** Mesurez : PDU monitorées, 1 mois de données, puis
décidez (section 158).

## 178. Unités et conversions : mémo

| Grandeur | Unité | Conversion |
|---|---|---|
| Puissance | watt (W) | 1 kW = 1000 W |
| Énergie | kilowattheure (kWh) | P × temps |
| Débit mémoire | MT/s puis Go/s | Go/s = MT/s × 8 / 1000 |
| Pression acoustique | dB(A) | +10 dB = 10× l'intensité |
| Température | °C | — |
| Intensité triphasée | ampère (A) | I = P/(√3×U×cos φ) |
| PUE | sans unité | ≥ 1,0 |

## 179. Erreurs d'interprétation fréquentes des fiches techniques

- **« Jusqu'à 6400 MT/s »** : c'est en 1 DPC, 1R/2R validé — pas avec
  n'importe quelle population (section 47).
- **« 500 W TDP »** : c'est le défaut ; le PPT peut dépasser, le cTDP peut
  être baissé (sections 9, 133).
- **« 2× plus rapide »** (marketing) : sur quel benchmark, quelle config,
  quelle version ? (section 103).
- **« CXL-ready »** : le CPU le supporte, mais il faut le BIOS, l'OS et le
  matériel — trois conditions cumulatives (section 97).
- **« 96 cœurs »** : avec ou sans HT ? P-core ou E-core ? La fiche ne dit
  pas tout (sections 11, 20).

## 180. Pour aller plus loin : les 10 réflexes du chef de service

1. Tout achat > 10 k€ = **TCO 5 ans** écrit (section 101).
2. Tout devis = **3 offres** comparées à config identique (section 139).
3. Toute baie = **calcul kW/rack** avant commande (section 78).
4. Tout serveur = **supervisé** dès le premier jour (section 161).
5. Toute RAM = **QVL + symétrie** vérifiées (sections 45, 115).
6. Tout onduleur = **testé** mensuellement (section 91).
7. Tout firmware = **par vagues** (section 164).
8. Toute salle = **PUE suivi** annuellement (section 82).
9. Tout projet = **mesuré** avant d'être optimisé (section 158).
10. Tout guide = **revérifié** à chaque génération (section 130).

---

*Fin de l'extension — 180 sections. Rédigé et vérifié le 27/09/2026.*

---

## 181. Réglages BIOS critiques : la check-list

| Réglage | Recommandé | Pourquoi |
|---|---|---|
| Profil énergie | « Performance » ou « Balanced Performance » | le mode « Power Saving » agressif ajoute de la latence |
| C-states | C1E activé, C6 désactivé (latence) | compromis conso/latence pour la virtualisation |
| Turbo Boost | activé | la perf mono-thread en dépend |
| Hyper-threading/SMT | activé (sauf licences au thread) | +20-30 % de débit en virtualisation |
| NUMA | activé (exposé à l'OS) | placement mémoire optimal |
| Virtualisation (VT-x/AMD-V) | activé | obligatoire pour l'hyperviseur |
| IOMMU (VT-d/AMD-Vi) | activé | SR-IOV, passthrough |
| Courbe ventilateurs | profil datacenter | pas de « silencieux » en salle |
| cTDP | nominal ou -15 % (section 133) | selon politique énergie |
| Secure Boot | activé | hygiène sécurité |

**Documentez les réglages** (export BIOS) : après un reset ou un
remplacement de carte mère, les défauts reviennent et les performances
changent silencieusement.

## 182. Tuning NUMA sous Linux : commandes utiles

```bash
# Voir la topologie
numactl --hardware
lscpu | grep -i numa

# Voir où tourne un processus et sa mémoire
numastat -p <pid>

# Lancer un workload sur un nœud précis
numactl --cpunodebind=0 --membind=0 ./mon_app

# Politique mémoire : interleave (répartir) vs local (défaut)
numactl --interleave=all ./mon_app   # HPC : bande passante max
```

Règles :
- Bases de données : **membind local** strict, une instance par socket si
  possible (2 instances sur 2S > 1 instance à cheval).
- HPC : interleave pour la bande passante agrégée, sauf codes NUMA-aware.
- Vérifiez avec `numastat` : un taux élevé de `numa_miss` = placement à
  revoir.

## 183. Paramètres kernel Linux pour serveurs

| Paramètre | Valeur serveur type | Effet |
|---|---|---|
| `vm.swappiness` | 10 (voire 1) | évite le swap intempestif |
| `vm.overcommit_memory` | 2 (strict) pour DB | pas de promesse intenable |
| `transparent_hugepage` | `madvise` (DB) / `always` (HPC) | TLB, à tester par workload |
| `intel_pstate` / `amd_pstate` | `performance` (DB/HPC) | gouverneur CPU |
| `mitigations` | `off` uniquement en labo isolé | les mitigations coûtent 2-10 % |

**Ne désactivez les mitigations Spectre/Meltdown qu'en environnement isolé
et documenté** : le gain (quelques %) ne vaut pas le risque en production
exposée.

## 184. Virtualisation : tuning de l'hyperviseur

- **vNUMA** : exposez la topologie au guest quand la VM dépasse ~8 vCPU —
  l'OS invité place alors correctement sa mémoire.
- **1 VM = 1 nœud NUMA** : dimensionnez les grosses VM en multiples de
  demi-socket (ex. ≤ 48 vCPU sur un socket 96 cœurs).
- **CPU pinning** : pour les VM critiques (DB, temps réel), épinglez les
  vCPU sur des cœurs physiques fixes — fini les migrations intempestives.
- **Ballooning** : utile en surallocation, à désactiver pour les VM
  critiques (latence imprévisible).
- **Paravirtualisation** : drivers virtio (KVM/Proxmox), VMware Tools —
  toujours installés et à jour, sinon I/O réseau/disque en chute libre.

## 185. SR-IOV et passthrough PCIe : quand les utiliser

| Technique | Principe | Usage |
|---|---|---|
| SR-IOV | une carte réseau se présente comme N cartes virtuelles | NFV, telco, trading |
| Passthrough (VFIO) | une carte dédiée à une VM | GPU, FPGA, cartes crypto |

Conditions : IOMMU activé (BIOS), carte et pilote compatibles, hyperviseur
qui le supporte. Gain : latence réseau divisée par 2-5, débit proche du
bare metal. Coût : la VM devient liée au nœud physique (plus de vMotion)
— **à réserver aux workloads qui le justifient**, pas par défaut.

## 186. DPDK : le réseau sans le kernel

DPDK (Data Plane Development Kit) traite les paquets en espace utilisateur,
sans passer par la pile réseau du kernel : 10-100× plus de paquets/seconde
par cœur.

Usage : pare-feu, routeurs virtuels, sondes, telco (avec les accélérateurs
DLB/QAT des Xeon, section 18). Exigences : cœurs dédiés (isolés via
`isolcpus`), cartes réseau compatibles, hugepages. **C'est une
spécialité** : ne l'envisagez que si votre équipe maîtrise déjà le réseau
haute performance.

## 187. Benchmarks de réception : protocole type

À chaque livraison de serveur, avant mise en production :
1. **memtest86+** (1 passe complète) — RAM.
2. **stress-ng / mprime** (1 h) — stabilité CPU + thermique (surveillez les
   températures, section 67).
3. **fio** (lecture/écriture aléatoire et séquentielle) — NVMe, alignement
   NUMA (section 117).
4. **iperf3** — débit réseau réel (pas le chiffre du port).
5. **Vérification ECC** : injection d'erreur si le BIOS le permet
   (section 147).
6. **Photo + fiche** : config as-built archivée.

Durée : ~1 journée par serveur, en grande partie automatisable (PXE +
scripts). **Un serveur non testé à réception est un serveur dont on
découvre les défauts en production.**

---

## 188. NVMe : bien dimensionner le stockage serveur

