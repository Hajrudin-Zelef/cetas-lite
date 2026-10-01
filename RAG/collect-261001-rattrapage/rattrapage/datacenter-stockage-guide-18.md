---
id: collect-261001-rattrapage/rattrapage/datacenter-stockage-guide-18
title: "STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter"]
source: docs/RAG/collect-261001-rattrapage/datacenter_stockage_guide.md
source_anchor: ""
source_lines: [2880, 3083]
sha256: c05e8a8c5fe0797bc0c49982872d7ffd645ec0ed162fd40ad242a278dd357bfe
---

# STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur

Ne qualifie jamais un seul modèle de SSD/HDD : une rupture ou un
bug firmware par lot te laisse sans option. Deux modèles qualifiés
(même capacité, perfs proches) = négociation + résilience. En
pratique : 70/30 sur les volumes, avec le 30 % qui tourne vraiment
(pas juste « qualifié sur papier »).

## 295. Reconditionné : quand c'est pertinent

HDD Exos/Ultrastar reconditionnés garantis 1-2 ans : pertinents pour
du **froid non critique** ou du lab (prix ~50 % du neuf — à vérifier).
Jamais pour : block.db, cache, production critique sans spare
massif. Vérifie : heures de vol (SMART 9), `percentage_used` (SSD),
et bannis les modèles à AFR élevé (section 177).

## 296. Le planning type d'un projet

| Phase | Durée indicative |
|---|---|
| Dimensionnement + devis | 3-4 semaines |
| Commande (allocation 2026) | 4-12 semaines |
| Réception + burn-in | 2-3 semaines |
| Installation + câblage | 1-2 semaines |
| Tests (bench + game day) | 2 semaines |
| Production progressive | 2-4 semaines |

Total : **4-6 mois** de la décision à la production. Le burn-in et
les tests ne se compressent pas : c'est là que les projets « rapides »
meurent.

## 297. Gabarit BOM : nœud Ceph HDD standard

```
Nœud Ceph HDD (×N) :
- Châssis 2U 12 baies 3,5" / backplane expander 12G
- 1× CPU 16c, 128 Go DDR5 ECC
- 1× HBA 9500-16i (IT)
- 12× HDD 20 To CMR (burn-in OK)
- 1× NVMe 3,84 To TLC (block.db, 4 DWPD si RGW)
- 2× 25G NIC (1 pub + 1 cluster) ou 2×25G ×2
- 2× PSU Titanium
- Rails + obturateurs
```

## 298. Gabarit BOM : nœud Ceph NVMe standard

```
Nœud Ceph NVMe (×N) :
- Châssis 2U 16-24 baies E3.S Gen5 (NVMe direct si possible)
- 2× CPU 16c ou 1× 32c, 256 Go DDR5 ECC
- 16× NVMe 7,68 To Gen5 E3.S TLC 1 DWPD
- 2× 100G NIC client + 2× 100G NIC cluster
- 2× PSU Titanium 1600W
- Câblage DAC/fibre + 2 switchs 100G MLAG (mutualisé)
```

## 299. Gabarit : fiche de réception

À chaque livraison : n° série par disque (photo labels), versions
firmware (HBA/expander/SSD), SMART initial archivé, test burn-in
planifié, étiquetage rack/U conforme au plan. **Aucun disque ne
rentre en production sans fiche de réception complète.**

## 300. Gabarit : plan de nommage

`dc1-r07-n03` (datacenter, rangée, nœud), `osd.45`, pools
`rbd-vm-prod-ssd`, `rgw-backup-hdd-ec83`. Un nommage **prédictible**
= des scripts qui marchent et des astreintes moins douloureuses.
Documente la convention, fais-la respecter par tous (humains et
Ansible).

## 301. Ansible : l'automatisation minimale

Playbooks : inventaire disques (`lsblk` + SMART), déploiement
cephadm, collecte `storcli`, mise à jour firmware par vagues,
vérification post-reboot (chrony, OSD up, HEALTH_OK). Tout ce qui
est fait 2× à la main doit être en playbook à la 3e fois.

## 302. La revue annuelle du stockage

