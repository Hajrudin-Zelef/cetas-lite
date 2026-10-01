---
id: collect-261001-rattrapage/rattrapage/les-18-principales-commandes-docker-pour-construire-executer-et-gerer-des-conteneurs-3
title: "syntax=docker/dockerfile:1"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/les-18-principales-commandes-docker-pour-construire-executer-et-gerer-des-conteneurs.md
source_anchor: ""
source_lines: [233, 365]
sha256: b34d0830332e04d40bfb0e5aaf7703d731ebf1bb0ce0afffee8a55ee5986f228
---

# syntax=docker/dockerfile:1

`docker network ls --no-trunc`
### docker network create <nom du réseau>

La commande `docker network create` crée un nouveau réseau Docker. Par défaut, il utilise le pilote `bridge`, à moins que vous n'en spécifiiez un autre à l'aide de l'indicateur `--driver` (ou `-d`).

Docker prend en charge les pilotes réseau intégrés tels que :

- `bridge` pour les réseaux à hôte unique.
- `overlay` pour les réseaux multi-hôtes en mode essaim.

Vous pouvez également utiliser des pilotes tiers ou personnalisés si nécessaire.

- Voici la syntaxe de base : `docker network create [OPTIONS] NETWORK`

La commande offre de nombreuses options pour différents objectifs. Consultez la documentation officielle pour voir la liste complète des options.

Voici un exemple de création d'un réseau de passerelles :

`docker network create -d bridge my-bridge-network`
Les réseaux Bridge sont limités à un seul moteur Docker, de sorte qu'ils ne connectent pas les conteneurs sur différents hôtes.

Une fois le mode Swarm activé, vous pouvez créer un réseau qui s'étend sur plusieurs hôtes Docker :

`docker network create --scope=swarm --attachable -d overlay my-multihost-network`
> Vous êtes curieux de savoir comment Docker se compare à Kubernetes ? Cette analyse *de Kubernetes vs Docker*couvre les principales différences et les cas d'utilisation.

## Volumes Docker

Les volumes Docker sont utilisés pour stocker des données qui doivent persister, même lorsqu'un conteneur s'arrête ou est supprimé. Vous pouvez les créer explicitement ou laisser Docker les créer automatiquement lors du démarrage d'un conteneur.

Les volumes sont stockés sur le système hôte mais sont isolés des fichiers principaux de l'hôte. Ils sont montés dans des conteneurs d'une manière similaire aux montages bind, mais avec une meilleure portabilité et une meilleure sécurité.

Explorons quelques commandes de volume docker pertinentes :

### docker volume ls

La commande `docker volume ls` répertorie tous les volumes connus de Docker. Vous pouvez également utiliser son alias : `docker volume list`.

- Sa syntaxe est la suivante : `docker volume ls [OPTIONS]`

Cette commande prend en charge quelques options facultatives pour vous aider à filtrer ou à formater la sortie - voici ce qu'elles sont :

Options pour la commande `docker volume ls`. Sou*rce : Documentation Docker*

### docker volume create <nom du volume>

La commande `docker volume create` crée un nouveau volume pour le stockage des données persistantes. Si vous ne fournissez pas de nom, Docker en génère un pour vous automatiquement.

La création de volumes est une étape courante lorsque vous souhaitez que les données persistent au-delà de la durée de vie d'un seul conteneur.

- Voici la syntaxe : `docker volume create [OPTIONS] [VOLUME_NAME]`

Voyons un exemple de création d'un volume et de configuration d'un conteneur pour l'utiliser :

```
docker volume create hello
docker run -d -v hello:/world busybox ls /world
```
Dans cet exemple, un volume appelé `hello` est créé. Il est ensuite monté dans un conteneur sur le site `/world`. Cela permet au conteneur d'écrire ou de lire des données sur ce volume.

Plusieurs conteneurs peuvent utiliser le même volume, ce qui est utile si un conteneur doit écrire des données pendant qu'un autre les lit.

Note : Les noms de volumes doivent être uniques d'un pilote à l'autre. Vous ne pouvez pas utiliser le même nom de volume dans deux pilotes de stockage différents.

## Commandes Docker Compose

