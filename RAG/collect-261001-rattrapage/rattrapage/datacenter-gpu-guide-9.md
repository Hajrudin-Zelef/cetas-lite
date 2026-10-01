---
id: collect-261001-rattrapage/rattrapage/datacenter-gpu-guide-9
title: "Les GPU datacenter / IA — présent et futur vérifié"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia", "SGLang", "TensorRT-LLM", "vLLM"]
dates: ["2026-08-18", "2026-09-27"]
keywords: ["datacenter", "gpu", "agents", "agi", "amd", "attention", "awq", "blackwell", "compute", "distribution", "fp4", "fp8"]
source: docs/RAG/collect-261001-rattrapage/datacenter_gpu_guide.md
source_anchor: ""
source_lines: [1145, 1276]
sha256: b2bd396a04595fd3862c34463ffe1a377cfbe5514951c42f58a165842e134fa2
---

# Les GPU datacenter / IA — présent et futur vérifié

1. **HBM** (High Bandwidth Memory) : mémoire 3D empilée sur interposeur, 3–20 To/s.
2. **HBM3 / HBM3e** : générations 2022/2024 (e = enhanced). HBM4 arrive avec Rubin/MI450.
3. **GDDR** : mémoire graphique classique (GDDR6/6X/7), moins chère que HBM, ~1–1,8 To/s.
4. **ECC** : correction d'erreurs mémoire. Obligatoire en datacenter.
5. **TDP** : puissance thermique de design — dimensionne le refroidissement.
6. **TBP/TGP** : puissance totale de la carte (proche du TDP).
7. **Tensor Core** : unités matricielles NVIDIA (FP16/FP8/FP4/INT8).
8. **Matrix Core** : équivalent AMD (CDNA).
9. **XMX** : unités matricielles Intel (Arc).
10. **FP8** : flottant 8 bits — le standard training/inférence 2024-2026.
11. **FP4** : flottant 4 bits — inférence Blackwell/CDNA 4+, qualité à valider par modèle.
12. **NVFP4** : format FP4 propriétaire NVIDIA (Rubin).
13. **MXFP4/6/8** : formats microscaling OCP, supportés par AMD CDNA 4.
14. **FNUZ** : format FP8 AMD pré-OCP (MI300X/MI325X) — incompatibilité logicielle avec OCP FP8.
15. **Quantification** : réduction de précision (GPTQ, AWQ, GGUF, FP8) pour tenir en VRAM.
16. **KV cache** : mémoire des clés/valeurs d'attention, croît avec le contexte.
17. **SXM** : format mezzanine NVIDIA (baseboard HGX).
18. **OAM** : format mezzanine ouvert OCP (AMD).
19. **UBB** : baseboard 8× OAM, équivalent ouvert du HGX.
20. **HGX** : baseboard NVIDIA 8× SXM + NVSwitch.
21. **DGX** : serveur clé en main NVIDIA (HGX + support + Base Command).
22. **NVLink** : interconnect GPU-à-GPU NVIDIA (900 Go/s gén. 4, 1,8 To/s gén. 5).
23. **NVSwitch** : commutateur tout-à-tout 8 GPU.
24. **xGMI** : liens directs GPU-à-GPU AMD.
25. **UALink** : standard ouvert de scale-up (AMD Helios, 260 To/s).
26. **NCCL / RCCL** : librairies de collectives multi-GPU (NVIDIA / AMD).
27. **MIG** : Multi-Instance GPU — découpage d'un GPU en instances isolées.
28. **PCIe 5.0 x16** : 128 Go/s — le goulot host↔GPU.
29. **PUE** : Power Usage Effectiveness — 1,1 (liquide) à 1,6 (air).
30. **DLC** : Direct Liquid Cooling — plaques froides sur les puces.
31. **CDU** : unité de distribution du liquide de refroidissement.
32. **vLLM** : serveur d'inférence LLM (PagedAttention), CUDA + ROCm.
33. **SGLang** : serveur d'inférence (RadixAttention), fort en agents.
34. **TensorRT-LLM** : optimiseur d'inférence NVIDIA, le plus rapide sur son silicium.
35. **MoE** (Mixture of Experts) : n'active qu'une fraction des paramètres par token.

## 43. Quiz — 10 questions + réponses

**Q1. Combien de VRAM faut-il pour les poids d'un Llama 70B en FP16, avec 15 % d'overhead ?**
R : 70 × 2 × 1,15 = **161 Go**.

**Q2. Pourquoi le H200 est-il meilleur que le H100 pour l'inférence 70B, à compute égal ?**
R : 141 Go HBM3e vs 80 Go — le 70B FP16 (161 Go avec overhead) tient sur **1 GPU**
au lieu de 2, supprimant le tensor-parallelisme et sa latence.

**Q3. Un serveur 8× B200 consomme combien au mur (ordre de grandeur) ?**
R : **~10,6 kW AC** (8 kW GPU + CPU/réseau/stockage + rendement PSU).

**Q4. Quelle est la différence entre SXM et PCIe pour un H100 ?**
R : SXM = module sur baseboard HGX, 700 W, NVLink natif ; PCIe = carte slot,
350 W (80 Go) ou 700 W (96 Go), pas de NVLink inter-GPU natif.

**Q5. Pourquoi une RTX 4090 est-elle déconseillée en training production ?**
R : Pas d'ECC (bitflips silencieux), pas de garantie datacenter 24/7, pas de NVLink.

**Q6. À 128K tokens, quel poste mémoire dépasse les poids sur un 70B ?**
R : Le **KV cache** (~160 Go en BF16 > 140 Go de poids FP16).

