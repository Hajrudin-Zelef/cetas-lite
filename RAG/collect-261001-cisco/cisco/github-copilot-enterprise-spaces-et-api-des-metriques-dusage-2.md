---
id: collect-261001-cisco/cisco/github-copilot-enterprise-spaces-et-api-des-metriques-dusage-2
title: "github-copilot-enterprise-spaces-et-api-des-metriques-dusage"
domain: cisco
role: reference
task: reference
actors: ["Microsoft"]
dates: ["2026-03-10"]
keywords: ["copilot", "agents", "mcp"]
source: docs/RAG/collect-261001-cisco/github-copilot-enterprise-spaces-et-api-des-metriques-dusage.md
source_anchor: ""
source_lines: [131, 297]
sha256: c8dd20844cc1b291d61bd0c085b601337402806eab76f0aaeb95b4477759278a
---

# github-copilot-enterprise-spaces-et-api-des-metriques-dusage

La structure initiale a tendance à perdurer. J’ai vu des organisations créer par inadvertance une « jungle documentaire » dans Spaces faute de règles de responsabilité claires dès le départ.

Tout le monde peut créer un Copilot Space, testons donc dans un dépôt personnel. Les étapes sont similaires au niveau Enterprise, avec quelques pages différentes.

### Configurer un Space

La création suit généralement ce workflow :

1. Accéder à la page Copilot Spaces dans la zone d’administration Enterprise
2. Créer un nouveau Space

1. Sélectionner les dépôts et sources de contenu, y compris les MCP et autres outils utiles

1. Ajouter des sources via le bouton « + Add sources » à droite

1. Choisir de partager le Space ou définir les paramètres de partage à ce stade

1. Vérifier que Copilot peut référencer le contenu pendant les conversations

Note pour les utilisateurs Enterprise : votre administrateur peut désactiver le partage des Spaces personnels. Si vous utilisez votre propre compte, cela peut limiter le partage d’un Copilot Space qui n’utilise pas les dépôts de l’entreprise.

Après la configuration, les administrateurs doivent tester le Space avec des invites concrètes.

Par exemple :

`How does our authentication middleware handle token refresh logic?`
Ou :

`Show me an example of how our backend services structure database migrations.`
Si Copilot ne peut pas répondre avec précision, la cause est généralement :

- Dépôts manquants
- Documentation de faible qualité
- Autorisations incorrectes
- Temps d’indexation insuffisant

### Partage et contrôles d’accès

Spaces prend en charge deux grands modèles de visibilité :

- Spaces individuels
- Spaces à l’échelle de l’organisation

Les membres d’une entreprise peuvent voir leurs espaces individuels gérés par les paramètres globaux de l’entreprise. Les administrateurs Enterprise peuvent aussi gérer centralement les politiques d’accès anticipé et de disponibilité des fonctionnalités.

Les Spaces privés conviennent aux équipes en expérimentation ou aux initiatives sensibles. Les Spaces organisationnels sont idéaux pour les standards d’ingénierie, l’onboarding ou les frameworks communs.

Une erreur fréquente est la sur-centralisation. Un unique Space tentaculaire et global devient vite bruyant et moins utile pour Copilot.

### Organiser les Spaces par équipe ou par domaine

Il n’existe pas une seule structure universelle.

Les schémas courants incluent un Space par équipe, un Space par produit, ou des Spaces de standards partagés. Chacun a une portée différente et exploite les mêmes réglages de façon spécifique.

#### Un Space par équipe

Utile lorsque les groupes d’ingénierie opèrent de façon relativement indépendante.

Exemples :

- Ingénierie plateforme
- Ingénierie des données
- Développement mobile

#### Un Space par produit

Utile pour les organisations structurées par produits plutôt que par départements.

Exemples :

- Paiements
- Analytique
- Infrastructure
- Plateforme client

#### Un Space de standards partagés

Beaucoup d’organisations maintiennent un Space partagé séparé pour :

- Lignes directrices de sécurité
- Conventions de code
- Workflows de déploiement
- Standards d’architecture

En pratique, les approches hybrides donnent généralement les meilleurs résultats : chaque équipe dispose de son Space, complété par de grands Spaces de standards partagés.

