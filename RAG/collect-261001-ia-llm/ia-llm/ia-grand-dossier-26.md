---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-26
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Cohere", "DeepSeek", "Google", "Meta", "Microsoft", "Mistral", "Moonshot", "OpenAI", "Z.ai", "xAI"]
dates: []
keywords: ["agent", "agents", "apache", "attention", "benchmark", "benchmarks", "claude", "cohere", "compute", "deepseek", "distillation", "embedding"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [1727, 1806]
sha256: 1b113a7a8558a35dd7d8ccfc95c116923339b70d265af0318f61662d60837afa
---

# IA — Le grand dossier

1. **Le test-time compute a-t-il un plafond ?** Les lois de Snell et al. (2024) tiennent jusqu'à quel budget ? L'« overthinking » (dégradation au-delà d'un certain thinking) est-il un artefact ou une limite fondamentale ?
2. **Le RL avec récompenses vérifiables scale-t-il au-delà des maths et du code ?** Les domaines sans vérificateur automatique (droit, médecine, stratégie) suivront-ils la même courbe ?
3. **Les données synthétiques peuvent-elles remplacer les données humaines ?** À quel ratio le « model collapse » apparaît-il, et les filtres par vérificateurs suffisent-ils à l'éviter ?
4. **Existe-t-il une architecture post-Transformer ?** Mamba/SSM, les hybrides, JEPA : lequel (s'il y en a un) cassera la complexité quadratique sans perdre la qualité ?
5. **Combien vaut vraiment un paramètre ?** La distillation et la quantification repoussent sans cesse le ratio qualité/taille : où est le plancher ?
6. **L'énergie est-elle le vrai mur ?** Si le training plafonne mais que l'inférence explose, la contrainte 2027-2030 est-elle le **TWh** plutôt que le token ?
7. **L'open source rattrapera-t-il durablement le fermé ?** Le cycle « percée fermée → réplique ouverte en 12-24 mois » (pattern 1.15) est-il structurel ou conjoncturel ?
8. **Les agents sont-ils économiquement viables ?** Un agent qui coûte 100-1 000× une requête simple doit créer 100-1 000× plus de valeur : où est la preuve à grande échelle ?
9. **L'évaluation suit-elle les capacités ?** Les benchmarks saturent plus vite que les modèles ne progressent : comment mesurer ce qu'on ne sait plus tester ?
10. **Que se passe-t-il quand les modèles s'entraînent sur le web qu'ils ont eux-mêmes écrit ?** La boucle « IA → contenu web → training » est déjà enclenchée : enrichissement ou appauvrissement informationnel à long terme ?

*Chacune de ces questions est un sujet de veille : quand l'une d'elles bascule (réponse partielle, résultat expérimental), c'est un signal stratégique — bien plus fiable que les annonces produits.*

---

## 22. Checklist : lire une annonce IA sans se faire avoir (10 points)

> À appliquer à chaque « révolution » annoncée par communiqué de presse. Inspirée des règles d'hygiène de la section 15.

1. **Qui parle ?** Le labo lui-même (biaisé) ou une évaluation indépendante ? Un communiqué n'est pas un papier.
2. **Quels benchmarks ?** Sont-ils publics, reproductibles, non saturés ? Chercher ceux où le modèle **perd**.
3. **Comparé à quoi ?** À la version précédente du même labo (facile) ou au meilleur concurrent au même prix (le vrai test) ?
4. **Le chiffre est-il un fait ou une estimation ?** « Coût d'entraînement » : run final ou programme complet ? « Paramètres » : totaux ou actifs ?
5. **Qu'est-ce qui n'est PAS dit ?** Pas d'architecture, pas de données, pas de section limites = signal d'alerte (cas GPT-4, 2023).
6. **Le gain est-il du scaling ou de l'évaluation ?** Un meilleur prompt, un benchmark plus facile ou une métrique différente peuvent simuler un progrès.
7. **Quelle est la licence d'accès ?** API fermée, poids ouverts, et **quelle licence exactement** (section 11, idée reçue n° 12) ?
8. **Quel est le coût d'usage réel ?** Prix/1M tokens in/out, latence, besoin en thinking tokens — le TCO, pas le score.
9. **Qui a intérêt à cette annonce ?** Levée de fonds en cours ? Concurrence d'un rival la même semaine ? Les annonces sont des armes.
10. **Que disent les réplications indépendantes, 2 semaines plus tard ?** La vérité d'une annonce se mesure à sa **survie** aux tests tiers, pas à son lancement.

**Exemple d'application :** « Notre modèle bat GPT-5 sur 12 benchmarks » → exiger la liste des 12, vérifier la saturation, attendre LMArena et les tests indépendants, comparer les prix au token. En général, le tableau est moins rose que le communiqué — mais parfois (R1, janv. 2025), il est **meilleur** que prévu : la checklist protège des deux erreurs.

---

---

## 23. Fiches modèles 2026 : les 12 à connaître par cœur

> Ce que chaque modèle fait le mieux, sa licence, son positionnement. « À vérifier » au jour près pour les versions.

**23.1. DeepSeek R1 / V4 (DeepSeek, MIT).** Le raisonnant ouvert de référence. R1 (janv. 2025) = niveau o1 en MIT ; V4 (2026) = -75 % de prix API. **Usage :** raisonnement, maths, code — en local (distillations 7B-70B) ou API pas chère. Le choix par défaut du budget serré.

**23.2. Qwen3 (Alibaba, Apache 2.0).** La famille open la plus complète (0,6B → 235B MoE), hybride thinking/non-thinking, **Qwen3-Coder** excellent en code. **Usage :** le couteau suisse du self-hosting ; le 32B est le meilleur rapport qualité/VRAM de 2025-2026.

**23.3. Llama 4 (Meta, licence communautaire).** Scout/Maverick : multimodaux natifs, contextes longs. **Usage :** là où l'écosystème Meta compte (outillage, fine-tunes communautaires). Attention à la licence (restrictions vs Apache 2.0).

**23.4. Mistral Large 3 / Magistral (Mistral, Apache 2.0).** Large 3 = le dense européen efficace ; Magistral = le raisonnant ouvert (2025). **Usage :** le choix souveraineté UE + français (excellent en français). Devin/Mistral Compute pour l'inférence européenne.

**23.5. GPT-5.x (OpenAI, fermé/API).** La frontière généraliste, « intelligence unifiée » (plus de séparation thinking/non-thinking côté utilisateur). **Usage :** quand il faut le meilleur score brut et que le budget suit ; l'API la plus chère, l'écosystème le plus riche.

**23.6. Claude 4.x (Anthropic, fermé/API).** La référence **code et agents** (SWE-bench, tool use), style d'écriture soigné, fenêtres de contexte longues. **Usage :** agents de dev, rédaction technique, tâches où la prudence/qualité prime sur le prix.

**23.7. Gemini 2.5/3 (Google DeepMind, fermé/API).** Contexte géant (1-2M tokens), multimodalité native, intégration Google (Search, Workspace). **Usage :** « avaler » des corpus entiers d'un coup, vidéo, usage grand public via l'écosystème Google.

**23.8. Kimi K2 (Moonshot, ouvert — licence à vérifier).** Le champion chinois du **contexte long et des agents** (1T paramètres MoE annoncé, prix agressifs). **Usage :** alternative asiatique crédible pour agents et longs documents.

**23.9. GLM-4.5/5 (Zhipu, ouvert).** Fort en chinois/anglais, stratégie open agressive. **Usage :** corpus multilingues incluant le chinois ; écosystème en croissance.

**23.10. Grok (xAI, partiellement ouvert).** Accès temps réel à X, ton moins filtré, cluster Colossus derrière. **Usage :** veille temps réel, cas où les garde-fous des concurrents gênent (avec discernement).

**23.11. Les petits efficaces : Qwen3-8B, Mistral Small, Phi (Microsoft), Gemma (Google).** 3B-14B, quantifiables en 4-8 Go VRAM. **Usage :** classification, extraction, résumé, RAG local sur machine modeste — 80 % des tâches d'entreprise n'ont pas besoin de plus.

**23.12. Les embeddings : text-embedding-3-small/large (OpenAI), BGE-M3, E5, Cohere Embed.** Le modèle de Zelef (`text-embedding-3-small`) est un bon défaut ; **BGE-M3** (multilingue, 100+ langues) mérite un A/B test pour un corpus français. **Règle :** l'embedding se choisit sur VOTRE corpus (benchmark de retrieval maison), pas sur un leaderboard.

**Matrice de choix rapide :**

| Besoin | Premier choix | Alternative budget |
|---|---|---|
| Raisonnement dur (maths, code) | Claude 4.x / GPT-5.x (API) | R1-Distill 32B local |
| RAG d'entreprise FR | Mistral Large 3 / Qwen3-32B local | Qwen3-8B |
| Agent autonome | Claude 4.x (tool use) | Kimi K2 / Qwen3 |
| Volume massif, coût mini | DeepSeek V4 (API) | Modèle local quantifié |
| Souveraineté UE | Mistral (Apache 2.0) | — |
| Contexte ≥ 1M tokens | Gemini | Kimi K2 |
| On-device / offline | Qwen3-8B Q4 / Phi | Gemma |

---

## 24. Les 10 erreurs classiques d'un déploiement RAG (et comment les éviter)

