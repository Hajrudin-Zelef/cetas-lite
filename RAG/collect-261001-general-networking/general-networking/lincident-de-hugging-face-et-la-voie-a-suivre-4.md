---
id: collect-261001-general-networking/general-networking/lincident-de-hugging-face-et-la-voie-a-suivre-4
title: "lincident-de-hugging-face-et-la-voie-a-suivre"
domain: general-networking
role: reference
task: reference
actors: ["ExploitGym", "Hugging Face", "OpenAI"]
dates: []
keywords: ["incident", "agents", "arr", "astra", "sol", "valuation"]
source: docs/RAG/collect-261001-general-networking/lincident-de-hugging-face-et-la-voie-a-suivre.md
source_anchor: ""
source_lines: [127, 168]
sha256: 05ce3fe05032d25e98a2c8f27e3e2b5af7f1150bcb8a32be6bef84013aba0a4c
---

# lincident-de-hugging-face-et-la-voie-a-suivre

Couverture des mesures de protection dans les évaluations internes

Outre le désalignement, un autre facteur ayant contribué à cet incident est que nous n’avions pas étendu à toutes les évaluations internes les solides mesures de protection que nous déployons pour nos modèles accessibles au public.

Premièrement, de nombreux mécanismes réduisent le désalignement dans les environnements de production de nos clients, notamment les prompts système, les harnais et les mécanismes de contrôle tels que nos modèles d’auto-évaluation et nos classificateurs de sécurité. Ces protections n’étaient pas appliquées dans l’environnement d’évaluation utilisé lors de l’incident. Après l’incident, nous avons créé une évaluation qui mesure la propension d’un modèle à compromettre l’infrastructure lors des évaluations ExploitGym. Nous avons constaté que cette propension pouvait diminuer de plus de 100 fois avec le harnais Codex de production et le prompt système. Nous avons également exécuté le système d’auto-évaluation de Codex sur un ensemble d’exécutions liées à l’incident de Hugging Face et constaté qu’il aurait signalé un grand nombre d’actions dangereuses des modèles.

Deuxièmement, avant l’incident, nous avions considérablement investi dans la surveillance de la chaîne de pensée, notamment en surveillant un grand nombre de nos exécutions d’entraînement RL de pointe ainsi qu’une nette majorité de l’utilisation interne des agents de codage. Ces dispositifs de surveillance n’étaient pas utilisés lors des évaluations liées à cet incident.

Lors d’évaluations rétrospectives, nous avons constaté que nos dispositifs de surveillance du CoT étaient très efficaces pour signaler les activités liées à l’incident. Si notre système de surveillance de la chaîne de pensée actuellement déployé avait été en fonctionnement au moment de l’incident, il aurait probablement alerté notre équipe de sécurité dans l’heure suivant le premier accès non autorisé par un tiers. Cela aurait eu lieu plus de 30 heures avant que les modèles ne compromettent les systèmes de Hugging Face. Nous travaillons également à réduire davantage la latence de ces alertes.

Troisièmement, bien que nous ayons testé et validé ce bac à sable, les agents ont pu enchaîner des vulnérabilités jusque-là inconnues (« 0-days ») dans le service de gestion des paquets exposé au sein du bac à sable afin de contourner les restrictions, comme le décrit le rapport d’incident technique. Depuis l’incident, nous avons réalisé des progrès importants dans le renforcement de nos bacs à sable internes afin de limiter les risques similaires lors des déploiements et des évaluations.

Perspectives

La mission d’OpenAI est de faire en sorte que l’intelligence artificielle générale profite à toute l’humanité. Le comportement de nos modèles décrit ici était très éloigné du niveau que nous souhaitons atteindre. Cet incident n’aurait pas dû se produire. Il a montré à quel point il est essentiel d’améliorer continuellement notre sécurité, notre surveillance et notre alignement, en particulier à mesure que nos modèles atteignent un niveau de capacités susceptible d’entraîner une véritable perte de contrôle.

