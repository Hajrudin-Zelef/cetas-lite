---
id: collect-261001-rattrapage/rattrapage/qu-est-ce-que-git-guide-du-debutant-sur-le-controle-de-version-git-1
title: "qu-est-ce-que-git-guide-du-debutant-sur-le-controle-de-version-git"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["open source"]
source: docs/RAG/collect-261001-rattrapage/qu-est-ce-que-git-guide-du-debutant-sur-le-controle-de-version-git.md
source_anchor: ""
source_lines: [1, 87]
sha256: 9abc6cfc5b71342c67b6559c841016aaf5e31289c650280e92b583efd7b9fd88
---

# qu-est-ce-que-git-guide-du-debutant-sur-le-controle-de-version-git

Si vous avez déjà lu quoi que ce soit sur le code, la programmation ou le développement logiciel, vous avez entendu parler de Git.

Pratique (et gratuit), cet outil est le système de contrôle de version le plus utilisé au monde. Il est si répandu que plus de 90 % des développeurs professionnels s'en servent, sans parler des experts d'autres domaines.

À bien des égards, Git est devenu synonyme de contrôle de version. Mais qu'est-ce que le contrôle de version et pourquoi est-ce si important ?

Plongez avec nous dans l'univers Git. Nous allons examiner de près ce qu'est Git, qui l'utilise et son histoire.

## Qu'est-ce que Git ?

Git est un système de contrôle de version distribué (dVCS). Comme son nom l'indique, le contrôle de version consiste à gérer et à suivre différentes versions d'un projet donné.

### Qu'est-ce qu'un système de contrôle de version (VCS) ?

Un VCS suit et enregistre les modifications apportées à n'importe quel fichier (ou groupe de fichiers), ce qui vous permet de retrouver des itérations spécifiques plus tard ou à la demande. Les VCS sont parfois appelés gestion de code source (SCM) ou systèmes de gestion de révisions (RCS).

Le contrôle de version permet à de nombreux membres d'une équipe de travailler ensemble sur un projet, même s'ils ne sont pas dans la même pièce, ni même dans le même pays.

Par exemple, disons que vous êtes auteur-compositeur. Vous travaillez chez vous sur une nouvelle chanson, mais vous n'en êtes pas tout à fait satisfait. Vous décidez donc de collaborer avec deux autres auteurs pour retravailler les passages à améliorer.

Vous et les deux autres auteurs commencez à retoucher les paroles et la partition, chacun travaillant de son côté. Lorsque les autres musiciens vous envoient leurs versions, certaines modifications vous plaisent, d'autres moins.

Imaginez maintenant que vous puissiez voir chaque changement dans chaque version, les tester pour entendre le résultat, puis synchroniser les retouches que vous retenez à travers les versions.

C'est exactement ce que permet Git. Chacun peut travailler localement (sur son propre ordinateur), enregistrer les changements qui fonctionnent, puis synchroniser ces changements vers un dépôt Git afin que les autres puissent voir la nouvelle version.

On pense souvent à Git comme à un outil de développement logiciel — ce qu'il est —, mais il peut être utilisé pour versionner tout type de fichier : lignes de code, maquette de site web ou même une chanson.

### Les avantages du contrôle de version

Au-delà de la collaboration, le contrôle de version présente d'autres atouts :

- **Modifications attribuables** . Chaque changement peut être rattaché à un membre de l'équipe.
- **Traçabilité fine et retour en arrière facilité** . Comme chaque changement est tracé, même les plus petits, il est facile de revenir à une version antérieure si nécessaire. Comme vous pouvez l'imaginer, c'est essentiel en développement logiciel.
- **Meilleure organisation et communication.** Les messages de commit — qui expliquent à l'équipe pourquoi vous avez effectué un changement — favorisent une bonne communication. Ils vous aident aussi à vous y retrouver si vous oubliez ce que vous avez modifié par le passé !
- **Concurrence** . Dans les projets logiciels, les développeurs modifient souvent le code source. En général, plusieurs personnes travaillent sur des sujets différents. L'une améliore la sécurité, l'autre développe une nouvelle fonctionnalité. Git leur permet de travailler en parallèle tout en limitant les conflits entre leurs changements.
- **Branches et fusions.** Les membres de l'équipe peuvent créer des branches séparées pour travailler, puis fusionner leurs modifications dans la branche principale. Les branches sont temporaires et peuvent être supprimées après la fusion.

