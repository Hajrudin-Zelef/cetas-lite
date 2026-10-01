---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-12
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: ["Broadcom", "Nvidia", "TSMC"]
dates: ["2025-10-08", "2026-09-27"]
keywords: ["datacenter", "arr", "cpo", "dsp", "ethernet", "gpu", "lpo", "npo", "nvidia", "optics", "packaging", "rubin"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [1650, 1798]
sha256: 34b493c1e7a86f91852c1326e732cf0a5da7280780665708c1a4331fa9cb9fec
---

# RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE

- 18 W dans ~10 cm³ : sans airflow ≥ 2 m/s (⚠️), la température boîtier
  dépasse 80 °C → le module **réduit son débit** (thermal throttling) ou
  coupe.
- **Allées chaudes/froides** : les switchs réseau soufflent généralement
  **vers l'arrière** (C2P : connector-to-power) — les positionner en
  cohérence avec les serveurs (qui soufflent aussi vers l'arrière en
  général). Un switch à l'envers dans une allée froide = il aspire son
  propre air chaud.
- **Sonde** : chaque baie réseau = 1 sonde de température en haut
  (l'air chaud monte : +5-10 °C entre bas et haut de baie, ⚠️).
- **Water-cooling** : en 2026, les racks IA sont à eau, mais les switchs
  restent à air — **prévoir des allées froides dédiées** aux baies réseau
  dans un datacenter à eau, sinon elles deviennent le point chaud.

## 93. PUE partiel réseau — le calcul honnête

PUE = Énergie totale / Énergie IT. Pour le **lot réseau seul** :

```
  PUE_réseau = (IT_réseau + clim_réseau + pertes_élec_réseau) / IT_réseau
```

Exemple pod §62 : 181 kW IT réseau, clim ~60 kW (COP ~3, ⚠️), pertes
~10 kW → PUE_réseau ≈ (181+60+10)/181 ≈ **1,39**.
**Le réseau a un meilleur PUE que le GPU** (moins de water-cooling, moins
de pertes de conversion) — mais il tourne **24/7 à charge quasi constante**,
contrairement aux GPU (pics de training). En TCO annuel, le réseau pèse
plus que son % de puissance ne le suggère.

## 94. Refroidissement : air vs liquide pour le réseau

| Solution | Pour le réseau | Note |
|---|---|---|
| Air (allées chaudes/froides) | **Standard** | Suffit jusqu'à ~15 kW/baie |
| Confinement d'allée | Recommandé | +10-15 % d'efficacité (⚠️) |
| Eau (plaques froides) | Rare sur switchs en 2026 | Les switchs restent à air |
| Immersion | ❌ | Incompatible avec les optiques enfichables |

**Règle** : une baie réseau 800G dense = **8-12 kW** (⚠️) — dimensionner la
clim de la rangée réseau comme une rangée serveurs, pas comme du « brassage
télécom ». L'erreur classique : la salle réseau dimensionnée pour du 3 kW/
baie des années 2010 (P35, §107).

## 95. Dimensionnement électrique d'une rangée IA (méthode)

1. **Inventaire** : switchs (max, pas typique) + NICs + optiques + 20 %.
2. **PDU** : 2× PDU par baie (A/B), chacune capable de **100 % de la charge**
   (si A tombe, B prend tout).
3. **Disjoncteurs** : calibre ≥ 125 % du max (appel au boot, §61).
4. **Onduleurs** : le réseau sur l'onduleur **prioritaire** — c'est lui qui
   porte la supervision d'arrêt (NUT, voir guide onduleurs) et le
   management. Un réseau qui tombe avant les serveurs = arrêt à l'aveugle.
5. **Mesure** : PDU metered par prise sur les baies réseau — sans mesure,
   pas d'optimisation.

## 96. CAS CHIFFRÉ — 1 MW de réseau : que peut-on alimenter ?

Avec 1 MW IT réseau (PUE 1,4 → 1,4 MW au compteur), on alimente au choix :

| Option | Composition | Endpoints |
|---|---|---|
| A. Backend IA 800G | ~650 switchs 64×800G ? Non — réaliste : ~200 switchs + 8000 NIC 800G + optiques | ~8000 GPU |
| B. Cloud 100G | ~1500 switchs 32×100G + 40 000 NIC 100G | ~40 000 serveurs |
| C. Mixte | 1 pod IA (272 kW, §62) + 3000 serveurs 100G | — |

**Lecture** : 1 MW réseau = l'ordre de grandeur d'un **datacenter IA
moyen** (10 MW IT total dont ~10 % réseau). À 0,15 €/kWh (⚠️) : **1,84 M€/an**
d'électricité pour le seul réseau. Le choix LPO/CPO (§97) = des centaines
de k€/an.

## 97. Leviers d'efficacité : CPO, LPO, 800G

| Levier | Gain (✅/⚠️) | Maturité |
|---|---|---|
| Passer de 400G à 800G | ~−40 % W/Gb/s ⚠️ | Immédiat |
| LPO vs DSP (intra-DC) | −30 à −50 % par module ⚠️ | Croissante |
| CPO (TH6-Davisson) | **3,5 W/port à 800G** vs ~15 W pluggable = **−77 %** ✅ | Early access (10/2025) |
| Quantum-X800 CPO | **9 W/port** vs 30 W pluggable = **−70 %** ✅ | Annoncé |
| DAC là où c'est possible | ~0 W vs 15 W | Toujours |

**Message pour la direction** : le réseau IA 2026-2028 va diviser par
**3 à 4** sa conso par bit grâce au CPO. Tout investissement 800G
pluggable doit intégrer cette trajectoire : prévoir des switchs
**compatibles CPO ou à cycle de remplacement 3 ans**, pas 7.

## 98. Pièges énergie — l'essentiel (détail §107)

- **P35** — Salle réseau climatisée pour 3 kW/baie : insuffisant en 800G.
- **P36** — Disjoncteurs au nominal : déclenchement au boot.
- **P37** — Réseau pas sur onduleur prioritaire : arrêt à l'aveugle.
- **P38** — PUE calculé sans les optiques : −50 % du poste oublié.

# PARTIE J — À VENIR : ANNONCES VÉRIFIÉES AU 27/09/2026

> Seules figurent ici des annonces officielles (constructeur, consortium)
> retrouvées le 27/09/2026. Toute donnée issue de la presse ou d'analystes
> est explicitement signalée (⚠️) et ne compte pas comme annonce officielle.

## 99. 1,6T Ethernet — statut vérifié

| Annonce | Détail | Statut |
|---|---|---|
| NVIDIA ConnectX-9 SuperNIC | **1,6 Tb/s** scale-out, associé à l'architecture **Rubin (2ᵉ sem. 2026)** | ✅ Annoncé |
| Broadcom Taurus | Premier DSP optique **400G/lane** → modules **1,6T** (et base 3,2T) | ✅ Dévoilé OFC 2026 (mars 2026) |
| Tomahawk 6 | 64 ports **1,6T** (ou 128×800G, 512×200G) | ✅ Annoncé |
| Modules 1,6T | En développement sur DSP Taurus | ⏳ Pas de volume au 27/09/2026 |

**Technique** : le 1,6T = 8 lanes × 200G PAM4 (électrique 200G/lane).
Le verrou n'est plus le SerDes mais **l'optique et la thermique** :
un module 1,6T DSP ≈ 25-30 W (⚠️ estimation) — d'où le CPO (§101).

## 100. Ultra Ethernet — roadmap après la 1.0

| Étape | Date | Statut |
|---|---|---|
| UEC Specification 1.0 | **11 juin 2025** | ✅ Publiée |
| UEC Specification 1.0.1 | Mi-2025 | ✅ |
| Premier NIC conforme (Thor Ultra 800G) | **14 octobre 2025** | ✅ |
| Première solution de validation (VIAVI) | **23 juin 2026** | ✅ |
| Interopérabilité multi-vendeurs large | 2027+ | ⏳ À venir |
| UEC 2.0 (1,6T, scale-up ?) | ❌ Non annoncée au 27/09/2026 | — |

**À surveiller** : les premiers fabrics UEC multi-vendeurs réels (pas les
démos), et la position NVIDIA (Spectrum-X vs UEC) — c'est le clivage
structurant du marché 2027-2028.

## 101. CPO (Co-Packaged Optics) — l'annonce la plus structurante

| Annonce | Détail | Statut |
|---|---|---|
| Broadcom Tomahawk 6-Davisson | **102,4T avec CPO**, 16 moteurs optiques 6,4T co-packagés ; **3,5 W/port à 800G** (−36 % vs TH5 CPO, −70 % vs pluggable) ; lasers **remplaçables par la face avant** ; 131 072 XPU en 2 tiers | ✅ Annoncé **08/10/2025**, early access |
| NVIDIA Quantum-X800 (variante CPO Q3450-LD) | **115,2T** InfiniBand, 144×800G ; **9 W/port** vs 30 W pluggable (⚠️ estimation analyste) | ✅ Datasheet NVIDIA (déc. 2024) |
| Partenaires Davisson | Celestica, Accton (box), Micas, Nexthop AI (SONiC durci), TSMC (packaging COUPE) | ✅ Cités par Broadcom |

**Pourquoi c'est structurant** : le CPO supprime le DSP + le connecteur
enfichable (les deux plus gros consommateurs et points de panne). La 3ᵉ
tentative de Broadcom (après 2 générations) semble la bonne : lasers
remplaçables (l'objection n°1 des hyperscalers est levée).
**Impact achat** : ne pas signer en 2026 des contrats 800G pluggables à
7 ans sans clause de sortie — le CPO change le TCO dès 2027-2028.

## 102. Après : 3,2T, 204,8T, l'ère 200T

| Annonce | Détail | Statut |
|---|---|---|
| Broadcom OFC 2026 | Feuille de route vers le **200T** : switchs **204,8T**, modules **3,2T** (sur DSP 400G/lane Taurus) | ✅ Annoncé mars 2026 |
| Broadcom NPO 3,2T | Optiques « near-packaged » VCSEL 3,2T | ✅ Annoncé OFC 2026 |
| Retimers 200G/lane, AEC | Jusqu'à 6 m en cuivre actif | ✅ Annoncé OFC 2026 |

