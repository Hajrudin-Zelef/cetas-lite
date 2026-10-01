---
id: collect-261001-ia-llm/ia-llm/ia-modeles-specialises-16
title: "Encyclopédie des modèles IA — Volume 3"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Moonshot"]
dates: ["2026-09-27"]
keywords: ["agi", "benchmark", "benchmarks", "deepseek", "embedding", "embeddings", "gemini", "kimi", "opus 4", "prefill", "reranker"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_specialises.md
source_anchor: ""
source_lines: [1596, 1709]
sha256: 7d992e5e97235787678bf3d1422c0daf9cc68bef934baba0958acb8da4a1d27a
---

# Encyclopédie des modèles IA — Volume 3

Le retrieval vectoriel pur a une faiblesse : il ne sait pas filtrer (« je veux la doc du **S310**, pas du AR720 »). Les métadonnées comblent ça avec le **filtrage hybride** : filtre exact sur les métadonnées + recherche vectorielle dans le sous-ensemble.

Métadonnées à stocker par chunk (recommandé pour ton corpus) :

| Champ | Exemple | Usage |
|---|---|---|
| `source` | `huawei_s310_guide.md` | Citer la source dans la réponse |
| `section` | `12. Configuration du PoE` | Affichage, navigation |
| `famille` | `S310`, `AP361`, `USG6000` | Filtre « ne cherche que dans… » |
| `type` | `procedure`, `spec`, `tableau`, `code`, `quiz`, `glossaire` | Boost selon la question (« donne-moi la procédure » → boost `procedure`) |
| `langue` | `fr` | Corpus multilingue futur |
| `date` | `2026-09-27` | Fraîcheur (firmware, gammes) |
| `page` / `image_path` | Pour les PDF/images | Retrouver l'original (section 36) |

Effet mesuré typiquement : le filtrage par `famille` avant la recherche vectorielle fait gagner +10 à +20 points de précision@5 sur les corpus multi-équipements — plus que n'importe quel changement d'embedding. C'est gratuit à l'indexation, ça ne coûte rien à la requête.

## 113. HyDE : le document hypothétique

**HyDE** (Hypothetical Document Embeddings, Gao et al., 2022) : au lieu d'embedder la question telle quelle, on demande d'abord au LLM de **rédiger une réponse hypothétique** (qui peut être fausse), on embedde *cette réponse*, et on cherche avec ce vecteur.

Pourquoi ça marche : l'embedding d'une question (« quel est le mode maintenance Kyocera ? ») vit dans l'espace des *questions* ; l'embedding d'un chunk vit dans l'espace des *réponses*. HyDE translate la question dans l'espace des réponses avant de chercher — le matching est meilleur.

Quand l'utiliser :

- ✅ Questions courtes et vagues (« problème de chauffe ? ») — HyDE les « déplie » en réponse détaillée cherchable.
- ✅ Corpus très technique où le vocabulaire question ≠ vocabulaire doc (« ça imprime flou » vs « densité du toner, U034 »).
- ❌ Questions précises déjà (« code erreur C6000 ») — HyDE n'ajoute rien et coûte un appel LLM.

Coût : 1 appel LLM supplémentaire par requête (quelques centaines de tokens). À activer comme option, à mesurer sur tes 50 questions — gain typique +3 à +8 points de rappel quand ça aide, 0 quand ça n'aide pas.

## 114. Le rerank dans le pipeline : où, combien, comment

Position canonique (détaillée section 101) : embedding → top-50 → rerank → top-5/10 → génération. Paramètres à régler :

| Paramètre | Défaut recommandé | Notes |
|---|---|---|
| Candidats embedding (k1) | 50 | 20 si corpus petit, 100 si très bruité |
| Candidats rerankés (k2) | 5–10 | = ce que voit le générateur |
| Seuil de score | Pas de seuil fixe | Les scores de rerank ne sont pas calibrés entre requêtes — toujours prendre le top-k, jamais « score > x » |
| Rerank listwise : taille de passe | 20–64 docs/passe | jina-reranker-v3.5 : ~64/passe dans 131K de contexte |

Piège : reranker **trop peu** de candidats (k1=10) — si le bon chunk est en position 30, le rerank ne le verra jamais. Le rerank ne **retrouve** pas, il **re-classe** : la qualité du top-k1 borne tout.

## 115. Fenêtres de contexte 1M : utiles ou gadget ?

En 2026, le 1M tokens est banalisé (Gemini 3.1 Pro, Kimi K3, DeepSeek V4, Qwen3.8-27B 262K…). La tentation : « je mets tout mon corpus dans le prompt, plus besoin de RAG ». Analyse froide :

| Argument « pour » | Réalité |
|---|---|
| Plus de retrieval à maintenir | Le retrieval est remplacé par un tri implicite — moins contrôlable |
| Le modèle « voit tout » | **Lost in the middle** : la qualité de rappel chute au milieu des longs contextes, même en 2026 |
| Simple conceptuellement | Coût : 1M tokens en entrée à $2–4/M = **$2–4 par question** en flagship ; cache de 320 Go pour un 70B (section 105) |
| — | Latence du prefill : des dizaines de secondes pour 1M tokens |

Verdict : **gadget pour le RAG documentaire, utile pour l'analyse de gros dossiers ponctuels** (« analyse ces 200 pages de specs et compare »). Pour ton usage quotidien (questions FR sur tes guides), le pipeline retrieval + rerank + 20–30K tokens de contexte reste 10 à 100× moins cher et plus précis. Le 1M est un outil d'analyse, pas une architecture.

## 116. Le prompt système d'un RAG : le template

Le prompt système fait la moitié de la qualité finale. Template recommandé (français, ton direct) :

```
Tu es un assistant technique francophone. Tu réponds UNIQUEMENT à partir
des extraits fournis ci-dessous.

Règles :
1. Cite chaque affirmation : [extrait N].
2. Si les extraits ne contiennent pas la réponse, dis exactement :
   « Je ne trouve pas cette information dans la documentation fournie. »
   N'invente rien, ne complète pas avec tes connaissances générales.
3. Réponds en français, de façon structurée : diagnostic, cause probable,
   procédure, vérification.
4. Les codes d'erreur, références de pièces et valeurs numériques doivent
   être recopiés à l'identique des extraits.

Extraits :
[extrait 1] (source : huawei_s310_guide.md, section 12) ...
[extrait 2] ...
Question : ...
```

Points clés : l'interdiction d'halluciner doit être **explicite et formulée comme une phrase à recopier** (les modèles suivent mieux une formulation imposée qu'une consigne abstraite) ; les sources sont numérotées pour la citation ; le format de réponse est imposé (diagnostic → vérification).

## 117. Benchmarks : MTEB et ses cousins

**MTEB** (Massive Text Embedding Benchmark) : 41 tâches d'embedding en anglais — retrieval, re-classement, clustering, classification, similarité (STS), résumé. C'est la vitrine des embeddings généralistes. **MMTEB** en est la version multilingue (dont le français), **CMTEB** la version chinoise.

Comment lire un score MTEB :

- Le score agrégé (ex : Qwen3-Embedding-8B : 75,22 en English v2) est une **moyenne** sur 41 tâches. Un modèle peut être excellent en retrieval et médiocre en clustering — la moyenne cache tout.
- Pour un RAG, seule la sous-partie **Retrieval** compte vraiment. Nemotron-3-Embed annonce 75,45 en **MMTEB Retrieval** (pas en MTEB global) : c'est la bonne granularité à comparer.
- Les écarts de 1–2 points entre modèles proches sont du bruit (splits, seeds). Un écart devient significatif à partir de ~3–5 points sur la même sous-tâche.

**MIRACL** (18 langues, retrieval) : plus proche d'un usage RAG multilingue réel que MTEB. Le 66,83 de jina-reranker-v3 sur MIRACL est un bon signal pour le français.

**RTEB** (2026, récent) : nouveau benchmark retrieval. Jeune = moins éprouvé, splits moins audités — voir l'alerte Voyage (section 6). À suivre, pas encore une référence.

## 118. Benchmarks : SWE-bench, Arena, GPQA et les autres

Le décodeur des benchmarks LLM que tu croiseras dans les fiches :

| Benchmark | Mesure | Exemples 2026 (fournisseurs) |
|---|---|---|
| **SWE-bench Verified** | Résoudre de vrais tickets GitHub (code) | Gemini 3.1 Pro 80,6 % ; DeepSeek V4-Pro 80,6 % ; Opus 4.8 88,6 % (non vérifié) |
| **GPQA Diamond** | QCM scientifiques niveau expert | Gemini 3.1 Pro 94,3 % |
| **HLE** (Humanity's Last Exam) | Questions académiques extrêmes | Gemini 3.1 Pro 44,4 % |
| **Terminal-Bench 2.0** | Tâches agentiques en terminal | Gemini 3.1 Pro 68,5 % |
| **ARC-AGI-2** | Raisonnement abstrait (nouveauté 2025) | Gemini 3.1 Pro 77,1 % |
| **BrowseComp** | Recherche web profonde | Gemini 3.1 Pro 85,9 % |
| **Codeforces** | Compétition d'algorithmique (Elo) | DeepSeek V4-Pro Max 3206 |
| **LMArena** | Préférence humaine (votes aveugles) | — |
| **Artificial Analysis** | Indépendant : qualité + vitesse + prix | Intelligence Index, tok/s, $/M pondéré |

Deux remarques :

