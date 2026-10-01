---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-25
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: ["Google", "Nvidia", "xAI"]
dates: []
keywords: ["blackwell", "capex", "distribution", "hyperscaler", "nvidia", "rubin", "training"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [3693, 3852]
sha256: 7957b610d78e5e7b518e8a99fdad921effa770700ac5037acc027fd70ff4c82e
---

# RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE

| | 25G | 100G | 200G | 400G | 800G | 1,6T |
|---|---|---|---|---|---|---|
| Format | SFP28 | QSFP28 | QSFP56 | QSFP112/DD | OSFP | OSFP-XD ⏳ |
| Modulation | NRZ | NRZ | PAM4 | PAM4 | PAM4 | PAM4 |
| Lanes | 1×25G | 4×25G | 4×50G | 4×100G | 8×100G | 8×200G |
| PCIe | Gen3 x8 | Gen4 x16 | Gen4 x16 | Gen5 x16 | Gen6 x16 | Gen7 x16 ⏳ |
| NIC réf. | XXV710 | E810 ✅ | CX-6 Dx | CX-7 / Thor 2 ✅ | CX-8 / Thor Ultra ✅ | CX-9 ⏳ |
| Switch 64p | — | 6,4T | 12,8T | 25,6T | 51,2T | 102,4T |
| Fibre intra | DAC/OM4 | DAC/SR4 | DAC/SR4 | DAC/SR8 | DAC/SR8 | ⏳ |
| Conso NIC (⚠️) | 10 W | 15-25 W | 20-30 W | 30-40 W | 50-75 W | ? |
| W/Gb/s (⚠️) | 0,4 | 0,2 | 0,15 | 0,1 | 0,08 | 0,06 ? |
| Usage 2026 | OOB | Front-end | Transition | Backend Hopper | Backend Blackwell | Rubin 2027+ |
| Statut | Legacy | Standard | Palier | Standard IA | Nouveau | Annoncé |

## 176. Contrôles de congestion — comparatif

| Algorithme | Principe | Où | Note |
|---|---|---|---|
| DCQCN | ECN + CNP, émetteur ralentit | RoCEv2 | Le standard, à régler |
| TIMELY | RTT mesuré, ralentit si ↑ | RoCEv2 (Google) | Précis, sensible au bruit |
| HPCC | Télémétrie in-band (INT) | Recherche | Optimal mais complexe |
| Swift | Délai + ECN | Google | Production hyperscaler |
| CC Spectrum-X | Adaptatif + télémétrie HW | NVIDIA fermé | 95 % (cas xAI) ✅ |
| CC UEC | Émetteur/récepteur programmable | Thor Ultra ✅ | Le futur ouvert |

**Règle** : ne pas mélanger deux CC sur le même fabric (ils se battent).
Un seul algorithme, bien réglé, validé au lab (§137).

## 177. Connecteurs et câbles — repères physiques

| Connecteur | Type | Usage |
|---|---|---|
| RJ45 | Cuivre 8P8C | 1G/10G, OOB, console |
| SFP+/28 | 1 lane | 10/25G |
| QSFP+/28/56 | 4 lanes | 40/100/200G |
| QSFP-DD/112 | 8 lanes | 400G |
| OSFP | 8 lanes, grand | 400/800G (thermique) |
| QSFP-DD800 | 8 lanes | 800G (format DD) |
| LC duplex | Fibre 2 brins | SR/DR/FR/LR |
| MPO-12/16/24 | Fibre multi-brins | SR4/SR8/DR4/DR8, trunks |
| CS / SN | Fibre compacte | 800G dense |

## 178. Template — plan de câblage (par baie)

```
Baie : __________    Date : __________    Technicien : __________
| U | Équipement | Port | Destination (baie/éq/port) | Câble | Long. | Étiquette |
|---|------------|------|------------------------------|-------|-------|-----------|
|42 | LF04       | e49  | S1-R03-B12-SRV07-ETH1         | DAC400| 2 m   | S1-R03... |
|...|            |      |                              |       |       |           |
Vérifié : [ ] polarité MPO  [ ] longueurs ≤ portée  [ ] mou 30 cm
          [ ] étiquettes 2 bouts  [ ] photo
```

## 179. Alimentations — repères par switch

| Switch | PSU | Config | Note |
|---|---|---|---|
| Leaf 1U 100G | 2× 500-800 W | 1+1 | — |
| Leaf 1U 400G | 2× 1000-1500 W | 1+1 | — |
| Leaf 2U 800G | 2× 2400 W | 1+1 | Arista 7060X6 ✅ |
| SN5600 | 2× (1+1) | 1+1 | ✅ |
| SN5610 | 4× (2+2) | 2+2 | ✅ |
| Châssis 7800R4 | N+1 par zone | N+1 | Jusqu'à 25 kW (⚠️) |

**Règle** : PDU A/B **indépendantes**, chacune = 100 % de la charge.
Disjoncteur ≥ 125 % du max mesuré au boot.

## 180. FAQ 2 — 20 questions courtes

**1.** Spectrum-X ou UEC ? → Fermé performant vs ouvert : selon l'équipe.
**2.** Un seul vendeur, c'est grave ? → Non si le contrat prévoit la sortie.
**3.** SONiC en prod IA ? → Distribution durcie seulement.
**4.** Cumulus ou SONiC sur NVIDIA ? → Cumulus (intégré Spectrum-X).
**5.** EVPN : eBGP ou iBGP ? → eBGP underlay, iBGP overlay (RR sur spines).
**6.** Combien de VTEP par fabric ? → Tous les leafs (anycast gateway).
**7.** BFD : quel timer ? → 100 ms × 3 (300 ms).
**8.** Jumbo frames : partout ? → Oui, MTU 9000 de bout en bout.
**9.** VXLAN sans EVPN ? → Possible (flood-and-learn) mais ne pas faire en grand.
**10.** PFC : sur quelles files ? → La seule file RoCE, jamais en global.
**11.** ECN : seuils ? → Marquage 30-50 %, pause 70-80 % (⚠️ à tuner).
**12.** Buffer : combien ? → BDP × flux (voir §44).
**13.** Cut-through ou store-and-forward ? → Cut-through (latence).
**14.** Latence leaf-spine-leaf ? → ~2-3 µs (cut-through 800G).
**15.** Gigue acceptable ? → < 1 µs pour NCCL synchrone (⚠️).
**16.** Câble : longueur max en baie ? → 3 m (DAC), au-delà AOC.
**17.** OS2 ou OM4 entre 2 rangées (60 m) ? → OM4 SR (moins cher), sauf si 800G+ → OS2.
**18.** Un transceiver peut-il servir 10 ans ? → Non : 5-7 ans, cycle avec le switch.
**19.** Revendre les vieux switchs ? → Oui (marché secondaire), après wipe config.
**20.** Le conseil ultime ? → §104, règle 1 : séparer les fabrics.

## 181. Fin de vie — décommission propre

- [ ] Configs sauvegardées (preuve), puis **wipe** (données, certificats)
- [ ] Optiques retirées, bouchons remis, stockées anti-statique
- [ ] Inventaire NetBox à jour (statut « retired »)
- [ ] Câbles : réutilisables ? (DAC oui, fibre si propre) sinon recyclage
- [ ] Revente : effacer les licences, transférer la garantie si possible
- [ ] DEEE : switchs = déchets électroniques (traçabilité)

---

*Fin du guide — **181 sections numérotées**, 80 termes de glossaire,
45 pièges, quiz, 50 FAQ, 100 acronymes, BOMs, checklists, templates.
Vérification finale : `wc -l` — objectif ≥ 4000 lignes.*

# PARTIE R — CLÔTURE OPÉRATIONNELLE

## 182. Matrice de décision — quel fabric pour quel workload

| Workload | Endpoints | Ratio | Débit | Topologie | NOS | Budget réseau (⚠️) |
|---|---|---|---|---|---|---|
| Bureautique / web | < 500 | 5:1 | 25G | 2 tiers | Au choix | 100-300 k€ |
| Virtualisation | 500-2000 | 3:1 | 25/100G | 2 tiers + EVPN | EOS/NX-OS | 0,5-2 M€ |
| Cloud privé | 2000+ | 3:1 | 100G | 2 tiers + EVPN | EOS/SONiC | 2-5 M€ |
| Stockage NVMe-oF | 100-500 | 2:1 | 100/200G | 2 tiers dédié | Au choix | 0,5-1,5 M€ |
| HPC classique | 100-1000 | 1:1 | 200/400G | 2 tiers / IB | Cumulus/EOS | 1-5 M€ |
| IA training (Hopper) | 256-2000 | 1:1 | 400G | Rail-opt. 2 tiers | EOS/Cumulus | 5-20 M€ |
| IA training (Blackwell) | 1000+ | 1:1 | 800G | Rail-opt. multi-plane | EOS/Cumulus | 20-90 M€ |
| Inférence | 50-500 | 2:1 | 100/400G | 2 tiers | Au choix | 1-5 M€ |
| Edge / CDN | 2-20/PoP | — | 100G | Simple | Au choix | 70-120 k€/PoP |

**Lecture** : il n'y a pas « un » bon fabric — il y a le bon fabric pour
le workload. Le surdimensionnement coûte du CAPEX, le sous-dimensionnement
coûte des semaines de production.

## 183. Extraits de configuration commentés

**Arista EOS — underlay eBGP + BFD** :
```
router bgp 65001
   router-id 10.0.0.1
   maximum-paths 64
   neighbor SPINE peer-group
   neighbor SPINE remote-as 65101
   neighbor SPINE bfd          ! détection 300 ms
   neighbor 10.0.1.0 peer-group SPINE
   !
   address-family ipv4
      network 10.0.0.1/32      ! loopback = VTEP
```

**Cumulus — ECN sur file RoCE** :
```
# /etc/cumulus/datapath/traffic.conf (extrait conceptuel)
ecn -q 3 -min 30 -max 80       ! marquage ECN file 3 (RoCE)
pfc -q 3 -xoff 75 -xon 60      ! PFC dernier recours
```

**Linux — offloads et MTU** :
```
ip link set eth0 mtu 9000
ethtool -K eth0 tx-checksum-ipv4 on tso on gro on
tc qdisc add dev eth0 root fq_codel ecn   ! ECN côté hôte
```

⚠️ Extraits pédagogiques : adapter au NOS et à la version exacte.

## 184. Les 10 erreurs humaines les plus chères (chiffrées, ⚠️)

