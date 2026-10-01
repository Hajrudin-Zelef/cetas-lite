---
id: collect-261001-rattrapage/rattrapage/datacenter-stockage-guide-9
title: "STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-27"]
keywords: ["apache", "arr", "gpu", "training"]
source: docs/RAG/collect-261001-rattrapage/datacenter_stockage_guide.md
source_anchor: ""
source_lines: [1387, 1559]
sha256: 88c4774f508fe14098be7317478d2b8f72acccdcb0f4748a25c1dc1b6e36735b
---

# STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur

Ne mets **jamais** le block.wal sur le HDD « pour économiser un SSD » :
le WAL est sur le chemin critique de chaque écriture. Le couple
wal+db va **toujours** sur le rapide (d'ailleurs le wal est absorbé
dans le db par défaut en BlueStore moderne).

## 143. Piège : mélanger les tailles de HDD

Des OSD de 12 To et de 20 To dans le même pool : Ceph remplit au
pourcentage, les 12 To se remplissent (en %) aussi vite → le cluster
s'arrête sur les petits disques. Règle : **une taille de disque par
pool**, ou des pools séparés par taille (device class custom ou
pools dédiés).

## 144. Migration hybride → full NVMe

Stratégie : ajoute des nœuds NVMe (device class `nvme`), crée les
nouveaux pools dessus, migre les données (rbd migration, RGW
resharding, rsync pour CephFS), puis retire les nœuds HDD avec
`ceph osd out` + rebalance. **Jamais de bascule « big bang »** :
la migration prend des jours sur des centaines de To — c'est normal,
planifie-la.

## 145. Cas d'usage : backup sur hybride

Backup (Veeam, Proxmox Backup Server, Borg) : écritures séquentielles,
lectures rares → **hybride HDD + db SSD**, EC 8+3 ou réplica 2 selon
criticité. Le db SSD absorbe les pics d'écriture des fenêtres de backup.
Dimensionne le réseau pour la **fenêtre de backup** (ex. 8 h pour
50 To = ~1,8 Go/s soutenus).

## 146. Cas d'usage : CephFS froid

Données utilisateurs, archives : pool data en EC sur HDD, pool metadata
en réplica 3 sur SSD (les métadonnées sont petites et chaudes). Le MDS
apprécie énormément un pool metadata sur SSD : le `ls -R` d'un gros
volume passe de minutes à secondes.

## 147. Ordre de grandeur prix hybride

Exemple section 140 : 40× HDD 20 To (~18 €/To ≈ 14 400 €) + 3× NVMe
7,68 To (~10 000 €) + 5 nœuds modestes + réseau 25G. Total ordre de
grandeur : **60-100 k€** pour 500 To utiles → **~120-200 €/To utile**,
soit ~10× moins cher que le full NVMe. **À vérifier** sur devis.

## 148. Checklist hybride

- [ ] block.db ≥ 2,5 % (4 % RGW), jamais sur HDD
- [ ] ≤ 15 OSD HDD par NVMe de db
- [ ] Device classes + règles CRUSH par tier
- [ ] Pas de cache tiering
- [ ] Une taille de HDD par pool
- [ ] Monitoring espace db

---

# PARTIE J — ALTERNATIVES À CEPH (SDS DU MARCHÉ)

## 149. Panorama : tableau comparatif (existence vérifiée le 27/09/2026)

| Solution | Type | Licence | Contrats | Positionnement |
|---|---|---|---|---|
| Ceph | SDS unifié | LGPLv2.1 | bloc/fichier/objet | généraliste, expert |
| MinIO | objet S3 | AGPLv3 | objet | S3 simple et rapide |
| Garage | objet S3 | AGPLv3 | objet | léger, géo-distribué |
| SeaweedFS | objet/fichier | Apache 2.0 | objet/fichier | petits fichiers, simple |
| RustFS | objet S3 | Apache 2.0 | objet | successeur potentiel MinIO, jeune |
| Apache Ozone | objet/fichier | Apache 2.0 | objet/fichier | Hadoop, exa-échelle |
| OpenStack Swift | objet | Apache 2.0 | objet | clouds privés historiques |
| JuiceFS | fichier sur objet | Apache 2.0 | fichier | POSIX au-dessus de S3 |
| Lustre | fichier parallèle | GPLv2 | fichier | HPC |
| WEKA | fichier parallèle | propriétaire | fichier/objet | IA/GPU, ultra-faible latence |
| VAST Data | fichier/objet | propriétaire | fichier/objet | IA, all-flash |
| Scality | objet | propriétaire | objet | S3 entreprise, multi-région |
| Cloudian | objet | propriétaire | objet | S3, backup, anti-ransomware |
| IBM Spectrum Scale | fichier | propriétaire | fichier | HPC/IA, POSIX |

