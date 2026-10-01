---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-18
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "ByteDance", "DeepSeek", "Google", "Meta", "Moonshot", "OpenAI", "Poolside", "Z.ai", "xAI"]
dates: ["2026-06-24", "2026-09-01", "2026-09-02", "2026-09-03", "2026-09-05", "2026-09-16", "2026-09-18", "2026-09-20", "2026-09-22", "2026-09-23", "2026-09-24", "2026-09-25", "2026-09-27"]
keywords: ["agents", "apache", "astra", "benchmark", "benchmarks", "claude", "deepseek", "diffusion", "fable 5", "gemini", "gemini 3.8", "gemini 4"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [2238, 2298]
sha256: 9a7787112b0f651ee70f4c57aece15901eaa4f956c4599114a7f3f212bd7c395
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

| Point | État au 27/09/2026 |
|---|---|
| Annonce | **Semaine du 19–25/09/2026** : Anthropic indique que **Sonnet 5.5 et Haiku 5.5 arrivent dans les prochaines semaines** (digest hebdomadaire EveryDev.ai, citant Anthropic). |
| Contexte déjà sorti | **Claude Opus 5.5** est sorti la semaine du 22/09/2026 : $4/$20 par million de tokens (entrée/sortie) contre $5/$25 pour Opus 5, lectures de cache à $0.20/M (contre $0.50), vitesse de sortie annoncée +30 % par rapport à Opus 5. Performances comparables au haut de gamme Fable 5.1 selon Anthropic, à coût d'exploitation inférieur. Prix vérifiés le 27/09/2026 via le digest EveryDev.ai citant Anthropic. |
| Autre fait officiel récent | **Claude Fable 5.1 et Claude Mythos 5.1**, lancés le **01/09/2026**. Mythos 5.1 est réservé aux institutions vérifiées via des programmes d'accès de confiance (cybersécurité, sciences de la vie) — c'est l'exemple type du « capability-tiered gating » de septembre 2026 : les capacités les plus risquées sont cloisonnées par l'éditeur, pas par l'utilisateur. |
| Ce qui n'est pas annoncé | **Sonnet 4.8** : rien d'officialisé au 27/09/2026 (reste dans la liste des « non trouvés » du guide). **Haiku 5** (sans le « .5 ») : rien d'officialisé ; seul Haiku 5.5 a été évoqué comme à venir. |

Sources (vérifiées le 27/09/2026) : EveryDev.ai « Weekly AI Dev News Digest: September 19 - 25, 2026 », TechTarget (24/09/2026, analyse des lancements OpenAI/Anthropic), wowtale.net (05/09/2026, gating par capacités des quatre labs).

### 159.4. OpenAI — GPT-6 Sol et Luna sortis le 22/09/2026 ; teaser DevDay ; GPT-6 Terra non annoncé

| Point | État au 27/09/2026 |
|---|---|
| Sorti le 22/09/2026 | **GPT-6 Sol** (code et agents, ~2× moins d'erreurs que son prédécesseur selon OpenAI, fiabilité niveau Astra) et **GPT-6 Luna** (volume bureautique : résumés, extraction de champs, questions courtes), **à moitié prix de la série 5.6** — économie attribuée par OpenAI à des améliorations de caching et d'inférence. Les prix par token ne figurent pas dans l'annonce. |
| Sorti le 03/09/2026 | **GPT-6 Astra**, décrit par OpenAI comme son système le plus capable et le plus aligné à ce jour. |
| Teaser officiel | OpenAI (via « Thibault », équipe OpenAI) a teasé des annonces **de niveau DevDay** pour la semaine du 16/09/2026 (digest AI Builders du 16/09/2026). Au 27/09/2026 : teaser confirmé, contenu non détaillé dans les sources consultées — à vérifier sur le blog OpenAI. |
| Non annoncé | **GPT-6 Terra** : rien d'officialisé au 27/09/2026. Le nom circule dans les listes de rumeurs mais n'apparaît dans aucune annonce officielle trouvée. **Ne pas ajouter ce modèle** (consigne explicite du guide). |
| Accès restreint | Les capacités de cybersécurité d'Astra ayant franchi le seuil « Critical » du Preparedness Framework d'OpenAI sont réservées aux entreprises vérifiées via le programme Daybreak (sur dossier). |

Sources (vérifiées le 27/09/2026) : EveryDev.ai (digest 19–25/09/2026, citant TechCrunch et OpenAI), startupfortune.com (semaine des quatre lancements), digest GitHub to-real/ai-builder-digest du 16/09/2026.

### 159.5. Meta — plus haut niveau de raisonnement RETENU (annoncé), Muse Spark 1.3 sorti

| Point | État au 27/09/2026 |
|---|---|
| Annoncé | Meta **retient son plus haut niveau de raisonnement** en attendant la fin de tests de sécurité supplémentaires (wowtale.net, 05/09/2026, dans le cadre de la couverture de la semaine des quatre lancements). C'est une annonce de roadmap : le niveau existe, sa diffusion générale est conditionnée aux tests. |
| Sorti le 02/09/2026 | **Muse Spark 1.3** (Meta), dans la vague de lancements de la semaine du 31/08 au 03/09/2026. |
| Non annoncé | **Llama 4 Behemoth** : rien d'officialisé au 27/09/2026. |

### 159.6. xAI — Grok 4.7 sorti ; rien d'annoncé au-delà

| Point | État au 27/09/2026 |
|---|---|
| Sorti | **Grok 4.7**, semaine du 19–25/09/2026. Tarification à deux étages, vérifiée le 27/09/2026 (digest EveryDev.ai citant xAI) : **en dessous de 200K tokens de prompt** : $2 entrée / $0.50 entrée en cache / $6 sortie par million ; **au-dessus de 200K** : $4 / $1 / $12. Fenêtre de contexte inchangée à 500K (identique à Grok 4.6) — seul le bout lointain de la fenêtre coûte plus cher. Aucune annonce de benchmark plaçant 4.7 devant Claude ou GPT-6 publiée par xAI. |
| Non annoncé | Aucune version ultérieure officialisée au 27/09/2026. |

### 159.7. Google — Gemini 3.8 Flash sorti ; Gemini 4 NON ANNONCÉ

| Point | État au 27/09/2026 |
|---|---|
| Sorti | **Gemini 3.8 Flash**, semaine du 31/08 au 03/09/2026 (vague des quatre lancements). Google a couplé son modèle orienté code à une variante cybersécurité distincte, réservée aux opérateurs gouvernementaux et d'infrastructures de confiance via le programme Fairwind. |
| Non annoncé | **Gemini 4** : rien d'officialisé au 27/09/2026. **Ne pas ajouter ce modèle** (consigne explicite du guide, comme GPT-6 Terra). |
| Non annoncé | **Veo 4** : rien d'officialisé au 27/09/2026. |

### 159.8. Chine : Alibaba, Moonshot, Zhipu, ByteDance — rien d'officialisé au-delà du sorti

| Éditeur | État au 27/09/2026 |
|---|---|
| Alibaba (Qwen) | Rien d'officialisé au-delà des modèles déjà sortis (Qwen 3.6/3.8, HY3/Hunyuan 3 en open-weight Apache 2.0). |
| Moonshot (Kimi) | Rien d'officialisé au-delà des modèles déjà sortis (K2.5). |
| Zhipu (GLM) | GLM 5.3 sorti (cité dans les comparatifs de benchmarks de septembre 2026). Rien d'officialisé au-delà. |
| ByteDance (Doubao) | **Doubao Seed 2.1 Pro et Turbo** sortis le **24/06/2026** (conférence Volcano Engine FORCE, 256K de contexte, Pro = deep-thinking, Turbo = moitié prix et basse latence). Tarifs Volcano Engine vérifiés le 27/09/2026 : Seed 2.1 Pro ¥6/¥30 par million (entrée/sortie), Turbo ¥3/¥15, entrée en cache ¥1.20 ; Seed 2.0 Pro ¥3.20/¥16 (entrées ≤ 32K). Via l'agrégateur ofox.ai (USD, vérifié le 27/09/2026) : Pro $0.884/$4.42, cache $0.177 ; Turbo $0.442/$2.212, cache $0.085. Rien d'officialisé au-delà de Seed 2.1 au 27/09/2026. |

### 159.9. RUMEURS non confirmées — à ne jamais mettre en prod

1. **RUMEUR non confirmée — Anthropic préparerait un nouveau modèle pour répondre à GPT-6 Astra.** Reuters a rapporté le 18/09/2026 qu'Anthropic « envisageait » un nouveau modèle, alors qu'Astra capterait ~13 % des dépenses IA business suivies par Ramp (contre ~8 % pour Claude Fable). Anthropic n'a rien confirmé. C'est une rumeur de presse, pas une annonce. (Sources : daylila.com 23/09/2026 citant Reuters ; aiagentsdirectory.com 20/09/2026.)
2. **RUMEUR non confirmée — fenêtre de sortie de DeepSeek V4.1 Pro.** Le billet orcarouter.ai évoque une « fenêtre qui se ferme le 30 septembre », mais DeepSeek n'a publié ni date ni fenêtre officielle. Le nom n'apparaît ni dans le changelog, ni sur la page pricing, ni dans un papier. Ne pas planifier de migration dessus.
3. **RUMEUR non confirmée — Poolside Malibu.** Nom circulant dans les listes de modèles non trouvés ; aucune annonce officielle au 27/09/2026.

### 159.10. Tableau récapitulatif — ce qui est annoncé vs ce qui ne l'est pas

