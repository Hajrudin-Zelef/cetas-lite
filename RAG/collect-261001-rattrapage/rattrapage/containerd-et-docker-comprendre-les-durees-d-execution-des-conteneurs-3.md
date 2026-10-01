---
id: collect-261001-rattrapage/rattrapage/containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs-3
title: "containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs"
domain: rattrapage
role: reference
task: reference
actors: ["AWS", "Google"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-rattrapage/containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs.md
source_anchor: ""
source_lines: [138, 213]
sha256: 1c3db13c59ad298096ebe43308b5cc8fb5e232c38f7cf35178fb63b9d7b6bdbc
---

# containerd-et-docker-comprendre-les-durees-d-execution-des-conteneurs

Containerd est fourni avec `ctr`, une interface CLI minimale destinée exclusivement au débogage et au test des fonctionnalités de bas niveau de containerd. L'outil `ctr` est intentionnellement peu convivial pour les développeurs. Il manque des fonctionnalités courantes telles que les raccourcis de mappage de ports, les politiques de redémarrage automatique et l'intégration avec les assistants d'identification. Il est conçu pour les développeurs de conteneurs, et non pour les développeurs d'applications.

#### Combler le fossé avec nerdctl

Cette lacune en matière d'utilisabilité a conduit à la création d'`nerdctl`, une interface CLI compatible avec Docker pour containerd. L'utilisation d'`nerdctl` est similaire à celle de Docker : même syntaxe de commande, mêmes indicateurs, même flux de travail, mais avec containerd comme environnement d'exécution sous-jacent. Cela fait d'`nerdctl`, un excellent outil de transition pour les équipes qui passent de Docker à containerd dans leurs environnements de développement.

Dans la pratique, les développeurs interagissent rarement directement avec containerd. Lorsque cela s'avère nécessaire, `nerdctl` fournit l'interface familière à laquelle ils s'attendent, tandis que les opérateurs et administrateurs de plateformes utilisent les API de containerd de manière programmatique via des systèmes d'orchestration tels que Kubernetes.

Afin d'illustrer les différences dans la pratique, voici une comparaison des opérations courantes sur les conteneurs entre les trois outils CLI :

| **Tâche** | **Docker** | **nerdctl** | **ctr** | 
| Exécuter le conteneur | `docker run -d -p 8080:80 nginx` | `nerdctl run -d -p 8080:80 nginx` | `ctr run --net-host -d docker.io/library/nginx:latest nginx_id` | 
| Liste des conteneurs | `docker ps` | `nerdctl ps` | `ctr tasks list` | 
| Construire une image | `docker build -t myapp .` | `nerdctl build -t myapp .` | *Non pris en charge* | 
| Consulter les journaux | `docker logs`  | `nerdctl logs`  | *Non pris en charge* | 
| Veuillez inspecter le conteneur. | `docker inspect`  | `nerdctl inspect`  | `ctr containers info`  | 
| Veuillez cliquer sur l'image | `docker pull nginx` | `nerdctl pull nginx` | `ctr images pull docker.io/library/nginx:latest` | 
| Composer le soutien | `docker compose up` | `nerdctl compose up` | *Non pris en charge* | 

Remarque : ctr ne dispose pas de mappage de ports (`-p`) et nécessite une mise en réseau hôte (`--net-host`) pour exposer les services. Il ne télécharge pas automatiquement les images.

### Aperçu des principales différences

Avant d'aborder les recommandations pour des cas d'utilisation spécifiques, récapitulons les différences que nous avons abordées jusqu'à présent. Le tableau suivant résume les principales différences fonctionnelles entre Docker et containerd, en mettant en évidence leurs capacités distinctes et leurs publics cibles :

| **Caractéristique** | **Docker** | **Conteneur** | 
| Construction d'image | Intégré (Dockerfiles, BuildKit) | Nécessite des outils externes (buildctl, nerdctl) | 
| Orchestration | Docker Swarm / Kubernetes | Aucun (utilisé par Kubernetes) | 
| Gestion du stockage | Gestion du volume | Système de capture d'écran | 
| Interface graphique | Docker Desktop | Aucun | 
| Cycle de vie des conteneurs | Gestion complète (via containerd) | Objectif principal (compatible CRI) | 
| Utilisateurs principaux | Développeurs d'applications | Opérateurs de clusters, développeurs de plateformes | 

## Pourquoi choisir Docker ?

Maintenant que nous avons abordé les différences techniques, examinons des scénarios pratiques dans lesquels chaque outil excelle. Malgré l'essor des conteneurs dans les environnements de production, Docker demeure le choix privilégié pour des scénarios spécifiques où l'expérience des développeurs et la disponibilité d'outils complets sont primordiales.

### Pour le développement local et le prototypage

Docker est particulièrement efficace lorsque vous avez besoin d'une solution « tout-en-un » pour écrire et tester du code. Grâce à la chaîne d'outils intégrée, les développeurs peuvent passer de zéro à l'exécution de conteneurs en quelques minutes, sans avoir à assembler plusieurs composants ni à configurer un réseau complexe.

L'écosystème Docker offre d'importants avantages en termes de productivité :

- **Docker Hub :** Des millions d'images prêtes à l'emploi pour les bases de données, les files d'attente de messages et les serveurs Web.
- **Docker Compose:** Définissez des applications multi-conteneurs dans un seul fichier YAML et lancez des environnements de développement complets à l'aide d'une seule commande.
- **Interface graphique Docker Desktop :** Gestion visuelle des conteneurs, exploration des volumes, ajustements des limites de ressources et prise en charge intégrée de Kubernetes.
- **Réduction des obstacles à l'entrée :** Des outils graphiques et des commandes intuitives rendent les conteneurs accessibles aux développeurs novices dans cette technologie.

Pour les équipes utilisant Docker Desktop, l'interface graphique offre des fonctionnalités supplémentaires qui accélèrent considérablement l'intégration et les flux de travail quotidiens.

### Pour les pipelines de construction complexes

Au-delà du développement, les capacités de compilation de Docker en font un choix naturel pour les workflows sophistiqués d'intégration et de déploiement continus.

Le BuildKit intégré à Docker offre des fonctionnalités avancées indispensables aux pipelines CI/CD modernes. Les constructions en plusieurs étapes réduisent la taille des images tout en conservant la lisibilité des fichiers Dockerfile. Les mécanismes de mise en cache de BuildKit réutilisent intelligemment les couches d'une compilation à l'autre, ce qui réduit considérablement les temps de compilation dans les environnements d'intégration continue.

La plupart des plateformes d'automatisation, telles que GitHub Actions, GitLab CI, Jenkins et d'autres, disposent d'intégrations Docker éprouvées et matures. Ces intégrations gèrent l'authentification, la mise en cache et la publication d'images sans aucune difficulté.

Bien que d'autres outils puissent offrir des résultats similaires, l'omniprésence de Docker signifie que des solutions et une assistance pour le dépannage sont facilement accessibles, ce qui constitue un autre avantage considérable.

## Pourquoi choisir Containerd ?

Containerd se distingue dans les scénarios de production où le minimalisme, les performances et la stabilité priment sur la commodité des outils intégrés.

### Pour les clusters Kubernetes de production

L'utilisation de containerd comme environnement d'exécution pour les nœuds Kubernetes apporte des avantages significatifs :

- **Réduction des frais généraux :** La suppression du démon Docker permet de réduire la consommation de ressources par nœud et de réaliser des économies significatives à grande échelle.
- **Stabilité améliorée :** Moins de composants mobiles signifie moins de points de défaillance potentiels dans votre infrastructure.
- **Surface d'attaque réduite :** Moins de code à auditer et moins de vulnérabilités potentielles en matière de sécurité
- **Débogage simplifié :** L'intégration directe de CRI élimine la couche de traduction dockershim, ce qui simplifie le dépannage.
- **Meilleures performances :** La pile d'exécution optimisée améliore les temps de démarrage et la réactivité des conteneurs.

Le statut de graduation CNCF de Containerd témoigne de sa maturité et de sa fiabilité. Les principaux fournisseurs de services cloud, notamment AWS, Google Cloud et Azure, ont adopté containerd comme norme pour leurs offres Kubernetes gérées, démontrant ainsi leur confiance dans sa capacité à être utilisé en production pour les infrastructures critiques.

