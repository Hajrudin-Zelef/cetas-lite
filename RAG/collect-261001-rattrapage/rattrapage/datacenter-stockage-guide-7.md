---
id: collect-261001-rattrapage/rattrapage/datacenter-stockage-guide-7
title: "STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-27"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/datacenter_stockage_guide.md
source_anchor: ""
source_lines: [1028, 1199]
sha256: d290e2ab74ea09c36f8d0203ae629b1ffc7da2bc39f1d452ffef93a398b2cf0f
---

# STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur

```
 ┌─────┐  ┌─────┐  ┌─────┐
 │ MON │  │ MGR │  │ RGW │   ← contrôle (léger)
 └──┬──┘  └──┬──┘  └──┬──┘
    └────────┼────────┘
        ┌────▼────┐
        │  RADOS  │  ← objet distribué, CRUSH
        └────┬────┘
   ┌─────────┼──────────┐
   ▼         ▼          ▼
 ┌────┐   ┌────┐    ┌────┐
 │OSD │   │OSD │    │OSD │  ← 1 daemon par disque
 └────┘   └────┘    └────┘
 RBD (bloc)   CephFS (fichier)   RGW (objet S3)
```

MON : quorum (3 ou 5). MGR : métriques, autoscaler. OSD : le daemon qui
sert un disque. **Règle d'acier : 1 OSD = 1 disque physique** (sauf cas
exotiques). Tout le dimensionnement part de là.

## 106. Pourquoi le NVMe change le design Ceph

Avec des HDD, le disque est le goulot (200 IOPS). Avec des NVMe Gen5
(2 600K IOPS, 14 Go/s), le goulot devient **le CPU, la RAM et le réseau**.
Un cluster full NVMe mal dimensionné en réseau/CPU est un cluster qui
coûte cher pour servir du 25 Gb/s par nœud. Le design full NVMe =
design **réseau + CPU** d'abord, disques ensuite.

## 107. 1 OSD par disque : la règle et ses exceptions

Un OSD par SSD NVMe. Exceptions :
- 2 OSD sur un très gros SSD (30 To+) pour mieux répartir les PG —
  technique avancée, à tester (le WAL/DB partagé peut devenir un goulot).
- Jamais 1 OSD pour 2 SSD (sauf LVM stripé, déconseillé : tu perds la
  granularité de panne).
En full NVMe : **simple, 1:1**.

## 108. PG par OSD : les chiffres officiels (vérifié)

Documentation Ceph / Red Hat (vérifié le 27/09/2026) :
- `mon_target_pg_per_osd` = **100 par défaut**,
- **200 recommandé** pour la plupart des clusters (sauf les plus petits),
- **> 500 = trop** : trafic de peering et RAM excessifs.
Avec le balancer, un cluster démarre vers 50-70 PG/OSD puis l'autoscaler
ajuste. Formule manuelle : `PG ≈ (100 × OSD) / taille_réplica`.

## 109. Calcul PG : exemple chiffré

Cluster : 6 nœuds × 16 NVMe = 96 OSD, réplica 3.
`PG = (100 × 96) / 3 = 3200` → arrondi à la puissance de 2 : **4096**
(ou 2048 si petit). Vérification : 4096 × 3 / 96 = **128 PG/OSD** —
dans la fourchette 100-200. En pratique : laisse **l'autoscaler en mode
`on`** (défaut sain) et marque les gros pools `bulk: true`.

## 110. Autoscaler : on / warn / off

| Mode | Comportement |
|---|---|
| `on` | ajuste pg_num automatiquement (défaut conseillé) |
| `warn` | propose, tu valides (clusters sensibles) |
| `off` | manuel uniquement (experts, calcul fait) |

Deux règles vérifiées (docs + retours terrain 2026) :
- pendant une **mise à jour**, l'autoscaler est en pause (un split de PG
  en pleine upgrade peut ajouter des jours sur un gros cluster),
- `target_size_ratio` > pg_num manuel : donne à l'autoscaler la part
  attendue du pool au lieu de coder des PG en dur.

## 111. RAM par OSD : règle de dimensionnement

Ordre de grandeur usuel (à vérifier selon version Ceph — Reef/Squid) :
- **HDD : ~4-5 Go par OSD** (RocksDB + cache),
- **NVMe : ~2-4 Go par OSD** (moins de latence à masquer, mais plus de PG).
Exemple : 16 NVMe × 3 Go = 48 Go + OS/MON/MGR (~16 Go) + marge → **128 Go
minimum, 256 Go confortable** par nœud 16 baies. Le manque de RAM se voit
en premier sur les pics de recovery (OOM killer sur les OSD).

## 112. CPU : cœurs par OSD

Ordre de grandeur (à vérifier sur ta version) :
- HDD : ~0,5-1 cœur par OSD,
- **NVMe : ~1-2 cœurs par OSD** (chiffrement msgr, BlueStore, recovery).
Exemple : 16 NVMe → 16-32 cœurs. Un double socket 16c/32c (32c/64t) est
le standard raisonnable pour un nœud 16-24 NVMe. En dessous, le recovery
sature le CPU avant le réseau.