Comme nous l’avons indiqué dans notre récente mise à jour sur le rythme de développement des modèles, nous avons suspendu l’entraînement par apprentissage par renforcement (RL) de nos derniers modèles destinés au déploiement. Nous avons mis ce temps à profit pour renforcer davantage la sécurité de nos environnements de recherche de pointe et les soumettre à des tests d’intrusion. Nous avons réaffecté du personnel à des travaux sur la sécurité, la sûreté et l’alignement. Notre plus grande exécution d’entraînement RL de pointe prévue reste suspendue, tandis que nous menons des entraînements et des évaluations à plus petite échelle afin d’évaluer le comportement des modèles, de valider nos mesures de protection et de réunir davantage de preuves d’alignement avant de poursuivre. Ce travail a nécessité d’importants efforts d’ingénierie, entraînant des coûts élevés et des retards pour la recherche de pointe.

Comme indiqué dans notre rapport, notre réponse vise à renforcer la sécurité et la surveillance de la recherche, à accélérer l’alignement des modèles et à améliorer notre manière d’identifier, de faire remonter et de traiter les incidents.

Sécurité et surveillance

Nous avons mis en œuvre des exigences de sécurité plus strictes pour les charges de travail de recherche de pointe. Par exemple :

Isolement des charges de travail : nous mettons en place des bacs à sable plus isolés pour les charges de travail qui exécutent du code généré par un modèle ou autrement non fiable, y compris les logiciels susceptibles d’être compromis lors du traitement des sorties d’un modèle.

Isolement réseau : nous avons mis en place davantage de mesures de contrôle afin d’isoler d’Internet les charges de travail à risque élevé et non fiables. Ces contrôles sont conçus pour garantir que la compromission d’une seule charge de travail ou d’un service sous-jacent ne permette pas, à elle seule, un accès non autorisé à Internet ou à d’autres réseaux internes.

Tests de sécurité en continu : nous améliorons notre capacité à collecter et à surveiller les journaux de sécurité et investissons dans l’automatisation à l’aide de nos modèles afin de tester en continu les limites d’isolation des charges de travail et du réseau au moyen d’attaques simulées.

Nous exigeons désormais également une surveillance de la chaîne de pensée pour tous les entraînements et évaluations par apprentissage par renforcement (RL) utilisant des outils et impliquant des modèles dotés de capacités de niveau GPT‑5.6 Sol ou supérieur. Pour nos modèles de classe Astra, qui pourraient disposer de capacités critiques en cybersécurité, cette exigence couvre également toutes les charges de travail d’inférence utilisant des outils.

Certaines charges de travail de recherche reprennent progressivement dans le cadre de ces nouvelles mesures de contrôle, tandis que d’autres nécessitent des mises à niveau supplémentaires.

Accélérer l’alignement

Nous avons également accéléré les travaux que nous menions déjà sur l’entraînement à l’alignement dans l’ensemble du pipeline de développement des modèles. L’incident Hugging Face nous a amenés à nous concentrer davantage sur les comportements spécifiques qui y ont contribué, notamment :

Triche, environnements défaillants et arrêt sécurisé. Lorsqu’une tâche est corrompue, défectueuse ou impossible à accomplir, les agents doivent demander des clarifications ou s’arrêter en toute sécurité, plutôt que de poursuivre des solutions de rechange de plus en plus douteuses. Nous étendons la couverture des systèmes d’évaluation, qui déterminent non seulement si une tâche a été accomplie, mais aussi comment. Ces systèmes d’évaluation récompensent les modèles qui identifient correctement les tâches défectueuses, demandent des précisions ou s’arrêtent en toute sécurité lorsque cela est nécessaire.

Alignement multi-agents. L’incident a révélé que nos agents ne font pas preuve d’un discernement suffisant quant aux personnes avec lesquelles ils collaborent. Nous développons activement des environnements d’apprentissage par renforcement (RL) qui apprennent à nos modèles à se méfier des instructions non autorisées.

