---
id: collect-261001-rattrapage/rattrapage/datacenter-stockage-guide-16
title: "STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-rattrapage/datacenter_stockage_guide.md
source_anchor: ""
source_lines: [2525, 2689]
sha256: e2abd2277ef54d25cd03ed1d757b3bb3831593efe36fec825bc607d52c118728
---

# STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur

1. `ceph osd out <id>` (le cluster recopie).
2. Attends `HEALTH_OK` (ou `active+clean` sur les PG concernées).
3. `ceph osd purge <id> --yes-i-really-mean-it` (note l'ID).
4. Remplace physiquement (LED locate via SES : `sg_ses --set=ident`).
5. `ceph-volume lvm create --data /dev/XXX` (recrée l'OSD).
6. Vérifie le crush (même host, même device class).
7. Burn-in du disque **retiré** avant de le mettre en spare.
Temps total : 10 min d'ops + heures de backfill. **Ne purge jamais
avant le `out` complet.**

## 260. Runbook : remplacement d'un disque RAID matériel

1. Identifie via `storcli /c0/eX/sY show` + LED.
2. `storcli /c0/eX/sY set offline` puis `set missing` si besoin.
3. Remplace à chaud.
4. Le rebuild démarre (auto ou `start rebuild`).
5. Surveille `%` et `storcli /c0 show rebuild`.
6. **Ne reboote pas** pendant le rebuild (sauf urgence).
Un rebuild interrompu repart de zéro sur beaucoup de cartes.

## 261. Runbook : upgrade Ceph (cephadm)

1. Sauvegarde la conf (`ceph config dump`) et vérifie les backups.
2. `ceph orch upgrade check --image ...`.
3. `ceph orch upgrade start --image ...` (par défaut : 1 daemon à
   la fois, avec `noout` géré).
4. Surveille `ceph orch upgrade status` + `ceph health`.
5. **Ne touche pas aux PG** pendant l'upgrade (autoscaler en pause,
   section 131).
6. Valide : `ceph versions`, bench rapide, puis rouvre les vannes.
Rollback : `ceph orch upgrade stop` — teste-le en lab d'abord.

## 262. Runbook : extension d'un JBOD

1. Câble le nouveau JBOD (domaines A/B, étiquettes).
2. Vérifie la découverte : `storcli /c0 show all` (nouveaux
   enclosures + disques).
3. Burn-in des nouveaux disques **avant** de les confier au cluster.
4. Ajoute au cluster (Ceph : `ceph orch daemon add osd` / `ceph-volume`).
5. Laisse le rebalance converger (norebalance en heures ouvrées).
6. Mets à jour le plan de câblage et la doc.

## 263. Runbook : panne HBA

Symptôme : tous les disques d'un domaine disparaissent d'un coup.
1. `lspci` : la carte est-elle visible ? Non → reseat, puis RMA.
2. Visible mais pas de disques → câble, puis expander, puis firmware.
3. Bascule sur le domaine B (multipath) le temps du diagnostic.
4. **Stock de rechange** : une HBA identique flashée en IT en spare
   (le flash prend 30 min que tu n'as pas en pleine panne).

## 264. Runbook : corruption suspectée (ZFS/Ceph)

1. `zpool scrub` / `ceph pg repair` sur le périmètre suspect.
2. Si erreurs persistantes : isole le disque (`smartctl`, `nvme
   error-log`), remplace.
3. Si erreurs sur plusieurs disques : suspecte **RAM** (memtest),
   HBA/câble, ou alimentation (tension instable = corruption).
4. Restaure depuis backup si la redondance est dépassée — d'où la
   règle : le RAID/Ceph n'est pas un backup (section 89).

## 265. Runbook : cluster Ceph en HEALTH_WARN

1. `ceph health detail` : lis **tout**, dans l'ordre.
2. `ceph -s` : OSD down ? PG degraded ? nearfull ?
3. `ceph osd tree` : localise (host ? rack ?).
4. Causes fréquentes : OSD down (disque/HBA), clock skew (chrony !),
   PG stuck (pg_num), nearfull (extension).
5. **Ne redémarre jamais tous les MON/OSD « pour voir »** : tu
   transformes un WARN en disaster.

## 266. Le clock skew : l'ennemi silencieux

Ceph exige des horloges synchronisées (MON quorum, heartbeats).
> 50 ms de décalage = WARN, > 1 s = OSD flapping. Cause n°1 :
**chrony mal configuré** ou VM sans horloge fiable. Vérifie
`chronyc tracking` sur **chaque** nœud après chaque reboot. Un cluster
qui flap « sans raison » a souvent une horloge qui dérive.

## 267. Supervision : quoi monitorer (Zabbix/Prometheus)

- `ceph health`, PG states, OSD down (via MGR/prometheus).
- SMART : attributs 5/187/197/198 (HDD), percentage_used (NVMe).
- Températures : disques, HBA (`storcli`), expander (SES).
- Réseau : erreurs CRC, paquets perdus (le RDMA n'aime pas la perte).
- Énergie : PDU par prise (dérive de conso = disque/ventilo qui meurt).
- Espace : 70 % warn, 85 % crit (section 229).
Voir ton `zabbix_guide.md` pour les templates.

## 268. Les logs à garder

- `ceph log` / `ceph audit log` : 90 jours mini.
- SMART de référence post burn-in : **à vie du disque**.
- Versions firmware par série : à vie du parc.
- Historique des pannes par modèle : le tableau qui justifie les
  bannissements (section 177).
En cas de litige garantie : le log fait foi, pas ta mémoire.

## 269. RMA : la discipline

1. Photo du label (série) avant démontage.
2. Secure erase si possible (section 305).
3. Emballage antistatique d'origine (un disque RMA mal emballé =
   refus de garantie).
4. Suivi : date d'envoi, n° RMA, délai. Un fournisseur qui dépasse
   15 jours ouvrés deux fois = renégocie le contrat ou change.
5. Le disque de remplacement subit un **burn-in** (neuf ou reconditionné).

## 270. Le stock de rechange : la formule

`spare = max(2, 5 % du parc)` par modèle de disque, + 1 HBA + câbles
de chaque type + 1 alimentation par modèle. Pour 600 HDD : 30 disques
en spare (~15 k€) contre une indisponibilité. Le spare se **teste**
(burn-in) et se **renouvelle** (FIFO : le plus vieux part en premier).

---

# PARTIE S — RÉSEAU DE STOCKAGE APPROFONDI

## 271. iSCSI : le bloc sur Ethernet

iSCSI = SCSI sur TCP/IP. Simple, partout, 10/25/100G. Latence : ~0,5-1 ms.
Usage : connecter des serveurs à une baie ou un cluster sans Fibre Channel.
Limites : le TCP ajoute de la latence vs FC/NVMe-oF ; le multipath iSCSI
se configure (2 sous-réseaux, pas de LACP sur une session). En 2026 :
correct pour la PME, dépassé pour le NVMe performant.

## 272. Fibre Channel : le roi historique du SAN

16/32/64G FC, latence ~0,2-0,5 ms, lossless par construction, zoning
matériel. Toujours présent dans les SAN entreprise classiques. Mais :
cartes + switchs propriétaires et chers, compétences rares, et le
NVMe-oF sur Ethernet grignote son terrain. Pour un **nouveau** projet
2026 : ne choisis le FC que si l'existant l'impose.

## 273. NVMe-oF : le NVMe sur le réseau

NVMe over Fabrics : le protocole NVMe encapsulé sur TCP, RDMA (RoCE/IB)
ou FC. Latence ajoutée : ~5-10 µs en RDMA — quasi du local. Usage :
désagréger le flash (baies NVMe partagées façon SAN moderne), ou
connecter des initiateurs à un cluster. En 2026 : de plus en plus
d'offres (baies EDSFF + NVMe-oF/TCP), support Linux mature
(`nvme-cli`, discovery).

## 274. NVMe/TCP vs NVMe/RDMA

| Transport | Latence ajoutée | Matériel | Complexité |
|---|---|---|---|
| NVMe/TCP | ~20-50 µs | Ethernet standard | faible |
| NVMe/RDMA (RoCE v2) | ~5-10 µs | NIC RDMA, lossless | élevée |

Pour 95 % des usages : **NVMe/TCP suffit** et s'opère comme de l'iSCSI.
Le RDMA se justifie quand chaque microseconde compte (trading, OLTP
extrême) — avec le coût ops du lossless Ethernet (PFC, ECN).

## 275. RoCE : le lossless à configurer

RDMA over Converged Ethernet exige un réseau **sans perte** : PFC
(Priority Flow Control) + ECN sur tous les switchs du chemin. Une
erreur de config PFC = deadlocks et perfs catastrophiques. Règle :
si tu n'as pas un admin réseau qui **maîtrise** le DCB, reste en
NVMe/TCP ou iSCSI. Le RDMA mal configuré est pire que le TCP.

## 276. Topologie réseau Ceph : le schéma complet

