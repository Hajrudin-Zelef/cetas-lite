---
id: collect-261001-rattrapage/rattrapage/datacenter-stockage-guide-1
title: "STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-27"]
keywords: ["datacenter", "compute", "dram", "gpu", "hbm", "nand", "training"]
source: docs/RAG/collect-261001-rattrapage/datacenter_stockage_guide.md
source_anchor: ""
source_lines: [1, 154]
sha256: 0fa73890969a0ba821d1f39830652dc8c364fd6260c04426bf270cc0d9ad5725
---

# STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur

**Guide technique — rédigé en français — vérifications web au 27/09/2026**
**Public :** Zelef, chef de service systèmes & énergies. Ce guide parle chiffres, BOM,
dimensionnement électrique et refroidissement. Les prix sont des ordres de grandeur
issus de sources publiques datées ; ils bougent vite (allocation tendue en sept. 2026).
**Convention de marquage :**
- « vérifié le 27/09/2026 » = source web consultée ce jour (fiche produit, doc constructeur, presse).
- « à vérifier » = ordre de grandeur usuel du métier, à valider sur devis / datasheet exacte.
- « non trouvé au 27/09/2026 » = recherché, introuvable ; à surveiller.

> **Lien avec tes autres guides :** ce document est le socle matériel de `proxmox_guide.md`
> (Ceph, ZFS) et de `onduleurs_ups_guide.md` (dimensionnement électrique, NUT).
> Le chapitre Énergie (sections 183-202) est écrit pour ton métier d'énergéticien.

---

## 1. Objectif du guide

Concevoir, acheter, câbler, exploiter et faire évoluer un stockage datacenter moderne :
serveurs de stockage, cartes HBA, SAS, NVMe (U.2/U.3/EDSFF), JBOD, RAID vs erasure coding,
Ceph full NVMe et hybride, alternatives SDS, fiabilité, énergie, et trajectoire 2026→2030.

## 2. Ce que ce guide n'est pas

Ce n'est pas un catalogue : les références produits changent. C'est une méthode de
dimensionnement + des repères vérifiés au 27/09/2026 + une matrice de décision.
Chaque fois qu'un chiffre peut faire basculer un achat, il est sourcé ou marqué « à vérifier ».

## 3. Architecture moderne du stockage (le paysage en 2026)

```
 ┌─────────────┐     ┌──────────────┐     ┌────────────────┐
 │  DAS        │     │  NAS         │     │  SAN / SDS     │
 │ Direct-     │     │ File         │     │ Block / Object │
 │ Attached    │     │ NFS / SMB    │     │ iSCSI, FC,     │
 │ HBA+JBOD    │     │              │     │ NVMe-oF, Ceph  │
 └─────────────┘     └──────────────┘     └────────────────┘
        \                    |                     /
         \                   |                    /
          └──────────┬───────┴───────┬───────────┘
                     │  NVMe + 100G  │
                     │  EDSFF + CXL  │  ← 2026 : le socle change
                     └───────────────┘
```

Le mouvement de fond : le SAS recule vers le capacitaire (HDD), le NVMe prend le
performant, l'EDSFF remplace l'U.2 dans le dense, et le logiciel (Ceph, SDS) remplace
de plus en plus le RAID matériel.

## 4. Les trois contrats de stockage

| Contrat | Unité | Opération typique | Exemples |
|---|---|---|---|
| Bloc | secteur 512 B / 4 KiB | réécriture in place | RBD, iSCSI, LVM, VM |
| Fichier | plage d'octets dans un chemin | POSIX, verrous, rename atomique | NFS, SMB, CephFS, Lustre |
| Objet | objet entier (immuable) | PUT/GET/DELETE, S3 | RGW, MinIO, Swift |

Règle d'or : **le contrat choisi conditionne tout le reste** (latence, cohérence,
scalabilité). On ne met pas des VM sur de l'objet pur, on ne met pas de l'archivage
froid sur du bloc répliqué 3x.

## 5. Les quatre métriques qui comptent

| Métrique | Unité | Ce qu'elle mesure | Piège classique |
|---|---|---|---|
| IOPS | op/s | transactions | gonflés à QD élevé, irréalistes |
| Débit | Go/s | bande passante | séquentiel ≠ aléatoire |
| Latence | µs / ms | temps de réponse | p99, pas la moyenne |
| Profondeur de file | QD | parallélisme | QD1 (client) vs QD32+ (serveur) |

Pour Ceph et les bases : **latence p99 + IOPS à QD réaliste** > débit séquentiel marketing.
Un SSD NVMe Gen5 fait ~14 500 Mo/s en séquentiel (vérifié : Solidigm D7-PS1010,
KIOXIA CD9P-R, 27/09/2026) mais ce chiffre ne dit rien de ton workload.

