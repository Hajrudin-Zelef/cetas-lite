---
id: collect-261001-ia-llm/ia-llm/mistral-small-3-vs-llama-4-scout-ia-locale-48-go-2026-1
title: "1. Installer Ollama (Linux/macOS)"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Meta", "Microsoft", "Mistral", "OpenAI"]
dates: []
keywords: ["llama", "apache", "benchmark", "benchmarks", "chatgpt", "claude", "mistral", "moe", "open source", "qwen", "scout"]
source: docs/RAG/collect-261001-ia-llm/mistral-small-3-vs-llama-4-scout-ia-locale-48-go-2026.md
source_anchor: ""
source_lines: [1, 40]
sha256: 8d1511cb7c49d2e011fb64d5d95396a0b85da85793e4d19a3ddc9144935f9ab4
---

# 1. Installer Ollama (Linux/macOS)

Un cabinet d’avocats parisien qui refuse d’envoyer ses dossiers clients vers un serveur situé aux États-Unis. Un service hospitalier qui veut résumer des comptes rendus médicaux sans exposer la moindre donnée de santé. Un développeur qui code dans le train, sans connexion internet fiable. Ces trois profils ont un point commun en 2026 : ils se tournent vers l’IA locale plutôt que vers ChatGPT ou Claude hébergés dans le cloud américain.

Trois modèles dominent désormais ce terrain, avec des philosophies et des besoins matériels très différents : **Mistral Small 3**, le pari français porté par une startup parisienne, **Llama 4 Scout**, la version allégée du géant Mixture-of-Experts de Meta, et **Qwen 3**, le polyglotte open source d’Alibaba. Voici ce que chacun vaut réellement, avec les chiffres à l’appui, la configuration PC nécessaire, et un guide pour migrer du cloud vers l’IA locale sans perdre en productivité.

## L’IA locale, la bascule qui redéfinit l’usage de l’IA en France en 2026

Pendant deux ans, choisir une IA a surtout voulu dire choisir un abonnement : ChatGPT Plus, Claude Pro, ou Le Chat Pro. En avril 2026, ce calcul a changé. L’IA locale, celle qui tourne directement sur un PC ou un serveur d’entreprise sans passer par une API externe, est devenue une option crédible et performante pour la majorité des usages professionnels, selon plusieurs cabinets d’analyse spécialisés en IA française. La raison n’est pas seulement technique. Elle est aussi réglementaire.

Le Règlement général sur la protection des données encadre strictement l’export de données personnelles hors de l’Union européenne. Un cabinet médical, un avocat ou une administration qui envoie des dossiers sensibles vers un modèle hébergé aux États-Unis prend un risque de conformité que beaucoup préfèrent éviter. Faire tourner le modèle en interne, sur du matériel que l’entreprise contrôle, supprime ce problème à la racine. C’est tout l’enjeu de l’IA open source déployée en local : garder la donnée chez soi.

L’État français a mis de l’argent sur la table pour accélérer ce mouvement. Le programme France 2030 a alloué 1,5 milliard d’euros à l’infrastructure de recherche, à la formation et à l’adoption de l’IA en entreprise sur la période 2025-2026. Mistral AI, seul éditeur français de modèles frontières, a vu sa base d’utilisateurs grimper de 300 % en 2025, portée par une demande d’entreprise pour une IA souveraine et optimisée pour le français.

Mais la souveraineté ne suffit pas si le modèle est mauvais. C’est là que Mistral Small 3, Llama 4 Scout et Qwen 3 entrent en jeu : trois approches différentes du même problème, faire tourner une IA compétente sans dépendre d’un centre de données étranger.

## Mistral Small 3 : le pari français pensé pour tourner en local

Mistral Small 3 est la réponse directe de la startup parisienne à une demande précise : un modèle assez petit pour tourner sur du matériel raisonnable, mais assez solide pour remplacer un abonnement cloud sur la majorité des tâches du quotidien.

