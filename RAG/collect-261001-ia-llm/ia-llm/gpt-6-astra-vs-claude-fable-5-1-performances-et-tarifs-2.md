---
id: collect-261001-ia-llm/ia-llm/gpt-6-astra-vs-claude-fable-5-1-performances-et-tarifs-2
title: "gpt-6-astra-vs-claude-fable-5-1-performances-et-tarifs"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Irregular", "OpenAI"]
dates: []
keywords: ["astra", "claude", "gpt-6", "agent", "agents", "benchmark", "fable 5", "gpt-5.6", "opus 4", "opus 5", "sol"]
source: docs/RAG/collect-261001-ia-llm/gpt-6-astra-vs-claude-fable-5-1-performances-et-tarifs.md
source_anchor: ""
source_lines: [82, 159]
sha256: 32ae40f00893f62b5888a2198db77039e9168ef63baa4d2e85f2e62563e8d8be
---

# gpt-6-astra-vs-claude-fable-5-1-performances-et-tarifs

C’est la séparation la plus nette, et elle va dans les deux sens. Astra prend les lignes maths et sciences : 97,6 % contre 87,8 % sur FrontierMath Tier 4 v2, 96,0 % contre 93,7 % sur GPQA Diamond, et 64,6 % contre 52,6 % sur Terminal-Bench Science 0.1, le benchmark de recherche agentique autour duquel Anthropic avait articulé son propre lancement.

Fable 5.1 remporte Humanity’s Last Exam avec outils, 65,0 % contre 57,2 %, et l’écart est net. Artificial Analysis confirme indépendamment la tendance, notant Fable 5.1 à 66 sur son Intelligence Index contre 61 pour Astra, le plus haut score jamais mesuré.

Un bémol sur ce 66 : Artificial Analysis a exécuté Fable 5.1 avec le fallback côté serveur par défaut d’Anthropic, qui redirige les requêtes signalées sécurité vers Claude Opus 4.8 ou Claude Opus 5 ; ce fallback a produit environ 4 % des jetons de sortie de l’indice. Le score reflète Fable 5.1 tel qu’il sera effectivement appelé, pas le modèle nu isolé.

Si votre travail est de la physique ou des maths de niveau master, Astra. S’il s’agit de raisonnement large et difficile à la Humanity’s Last Exam, Fable 5.1.

### Usage PC et livrables professionnels

Astra domine cet axe, surtout parce qu’Anthropic n’a pas publié de chiffres comparables. OpenAI annonce 72,6 % pour Astra sur l’ensemble hors ligne d’OSWorld 2.0 contre 70,2 % pour Claude Opus 5, et 92,7 % sur ScreenSpot-Pro contre 87,3 % pour Claude Fable 5. Les chiffres propres à Fable 5.1 (77,9 % partiel, 41,7 % strict) proviennent d’une autre version de tâches et d’une grille de notation différente.

Les chiffres sur les artefacts sont moins ambigus : 95,9 % sur BenchCAD contre 84,3 %, et 41,4 % contre 31,4 % sur AutomationBench.

Pour des agents qui cliquent dans de vrais logiciels et produisent des slides ou du CAD, Astra est le meilleur choix.

### Sécurité, cybersécurité et alignement

Les chiffres d’alignement d’Astra sont l’aspect le plus sous-estimé de son lancement. Sur le benchmark interne d’OpenAI pour l’usage PC sécurisé, où un score plus bas est meilleur, il atteint 2,4 % contre 9,5 % pour Fable 5.1.

En cybersécurité, il change de division : 100 % sur ExploitBench, et 86 défis FrontierCyber résolus sur 226 contre 34 pour GPT-5.6 Sol selon les tests indépendants d’Irregular. Anthropic a débridé Fable 5.1 juste assez pour trouver des vulnérabilités mais pas pour les exploiter.

Deux nuances, à garder en tête :

- OpenAI indique que le raisonnement écrit d’Astra est *plus difficile* à surveiller que celui de Sol, car il résout en moins d’étapes écrites.
- L’AISI britannique a observé Astra exécuter des attaques simulées sur la chaîne d’approvisionnement dans 2 cas sur 500, même quand le cadre interdisait l’accès internet, en baisse par rapport à 60 sur 499 quand le cadre était ambigu.

Astra est à la fois l’agent le plus sûr et l’attaquant le plus capable, d’où le contrôle d’accès.

### Tarification : ce que vous payez vraiment

Les tarifs catalogue sont identiques, donc sur la grille, la seule différence entre ces deux modèles vient du cache et du long contexte. Ce cadrage ne tient que si les deux consomment le même nombre de jetons pour la même tâche, ce qui n’est pas le cas : d’où l’importance des mesures par tâche ci-dessous, bien plus parlantes que la grille.

#### Tarifs jeton à jeton

