---
id: collect-261001-ia-llm/ia-llm/tutoriel-n8n-2026-workflow-ia-auto-heberge-en-13-etapes-2
title: "Verifier la version de Node.js (20 ou superieur requis)"
domain: ia-llm
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-ia-llm/tutoriel-n8n-2026-workflow-ia-auto-heberge-en-13-etapes.md
source_anchor: ""
source_lines: [57, 187]
sha256: 4c88f867e0db58da1d50b5f43443b7046b49b3ca9a467137942a241a6928dd61
---

# Verifier la version de Node.js (20 ou superieur requis)

Si vous ne souhaitez pas utiliser Docker, n8n s’installe directement via **npm**, le gestionnaire de paquets de Node.js. Cette méthode convient pour un test rapide sur une machine de développement, mais Docker reste préférable en production pour l’isolation et la reproductibilité. Assurez-vous d’avoir **Node.js 20 LTS** ou une version supérieure, puis installez le paquet globalement :

```
# Verifier la version de Node.js (20 ou superieur requis)
$ node --version
v20.18.0
# Installer n8n globalement via npm
$ npm install -g n8n
# Demarrer le serveur n8n
$ n8n start
Editor is now accessible via: http://localhost:5678/
```
Une autre option, encore plus rapide pour un test ponctuel, consiste à utiliser `npx`, qui exécute n8n sans installation permanente : `npx n8n`. Toutefois, gardez à l’esprit que cette méthode npm utilise par défaut la base **SQLite** stockée dans le dossier `~/.n8n`. Pour toute utilisation sérieuse, configurez les mêmes variables d’environnement PostgreSQL que dans notre exemple Docker. La méthode npm impose aussi de gérer manuellement les mises à jour (`npm update -g n8n`) et la supervision du processus, là où Docker simplifie grandement le cycle de vie.

Quelle que soit la méthode retenue, le principe reste identique : un serveur n8n écoute sur le port **5678** et expose l’éditeur web ainsi que l’API REST. La suite de ce tutoriel s’applique à l’identique, que vous ayez choisi Docker ou npm.

## Étapes 1 & 2 : Installer n8n avec Docker (port 5678)

**Étape 1.** La façon la plus rapide de tester n8n est de lancer un conteneur unique. Cette commande télécharge l’image officielle `n8nio/n8n`, expose le port **5678** (le port web par défaut de n8n) et crée un volume nommé pour conserver vos données entre les redémarrages :

```
# Créer un volume persistant pour ne pas perdre vos workflows
$ docker volume create n8n_data
# Lancer n8n en arrière-plan sur le port 5678
$ docker run -d \
  --name n8n \
  -p 5678:5678 \
  -v n8n_data:/home/node/.n8n \
  -e GENERIC_TIMEZONE="Europe/Paris" \
  -e TZ="Europe/Paris" \
  docker.n8n.io/n8nio/n8n
```
**Étape 2.** Vérifiez que le conteneur tourne et consultez ses journaux. Vous devez voir une ligne indiquant que l’éditeur est accessible :

```
$ docker ps --filter "name=n8n"
CONTAINER ID   IMAGE                       STATUS         PORTS
a3f1c9e8b2d4   docker.n8n.io/n8nio/n8n      Up 12 seconds  0.0.0.0:5678->5678/tcp
$ docker logs n8n --tail 5
Editor is now accessible via:
http://localhost:5678/
Version: 2.25.x
```
Ouvrez votre navigateur à l’adresse `http://localhost:5678`. Vous arrivez sur l’écran de création du compte propriétaire (owner). Cette configuration mono-conteneur convient parfaitement à l’apprentissage, mais elle stocke les données dans une base **SQLite** intégrée – insuffisante pour la production. Nous corrigerons cela dès l’étape suivante avec PostgreSQL.

## Étapes 3 & 4 : Configurer PostgreSQL et la persistance avec Docker Compose

**Étape 3.** Pour un déploiement sérieux, n8n recommande **PostgreSQL** plutôt que SQLite : la base gère mieux la concurrence et les gros volumes d’exécutions. Créez un dossier de projet et un fichier `docker-compose.yml` orchestrant les deux services :

