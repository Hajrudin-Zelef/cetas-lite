---
id: collect-261001-ia-llm/ia-llm/pourquoi-les-agents-dia-peinent-davantage-avec-le-code-data-science-quavec-le-code-genie-l-1
title: "pourquoi-les-agents-dia-peinent-davantage-avec-le-code-data-science-quavec-le-code-genie-logiciel"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "distribution"]
source: docs/RAG/collect-261001-ia-llm/pourquoi-les-agents-dia-peinent-davantage-avec-le-code-data-science-quavec-le-code-genie-logiciel.md
source_anchor: ""
source_lines: [1, 104]
sha256: dddd872d788f425114c241e0415bb7c6136acc135738edc163269d6e18301916
---

# pourquoi-les-agents-dia-peinent-davantage-avec-le-code-data-science-quavec-le-code-genie-logiciel

Cursus

J’ai analysé des centaines de dépôts GitHub pour comprendre pourquoi les agents de codage IA paraissent souvent beaucoup plus compétents en génie logiciel qu’en data science. Ce que j’ai découvert montre qu’il ne s’agit pas seulement d’un écart d’outillage. Cela renvoie à une différence plus profonde : l’endroit où se loge le sens dans chaque type de code.

## De quoi parle cet article

Si vous avez utilisé un agent de codage IA sur une base de code de génie logiciel, vous avez sans doute constaté son efficacité. L’agent navigue dans l’architecture, suit les abstractions et apporte des modifications qui s’intègrent étonnamment bien au reste du système.

Puis vous ouvrez un notebook de data science, et l’expérience change souvent.

L’agent sait toujours écrire du code valide. Il sait toujours suivre des instructions. Mais il ne saisit pas toujours l’essentiel : pourquoi ce jeu de données, pourquoi ce filtre, pourquoi cette fenêtre temporelle, pourquoi cette sortie a changé l’orientation de l’analyse.

Il traite le notebook comme un projet logiciel, alors que ce n’en est qu’une partie.

Je voulais comprendre pourquoi. J’ai donc analysé des centaines de dépôts GitHub en data science et en génie logiciel, en mesurant l’entropie, les schémas de référence et les comportements de couplage.

Mes conclusions m’ont surpris, et je pense qu’elles ont de vraies implications pour quiconque construit ou utilise des agents d’IA dans des travaux analytiques.

## L’inversion d’entropie : le code data science n’est pas ce qu’il paraît

### Ce que j’ai réellement mesuré

J’ai mesuré l’entropie de Shannon à trois niveaux d’abstraction pour chaque dépôt : au niveau caractère, au niveau jeton, et au niveau AST. Chacun capture une dimension différente de la variation du code.


Distribution de l’entropie du code : data science vs génie logiciel (violons)

Le motif qui émerge est ce que j’appelle une inversion d’entropie.

### La surface semble complexe, la structure l’est souvent moins

Au niveau caractère, le code de data science a tendance à présenter une entropie plus élevée que le code de génie logiciel. Cela paraît intuitif. La data science fourmille de noms de colonnes variés, d’étiquettes de jeux de données, de variables ad hoc et d’identifiants spécifiques au domaine qui rendent le code bruyant et irrégulier.

Au niveau jeton, les deux domaines sont bien plus proches. Ils s’appuient en grande partie sur les mêmes briques syntaxiques.

Mais au niveau AST, où l’on observe la diversité structurelle, l’image s’inverse.

Le code de génie logiciel encode généralement bien plus de variation structurelle. Il crée des comportements distincts via des abstractions, interfaces, modules et une logique interne.

Le code de data science, à l’inverse, réutilise souvent un plus petit ensemble d’opérations dans des contextes changeants : charger, filtrer, grouper, agréger, visualiser, inspecter, ajuster.

En bref, le code de data science paraît souvent plus complexe en surface, tandis que le code de génie logiciel porte davantage de complexité dans sa structure.

Ce n’est pas qu’une différence de style. Cela révèle une divergence plus profonde dans la manière dont chaque type de travail stocke le sens.