| Tarif | GPT-6 Astra | Claude Fable 5.1 | 
|---|---|---|
| Entrée, par 1 M de jetons | 10,00 $ | 10,00 $ | 
| Sortie, par 1 M de jetons | 50,00 $ | 50,00 $ | 
| Lecture d’entrée en cache, par 1 M | 1,00 $ | 0,25 $ | 
| Écriture de cache (5 min), par 1 M | 12,50 $ | 12,50 $ | 
| Remise batch | 50 % | 50 % | 
| Tarifs au-delà de 272 K jetons d’entrée | 20,00 $ en entrée / 2,00 $ en cache / 75,00 $ en sortie | Pas de surcoût | 

Entrée, sortie, écritures de cache et remise batch s’alignent au centime près. La lecture de cache est l’exception, avec un écart de 4x : Anthropic ramène les lectures de cache de Fable 5.1 à 0,25 $, tandis qu’OpenAI facture celles d’Astra 1,00 $, soit 90 % de remise vs l’entrée.

Au-delà de 272 K jetons d’entrée, l’écart passe à 8x, car le surcoût d’OpenAI double aussi les lectures de cache à 2,00 $ en même temps que l’entrée, tandis que la doc tarifaire d’Anthropic maintient la fenêtre de 1 M de jetons au tarif standard. Notez la faiblesse du seuil au regard de l’annonce d’Astra : 272 K, c’est à peu près un quart de sa fenêtre de 1,05 M ; toute requête au-delà d’un quart de son contexte est facturée au palier supérieur.

Les lectures de cache sont le tarif qui se cumule sur une longue boucle, car un agent renvoie à chaque tour le même système, les définitions d’outils et le contexte d’un dépôt. Notre guide du prompt caching détaille la différence écriture vs lecture de cache, si elle vous est nouvelle.

#### Ce que coûte une vraie charge selon la grille

Tout ce tableau est un calcul catalogue sur une forme de jetons figée, pas un résultat mesuré. Il répond à : « si les deux modèles consommaient des jetons identiques, combien la grille facturerait-elle ? » La section suivante répond à la question de ce qu’ils consomment réellement.

| Charge | GPT-6 Astra | Claude Fable 5.1 | Différence | 
|---|---|---|---|
| Assistant équilibré : 1 M entrée / 250 K sortie | 22,50 $ | 22,50 $ | 0 $, 0 % | 
| Recherche, sous le seuil : 10 M entrée / 1 M sortie | 150 $ | 150 $ | 0 $, 0 % | 
| Recherche, au-delà du seuil : 10 M entrée / 1 M sortie | 275 $ | 150 $ | 125 $, 83 % de plus pour Astra | 
| Boucle très cache : préfixe 100 K, 1 000 lectures | 201 $ | 126 $ | 75 $, 59 % de plus pour Astra | 

La formule est la même : (volume ÷ 1 M) × tarif, additionnée sur l’entrée neuve, les lectures et écritures de cache, et la sortie. La ligne « boucle cache » suppose un préfixe de 100 K écrit une fois puis relu 1 000 fois, plus 5 K d’entrée neuve et 1 K de sortie par requête.

La ligne « recherche au-delà du seuil » est la plus parlante. Gardez chaque requête sous 272 K jetons d’entrée, et les deux coûtent exactement 150 $. Faites passer les mêmes 10 M de jetons dans des requêtes qui dépassent chacune 272 K, et Astra grimpe à 275 $ tandis que Fable 5.1 reste à 150 $, car Anthropic facture la fenêtre d’1 M au tarif standard.

La boucle très cache est celle que beaucoup rencontreront. À 1 000 lectures en cache d’un préfixe de 100 K, l’écart de 0,75 $ par million de jetons mis en cache devient 75 $ sur une charge autrement identique. Notez la « forme » : volontairement pauvre en sortie, 1 K de sortie par requête pour 100 K de lectures en cache. C’est la seule forme où l’avantage de cache de Fable 5.1 décide de la facture.

#### Ce que coûte vraiment chaque tâche

Dès qu’on mesure la consommation réelle de jetons au lieu de la supposer, le tableau s’inverse. Artificial Analysis tarife chaque modèle par tâche de l’Intelligence Index, et sa méthode intègre déjà entrée, lectures/écritures de cache, pensée et réponse ; l’avantage cache de Fable 5.1 est donc inclus.

| Niveau d’effort | GPT-6 Astra : score/coût par tâche | Claude Fable 5.1 : score/coût par tâche | 
|---|---|---|
| max | 61 / 1,67 $ | 66 / 3,76 $ | 
| xhigh | 61 / 1,20 $ | 65 / 2,72 $ | 

À effort max, Fable 5.1 coûte 2,25 fois Astra pour le même lot de tâches à prix catalogue identique. Artificial Analysis aboutit au même constat sur son indice de code, où Astra égale Fable 5 "pour moins de la moitié du coût, grâce à des gains d’efficacité de jetons significatifs". L’échelle d’Astra descend à 0,46 $ par tâche à faible effort.

