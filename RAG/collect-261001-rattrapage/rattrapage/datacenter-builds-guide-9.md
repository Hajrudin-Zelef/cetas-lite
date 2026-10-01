---
id: collect-261001-rattrapage/rattrapage/datacenter-builds-guide-9
title: "Datacenter Builds — Le guide des BOMs"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Nvidia", "vLLM"]
dates: ["2026-09-27"]
keywords: ["datacenter", "amd", "attention", "awq", "benchmark", "blackwell", "capex", "fp8", "gptq", "gpu", "hbm", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/datacenter_builds_guide.md
source_anchor: ""
source_lines: [1438, 1631]
sha256: 7604536fe5903ea5eba8b8fb6f81594b44a889ef02a0c97975c82701bbe85448
---

# Datacenter Builds — Le guide des BOMs

- **Débit** : un 70B FP8 sur 2× MI300X sort ≈ 100–200 tok/s en batch modeste
  (à vérifier sur votre modèle — mesurez avec un benchmark).
- **Parallélisme** : tensor-parallel sur les GPU du nœud, data-parallel entre
  réplicas pour la concurrence.
- **Quantification** : AWQ/GPTQ (poids seuls) ou FP8 (poids+activations) —
  le FP8 natif Blackwell/MI300 est le plus simple.
- **RAM CPU** : 2× la taille du modèle pour le chargement/déchargement
  (70B → 140 Go RAM CPU mini).

---

## 87. Cas chiffré : servir un 70B FP8 à 50 utilisateurs

| Poste | Détail | Total |
|---|---|---|
| 1× serveur 2× MI300X (section 82) | — | 77 560 € |
| Switch 100G (partagé) | quote-part | 5 000 € |
| Électricité (1,8 kW × 8 760 h × 0,20 €) | /an | 3 150 €/an |
| **CAPEX** | | **≈ 82 500 €** |
| **Coût/1k tokens** (10 M tok/j) | 3 150 / 365 / 10 000 | **≈ 0,0009 €** |

À comparer à l'API cloud : 70B à ≈ 0,50–0,90 $/M tokens → 10 M tok/j =
**5 000–9 000 $/j**. Le serveur est amorti en **moins de 3 semaines** à ce
rythme. **L'on-premise gagne dès que l'usage est soutenu** — c'est le calcul
à refaire pour chaque projet IA.

---

## 88. Refroidissement nœud MI300X

- 2× MI300X (2,3 kW) : air possible en 2U bien ventilé, RDHx conseillé.
- 4× MI300X (4,4 kW) : **DLC ou RDHx obligatoire** — l'air seul ne suit plus
  dans un datacenter à 25 °C.
- Les HBM sont sensibles à la température : **T_junction < 95 °C**,
  throttling au-delà. Surveillez via `rocm-smi`.

---

## 89. Pièges terrain — Inférence AMD

1. **Acheter avant de tester le modèle** : 5 % des architectures ne passent
   pas sous ROCm. Louez 10 h à 1,99 $/h pour valider AVANT la BOM.
2. **Oublier la RAM CPU** : le chargement d'un 405B demande 800 Go de RAM
   système — sinon OOM au démarrage.
3. **NVLink vs Infinity Fabric** : la bande passante inter-GPU AMD (Infinity
   Fabric) est bonne en intra-nœud, mais l'écosystème multi-nœuds reste
   derrière NVLink. Restez en intra-nœud pour l'inférence.
4. **Pilotes ROCm** : épinglez la version (6.x validée avec votre vLLM),
   ne suivez pas « latest » en prod.

---

## 90. Ce qu'il faut retenir — Inférence AMD

- 192 Go HBM à 10–12 k$ : le meilleur €/Go VRAM du marché 2026.
- 2× MI300X = 70B FP8 pour ≈ 78 k€ ; 4× = 405B FP8 pour ≈ 144 k€.
- Validez le modèle à l'heure en cloud avant d'acheter.
- 4,4 kW en 4U : prévoyez le liquide dès la commande.

---

## 91. Workload 8 — Inférence NVIDIA : panorama 2026

| GPU | VRAM | TDP | Prix constaté | Positionnement |
|---|---|---|---|---|
| L40S | 48 Go | 350 W | 8 610–8 900 $ | Le standard inférence PCIe |
| RTX PRO 6000 Blackwell | 96 Go | 600 W | 9 450–9 800 $ | 70B sur 1 carte |
| RTX 6000 Ada | 48 Go | 300 W | 7 400–7 800 $ | Alternative L40S (occasion intéressante) |
| H100 SXM | 80 Go | 700 W | 25–35 k$ neuf | Haute concurrence, SLA prod |
| H200 | 141 Go | 700 W | 30–40 k$ | Gros modèles, training |
| B200 | 192 Go | 1 000 W | 30–40 k$ | Génération Blackwell, délais 3–7 mois |

Tous prix vérifiés le 27/09/2026 (snapshot mars 2026 + MSRP). L'écosystème
CUDA reste l'argument n° 1 : tout tourne, tout est documenté.

---

## 92. Serveur 2× L40S — BOM

| Composant | Référence type | Qté | Prix |
|---|---|---|---|
| Châssis 2U 4× GPU double-slot | ASUS ESC4000A-E12 nu | 1 | ≈ 5 000 € |
| EPYC 9455P (48c) | — | 1 | 3 406 € |
| RAM 12× 64 Go = 768 Go | DDR5 ECC | 12 | 16 560 € |
| 2× NVMe 3,84 To (modèles) | PM9A3 | 2 | 6 070 € |
| 2× NVIDIA L40S 48 Go | — | 2 | ≈ 8 000 € × 2 = 16 000 € |
| NIC 2× 25 GbE (+ 2× 100 GbE option) | E810 | 1 | 450 € |
| PSU 2× 2 600 W | incluses | — | — |
| **TOTAL** | | | **≈ 47 485 €** |
| P_max / réaliste | 300+120+60+700+50+150 | | **≈ 1,38 kW / 1,05 kW** |

96 Go VRAM : 70B en FP8 sur les 2 cartes (tensor-parallel), ou 2× 32B
indépendants.

---

## 93. Serveur 4× L40S — BOM (« quad »)

| Composant | Référence type | Qté | Prix |
|---|---|---|---|
| Châssis 2U 4× GPU double-slot | ASUS ESC4000A-E12 nu | 1 | ≈ 5 000 € |
| EPYC 9555P (64c) | 2× | 2 | 4 830 € × 2 = 9 660 € |
| RAM 24× 64 Go = 1,5 To | DDR5 ECC | 24 | 33 120 € |
| 4× NVMe 7,68 To (modèles) | PM9A3 | 4 | 15 660 € |
| 4× NVIDIA L40S 48 Go | — | 4 | 8 000 € × 4 = 32 000 € |
| NIC 2× 100 GbE | CX6-DX | 1 | 1 100 € |
| PSU 2× 2 600 W | incluses | — | — |
| **TOTAL** | | | **≈ 96 540 €** |
| P_max / réaliste | 720+240+120+1400+60+150 | | **≈ 2,69 kW / 2,0 kW** |

192 Go VRAM : 2× 70B FP8 en parallèle, ou 1× 70B FP16 confortable.

---

## 94. Serveur 2× RTX PRO 6000 96 Go — BOM

| Composant | Référence type | Qté | Prix |
|---|---|---|---|
| Châssis 2U 4× GPU (600 W/GPU !) | plateforme 600W-capable (devis — à vérifier) | 1 | ≈ 6 500 € |
| EPYC 9455P (48c) | — | 1 | 3 406 € |
| RAM 12× 64 Go = 768 Go | DDR5 ECC | 12 | 16 560 € |
| 2× NVMe 3,84 To | PM9A3 | 2 | 6 070 € |
| 2× RTX PRO 6000 96 Go | — | 2 | ≈ 8 800 € × 2 = 17 600 € |
| NIC 2× 25 GbE | E810 | 1 | 450 € |
| PSU 2× 3 000 W | — | — | incluses |
| **TOTAL** | | | **≈ 50 585 €** |
| P_max / réaliste | 300+120+60+1200+50+150 | | **≈ 1,88 kW / 1,4 kW** |

**Attention : 600 W par GPU** — vérifiez que le châssis et le refroidissement
sont qualifiés 600 W/GPU. Tous les 2U 4-GPU ne le sont pas.

---

## 95. Serveur 4× RTX PRO 6000 96 Go — BOM (« quad »)

| Composant | Référence type | Qté | Prix |
|---|---|---|---|
| Châssis 4U 4× GPU 600 W | plateforme qualifiée (devis) | 1 | ≈ 9 000 € |
| 2× EPYC 9555P (64c) | — | 2 | 9 660 € |
| RAM 24× 64 Go = 1,5 To | DDR5 ECC | 24 | 33 120 € |
| 4× NVMe 7,68 To | PM9A3 | 4 | 15 660 € |
| 4× RTX PRO 6000 96 Go | — | 4 | 8 800 € × 4 = 35 200 € |
| NIC 2× 100 GbE | CX6-DX | 1 | 1 100 € |
| PSU 4× 3 000 W (N+1) | incluses | — | — |
| **TOTAL** | | | **≈ 103 740 €** |
| P_max / réaliste | 720+240+120+2400+60+200 | | **≈ 3,74 kW / 2,8 kW** |

384 Go VRAM sur 4 cartes : 4× 70B FP8 indépendants, ou 405B FP8 en
tensor-parallel sur le nœud. 3,74 kW : **DLC/RDHx obligatoire**.

---

## 96. H100 / H200 : quand les choisir

| Critère | L40S / RTX PRO 6000 | H100 / H200 |
|---|---|---|
| Prix/GPU | 8–10 k$ | 25–40 k$ |
| Bande passante mémoire | 864–1 792 Go/s | 2 000–4 800 Go/s |
| Concurrence (req/s) | moyenne | très haute |
| NVLink | non (PCIe) | oui (SXM) |
| Quand | inférence classique | SLA prod, batch énorme |

Règle : **tant que la concurrence < 50 req/s sur un 70B, le PCIe suffit**.
Le H100 ne se justifie que par la bande passante HBM ou le NVLink 8-GPU.

---

## 97. Tableau de choix GPU inférence 2026

| Besoin | Choix prix/perf | Budget | P_max |
|---|---|---|---|
| 8–32B, dev | 2× L40S | ≈ 47 k€ | 1,4 kW |
| 70B FP8, 1 carte | 1–2× RTX PRO 6000 96 Go | ≈ 35–50 k€ | 1,2–1,9 kW |
| 70B, petit prix | 2× MI300X (AMD) | ≈ 78 k€ | 2,3 kW |
| 70B haute concurrence | 4× L40S | ≈ 97 k€ | 2,7 kW |
| 405B FP8 | 4× MI300X ou 4× RTX PRO 6000 | ≈ 104–144 k€ | 3,7–4,4 kW |

---

## 98. Cas comparé : 70B FP8 — 1× RTX PRO 6000 vs 2× L40S

| Critère | 1× RTX PRO 6000 96 Go | 2× L40S 48 Go |
|---|---|---|
| VRAM totale | 96 Go (1 carte) | 96 Go (2 cartes) |
| Prix GPU | ≈ 8 800 € | ≈ 16 000 € |
| Serveur total | ≈ 35 000 € | ≈ 47 500 € |
| P_max GPU | 600 W | 700 W |
| Simplicité | 1 carte, pas de TP | tensor-parallel requis |
| Verdict | **gagne** si le modèle tient en 96 Go | utile si déjà en parc |

---

## 99. Pièges terrain — Inférence NVIDIA