## Indexical vs symbolique : où se loge le sens

### Deux natures de code différentes

L’inversion d’entropie s’explique mieux si l’on considère ce que fait réellement chaque type de code.

En génie logiciel, le sens est souvent compressé dans la structure. Fonctions, interfaces, modules, types et frontières des classes font l’essentiel du travail. Une fois ces abstractions établies, elles stabilisent les comportements et réduisent l’incertitude future.

Une grande part du sens est à l’intérieur même du code.

En data science, le sens reste bien plus étroitement lié au contexte externe. Il dépend du jeu de données, des colonnes, des sorties intermédiaires, des hypothèses derrière une transformation et de la question qui évolue au fil de l’analyse. Le code n’exprime pas seulement une logique ; il renvoie à une situation analytique précise.

Cette différence aide à comprendre pourquoi les agents se comportent souvent si différemment dans les deux domaines.

### On le voit à l’endroit où le code « pointe »

J’ai également mesuré les densités de références externes et internes pour 100 lignes de code dans les deux groupes.


Densités de références externes et internes avec significativité statistique

La différence est nette. Le code de data science renvoie plus souvent vers l’extérieur : jeux de données, tables, colonnes, objets temporaires et états qui existent en dehors du code lui‑même.

Le code de génie logiciel est plus auto‑référentiel. Il construit plus souvent le sens en pointant vers des fonctions, classes, modules et abstractions définis ailleurs dans la base de code.

Le code « data » pointe vers l’extérieur. Le code logiciel pointe vers l’intérieur.

Et cela compte pour les agents. Dans un cas, une grande partie du contexte pertinent est dans la base de code. Dans l’autre, il réside dans l’état analytique environnant.

## Pourquoi cela fait agir l’abstraction différemment

### Le problème de couplage dans les notebooks

L’une des observations intéressantes concerne le couplage. Les workflows de type notebook montrent un couplage nettement plus serré pour 100 lignes de code que les bases de code en génie logiciel.

Ce n’est pas nécessairement une mauvaise conception. Cela reflète quelque chose de fondamental dans l’exploration : la question elle‑même évolue encore. Vous testez des hypothèses, suivez des résultats inattendus, vérifiez des cas limites et changez de direction au fur et à mesure que vous apprenez.

Dans ce contexte, l’abstraction ne paie pas forcément comme en génie logiciel. Structurer trop tôt peut réduire les options avant même de comprendre l’essentiel. Un couplage plus serré est souvent une conséquence de l’exploration, pas simplement d’une mauvaise ingénierie.

### Le rôle des commentaires et de l’état

Il existe aussi une différence importante dans la manière dont les deux domaines s’expliquent.

En génie logiciel, le code s’explique souvent par sa structure. Les types, interfaces et abstractions portent l’essentiel du sens, et les commentaires sont généralement secondaires.

En data science, les commentaires, les sorties et l’état portent souvent une partie du sens. Une note du type « suppression des valeurs aberrantes au‑delà du 99e percentile, confirmé que cela n’affecte pas la cohorte principale » n’est pas une simple documentation.

Elle capture une décision analytique qui peut ne pas être récupérable à partir du seul code. Une table ou un graphique au milieu du notebook peut expliquer pourquoi la suite de l’analyse existe tout court.

Cela compte pour les agents. Un agent qui ignore les commentaires, les résultats en ligne et l’état évolutif d’un notebook passe à côté d’une partie de l’analyse. Dans une base de code logicielle, ces informations sont souvent périphériques. En data science, souvent non.

## Ce que cela implique pour les agents d’IA

### Par défaut, les agents sont mieux adaptés à un des deux domaines

Conséquence pratique : les agents fonctionnent au mieux lorsque leurs outils correspondent à l’endroit où se trouve le sens.

En génie logiciel, un agent efficace navigue dans l’architecture. Il suit le graphe d’appels, respecte les interfaces et maintient la cohérence des changements dans toute la base de code. Cela fonctionne parce qu’une grande partie du sens est encodée dans la structure.

