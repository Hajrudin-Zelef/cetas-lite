---
id: collect-261001-huawei/huawei/les-26-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-concernant-d-2
title: "Step 1: Choose a base image"
domain: huawei
role: reference
task: reference
actors: ["AWS", "Google"]
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-huawei/les-26-questions-et-reponses-les-plus-frequentes-lors-d-entretiens-d-embauche-concernant-docker-pour.md
source_anchor: ""
source_lines: [122, 204]
sha256: 9af3b23ea5bffc7346a28cde1fa38d317c7c2d296e277f369e87d534aa91fb4d
---

# Step 1: Choose a base image

Nous utilisons les volumes Docker pour assurer la sécurité des données en dehors des conteneurs Docker. Ils fournissent un emplacement distinct sur les hôtes où les données sont conservées même si le conteneur est supprimé. De plus, il est plus facile de gérer, de sauvegarder et de partager les volumes entre les conteneurs.

### 9. Que sont les montages liés Docker et pourquoi privilégions-nous les volumes aux montages liés ?

Grâce aux montages liés Docker, il est possible de partager des fichiers entre la machine hôte et un conteneur. Ils relient un fichier spécifique sur le système hôte à un emplacement dans le conteneur. Si nous apportons des modifications aux fichiers, elles apparaîtront immédiatement dans le conteneur.

Les montages Docker sont adaptés au partage de fichiers en temps réel, mais ils dépendent du système d'exploitation hôte, ce qui soulève des questions de sécurité.

Au contraire, comme les volumes Docker fonctionnent de manière indépendante, ils sont plus sécurisés que les montages.

*Schéma des montages et volumes Docker Bind. Source de l'image : Docker*

### 10. Qu'est-ce que Docker Swarm ?

Docker Swarm est un outil d'orchestration de conteneurs qui gère et déploie des services sur un cluster de nœuds Docker. Il offre une haute disponibilité, une évolutivité et un équilibrage de charge, permettant à plusieurs hôtes d'agir comme un seul moteur Docker virtuel.

### 11. Est-il possible de mettre en place une mise à l'échelle automatique de Docker Swarm ?

Non, Docker Swarm ne prend pas en charge de manière native la mise à l'échelle automatique. Pour réaliser l'autoscaling, il est nécessaire d'intégrer des outils de surveillance et d'utiliser des scripts pour ajuster manuellement le nombre d'instances. Voici comment procéder :

- Veuillez installer un outil de surveillance, tel que Prometheus ou Grafana, afin de suivre l'utilisation des ressources, telles que le processeur et la mémoire.
- Veuillez définir les déclencheurs de mise à l'échelle. Par exemple, nous pouvons définir qu'une utilisation du processeur supérieure à 82 % déclenchera une augmentation de la capacité.
- Ensuite, veuillez rédiger un script à l'aide de la commande `docker service scale` pour ajuster le nombre de répliques. Par exemple, pour faire évoluer un service vers 5 répliques :`docker service scale =5`

En combinant des outils de surveillance, des déclencheurs et des scripts, il est possible de mettre en œuvre une forme d'autoscaling dans Docker Swarm, même si cette fonctionnalité n'est pas intégrée.

### 12. Comment utiliseriez-vous Docker Compose pour faire évoluer les services ?

Pour faire évoluer les services à l'aide de Docker Compose, nous pouvons utiliser le drapeau ` `--scale` ` avec la commande ` `docker-compose up` `. Ceci est généralement utilisé pour les services sans état tels que les serveurs web. Par exemple, pour faire évoluer un service Web vers 3 instances :

`docker-compose up --scale web=3`

Il est essentiel de s'assurer que le fichier `docker-compose.yml` définit correctement les services et utilise un équilibreur de charge externe ou prend en charge les instances à échelle variable. La mise à l'échelle des services avec état (par exemple, les bases de données) nécessite une configuration supplémentaire afin de garantir la cohérence des données.

### 13. Un conteneur peut-il redémarrer de lui-même ? Définissez ses politiques par défaut et permanentes.

Oui, un conteneur peut redémarrer de manière autonome. Cependant, il est nécessaire de définir une politique de redémarrage à cet effet.

