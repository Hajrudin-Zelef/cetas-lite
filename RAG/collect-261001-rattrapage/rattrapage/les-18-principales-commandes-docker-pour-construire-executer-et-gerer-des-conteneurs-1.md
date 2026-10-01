---
id: collect-261001-rattrapage/rattrapage/les-18-principales-commandes-docker-pour-construire-executer-et-gerer-des-conteneurs-1
title: "syntax=docker/dockerfile:1"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/les-18-principales-commandes-docker-pour-construire-executer-et-gerer-des-conteneurs.md
source_anchor: ""
source_lines: [1, 95]
sha256: e6b134569e6dc37b572f55a4665c49d4e1ac55d22a316b59a03ee37487b9cfe4
---

# syntax=docker/dockerfile:1

Cours

Docker est devenu un outil incontournable pour les développeurs et les professionnels des données qui ont besoin de créer, de partager et d'exécuter des applications de manière cohérente dans différents environnements. Que vous créiez des conteneurs pour le développement local ou que vous déployiez des microservices en production, il est essentiel de maîtriser les commandes Docker.

Dans ce guide, je vais vous présenter 18 commandes Docker essentielles, couvrant les images, les conteneurs, la mise en réseau, les volumes et Compose, qui amélioreront votre flux de travail et rendront votre travail avec Docker plus fluide et plus efficace.

## Qu'est-ce que Docker ?

Docker est une plateformetform pour le développement, l'expédition et l'exécution d'applications. Il vous permet de séparer votre application de l'infrastructure sous-jacente, ce qui accélère la livraison des logiciels et vous permet de gérer votre configuration de la même manière que vous gérez vos applications.

Docker exécute des applications dans des environnements isolés à l'aide de paquets légers appelés conteneurs, qui comprennent tout ce dont une application a besoin pour fonctionner, comme les dépendances et les installations, ce qui vous aide à économiser les ressources du système. Nous pouvons facilement partager des conteneurs avec des coéquipiers, en exécuter plusieurs simultanément et les gérer tous à l'aide des outils et de la plateforme de Docker.

Nous pouvons utiliser Docker pour beaucoup de choses, y compris :

- Déploiement et mise à l'échelle réactifs.
- Exécution d'un plus grand nombre de charges de travail sur le même matériel.
- Livraison rapide et cohérente des applications.

Lorsque vous utilisez Docker, vous travaillez avec différents objets Docker tels que des images, des conteneurs, des réseaux, des plugins et des volumes. Ceux-ci constituent les éléments de base de l'installation Docker. Sous le capot, Docker utilise les caractéristiques du noyau Linux pour faire fonctionner tout cela. Nous interagissons avec lui en utilisant des commandes simples dans le terminal, et chaque commande Docker commence par `docker`.

**> Si vous débutez, l'introduction à Docker vous aidera à vous familiariser avec Docker.***e cours Introduction à Docker offre une base pratique pour apprendre les bases de la conteneurisation.***rse offre une base pratique pour apprendre les principes de base de la conteneurisation.**

## Commandes de base de Docker

Maintenant que nous avons expliqué ce qu'est Docker et comment il fonctionne, examinons quelques-unes des commandes les plus courantes. Ils vous aideront à construire, exécuter et gérer des conteneurs dans votre travail quotidien.

### docker --version et docker info

Dans les commandes Docker, tout ce qui commence par `--` est appelé un drapeau. 

Par exemple, `--version` est un drapeau qui indique la version du CLI de Docker que vous utilisez. Vous pouvez également utiliser `docker version` (sans le drapeau) pour obtenir des informations détaillées sur la version de tous les composants Docker. 

La sortie est divisée en deux parties :

- Client présente des informations sur le CLI de Docker et les outils associés.
- Serveur affiche des détails sur le moteur Docker et ce sur quoi il s'exécute.

Vous pouvez également formater cette sortie en utilisant l'option `--format` avec un modèle personnalisé.

La commande `docker info` vous donne un aperçu complet de votre configuration Docker. C'est la même chose que d'aller sur le site `docker system info`, mais avec un nom plus court. Vous verrez des détails comme la version de votre noyau, le nombre de conteneurs et d'images, ainsi que d'autres détails concernant le système. 

