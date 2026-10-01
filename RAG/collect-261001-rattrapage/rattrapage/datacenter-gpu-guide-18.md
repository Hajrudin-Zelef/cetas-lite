---
id: collect-261001-rattrapage/rattrapage/datacenter-gpu-guide-18
title: "Les GPU datacenter / IA — présent et futur vérifié"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Alibaba", "DeepSeek", "Mistral", "Nvidia", "SGLang", "vLLM"]
dates: ["2026-09-27"]
keywords: ["gpu", "agent", "agents", "amd", "arr", "compute", "deepseek", "embedding", "embeddings", "fp8", "helios", "kv cache"]
source: docs/RAG/collect-261001-rattrapage/datacenter_gpu_guide.md
source_anchor: ""
source_lines: [2608, 2768]
sha256: 0fd57e889f6c0dbef486b380062f43071f4de39896a037d4e77e588d67fc36a3
---

# Les GPU datacenter / IA — présent et futur vérifié

## 128. Calculateur VRAM : modèles populaires (poids + 15 %)

| Modèle | Params | BF16 | FP8 | Q4 |
|---|---|---|---|---|
| Qwen3 8B | 8B | 18 Go | 9 Go | 5 Go |
| Gemma 3 12B | 12B | 28 Go | 14 Go | 7 Go |
| Qwen3 32B | 32B | 74 Go | 37 Go | 18 Go |
| Mistral Large 3 123B | 123B | 283 Go | 141 Go | 71 Go |
| Llama 3.1 8B | 8B | 18 Go | 9 Go | 5 Go |
| Llama 3.1 70B | 70B | 161 Go | 80 Go | 40 Go |
| Llama 3.1 405B | 405B | 931 Go | 466 Go | 233 Go |
| DeepSeek V3/R1 671B (MoE) | 671B | 1 543 Go | 772 Go | 386 Go |
| Qwen3 235B (MoE) | 235B | 540 Go | 270 Go | 135 Go |
| Mixtral 8×22B (MoE) | 141B | 324 Go | 162 Go | 81 Go |

Ajouter le KV cache selon le contexte visé (§23) — c'est lui qui fait déborder.

## 129. Quel GPU pour quel modèle (1 GPU, ctx 8K, batch 8)

| Modèle / précision | GPU minimum | GPU confortable |
|---|---|---|
| 8B BF16 | RTX 4090 (24 Go) | L40S |
| 32B BF16 | RTX PRO 6000 (74 Go) | RTX PRO 6000 |
| 70B FP8 | RTX PRO 6000 (90 Go) | H200 |
| 70B BF16 | H200 NVL (282 Go) | MI300X (192 Go) |
| 123B FP8 | H200 (141 Go : limite) | MI300X |
| 405B FP8 | 4× H200 (564 Go) | 8× H100 |
| 671B MoE Q4 | 4× H200 NVL | 8× B200 |
| 70B BF16 + 128K ctx | MI325X (256 Go) | MI355X |

## 130. Workloads agentiques 2026 : le nouveau dimensionnement

- **Pattern** : 1 requête utilisateur = 10–50 appels LLM (planification, outils,
  vérification). Le débit requis est **10–50×** celui du chatbot simple.
- **Conséquence** : le dimensionnement « 100 users → 1 GPU » devient
  « 100 users agentiques → 4–8 GPU ».
- **KV cache** : le prefix caching (SGLang RadixAttention) est **critique** —
  les prompts système d'agent sont longs et répétés.
- **Latence** : les agents sont sensibles au TTFT (time to first token) —
  privilégier la bande passante (B200/MI355X) au $/token pur.
- **Chiffre** : un agent « employé virtuel » consomme ~0,5–2 M tokens/jour —
  à 5 $/M tokens (B200), c'est 2,50–10 $/jour/agent en compute pur.

## 131. RAG à l'échelle : architecture type

```
Documents ──► Chunking ──► Embeddings (1× L40S) ──► Vector DB (NVMe)
                                                        │
Requête ──► Embedding ──► Recherche ──► Rerank (1× L40S) ──┤
                                                        ▼
                                              LLM 70B FP8 (1× RTX PRO 6000
                                              ou 1× H200) ──► Réponse
```

- **3 GPU** suffisent pour un RAG 70B à 500 req/h : 1 L40S (embed), 1 L40S (rerank),
  1 RTX PRO 6000 (génération).
- **Coût** : ~35 k$ de GPU — 10× moins qu'un « tout H100 » pour le même service.
- **Leçon** : découper par workload au lieu de tout mettre sur le même gros GPU.

## 132. Câblage optique vs cuivre : règles

| Lien | Cuivre (DAC) | Fibre (AOC/transceiver) |
|---|---|---|
| 400G, < 3 m | OK (DAC passif) | Inutile |
| 400G, 3–7 m | AEC (actif) | Possible |
| 400G, > 7 m | Non | Fibre OM4/OS2 |
| 800G, < 2 m | DAC | — |
| 800G, > 2 m | Non | Fibre obligatoire |

**Budget optique** : 8 nœuds × 8 liens 800G = 64 transceivers × ~500 $ = 32 k$
rien qu'en optiques — à intégrer au BOM réseau.

## 133. Qualité de l'énergie : harmoniques et facteur de puissance

- Les PSU à découpage des serveurs GPU génèrent des **harmoniques** (THD 5–15 %).
- **Conséquences** : échauffement du neutre, déclenchements intempestifs,
  pénalités du distributeur si THD > 8 %.
