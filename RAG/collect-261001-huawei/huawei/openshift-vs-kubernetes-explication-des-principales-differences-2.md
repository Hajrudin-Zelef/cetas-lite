---
id: collect-261001-huawei/huawei/openshift-vs-kubernetes-explication-des-principales-differences-2
title: "openshift-vs-kubernetes-explication-des-principales-differences"
domain: huawei
role: reference
task: reference
actors: ["AWS", "Google"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-huawei/openshift-vs-kubernetes-explication-des-principales-differences.md
source_anchor: ""
source_lines: [77, 151]
sha256: 25c746863a048cfdebf77df6c82ee018756e78220a0161919f0e9614f345e57e
---

# openshift-vs-kubernetes-explication-des-principales-differences

- CI/CD intégré: OpenShift inclut un support natif pour les pipelines Tekton et les flux de travail GitOps.
- Console Web: Une interface web propre et conviviale pour les développeurs et les administrateurs de clusters. Il rend la gestion des déploiements, la surveillance des ressources et l'affichage des journaux beaucoup plus accessibles.
- Améliorations de la sécurité: Des éléments tels que les contraintes de contexte de sécurité (SCC), l'analyse d'images et le RBAC intégré sont plus stricts et plus standardisés que Kubernetes vanille.
- Red Hat Enterprise Linux CoreOS (RHCOS): Un système d'exploitation spécialement conçu pour fonctionner sur les nœuds OpenShift, optimisé pour les charges de travail des conteneurs.
- OperatorHub: Une place de marché curatée pour installer et gérer les opérateurs Kubernetes, qui automatisent les tâches du cycle de vie des applications.

Ces ajouts rendent OpenShift particulièrement intéressant pour les grandes équipes ou les organisations qui ont besoin de plus qu'une plateforme Kubernetes de base.

### Saveurs OpenShift

Selon la façon dont vous souhaitez déployer et gérer OpenShift, il existe quelques saveurs différentes :

- OpenShift Container Platform: La version autogérée que vous pouvez exécuter sur votre infrastructure, qui est populaire dans les configurations sur site et hybrides.
- OpenShift Online: Un service OpenShift entièrement hébergé et géré par Red Hat, idéal pour démarrer sans rien mettre en place soi-même.
- OpenShift Dedicated: Une offre OpenShift gérée, hébergée sur des fournisseurs de cloud public comme AWS et Google Cloud, mais gérée par Red Hat en votre nom.

Chacun d'entre eux répond à des besoins différents, qu'il s'agisse de startups souhaitant être rapidement opérationnelles ou d'entreprises ayant des exigences strictes en matière de sécurité et de conformité.

## Comparaison des architectures

Si OpenShift et Kubernetes partagent une base commune, leurs approches architecturales divergent dès que l'on dépasse les bases.

Kubernetes offre un modèle modulaire, à construire soi-même, tandis qu'OpenShift fournit une pile plus intégrée, fondée sur des opinions. Il est essentiel de comprendre en quoi les deux plateformes diffèrent sous le capot pour savoir laquelle conviendra le mieux à votre équipe ou à votre organisation.

### Plate-forme de base

À la base, les deux plateformes exécutent Kubernetes. OpenShift est construit directement au-dessus de Kubernetes et s'en tient à ses API et composants en amont. Un cluster OpenShift standard comprend tous les éléments Kubernetes essentiels, comme le serveur API, etcd, le planificateur et les kubelets. Cela signifie qu'il est entièrement compatible avec les charges de travail et les outils Kubernetes.

Ce qui distingue OpenShift, c'est l'écosystème que Red Hat construit autour de cette base. Il introduit des fonctionnalités telles que Red Hat Enterprise Linux CoreOS (RHOCS) en tant que système d'exploitation du nœud, un registre de conteneurs intégré et des services de plateforme supplémentaires étroitement couplés à Kubernetes.

Ces intégrations réduisent la quantité de travail que les équipes doivent effectuer pour créer et sécuriser un environnement de production.

Aperçu de l'architecture de la plateforme de conteneurs OpenShift. Image fournie par la documentation de Red Hat.

### Installation et configuration

L'une des différences architecturales les plus importantes est la méthode d'installation de chaque plateforme.

- Kubernetes offre une grande souplesse de déploiement. Vous pouvez mettre en place un cluster local à l'aide d'outils comme minikube, de services gérés dans le cloud comme GKE ou EKS, ou l'installer manuellement à l'aide d'outils comme kubeadm. Bien que cela vous donne un contrôle total, cela signifie également que vous êtes responsable de tout, y compris de la mise en place du réseau, de la configuration du stockage et de l'installation d'outils tels que la surveillance et l'enregistrement.
- OpenShift, en revanche, offre un processus d'installation plus guidé et automatisé. OpenShift propose plusieurs méthodes d'installation pour s'adapter aux différents environnements d'infrastructure et aux préférences des utilisateurs.
- La première est l'infrastructure fournie par l'installateur, qui automatise la création de grappes sur les plates-formes prises en charge.
- L'autre est l'infrastructure fournie par l'utilisateur, qui donne plus de contrôle aux équipes d'infrastructure.
- Vous pouvez en savoir plus sur l'installation d'OpenShift dans la documentation utilisateur officielle.

### Outils de déploiement et de gestion

Kubernetes et OpenShift fournissent tous deux des outils en ligne de commande pour interagir avec le cluster, mais ils diffèrent en termes de portée et de facilité d'utilisation :

- Kubernetes s'appuie sur `kubectl` , un outil CLI puissant qui interagit directement avec l'API Kubernetes. Il est flexible et largement utilisé, mais suppose un certain niveau de familiarité avec YAML et l'architecture sous-jacente.
- OpenShift inclut `oc` , un CLI qui étend`kubectl` et possède des capacités supplémentaires. Il prend en charge les mêmes commandes de base et ajoute des fonctionnalités telles que l'accès par projet, la gestion des flux d'images et des flux de connexion simplifiés.

OpenShift comprend également une console web complète qui facilite la gestion des applications, le suivi des charges de travail et le contrôle des accès, sans qu'il soit nécessaire d'écrire YAML à la main.

Kubernetes propose également un tableau de bord, mais sa configuration et sa sécurisation nécessitent plus de travail.

Les deux plates-formes prennent en charge Helm pour manager les paquets d'applications. OpenShift s'intègre également étroitement avec les opérateurs, disponibles via l'OperatorHub intégré.

## Sécurité et gouvernance

La sécurité est souvent un facteur de différenciation clé entre Kubernetes et OpenShift, en particulier pour les équipes opérant dans des environnements réglementés ou les entreprises ayant des exigences de conformité strictes.

Comprendre comment les deux plateformes gèrent la sécurité, le contrôle d'accès et la conformité peut vous aider à décider laquelle est la mieux adaptée au profil de risque et aux besoins de gouvernance de votre organisation.

### Politiques de sécurité par défaut

Alors que Kubernetes fournit un cadre de sécurité flexible, OpenShift s'appuie sur lui avec une approche plus stricte et plus sécurisée par défaut.

- Kubernetes offre beaucoup de flexibilité en matière de sécurité, mais il vous faut surtout tout construire à partir de zéro. Kubernetes n'applique pas de règles de sécurité strictes au niveau de l'exécution. Il fournit des fonctionnalités telles que PodSecurity Admission (PSA) et Security Contexts, qui doivent être explicitement configurées et gérées.

- OpenShift est plus proactif en matière de sécurité. Il est livré avec les contraintes de contexte de sécurité (SCC) activées par défaut. Ils définissent ce qu'un pod peut et ne peut pas faire, par exemple s'il peut fonctionner en tant que root ou utiliser le réseau de l'hôte. En appliquant ces contraintes, OpenShift permet de réduire le risque de mauvaise configuration et d'escalade des privilèges, ce qui est particulièrement important dans les clusters multi-tenants.

En pratique, les charges de travail qui fonctionnent bien sur Kubernetes pourraient échouer sur OpenShift jusqu'à ce qu'elles répondent à ses normes de sécurité plus strictes. Cela peut sembler un obstacle à première vue, mais cela oblige les équipes à adopter les meilleures pratiques dès le début du cycle de développement.

### Contrôle d'accès basé sur les rôles (RBAC)

