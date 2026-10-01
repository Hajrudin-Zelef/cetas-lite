---
id: collect-261001-rattrapage/rattrapage/formules-github-copilot-guide-des-fonctionnalites-et-de-l-administration-1
title: "Ignore the /src/some-dir/kernel.rs file in this repository."
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["attribution", "copilot"]
source: docs/RAG/collect-261001-rattrapage/formules-github-copilot-guide-des-fonctionnalites-et-de-l-administration.md
source_anchor: ""
source_lines: [1, 73]
sha256: 3984020a2a69c58c9b4222d0792ddf5f40d6e5b785f4b4473d1c684122304602
---

# Ignore the /src/some-dir/kernel.rs file in this repository.

Cursus

Votre équipe vient d’obtenir un budget pour déployer GitHub Copilot à l’échelle de l’organisation d’ingénierie. Pour en tirer le meilleur parti, il faut comprendre comment s’articulent les paramètres de politique, les exclusions de fichiers et les requêtes de journaux d’audit, car c’est là que se trouve la vraie valeur de la plateforme.

La surface de configuration est vaste parce que les besoins le sont tout autant. Un développeur solo sur des projets personnels ne fait pas face aux mêmes enjeux de confidentialité et de conformité qu’un administrateur d’entreprise qui gère des milliers de sièges sur des dépôts réglementés. La structure de formules par paliers de GitHub Copilot est conçue pour couvrir cet éventail.

Ce guide présente chaque niveau de formule Copilot, les frontières de confidentialité et de propriété intellectuelle (PI) qui les distinguent, ainsi que les mécanismes d’administration nécessaires pour passer à l’échelle dans une organisation.

Avant d’entrer dans l’administration, vous devriez déjà être à l’aise avec les organisations, les dépôts et le système d’autorisations de GitHub. Si vous découvrez l’écosystème, commencez par notre guide Comment utiliser GitHub Copilot.

Si vous hésitez encore entre Copilot et le reste du marché, notre sélection des 13 meilleurs assistants de programmation IA en 2026 couvre l’ensemble du paysage concurrentiel. Pour une comparaison ciblée avec l’un des principaux concurrents, consultez notre guide Cursor vs. GitHub Copilot.

## En bref

- GitHub propose quatre niveaux individuels (Free, Student, Pro, Pro+) et deux niveaux organisationnels (Business et Enterprise) pour Copilot, chacun avec des périmètres distincts de confidentialité, de gouvernance et d’usage.
- Les formules Business et Enterprise apportent des garanties contractuelles que les données d’interaction ne sont jamais utilisées pour l’entraînement, tandis que les formules individuelles sont en opt-out par défaut depuis avril 2026.
- Choisissez d’abord votre formule GitHub Copilot selon vos exigences de conformité et de gouvernance ; optimisez ensuite le choix des modèles et les quotas d’usage.
- Les règles d’exclusion de fichiers et les paramètres de politique à l’échelle de l’organisation ne sont disponibles que sur Business et Enterprise, ce qui en fait la base pour les équipes qui traitent du code propriétaire.
- GitHub Copilot Enterprise nécessite un abonnement actif à GitHub Enterprise Cloud, ce qui porte le coût minimal réel à 60 $ par utilisateur et par mois.
- La gestion des sièges, les requêtes de journaux d’audit et l’application des politiques peuvent être automatisées via l’API REST, transformant les licences en infrastructure-as-code.

## Aperçu des formules GitHub Copilot

GitHub propose plusieurs paliers pour son écosystème. Notamment, la plateforme finalise sa transition vers une facturation à l’usage, remplaçant l’ancien cadre « Premium Request Unit » (PRU) par les GitHub AI Credits en juin 2026.

Dans le nouveau système, les complétions de code de base et les suggestions « Next Edit » restent illimitées et ne consomment pas de crédits.

En revanche, les opérations avancées comme le chat multi‑fichiers, les workflows agentiques, les sessions de codage longues et les relectures de code approfondies consomment des AI Credits en fonction des jetons (entrée, sortie et cache) par rapport aux tarifs API publiés du modèle concerné.

Les prix d’abonnement mensuels de base sont restés stables, mais ce changement modifie la manière dont les administrateurs budgètent les dépassements et suivent l’usage effectif.

