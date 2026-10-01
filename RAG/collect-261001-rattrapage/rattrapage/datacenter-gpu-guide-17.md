---
id: collect-261001-rattrapage/rattrapage/datacenter-gpu-guide-17
title: "Les GPU datacenter / IA — présent et futur vérifié"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "AWS", "CoreWeave", "DeepSeek", "Lambda", "Nvidia", "SGLang", "TensorRT-LLM", "vLLM"]
dates: []
keywords: ["gpu", "amd", "attention", "aws", "benchmarks", "blackwell", "deepseek", "distribution", "ethernet", "fine-tuning", "fp8", "gqa"]
source: docs/RAG/collect-261001-rattrapage/datacenter_gpu_guide.md
source_anchor: ""
source_lines: [2454, 2607]
sha256: 9c48c031d54ae9bbefafc2a4bd8b47c6ce8926823c3260a398ddd7445ab35f83
---

# Les GPU datacenter / IA — présent et futur vérifié

Un rack 8U HGX pèse 150–250 kg et fait 1 200 mm de profondeur. Vérifier :
largeur des portes, ascenseur (charge), dalle (portance 500 kg/m² min),
rayon de braquage dans les couloirs. Classique des installations en sous-sol.

### Piège n°35 : le contrat de maintenance qui ne couvre pas le GPU

Certains contrats serveurs excluent « les accélérateurs » ou les limitent à
1 an. Lire les exclusions — un H100 hors garantie = 30 k$ à remplacer.

### Piège n°36 : les benchmarks vendeurs sur vos données

Un bench vendeur sur Llama 2 70B ne prédit pas vos perfs sur votre modèle
fine-tuné en français avec vos prompts. Exiger un **POC sur vos workloads**
(2–4 semaines, matériel prêté ou cloud équivalent) avant signature.

### Piège n°37 : l'obsolescence du dataset

Un cluster dimensionné pour des contextes 8K en 2024 est sous-dimensionné pour
du 128K en 2026 (KV cache ×16). Prévoir 30 % de marge VRAM pour l'évolution des
usages — ou un plan d'extension (slots, PDU, froid) dès le jour 1.

### Piège n°38 : oublier le backup des modèles fine-tunés

Un fine-tuning à 25 k$ dont les poids ne sont que sur le NVMe du nœud :
un disque mort = 25 k$ à refaire. Versionner les modèles (registry MLflow/
Weights & Biases + stockage répliqué), comme du code.

### Piège n°39 : la latence réseau applicative

Le GPU génère à 100 tok/s, mais l'API répond en 3 s : le goulot est le réseau
applicatif (TLS, proxy, sérialisation), pas le GPU. Mesurer de bout en bout
avant d'acheter plus de GPU.

### Piège n°40 : le « on fera du multi-cloud GPU » sans abstraction

Chaque cloud a ses images, ses drivers, ses quotas. Sans abstraction (Kubernetes +
opérateurs GPU, ou une couche type SkyPilot), le multi-cloud GPU = 3× le travail
d'exploitation. Choisir : un cloud principal + abstraction, ou un seul.

## 122. Glossaire étendu (15 termes de plus)