- **Solutions** : PSU à PFC actif (standard en 80+ Titanium), filtre d'harmoniques
  au TGBT si > 100 kW de charges non-linéaires, neutre surdimensionné (200 %).
- **Mesure** : analyseur de réseau pendant 1 semaine à la mise en service —
  exiger le rapport dans le DOE (dossier des ouvrages exécutés).

## 134. Coordination onduleur / groupe / DLC

```
Coupure réseau
  ├── 0–20 ms : onduleur prend la charge (VFI, zéro coupure)
  ├── 0–10 s  : groupe démarre, se synchronise
  ├── 10–30 s : bascule onduleur → groupe (bypass statique)
  └── En parallèle : la DLC tourne sur l'onduleur aussi !
```

**Point critique** : les pompes DLC et le dry cooler doivent être **secourus**
par l'onduleur/groupe — sinon les GPU cuisent en 2 minutes sans refroidissement
alors que l'IT tourne encore. Dimensionner l'onduleur **avec** le froid.

## 135. Commissioning DLC : checklist

- [ ] Rinçage et remplissage (eau déminéralisée + inhibiteurs)
- [ ] Test d'étanchéité 24 h à 1,5× la pression de service
- [ ] Équilibrage hydraulique par rack (débits nominaux)
- [ ] Test de fuite simulée (vanne d'isolement auto < 30 s)
- [ ] Montée en charge progressive : 25 % → 50 % → 100 % (paliers 1 h)
- [ ] Relevé des ΔT à pleine charge (conforme au design ±2 °C)
- [ ] Formation des exploitants (appoint, purge, alarmes)

## 136. Scénarios d'évolution : prévoir l'extension

| Scénario | Prévoir dès le jour 1 |
|---|---|
| +1 nœud dans 12 mois | Slots PDU libres, 20 % de froid en réserve, baie réseau |
| Passage air → DLC | Dalles avec passages hydrauliques, local CDU réservé |
| Rubin/MI450 dans 2 ans | Le rack NVL72/Helios ne tient pas dans une baie 19" standard — prévoir une zone |
| Doublement du contexte | +30 % VRAM ou nœuds supplémentaires au budget N+2 |

Le coût de la « réserve » (20 % de marge élec/froid) est 10× inférieur au coût
d'un retrofit.

## 137. Comparatif final : le podium par catégorie (27/09/2026)

| Catégorie | Or | Argent | Bronze |
|---|---|---|---|
| $/Go VRAM (neuf) | RTX 5090 (~81 $) | RTX 4090 (~100 $) | RTX PRO 6000 (~125 $) |
| $/Go VRAM (pro, ECC) | RTX PRO 6000 | L40S (~208 $) | B200 (~208 $) |
| Perf/watt FP8 | B200 (4,5) | MI350X (4,6) | MI300X (3,5) |
| Mémoire / 1 GPU | MI355X (288 Go) | MI325X (256 Go) | B200 / MI300X (192 Go) |
| Bande passante | MI455X (19,6 To/s, H2 2026) | B200/B300 (8 To/s) | MI355X (8 To/s) |
| Écosystème logiciel | NVIDIA CUDA | AMD ROCm (en progrès) | — |
| $/token (estimation) | B200 | MI355X | RTX PRO 6000 (petits volumes) |
| Disponibilité immédiate | RTX 4090/5090 | L40S | H100 (fin de cycle) |

## 138. Les 10 chiffres à retenir

1. **1 979** : TFLOPS FP8 d'un H100 — l'unité de compte du marché.
2. **141** : Go du H200 — le seuil du 70B FP16 en 1 GPU (presque).
3. **96** : Go de la RTX PRO 6000 — le 70B FP8 sur une carte à 12 k$.
4. **1 000** : watts du B200 — le mur thermique de la génération.
5. **8** : To/s de bande passante du B200/MI355X — le vrai moteur des tok/s.
6. **10,6** : kW au mur d'un nœud 8× B200 — le chiffre pour l'électricien.
7. **1,1** : le PUE d'une salle DLC bien conçue (vs 1,5 en air).
8. **15** : mois pour rentabiliser un 8× H100 acheté vs loué en 24/7.
9. **60** : % d'utilisation en dessous duquel la location gagne toujours.
10. **6–9** : mois d'un projet salle GPU, dont 4–9 pour l'électrique seul.
## 139. Test d'acceptation (recette) : protocole type

| Test | Critère d'acceptation |
|---|---|
| Burn-in 24 h tous GPU à 100 % | 0 Xid, temp < 85 °C, 0 throttling |
| NCCL all-reduce 8 GPU | ≥ 80 % de la bande passante théorique NVLink |
| Bande passante PCIe H→D | ≥ 90 % du négocié (ex. 115 Go/s en gen5 x16) |
| Inférence 70B FP8 (vLLM) | ≥ 50 tok/s batch 1, 0 erreur sur 10 k requêtes |
| Coupure réseau simulée | Bascule onduleur sans arrêt GPU |
| Arrêt/démarrage complet | Procédure §84 exécutée en < 2 h |
| DCGM | Toutes métriques remontées, alertes testées |
| Documentation | DOE + runbook + mots de passe sous scellés |

**Ne pas signer la réception sans le burn-in 24 h.** Un GPU faible sur 8 se
voit en 24 h, pas en 24 minutes.

## 140. Stockage : dimensionnement détaillé

