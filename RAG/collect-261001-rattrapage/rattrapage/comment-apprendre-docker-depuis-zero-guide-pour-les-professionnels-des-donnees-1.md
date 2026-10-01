---
id: collect-261001-rattrapage/rattrapage/comment-apprendre-docker-depuis-zero-guide-pour-les-professionnels-des-donnees-1
title: "Use an official Python runtime as a parent image"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["apache", "distribution"]
source: docs/RAG/collect-261001-rattrapage/comment-apprendre-docker-depuis-zero-guide-pour-les-professionnels-des-donnees.md
source_anchor: ""
source_lines: [1, 161]
sha256: 06588dddc17089dadd500317dbedb37f3a00b2a5ce00ca42658057dd03a5e82d
---

# Use an official Python runtime as a parent image

Cursus

La conteneurisation a transformé la façon dont les équipes d’ingénierie gèrent et font évoluer les applications, en particulier pour la gestion des données, l’analytique et l’apprentissage automatique. En emballant les applications dans des environnements isolés et légers, les conteneurs garantissent des performances cohérentes du développement à la production.

Parmi les différentes plateformes disponibles, Docker s’impose comme la solution la plus populaire. Sa flexibilité et sa simplicité permettent aux professionnels des données de créer des pipelines reproductibles, évolutifs et efficaces, tout en favorisant la collaboration.

Dans cet article, nous vous proposons un plan d’apprentissage concret de Docker, avec des étapes pour déployer votre première application simple. Allons-y !

## Qu’est-ce que Docker et pourquoi l’apprendre ?

Docker est une plateforme open-source qui simplifie le déploiement, la mise à l’échelle et la gestion des applications via la conteneurisation.

Les conteneurs sont des environnements portables et légers qui regroupent tout le nécessaire pour exécuter une application — code, runtime, bibliothèques et réglages — afin d’assurer des performances constantes sur différents systèmes. Dans les projets data, Docker sert à construire et gérer ces conteneurs, permettant aux applications de tourner de manière fiable sur toute infrastructure.

À la différence des machines virtuelles (VM), qui nécessitent leur propre système d’exploitation et un hyperviseur, Docker ne virtualise que la couche applicative. Résultat : des conteneurs plus rapides à démarrer, moins gourmands en ressources et plus simples à configurer.

*Applications conteneurisées versus machines virtuelles. Source : Docker*

Pour les professionnels des données, Docker permet de créer des environnements reproductibles afin d’exécuter les pipelines avec la même fiabilité du développement à la production. Il réduit les problèmes de dépendances, fluidifie les workflows et favorise la collaboration grâce à des environnements standardisés et partageables.

De plus, Docker s’intègre avec des outils data populaires comme Jupyter, TensorFlow, et Apache Hadoop.

Maîtriser Docker peut accélérer votre productivité, optimiser vos workflows et rendre vos projets facilement déployables et scalables !

## Apprendre Docker depuis zéro : votre premier déploiement

La meilleure façon d’apprendre Docker, c’est de pratiquer. Laissez-nous vous guider dans un premier déploiement simple. Ensuite, nous verrons des plans d’apprentissage pour approfondir vos connaissances.

### Étape 1 : comprendre les concepts clés

Avant de passer à la pratique, il est important d’assimiler quelques fondamentaux. Voici les principaux concepts Docker :

- Conteneurs : unités légères et isolées qui embarquent une application avec toutes ses dépendances, garantissant un fonctionnement identique dans différents environnements.
- Images : modèle en lecture seule utilisé pour créer des conteneurs. Il inclut tout le nécessaire (code, bibliothèques, outils système). Les images sont généralement construites à partir de Dockerfiles.
- Dockerfile : fichier texte contenant les instructions de construction d’une image Docker (installation de logiciels, copie de fichiers, configuration de l’environnement, etc.).
- Docker Hub : Docker Hub est un registre public pour stocker, partager et télécharger des images Docker. Il facilite la distribution et la réutilisation d’environnements préconfigurés.
- Volumes : mécanisme pour persister les données générées et utilisées par les conteneurs Docker, en les stockant hors du cycle de vie des conteneurs afin d’éviter toute perte.
- Réseaux : les réseaux Docker facilitent la communication entre conteneurs. Chaque conteneur peut être connecté à un ou plusieurs réseaux pour échanger des données en toute sécurité.

