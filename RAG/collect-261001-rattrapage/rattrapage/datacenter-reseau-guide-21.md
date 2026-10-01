---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-21
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["arr", "asic", "cpo", "nvidia", "rubin", "wavelength"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [3049, 3168]
sha256: 3107bff2843772e6a5565934787c1460de5e1cd861a8f7f9c3394a96d91dad8e
---

# ECMP déséquilibré (écart-type des compteurs par lien)
- alert: EcmpImbalance
  expr: stddev by (device) (rate(if_out_octets[5m])) / avg by (device) (rate(if_out_octets[5m])) > 0.3
  for: 30m
  labels: { severity: warning }
```

## 150. FAQ terrain — 30 questions/réponses courtes

**1.** DAC ou AOC en intra-rack 400G ? → DAC si < 3 m, AOC sinon.
**2.** OM4 ou OS2 pour une rocade neuve ? → OS2 (15 ans), OM4 vers les baies.
**3.** 400G ou 800G pour un cluster neuf en 2026 ? → 800G (coût/Gb/s inférieur).
**4.** RoCEv2 sans PFC, ça marche ? → Non en production (il faut le lossless).
**5.** UEC remplace RoCEv2 ? → À terme ; en 2026 c'est une assurance, pas un remplacement.
**6.** Combien de cœurs CPU pour 400G sans offload ? → 8-16 (d'où les offloads).
**7.** PCIe Gen4 x16 pour une NIC 400G ? → Non, bridé à ~200G.
**8.** MTU pour RoCEv2 ? → 9000 partout, sans exception.
**9.** ECN ou PFC en premier ? → ECN (marquage), PFC en dernier recours.
**10.** Un switch 800G chauffe combien ? → ~1-2 kW tout-optique.
**11.** Nettoyer une fibre avec quoi ? → Cassette à sec d'abord (§24).
**12.** MPO mâle + MPO mâle ? → Jamais (casse les broches).
**13.** Breakout 800G → 8×100G : piège ? → Certains ports partagent des ressources (§23).
**14.** Oversubscription du backend IA ? → 1:1, toujours.
**15.** ECMP + un seul gros flux ? → Ne se répartit pas (spraying nécessaire).
**16.** DPU pour un serveur web ? → Non, surcoût inutile.
**17.** BlueField-3 ou 4 en 2026 ? → BF-3 dispo, BF-4 early availability 2026.
**18.** Chiffrer le backend NCCL en soft ? → Non (tue la perf), HW si exigé.
**19.** 1,6T : quand ? → ConnectX-9 / Rubin, H2 2026 (annoncé).
**20.** CPO : hype ou réel ? → Réel : Davisson early access 10/2025, −70 % conso.
**21.** Prix d'un switch 64×800G ? → Non public ; fourchette 150-300 k€ (⚠️).
**22.** Durée de vie d'un fabric IA ? → 3-4 ans, pas 7.
**23.** Tester un fabric : combien de temps ? → 72 h de burn-in minimum.
**24.** NCCL lent : premier réflexe ? → `NCCL_DEBUG=INFO` (quelle interface ?).
**25.** Un lien « monte » mais BER élevée ? → Nettoyer, puis budget optique.
**26.** Goulotte remplie à combien ? → 50 % max.
**27.** Étiquette : quel format ? → `SALLE-RANGÉE-BAIE-ÉQ-PORT` (TIA-606).
**28.** Onduleur du réseau : séparé ? → Oui, prioritaire, en dernier à l'arrêt.
**29.** IPv6 : quand ? → Plan prévu dès maintenant, activé plus tard.
**30.** La règle d'or ? → Mesurer avant de changer ; une chose à la fois.

---

*Fin du guide — 150 sections. Vérification : `wc -l` — objectif ≥ 4000 lignes.*

# PARTIE O — COMPLÉMENTS ULTRA-DENSES

## 151. Comparatif switches — 20 modèles du marché

Capacité = agrégée bidirectionnelle. Prix ⚠️ fourchettes marché.

| Modèle | ASIC | Ports max | Capacité | Format | Conso (⚠️) | Usage | Prix (⚠️) |
|---|---|---|---|---|---|---|---|
| NVIDIA SN5600 | Spectrum-4 | 64×800G | 51,2T | 2U | 0,94-2 kW | Backend IA | 150-300 k€ |
| NVIDIA SN5610 | Spectrum-4 | 64×800G | 51,2T | 2U | 0,9-2,08 kW | Backend IA | 150-300 k€ |
| Arista 7060X6 | TH5 | 64×800G | 51,2T | 2U | ~1,5-2 kW | Leaf IA | 150-300 k€ |
| Arista 7060X5 | TH4 | 32×800G | 25,6T | 1U | ~1-1,5 kW | Leaf | 80-150 k€ |
| Arista 7800R4 | J3-AI | 576×800G | 460T | Châssis | ~15-25 kW | Spine IA | 1-2 M€ |
| Arista 7280R4 | TH5 | 48×800G | 51,2T | 2U | ~1,5 kW | Leaf/spine | 150-250 k€ |
| Cisco 8102-64H | S1 G200 | 64×800G | 51,2T | 2U | ~1,5 kW | SP/IA | ⚠️ |
| Juniper QFX5240 | TH5 | 64×800G | 51,2T | 2U | ~1,5 kW | DC | 120-250 k€ |
| Edgecore DCS810 | TH5 | 64×800G | 51,2T | 2U | ~1,5 kW | Whitebox | 80-150 k€ |
| Celestica DS5000 | TH5 | 64×800G | 51,2T | 2U | ~1,5 kW | Whitebox | 80-150 k€ |
| Dell Z9664F | TH5 | 64×800G | 51,2T | 2U | ~1,5 kW | Enterprise | 120-250 k€ |
| HPE 9300S | TH5 | 64×800G | 51,2T | 2U | ~1,5 kW | Enterprise | 120-250 k€ |
| Arista 7060CX2 | TH3 | 32×100G | 6,4T | 1U | ~300 W | Leaf 100G | 15-30 k€ |
| Cisco N9K-C9336C | — | 36×100G | 7,2T | 1U | ~400 W | Enterprise | 20-40 k€ |
| Juniper QFX5120 | TH3 | 48×25G+8×100G | 4T | 1U | ~300 W | Leaf | 12-25 k€ |
| Edgecore DCS203 | TH3 | 32×400G | 12,8T | 1U | ~500 W | Leaf 400G | 25-50 k€ |
| NVIDIA SN4700 | Spectrum-3 | 32×400G | 12,8T | 1U | ~500 W | Leaf 400G | 40-80 k€ |
| Arista 7050X4 | TH4 | 32×400G | 12,8T | 1U | ~450 W | Leaf 400G | 40-70 k€ |
| FS S5860 | TH3 | 32×100G | 6,4T | 1U | ~300 W | Budget | 8-15 k€ |
| MikroTik CRS518 | — | 16×100G | 3,2T | 1U | ~150 W | Lab/PME | 3-6 k€ |

**Lecture** : à ASIC égal (TH5), l'écart whitebox→marque = **±50 %** —
c'est le support, le NOS et la validation qui se paient, pas le silicium.

## 152. Lire une datasheet optique — les 12 champs qui comptent

| Champ | Ce qu'il dit | Piège |
|---|---|---|
| Form factor | QSFP-DD, OSFP… | Compatibilité mécanique |
| Data rate | 400GBASE-DR4 | Le « 400G » seul ne suffit pas |
| Wavelength | 1310 nm / 850 nm / CWDM | Mélanger les λ = pas de link |
| Fiber type | SMF / MMF | Jaune vs aqua (§18) |
| Distance | 500 m, 2 km… | « Jusqu'à » = conditions idéales |
| Tx power (min/max) | −2,9 à +2 dBm (⚠️ ex.) | Sert au budget (§20) |
| Rx sensitivity | −5,4 dBm (⚠️ ex.) | Sert au budget |
| Link budget | 4 dB (⚠️ ex.) | Tx_min − Rx_sens |
| Power consumption | 9 W | Dimensionnement thermique |
| FEC | RS(544,514) requis | Les 2 bouts doivent l'activer |
| Operating temp | 0-70 °C (commercial) | Industriel = −40/+85 °C |
| Coding | Pré-codé constructeur ? | Lock optique (§11) |

**Règle** : exiger la datasheet **avant** l'achat, pas après la panne.
Un module sans datasheet = un module qu'on ne dépanne pas.

## 153. Budgets optiques — 3 exemples pas-à-pas

**Exemple A — 400G-DR4, 80 m, 1 brassage** :
- Module : Tx −2,9 dBm, sensibilité −5,4 dBm → budget 2,5 dB (⚠️).
- Fibre 80 m OS2 : 0,03 dB. 4× MPO : 4×0,4 = 1,6 dB. 1 brassage (+2 MPO) : +0,8 dB.
- Perte totale : ~2,4 dB. **Marge : 0,1 dB → INSUFFISANT.** Retirer le
  brassage (direct) ou passer en FR4 (budget ~4 dB).

**Exemple B — 100G-SR4, 60 m, direct** :
- Budget module : ~1,9 dB (⚠️). Fibre 60 m OM4 : 0,18 dB. 2× MPO : 0,8 dB.
- Perte : ~1 dB. **Marge : 0,9 dB → limite mais OK** si connecteurs propres.
  Prévoir nettoyage semestriel.

**Exemple C — 800G-DR8, 300 m, 1 brassage** :
- Budget : ~4 dB (⚠️). Fibre 300 m : 0,12 dB. 6× MPO : 2,4 dB.
- Perte : ~2,5 dB. **Marge : 1,5 dB → OK**, mais surveiller le pré-FEC :
  à 800G, 1,5 dB de marge part vite avec la poussière.

**Méthode** : toujours calculer le **pire cas** (connecteurs à 0,5 dB, pas
0,3) et exiger ≥ 3 dB sur les liens critiques. Un lien à 1 dB de marge
neuf sera en panne dans 2 ans.

## 154. Table câbles — 20 références types

