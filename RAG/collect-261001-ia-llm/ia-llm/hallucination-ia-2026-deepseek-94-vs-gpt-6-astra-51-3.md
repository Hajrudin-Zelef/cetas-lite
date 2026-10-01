---
id: collect-261001-ia-llm/ia-llm/hallucination-ia-2026-deepseek-94-vs-gpt-6-astra-51-3
title: "Comparaison basique du taux de réponses \"je ne sais pas\" entre deux fournisseurs"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Mistral", "OpenAI", "xAI"]
dates: []
keywords: ["astra", "benchmarks", "claude", "deepseek", "distribution", "fable 5", "gemini", "gpt-5.6", "gpt-6", "grok", "grok 4", "leaderboard"]
source: docs/RAG/collect-261001-ia-llm/hallucination-ia-2026-deepseek-94-vs-gpt-6-astra-51.md
source_anchor: ""
source_lines: [65, 92]
sha256: 7674e8d3296c70c279dd9c7b7fc7c638448f075ef3b01f6af9bfe336fe12ca5c
---

# Comparaison basique du taux de réponses "je ne sais pas" entre deux fournisseurs

Gemini 3 Deep Think, le mode de raisonnement approfondi de Google, présente un résultat plus surprenant : sa précision sur AA-Omniscience atteint 41 %, un chiffre inférieur à celui de Gemini 3.1 Pro en mode standard (55 %). Ce paradoxe apparent s’explique probablement par la nature du test : un mode de raisonnement plus long peut, sur certaines tâches de rappel factuel pur, introduire davantage d’occasions de dérive que sur des questions purement logiques ou mathématiques, où Deep Think excelle par ailleurs. Aucun taux d’hallucination isolé n’a été publié pour ce mode spécifique dans les sources consultées.

Côté tarification, Google applique une grille à deux paliers selon la page de tarification officielle de l’API Gemini : 2 $ par million de tokens en entrée et 12 $ en sortie pour les requêtes de moins de 200 000 tokens, puis 4 $ et 18 $ au-delà de ce seuil, les tokens de “réflexion” (thinking tokens) de Deep Think étant facturés au même tarif que les tokens de sortie classiques. La fenêtre de contexte annoncée est de 1 million de tokens pour les deux variantes, un chiffre qui varie toutefois selon les sources et les canaux de distribution (API directe, Vertex AI), ce qui justifie de vérifier la documentation au moment de l’intégration plutôt que de se fier à un chiffre relayé par un comparatif tiers, y compris celui-ci.

## DeepSeek V4 : l’open source le moins cher du marché, et le moins sûr sur les faits

DeepSeek a publié DeepSeek V4 Pro et DeepSeek V4 Flash en avril 2026, avant de sortir une variante DeepSeek-V4.1-Flash le 10 septembre 2026, présentée par l’entreprise dans son propre journal des mises à jour et de sa grille tarifaire officielle comme le plus petit modèle de sa nouvelle famille d’architecture, avec une compréhension visuelle multimodale native. Cette trajectoire de prix cassés prolonge celle déjà observée dans notre comparatif DeepSeek V4 vs Claude Fable 5 vs GPT-5.6. Sur le plan du prix, DeepSeek reste imbattable : DeepSeek V4 Pro coûte 0,435 $ par million de tokens en entrée et 0,87 $ en sortie, tandis que DeepSeek V4 Flash descend à 0,14 $ et 0,28 $, soit environ 30 à 70 fois moins cher que Claude Opus 5 ou GPT-6 Astra selon le sens de la comparaison.

Ce prix a une contrepartie factuelle nette. Sur AA-Omniscience, DeepSeek V4 Pro obtient un indice de -10, une amélioration de 11 points par rapport à son prédécesseur DeepSeek V3.2 Reasoning (-21), portée principalement par une meilleure précision de connaissance selon Artificial Analysis. Une étape intermédiaire, DeepSeek V3.1-Terminus, avait pourtant déjà réduit de 38 % le taux d’hallucination de son prédécesseur dès août 2026 selon ZoomBangla, une trajectoire que la génération V4 ne prolonge cependant pas sur l’indicateur d’hallucination brut. Cette progression s’accompagne en effet d’un taux d’hallucination de 94 % pour V4 Pro et 96 % pour V4 Flash, les deux scores les plus élevés de ce comparatif, loin devant DeepSeek V3.2 Reasoning qui plafonnait déjà à 82 %. Autrement dit, quand DeepSeek V4 ne connaît pas une réponse, il l’invente dans la quasi-totalité des cas plutôt que de reconnaître son incertitude.

