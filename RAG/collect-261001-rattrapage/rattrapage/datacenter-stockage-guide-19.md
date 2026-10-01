---
id: collect-261001-rattrapage/rattrapage/datacenter-stockage-guide-19
title: "STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "arr", "memory", "revenue"]
source: docs/RAG/collect-261001-rattrapage/datacenter_stockage_guide.md
source_anchor: ""
source_lines: [3084, 3264]
sha256: 3182d93644ce1d3709690c58b1223ed040f6b28d2549e3637c41a893fc3348ab
---

# STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur

- `nr_requests` : 256-1024 sur SSD/NVMe (plus = plus de parallélisme).
- `read_ahead_kb` : 128 par défaut ; 1024-4096 en séquentiel pur
  (backup), 0-128 en aléatoire (bases).
- `max_sectors_kb` : 1024+ sur HBA modernes.
Ces réglages se mettent en **udev** (persistants), pas à la main.

## 316. Tuning Linux : l'affinité IRQ

`irqbalance` suffit souvent, mais sur nœud NVMe chargé : affinité
manuelle des queues NVMe (`/proc/irq/*/smp_affinity`) sur des cœurs
dédiés, séparés des OSD. Mesure avec `mpstat -P ALL` : un cœur à
100 % softirq pendant que 31 autres dorment = affinité à revoir.

## 317. Huge pages et Ceph

Les OSD allouent beaucoup de petits buffers : les huge pages (2 Mo)
réduisent la pression TLB sur gros nœuds. À tester par workload —
le gain est réel sur 32+ OSD par nœud, marginal en dessous. Comme
tout tuning : **bench avant/après**, pas de cargo cult.

## 318. La déduplication : le calcul honnête

La dédup (ZFS, VAST) économise 2-5× sur VDI/sauvegardes, ~1,2× sur
données généralistes. Coût : **RAM** (table DDT : ~5 Go par To
dédupliqué sous ZFS, ordre de grandeur) + CPU. Règle : ne déduplique
que si le ratio mesuré en lab dépasse 2× **et** que tu as la RAM.
Sinon : compression seule.

## 319. La compression : lz4/zstd

- **lz4** : quasi gratuit en CPU, ratio ~2× sur texte/logs/VM.
- **zstd** : meilleur ratio, coût CPU modéré (niveaux 1-3 en temps réel).
Sur Ceph BlueStore : compression par pool (`compression_algorithm zstd`).
Sur ZFS : `compression=zstd-3`. **Toujours activer** sauf données déjà
compressées/chiffrées (médias, backups chiffrés : la compression ne
sert à rien et coûte du CPU — mesure).

## 320. Le chiffrement et la compression : l'ordre compte

Chiffré puis compressé = incompressible. Donc : compresse **avant**
de chiffrer (ZFS le fait dans le bon ordre en natif). Si tes backups
sont chiffrés côté client avant envoi : désactive la compression côté
stockage, c'est du CPU perdu.

## 321. La bande magnétique : LTO en 2026

LTO-9 : 18 To natifs (45 To compressés), ~400 Mo/s. Usage : archivage
froid ultime, air gap anti-ransomware. Coût : ~10 €/To le média
(ordre de grandeur, à vérifier) + librairie. Le **vrai** 3-2-1
d'un gros site inclut encore de la bande en 2026. Inconvénient :
temps d'accès (minutes), gestion de librairie = métier.

## 322. Archivage : la pyramide des tiers

```
 Chaud : NVMe réplica 3      < 1 ms    ~2000 €/To
 Tiède : TLC EC 4+2          ~2 ms     ~500 €/To
 Froid : QLC EC 8+3          ~5 ms     ~250 €/To
 Glacial : HDD EC 8+3        ~10 ms    ~150 €/To
 Air gap : LTO               minutes   ~30 €/To (média)
```
(Ordres de grandeur, à vérifier.) Le tiering = placer chaque donnée
au bon étage. 80 % des données ont plus d'un an : elles n'ont rien
à faire sur du NVMe.

## 323. RGW avancé : multisite