**Q7. Air ou liquide pour un rack 8× MI355X (11,2 kW GPU) ?**
R : **Liquide (DLC)** — 13,5 kW nœud complet, au-delà de l'air classique (15–20 kW/rack
serait théoriquement OK pour 1 nœud, mais 2 nœuds par rack = 27 kW → DLC).

**Q8. Que signifie « FP4 sparse 18 000 TOPS » sur un B200 ?**
R : Débit pic avec sparsité structurée 2:4 (une valeur sur deux forcée à zéro) —
le débit dense réel est 9 000 TFLOPS, et le débit applicatif encore en dessous.

**Q9. Quel est le piège d'acheter un H100 « pas cher » sur le marché gris ?**
R : Mauvais format (SXM vs PCIe), vBIOS non datacenter, carte rebadgée, pas de
garantie — vérifier Device ID via nvidia-smi et acheter via distribution agréée.

**Q10. Rubin apporte quoi vs Blackwell, en chiffres NVIDIA ?**
R : NVLink 6 (3,6 To/s/GPU), 50 PFLOPS NVFP4 par GPU, coût/token ÷10, 4× moins de
GPU pour entraîner un MoE (claims constructeur, GTC 2026).

## 44. Sources

Recherche web effectuée le 27/09/2026. Principales sources recoupées :

- Fiches produits NVIDIA (H100/H200/B200, L40S, RTX PRO 6000) et AMD (MI300X/MI325X/
  MI350X/MI355X) — specs constructeur.
- Spheron Network (prix location live au 27/09/2026 : H100/H200/B200/5090/L40S).
- GPUaaS.com — comparatif H100/H200/B200, prix août 2026.
- MillionMiner — AI Server Price Guide 2026 (prix serveurs 8×).
- DirectMacro — comparatif DGX B300/B200/H200/H100.
- Tech-insider.org — GB300 production, comparatif RTX PRO 6000/5090/W7900.
- SLYD.com — specs AMD Instinct (revue 18/08/2026).
- GitHub : kinvert/knowledge-base (BOM H200), biterik/lammps-compile-n-bench
  (prix mi-2026), vladislavduma/gpu_specs (tableau perfs), wty2003328/silicon-to-serving
  (architecture MI300), amd-agi/geak (tableau cartes AMD), redhat-et/physical-ai-platform-intel
  (deep-dive AMD, roadmap MI450/Helios).
- Fierce Network / SDxCentral / HPE — lancement AMD Helios + MI450 (Advancing AI 2026).
- Medium/CryptoBriefing/Securities.io/HotHardware — NVIDIA Vera Rubin GTC 2026.
- RunPod, ofzenandcomputing.com, ammarandi.com — RTX 4090/5090 specs et prix sept. 2026.
- VideoCardz / ThinkComputers — prix rue RTX 5090 (Walmart 4 299 $, sept. 2026).
- Axitech.be / Marigold Systems / Ahead-IT / eBay — prix L40S sept. 2026.
- Hot Aisle — prix cloud MI300X (1,99 $/h, juin 2025).
- Wccftech / Newegg / Tech-insider — DLSS 4/5 vs FSR 4 vs XeSS 2/3 (2026).

> Prix et disponibilités : photographie au 27/09/2026. Re-vérifier avant tout achat —
> le marché GPU bouge au trimestre, parfois à la semaine (pénurie GDDR7 constatée mi-2026).
## 45. Débits d'inférence mesurés par GPU (ordres de grandeur vLLM)

Conditions : batch 1, contexte 4K, génération 256 tokens, FP16 sauf mention.
Les chiffres varient ±30 % selon le moteur et la quantification — ce sont des
repères terrain, pas des specs.

| GPU | Llama 8B (tok/s) | Llama 70B Q4 (tok/s) | Llama 70B FP8 (tok/s) | 405B FP8 (tok/s, multi-GPU) |
|---|---|---|---|---|
| RTX 4090 | ~90–120 | — (tient pas : 40 Go > 24) | — | — |
| RTX 5090 | ~130–170 | ~35–45 | — (96 Go requis) | — |
| L40S | ~70–95 | ~25–35 | — | — |
| RTX PRO 6000 | ~120–160 | ~40–55 | ~45–60 | ~15–20 (8×, pipeline) |
| H100 | ~110–150 | ~45–60 | ~50–65 | ~25–35 (8×) |
| H200 | ~115–155 | ~50–65 | ~55–70 | ~30–40 (8×) |
| B200 | ~180–240 | ~80–110 | ~90–120 | ~50–70 (8×) |
| MI300X | ~120–160 | ~55–70 | ~60–75 | ~35–45 (8×) |
| MI355X | ~170–220 | ~85–110 | ~95–125 | ~55–75 (8×) |

**Lecture** : à batch 1, le débit suit la bande passante mémoire (B200 8 To/s ≈ 2,4×
H100 3,35 To/s → ~2× tok/s). À batch élevé, le compute prend le relais et l'écart
se creuse en faveur du FP8/FP4.

## 46. Retour terrain : 2× RTX 5090 en lab (Zelef-style)

- **Config** : 2× RTX 5090 (64 Go cumulés), Threadripper/EPYC, 128 Go RAM, 2 kW.
- **Ce qui tourne** : 70B Q4 (~40 Go) en tensor-parallel sur 2 GPU via vLLM,
  ~60–80 tok/s ; 32B FP16 (~70 Go avec overhead : non, 64 Go < 70 Go → Q4 ou FP8).
- **Coût** : ~5 200 $ GPU + ~3 000 $ système = ~8 200 $, ~150 €/an d'élec.
- **Limites** : pas d'ECC, pas de NVLink (PCIe P2P seulement), driver gaming
  (pas de persistence garantie au reboot — activer `nvidia-persistenced`).
- **Verdict** : parfait pour dev/test, inacceptable pour prod 24/7 critique.

