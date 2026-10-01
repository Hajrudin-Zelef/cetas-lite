---
id: collect-261001-ia-llm/ia-llm/gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces-2
title: "gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces"
domain: ia-llm
role: reference
task: reference
actors: ["Google", "OpenAI", "SpaceX", "xAI"]
dates: []
keywords: ["benchmark", "benchmarks", "gemini", "agent", "agents", "astra", "gemini 3.8", "gpt-live", "grok", "leaderboard", "voice"]
source: docs/RAG/collect-261001-ia-llm/gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces.md
source_anchor: ""
source_lines: [55, 130]
sha256: ef8202a7d87e742102d25153556f9e17203901ffcb5d4144411097e46dbf9ae0
---

# gemini-3-8-live-fonctionnalites-benchmarks-tarifs-acces

Deux petits changements affectent le code existant. L’audio proactif, où le modèle peut décider de ne pas répondre à une parole qui ne lui est pas destinée, est désormais activé en permanence, et le régler à `false` renvoie une erreur. Le dialogue affectif a été entièrement retiré de l’API ; toute configuration `enable_affective_dialog` doit être supprimée.

## Comment Gemini 3.8 Live se comporte-t-il sur les benchmarks ?

Extended Thinking mène tous les tableaux de qualité et d’agentique publiés par Google, mais de peu devant GPT-Live-1 Astra d’OpenAI, tandis que le modèle 3.8 Live de base l’emporte largement sur le coût mais perd sur les tâches agentiques.

L’index, τ-Voice, Big Bench Audio et les chiffres de coût ci-dessous proviennent tous du leaderboard public d’Artificial Analysis, donc mesurés indépendamment et non déclarés par les vendeurs. Les chiffres τ³-Banking viennent du graphique de lancement de Google citant Sierra : considérez cette ligne comme déclarée par le fournisseur jusqu’à ce que le tableau de Sierra reflète les mêmes résultats.

### Accomplissement de tâches agentiques

Gemini 3.8 Live Extended Thinking est en tête sur τ-Voice, le benchmark de Sierra qui mesure si un agent vocal mène à bien des tâches multi-étapes avec des outils, tel que mesuré par Artificial Analysis. GPT-Live-1 Astra suit de près, le reste du peloton décroche :

- **Gemini 3.8 Live Extended Thinking :** 68,6 %
- **GPT-Live-1 Astra :** 67,9 %
- **Grok Voice Think Fast 2.0** : 56,5 %
- **Gemini 3.1 Flash Live :** 37,7 %
- **Gemini 3.8 Live (version de base) :** 30,1 %

Le chiffre surprenant est que le score du modèle de base sur τ-Voice est inférieur à celui de son propre prédécesseur. Google est clair : 3.8 Live est conçu pour l’échelle et l’efficacité des coûts plutôt que pour des workflows complexes, mais un écart de plus de 38 points entre les deux variantes signifie que le choix de palier n’est pas une nuance. Si votre agent enchaîne des outils, le modèle de base n’est pas le bon défaut.

Sur le leaderboard τ³-Banking de Sierra, un benchmark service client plus ardu centré sur des workflows bancaires, Extended Thinking obtient 35,1 % contre 32,0 % pour GPT-Live-1 Astra, 16,5 % pour Grok Voice Think Fast 2.0 de SpaceXAI, 11,3 % pour Gemini 3.1 Flash Live et 10,3 % pour GPT-Realtime 2. Tous les modèles de ce tableau échouent la plupart du temps, signe de la distance qui reste avant un agent vocal autonome en banque, mais tripler le score de la génération Gemini précédente est un vrai progrès.

### Qualité vocale et raisonnement

Sur le Speech to Speech Index, Extended Thinking devance lui aussi la concurrence :

- **Gemini 3.8 Live Extended Thinking :** 82,6
- **GPT-Live-1 Astra :** 81,5
- **Grok Voice Think Fast 2.0 :** 81,3
- **Gemini 3.8 Live (base) :** 76,0
- **Gemini 3.1 Flash Live :** 71,5

Un avantage de 1,1 point sur OpenAI reste un avantage, mais je ne fonderais pas une décision d’achat uniquement dessus.

Pour le raisonnement pur sur entrée parlée, Google annonce 97,7 % sur Big Bench Audio pour Extended Thinking. Google indique aussi que les deux modèles repoussent la frontière de Pareto sur EVA-Bench de ServiceNow pour les agents vocaux, en équilibrant précision et qualité conversationnelle, même si ce test a été réalisé via l’API Live à travers la Gemini Enterprise Agent Platform plutôt que l’API publique.