Le modèle compte 24 milliards de paramètres, une taille pensée pour l’inférence locale plutôt que pour battre des records absolus de performance. Mistral affirme, benchmarks à l’appui, que Small 3 atteint un niveau de performance comparable à Llama 3.3 70B, un modèle presque trois fois plus gros, tout en offrant une vitesse d’inférence nettement supérieure sur le même matériel. Concrètement, cela veut dire des réponses plus rapides sur un PC ou un Mac équipé d’une carte graphique grand public, là où un modèle de 70 milliards de paramètres exigerait un poste bien plus musclé.

Mistral Small 3 est distribué sous licence Apache 2.0, la licence open source la plus permissive du marché. Une entreprise peut le télécharger, le modifier, l’intégrer dans un produit commercial et le redistribuer sans payer de redevance ni demander d’autorisation. C’est un choix stratégique assumé par Mistral AI, qui mise sur l’adoption large plutôt que sur la licence fermée pour construire son écosystème d’IA open source.

Le positionnement produit est clair : Mistral Small 3 vise les PME, les cabinets professionnels et les développeurs qui veulent un modèle capable de tourner sur un poste de travail standard, sans salle serveur dédiée. Pour ceux qui préfèrent malgré tout une IA en ligne, Mistral propose aussi Le Chat, sa version cloud gratuite sans limite de messages annoncée, ainsi qu’une offre Pro à 20 dollars (18 euros) par mois lancée en février 2026, qui dépasse GPT-4o sur le benchmark MMLU avec un score de 88 % contre 86 %.

Sur le papier, Mistral Small 3 n’est pas le modèle le plus puissant de ce comparatif. Il n’a ni la taille de Llama 4 Scout, ni l’éventail de tailles de Qwen 3. Mais c’est le seul des trois pensé, dès le départ, pour l’équilibre entre performance et accessibilité matérielle, avec en prime un ancrage français qui pèse lourd pour les acheteurs publics et les secteurs régulés.

## Llama 4 Scout : la puissance Mixture-of-Experts de Meta, au prix de la RAM

Llama 4 Scout part d’une philosophie opposée. Plutôt que de viser la légèreté, Meta a construit un modèle Mixture-of-Experts (MoE) de 109 milliards de paramètres au total, dont seulement 17 milliards sont activés à chaque inférence.

Le principe du Mixture-of-Experts change la donne par rapport à un modèle dense comme Mistral Small 3. Au lieu de faire transiter chaque requête par l’intégralité du réseau de neurones, l’architecture MoE ne réveille qu’un sous-ensemble de modules internes spécialisés, sélectionnés selon la nature de la requête. Le résultat : une capacité de connaissance proche d’un modèle à 109 milliards de paramètres, avec un coût de calcul par requête plus proche de celui d’un modèle à 17 milliards. En théorie, c’est le meilleur des deux mondes.

En pratique, ce compromis a un prix : la mémoire. Même si seuls 17 milliards de paramètres s’activent à chaque calcul, l’intégralité des 109 milliards doit rester chargée en mémoire pour que le routage vers les bons modules fonctionne correctement. Résultat, Llama 4 Scout exige un minimum de 48 Go de RAM pour fonctionner dans de bonnes conditions, ce qui le classe d’emblée dans la catégorie « production avancée » plutôt que dans celle des modèles utilisables sur un simple ordinateur portable grand public.

Ce seuil de 48 Go change complètement le public cible. Un particulier avec un ordinateur portable standard doté de 16 ou 32 Go de RAM ne pourra pas faire tourner Llama 4 Scout dans de bonnes conditions. Une PME qui investit dans un poste de travail dédié à l’IA, ou une entreprise qui déploie le modèle sur un serveur interne, y trouvera en revanche un modèle capable de rivaliser avec des solutions cloud bien plus chères à l’usage.

Meta a conçu la famille Llama 4 comme nativement multimodale, capable de traiter texte et image dans un même flux de raisonnement, un atout que ni Mistral Small 3 ni les petites variantes de Qwen 3 ne revendiquent au même niveau. La contrepartie de cette ambition, c’est la facture matérielle. Llama 4 Scout est le modèle le plus exigeant des trois comparés ici, et le seul qui pousse clairement vers un investissement en RAM au-delà du PC grand public.

## Qwen 3 : le polyglotte open source d’Alibaba