### Git est-il le seul système de contrôle de version ?

Non, Git n'est pas le seul VCS, mais c'est le plus populaire et il est considéré comme l'outil de référence. Parmi les autres systèmes connus, citons Fossil, Mercurial et Subversion.

Ces systèmes présentent des nuances — par exemple dans la gestion des fonctions clés comme le branchement et la fusion —, mais l'idée générale reste la même. La grande différence tient au fait qu'ils soient centralisés ou distribués.

#### Systèmes de contrôle de version centralisés et distribués

Les systèmes centralisés et distribués, comme Git, remplissent la même fonction.

La différence principale : dans les systèmes centralisés, un serveur central reçoit les dernières versions sur lesquelles les membres de l'équipe poussent leur travail. Vous pouvez l'imaginer comme un projet unique et partagé par tous.

Avec les VCS distribués, chaque membre possède une copie locale (clone) de l'historique complet du projet sur son appareil. Il n'a donc pas besoin d'être en ligne pour modifier ou travailler sur son code. Au lieu d'un serveur central, il récupère ce clone depuis un dépôt en ligne.

Avec Git, le clone de chaque membre est un dépôt qui peut contenir tous les changements depuis le début du projet.

## L'histoire de Git

Git a été développé en 2005 par l'ingénieur logiciel finlandais Linus Torvalds, également à l'origine du noyau du système d'exploitation Linux.

Git a été créé pour répondre à un besoin immédiat. Avant son invention, les développeurs Linux du monde entier utilisaient le logiciel propriétaire BitKeeper, lui-même un dVCS.

Comme ce logiciel appartenait à une entreprise, cela créait des tensions au sein de la communauté Linux, majoritairement attachée à l'open source.

En échange d'une utilisation gratuite, BitMover, l'entreprise derrière BitKeeper, imposait des restrictions à la communauté Linux. Selon le Linux Journal, l'une d'elles interdisait de travailler sur des projets concurrents de contrôle de version.

Comme on pouvait s'y attendre, un développeur Linux a commencé à faire de l'ingénierie inverse de BitKeeper pour créer un produit open source. Fidèle à son avertissement, BitMover a cessé de fournir ses services au noyau Linux, plongeant le système de développement distribué dans l'incertitude.

Pour résoudre l'impasse, Torvalds a mis en pause, pour la première fois depuis 1991, ses travaux sur Linux et a créé Git, dont une version stable est sortie quelques mois seulement après le début du développement.

Fait intéressant : avant même l'adoption de BitKeeper, les développeurs envoyaient à Torvalds leurs « patches » (modifications) individuellement, qu'il intégrait au fil de l'eau. Et en 2016, 11 ans après la sortie de Git, BitKeeper est devenu open source.

### Pourquoi Git s'appelle-t-il Git ?

Dans son tout premier commit de code sur Git en 2005, Linus Torvalds a ajouté un fichier readme qui éclaire l'origine du nom. En voici un extrait :

À moins de préférer la version « Global Information Tracker », le nom Git est un clin d'œil un brin moqueur à ses capacités — ou à leur supposée absence.

### Brève histoire des VCS

Les systèmes de contrôle de version existent bien avant Git ou même BitKeeper. Voici une courte chronologie :

- 1972 - SCCS, premier VCS, est créé par Bell Labs. Il ressemble peu aux systèmes actuels.
- 1982 - Le système de gestion de révisions (RCS) est développé par un informaticien de l'université Purdue.
- 1986 - Le concurrent versions system (CVS) est développé. C'est le premier VCS à proposer un dépôt centralisé accessible à plusieurs utilisateurs.
- 1995 - Développement de Perforce, un VCS encore populaire aujourd'hui.
- 2000 - Apparition de Subversion (ou SVN), plus sophistiqué. Et de BitKeeper, l'un des premiers dVCS, qui a popularisé les systèmes distribués.
- 2005 - Invention de Git, qui devient rapidement l'outil privilégié des développeurs du monde entier.

