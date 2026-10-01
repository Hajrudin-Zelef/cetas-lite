---
id: collect-261001-ia-llm/ia-llm/gpt-6-astra-vs-sol-vs-luna-prix-x100-2026-1
title: "Avant : appel avec GPT-5.6 Sol"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Google", "Microsoft", "Mistral", "OpenAI"]
dates: []
keywords: ["gpt-5.6", "sol", "agents", "agi", "astra", "attention", "bedrock", "benchmarks", "chatgpt", "claude", "fable 5", "gemini"]
source: docs/RAG/collect-261001-ia-llm/gpt-6-astra-vs-sol-vs-luna-prix-x100-2026.md
source_anchor: ""
source_lines: [1, 54]
sha256: f49b3db2b147b598dd775fae72e0a649f741194970cd62bc450e7e10f46fa8f8
---

# Avant : appel avec GPT-5.6 Sol

Le 22 septembre 2026, OpenAI a discrètement réorganisé toute sa gamme GPT-6 en publiant deux nouveaux modèles, GPT-6 Sol et GPT-6 Luna, à peine trois semaines après le lancement de GPT-6 Astra le 3 septembre. Résultat : la famille GPT-6 compte désormais trois membres aux prix qui varient d’un facteur 100, et aucun guide ne permettait jusqu’ici de choisir entre eux. Pour les développeurs et les entreprises françaises et européennes qui budgétisent leurs projets d’intelligence artificielle, la question n’est plus seulement “faut-il utiliser GPT-6”, mais “lequel des trois GPT-6”. Ce comparatif détaille les spécifications, les prix, les benchmarks et les cas d’usage réels de GPT-6 Astra, GPT-6 Sol et GPT-6 Luna, avec des chiffres vérifiés et des recommandations concrètes selon votre budget et votre charge de travail.

## GPT-6 Astra, Sol et Luna : trois modèles, une seule famille

GPT-6 Astra est arrivé en premier, le 3 septembre 2026, comme modèle phare d’OpenAI destiné au raisonnement complexe, à la génération de code et à l’usage d’ordinateur (computer use). Son déploiement a commencé auprès d’un nombre limité d’organisations avant de s’étendre progressivement aux comptes ChatGPT Plus, Pro, Business et Enterprise, ainsi qu’à l’API OpenAI, à Microsoft Azure et à Amazon Bedrock. Selon les notes de version d’OpenAI, Astra n’était encore “pas disponible à grande échelle” à la mi-septembre, ce qui explique pourquoi une bonne partie des utilisateurs Plus ou Pro n’y ont pas encore accès trois semaines après l’annonce.

Dix-neuf jours plus tard, OpenAI a comblé le vide en dessous d’Astra avec deux modèles frères. GPT-6 Sol cible le travail quotidien : rédaction, analyse, code courant, avec une capacité de raisonnement proche d’Astra mais à un tarif bien plus faible. GPT-6 Luna, de son côté, privilégie la vitesse et le volume : résumés, extraction d’informations, classification de texte, réponses rapides à faible coût. OpenAI présente les deux nouveaux venus comme entraînés “avec des méthodes similaires à celles de GPT-6 Astra”, en reprenant une partie de ses avancées en matière de fiabilité factuelle et de programmation, mais avec une architecture optimisée pour l’inférence à grande échelle.

Cette structure à trois étages n’est pas une nouveauté chez OpenAI. Elle reproduit la logique déjà testée avec GPT-5.6, qui proposait lui aussi les variantes Sol et Luna. Ce qui change avec GPT-6, c’est l’écart de prix : OpenAI a réduit de moitié les tarifs API de Sol et Luna par rapport aux prix promotionnels de GPT-5.6, tout en revendiquant des gains de précision mesurables sur plusieurs benchmarks d’agents.

## Tableau comparatif : les spécifications de GPT-6 Astra, Sol et Luna

Voici la fiche technique complète des trois modèles GPT-6, compilée à partir des annonces officielles d’OpenAI et des premières analyses indépendantes publiées entre le 3 et le 23 septembre 2026.