36. **GQA** : Grouped-Query Attention — divise le KV cache par 4–8.
37. **MLA** : Multi-head Latent Attention (DeepSeek) — KV cache compressé appris.
38. **PagedAttention** (vLLM) : gestion du KV cache par pages, anti-fragmentation.
39. **RadixAttention** (SGLang) : partage du cache des préfixes entre requêtes.
40. **Sparsité 2:4** : 2 valeurs non-nulles sur 4 — le format du « sparse » NVIDIA.
41. **MFU** : Model FLOPs Utilization — rendement réel du training (30–60 %).
42. **All-reduce** : collective qui somme les gradients sur tous les GPU.
43. **Checkpointing** : sauvegarde périodique (poids + optimiseur) en training.
44. **Activation checkpointing** : recalcule les activations au lieu de les stocker.
45. **AITER** : kernels ROCm optimisés pour vLLM/SGLang sur AMD.
46. **UEC** : Ultra Ethernet Consortium — l'Ethernet pour le training.
47. **Spectrum-X** : pile Ethernet IA de NVIDIA (alternative à IB).
48. **CDU** : Coolant Distribution Unit — le cœur hydraulique du DLC.
49. **Dry cooler** : refroidisseur à air sec (pas d'eau consommée).
50. **PUE partiel (pPUE)** : PUE d'une zone (ex. salle GPU seule).

## 123. Quiz — 10 questions de plus + réponses

**Q11. Un 70B FP16 + 15 % overhead tient-il sur un H200 (141 Go) ?**
R : Non : 161 Go > 141 Go. Il faut le H200 NVL (282 Go) ou du FP8 (90 Go).

**Q12. Quelle économie d'énergie entre PUE 1,5 et 1,1 sur 100 kW IT ?**
R : 40 kW × 8 760 h = 350 MWh/an, soit ~52 500 €/an à 0,15 €/kWh.

**Q13. Pourquoi 2 GPU PCIe sans NVLink sont-ils mauvais en tensor-parallel ?**
R : Le PCIe (64–128 Go/s) est 7–14× plus lent que NVLink (900 Go/s) — chaque
couche attend l'all-reduce, la latence explose.

**Q14. Que faire devant un Xid 79 répété ?**
R : Vérifier alim et connecteur 16-pin d'abord (80 % des cas), reseat la carte,
puis seulement envisager le RMA.

**Q15. Combien de temps pour rentabiliser un serveur 8× H100 acheté vs loué ?**
R : ~15 mois à utilisation 24/7 (285 k$ vs 3,33 $/GPU/h). En dessous de 60 %
d'utilisation : la location gagne.

**Q16. Quelle est la différence entre FP8 E4M3 et E5M2 ?**
R : E4M3 = 4 bits d'exposant, plus de précision (inférence) ; E5M2 = 5 bits
d'exposant, plus de dynamique (gradients en training).

**Q17. Pourquoi le KV cache croît-il avec le contexte ?**
R : Chaque token généré stocke ses clés/valeurs d'attention pour tous les tokens
précédents — mémoire proportionnelle au nombre de tokens.

**Q18. Un rack Rubin NVL72 consomme quel ordre de grandeur ?**
R : 120–200 kW estimés (72 GPU + 36 CPU + réseau, 100 % liquide).

**Q19. Que signifie « 30 % de tokens/$ de plus » (claim AMD Helios vs Rubin) ?**
R : À budget égal, 30 % de tokens générés en plus — chiffre constructeur AMD,
à valider en POC sur vos workloads.

**Q20. Quelle est la première chose à vérifier sur un devis GPU ?**
R : La référence exacte du GPU (modèle + format SXM/PCIe/OAM) et le détail
ligne par ligne (GPU, réseau, bridges, garantie) — pas de forfait global.

## 124. Planning projet type : du besoin au premier token

| Phase | Durée | Jalons |
|---|---|---|
| Expression du besoin + dimensionnement | 3–4 sem. | Note de calcul VRAM/énergie validée |
| Étude électrique + froid | 4–6 sem. | Devis TGBT/onduleur/DLC |
| Appel d'offres OEM (3 mini) | 4–6 sem. | Devis comparés ligne à ligne |
| Commande | 1 sem. | Bon de commande + clause délai |
| Travaux salle (élec, froid, réseau) | 8–16 sem. | Salle réceptionnée (en parallèle des délais GPU) |
| Livraison + installation | 2–4 sem. | Serveurs rackés, câblés |
| Mise en service + burn-in | 2 sem. | Checklist §84 validée |
| POC workloads | 2–4 sem. | Perfs mesurées vs objectifs |
| Production | — | Runbook §84 + monitoring §107 |

**Durée totale typique : 6 à 9 mois** de la décision au premier token en prod.
Le chemin critique est presque toujours l'électrique, pas le GPU.

## 125. Contacts et ressources (points d'entrée)

| Besoin | Point d'entrée |
|---|---|
| Devis serveurs NVIDIA | Commerciaux Dell/HPE/Supermicro/Lenovo (demander le « AI team ») |
| Devis AMD | Mêmes OEM, demander la « plateforme UBB Instinct » |
| Cartes PCIe pro | PNY (distribution), revendeurs agréés |
| Location ponctuelle | Runpod, Vast.ai, Lambda Labs |
| Location entreprise | CoreWeave, AWS, Azure, GCP |
| Benchmarks indépendants | MLPerf (mlcommons.org), Artificial Analysis |
| Prix marché gris (référence) | eBay « sold listings », Jawa |

## 126. Sigles constructeurs (pense-bête)

- **NVIDIA** : HGX (baseboard), DGX (serveur clé en main), NGC (catalogue),
  DCGM (monitoring), NCCL (collectives), MIG (partitionnement).
- **AMD** : UBB (baseboard), ROCm (stack), RCCL (collectives), AITER (kernels),
  xGMI (liens), UALink (scale-up ouvert).
- **OEM** : iDRAC (Dell), iLO (HPE), XCC (Lenovo), SMCIPMITool (Supermicro).
- **Réseau** : IB (InfiniBand), NDR/XDR (générations), RoCE (RDMA sur Ethernet),
  UEC (Ultra Ethernet Consortium).
## 127. Formats numériques basse précision : le détail

| Format | Bits (signe/exp/mantisse) | Dynamique | Précision | Support |
|---|---|---|---|---|
| BF16 | 1/8/7 | Grande | Moyenne | Tous (2018+) |
| FP16 | 1/5/10 | Moyenne | Bonne | Tous |
| FP8 E4M3 | 1/4/3 | Moyenne | Moyenne | Hopper+, CDNA 4+ |
| FP8 E5M2 | 1/5/2 | Grande | Faible | Hopper+ |
| FP8 FNUZ (AMD) | 1/4/3 (variante) | Moyenne | Moyenne | CDNA 3 (MI300X) |
| MXFP8 (OCP) | Échelle partagée ×32 | Grande | Moyenne | Blackwell, CDNA 4 |
| MXFP6 E3M2/E2M3 | 1/3/2 ou 1/2/3 | Moyenne | Faible | Blackwell, CDNA 4 |
| MXFP4 E2M1 | 1/2/1 | Faible | Très faible | Blackwell, CDNA 4 |
| NVFP4 (NVIDIA) | 1/2/1 + échelle fine | Faible | Très faible | Blackwell, Rubin |
| INT8 | Entier signé | Fixe | Moyenne | Tous (Tensor) |
| INT4 | Entier signé | Fixe | Faible | Tous (via quantif.) |

**Règle** : E4M3 pour l'inférence (précision), E5M2 pour les gradients (dynamique).
MXFP4/MXFP6 exigent un rescaling par blocs — le framework doit le supporter
(vLLM/SGLang/TensorRT-LLM : oui en 2026).

