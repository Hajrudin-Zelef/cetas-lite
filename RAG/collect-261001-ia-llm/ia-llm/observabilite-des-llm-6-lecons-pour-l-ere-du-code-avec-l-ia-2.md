---
id: collect-261001-ia-llm/ia-llm/observabilite-des-llm-6-lecons-pour-l-ere-du-code-avec-l-ia-2
title: "observabilite-des-llm-6-lecons-pour-l-ere-du-code-avec-l-ia"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "agents", "arr", "claude"]
source: docs/RAG/collect-261001-ia-llm/observabilite-des-llm-6-lecons-pour-l-ere-du-code-avec-l-ia.md
source_anchor: ""
source_lines: [98, 158]
sha256: 7215abadfcfb93e6a191ddfe45b27c6105108b221091619895283d05e2c09e97
---

# observabilite-des-llm-6-lecons-pour-l-ere-du-code-avec-l-ia

Il cite l'optimisation de requêtes base de données comme exemple. N'importe quel modèle peut réécrire une requête lente ; le plus difficile est de prouver que la version réécrite est plus rapide et sûre avant d'atteindre la production. Datadog la teste donc sur une copie réaliste des données de production, puis fournit une pull request assortie des preuves.

### La boucle opérations et sécurité

L'autre boucle tourne en parallèle, chez les mêmes personnes ou une équipe différente :

1. Détecter
2. Enquêter
3. Corriger
4. Répéter

C'est là que l'AI Guard de Datadog priorise les événements de sécurité et bloque les attaques plus vite qu'un analyste ne le ferait à la main. Les agents peuvent également gérer des tâches opérationnelles routinières que les ingénieurs effectuent chaque jour sans enthousiasme, comme redimensionner ce fameux pod Kubernetes.

Dans les deux boucles, Lê-Quôc reste clair sur l'ordre des priorités. Datadog ne part pas de « voici l'IA, quel problème peut-elle résoudre ? » : on part d'une douleur client avérée, généralement une variante de « je ne veux plus faire cette tâche répétitive », puis on évalue si un agent est digne de confiance pour s'en charger.

## Leçon 5 : maîtrisez les dépenses d'IA

Le coût est la contrainte jumelle de la sécurité, et contenir le prix de la mise en production des grands modèles de langage devient une discipline à part entière. La réponse présentée par Lê-Quôc à DASH : l'Agent Console de Datadog.

Demandez à un développeur quel modèle il lui faut : souvent, il citera le plus puissant (et le plus cher). Parfois, c'est le bon choix, mais une grande part du travail est du générique qu'un modèle plus économique et rapide gère tout aussi bien. Les distinguer suppose d'analyser les trajectoires des agents d'une organisation : quels outils ils appellent, à quelle fréquence ils réussissent, jusqu'à faire apparaître des motifs.

Ces motifs deviennent des heuristiques plutôt que des règles : un modèle de pointe comme le dernier Claude Opus ou les modèles GPT pour la planification, un modèle économique comme Claude Haiku pour générer des tests.

| Tâche | Niveau de modèle | Pourquoi | 
|---|---|---|
| Planification et raisonnement complexe | Modèle de pointe (ex. : Claude Opus, GPT) | La meilleure capacité de raisonnement se rentabilise ici | 
| Code routinier, générique | Niveau intermédiaire (ex. : Claude Sonnet, GPT-mini) | Assez performant, et bien moins coûteux à exécuter fréquemment | 
| Génération de tests et transformations simples | Rapide et peu cher (ex. : Claude Haiku, GPT-nano) | La vitesse et le prix l'emportent tant que la qualité tient | 

Le principe sous-jacent concerne la propriété de la décision. Si vous remontez le coût à un seul chiffre, vous obtenez ce que Lê-Quôc appelle une « très faible actionnabilité » : soit tout le monde coupe les dépenses, et on tue des travaux utiles, soit tout le monde continue, et l'entreprise ne peut pas suivre. Il préfère mettre les données sous les yeux des développeurs et SRE qui choisissent les modèles.

## Leçon 6 : apprenez à apprendre

Interrogé sur ce que les nouveaux ingénieurs devraient étudier, Lê-Quôc donne une réponse qui paraît ancienne, mais ne l'est pas.

Vous devez apprendre à apprendre.

Alexis Lê-Quôc, CTO at Datadog 

Les modèles sont les tuteurs les plus patients jamais inventés, capables d'expliquer n'importe quoi à n'importe quel rythme — un niveau d'accès autrefois réservé à quelques privilégiés. Mais un tuteur n'est utile que si vous l'interrogez. La compétence, c'est savoir quoi demander et comment vérifier la réponse.

Il recommande de comprendre l'informatique couche par couche plutôt que de la traiter comme de la magie. Prenez un ordonnanceur, un équilibreur de charge, un bac à sable, et demandez à un modèle d'expliquer son fonctionnement, puis creusez :

- Que signifie ce terme ?
- Comment le mesurer ?
- Quelles sont les bases mathématiques ?
- Comment savoir si cela fonctionne bien ?

Étudier les classiques de cette manière est volontairement lent. Il compare cela à l'apprentissage d'un instrument : vous pouvez écouter de la musique toute la journée, mais pour jouer du piano, il faut poser les mains sur le clavier.

Même chose pour le code écrit par l'IA. Le vibe coding est très bien, dit-il, à condition d'y revenir et de demander pourquoi ça a marché : pourquoi tel choix d'architecture, existe-t-il de meilleures approches, sur quoi cela s'est-il fondé. Le but n'est pas d'écrire moins de code avec l'IA, mais de mieux comprendre le code que vous produisez désormais en bien plus grande quantité.

## En conclusion

Le message central de Lê-Quôc : la boucle n'a pas changé, mais la vitesse, si. Désormais, aucun humain ne peut observer d'assez près à la cadence de l'IA : la surveillance, et une part croissante de la construction, passent à des agents qui ne se fatiguent ni ne paniquent.

Il plaide pour traiter l'observabilité comme un plan de contrôle, pas comme une collection de graphiques. Si des agents écrivent, testent, livrent et exploitent du logiciel, ils ont besoin du même ancrage dans les données de production réelles que les bons ingénieurs, avec en plus une personne qui garde le jugement et le bouton d'arrêt. Datadog positionne l'observabilité comme la couche qui rend cet équilibre sûr.

La compétence attendue des ingénieurs est claire : lire les systèmes à travers leur comportement en production, pas seulement via leur source. Pour ancrer cette habitude, notre parcours de compétences Machine Learning in Production est un bon point de départ.

**Rédacteur en chef Data Science chez DataCamp |** **Je suis passionné par la prévision et le développement à l'aide d'API.**
