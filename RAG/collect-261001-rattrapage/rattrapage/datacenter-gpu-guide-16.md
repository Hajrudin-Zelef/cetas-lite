---
id: collect-261001-rattrapage/rattrapage/datacenter-gpu-guide-16
title: "Les GPU datacenter / IA — présent et futur vérifié"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Nvidia", "United States"]
dates: []
keywords: ["gpu", "amd", "blackwell", "capex", "compute", "fp4", "fp8", "hbm3", "hbm4", "kv cache", "lora", "mi455x"]
source: docs/RAG/collect-261001-rattrapage/datacenter_gpu_guide.md
source_anchor: ""
source_lines: [2274, 2453]
sha256: 588b4c40cd9d9f87dddeb494f8ad50880ee809c6cef1ff22b524aea58d8483b8
---

# Les GPU datacenter / IA — présent et futur vérifié

| Mode | Avantages | Inconvénients | Quand |
|---|---|---|---|
| Achat cash | Pas d'intérêts, actif amortissable | Capex massif, risque obsolescence | Trésorerie saine, usage > 70 % |
| Leasing 36 mois | OPEX, renouvellement programmé | Coût total +15–25 % | Usage > 60 %, cycle 3 ans |
| Location cloud | Zéro capex, flexibilité | 2× plus cher à 24/7 | Usage < 60 %, burst, test |
| Crédit-bail | Mixte | Complexité comptable | Cas spécifiques |

**Règle** : le GPU se déprécie de ~30 %/an — tout financement qui dépasse 36 mois
fait payer un actif qui ne vaut plus rien à la fin.
## 113. Note de calcul modèle : salle 2 nœuds 8× B200 (cas complet)

### 113.1. Hypothèses

- 2 nœuds 8× B200 HGX, DLC, PUE 1,15, 0,15 €/kWh, utilisation 80 %.
- Durée d'analyse : 5 ans.

### 113.2. Bilan de puissance

| Poste | Détail | Puissance |
|---|---|---|
| GPU | 16 × 1 000 W | 16,0 kW |
| CPU | 4 × 350 W | 1,4 kW |
| NIC 800G | 16 × 30 W | 0,5 kW |
| Ventilateurs + stockage | — | 1,0 kW |
| Switches IB (2) | 2 × 1,5 kW | 3,0 kW |
| **IT total** | | **21,9 kW** |
| Refroidissement (PUE 1,15) | 0,15 × 21,9 | 3,3 kW |
| **Total site** | | **25,2 kW** |

### 113.3. Électrique

- Courant triphasé : 25 200 / (√3 × 400 × 0,9) ≈ **40 A/phase** → TGBT 100 A.
- Onduleur : 2× 30 kVA N+1 (ou 1× 60 kVA modulaire).
- Groupe : 75 kVA (1,5× IT + marge démarrage).

### 113.4. Coûts 5 ans

| Poste | Montant |
|---|---|
| 2 serveurs 8× B200 | 920 000 $ |
| Réseau IB | 120 000 $ |
| DLC (CDU + dry cooler) | 180 000 $ |
| Élec (TGBT, onduleur, groupe) | 150 000 $ |
| Électricité 5 ans (80 %, PUE 1,15) | 25,2 × 0,8 × 8 760 × 5 × 0,15 = **132 000 €** |
| Maintenance 5 ans (5 %/an du HW) | 230 000 $ |
| **Total 5 ans** | **~1 730 000 $** |

Soit ~108 k$/GPU/5 ans tout compris. Le chiffre à comparer à la location :
16 GPU × 5,50 $/h × 8 760 × 5 × 0,8 = **3 080 000 $** — l'achat gagne de **43 %**.

## 114. Contexte long : dimensionnement détaillé par GPU (70B FP8)

Rappel : poids 80 Go + overhead 15 % = 92 Go fixes. Reste pour KV cache (BF16) :

