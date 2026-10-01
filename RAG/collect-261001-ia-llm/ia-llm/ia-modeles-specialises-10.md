---
id: collect-261001-ia-llm/ia-llm/ia-modeles-specialises-10
title: "Encyclopédie des modèles IA — Volume 3"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Falcon", "Moonshot", "Nvidia"]
dates: []
keywords: ["attention", "awq", "benchmarks", "datacenter", "deepseek", "fine-tuning", "fp4", "fp8", "gguf", "gptq", "gpu", "gqa"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_specialises.md
source_anchor: ""
source_lines: [1028, 1113]
sha256: 86837cac5201ea68d1dedb8486a55dd607b0d0933aae29869dd639f644a46ba8
---

# Encyclopédie des modèles IA — Volume 3

Règle de lecture : le **total** détermine la VRAM (il faut charger tous les experts), les **actifs** déterminent la vitesse et le coût de calcul par token. Un « 1T » qui n'active que 32B se sert comme un 32B dense… mais se charge comme un 1T. C'est toute l'implication serving de la section suivante.

## 74. Implications serving : tout en VRAM, calcul partiel

Le paradoxe du MoE en production :

- **Mémoire** : il faut **tous** les experts en VRAM (ou en RAM avec offload, au prix de la vitesse). Kimi K2.6 : 1T paramètres = ~2 To en BF16, ~1 To en FP8, ~500 Go en INT4. D'où les déploiements « 8× H200 » ou « 64+ accélérateurs » (Kimi K3).
- **Calcul** : seuls les experts sélectionnés calculent → le nombre de FLOPS par token suit les actifs, pas le total.
- **Bande passante** : le facteur limitant en génération (décodage autoregressif) est la **bande passante mémoire** — lire les poids depuis la VRAM à chaque token. Pour un MoE, on ne lit que les experts actifs… mais le routeur peut activer des experts différents à chaque token, donc en pratique on lit beaucoup.

Conséquences :

1. **Le MoE adore le batching.** Avec un gros batch, les mêmes experts servent plusieurs tokens : le coût de lecture des poids s'amortit. À batch=1, le MoE est inefficace ; à batch=64+, il brille. D'où l'usage datacenter/API plutôt que local solo.
2. **L'expert parallelism (EP)** : on répartit les experts sur plusieurs GPU (chaque GPU héberge un sous-ensemble d'experts), reliés en NVLink. DeepSeek V4 documente l'EP sur NVL72. Sans interconnexion rapide, le MoE multi-GPU s'effondre.
3. **En local, le MoE se quantize agressivement.** Kimi K2 est **livré nativement en INT4** (pas de release haute précision) : 1T paramètres ≈ 500 Go en INT4 — encore trop pour du local, mais l'intention est claire.

## 75. Exemples 2026 : DeepSeek V4, Kimi K3, Qwen4-Exp

Trois MoE représentatifs, trois philosophies :

