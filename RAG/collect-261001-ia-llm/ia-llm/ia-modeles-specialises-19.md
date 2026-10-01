---
id: collect-261001-ia-llm/ia-llm/ia-modeles-specialises-19
title: "Encyclopédie des modèles IA — Volume 3"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Google", "Huawei", "Moonshot", "OpenAI"]
dates: ["2026-09-24"]
keywords: ["apache", "awq", "benchmark", "deepseek", "embedding", "embeddings", "fine-tuning", "fp4", "fp8", "gemini", "gptq", "gqa"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_specialises.md
source_anchor: ""
source_lines: [1970, 2044]
sha256: 942fbbf0d1b39d14ee4038e8db613028c9aca35caf9f5fac3e962494fcb81ede
---

# Encyclopédie des modèles IA — Volume 3

**Q5.** Pourquoi le MoE est-il excellent à gros batch mais inefficace à batch=1 ?

**Q6.** Classe MHA, MQA, GQA, MLA par taille de KV cache croissante et explique en une phrase chaque mécanisme.

**Q7.** Tu as une RTX 4090 (24 Go). Tu veux servir un 27B en Q4_K_M (18 Go). Combien de VRAM reste-t-il pour le KV cache, et combien de tokens de contexte (FP16, 48 couches, GQA-8, head_dim 128) peux-tu tenir ?

**Q8.** Quelle est la différence entre GPTQ et AWQ, et dans quel cas préférer le checkpoint officiel FP8/INT4 à un PTQ communautaire ?

**Q9.** Tu veux que ton RAG « connaisse » tes guides Huawei ET qu'il réponde toujours avec la structure « Diagnostic → Cause → Procédure → Vérification ». RAG, LoRA, ou les deux ? Justifie.

**Q10.** Qu'est-ce que HyDE, et dans quel cas l'activer (ou pas) ?

**Q11.** Ton corpus contient 12 620 commandes CLI Huawei (1 fichier .md par commande). Propose une stratégie de chunking et justifie-la.

**Q12.** Pourquoi faut-il mettre le prompt système et les instructions AVANT la question dans le prompt, dans le contexte du prompt caching ?

**Q13.** Un fournisseur annonce « 78,46 au benchmark RTEB » pour sa famille d'embeddings. Cite trois vérifications de la checklist (section 120) avant de prendre ce chiffre au sérieux.

**Q14.** Compare le coût de 1000 questions/mois en mode « flagship API » vs « pipeline éco » (ordres de grandeur), et cite les deux leviers qui réduisent le plus la facture.

**Q15.** Sora 2 a vu son API fermée le 24/09/2026. Quelle est la réserve méthodologique sur ce statut, et quelle conséquence pratique en tires-tu pour un nouveau projet ?

## 127. Réponses du quiz

**R1.** Point faible : modèle de décembre 2023, deux générations derrière (Gemini Embedding 2, Nemotron-3, Jina v5) ; pas de Matryoshka natif. Pas urgent : il est très bon marché (~$0,02/M), stable, ton index existe déjà — on ne migre que si un challenger gagne +5 pts de rappel@5 sur tes 50 questions FR (section 95).

**R2.** CC BY-NC 4.0 interdit l'usage commercial. Pour un RAG personnel : aucun problème, usage non commercial. Pour un outil d'entreprise : illicite sans licence commerciale — il faut basculer sur Apache 2.0 (Qwen3-Embedding-8B, BGE-M3).

**R3.** Par token : 2 × 32 × 8 × 128 × 2 = 131 072 octets = 128 Kio. À 32K tokens : 128 Kio × 32 768 = 4 294 967 296 octets ≈ **4 Go**.

**R4.** Charger : 35B paramètres (~70 Go FP16, ~18–20 Go en INT4). Calcul : 3B actifs par token. Conséquence : la VRAM suit le total (lourd à charger), la vitesse suit les actifs (rapide) — le MoE exige du batch et de l'expert parallelism pour être rentable (section 74).

**R5.** À batch=1, chaque token active des experts différents : on lit beaucoup de poids pour peu de calcul utile. À gros batch, les mêmes experts servent plusieurs tokens simultanément : le coût de lecture des poids s'amortit sur le batch.

**R6.** Ordre croissant : MLA (~7 % du MHA) < MQA (~3 %… en fait MQA < MLA selon les configs ; l'ordre typique est MQA < MLA < GQA < MHA). MHA : 1 tête KV par tête Q. MQA : 1 seule tête KV partagée. GQA : groupes de têtes Q partagent des têtes KV. MLA : K/V compressés en vecteur latent. (L'ordre exact MQA vs MLA dépend des dimensions — ce qui compte c'est le principe.)

**R7.** Reste : 24 − 18 = 6 Go, moins ~2 Go de marge moteur → ~4 Go utiles. Par token : 2 × 48 × 8 × 128 × 2 = 196 608 octets = 192 Kio. 4 Go / 192 Kio ≈ **21 800 tokens** de contexte. Suffisant pour un RAG standard (10–20 chunks).

**R8.** GPTQ compense l'erreur d'arrondi couche par couche (Hessien) ; AWQ protège les poids saillants par scaling avant quantization. Préférer le checkpoint officiel basse précision (QAT/natif : DeepSeek V4 FP4, Kimi K2 INT4, Gemma NVFP4) quand il existe : qualité meilleure à taille égale qu'un PTQ appliqué après coup.

**R9.** Les deux : **RAG pour les faits** (guides qui changent, réponses citables, 25 000 lignes impossibles à « apprendre » fiablement) + **LoRA pour le comportement** (le format de réponse standardisé s'apprend très bien en PEFT sur ~500 exemples). Le fine-tuning seul hallucinerait des faits avec aplomb.

**R10.** HyDE = embedder une réponse hypothétique générée par le LLM plutôt que la question brute, pour chercher dans l'espace des réponses. L'activer pour les questions courtes/vagues ou quand le vocabulaire question ≠ vocabulaire doc ; inutile pour les questions précises (« code C6000 ») — ça coûte un appel LLM.

**R11.** Ne pas chunker : 1 commande = 1 chunk (l'unité naturelle est déjà parfaite). Indexer chaque .md tel quel avec métadonnées `commande`, `famille`, `view`, `niveau`. Ne jamais couper les tableaux de paramètres. Volume : 12 620 chunks ≈ 52 Mo d'index à 1024 dims — négligeable.

**R12.** Le prompt caching fonctionne par **préfixe** : le cache n'est réutilisé que si le début du prompt est identique. Le stable (système, instructions) en premier = 100 % de hit ; la question variable en dernier. Inverser l'ordre invaliderait le cache à chaque requête.

**R13.** (1) Le score concerne-t-il ce checkpoint précis ou la famille ? (2) Le benchmark est-il standard et l'accès au split est-il sain (alerte Voyage/RTEB) ? (3) Existe-t-il une mesure indépendante qui confirme l'ordre de grandeur ?

**R14.** Flagship : ~$0,10/question → ~$100/mois pour 1000 questions. Éco : ~$0,005–0,015/question → ~$5–15/mois. Leviers : (1) le choix du générateur (90 % du coût — passer en local = ~0 €) ; (2) le prompt caching (−90 % sur le préfixe en cache).

**R15.** Réserve : le statut repose sur un miroir GitHub des docs OpenAI, pas sur openai.com ouvert en direct — à confirmer avant d'en faire un fait définitif. Conséquence : ne pas intégrer Sora à un nouveau projet ; choisir Seedance 2.0, Kling 3.0 ou Runway selon le besoin.

## 128. Checklist finale : ton RAG en 2026, étape par étape

L'ordre d'exécution recommandé — ne passe à l'étape suivante que si la précédente plafonne :

- [ ] **1. Corpus propre.** Tes guides sont déjà en markdown structuré — c'est 50 % du travail. Vérifie : titres hiérarchisés, tableaux non coupés, pas de résidus de scrape.
- [ ] **2. Chunking par sections** (sections 109–111). Frontières sur les `##`, garde-fous 200–2500 tokens, tableaux atomiques, 1 commande = 1 chunk pour les CLI.
- [ ] **3. Métadonnées** (section 112). `famille`, `type`, `section`, `source`, `date` — le filtrage hybride avant tout changement d'embedding.
- [ ] **4. Embedding actuel** (`text-embedding-3-small`). Indexe tout, mesure le rappel@5 sur 50 questions FR réelles.
- [ ] **5. Si rappel@5 < 90 %** : essaie HyDE (section 113) puis un reranker 0,6B local (section 98). Re-mesure.
- [ ] **6. Si ça plafonne encore** : challenger d'embedding (section 95) ou fine-tuning contrastif (section 92). Re-mesure.
- [ ] **7. Générateur** : commence en API éco (Qwen tiers) ou local 27B Q4 selon ton hardware. Prompt système avec citations obligatoires (section 116).
- [ ] **8. Évaluation continue** : garde ton jeu de 50–100 questions versionné ; rejoue-le à chaque changement (modèle, chunking, prompt). En batch API (−50 %).
- [ ] **9. Optimise les coûts** : prompt caching (stable en premier), reranker local, batch pour l'éval (sections 122–124).
- [ ] **10. Durcis** : « je ne sais pas » obligatoire, citations vérifiées, logs des questions sans réponse → nouveaux chunks ou corrections du corpus.

Règle d'or : **mesure avant d'optimiser**. Chaque étape se valide par un chiffre (rappel@5, nDCG, taux de « je ne sais pas » légitimes), pas par une impression.

## 129. Index vectoriels : Qdrant, pgvector, FAISS

Le vecteur ne sert à rien sans un index pour chercher dedans vite. Trois options :

