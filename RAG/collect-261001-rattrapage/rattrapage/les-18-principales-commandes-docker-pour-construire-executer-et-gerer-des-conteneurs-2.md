---
id: collect-261001-rattrapage/rattrapage/les-18-principales-commandes-docker-pour-construire-executer-et-gerer-des-conteneurs-2
title: "syntax=docker/dockerfile:1"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/les-18-principales-commandes-docker-pour-construire-executer-et-gerer-des-conteneurs.md
source_anchor: ""
source_lines: [96, 232]
sha256: cc3fa6f23ca6f2a1a8dc0a96f47e8d999656a1e957d5298f7e43d9223c919df5
---

# syntax=docker/dockerfile:1

La commande `docker build` est l'une des fonctionnalités les plus utilisées de Docker. Bien qu'il fasse partie d'un écosystème plus large qui prend en charge des cas d'utilisation avancés, nous nous concentrerons sur la façon de l'utiliser pour construire une image à partir d'un simple fichier Docker. 

Un `Dockerfile` est un fichier texte brut (sans extension) qui contient des instructions étape par étape que Docker utilise pour construire une image. Voici comment en créer un :

1. Dans le répertoire racine de votre application, créez un fichier nommé `Dockerfile` avec le contenu suivant :

```
# syntax=docker/dockerfile:1
FROM node:lts-alpine
WORKDIR /app
COPY . .
RUN yarn install --production
CMD ["node", "src/index.js"]
EXPOSE 3000
```
Ce fichier Docker commence par une image de base légère qui inclut Node.js et Yarn. Il copie le code source de votre application dans l'image, installe les dépendances et définit comment démarrer l'application.

1. Maintenant, construisez l'image à l'aide de la commande suivante :

`docker build -t getting-started.`
Le drapeau `-t` vous permet de marquer l'image. Dans ce cas, nous l'avons nommé `getting-started.` Le `.` à la fin indique à Docker de rechercher le Dockerfile dans le répertoire actuel.

### images docker

La commande `docker images` répertorie toutes vos images de premier niveau, leur référentiel, leurs balises et leur taille. Vous pouvez également utiliser ses alias :

- `docker image list`
- `docker image ls`

Sa syntaxe est la suivante : `docker image ls [OPTIONS] [REPOSITORY[:TAG]]`

Cette commande comporte plusieurs options. Par exemple, pour afficher toutes les images, y compris les images intermédiaires, vous pouvez ajouter l'indicateur `-a` ou `--all` comme suit : `docker images -a`.

### docker rmi <image>

La commande `docker rmi`  supprime une ou plusieurs images de votre système. Si une image comporte plusieurs balises, l'exécution de cette commande avec une balise spécifique ne supprimera que cette balise. Mais si la balise est la seule liée à l'image, la balise et l'image seront toutes deux supprimées.

Vous pouvez également utiliser l'un de ces alias :

- `docker image remove`
- `docker image rm`

Sa syntaxe est la suivante : `docker image rm [OPTIONS] IMAGE [IMAGE...]`

Si vous devez supprimer une image encore utilisée par un conteneur en cours d'exécution, vous devrez la forcer en ajoutant l'option `-f` ou `--force`.

## Gestion des conteneurs Docker

Nous devons souvent gérer des conteneurs - les démarrer, les arrêter, les inspecter ou les supprimer au fur et à mesure de l'évolution de notre application. Je vais donc vous présenter les commandes Docker les plus utiles pour gérer les conteneurs dans votre travail quotidien.

### docker exec <conteneur> <commande>

La commande `docker exec` vous permet d'exécuter une commande à l'intérieur d'un conteneur en cours d'exécution sans le redémarrer. C'est particulièrement utile pour déboguer ou vérifier manuellement quelque chose à l'intérieur du conteneur. Vous pouvez également utiliser son alias : `docker container exec`.

La commande ne fonctionne que si le processus principal du conteneur (PID 1) est en cours d'exécution. Il ne s'exécute pas automatiquement si le conteneur redémarre.

Sa syntaxe est la suivante : `docker exec [OPTIONS] CONTAINER COMMAND [ARG...]`

Vous pouvez utiliser plusieurs drapeaux optionnels. Par exemple, `--privileged` donne à la commande des autorisations étendues à l'intérieur du conteneur. Pour obtenir la liste complète des optionsions, consultez les documents officiels.

