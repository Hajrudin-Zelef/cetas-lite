---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-11
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: ["Broadcom", "Intel", "Nvidia"]
dates: ["2026-09-27"]
keywords: ["cpo", "distribution", "gpu", "intel", "lpo", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [1485, 1649]
sha256: 5cc4dccc3afe6a38f6251750a8e9bd0a3d2bf95f393aba1438db8889227f929e
---

# RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE

**Règles** :
- **/31** sur les point-à-point : divise par 2 la consommation d'adresses
  et évite les erreurs de masque. Supporté partout en 2026.
- **ASN privés** : 64512-65534 (16 bits) ou 4200000000-4294967294 (32 bits).
  Convention : spines 65101/65102…, leafs 65001+ (voir §38).
- **IPv6** : prévoir le plan v6 dès maintenant (loopbacks /128, liens /127)
  même si on ne l'active pas — le renuméroter après = cauchemar.
- **Un seul fichier source de vérité** (NetBox — voir guide NetBox) :
  jamais d'IP « dans la tête de Kevin ».

## 82. Câblage : les bonnes pratiques qui évitent 90 % des pannes

1. **Séparer les chemins** : fibre optique d'un côté du chemin de câbles,
   cuivre/élec de l'autre. Jamais de fibre dans le même faisceau que du
   400 V.
2. **Longueurs justes** : 30 cm de mou, pas 3 m de spaghetti. Le surplus
   se range dans des gestionnaires verticaux, pas en boule derrière le switch.
3. **Top-of-Rack** : leafs en haut de chaque baie → DAC courts (1-2 m),
   pas de brassage horizontal pour les serveurs.
4. **Couleurs par usage** : ex. bleu = data, jaune = OS2 (norme), rouge =
   OOB, vert = stockage. **Documenter la légende** et s'y tenir.
5. **Jamais de tension** sur un connecteur : un MPO tiré = micro-courbures
   = pertes. Laisser du mou aux deux extrémités.

## 83. Longueurs et rayons de courbure — les chiffres

| Câble | Rayon de courbure min (⚠️ usuel) | Longueur max pratique |
|---|---|---|
| Fibre OM4 (G.651) | 15 mm (insensible ~7,5 mm) | 100 m (SR8) |
| Fibre OS2 (G.652D) | 30 mm (G.657 : 7,5-10 mm) | km |
| DAC 400G/800G | 50-70 mm (rigide !) | 2-3 m |
| AOC | 30 mm | 30-100 m |
| Cuivre Cat6A | 4× diamètre (~30 mm) | 100 m |

**Micro-courbures** : un rayon trop serré sur de l'OS2 = la lumière fuit
dans la gaine = **pertes de 1-3 dB invisibles à l'œil**. Symptôme typique :
lien qui monte mais BER pré-FEC élevée (§27). **Ne jamais coincer une fibre
sous un rail ou une porte de baie** (P31, §107).

## 84. Brassage : l'architecture MDA (Main Distribution Area)

```
  [Rangées serveurs] ──trunks MPO──→ ┌─────────────┐
                                     │  MDA / EDA  │  ← brassage central
  [Rangées serveurs] ──trunks MPO──→ │  panneaux   │
                                     │  MPO + LC   │
                                     └──────┬──────┘
                                            │ uplinks
                                            ▼
                                     [Spines / cœur]
```

- **Avantage** : tout brassage au même endroit, MAC (moves/adds/changes)
  sans toucher aux rangées, documentation unique.
- **Inconvénient** : 2 connecteurs de plus par lien (donc +0,6-1 dB, §20) —
  à intégrer dans le budget optique.
- **Règle** : tout lien > 30 m passe par le brassage ; les DAC intra-rack
  restent en direct (pas de brassage pour 2 m de cuivre).

## 85. Étiquetage : norme TIA-606, sans exception

Format recommandé : `SALLE-RANGÉE-BAIE-ÉQUIPEMENT-PORT` aux deux bouts,
**identique des deux côtés** (pas « A » d'un côté et « B » de l'autre).

Exemple : `S1-R03-B12-LF04-ETH49` = Salle 1, rangée 3, baie 12, leaf 4,
port 49. L'autre bout : `S1-R03-B12-SRV07-ETH1`.

| Support | Usage |
|---|---|
| Étiquettes auto-laminées | Câbles (tiennent 10 ans) |
| Manchons thermorétractables | Jarretières fibre (propre) |
| Plaques gravées | Baies, panneaux de brassage |
| QR code → NetBox | Équipements (inventaire) |

**Règle d'or** : un câble non étiqueté = un câble à retirer à la prochaine
intervention. Le coût d'une étiquette (~0,50 €) vs le coût d'une heure de
debug (150 €) : le calcul est vite fait.

## 86. Gestion des câbles : capacité des chemins

| Chemin | Règle de remplissage (⚠️ usuel) |
|---|---|
| Chemin de câbles (goulotte) | **≤ 50 %** de la section (réserve pour croissance) |
| Faisceau vertical | Attaches velcro, jamais de colliers nylon serrés sur fibre |
| Traversées coupe-feu | Manchons coupe-feu, **reboucher après chaque passage** |

**Pire ennemi** : le « juste un câble de plus » répété 200 fois = goulotte
à 120 %, fibres écrasées en bas du faisceau, pertes diffuses sur des
dizaines de liens. **Auditer le remplissage une fois par an.**

## 87. CAS CHIFFRÉ — plan d'adressage d'un pod 256 GPU

| Usage | Plage | Détail |
|---|---|---|
| Loopbacks leafs+spines | 10.100.0.0/24 | 24 équipements → /32 chacun |
| Liens leaf-spine | 10.100.1.0/24 | 8 leafs × 16 spines = 128 liens → 128× /31 = 256 adresses ✅ (un /24 = 256 /31) |
| VTEP / anycast | = loopbacks | — |
| Management | 192.168.100.0/24 | 24 équipements + 2 ToR mgmt |
| ASN | Spines 65101-65116, leafs 65001-65008 | eBGP (voir §38) |

**Vérification** : 128 liens /31 = pile un /24 — **prévoir le /23** pour la
croissance (P13). Un plan d'adressage trop juste = renumérotation = nuit
blanche.

## 88. Pièges câblage — l'essentiel (détail §107)

- **P31** — Fibre coincée sous un rail : pertes invisibles, BER élevée.
- **P32** — Goulotte à 120 % : dizaines de liens dégradés en même temps.
- **P33** — Câble non étiqueté : 1 h de debug à 150 €/h.
- **P34** — Plan d'adressage trop juste : renumérotation nocturne.

# PARTIE I — ÉNERGIE : LE RÉSEAU SE CHIFFRE EN kW

## 89. Conso des NICs par débit — tableau de référence

| NIC | Débit | Conso carte (⚠️/✅) | Note |
|---|---|---|---|
| Intel E810 2×100G | 100G | **15 W** ✅ (27,1 W avec AOC ✅) | La plus sobre du 100G |
| NVIDIA CX-6 Dx 2×100G | 100G | ~20-25 W ⚠️ | — |
| NVIDIA CX-7 1×400G | 400G | ~30-40 W ⚠️ | + optique ~9 W |
| Broadcom Thor 2 1×400G | 400G | ~25-35 W ⚠️ | 5 nm, argument efficacité ✅ |
| NVIDIA CX-8 1×800G | 800G | ~50-75 W ⚠️ | + optique ~15 W |
| BlueField-3 | 400G | ~75 W ⚠️ | = NIC + ARM + accélérateurs |
| BlueField-4 | 800G | ❌ non publiée au 27/09/2026 | Grace 64 cœurs : prévoir large |

**Efficacité par bit** (ordre de grandeur ⚠️) :
- 100G : ~0,2 W/Gb/s → 400G : ~0,1 W/Gb/s → 800G : ~0,08 W/Gb/s.
**Monter en débit divise par ~2,5 le joule par bit.** C'est l'argument
n°1 en comité d'investissement : le 800G n'est pas un luxe, c'est du
PUE.

## 90. Conso des switches par débit — tableau de référence

| Switch | Ports | Conso typique | Conso max optiques | Source |
|---|---|---|---|---|
| Leaf 48×25G + 8×100G | 100G | ~300-400 W ⚠️ | ~500 W ⚠️ | Usuel |
| Leaf 32×400G | 400G | ~600-800 W ⚠️ | ~1,2 kW ⚠️ | Usuel |
| SN5600 64×800G | 800G | **940 W** (cuivre) ✅ | **~2 kW** (optiques) ⚠️ | Datasheet ✅ |
| SN5610 64×800G | 800G | **0,9 kW** ✅ | **2,08 kW** (64 optiques) ✅ | Datasheet ✅ |
| Arista 7060X6 64×800G | 800G | (2× 2400 W PSU) ✅ | Dimensionné pire cas ✅ | Datasheet ✅ |
| Arista 7060X5 32×800G | 800G | (2× 1500 W PSU) ✅ | ✅ | Datasheet ✅ |

**Lecture** : sur un switch 800G tout-optique, **les optiques consomment
autant que le switch lui-même**. Tout levier sur les optiques (LPO §28,
CPO §101) = −30 à −70 % sur la moitié de la facture du switch.

## 91. Conso des optiques et câbles — le détail qui fait le kW

| Média | Conso / extrémité (⚠️ sauf note) |
|---|---|
| DAC 100G/400G/800G | 0,1-0,5 W |
| AEC 800G | 3-4 W |
| AOC 100G/400G | 1-2 W |
| 100G-SR4 | ~2,5 W |
| 400G-DR4 | ~8-10 W |
| 400G-SR8 | ~8-10 W |
| 800G-SR8/DR8 | ~12-16 W |
| 800G (max port OSFP) | **18 W** ✅ (SN5600, vérifié 27/09/2026) |

**Calcul éclair** : 1000 liens 800G-DR8 = 2000 extrémités × 15 W =
**30 kW** rien que d'optiques. Sur un fabric 10k GPU, les optiques =
**plusieurs centaines de kW**. C'est un poste PUE à part entière.

## 92. Chaleur des OSFP : à prévoir, pas à subir