Docker Compose facilite la gestion des applications multi-conteneurs à l'aide d' un simple fichier YAML. Il supporte différents environnements tels que le développement, les tests, la mise en scène, la production et l'analyse critique. En une seule commande, vous pouvez contrôler les services, configurer les réseaux et gérer les volumes, le tout en un seul endroit.

Examinons quelques commandes de composition courantes :

### docker-compose up

La commande `docker compose up` construit, (re)crée, démarre et attache des conteneurs pour un service. Si les conteneurs ne sont pas encore en cours d'exécution, il démarre également tous les services liés automatiquement.

- Voici la syntaxe : `docker compose up [OPTIONS] [SERVICE...]`

Par défaut, cette commande combine les résultats de tous les conteneurs. Si vous souhaitez vous concentrer sur des services spécifiques, vous pouvez le faire :

- Utilisez le drapeau `--attach` pour vous attacher à certains services.
- Utilisez l'option `--no-attach` pour exclure d'autres personnes

Par exemple, `docker compose up --no-attach`  démarre tous les services à l'exception de celui que vous avez exclu des journaux.

Lorsque la commande se termine, les conteneurs s'arrêtent également. Pour qu'ils continuent à fonctionner en arrière-plan, utilisez l'option `--detach`: 

`docker compose up --detach`
Avant d'exécuter cette commande, assurez-vous d'avoir navigué (`cd`) jusqu'au répertoire où se trouve votre fichier `docker-compose.yml`.

### docker-compose down

La commande `docker compose down` arrête les conteneurs et supprime les conteneurs, les images, les réseaux et les volumes créés par `docker compose up`. 

- Sa syntaxe est la suivante : `docker compose down [OPTIONS] [SERVICES]`

Vous pouvez utiliser différentes options, dont les suivantes :

Options pour la commande `docker compose down`. Source : Docker docs.

Par défaut, la commande supprime les éléments suivants :

- Conteneurs pour les services définis dans le fichier Compose.
- Réseaux définis dans la section des réseaux du fichier Compose.
- Le réseau par défaut, s'il y en a un.

Les éléments suivants ne sont pas supprimés par défaut:

- Réseaux et volumes définis comme externes.
- Les volumes anonymes, c'est-à-dire les volumes qui n'ont pas de nom.

Les volumes anonymes ne sont pas montés automatiquement lorsque vous exécutez à nouveau `docker compose up`, car ils n'ont pas de nom. Si vous avez besoin d'un stockage de données persistant, il est préférable d'utiliser des volumes nommés ou des montages bind.

## Meilleures pratiques pour l'utilisation des commandes Docker

Docker est un moyen puissant de créer, d'expédier et d'exécuter des applications dans des conteneurs. Mais pour l'utiliser à bon escient et faire en sorte que votre installation soit efficace et facilement extensible, vous devez suivre certaines bonnes pratiques.

### Utilisez les volumes Docker pour les données persistantes

Par défaut, les fichiers du conteneur sont stockés dans une couche inscriptible qui est perdue lorsque le conteneur est supprimé ! Cette couche est propre à chaque conteneur et n'est pas facilement accessible. Docker utilise différents types de montages pour conserver les données, l'un desch étant les volumes, gérés par le démon Docker et stockés sur l'hôte.

Les volumes nous le permettent :

- Conservez les données même après la suppression d'un conteneur.
- Stockez les données critiques pour les performances avec une vitesse de niveau hôte.
- Gérez facilement le stockage grâce à Docker.

Ils sont idéaux pour les données à long terme ou lorsque plusieurs conteneurs doivent partager l'accès. Gardez à l'esprit que si vous avez besoin d'accéder à des fichiers directement à partir de l'hôte, les montages bind peuvent être mieux adaptés, car les volumes sont entièrement gérés par Docker.

### Automatiser avec Docker Compose

La gestion manuelle de plusieurs conteneurs peut s'avérer fastidieuse. Docker Compose simplifie les choses en vous permettant de tout définir dans un seul fichier YAML afin que vous puissiez vous concentrer sur la construction.

Voici pourquoi il s'agit d'une bonne pratique à suivre :

