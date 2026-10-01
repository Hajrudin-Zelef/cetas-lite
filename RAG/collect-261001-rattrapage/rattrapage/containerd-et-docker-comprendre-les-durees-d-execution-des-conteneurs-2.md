---
id: collect-261001-rattrapage/rattrapage/containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs-2
title: "containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "arr"]
source: docs/RAG/collect-261001-rattrapage/containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs.md
source_anchor: ""
source_lines: [68, 137]
sha256: 0a7971413591fa3007800fcd6495f5bd486b919c659bff6ec072b598f684c92c
---

# containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs

- **Couche API gRPC :** Fournit un modèle client-serveur dans lequel plusieurs clients (Docker, Kubernetes, outils personnalisés) communiquent avec une seule instance containerd.
- **Démon Containerd :** Gère l'état des conteneurs en cours d'exécution, administre le stockage des images via des snapshotters et coordonne la mise en réseau via des plugins.
- **runc runtime :** Un environnement d'exécution léger et conforme à l'OCI qui s'interface directement avec les fonctionnalités du noyau Linux, créant des espaces de noms, configurant des cgroups et lançant des processus de conteneurs.
- **Plugins modulaires :** Des snapshotters personnalisés pour le stockage spécialisé, des environnements d'exécution alternatifs (gVisor, Kata Containers) et des plugins réseau peuvent être intégrés sans modification de containerd.

Lorsque containerd doit démarrer un conteneur, il lance runc, qui effectue le travail réel consistant à créer des espaces de noms isolés, à configurer des cgroups pour les limites de ressources et à lancer le processus du conteneur.

Cette séparation des préoccupations rend containerd hautement extensible. Les organisations peuvent personnaliser presque tous les aspects sans modifier le code conteneur principal.

## Containerd et Docker : principales différences

Maintenant que nous avons une bonne compréhension du fonctionnement de chaque outil, examinons leurs différences dans la pratique, de l'architecture aux modèles d'intégration.

### Exécution de conteneurs par rapport à la plateforme

La différence fondamentale réside dans la portée et l'objectif.

Docker utilise une architecture démon centralisée dans laquelle `dockerd` coordonne de nombreux aspects différents, tels que :

- Construction d'image
- Exécution de conteneurs
- Réseautage
- Gestion du volume

Cette conception tout-en-un simplifie l'expérience des développeurs, mais introduit une surcharge due à la multiplicité des couches d'abstraction.

Containerd, en revanche, se concentre exclusivement sur l'exécution des conteneurs. Il n'inclut pas de fonctionnalités de création d'images, d'orchestration ou d'interface graphique. Pour la mise en réseau, Docker intègre une interface réseau de conteneur ( `Libnetwork` ) dans son démon, tandis que containerd s'appuie sur des plugins externes d'interface réseau de conteneur (CNI) qui peuvent être remplacés en fonction des besoins.

Cette différence architecturale a des implications en termes de performances. Dans les environnements à forte rotation où les conteneurs démarrent et s'arrêtent fréquemment, tels que les clusters Kubernetes à dimensionnement automatique, la conception rationalisée de containerd peut se traduire par des temps de démarrage plus rapides et une consommation de ressources réduite. La réduction des frais généraux signifie que containerd utilise moins de mémoire et de CPU, ce qui devient significatif à grande échelle.

### Intégration Kubernetes

En ce qui concerne Kubernetes, la relation entre les environnements d'exécution de conteneurs et les plateformes d'orchestration a considérablement évolué, notamment en ce qui concerne la manière dont Kubernetes se connecte à containerd.

La relation entre Kubernetes et les environnements d'exécution de conteneurs a connu un changement majeur avec la version 1.24 de Kubernetes, publiée en 2022. Cette version a supprimé « Dockershim », une couche de compatibilité qui permettait à Kubernetes d'utiliser Docker comme environnement d'exécution de conteneurs.