**DeepSeek V4-Pro** (24 avril 2026, GA 13 août 2026) : ~1,6T totaux / 49B actifs. Poids ouverts **MIT** (le plus permissif du lot). Formats BF16/FP8/**FP4 natifs** (NVFP4). Entraîné sur 32T+ tokens. Prix API : $0,435/M input (cache miss) mais **$0,003625/M en cache hit** — 120× moins cher : le prompt caching agressif est la vraie proposition de valeur. Benchmarks fournisseur : Codeforces 3206, SWE-bench Verified 80,6 %.

**Kimi K3** (juillet 2026) : ~2,8T / 104B, 896 experts. Le plus gros MoE ouvert du snapshot. Checkpoint 1,56 To, 64+ accélérateurs conseillés. Licence custom (pas OSI). Positionnement : frontière de la capacité, pas de l'efficacité.

**Qwen4-Exp** (août 2026) : 125B / 6B — le ratio le plus extrême (5 % d'actifs). Hybride Gated DeltaNet + attention (section 64), tête MTP 4B pour le décodage spéculatif. Positionnement : efficacité maximale, preview technique.

## 76. MoE vs dense : le verdict pratique

| Critère | Dense (ex : Llama 70B, Gemma 4 31B) | MoE (ex : Mixtral, Qwen3.6-35B-A3B) |
|---|---|---|
| Qualité à paramètres actifs égaux | Référence | Équivalent ou meilleur (plus de capacité) |
| VRAM (poids) | Proportionnelle aux paramètres | Proportionnelle au **total** (lourd) |
| Vitesse par token | Proportionnelle aux paramètres | Proportionnelle aux **actifs** (rapide) |
| Batch=1 (local) | Efficace | Gaspille (experts chargés, peu utilisés) |
| Gros batch (serveur) | Linéaire | Excellent (amortissement) |
| Fine-tuning (LoRA) | Simple, bien outillé | Plus délicat (routeur, équilibrage) |
| Quantization | Mature (GPTQ/AWQ/GGUF) | Possible mais moins standard |

Verdict :

- **En local, solo, sur un GPU** : dense. Un Qwen3.5-27B dense (18 Go en Q4_K_M) bat un MoE équivalent en simplicité et souvent en vitesse réelle.
- **En serveur, avec du batch** : MoE. C'est pour ça que les API pas chères (DeepSeek $0,435/M) sont des MoE : le coût par token servi s'effondre à gros batch.
- **Pour ton RAG** : le générateur sera probablement appelé en batch=1 (une question à la fois). Privilégie un dense de bonne qualité — sauf si tu batches les requêtes (file d'attente), auquel cas un petit MoE (Qwen3.6-35B-A3B, 3B actifs) devient intéressant.

## 77. L'attention, en une page

Le mécanisme d'attention (2017, « Attention Is All You Need ») répond à une question simple : pour prédire le token suivant, **quels tokens précédents sont pertinents, et à quel degré** ? Chaque token émet une **requête** (Query) qui se compare à toutes les **clés** (Keys) des tokens précédents ; les scores de similarité pondèrent les **valeurs** (Values) pour construire le contexte.

Coût : pour n tokens, n² comparaisons — **quadratique**. 4K tokens = 16M comparaisons ; 128K = 16Md. C'est le mur fondamental du long contexte, et toute l'histoire des sections 62–66 est une histoire de contournement de ce mur : moins de têtes KV (MQA/GQA), compression (MLA), fenêtres (sliding), couches linéaires (Mamba, Gated DeltaNet).

Ce qu'il faut retenir : quand une fiche vante « 1M de contexte », la vraie question est **à quel prix** — VRAM du cache (section 58), qualité du rappel à longue distance (les benchmarks « needle » vs « vraies tâches »), et latence du premier token (le *prefill* de 1M tokens prend des dizaines de secondes même sur un H100).

## 78. MHA → MQA → GQA : la généalogie

| Année | Mécanisme | Idée | Cache relatif |
|---|---|---|---|
| 2017 | MHA | 1 tête KV par tête Q | 100 % (référence) |
| 2019–2023 | MQA | 1 seule tête KV partagée | ~3 % (÷32) |
| 2023 | GQA | Groupes de têtes Q partagent des têtes KV | ~12–25 % (÷4 à ÷8) |
| 2024 | MLA (DeepSeek) | K/V compressés en latent | ~7 % (×14 de réduction) |
| 2025–2026 | Hybrides | Attention globale seulement sur une minorité de couches | Variable, jusqu'à ~10 % du MLA |

Le MQA (Shazeer, 2019) a d'abord servi aux modèles rapides (Falcon, StarCoder). Le GQA (Ainslie et al., 2023) est devenu le standard parce qu'il ne perd quasi rien en qualité. Le MLA est le pari de DeepSeek. Les hybrides 2026 (DeepSeek V4, Qwen4-Exp) combinent tout : GQA sur les couches d'attention restantes + couches linéaires ailleurs.

Point pratique : ces mécanismes sont **figés à l'entraînement**. Tu ne peux pas convertir un MHA en GQA après coup (enfin, il existe des techniques de « uptraining » GQA, mais c'est du post-training lourd). Le choix se fait donc à la sélection du modèle, pas au déploiement.

## 79. MLA en détail (pour les curieux)

Le MLA compresse : au lieu de stocker K (n_têtes × dim) et V (n_têtes × dim), il projette la séquence en un **vecteur latent** c de dimension réduite (ex : 512 au lieu de 8192), stocke c dans le cache, et reconstruit K/V par couche via des projections apprises.

Schéma simplifié :

```
Standard :  token → [K_1..K_h, V_1..V_h]  → cache (gros)
MLA :       token → c (latent compact)     → cache (petit)
                    puis par couche : c → [K_1..K_h, V_1..V_h] (décompression)
```

Le découplage RoPE (l'encodage positionnel rotatif) est la subtilité technique : DeepSeek a dû séparer la partie positionnelle (non compressible) de la partie contenu (compressible). C'est élégant, mais ça complexifie l'implémentation — d'où l'abandon du MLA pur sur V4 au profit de l'hybride plus simple à opérer.

Pourquoi c'est important pour toi : le MLA prouve qu'on peut diviser le cache par 10 **sans changer l'interface** du modèle. Les futurs modèles « long contexte pas cher » utiliseront ces techniques — ton dimensionnement VRAM d'aujourd'hui sera peut-être divisé par 5 demain.

## 80. Pourquoi tout ça compte pour ton RAG

Synthèse opérationnelle des sections 57–79 :

