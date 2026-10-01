---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-6
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: ["Nvidia", "OpenAI"]
dates: ["2026-09-27"]
keywords: ["arr", "cost", "gpu", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [711, 869]
sha256: 437feac227fe6f6c5b94445b044589100184d4aa7537bb5256a4e1a31fb8c7df
---

# RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE

Au-delà de ~1-2k GPU, on passe en 3 tiers : des **pods** (2 tiers) reliés
par des **super-spines**. Les pods = unités de déploiement (ex. 256 GPU).

```
              ┌─────────────────────────────┐
              │      SUPER-SPINES (×8)       │  ← 800G, Jericho3-AI / TH5
              │  64×800G, deep buffer       │
              └──────┬──────────────┬───────┘
                     │              │
        ┌────────────┴───┐   ┌──────┴────────┐
        │    POD 1       │   │    POD 2      │  ← 256 GPU chacun
        │  (rail-opt.)   │   │  (rail-opt.)  │
        └────────────────┘   └───────────────┘

  Détail d'un POD (rail-optimized, voir §39) :

        Spine P1    Spine P2          ← 1 spine par "rail" (plan)
         │  │        │  │
    ┌────┘  │        │  └────┐
    │  ┌────┘        └────┐  │
 ┌──┴──┴──┐           ┌──┴──┴──┐
 │ Leaf R1 │           │ Leaf R2 │  ← leafs dédiés au rail 1 / rail 2
 └──┬───┬──┘           └──┬───┬──┘
  GPU0 GPU1            GPU2 GPU3   ← NIC#0 des GPU → rail 1, NIC#1 → rail 2
```

**Règle de pouce** : 2 tiers jusqu'à ~2 000 endpoints, 3 tiers au-delà.
Le 3ᵉ tier ajoute 1 saut (~1 µs) mais divise le nombre de ports spine par pod.

## 35. Dimensionnement : exemple 512 serveurs généralistes (3:1)

| Paramètre | Valeur |
|---|---|
| Serveurs | 512 × 2×25G (rouge/bleu) |
| Leaf | 48×25G + 8×100G → 24 ports serveurs/leaf en 3:1 ? |

Calcul : 48×25G = 1,2 Tb/s down. Pour 3:1 → uplink = 400 Gb/s = 4×100G.
Ports serveurs utilisables : 48 (le ratio se joue sur les uplinks).
- Leafs : 512 / 48 ≈ **11 leafs** (prendre 12, pair).
- Uplinks par leaf : 4×100G → 48 uplinks au total.
- Spines : switch 32×100G → **2 spines** (48 ports utilisés).
- Ratio vérifié : (48×25)/(4×100) = 1200/400 = **3:1** ✅

**BOM** : 12 leafs + 2 spines + 1024 DAC 25G + 48 optiques 100G.
(⚠️ : prévoir 20 % de ports libres pour croissance — P13, §107.)

## 36. ECMP : le routage qui répartit

En spine-leaf, **pas de STP** : tous les liens sont actifs grâce à ECMP
(Equal-Cost Multi-Path). Chaque leaf installe N routes vers chaque
destination (N = nombre de spines) et **hashe** les flux (5-tuple
IP src/dst + ports) sur les liens.

Points critiques :
- **Polarisation ECMP** : si deux équipements hashent pareil, un lien se
  sature pendant que d'autres sont vides. Activer le **hash seed aléatoire**
  par switch (option sur Arista EOS, Cisco NX-OS, Cumulus).
- **Flots éléphants** (IA) : un seul flux NCCL de 400G ne se répartit pas —
  d'où le **packet spraying** (Spectrum-X) ou l'**adaptive routing** : on
  répartit à la granularité paquet, pas flux.
- **BGP PIC / fast-reroute** : prévoir la convergence < 1 s en cas de perte
  d'un spine (BFD, voir §38).

## 37. EVPN/VXLAN en bref : l'overlay qui découple

**Problème** : les VM/conteneurs migrent, les VLAN sont limités à 4094, le
spanning-tree ne passe pas l'échelle.
**Solution** : VXLAN (encapsulation L2 dans UDP, 24 bits = 16M de segments)
+ EVPN (BGP qui distribue les MAC/IP comme des routes).

```
  VM-A (VNI 5001)                                    VM-B (VNI 5001)
      │                                                  │
  ┌───┴────┐                                        ┌───┴────┐
  │ Leaf 1 │═══ VXLAN tunnel (VNI 5001) ════════════│ Leaf 2 │  ← VTEP
  └───┬────┘   (sous-réseau IP underlay)            └───┬────┘
      │                routé en ECMP                     │
   ───┴────────────── SPINE ────────────────────────────┴───
```

- **Underlay** : BGP numéroté (souvent /31 point-à-point, RFC 5549), eBGP
  entre leaf et spine avec ASN privés (ex. leafs 65001, spines 65101).
- **Overlay** : iBGP EVPN entre leafs (VTEP), route-reflectors sur les spines.
- **Anycast gateway** : la même IP/MAC de passerelle sur tous les leafs —
  la VM garde sa gateway en migrant.

**Quand s'en passer** : backend IA pur (pas de VM, pas de migration) →
routage simple, pas de VXLAN : moins d'overhead, moins de MTU à gérer
(VXLAN = +50 octets → MTU 9000+50, P14, §107).

## 38. BGP underlay : le plan de routage qui ne tombe jamais

Config type (eBGP, ASN privés) :

```
  Leaf 1 : ASN 65001                    Spine 1 : ASN 65101
  interfaces /31 : 10.0.0.0/31, 10.0.0.2/31...
  Spine 2 : ASN 65102 (ASN différent par spine = ECMP naturel)
```

- **unnumbered BGP** (RFC 5549) : possible avec IPv6 link-local — simplifie
  le plan d'adressage, supporté par EOS/Cumulus/NX-OS.
- **BFD** : détection de panne en ~300 ms (3×100 ms) au lieu des 180 s du
  hold-time BGP. **Obligatoire** sur les liens leaf-spine.
- **Pas de redistribution hasardeuse** : l'underlay ne porte que les
  loopbacks (/32) et les liens ; tout le reste passe par l'overlay EVPN.
- **EBGP multipath** : `maximum-paths 64` pour installer tous les spines.

## 39. Rail-optimized : la topologie des clusters GPU

Dans un serveur 8 GPU, chaque GPU a sa NIC (rail). Le **rail-optimized**
connecte toutes les NIC#0 au même plan réseau, toutes les NIC#1 à un autre
plan, etc. — **jamais de mélange entre rails**.

```
  Serveur A (8 GPU)              Serveur B (8 GPU)
  GPU0-NIC0 ──→ PLAN 0 (leafs+spines dédiés)
  GPU1-NIC1 ──→ PLAN 1
  ...
  GPU7-NIC7 ──→ PLAN 7

  Chaque plan = un mini spine-leaf indépendant :
  ┌─────────────────────────────────────────┐
  │ PLAN k :  Spines-k ─── Leafs-k ─── NICk │  × 8 plans
  └─────────────────────────────────────────┘
```

**Pourquoi** : le trafic dominant (all-reduce en anneau NCCL) circule entre
les mêmes indices de GPU. Séparer les plans = pas de contention entre
rails, pas de head-of-line blocking inter-rails, et un plan en panne
n'arrête pas tout (dégradé, pas mort).
✅ C'est la topologie par défaut de Spectrum-X et des références NVIDIA
(vérifié 27/09/2026). **Multi-plane** = version où l'équilibrage entre
plans est accéléré en hardware (utilisé avec OpenAI, d'après la doc
d'écosystème).

## 40. Multi-plane et topologies alternatives

| Topologie | Principe | Quand |
|---|---|---|
| Rail-optimized multi-plane | 1 plan par indice NIC, LB HW inter-plans | Clusters IA ≥ 1k GPU |
| Fat-tree / Clos classique | Tout vers tout via ECMP | Généraliste, cloud |
| Dragonfly / SlimFly | Groupes denses, liens inter-groupes | HPC extrême (rare en enterprise) |
| Rail-only (sans spine) | NICk↔NICk en direct (câbles) | Petits clusters ≤ 32 nœuds, éco |

**Règle** : < 256 GPU → 2 tiers rail-optimized simple ; 256-4 000 →
multi-plane ; > 4 000 → 3 tiers + rail-optimized par pod.

## 41. CAS CHIFFRÉ — BOM spine-leaf 1024 ports 400G (backend IA)

Hypothèses : 256 serveurs GPU × 4×400G ? Non — prenons 128 nœuds × 8 NIC
400G = **1024 ports 400G**, 1:1 non-bloquant, 2 tiers, switchs 64×800G
utilisés en 2×400G (128 ports 400G par switch).

- Leafs : 1024 / 128 = **8 leafs** (64×800G → 128×400G).
  Uplinks : 64×400G par leaf → 512 uplinks.
- Spines : 512 / 128 = **4 spines** 64×800G (128×400G).
- Vérification : chaque leaf 64×400G up / 64×400G down = **1:1** ✅