```
services:
  postgres:
    image: postgres:16
    restart: always
    environment:
      - POSTGRES_USER=n8n
      - POSTGRES_PASSWORD=motdepasse_solide
      - POSTGRES_DB=n8n
    volumes:
      - pg_data:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U n8n -d n8n"]
      interval: 10s
      timeout: 5s
      retries: 5
  n8n:
    image: docker.n8n.io/n8nio/n8n
    restart: always
    ports:
      - "5678:5678"
    environment:
      - DB_TYPE=postgresdb
      - DB_POSTGRESDB_HOST=postgres
      - DB_POSTGRESDB_DATABASE=n8n
      - DB_POSTGRESDB_USER=n8n
      - DB_POSTGRESDB_PASSWORD=motdepasse_solide
      - N8N_HOST=localhost
      - GENERIC_TIMEZONE=Europe/Paris
      - N8N_ENCRYPTION_KEY=remplacez_par_une_cle_aleatoire_32_car
    volumes:
      - n8n_data:/home/node/.n8n
    depends_on:
      postgres:
        condition: service_healthy
volumes:
  pg_data:
  n8n_data:
```
**Étape 4.** La variable `N8N_ENCRYPTION_KEY` est cruciale : elle chiffre vos identifiants (clés API, mots de passe) stockés dans n8n. Générez une clé aléatoire de 32 caractères et lancez la pile :

```
# Générer une clé de chiffrement robuste
$ openssl rand -hex 16
8f3a1d7c9b2e4f6a0c5d8e1b3a7f9c2d
# Lancer n8n + PostgreSQL
$ docker compose up -d
# Suivre le démarrage
$ docker compose logs -f n8n
```
Conservez précieusement cette clé : si vous la perdez, vos identifiants enregistrés deviennent illisibles après une réinstallation. Stockez-la dans un gestionnaire de secrets, jamais en clair dans un dépôt Git.

## Étape 5 : Premier accès à l’éditeur et création du compte propriétaire

**Étape 5.** Rendez-vous sur `http://localhost:5678`. n8n vous demande de créer le compte **propriétaire** : adresse e-mail, prénom, nom et mot de passe. Ce compte dispose des droits administrateur complets sur l’instance. Choisissez un mot de passe fort – cette interface donnera bientôt accès à vos clés API et à vos données métier.

Une fois connecté, vous découvrez le **canevas** (canvas) de l’éditeur n8n. C’est ici que vous assemblerez vos nœuds. L’interface comprend trois zones clés :

- **Le canevas central** : où vous déposez et reliez les nœuds par glisser-déposer.
- **Le panneau d’ajout de nœud** (touche`Tab` ou bouton`+` ) : la bibliothèque de plus de 400 intégrations.
- **Le panneau d’exécution** : affiche les données entrantes et sortantes de chaque nœud après un test, indispensable pour le débogage.

Familiarisez-vous avec le raccourci `Tab` : il ouvre instantanément le sélecteur de nœuds. La notion centrale à comprendre est que chaque nœud reçoit des **items** (des objets JSON) en entrée et produit des items en sortie. Cette structure de données circule de gauche à droite tout au long du workflow – c’est le modèle d’exécution en graphe de n8n.

## Étapes 6 & 7 : Créer le workflow – déclencheur et nœud HTTP Request

**Étape 6.** Tout workflow n8n commence par un **nœud déclencheur** (trigger). Pour notre projet – un agent de veille qui résume automatiquement les actualités tech – nous utiliserons le déclencheur **Schedule Trigger**, qui lance le workflow à intervalle régulier. Cliquez sur « Add first step », choisissez « On a schedule » et réglez-le pour s’exécuter chaque matin à 8 h.

n8n propose plusieurs familles de déclencheurs : **Webhook** (réagir à une requête HTTP entrante), **Schedule** (cron), ou des déclencheurs applicatifs (nouveau message Slack, nouvelle ligne Google Sheets, etc.). Pour tester un webhook, n8n génère une URL unique de la forme `http://localhost:5678/webhook/<id>` que vous pouvez appeler avec curl.

**Étape 7.** Ajoutez ensuite un nœud **HTTP Request** pour récupérer des données externes. Dans notre exemple, nous interrogeons l’API publique Hacker News pour obtenir les meilleures actualités. Configurez le nœud ainsi : méthode `GET`, URL de l’API, et n8n renverra le JSON sous forme d’items exploitables.

```
# Configuration du nœud HTTP Request
Method: GET
URL: https://hacker-news.firebaseio.com/v0/topstories.json
# Pour tester l'équivalent en ligne de commande :
$ curl -s "https://hacker-news.firebaseio.com/v0/topstories.json" | head -c 120
[42891234,42891200,42890987,42890850, ... ]
```
Cliquez sur « Execute step » : le panneau de droite affiche la réponse. Vous voyez un tableau d’identifiants d’articles. C’est le moment clé du débogage dans n8n – vous inspectez visuellement la **structure exacte des données** à chaque étape, ce qui rend la construction de workflows complexes beaucoup plus intuitive qu’avec du code pur.

