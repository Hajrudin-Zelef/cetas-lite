---
id: collect-261001-rattrapage/rattrapage/datacenter-stockage-guide-17
title: "STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-27"]
keywords: ["datacenter", "attention", "capex", "ethernet", "incident", "nand"]
source: docs/RAG/collect-261001-rattrapage/datacenter_stockage_guide.md
source_anchor: ""
source_lines: [2690, 2879]
sha256: b3e3351ec67dfafc2c2ba5b170c4a88f7324095e8618e639a16718ddb42a3128
---

# STOCKAGE DATACENTER — Serveurs, HBA, SAS, NVMe, JBOD, Ceph : présent et futur

```
                    ┌─────────────┐
                    │ 2× switch   │
                    │ 100G (MLAG) │
                    └──────┬──────┘
              ┌────────────┼────────────┐
              ▼            ▼            ▼
        ┌──────────┐ ┌──────────┐ ┌──────────┐
        │ Nœud 1   │ │ Nœud 2   │ │ Nœud 3   │ ...
        │2×100G clu│ │2×100G clu│ │2×100G clu│  ← VLAN cluster
        │2×100G pub│ │2×100G pub│ │2×100G pub│  ← VLAN client
        └──────────┘ └──────────┘ └──────────┘
```
2 switchs en MLAG (pas de SPOF), LACP 2 liens par VLAN par nœud,
MTU 9000 (jumbo) sur le VLAN cluster **de bout en bout** (un seul
équipement en 1500 = fragmentation silencieuse).

## 277. MTU 9000 : le checklist

Jumbo frames : -10 à -20 % de CPU réseau, +débit. Mais : **tous** les
équipements du VLAN (NIC, switchs, BMC si partagé — ne partage pas),
et vérifié avec `ping -M do -s 8972`. Le classique : ça marche en
lab (1 switch) et ça casse en prod (un uplink oublié en 1500).

## 278. Bande passante : Ethernet 2026

| Débit | Usage stockage 2026 | Note |
|---|---|---|
| 25G | nœud HDD/hybride, client | le standard économique |
| 100G | nœud NVMe, cluster Ceph | le standard perf |
| 200/400G | gros nœuds 32 NVMe Gen5, IA | prix en baisse, à vérifier |
| 800G | inter-switch, backbone IA | early 2026 |

Règle : le réseau se choisit **après** le calcul de la section 113,
pas avant. Un cluster full NVMe en 25G, c'est une Ferrari avec des
roues de vélo.

## 279. Latence réseau : les ordres de grandeur

| Segment | Latence typique |
|---|---|
| NVMe local | 50-100 µs |
| NVMe-oF RDMA | +5-10 µs |
| iSCSI 25G | 0,5-1 ms |
| Ceph RBD réplica 3 (100G) | 0,5-2 ms |
| Ceph EC (encodage + fan-out) | 2-5 ms |
| Inter-sites 50 km | +0,5 ms (RTT) |

Si ton application exige < 1 ms : NVMe local ou NVMe-oF/RDMA, **pas**
Ceph réplica 3 (sauf cluster très bien réglé et peu chargé).

## 280. Le réseau et l'énergie

4 NIC 100G ≈ 80 W par nœud (section 185) + 2 switchs 32×100G à
~300 W chacun. Sur un cluster 8 nœuds : ~1,2 kW rien que pour le
réseau, soit ~15 % du budget. Les NIC 400G consomment plus (~30-40 W
par port, ordre de grandeur à vérifier) : le réseau du stockage IA
est un poste énergétique à part entière.

## 281. Erreurs réseau : les signes avant-coureurs

`ethtool -S` : CRC, dropped, fifo_errors. En stockage : **0 erreur
tolérée** en régime établi. Une erreur CRC de temps en temps =
câble/fibre/ transceiver à changer **maintenant**. Les problèmes
intermittents réseau donnent des symptômes « disque lent » (slow ops)
qui font accuser les SSD à tort : vérifie le réseau **avant** de RMA
un SSD.

## 282. Sécurisation du réseau stockage

VLAN dédiés, pas de routage vers Internet, ACL sur les switchs,
authentification Ceph (cephx) activée, et pour RGW : TLS partout.
Le réseau cluster Ceph en clair sur un VLAN dédié = acceptable ;
le même en clair sur le LAN bureautique = non.

---

# PARTIE T — SÉCURITÉ, CHIFFREMENT, FIN DE VIE

## 283. Le chiffrement : où et pourquoi

| Couche | Solution | Protège contre |
|---|---|---|
| Disque (SED/Opal) | chiffrement matériel | vol du disque |
| Ceph (dmcrypt/LUKS) | OSD chiffrés | vol du nœud |
| Réseau (TLS/IPsec) | RGW TLS, msgr v2 | écoute réseau |
| Application | chiffrement applicatif | tout (bout-en-bout) |

