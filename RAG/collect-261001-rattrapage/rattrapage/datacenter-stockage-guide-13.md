---
id: collect-261001-rattrapage/rattrapage/datacenter-stockage-guide-13
title: "STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "arr"]
source: docs/RAG/collect-261001-rattrapage/datacenter_stockage_guide.md
source_anchor: ""
source_lines: [2096, 2221]
sha256: bd6f25c45f726edc73bbee95ffdd15e1f42abeb2d02a46789993dd708b60be78
---

# STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur

Symptôme : Ceph/ZFS ne voit qu'un « volume » ou rien. Cause : firmware
MegaRAID au lieu de HBA. Solution : reflasher en IT **avant** toute mise
en production, et vérifier `storcli /c0 show` (personality). Coût de
l'oubli : reconstruire le nœud.

## 218. Piège n°2 : le câble SlimSAS au mauvais brochage

Symptôme : lien qui ne monte pas, ou SSD NVMe vu en SAS. Cause : câble
SFF-8654 câblé PCIe branché sur backplane SAS (ou l'inverse). Solution :
commander les câbles **par référence constructeur du châssis**, pas « un
SlimSAS générique ». Garde la doc de brochage dans le dossier du rack.

## 219. Piège n°3 : le slot PCIe bridé

Symptôme : HBA 16 voies qui ne débite que la moitié. Cause : slot x16
mécanique mais **x8 ou x4 électrique** (fréquent sur cartes mères
denses). Solution : vérifier `lspci -vv | grep LnkSta` (vitesse **et**
largeur négociées) à l'installation. Un `x8` au lieu de `x16` sur une
9670 = 50 % du débit en moins.

## 220. Piège n°4 : le RAID 5 sur disques de 20 To

Symptôme : rebuild qui échoue à 80 %. Cause : URE pendant le rebuild
(section 91). Solution : RAID 6 minimum sur gros HDD, et au-delà d'une
baie : erasure coding logiciel. Ce piège a coûté des volumes entiers :
**ne transige pas**.

## 221. Piège n°5 : le block.db sous-dimensionné

Symptôme : OSD en erreur avec des HDD à 50 % (section 141). Cause : db
à 2,5 % sur un workload RGW qui demandait 4 %+. Solution : dimensionne
large dès le départ (4 % pour RGW), monitor l'espace db. Il n'y a pas
de « petit rajout » : c'est une recréation d'OSD.

## 222. Piège n°6 : le SSD QLC comme block.db

Symptôme : latences d'écriture qui explosent par vagues. Cause : QLC à
0,5 DWPD dont le cache SLC sature sous le journal Ceph. Solution :
**TLC 3 DWPD minimum** pour les db/wal (section 170). Le QLC reste sur
le froid.

## 223. Piège n°7 : le spin-down sur Ceph actif

Symptôme : latences aléatoires de plusieurs secondes. Cause : les
disques s'endorment et Ceph les réveille en permanence (scrub,
heartbeat). Solution : **désactive le spin-down** sur tout cluster
Ceph/ZFS actif ; réserve-le à l'archivage WORM pur (section 186).

## 224. Piège n°8 : le JBOD sans multipath

Symptôme : corruption après un flap de câble. Cause : dual-path SAS
sans `multipathd` → le même disque vu deux fois → écritures concurrentes.
Solution : multipath configuré **et testé** (débranche un câble en
burn-in) avant production (section 36).

## 225. Piège n°9 : l'expander en cascade triple

Symptôme : rebuild 3× plus lent que prévu. Cause : 3 JBOD en cascade
sur 8 voies = bande passante divisée + latence cumulée. Solution :
**2 niveaux de cascade max**, et de préférence une HBA par JBOD en
production (section 76).

## 226. Piège n°10 : le firmware flashé un vendredi

Symptôme : 60 disques qui disparaissent après un flash d'expander.
Cause : flash en production sans fenêtre, sans rollback. Solution :
firmware par vagues (1 nœud pilote → 48 h → reste), en fenêtre de
maintenance, avec l'ancien firmware sous la main (sections 37, 85, 173).

## 227. Piège n°11 : le PG count oublié

Symptôme : cluster à 30 % avec des OSD à 95 % (hot spots). Cause :
pool créé sans `bulk: true`, autoscaler jamais convergé, 32 PG pour
96 OSD. Solution : `bulk: true` sur les gros pools, vérifie
`ceph osd pool autoscale-status`, vise 100-200 PG/OSD (sections 108-110).

## 228. Piège n°12 : le réseau client/cluster mélangé

Symptôme : les VM rament pendant chaque recovery. Cause : un seul
réseau 25G pour clients + réplication. Solution : **sépare** (VLAN
mini, physique idéalement), 100G sur le cluster en full NVMe
(sections 114, 118).

## 229. Piège n°13 : le remplissage à 95 %

Symptôme : Ceph passe en `nearfull` puis `full`, écritures bloquées.
Cause : pas de marge (on a « optimisé » le dimensionnement). Solution :
**alerte à 70 %, critique à 85 %**, et achète l'extension à 70 %
(délais d'allocation 2026 : plusieurs semaines — section 213).

## 230. Piège n°14 : le disque SMR dans le RAID

Symptôme : rebuild à 5 Mo/s pendant des jours. Cause : SMR shingled
(section 103). Solution : **CMR uniquement** en datacenter ; vérifie
la fiche disque (les gammes grand public mélangent CMR/SMR sans le
crier).

## 231. Piège n°15 : l'arrêt électrique non séquencé

Symptôme : après une coupure, des OSD à reconstruire et des SSD
corrompus. Cause : extinction brutale, pas de séquence NUT (section
190). Solution : séquence écrite, **testée 2×/an**, onduleur
dimensionné pour le pic de spin-up (section 200).

## 232. Piège n°16 : le SSD grand public sans power-loss protection

Symptôme : filesystem corrompu après micro-coupure. Cause : SSD client
sans condensateurs de protection (PLP). Solution : en datacenter,
**SSD entreprise avec PLP** (tous les modèles cités en section 70
l'ont — vérifié). Le « bon plan » SSD gaming pour un serveur, c'est
une corruption en attente.

## 233. Piège n°17 : le prix au To brut

Symptôme : le « moins cher » coûte 2× plus cher à 5 ans. Cause :
comparaison en €/To brut sans énergie, sans surcoût de protection,
sans maintenance. Solution : **TCO 5 ans en €/To utile** incluant
kWh (PUE), RMA, et extensibilité (sections 121, 197). Fais signer ce
tableau avant l'achat : il évite 90 % des regrets.

---

# PARTIE O — GLOSSAIRE (40 TERMES)

## 234. Glossaire

