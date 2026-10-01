---
id: collect-261001-ia-llm/ia-llm/mistral-medium-3-5-77-6-swe-bench-1-50-api-2026-1
title: "mistral-medium-3-5-77-6-swe-bench-1-50-api-2026"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "Hugging Face", "Mistral", "Nvidia", "OpenAI", "United States"]
dates: []
keywords: ["mistral", "agent", "agents", "benchmark", "benchmarks", "claude", "deepseek", "gemini", "gpu", "llama", "moe", "multimodal"]
source: docs/RAG/collect-261001-ia-llm/mistral-medium-3-5-77-6-swe-bench-1-50-api-2026.md
source_anchor: ""
source_lines: [1, 33]
sha256: 64d7ee2f731aceffee2f4cdf8ddb3ffa83d05d33774f06cfb13608becd40e5f8
---

# mistral-medium-3-5-77-6-swe-bench-1-50-api-2026

Paris, 29 avril 2026 – La start-up française **Mistral AI** a frappé un grand coup hier en dévoilant simultanément **Mistral Medium 3.5**, un modèle dense de 128 milliards de paramètres en poids ouverts, un nouvel agent de codage à distance baptisé **Vibe**, et un mode **Work** dans son assistant grand public **Le Chat**. L’annonce, publiée sur le blog officiel de l’entreprise et confirmée par sa fiche Hugging Face datée du 29 avril 2026, marque la consolidation la plus ambitieuse jamais tentée par le champion européen de l’intelligence artificielle générative. Sur le banc d’essai **SWE-Bench Verified**, Medium 3.5 atteint **77,6 %**, et sur le test agentique **τ³-Telecom**, le score grimpe à **91,4 %**, des chiffres qui placent la jeune pousse parisienne au coude-à-coude avec les meilleurs modèles propriétaires américains, à un tarif API près de trois fois inférieur.

Dix-huit mois après son tour de table record de **1,7 milliard d’euros** en série C valorisant l’entreprise à **11,7 milliards d’euros post-money**, mené par ASML avec la participation de DST Global, Andreessen Horowitz, Bpifrance, General Catalyst, Index Ventures, Lightspeed et NVIDIA, Mistral devait prouver qu’il pouvait transformer ce capital en feuille de route produit cohérente. Le lancement du 29 avril est sa réponse : un seul modèle généraliste — commercialisé sous la référence de build v26.04 selon la fiche Hugging Face — qui remplace simultanément Magistral et Medium 3.1 dans Le Chat ainsi que Devstral 2 dans Vibe, et qui devient l’ossature d’une suite intégrée pour le grand public, les développeurs et les entreprises. Cette stratégie de modèle unique tranche radicalement avec la fragmentation que l’on observe chez OpenAI ou Google.

## Mistral Medium 3.5 : la fiche technique d’un modèle dense de 128 milliards de paramètres

Le cœur du lancement est **Mistral Medium 3.5**, un modèle de langage *dense*, c’est-à-dire *non* Mixture-of-Experts, ce qui le distingue immédiatement de DeepSeek V4, Mixtral et de la plupart des nouveaux modèles à très grande échelle. Avec ses 128 milliards de paramètres tous activés à chaque inférence, Medium 3.5 sacrifie l’efficacité économique du sparse pour la stabilité et la prévisibilité du dense. La fiche Hugging Face publiée le 29 avril 2026 confirme une fenêtre de contexte de **256 000 tokens**, soit environ 200 000 mots, suffisamment pour ingérer un rapport annuel complet du CAC 40 ou un code source de plusieurs centaines de fichiers en une seule requête. Selon le guide de déploiement publié par Clore.ai en août 2026, ce transformeur dense de 128 milliards de paramètres affiche un temps jusqu’au premier token (TTFT) compris entre 50 et 250 millisecondes en mode « instant », une latence qui le rend directement exploitable pour des interfaces conversationnelles temps réel.

Le modèle est **multimodal** : il accepte texte et image en entrée et produit du texte. Il intègre nativement l’**appel de fonctions** (function calling) et la sortie JSON structurée, deux fonctionnalités désormais indispensables pour les workflows agentiques. Mistral introduit également un paramètre original : un **réglage de l’effort de raisonnement** configurable par requête, qui permet au développeur d’arbitrer en temps réel entre latence et profondeur d’analyse, sans avoir à basculer entre un modèle « rapide » et un modèle « lent » comme c’est le cas avec OpenAI o3 et GPT-5.5.

