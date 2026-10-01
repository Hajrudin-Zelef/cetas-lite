---
id: collect-261001-rattrapage/rattrapage/qu-est-ce-que-github-le-guide-ultime-1
title: "qu-est-ce-que-github-le-guide-ultime"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["open source"]
source: docs/RAG/collect-261001-rattrapage/qu-est-ce-que-github-le-guide-ultime.md
source_anchor: ""
source_lines: [1, 106]
sha256: f2d2cbb77200620dabdabb57f27a6c7dfb80b58227f76f77b5e9d94ceea54bf3
---

# qu-est-ce-que-github-le-guide-ultime

Cours

*Logo GitHub. Source : GitHub Logos and Usage* 

Imaginez que vous travaillez sur un projet de data science et que vous avez bien avancé. Soudain, un bug survient. Vous aimeriez revenir à la dernière version fonctionnelle, mais vous ne vous souvenez plus de toutes les modifications effectuées. Ou peut-être collaborez-vous avec d'autres personnes, et fusionner les contributions de chacun devient un casse-tête. Si ces situations vous parlent, vous n'êtes pas seul.

Ces problèmes courants se résolvent avec GitHub, la plateforme de référence pour la gestion de versions et la collaboration. Dans cet article, nous verrons comment GitHub peut transformer la manière dont vous pilotez vos projets data. Nous explorerons également des techniques de collaboration et des stratégies pour gagner en productivité.

Commençons par les fondamentaux de la gestion de versions.

## Qu'est-ce que la gestion de versions ?

La gestion de versions est un système qui enregistre l'historique des modifications apportées à des fichiers dans le temps. Elle permet à plusieurs personnes de travailler ensemble sur un projet tout en conservant une trace précise des évolutions. Sans gestion de versions, le suivi des changements de code devient vite chaotique et source d'erreurs, surtout en équipe lorsque plusieurs contributeurs travaillent en parallèle sur des parties différentes du code.

## À quoi sert GitHub ?

Comme vous vous en doutez, GitHub excelle en gestion de versions. Mais la plateforme va bien au-delà, avec de nombreux autres usages, notamment :

- 
**Créer un portfolio de projets :** GitHub vous permet de créer un profil public pour présenter vos compétences et vos projets data à des recruteurs ou à vos collègues.
- 
**Collaborer :** GitHub facilite le travail d'équipe sur des projets, le partage d'extraits de code et la relecture mutuelle des contributions.
- 
**Contribuer à l'open source :** Avec GitHub, vous pouvez explorer et contribuer à des projets open source en data science, accélérant ainsi votre apprentissage et l'innovation.

## Comment fonctionne GitHub ?

Pour tirer pleinement parti de GitHub, il est essentiel d'en comprendre les composants clés et leur fonctionnement conjoint.

- **Dépôts (repositories) :** Ce sont des dossiers qui stockent les fichiers de votre projet et l'historique de leurs versions. Pensez-y comme à une armoire numérique pour vos projets data. Chaque dépôt a une URL unique et contient des fichiers, des branches et des commits.
- **Forks :** Un fork est une copie personnelle du dépôt d'un autre utilisateur. Vous pouvez y apporter des changements de manière indépendante, puis proposer de les réintégrer au dépôt d'origine.
- **Pull requests :** Les PR sont un moyen formel de proposer vos modifications au propriétaire du projet pour relecture et fusion. Elles facilitent la revue de code et la collaboration.
- **Issues :** Elles servent à suivre des tâches, des bugs ou des améliorations.
- **Branches :** Une branche est une version parallèle d'un dépôt. Créez des branches pour développer une fonctionnalité ou corriger un problème, puis fusionnez-les dans la branche principale lorsqu'elles sont prêtes. Pour en savoir plus, consultez ce tutoriel sur Git Clone Branch.
- **Fusion (merging) :** La fusion combine vos modifications avec le projet d'origine pour garder le tout organisé et à jour, par exemple en fusionnant une branche de fonctionnalité dans la branche principale.

## Git vs. GitHub

Vous vous demandez peut-être quel est le lien entre Git et GitHub. Ces deux termes sont parfois confondus, mais il existe une différence essentielle.

