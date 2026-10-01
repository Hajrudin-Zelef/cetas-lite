---
id: collect-261001-rattrapage/rattrapage/docker-vs-podman-quel-outil-de-conteneurisation-vous-convient-le-mieux-2
title: "docker-vs-podman-quel-outil-de-conteneurisation-vous-convient-le-mieux"
domain: rattrapage
role: reference
task: reference
actors: ["AWS", "Google"]
dates: []
keywords: ["arr", "aws"]
source: docs/RAG/collect-261001-rattrapage/docker-vs-podman-quel-outil-de-conteneurisation-vous-convient-le-mieux.md
source_anchor: ""
source_lines: [71, 158]
sha256: fac70531d63041db3058e054e081dfb82f5bb87c2b26bab77e254dfef588a451
---

# docker-vs-podman-quel-outil-de-conteneurisation-vous-convient-le-mieux

Il existe plusieurs façons de travailler avec Docker. Les praticiens expérimentés des logiciels et des données s'appuient généralement sur le CLI Docker (AKA le client Docker) pour interagir avec leurs images et leurs conteneurs Docker.

Cependant, il existe un moyen encore plus simple de commencer, et c'est Docker Desktop.

Docker Desktop est un outil gratuit basé sur une interface graphique qui permet aux utilisateurs de créer et de gérer les images et les conteneurs qui exécutent leurs applications ou leurs charges de travail. Un ingénieur des données peut utiliser Docker Desktop pour afficher les images disponibles sur sa machine et transformer cette image en conteneur. De même, un développeur de logiciels peut télécharger une image à partir de Docker Hub pour l'utiliser dans le cadre de son prochain projet.

L'interface utilisateur est simple et intuitive, tout en conservant une visibilité et un contrôle complets sur votre environnement Docker.

Cependant, Docker Desktop ne se limite pas à l'affichage et à la gestion des objets Docker.

