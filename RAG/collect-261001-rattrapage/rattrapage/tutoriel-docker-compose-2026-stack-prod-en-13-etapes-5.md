---
id: collect-261001-rattrapage/rattrapage/tutoriel-docker-compose-2026-stack-prod-en-13-etapes-5
title: "Mise à jour des dépôts et installation des dépendances"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agents", "apache", "arr", "attention"]
source: docs/RAG/collect-261001-rattrapage/tutoriel-docker-compose-2026-stack-prod-en-13-etapes.md
source_anchor: ""
source_lines: [538, 631]
sha256: 2436448798d5c993782b4c47ff1c1f92efb1caf6939609e4e8ef93e8005acfcc
---

# Mise à jour des dépôts et installation des dépendances

- **Piège n° 3 – Croire que `depends_on` attend la disponibilité applicative.** Sans`condition: service_healthy` ,`depends_on` attend uniquement que le conteneur*démarre* , pas que PostgreSQL accepte les connexions. Utilisez systématiquement les healthchecks pour les bases de données.
- **Piège n° 4 – Bind mounts en lecture-écriture par défaut.** Tout fichier de configuration doit se monter en`:ro` . Sinon, un conteneur compromis peut réécrire votre`nginx.conf` et tenir un point d’ancrage.
- **Piège n° 5 – Exposer trop de ports avec `ports:`.** En production, seul le reverse proxy doit avoir un mapping`HOST:CONTAINER` . Les bases de données et caches doivent rester sur le réseau interne uniquement.
- **Piège n° 6 – Oublier `restart:`.** Sans politique, un conteneur qui crashe ne redémarre jamais. Préférez`unless-stopped` pour la production.
- **Piège n° 7 – Mettre des secrets dans `environment:`.** Les variables d’environnement sont visibles via`docker inspect` et fuient dans les logs et les rapports d’erreur. Utilisez`secrets:` .
- **Piège n° 8 – Ne pas pin les images.**`postgres:latest` peut casser une migration majeure du jour au lendemain. Pin sur`postgres:17.2-alpine` et planifiez les upgrades.
- **Piège n° 9 – Utiliser `network_mode: host` par confort.** Vous perdez l’isolation, le DNS interne et la portabilité. Ne le faites que pour les agents système (collecteurs réseau, etc.).
- **Piège n° 10 – Lancer `down -v` sans précaution.** Le flag`-v` détruit irrémédiablement les volumes nommés. Catastrophe garantie en prod.

## Dépannage : 8 Problèmes Fréquents et Leur Solution

| Symptôme | Cause Probable | Solution | 
|---|---|---|
| `Cannot connect to the Docker daemon` | Daemon arrêté ou utilisateur non dans le groupe `docker` | `sudo systemctl start docker` +`sudo usermod -aG docker $USER` puis se relogger | 
| `port is already allocated` | Un autre service écoute sur le port hôte | `sudo ss -tlnp \| grep 80` pour identifier ; changer le mapping ou tuer le service | 
| `service "api" failed to build: COPY failed` | Chemin Dockerfile incorrect ou `.dockerignore` trop large | Vérifier `build.context` et le contenu de`.dockerignore` | 
| `Database "appdb" does not exist` | Init script jamais exécuté (volume préexistant) | `docker compose down -v` (attention, perte de données) puis relancer | 
| Conteneur en boucle `Restarting` | Crash de l’application au démarrage | `docker compose logs --tail=100 service` pour la stack trace | 
| Healthcheck reste `unhealthy` | Curl ou wget absent de l’image Alpine | Ajouter `RUN apk add --no-cache curl` ou utiliser un test natif (`nc -z` ,`pg_isready` ) | 
| `No space left on device` | Volumes orphelins, builds intermédiaires | `docker system prune -a --volumes` (vérifier avant !) | 
| Variables `.env` ignorées | Mauvais emplacement du fichier | Le `.env` doit être à côté du`docker-compose.yml` , pas dans`./app/` | 

Pour un debug en profondeur, activez la verbosité maximale du moteur Compose : `COMPOSE_PROGRESS=plain docker compose --verbose up`. Vous verrez chaque étape (résolution du DNS, montage des volumes, exécution du healthcheck) en clair, ce qui révèle 90 % des problèmes en moins de cinq minutes.

## Optimisation Avancée : Bake, Caches et Builds Multi-Plateforme

Depuis 2025, Compose délègue les builds à **Docker Bake** par défaut, ce qui apporte la parallélisation, le cache distant et le multi-plateforme natif. Pour exploiter Bake explicitement, créez un fichier `docker-bake.hcl` en complément.

