---
id: collect-261001-rattrapage/rattrapage/datacenter-gpu-guide-1
title: "Les GPU datacenter / IA — présent et futur vérifié"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "AWS", "CoreWeave", "Nvidia", "TSMC", "United States"]
dates: ["2026-09-27"]
keywords: ["datacenter", "gpu", "amd", "arr", "awq", "aws", "blackwell", "distribution", "fine-tuning", "fp4", "fp8", "gguf"]
source: docs/RAG/collect-261001-rattrapage/datacenter_gpu_guide.md
source_anchor: ""
source_lines: [1, 140]
sha256: 34f82f14966b8b79edb6a70e0a519fb45b0353d30855a26397a57a4309e47e41
---

# Les GPU datacenter / IA — présent et futur vérifié

**Guide technique pour Zelef — Rédigé et vérifié le 27/09/2026**

> Chiffres vérifiés le 27/09/2026 via recherche web. Les prix datacenter sont indicatifs et volatils
> (marqués « à vérifier »). Référence introuvable = « non trouvée au 27/09/2026 » — jamais d'invention.

## Sommaire

- [1. Lecture du guide](#1-lecture-du-guide)
- [2. RTX 4090](#2-nvidia-geforce-rtx-4090)
- [3. RTX 5090](#3-nvidia-geforce-rtx-5090)
- [4. L40S](#4-nvidia-l40s)
- [5. RTX PRO 6000 Blackwell](#5-nvidia-rtx-pro-6000-blackwell)
- [6. H100](#6-nvidia-h100)
- [7. H200](#7-nvidia-h200)
- [8. B200](#8-nvidia-b200)
- [9. MI300X](#9-amd-instinct-mi300x)
- [10. MI325X](#10-amd-instinct-mi325x)
- [11. MI350X / MI355X](#11-amd-instinct-mi350x-et-mi355x)
- [12. MI300A](#12-amd-instinct-mi300a)
- [13. Tableau : mémoire](#13-tableau-comparatif-memoire)
- [14. Tableau : TDP / perf](#14-tableau-comparatif-tdp-performance)
- [15. Tableau : prix / perf](#15-tableau-comparatif-prix-performance)
- [16. Tableau : cas d'usage](#16-tableau-comparatif-par-cas-dusage)
- [17. Formats SXM vs PCIe vs OAM/UBB](#17-formats-physiques-sxm-vs-pcie-vs-oamubb)
- [18. NVLink / NVSwitch](#18-nvlink-et-nvswitch)
- [19. Infinity Fabric / xGMI](#19-infinity-fabric-et-xgmi-cote-amd)
- [20. Serveurs 4/8 GPU](#20-serveurs-48-gpu-hgx-dgx-et-equivalents-amd)
- [21. DLSS / FSR / XeSS](#21-dlss--fsr--xess-section-courte)
- [22. Dimensionnement VRAM](#22-dimensionnement-vram-pour-les-llm)
- [23. KV cache](#23-le-kv-cache-le-second-compteur)
- [24. Exemples chiffrés 7B→405B](#24-exemples-chiffres-de-7b-a-405b)
- [25. Contexte 128K+](#25-impact-du-contexte-long-128k-1m-tokens)
- [26. Refroidissement air vs liquide](#26-refroidissement-air-vs-liquide-par-modele)
- [27. TDP réels](#27-tdp-reels-mesures-vs-spec-sheet)
- [28. Conso 8×GPU](#28-consommation-par-config-8xgpu)
- [29. Énergie et coûts](#29-energie-pue-et-cout-electrique)
- [30. Pièges 1-5](#30-pieges-terrain-15-erreurs-reelles)
- [31. Pièges 6-10](#31-pieges-terrain-suite)
- [32. Pièges 11-15+](#32-pieges-terrain-fin)
- [33. Où acheter](#33-ou-acheter-circuits-dachat)
- [34. BOM 4×H200](#34-bom-type-serveur-4xh200-pcie)
- [35. BOM 8×RTX PRO 6000](#35-bom-type-serveur-8xrtx-pro-6000)
- [36. BOM AMD MI350X](#36-bom-type-serveur-amd-8xmi350x)
- [37. Logiciels](#37-pile-logicielle-cuda-rocm-frameworks)
- [38. Cloud vs achat](#38-rent-vs-buy-seuil-de-rentabilite)
- [39. À venir NVIDIA](#39-a-venir-roadmap-nvidia-verifiee)
- [40. À venir AMD](#40-a-venir-roadmap-amd-verifiee)
- [41. Rumeurs](#41-rumeurs-non-confirmees-marquees-rumeur)
- [42. Glossaire](#42-glossaire-35-termes)
- [43. Quiz](#43-quiz-10-questions-reponses)
- [44. Sources](#44-sources)

---

## 1. Lecture du guide

Ce guide fait 100+ sections numérotées, mais il se lit en diagonale :

- **Sections 2 à 12** : fiches techniques de chaque GPU (architecture, mémoire, TDP, format,
  précisions supportées, cas d'usage, prix indicatifs vérifiés le 27/09/2026, où acheter).
- **Sections 13 à 20** : comparatifs et infrastructures (formats, NVLink, serveurs).
- **Sections 22 à 29** : dimensionnement VRAM et énergie — le cœur du métier de Zelef.
- **Sections 30 à 32** : 15+ pièges terrain vérifiés.
- **Sections 39 à 41** : roadmaps officielles + rumeurs explicitement marquées.

Conventions : TDP = Thermal Design Power (puissance dissipée thermique), TBW = bande passante
mémoire, PFLOPS = 10^15 opérations flottantes par seconde. Prix en USD sauf mention contraire,
« à vérifier » quand le prix n'a pas pu être recoupé le 27/09/2026.

---

## 2. NVIDIA GeForce RTX 4090

- **Référence exacte** : NVIDIA GeForce RTX 4090 24GB GDDR6X (vérifiée le 27/09/2026).
- **Architecture** : Ada Lovelace (GPU AD102), gravé TSMC 4N (classe 5 nm), 76,3 milliards de transistors.
- **Cœurs** : 16 384 CUDA, 512 Tensor (4e génération), 128 RT (3e génération).
- **Mémoire** : 24 Go GDDR6X, bus 384-bit, **1 008 Go/s** de bande passante, sans ECC.
- **TDP** : **450 W** (connecteur 16-pin PCIe 5.0 / 12VHPWR, adaptateur fourni).
- **Format** : PCIe 4.0 x16, carte 3 à 3,5 slots selon AIB, longueur ~336 mm.
- **Précisions IA** : FP32 ~82,6 TFLOPS ; FP16 Tensor ~330 TFLOPS dense ; FP8 Tensor supporté
  (4e gén. Tensor Cores) ; INT8 ~1 321 AI TOPS (référence Newegg/2026 : 1 321 AI TOPS).
- **Cas d'usage** : inférence locale, fine-tuning léger (LoRA/QLoRA), développement, labs.
  Pas d'ECC, pas de garantie datacenter : usage « workstation/lab », pas production critique.
- **Prix** : MSRP lancement 1 599 $ (oct. 2022). Au 27/09/2026 : **1 800 à 3 500 $** selon modèle
  (bande 2 400–2 800 $ la plus fréquente chez les retailers US) ; Amazon liste des cartes
  jusqu'à ~4 390–4 840 $ (revendeurs). Fin de vie commerciale, production arrêtée.
- **Où l'acheter** : retailers grand public (Amazon, Newegg), marché secondaire (eBay, Jawa).
  Pas de canal OEM datacenter (carte gaming).

---

## 3. NVIDIA GeForce RTX 5090

- **Référence exacte** : NVIDIA GeForce RTX 5090 32GB GDDR7 (vérifiée le 27/09/2026).
- **Architecture** : Blackwell (GPU GB202), TSMC 4N, 92,2 milliards de transistors.
- **Cœurs** : 21 760 CUDA, 680 Tensor (5e génération), 136–170 RT (4e génération selon source).
- **Mémoire** : **32 Go GDDR7**, bus 512-bit, **1 792 Go/s** (+78 % vs 4090), sans ECC.
- **TDP** : **575 W** (connecteur 16-pin 12V-2×6). PSU recommandée : 1 000 W minimum.
- **Format** : PCIe 5.0 x16, carte ~3,6 slots, Founders Edition double-flow-through.
- **Précisions IA** : FP32 ~104,8 TFLOPS ; FP16 Tensor dense ~209,5 TFLOPS ;
  FP8 Tensor dense ~419 TFLOPS ; INT8 dense ~838 TOPS ; **FP4 sparse ~3 352 TOPS**
  (données Spheron, vérifiées le 27/09/2026). DLSS 4 Multi Frame Generation (jusqu'à 3 images
  générées par image rendue).
- **Cas d'usage** : inférence locale 7B–32B en quantifié, fine-tuning QLoRA, rendu,
  postes de dev IA. Le ratio €/Go VRAM gaming le rend populaire en lab.
- **Prix** : MSRP 1 999 $ (janv. 2025). Rue au 27/09/2026 : **2 200–2 800 $** typique,
  jusqu'à **4 299 $** chez Walmart (promo GeForce Week, épuisé) et 3 899–6 000 $ chez les
  revendeurs (Jawa). Hausse +145 % vs MSRP constatée début sept. 2026.
- **Où l'acheter** : retailers grand public. Serveurs 8×RTX 5090 proposés par des intégrateurs
  chinois (~29 500 $ le 8U, source Accio, à vérifier).

---

## 4. NVIDIA L40S

- **Référence exacte** : NVIDIA L40S 48GB GDDR6 (réf. commerciale **900-2G133-0080-000**,
  vérifiée le 27/09/2026).
- **Architecture** : Ada Lovelace (AD102 complet, 18 176 CUDA), TSMC 4N.
- **Cœurs** : 18 176 CUDA, 568 Tensor (4e gén.), 142 RT (3e gén.).
- **Mémoire** : **48 Go GDDR6 avec ECC**, bus 384-bit, **864 Go/s**.
- **TDP** : **350 W** (1× 16-pin). Le TDP le plus sobre de cette liste côté pro.
- **Format** : PCIe 4.0 x16, double-slot, version **passive** (datacenter) ou active selon OEM.
  Pas de sorties vidéo utilisables en datacenter (4× DP 1.4a désactivées par défaut).
- **Précisions IA** : FP32 91,6 TFLOPS ; FP16/BF16 Tensor dense ~366 TFLOPS (sparse ~733) ;
  FP8 Tensor dense ~733 TFLOPS (sparse ~1 466) ; INT8 dense ~733 TOPS (sparse ~1 466).
  Quantizations supportées : GPTQ, AWQ, GGUF, FP8.
- **Cas d'usage** : **inférence** (le roi du $/token en datacenter 2024-2025), graphisme
  professionnel, rendu, VDI. Pas NVLink : scale-up limité au PCIe.
- **Prix** : neuf **9 000–11 000 $** (référence GitHub gpu_specs, mi-2026) ; PNY listé à
  10 819 € (axitech.be, sept. 2026) ; reconditionné Dell ~8 999 $ (eBay). Location cloud :
  ~0,96 $/h dédié (Spheron), ~1,86–1,89 $/h chez AWS/CoreWeave.
- **Où l'acheter** : PNY (distribution), Dell, HPE, Lenovo, Supermicro, revendeurs
  (Marigold Systems ~12 999 $ neuf). Carte « universelle » : 4× MIG possibles selon NVIDIA.

---

## 5. NVIDIA RTX PRO 6000 Blackwell

