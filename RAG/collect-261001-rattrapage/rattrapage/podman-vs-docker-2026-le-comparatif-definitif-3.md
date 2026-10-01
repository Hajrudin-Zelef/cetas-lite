---
id: collect-261001-rattrapage/rattrapage/podman-vs-docker-2026-le-comparatif-definitif-3
title: "Créer un pod avec Podman"
domain: rattrapage
role: reference
task: reference
actors: ["JFrog"]
dates: []
keywords: ["agents", "open source"]
source: docs/RAG/collect-261001-rattrapage/podman-vs-docker-2026-le-comparatif-definitif.md
source_anchor: ""
source_lines: [122, 214]
sha256: 354f3e8a22df672daab30f589d7df4df22ad487f09e2cfb0476455ad5662f0ee
---

# Créer un pod avec Podman

Docker propose Docker Compose pour le développement local et Docker Swarm pour l’orchestration, mais aucun de ces outils ne génère nativement du YAML Kubernetes. Des outils tiers comme **Kompose** permettent de convertir des fichiers docker-compose.yml en manifestes Kubernetes, mais cette étape supplémentaire introduit de la complexité et des risques d’incompatibilité.

Pour les entreprises européennes qui utilisent massivement Kubernetes en 2026 – selon le rapport CNCF 2025, 78 % des organisations de plus de 500 employés en Europe utilisent Kubernetes en production – la compatibilité native de Podman représente un avantage opérationnel tangible qui réduit les frictions entre développement et déploiement.

## Intégration CI/CD : GitHub Actions, GitLab CI et Jenkins

L’intégration dans les pipelines d’intégration et de déploiement continus est essentielle pour tout outil de conteneurisation. En 2026, Docker reste le standard de facto dans les environnements CI/CD, mais **Podman** gagne rapidement du terrain grâce à sa compatibilité CLI de 95 %.

### GitHub Actions et GitLab CI

Les runners Ubuntu de GitHub Actions incluent Podman préinstallé depuis 2024. Les workflows existants peuvent utiliser Podman en remplacement direct de Docker pour la plupart des opérations. La migration se résume souvent à un simple alias :

```
# .github/workflows/build.yml
name: Build avec Podman
on: [push]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Build image
        run: |
          podman build -t mon-app:latest .
          podman push mon-app:latest ghcr.io/${{ github.repository }}/mon-app:latest
```
L’avantage principal de Podman dans les pipelines GitHub Actions est la **sécurité rootless**. Les runners s’exécutent dans des environnements partagés, et l’absence de daemon root réduit les risques d’évasion de conteneur. Pour les projets open source avec des contributeurs externes, cette couche de sécurité supplémentaire est non négligeable.

GitLab CI supporte Podman comme runtime alternatif, bien que la configuration par défaut utilise Docker. Jenkins, via le plugin Podman, permet d’exécuter des builds de conteneurs sans Docker installé sur les agents. Les deux plateformes bénéficient de la compatibilité CLI de Podman, qui accepte les mêmes commandes `build`, `push`, `pull` et `run` que Docker.

Les quelques incompatibilités concernent principalement Docker Swarm (non supporté par Podman), certaines fonctionnalités avancées de BuildKit et la gestion des volumes (Podman utilise des bind mounts légers là où Docker offre une gestion de volumes plus complète). Pour les pipelines standard de build-test-push, la migration est généralement transparente.

## Écosystème et Support Compose : Docker Compose vs Podman Compose

L’écosystème est historiquement le point fort de Docker, et cette réalité persiste en 2026. **Docker Hub** héberge des millions d’images officielles et communautaires, et Docker Compose V2 est devenu un standard de fait pour la définition d’applications multi-conteneurs en développement local.

Podman offre une compatibilité d’environ **90 %** avec Docker Compose via la commande intégrée `podman compose` (disponible depuis Podman 4.1, 2022) ou l’outil communautaire **podman-compose**. La majorité des fichiers `docker-compose.yml` fonctionnent sans modification, mais certaines fonctionnalités avancées (réseaux overlay, configurations Swarm, certains plugins de volumes) peuvent nécessiter des ajustements.