## 6. To vs TiB : la confusion qui coûte 9 %

1 To (décimal) = 1 000 Go. 1 Tio (binaire) = 1 099,5 Go. Un « 7,68 To » constructeur =
7,68 × 10¹² octets = **6,99 Tio** utiles bruts avant formatage et overhead.
Sur 1 Po annoncé, l'écart To/TiB mange ~9 %. **Toujours convertir en Tio avant de
dimensionner**, et toujours déduire : over-provisioning (~7-14 %), filesystem (~2-5 %),
métadonnées Ceph/RAID.

## 7. Patterns de charge (workloads)

| Pattern | Lecture/écriture | Taille IO | Sensible à |
|---|---|---|---|
| VM / bases OLTP | 70/30, aléatoire | 4-16 KiB | latence, IOPS |
| Backup / archivage | séquentiel écriture | 128 KiB - 1 MiB | débit, €/To |
| Objet S3 / RGW | gros objets, WORM | 1-64 MiB | débit, capacité |
| IA / training | lecture séquentielle massive | gros blocs, QD très élevé | débit agrégé, GPU-feed |
| VDI | pics matinaux, aléatoire | 4-8 KiB | IOPS burst |

Depuis 2024-2025, les pipelines IA pilotés par GPU poussent des profondeurs de file
de plusieurs milliers (QD4096 en tests — vérifié : méthodologie TweakTown 2026).
Si ton stockage doit nourrir des GPU, dimensionne le **débit agrégé**, pas les IOPS unitaires.

## 8. Échelles : du lab au pétaoctet

| Échelle | Capacité utile | Architecture typique | Budget ordre de grandeur |
|---|---|---|---|
| Lab / PME | 10-50 To | 1-3 nœuds, ZFS ou Ceph HCI | à vérifier (devis) |
| Datacenter moyen | 100 To - 1 Po | 4-12 nœuds Ceph, EC ou réplica 3 | à vérifier |
| Cloud / IA | 1-50 Po | Ceph multi-sites, EDSFF, 400G | à vérifier |

Ce guide détaille le dimensionnement du cas « 100 To utiles » (sections 115-119)
et donne les formules pour extrapoler.

## 9. Méthode de lecture des tableaux « vérifié »

Quand tu lis « vérifié le 27/09/2026 », la valeur vient d'une fiche produit, d'un
communiqué ou d'une doc consultée ce jour-là. Le marché NAND/DRAM est en allocation
tendue en septembre 2026 (source : Fusion Worldwide Greensheet, sept. 2026 — vérifié) :
**les prix bougent, les délais s'allongent** sur SSD entreprise PCIe 5.0, HBM et DDR5
haute densité. Demande toujours un devis daté.

## 10. Plan du guide

Serveurs (11-26) → HBA (27-42) → SAS (43-58) → U.2/U.3/EDSFF (59-72) →
JBOD (73-88) → RAID vs EC vs JBOD (89-104) → Ceph full NVMe (105-132) →
Ceph hybride (133-148) → Alternatives SDS (149-164) → Fiabilité (165-182) →
Énergie (183-202) → À venir (203-216) → Pièges terrain (217-233) →
Glossaire (234) → Quiz (235-236).

---

# PARTIE B — SERVEURS DE STOCKAGE

## 11. Formats de châssis : 1U, 2U, 4U

| Format | Usage stockage typique | Contrainte |
|---|---|---|
| 1U | compute dense, EDSFF E1.S (16 baies) | refroidissement, peu de PCIe |
| 2U | standard : 12 à 32 baies NVMe/SAS | le meilleur compromis |
| 4U | capacitaire : 60 baies 3,5" | poids, profondeur, puissance |

En 2026, le 1U EDSFF monte à **un demi-pétaoctet en 1U16 E3.S** et **1 Po en 2U32 E3.S**
(vérifié : Supermicro, communiqué 2023 toujours d'actualité catalogue 2026).

## 12. Topologie 2U 12 baies

Le classique « 2U 12 baies 3,5" » : 12 disques en façade, carte mère + 2 CPU,
1-2 HBA internes. Idéal pour nœud Ceph HDD, backup, ou tête NVMe U.2 (12 baies).
Bande passante : 1 HBA 12G x8 = ~9,6 Go/s théoriques partagés — largement assez
pour 12 HDD (~2,4 Go/s agrégés).

## 13. Topologie 2U 24 baies

2U 24 baies 2,5" : le format SDS universel (Ceph, vSAN, HCI). 24 SSD SATA/SAS ou
NVMe U.2 selon backplane. Vérifie le **backplane** : simple expander 12G ou
tri-mode NVMe direct — ça change tout le câblage et le débit par baie.

## 14. Topologie 4U 60 baies (top-load)

