---
id: collect-261001-rattrapage/rattrapage/datacenter-cpu-ram-guide-21
title: "CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: ["2026-06-01", "2026-07-23", "2026-09-27"]
keywords: ["capex", "clearwater forest", "datacenter", "ethernet", "gpu", "inference", "intel", "nvlink", "training"]
source: docs/RAG/collect-261001-rattrapage/datacenter_cpu_ram_guide.md
source_anchor: ""
source_lines: [3349, 3520]
sha256: 5c47256c29bc30e5fa272bbd6722879a649363a152fd4209d419d8ea3bfe1307
---

# CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION

| Critère | Repère |
|---|---|
| Endurance (DWPD) | ≥ 1 (virtualisation), ≥ 3 (DB/write-intensive) |
| Latence | < 100 µs en lecture aléatoire (datacenter NVMe) |
| Format | U.2/U.3 (hot-swap) en serveur ; M.2 pour boot uniquement |
| PCIe | Gen4 minimum en 2026, Gen5 pour les exigeants |
| Over-provisioning | 10-20 % réservés (endurance et perf. stables) |

Règles :
- **Séparez les usages** : OS/boot, données, WAL/journal, cache — des
  pools ou des disques distincts, pas tout sur le même volume.
- **Le RAID n'est pas une sauvegarde** (et le RAID5/6 sur gros SSD pose
  des questions de rebuild : préférez miroir ou erasure coding distribué).
- Surveillez l'**usure** (SMART : pourcentage used) : un NVMe à 90 % d'usure
  se remplace **avant** la panne, en heures ouvrées.

## 189. Débit PCIe par génération : mémo

| Génération | Débit/ligne | ×16 (GPU/NIC) | ×4 (NVMe) |
|---|---|---|---|
| Gen3 | 8 GT/s (~1 Go/s) | ~16 Go/s | ~4 Go/s |
| Gen4 | 16 GT/s (~2 Go/s) | ~32 Go/s | ~8 Go/s |
| Gen5 | 32 GT/s (~4 Go/s) | ~64 Go/s | ~16 Go/s |
| Gen6 | 64 GT/s (~8 Go/s) | ~128 Go/s | ~32 Go/s |

Un NVMe Gen4 (7 Go/s) sature un lien ×4 Gen4 : le passer en Gen5 ne sert à
rien sans disque Gen5. **Alignez les générations** (CPU, slot, carte,
disque) au lieu de payer une génération inutilisée.

## 190. Réseau serveur : 25/100/200 GbE, que choisir ?

| Débit | Usage 2026 | Quand |
|---|---|---|
| 2× 25 GbE | standard virtualisation | défaut raisonnable |
| 2× 100 GbE | DB, SDS, virtualisation dense | > 50 VM exigeantes/nœud |
| 200 GbE / InfiniBand | HPC (MPI), IA | calcul distribué |
| 400/800 GbE | backbone, futurs GPU | à l'étude, pas en serveur standard |

Règles :
- **Séparez les plans** : production, management (BMC), vMotion/réplication,
  stockage — VLANs au minimum, liens physiques séparés pour le critique.
- **2 liens minimum** en LACP ou actif/passif : une carte ou un câble qui
  lâche ne doit jamais isoler un serveur.
- Le réseau de **stockage** (Ceph, vSAN) est le plus gourmand : 25 GbE
  minimum par nœud, 100 GbE dès que les NVMe sont nombreux.

## 191. RDMA, RoCE, InfiniBand : panorama

| Technologie | Principe | Usage |
|---|---|---|
| InfiniBand | réseau dédié HPC, très faible latence | HPC, IA training |
| RoCE v2 | RDMA sur Ethernet | SDS, HPC sur Ethernet |
| TCP classique | pile standard | tout le reste |

Le RDMA permet aux cartes réseau d'accéder directement à la mémoire
distante sans passer par le CPU : indispensable en HPC (MPI) et de plus en
plus en stockage (NVMe-oF). **Coût** : switches et cartes spécifiques,
compétences réseau pointues. Pour de la virtualisation classique, c'est
inutile.

## 192. Câbles DAC, AOC, optique : que choisir en baie

| Type | Distance | Coût (indicatif) | Usage |
|---|---|---|---|
| DAC (cuivre) | ≤ 3-5 m | ~30-80 € | ToR dans la même baie |
| AOC (fibre active) | ≤ 30 m | ~150-300 € | inter-baies proches |
| Optique (transceivers) | > 30 m | ~200-600 €/paire | campus, longues distances |

En baie : **DAC par défaut** (pas cher, fiable, faible latence). Vérifiez
la **compatibilité** (codage du constructeur du switch) : un DAC non codé
pour votre switch peut être rejeté au link-up.

## 193. Switch top-of-rack (ToR) : dimensionnement

| Critère | Repère |
|---|---|
| Ports | nb serveurs × liens + 20 % de réserve + uplinks |
| Débit | non-bloquant sur les ports serveurs (1:1) |
| Uplinks | ≥ 2× la somme des ports en oversubscription 3:1 max |
| Redondance | 2 ToR en MLAG/vPC par baie critique |
| Management | hors-bande (réseau BMC séparé) |

