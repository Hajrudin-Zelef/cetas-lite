---
id: collect-261001-rattrapage/rattrapage/datacenter-gpu-guide-24
title: "Les GPU datacenter / IA — présent et futur vérifié"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Broadcom", "Nvidia"]
dates: ["2026-09-27"]
keywords: ["gpu", "amd", "benchmarks", "capex", "compute", "embeddings", "ethernet", "fine-tuning", "fp8", "hbm", "hbm3", "helios"]
source: docs/RAG/collect-261001-rattrapage/datacenter_gpu_guide.md
source_anchor: ""
source_lines: [3620, 3777]
sha256: ebff3db2a6870f5216fd484b6457b6b2dfb9479c116ce1c07195b0ecadcd1d54
---

# Les GPU datacenter / IA — présent et futur vérifié

Préparer le dossier **avant** les travaux — un avis défavorable bloque la mise
en service de 4 à 8 semaines.

## 205. Assurance : le questionnaire type de l'assureur

- Valeur totale des actifs (neuf) : ___ €
- Mesures anti-vol : contrôle d'accès, vidéo, alarme : ___
- Protection incendie : détection, extinction : ___
- Protection électrique : onduleur, parafoudres : ___
- Historique sinistres 5 ans : ___
- Maintenance : contrat ? Fréquence ? ___

Un dossier solide = -15 à -25 % de prime. Un dossier vide = refus ou surprime.

## 206. Le mot de la fin : l'IA est un métier de l'énergie

On parle de TFLOPS, mais on paie des MWh. Le GPU le plus rapide du monde ne sert
à rien sans le megawatt derrière — et le megawatt le moins cher est celui qu'on
ne consomme pas (PUE, free cooling, batching, quantification).

Zelef, avec ton double regard systèmes + énergies, tu es exactement au bon
endroit : 80 % des acheteurs de GPU sous-dimensionnent l'énergie. Toi, tu
commences par là — c'est ton avantage.

*Les chiffres de ce guide sont une base de travail vérifiée le 27/09/2026.
Le terrain a toujours raison : mesure, compare, ajuste.*
## 207. Interconnexions : le comparatif exhaustif

| Techno | Débit / GPU | Latence | Topologie | Écosystème |
|---|---|---|---|---|
| PCIe 4.0 x16 | 64 Go/s | ~1 µs | Étoile (CPU) | Universel |
| PCIe 5.0 x16 | 128 Go/s | ~1 µs | Étoile (CPU) | Universel |
| NVLink 4 | 900 Go/s | ~0,5 µs | Tout-à-tout (NVSwitch) | NVIDIA H100/H200 |
| NVLink 5 | 1,8 To/s | ~0,5 µs | Tout-à-tout (NVSwitch) | NVIDIA B200/B300 |
| NVLink 6 | 3,6 To/s | ~0,5 µs | Rack (NVL72) | NVIDIA Rubin |
| xGMI (MI300X) | ~896 Go/s agrégés | ~0,6 µs | Tout-à-tout (plateforme) | AMD |
| UALink | 260 To/s agrégés (rack) | ~0,6 µs | Rack (Helios) | Standard ouvert |
| IB NDR 400G | 50 Go/s / lien | ~1 µs | Fat-tree | NVIDIA/Mellanox |
| IB XDR 800G | 100 Go/s / lien | ~1 µs | Fat-tree | NVIDIA/Mellanox |
| Ethernet UEC 800G | 100 Go/s / lien | ~2 µs | Fat-tree | AMD/Broadcom/Cisco |

**Hiérarchie des débits** : HBM (8 To/s) > NVLink (1,8 To/s) > PCIe (128 Go/s) >
IB (100 Go/s). Tout algorithme distribué doit respecter cette hiérarchie —
les collectives lourdes restent intra-nœud (NVLink), le léger traverse l'IB.

## 208. Scénarios complets : 6 budgets, 6 architectures

### Scénario A — 10 k$ : lab 1× RTX 5090
32 Go, 575 W, 130 tok/s sur 8B. Idéal dev. Limite : pas d'ECC.

### Scénario B — 25 k$ : inférence 70B 1× RTX PRO 6000
96 Go ECC, ~700 W nœud, 50 tok/s sur 70B FP8. Le meilleur $/token sur site.

### Scénario C — 120 k$ : RAG 8× L40S
384 Go ECC, 4,8 kW, embeddings + rerank + LLM. ROI 20 mois vs cloud.

### Scénario D — 400 k$ : training 8× H100 + IB + élec
640 Go HBM3, 8 kW, 15,8 PFLOPS FP8. Le standard éprouvé, prix en baisse.

### Scénario E — 1,4 M$ : cluster 2 nœuds 8× B200 + DLC
3 To HBM3e, 25 kW site, 72 PFLOPS FP8. La prod frontier 2026.

### Scénario F — 0 $ capex : 100 % cloud
8× H100 réservés à ~20 k$/mois. Gagnant si usage < 60 % ou besoin < 18 mois.

## 209. FAQ technique (3/3)

