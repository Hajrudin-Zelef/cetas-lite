---
id: collect-261001-rattrapage/rattrapage/docker-vs-podman-quel-outil-de-conteneurisation-vous-convient-le-mieux-1
title: "docker-vs-podman-quel-outil-de-conteneurisation-vous-convient-le-mieux"
domain: rattrapage
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-rattrapage/docker-vs-podman-quel-outil-de-conteneurisation-vous-convient-le-mieux.md
source_anchor: ""
source_lines: [1, 70]
sha256: 4fe033043f9dbdfe68618822f425ca5b47ad9870800c14a1d4ab0369d254f747
---

# docker-vs-podman-quel-outil-de-conteneurisation-vous-convient-le-mieux

Cours

Les conteneurs exécutent les applications et les charges de travail de données du monde entier. Conçus pour la première fois dans les années 1970, les conteneurs permettent de regrouper en un seul objet tout ce qui est nécessaire à l'exécution d'une application ou d'une charge de travail. Les conteneurs permettent de résoudre le problème du "ça tourne sur ma machine" en offrant une solution isolée et portable pour développer, tester et expédier du code. Des outils comme Kubernetes s'appuient fortement sur les conteneurs en tant qu'élément central de leur architecture. Pour l'instant, les conteneurs ne vont nulle part.

Pour faire fonctionner ces conteneurs, vous aurez besoin d'une solution de gestion de conteneurs. Entrez dans Docker et Podman.

Docker et Podman sont utilisés pour construire, gérer et déployer des conteneurs. Ensemble, nous allons analyser les similitudes et les différences entre Docker et Podman, ainsi que les fonctionnalités uniques de chacun d'entre eux. Nous explorerons des sujets tels que l'architecture avec ou sans démon, la gestion multi-conteneurs et l'intégration multiplateforme. À la fin, vous serez armé des informations dont vous avez besoin pour choisir la solution de conteneur parfaite pour vos besoins.

**Si vous ne connaissez pas encore ces outils, vous pouvez également consulter notre Introduction à Docker et notre tutoriel Introduction à Podman pour l'apprentissage automatique.** 

## Devenez ingénieur en données

## Que sont Podman et Docker ?

Commençons par une vue d'ensemble de ces outils afin d'entamer notre comparaison :

### Aperçu de Docker

Docker est la norme de facto pour la construction, l'exécution et l'expédition de conteneurs. Les conteneurs sont des objets qui combinent des dépendances au niveau du système d'exploitation et une sorte de code d'application pour emballer et exécuter des éléments tels que des applications complètes ou des pipelines ETL dans leur propre environnement isolé. Les conteneurs sont comme de petits ordinateurs qui ne disposent que de l'essentiel pour exécuter une sorte de code.

Docker est assez jeune et a été publié pour la première fois en tant que projet open-source en 2013. Depuis, le projet a explosé.

Lorsqu'il s'agit d'exécuter des conteneurs dans une entreprise, presque toutes les équipes chargées des logiciels et des données utilisent Docker.

Les développeurs peuvent utiliser Docker sur les trois principaux systèmes d'exploitation, et il s'intègre parfaitement à presque toutes les technologies modernes. Cela signifie qu'un ingénieur de données peut écrire et empaqueter un pipeline de données en utilisant un conteneur Docker sur son Mac local, et expédier ce conteneur pour qu'il s'exécute sur AWS ECS.

Des outils tels que Docker CLI, Docker Desktop et Docker Hub permettent aux développeurs de tous niveaux de démarrer facilement.

Si vous cherchez un moyen plus pratique d'apprendre Docker, nous avons plusieurs projets Docker et des informations sur les certifications Docker pour vous aider à améliorer vos compétences en matière de Docker !

### Aperçu de Podman

Comme Docker, Podman est un outil open-source pour le développement et la gestion de conteneurs. Podman a été développé à l'origine par Red Hat en tant qu'alternative Linux-native à Docker et a été publié en 2019.