En fonction de votre pilote de stockage, vous pouvez également voir des informations telles que les noms des pools et des fichiers de données. Comme pour `docker version`, vous pouvez formater la sortie en utilisant `--format` ou `-f`.

### docker pull <image>

La commande pull télécharge une image Docker à partir d'un registre, généralement Docker Hub, une bibliothèque publique d'images préconstruites que vous pouvez utiliser sans rien configurer vous-même. Vous pouvez l'exécuter en tant que `docker pull`  ou `docker pull`, et les deux feront la même chose.

- La syntaxe complète se présente comme suit : `docker image pull [OPTIONS] NAME[:TAG|@DIGEST]`

Si vous ne spécifiez pas de balise, Docker utilisera `:latest` par défaut. Par exemple : `docker image pull debian` tire l'image `debian:latest`.

Vous pouvez également ajouter des options après la commande pour personnaliser la manière dont l'image est extraite, par exemple en limitant la bande passante ou en ignorant la vérification de l'image. L'image ci-dessous présente toutes les options disponibles et leurs fonctions.

Options pour la commande docker pull. Sour*ce : Documentation Docker*

### docker run <image>

La commande `docker run`  crée et démarre un nouveau conteneur à partir d'une image spécifiée, c'est-à-dire qu'elle exécute l'image dans un nouveau conteneur. Il s'agit d'un raccourci pour `docker container run`, et les deux fonctionnent de la même manière.

- Voici la syntaxe de base : `docker container run [OPTIONS] IMAGE [COMMAND] [ARG...]`

Si vous avez déjà lancé un conteneur et que vous souhaitez le relancer avec toutes les modifications précédentes, utilisez : `docker start` .

La commande run propose de nombreuses options permettant de personnaliser l'exécution de votre conteneur. Nous allons passer en revue quelques-unes des plus courantes en les illustrant par des exemples :

| Flag | Exemple de commande | Description | 
| `-- name` | `docker run --name test -d nginx:alpine` | L'identifiant personnalisé spécifié pour le conteneur nommé test utilisant l'image `nginx:alpine` . | 
| `-w` ,`--workdir` | `docker run -w /path/to/dir/ -i -t ubuntu pwd` | Exécute la commande dans le répertoire spécifié, dans cet exemple, `/path/to/dir/` | 
| `--pid` | `docker run --rm -it --pid=host alpine` | Par défaut, l'espace de noms PID est activé dans tous les conteneurs, ce qui permet de séparer les processus. L'exemple utilise un conteneur alpin avec l'option `--pid = host` . | 
| `--cidfile` | `docker run --cidfile /tmp/docker_test.cid ubuntu echo "test"` | Ceci crée un conteneur et imprime un test sur la console L'option `cidfile` permet à Docker de créer un nouveau fichier et d'y inscrire l'identifiant du conteneur. | 

### docker stop <container> et docker start <container> :

La commande `docker start`  démarre un ou plusieurs conteneurs arrêtés. Par exemple, dans `docker start my_container`, `my_container` est le nom du conteneur que nous voulons démarrer. Nous pouvons également utiliser son alias : `docker container start`.

- Sa syntaxe complète est la suivante : `docker container start [OPTIONS] CONTAINER [CONTAINER...]`

De même, `docker stop` arrête un ou plusieurs conteneurs en cours d'exécution. Par exemple, dans `docker stop my_container`, `my_container` est le nom d'un conteneur en cours d'exécution.

Il a également un alias : `docker container stop`. 

- Sa syntaxe complète se présente comme suit : `docker container stop [OPTIONS] CONTAINER [CONTAINER...]`

Comme `start`, la commande `stop` comporte des options permettant de personnaliser la manière dont les conteneurs sont arrêtés. Nous allons maintenant vous en présenter quelques-uns :

Options pour `docker stop`. Source*: Documentation Docker*

## Maîtriser Docker et Kubernetes

## Travailler avec des images Docker

Les images sont la base de tout conteneur. Dans cette section, nous verrons comment créer, gérer et inspecter les images Docker à l'aide de ses commandes courantes.

### construction de docker

