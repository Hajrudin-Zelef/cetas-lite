---
id: collect-261001-ia-llm/ia-llm/deepseek-v4-vs-qwen3-8-max-vs-glm-5-3-x14-de-prix-1
title: "DeepSeek V4 Flash 0731 (via API officielle)"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Hugging Face", "Mistral", "OpenAI", "OpenRouter", "Together AI", "Z.ai"]
dates: []
keywords: ["deepseek", "agents", "benchmarks", "claude", "cost", "cyber", "glm", "gpt-5.6", "mistral", "moe", "opus 5", "qwen"]
source: docs/RAG/collect-261001-ia-llm/deepseek-v4-vs-qwen3-8-max-vs-glm-5-3-x14-de-prix.md
source_anchor: ""
source_lines: [1, 36]
sha256: 5b4e42796223d1a99ea90dadc29d282c4790ac28f30c03d83ea3bb3c97a3e560
---

# DeepSeek V4 Flash 0731 (via API officielle)

Trois lancements en deux semaines, trois architectures MoE, un seul objectif affiché : reprendre l’initiative face à OpenAI et Anthropic sur le terrain du prix et du code. Entre le 31 juillet et le 15 août 2026, DeepSeek, Alibaba et Zhipu AI (Z.ai) ont chacun publié leur nouveau modèle phare : DeepSeek V4 Flash 0731, Qwen3.8 Max et GLM-5.3. Les trois visent le même public, celui des développeurs et des entreprises qui veulent des modèles agentiques puissants sans payer le tarif Claude Opus 5 ou GPT-5.6. Mais derrière cette promesse commune se cachent des écarts considérables : de 284 milliards à 2,4 billions de paramètres, et un rapport de prix qui va jusqu’à 14 fois entre le moins cher et le plus cher sur les tokens de sortie. Ce comparatif détaille les architectures, les tarifs, les benchmarks disponibles et les cas d’usage réels pour aider les équipes françaises et européennes à choisir le bon modèle avant la rentrée.

## Trois lancements en deux semaines : ce qui s’est passé en août 2026

La séquence est serrée. DeepSeek a publié la version **0731** de son modèle Flash le 31 juillet, avec un focus explicite sur les gains agentiques plutôt que sur une nouvelle architecture. Deux jours plus tard, le 2 août, Alibaba a dévoilé Qwen3.8 Max, un modèle massif de 2,4 billions de paramètres présenté comme capable de mener des tâches complexes en autonomie sur plusieurs jours. Puis, le 14 août, Zhipu AI a sorti GLM-5.3 avec un positionnement marketing sans détour : « Built to Code. Ready for Cyber Defense. »

Les trois laboratoires partagent un point commun frappant : aucun n’a changé de base pré-entraînée. DeepSeek V4 Flash 0731 garde l’architecture MoE à 284 milliards de paramètres de la préversion d’avril, réentraînée uniquement sur des données orientées agents. GLM-5.3 reprend le même socle à 743 milliards de paramètres que GLM-5.2, amélioré « entièrement par du post-entraînement à grande échelle », selon les documents techniques de Z.ai. Seul Qwen3.8 Max change d’échelle avec un saut massif à 2,4 billions de paramètres, ce qui en fait de loin le plus gros des trois modèles de cette génération.

Ce virage vers le post-entraînement plutôt que le pré-entraînement coûteux illustre une tendance de fond en Chine en 2026 : les laboratoires optimisent l’existant pour gagner sur des benchmarks agentiques précis (code, terminal, sécurité) au lieu de relancer des cycles d’entraînement à plusieurs centaines de millions de dollars. Pour l’écosystème européen, qui suit de près les alternatives à Mistral Large 3, cette bataille de prix change directement le calcul de rentabilité d’un déploiement d’agents IA à grande échelle.

## DeepSeek V4 Flash 0731 : la référence low-cost à poids ouverts

DeepSeek V4 Flash 0731 reste fidèle à la philosophie qui a fait le succès de DeepSeek depuis 2025 : un modèle Mixture-of-Experts à **284 milliards de paramètres totaux** dont seulement **13 milliards sont actifs par token**. Cette sparsité extrême explique pourquoi le modèle tient un prix aussi bas tout en offrant une fenêtre de contexte d’un million de tokens et une sortie maximale de 384 000 tokens.

Le vrai changement de la version 0731 ne touche pas l’architecture mais l’entraînement. DeepSeek a reposté le modèle uniquement sur des données orientées agents, ce qui se traduit par un score de **82,7 % sur Terminal-Bench 2.1**, contre 61,8 % pour la préversion d’avril selon les données publiées par MarkTechPost. C’est un bond de plus de 20 points sans changer un seul paramètre du réseau.

