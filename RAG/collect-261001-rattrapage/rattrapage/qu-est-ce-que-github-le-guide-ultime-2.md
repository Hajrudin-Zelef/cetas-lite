---
id: collect-261001-rattrapage/rattrapage/qu-est-ce-que-github-le-guide-ultime-2
title: "qu-est-ce-que-github-le-guide-ultime"
domain: rattrapage
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["agent", "agents", "aws", "open source"]
source: docs/RAG/collect-261001-rattrapage/qu-est-ce-que-github-le-guide-ultime.md
source_anchor: ""
source_lines: [107, 218]
sha256: ea86e93060615da8221a5bf7fc7c2ba78a427d4b5578662131d191584fdfc7fd
---

# qu-est-ce-que-github-le-guide-ultime

1. 
Accédez au fichier que vous souhaitez modifier dans votre dépôt.
2. 
Modifiez-le en cliquant sur l'icône crayon, puis effectuez vos changements dans l'éditeur de texte.
3. 
Cliquez sur **Commit****Changes** . Faites défiler jusqu'à la section**Commit****changes** . Renseignez un message de commit décrivant précisément vos modifications — c'est essentiel.

*Commits. Image par l'auteur*

### Créer une pull request

Une fois vos commits réalisés, créez une pull request. Procédure :

1. 
Allez dans l'onglet **Pull requests** de votre dépôt.
2. 
Cliquez sur **New****pull****request** . GitHub comparera automatiquement les changements entre les branches.
3. 
Vérifiez les différences pour confirmer que tout est correct, en comparant votre branche avec la branche principale.
4. 
Créez la pull request en cliquant sur **Create****pull****request** . Ajoutez un titre et une description.
5. 
Ajoutez des reviewers si nécessaire, puis soumettez. C'est facultatif, mais très utile si vous souhaitez un regard expert et des retours avant fusion.

Création d'une pull request. Image par l'auteur

### Fusionner des branches

Après relecture et approbation de votre pull request, il ne reste plus qu'à fusionner. Étapes :

1. Une fois la pull request approuvée, ouvrez-la dans l'onglet **Pull****requests** .
2. Cliquez sur le bouton **Merge****pull****request** .
3. Confirmez la fusion avec **Confirm****merge** .

Au-delà de l'interface GitHub, vous pouvez réaliser la plupart des actions avec Git en ligne de commande. Par exemple, créez un dépôt Git en local et récupérez les changements d'un dépôt distant avec la commande `git pull`. Se familiariser avec cet environnement devient crucial au fil de votre carrière. Le cours Introduction to Git de DataCamp est une excellente ressource pour progresser. Vous pouvez aussi commencer avec notre GitHub and Git Tutorial for Beginners. 

## Alternatives à GitHub

GitHub est la plateforme la plus populaire pour la gestion de versions et la collaboration. Cependant, plusieurs alternatives proposent des approches et avantages spécifiques. Il est important, pour les professionnels de la data, de connaître ces solutions susceptibles de mieux répondre à certains besoins.

Voici un aperçu de quelques alternatives à GitHub.

### GitLab

Il s'agit d'une plateforme DevOps qui offre la gestion de dépôts Git, le suivi des issues et des pipelines CI/CD. GitLab est apprécié pour son approche tout-en-un intégrant un large éventail d'outils et de services.

*Logo GitLab*

Quelques fonctionnalités clés :

- **CI/CD intégrée :** Des pipelines intégrés pour l'intégration et le déploiement continus, automatisant les tests et la mise en production.
- **Suivi des issues :** Un système pour gérer tâches, bugs et demandes de fonctionnalités, avec tableaux, jalons et étiquettes.
- **Auto DevOps :** Configure automatiquement les pipelines selon la structure du projet, ce qui accélère la mise en place.
- **Registry de conteneurs :** Un registre intégré pour gérer les images Docker.
- **Fonctions de sécurité :** Des outils de tests de sécurité intégrés pour le monitoring continu et la détection de vulnérabilités dans votre code.

### Bitbucket

Bitbucket est un produit de la suite Atlassian qui gère des dépôts Git, y compris des dépôts Mercurial. Il s'intègre naturellement avec d'autres produits Atlassian comme Jira et Trello.

