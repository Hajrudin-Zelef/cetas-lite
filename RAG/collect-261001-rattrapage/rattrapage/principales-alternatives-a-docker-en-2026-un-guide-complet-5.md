---
id: collect-261001-rattrapage/rattrapage/principales-alternatives-a-docker-en-2026-un-guide-complet-5
title: "Buildah scripting approach with CI integration"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-rattrapage/principales-alternatives-a-docker-en-2026-un-guide-complet.md
source_anchor: ""
source_lines: [237, 305]
sha256: 54b13fd6bf435b97ddf37bb0d802ee6aeb26f9083065533ff1f83050f88ee7bb
---

# Buildah scripting approach with CI integration

La surcharge mémoire varie selon les environnements d'exécution des conteneurs, et ces différences deviennent critiques dans les environnements aux ressources limitées ou les déploiements à haute densité.

La surcharge de base d'exécution diffère considérablement :

- s sur le moteur Docker: Nécessite une surcharge du démon ainsi qu'une surcharge par conteneur.
- Podman: Aucune surcharge liée aux démons grâce à une architecture sans démon, surcharge minimale par conteneur
- s sur Containerd: Charge modérée du démon et charge minimale par conteneur
- S SUR LE CRI-O: Faible surcharge du démon et surcharge minimale par conteneur

La déduplication des couches d'image permet d'économiser de la mémoire lors de l'exécution de plusieurs conteneurs à partir d'images associées. Les environnements d'exécution de conteneurs utilisent des systèmes de fichiers de type « copy-on-write » (copie à l'écriture) dans lesquels les couches partagées ne consomment de la mémoire qu'une seule fois pour l'ensemble des conteneurs. Un cluster exécutant de nombreux conteneurs à partir d'images de base similaires peut réaliser d'importantes économies de mémoire grâce à la déduplication.

L'optimisation du mappage mémoire dans les environnements d'exécution modernes réduit l'utilisation de la mémoire résidente. Des outils tels que crun exécutent les fichiers directement à partir du stockage plutôt que de les charger en mémoire, ce qui réduit l'empreinte mémoire des conteneurs contenant des binaires volumineux.

La comptabilité mémoire Cgroup permet un contrôle précis des limites de mémoire des conteneurs, mais les différents environnements d'exécution gèrent différemment la pression mémoire. Certains environnements d'exécution optimisent la récupération de mémoire en cas de forte sollicitation, tandis que d'autres fournissent des rapports plus précis sur l'utilisation de la mémoire pour faciliter les décisions d'autoscaling.

L', qui ne nécessite pas de mémoire, privilégie la sécurité au détriment de l'efficacité. Les conteneurs sans racine nécessitent des processus supplémentaires pour la gestion de l'espace de noms utilisateur et la mise en réseau, ce qui ajoute généralement une surcharge par rapport au fonctionnement avec racine.

Le choix entre les environnements d'exécution se résume souvent à trouver un équilibre entre l'efficacité de la mémoire et les exigences en matière de fonctionnalités. CRI-O offre une faible surcharge pour les charges de travail Kubernetes, tandis que Podman sacrifie une partie de son efficacité au profit de la sécurité et de la compatibilité.

## Intégration du flux de travail de développement

Le meilleur environnement d'exécution de conteneurs n'a aucune valeur s'il ne correspond pas à votre flux de travail de développement. Les alternatives à Docker ont mis en place des outils qui surpassent souvent l'expérience développeur offerte par Docker dans des scénarios spécifiques.

### Environnements Kubernetes locaux

Le développement local de Kubernetes a évolué au-delà de l'approche de machine virtuelle de minikube vers des solutions plus efficaces qui s'intègrent directement aux environnements d'exécution de conteneurs. Le choix de l'environnement local a un impact significatif sur la vitesse de développement et la consommation des ressources.