**Git est un système de gestion de versions distribué (DVCS) qui aide les développeurs à gérer leur code. Il suit les changements et permet de créer différentes versions, ou branches, du code, ce qui facilite le travail collaboratif. Git propose aussi des fonctionnalités comme la zone d'index (staging) et l'historique des commits, offrant une traçabilité fine des modifications.**

GitHub, pour sa part, ajoute des fonctionnalités comme le contrôle d'accès, le suivi des bugs, la gestion des tâches et des wikis, ce qui simplifie la collaboration sur les projets. Avec GitHub, vous pouvez gérer votre code, suivre les changements, relire les contributions et discuter des problèmes — le tout au même endroit. La plateforme s'intègre en outre à de nombreux outils et services pour fluidifier les workflows de développement. 

| Catégorie | Git | GitHub | 
| Définition | Système de gestion de versions distribué | Plateforme web construite au-dessus de Git | 
| Objectif | Aide à gérer le code, suivre les changements et créer des branches | Héberge des dépôts Git et fournit des outils de collaboration supplémentaires | 
| Fonctionnalités | Zone d'index, historique des commits, branches et fusion | Contrôle d'accès, suivi des bugs, gestion des tâches, wikis et intégrations | 
| Bénéfice | Permet le travail collaboratif et un suivi détaillé des modifications de code | Améliore la collaboration, la gestion de projet et les processus de revue de code | 

## Comment utiliser GitHub

Jusqu'ici, nous avons défini GitHub et la gestion de versions, et comparé Git à GitHub. Passons maintenant à la pratique.

Nous allons d'abord voir comment créer un compte GitHub, personnaliser votre expérience et choisir une formule. Ensuite, nous créerons un dépôt, en le configurant, en ajoutant une description et en gérant sa visibilité. Puis nous aborderons la création de branches pour travailler sur différentes versions du projet, avant de passer aux commits, où nous apprendrons à modifier des fichiers et à documenter les changements.

### Créer un compte

Voici les étapes pour créer un compte GitHub et démarrer :

1. Rendez-vous sur GitHub et cliquez sur le bouton **Sign****up** .
2. Suivez les instructions pour créer votre compte. Indiquez votre adresse e-mail, choisissez un nom d'utilisateur et un mot de passe.
3. Personnalisez votre expérience en choisissant une formule adaptée et en ajustant vos préférences lors de la configuration. L'offre gratuite suffit largement aux débutants et aux profils data juniors.

    *Créer un compte GitHub. Image par l'auteur*

### Créer un dépôt

Après la création du compte, la prochaine étape est de créer un dépôt. Procédez ainsi pour votre premier dépôt :

1. 
Cliquez sur l'icône **+** en haut à droite et sélectionnez**New****repository** .
2. 
Ajoutez un nom et une description, puis choisissez si le dépôt doit être public ou privé. Les dépôts publics sont visibles de tous. Les dépôts privés ne sont accessibles qu'à vous et aux collaborateurs que vous invitez.
3. 
Vous pouvez, en option, ajouter un fichier README, un fichier `.gitgnore` et une licence. Vous pourrez aussi les ajouter plus tard.
4. 
Cliquez sur **Create****repository** .

*Création d'un dépôt. Image par l'auteur*

Une fois ces étapes suivies, une fenêtre de configuration rapide pour votre nouveau dépôt s'affiche. Vous pouvez démarrer en créant un nouveau fichier ou en téléversant un fichier existant dans le dépôt.

*Configuration d'un nouveau dépôt. Image par l'auteur*

*Téléversement de fichiers. Image par l'auteur*

### Créer des branches

Une fois le dépôt prêt, créez des branches. Voici comment procéder :

1. 
Dans votre dépôt, cliquez sur **Branch:main** , près du haut de la page.
2. 
Cliquez ensuite sur le bouton **New branch** en haut à droite.
3. 
Saisissez un nom de branche et cliquez sur **Create new branch** .
4. 
Vous pouvez basculer d'une branche à l'autre via le menu déroulant des branches et sélectionner celle sur laquelle vous voulez travailler.

*Création de branches. Image par l'auteur*

### Effectuer des commits

Après la création de branches, passez aux commits. Étapes à suivre :

