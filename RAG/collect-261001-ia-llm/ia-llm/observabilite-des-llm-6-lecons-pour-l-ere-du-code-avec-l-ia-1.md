---
id: collect-261001-ia-llm/ia-llm/observabilite-des-llm-6-lecons-pour-l-ere-du-code-avec-l-ia-1
title: "observabilite-des-llm-6-lecons-pour-l-ere-du-code-avec-l-ia"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "arr", "incident", "valuation"]
source: docs/RAG/collect-261001-ia-llm/observabilite-des-llm-6-lecons-pour-l-ere-du-code-avec-l-ia.md
source_anchor: ""
source_lines: [1, 97]
sha256: 036868fe5c8dd59c141a027ece8c0198366e3c04a671c9b06d3cc63042422211
---

# observabilite-des-llm-6-lecons-pour-l-ere-du-code-avec-l-ia

Cursus

Les équipes d'ingénierie livrent aujourd'hui plus de code qu'elles ne peuvent en lire. Les assistants IA en écrivent désormais une grande partie, plus vite que n'importe quel relecteur ne peut suivre ligne par ligne. Ce basculement sert de toile de fond à la conférence DASH de Datadog à New York cette semaine, où le cofondateur et CTO Alexis Lê-Quôc anime une session intitulée « The New Shape of Engineering ».

Son constat est simple : la façon d'exploiter les logiciels n'a pas changé : on livre une modification, on la déploie, puis on observe. Ce qui change, c'est le volume et la cadence — et ça modifie ce qui garantit la sécurité de l'ensemble.

Dans cet article, je synthétise sa pensée en six leçons clés : évolution de la revue de code, production comme test ultime, et ce que vous devez en retenir.

Si vous découvrez l'observabilité des LLM, nous vous recommandons de lire nos guides pour bien démarrer avec le MLOps et l'évaluation des LLM comme point de départ.

## En bref

Le fil conducteur de Lê-Quôc : l'observabilité devient la couche de contrôle du logiciel écrit, testé et déployé par l'IA, au service des opérateurs comme des agents eux-mêmes.

Les six leçons, en résumé :

- **La revue se déplace hors du code lui-même.** Il y a trop de code généré par l'IA pour le lire ligne à ligne ; le véritable contrôle, ce sont les tests, spécifications et preuves que vous concevez en amont, en prévoyant aussi des garde-fous contre des agents qui chercheraient à biaiser ces tests.
- **La production est le seul test qui compte.** Un pipeline CI tout au vert ne prouve pas grand-chose face aux usages réels que vous n'aviez pas pu anticiper, et la sortie d'un modèle n'est jamais totalement certaine ; il faut donc la surveiller en direct et garder un bouton d'arrêt.
- **Laissez les agents prendre la pénibilité.** Confiez-leur la surveillance de tableaux de bord et la vérification d'hypothèses qui épuisent les humains, et gardez les personnes pour les décisions à fort jugement.
- **Séparez le travail en deux boucles :** une boucle de développement (écrire, livrer, vérifier, corriger) et une boucle opérations et sécurité (détecter, enquêter, résoudre).
- **Maîtrisez les dépenses d'IA.** Dimensionnez quel modèle fait quelle tâche à partir des trajectoires des agents, et laissez cette décision aux développeurs et SRE qui la prennent.
- **Apprenez à apprendre.** Les modèles sont des tuteurs infatigables, mais la compétence est de les questionner : comprendre les systèmes couche par couche et demander pourquoi le code qu'ils ont écrit a vraiment fonctionné.

## Améliorez les compétences de votre équipe en matière d'IA

Transformez votre entreprise en dotant votre équipe de compétences avancées en matière d'IA grâce à DataCamp for Business. Améliorez vos connaissances et votre efficacité.

## Leçon 1 : l'IA a fait voler en éclats l'ancienne revue de code

Commençons par la pression qui conditionne tout le reste : il y a plus de code qu'aucun humain ne peut en lire.

Lê-Quôc est direct : le modèle historique, un humain lisant une pull request ligne à ligne, ne résiste pas au développement assisté par IA. L'inquiétude qu'il entend dans tout le secteur porte sur l'impossibilité de la revue : il se passe trop de choses pour suivre en lisant des PR.

Sa réponse n'est pas de demander aux gens de lire plus vite, mais de déplacer la revue ailleurs.