Sur le plan opérationnel, Mistral revendique la possibilité de *self-host* Medium 3.5 sur **seulement quatre GPU**, un argument clé pour les administrations européennes et les grandes entreprises soumises à des obligations de souveraineté des données. Le 1er juillet 2026, NVIDIA a publié sur Hugging Face une version quantisée **Mistral-Medium-3.5-128B-NVFP4**, conservant la fenêtre de contexte complète de 256 000 tokens tout en réduisant l’empreinte mémoire nécessaire au déploiement on-premises. Le modèle est distribué sous une **licence MIT modifiée**, ce qui le maintient dans la catégorie des *open weights* mais avec des restrictions d’usage commercial qui le rapprochent davantage de Llama 4 que d’un véritable open source au sens OSI. Les poids sont disponibles publiquement sur Hugging Face dès le jour du lancement.

## SWE-Bench 77,6 %, τ³-Telecom 91,4 % : pourquoi ces scores changent la donne

Les chiffres communiqués par Mistral le 29 avril 2026 placent Medium 3.5 dans le peloton de tête des modèles agentiques. Sur **SWE-Bench Verified**, le test de référence qui mesure la capacité d’un modèle à résoudre des bugs réels issus de dépôts GitHub open source en autonomie, le score de **77,6 %** dépasse celui de Mistral Large 3 et talonne directement les modèles propriétaires les mieux placés du marché. Ce résultat est d’autant plus significatif qu’il provient d’un modèle *en poids ouverts* : aucun modèle ouvert n’avait jusqu’à présent franchi ce seuil symbolique avec une licence permissive sur quatre GPU. Ce score de 77,6 % a été confirmé indépendamment en juillet 2026 par le classement AI Rankings, qui situe Medium 3.5 nettement devant son prédécesseur Devstral 2 (environ 72,2 % sur le même test), validant ainsi le remplacement opéré au sein de Vibe.

Le score de **91,4 %** sur **τ³-Telecom** est encore plus révélateur de l’orientation produit. τ³-Telecom, troisième itération du benchmark tau-bench développé pour évaluer les agents conversationnels en situation client réelle (assistance technique, gestion de contrats, escalade), est devenu en 2025 le standard de référence pour les déploiements en centre de contacts. Un score supérieur à 90 % signifie que l’agent gère sans intervention humaine la quasi-totalité des conversations de niveau 1 et 2. Selon les retours préliminaires de la communauté testeur, Medium 3.5 s’impose comme l’un des deux ou trois modèles capables de soutenir ce seuil de manière reproductible.

### Comparaison des benchmarks face aux modèles concurrents

| Modèle | Éditeur | Paramètres | Contexte | SWE-Bench Verified | Prix entrée / 1M tokens | Prix sortie / 1M tokens | 
|---|---|---|---|---|---|---|
| Mistral Medium 3.5 | Mistral AI (FR) | 128 Md (dense) | 256K | 77,6 % | 1,50 $ | 7,50 $ | 
| Mistral Large 3 | Mistral AI (FR) | Non divulgué | 256K | ~71 % | 2,00 $ | 6,00 $ | 
| Claude Opus 4.7 | Anthropic (US) | Non divulgué | 200K-1M | ~80 % | 15,00 $ | 75,00 $ | 
| GPT-5.5 | OpenAI (US) | Non divulgué | 400K | ~78 % | 5,00 $ | 15,00 $ | 
| Gemini 3.5 Flash | Google (US) | Non divulgué | 1M | ~72 % | 2,00 $ | 8,00 $ | 
| DeepSeek V4 | DeepSeek (CN) | 1,6 T (MoE) | 128K | ~74 % | 0,55 $ | 1,74 $ | 

Le tarif API de **1,50 dollar par million de tokens en entrée** et **7,50 dollars par million de tokens en sortie** positionne Medium 3.5 sur un segment économique radicalement différent de Claude Opus 4.7 et GPT-5.5. Ce tarif a d’ailleurs évolué depuis le lancement : selon le suivi de ConductAtlas daté du 12 au 15 août 2026, la fiche modèle est passée d’une tarification initiale en euros (1,25 €/6,40 € par million de tokens) à la grille actuelle de 1,50 $/7,50 $, un ajustement qui confirme la stratégie de Mistral d’aligner ses prix sur le marché américain en dollars. Pour un usage d’agent de codage produisant en moyenne 200 000 tokens de sortie par tâche, la facture Mistral s’élève à environ 1,50 dollar contre 15 dollars chez Anthropic. Sur une équipe de 50 développeurs effectuant 20 tâches par jour, l’écart annuel dépasse les 1,2 million d’euros.

## Vibe : l’agent de codage à distance qui défie Cursor et Claude Code

