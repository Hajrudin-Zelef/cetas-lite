---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-22
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Glasswing", "Google", "Meta", "Microsoft", "MiniMax", "Nvidia", "OpenAI", "Poolside", "United States", "xAI"]
dates: ["2026-02-19", "2026-09-02", "2026-09-03", "2026-09-09", "2026-09-11", "2026-09-27", "2026-12-31", "2027-02-04"]
keywords: ["agent", "agentic", "agents", "agi", "asl", "astra", "attention", "benchmark", "benchmarks", "blackwell", "chatgpt", "claude"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [1717, 1789]
sha256: 9f87d17154d288790be9951f4159df5bcf2aa525f43d0e8b286f3cc7d429a043
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident

1. **Agent** : système IA qui perçoit, raisonne, utilise des outils et agit de façon autonome sur plusieurs étapes pour atteindre un objectif.
2. **Agentic coding** : génération de code par agents autonomes (planification, édition multi-fichiers, exécution de tests, boucle de correction).
3. **ARC-AGI** : benchmark de raisonnement abstrait (grilles) ; la version 3 teste les agents interactifs.
4. **ASL (AI Safety Level)** : échelle de sécurité d'Anthropic ; ASL-3 = modèle à capacités dangereuses potentielles, déploiement renforcé.
5. **Batch API** : traitement asynchrone de requêtes par lots, généralement à −50 % du prix temps réel.
6. **Benchmark** : test standardisé mesurant une capacité ; « vendor-déclaré » = mesuré par le vendeur lui-même, à manier avec prudence.
7. **Canary test** : déploiement test furtif d'un modèle non annoncé, repéré par la communauté (ex. Sonnet 5.5 fin sept. 2026).
8. **Checkpoint** : snapshot des poids d'un modèle à un stade d'entraînement ; on peut « re-télécharger » un checkpoint mis à jour (ex. Laguna S 2.1, août 2026).
9. **Computer use** : capacité d'un modèle à opérer un ordinateur comme un humain (pixels d'écran, souris, clavier, navigateur).
10. **Context compaction** : résumé/compression automatique de l'historique d'une conversation pour tenir dans la fenêtre de contexte.
11. **Context window (fenêtre de contexte)** : quantité maximale de tokens (entrée + sortie) que le modèle peut traiter d'un coup ; 1M ≈ un gros roman.
12. **CyberGym / ExploitBench** : benchmarks de cybersécurité offensive (découverte/exploitation de vulnérabilités) — au cœur des classements « Critical ».
13. **Daybreak (Red)** : programme d'accès contrôlé d'OpenAI aux capacités cyber offensives (variantes -Cyber).
14. **Deep Think / thinking** : mode de raisonnement long et parallèle ; en 2026, devenu un curseur d'effort réglable (low → max).
15. **Distillation** : entraîner un petit modèle à imiter un grand ; « codistillation » (Maverick ← Behemoth) en est une variante.
16. **ELO (Arena)** : score de préférence humaine sur LMArena ; mesure la popularité, pas une capacité brute.
17. **Effort (paramètre)** : curseur low/medium/high/xhigh/max arbitrant qualité, latence et coût du raisonnement.
18. **Fairwind Program** : programme d'accès restreint de Google aux variantes Gemini Flash Cyber.
19. **Fine-tuning** : réentraînement d'un modèle sur des données spécifiques ; absent de l'offre publique Anthropic 2026, non supporté sur GPT-6 Sol/Luna.
20. **FP8 / NVFP4** : formats de quantification sur 8/4 bits (NVFP4 = format NVIDIA Blackwell) réduisant mémoire et coût d'inférence.
21. **Frontier model** : modèle à la frontière des capacités connues ; « frontier » est aussi un argument marketing.
22. **GGUF** : format de quantification communautaire pour exécution locale (llama.cpp, Ollama).
23. **Glasswing (Project)** : programme d'accès restreint d'Anthropic à Claude Mythos (~40 organisations, NDA).
24. **GQA (Grouped-Query Attention)** : attention avec moins de têtes clé-valeur que de têtes query — réduit le KV cache.
25. **HLE (Humanity's Last Exam)** : benchmark de questions académiques extrêmes ; score de référence du raisonnement.
26. **Hybrid reasoning** : modèle pouvant alterner mode réponse directe et mode raisonnement explicite (ex. Hermes, Hy3).
27. **iRoPE** : variante d'encodage positionnel de Llama 4 (couches sans embeddings positionnels entrelacées).
28. **KV cache** : mémoire des clés/valeurs d'attention mise en cache entre tokens ; son coût domine l'inférence long-contexte.
29. **LTS (Long-Term Support)** : version garantie dans le temps (ex. GPT-5.3-Codex, garanti jusqu'au 04/02/2027) — une première chez OpenAI.
30. **Mamba** : architecture d'espace d'états (alternative/complément au Transformer), efficace en long-contexte ; au cœur des Nemotron.
31. **MCP (Model Context Protocol)** : protocole ouvert d'Anthropic pour brancher des outils/sources aux modèles ; supporté nativement par Opus 5, Astra, Grok 4.1 Fast.
32. **MoE (Mixture of Experts)** : architecture où seuls quelques « experts » (sous-réseaux) s'activent par token ; ex. 17B actifs / 400B total (Maverick).
33. **MTP (Multi-Token Prediction)** : prédiction de plusieurs tokens d'un coup — accélère le décodage spéculatif (Nemotron, Gemma 4, MiniMax).
34. **Muon** : optimiseur d'entraînement alternatif à AdamW, utilisé par Poolside (et DeepSeek V4).
35. **Omnimodal** : modèle traitant nativement texte/image/audio/vidéo de bout en bout (ex. GPT-5.5, Gemini Omni).
36. **Open-weight** : poids publiés (téléchargeables, fine-tunables) sous licence plus ou moins permissive — ≠ « open source » au sens OSI strict.
37. **OSWorld** : benchmark d'agents computer-use sur OS réel ; devenu le standard du « bureau autonome ».
38. **Preparedness Framework** : cadre d'évaluation des risques d'OpenAI ; GPT-6 Astra est son premier modèle classé « Critical » (cyber).
39. **Prompt caching** : réutilisation du préfixe déjà traité, facturée ~10 % du prix input (voire 2,5 % chez Fable 5.1).
40. **SWE-bench** : benchmark de résolution de vraies issues GitHub ; la version « Pro »/« Verified » est la référence du code agentique en 2026.

## 142. Quiz — 15 questions + réponses

**Q1.** Que signifient « Sol », « Terra » et « Luna » chez OpenAI ?
**R1.** Ce sont les noms officiels publics des trois tiers de la famille GPT-5.6 (puis GPT-6) : Sol = flagship/capacité max, Terra = équilibré, Luna = économique. Pas des noms de code internes.

**Q2.** « GPT Pro » est-il un modèle ?
**R2.** Non. Le terme recouvre l'abonnement ChatGPT Pro ($100/$200 par mois — le $200 suspendu aux nouvelles souscriptions le 10–11/09/2026) et les variantes « -Pro » des modèles (5.4 Pro, 5.5 Pro, Astra Pro…).

**Q3.** Quel est le premier modèle OpenAI classé « Critical » en cybersécurité, et que change ce classement ?
**R3.** GPT-6 Astra (03/09/2026). Ses capacités offensives les plus avancées sont placées en accès contrôlé (Daybreak Red).

**Q4.** Quelle est la différence entre Muse et Muse Spark ?
**R4.** Muse = l'agent personnel de Meta (produit, lancé 08–09/09/2026, US only) ; Muse Spark = la famille de modèles qui le propulse (versions 1.0 → 1.3).

**Q5.** Pourquoi Meta a-t-il abandonné l'open-weight en 2026 ?
**R5.** Stratégie assumée depuis avril 2026 : pivot vers des modèles propriétaires sous la marque Muse (Meta Superintelligence Labs). Llama 4 Scout/Maverick (avril 2025) restent les derniers poids ouverts Meta.

**Q6.** Qu'est-ce que Claude Mythos, et pourquoi n'est-il pas public ?
**R6.** Le modèle le plus puissant d'Anthropic (nom de code Capybara), volontairement non publié pour son potentiel cyber destructeur (découverte autonome de zero-days). Accès restreint via Project Glasswing (~40 organisations, NDA).

**Q7.** « Claude Sonnet 4.8 » et « Claude Haiku 5 » existent-ils ?
**R7.** Non, les deux sont introuvables : la lignée Sonnet est 4.5 → 4.6 → 5 (5.5 annoncé), et Haiku passe de 4.5 à 5.5 (annoncé). Le catalogue Vertex rejette `claude-haiku-5`.

**Q8.** Quelle est la particularité tarifaire de Claude Fable 5.1 sur le cache ?
**R8.** Lectures de cache à 0,025× (0,25 $/M, −75 % vs le standard 0,1×) — la principale baisse de coût, jusqu'à −45 % sur charges agentiques.

**Q9.** « Gemini 3.5 Pro » et « Gemini 4 » existent-ils au 27/09/2026 ?
**R9.** Non. La ligne Pro est restée à Gemini 3.1 Pro Preview depuis le 19/02/2026 ; Google n'a itéré que sur les Flash (3.5 → 3.8).

**Q10.** Quel est le modèle texte Gemini recommandé par Google au 27/09/2026, et son prix ?
**R10.** Gemini 3.8 Flash (02/09/2026), en promo à $0,75/$3,75 par MTok jusqu'au 31/12/2026.

