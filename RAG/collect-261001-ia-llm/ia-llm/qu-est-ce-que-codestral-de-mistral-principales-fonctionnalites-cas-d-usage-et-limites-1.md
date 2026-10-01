---
id: collect-261001-ia-llm/ia-llm/qu-est-ce-que-codestral-de-mistral-principales-fonctionnalites-cas-d-usage-et-limites-1
title: "qu-est-ce-que-codestral-de-mistral-principales-fonctionnalites-cas-d-usage-et-limites"
domain: ia-llm
role: reference
task: reference
actors: ["DeepSeek", "Mistral"]
dates: []
keywords: ["mistral", "benchmark", "benchmarks", "deepseek", "exploit", "llama", "open-weight"]
source: docs/RAG/collect-261001-ia-llm/qu-est-ce-que-codestral-de-mistral-principales-fonctionnalites-cas-d-usage-et-limites.md
source_anchor: ""
source_lines: [1, 108]
sha256: d7236434876f16ed377c0575399634acd82ddc0149d79e11320c5fa395ee6e32
---

# qu-est-ce-que-codestral-de-mistral-principales-fonctionnalites-cas-d-usage-et-limites

Cours

La génération de code prend une place grandissante dans le développement logiciel moderne. Mistral AI a récemment investi ce domaine avec **Codestral**, son tout premier modèle spécialisé pour le code.

Dans cet article, nous vous proposons un tour d'horizon complet de Codestral, en explorant ses fonctionnalités et son fonctionnement.

Je partagerai également des retours issus de mon expérience pratique avec le modèle, avec des exemples concrets de génération de code.

Pour en savoir plus sur Mistral, consultez ce guide complet pour travailler avec le modèle Mistral Large.

## Qu'est-ce que Codestral ?

Codestral est un modèle d'IA générative open-weight, conçu spécifiquement pour les tâches de génération de code — « open-weight » signifie que les paramètres appris du modèle sont librement accessibles à des fins de recherche et d'usage non commercial, offrant plus d'accessibilité et de possibilités de personnalisation.

Codestral propose aux développeurs une approche flexible pour écrire et interagir avec le code via un point de terminaison API commun pour l'instruction et la complétion. Concrètement, on peut fournir à Codestral des consignes en langage naturel ou des extraits de code, et il génère le code correspondant.

La capacité unique de Codestral à comprendre à la fois le code et le langage naturel en fait un outil polyvalent pour des tâches comme la complétion de code, la génération à partir de descriptions en langage courant, ou encore les questions-réponses sur des snippets. Cela ouvre la voie à de nombreux outils dopés à l'IA capables de fluidifier nos flux de développement.

## Fonctionnalités clés de Codestral

Codestral offre plusieurs fonctionnalités marquantes qui renforcent son intérêt pour la génération de code. Passons-les en revue.

### Maîtrise de plus de 80 langages de programmation

L'une des capacités les plus impressionnantes de Codestral est sa maîtrise de plus de 80 langages de programmation. Cette couverture inclut non seulement les langages populaires comme Python, Java, C, C++ et JavaScript, mais aussi des langages plus spécialisés utilisés dans certains domaines ou pour des besoins de niche (comme Swift ou Fortran).

Cette polyvalence en fait un atout pour les projets multi-langages ou les équipes où les développeurs travaillent avec des écosystèmes différents. Qu'il s'agisse d'un projet data science en Python, d'une application web en JavaScript ou d'un développement système en C++, Codestral s'adapte et apporte une aide à la génération de code sur un large éventail de langages.

### Génération de code

La fonction cœur de Codestral est la génération de code. Il vise à rationaliser notre travail en automatisant des tâches telles que la complétion de fonctions, la génération de cas de test ou le remplissage de segments manquants.

Le mécanisme de complétion « au milieu » est conçu pour aider sur des bases de code complexes ou dans des langages peu familiers. Bien exploitées, ces fonctionnalités peuvent libérer du temps pour la conception et la réflexion de plus haut niveau, accélérant les cycles de développement et améliorant la fiabilité du code.

### Open-weight

Un aspect notable de Codestral est sa nature open-weight. Les paramètres appris du modèle sont librement accessibles pour la recherche et l'usage non commercial.

