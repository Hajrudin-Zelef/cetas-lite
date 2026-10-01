---
id: collect-261001-ia-llm/ia-llm/deepseek-v4-1-flash-vs-glm-5-3-flash-vs-gemini-3-8-flash-2
title: "Estimation du coût mensuel par modèle (en dollars)"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "OpenAI", "Z.ai"]
dates: ["2025-10-15", "2026-07-30", "2026-09-09", "2026-09-10", "2026-12-31"]
keywords: ["agent", "agents", "benchmark", "benchmarks", "claude", "cyber", "deepseek", "gemini", "gemini 3.8", "glm", "gpt-5.6", "luna"]
source: docs/RAG/collect-261001-ia-llm/deepseek-v4-1-flash-vs-glm-5-3-flash-vs-gemini-3-8-flash.md
source_anchor: ""
source_lines: [27, 89]
sha256: 8b178a755050aca57f407c5a16bb59270148f4ff98b62ed114fcd83a59a38ddb
---

# Estimation du coût mensuel par modèle (en dollars)

Sur le plan tarifaire, Gemini 3.8 Flash reprend exactement le même tarif introductif que son prédécesseur Gemini 3.7 Flash : 0,75 $ par million de tokens en entrée et 3,75 $ en sortie (les tokens de raisonnement interne sont facturés comme de la sortie), valable jusqu’au 31 décembre 2026. La lecture en cache coûte 0,075 $ par million de tokens sur cette période. Cette tarification double automatiquement au 1er janvier 2027, passant à 1,50 $ l’entrée et 7,50 $ la sortie, une hausse déjà annoncée par Google et confirmée par plusieurs analyses de prix indépendantes. Le tarif Batch/Flex, réservé aux traitements différés, reste à environ la moitié du tarif standard (0,375 $ / 1,875 $), tandis que le tarif Priority grimpe à 1,35 $ / 6,75 $. Contrairement à GLM-5.3-Flash et DeepSeek V4.1-Flash, Google n’a pas communiqué de chiffre précis de fenêtre de contexte pour cette version spécifique dans sa documentation publique ; le modèle s’inscrit toutefois dans la continuité des générations Flash précédentes, positionnées sur de grandes fenêtres de contexte adaptées aux agents et au code long.

Google détaille l’intégralité de sa grille tarifaire, y compris les paliers Batch, Flex et Priority, dans sa documentation officielle de tarification de l’API Gemini, ainsi que dans l’article de blog annonçant Gemini 3.8 Flash et sa variante Gemini 3.8 Flash Cyber, une déclinaison orientée cybersécurité lancée le même jour. Pour les entreprises qui déploient leurs agents via la plateforme Gemini Enterprise Agent, la page de tarification Google Cloud dédiée confirme que Gemini 3.8 Flash, 3.7 Flash et 3.6 Flash partagent exactement la même grille introductive jusqu’à fin 2026, avant un doublement simultané au 1er janvier 2027 pour l’ensemble de ces trois versions.

## Tableau comparatif : les spécifications techniques face à face

Le tableau suivant réunit les caractéristiques techniques publiées pour chacun des trois modèles économiques, avec deux modèles de référence (GPT-5.6 Luna et Claude Haiku 4.5) pour situer le niveau de prix par rapport au reste du marché.

| Caractéristique | GLM-5.3-Flash | DeepSeek V4.1-Flash | Gemini 3.8 Flash | 
|---|---|---|---|
| Éditeur | Zhipu AI (Z.AI) | DeepSeek |  | 
| Date de sortie | 26 août 2026 | 10 septembre 2026 | 2 septembre 2026 | 
| Licence | Open weights (MIT) | Open weights (MIT) | Propriétaire | 
| Architecture | MoE | MoE, 552 Md paramètres | Non communiquée | 
| Paramètres actifs | Non communiqués | ~8 Md (entrée) / 16 Md (sortie) | Non communiqués | 
| Fenêtre de contexte | 1 000 000 tokens | ~1 000 000 tokens | Non communiquée précisément | 
| Sortie maximale | 131 100 tokens | Non précisée séparément | Non précisée séparément | 
| Prix entrée (par 1M tokens) | 0,15 $ | 0,15 $ (hors pointe) / 0,30 $ (pointe) | 0,75 $ (jusqu’au 31/12/2026) | 
| Prix sortie (par 1M tokens) | 0,50 $ | 0,60 $ (hors pointe) / 1,20 $ (pointe) | 3,75 $ (jusqu’au 31/12/2026) | 
| Lecture en cache (par 1M) | 0,03 $ | 0,003 $ (hors pointe) / 0,006 $ (pointe) | 0,075 $ | 
| MMLU-Pro | Non documenté publiquement | 83,0 % à 86,4 % | Non documenté publiquement | 
| LiveCodeBench | Non documenté publiquement | Jusqu’à 91,6 % | Non documenté publiquement | 
| Accès | API Z.AI, DeepInfra, Novita | API DeepSeek, hébergeurs tiers | API Gemini, Google AI Studio, Vertex AI | 