## 113. Réseau : débit par OSD (le calcul clé)

En réplica 3, chaque octet écrit par le client génère ~3 octets de trafic
cluster (réplication) + overhead. Formule terrain :
`débit_réseau_nœud ≈ débit_disques_nœud × facteur_réplication × 1,2`.
Exemple : 16 NVMe × 3 Go/s (soutenable, pas le pic) = 48 Go/s de disques ;
× 3 × 1,2 = **172 Go/s = 1 376 Gb/s** théoriques. Personne ne met 14×100G
par nœud : en pratique on dimensionne sur le **débit client cible**
(voir section 118), pas sur le pic disque.

## 114. Réseau : topologie recommandée

| Lien | Débit 2026 | Usage |
|---|---|---|
| Client (public) | 2× 25G ou 2× 100G | accès RBD/RGW/CephFS |
| Cluster (privé) | 2× 100G | réplication, recovery, heartbeat |
| (Option) | 400G | gros nœuds 32 NVMe Gen5 |

Sépare **toujours** le réseau cluster du réseau client (VLAN ou physique).
Le recovery d'un nœud 32×7,68 To peut saturer 100G pendant des heures :
si c'est le même lien que les clients, tout le monde le sent.

## 115. Dimensionnement « 100 To utiles » : hypothèses

- Besoin : **100 To utiles** (Tio, après conversion !), workload mixte VM.
- Protection : réplica 3 (latence faible).
- Disques : NVMe E3.S 7,68 To Gen5 (6,99 Tio bruts).
- Overhead : 10 % (métadonnées, marge de remplissage — on ne remplit
  jamais un cluster Ceph au-delà de ~70-80 %).

## 116. Dimensionnement : nombre de disques

Brut nécessaire = 100 × 3 × 1,1 ≈ **330 To** → / 6,99 Tio ≈ **48 disques**.
Arrondi pratique : **48 NVMe de 7,68 To**. Avec des 15,36 To : 24 disques
(moins de granularité PG — voir section 108 : 24 OSD, c'est peu pour
4096 PG ; préfère plus de petits disques que peu de gros).

## 117. Dimensionnement : nombre de nœuds

Contraintes : failure domain = le nœud ; réplica 3 → **minimum 3 nœuds**,
recommandé 4-5 pour survivre à 1 panne + 1 maintenance.
Choix : **4 nœuds × 12 NVMe** = 48 OSD.
Par nœud : 12 × 6,99 = 83,9 Tio bruts ; cluster brut = 335,5 Tio ;
utiles ≈ 335,5 / 3 × 0,9 ≈ **100 Tio**. Le calcul tombe juste.
Alternative 5 nœuds × 10 : plus résilient, un peu plus cher.

## 118. Dimensionnement : réseau chiffré

Débit client cible : 4 nœuds × 25 Go/s (2×100G LACP par nœud, réaliste
soutenable) = **100 Go/s agrégés** côté client. Réseau cluster : 2×100G
par nœud également (le recovery doit suivre). Total par nœud : **4×100G**
(2 client + 2 cluster) ou 2×100G partagés avec QoS stricte (déconseillé
en production). Budget switch : 2× 32 ports 100G (MLAG) — **à vérifier**
sur devis (~15-30 k€ pièce en ordre de grandeur public).

## 119. Dimensionnement : RAM et CPU chiffrés

Par nœud (12 NVMe) : RAM = 12 × 3 Go + 16 Go système ≈ **64 Go → 128 Go
installés** (marge recovery). CPU : 12 × 1,5 cœur ≈ 18 cœurs → **2× 16c**
ou 1× 32c moderne. Conso nœud estimée : voir section 187.

## 120. Exemple BOM : nœud Ceph NVMe (×4)

| Ligne | Qté/nœud | Ordre de grandeur |
|---|---|---|
| Châssis 2U 12-24 baies E3.S/U.2 Gen5 | 1 | à vérifier |
| 2× CPU 16c / 1× 32c | 1 lot | à vérifier |
| 128 Go DDR5 ECC | 1 lot | à vérifier |
| 12× NVMe 7,68 To Gen5 E3.S | 12 | ~350-600 €/To → 32 000-55 000 € |
| 2× NIC 100G (client) + 2× NIC 100G (cluster) | 4 | à vérifier |
| 2× alimentations Titanium | inclus | — |

Ordre de grandeur par nœud (sans les NIC/switch) : **40-70 k€** ;
cluster 4 nœuds : **160-280 k€** hors réseau. **À vérifier** sur devis
(allocation tendue sept. 2026).

## 121. Coût au To utile : l'indicateur à suivre

`€/To utile = coût total cluster / To utiles`. Exemple ci-dessus :
~200 k€ / 100 To ≈ **2 000 €/To utile** (ordre de grandeur, à vérifier).
Compare toujours en €/To **utile** (pas brut) et en €/To/an sur 5 ans
en ajoutant énergie + maintenance. Le NVMe full perf coûte ~5-10× le
HDD capacitaire au To : c'est normal, ce n'est pas le même service.

## 122. Recovery et backfill : l'épreuve du feu