Cette ouverture favorise la collaboration : développeurs et chercheurs peuvent expérimenter, l'affiner pour des tâches spécifiques et contribuer à son évolution.

Elle démocratise l'accès à de puissantes capacités de génération de code et encourage transparence et innovation au sein de la communauté IA.

### Performance et efficacité

Mistral AI affirme que Codestral fixe un nouveau standard en matière de performance et de latence pour la génération de code, surpassant d'autres modèles sur certains benchmarks. Sa large fenêtre de contexte (32 000 tokens) renforcerait sa capacité à gérer des tâches de complétion longue portée.

Abordons plus en détail la performance et l'efficacité dans la section suivante.

## Comparaison de Codestral avec d'autres modèles

Pour mieux appréhender les capacités de Codestral, comparons ses résultats à ceux d'autres modèles de référence en génération de code. Les sections suivantes présentent des benchmarks spécifiques et mettent en lumière les principales différences.

### Fenêtre de contexte

Commençons par ces résultats :

Codestral se distingue par ses performances sur les tâches de complétion longue portée (RepoBench), probablement grâce à sa fenêtre de contexte étendue à 32 000 tokens. Cette fenêtre plus large lui permet de prendre en compte davantage de code environnant, améliorant la prédiction. Codestral excelle aussi sur le benchmark HumanEval en Python, démontrant sa capacité à produire un code précis.

Si Codestral brille sur certains axes, d'autres modèles comme DeepSeek Coder performent mieux sur d'autres benchmarks (MBPP).

Bien que plus compact que nombre de LLM concurrents, Codestral affiche des performances souvent supérieures ou au moins comparables à des modèles bien plus grands comme Llama 3 70B, et ce, sur l'ensemble des langages en génération et en complétion « au milieu ».

### Performance en complétion au milieu

Regardons maintenant la performance en fill-in-the-middle (FIM) :

Codestral 22B obtient des résultats nettement supérieurs sur les trois langages (Python, JavaScript et Java) et sur la moyenne FIM globale par rapport à DeepSeek Coder 33B. Cela suggère une bonne compréhension du contexte et une grande précision pour combler les segments manquants.

Cependant, ce benchmark ne compare que Codestral à DeepSeek Coder 33B et n'intègre pas CodeLlama 70B ni Llama 3 70B, ce qui limite la portée des conclusions.

### HumanEval

Le benchmark HumanEval évalue la précision de génération en testant la capacité des modèles à produire un code qui réussit des tests unitaires rédigés par des humains à partir de descriptions de fonctions. Voyons comment Codestral se positionne face à d'autres modèles sur HumanEval :

Codestral affiche les meilleures performances en Python, bash, Java et PHP. Si d'autres modèles dominent sur certains langages, la moyenne générale de Codestral est en tête, démontrant une forte capacité à générer un code juste dans plusieurs langages.

## Cas d'usage de Codestral

La diversité des capacités de Codestral se prête à de nombreuses applications concrètes tout au long du cycle de vie logiciel. Voici quelques cas d'usage où Codestral peut apporter un impact notable.

### Complétion et génération de code

Codestral excelle en complétion et génération de code, son cas d'usage principal. Les développeurs peuvent s'appuyer sur Codestral pour proposer des complétions en fonction du contexte, accélérant l'écriture et réduisant les erreurs.

Il peut également générer des extraits entiers à partir de descriptions ou d'instructions en langage naturel, fluidifiant davantage le développement et améliorant la productivité.

Voici un exemple rapide de ce que j'ai demandé à Codestral de générer :

```
prompt = "Please write me a function that adds up two numbers"
data = {
    "model": "codestral-latest",
    "messages": [
        {
            "role": "user",
            "content": prompt
        }
    ],
    "temperature": 0
}
response = call_chat_instruct_endpoint(api_key, data)
```
### Génération de tests unitaires

Codestral facilite aussi la génération de tests unitaires pour du code existant. Cette automatisation fait gagner un temps précieux et contribue à améliorer la qualité du code tout en réduisant le risque de bugs, pour des projets plus robustes et maintenables.

J'ai demandé à Codestral de générer un test unitaire simple pour la fonction précédente :