```
# docker-compose.yml compatible Podman
version: "3.9"
services:
  web:
    image: nginx:alpine
    ports:
      - "8080:80"
    volumes:
      - ./html:/usr/share/nginx/html:Z
  db:
    image: postgres:16
    environment:
      POSTGRES_PASSWORD: secret
    volumes:
      - pgdata:/var/lib/postgresql/data:Z
volumes:
  pgdata:
# Lancer avec Podman
# podman compose up -d
```
Le suffixe `:Z` sur les volumes est une particularité Podman/SELinux qui applique le bon contexte de sécurité aux montages. C’est l’un des rares ajustements nécessaires lors de la migration d’un fichier Compose de Docker vers Podman.

Concernant les registres d’images, Podman supporte nativement Docker Hub, Quay.io (Red Hat), GitHub Container Registry et tout registre compatible OCI. La configuration multi-registre est même plus flexible que celle de Docker : Podman peut chercher une image simultanément dans plusieurs registres configurés dans `/etc/containers/registries.conf`, là où Docker se connecte par défaut uniquement à Docker Hub.

L’écosystème de plugins et d’extensions reste plus riche côté Docker. Docker Desktop propose des extensions pour le monitoring, la sécurité, le scanning d’images (Snyk, JFrog) qui n’ont pas toujours d’équivalent direct dans Podman Desktop. Toutefois, l’écart se réduit rapidement : Podman Desktop a ajouté le support des extensions en 2025, et la communauté Red Hat contribue activement à enrichir l’écosystème.

## Intégration Systemd : L’Atout Production de Podman

L’intégration avec **systemd**, le gestionnaire de services standard de Linux, est un avantage souvent sous-estimé de Podman. Là où Docker nécessite des outils tiers ou des scripts personnalisés pour gérer les conteneurs comme des services système, Podman génère nativement des fichiers unit systemd.

```
# Générer un fichier unit systemd pour un conteneur Podman
podman generate systemd --new --name mon-app > ~/.config/systemd/user/mon-app.service
# Activer et démarrer le service
systemctl --user enable --now mon-app.service
# Vérifier le statut
systemctl --user status mon-app.service
```
Cette intégration permet de gérer les conteneurs Podman exactement comme n’importe quel autre service Linux : démarrage automatique au boot, redémarrage en cas d’échec, gestion des dépendances entre services, journalisation via journald. Pour les administrateurs système habitués à systemd, cette approche est naturelle et ne nécessite aucun apprentissage supplémentaire.

En production, cette capacité est particulièrement utile pour les déploiements sur des serveurs bare metal ou des machines virtuelles où Kubernetes serait surdimensionné. Un serveur web avec quelques conteneurs applicatifs peut être entièrement géré via systemd et Podman, sans orchestrateur supplémentaire. C’est un cas d’usage très courant dans les PME européennes qui n’ont pas les ressources pour maintenir un cluster Kubernetes.

Docker a introduit un support basique de systemd, mais il nécessite que le daemon Docker soit lui-même lancé comme service. Cela crée une dépendance en chaîne (systemd → dockerd → conteneur) là où Podman offre une relation directe (systemd → conteneur). En cas de problème avec le daemon Docker, tous les conteneurs gérés sont impactés. Avec Podman, chaque conteneur est indépendant.

## 5 Cas d’Usage Réels : Quand Choisir Docker ou Podman

Au-delà des spécifications techniques, le choix entre Docker et Podman dépend largement de votre contexte d’utilisation. Voici cinq scénarios concrets avec des recommandations argumentées pour chaque cas.

**1. Développement local sur macOS/Windows – Recommandation : Podman Desktop**

Pour les développeurs individuels ou les petites équipes travaillant sur macOS ou Windows, Podman Desktop offre une expérience comparable à Docker Desktop sans aucun coût de licence. L’application fournit une interface graphique pour gérer les conteneurs, les images et les pods, avec la même compatibilité CLI. L’économie de 9 à 24 $/utilisateur/mois par rapport à Docker Desktop est un argument de poids, surtout pour les freelances et les startups françaises en phase d’amorçage.

**2. Environnements RHEL/Fedora en entreprise – Recommandation : Podman**