Côté licence, DeepSeek V4 Flash 0731 est publié en **poids ouverts sous licence MIT**, disponible le jour même de l’annonce selon la fiche technique publiée sur Hugging Face. C’est un argument de poids pour les équipes qui veulent héberger le modèle en interne, un critère de plus en plus scruté en France depuis l’entrée en vigueur des obligations de transparence de l’article 50 de l’AI Act européen, effective depuis le 2 août 2026. Le modèle est aussi servi via l’API officielle DeepSeek, via OpenRouter et via Together AI, avec une limite de concurrence fixée à 2 500 requêtes simultanées sur l’API directe.

Le point fort commercial reste le cache : 0,14 $ par million de tokens d’entrée en cas de défaut de cache, mais seulement **0,0028 $** si le contenu est déjà en cache, soit une remise de 98 %. Pour des applications à contexte répétitif (agents qui relisent le même code, chatbots avec un prompt système long), cette architecture de tarification rend DeepSeek V4 Flash quasiment imbattable sur le coût réel par requête.

## Qwen3.8 Max : le mastodonte à 2,4 billions de paramètres

Alibaba change de catégorie avec Qwen3.8 Max. Le modèle affiche **2,4 billions de paramètres totaux**, avec 95 milliards de paramètres actifs par requête selon les données publiées au lancement le 2 août 2026. C’est près de neuf fois la taille totale de DeepSeek V4 Flash, ce qui en fait l’un des plus gros modèles disponibles publiquement en 2026, comme le souligne notre actualité dédiée au lancement de Qwen3.8-Max.

Sur le plan des capacités, Qwen3.8 Max se distingue par deux choses. D’abord, la **modalité native** : le modèle accepte du texte, de l’image et de la vidéo en entrée, ce qu’aucun des deux autres modèles de ce comparatif ne propose. Ensuite, un positionnement explicite sur les tâches longues : Alibaba met en avant sa capacité à « gérer des tâches complexes de façon autonome sur plusieurs jours », y compris la reproduction de résultats de recherche ou la conception de circuits, selon la couverture de lancement.

Sur les benchmarks, Qwen3.8 Max affiche des scores solides là où les données existent : **92,6** sur GPQA Diamond, **67,7** sur SWE-Bench Pro, **86,6** sur Terminal-Bench 2.1 et **86,1** sur OSWorld-Verified, chiffres vendeur-déclarés relayés par plusieurs agrégateurs de spécifications. Ce sont, à ce jour, les seuls scores GPQA Diamond et SWE-Bench disponibles pour les trois modèles de ce comparatif, ce qui fait de Qwen3.8 Max le mieux documenté sur le raisonnement scientifique.

Côté tarif, deux chiffres circulent et il faut les distinguer. Le tarif officiel Alibaba Cloud Model Studio annonce 12 yuans pour 1 million de tokens d’entrée et 36 yuans en sortie, soit environ **1,67 $ et 5,00 $** au taux de change courant, identique au tarif de Qwen3.7-Max. Mais la majorité des agrégateurs API et guides développeurs cités début août listent plutôt **2 $ en entrée et 6 $ en sortie** par million de tokens. C’est ce second tarif, plus élevé, qui sert de référence dans ce comparatif car c’est celui rencontré par la majorité des intégrateurs tiers.

Sur la licence, Qwen3.8 Max est disponible en API dès le lancement via Alibaba Cloud Model Studio et QwenCloud, mais les poids ouverts n’étaient pas encore publiés au moment de l’annonce. Alibaba a communiqué une mise à disposition « la semaine suivante » sur Hugging Face et ModelScope, ce qui en ferait le premier modèle de la gamme Qwen-Max à passer en poids ouverts. Le nom exact de la licence n’a pas été précisé dans les sources disponibles au moment de la rédaction.

## GLM-5.3 : le spécialiste code et cyberdéfense de Zhipu AI

Zhipu AI, qui opère sous la marque Z.ai, a choisi un axe différent avec GLM-5.3 : ne pas retoucher la base et concentrer tous les efforts de post-entraînement sur le code et les tâches longues. Le modèle garde une architecture MoE à environ **743 milliards de paramètres totaux** (une source la donne à 753 milliards) avec **40 milliards de paramètres actifs par token**, un chiffre identique à celui de GLM-5.2 sorti en juin.