```
# docker-bake.hcl
group "default" {
  targets = ["api"]
}
target "api" {
  context  = "./app"
  platforms = ["linux/amd64", "linux/arm64"]
  tags      = ["registry.example.com/myapp/api:latest"]
  cache-from = ["type=registry,ref=registry.example.com/myapp/api:cache"]
  cache-to   = ["type=registry,ref=registry.example.com/myapp/api:cache,mode=max"]
}
```
```
# Build multi-arch poussé directement vers le registre
docker buildx bake --push
# Combinaison Compose + Bake
COMPOSE_BAKE=true docker compose build  # Active Bake explicitement
```
Pour réduire encore le temps de build, activez le **BuildKit cache mount** dans votre Dockerfile : `RUN --mount=type=cache,target=/root/.cache/pip pip install -r requirements.txt`. Le cache pip est conservé entre deux builds, divisant le temps par 5 en moyenne. Pour Node.js, ciblez `/root/.npm` ; pour Go, `/go/pkg/mod`.

## Comparatif Docker Compose vs Alternatives en 2026

| Outil | Format | Cas d’usage idéal | Courbe d’apprentissage | 
|---|---|---|---|
| **Docker Compose v2** | YAML (Compose Spec) | Dev local, CI, mono-hôte prod | Faible (1-2 j) | 
| Docker Swarm | YAML Compose étendu | Cluster Docker 2-10 nœuds | Moyenne (1 sem) | 
| Kubernetes | YAML manifests | Production multi-nœuds, > 20 services | Élevée (1-3 mois) | 
| Podman Compose | YAML compatible Compose | Environnements rootless, RHEL/CentOS | Faible (1-2 j) | 
| Nomad | HCL | Workloads mixtes (Docker, JVM, binaires) | Moyenne | 
| Nerdctl Compose | YAML compatible Compose | Containerd direct, sans Docker | Faible | 

Si votre stack dépasse une vingtaine de conteneurs ou nécessite du *scaling* automatique, migrez sur Kubernetes : l’outil Kompose convertit un fichier Compose en manifests Kubernetes en une seule commande. Pour une transition douce, jetez aussi un œil à notre tutoriel ArgoCD GitOps.

## Variables d’Environnement et Fichier .env

Compose lit automatiquement un fichier `.env` à la racine du projet et substitue toute occurrence `${VAR}` dans le YAML. C’est l’idéal pour *parameterize* versions d’image, ports et tags sans toucher au fichier principal.

```
# .env (à la racine, jamais commité)
POSTGRES_VERSION=17-alpine
REDIS_VERSION=7.4-alpine
NGINX_VERSION=1.27-alpine
API_PORT=8000
COMPOSE_PROJECT_NAME=monprojet
```
Dans le YAML : `image: postgres:${POSTGRES_VERSION}`. Pour passer ponctuellement une variable depuis le shell, exportez-la avant : `POSTGRES_VERSION=17.2-alpine docker compose up -d`. La variable shell prend la priorité sur celle du `.env`.

La variable spéciale `COMPOSE_PROJECT_NAME` remplace le préfixe par défaut des conteneurs et volumes. `COMPOSE_FILE` permet de lister plusieurs fichiers (séparés par `:` sous Linux/macOS, `;` sous Windows). `COMPOSE_PROFILES` active des profils sans flag CLI.

## FAQ : Questions Fréquentes sur Docker Compose

### Docker Compose est-il toujours gratuit ?

Oui. Le plugin `docker compose` est sous licence Apache 2.0 et reste librement utilisable. Seul Docker Desktop (l’environnement clé en main pour macOS/Windows) requiert un abonnement payant pour les entreprises de plus de 250 employés ou de plus de 10 millions de dollars de revenus annuels. Sous Linux, tout est gratuit.

### Quelle est la différence entre `docker-compose` et `docker compose` ?

`docker-compose` (avec tiret) est la **v1** écrite en Python, dépréciée depuis juin 2023 et plus du tout maintenue. `docker compose` (sans tiret) est la **v2**, écrite en Go, livrée comme plugin officiel de la CLI Docker. Elle est 3 à 4 fois plus rapide, gère mieux les builds parallèles et est la seule à recevoir les correctifs de sécurité.

### Compose remplace-t-il Kubernetes ?

Non. Compose orchestre des conteneurs sur *un seul hôte* (ou un petit cluster Swarm). Kubernetes orchestre des conteneurs sur des centaines de nœuds, avec auto-scaling, self-healing avancé, secrets distribués et networking complexe. Pour le développement local, Compose reste le choix optimal même dans une équipe Kubernetes – d’où son intégration native dans des outils comme Tilt et Skaffold.

### Comment exécuter une commande dans un conteneur en cours ?

