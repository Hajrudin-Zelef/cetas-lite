---
id: collect-261001-rattrapage/rattrapage/datacenter-stockage-guide-6
title: "STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter"]
source: docs/RAG/collect-261001-rattrapage/datacenter_stockage_guide.md
source_anchor: ""
source_lines: [844, 1027]
sha256: 1340731a34ba5c4299b3ff377885c17df85f6ffc75c8f9d6e1809efa70f19885
---

# STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur

L'expander a son firmware, avec ses bugs (perte de disques fantôme,
négociation bloquée). Symptôme typique : des disques qui « disparaissent
et réapparaissent ». Avant de suspecter les disques, **vérifie la version
firmware de l'expander** et les errata du constructeur. Flash en fenêtre
de maintenance : un flash d'expander coupe tous ses disques.

## 86. Câblage : bonnes pratiques

- Étiquette **chaque** câble (A1, A2, B1, B2 : domaine + numéro).
- Ne mélange pas les générations sur un même domaine sans test.
- Serre les vis des SFF-8644/8674 (ils se débranchent avec les vibrations).
- Photo du câblage dans la doc d'exploitation, mise à jour à chaque
  intervention.
- Un câble de rechange de chaque type **dans la baie**, pas au dépôt.

## 87. Exemple BOM : extension JBOD 60 baies

| Ligne | Qté |
|---|---|
| JBOD 60 baies 12G double expander, 2×1600 W | 1 |
| Câbles SFF-8644 2 m | 4 (2 par domaine) |
| HDD 20 To CMR | 60 |
| Rails + kit | 1 |

Coût disques seul : 60 × 20 To × ~18 €/To ≈ 21 600 € (ordre de grandeur,
**à vérifier**). Le JBOD nu (sans disques) : **à vérifier** sur devis.

## 88. Tableau : JBOD vs tout-en-un

| Besoin | JBOD + tête | Serveur 60 baies |
|---|---|---|
| €/To capacitaire | souvent meilleur | plus simple |
| Domaine de panne | tête unique | nœud unique (pareil) |
| Évolutivité | ajoute des JBOD | ajoute des nœuds |
| Ops | 2 équipements à gérer | 1 seul |

---

# PARTIE G — RAID vs ERASURE CODING vs JBOD

## 89. RAID : le principe en une phrase

Le RAID répartit données + parité sur N disques pour survivre à 1-2 pannes,
au prix d'un surcoût et d'un rebuild. Il ne protège **ni** de la corruption
silencieuse (sans checksums), **ni** de la perte du site, **ni** de l'erreur
humaine (`rm -rf`).

## 90. Tableau RAID 0 / 1 / 10

| Niveau | Surcoût | Pannes tolérées | Usage |
|---|---|---|---|
| RAID 0 | 0 % | 0 | scratch, cache (données reproductibles) |
| RAID 1 | 50 % | 1 par miroir | boot, petits volumes critiques |
| RAID 10 | 50 % | 1+ (selon placement) | bases OLTP, perf + résilience |

## 91. RAID 5 : le write hole et les URE

Le RAID 5 (1 parité) a deux défauts mortels sur gros disques :
1. **Write hole** : coupure pendant l'écriture → parité incohérente
   (atténué par journal/cache protégé, jamais totalement supprimé).
2. **URE pendant le rebuild** : avec des disques 20 To (BER ~10⁻¹⁴),
   la probabilité de rencontrer une erreur de lecture irrécupérable
   pendant un rebuild de 20 To est **significative** → rebuild qui échoue
   → volume perdu. **Règle : plus de RAID 5 au-delà de ~4-6 To par disque.**

## 92. RAID 6 : la double parité

Deux parités → survit à 2 pannes. Le standard capacitaire HDD en RAID
matériel. Surcoût : 2 disques par groupe. Reste exposé au temps de rebuild
(section 93) et ne scale pas au-delà d'une baie. Pour du multi-nœuds :
l'erasure coding logiciel (section 94) le remplace.

## 93. Temps de rebuild : le calcul qui fait peur

`temps = capacité_disque / débit_rebuild_effectif`.
Exemple : 20 To à 200 Mo/s effectifs (rebuild partagé) = 100 000 s ≈
**28 heures**. Pendant ces 28 h, le groupe est dégradé ; une 2e panne
(sur RAID 5) ou 3e (sur RAID 6) = perte. Avec 60 disques de 20 To, la
probabilité d'une 2e panne pendant la fenêtre n'est pas négligeable :
c'est l'argument mathématique pour l'erasure coding distribué.

## 94. Erasure coding : le principe

