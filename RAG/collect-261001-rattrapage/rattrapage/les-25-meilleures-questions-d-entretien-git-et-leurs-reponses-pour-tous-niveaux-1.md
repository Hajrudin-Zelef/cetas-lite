---
id: collect-261001-rattrapage/rattrapage/les-25-meilleures-questions-d-entretien-git-et-leurs-reponses-pour-tous-niveaux-1
title: "Git extrait un commit intermédiaire ; vous testez, puis :"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/les-25-meilleures-questions-d-entretien-git-et-leurs-reponses-pour-tous-niveaux.md
source_anchor: ""
source_lines: [1, 99]
sha256: 5b48a64a33d1d583f27a607714214384b6dce66b1434c38cebb3230bd659609b
---

# Git extrait un commit intermédiaire ; vous testez, puis :

Cours

Git est un outil incontournable dans la boîte à outils des développeurs modernes, reconnu pour ses puissantes capacités de gestion de versions. Créé par Linus Torvalds en 2005 pour soutenir le développement du noyau Linux, Git est depuis devenu l’ossature d'innombrables projets logiciels à travers le monde. Son efficacité et sa flexibilité dans la gestion des versions, associées à un solide support de la collaboration, en font un indispensable pour des équipes de toutes tailles.

Cet article vise à vous préparer aux entretiens techniques en couvrant les 20 principales questions d'entretien Git, du niveau débutant au niveau avancé. Que vous découvriez Git ou que vous cherchiez à approfondir vos connaissances, ces questions-réponses vous aideront à démontrer votre maîtrise et à réussir votre entretien.

## Devenez ingénieur en données

## Questions Git pour débuter

Si vous débutez avec Git, certaines questions d'entretien porteront probablement sur des notions de base et des cas d'usage simples. Pour réviser ces fondamentaux, consultez le cours Introduction to Git de DataCamp.

### Qu’est-ce qu’un dépôt Git ?

Un dépôt Git stocke les fichiers d’un projet et l’historique de leurs révisions, et facilite le contrôle de version en suivant les modifications dans le temps. Il peut être hébergé localement dans un dossier sur votre machine ou en ligne sur une plateforme comme GitHub. Cela permet de collaborer, de revenir à des versions antérieures et de gérer efficacement le développement via des commandes comme commit, push et pull.

### Comment fonctionne Git ?

Git enregistre les changements apportés aux fichiers et répertoires d’un projet sous forme d’instantanés successifs. Vous pouvez suivre les modifications, créer des branches pour développer en parallèle, fusionner des branches et revenir à des états antérieurs si nécessaire. Il favorise la collaboration et assure un contrôle de version efficace dans les projets logiciels.

### Qu’est-ce que git add ?

La commande `git add` sert à mettre en scène (stager) des modifications en vue du prochain commit. Elle prépare les ajouts, suppressions ou modifications effectués dans l’espace de travail afin qu’ils soient inclus dans le prochain instantané. Notez qu’elle ne valide pas les changements : elle se contente de les placer dans l’index.

### Qu’est-ce que git push ?

La commande `git push` permet d’envoyer le contenu du dépôt local vers un dépôt distant. Elle transfère les commits du dépôt local vers un serveur distant, typiquement GitHub ou GitLab. Cette commande favorise la collaboration en partageant vos changements avec les autres membres du projet.

Pour en savoir plus sur git push et git pull, consultez notre tutoriel dédié.

### Qu’est-ce que git status ?

La commande `git status` affiche l’état courant du dépôt. Elle indique quels fichiers ont été modifiés, lesquels sont prêts à être commités et lesquels ne sont pas suivis. Elle aide à suivre l’avancement du travail et à identifier ce qui doit être ajouté à l’index ou validé.

### Qu’est-ce qu’un commit dans Git ?

Un commit représente un instantané des changements apportés aux fichiers d’un dépôt à un moment donné. Lorsque vous validez vos modifications, vous enregistrez l’état actuel de vos fichiers et vous pouvez ajouter un message descriptif expliquant les changements (fortement recommandé).

