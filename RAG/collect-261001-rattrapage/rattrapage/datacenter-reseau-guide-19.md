---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-19
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["cpo", "ethernet", "intel", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [2675, 2851]
sha256: d20b1c6e132c495897cb8b261dc796c725dfba83c9e7a4dbe89fae6ed7293cb3
---

# RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE

## 138. Annexe — commandes de diagnostic (pense-bête)

**NIC NVIDIA (Linux)** :
```
ethtool -m ethX            # DOM : températures, puissances Tx/Rx
ethtool -S ethX | grep -i -E 'fec|cnp|retrans|pause'   # compteurs clés
mlxlink -d /dev/mst/mt4125_pciconf0   # état du lien, BER, FEC
lspci -vvv -s <bdf> | grep -E 'LnkSta|LnkCap'  # PCIe : génération × lanes
ib_write_bw -d mlx5_0      # bande passante RDMA
ib_write_lat -d mlx5_0     # latence RDMA
```

**NIC Intel** :
```
ethtool -k ethX            # offloads on/off
ethtool -S ethX | grep -v ': 0'   # compteurs non nuls
```

**Arista EOS** :
```
show interfaces transceiver dom   # DOM optiques
show interfaces counters errors   # erreurs/FCS
show priority-flow-control        # PFC par file
show qos interfaces               # ECN, files
show ip bgp summary               # underlay
bash tcpdump -i ethX              # capture (shell Linux)
```

**Cumulus / SONiC** :
```
nv show interface                 # Cumulus : état, DOM, compteurs
sonic-clear counters ; show interfaces counters  # SONiC
```

**Test rapide « c'est le réseau ou pas »** :
1. `ib_write_bw` entre 2 nœuds : si < 90 % → réseau.
2. `ethtool -S` : si pre-FEC BER > 1e-9 → fibre/optique.
3. Compteurs ECN/CNP : si storm → congestion/CC.
4. `NCCL_DEBUG=INFO` : si mauvaise interface → config.

## 139. Annexe — table maître débits / formats / fibres / distances

| Débit | Format NIC/switch | Fibre intra-DC | Distance | Inter-DC |
|---|---|---|---|---|
| 25G | SFP28 | DAC / OM4 SR | 5 m / 100 m | — |
| 100G | QSFP28 | DAC / AOC / OM4 SR4 / OS2 DR | 5 m / 100 m / 500 m | FR1 2 km, LR4 10 km |
| 200G | QSFP56 | DAC / AOC / OM4 SR4 | 3 m / 100 m | — |
| 400G | QSFP112/DD, OSFP | DAC / AOC / OM4 SR8 / OS2 DR4 | 3 m / 100 m / 500 m | FR4 2 km, LR4 10 km, ZR 80 km |
| 800G | OSFP, QSFP-DD800 | DAC / AEC / AOC / OM4 SR8 / OS2 DR8 | 2 m / 6 m / 30 m / 100 m / 500 m | 2×FR4 2 km |
| 1,6T | (OSFP-XD) | (AOC / OS2) | (⏳ 2027+) | — |

## 140. Index — retrouver vite

| Besoin | Sections |
|---|---|
| Choisir une NIC | §2, §4, §6, §7, §11, fiches §126-129, §133 |
| Valider le PCIe | §2, §8 |
| Comprendre RoCEv2 / UEC | §10, §45, §46, §54 |
| Choisir fibre / optique | §16-23, §28, §139 |
| Nettoyer / diagnostiquer la fibre | §24, §25, §27, §138 |
| Dimensionner un spine-leaf | §32, §33, §34, §35, §53 |
| Topologie IA (rail-optimized) | §39, §40 |
| Choisir un switch IA | §43-52, fiches §130-132 |
| InfiniBand vs Ethernet | §47, §48 |
| DPU : en mettre ou pas | §64-72, fiche §129 |
| Dimensionner un CDN/PoP | §73-80 |
| Adressage / câblage / étiquetage | §81-88 |
| Chiffrer l'énergie | §89-98, §112 |
| Futur (1,6T, UEC, CPO) | §99-103 |
| Glossaire / quiz / pièges | §105, §121, §122 |
| BOMs / checklist / sources | §123, §124, §125 |
| Dimensionnements pas-à-pas | §134-137 |

# PARTIE N — ANNEXES TECHNIQUES DENSES

## 141. Catalogue transceivers — 25 références types

Valeurs usuelles IEEE/MSA (⚠️ vérifier la datasheet du module exact) :

| Référence type | Débit | Format | Fibre | Connecteur | Distance | Conso (⚠️) |
|---|---|---|---|---|---|---|
| 10GBASE-SR | 10G | SFP+ | OM3/OM4 | LC | 300/400 m | 1 W |
| 10GBASE-LR | 10G | SFP+ | OS2 | LC | 10 km | 1 W |
| 25GBASE-SR | 25G | SFP28 | OM3/OM4 | LC | 70/100 m | 1 W |
| 25GBASE-LR | 25G | SFP28 | OS2 | LC | 10 km | 1,5 W |
| 100GBASE-SR4 | 100G | QSFP28 | OM4 (8 brins) | MPO-12 | 100 m | 2,5 W |
| 100GBASE-PSM4 | 100G | QSFP28 | OS2 (8 brins) | MPO-12 | 500 m | 3,5 W |
| 100GBASE-DR | 100G | QSFP28 | OS2 (2 brins) | LC | 500 m | 3,5 W |
| 100GBASE-FR1 | 100G | QSFP28 | OS2 | LC | 2 km | 4 W |
| 100GBASE-LR4 | 100G | QSFP28 | OS2 | LC | 10 km | 4,5 W |
| 100G-BiDi (SWDM4) | 100G | QSFP28 | OM5 (2 brins) | LC | 150 m | 3,5 W |
| 200GBASE-SR4 | 200G | QSFP56 | OM4 | MPO-12 | 100 m | 4 W |
| 200GBASE-DR4 | 200G | QSFP56 | OS2 (8 brins) | MPO-12 | 500 m | 5 W |
| 400GBASE-SR8 | 400G | QSFP-DD/OSFP | OM4 (16 brins) | MPO-16 | 100 m | 9 W |
| 400GBASE-DR4 | 400G | QSFP-DD/OSFP | OS2 (8 brins) | MPO-12 | 500 m | 9 W |
| 400GBASE-FR4 | 400G | QSFP-DD/OSFP | OS2 (2 brins) | LC | 2 km | 10 W |
| 400GBASE-LR4 | 400G | QSFP-DD/OSFP | OS2 | LC | 10 km | 12 W |
| 400GBASE-ZR | 400G | QSFP-DD | OS2 | LC | 80 km | 15 W |
| 800GBASE-SR8 | 800G | OSFP | OM4 (16 brins) | MPO-16 | 100 m | 13 W |
| 800GBASE-DR8 | 800G | OSFP | OS2 (16 brins) | MPO-16 | 500 m | 15 W |
| 800GBASE-2×FR4 | 800G | OSFP | OS2 (4 brins) | 2×CS | 2 km | 16 W |
| 800GBASE-LR8 | 800G | OSFP | OS2 | LC | 10 km | 18 W |
| DAC 100G 3 m | 100G | QSFP28 | Cuivre | — | 3 m | 0,3 W |
| DAC 400G 2 m | 400G | QSFP-DD | Cuivre | — | 2 m | 0,5 W |
| AOC 400G 30 m | 400G | QSFP-DD | Fibre active | — | 30 m | 2 W |
| AOC 800G 30 m | 800G | OSFP | Fibre active | — | 30 m | 3 W |

**Lecture** : à 800G, l'écart DAC (0,5 W) vs DR8 (15 W) = **×30**. Chaque
lien 800G qui peut rester en cuivre fait économiser ~30 W (les 2 bouts).

## 142. Normes IEEE 802.3 — repères par génération

| Norme | Année | Débit | Modulation | Média |
|---|---|---|---|---|
| 802.3ae | 2002 | 10G | NRZ | Fibre |
| 802.3ba | 2010 | 40/100G | NRZ | Fibre/Cuivre |
| 802.3bm | 2015 | 100G | NRZ | — |
| 802.3bs | 2017 | 200/400G | PAM4 | Fibre |
| 802.3cd | 2018 | 50/100/200G | PAM4 | — |
| 802.3cu | 2021 | 100G monomode | PAM4 | OS2 |
| 802.3db | 2022 | 100/200/400G SR | PAM4 | OM4 |
| 802.3ck | 2022 | 100/200/400G élec. | PAM4 100G/lane | Backplane |
| 802.3df | 2024 | 800G | PAM4 100G/lane | Fibre |
| 802.3dj (⏳) | 2026+ | 1,6T | PAM4 200G/lane | En cours |

**À retenir** : chaque génération double le débit en doublant les lanes
ou le baud rate, en gardant la compatibilité descendante des formats.

## 143. Arbre de dépannage — « le réseau est lent »

```
  LE RÉSEAU EST LENT
  │
  ├─ Tous les nœuds ? ── OUI ──→ Fabric : ECMP ? (§36) / CC ? (§46) / spine down ?
  │                        NON ──→ Ce nœud : voir ci-dessous
  │
  ├─ ib_write_bw < 90 % ? ── OUI ──→ Lien/NIC : ethtool -S (BER ? FEC ?)
  │   └─ BER > 1e-9 ? ── OUI ──→ Nettoyer fibre (§24) → budget optique (§20)
  │                        NON ──→ PCIe ? (lspci : Gen × lanes, §8)
  │
  ├─ NCCL lent seul ? ── OUI ──→ NCCL_SOCKET_IFNAME (§108) : bonne interface ?
  │   └─ busbw < 70 % ? ── OUI ──→ ECMP/CC sur le chemin (§107)
  │
  ├─ Pauses PFC ? ── OUI ──→ Deadlock : quelle file ? Quel émetteur ? (§45)
  │
  ├─ CNP storm ? ── OUI ──→ DCQCN mal réglé ou incast réel (§46)
  │
  └─ Aléatoire ? ──→ FW hétérogènes (§13) / ECMP polarisé (§36) / thermique (§60)
```

**Règle** : toujours mesurer **avant** de changer quoi que ce soit, et ne
changer **qu'une chose à la fois**. Le debug « on a tout redémarré et ça
remarche » n'apprend rien et ça reviendra.

## 144. Templates — plan d'adressage, BOM, test de réception

**Template plan d'adressage** (à remplir par site) :
```
Site : __________    Date : __________    Responsable : __________
Loopbacks      : 10.__.0.0/24   (équipement : __________)
Liens /31      : 10.__.1.0/__   (____ liens prévus, ____ de réserve)
Management     : 192.168.__.0/24
ASN spines     : 651__
ASN leafs      : 650__
VTEP           : = loopbacks
IPv6 loopbacks : ____::/64  (prévu : oui/non)
NetBox         : URL __________
```

**Template BOM** (une ligne par poste) :
```
| Poste | Référence | Qté | PU HT | Total HT | Fournisseur | Délai |
|-------|-----------|-----|-------|----------|-------------|-------|
|       |           |     |       |          |             |       |
Réserve spares : ____%    Total HT : ______    Total TTC : ______
```