Plus particulièrement, l'architecture sous-jacente des deux systèmes d'exécution de conteneurs diffère ; alors que Docker utilise des démons, Podman fonctionne sans démon (plus d'informations à ce sujet plus loin).

Contrairement à Docker, Podman ne nécessite pas d'accès root à la machine sur laquelle s'exécutent les pods qu'il gère, ce qui en fait une option plus respectueuse de la sécurité pour les équipes qui utilisent des conteneurs pour exécuter leurs applications et leurs charges de travail.

Les utilisateurs de Podman bénéficient d'une expérience utilisateur similaire à celle de Docker ; les développeurs peuvent utiliser un CLI ou une interface graphique (Podman Desktop) pour interagir avec Podman dans leur environnement local.

Les utilisateurs de Linux, Mac et Windows peuvent utiliser Podman pour construire et tester leurs conteneurs localement avant de les déployer dans une sorte d'environnement distant, comme Kubernetes.

## Principales différences entre Podman et Docker

### Architecture avec ou sans démon

La plus grande différence entre Docker et Podman est l'architecture sous-jacente sur laquelle ils sont construits. Docker s'appuie fortement sur un démon, tandis que Podman est sans démon.

Vous pouvez considérer un démon comme un processus qui s'exécute en arrière-plan sur le système d'exploitation hôte. Dans le cas de Docker, son démon est responsable de la gestion des objets Docker (images et conteneurs) et de la communication avec d'autres systèmes. Pour exécuter son démon, Docker utilise un paquetage appelé dockerd.

Pourquoi est-ce important ? Tout d'abord, les démons nécessitent généralement un accès au niveau racine de la machine sur laquelle ils s'exécutent. Cette situation est propice aux failles de sécurité : si un acteur mal intentionné parvient à accéder à un démon, il a désormais accès à l'ensemble de la machine.

L'architecture sans démon de Podman présente quelques avantages. Étant donné que l'exécution des démons nécessite presque toujours les privilèges de l'administrateur, une architecture sans démon peut être considérée comme "sans racine". Cela signifie que les utilisateurs qui n'ont pas d'accès au niveau du système à la machine sur laquelle leurs conteneurs s'exécutent peuvent toujours utiliser Podman, ce qui n'est pas toujours le cas avec Docker.

Au lieu d'un démon, Podman utilise un paquetage Linux connu sous le nom de systemd. Étant donné que systemd fait partie intégrante du système d'exploitation Linux, Podman est souvent considéré comme plus "léger" que Docker ; les utilisateurs de Podman constatent généralement des temps de démarrage des conteneurs plus rapides que lorsqu'ils utilisent Docker.

### Construire des images et des conteneurs

Malgré leurs architectures fondamentalement différentes, Docker et Podman partagent le même objectif principal : créer et exécuter des images et des conteneurs. Cependant, leurs approches de ce processus diffèrent légèrement.

Avec Docker, une image est construite en ajoutant d'abord des commandes à un Dockerfile. Ensuite, une commande telle que docker build est exécutée. Cette opération appelle chacune des instructions du fichier Docker , ce qui permet de créer une image. Une image peut ensuite être "exécutée" en tant que conteneur. Comme vous l'avez peut-être deviné, cela se fait en utilisant la commande docker run, et en spécifiant un ID d'image ou un tag. Pour construire et exécuter plusieurs conteneurs, nous utiliserons un outil spécial appelé docker-compose, que nous explorerons un peu plus loin.

Le processus de construction des images et de leur exécution en tant que conteneurs est presque identique dans Podman. Plutôt qu'un fichier Docker (bien que ce nom de fichier fonctionne toujours), les utilisateurs de Podman créeront un fichier Container. La syntaxe de composition de l'image est la même. Une fois les commandes appropriées ajoutées aufichier de conteneurs , l'image peut être construite et exécutée à l'aide de l'interface de programmation Podman.

Pour l'essentiel, Podman est compatible avec la plupart des éléments de Docker. Vous trouverez des différences ici et là, mais pour l'essentiel, l'interface de programmation de Docker peut être remplacée par l'interface de programmation de Podman sans problème.

## Podman vs Docker Desktop

### Docker Desktop pour un accès multiplateforme simplifié

