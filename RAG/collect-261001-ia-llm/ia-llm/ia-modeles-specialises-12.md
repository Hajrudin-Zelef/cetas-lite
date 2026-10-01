---
id: collect-261001-ia-llm/ia-llm/ia-modeles-specialises-12
title: "Encyclopédie des modèles IA — Volume 3"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Huawei", "Mistral", "OpenAI", "Unsloth", "vLLM"]
dates: []
keywords: ["attention", "claude", "embedding", "embeddings", "fine-tuning", "fp8", "gguf", "gpu", "gqa", "kv cache", "llama", "llama.cpp"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_specialises.md
source_anchor: ""
source_lines: [1217, 1315]
sha256: 9026bc947aa3c78fa47aca4e1a2fe9b6ef36ced349b37a76b0164110dd485d95
---

# Encyclopédie des modèles IA — Volume 3

```
Combien de VRAM as-tu ?
├── ≥ 80 Go (H100, 2× 4090...) → 70B en Q4/Q5 GGUF ou FP8 vLLM
├── 24 Go (RTX 4090, 3090) → 27B–32B en Q4_K_M (ex: Qwen3.8-27B = 18 Go)
│   └── + besoin long contexte → Q4 + cache FP8 (vLLM)
├── 16 Go (4060 Ti, M1 Pro...) → 9B–14B en Q4_K_M
├── 8–12 Go → 7B–8B en Q4_K_M (ex: ~4–5 Go)
└── CPU uniquement → 7B en Q4_K_M via llama.cpp (lent mais fonctionnel,
    ~5–15 tok/s sur un bon CPU moderne)
```

Tableau récapitulatif :

| Budget VRAM | Modèle type | Format | Usage RAG |
|---|---|---|---|
| 8 Go | 7B–8B | Q4_K_M (~4–5 Go) | Générateur correct, contexte ≤ 8K |
| 16 Go | 9B–14B | Q4_K_M (~8–10 Go) | Bon générateur FR, contexte 16–32K |
| 24 Go | 27B–32B | Q4_K_M (~15–20 Go) | Excellent générateur, contexte 32K+ |
| 48 Go | 70B | Q4_K_M (~40 Go) | Très bon, ou 2× 27B |
| 80 Go+ | 70B FP8 / 122B-A10B | FP8 / Q4 | Serveur vLLM, batch |

Règle d'or : **laisse 20–30 % de VRAM libre pour le KV cache** (section 69). Un 27B Q4_K_M (18 Go) sur une 24 Go laisse 6 Go de cache — soit ~47K tokens de cache en FP16 pour un 27B GQA. Ça passe pour un RAG standard (10–20 chunks).

## 88. Fine-tuning : le principe

Le fine-tuning, c'est **continuer l'entraînement** d'un modèle déjà entraîné, sur tes données, pour l'adapter : ton vocabulaire métier, ton format de réponse, tes procédures. Trois régimes :

| Régime | Données | Coût | Effet |
|---|---|---|---|
| **Pre-training** (from scratch) | Milliards de tokens | Millions de $ | Nouveau modèle — hors sujet pour toi |
| **Fine-tuning complet** | Dizaines de milliers d'exemples | Dizaines de milliers de $ (GPU) | Change le modèle en profondeur, risque d'oubli catastrophique |
| **PEFT** (LoRA & co) | Centaines à milliers d'exemples | Quelques $ à centaines de $ | Adapte le style/comportement, préserve les capacités |

En 2026, le fine-tuning « normal » pour un particulier/une PME, c'est le **PEFT** (Parameter-Efficient Fine-Tuning) : on n'entraîne qu'une petite fraction des paramètres. LoRA en est la méthode reine.

Note : Anthropic **n'offre pas** de fine-tuning public (relevé dans la recherche Claude) — si tu veux fine-tuner, c'est sur poids ouverts (Qwen, Llama, Mistral, Gemma) ou via les API qui le proposent (OpenAI sur certains modèles).

## 89. LoRA expliqué simplement

**LoRA** (Low-Rank Adaptation, Hu et al., 2021) : au lieu de modifier la matrice de poids W du modèle (ex : 8192×8192 = 67M paramètres par couche), on apprend une **correction** ΔW = A×B, où A et B sont deux petites matrices (ex : 8192×16 et 16×8192 = 262K paramètres — **256× moins**).

Schéma :

```
W_original (figé, 67M params)  +  A×B (entraîné, 262K params)  =  W_adapté
```

Points clés :

- **Le modèle de base ne bouge pas.** L'adaptateur LoRA est un petit fichier séparé (quelques dizaines de Mo) qu'on charge par-dessus. On peut empiler/retirer des adaptateurs à volonté.
- **Le rang r** (16 dans l'exemple) contrôle la capacité d'adaptation : r=8 pour un style léger, r=32–64 pour un domaine technique pointu. Plus r est grand, plus l'adaptateur est gros et plus il faut de données.
- **Quelles couches ?** En pratique on applique LoRA aux projections d'attention (q, k, v, o) et souvent au MLP. Les recettes standard (axolotl, Unsloth) le font par défaut.
- **Inférence** : l'adaptateur peut être **fusionné** dans les poids (coût nul à l'inférence) ou chargé dynamiquement (multi-adaptateurs servis par le même modèle de base — utile pour servir plusieurs métiers).

Ordre de grandeur : fine-tuner un 7B en LoRA r=16 demande ~500–2000 exemples de qualité et tient sur **une RTX 4090** en quelques heures.

## 90. QLoRA : LoRA sur modèle quantizé

**QLoRA** (Dettmers et al., 2023) : on combine LoRA avec un modèle de base **quantizé en 4-bit (NF4)**. Le modèle gelé occupe ~4× moins de VRAM, ce qui libère la place pour l'entraînement des adaptateurs + les gradients + l'optimiseur.

Le tour de force de QLoRA : fine-tuner un **65B sur un seul GPU 48 Go** (dans le papier originel) — impensable en fine-tuning classique. En 2026, les recettes sont industrialisées :

- **Unsloth** : kernels optimisés, fine-tuning 2–5× plus rapide, GGUF exportables directement. La référence pour le local (Qwen, Llama, Gemma, Mistral).
- **axolotl** : configs YAML, recettes éprouvées par modèle (la recherche Qwen cite `axolotl-ai-cloud/axolotl` pour Qwen3.5).
- **Ollama / llama.cpp** : servent les GGUF issus d'Unsloth sans conversion.

Limites honnêtes : QLoRA n'apprend pas de **faits nouveaux** aussi bien qu'un fine-tuning complet — il excelle à apprendre des *formats*, des *styles*, des *comportements*. Pour injecter des connaissances (tes 25 000 lignes de guides), le RAG reste supérieur (section 91).

## 91. Fine-tuner vs RAG : la règle de décision

La question que tout le monde se pose. La réponse en tableau :

| Objectif | RAG | Fine-tuning |
|---|---|---|
| Le modèle doit **connaître** tes docs (faits, procédures, specs) | ✅ Idéal — les faits restent dans l'index, à jour | ⚠️ Possible mais les faits « appris » se périment et s'hallucinent |
| Le modèle doit **parler comme** ton équipe (ton, format de réponse) | ❌ Le prompt système suffit souvent | ✅ Le cas d'usage roi du LoRA |
| Les docs **changent souvent** (firmwares, tarifs, procédures) | ✅ Re-indexer = quelques minutes | ❌ Re-fine-tuner à chaque changement |
| Traçabilité (« d'où vient cette réponse ? ») | ✅ Chunks cités, sources vérifiables | ❌ Boîte noire |
| Coût marginal par question | Faible (embedding + retrieval) | Nul (pas de retrieval) mais latence du gros modèle |
| Mise en œuvre | Index + prompt | Données d'entraînement + GPU + évaluation |

La règle :

1. **Toujours RAG d'abord** pour des connaissances. Tes guides Huawei, onduleurs, Windows : c'est du RAG, pas du fine-tuning. Les faits changent (nouveaux firmwares, nouvelles gammes), doivent être cités, et tu en as des dizaines de milliers de lignes — le fine-tuning n'apprendrait qu'une fraction.
2. **Fine-tune (LoRA) pour le comportement** : si tu veux que le modèle réponde toujours en français technique avec la structure « Diagnostic → Cause probable → Procédure → Vérification », un LoRA léger sur 500 exemples de ce format fera mieux qu'un long prompt système — et coûtera moins de tokens à chaque requête.
3. **Les deux se combinent** : RAG pour les faits + LoRA pour le format. C'est l'architecture « RAG + adaptateur métier », la plus robuste en production.

Le piège à éviter : fine-tuner pour « apprendre » tes docs puis découvrir que le modèle hallucine des références de cartes ou des codes d'erreur avec aplomb. Le fine-tuning rend le modèle **plus confiant**, pas plus exact, sur les faits.

## 92. Fine-tuning d'embeddings : le levier sous-estimé

On fine-tune les LLM, mais on peut aussi fine-tuner les **embeddings** — et pour un RAG, l'impact est souvent supérieur à celui d'un LoRA sur le générateur.

Principe : on entraîne l'embedding sur des paires (question, passage pertinent) et des contre-exemples (question, passage non pertinent) avec une **loss contrastive** (InfoNCE) : rapprocher les bons, éloigner les mauvais. Données nécessaires : quelques centaines à quelques milliers de paires — que tu peux générer semi-automatiquement (questions synthétiques à partir de tes chunks, validées par échantillonnage).

Effets mesurés typiquement : +5 à +15 points de rappel@10 sur ton domaine, parce que l'embedding apprend ton vocabulaire métier (« U163 », « DK-8350 », « VFI », « THDi ») que l'embedding générique ne pondère pas bien.

Recette pragmatique 2026 :

