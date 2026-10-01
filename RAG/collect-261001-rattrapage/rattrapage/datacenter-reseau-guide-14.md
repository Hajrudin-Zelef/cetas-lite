---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-14
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "capex", "dci", "ethernet", "gpu", "training"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [1937, 2089]
sha256: 78b93a1bcac21fe58d3bd005d73c770f633b1457c8ee134a467316026f4dc52b
---

# RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE

**Chiffres** : sur un lien 400G, NCCL atteint ~90-95 % du débit utile en
all-reduce bien réglé (busbw). En dessous de 70 % : problème réseau
(ECMP, CC, NUMA) — pas problème GPU.
**Variables d'environnement clés** :
- `NCCL_SOCKET_IFNAME` : forcer l'interface backend (pas le front-end !).
- `NCCL_IB_DISABLE=0/1` : choisir IB vs Ethernet.
- `NCCL_DEBUG=INFO` : voir quelle interface/NIC est utilisée — **le premier
  debug quand « c'est lent »**.
**Piège classique** : NCCL qui passe par le réseau front-end 100G au lieu
du backend 400G parce que l'IP du backend n'est pas dans le bon sous-réseau
(P40, §122). Symptôme : 4× moins vite que prévu, pile le ratio 400/100.

## 109. Stockage et réseau : NVMe-oF, RDMA et dimensionnement

Le stockage moderne (NVMe-oF sur RoCEv2) partage le réseau — ou pas :

| Architecture | Réseau stockage | Avantage | Inconvénient |
|---|---|---|---|
| Fabric dédié | 100/200G séparé | Isolation, pas de contention | 2ᵉ fabric à payer |
| Converged | Même fabric, QoS | 1 seul réseau | Le training peut noyer le stockage |
| DPU offload | NVMe-oF sur DPU | CPU libéré (§67) | Coût DPU |

**Dimensionnement** : 1 SSD NVMe Gen4 ≈ 7 Go/s ≈ 56 Gb/s. Une baie de
24 SSD = **1,3 Tb/s** théoriques → il faut **4× 400G** ou 2× 800G par baie
pour ne pas brider (en pratique : 2× 400G avec un ratio réaliste, les SSD
ne poussent pas tous à fond en même temps).
**RDMA obligatoire** : NVMe-oF sur TCP à 100G+ = CPU saturé ; sur RoCEv2 =
zéro-copy, latence ~10 µs de bout en bout (⚠️).
**Checkpoints IA** : un training qui checkpoint 1 To toutes les heures sur
un stockage à 100 Go/s = 10 s de pause — dimensionner le stockage
**pour le checkpoint**, pas pour la moyenne (P41, §122).

## 110. Sécurité du fabric : segmentation, ACLs, zero-trust

Le réseau IA est souvent « plat et ouvert » par performance — c'est une
dette sécurité :