| GPU | VRAM | Reste pour KV | Tokens max (BF16) | Tokens max (KV FP8) |
|---|---|---|---|---|
| RTX PRO 6000 (96 Go) | 96 | ~4 Go | ~3K | ~6K |
| H100 80 Go | 80 | — (ne tient pas) | — | — |
| H200 (141 Go) | 141 | ~49 Go | ~39K | ~78K |
| MI300X (192 Go) | 192 | ~100 Go | ~80K | ~160K |
| MI325X (256 Go) | 256 | ~164 Go | ~131K | ~262K |
| MI355X (288 Go) | 288 | ~196 Go | ~157K | ~314K |
| B200 (192 Go) | 192 | ~100 Go | ~80K | ~160K |

**Lecture** : pour du 128K+ sérieux sur 70B, il faut MI325X/MI355X en 1 GPU,
ou 2× H200/B200 en tensor-parallel. La RTX PRO 6000 est exclue du contexte long 70B.

## 115. Coût par million de tokens : estimation par GPU

Hypothèses : 70B FP8, batch optimal, GPU à 80 %, élec 0,15 €/kWh, amortissement
3 ans, PUE 1,4.

| GPU | Débit (tok/s) | Coût GPU/h (amorti) | Coût/M tokens |
|---|---|---|---|
| RTX PRO 6000 | ~55 | 0,55 $ | ~2,78 $ |
| H100 | ~60 | 1,37 $ | ~6,34 $ |
| H200 | ~65 | 1,60 $ | ~6,84 $ |
| B200 | ~110 | 1,83 $ | ~4,62 $ |
| MI300X | ~68 | ~1,60 $ (est.) | ~6,53 $ |
| MI355X | ~110 | ~1,90 $ (est.) | ~4,80 $ |