Kind (Kubernetes dans Docker) permet de créer des clusters Kubernetes en utilisant des nœuds de conteneurs plutôt que des machines virtuelles. Le temps d'installation est généralement de 1 à 2 minutes, avec une charge mémoire modérée par nœud. Kind est compatible avec tous les environnements d'exécution compatibles avec Docker, vous pouvez donc l'utiliser avec Podman (`kind create cluster --runtime podman`) pour le développement Kubernetes sans root.

K3s offre une option légère, exécutant une distribution Kubernetes complète avec une utilisation minimale de la mémoire. Il démarre rapidement et comprend un stockage intégré, une mise en réseau et des contrôleurs d'entrée. K3s fonctionne efficacement avec containerd et peut être exécuté sur des machines de développement aux ressources limitées.

L' MicroK8s de Canonical offre un compromis avec une utilisation modérée de la mémoire et des modules complémentaires modulaires. Il s'intègre parfaitement à containerd et offre des fonctionnalités similaires à celles utilisées en production, sans la surcharge liée aux machines virtuelles. Le temps de démarrage est raisonnable pour un cluster complet.

Rancher Desktop combine K3s avec les backends containerd ou Dockerd, offrant une alternative à Docker Desktop qui utilise moins de ressources. Il comprend une fonctionnalité intégrée de numérisation d'images et une intégration au tableau de bord Kubernetes.

Les pods Podman offrent une alternative unique : vous pouvez développer des applications multi-conteneurs en utilisant le concept de pod de Podman, qui reflète le comportement des pods Kubernetes. Générez directement du code YAML Kubernetes à partir de pods en cours d'exécution à l'aide d'`podman generate kube`, créant ainsi un cheminement fluide entre le développement local et le déploiement en cluster.

### Optimisation du pipeline CI/CD

Les pipelines CI/CD traditionnels basés sur Docker rencontrent des limitations dans les environnements conteneurisés où l'exécution de Docker-in-Docker pose des défis en matière de sécurité et de performances. Les alternatives modernes offrent de meilleures solutions pour créer et déployer des images de conteneurs dans les systèmes d'intégration continue.

Buildah excelle dans les environnements CI car il ne nécessite ni démon ni privilèges root. Vous pouvez créer des images conformes à l'OCI à l'aide de scripts shell qui sont plus faciles à auditer que les fichiers Dockerfile. L'approche de script de Buildah permet la construction dynamique d'images basée sur des variables CI, ce qui la rend idéale pour les processus de construction complexes nécessitant une logique conditionnelle.

Image 7 - Page d'accueil de Buildah

À titre de comparaison, les fichiers Dockerfile traditionnels utilisent des instructions déclaratives :

```
FROM alpine:latest
RUN apk add --no-cache nodejs npm
COPY package.json /app/
WORKDIR /app
```
Buildah utilise des commandes shell impératives pouvant inclure des variables et une logique conditionnelle :

```
# Buildah scripting approach with CI integration
buildah from alpine:latest
buildah run $container apk add --no-cache nodejs npm
buildah copy $container package.json /app/
buildah config --workingdir /app $container
buildah commit $container myapp:${CI_COMMIT_SHA}
```
Cette flexibilité de script vous permet de sélectionner dynamiquement des images de base, d'installer des paquets de manière conditionnelle en fonction des noms de branche ou de modifier les étapes de compilation en fonction des variables d'environnement CI, des capacités qui nécessitent des solutions de contournement complexes dans les fichiers Dockerfile traditionnels.

Kaniko résout le problème de l'e et du Docker-in-Docker en créant des images entièrement dans l'espace utilisateur au sein d'un conteneur. Il fonctionne dans des pods Kubernetes sans nécessiter d'accès privilégié ni de démon Docker. Kaniko est efficace dans les pipelines GitLab CI et Jenkins X où les politiques de sécurité empêchent les conteneurs privilégiés.

Cet outil extrait les images de base, applique les instructions Dockerfile de manière isolée et transfère les résultats directement vers les registres. Les temps de compilation sont comparables à ceux de Docker, mais avec une sécurité nettement améliorée dans les environnements orchestrés.

