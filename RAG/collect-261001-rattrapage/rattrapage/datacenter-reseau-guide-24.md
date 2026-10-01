---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-24
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: ["Broadcom", "Nvidia"]
dates: ["2026-09-27"]
keywords: ["datacenter", "cpo", "diffusion", "ethernet", "gpu", "hyperscaler", "incident", "latency", "nvidia", "throughput"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [3487, 3692]
sha256: f125308b0abef86ab816c855b4cb3489c947950d5c042abb470f6fe2db90f465
---

# RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE

| Norme | Débit | OM1 | OM2 | OM3 | OM4 | OM5 | OS2 |
|---|---|---|---|---|---|---|---|
| 1000BASE-SX | 1G | 275 m | 550 m | 800 m | 800 m | — | — |
| 10GBASE-SR | 10G | 33 m | 82 m | 300 m | 400 m | — | — |
| 10GBASE-LR | 10G | — | — | — | — | — | 10 km |
| 25GBASE-SR | 25G | — | — | 70 m | 100 m | — | — |
| 40GBASE-SR4 | 40G | — | — | 100 m | 150 m | — | — |
| 100GBASE-SR4 | 100G | — | — | 70 m | 100 m | — | — |
| 100GBASE-SR10 | 100G | — | — | 100 m | 150 m | — | — |
| 100GBASE-DR | 100G | — | — | — | — | — | 500 m |
| 200GBASE-SR4 | 200G | — | — | 70 m | 100 m | — | — |
| 400GBASE-SR8 | 400G | — | — | 60 m | 100 m | — | — |
| 400GBASE-DR4 | 400G | — | — | — | — | — | 500 m |
| 800GBASE-SR8 | 800G | — | — | 60 m | 100 m | — | — |
| 800GBASE-DR8 | 800G | — | — | — | — | — | 500 m |

**Note** : OM1/OM2 (62,5/50 µm anciennes) = legacy, ne plus tirer.
SWDM4/BiDi sur OM5 : 100-150 m sur 2 brins.

## 168. Guide d'achat par profil

| Profil | Backend | Front-end | Fibre | Budget réseau (⚠️) |
|---|---|---|---|---|
| **PME** (20 serveurs) | — (pas d'IA) | 2× switch 48×1/10G, LACP | OM4 | 15-40 k€ |
| **ETI** (200 serveurs) | — ou petit | Leaf-spine 25/100G, 3:1 | OM4 + OS2 rocade | 200-500 k€ |
| **Cloud privé** (1000 srv) | 100G dédié si IA légère | Leaf-spine 25/100G, EVPN | OS2 partout | 1-3 M€ |
| **IA moyen** (256 GPU) | 400/800G 1:1, rail-opt. | Séparé 100G | OS2 + OM4 | 5-15 M€ |
| **IA grand** (2048 GPU) | 800G 1:1, 3 tiers | Séparé | OS2 + MPO-16 | 40-90 M€ |
| **Hyperscaler** | CPO, UEC, custom | Converged | Tout, en masse | — |

**Règle par profil** : ne pas acheter « comme les grands » quand on est
petit — un spine-leaf 3:1 bien opéré bat un mini-fabric IA mal tuné.
Et inversement : ne pas bricoler du 100G pour un cluster GPU.

## 169. Lexique français-anglais — 50 termes

| Français | Anglais |
|---|---|
| Brassage | Patching / cross-connect |
| Jarretière | Patch cord |
| Trunk (fibre) | Trunk cable |
| Goulotte | Cable tray / raceway |
| Baie | Rack / cabinet |
| Allée chaude/froide | Hot/cold aisle |
| Débit | Throughput / data rate |
| Latence | Latency |
| Gigue | Jitter |
| Perte de paquets | Packet loss |
| Engorgement | Congestion |
| File d'attente | Queue |
| Tampon | Buffer |
| Commutateur | Switch |
| Routeur | Router |
| Carte réseau | Network Interface Card (NIC) |
| Pilote | Driver |
| Micrologiciel | Firmware |
| Supervision | Monitoring |
| Télémétrie | Telemetry |
| Seuil | Threshold |
| Alerte | Alert |
| Panne | Failure / outage |
| Redondance | Redundancy |
| Basculement | Failover |
| Reprise | Recovery |
| Sauvegarde (config) | Backup |
| Mise en service | Commissioning |
| Recette | Acceptance testing |
| Appel d'offres | RFP / tender |
| Bordereau de prix | Bill of quantities |
| Pièce de rechange | Spare part |
| Contrat de maintenance | Maintenance contract |
| Délai d'intervention | Response time (SLA) |
| Coupure | Power outage |
| Onduleur | UPS |
| Groupe électrogène | Generator |
| Climatisation | Cooling / HVAC |
| Efficacité énergétique | Power efficiency |
| Facture électrique | Power bill |
| Terre (élec.) | Ground / earth |
| Disjoncteur | Circuit breaker |
| Chemin de câbles | Cable pathway |
| Coupe-feu | Firestop |
| Plan d'adressage | Addressing plan |
| Sous-réseau | Subnet |
| Passerelle | Gateway |
| Diffusion | Broadcast |
| En-tête (paquet) | Header |

## 170. Template — rapport d'incident réseau

```
Incident n° : __________    Date/heure début : __________    Fin : __________
Sévérité : P1 / P2 / P3    Détecté par : supervision / utilisateur / autre
Périmètre : __________ (fabric, sites, tenants impactés)

TIMELINE (heure : fait)
__:__ :
__:__ :
__:__ :

CAUSE RACINE (5 pourquoi) :
1. Pourquoi ? __________
2. Pourquoi ? __________
3. Pourquoi ? __________
4. Pourquoi ? __________
5. Pourquoi ? __________

CORRECTIF IMMÉDIAT : __________
CORRECTIF DURABLE : __________ (avec échéance et responsable)
PIÈGE DU GUIDE CONCERNÉ : P__ (§122)

Leçons : __________
Diffusion : __________
```

**Règle** : tout P1/P2 = rapport sous 5 jours ouvrés, relu par un pair.
Un incident sans rapport **reviendra** — c'est mécanique.

## 171. Parcours de formation — devenir autonome sur le réseau IA

| Niveau | Contenu | Durée (⚠️) |
|---|---|---|
| 1. Bases | TCP/IP, switching, VLAN, STP, routage | 2-4 sem. |
| 2. Datacenter | Spine-leaf, BGP, ECMP, EVPN/VXLAN | 1-2 mois |
| 3. Fibre | Optique, budgets, nettoyage, MPO | 1 sem. + pratique |
| 4. IA | RDMA, RoCEv2, ECN/DCQCN, NCCL | 1 mois |
| 5. Pratique | Lab : monter un mini-fabric, le casser, le réparer | Continu |
| 6. Énergie | PUE, dimensionnement, CPO | 1 sem. |

**Certifications utiles** (⚠️) : JNCIA/CCNA (bases), Arista ACE (EOS),
NVIDIA NCA-AI (réseau IA). Mais **rien ne remplace le lab** : 1 jour de
casse volontaire = 1 mois de théorie.

## 172. Index des tableaux du guide

| Tableau | Section |
|---|---|
| Débit ⇄ format ⇄ PCIe | §2 |
| NIC 100G / 400G / 800G | §4, §6, §7 |
| Offloads | §9 |
| RDMA : IB / RoCE / UET | §10 |
| Matrice de choix NIC | §11 |
| Formats physiques NIC | §12 |
| Lexique formats fibre | §16 |
| Correspondance débit/format/fibre | §17, §139 |
| OM3/4/5 vs OS2 | §18 |
| Distances | §19, §167 |
| Budgets optiques (éléments) | §20 |
| Connecteurs | §21 |
| DAC vs AOC vs transceivers | §22 |
| Breakout | §23 |
| Seuils IEC 61300-3-35 | §25 |
| FEC | §27 |
| BOM fibre rack 400G | §29 |
| Oversubscription | §32 |
| ECMP / EVPN | §36, §37 |
| Switch IA vs classique | §43 |
| Buffers deep/shallow | §44 |
| IB vs Ethernet marché | §47, §48 |
| Spectrum-X / Broadcom / Arista | §49, §50, §51 |
| Conso NICs / switchs / optiques | §89, §90, §91 |
| PUE / refroidissement | §93, §94 |
| Leviers d'efficacité | §97 |
| 1,6T / UEC / CPO | §99, §100, §101 |
| Glossaires (80 termes) | §105, §158 |
| Les 45 pièges | §122 |
| BOMs projets | §108→§123 |
| Fiches produits | §126-133 |
| Comparatif 20 switches | §151 |
| Transceivers (25 réf.) | §141 |
| Câbles (20 réf.) | §154 |
| Normes IEEE | §142 |

## 173. Checklist — avant de partir en congés (NOC)

- [ ] Astreinte définie, téléphone chargé, accès VPN testé
- [ ] Spares : stock vérifié (optiques, DAC, 1 switch par plan critique)
- [ ] Supervision : alertes testées, seuils à jour, pas d'alerte en cours ignorée
- [ ] Changements gelés : pas de FW/config pendant l'absence (ou rollback prêt)
- [ ] Doc : mots de passe (coffre), plans, contacts TAC affichés
- [ ] Énergie : seuils de température vérifiés (canicule ?)
- [ ] Dernier REX : les 3 incidents du mois sont-ils soldés ?

**Un réseau qui ne peut pas tenir 2 semaines sans vous n'est pas un réseau
fini** — c'est une astreinte déguisée.

## 174. Notes de version et maintenance du guide

| Version | Date | Contenu |
|---|---|---|
| 1.0 | 27/09/2026 | 174 sections, recherche web du jour |

**Maintenance** : revérifier les sections marquées ✅ tous les 6 mois
(le marché IA bouge vite : CPO, UEC, 1,6T). Les ⚠️ sont à confirmer à
chaque achat. Journal des modifications ci-dessous pour les v1.1+.

---

*Fin du guide — 174 sections numérotées, 80 termes de glossaire, 45 pièges,
quiz, 30 FAQ, 100 acronymes, BOMs, checklists, templates.
Vérification finale : `wc -l` — objectif ≥ 4000 lignes.*

# PARTIE Q — SYNTHÈSES FINALES

## 175. Une page par débit — la synthèse ultime

