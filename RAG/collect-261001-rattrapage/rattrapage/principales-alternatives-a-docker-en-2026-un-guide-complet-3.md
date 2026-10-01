---
id: collect-261001-rattrapage/rattrapage/principales-alternatives-a-docker-en-2026-un-guide-complet-3
title: "Buildah scripting approach with CI integration"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["gpu"]
source: docs/RAG/collect-261001-rattrapage/principales-alternatives-a-docker-en-2026-un-guide-complet.md
source_anchor: ""
source_lines: [133, 190]
sha256: 9ece0b5a28c270f14329ebbd732ca44dda4a066324e906ea519baa1f807d642f
---

# Buildah scripting approach with CI integration

runC sert d'implémentation de référence de l' pour la spécification OCI Runtime. Il s'agit d'un exemple standard illustrant la manière dont les conteneurs devraient fonctionner. La plupart des plateformes de conteneurs utilisent runC comme moteur d'exécution, notamment Docker, containerd, CRI-O et Podman. Lorsque vous démarrez un conteneur à l'aide de l'un de ces outils, runC est susceptible de gérer la création et l'isolation réelles du processus.

Le runtime implémente les primitives de conteneur de base : création d'espaces de noms Linux pour l'isolation, configuration de cgroups pour les limites de ressources et configuration des contextes de sécurité. Il est écrit en Go et conçu pour être simple, fiable et conforme aux spécifications plutôt que riche en fonctionnalités.

runC est particulièrement performant dans les systèmes embarqués et les piles de conteneurs personnalisées où un comportement prévisible et des dépendances minimales sont requis. Comme il ne gère que l'exécution des conteneurs, il est possible de créer des plateformes de conteneurs spécialisées autour de lui sans hériter d'une complexité inutile. Les appareils IoT, les plateformes informatiques de pointe et les systèmes d'orchestration personnalisés utilisent souvent runC directement plutôt que des environnements d'exécution de niveau supérieur.

Cet outil fonctionne comme un utilitaire en ligne de commande qui lit les spécifications des paquets OCI et crée des conteneurs en conséquence. Cela le rend idéal pour l'intégration dans des systèmes existants ou pour la création d'outils personnalisés de gestion des conteneurs.

### Youki : Performance basée sur Rust

Youki réimplémentela spécification OCI Runtime dans Rust, en mettant l'accent sur la sécurité de la mémoire et les performances. L'implémentation Rust élimine des catégories entières de vulnérabilités de sécurité pouvant affecter les environnements d'exécution C et Go, tout en améliorant les temps de démarrage des conteneurs grâce à une meilleure gestion de la mémoire et à une réduction de la surcharge.

Les tests de performance indiquent que Youki démarre les conteneurs plus rapidement que runC dans de nombreux cas, bien que l'amélioration exacte varie en fonction de la charge de travail. Cette amélioration est due aux abstractions sans coût et aux modèles d'allocation de mémoire plus efficaces de Rust. Pour les applications qui créent et détruisent de nombreux conteneurs à courte durée de vie, ces améliorations du temps de démarrage peuvent être significatives.

Youki est entièrement compatible avec la spécification OCI Runtime, ce qui lui permet de remplacer runC dans la plupart des plateformes de conteneurs. Docker Engine, containerd et d'autres environnements d'exécution de haut niveau peuvent utiliser Youki sans modification de leur configuration ou de leurs API.

Le runtime offre des avantages pour les charges de travail critiques en termes de performances, telles que les fonctions sans serveur, les pipelines CI/CD avec de nombreuses constructions de conteneurs et les architectures de microservices avec des événements de mise à l'échelle fréquents. Dans ces scénarios, des temps de démarrage plus rapides se traduisent directement par une réduction de la latence au démarrage à froid et une meilleure utilisation des ressources.

Youki inclut également des fonctionnalités telles que l'optimisation cgroup v2 et une prise en charge améliorée des conteneurs rootless qui exploitent le système de types de Rust pour éviter les erreurs de configuration lors de la compilation.

