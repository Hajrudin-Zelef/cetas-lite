---
id: collect-261001-ia-llm/ia-llm/ia-modeles-specialises-21
title: "Encyclopédie des modèles IA — Volume 3"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "OpenAI", "SGLang", "vLLM"]
dates: []
keywords: ["embedding", "fp8", "gqa", "kv cache", "lora", "multimodal", "reranker", "sglang", "speculative decoding", "valuation", "vllm"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_specialises.md
source_anchor: ""
source_lines: [2113, 2206]
sha256: e745ccfd9069370b97a2562571ae90c9912835315e3b459da11d41a3d9119efd
---

# Encyclopédie des modèles IA — Volume 3

**Méthode.** Aucune spec inventée : quand l'info manquait, c'est marqué. Quand les sources se contredisaient, c'est signalé, pas tranché. Les recommandations pour ton RAG (sections 94–102) découlent de tes contraintes réelles (français technique, crédits limités, corpus existant), pas d'un classement abstrait.

*Fin du volume 3 — 134 sections.*

## 135. Cas pratique : dimensionner ton serveur RAG local (24 Go)

Exercice complet : tu as une RTX 4090 (24 Go), tu veux servir en local un RAG avec générateur 27B + embedding + reranker. Faisons les comptes.

**Étape 1 — le générateur.** Qwen3.8-27B en Q4_K_M = 18 Go (chiffre Ollama réel). Reste : 24 − 18 = 6 Go.

**Étape 2 — le KV cache.** Architecture type 27B : 48 couches, GQA 8 têtes KV, head_dim 128, cache FP16.
Par token : 2 × 48 × 8 × 128 × 2 = 192 Kio.
Avec 4 Go utiles (6 Go moins ~2 Go de marge moteur/fragmentation) : 4 Go / 192 Kio ≈ 21 800 tokens de contexte max par requête.

**Étape 3 — traduire en usage RAG.** Ton prompt typique : 2K (système + instructions) + 15 chunks × 1200 tokens (18K) + question ≈ 20–21K tokens. Ça passe — tout juste. Avec vLLM et le partage de préfixes (les 2K du système en cache une fois), tu gagnes de la marge.

**Étape 4 — l'embedding et le reranker.** Qwen3-Embedding-8B en FP16 = 16 Go : **impossible** en même temps que le générateur. Options :
- (a) Embedding en API (text-embedding-3-small actuel) + reranker 0,6B local (~1,5 Go en Q4) : total 18 + 1,5 + cache ≈ 23 Go — ça passe en serré, batch=1.
- (b) Tout en local mais générateur 14B (Q4 ≈ 8–10 Go) + embedding 8B FP16 (16 Go) : 26 Go — **ne passe pas**. Il faudrait l'embedding en INT8 (~8 Go) : 10 + 8 = 18 Go + cache — ça passe.
- (c) Embedding léger local (BGE-M3, 568M ≈ 1,2 Go FP16) + reranker 0,6B (1,5 Go) + générateur 27B (18 Go) : ~21 Go + cache — confortable.

**Verdict :** l'option (a) ou (c). Le générateur est le poste prioritaire : ne sacrifie jamais sa qualité pour loger l'embedding en local — l'embedding en API coûte des poussières (section 102).

**Étape 5 — le batch.** Tout ce qui précède est à batch=1 (une question à la fois, ton usage). Si un jour tu sers 4 utilisateurs : 4× le cache (4 × 4 Go = 16 Go) — impossible sur 24 Go avec un 27B. Solutions : modèle plus petit (14B), cache FP8 (÷2), ou carte 48 Go. Le dimensionnement multi-utilisateur se fait **toujours** sur le cache, jamais sur les poids.

## 136. SGLang vs vLLM : le match 2026

vLLM a dominé 2023–2025. SGLang (LMSYS/Berkeley) est le challenger 2026 — la recherche Granite cite un **support day-0 SGLang** (FP8 + speculative decoding) sur les nouveautés. Comparatif :

| Critère | vLLM | SGLang |
|---|---|---|
| Maturité / écosystème | Le standard, immense communauté | Montant vite, communauté plus petite |
| PagedAttention / batching continu | Oui (l'inventeur) | Oui (équivalent, RadixAttention) |
| **RadixAttention** | Non | Oui — partage de cache par **arbre de préfixes** (mieux que le préfixe linéaire de vLLM quand les prompts partagent des sous-parties) |
| Support day-0 nouveautés | Bon | Excellent en 2026 (Qwen3.5, FP8, MTP) |
| Speculative decoding | Oui | Oui (EAGLE, MTP natif) |
| Structured output / constrained decoding | Oui | Oui (le point fort historique de SGLang : grammaires, JSON forcé) |
| Langage | Python + CUDA | Python + CUDA (+ backend plus modulaire) |
| Courbe d'apprentissage | Docs abondantes | Plus raide |

Pour ton RAG, le différenciateur concret : **RadixAttention**. Tes requêtes partagent le prompt système + souvent les mêmes chunks populaires (les guides les plus interrogés) : l'arbre de préfixes de SGLang mutualise le cache là où vLLM ne partage que le préfixe strictement identique. En pratique, sur un RAG à corpus stable, ça peut faire gagner 20–40 % de VRAM de cache.

Recommandation : vLLM par défaut (écosystème, docs, stabilité), SGLang à évaluer si tu sers en continu et que le cache devient le goulot. Les deux parlent l'API OpenAI — la migration est un changement d'image Docker, pas un rewrite.

## 137. Prompts avancés : few-shot, citations, anti-hallucination

Le template de base (section 116) suffit pour démarrer. Trois upgrades quand tu veux durcir :

**1. Le few-shot comportemental.** Au lieu de décrire le format voulu, montre-le — 2 exemples question/réponse dans le prompt système :

```
Exemple :
Question : « Quelle est la procédure de reset du mot de passe eKit ? »
Réponse :
**Diagnostic** : mot de passe d'administration perdu sur AP361.
**Cause probable** : …
**Procédure** : 1. … [extrait 2] 2. … [extrait 2]
**Vérification** : …
```

Coût : ~500–800 tokens de plus par requête (cachés à 90 % via prompt caching — section 122). Gain : format constant sans LoRA.

**2. Les citations vérifiables.** Exige le format `[extrait N]` **et** fais vérifier en post-traitement : un script contrôle que chaque citation pointe vers un extrait réellement fourni (regex + index). Les citations inventées (`[extrait 9]` alors qu'il n'y en a que 5) sont le signal d'hallucination le plus facile à détecter automatiquement.

**3. Le double passage anti-hallucination.** Pour les réponses critiques (procédures de consignation électrique, valeurs de protection) : second appel au modèle avec « Vérifie chaque valeur numérique de la réponse ci-dessus contre les extraits. Liste les écarts. » Coût : ×2 sur ces requêtes — à réserver aux réponses à risque, pas au tout-venant.

**4. La température basse + top-p resserré.** RAG factuel : temperature 0,1–0,3, top_p 0,9. La créativité est l'ennemie de la fidélité aux extraits. (La température n'empêche pas l'hallucination à elle seule — c'est le « je ne sais pas » obligatoire qui fait le travail.)

## 138. Migrer d'embedding sans tout casser : le protocole

Le jour où tu voudras changer d'embedding (section 95 : seulement si +5 pts de rappel@5), voici la procédure sans interruption :

**Phase 1 — double index (1–2 jours).**
1. Garde l'index actuel en production (text-embedding-3-small).
2. Vectorise tout le corpus avec le challenger dans un **second index** (Qdrant gère plusieurs collections).
3. Coût : une re-vectorisation ≈ $0,04–0,40 en API (section 102) — négligeable.

**Phase 2 — évaluation comparative (1 jour).**
4. Joue tes 50–100 questions de test sur les deux index (même rerank, même générateur, même prompt).
5. Compare rappel@5 et qualité finale. Décision chiffrée : migration seulement si le challenger gagne **≥ 5 points** de façon stable.

**Phase 3 — bascule progressive (1 jour).**
6. Bascule 10 % du trafic (ou tes propres questions tests) sur le nouvel index, surveille les logs une semaine.
7. Si OK : bascule 100 %, archive l'ancien index (ne le supprime pas tout de suite — rollback possible).

**Pièges :**
- Les dimensions changent (1536 → 1024/4096) : les deux index sont incompatibles, pas de migration « en place ».
- Si tu avais fine-tuné l'ancien embedding (section 92), le nouveau repart de zéro — prévois de régénérer les paires.
- Matryoshka (Jina v5, Qwen3-Embedding) : choisis la dimension de stockage **avant** d'indexer (256/512/1024) — changer après = re-vectoriser.

## 139. Le RAG multimodal : brancher la vision sur ton corpus

Tes guides contiennent des images (faces avant d'équipements, schémas, captures CLI, topologies eNSP). Trois niveaux d'intégration, du simple au avancé :