Le minimum datacenter 2026 : **SED ou LUKS** sur les disques +
TLS sur RGW. Le vol d'un JBOD de 60 disques non chiffrés, c'est
1,8 Po de données en clair sur leboncoin.

## 284. SED et TCG Opal (vérifié)

Les SSD entreprise cités (D7-PS1010, CD9P-R, D5-P5430) proposent
**TCG Opal 2.0** (vérifié 27/09/2026), certains en FIPS 140-2/3.
Le chiffrement est matériel (pas de perte de perfs), piloté par
`sedutil-cli`. Point d'attention : la **gestion des clés** (un SED
dont on perd la clé = une brique). Intègre les clés au coffre
(Vault) avec procédure de recouvrement.

## 285. LUKS sur OSD Ceph

`ceph-volume lvm create --dmcrypt` : chaque OSD chiffré. Coût CPU :
~5-15 % sur workloads chiffrés (AES-NI), plus sur petits IO. En
full NVMe : mesure avant/après — sur des CPU modernes le coût est
souvent acceptable. Avantage : pas de dépendance au firmware du
disque, clés centralisables.

## 286. Secure erase : la procédure

Avant RMA, revente ou réforme :
1. `nvme format --ses=2` (crypto erase, instantané sur SED) ou
   `hdparm --security-erase` (HDD),
2. vérifie (lecture de contrôle),
3. **certificat de destruction** pour les disques sensibles
   (dégaussage/broyage si secret défense — filière agréée).
Un disque « formaté » n'est pas un disque effacé : le formatage
rapide ne touche pas les blocs réalloués (SMART 5 !).

## 287. Blocs réalloués et données résiduelles

Les secteurs réalloués (attribut 5) sont **illisibles par l'OS**
mais physiquement présents sur les plateaux/NAND. Sur disque
sensible : crypto-erase (rend tout illisible via la clé) ou
destruction physique. C'est l'argument massue pour le **SED
systématique** : l'effacement = destruction de la clé.

## 288. RGW : Object Lock et immuabilité

S3 Object Lock (WORM) : l'arme anti-ransomware pour les backups
(section 161). Configure : retention compliance sur les buckets
de backup, **avec une durée supérieure à la rétention**. Teste la
restauration **avec** le lock (certains outils de backup ne
gèrent pas la suppression des versions verrouillées).

## 289. Snapshots vs backups

Un snapshot (Ceph RBD, ZFS) n'est **pas** un backup : il est sur le
même cluster (incendie, ransomware, `rbd snap purge` par erreur).
Règle 3-2-1 : 3 copies, 2 supports, 1 hors site. Pour Ceph : export
rbd vers un cluster distant, ou RGW avec réplication + Object Lock.

## 290. Contrôle d'accès Ceph

cephx : chaque client a une clé avec des **capabilities minimales**
(`mon 'profile rbd' osd 'profile rbd pool=vm'`). Ne distribue jamais
la clé admin. Pour RGW : IAM-like (users, buckets policies). Audite
les clés 1×/an : une clé d'ancien prestataire qui traîne, c'est
classique.

## 291. Audit et traçabilité

`ceph audit log` : qui a fait quoi (création pool, purge OSD...).
Envoie-le au SIEM. En cas d'incident : c'est la première source.
Conserve 1 an mini (contrainte réglementaire selon secteur).

---

# PARTIE U — TCO, ACHATS, GABARITS

## 292. TCO 5 ans : le gabarit

| Poste | Année 1 | Années 2-5/an |
|---|---|---|
| CAPEX matériel | X | — |
| Énergie (kWh × PUE × prix) | — | Y |
| Maintenance / support | — | Z |
| RMA / spares (5 % du parc) | — | W |
| Ops (heures × TJM interne) | — | V |
| Extension prévue (année 3 ?) | — | E |

`TCO = X + 5×(Y+Z+W+V) + E`. Compare les architectures sur ce
tableau, pas sur le CAPEX. Le NVMe perd sur le CAPEX et gagne sur
les ops (moins de baies, moins de pannes mécaniques) : le TCO
arbitre honnêtement.

## 293. L'appel d'offres : les clauses qui comptent

- Firmware **IT** imposé (HBA), UBM pour les baies tri-mode.
- Garantie 5 ans disques, avec **délai de remplacement écrit**
  (J+1 ? J+5 ?).
- SLA pièces : stock local ou expédition 48 h ?
- Droit au firmware sans contrat de support (certains OEM
  verrouillent les téléchargements).
- Pénalités de retard (allocation tendue 2026 — section 213).
- **Deux fournisseurs qualifiés** minimum par famille de disques.

## 294. Le double sourcing