## Conteneurs système : LXC/LXD par rapport aux conteneurs d'applications

La conteneurisation ne doit pas nécessairement se concentrer sur des applications individuelles ; il est parfois nécessaire d'exécuter des systèmes d'exploitation complets dans des conteneurs. LXC et LXD offrent une conteneurisation au niveau du système qui diffère fondamentalement de l'approche axée sur les applications de Docker.

### LXC : Virtualisation au niveau du système d'exploitation

LXC (Linux Containers)crée des conteneurs qui se comportent comme des systèmes Linux complets plutôt que comme des processus d'application isolés.

Image 5 - Page d'accueil de LXC

Chaque conteneur LXC exécute son propre système d'initialisation, peut héberger plusieurs services et fournit un environnement utilisateur complet qui est pratiquement impossible à distinguer d'une machine virtuelle.

Cette approche est particulièrement efficace pour les charges de travail existantes qui n'ont pas été conçues pour la conteneurisation. Les applications qui prévoient d'écrire dans le répertoire d'accueil ( `/etc`), d'exécuter des services système ou d'interagir avec l'intégralité de la hiérarchie du système de fichiers fonctionnent de manière transparente dans les conteneurs LXC. Il est possible de transférer l'intégralité des configurations de serveurs vers LXC sans avoir à refactoriser les applications pour les architectures de microservices.

Les conteneurs LXC partagent le noyau hôte, mais offrent une isolation plus forte que les conteneurs d'applications. Chaque conteneur dispose de sa propre pile réseau, de son propre arbre de processus et de son propre espace de noms de système de fichiers, créant ainsi une isolation similaire à celle d'une machine virtuelle sans la surcharge liée à la virtualisation matérielle.

LXD, désormais disponible sursous Canonical, ajoute une couche de gestion puissante au-dessus de LXC. LXD fournit des API REST, la gestion des images et des fonctionnalités avancées telles que la migration en direct entre les hôtes. Vous pouvez déplacer des conteneurs en cours d'exécution d'une machine physique à une autre sans interruption de service, de manière similaire à VMware vMotion, mais avec des conteneurs.

Les capacités de transfert matériel permettent aux conteneurs LXD d'accéder directement aux GPU, aux périphériques USB et à d'autres matériels. Cela le rend adapté aux charges de travail qui nécessitent un accès matériel spécialisé tout en conservant les avantages des conteneurs, tels que la densité et le provisionnement rapide.

LXD prend également en charge la mise en cluster, ce qui vous permet de gérer plusieurs hôtes comme une seule unité logique avec des capacités de placement et de basculement automatisés.

### Analyse comparative

Voici une comparaison entre les conteneurs système et les conteneurs d'application en fonction de leurs principales caractéristiques :

Image 6 - Aperçu des principales caractéristiques de Docker, LXC et des machines virtuelles

Les conteneurs système comblent le fossé entre les conteneurs d'applications légers et les machines virtuelles lourdes. Ils sont particulièrement adaptés lorsque vous avez besoin de fonctionnalités similaires à celles d'une machine virtuelle avec l'efficacité d'un conteneur, ou lorsque vous migrez des applications existantes qui ne peuvent pas être facilement décomposées en microservices.

Le choix entre les conteneurs système et les conteneurs d'application dépend davantage des caractéristiques de votre charge de travail et de vos exigences opérationnelles que de leur supériorité technique. Ils répondent à des besoins différents dans le domaine de la conteneurisation.

## Architectures de sécurité dans les environnements d'exécution modernes

La sécurité est passée d'une considération secondaire à un principe de conception fondamental dans les environnements d'exécution de conteneurs modernes. Les plateformes actuelles mettent en œuvre des stratégies de défense en profondeur qui partent du principe que les conteneurs seront compromis et se concentrent sur la limitation du rayon d'action.

### Fonctionnement sans racine