Docker dispose de différentes politiques de redémarrage qui déterminent quand et comment les conteneurs doivent redémarrer. La politique par défaut est « non », ce qui signifie qu'un conteneur ne redémarrera pas s'il s'arrête. Avec la politique « always », Docker redémarrera automatiquement le conteneur chaque fois qu'il s'arrête.

Nous pouvons utiliser cette commande pour appliquer la politique « toujours » :

`docker run --restart=always` 

## Questions d'entretien avancées sur Docker

Passons maintenant aux questions d'entretien avancées sur Docker.

### 14. Veuillez expliquer le cycle de vie des conteneurs Docker.

Un conteneur Docker suit un cycle de vie qui définit les états dans lesquels il peut se trouver et son fonctionnement dans ces états. Les étapes du cycle de vie d'un conteneur Docker sont les suivantes :

- Créer : Dans cet état, nous configurons un conteneur à partir d'une image à l'aide de la commande `docker create` .
- **Exécuter** : Ici, nous utilisons la commande`docker start` pour exécuter le conteneur, qui effectue les tâches jusqu'à ce que nous l'arrêtions ou le mettions en pause.
- **Pause** : Nous utilisons la commande «`docker pause` » pour interrompre le processus. Cet état préserve la mémoire et le disque. Si vous souhaitez redémarrer le conteneur, veuillez utiliser la commande`docker unpause` .
- **Arrêt** : Si le conteneur est inactif, il entre en phase d'arrêt, mais cela peut se produire pour plusieurs raisons :
- **Arrêt immédiat** : La commande ``docker kill` ` arrête le conteneur sans nettoyage.
- **Achèvement du processus** : Une fois la tâche terminée, le conteneur s'arrête automatiquement.
- **Mémoire insuffisante** : Le conteneur s'arrête lorsqu'il utilise une quantité excessive de mémoire.
- **Supprimer** : Dans la dernière étape, nous supprimons le conteneur arrêté ou créé à l'aide de la commande`docker rm` .

### 15. Qu'est-ce qu'un référentiel d'images Docker ?

Un référentiel d'images Docker stocke et partage plusieurs images de conteneurs portant le même nom avec les clients ou la communauté. Nous pouvons les étiqueter à l'aide de balises afin de distinguer leurs différentes versions. Par exemple, `app/marketing_campaign:v1` sera la première version d'une application marketing, et `app/marketing_campaign:v2` sera la deuxième version.

Docker Hub, le référentiel d'images Docker le plus populaire, permet aux utilisateurs d'héberger, de partager et de récupérer des images de conteneurs de manière publique ou privée. D'autres alternatives incluent Amazon ECR, Google Artifact Registry et GitHub Container Registry.

### 16. Veuillez me présenter trois bonnes pratiques pour assurer la sécurité d'un conteneur Docker.

Afin de renforcer la sécurité des conteneurs et de minimiser les vulnérabilités courantes, je respecte les meilleures pratiques suivantes :

1. **Veuillez sélectionner des images légères** : Veuillez utiliser des images de base minimales telles qu'Alpine afin de réduire la surface d'attaque.
2. **Limiter les appels système** : Étant donné que les conteneurs Docker peuvent accéder à des appels non nécessaires, il est recommandé d'utiliser des outils tels que Seccomp pour limiter ces appels.
3. **Sécurisez les données sensibles** : Veuillez utiliser les secrets Docker pour gérer les clés API ou les mots de passe. Ils cryptent les secrets et ne les rendent accessibles que pendant l'exécution.

### 17. Pourquoi les conteneurs Docker nécessitent-ils des contrôles de santé ?

Les conteneurs Docker dépendent de contrôles de santé pour garantir leur bon fonctionnement. Le déploiement d'un conteneur qui est en cours d'exécution mais qui ne traite pas les demandes peut poser des problèmes aux équipes de déploiement. Les contrôles de santé surveillent ces problèmes en temps réel et nous informent immédiatement.

Par exemple, un contrôle de santé peut être ajouté dans un fichier Dockerfile comme suit :

`HEALTHCHECK --interval=30s --timeout=10s --retries=3 CMD curl -f http://localhost:8080/health || exit 1`