## L’API des métriques d’usage de Copilot

Spaces résout le problème du contexte. L’API des métriques d’usage résout le problème de la mesure. Elle a remplacé plusieurs anciens systèmes de télémétrie que GitHub a retirés lors de la consolidation des API en 2026.

Sans mesures claires, les organisations perdent rapidement de vue le succès de l’adoption de Copilot. Les directions veulent des preuves que l’investissement améliore les workflows des développeurs, et pas seulement une ligne d’abonnement supplémentaire.

Le tableau de bord est en disponibilité générale depuis février 2026 et accessible via votre compte enterprise → AI Controls → Copilot → Metrics → Copilot usage metrics dans l’onglet Insights.

### Ce que mesure l’API

L’API des métriques d’usage expose plusieurs catégories de télémétrie opérationnelle.

Parmi les métriques courantes :

- Utilisateurs actifs
- Lignes de code proposées vs lignes de code acceptées
- Profils d’usage des IDE
- Usage des modèles
- Interactions avec les agents
- Répartition par langage

Cela offre une vision bien plus fine que de simples comptages de licences.

Une équipe avec 100 licences attribuées mais seulement 15 utilisateurs actifs n’a pas le même profil d’adoption qu’une équipe avec un usage quotidien régulier et des taux d’acceptation élevés.

### La transition d’API en 2026

GitHub a retiré plusieurs anciennes API de télémétrie (User-level Feature Engagement Metrics API, Direct Data Access API, Copilot Metrics API) entre 2025 et 2026, pour une extinction complète en avril 2026.

Elles comprenaient :

- L’ancienne Metrics API
- Les Feature Engagement APIs
- Les Direct Data Access APIs

Les nouveaux endpoints de métriques d’usage, disponibles depuis février 2026, ont unifié ces systèmes de reporting dans un modèle plus cohérent, avec versionnage en cas de changements incompatibles.

C’est important car beaucoup d’anciens billets et exemples GitHub référencent encore des endpoints dépréciés. Avant d’industrialiser vos intégrations, vérifiez systématiquement la documentation la plus récente.

## Interroger l’API des métriques d’usage

Maintenant que le but de l’API des métriques est clair, voyons comment l’utiliser concrètement.

### Authentification et autorisations

Les endpoints des métriques d’usage GitHub Copilot nécessitent généralement quelques autorisations sur votre Personal Access Token (PAT), en PAT classique ou finement granulé.

- 
Pour les PAT classiques, votre admin enterprise doit vous attribuer les permissions `manage_billing:copilot` et`read:org` .
- 
Pour les tokens finement granulés, utilisez un GitHub App user access token ou installation access token avec la permission `Enterprise Copilot metrics enterprise permissions (read)` .

En général, les tokens finement granulés sont à privilégier car ils limitent l’exposition inutile des permissions.

### Endpoints au niveau organisation

Les deux rapports les plus courants au niveau organisation sont :

- 
`organization-1-day`
- 
`organization-28-day`

#### Rapport organisationnel sur une journée

Le rapport sur une journée est idéal pour le suivi opérationnel et l’analyse de tendances à court terme. L’historique remonte au 10 octobre 2025 et reste accessible pendant un an à partir de la date courante.

La commande curl ci-dessous appelle l’API de rapport 1 jour et renvoie une réponse JSON avec des liens de téléchargement. Renseignez `YOUR_TOKEN` pour le Bearer et choisissez un `DAY` au format `YYYY-MM-DD`.

```
curl -L \
 -H "Accept: application/vnd.github+json" \
 -H "Authorization: Bearer <YOUR_TOKEN>" \
-H “X-GitHub-Api-Version: 2026-03-10” \
"https://api.github.com/enterprises/ENTERPRISE/copilot/metrics/reports/enterprise-1-day?day=DAY"
```
Les URLs dans `download_links` sont signées et à durée limitée : elles expirent rapidement après leur génération. Votre workflow doit récupérer l’URL puis télécharger le fichier immédiatement, dans la même exécution.

La réponse peut ne contenir que `download_links` et `report_day`, mais voici le schéma complet possible :