### Coût par heure d’audio

C’est le graphique que Google veut surtout vous montrer. Dans le test de coût d’Artificial Analysis sur un sous-ensemble de Big Bench Audio, Gemini 3.8 Live traite une heure d’audio en entrée pour 0,84 $.

Extended Thinking coûte 3,50 $ pour la même heure, Grok Voice Think Fast 2.0 coûte 4,80 $, et GPT-Live-1 Astra 5,83 $. Gemini 3.1 Flash Live se situait à 1,50 $ et 1,75 $ selon le niveau de « thinking ».

En lisant les deux barres Gemini ensemble, la stratégie produit saute aux yeux. Le modèle de base écrase tous les autres sur le prix, et le modèle de raisonnement qui rivalise avec OpenAI en qualité coûte encore 40 % de moins par heure d’entrée. Notez que le modèle de raisonnement consomme plus de tokens aux niveaux de thinking plus élevés, d’où un coût environ 4 fois supérieur à son jumeau malgré une grille tarifaire commune.

## Quelle variante choisir ?

Gemini 3.8 Live est le bon défaut pour la plupart des agents vocaux, et Extended Thinking ne justifie sa consommation de tokens plus élevée que lorsque l’agent doit planifier, comparer ou attendre des outils lents. Les recommandations de Google tracent la frontière sur la latence des outils : si vos fonctions répondent en millisecondes, restez sur le modèle de base.

Les deux variantes partagent la même grille tarifaire (voir ci-dessous), des limites de tokens de 131 072 en entrée et 65 536 en sortie, et le même endpoint WebSocket. Ce qui les distingue, c’est l’architecture de raisonnement.

Gemini 3.8 Live utilise un raisonnement entrelacé avec une latence fixe et aucun réglage `thinking_level`. Extended Thinking expose des niveaux de raisonnement de fond `low`, `medium` et `high`, diffuse des « remplissages » conversationnels pendant qu’il travaille et exige des outils asynchrones. Ces niveaux de thinking sont le seul levier d’effort de cette version.

| Cas d’usage | Choix | Pourquoi | 
|---|---|---|
| Tri en service client, recherche vocale, pratique des langues | Gemini 3.8 Live | L’échange immédiat compte plus que la profondeur, et chaque tour utilisateur appelle une unique réponse | 
| Contrôle d’objets connectés, lecture de capteurs | Gemini 3.8 Live | Les outils répondent en millisecondes, rien à masquer | 
| Support technique sur logs, codes d’erreur et vérifs de configuration | Extended Thinking | Un diagnostic multi-étapes nécessite du raisonnement de fond entre les énoncés | 
| Voyage et réservation interrogeant vols et hôtels en parallèle | Extended Thinking | Appels asynchrones parallèles avec retours d’avancement parlés plutôt que du silence | 
| Accompagnement STEM et code | Extended Thinking | Le modèle vérifie des formules ou débugge du code avant d’énoncer une explication | 

Dernier point : la complexité côté client. Passer à Extended Thinking implique de réécrire votre gestion d’état autour de `interaction_status` ; si vous avez une appli 3.1 Flash Live fonctionnelle, le modèle de base est la mise à niveau à l’identique, et le modèle de raisonnement requiert un petit refactoring.

## Tarifs et disponibilité de Gemini 3.8 Live

Les deux modèles partagent la grille de prix de Gemini 3.1 Flash Live Preview, donc la mise à niveau est gratuite. Sur l’offre payante, l’audio en entrée coûte 3,00 $ par million de tokens (Google l’estime à 0,005 $ la minute) et l’audio en sortie 12,00 $ par million de tokens, tokens de thinking inclus (environ 0,018 $ la minute). Les tokens de thinking sont facturés au tarif de sortie, ce qui explique le coût d’exécution plus élevé d’Extended Thinking.

| Modalité | Offre payante, par 1 M de tokens | Estimation par minute | 
|---|---|---|
| Texte en entrée | 0,75 $ | n/a | 
| Audio en entrée | 3,00 $ | 0,005 $/min | 
| Image et vidéo en entrée | 1,00 $ | 0,002 $/min | 
| Texte en sortie | 4,50 $ | n/a | 
| Audio en sortie (tokens de thinking inclus) | 12,00 $ | 0,018 $/min | 

L’offre gratuite couvre les deux modèles sans frais pour l’entrée et la sortie, avec la réserve habituelle : Google utilise les données du palier gratuit pour améliorer ses produits ; celles du palier payant ne sont pas utilisées à cet effet.