Ce tableau met en évidence une asymétrie nette : GLM-5.3-Flash et DeepSeek V4.1-Flash publient leurs poids et jouent la carte de la transparence tarifaire simple, quand Gemini 3.8 Flash reste fermé et applique un système de paliers (standard, batch, flex, priority) plus complexe à modéliser pour un budget prévisionnel. À l’inverse, DeepSeek complique l’équation avec son système peak/off-peak, qui demande une vraie discipline de planification pour en tirer le meilleur prix.

## Combien coûte réellement chaque modèle : le comparatif des prix

Pour resituer ces trois modèles économiques dans le paysage plus large des offres IA, voici un comparatif incluant deux références couramment utilisées comme benchmark de prix : GPT-5.6 Luna d’OpenAI et Claude Haiku 4.5 d’Anthropic.

| Modèle | Entrée ($/1M) | Sortie ($/1M) | Contexte | Dernière évolution tarifaire | 
|---|---|---|---|---|
| GLM-5.3-Flash | 0,15 $ | 0,50 $ | 1M tokens | Promo à 0,075 $/0,25 $ expirée le 9/09/2026, retour au tarif liste | 
| DeepSeek V4.1-Flash (hors pointe) | 0,15 $ | 0,60 $ | ~1M tokens | Lancé le 10/09/2026 avec système peak/off-peak | 
| DeepSeek V4.1-Flash (pointe) | 0,30 $ | 1,20 $ | ~1M tokens | Tarif doublé aux heures de forte demande | 
| GPT-5.6 Luna | 0,20 $ | 1,20 $ | ≥272K tokens | Baisse de 80 % le 30/07/2026 (1 $/6 $ → 0,20 $/1,20 $) | 
| Gemini 3.8 Flash (2026) | 0,75 $ | 3,75 $ | Non précisée | Tarif introductif jusqu’au 31/12/2026 | 
| Gemini 3.8 Flash (dès 2027) | 1,50 $ | 7,50 $ | Non précisée | Doublement programmé au 1er janvier 2027 | 
| Claude Haiku 4.5 | 1,00 $ | 5,00 $ | 200K tokens | Tarif stable depuis son lancement le 15/10/2025 | 

L’écart entre le modèle le moins cher en sortie (GLM-5.3-Flash à 0,50 $) et le plus cher (Gemini 3.8 Flash à 3,75 $) atteint un facteur x7,5 sur le seul prix de la sortie, avant même le doublement prévu par Google en 2027. Si l’on compare Claude Haiku 4.5 à GLM-5.3-Flash, l’écart grimpe à x10 en sortie. Ces chiffres confirment une tendance déjà documentée sur ce site : DeepSeek V4-Flash affichait déjà un écart de prix de 21x face à Claude Haiku 4.5 et GPT-5.6 Luna lors de sa première comparaison, un rapport qui reste globalement d’actualité avec cette nouvelle génération V4.1.

Pour visualiser l’impact concret sur une facture mensuelle, voici un script simple qui calcule le coût pour un volume donné de tokens, à adapter avec vos propres chiffres d’usage :

```
# Estimation du coût mensuel par modèle (en dollars)
tarifs = {
    "GLM-5.3-Flash":            {"entree": 0.15, "sortie": 0.50},
    "DeepSeek V4.1-Flash (hors pointe)": {"entree": 0.15, "sortie": 0.60},
    "Gemini 3.8 Flash (2026)":  {"entree": 0.75, "sortie": 3.75},
    "GPT-5.6 Luna":             {"entree": 0.20, "sortie": 1.20},
    "Claude Haiku 4.5":         {"entree": 1.00, "sortie": 5.00},
}
tokens_entree_millions = 10   # ex : 10 millions de tokens en entrée / mois
tokens_sortie_millions = 2    # ex : 2 millions de tokens en sortie / mois
for modele, prix in tarifs.items():
    cout = (tokens_entree_millions * prix["entree"]) + (tokens_sortie_millions * prix["sortie"])
    print(f"{modele} : {cout:.2f} $ / mois")
```
Sur ce volume illustratif de 10 millions de tokens en entrée et 2 millions en sortie par mois, GLM-5.3-Flash revient à 2,50 $, DeepSeek V4.1-Flash hors pointe à 2,70 $, GPT-5.6 Luna à 4,40 $, Gemini 3.8 Flash à 15,00 $ et Claude Haiku 4.5 à 20,00 $. L’écart se creuse mécaniquement à mesure que le volume augmente, ce qui explique pourquoi les équipes à fort trafic (support client, modération, extraction de données) surveillent de très près ces grilles tarifaires.

## Benchmarks : MMLU-Pro, LiveCodeBench et la fiabilité du code

