---
id: collect-261001-ia-llm/ia-llm/pourquoi-les-agents-dia-peinent-davantage-avec-le-code-data-science-quavec-le-code-genie-l-2
title: "pourquoi-les-agents-dia-peinent-davantage-avec-le-code-data-science-quavec-le-code-genie-logiciel"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents"]
source: docs/RAG/collect-261001-ia-llm/pourquoi-les-agents-dia-peinent-davantage-avec-le-code-data-science-quavec-le-code-genie-logiciel.md
source_anchor: ""
source_lines: [105, 171]
sha256: d4c81d5db16892ff343c4688324eaded6a61786866d3c8d72066e5064476e1d7
---

# pourquoi-les-agents-dia-peinent-davantage-avec-le-code-data-science-quavec-le-code-genie-logiciel

En data science, un agent efficace doit faire plus que naviguer dans le code. Il doit comprendre ce que contiennent réellement les données, suivre l’état au fil des étapes, retracer la provenance des variables et raisonner sur les raisons d’une transformation ou d’un filtre appliqué plusieurs étapes plus tôt.

Un agent conçu pour le génie logiciel ne se transpose pas automatiquement à la data science. Le problème n’est pas seulement syntaxique. La structure informationnelle sous‑jacente est différente.

### Pourquoi les workflows de type notebook persistent

Cela aide aussi à expliquer ce qui déroute souvent les ingénieurs : pourquoi les notebooks restent centraux malgré leurs limites évidentes.

La raison est que les notebooks gardent le sens au plus près des données tant que la question analytique évolue. Un notebook n’est pas simplement un module Python mal structuré. C’est un artefact différent, conçu pour une phase de travail différente. Il maintient code, sorties et décisions au même endroit tant que l’analyse se construit.

C’est pourquoi un notebook peut sembler intuitif à l’analyste immergé dans le contexte, et maladroit pour un agent qui ne voit que le code.

Refactorer un notebook en code propre et modulaire avant que la question analytique ne soit stabilisée n’est pas toujours une amélioration. Cela peut détacher la logique du contexte qui lui donnait son sens.

### L’opportunité : combler le fossé avec la production

La vraie opportunité pour les agents d’IA en data science n’est pas de copier ce qui marche en génie logiciel. C’est d’aider à combler le fossé de mise en production : la distance entre un insight découvert en exploration et un workflow reproductible et déployable.

Ce fossé existe parce que les notebooks préservent le contexte analytique au prix d’une propreté structurelle. Un agent capable de comprendre à la fois le contexte analytique et les exigences structurelles des systèmes de production pourrait combler ce fossé sans forcer l’analyste humain à quitter trop tôt le mode exploratoire.

C’est un problème plus difficile que de naviguer dans une base de code. Mais c’est aussi le plus important.

## Points clés à retenir

En résumé, ce que suggère l’analyse des dépôts étudiés :

- Le code de data science présente souvent une entropie plus élevée en surface, tandis que le code de génie logiciel présente souvent une entropie plus élevée au niveau structurel.
- Le code de data science se réfère plus souvent à un état externe (jeux de données, colonnes, tables, sorties intermédiaires), tandis que le code de génie logiciel est plus auto‑référentiel.
- Les workflows de type notebook tendent à montrer un couplage plus serré, souvent conséquence de l’exploration plutôt que d’une simple mauvaise pratique.
- Les commentaires, les sorties et l’état évolutif font partie du contenu sémantique du travail de data science, pas de simples détails périphériques.
- Les agents optimisés pour des contextes de génie logiciel sous‑performeront souvent en data science s’ils ne sont pas conçus pour raisonner sur les données, les sorties et le contexte en plus de la structure du code.

## Conclusion

Au lancement de cette analyse, je m’attendais à constater que le code de data science était simplement moins bien structuré que le code de génie logiciel. J’ai au contraire trouvé qu’il est structuré différemment pour de bonnes raisons, le sens résidant souvent ailleurs.

Cela a des conséquences concrètes pour quiconque conçoit ou évalue des agents d’IA pour des travaux analytiques. La vraie question n’est pas de savoir si un agent peut écrire du Python correct, mais s’il comprend ce que sont les données, pourquoi l’analyse prend telle forme et quelles décisions ont été prises en chemin et qui ne sont pas visibles dans le seul code.

Le code de data science n’est pas un génie logiciel immature. C’est une optimisation d’une autre nature : plus dépendante de l’état externe, moins tributaire d’une structure interne durable pendant l’exploration, et façonnée autant par le contexte que par le code.

Une fois que l’on voit cela, les notebooks cessent de ressembler à des projets logiciels ratés et apparaissent pour ce qu’ils sont : des surfaces de travail pour raisonner sur des données changeantes.

**Je poursuis cette analyse avec un corpus élargi d’au moins 1 000 dépôts et des métriques supplémentaires. Si vous souhaitez approfondir la manière dont les agents d’IA sont conçus pour travailler avec les données dans leur contexte (y compris des architectures d’exécution persistantes qui maintiennent l’état des données entre les sessions), le parcours AI agent fundamentals de DataCamp est un excellent point de départ.**

## FAQs

### Qu’est-ce que l’entropie du code et pourquoi est-ce important pour les agents d’IA ?

**L’entropie de Shannon appliquée au code est une façon de mesurer la variation ou l’imprévisibilité à un niveau d’abstraction donné. Une entropie plus élevée au niveau structurel peut indiquer que le code exprime une gamme plus large de comportements internes. Pour les agents d’IA, ces schémas comptent car ils déterminent le type de contexte que l’agent doit comprendre.**

### Quelle est la différence entre code indexical et code symbolique ?

**Cette distinction est un raccourci utile. Le code de génie logiciel porte souvent davantage son sens en interne via des fonctions, interfaces et abstractions. Le code de data science dépend plus fortement du contexte externe : jeux de données, colonnes, sorties et état évolutif de l’analyse.**

### Cela signifie‑t‑il que le code de data science est de moindre qualité que le code de génie logiciel ?

**Non. Il ne s’agit pas de dire que l’un est meilleur que l’autre. Ils sont optimisés pour des objectifs différents. L’exploration en data science privilégie souvent la rapidité d’itération et la préservation du contexte, tandis que le génie logiciel privilégie généralement la structure, la réutilisation et la maintenabilité.**

### Pourquoi les agents de codage IA semblent‑ils souvent moins efficaces dans les notebooks de data science ?

**Parce que nombre d’agents actuels sont optimisés pour naviguer dans la structure du code. En data science, une large part du sens se situe en dehors du code lui‑même : état des données, sorties, commentaires et décisions analytiques. Si un agent ne sait pas raisonner sur ce contexte, il passe à côté d’une partie de la tâche.**

### À quoi ressemblerait un agent d’IA natif pour la data science ?

**Il devrait maintenir une conscience de l’état des données, des sorties et de la provenance des variables, pas seulement du code source. Il devrait traiter les commentaires et résultats intermédiaires comme un contexte significatif et aider à faire passer le travail exploratoire en workflows reproductibles de production.**

Jason Hillary a cofondé Zerve après avoir constaté à quel point les frictions ralentissaient le travail autour des données. Titulaire d'un doctorat en ingénierie de l'University of Limerick, il a passé des années à concevoir des systèmes d'IA réellement opérationnels en production, et pas seulement en théorie. Avant de créer Zerve, Jason a travaillé dans tout l'écosystème data et IA, avec pour objectif de rendre les tâches techniques moins pénibles et plus productives.
