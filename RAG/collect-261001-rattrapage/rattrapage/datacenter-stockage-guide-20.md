---
id: collect-261001-rattrapage/rattrapage/datacenter-stockage-guide-20
title: "STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: ["2026-09-27"]
keywords: ["datacenter", "amd", "arr", "intel", "nvidia", "training"]
source: docs/RAG/collect-261001-rattrapage/datacenter_stockage_guide.md
source_anchor: ""
source_lines: [3265, 3449]
sha256: 459c0519b2ab9483cf3548687c57c8f72f658f773cce738f84be42819dcc431b
---

# STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur

- [ ] Burn-in OK, fiches de réception complètes
- [ ] Firmware inventorié, mode IT vérifié
- [ ] PG/OSD dans 100-200, autoscaler on
- [ ] Réseau : MTU, LACP, VLAN séparés, 0 erreur
- [ ] Électrique : PDU A/B, NUT testé 2× (arrêt + redémarrage)
- [ ] Supervision : seuils, alertes testées
- [ ] Runbooks imprimés (pas juste dans le wiki)
- [ ] Game day n°1 planifié (kill OSD)
- [ ] Spares en stock, RMA testé une fois
- [ ] Doc d'exploitation relue par un tiers

## 337. La checklist d'audit annuel

- [ ] AFR par modèle, bannissements éventuels
- [ ] Garanties : expirations à 12 mois
- [ ] TCO réel vs prévisionnel
- [ ] Croissance vs capacité : date de saturation
- [ ] Firmware : versions à jour ?
- [ ] Game day réalisé, résultats archivés
- [ ] Clés cephx / RGW : rotation et nettoyage
- [ ] Test de restauration backup (pas juste « backup OK »)

## 338. Lexique énergie du stockage (pour ton métier)

| Terme | Sens stockage |
|---|---|
| W/To | puissance par To utile (1 = HDD froid, 30+ = NVMe) |
| PUE | multiplicateur salle (1,3-1,6 typique) |
| Spin-up | pic 2× au démarrage des HDD (10-20 s) |
| Titanium | 96 % de rendement alim (imposé au neuf) |
| kWh/To/an | l'unité du TCO énergie (~15 pour HDD, ~700 pour NVMe) |
| ΔT air | 1,7 m³/h par watt pour 2 °C (règle terrain) |

## 339. Formules à garder sous la main

```
PG            = (100 × OSD) / réplica        → puissance de 2
Brut utile    = To_utiles × (1+surcoût) × 1,15 (marge)
Disques       = Brut / To_par_disque (en Tio !)
block.db      = HDD × 2,5 % (4 % RGW)
RAM nœud      = OSD × 3-4 Go + 16 Go
CPU nœud      = OSD × 1-2 cœurs
Réseau nœud   = débit_disques × réplica × 1,2 (pire cas)
W_toit        = W_nœuds × PUE ; Onduleur ≥ 1,5 × W_toit (+ pic spin-up)
Rebuild       = To_disque / débit_effectif
€/To/an       = TCO_5ans / 5 / To_utiles
```

## 340. Les 10 commandements du stockage datacenter

1. Un OSD par disque, toujours.
2. IT mode ou rien (sous Ceph/ZFS).
3. 100-200 PG par OSD, autoscaler on.
4. Le réseau se dimensionne après les disques, jamais avant.
5. Burn-in avant production, sans exception.
6. Le RAID n'est pas un backup ; le snapshot non plus.
7. 70 % = alerte, 85 % = critique, jamais 95 %.
8. Titanium, PDU A/B, NUT testé 2×/an.
9. Deux fournisseurs qualifiés, toujours.
10. Si ce n'est pas écrit (doc, runbook), ça n'existe pas.

## 341. Tableau final : que choisir en 2026 (résumé décideur)

| Besoin | Architecture | Ordre de grandeur €/To utile |
|---|---|---|
| VM critiques, < 1 ms | Ceph NVMe réplica 3 / HCI | ~2 000 € |
| S3 / backup chaud | Ceph EC sur TLC ou QLC | ~300-500 € |
| Archivage froid | HDD EC 8+3 / LTO | ~150 € / ~30 € |
| IA training | NVMe Gen5 + 400G / WEKA | à chiffrer |
| PME simple | 3 nœuds HCI NVMe | ~300 € |

## 342. Dernier mot : la méthode Zelef

Ton métier, c'est l'énergie et les systèmes : applique la même
rigueur au stockage qu'à une note de calcul électrique. Chaque
dimensionnement de ce guide suit la même méthode : **hypothèses
écrites → formules → chiffres → marges → vérification terrain**.
Un dimensionnement sans hypothèses écrites est une opinion ; avec
des hypothèses écrites, c'est un document d'ingénierie. Et quand
un commercial te dit « ça passe », demande-lui de signer le calcul.

---

**FIN DU GUIDE — 342 sections. Vérifications web au 27/09/2026.**
**Rappel : prix en allocation tendue (sept. 2026) — devis datés exigés, à vérifier.**
**Guides liés : `proxmox_guide.md`, `onduleurs_ups_guide.md`, `zabbix_guide.md`, `debian_ubuntu_guide.md`.**