Cette combinaison prix imbattable et taux d’hallucination élevé résume assez bien le compromis que proposent aujourd’hui les modèles ouverts chinois les moins chers : ils conviennent très bien à des tâches où une erreur factuelle a un coût faible et où un mécanisme de vérification externe (RAG, base de connaissances contrôlée, relecture humaine) est déjà en place, mais ils sont un choix risqué pour toute application qui expose directement une réponse non vérifiée à un client final ou à un décideur, en particulier dans un contexte réglementé comme celui de la santé, du droit ou de la finance en France.

## Mistral Large 3 : l’option européenne et son angle mort statistique

Mistral Large 3, sorti le 2 décembre 2025, reste le seul modèle produit par un éditeur européen dans ce comparatif, un point qui compte pour les organisations françaises soumises au RGPD ou soucieuses de souveraineté numérique, un enjeu déjà abordé dans notre comparatif Mistral vs Claude Sonnet 5. Sur le plan tarifaire, Mistral AI reste extrêmement compétitif selon sa page de tarification officielle : 0,50 $ par million de tokens en entrée, 0,05 $ pour les tokens mis en cache, et 1,50 $ en sortie, pour une fenêtre de contexte d’environ 262 144 tokens, plus modeste que le million de tokens proposé par la plupart des concurrents américains et chinois de ce comparatif.

Le point le plus notable concernant Mistral Large 3 dans le cadre de cet article n’est pas un chiffre, c’est l’absence de chiffre. Aucune des sources consultées, y compris le classement AA-Omniscience d’Artificial Analysis et le Vectara Hallucination Leaderboard, ne publie de score d’hallucination indépendant pour Mistral Large 3 à la date du 12 septembre 2026. Le modèle apparaît dans des comparatifs de prix et de capacités générales, mais pas dans les benchmarks de fiabilité factuelle qui font référence pour les autres modèles de ce panel. Ce vide n’est pas nécessairement le signe d’un problème caché : il peut simplement refléter le fait que Mistral Large 3 n’a pas encore été intégré au protocole de test de ces laboratoires indépendants, ou que l’éditeur n’a pas communiqué de partenariat d’évaluation avec eux.

Pour une équipe française ou européenne qui doit documenter ses choix technologiques dans le cadre de l’AI Act, cette absence de données indépendantes constitue en soi un critère de décision : elle impose de mener ses propres tests de fiabilité factuelle avant tout déploiement en production, plutôt que de s’appuyer sur un classement tiers qui, pour l’instant, ne couvre pas ce modèle. C’est un exercice que ce même comparatif recommande d’ailleurs pour tous les modèles listés ici, dans la section méthodologie plus haut.

## Grok 4.6, Grok 4.5 et Qwen3.8-Max : les autres prétendants à surveiller

Grok 4.6, lancé par xAI le 12 août 2026, obtient un indice AA-Omniscience de 30,5, mesuré le 12 août 2026 en configuration à effort de raisonnement élevé selon Artificial Analysis, sans qu’un taux d’hallucination isolé ne soit publiquement détaillé pour ce modèle précis dans les sources disponibles. C’est le même Grok 4.6 déjà positionné face à Mistral Large 3 dans notre comparatif Grok 4.6 vs Qwen3.8-Max vs Mistral Large 3. Son prix suit une grille à seuil selon la documentation officielle de tarification xAI : 2 $ par million de tokens en entrée et 6 $ en sortie jusqu’à 200 000 tokens de contexte, puis 4 $ et 12 $ au-delà, pour une fenêtre de contexte totale de 500 000 tokens, nettement inférieure au million de tokens proposé par Claude, Gemini ou DeepSeek.

Qwen3.8-Max, publié par Alibaba le 3 août 2026, n’a pas encore reçu de score AA-Omniscience au moment de la rédaction : plusieurs mirroirs du classement le signalent explicitement comme “non noté” sur cet indicateur précis, alors qu’il obtient par ailleurs 56 points sur l’indice de raisonnement général d’Artificial Analysis, un score qui le place au niveau de Claude Opus 4.8 selon le média spécialisé The Decoder. Son prix, 2 $ en entrée et 6 $ en sortie par million de tokens pour une fenêtre de contexte proche du million de tokens, en fait une alternative crédible sur le plan du rapport capacités-prix, mais l’absence de données sur sa tendance à halluciner empêche toute conclusion ferme sur sa fiabilité factuelle par rapport aux autres modèles de ce comparatif.

## Tableau des prix : combien coûte chaque modèle par million de tokens

