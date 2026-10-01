---
id: collect-261001-ia-llm/ia-llm/claude-opus-5-en-tete-des-llm-63-1-vs-gpt-5-6-2026-1
title: "claude-opus-5-en-tete-des-llm-63-1-vs-gpt-5-6-2026"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Meta", "Microsoft", "Moonshot", "OpenAI", "Z.ai"]
dates: []
keywords: ["claude", "apache", "benchmark", "benchmarks", "deepseek", "fable 5", "gemini", "glm", "gpt-5.6", "gpu", "kimi", "llama"]
source: docs/RAG/collect-261001-ia-llm/claude-opus-5-en-tete-des-llm-63-1-vs-gpt-5-6-2026.md
source_anchor: ""
source_lines: [1, 40]
sha256: 933d0a5c662188c98e50b081c4bf473d54cb18972a01926562e56aba76545aeb
---

# claude-opus-5-en-tete-des-llm-63-1-vs-gpt-5-6-2026

Le 1er août 2026, le classement Intelligence Index v4.1 d’Artificial Analysis a basculé. **Claude Opus 5** d’Anthropic est passé en tête devant tous les autres grands modèles de langage, avec un score de 63,1 points relevé le 15 août. Une semaine plus tôt, OpenAI avait sorti **GPT-5.6** en trois variantes. Deux semaines avant ça, le modèle chinois **Kimi K3**, doté de 2 800 milliards de paramètres, était devenu le plus gros modèle à poids ouverts jamais publié. Et le 2 août, les obligations de transparence de l’AI Act européen sont entrées en vigueur avec force contraignante, en plein cœur de cette bataille de classements.

Pour les équipes techniques françaises et européennes qui doivent choisir un modèle d’ici la rentrée, la situation est à la fois plus riche et plus confuse qu’il y a un an. Six familles de modèles majeurs ont changé de génération en moins de huit semaines. Voici ce que ces mouvements signifient concrètement, avec les chiffres qui comptent.

## Claude Opus 5 prend la tête du classement mondial des LLM

Anthropic a mis en ligne Claude Opus 5 le 24 juillet 2026. L’entreprise le présente comme une intelligence proche de son modèle frontière Fable 5, mais deux fois moins chère à l’usage, avec un curseur d’effort réglable qui permet d’ajuster la profondeur de raisonnement selon la tâche. Un mois plus tard, ce pari technique s’est traduit en chiffres concrets : au 15 août 2026, Claude Opus 5 (max) occupe la première place du classement Intelligence Index d’Artificial Analysis, avec un score de 63,1, une vitesse de génération de 46 tokens par seconde et un tarif d’environ 4,32 € en entrée et 21,61 € en sortie par million de tokens, sur une fenêtre de contexte d’un million de tokens.

Ce résultat place Opus 5 devant Claude Fable 5, le modèle « Mythos » d’Anthropic sorti en juin, qui conserve un score autour de 60. Il devance aussi GPT-5.6 Sol dans sa configuration maximale, classé troisième début août. C’est la première fois depuis le lancement de Fable 5 qu’un autre modèle Anthropic repasse devant lui sur un classement de référence indépendant, et non plus seulement sur LMArena, le classement communautaire déjà remporté par Fable 5 à la mi-juin.

## Pourquoi Fable 5 reste absent des bureaux européens

La montée d’Opus 5 prend tout son sens à la lumière d’un épisode que nous avions couvert en juin. Trois jours après le lancement de Fable 5, le 12 juin 2026, le département du Commerce américain a coupé l’accès au modèle pour tout utilisateur situé hors des États-Unis, invoquant des restrictions à l’exportation. Deux mois plus tard, cette limitation n’a pas été levée. Pour une entreprise française qui veut un modèle Anthropic de pointe accessible sans dépendre d’une décision politique américaine, Opus 5 est donc devenu, de fait, le choix par défaut plutôt qu’une simple alternative économique.

Cette situation illustre un phénomène que les équipes IT commencent à intégrer dans leurs plans de continuité : la disponibilité d’un modèle propriétaire peut changer du jour au lendemain, indépendamment de ses performances techniques. C’est un argument de poids en faveur d’une architecture multi-fournisseurs, plutôt que d’un verrouillage sur un seul modèle frontière.

## GPT-5.6 : trois variantes et une baisse de prix de 80 %