Chaque commit possède un identifiant unique, ce qui permet de retracer l’historique du dépôt. Les commits sont essentiels au contrôle de version : ils permettent de revenir en arrière, de relire l’historique des modifications et de collaborer en partageant des mises à jour.


*Consultez l’aide-mémoire Git de DataCamp pour préparer votre entretien*

### Qu’est-ce que le branchement dans Git ?

Le branchement consiste à diverger de la ligne principale de développement (généralement appelée `main`, anciennement `master`) afin de travailler sur de nouvelles fonctionnalités, des correctifs ou des expérimentations sans impacter la base de code principale. Cela permet de faire coexister plusieurs lignes de développement en parallèle au sein d’un même dépôt.

Chaque branche suit sa propre lignée de commits, ce qui permet à plusieurs développeurs de travailler simultanément sur des sujets distincts. Le branchement facilite la collaboration, l’expérimentation et l’organisation : une fois les changements terminés et testés, ils peuvent être fusionnés dans la branche principale.

### Qu’est-ce qu’un conflit dans Git ?

Les conflits surviennent lorsque des modifications incompatibles sont apportées à la même portion d’un fichier par des contributeurs différents, généralement lors d’une fusion (merge) ou d’un rebase. Git ne peut pas les résoudre automatiquement et requiert une intervention manuelle.

Pour résoudre un conflit, ouvrez le fichier concerné : Git marque les sections en conflit avec les indicateurs `<<<<<<<`, `=======` et `>>>>>>>`. Éditez le fichier pour conserver la bonne version, supprimez les marqueurs, puis :

```
git add <resolved-file>
git commit
```
`git mergetool` rendent ce processus plus visuel et facile à suivre.
### Qu’est-ce que la fusion (merge) dans Git ?

La fusion est une opération fondamentale qui facilite la collaboration et l’intégration des changements entre différentes branches d’un projet. Elle consiste à combiner les modifications de plusieurs branches dans une seule, généralement la branche principale (par ex. master ou main).

Une fusion intègre les changements d’une branche dans une autre et crée un nouveau commit qui rassemble les historiques des deux branches. Pour en savoir plus sur la résolution des conflits de fusion, consultez notre tutoriel dédié.

## Obtenez une certification pour le poste de Data Engineer de vos rêves

Nos programmes de certification vous aident à vous démarquer et à prouver aux employeurs potentiels que vos compétences sont adaptées à l'emploi.

## Questions Git intermédiaires

### Qu’est-ce qu’un remote dans Git ?

Un remote est un dépôt hébergé sur un serveur ou un autre ordinateur pour collaborer et partager du code. Il sert de point central où les développeurs peuvent pousser leurs changements locaux et récupérer ceux des autres.

Les remotes sont généralement configurés sur des plateformes comme GitHub, GitLab ou Bitbucket. Ils permettent un développement distribué et facilitent le travail en équipe en offrant un emplacement commun pour stocker et synchroniser le code entre plusieurs contributeurs.

### Comment annuler un commit déjà poussé et rendu public ?

La commande `git revert <commit-hash>` permet d’annuler un commit déjà poussé et public.

Procédure pas à pas :

1. Identifiez le commit à annuler en trouvant son hash. Utilisez `git log` pour parcourir l’historique et récupérer le hash souhaité.

2. Une fois le hash obtenu, exécutez la commande git revert suivie de ce hash pour créer un nouveau commit qui annule les changements introduits par le commit ciblé. Par exemple :

`git revert <commit-hash>`
3. Git ouvre un éditeur pour saisir le message de revert. Modifiez-le si nécessaire, puis enregistrez et fermez.

4. Après l’enregistrement, Git crée un nouveau commit qui annule effectivement les changements introduits par le commit d’origine. Ce nouveau commit s’ajoute à l’historique.

5. Enfin, poussez le nouveau commit vers le dépôt distant pour rendre le revert public :

`git push origin <branch-name>` Avec `git revert`, vous créez un commit d’annulation sans réécrire l’historique. C’est plus sûr que `git reset` ou `git amend`, qui peuvent modifier l’historique et perturber des collaborateurs ayant déjà récupéré les changements.

### Qu’est-ce que git stash ?

