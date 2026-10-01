---
id: collect-261001-rattrapage/rattrapage/tutoriel-docker-compose-2026-guide-complet-multi-conteneurs-6
title: "Vérifier les versions installées"
domain: rattrapage
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["arr", "cyber", "gpu", "memory", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/tutoriel-docker-compose-2026-guide-complet-multi-conteneurs.md
source_anchor: ""
source_lines: [639, 712]
sha256: 0709717933b7ea16bdff7263942ca31171ac3d41f774ca3f69028a72eca4fa59
---

# Vérifier les versions installées

**Cause :** Vous utilisez un Mac Apple Silicon (M1/M2/M3/M4) et l’image Docker n’est pas disponible pour l’architecture ARM64. Fréquent avec les images communautaires non officielles.

**Solution :** Ajoutez `platform: linux/amd64` au service concerné dans `compose.yaml`. Docker Desktop émuler l’architecture x86 via Rosetta 2, avec une légère perte de performance. Privilégiez toujours les images officielles qui supportent le multi-architecture.

### Problème 3 : L’API ne se connecte pas à la base de données

**Cause :** L’API démarre avant que PostgreSQL ne soit prêt, ou les variables d’environnement sont incorrectes. Sans health check, `depends_on` ne garantit que l’ordre de démarrage, pas la disponibilité.

**Solution :** Vérifiez que vos health checks sont en place avec `condition: service_healthy`. Exécutez `docker compose config` pour vérifier que les variables d’environnement sont correctement résolues. Testez la connectivité manuellement : `docker compose exec api python -c "import psycopg2; psycopg2.connect(host='db', user='appuser', password='...', dbname='appdb')"`.

### Problème 4 : Les modifications du code ne sont pas prises en compte

**Cause :** Le cache de build Docker réutilise une couche précédente. C’est le cas lorsque vous modifiez le code source mais que le Dockerfile n’a pas changé au niveau de l’instruction `COPY`.

**Solution :** Utilisez `docker compose up --build` pour forcer la reconstruction, ou `docker compose build --no-cache api` pour un rebuild complet sans cache. En développement, activez Compose Watch avec `docker compose watch` pour la synchronisation automatique des fichiers.

### Problème 5 : “insufficient memory” ou conteneur qui crashe

**Cause :** Docker Desktop est configuré avec des ressources insuffisantes. Par défaut, Docker Desktop limite la mémoire à 2 Go, ce qui peut être insuffisant pour une stack avec PostgreSQL, Redis et une API.

**Solution :** Dans Docker Desktop, allez dans Settings → Resources et augmentez la mémoire à 4 Go minimum. Sur Linux, vérifiez avec `docker info | grep -i memory`. Ajoutez des limites de ressources dans `compose.yaml` pour éviter qu’un service ne monopolise toute la mémoire disponible.

### Problème 6 : Les données disparaissent après un redémarrage

**Cause :** Vous utilisez `docker compose down -v` qui supprime les volumes, ou vos données sont stockées dans le conteneur au lieu d’un volume nommé.

**Solution :** Vérifiez que vos volumes sont correctement définis dans la section `volumes:` du `compose.yaml`. Utilisez `docker compose down` (sans `-v`) pour conserver les données. Vérifiez avec `docker volume ls` que les volumes existent.

### Problème 7 : “network compose-tutorial_backend not found”

**Cause :** Un réseau orphelin d’une exécution précédente interfère avec le nouveau déploiement. Cela se produit souvent après un crash ou un arrêt forcé.

**Solution :** Nettoyez les réseaux orphelins avec `docker network prune`. Si le problème persiste, supprimez manuellement le réseau avec `docker network rm compose-tutorial_backend` puis relancez `docker compose up`.

### Problème 8 : Compose ne détecte pas les changements dans le fichier .env

**Cause :** Docker Compose charge le fichier `.env` au moment de l’analyse du fichier Compose, pas au démarrage des conteneurs. Si vous modifiez `.env` pendant que les services tournent, les changements ne sont pas appliqués.

**Solution :** Après toute modification du fichier `.env`, vous devez recréer les conteneurs avec `docker compose up -d`. Pour vérifier que les variables sont correctement chargées, utilisez `docker compose config` qui affiche la configuration résolue avec toutes les substitutions appliquées.

## Astuces Avancées Docker Compose en 2026

Au-delà des bases, Docker Compose offre des fonctionnalités avancées qui transforment votre workflow de développement et de déploiement. Voici les techniques que les ingénieurs DevOps les plus expérimentés utilisent en 2026.

**Compose Bridge vers Kubernetes :** L’outil Compose Bridge, mis en avant dans la documentation Docker 2026, transforme vos fichiers `compose.yaml` en manifestes Kubernetes. C’est un excellent moyen de migrer progressivement vers Kubernetes sans réécrire toute votre configuration. Exécutez `docker compose bridge` pour générer les fichiers YAML Kubernetes correspondant à votre stack.

**Support SDK officiel :** Docker Compose v5.0.0 introduit un SDK officiel qui permet d’intégrer Compose dans des outils tiers via des paramètres fonctionnels et des API documentées. Cela ouvre la porte à l’automatisation programmatique de vos déploiements, par exemple dans des scripts Python ou des workflows CI/CD personnalisés.

**Ressources OCI et Git distantes :** La v5.0.0 apporte le support officiel des ressources OCI et Git distantes dans les fichiers Compose. Vous pouvez désormais référencer des configurations stockées dans un registre OCI ou un dépôt Git distant, facilitant le partage de configurations standardisées à travers plusieurs équipes et projets.

**Builds avec Docker Bake :** La délégation des builds à Docker Bake dans la v5.0.0 aligne le comportement de Compose avec celui de `docker build`. Docker Bake offre des fonctionnalités avancées comme les builds parallèles, les matrices de build et les cibles multiples, ce qui accélère considérablement les temps de construction pour les projets multi-services.

**Inclure des fichiers Compose externes :** La directive `include` permet de décomposer votre configuration en fichiers modulaires. Par exemple, vous pouvez avoir un fichier `compose.monitoring.yaml` pour Prometheus/Grafana et l’inclure uniquement dans certains environnements. C’est une approche modulaire qui améliore la maintenabilité des configurations complexes.

## Comparaison : Docker Compose vs Alternatives en 2026

Pour situer Docker Compose dans l’écosystème des outils d’orchestration en 2026, voici un comparatif avec les principales alternatives. Chaque outil a ses forces et ses cas d’usage privilégiés. Docker Compose excelle dans le développement local et les déploiements simples, tandis que Kubernetes domine pour l’orchestration à grande échelle en production.

| Critère | Docker Compose | Kubernetes | Docker Swarm | Podman Compose | 
|---|---|---|---|---|
| Courbe d’apprentissage | Faible | Élevée | Moyenne | Faible | 
| Cas d’usage principal | Développement local, CI/CD, petite production | Production à grande échelle | Production moyenne | Développement local sans daemon | 
| Scaling automatique | Manuel ( `--scale` ) | Automatique (HPA) | Basique | Manuel | 
| Haute disponibilité | Non native | Oui (multi-nœuds) | Oui (multi-nœuds) | Non native | 
| Configuration | 1 fichier YAML | Multiples manifestes | 1 fichier YAML + init | 1 fichier YAML | 
| Communauté (2026) | Très large | Massive | En déclin | Croissante | 
| Support GPU | Oui (runtime nvidia) | Oui (device plugins) | Limité | Oui | 

En 2026, la trajectoire est claire : Docker Compose pour le développement et les environnements simples, Kubernetes pour la production à grande échelle. Docker Swarm, bien que toujours fonctionnel, voit sa communauté décliner au profit de Kubernetes. Podman Compose gagne en popularité dans les environnements où un daemon root n’est pas souhaitable, notamment dans les entreprises soumises à des contraintes de sécurité strictes comme celles régies par la directive NIS2 en Europe.

## Bonnes Pratiques de Sécurité Docker Compose

La sécurité des conteneurs est devenue une priorité absolue en 2026, notamment en Europe avec l’entrée en vigueur de la directive NIS2 et le renforcement du Cyber Resilience Act. Voici les bonnes pratiques de sécurité à appliquer systématiquement dans vos projets Docker Compose.