On découpe les données en **k** fragments + **m** fragments de parité
(Reed-Solomon). N'importe quels k fragments parmi k+m reconstruisent les
données. Surcoût = m/(k+m). Tolère **m** pannes **distribuées sur le cluster**,
pas sur un groupe de disques.

## 95. Profils EC courants

| Profil | Surcoût | Tolère | Usage typique |
|---|---|---|---|
| 4+2 | 33 % | 2 pannes | petit cluster, objet |
| 8+3 | 27 % | 3 pannes | standard S3/RGW |
| 8+2 | 20 % | 2 pannes | froid, gros cluster |

Comparé à réplica 3 (surcoût 200 %), l'EC divise le coût capacitaire par ~2.
Prix à payer : **CPU** (encodage), **latence d'écriture**, et rebuilds
gourmands en réseau.

## 96. Surcoût comparé (tableau)

| Méthode | Surcoût | 100 To utiles → brut |
|---|---|---|
| Réplica 3 | 200 % | 300 To |
| EC 4+2 | 50 % | 150 To |
| EC 8+3 | 37,5 % | 137,5 To |
| RAID 6 (8+2 local) | 25 % | 125 To |

Le RAID 6 local reste le moins cher en brut, mais il ne protège que dans
la baie. L'EC protège dans tout le cluster : ce n'est pas le même service.

## 97. JBOD + réplication logicielle

Le « JBOD » comme stratégie = pas de RAID du tout, la redondance est
gérée par le logiciel (Ceph réplica, ZFS copies=2, MinIO erasure sets).
Avantage : le logiciel voit les disques et leurs erreurs. Inconvénient :
il faut un logiciel qui sait le faire (pas un simple montage ext4).

## 98. Comparatif honnête

| Critère | RAID HW | EC logiciel | JBOD + réplica SW |
|---|---|---|---|
| Surcoût typique | 25-50 % | 27-50 % | 100-200 % |
| Échelle | 1 serveur | cluster | cluster |
| Rebuild | lent, local | rapide, distribué | rapide, distribué |
| Corruption silencieuse | non détectée | détectée (checksums) | détectée |
| Complexité ops | faible | élevée | moyenne |
| Coût logiciel | carte RAID | expertise | expertise |

## 99. Quand choisir le RAID hardware

- Serveur unique, applicatif classique, pas de SDS (PME, boot, petits volumes).
- RAID 1 (boot/OS) et RAID 10 (bases sur serveur unique).
- Quand l'équipe n'a pas l'expertise SDS et que le besoin tient dans une baie.
**Jamais** sous Ceph/ZFS/vSAN sur les disques de données.

## 100. Quand choisir l'erasure coding

- Objet/S3, backup, archivage : données froides, gros volumes, écriture
  peu fréquente → EC 8+3.
- Cluster ≥ 8-10 nœuds (l'EC a besoin de largeur).
- Quand le €/To prime sur la latence d'écriture.

## 101. Quand choisir JBOD + réplication

- Bloc (RBD, iSCSI) et fichiers : la réplication 3x reste le standard
  (latence d'écriture faible).
- Petits clusters (< 8 nœuds) où l'EC serait trop étroit.
- Workloads mixtes sans expertise EC.

## 102. RAID et SSD : particularités

Le rebuild SSD est rapide (pas de mécanique), mais :
- l'usure s'accélère pendant le rebuild (écritures massives),
- le RAID 5/6 sur SSD QLC à faible endurance = usure prématurée,
- le TRIM/discard ne traverse pas toujours la couche RAID → perfs qui
  se dégradent avec le temps. Vérifie le support du TRIM chez le
  constructeur de la carte.

## 103. RAID et SMR : ne pas

Les disques SMR (enregistrement magnétique shingled, ex. certains
Barracuda grand public) sont **incompatibles** avec le RAID et ZFS/Ceph
en écriture aléatoire : le réécriture de bandes fait chuter le débit à
quelques Mo/s. Pour le datacenter : **CMR uniquement** (Exos, Ultrastar
DC, IronWolf Pro). Le SMR n'a sa place qu'en écriture séquentielle pure
(archivage WORM maîtrisé).

## 104. Matrice de décision finale

```
Une seule baie, équipe réduite ──▶ RAID 1/10/6 matériel
Cluster, bloc/fichier, latence ──▶ Ceph réplica 3 / ZFS miroir
Cluster, objet/froid, €/To ──▶ Ceph EC 8+3 / MinIO erasure sets
Données reproductibles ──▶ RAID 0 / JBOD pur (avec réplication ailleurs)
```

---

# PARTIE H — CEPH FULL NVMe : DESIGN COMPLET

## 105. Architecture Ceph en 30 secondes