Le B200 et le MI355X écrasent le coût/token grâce au FP4/FP8 + 8 To/s.
La RTX PRO 6000 reste compétitive pour les petits volumes (pas de DLC, pas d'IB).

## 116. Training : calculateur de coût par run

```
coût_run = N_GPU × heures × prix_h × (1 + overhead 15 %)
```

| Run | Config | Durée | Coût (location 3,33 $/h) |
|---|---|---|---|
| LoRA 8B (100M tokens) | 1× 4090 | 12 h | ~5 $ (0,34 $/h) |
| Full FT 8B (10Md tokens) | 8× H100 | 2 j | ~1 280 $ |
| Full FT 70B (100Md tokens) | 32× H100 | 10 j | ~25 600 $ |
| Pre-train 70B (1 000Md) | 256× H100 | 30 j | ~614 000 $ |
| Pre-train 70B (1 000Md) | 128× B200 | 14 j | ~430 000 $ (5,50 $/h) |

Le B200 divise le coût des gros runs par ~1,4 malgré un prix/h supérieur —
c'est l'arithmétique du MFU et du FP8.

## 117. Tableau : générations NVIDIA — l'évolution en chiffres

| Génération | GPU | Année | Mémoire | BW | FP8 dense | TDP | NVLink |
|---|---|---|---|---|---|---|---|
| Pascal | P100 | 2016 | 16 Go HBM2 | 732 Go/s | — | 300 W | 160 Go/s |
| Volta | V100 | 2017 | 16/32 Go HBM2 | 900 Go/s | — | 300 W | 300 Go/s |
| Ampere | A100 80 Go | 2020 | 80 Go HBM2e | 2 To/s | — | 400 W | 600 Go/s |
| Hopper | H100 | 2022 | 80 Go HBM3 | 3,35 To/s | 1 979 TF | 700 W | 900 Go/s |
| Hopper+ | H200 | 2024 | 141 Go HBM3e | 4,8 To/s | 1 979 TF | 700 W | 900 Go/s |
| Blackwell | B200 | 2024 | 192 Go HBM3e | 8 To/s | 4 500 TF | 1 000 W | 1,8 To/s |
| Blackwell Ultra | B300 | 2026 | 288 Go HBM3e | 8 To/s | 7 000 TF | 1 400 W | 1,8 To/s |
| Rubin | Rxxx | 2026 | HBM4 (20,7 To/rack) | n.c. | 50 PF NVFP4 | n.c. | 3,6 To/s |

**Tendance** : ×12 la mémoire et ×3,5 le FP8 en 6 ans (H100 → B300), TDP ×2.
La mémoire croît plus vite que le compute — le marché paie pour la mémoire.

## 118. Tableau : générations AMD — l'évolution en chiffres

| Génération | GPU | Année | Mémoire | BW | FP8 dense | TDP |
|---|---|---|---|---|---|---|
| CDNA 2 | MI250X | 2021 | 128 Go HBM2e | 3,2 To/s | — | 560 W |
| CDNA 3 | MI300X | 2023 | 192 Go HBM3 | 5,3 To/s | 2 615 TF | 750 W |
| CDNA 3+ | MI325X | 2024 | 256 Go HBM3e | 6 To/s | 2 615 TF | 1 000 W |
| CDNA 4 | MI350X | 2025 | 288 Go HBM3e | 8 To/s | ~4 600 TF | 1 000 W |
| CDNA 4+ | MI355X | 2025 | 288 Go HBM3e | 8 To/s | ~5 000 TF | 1 400 W |
| CDNA 5 | MI455X | 2026 | 432 Go HBM4 | 19,6 To/s | 20 000 TF | ~1 200–1 500 W |

AMD a comblé l'écart compute en 2 générations (CDNA 3 → 4) et mène sur la
mémoire par GPU depuis le MI300X.

## 119. Matrice de décision : quel GPU pour quel budget (achat)

| Budget | Config | VRAM totale | Usage |
|---|---|---|---|
| < 3 k$ | 1× RTX 4090 (occasion) | 24 Go | Lab, 8B–13B |
| 3–8 k$ | 1–2× RTX 5090 | 32–64 Go | Lab, 32B Q4 |
| 8–15 k$ | 1× L40S | 48 Go | Inférence pro légère |
| 15–25 k$ | 1× RTX PRO 6000 | 96 Go | 70B FP8 1 GPU |
| 25–60 k$ | 4× L40S / 2× PRO 6000 | 192 Go | RAG, multi-modèles |
| 100–200 k$ | 4× H200 NVL | 564 Go | 405B Q4, 70B gros ctx |
| 250–350 k$ | 8× H100 | 640 Go | Training standard |
| 350–500 k$ | 8× H200 / 8× MI300X | 1,1–1,5 To | Training + gros ctx |
| 450–600 k$ | 8× B200 / 8× MI355X | 1,5–2,3 To | Frontier |
| > 1 M$ | Multi-nœuds B200 + IB + DLC | 3 To+ | Prod à l'échelle |

## 120. Matrice de décision : quel GPU pour quel budget (location / mois 24/7)

| Budget/mois | Config | Usage |
|---|---|---|
| < 300 $ | 1× 4090 spot | Dev perso |
| 300–800 $ | 1× L40S / 1× 5090 | Tests, petits modèles |
| 800–2 500 $ | 1× H100 / 2× L40S | Inférence 70B |
| 2 500–5 000 $ | 1× H200 / 1× B200 spot | Gros modèles ponctuels |
| 5 000–20 000 $ | 8× H100 réservé | Training sérieux |
| > 20 000 $ | 8× B200 / cluster | Frontier |

En dessous de 60 % d'utilisation mensuelle, la location gagne toujours.
## 121. Pièges terrain (5/5) — derniers

### Piège n°31 : le « GPU pas cher » sans le serveur autour

Un B200 à 40 k$ sans HGX : inutilisable seul (pas de refroidissement, pas d'alim,
pas de baseboard). Le prix « carte » n'a de sens que dans le système. Toujours
chiffrer le **système complet**.

### Piège n°32 : la TVA et les douanes sur les importations

Serveur 8× H100 importé hors UE : +20 % TVA + droits de douane éventuels +
délai dédouanement 2–4 semaines. Un devis US à 285 k$ = ~350 k€ rendu France.
Faire chiffrer en DDP (rendu droits acquittés).

### Piège n°33 : l'assurance transport

Un serveur 500 k$ qui voyage : assurance ad valorem (~0,3–0,5 % de la valeur).
Un choc pendant le transport = baseboard fissurée = 100 k$ de dégâts invisibles.
Exiger le transport sur palette amortissante + assurance tous risques.

### Piège n°34 : le rack qui ne passe pas la porte

