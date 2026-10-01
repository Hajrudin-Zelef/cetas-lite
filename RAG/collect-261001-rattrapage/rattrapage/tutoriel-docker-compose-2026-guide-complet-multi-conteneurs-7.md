---
id: collect-261001-rattrapage/rattrapage/tutoriel-docker-compose-2026-guide-complet-multi-conteneurs-7
title: "Vérifier les versions installées"
domain: rattrapage
role: reference
task: reference
actors: ["AWS", "Google", "Microsoft"]
dates: []
keywords: ["arr", "aws", "open source", "scout"]
source: docs/RAG/collect-261001-rattrapage/tutoriel-docker-compose-2026-guide-complet-multi-conteneurs.md
source_anchor: ""
source_lines: [713, 770]
sha256: 9e3706d32ce649aecd8126aec3aef4baa43b4556536bd43264f3cc2f14013cc5
---

# Vérifier les versions installées

Premièrement, ne jamais exécuter de processus en tant que root dans vos conteneurs. Notre Dockerfile inclut déjà la création d’un utilisateur dédié avec `USER appuser`. Cette pratique limite l’impact d’une éventuelle compromission : même si un attaquant exploite une vulnérabilité dans votre application, il n’aura pas les privilèges root dans le conteneur.

Deuxièmement, utilisez toujours des images avec des tags spécifiques plutôt que `:latest`. La notation `postgres:16-alpine` est préférable à `postgres:latest` car elle garantit la reproductibilité de vos déploiements. Un changement majeur de version dans une image `:latest` peut casser votre application sans avertissement. Pour une sécurité maximale, utilisez le digest SHA256 de l’image : `postgres@sha256:abc123...`.

Troisièmement, scannez régulièrement vos images avec Docker Scout ou Trivy. Docker Scout, intégré à Docker Desktop depuis 2025, analyse automatiquement vos images et signale les vulnérabilités CVE connues. Exécutez `docker scout quickview` pour obtenir un résumé rapide des vulnérabilités dans vos images.

Quatrièmement, limitez les capacités système de vos conteneurs. Par défaut, Docker accorde un ensemble de capacités Linux qui ne sont pas nécessaires pour la plupart des applications. Ajoutez `cap_drop: [ALL]` puis `cap_add` uniquement les capacités requises. Pour une API web, vous n’avez généralement besoin d’aucune capacité supplémentaire au-delà de `NET_BIND_SERVICE`.

Cinquièmement, activez le mode `read_only: true` pour le système de fichiers du conteneur lorsque c’est possible, et définissez des volumes `tmpfs` pour les répertoires qui nécessitent une écriture temporaire. Cette approche renforce considérablement la posture de sécurité de vos conteneurs en empêchant toute modification non autorisée du système de fichiers.

## FAQ : Questions Fréquentes sur Docker Compose

### Quelle est la différence entre docker-compose et docker compose ?

`docker-compose` (avec tiret) est l’ancienne version standalone écrite en Python, officiellement dépréciée depuis 2023. `docker compose` (avec espace) est la version V2 intégrée comme plugin à Docker CLI, écrite en Go. En 2026, vous devez utiliser exclusivement `docker compose`. Si vous avez des scripts utilisant l’ancienne syntaxe, mettez-les à jour immédiatement.

### Docker Compose est-il adapté à la production ?

Oui, pour des déploiements de petite à moyenne envergure sur un serveur unique. Docker Compose gère très bien les applications avec 5 à 15 services sur une même machine. Pour des architectures nécessitant du scaling horizontal, de la haute disponibilité multi-nœuds ou des déploiements blue-green, migrez vers Kubernetes. L’outil Compose Bridge facilite cette transition en convertissant vos fichiers Compose en manifestes Kubernetes.

### Comment mettre à jour un seul service sans interrompre les autres ?

Utilisez `docker compose up -d --no-deps --build api`. Le flag `--no-deps` empêche Docker Compose de recréer les services dont dépend l’API. Le flag `--build` force la reconstruction de l’image. Les autres services continuent de fonctionner normalement pendant la mise à jour.

### Peut-on utiliser Docker Compose avec Docker Swarm ?

