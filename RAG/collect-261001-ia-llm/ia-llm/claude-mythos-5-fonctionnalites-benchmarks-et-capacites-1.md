---
id: collect-261001-ia-llm/ia-llm/claude-mythos-5-fonctionnalites-benchmarks-et-capacites-1
title: "claude-mythos-5-fonctionnalites-benchmarks-et-capacites"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Glasswing", "Google", "OpenAI", "Stripe"]
dates: []
keywords: ["benchmarks", "claude", "fable 5", "gemini", "mythos 5", "opus 4", "reasoning", "valuation"]
source: docs/RAG/collect-261001-ia-llm/claude-mythos-5-fonctionnalites-benchmarks-et-capacites.md
source_anchor: ""
source_lines: [1, 57]
sha256: 0e0b724d4748012b491ab5c48f4aab43200e367392b21881d3529764437275fe
---

# claude-mythos-5-fonctionnalites-benchmarks-et-capacites

Cursus

**Mise à jour (1er juillet 2026) :** l'accès à Claude Fable 5 a été rétabli et est désormais disponible dans le monde entier sur Claude Platform, Claude.ai, Claude Code et Claude Cowork, suite à la levée de l'ordonnance de contrôle des exportations. Mythos 5 reste réservé à des partenaires Project Glasswing dûment approuvés.

Anthropic a lancé deux modèles le 9 juin 2026 : Claude Fable 5, la version publique de la classe Mythos avec des garde-fous de sécurité conservateurs, et Claude Mythos 5, le même modèle sous-jacent avec ces garde-fous levés pour un groupe restreint de partenaires de confiance. Cet article porte sur Mythos 5, la version qu’Anthropic décrit comme ayant « les capacités de cybersécurité les plus avancées au monde ».

Dans cet article, nous verrons ce qu'est Claude Mythos 5, ce qu'il sait faire en génie logiciel, sciences de la vie et recherche scientifique, ses performances sur les benchmarks, ainsi que les modalités d'accès. Vous pouvez aussi consulter notre analyse de Claude Opus 4.8 pour situer Mythos 5 au sein de la gamme Anthropic. Si vous découvrez Claude et souhaitez démarrer rapidement, nous recommandons notre guide pour apprendre à utiliser Claude.

Restez au fait des dernières évolutions de l’IA. Abonnez-vous à *The Median*, notre newsletter gratuite du vendredi qui décrypte l’actu de la semaine. Faites votre veille en quelques minutes.


## Qu'est-ce que Claude Mythos 5 ?

Claude Mythos 5 est le modèle le plus performant d’Anthropic, positionné au-dessus de la classe Opus dans ce qu’Anthropic appelle le niveau Mythos. Le premier modèle de cette classe, Claude Mythos Preview, est sorti en avril 2026 dans le cadre de Project Glasswing, une collaboration avec le gouvernement américain axée sur la cybersécurité. Mythos 5 est la deuxième version de ce niveau et une mise à jour directe de Mythos Preview.

Mythos 5 et Fable 5 partagent la même architecture. La différence tient aux garde-fous : Fable 5 embarque des classifieurs qui redirigent les requêtes sensibles en cybersécurité et en biologie vers Claude Opus 4.8. Mythos 5 lève ces classifieurs sur des domaines précis pour des partenaires évalués dans le cadre du programme d’accès de confiance. Anthropic précise que la différence de nom reflète les garde-fous, pas les capacités.

Le principal argument chiffré : Mythos 5 obtient 80,3 % sur SWE-bench Pro, contre 77,8 % pour Mythos Preview et 69,2 % pour Opus 4.8. Sur Humanity's Last Exam avec outils, il atteint 64,5 %, devant les 57,9 % d’Opus 4.8 et les 52,2 % de GPT-5.5. Ce ne sont pas des gains marginaux par rapport à la classe Opus.

## Présentation des modèles Claude

## Quoi de neuf avec Claude Mythos 5 ?

Mythos 5 marque une progression par rapport à Mythos Preview sur toutes les grandes capacités testées par Anthropic. Les gains sont particulièrement visibles sur les travaux autonomes de longue haleine, notamment pour le raisonnement scientifique et les tâches de vision. Voici ce que cela donne concrètement.

### Ingénierie logicielle autonome, sécurisée et à grande échelle

Mythos 5 peut travailler de manière autonome sur de vastes bases de code plus longtemps que tout autre modèle Claude. Stripe a indiqué que le modèle a compressé des mois de travail d’ingénierie en quelques jours, en menant une migration à l’échelle d’un code Ruby de 50 millions de lignes en une seule journée. Sur FrontierCode (Diamond), il obtient les meilleurs scores parmi les modèles de pointe, même à effort moyen.

Pour les missions de sécurité, Mythos 5 étend les capacités qui avaient fait la valeur de Mythos Preview pour les partenaires de Project Glasswing. Ces partenaires ont utilisé Mythos Preview pour identifier plus de 10 000 vulnérabilités majeures et critiques sur des systèmes en production.

### Conception de médicaments et ingénierie des protéines

L’équipe interne de conception de protéines d’Anthropic a utilisé Mythos 5 pour accélérer la conception de médicaments par un facteur d’environ dix. Dans une comparaison contrôlée, Mythos 5 a égalé ou surpassé des opérateurs humains expérimentés sur 14 cibles protéiques pour l’ensemble du pipeline :

- choix des sites de liaison
- sélection des outils
- récupération après échec

Neuf ont produit des candidats médicaments prometteurs actuellement à l’étude.

### Génération d’hypothèses scientifiques inédites

Mythos 5 est le premier modèle d’Anthropic à produire régulièrement des hypothèses scientifiques originales, et non de simples synthèses de la littérature. Lors de comparaisons en aveugle, les scientifiques d’Anthropic ont préféré ses hypothèses en biologie moléculaire dans environ 80 % des cas, et plusieurs ont été avancées à l’évaluation expérimentale. Une hypothèse concernant un nouveau mécanisme protéique chez *E. coli* a été corroborée indépendamment par un laboratoire travaillant sur le même sujet.

### Recherche autonome en génomique

Mythos 5 a mené une recherche génomique inédite pendant plus d’une semaine de travail largement autonome, en assemblant des données monocellulaires pour des millions de cellules sur 138 espèces animales et en entraînant un modèle de ML personnalisé pour identifier des types cellulaires équivalents chez des organismes éloignés. Le modèle entraîné a surpassé un modèle récemment publié dans Science tout en étant 100 fois plus petit.

### Vision et performance en long contexte

Mythos 5 atteint 93,2 % sur CharXiv Reasoning avec outils et peut extraire des chiffres précis à partir de figures scientifiques détaillées ou reconstruire une application web à partir de simples captures d’écran. Sur les tâches à long contexte, l’ajout d’une mémoire basée sur des fichiers a triplé son gain de performance par rapport à une configuration identique avec Opus 4.8, et il a atteint l’acte final de Slay the Spire trois fois plus souvent.

## Benchmarks de Claude Mythos 5

Mythos 5 mène ou égalise sur quasiment tous les benchmarks testés par Anthropic, avec des gains sur Opus 4.8 cohérents entre catégories plutôt que concentrés sur un seul domaine. Le tableau comparatif l’oppose à Claude Mythos Preview, Claude Opus 4.8, GPT 5.5 et Gemini 3.1 Pro.