OpenAI a lancé GPT-5.6 le 9 juillet 2026, après une préversion réservée à des partenaires de confiance le 26 juin. Le modèle se décline en trois profils : Sol, taillé pour le raisonnement complexe et le code, Terra, positionné en usage généraliste, et Luna, la version économique destinée aux tâches simples et au volume. Le 30 juillet, OpenAI a réduit le prix de Luna de 80 %, un geste tarifaire agressif qui vise directement les usages d’automatisation à grande échelle : support client, classification de tickets, extraction de données.

Sur les classements de qualité pure, GPT-5.6 Sol en configuration maximale reste sur le podium, juste derrière Opus 5 et Fable 5. Mais c’est sur le rapport prix-performance que la famille GPT-5.6 marque des points : Luna devient, après sa baisse de prix, l’un des modèles propriétaires les moins chers du marché pour les tâches qui ne demandent pas un raisonnement poussé. Cette segmentation en trois niveaux, plutôt qu’un modèle unique, devient la norme chez les grands fournisseurs : Anthropic distingue déjà Opus et Fable, Google fait de même avec sa gamme Gemini.

## Kimi K3 : le plus gros modèle à poids ouverts jamais publié

Côté modèles ouverts, l’événement du mois vient de Chine. Kimi K3, développé par Moonshot AI, a vu ses poids mis en ligne le 26 juillet 2026. Avec 2 800 milliards de paramètres, il s’agit du plus gros modèle librement téléchargeable jamais publié, dépassant largement GLM-5.2, DeepSeek V4 Pro et Qwen3.7 Max. Sur les benchmarks de code, la démonstration de force est nette : Kimi K3 établit un record de 96 % sur SWE-bench Verified et atteint 79,2 % sur SWE-bench Pro, contre 69,2 % pour Claude Opus 4.8, la génération précédente d’Anthropic.

Pour les entreprises européennes soumises à l’AI Act, un modèle à poids ouverts comme Kimi K3 présente un intérêt réglementaire direct : il peut être audité, hébergé localement et documenté plus facilement qu’un modèle propriétaire fermé. C’est un argument que l’on retrouve de plus en plus dans les appels d’offres publics en France, où la traçabilité des données d’entraînement devient un critère de sélection à part entière.

## DeepSeek V4 et la pression continue des modèles ouverts chinois

Kimi K3 n’est pas seul sur ce terrain. DeepSeek V4 est disponible depuis fin juillet en aperçu à poids ouverts, ce qui en fait une alternative crédible aux modèles fermés sur les tâches de code et de raisonnement technique. Depuis l’irruption de DeepSeek V3 fin 2024, le laboratoire chinois a maintenu un rythme de publication qui force les acteurs occidentaux à justifier leurs prix. La bataille ne se joue plus seulement sur le score brut, mais sur le coût par tâche réellement accomplie, un indicateur que les équipes de DevOps et de finance suivent désormais de près dans leurs feuilles de calcul de FinOps IA.

## Gemini 3.6 Flash et l’offensive open source de Meta

Google positionne Gemini 3.6 Flash sur un créneau différent : un score d’intelligence de 50, une fenêtre de contexte d’un million de tokens et un tarif de 1,50 $ en entrée et 7,50 $ en sortie par million de tokens. Ce n’est pas le modèle le plus intelligent du marché, mais c’est l’un des plus rapides et des moins chers pour de la génération de contenu à grand volume, ce qui explique sa popularité croissante chez les éditeurs de sites et les agences qui doivent produire en masse.

Meta, de son côté, a choisi la carte de l’ouverture totale. Le 10 août 2026, l’entreprise a dévoilé Muse Glimmer, un modèle multimodal open source de 30 milliards de paramètres, sous licence Apache 2.0, capable de tourner sur un seul GPU grand public une fois compressé. Meta le présente comme un modèle « agentique », conçu pour exécuter des workflows automatisés plutôt que pour répondre à des questions ponctuelles. En parallèle, la génération Llama 4 continue d’être citée comme le champion open source généraliste, avec une variante Scout qui offre une fenêtre de contexte record de 10 millions de tokens et une version 405B qui a obtenu un score de 81,1 sur un benchmark francophone en juillet.

## EuroLLM-22B, la réponse européenne à la dépendance américaine