| **Niveau de formule** | **Public visé** | **Tarif de base** | **Volume mensuel inclus** | **Différenciateurs clés** | 
| Free | Utilisateurs individuels occasionnels | Gratuit | AI Credits limités | Accès basique aux complétions et au Chat. | 
| Student | Étudiants et enseignants vérifiés | Gratuit | AI Credits étendus | Accès plus large aux modèles pour l’apprentissage. | 
| Pro | Développeurs individuels | 10 $ / mois | 1 000 Base + 500 Flex (1 500 au total) | Intégrations IDE étendues et support multi‑modèles. | 
| Pro+ | Utilisateurs individuels intensifs | 39 $ / mois | 3 900 Base + 3 100 Flex (7 000 au total) | Gros quotas de jetons ; inclut l’accès à GitHub Spark. | 
| Business | Équipes et organisations | 19 $ / utilisateur / mois | 1 900 crédits / utilisateur (3 000 du 1er juin au 1er sept. 2026) | Gestion centralisée des sièges, journaux d’audit, exclusions de fichiers, indemnisation PI. | 
| Enterprise | Grandes entreprises | 39 $ / utilisateur / mois | 3 900 crédits / utilisateur (7 000 du 1er juin - 1er sept. 2026) | Indexation des dépôts, affinement personnalisé, gouvernance globale. | 

### Formules individuelles : Free, Student, Pro et Pro+

Les formules individuelles diffèrent par l’accès aux modèles, les limites d’usage et les capacités expérimentales. Par exemple, la formule Free permet l’exploration de base, tandis que Pro+ donne accès à GitHub Spark, un environnement conçu pour créer des applications assistées par l’IA.

Actuellement, les nouvelles inscriptions aux comptes payants individuels de GitHub, comme Pro, Pro+ et Student, sont suspendues. Les comptes existants peuvent passer de Pro à Pro+, mais les nouveaux comptes ne peuvent pas s’inscrire tant que GitHub n’a pas achevé la transition vers la facturation aux AI Credits.

### Business et Enterprise

Avec Business et Enterprise, les formules GitHub Copilot passent d’une simple extension d’IDE à un véritable actif d’infrastructure d’entreprise, entièrement auditable.

GitHub Copilot Business introduit des fonctionnalités essentielles de gestion :

- Attribution et retrait centralisés des sièges.
- Politiques de référence à l’échelle de l’organisation.
- Journaux d’audit structurés et suivi des événements de conformité.
- Exclusions de contenu et de fichiers de dépôts.
- Indemnisation commerciale en propriété intellectuelle.

GitHub Copilot Enterprise va plus loin en matière de contrôle et de capacités :

- Copilot Spaces : un hub de connaissances permettant d’interroger Copilot sur la documentation interne, les wikis et les normes de code système.
- Intégration renforcée de GitHub.com Chat.
- Héritage hiérarchique des politiques entre organisations filles.

GitHub Copilot Enterprise nécessite un abonnement actif à GitHub Enterprise Cloud. Comme GitHub Enterprise Cloud coûte 21 $ par utilisateur et par mois et que la licence Copilot Enterprise est à 39 $ par utilisateur et par mois, le coût minimal réel est de 60 $ par utilisateur et par mois pour Enterprise. Cela ne s’applique pas à la formule GitHub Copilot Business, qui peut être achetée nativement par des organisations utilisant GitHub Free ou GitHub Team.

Les organisations n’ont pas accès aux avantages Enterprise comme l’héritage des politiques, mais disposent tout de même de l’indemnisation PI, de l’audit, de l’exclusion de fichiers et de la gestion des politiques au niveau organisationnel ; une bonne alternative pour des équipes d’ingénierie de taille moyenne.

Si vous envisagez un abonnement Enterprise, notre guide GitHub Copilot Enterprise vous montrera comment exploiter ses fonctionnalités, comme Copilot Spaces et la nouvelle Usage Metrics API.

## Ce qui distingue les formules individuelles des formules Business

La gestion des données, l’indemnisation PI et la facturation sont les principaux domaines où les formules individuelles et Business diffèrent fortement. Les fonctionnalités supplémentaires pour les utilisateurs sont utiles, mais comprendre ces écarts est essentiel pour décider entre gérer une pile de licences Pro personnelles et souscrire une formule Business.

