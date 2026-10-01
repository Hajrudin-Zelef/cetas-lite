---
id: collect-261001-rattrapage/rattrapage/datacenter-gpu-guide-20
title: "Les GPU datacenter / IA — présent et futur vérifié"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Anthropic", "Broadcom", "CoreWeave", "Crusoe", "DeepSeek", "Fireworks AI", "Google", "Intel", "Lambda", "Meta", "Microsoft", "Nebius", "Nvidia", "OpenAI", "Oracle", "SGLang", "TSMC", "Together AI", "vLLM"]
dates: []
keywords: ["gpu", "amd", "apache", "attention", "capex", "compute", "deepseek", "ethernet", "fp4", "fp8", "hbm4", "helios"]
source: docs/RAG/collect-261001-rattrapage/datacenter_gpu_guide.md
source_anchor: ""
source_lines: [2924, 3091]
sha256: f41cbacd15f7ab7220915fb58bb1699210b9d39c52e3852a69921d6a441a1f32
---

# Les GPU datacenter / IA — présent et futur vérifié

| Composant | Prix indicatif | Note |
|---|---|---|
| AI Enterprise (par GPU/an) | ~3 500–4 500 $ | Support + vGPU + Base Command |
| Base Command Manager | Inclus DGX | Orchestration |
| NeMo / NIM | Inclus (licence d'usage) | Microservices d'inférence |
| vGPU (vPC/vWS) | ~250–450 $/GPU/an | Virtualisation poste |

Sur 8 GPU × 3 ans : **84 000–108 000 $** de licences — 20–25 % du prix du serveur.
Alternative : stack open-source (vLLM + Kubernetes) = 0 $ de licence, 1 ETP MLOps.

## 152. DGX Cloud vs on-prem : le calcul

- **DGX Cloud** : du H100/B200 « managé » chez les partenaires (Oracle, Lambda…),
  ~4–6 $/GPU/h tout compris.
- **Break-even vs achat** : ~18–24 mois à 80 % d'utilisation.
- **Intérêt** : zéro capex, scaling instantané, pas d'équipe infra.
- **Limite** : 2× le coût à 3 ans en 24/7, données chez un tiers, disponibilité
  par quotas.

## 153. Slurm : l'ordonnanceur des clusters GPU

- **Standard** du HPC et du training : file d'attente, allocation exclusive,
  comptabilité par projet.
- **Config GPU** : `Gres=gpu:h100:8`, partitions par génération (pas de mélange).
- **Alternative cloud-native** : Kubernetes + NVIDIA GPU Operator (plus souple,
  moins efficace en all-reduce pur).
- **Règle** : training > 32 GPU → Slurm ; inférence multi-tenant → Kubernetes.

## 154. Kubernetes + GPU Operator

- **GPU Operator** : déploie drivers, DCGM, device plugin en DaemonSet.
- **Time-slicing** : partage logiciel d'un GPU (dev/test, pas de prod facturée).
- **MIG + K8s** : 1 instance MIG = 1 ressource `nvidia.com/mig-1g.5gb`.
- **Limite** : le réseau IB en K8s (SR-IOV, Macvlan) reste plus fragile qu'en
  Slurm/bare metal pour le training.

## 155. Capacity planning : modéliser la croissance

```
GPU_requis(t) = trafic(t) × tokens/requête / (débit/GPU × utilisation_cible)
```

- **Trafic** : modéliser en requêtes/heure pointe (pas en moyenne).
- **Utilisation cible** : 60 % (marge burst + maintenance).
- **Exemple** : 10 000 req/h × 2 000 tokens = 20 M tokens/h.
  À 200 tok/s/GPU (batché) × 0,6 = 120 tok/s utiles → 46 GPU → **6 nœuds 8×**.
- **Réviser trimestriellement** : les usages IA doublent tous les 6–12 mois.
## 156. Deep dive : CDNA 4 (MI350X/MI355X)

- **Refonte des matrix cores** : support natif MXFP4/MXFP6/MXFP8 (format OCP),
  là où CDNA 3 plafonnait au FP8 FNUZ.
- **Gravure** : TSMC N3 pour les XCD (vs N5 sur CDNA 3) — +30 % d'efficacité.
- **256 CU** (vs 304 sur MI300X) mais des CU plus larges en débit matriciel.
- **8 To/s** : la HBM3e 12-high qui aligne AMD sur le B200 en bande passante.
- **Différence 350X/355X** : le 355X est le même silicium avec TDP 1 400 W
  (liquide) qui libère tout le débit FP4 — le 350X air est bridé thermiquement
  sur les pics FP4 longs.

## 157. ROCm : état des lieux 2026

| Version | Apport majeur |
|---|---|
| ROCm 6.0–6.2 | Support MI300X stable, PyTorch day-0 Llama 3 |
| ROCm 6.3–6.4 | FP8 OCP, kernels FlashAttention-3 |
| ROCm 7.x (2026) | gfx950 (CDNA 4), AITER intégré, vLLM parité ~95 % |

- **Points forts** : open-source (MIT/Apache), pas de licence, support CPU+GPU.
- **Points faibles** : les nouveaux modèles mettent 2–8 semaines à être optimisés,
  la doc reste en retrait de CUDA, le support Windows est anecdotique.
- **Verdict 2026** : production-ready pour vLLM/SGLang sur les modèles mainstream,
  pas encore pour la R&D kernels.

## 158. AITER : l'arme secrète d'AMD

- **AITER** (AI Tensor Engine for ROCm) : bibliothèque de kernels fusionnés
  (MoE, MLA, FlashAttention) optimisés pour CDNA.
- **ATOM** : son successeur/complément pour les opérateurs d'attention.
- **Activation** : variables d'environnement vLLM (`VLLM_ROCM_USE_AITER=1`).
- **Gain mesuré** : +20 à +40 % de tok/s sur les MoE (DeepSeek) vs kernels
  génériques — c'est ce qui rend le MI300X compétitif face au H100.

## 159. Deep dive : Helios (rack AMD MI450)

- **72× MI455X**, 31 To HBM4, 1,4 Po/s agrégés, 2,9 exaflops FP4.
- **UALink** : 260 To/s de scale-up ouvert (vs NVLink propriétaire).
- **Scale-out** : 43 To/s Ethernet UEC via Pensando Vulcano.
- **CPU** : EPYC Venice (6e gén) — 6 750 000 unités prévues 2027 (Morgan Stanley).
- **Électrique** : design Schneider **246 kW par rack**, 100 % liquide.
- **Clients** : Anthropic (2 GW), OpenAI (6 GW, 3 mois de prod), Meta, Oracle.
- **Claim** : +30 % tokens/$ vs Rubin NVL72 (chiffre AMD, à valider en POC).

## 160. Comparatif rack-à-rack : NVL72 Rubin vs Helios

| Critère | Rubin NVL72 (NVIDIA) | Helios (AMD) |
|---|---|---|
| GPU | 72× Rubin | 72× MI455X |
| Mémoire totale | 20,7 To HBM4 | 31 To HBM4 |
| Compute FP4 | 3,6 exaflops NVFP4 | 2,9 exaflops FP4 |
| Scale-up | NVLink 6 (propriétaire) | UALink (ouvert) |
| Scale-out | ConnectX-9 / Spectrum-6 | UEC / Vulcano |
| Puissance rack | ~120–200 kW (est.) | 246 kW (design Schneider) |
| Refroidissement | Liquide sec, 45 °C inlet | Liquide, quick-disconnect |
| Clients | OpenAI, CoreWeave, Azure… | Anthropic, OpenAI, Meta… |
| Disponibilité | H2 2026 (ramp Q4) | Fin Q3 2026 (ramp 2027) |

Deux philosophies : NVIDIA = intégration verticale propriétaire, AMD = standards
ouverts. Le TCO se jouera sur le $/token réel en 2027, pas sur les claims 2026.

## 161. Le pari UALink : enjeux

- **Membres** : AMD, Broadcom, Cisco, Google, HPE, Intel, Meta, Microsoft…
  (tous sauf NVIDIA).
- **Objectif** : un NVLink ouvert — 200 GT/s par lane, scale-up jusqu'à 1 024 GPU.
- **Risque** : standard jeune (1.0 en 2025), interopérabilité réelle à prouver
  en 2026-2027.
- **Enjeu pour Zelef** : si UALink perce, les clusters multi-fournisseurs
  deviennent possibles — fin du lock-in interconnect. À suivre, pas à parier
  un achat 2026 dessus.

## 162. Poids du logiciel dans le choix AMD vs NVIDIA

| Critère | NVIDIA | AMD |
|---|---|---|
| Maturité kernels | 10 ans d'avance | Rattrapage (AITER) |
| Nouveaux modèles | Day-0 | Day-0 à +8 semaines |
| Outils profilage | Nsight (excellent) | ROCm profiler (correct) |
| Communauté | Immense | Croissante |
| Coût licence | $$ (AI Enterprise) | 0 $ |
| Risque projet | Faible | Moyen (prévoir POC) |

**Décision** : si l'équipe est junior ou pressée → NVIDIA. Si l'équipe est
senior Linux et le budget serré → AMD avec POC de 4 semaines minimum.
## 163. Prix historiques : la courbe du H100 (2023–2026)

| Période | Prix carte (indicatif) | Contexte |
|---|---|---|
| T1 2023 | 35 000–40 000 $ | Pénurie, files d'attente 6 mois |
| T4 2023 | 30 000–35 000 $ | Production à plein régime |
| T2 2024 | 25 000–30 000 $ | Offre normalisée |
| T1 2025 | 25 000–31 000 $ | Arrivée B200, stabilisation |
| T3 2026 | 20 000–28 000 $ (négocié) | Fin de cycle, Rubin en prod |

**Leçon** : acheter au pic (2023) = -40 % de valeur en 3 ans. Le timing d'achat
est un levier aussi puissant que la négociation.

## 164. Les néoclouds GPU : qui sont-ils

| Acteur | Modèle | Prix H100 |
|---|---|---|
| CoreWeave | Contrats entreprise, IPO 2025 | ~3,50–6 $/h |
| Lambda Labs | Dev-friendly | ~2,90 $/h |
| Together AI | Inférence serverless | Au token |
| Fireworks AI | Inférence serverless | Au token |
| Nebius | Ex-Yandex, Europe | ~2–3 $/h |
| Crusoe | Énergie + compute | Contrats |

Les néoclouds ont capté ~30 % du marché de la location GPU en 2025-2026 en
cassant les prix des hyperscalers (-30 à -50 %).

## 165. Inférence serverless : le $/token en 2026

| Provider | 70B $/M tokens (in) | 70B $/M tokens (out) |
|---|---|---|
| Hyperscaler premium | ~2,50 $ | ~2,50 $ |
| Néocloud | ~0,90 $ | ~0,90 $ |
| Self-hosted B200 (amorti) | ~4,60 $ | ~4,60 $ |
| Self-hosted RTX PRO 6000 | ~2,80 $ | ~2,80 $ |

