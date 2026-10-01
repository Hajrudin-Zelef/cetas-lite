---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-5
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: ["Broadcom"]
dates: ["2026-09-27"]
keywords: ["datacenter", "dsp", "gpu", "lpo", "optics", "serdes", "training"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [551, 710]
sha256: 032d057e3bb46ec575186605a7061bc5fc78b9cb4bcf3826ca63a0d6d2bfd65f
---

# RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE

Le PAM4 (4 niveaux) a un rapport signal/bruit ~9,6 dB moins bon que le NRZ.
Sans FEC, le 400/800G ne fonctionnerait pas. Points clés :

| Paramètre | Valeur usuelle (⚠️) |
|---|---|
| FEC 400/800G | RS(544,514) — Reed-Solomon |
| Seuil pré-FEC tolérable | ~1e-4 à 2,4e-4 BER |
| Post-FEC visé | < 1e-12 (souvent < 1e-15) |
| Latence FEC | ~100 ns par traversée |

**Supervision** : exposer les compteurs `pre-FEC BER` / `FEC corrected` en
SNMP/telemetry. Seuils d'alerte :
- 🟡 BER pré-FEC > 1e-6 : surveiller, planifier nettoyage.
- 🔴 BER pré-FEC > 1e-5 : intervenir (nettoyage, puis remplacement).
- ⛔ Erreurs **non corrigées** > 0 : le lien est en échec, basculer.

**Piège** : le FEC masque une fibre dégradée pendant des mois, puis le lien
s'effondre d'un coup quand on franchit le seuil. D'où la supervision
préventive — un lien « qui marche » n'est pas un lien « sain ».

## 28. LPO / LRO : les optiques linéaires basse consommation

- **LPO (Linear Pluggable Optics)** : supprime le DSP du module, le traitement
  du signal est fait par le SerDes du switch/NIC. Gain : **−30 à 50 % de
  conso** et −latence (~dizaines de ns).
- **LRO (Linear Receive Optics)** : moitié linéaire (côté réception).
- ✅ Broadcom pousse les LPO sur Thor 2 (vérifié 27/09/2026) ; écosystème en
  croissance pour le 800G intra-DC.
- **Contrainte** : LPO exige des SerDes compatibles des deux côtés
  (interopérabilité limitée en 2026) et des liens courts/propres. **Ne pas
  mélanger LPO et DSP sur un même lien.**
- **Lecture énergie** : à l'échelle d'un fabric 800G (10 000 modules),
  passer de 15 W à 8 W par module = **70 kW économisés** — un levier PUE
  direct (voir §97).

## 29. CAS CHIFFRÉ — BOM fibre d'un rack 400G (8 serveurs)

Hypothèses : 8 serveurs × 2×400G (rouge/bleu), leaf en tête de rack,
liens < 3 m.

| Poste | Qté | Choix | Prix unit. (⚠️) | Total (⚠️) |
|---|---|---|---|---|
| DAC 400G QSFP-DD 2 m | 16 | Passif, codé switch | 150-350 € | 2 400-5 600 € |
| Uplinks leaf→spine 400G DR4 | 4 | Transceiver + 50 m OS2 | 1 500-4 000 € | 6 000-16 000 € |
| Jarretières OS2 LC-LC 50 m | 4 | — | 40-80 € | 160-320 € |
| Kit nettoyage + microscope | 1 | Mutualisé salle | 1 500-3 000 € | 1 500-3 000 € |
| Étiquettes + platine brassage | 1 | TIA-606 (§85) | 300-600 € | 300-600 € |
| **Total rack** | | | | **~10-25 k€** |

**Enseignement** : le câblage cuivre intra-rack coûte 10× moins cher que
l'optique — **rapprocher les leafs des serveurs** (top-of-rack) est une
décision économique autant que technique. Chaque mètre de fibre en plus =
optiques + brassage + points de panne.

## 30. Pièges fibre — l'essentiel (détail §107)

- **P6** — OSFP flat-top sur switch (au lieu de finned-top) : surchauffe.
- **P7** — Jarretière jaune (OS2) sur port SR (multimode) : lien muet.
- **P8** — Polarité MPO non documentée : heures de debug nocturne.
- **P9** — DAC d'un côté, transceiver de l'autre : pas de link.
- **P10** — Breakout non supporté par le port : ports morts après achat.
- **P11** — FEC qui masque une fibre sale : panne brutale 6 mois plus tard.

# PARTIE C — ARCHITECTURES SPINE-LEAF

## 31. Principe : le Clos à deux tiers

Le spine-leaf est un **réseau de Clos** : chaque leaf est connecté à **tous**
les spines, aucun leaf n'est connecté à un autre leaf, aucun serveur n'est
branché sur un spine.

Propriétés fondamentales :
- **Tout trajet serveur→serveur = 3 sauts** (leaf→spine→leaf) : latence
  prévisible, ~1-2 µs par saut en cut-through.
- **Bande passante horizontale** (est-ouest) : c'est le trafic dominant du
  datacenter moderne (VM↔VM, GPU↔GPU, stockage). Le vieux modèle 3 tiers
  (accès/agrégation/cœur) était pensé pour le nord-sud.
- **Évolutivité** : on ajoute des leafs (capacité serveurs) ou des spines
  (bande passante) indépendamment.

```
        TRAFIC NORD-SUD (vers WAN/Internet)
                    ▲  │  ▲
                    │  ▼  │
              ┌───────────────┐
              │   Routeurs    │  ← bordure (peering, firewall)
              │   de bordure  │
              └──────┬───┬────┘
                     │   │
   ═════════════ TIER SPINE ═════════════
        ┌────────┴───┴────────┐
   ┌────┴────┐          ┌────┴────┐
   │ Spine 1 │          │ Spine 2 │   ← jamais reliés entre eux
   └──┬───┬──┘          └──┬───┬──┘
      │   │    ╲     ╱    │   │
      │   │     ╲ ╱       │   │      ← chaque leaf vers TOUS les spines
      │   │     ╱ ╲       │   │
   ┌──┴───┴──┐ ┌──┴───┴──┐ ┌──┴───┴──┐
   │ Leaf 1  │ │ Leaf 2  │ │ Leaf 3  │  ← ToR (top-of-rack)
   └──┬───┬──┘ └──┬───┬──┘ └──┬───┬──┘
     ▓▓▓ ▓▓▓     ▓▓▓ ▓▓▓     ▓▓▓ ▓▓▓
     serveurs    serveurs    serveurs
```

## 32. Oversubscription : le ratio qui dimensionne tout

**Oversubscription = capacité downlink (vers serveurs) ÷ capacité uplink
(vers spines).**

| Ratio | Signification | Usage |
|---|---|---|
| 1:1 | Non-bloquant | Backend IA (GPU), HPC |
| 3:1 | Standard | Virtualisation généraliste |
| 5:1 à 10:1 | Économique | Bureautique, web frontal |

**Exemple de calcul** — leaf 48×25G serveurs + 8×100G uplinks :
- Downlink : 48 × 25 = 1 200 Gb/s
- Uplink : 8 × 100 = 800 Gb/s
- **Ratio = 1 200 / 800 = 1,5:1**

**Exemple classique 3:1** — leaf 48×10G + 4×40G :
- 480 / 160 = **3:1** ✅

**Règle IA** : le backend GPU exige **1:1 non-bloquant** (un all-reduce
NCCL sature tous les liens en même temps). Le front-end (management,
stockage) peut être à 3:1. **Ne jamais mélanger les deux trafics sur le
même fabric** : le burst d'un training noie le stockage (P12, §107).

## 33. Schéma ASCII : 2 tiers, 1024 serveurs non-bloquants

Dimensionnement Clos non-bloquant : avec des switchs de **n** ports,
on met n/2 ports vers le bas et n/2 vers le haut par leaf.
Capacité : **(n/2) leafs × (n/2) serveurs = n²/4 serveurs**, avec n/2 spines.

Exemple n = 64 (switch 64×100G) → 32×32 = **1024 serveurs**, 32 spines :

```
                    32 SPINES (64×100G)
              ┌───┬───┬───┬───┬───┬───┐
              │ S1│ S2│...│...│S31│S32│
              └─┬─┴─┬─┴───┴───┴─┴─┬─┘
                │   │             │
      ┌─────────┘   │     ┌───────┴─────────┐
      │  32 uplinks │     │  32 uplinks     │   ← 32×100G = 3,2 Tb/s par leaf
 ┌────┴─────┐  ┌────┴─────┐        ┌────┴─────┐
 │ Leaf 1   │  │ Leaf 2   │  ...   │ Leaf 32  │  ← 32 leafs
 └────┬─────┘  └────┬─────┘        └────┬─────┘
   32×100G       32×100G              32×100G   ← vers serveurs
   1024 serveurs au total, 1:1 non-bloquant
```

**Vérification** : chaque leaf a 32×100G = 3,2 Tb/s vers le haut et
3,2 Tb/s vers le bas → 1:1 ✅. Chaque spine a 32×100G = 3,2 Tb/s, tous
vers des leafs distincts ✅.

**Coût** : 64 switchs 64×100G + 2048 optiques/câbles. C'est le prix du
non-bloquant : ~2× plus de ports switch que de ports serveurs.

## 34. Schéma ASCII : 3 tiers (pod + super-spine) pour l'IA