**Q : Quelle différence entre H100 SXM et H100 PCIe 96 Go ?**
R : Le SXM (700 W, 80 Go, 3,35 To/s) va sur baseboard HGX avec NVLink natif.
Le PCIe 96 Go (700 W, 1,68 To/s) est une carte slot — plus de VRAM mais moins
de bande passante et pas de NVLink inter-GPU natif.

**Q : Peut-on faire du training sur des RTX PRO 6000 ?**
R : Oui pour du fine-tuning et du petit training (8× = 768 Go, 6,4 PFLOPS FP8),
mais sans NVLink le multi-GPU est 3–5× moins efficace qu'en HGX. Réservé au
data-parallel simple ou au single-GPU.

**Q : Combien de temps pour rentabiliser le DLC ?**
R : Sur 100 kW IT : 40 kW économisés (PUE 1,5→1,1) = 52 k€/an. Surcoût DLC
~200–400 $/kW = 20–40 k€. **ROI < 1 an** en climat tempéré avec free cooling.

**Q : Le 800VDC est-il obligatoire pour Rubin ?**
R : Non — c'est une option sidecar. Mais les racks > 100 kW y gagnent 3–5 % de
rendement. À prévoir dans les salles neuves (chemins de câbles adaptés).

**Q : AMD ROCm supporte-t-il les containers Docker ?**
R : Oui — `rocm/pytorch` et `rocm/vllm` sur Docker Hub, avec `--device=/dev/kfd
--device=/dev/dri`. Vérifier le tag gfx (942/950) selon le GPU.

## 210. Références et documents à garder sous la main

- Fiches produits NVIDIA (nvidia.com — H100/H200/B200/L40S/RTX PRO).
- Fiches AMD Instinct (amd.com — MI300X/MI350X/MI355X).
- MLPerf results (mlcommons.org) — les seuls benchmarks indépendants.
- Spheron GPU pricing (spheron.network) — prix location live.
- Manuels d'installation OEM (Dell/HPE/Supermicro) — PDF à archiver avec le DOE.
- Ce guide (imprimer §173, §194, §68).
## 211. Modèle de PV de recette (extrait)

```
PROCÈS-VERBAL DE RECETTE — Cluster GPU
Date : __/__/____    Lieu : ______
Fournisseur : ______    Client : ______

Matériel réceptionné :
□ [___]× GPU [référence exacte] — n° série : ___
□ Serveur(s) [modèle] — n° série : ___
□ Switches [modèle] — n° série : ___
□ Onduleur [modèle] — n° série : ___

Tests (cf. §139) :
□ Burn-in 24 h : RÉUSSI / ÉCHEC (détails : ___)
□ NCCL all-reduce : ___ Go/s (attendu ≥ ___)
□ PCIe H→D : ___ Go/s (attendu ≥ ___)
□ Inférence 70B : ___ tok/s, 0 erreur / 10 k req : □ oui □ non
□ Bascule onduleur : □ oui □ non
□ DCGM : □ oui □ non

Réserves : ___
Réception : □ sans réserve □ avec réserves (levée sous ___ j)

Signatures : _______________ / _______________
```

## 212. KPI d'exploitation : les 12 à suivre

| KPI | Formule | Cible |
|---|---|---|
| Disponibilité GPU | 1 − (heures HS / heures totales) | > 99,5 % |
| Utilisation | Heures compute / heures totales | 60–80 % |
| PUE | Énergie totale / énergie IT | < 1,3 |
| tok/s/GPU | Tokens / (GPU × s) | Selon modèle |
| Coût/token | Coût total / tokens | < objectif |
| Latence p99 | Percentile 99 TTFT+TPOT | < SLA |
| Taux d'erreur | 5xx / requêtes | < 0,1 % |
| MTTR | Temps moyen de réparation | < 4 h |
| RMA/an | Cartes remplacées / parc | < 2 % |
| Efficacité énergétique | Tokens / kWh | En hausse |
| Taux de remplissage VRAM | VRAM utilisée / totale | 70–90 % |
| Satisfaction utilisateurs | Enquête trimestrielle | > 4/5 |

## 213. Antisèche : les formules du guide (toutes en une page)

```
VRAM_poids     = N_params × octets/param × 1,15
KV_cache       = 2 × couches × têtes_kv × dim × tokens × octets
VRAM_inf       = VRAM_poids + KV_cache + 3 Go
VRAM_train     ≈ N_params × 18 octets (Adam, mixed)
Temps_train    = 6 × N × D / (GPU × FLOPS × MFU)
Coût/an élec   = P_IT × PUE × 8760 × prix_kWh
I_tri (A)      = P_W / (692)   [400 V, cos φ 0,9]
I_mono (A)     = P_W / (207)   [230 V, cos φ 0,9]
Seuil_achat    = prix_achat / (prix_loc × 0,7 × 730)
TCO_3ans       = serveur + 3×(élec + 5% maintenance) + réseau
```

## 214. Dernier tableau : ce guide en 20 lignes