---

# PARTIE W — TUNING AVANCÉ ET SUJETS EXPERTS

## 343. msgr v2 et chiffrement du trafic cluster

Ceph Messenger v2 : protocole sécurisé (authentification + chiffrement
optionnel du trafic OSD). Active `ms_cluster_mode secure` sur les
clusters sensibles. Coût : CPU (AES-NI quasi gratuit sur CPU modernes).
Sur 100G+, mesure avant/après : le chiffrement peut coûter 5-10 %
de débit sur petits IO.

## 344. Plugins d'erasure coding : ISA-L vs Jerasure

| Plugin | Vitesse | Usage |
|---|---|---|
| `isa-l` (Intel) | rapide (SIMD) | **défaut recommandé** x86 |
| `jerasure` | portable | non-x86, compatibilité |
| `shec` | tolère plus | cas exotiques |

`ceph osd erasure-code-profile set` : choisis `isa-l` sur x86_64,
toujours. Un profil EC créé en `jerasure` par défaut sur vieux
tutos = 2-3× plus lent à l'encodage.

## 345. Crimson OSD : l'avenir (à surveiller)

Crimson = la réécriture des OSD sur Seastar (C++ asynchrone), pour
tirer parti du NVMe (moins de threads, plus de débit par cœur). État
2026 : en développement actif, pas encore le défaut. À surveiller :
quand Crimson deviendra le défaut, les règles CPU/OSD de ce guide
(section 112) changeront (moins de cœurs par OSD).

## 346. NVMe ZNS : les zones

Zoned Namespaces : le SSD expose des zones à écriture séquentielle,
l'hôte gère le placement (fini l'amplification d'écriture cachée).
Intérêt : endurance ×2-4 sur workloads adaptés, QoS prédictible.
État 2026 : support Linux OK, adoption applicative faible (il faut
des applis « zone-aware »). Pour Ceph : expérimental — **à surveiller**,
pas à déployer en 2026.

## 347. DPU / SmartNIC : l'offload

Les DPU (NVIDIA BlueField, AMD Pensando) déportent réseau, chiffrement
et parfois stockage du CPU. Usage stockage : NVMe-oF target offload,
chiffrement IPSec/TLS, isolation. En 2026 : pertinent sur les nœuds
très denses (32 NVMe) où le CPU est le goulot. Coût : prix + complexité.
À évaluer quand le CPU du nœud sature avant le réseau.

## 348. SPDK : le fast path

Storage Performance Development Kit : drivers userspace poll-mode
(zéro interruption, zéro copie). Utilisé par les targets NVMe-oF
haute performance et certains SDS. Pour l'exploitant Ceph classique :
pas d'action directe, mais c'est ce qui tourne sous les baies NVMe-oF
que tu achètes. Retiens : si un vendeur annonce « 5 M IOPS », c'est
souvent du SPDK + RDMA.

## 349. La loi de Little : le modèle mental

`IOPS × latence = profondeur_de_file`. Si tu veux 100K IOPS à 1 ms :
il faut QD = 100. Si ton application ne fait que QD=8 : max 8K IOPS
à 1 ms, **quel que soit le SSD**. 90 % des « problèmes de perfs
stockage » sont des problèmes de parallélisme applicatif, pas de
disques. Mesure le QD réel (`iostat -x`, `avgqu-sz`) avant d'acheter
du plus rapide.

## 350. Le dimensionnement par la file d'attente

Formule inverse : `latence_cible = QD_appli / IOPS_cibles`. Exemple :
appli à QD=32 qui veut 50K IOPS → latence max = 32/50000 = 0,64 ms.
Si ton Ceph fait 2 ms p99 : impossible, il faut plus de parallélisme
ou moins de latence. Ce calcul tient sur un post-it et évite des
achats inutiles.

## 351. NFS Ganesha sur CephFS

Pour servir du NFSv4 depuis CephFS : NFS-Ganesha avec backend FSAL_CEPH.
Usage : applis legacy qui ne parlent que NFS. Pièges : le Ganesha est
stateful (HA = failover actif/passif ou cluster), et les verrous NFS
sur CephFS distribué ont leurs quirks. À tester par workload.

## 352. Samba / SMB sur CephFS

De même avec Samba (`vfs_ceph`). Usage : homes Windows, bureautique.
Le duo Samba+CTDB+CephFS donne du SMB HA. En 2026 c'est mature pour
de la bureautique ; pour du très gros débit SMB, une appliance
(PowerScale) reste plus simple.

## 353. iSCSI gateway Ceph (tcmu)

`ceph-iscsi` (tcmu-runner) : expose des images RBD en iSCSI avec
multipath et HA (2+ gateways). Usage : hyperviseurs ou applis qui
ne parlent que iSCSI (VMware !). Latence ajoutée ~0,5 ms. C'est la
passerelle propre pour intégrer Ceph à un existant VMware/Hyper-V.

## 354. NVMe-oF gateway Ceph