Toutes ces solutions **existent** (vérifié via comparatifs et docs, 27/09/2026).
GlusterFS existe aussi mais est en déclin (Red Hat l'a déprécié au profit de Ceph/ODF).

## 150. MinIO : l'état en 2026

MinIO = le S3 self-hosted le plus simple : erasure coding inline, pas
d'index séparé, déploiement en minutes. Points de vigilance 2026
(presse/communauté, vérifié) : licence **AGPLv3** (contamination IP à
évaluer avec ton juridique), et un essoufflement perçu de la version
communautaire (des équipes migrent vers Garage/SeaweedFS/RustFS ou Ceph
RGW). Pour un S3 simple, mono-site, équipe réduite : reste un choix
valable. Pour du multi-site critique : évalue Ceph RGW ou Scality.

## 151. Garage : le léger géo-distribué

Écrit en Rust, **très faible empreinte** (tourne sur ARM, petits nœuds),
réplication multi-sites pensée dès le départ. Limites : pas pour le
multi-pétaoctet, S3 partiel. Idéal : edge, sites distants, S3 léger
d'entreprise. Licence AGPLv3.

## 152. SeaweedFS : les petits fichiers

Modèle inspiré de Haystack (Facebook) : excellent sur les **milliards
de petits fichiers** où Ceph RGW peine (overhead par objet). Licence
Apache 2.0 (business-friendly), S3 + FUSE + HDFS. À évaluer quand ton
workload = beaucoup de petits objets.

## 153. RustFS : le jeune challenger

Rust, Apache 2.0, annoncé ~2,3× plus rapide que MinIO sur petits payloads
(selon le projet — **à vérifier** en bench propre), zéro télémétrie.
Statut : jeune (alpha 2026), une faille de sécurité a refroidi des
adoptants. À surveiller, pas à mettre en production critique en 2026.

## 154. Apache Ozone : l'exa-échelle Hadoop

Ozone = l'objet « Hadoop moderne » : S3 + HDFS API, **forte cohérence**,
échelle exaoctet / dizaines de milliards de clés (vérifié : comparatif
Apache Ozone). Si ton écosystème est Hadoop/Spark : Ozone plutôt que
Ceph RGW. Sinon : complexité injustifiée.

## 155. OpenStack Swift : l'historique

L'objet d'OpenStack : anneaux (rings), cohérence à terme, multi-tenant.
Toujours maintenu, pertinent si tu es déjà sur OpenStack. Sinon, les
alternatives modernes (RGW, SeaweedFS) sont plus simples à opérer.

## 156. JuiceFS : du POSIX sur de l'objet

JuiceFS met un filesystem POSIX **au-dessus** d'un stockage objet (S3,
etc.) avec Redis/métadonnées séparées. Cas d'usage : partager un volume
POSIX entre clouds, ou donner du filesystem à des applis sur un S3
existant. Ce n'est pas un SDS complet : il faut l'objet en dessous.

## 157. Lustre : le HPC

Le filesystem parallèle des supercalculateurs : débit agrégé monstrueux,
POSIX, mais administration experte et pas d'objet. Pour un cluster IA/GPU
où le checkpoint/restore et le training dictent tout : Lustre ou WEKA,
pas Ceph.

## 158. WEKA : la performance propriétaire (vérifié)

Filesystem parallèle propriétaire, positionné **IA/HPC** : latence
ultra-faible, IOPS extrêmes, protocoles POSIX/NFS/SMB/S3. Modèle :
logiciel sur ton hardware (ou appliance). Coût : licence au To —
**à vérifier** sur devis, c'est le haut du panier tarifaire. À évaluer
quand le stockage est le goulot du training IA.

## 159. VAST Data : l'all-flash IA (vérifié)

Plateforme propriétaire **all-flash NVMe** (NFS/S3), pensée pour l'IA et
l'analytique : faible latence, réduction de données agressive. Positionné
contre les baies traditionnelles. Comme WEKA : excellent, cher, à
chiffrer contre un Ceph full NVMe bien conçu (l'écart se réduit si tu
as l'expertise Ceph en interne).

## 160. Scality : l'objet entreprise (vérifié)

S3 entreprise : multi-région, durabilité, cas backup/archive et cloud
hybride. Le choix « S3 sans surprise » quand on veut du support
contractuel et pas d'équipe SDS interne. Concurrent direct de Ceph RGW
+ support Red Hat/IBM.

## 161. Cloudian : l'objet anti-ransomware (vérifié)

S3 compatible avec fonctions **protection ransomware** (verrouillage
d'objets, immuabilité), backup et hybride. Pertinent quand la menace
n°1 est le chiffrement malveillant des backups : l'immutabilité S3
(Object Lock) est une exigence, quel que soit le produit.

## 162. IBM Spectrum Scale (ex-GPFS) : le POSIX scale-out (vérifié)

Filesystem parallèle IBM : POSIX strict, tiering par politiques, HPC/IA.
L'option « grande entreprise IBM » quand le support et la conformité
priment sur le coût. Toujours vivant en 2026 dans le HPC et la finance.

