---
id: collect-261001-rattrapage/rattrapage/datacenter-builds-guide-16
title: "Datacenter Builds — Le guide des BOMs"
domain: rattrapage
role: reference
task: reference
actors: ["Nvidia"]
dates: ["2026-09-27"]
keywords: ["datacenter", "arr", "blackwell", "capex", "compute", "cost", "distribution", "dram", "ethernet", "fp8", "gpu", "kv cache"]
source: docs/RAG/collect-261001-rattrapage/datacenter_builds_guide.md
source_anchor: ""
source_lines: [2820, 2950]
sha256: a639a764e8af46f711fafcba4a3cec2376f1dc308f4fe8363927ca9fd9de8cee
---

# Datacenter Builds — Le guide des BOMs

**À NE PAS copier :**
- La densité extrême (100 kW/rack) : sans l'équipe et l'infra qui vont avec,
  c'est un incendie en attente.
- Le hardware custom : pas de pièces de rechange, pas de support.
- L'optimisme des délais GPU : eux ont des allocations, vous non.

---

## 179. Section « À venir » — annonces vérifiées au 27/09/2026

> Uniquement des annonces officielles ou des faits sourcés. Le reste est
> marqué « non trouvé » ou « à vérifier ».

1. **800VDC commercial** : Vertiv, Schneider Electric, Eaton, Delta —
   produits attendus H2 2026 (source : analyses industrie, §176).
2. **Thales Luna 8** : HSM FIPS 140-3 niveau 3, support post-quantique,
   disponible EMEA en 2026 (annonce Thales — vérifié le 27/09/2026).
3. **Densités Rubin** : racks 100–240 kW en préparation (trajectoire
   documentée par les équipementiers — à vérifier au cas par cas).
4. **Immersion monophasée** : progression chez les hyperscalers, le diphasique
   freiné par la réglementation PFAS (reporting EPA janvier 2027).
5. **DLC monophasé 55 % du marché DTC** : le standard de facto pour les
   nouveaux builds IA (vérifié le 27/09/2026).
6. **Délais GPU** : normalisation H100 (2–3 mois), tension persistante sur
   Blackwell PRO (3–7 mois début 2026 — à revérifier au devis).
7. **Prix DRAM** : TrendForce projette +13 à 18 % au T3 2026 sur la DRAM
   serveur — **achetez la RAM tôt** (vérifié le 27/09/2026).

Non trouvé au 27/09/2026 (prix publics) : MI325X/MI355X standalone, serveurs
8-GPU HGX en prix catalogue, PDU 800VDC en vente libre.

---

## 180. Veille : la méthode

- **Trimestrielle** : prix DRAM (TrendForce), GPU (snapshots marché),
  annonces refroidissement (AFCOM, Data Center World).
- **À chaque achat** : 3 devis, délais écrits, firmware notés.
- **Ce guide** : relisez la section 3 avant chaque budget annuel — le marché
  2026 bouge par trimestre, pas par an.

---

## Glossaire (35 termes)

| Terme | Définition |
|---|---|
| **BOM** | Bill of Materials : liste complète des composants d'un build, avec quantités et prix. |
| **TDP** | Thermal Design Power : puissance thermique max d'un composant, base du dimensionnement électrique et froid. |
| **PUE** | Power Usage Effectiveness : énergie totale / énergie IT. 1,0 = parfait, 1,5 = correct, 2,0 = à améliorer. |
| **DWPD** | Drive Writes Per Day : endurance d'un SSD (1 = on peut réécrire tout le disque 1×/jour pendant la garantie). |
| **RDIMM** | Mémoire serveur ECC à registres (buffered). Obligatoire en serveur. |
| **DPC** | DIMM Per Channel : nombre de barrettes par canal mémoire (1 DPC = pleine vitesse). |
| **NUMA** | Non-Uniform Memory Access : en bi-socket, la mémoire de l'autre CPU est plus lente. |
| **U.2** | Format SSD 2,5" NVMe hot-swap, standard datacenter. |
| **E3.S / EDSFF** | Format SSD nouvelle génération, plus dense que l'U.2. |
| **OCP 3.0** | Format de carte réseau mezzanine standard serveurs récents. |
| **RoCE** | RDMA over Converged Ethernet : accès mémoire distant sur Ethernet, pour Ceph/HPC. |
| **DAC** | Direct Attach Cable : câble cuivre avec transceivers intégrés, ≤ 5–7 m. |
| **AOC** | Active Optical Cable : câble optique actif, souple, jusqu'à 30 m+. |
| **ToR** | Top of Rack : switch en haut du rack, 1–2 par rack. |
| **MLAG** | Multi-Chassis Link Aggregation : 2 switchs vus comme un seul (redondance sans STP). |
| **PDU** | Power Distribution Unit : multiprise de rack intelligente (metered/switched). |
| **ATS** | Automatic Transfer Switch : bascule auto entre 2 arrivées électriques. |
| **N+1** | Redondance : N équipements pour la charge + 1 de secours. |
| **CARP** | Common Address Redundancy Protocol : VIP partagée entre 2 firewalls (HA). |
| **RDHx** | Rear Door Heat Exchanger : porte arrière à eau qui refroidit l'air sortant du rack. |
| **DLC** | Direct Liquid Cooling : refroidissement liquide direct des puces (cold plates). |
| **CDU** | Coolant Distribution Unit : échangeur qui sépare boucle bâtiment et boucle techno. |
| **PFAS** | « Polluants éternels » : freinent le refroidissement diphasique (réglementation 2027). |
| **EC 4+2** | Erasure Coding : 4 fragments data + 2 parité (économe en capacité, Ceph). |
| **RBD** | RADOS Block Device : volumes bloc Ceph. |
| **BlueStore** | Moteur de stockage Ceph (remplace FileStore), avec DB/WAL sur NVMe. |
| **OSD** | Object Storage Daemon : processus Ceph par disque. |
| **MON/MGR** | Moniteurs et managers Ceph (cluster, placement, métriques). |
| **vGPU** | GPU virtualisé NVIDIA (vPC/vWS) : partage d'un GPU entre VMs. |
| **EPS** | Événements Par Seconde : métrique de dimensionnement SIEM. |
| **HSM** | Hardware Security Module : boîtier inviolable pour les clés cryptographiques. |
| **PQC** | Post-Quantum Cryptography : algos résistants à l'ordinateur quantique (ML-KEM, ML-DSA). |
| **TCO** | Total Cost of Ownership : CAPEX + 5 ans d'OPEX (énergie, maintenance, licences). |
| **Burn-in** | Test de vieillissement accéléré (48–72 h) avant mise en prod. |
| **As-built** | Dossier « tel que construit » : plans, versions, mots de passe, procédures. |

