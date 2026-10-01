---
id: collect-261001-ia-llm/ia-llm/ia-modeles-specialises-13
title: "Encyclopédie des modèles IA — Volume 3"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Cohere", "Google", "Huawei", "OpenAI"]
dates: []
keywords: ["apache", "cohere", "embedding", "embeddings", "fine-tuning", "gemini", "gpu", "lora", "moe", "multimodal", "open-weight", "reranker"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_specialises.md
source_anchor: ""
source_lines: [1316, 1391]
sha256: f69fce39445b7f66a5cc166dfd2921be256298c399058191222b9468937f0c1c
---

# Encyclopédie des modèles IA — Volume 3

1. Prends un embedding open-weight léger (BGE-M3, Jina v5 small si usage perso).
2. Génère 1000–3000 paires question/réponse depuis tes guides (un LLM génère les questions, tu valides un échantillon).
3. Fine-tune contrastif (sentence-transformers rend ça accessible en quelques lignes).
4. Évalue sur tes 50 questions de test (section 11) : si le rappel@5 progresse, tu gardes.

Coût : quelques heures de GPU modeste. C'est le meilleur ROI « entraînement » pour un RAG.

## 93. Les pièges du fine-tuning

Quatre échecs classiques, pour ne pas les payer :

1. **Trop peu de données, trop de rangs.** 200 exemples avec r=64 : le modèle apprend le bruit. Règle : r=8–16 pour < 1000 exemples, r=32+ au-delà de 5000.
2. **L'oubli catastrophique (même en LoRA).** Un LoRA agressif sur un seul domaine peut dégrader les capacités générales (le modèle devient moins bon en dehors de ton domaine). Mitigation : mélanger 10–20 % de données généralistes dans l'entraînement.
3. **Évaluer à l'œil.** « Ça a l'air mieux » n'est pas une évaluation. Construis un jeu de test **avant** de fine-tuner (50–100 exemples avec réponses attendues), mesure avant/après. Sans ça, tu ne sais pas si tu as progressé ou régressé.
4. **Fine-tuner un modèle déjà mauvais pour ton usage.** Le fine-tuning amplifie, il ne ressuscite pas. Si le modèle de base est nul en français technique, un LoRA ne le sauvera pas — change de base.

Et le piège organisationnel : le fine-tuning crée une **dette de maintenance**. Chaque nouvelle version du modèle de base = re-fine-tuner, re-évaluer. Le RAG, lui, survit aux changements de modèle générateur sans toucher à l'index (juste le prompt à ajuster).

## 94. Choisir un embedding pour ton RAG : les critères

Ton cas : corpus français technique (guides 1600–4200 lignes, CLI Huawei, docs réseau), usage personnel intensif, contrainte crédits tokens (tu fais la récup toi-même). Les critères, par ordre d'importance :

1. **Rappel@k sur TON corpus en français.** Le reste est secondaire. Construis ton jeu de 50 questions (section 11) et mesure. Un embedding à 70 en MMTEB qui comprend « U163 reset fusion » bat un embedding à 78 qui ne le comprend pas.
2. **Contexte d'entrée ≥ taille de tes chunks.** Si tu chunkes à 1500–2000 tokens, élimine tout modèle < 2048 tokens de contexte (nomic v2 MoE : 512 — éliminé pour toi, sauf chunks très courts).
3. **Coût total.** En API : ~$0,02–0,20/M tokens. Ton corpus : ~100 000 lignes de guides ≈ 1,5–2M tokens ≈ **$0,04 à $0,40 pour vectoriser tout le corpus une fois**. C'est négligeable — le coût API des embeddings n'est pas ton problème. Le vrai coût serait en self-host (GPU).
4. **Licence.** Usage perso : tout est permis. Si un jour le RAG devient un outil d'équipe/d'entreprise : Apache 2.0 requis → Qwen3-Embedding-8B, BGE-M3, nomic v2. Élimine Jina (CC BY-NC) et les API fermées si tu veux l'option.
5. **Dimensions et Matryoshka.** Stocker 2M tokens de chunks à 4096 dims en float32 = 2M × 4096 × 4 octets ≈ 33 Go d'index. À 256 dims (Matryoshka) : ~2 Go. Pour un index local (Qdrant, pgvector), ça compte.

## 95. Recommandation embeddings 2026 (pour ton usage)

| Profil | Choix | Pourquoi |
|---|---|---|
| **Statu quo (recommandé à court terme)** | Garder `text-embedding-3-small` | Il marche, il est pas cher (~$0,02/M), ton index existe déjà. Ne change pas ce qui marche sans mesure. |
| **Upgrade API, FR/multilingue** | Gemini Embedding 2 ou Voyage 4 | 100+ langues, multimodal (PDF natif pour Gemini) ; ~$0,02–0,20/M |
| **Self-host, licence propre** | Qwen3-Embedding-8B (Apache 2.0) | 8B, 4096 dims MRL, 32K, scores MMTEB publiés ; ~16 Go VRAM en FP16 |
| **Self-host léger, perso** | Jina v5 small (239M, 32K, Matryoshka) | Tourne sur CPU ; mais CC BY-NC |
| **Baseline Apache 2.0 légère** | BGE-M3 (568M) | La valeur sûre permissive, mais contexte 8K à surveiller |

Protocole de décision : ne migre que si un challenger bat ton actuel de **+5 points de rappel@5** sur tes 50 questions FR. En dessous, le coût de migration (re-vectoriser tout le corpus, revalider) n'est pas rentabilisé.

## 96. Le cas text-embedding-3-small (ton actuel)

Tu utilises `text-embedding-3-small` d'OpenAI depuis le début du pipeline. Évaluation honnête :

- **Points forts** : 1536 dimensions (bon compromis), prix très bas (~$0,02/M tokens — vectoriser 2M tokens = $0,04), API stable et rapide, support multilingue correct dont le français. C'est le choix « ça marche » de 80 % des RAG en production.
- **Limites connues** : modèle de décembre 2023 — deux générations derrière l'état de l'art 2026 (Gemini Embedding 2, Nemotron-3, Jina v5). Contexte 8192 tokens (suffisant pour tes chunks). Pas de Matryoshka officiel (dimensions fixes 1536, réductible à 512 via le paramètre `dimensions` — avec perte).
- **Verdict** : pas d'urgence à migrer. Le jour où tu mesures un plafond de rappel sur tes questions FR techniques, les challengers sont identifiés (section 95). Le fine-tuning d'embeddings (section 92) sur une base open serait l'étape d'après, pas un changement d'API.

## 97. Reranker : quand ça vaut le coup (pour toi)

Reprenons la section 22 avec ton cas concret. Ton corpus a les trois propriétés qui rendent le rerank rentable :

1. **Vocabulaire répétitif** : « port », « VLAN », « PoE », « onduleur », « batterie » apparaissent dans des dizaines de chunks — l'embedding seul peine à départager.
2. **Chunks longs** : tes guides font 1600–4200 lignes ; même chunkés à 1500 tokens, le signal se dilue.
3. **Semi-structuré** : tableaux de specs, tableaux de commandes CLI, codes d'erreur — là où le rerank listwise gagne +9,6 pts (Jina v3.5).

Protocole : retrieval top-50 par embedding → rerank → top-5 au générateur. Surcoût par question : ~0,5–2 s et quelques centièmes de centime (API) ou un petit GPU en local (0,6B = quelques Go VRAM).

Quand t'en passer : si tes questions sont factuelles simples (« quel est le code U163 ? ») et que le rappel@5 est déjà > 95 %, le rerank n'ajoute rien. Mesure d'abord.

## 98. Recommandation reranker (pour ton usage)

| Profil | Choix | Pourquoi |
|---|---|---|
| **API, sans infra** | Cohere Rerank 4.0 Fast ($2,00/1000 search units) | 100+ langues, 32K, pensé pour RAG agentiques |
| **API, éco** | Voyage rerank-2.5-lite (~$0,02/M, à vérifier) | Le moins cher si l'unité se confirme |
| **Self-host, qualité max, perso** | jina-reranker-v3.5 (0,6B, listwise, 131K) | BEIR 63,20 ; 93 langues ; ~4 Go VRAM |
| **Self-host, licence propre** | Qwen3-Reranker-0.6B (Apache 2.0) | Même gabarit que Jina, usage commercial OK |
| **Pipeline 100 % Apache 2.0** | bge-reranker-v2-m3 (568M) | Avec BGE-M3 en embedding : stack permissive complète |

Note d'architecture : le reranker 0,6B se sert **sur le même GPU** que l'embedding (ou même sur CPU en mode dégradé). Pas besoin d'une carte dédiée. En revanche, ne mets pas le reranker et le générateur 27B sur la même 24 Go — le générateur + son cache priment.

## 99. Choisir le LLM générateur : les critères

Le générateur reçoit les chunks + la question et rédige la réponse. Critères :