1×/an : AFR par modèle, croissance des données vs prévision,
TCO réel vs budget, état des garanties (renouvellement ?),
roadmap (les 245 To arrivent — faut-il attendre pour l'extension ?),
exercice de panne (game day). C'est cette revue qui transforme
l'exploitation en **pilotage**.

## 303. Les indicateurs pour la direction

- €/To utile (TCO 5 ans), par tier (chaud/froid).
- Disponibilité (9x) : mesurée, pas promise.
- Temps moyen de recovery (issu des game days).
- kWh/To/an et PUE du stockage.
- Âge moyen du parc et fin de garantie (risque capacitaire).
Parle la langue de la direction : **€, risques, délais** — pas IOPS.

## 304. Erreurs d'achat classiques (récap)

1. Acheter au €/To brut sans TCO.
2. Un seul fournisseur qualifié.
3. Baies U.2 pour du neuf dense (vs E3.S).
4. HBA en IR « parce que c'était le défaut ».
5. Réseau 25G pour du full NVMe.
6. Pas de spare (le premier RMA arrive toujours un vendredi).
7. Oublier le pic de spin-up dans l'onduleur.
8. Cache tiering Ceph (déprécié).
9. RAID 5 sur 20 To.
10. Pas de burn-in « pour gagner du temps ».

## 305. Ce que Zelef doit exiger de ses fournisseurs

Firmware IT par écrit, baies UBM/E3.S, alimentations Titanium,
garantie 5 ans avec SLA J+1, double sourcing, devis datés avec
délais fermes (allocation 2026), et **un interlocuteur technique**
qui connaît la différence entre un SFF-8654 et un SFF-8644. Si le
commercial ne la connaît pas : change d'interlocuteur.

## 306. Pour aller plus loin (dans tes guides)

- `proxmox_guide.md` : Ceph HCI, NUT, ZFS.
- `onduleurs_ups_guide.md` : dimensionnement électrique, groupes.
- `zabbix_guide.md` : supervision (templates SMART/Ceph).
- `debian_ubuntu_guide.md` : tuning noyau, multipath, chrony.

---

# PARTIE V — APPROFONDISSEMENTS

## 307. ZFS vs Ceph : le duel honnête

| Critère | ZFS | Ceph |
|---|---|---|
| Échelle | 1 serveur (jusqu'à ~1 Po) | cluster (multi-Po) |
| Redondance | RAID-Z / miroir local | réplica / EC distribués |
| Survie nœud | non (sauf réplication) | oui |
| Complexité | modérée | élevée |
| Cas roi | NAS, backup, PME | cloud, S3, VM à l'échelle |

Les deux ne s'opposent pas : ZFS **dans** le nœud (boot, cache),
Ceph **entre** les nœuds. Et Proxmox parle aux deux (voir ton
`proxmox_guide.md`).

## 308. ZFS : recordsize et volblocksize

Le réglage qui change tout : `recordsize=1M` pour du séquentiel
(médias, backup), `recordsize=16K` ou `volblocksize=16K` pour des
VM/bases. Un mauvais recordsize = amplification d'écriture ×8.
Règle : aligne sur la taille d'IO de l'application, et ne change
jamais sur un dataset plein (réécriture nécessaire).

## 309. ZFS : ARC, L2ARC, ZIL/SLOG

- **ARC** : cache lecture en RAM (50 % de la RAM par défaut — à
  ajuster : `zfs_arc_max`).
- **L2ARC** : extension sur SSD (utile si ARC saturé **et**
  workload relisible — sinon du gaspillage).
- **SLOG** : device séparé pour le ZIL (écritures synchrones :
  NFS, bases). **Le SLOG doit être le SSD le plus endurant et le
  plus faible en latence** (3 DWPD+, PLP) — c'est le block.db de ZFS.
Sans SLOG, les écritures sync sur HDD = ~100 IOPS. Avec : ~10 000.

## 310. ZFS : copies=2 et ditto blocks

`zfs set copies=2` : chaque bloc écrit 2× **sur le même pool**.
Protège contre la corruption/bitrot, pas contre la panne disque.
Utile sur du miroir critique ou du single-disk. Ne remplace pas
un vrai miroir.

## 311. mdadm : le RAID logiciel Linux

Toujours pertinent en 2026 pour : boot en RAID 1, petits volumes,
serveurs sans ZFS/Ceph. `mdadm --create /dev/md0 --level=6
--raid-devices=8`. Limites : pas de checksums (corruption silencieuse
possible), rebuild lent comme tout RAID. Pour du sérieux : ZFS.

## 312. LVM : thin provisioning et cache

LVM thin : sur-allocation (pratique en virtualisation). LVM cache :
SSD devant HDD (le « tiering » local qui marche, contrairement au
cache tiering Ceph — car local et simple). `lvmcache` en writethrough
pour la lecture, writeback avec prudence (perte du cache = perte de
données en writeback sans redondance).

## 313. Filesystems : ext4 vs XFS en 2026

| Critère | ext4 | XFS |
|---|---|---|
| Maturité | maximale | maximale |
| Gros fichiers | bien | **excellent** |
| Parallélisme | correct | meilleur |
| Réparation (fsck) | rapide | plus lent sur gros volumes |

Pour OSD Ceph : BlueStore, pas de FS. Pour le reste (boot, données
applicatives) : XFS par défaut en 2026 sur RHEL-like, ext4 ailleurs.
Les deux sont sains ; btrfs reste déconseillé en production RAID.

## 314. Tuning Linux : le scheduler

- NVMe : `none` (pas de scheduler — le SSD gère).
- SSD SATA/SAS : `none` ou `mq-deadline`.
- HDD : `mq-deadline` ou `bfq` (ce dernier si multi-tenant).
Vérifie : `cat /sys/block/sdX/queue/scheduler`. Un `cfq` oublié sur
du NVMe = latence +30 % (vécu).

## 315. Tuning Linux : nr_requests et read_ahead