Dockershim a toujours été conçu comme une solution temporaire. Lorsque Kubernetes a introduit l'interface CRI (Container Runtime Interface) afin de normaliser la communication avec les environnements d'exécution, Docker n'a pas pu l'implémenter directement, car Docker était antérieur à la conception de CRI. Dockershim assure la traduction entre Kubernetes et Docker, ajoutant ainsi une couche de traduction qui pourrait s'avérer superflue.

Les versions modernes de Kubernetes communiquent directement avec containerd via CRI, éliminant ainsi complètement la couche de traduction Docker. Cette simplification apporte des avantages concrets :

- **Réduction de la latence :** La communication directe élimine les frais de traduction dans les opérations de conteneurs.
- **Stabilité améliorée :** Moins de pièces mobiles signifie moins de points de défaillance potentiels.
- **Meilleures performances :** La pile d'exécution rationalisée est particulièrement avantageuse pour les déploiements à grande échelle.
- **Débogage simplifié :** L'intégration directe du CRI simplifie le dépannage.

Pour les utilisateurs de Kubernetes, ce changement est largement transparent. Les images Docker restent entièrement compatibles car elles respectent les normes OCI. Concrètement, cela signifie que les clusters Kubernetes de production fonctionnent désormais plus efficacement en utilisant directement containerd, tandis que les développeurs peuvent continuer à utiliser Docker localement pour la construction et les tests.

Si vous n'êtes pas certain des avantages et des inconvénients liés à l'utilisation de Kubernetes, veuillez consulter cette comparaison entre Docker Compose et Kubernetes.

### Construction et gestion de l'image

Bien que l'intégration à l'exécution soit essentielle pour l'orchestration, le flux de travail des développeurs dépend fortement de la manière dont chaque outil gère la création et le stockage des images. La création d'images représente un écart de capacité significatif entre Docker et containerd.

Docker fournit un système de compilation intégré via Dockerfiles et BuildKit, permettant aux développeurs de créer des compilations complexes en plusieurs étapes avec mise en cache, parallélisation et fonctionnalités avancées telles que les secrets de compilation et le transfert d'agent SSH.

Conformément à sa conception, Containerd n'inclut aucun workflow natif de création d'images. Pour créer des images avec containerd, les développeurs doivent utiliser des outils externes. Les options incluent l'exécution de BuildKit en tant que démon distinct et l'utilisation de `buildctl` pour les compilations en ligne de commande, ou l'adoption de `nerdctl`, une interface CLI compatible avec Docker qui intègre BuildKit.

Les mécanismes de stockage diffèrent également dans leur approche. La gestion des volumes de Docker offre une abstraction qui semble naturelle pour les développeurs, avec des volumes nommés qui conservent les données indépendamment du cycle de vie des conteneurs. Containerd utilise un système de capture d'écran de niveau inférieur, dans lequel différents pilotes de capture d'écran peuvent être connectés pour gérer différemment les systèmes de fichiers en couches en fonction des exigences de stockage sous-jacentes.

Cette différence reflète les publics visés par ces outils :

- Docker optimise la productivité des développeurs grâce à des fonctionnalités intégrées pratiques.
- Containerd fournit des primitives flexibles que les développeurs de plateformes peuvent assembler en fonction de leurs besoins spécifiques.

### CLI, expérience développeur et nerdctl

Au-delà de l'architecture et des fonctionnalités, l'expérience quotidienne des développeurs est directement influencée par l'interface de ligne de commande fournie par chaque outil. L'expérience de la ligne de commande met clairement en évidence les différentes philosophies de conception.

L'interface CLI de Docker est réputée pour sa convivialité. Les commandes telles que `docker run`, `docker build` et `docker logs` sont intuitives, bien documentées et conçues pour les utilisateurs. L'interface CLI comprend des paramètres par défaut utiles, des messages d'erreur clairs et des options étendues qui couvrent la plupart des cas d'utilisation.

