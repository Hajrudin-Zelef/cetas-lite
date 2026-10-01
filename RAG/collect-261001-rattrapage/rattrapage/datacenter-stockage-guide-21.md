---
id: collect-261001-rattrapage/rattrapage/datacenter-stockage-guide-21
title: "STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-27"]
keywords: ["cyber", "incident", "latency", "packaging"]
source: docs/RAG/collect-261001-rattrapage/datacenter_stockage_guide.md
source_anchor: ""
source_lines: [3450, 3628]
sha256: 6003cfa8b921d392541e2db19d86bd6f35b46dd05a23e007215805a56a33cca3
---

# STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur

La passerelle NVMe/TCP pour RBD (projet en maturation 2025-2026) :
le successeur logique de l'iSCSI gateway, avec la latence NVMe.
État : à qualifier sur ta version Ceph (**à vérifier** au 27/09/2026
selon release). Si tu dois connecter des initiateurs NVMe à Ceph en
2026 : pilote d'abord, production ensuite.

## 355. Le client kernel CephFS vs FUSE

| Client | Perf | Usage |
|---|---|---|
| kernel (`mount -t ceph`) | meilleure | serveurs Linux, stable |
| FUSE (`ceph-fuse`) | moindre | debug, non-root, tests |

En production : **toujours le client kernel**. Le FUSE ne sert qu'au
dépannage. Et le client kernel doit matcher à peu près la version du
cluster (pas 3 versions d'écart).

## 356. Telemetry Ceph : opt-in/out

Ceph collecte des stats anonymes (opt-in). En entreprise : décide
explicitement (souvent opt-out pour des raisons de conformité), et
documente le choix. Les données de télémétrie sont utiles à la
communauté — si ta politique le permet, l'opt-in aide tout le monde.

## 357. SELinux / AppArmor et Ceph

Sur RHEL : SELinux en enforcing avec les policies Ceph = OK en 2026
(les soucis historiques sont réglés). Ne le désactive pas « parce que
le tuto le dit » : un cluster qui ne marche qu'en permissive a un
problème de packaging, pas de SELinux. Sur Ubuntu : AppArmor, même
logique.

## 358. Les versions Ceph : Reef, Squid, Tentacle

Releases récentes (noms de céphalopodes) : Quincy → Reef → Squid →
Tentacle (2025). Règle : reste sur une version **supportée** (N et
N-1), upgrade 1×/an. Les nouveautés (Crimson, NVMe-oF gateway)
arrivent par release : lis les release notes avant chaque upgrade,
surtout les « upgrade notes » (changement de défauts).

## 359. Le dashboard Ceph : ce qu'il vaut

Le dashboard (MGR) : santé, PG, OSD, perf, RGW. Utile pour le
quotidien et les démos. Limites : pas un vrai monitoring (pas
d'alerting fin, pas d'historique long) → **toujours** doublé par
Prometheus/Grafana ou Zabbix (voir ton `zabbix_guide.md`).

## 360. Prometheus + Grafana : les métriques d'or

`ceph_osd_op_w_latency`, `ceph_osd_op_r_latency` (p99 !),
`ceph_pg_degraded`, `ceph_osd_numpg`, `ceph_cluster_total_bytes` vs
`ceph_cluster_total_used_bytes`. Dashboard : santé + latence p99 par
pool + remplissage + recovery. Alerte : latence p99 > 2× la baseline
pendant 15 min = page.

## 361. Le benchmarking continu

Un bench mensuel automatisé (même `rados bench` 5 min sur pool dédié,
même heure) crée une **baseline**. Quand la latence dérive de 20 %
sur 3 mois : enquête (SSD qui vieillissent ? fragmentation ? réseau ?).
Sans baseline, toute dérive lente est invisible jusqu'à la plainte
utilisateur.

## 362. PRA : le plan de reprise d'activité

Le stockage est au cœur du PRA : RPO/RTO par application, pas par
technologie. Tableau type :

| Application | RPO | RTO | Moyen |
|---|---|---|---|
| VM critiques | 15 min | 1 h | réplication RBD + snapshots |
| S3 backup | 24 h | 4 h | RGW multisite |
| Archivage | 7 j | 72 h | bande hors site |

**Teste le PRA 1×/an** en vraie bascule (pas en « exercice papier »).
Un PRA non testé est une fiction coûteuse.

## 363. L'exercice de bascule : le protocole

1. Annonce, freeze des changements.
2. Bascule (promote rbd-mirror / DNS RGW / mount CephFS distant).
3. Validation applicative (pas juste « le ping passe »).
4. Mesure du RTO réel vs cible.
5. **Failback** (le retour est souvent plus dur que l'aller).
6. Post-mortem sans blâme, actions correctives datées.

## 364. Post-mortem sans blâme

Après chaque incident significatif : timeline factuelle, causes
racines (les 5 pourquoi), actions (qui, quoi, quand). **Jamais** de
nom de coupable dans le document. Un post-mortem qui cherche un
coupable produit des gens qui cachent les incidents ; un post-mortem
sans blâme produit des systèmes qui s'améliorent.

## 365. L'astreinte : l'organisation

Rotation, pas de « toujours le même ». Runbooks imprimés (section 336),
accès VPN + IPMI + dashboard testés **hors incident**, escalade écrite
(quand appeler le N+1 ? le fournisseur ?). Et : une astreinte qui sonne
plus de 2×/mois = un problème de fond (seuils, spares, ou architecture),
pas de malchance.

## 366. La formation de l'équipe

Niveau 1 (tous) : SMART de base, remplacement disque, lecture `ceph -s`.
Niveau 2 (référents) : CRUSH, PG, recovery, upgrade. Niveau 3 (expert) :
tuning, debug profond, relation éditeur. **Au moins 2 personnes** par
niveau critique (le bus factor de 1, c'est une panne en attente).

## 367. Green IT : le carbone du stockage

Deux postes : carbone **opérationnel** (kWh × intensité du mix) et
carbone **embarqué** (fabrication : ~50-100 kg CO₂e par SSD, ordre de
grandeur à vérifier — un HDD c'est ~20-40 kg). Sur 5 ans, un nœud
NVMe à 800 W en France (~50 g/kWh) = ~1,75 t CO₂e opérationnel.
Le levier n°1 : **allonger la durée de vie** (6-7 ans au lieu de 5)
et **ne pas sur-dimensionner** (le To inutile est du carbone inutile).

## 368. Le reconditionné : le bilan carbone

Un HDD reconditionné évite la fabrication d'un neuf (~30 kg CO₂e).
Pour du froid non critique : le reconditionné garanti est le meilleur
rapport €/To **et** CO₂/To. Intègre-le à la politique d'achat avec des
règles claires (section 295) au lieu de le subir en « bon plan ».

## 369. La fin de vie : DEEE et effacement

Disques en fin de vie : effacement certifié (section 286) puis filière
DEEE agréée (bordereau de suivi). Les aimants des HDD et les métaux
des SSD se recyclent ; les plateaux partent en filière. **Trace** :
n° série → certificat d'effacement → BSD. Un disque « jeté » sans
traçabilité, c'est un risque juridique.

## 370. L'assurance : le risque résiduel

Même bien conçu, un stockage peut perdre des données (incendie,
ransomware, erreur humaine en chaîne). L'assurance cyber/casse couvre
le financier, pas les données. Traduction : l'assurance ne remplace
**jamais** le 3-2-1 et le PRA testé. Elle paie la facture, pas la
reconstruction.

## 371. La conformité : ce qui s'applique

Selon secteur : RGPD (droit à l'effacement → crypto-shredding via
SED, section 287), HDS/secret (hébergement agréé), bancaire (DORA),
industrie (NIS2). Le stockage est concerné par : chiffrement,
traçabilité des accès (audit log), durées de rétention, et **droit
à l'effacement effectif** (un backup immuable Object Lock de 7 ans
vs un droit à l'effacement = tension juridique à arbitrer en amont).

## 372. Le chiffrement et le droit à l'effacement

Le crypto-shredding (destruction de la clé) est la réponse technique
au droit à l'effacement sur stockage distribué : impossible d'aller
réécrire chaque réplique, mais sans la clé, c'est du bruit. D'où
l'intérêt du chiffrement **par tenant/projet** (clés distinctes) :
effacer = détruire la clé du tenant.

## 373. Le capacity planning : la méthode

1. Mesure la croissance réelle (Go/jour par pool, 6 mois d'historique).
2. Projette linéairement + marge 30 %.
3. Date de saturation = (capacité × 0,7 - utilisé) / croissance.
4. Commande à **saturation - délai d'appro** (12 semaines en 2026 !).
5. Revois trimestriellement.
Un tableau, 5 lignes, mis à jour chaque trimestre : c'est tout ce
qu'il faut. La plupart des saturations « surprises » sont des
plannings non mis à jour.

## 374. La déduplication des projets

Avant d'acheter de l'extension : chasse le gaspillage. Les classiques :
snapshots oubliés (RBD/ZFS), buckets RGW sans lifecycle, VM éteintes
avec disques alloués, 3 copies du même dataset. Un « nettoyage de
printemps » annuel récupère souvent **10-20 %** — soit un an de
croissance offert.

## 375. Les politiques de lifecycle S3