---

## Quiz — 10 questions + réponses

**Q1.** Un module DDR5 RDIMM 64 Go coûtait ≈ 1 500 $ en contrat en septembre 2026.
Quel était son prix un an plus tôt, et quelle en est la conséquence sur les BOMs ?
**R1.** ≈ 272 $ (5,5× moins). Conséquence : la RAM est redevenue le premier poste
de coût d'un serveur — on la dimensionne au working set mesuré, pas « au cas où ».

**Q2.** Un serveur affiche : 2× CPU 400 W, 12 DIMM, 6 NVMe, 2 NIC 100G, base 150 W.
Calculez P_max avec la formule §13.
**R2.** 800 + 120 + 108 + 85 + 150 ≈ 1 260 W (proche du Compute M, §21).

**Q3.** Pourquoi le WAL PostgreSQL va-t-il sur des disques dédiés ?
**R3.** Les écritures séquentielles du WAL ne sont pas perturbées par les lectures
aléatoires des données : +20 à 40 % de TPS sur charges à écritures intenses (§28).

**Q4.** 5 nœuds Ceph de 123 To bruts en réplica 3 : combien de To utiles ?
**R4.** 615 × 0,85 / 3 ≈ 174 To (§46). Ne jamais dépasser 85 % de remplissage.

**Q5.** Un rebuild d'un HDD 20 To à 200 Mo/s prend combien de temps, et quel est
le risque pendant ce temps ?
**R5.** ≈ 28 h. Risque : un 2ᵉ disque qui lâche pendant le rebuild = perte de
données si le placement group n'a plus de réplica sain (§59).

**Q6.** 200 users VDI bureautiques : RAM totale et IOPS au boot storm ?
**R6.** 200 × 8 Go = 1,6 To RAM ; 200 × 100 IOPS = 20 000 IOPS pendant ~10 min (§66).

**Q7.** Un 70B en FP8 tient-il sur 1× RTX PRO 6000 96 Go ? Et sur 2× L40S ?
**R7.** 70B FP8 ≈ 70 Go + 30 % KV cache ≈ 91 Go ≤ 96 Go : oui, sur 1 carte.
Sur 2× L40S (96 Go au total) : oui aussi, en tensor-parallel (§85, §98).

**Q8.** Un rack de 25 kW : air seul ou liquide ? Justifiez avec le calcul d'air.
**R8.** Liquide (RDHx mini). 25 kW × 300 m³/h = 7 500 m³/h d'air — irréaliste en
air seul (§149–150).

**Q9.** Pourquoi dimensionne-t-on un onduleur en kW et pas en kVA ?
**R9.** Les serveurs à PFC actif consomment des watts réels ; un 60 kVA à fp 0,8
ne fournit que 48 kW — sous-dimensionné, il disjoncte (§166).

**Q10.** Citez 3 éléments obligatoires du dossier as-built.
**R10.** Plan de rack + plan de câblage, versions firmware par n° de série,
procédures d'arrêt/démarrage et de failover (§172).

---

## Top 15 des pièges terrain — récapitulatif