Exemple : 20 serveurs × 2× 25 GbE = 40 ports → switch 48× 25 GbE + 6× 100 GbE
uplinks, en double pour la redondance. **Le switch ToR redondant coûte
moins cher qu'une heure d'indisponibilité de la baie.**

## 194. Tableau : bande passante réseau vs stockage

| Config nœud | Débit disque agrégé | Réseau minimal |
|---|---|---|
| 4× NVMe SATA-like | ~2 Go/s | 25 GbE |
| 8× NVMe Gen4 | ~20 Go/s | 2× 100 GbE |
| 8× NVMe Gen5 | ~40 Go/s | 400 GbE ou RDMA |

Si votre stockage local peut sortir 20 Go/s mais que votre réseau n'en
transporte que 3 Go/s (25 GbE), la réplication Ceph/vSAN sera bridée par le
réseau, pas par les disques. **Dimensionnez le réseau sur le stockage,
pas l'inverse.**

---

## 195. Synthèse : l'arbre de décision en une page

```
 BESOIN ?
 |-- Virtualisation generaliste --> EPYC Zen5c / Xeon 6700P, 4-8 Go/coeur,
 |                                   12 canaux remplis, 2U, ~700-900 W
 |-- Base de donnees OLTP ---------> CPU "F" haute frequence, 1S,
 |                                   dataset + 30 % + croissance (sec. 50)
 |-- Base analytique / HPC --------> Xeon 6980P + MRDIMM-8800 (ou Venice 2027),
 |                                   BP/coeur >= 5 Go/s, DLC si 500 W
 |-- Inference IA -----------------> GPU d'abord ; sinon Xeon P-core (AMX)
 |-- Scale-out / cloud -------------> Xeon 6+ / EPYC 9965, densite max,
 |                                   perf/watt, 1U si salle adaptee
 |-- Edge -------------------------> EPYC 8004 / Xeon 6500, TDP contenu
 |
 +--> TOUJOURS : licences chiffrees d'abord, canaux remplis,
      onduleur sur pic +25 %, PUE suivi, TCO 5 ans.
```

## 196. Les 10 chiffres à retenir par cœur

1. **614 Go/s** : bande passante théorique d'un socket 12 canaux DDR5-6400.
2. **3,2 Go/s/cœur** : le ratio d'un EPYC 9965 — dense mais « affamé ».
3. **500 W** : le TDP des amiraux 2026 (9965, 9755, 6980P, 6990E+).
4. **8800 MT/s** : MRDIMM Gen1 — +24 % mesurés sur HPCG, +437 W mesurés.
5. **1,5** : le PUE à ne pas dépasser en rénovation.
6. **300 W** : le socle idle d'un 2S moderne — éteignez l'inutile.
7. **25 %** : la marge minimale sur le dimensionnement onduleur.
8. **12** : le nombre de barrettes pour remplir un socket SP5.
9. **5 ans** : l'horizon du TCO — le CAPEX n'est que le début.
10. **2** : le nombre de devis minimum… non : **3 devis minimum.**

## 197. Ce que ce guide ne couvre pas (périmètre assumé)

- Le détail des **GPU** (un guide dédié serait nécessaire : H100/H200,
  MI300, consommation 700-1000 W/GPU, NVLink).
- Le **stockage SAN/NAS** d'entreprise (baies, réplication).
- La **sécurité** avancée (chiffrement, HSM, segmentation).
- Les **CPU ARM** serveur en détail (section 137).
- Les **prix exacts** : le marché fluctue, les devis font foi.

## 198. Historique des vérifications de ce guide

| Date | Objet |
|---|---|
| 27/09/2026 | EPYC 9005 Turin : références, TDP, prix (lancement 10/2024) |
| 27/09/2026 | Xeon 6 6700/6900 : P-core/E-core, sockets, TDP |
| 27/09/2026 | Xeon 6+ Clearwater Forest : lancement 01/06/2026, specs, prix |
| 27/09/2026 | EPYC 9006 Venice : annonce 22-23/07/2026, specs, dispo T4 2026 (presse) |
| 27/09/2026 | Diamond Rapids (Xeon 7) : 2027, PCIe 6.0 (confirmé, détails à venir) |
| 27/09/2026 | MRDIMM : Gen1 8800 en production, Gen2 12800 T1 2027, mesures Phoronix |
| 27/09/2026 | DDR6 : JEDEC non finalisée, production 2028-2029 |
| 27/09/2026 | CXL : 2.0 en production (Azure 11/2025), 3.x en échantillonnage |

## 199. Comment maintenir ce guide à jour

- **Chaque trimestre** : vérifier les prix CPU/RAM (marché volatil).
- **À chaque annonce** (Advancing AI, Computex, Intel Innovation) :
  mettre à jour la partie I et les tableaux de références.
- **Annuellement** : relire les cas chiffrés énergie avec le prix du kWh
  réel et le PUE mesuré de vos salles.
- **En continu** : noter les retours terrain (section 140) — un guide qui
  n'apprend pas de ses pannes est un catalogue.

## 200. Le mot de la fin : l'énergie est le premier critère

On achète un serveur pour sa puissance ; on le paie pour son énergie.
Entre deux configurations à performance égale, **la moins gourmande gagne
toujours au TCO** — et c'est aussi celle qui vieillit le mieux (moins de
chaleur = moins de pannes = moins de bruit).