| Caractéristique | GPT-6 Astra | GPT-6 Sol | GPT-6 Luna | 
|---|---|---|---|
| Date de lancement | 3 septembre 2026 (déploiement progressif) | 22 septembre 2026 | 22 septembre 2026 | 
| Identifiant API | gpt-6-astra | gpt-6-sol | gpt-6-luna | 
| Prix en entrée (par 1M tokens) | 10,00 $ | 2,00 $ | 0,10 $ | 
| Prix en sortie (par 1M tokens) | 50,00 $ | 10,00 $ | 0,50 $ | 
| Prix en cache (par 1M tokens) | Non communiqué | 0,20 $ | 0,01 $ | 
| Baisse de prix vs génération précédente | Nouveau modèle | -50% vs GPT-5.6 Sol | -50% à -58% vs GPT-5.6 Luna | 
| Seuil de surtaxe longue contexte | Non communiqué | Au-delà de 272 000 tokens en entrée, tarif doublé | Non communiqué | 
| Accès ChatGPT Free / Go | Non | Non | Oui, via l’application de bureau uniquement | 
| Accès ChatGPT Plus / Pro / Business / Enterprise / Edu | Oui (déploiement progressif) | Oui, dans ChatGPT Work et Codex | Oui, dans ChatGPT Work et Codex | 
| Accès dans le chat ChatGPT classique | Oui | Pas encore, selon l’annonce du 22 septembre | Pas encore, selon l’annonce du 22 septembre | 
| Accès cloud tiers | Azure, Amazon Bedrock | API OpenAI | API OpenAI | 
| Score Artificial Analysis Intelligence Index v4.3 (7 sept. 2026) | 53 points | Pas encore évalué séparément | Pas encore évalué séparément | 
| Score ARC-AGI-2 | 95% | Non communiqué | Non communiqué | 
| Score GPQA | 96% | Non communiqué | Non communiqué | 
| Positionnement officiel OpenAI | Raisonnement de pointe, code, usage d’ordinateur | Travail quotidien, bon compromis coût/raisonnement | Réponses rapides, gros volumes, résumé et extraction | 

Deux lignes méritent une attention particulière. D’abord, le seuil de 272 000 tokens qui déclenche un doublement du tarif d’entrée sur GPT-6 Sol : au-delà de ce volume, chaque requête coûte deux fois plus cher, un détail que beaucoup d’équipes techniques risquent de découvrir uniquement sur leur première facture. Ensuite, la disponibilité dans le “chat” classique de ChatGPT : au 22 septembre 2026, Sol et Luna existent uniquement dans ChatGPT Work et dans Codex, l’environnement de développement d’OpenAI, pas encore dans l’interface de conversation grand public que la plupart des utilisateurs Plus ouvrent chaque jour.

## Prix et tarification API : un écart qui va jusqu’à x100

C’est la donnée la plus frappante de ce lancement. Entre GPT-6 Astra et GPT-6 Luna, le tarif d’entrée passe de 10 $ à 0,10 $ par million de tokens, et le tarif de sortie de 50 $ à 0,50 $. Dans les deux cas, l’écart est exactement de 100 fois. Un développeur qui traite le même volume de texte avec Luna plutôt qu’avec Astra divise sa facture API par 100, au prix d’une capacité de raisonnement nettement inférieure sur les tâches complexes.

| Modèle | Prix entrée / 1M tokens | Prix sortie / 1M tokens | Évolution | 
|---|---|---|---|
| GPT-6 Astra | 10,00 $ | 50,00 $ | Nouveau tarif de lancement | 
| GPT-6 Sol | 2,00 $ | 10,00 $ | -50% vs GPT-5.6 Sol (4,00 $ / 20,00 $) | 
| GPT-6 Luna | 0,10 $ | 0,50 $ | -50% vs GPT-5.6 Luna (0,20 $ / 1,20 $) | 
| GPT-5.6 Sol (ancienne génération) | 4,00 $ | 20,00 $ | Remplacé par GPT-6 Sol | 
| GPT-5.6 Luna (ancienne génération) | 0,20 $ | 1,20 $ | Remplacé par GPT-6 Luna | 
| Claude Fable 5.1 (indicatif, Anthropic) | ~10,00 $ | ~10,00 $ | Tarif de référence pour un modèle de raisonnement haut de gamme | 
| Gemini 3.8 Flash (indicatif, Google) | ~0,75 $ | ~3,75 $ | Positionné entre Sol et Luna | 

OpenAI justifie cette baisse de moitié sur Sol et Luna par des gains d’efficacité côté infrastructure. Dans son annonce du 22 septembre, l’entreprise explique avoir amélioré la mise en cache et l’inférence, ce qui lui permet de répercuter ces économies sur les tarifs API sans changer sa marge. Concrètement, le prix en cache de GPT-6 Sol tombe à 0,20 $ par million de tokens, et celui de Luna à seulement 0,01 $, soit une réduction de 90% par rapport au tarif standard d’entrée pour les requêtes qui réutilisent un contexte déjà traité. Pour une entreprise qui interroge répétitivement les mêmes documents de référence (base de connaissances interne, catalogue produit, documentation technique), ce mécanisme de cache change concrètement l’équation économique, bien plus que le prix nominal affiché en tête de grille tarifaire.

À titre de comparaison, le comparatif publié sur ce site entre les abonnements ChatGPT, Claude, Gemini et Mistral montre qu’un abonnement individuel Plus reste fixé indépendamment du choix de modèle sous-jacent. La vraie sensibilité au prix se joue côté API, là où les volumes se comptent en dizaines de millions de tokens par mois.

## Qui a accès à quoi ? ChatGPT Free, Plus, Pro, Business, Enterprise