*Logo Bitbucket*

Fonctionnalités clés :

- **Permissions par branche :** Contrôle fin de l'accès et des droits de modification sur les branches, pour plus de sécurité.
- **Intégration CI/CD :** Intégration avec Bitbucket Pipelines pour l'intégration et la livraison continues, avec tests et déploiements automatisés.
- **Intégration Jira :** Connexion directe à Jira pour une gestion de projet complète et le suivi de l'avancement du développement.
- **Pull requests :** Commentaires en ligne et revues de PR pour des revues de code efficaces.
- **Déploiement :** Prise en charge des déploiements via Bitbucket Pipelines pour des mises en production fluides.

### SourceForge

SourceForge est l'une des premières plateformes d'hébergement et de gestion de projets open source. Elle propose de nombreux outils pour les développeurs, dont des dépôts de code, le suivi des bugs et des fonctions de gestion de projet.

Fonctionnalités principales :

- **Hébergement de projets :** Hébergement gratuit pour les projets open source, avec bande passante et stockage illimités.
- **Statistiques de téléchargement :** Des statistiques détaillées pour suivre la popularité et la progression des projets.
- **Forums de discussion :** Des forums intégrés pour animer la communauté et assurer le support.
- **Système de mise à disposition de fichiers :** Un système avancé pour gérer les téléchargements et les versions.
- **Support SVN et Git :** Prise en charge de Subversion (SVN) et de Git.

### AWS CodeCommit

AWS CodeCommit est un service de gestion de code source entièrement managé par Amazon Web Services. Il permet de stocker et gérer votre code de manière sécurisée dans le cloud, avec une intégration fluide aux autres services AWS.

Fonctionnalités principales :

- **Scalabilité :** S'adapte à vos besoins de dépôts sans gestion d'infrastructure.
- **Sécurité :** Intégration avec AWS Identity and Access Management (IAM) pour un contrôle d'accès granulaire.
- **Intégration :** Fonctionne avec CodePipeline, CodeBuild et CodeDeploy pour une chaîne CI/CD complète.
- **Haute disponibilité :** Bâti sur l'infrastructure AWS pour une haute disponibilité et durabilité de vos dépôts.
- **Chiffrement :** Données chiffrées en transit et au repos pour renforcer la sécurité.

### Cursor Origin

Cursor Origin est une forge Git proposée par Anysphere, l'équipe derrière l'éditeur de code IA Cursor. À la différence des plateformes plus anciennes, elle est pensée dès l'origine pour le développement assisté par IA : elle héberge vos dépôts comme toute forge Git, mais traite les commits et pull requests générés par des agents comme des éléments de première classe. Vous pouvez aussi mirrorer vos dépôts GitHub existants vers Origin, ce qui permet de l'adopter sans bouleverser votre workflow actuel.

Fonctionnalités principales :

- **Conception "AI-agent-first" :** Les agents de codage de Cursor peuvent pousser des commits, ouvrir des pull requests et agir directement sur les dépôts, idéal pour le développement piloté par agents.
- **Mirroring GitHub :** Miroir de vos dépôts GitHub existants vers Origin pour une adoption progressive plutôt qu'une migration complète.
- **Intégration Cursor :** Intégration étroite avec l'éditeur Cursor et ses agents cloud pour réunir hébergement, édition et exécution d'agents au sein d'un même écosystème.
- **Workflows Git standard :** Dépôts, branches et pull requests fonctionnent comme prévu, avec une mise en route pilotée en CLI.
- **Revue de pull requests :** Relecture intégrée pour valider les PR générées par IA avant leur fusion dans la branche principale.

Pour un pas-à-pas via la CLI et le mirroring d'un dépôt GitHub, consultez notre tutoriel Cursor Origin.

## Conclusion

Dans cet article, nous avons passé en revue les bases de GitHub, son importance et son fonctionnement. Nous avons d'abord clarifié la notion de gestion de versions et son rôle pour maîtriser les évolutions du code et des données. Enfin, nous avons proposé un guide pratique de GitHub : inscription, création de dépôts, commits, gestion des pull requests et fusion des branches.