Les utilisateurs peuvent notamment gérer (à l'octet près) les ressources disponibles pour leurs objets Docker, s'attacher à un conteneur en cours d'exécution ou lancer un cluster Kubernetes sur leur machine locale. Les utilisateurs de Docker Desktop peuvent choisir parmi des centaines d'extensions, ou commencer à utiliser Docker grâce à des tutoriels utiles et des exemples d'environnements. Heureusement pour vous, Docker Desktop est largement accessible et fonctionne sur Mac, Windows ou Linux.

### Caractéristiques et limitations de Podman Desktop

Podman Desktop ressemble beaucoup à son homologue Docker. À partir de l'interface utilisateur de Podman Desktop, les utilisateurs peuvent afficher et gérer des conteneurs, des images, des pods et des volumes. Comme pour Docker, Podman prend en charge des plugins et des intégrations permettant de faire fonctionner localement un cluster Red Hat OpenShift ou de travailler avec des LLM en utilisant le Podman AI Lab.

Si vous souhaitez utiliser un plugin personnalisé, vous pouvez l'installer à partir de Podman Desktop.

Si vous êtes à la fois un utilisateur de Docker et de Podman, vous serez peut-être surpris de voir des objets Docker et Podman dans l'interface utilisateur de Podman Desktop. Ce n'est pas une coïncidence ! Nous examinerons plus en détail ce que cela implique. En attendant, cela signifie que les utilisateurs peuvent interagir à la fois avec leurs objets Podman ET Docker, le tout via une seule vitre.

Le cas d'utilisation le plus courant des conteneurs consiste à les faire fonctionner via Kubernetes. Malgré le titre de norme industrielle de Docker pour la conteneurisation, Podman offre une expérience Kubernetes plus robuste sur Podman Desktop.

La possibilité d'afficher et de gérer les ressources Kubernetes comme les nœuds, les pods, les déploiements (et bien plus encore) fait de l'administration et du développement de Kubernetes un citoyen de première classe sur Podman Desktop. Ces outils, ainsi que des plugins comme l'intégration Red Hat OpenShift mentionnée plus haut, différencient Podman en tant qu'outil orienté vers les ateliers Kubernetes.

## Podman Compose vs Docker Compose

### Définir et gérer des applications multi-conteneurs avec Docker

Certaines applications et charges de travail peuvent être regroupées dans un seul conteneur. Certains ne le peuvent pas. Pour faciliter la gestion de plusieurs conteneurs, Docker propose un outil appelé Docker Compose. Docker Compose utilise un seul fichier YAML pour définir les composants de votre application.

Ensuite, à l'aide de la CLIdocker-compose, ces conteneurs et services peuvent être démarrés, arrêtés ou reconstruits. Un fichier YAML de Docker Compose peut ressembler à ceci :

```
yaml
version: '3'
services:
	app:
		image: python:3.10
		container_name: app
		command: run app --host=0.0.0.0
	database:
		image: postgres:13
		container_name: database
		ports: 5432
		volumes:
			- postgres_data:/var/lib/postgresql/data
volumes:
	postgres_data
```
Il se passe beaucoup de choses, mais ce que Docker Compose nous permet de faire, c'est de définir un fichier YAML avec deux services et un volume. Ensuite, la commande docker-compose up lancera ces conteneurs et nous aurons une application en cours d'exécution .

Pour les équipes chargées des logiciels et des données qui gèrent des applications et des charges de travail volumineuses, Docker Compose facilite le développement local, ainsi que l'expédition et l'exécution du code dans un environnement de production.

### L'approche de Podman pour les applications multi-conteneurs

L'exécution d'applications multi-conteneurs avec Podman ressemble à s'y méprendre à celle de Docker. Podman le fait en utilisant Podman Compose. Comme pour Docker Compose, Podman Compose utilise des fichiers YAML pour définir les composants d'une application de manière déclarative.

Le podman-compose peut alors être utilisé pour démarrer, arrêter ou redémarrer les services définis dans le fichier YAML.

Dans la plupart des cas, podman-compose peut être utilisé à la place de docker-compose (il y a quelques incompatibilités ici et là). Comme avec Docker, l'utilisation de Podman Compose permet de gérer des applications multi-conteneurs de manière indépendante et flexible.

Vous trouverez ci-dessous un tableau comparant Docker et Podman.

| **Caractéristique/aspect** | **Docker** | **Podman** | 
| **Architecture** | Docker s'appuie sur un démon en tant que composant architectural de base. | Architecture sans démon. | 
| **Sécurité** | Nécessite les privilèges de root pour construire, exécuter et gérer les conteneurs. | La nature sans démon de l'architecture de Podman en fait un outil de gestion de conteneurs plus respectueux de la sécurité. | 
| **Outil de l'utilisateur** | Docker Desktop, docker CLI | Podman Desktop, podman CLI | 
| **Compatibilité** | Windows, Mac, Linux | Native de Linux, disponible pour Windows et Mac. | 
| **Adoption** | Standard industriel pour l'orchestration de conteneurs avec une communauté massive et une compatibilité quasi universelle. | Alternative à Docker avec une communauté plus petite, mais en pleine croissance. | 

## Cas d'utilisation et meilleurs scénarios pour Podman vs Docker

Examinons maintenant la question clé que vous vous posez peut-être : quand utiliser Docker et quand utiliser Podman ? Regardons de plus près.

### Quand utiliser Docker ?

Docker est la norme de facto pour la construction, l'exécution et l'expédition de conteneurs. Si vous débutez dans la conteneurisation (surtout sur votre machine personnelle), essayez d'utiliser Docker.

Il est facile de mettre en place votre premier (ou cinquantième) conteneur et de le faire fonctionner à l'aide d'outils tels que Docker Desktop ou Docker CLI. Docker dispose d'une communauté massive, et il y a de fortes chances que ce que vous essayez de faire ait déjà été fait. Cela permet de faciliter les opérations de dépannage.

Docker offre une plus grande cohérence multiplateforme que Podman. Plus important encore, Docker s'intègre à presque tous les services basés sur des conteneurs, notamment AWS ECS, Azure AKS et Google Cloud Run.

Cela signifie que lorsque le moment est venu d'exécuter vos conteneurs en production, vous êtes en mesure d'intégrer facilement le service de votre choix. La possibilité de passer du développement local à la production est l'un des aspects les plus puissants de la conteneurisation de votre code avec Docker.

Les équipes d'ingénierie logicielle et de données ne sont pas les seules à utiliser Docker. Ingénieurs en IA et ML, scientifiques des donnéeset même les analystes de données utilisent Docker pour améliorer leur travail !

### Quand utiliser Podman

