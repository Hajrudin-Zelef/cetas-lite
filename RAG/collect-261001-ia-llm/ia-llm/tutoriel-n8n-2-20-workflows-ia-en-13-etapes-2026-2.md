---
id: collect-261001-ia-llm/ia-llm/tutoriel-n8n-2-20-workflows-ia-en-13-etapes-2026-2
title: "Nœud Code – Langage Python"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "parameters", "sandbox"]
source: docs/RAG/collect-261001-ia-llm/tutoriel-n8n-2-20-workflows-ia-en-13-etapes-2026.md
source_anchor: ""
source_lines: [49, 199]
sha256: 757f15d8a1c6a6172632b17617552dee68cfedbd72b94e035f83decd5c11a3d1
---

# Nœud Code – Langage Python

```
mkdir -p ~/n8n-prod/data ~/n8n-prod/postgres
cd ~/n8n-prod
cat > docker-compose.yml << 'EOF'
version: "3.9"
services:
  postgres:
    image: postgres:17-alpine
    restart: unless-stopped
    environment:
      POSTGRES_DB: n8n
      POSTGRES_USER: n8nuser
      POSTGRES_PASSWORD: motdepasse_fort_2026
    volumes:
      - ./postgres:/var/lib/postgresql/data
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U n8nuser -d n8n"]
      interval: 10s
      timeout: 5s
      retries: 5
  n8n:
    image: n8nio/n8n:2.20.9
    restart: unless-stopped
    ports:
      - "5678:5678"
    environment:
      DB_TYPE: postgresdb
      DB_POSTGRESDB_HOST: postgres
      DB_POSTGRESDB_PORT: 5432
      DB_POSTGRESDB_DATABASE: n8n
      DB_POSTGRESDB_USER: n8nuser
      DB_POSTGRESDB_PASSWORD: motdepasse_fort_2026
      N8N_HOST: localhost
      N8N_PROTOCOL: http
      WEBHOOK_URL: http://localhost:5678/
      GENERIC_TIMEZONE: Europe/Paris
      N8N_ENCRYPTION_KEY: cle_chiffrement_32_caracteres_min
    volumes:
      - ./data:/home/node/.n8n
    depends_on:
      postgres:
        condition: service_healthy
EOF
docker compose up -d
```
La commande `docker compose up -d` télécharge les images officielles depuis Docker Hub puis lance les conteneurs en arrière-plan. Comptez environ deux minutes selon votre bande passante. Vérifiez ensuite que les deux services tournent avec `docker compose ps`. Vous devriez voir les conteneurs `postgres` et `n8n` en état *healthy*.

Le choix de **PostgreSQL 17** comme stockage est crucial dès la production. SQLite, utilisé par défaut, plafonne à quelques dizaines de workflows actifs avant de provoquer des verrouillages. PostgreSQL gère sans difficulté plusieurs centaines de workflows concurrents, comme le documente le guide de scaling officiel.

## Étape 2 : Configurer l’éditeur et créer votre compte propriétaire

Ouvrez votre navigateur sur `http://localhost:5678`. n8n vous demande à la première connexion de créer un compte propriétaire (owner). Renseignez un e-mail valide, un mot de passe robuste de 8 caractères minimum et un nom d’affichage. Cette étape configure l’instance en mode multi-utilisateur, indispensable pour partager des workflows en équipe.

Une fois connecté, n8n propose un tour guidé. Survolez la barre latérale gauche pour identifier les sections clés : **Workflows**, **Credentials**, **Executions** et **Settings**. Le panneau central affiche le canevas où vous glisserez les nœuds. Le panneau droit liste les centaines de nœuds intégrés, classés par catégorie : Action, Trigger, Transform, AI.

Pensez à activer la **télémétrie anonyme** dans Settings > Usage and plan si vous souhaitez recevoir les notifications de mise à jour. Pour les environnements sensibles RGPD, désactivez-la via la variable d’environnement `N8N_DIAGNOSTICS_ENABLED=false` dans votre `docker-compose.yml`.

## Étape 3 : Créer votre premier workflow avec un nœud Schedule

Cliquez sur **Add workflow** dans le tableau de bord. Renommez le workflow en *Veille quotidienne*. Le canevas est vide ; un bouton circulaire **+** au centre permet d’ajouter le premier nœud. Sélectionnez **Schedule Trigger** dans la section Triggers.