| Mesure | Où | Coût perf |
|---|---|---|
| Management VRF séparé | Tous switchs | Nul |
| ACLs infra (n'accepter que le NOC) | Interfaces de management | Nul |
| Secure boot + image signée | Switchs/NIC | Nul (au boot) |
| 802.1X sur ports serveurs | Leafs | Faible (au link-up) |
| Micro-segmentation (DPU) | Serveurs multi-tenant | Faible (HW) |
| Chiffrement MACsec 400/800G | Liens inter-sites | Faible (HW si supporté) |
| Chiffrement IPsec inline (CX-8 ✅) | Backend sensible | Nul (HW) |

**Règles** :
- **Jamais de management sur le fabric data** : un plan OOB dédié (1G/25G),
  même en petit DC. Le jour où le fabric data tombe, c'est l'OOB qui permet
  de le réparer.
- **AAA** (TACACS+/RADIUS) partout, pas de compte `admin/admin` — auditer
  avec un scan trimestriel.
- **Firmware signé** : exiger la vérification de signature sur switchs et
  DPU (supply-chain).
- **Ne pas chiffrer le backend IA par défaut** : le chiffrement logiciel
  tue la perf NCCL ; le HW (CX-8) seulement si l'exigence est réelle
  (données de santé, défense…).

## 111. Automation : ZTP, Ansible, NetBox, CI/CD réseau

| Couche | Outil | Usage |
|---|---|---|
| Source de vérité | **NetBox** | IP, baies, câbles, ASN — tout y est |
| Provisioning | ZTP (EOS, Cumulus, NX-OS) | Switch branché → config auto |
| Config management | Ansible / Nornir | Templates Jinja, idempotence |
| Validation | Batfish / SuzieQ | « La config fait-elle ce qu'on croit ? » |
| CI/CD | Git + pipeline | MR → lab → prod, rollback auto |

**Workflow type** :
1. NetBox = entrée (nouveau leaf : baie, ports, ASN).
2. Ansible génère la config depuis les templates + NetBox.
3. Batfish valide (pas de route qui fuit, pas d'ACL qui bloque le BGP).
4. Déploiement ZTP ou Ansible, vérification post-change (tests §107).
5. Tout est commité : **le réseau se versionne comme du code**.

**Piège** : automatiser un process cassé = casser plus vite. Stabiliser
manuellement sur 1 pod, **puis** automatiser (P42, §122).
**Lien avec les guides existants** : Ansible (guide dédié), NetBox (guide
dédié) — ce guide suppose ces bases acquises.

## 112. CAS CHIFFRÉ — TCO 5 ans : fabric 400G vs 800G (128 nœuds GPU)

Hypothèses : 128 nœuds × 8 GPU, 1 NIC/GPU, 1:1, élec 0,15 €/kWh (⚠️),
PUE 1,4 (⚠️). Prix = fourchettes marché ⚠️ à vérifier.

| Poste (5 ans) | 400G (1024 NIC CX-7) | 800G (1024 NIC CX-8) |
|---|---|---|
| Switchs (16× 64×400G vs 8× 64×800G) | 16 × 80-150 k€ = 1,3-2,4 M€ | 8 × 150-300 k€ = 1,2-2,4 M€ |
| NIC | 1024 × 1,5-3 k€ = 1,5-3,1 M€ | 1024 × 3-6 k€ = 3,1-6,1 M€ |
| Câbles/optiques | ~0,5-1 M€ | ~0,8-1,5 M€ |
| **CAPEX** | **~3,3-6,5 M€** | **~5,1-10 M€** |
| Élec réseau (5 ans) | ~90 kW × 1,4 × 8760 h × 5 × 0,15 € ≈ 0,83 M€ | ~154 kW → ≈ 1,42 M€ |
| Maintenance (15 %/an ⚠️) | ~2,5-4,9 M€ | ~3,8-7,5 M€ |
| **TCO 5 ans** | **~6,6-12,2 M€** | **~10,3-19 M€** |

**Lecture** : le 800G coûte ~1,6× plus cher mais pousse 2× plus de bits :
**le coût par Gb/s sur 5 ans est inférieur en 800G** (~−20 %, ⚠️). Et le
400G sera en fin de vie avant 5 ans (cycle 3-4 ans, §102) → prévoir un
refresh. **Décision** : cluster neuf en 2026 → 800G directement, sauf
contrainte budgétaire CAPEX dure.

## 113. Migration 100G → 400G → 800G : stratégie par étapes

| Étape | Action | Durée type |
|---|---|---|
| 1. Audit | Inventaire ports, PCIe, fibre (OM3 ? OS2 ?) | 2-4 sem. |
| 2. Fibre d'abord | Tirer OS2/OM4 là où il manque — **avant** les switchs | 1-3 mois |
| 3. Breakout | Switchs 400G en 4×100G vers serveurs existants | Au fil de l'eau |
| 4. Backend IA | Nouveau fabric 400/800G **séparé** (§32) | Projet |
| 5. Bascule | Migrer les workloads, pas les câbles | Par vague |
| 6. Décommission | Revendre/déclasser le 100G en OOB | Fin |

**Règles** :
- **Ne jamais mélanger les générations sur un même fabric** : un lien 100G
  dans un fabric 400G = ECMP asymétrique = polarisation (§36).
- La fibre se tire **une fois pour 15 ans**, les switchs se changent tous
  les 3-4 ans : **surdimensionner la fibre** (24-48 brins par trunk même si
  on n'en utilise que 8).
- **Breakout = l'arme de la migration douce** : un port 400G qui sert 4
  serveurs 100G rentabilise le switch avant la fin de la migration.

## 114. Multi-site / DCI : interconnecter des datacenters

| Distance | Technologie | Débit | Note |
|---|---|---|---|
| < 2 km (campus) | 400G-FR4 / 800G-2×FR4 | 400/800G | Fibre propre, pas d'ampli |
| 2-10 km (métro) | 400G-LR4 / 800G-DR8+ | 400/800G | — |
| 10-80 km (régional) | **400G-ZR / 800G-ZR** (cohérent) | 400/800G | Optique cohérente enfichable ✅ |
| > 80 km | DWDM + transpondeurs | N×400/800G | Baie optique dédiée |

**400G-ZR / 800G-ZR** : optiques cohérentes **enfichables directement dans
le routeur/switch** (pas de transpondeur externe) — a divisé par ~3 le coût
du DCI métro (⚠️). Standard : OpenZR+ / OIF.
**Règles DCI** :
- **Chiffrer** (MACsec ou IPsec) : la fibre inter-sites n'est pas sous votre
  contrôle physique.
- **Latence** : 1 ms ≈ 200 km de fibre — au-delà de ~50 ms RTT, le
  synchrone (stretch cluster) devient impossible.
- **Ne pas étendre le fabric IA entre sites** : l'all-reduce inter-sites
  s'effondre avec la latence — chaque site = 1 fabric, réplication
  asynchrone entre sites (P43, §122).

## 115. Edge IA et 5G : contraintes spécifiques

| Contrainte | Edge / 5G | Datacenter |
|---|---|---|
| Profondeur de baie | 600 mm (télécom) | 1000-1200 mm |
| Température | 35 °C ambiant (clim confort) | 25 °C contrôlés |
| Accès | Site distant, 4 h d'intervention | Sur place |
| Alimentation | −48 V DC (télécom) ou 230 V | 400 V / busway |
| Poussière | Possible (shelter) | Salle blanche |