RGW multisite : réplication asynchrone entre zones (active-active ou
active-passive). Usage : PRA inter-sites S3. Pièges : le lag de
réplication (surveille `radosgw-admin sync status`), les conflits
de versions (politique last-write-wins par défaut), et le **coût
réseau** inter-sites (facturé au Go chez beaucoup d'opérateurs).

## 324. RGW avancé : le sharding des buckets

Un bucket avec des millions d'objets sans sharding = index qui rame.
`rgw_override_bucket_index_max_shards` et resharding en ligne :
à planifier **avant** que le bucket n'explose (le resharding d'un
bucket de 100 M d'objets prend des heures). Règle : shard dès la
création pour les buckets applicatifs à forte croissance.

## 325. CephFS avancé : les MDS en pratique

1 MDS actif + 2 standby par filesystem en production. Le MDS tient
tout son cache en RAM : **64-128 Go** par MDS actif sur gros volumes.
Symptôme d'un MDS sous-dimensionné : `mds_cache_memory_limit`
dépassé, latences `ls` catastrophiques. Sépare les filesystems par
usage (homes vs scratch) : un `find /` d'un utilisateur ne doit pas
évincer le cache d'un autre.

## 326. CephFS : les snapshots

Snapshots au niveau répertoire (`mkdir .snap/x`) : instantanés,
cohérents, pas de surcoût initial. Idéal pour les homes et les
dépôts. Limite : le MDS doit suivre les snaps (ne pas en garder
des centaines). Ce n'est toujours pas un backup (section 289).

## 327. RBD avancé : mirroring inter-sites

`rbd-mirror` : réplication asynchrone d'images entre clusters
(journal-based). RPO : secondes à minutes. Usage : PRA VM
inter-sites. À tester : le failover (promote) et le failback —
un mirroring jamais testé en bascule est un espoir, pas un PRA.

## 328. RBD : les clones et le flatten

Clone = image COW sur snapshot parent. Pratique pour déployer 100 VM
depuis un template (espace partagé). Piège : le parent ne peut pas
être supprimé tant que des clones vivent, et un parent sur un pool
lent ralentit tous les clones → **flatten** les clones devenus
« adultes » (`rbd flatten`).

## 329. Les pools : un par usage

Un pool = une politique (device class, réplica/EC, PG). Crée :
`rbd-vm-ssd` (réplica 3), `rbd-backup-hdd` (EC), `rgw-data` (EC 8+3),
`cephfs-data-froid` (EC), `cephfs-metadata` (réplica 3 SSD).
**Jamais** de pool fourre-tout : on ne change pas la politique d'un
pool plein sans migration.

## 330. Le crush map : les buckets

```
host → chassis → rack → row → room → datacenter → root
```
CRUSH place les répliques en remontant les buckets : `chooseleaf`
sur `host` = 1 réplique par host. Pour du stretch : des buckets
`datacenter`. **Dessine ton crush map avant** de créer les règles :
un bucket oublié (rack non déclaré) = placement qui ignore la
topologie réelle.

## 331. Failure domains : la hiérarchie

| Domaine | Protège contre | Coût |
|---|---|---|
| osd | panne disque | 0 (défaut logique) |
| host | panne nœud | nœuds ≥ réplica |
| rack | panne rack/switch | câblage inter-racks |
| datacenter | perte de site | 2e site |

Monte d'un cran quand le budget le permet. La plupart des PME
s'arrêtent à `host` — c'est déjà très bien avec 4+ nœuds.

## 332. Étude de cas : la panne silencieuse

Scénario vécu (type) : un SSD NVMe « lent » (pas mort) fait grimper
la latence p99 de tout le cluster pendant 3 semaines avant d'être
identifié. Le SMART était OK, mais `bluestore_slow_ops` crachait.
Leçon : surveille les **slow ops**, pas seulement les pannes franches
(section 130). Le disque a été remplacé, p99 revenue à la normale
en 1 h. Coût de la non-détection : 3 semaines de perfs dégradées.

## 333. Étude de cas : le câble

Scénario type : après un déménagement de rack, des erreurs CRC
intermittentes sur un domaine SAS. 2 semaines de diagnostic (disques
suspectés, firmware reflashé) avant de trouver : un SFF-8644 **mal
enclenché** (vis non serrées, vibrations). Leçon : après chaque
intervention physique, `storcli` + compteur d'erreurs à J+1, J+7.
Et serre les vis (section 86).

## 334. Étude de cas : l'orage

Scénario type : micro-coupures répétées, onduleur sous-dimensionné
pour le pic de spin-up (section 200). Les nœuds tombent un par un,
Ceph passe en degraded, puis les NUT arrêtent le reste. Au redémarrage :
spin-up simultané → l'onduleur disjoncte → re-belote. Leçon : **teste
le redémarrage complet après coupure**, pas seulement l'arrêt. C'est
le test que personne ne fait et qui sauve les mises en production.

## 335. Étude de cas : le firmware tueur

Scénario type (documenté industrie) : un bug firmware sur un lot de
SSD provoque des pannes corrélées (même heure, même modèle). Le
cluster survit parce que les disques étaient **mélangés par lot
d'achat** (2 fournisseurs, section 294). Leçon : le double sourcing
n'est pas du luxe, c'est de l'assurance. Et ne flashe jamais tout le
parc le même jour (section 226).

## 336. La checklist de mise en production