Oui, mais avec des limitations. `docker stack deploy -c compose.yaml mystack` permet de déployer un fichier Compose sur un cluster Swarm. Cependant, certaines directives comme `build` ou `depends_on` ne sont pas supportées en mode Swarm. En 2026, la plupart des équipes choisissent Kubernetes plutôt que Swarm pour les déploiements multi-nœuds.

### Comment accéder à la base de données depuis un outil externe ?

Si le port est mappé dans `compose.yaml` (par exemple `"5432:5432"`), vous pouvez accéder à PostgreSQL depuis l’hôte via `localhost:5432` avec n’importe quel client SQL (DBeaver, pgAdmin, DataGrip). En production, retirez ce mapping de port pour des raisons de sécurité et accédez uniquement via `docker compose exec`.

### Docker Compose est-il gratuit ?

Docker Compose en tant que plugin CLI est entièrement gratuit et open source. Docker Desktop, qui inclut Compose, est gratuit pour les particuliers, l’éducation et les petites entreprises (moins de 250 employés et moins de 10 millions de dollars de revenus). Les grandes entreprises doivent souscrire à un abonnement Docker Business (à partir de 24 $/mois par utilisateur en 2026). Sur Linux, vous pouvez utiliser Docker Engine + Compose plugin sans Docker Desktop, entièrement gratuitement.

### Comment déboguer un conteneur qui ne démarre pas ?

Suivez cette séquence de diagnostic : 1) `docker compose logs service_name` pour voir les messages d’erreur. 2) `docker compose config` pour vérifier la configuration résolue. 3) `docker compose run --rm service_name sh` pour ouvrir un shell dans un nouveau conteneur avec la même configuration. 4) `docker inspect container_name` pour voir l’état détaillé du conteneur, incluant le code de sortie et la raison de l’arrêt.

### Quelle est la différence entre volumes et bind mounts ?

Les **volumes nommés** (`postgres_data:/var/lib/postgresql/data`) sont gérés par Docker, stockés dans `/var/lib/docker/volumes/`, et offrent les meilleures performances. Les **bind mounts** (`./config:/etc/app/config`) lient un chemin de l’hôte au conteneur, idéal pour les fichiers de configuration et le développement. Utilisez les volumes nommés pour les données persistantes (bases de données) et les bind mounts pour les fichiers que vous modifiez fréquemment.

## Couverture Associée

Pour approfondir vos connaissances sur Docker, le DevOps et l’écosystème cloud en 2026, consultez nos articles connexes :

- Docker vs Kubernetes 2026 : Le Comparatif Définitif des Conteneurs – Comprendre quand passer de Compose à Kubernetes
- Tutoriel Ansible 2026 : Automatiser Votre Infrastructure de Zéro à la Production – Compléter Docker Compose avec l’automatisation Ansible
- GitLab vs GitHub 2026 : Le Comparatif Définitif des Plateformes DevOps – Choisir la meilleure plateforme CI/CD pour vos workflows Docker
- Tutoriel FastAPI Python 2026 : Créer une API REST Complète – Approfondir FastAPI, le framework utilisé dans notre tutoriel
- Cloud Souverain vs Cloud Public 2026 : Le Comparatif pour la France et l’Europe – Comprendre les enjeux du cloud pour les déploiements Docker en entreprise
- AWS vs Azure vs Google Cloud 2026 : Le Comparatif Définitif – Choisir le bon cloud provider pour héberger vos conteneurs

Pour aller plus loin avec Docker Compose, consultez la documentation officielle Docker Compose, le dépôt GitHub de Docker Compose et le Docker Hub pour trouver des images officielles. La communauté Docker Forums est également une excellente ressource pour obtenir de l’aide sur des problèmes spécifiques.

*Dernière mise à jour : 19 août 2026, suite à la sortie de Docker Compose v5.5.0 le 17 août 2026. Entre-temps, l’outil a connu une cadence de publication soutenue : Docker Desktop 4.66.1 (26 mars 2026) embarquait la v5.1.1, suivi de Docker Desktop 4.71.0 (20 avril 2026) avec la v5.1.3. Ce tutoriel est maintenu et mis à jour régulièrement pour refléter les dernières versions de Docker Compose et les bonnes pratiques actuelles.*