Le nœud Schedule remplace l’ancien nœud Cron de n8n 1.x. Il offre une interface plus claire pour les cadences récurrentes. Configurez-le pour s’exécuter **tous les jours à 8h00 Europe/Paris** ; le tutoriel officiel « Build your first workflow », mis à jour le 26 juin 2026 par n8n Docs, illustre d’ailleurs une variante hebdomadaire à 9h00 (minute 0) pour ce même nœud, si vous préférez une cadence moins fréquente. Le nœud se serre automatiquement à la timezone définie globalement via `GENERIC_TIMEZONE` dans le compose.

```
{
  "trigger": "scheduleTrigger",
  "rule": {
    "interval": [
      {
        "field": "cronExpression",
        "expression": "0 8 * * *"
      }
    ]
  },
  "timezone": "Europe/Paris"
}
```
Ajoutez ensuite un nœud **HTTP Request** connecté au Schedule. Pointez l’URL vers `https://hacker-news.firebaseio.com/v0/topstories.json` en GET. Cliquez sur *Test step* : n8n récupère immédiatement la liste des identifiants des articles populaires. Ce premier workflow démontre la simplicité du modèle nœud-à-nœud où chaque sortie devient l’entrée du nœud suivant.

## Étape 4 : Maîtriser le nœud HTTP Request avec authentification

Le nœud HTTP Request est probablement le plus utilisé de n8n. La version 2026 introduit un sélecteur d’authentification amélioré avec prise en charge native d’OAuth 2.0, JWT, mTLS et signature HMAC. Pour configurer un appel authentifié, créez d’abord une **credential** dans Settings > Credentials.

Prenons un exemple concret : récupérer les dépôts publics d’un utilisateur GitHub via l’API REST avec un Personal Access Token. Créez une credential de type *HTTP Header Auth*, nommez-la *github-pat*, et renseignez le nom `Authorization` avec la valeur `Bearer ghp_votre_token`.

```
{
  "node": "HTTP Request",
  "method": "GET",
  "url": "https://api.github.com/users/n8n-io/repos",
  "authentication": "predefinedCredentialType",
  "nodeCredentialType": "httpHeaderAuth",
  "options": {
    "response": {
      "response": {
        "responseFormat": "json"
      }
    },
    "queryParameters": {
      "parameters": [
        { "name": "per_page", "value": "30" },
        { "name": "sort", "value": "updated" }
      ]
    }
  }
}
```
Connectez ce nœud à un nœud **Split In Batches** pour itérer sur chaque dépôt sans saturer la mémoire. C’est l’équivalent du paradigme *fan-out* de Apache Airflow, mais sans Python ni DAG explicite. n8n gère automatiquement la pagination si vous activez l’option *Pagination > Response Contains Next URL*.

## Étape 5 : Manipuler les données avec le nœud Code (JavaScript et Python)

Le nœud Code est ce qui distingue n8n de ses concurrents low-code traditionnels. Vous pouvez y écrire du JavaScript natif (V8 isolate) ou du Python (Pyodide). Le contexte d’exécution expose l’objet `items`, tableau des entrées du nœud précédent, et attend en retour un tableau d’objets `{ json: {...} }`.

Exemple JavaScript pour filtrer et enrichir des dépôts GitHub : ne garder que ceux ayant plus de 1 000 étoiles, ajouter une colonne `density` qui calcule le ratio étoiles / taille.

```
// Nœud Code – Mode "Run Once for All Items"
const seuil = 1000;
const enrichis = items
  .filter(item => item.json.stargazers_count >= seuil)
  .map(item => {
    const r = item.json;
    return {
      json: {
        name: r.name,
        stars: r.stargazers_count,
        forks: r.forks_count,
        size_kb: r.size,
        density: r.size > 0 ? (r.stargazers_count / r.size).toFixed(3) : 0,
        last_push: r.pushed_at,
        url: r.html_url
      }
    };
  })
  .sort((a, b) => b.json.stars - a.json.stars);
return enrichis;
```
Pour Python, la syntaxe diffère légèrement : la variable globale s’appelle `_input` et la sortie doit être assignée à `_output`. La mise à jour récente du nœud Code, documentée par Latenode, étend significativement les bibliothèques disponibles dans le sandbox Pyodide. Vous pouvez désormais utiliser **pandas**, **numpy** et **requests** directement, ce qui ouvre la voie à des transformations de données auparavant réservées à DuckDB ou Pandas.

```
# Nœud Code – Langage Python
import pandas as pd
df = pd.DataFrame([item['json'] for item in _input.all()])
df['density'] = df['stargazers_count'] / df['size'].clip(lower=1)
top = df.nlargest(10, 'stargazers_count')[
    ['name', 'stargazers_count', 'forks_count', 'density', 'html_url']
]
_output = [{'json': row} for row in top.to_dict(orient='records')]
```
## Étape 6 : Connecter PostgreSQL et persister vos résultats

