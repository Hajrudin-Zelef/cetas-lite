---
id: collect-261001-rattrapage/rattrapage/datacenter-stockage-guide-8
title: "STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-27"]
keywords: ["attention"]
source: docs/RAG/collect-261001-rattrapage/datacenter_stockage_guide.md
source_anchor: ""
source_lines: [1200, 1386]
sha256: b7b22b180c0289dbac93d7d23a004500d5025872838bd0cebd1db99021ee3154
---

# STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur

Quand un OSD meurt, Ceph recopie ses PG sur les survivants (backfill).
Sur full NVMe, le recovery est **très rapide** (bon) mais **sature le
réseau et le CPU** (mauvais si sous-dimensionnés). Réglages : limite
`osd_max_backfills` et `osd_recovery_max_active` pendant les heures
ouvrées, ouvre les vannes la nuit. Teste **avant** la production :
tue un OSD en burn-in et mesure l'impact client.

## 123. BlueStore : ce qu'il faut savoir

BlueStore écrit directement sur le block device (pas de filesystem
intermédiaire), avec RocksDB pour les métadonnées (omap) et checksums
intégraux. Filestore est **obsolète** (vérifié : docs 2026). Conséquence
pratique : pas de `mkfs` sur les disques OSD, et les métadonnées
RocksDB méritent un device rapide (d'où le block.db en hybride).

## 124. WAL/DB en full NVMe : colocalisés

En full NVMe, pas de device séparé : `block.wal` et `block.db` sont
**colocalisés** sur le même SSD (comportement par défaut de BlueStore
quand il n'y a pas de mix de médias — vérifié). Ne crée pas de LV
séparés « pour optimiser » : tu fragmenterais pour rien.

## 125. mClock et QoS (aperçu)

dmClock/mClock : l'ordonnanceur QoS de Ceph (client vs recovery vs
scrub). Sur full NVMe, active et profile le mClock : c'est ce qui
empêche un scrub ou un backfill d'écraser la latence client. Profils :
`high_recovery_ops` en journée, équilibrage la nuit. À tester par
workload — il n'y a pas de profil universel.

## 126. RGW sur NVMe : particularités

Le RGW (S3) génère énormément de petites clés omap → RocksDB sollicité.
En full NVMe ça tient, mais prévois : **plus de RAM** (cache omap),
et si tu fais de l'EC 8+3 sur RGW, le CPU d'encodage par nœud (voir
section 112, prends la fourchette haute).

## 127. CephFS sur NVMe

CephFS = MDS (métadonnées) + RADOS (données). Sur NVMe : MDS très
réactifs, idéal pour homes, builds, VDI. Points d'attention : dimensionne
les **MDS** (RAM ! les MDS adorent la RAM — 64 Go+ par MDS actif),
et sépare le pool `cephfs_metadata` (réplica 3, petits objets) du pool
data (EC possible pour le froid).

## 128. RBD sur NVMe

Le cas roi du full NVMe : images VM (Proxmox !), bases. Latence typique
bien réglée : < 1 ms p99 en lecture. Réglages qui comptent :
`rbd_cache` côté client, taille d'objet 4 Mo (défaut sain), et surtout
**pas de cache tiering** (déprécié — section 137).

## 129. Erasure coding sur NVMe : le piège latence

L'EC sur NVMe économise des To mais **ajoute de la latence d'écriture**
(encodage + fan-out réseau k+m). Sur du bloc synchrone (VM, bases),
l'EC 4+2 peut doubler la latence d'écriture vs réplica 3. Règle : **EC
pour l'objet et le froid, réplica pour le bloc chaud**. Mesure avant
de décider : `rados bench` sur les deux profils.

## 130. Monitoring : les métriques qui comptent

| Métrique | Seuil d'alerte |
|---|---|
| `ceph health` != HEALTH_OK | immédiat |
| PG `degraded` / `stuck` | immédiat |
| Latence op OSD p99 (commit/apply) | > 50 ms NVMe = anomalie |
| `bluestore_slow_ops` | > 0 répété = disque qui rame |
| Remplissage pool | > 70 % warning, > 85 % critique |
| Recovery/backfill en cours | info, corréler avec la latence client |

Resserre `bluestore_slow_ops_warn_lifetime` (défaut 86400 s) pour détecter
un SSD qui rame **avant** la cascade (retour terrain vérifié 2026 :
lifetime 300 s, threshold 5 — à ajuster).

## 131. Mises à jour : l'autoscaler en pause (vérifié)

Pendant une upgrade Ceph, l'autoscaler est **mis en pause automatiquement**
par cephadm (sauf `mgr/cephadm/pg_autoscale_during_upgrade`). Un split de
PG à mi-upgrade peut ajouter des jours sur un gros cluster (retour
terrain 2026). Procédure : upgrade → vérifie HEALTH_OK → réactive →
laisse l'autoscaler converger → seulement ensuite la charge.

## 132. Checklist design full NVMe

- [ ] 1 OSD = 1 SSD, PG/OSD dans 100-200
- [ ] Autoscaler `on`, pools `bulk: true`
- [ ] RAM : 2-4 Go/OSD + 16 Go système
- [ ] CPU : 1-2 cœurs/OSD
- [ ] Réseau cluster séparé, 100G mini par nœud
- [ ] Recovery testé en burn-in (kill OSD)
- [ ] mClock profilé, monitoring en place
- [ ] Remplissage max 70-80 % en exploitation

---

# PARTIE I — CEPH HYBRIDE (SSD + HDD)

## 133. Principe de l'hybride

Les HDD portent les données (block), les SSD portent les métadonnées
RocksDB + le journal (block.db, qui absorbe le block.wal). Résultat :
la latence des petites écritures devient celle du SSD, le €/To reste
celui du HDD. C'est **l'architecture Ceph capacitaire standard**.

## 134. Sizing block.db / block.wal (vérifié)

Documentation Ceph actuelle (vérifié le 27/09/2026 via docs + retours) :
- **≥ 2,5 %** de la taille du `block` (RocksDB compresse depuis Squid —
  l'ancienne règle « 4 % partout » est obsolète),
- **≥ 4 % pour RGW** (volume de clés omap),
- **1-2 % suffisent pour RBD**.
Exemple : HDD 20 To en RBD → db de 400 Go ; en RGW → 800 Go. **Ne
sous-dimensionne jamais** : un block.db plein = OSD en erreur.

## 135. Ratio OSD par device rapide (vérifié)

- **4-5 OSD HDD par SSD SATA**,
- **≤ 15 OSD HDD par NVMe**.
Exemple : 60 HDD → 12-15 SSD SATA de db, ou 4 NVMe. Le device rapide est
un **domaine de panne** : s'il meurt, tous ses OSD meurent → intègre-le
dans le design CRUSH (ne mets pas les 4 NVMe de db sur les mêmes 4 nœuds
que... enfin, réfléchis à la blast radius).

## 136. Le device rapide comme failure domain

Si 1 NVMe porte le db de 15 OSD et qu'il meurt : 15 OSD down d'un coup =
potentiellement des PG indisponibles si la réplication est mal répartie.
Mitigations : device classes CRUSH séparées, `ceph osd out` rapide
automatisé, et stock de SSD de rechange **identiques** (même taille de
db → remplaçable sans recalcul).

## 137. Cache tiering : ne pas utiliser

Le cache tiering (pool SSD chaud devant pool HDD froid) est **déprécié/
abandonné** dans Ceph moderne : bugs de promotion, pas de maintenance.
Les alternatives saines : **device classes + règles CRUSH** (données
chaudes sur SSD, froides sur HDD, déplacement par politique applicative),
ou deux clusters. Si un commercial te propose du cache tiering Ceph en
2026 : non.

## 138. CRUSH : les device classes

```
ceph osd crush set-device-class hdd osd.0
ceph osd crush set-device-class ssd osd.60
```
Puis des règles `replicated_rule_ssd` / `replicated_rule_hdd`. C'est le
mécanisme propre pour : pool RBD chaud sur SSD, pool backup sur HDD,
le tout dans **un seul cluster**.

## 139. Exemple de règle CRUSH hybride

```
rule rbd-ssd {
    ruleset 1
    type replicated
    min_size 1 max_size 10
    step take default class ssd
    step chooseleaf firstn 0 type host
    step emit
}
```
Traduction : les PG de ce pool ne sont placés que sur des OSD de classe
`ssd`, un par host. Duplique pour `hdd`. Vérifie avec
`ceph osd crush rule dump`.

## 140. Dimensionnement hybride chiffré : 500 To utiles froids

Hypothèses : 500 To utiles, EC 8+3 (37,5 % de surcoût), HDD 20 To.
Brut = 500 × 1,375 × 1,1 ≈ 756 To → **40 HDD de 20 To** (800 To bruts).
Nœuds : 5 nœuds × 8 HDD (blast radius raisonnable, EC 8+3 tient sur 5
nœuds avec min_size adapté — vérifie k+m ≤ nœuds).
DB : 40 × 20 To × 2,5 % = 20 To de db → **3 NVMe de 7,68 To** (ou 2×
15,36 To), répartis. RAM : 40 × 4 Go + système → **256 Go/nœud**.
Réseau : 2×25G par nœud suffisent (froid), cluster séparé.

## 141. Piège : le block.db plein

Symptômes : OSD qui flap, erreurs ENOSPC alors que le HDD est à 50 %.
Cause : db sous-dimensionné ou workload plus « omap » que prévu (RGW !).
Prévention : dimensionne à **4 % pour RGW dès le départ**, monitor
`ceph daemon osd.X perf dump` (espace db), et prévois la procédure
d'extension (recréation d'OSD — il n'y a pas d'extension en place simple).

## 142. Piège : le WAL sur device lent