La revue n'est plus à la ligne de code ; il y en a trop, vous ne pouvez pas suivre. Il s'agit des tests que nous conço ns en amont, et d'interdire à l'agent de les contourner.

Alexis Lê-Quôc, CTO at Datadog 

Ce dernier point est facile à manquer. Dès lors que vous orchestrez un agent pour planifier, un autre pour écrire et un troisième pour tester, vous devez aussi empêcher l'écrivain de jouer avec les tests automatisés au lieu de résoudre le problème.

Il va au-delà des tests. Datadog ajoute désormais des preuves semi-formelles et formelles que la spécification fait bien ce qu'elle doit faire — une démarche trop coûteuse à généraliser avant que les agents ne prennent le gros du travail. Cela fonctionne particulièrement bien sur les systèmes backend et de coordination, où le comportement est suffisamment mathématique pour raisonner avec précision.

## Leçon 2 : la production est le seul test qui compte

Réussir tous les tests en CI est nécessaire, et très loin d'être suffisant. Les défaillances qui comptent arrivent plus tard.

Là où ça compte vraiment, c'est en production.

Alexis Lê-Quôc, CTO at Datadog 

Chaque livraison repose sur des hypothèses impossibles à vérifier pleinement en amont : la forme des données et le comportement des utilisateurs. Exposez ces hypothèses à suffisamment de trafic réel, et les cas rares cessent de l'être : ils deviennent les lenteurs et erreurs quotidiennes de la dérive des données et des modèles.

Les LLM compliquent encore la donne : avec du code classique, on peut au moins raisonner sur chaque branche. Personne ne peut expliquer mécaniquement pourquoi un modèle renvoie telle réponse ; le même input ne garantit jamais le même output. Les résultats étranges occasionnels ne peuvent pas être éradiqués par l'ingénierie.

Vous cessez donc d'essayer de prouver qu'un système est correct avant la mise en production. À la place, vous :

- Rédigez des évaluations du comportement attendu
- Le surveillez en production
- Gardez un dispositif d'arrêt pour un déploiement qui tourne mal.

La question n'est plus de savoir s'il a réussi, mais si un problème est isolé ou le début d'une tendance.

Ce signal en direct n'est pas qu'un tableau de bord pour humains. Branché au système de déploiement, il permet à un agent de dérouler un changement comme le ferait un ingénieur prudent : 1 % des utilisateurs, puis 5 %, en jugeant sur données réelles si la modification produit l'effet attendu.

## Leçon 3 : laissez les agents prendre la pénibilité

Pour Lê-Quôc, les agents ne remplacent pas les ingénieurs : ils prennent les tâches qui usent.

Diagnostiquer un incident, c'est tester des hypothèses face à un symptôme et, lors des incidents longs, c'est souvent l'hypothèse improbable qui s'avère juste. L'agent Bits AI de Datadog les vérifie toutes en parallèle, en amont de l'ingénieur, tandis que la personne l'oriente vers l'intuition qu'aucun dashboard ne ferait ressortir.

Le fond de l'affaire, c'est la fatigue. Un déploiement en astreinte, c'est des pics d'alerte suivis d'heures de vide, répétés jusqu'à éroder le jugement.

Vous êtes en mode alerte maximale, puis vous regardez la peinture sécher.

Alexis Lê-Quôc, CTO at Datadog 

Un agent ne s'en formalise pas et ne décline pas après quatre heures à fixer des chiffres. Le stress et la fatigue dégradent les performances humaines, raison pour laquelle on fait tourner les personnes en astreinte.

Confiez la veille inlassable à une machine, et les équipes reviennent reposées pour les décisions qui comptent. Même logique en sécurité, où les analystes s'épuisent à trier faux positifs et vraies menaces.

## Leçon 4 : séparez le travail en deux boucles

Lê-Quôc organise le travail des agents chez Datadog autour de deux boucles.

### La boucle de développement

La première boucle sera familière à la plupart des ingénieurs :

1. Écrire du code
2. Le livrer
3. Vérifier si ça fonctionne
4. Corriger
5. Répéter

L'angle de Datadog : un problème qui naît dans le code se répare généralement dans le code. La plateforme tente donc de vous proposer cette correction, à partir de ce qu'elle sait de l'application : propriétaires, changements récents, erreurs remontées.