Voici un exemple d'exécution d'une commande sur le conteneur :

`docker exec -d mycontainer touch /tmp/execWorks`
La commande `touch` crée un nouveau fichier `/tmp/execWorks` dans le conteneur en cours d'exécution `mycontainer`, en arrière-plan.

### docker logs <container>

La commande `docker logs` vous permet de consulter les journaux d'un conteneur spécifique et d'afficher tout ce qui est imprimé sur la sortie standard et les erreurs au moment de l'exécution. Vous pouvez également utiliser son alias : `docker container logs`

Sa syntaxe est la suivante : `docker container logs [OPTIONS] CONTAINER`

Vous pouvez ajouter un certain nombre d'options utiles. Par exemple :

- `--details` affiche des attributs supplémentaires tels que les variables d'environnement et les étiquettes.
- `--until` vous permet de récupérer les journaux jusqu'à un moment précis.
- `docker logs -f --until=2s test` suit la sortie des journaux du conteneur`test` et s'arrête après avoir affiché les deux dernières secondes de journaux.

Voici toutes les options que nous pouvons utiliser avec `docker logs`:

Options pour `docker log` . Donc*urce : Documentation Docker*

### docker rm <container>

La commande `docker rm`  supprime un ou plusieurs conteneurs de votre système. Vous pouvez également utiliser ses alias :

- `docker container remove`
- `docker container rm`

Sa syntaxe est la suivante : `docker container rm [OPTIONS] CONTAINER [CONTAINER...]`

Cette commande comporte quelques options que vous pouvez utiliser. En voici quelques-unes :

Options pour `docker rm` . Source*e : Documentation Docker*

Par exemple, vous pouvez utiliser `docker rm /redis` pour supprimer le conteneur identifié par le lien `/redis`. Notez toutefois que cette commande ne supprime que les conteneurs en cours d'exécution. 

Si vous souhaitez supprimer les conteneurs arrêtés, vous devez utiliser `docker container prune`. Pour garder votre environnement propre, apprenez à supprimer en toute sécurité les ressources Docker inutilisées grâce à ce tutoriel Docker prune.

### docker restart <container>

La commande `docker restart`  arrête puis redémarre un ou plusieurs conteneurs. Vous pouvez également utiliser son alias : `docker container restart`. 

Sa syntaxe est la suivante : `docker restart [OPTIONS] CONTAINER [CONTAINER...]`

La commande restart dispose de plusieurs options utiles pour personnaliser le redémarrage des conteneurs. Vous trouverez ci-dessous les deux options, accompagnées d'exemples.

| Option | Description | Exemple | 
| -s, --signal | Signal à envoyer au conteneur. | `docker restart -s SIGTERM mycontainer` | 
| -t, --timeout | Nombre de secondes à attendre avant de tuer le conteneur. | `docker restart -t 10 mycontainer` | 

> Pour une pratique concrète, explorez ces idées de projets Docker.*ces idées de projets Docker qui* qui vont du niveau débutant au niveau avancé.

## Réseau Docker

La mise en réseau des conteneurs permet la communication entre les conteneurs et les charges de travail externes. Par défaut, les conteneurs disposent d'un réseau activé et peuvent établir des connexions sortantes, mais ils ne savent pas automatiquement sur quel type de réseau ils se trouvent ni à quelles autres charges de travail ils sont connectés.

À moins que vous n'utilisiez le pilote réseau `none` (qui désactive le réseau), les conteneurs peuvent interagir avec des éléments de réseau tels que les adresses IP, les passerelles et les DNS. 

Explorons quelques commandes courantes de mise en réseau de dockers que vous pouvez utiliser.

### docker network ls

La commande `docker network ls` répertorie tous les réseaux connus du moteur Docker, y compris ceux qui se trouvent sur plusieurs hôtes d'un cluster. Son alias est : `docker network list`.

- Voici sa syntaxe : `docker network ls [OPTIONS]`

Vous pouvez utiliser plusieurs options, comme le montre l'image ci-dessous :

Par défaut, cette commande affiche le nom de chaque réseau :

- ID
- Nom
- Conducteur
- Champ d'application

Vous pouvez utiliser des drapeaux pour personnaliser la sortie. Par exemple, le drapeau `--no-trunc` affiche les identifiants complets du réseau au lieu des identifiants abrégés. Voici comment utiliser cette commande :