Aperçu de l’architecture Docker. Source : *Docker*

Comprendre ces notions de base est indispensable avant de déployer des applications avec Docker. Les maîtriser vous donnera des fondations solides et rendra la pratique beaucoup plus efficace.

Le cours Introduction to Docker peut vous aider de manière significative à consolider vos acquis.

### Étape 2 : installer Docker

Pour commencer à utiliser Docker, vous devez l’installer sur votre système. Voici les instructions selon votre plateforme. Pour plus de détails, reportez-vous à la documentation officielle Docker via les liens fournis.

#### 1. Installer Docker sur Windows

Prérequis :

- Windows 10 64 bits : Pro, Enterprise ou Education (Build 19041 ou supérieur)
- Windows 11 64 bits : Home, Pro, Enterprise ou Education
- Backend WSL 2

Étapes :

1. Activer WSL 2 (Windows Subsystem for Linux) :

- Ouvrez PowerShell en tant qu’administrateur.
- Exécutez les commandes suivantes :

```
dism.exe /online /enable-feature /featurename:Microsoft-Windows-Subsystem-Linux /all /norestart 
dism.exe /online /enable-feature /featurename:VirtualMachinePlatform /all /norestart wsl --set-default-version 2
```
2. Installer Docker Desktop pour Windows :

- Téléchargez Docker Desktop depuis le site officiel de Docker.
- Lancez l’installeur et suivez les instructions.
- Choisissez l’option WSL 2 comme backend par défaut pendant l’installation.

3. Démarrer Docker Desktop :

- Lancez Docker Desktop depuis le menu Démarrer.
- Docker devrait démarrer automatiquement ; sinon, lancez-le manuellement.

4. Vérifier l’installation :

- Ouvrez PowerShell ou l’invite de commandes.
- Exécutez la commande suivante pour vérifier que Docker est bien installé :

`sudo docker --version`
Documentation officielle : Docker Desktop for Windows

#### 2. Installer Docker sur macOS

Prérequis :

- macOS 10.15 ou version ultérieure

Étapes :

1. Télécharger Docker Desktop pour macOS :

- Rendez-vous sur le site officiel de Docker et téléchargez l’installateur Docker Desktop.

2. Installer Docker Desktop :

- Ouvrez le fichier `.dmg` téléchargé.
- Glissez l’icône Docker dans le dossier Applications.

3. Démarrer Docker Desktop :

- Ouvrez Docker depuis le dossier Applications.
- Suivez l’assistant pour finaliser l’installation.

4. Vérifier l’installation :

- Ouvrez un terminal.
- Exécutez la commande suivante pour vérifier l’installation de Docker :

`docker --version`
Documentation officielle : Docker Desktop for Mac

#### 3. Installer Docker sur Linux

Distributions prises en charge :

- Ubuntu
- Debian
- Fedora
- CentOS
- RHEL

Étapes pour Ubuntu/Debian :

1. Désinstaller les anciennes versions :

- Avant d’installer Docker Engine, supprimez les paquets en conflit avec cette commande :

`for pkg in docker.io docker-doc docker-compose podman-docker containerd runc; do sudo apt-get remove $pkg; done`
2. Configurer le dépôt apt de Docker :

- Exécutez les commandes suivantes :

```
sudo apt-get update
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc
echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu $(. /etc/os-release && echo \"$VERSION_CODENAME\") stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
```
3. Installer les paquets Docker :

- Installez Docker Engine et les composants nécessaires :

`sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin`
4. Vérifier l’installation :

- Exécutez la commande suivante pour vérifier que Docker est installé :

`sudo docker --version`
Documentation officielle : Docker Engine on Debian

## Devenez ingénieur en données

