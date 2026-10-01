---
id: collect-261001-general-networking/general-networking/tutoriel-express-js-5-api-rest-node-js-en-12-etapes-2026-5
title: "v22.x.x (minimum v18)"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/tutoriel-express-js-5-api-rest-node-js-en-12-etapes-2026.md
source_anchor: ""
source_lines: [763, 925]
sha256: 225760114a79ebe1d65898a637ab14117e4e596d98130f8fc34c36bf6ff6fec0
---

# v22.x.x (minimum v18)

```
$ npm test
PASS  tests/product.test.js
  POST /api/products
    ✓ doit créer un produit avec un token valide (156 ms)
    ✓ doit refuser la création sans authentification (23 ms)
    ✓ doit refuser un produit avec un prix négatif (34 ms)
  GET /api/products
    ✓ doit retourner la liste paginée des produits (45 ms)
    ✓ doit filtrer par catégorie (28 ms)
Test Suites: 1 passed, 1 total
Tests:       5 passed, 5 total
Time:        2.341 s
```
## Étape 12 : Préparer le Déploiement en Production

Le déploiement d’une API Express.js en production nécessite plusieurs ajustements : compression des réponses, logging structuré, gestion du processus et configuration HTTPS. Voici la configuration finale de production.

`npm install compression morgan````
// Ajouter dans src/app.js pour la production
import compression from 'compression';
import morgan from 'morgan';
// Logging des requêtes HTTP
if (process.env.NODE_ENV === 'development') {
  app.use(morgan('dev'));
} else {
  app.use(morgan('combined'));
}
// Compression gzip des réponses
app.use(compression());
// Endpoint de readiness pour les orchestrateurs
app.get('/api/ready', async (req, res) => {
  const dbState = mongoose.connection.readyState;
  if (dbState === 1) {
    res.json({ status: 'ready', db: 'connected' });
  } else {
    res.status(503).json({ status: 'not ready', db: 'disconnected' });
  }
});
```
Pour le déploiement avec Docker, créez un `Dockerfile` optimisé :

```
# Dockerfile
FROM node:22-alpine AS base
WORKDIR /app
COPY package*.json ./
RUN npm ci --only=production
COPY src/ ./src/
ENV NODE_ENV=production
EXPOSE 3000
USER node
HEALTHCHECK --interval=30s --timeout=3s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:3000/api/health || exit 1
CMD ["node", "src/server.js"]
```
```
# docker-compose.yml
services:
  api:
    build: .
    ports:
      - "3000:3000"
    environment:
      - PORT=3000
      - MONGODB_URI=mongodb://mongo:27017/express-api
      - JWT_SECRET=${JWT_SECRET}
      - NODE_ENV=production
    depends_on:
      mongo:
        condition: service_healthy
    restart: unless-stopped
  mongo:
    image: mongo:8
    volumes:
      - mongo_data:/data/db
    healthcheck:
      test: ["CMD", "mongosh", "--eval", "db.adminCommand('ping')"]
      interval: 10s
      timeout: 5s
      retries: 5
volumes:
  mongo_data:
```
Lancez l’application complète avec `docker compose up -d`. Le healthcheck intégré garantit que Docker redémarrera automatiquement le conteneur si l’API ne répond plus.

## Express.js 5 vs Express 4 : Les Différences Clés

La migration d’Express 4 vers Express 5 apporte des changements significatifs. Voici un résumé des différences les plus importantes pour les développeurs qui maintiennent des projets existants.

| Fonctionnalité | Express 4 | Express 5 | 
|---|---|---|
| Support des Promises | Non natif (try-catch requis) | Natif (rejet automatique capturé) | 
| Node.js minimum | 0.10+ | 18+ | 
| Routage (path-to-regexp) | v1.x (regex non sécurisé) | v8.x (protection ReDoS) | 
| app.del() | Disponible (alias) | Supprimé (utiliser app.delete()) | 
| req.host | Retourne hostname:port | Retourne uniquement hostname | 
| req.query | Setter disponible | Getter uniquement | 
| res.json(obj, status) | Signature à deux arguments | Supprimée (utiliser res.status().json()) | 
| Gestion d’erreurs async | Crash silencieux possible | Erreurs interceptées automatiquement | 

Le passage à path-to-regexp v8 est le changement de sécurité le plus important, d’autant plus critique qu’Express.js équipe désormais 138 238 domaines actifs en août 2026 selon TechnologyChecker, contre 90 377 seulement deux mois plus tôt en juin 2026 — une surface d’exposition en forte croissance. Les sous-expressions regex sont supprimées pour éliminer les vecteurs d’attaques ReDoS. Si vos routes Express 4 utilisent des patterns regex complexes, vous devrez les adapter lors de la migration vers Express 5.

## Dépannage : 8 Problèmes Courants et Solutions

Voici les problèmes les plus fréquemment rencontrés lors du développement avec Express.js 5, accompagnés de leurs solutions détaillées.

### 1. ERR_MODULE_NOT_FOUND : Cannot find module

**Cause :** Avec les ES modules, les extensions de fichier sont obligatoires dans les imports. `import app from './app'` ne fonctionne pas — vous devez écrire `import app from './app.js'`.

**Solution :** Ajoutez `.js` à tous vos imports locaux. Vérifiez que `"type": "module"` est présent dans votre `package.json`.

### 2. MongoServerError : E11000 duplicate key

**Cause :** Tentative d’insertion d’un document avec une valeur unique déjà existante (typiquement l’email). Cela arrive souvent pendant les tests quand la base n’est pas nettoyée entre les exécutions.

**Solution :** Utilisez `Model.deleteMany({})` dans le `beforeAll` de vos tests. En production, renvoyez un message d’erreur clair grâce au middleware errorHandler (code 11000).

### 3. TypeError : app.del is not a function

**Cause :** Express 5 a supprimé l’alias `app.del()`. Ce changement breaking affecte le code migré depuis Express 4.

**Solution :** Remplacez toutes les occurrences de `app.del()` par `app.delete()`.

### 4. Error : Route path must not include “?” (Express 5)

**Cause :** Express 5 utilise path-to-regexp v8 qui interdit les sous-expressions regex comme `/users/:id?` pour les paramètres optionnels.

**Solution :** Utilisez la syntaxe `{:id}` pour les paramètres optionnels : `/users{/:id}`.

### 5. JsonWebTokenError : jwt malformed

**Cause :** Le token envoyé dans l’en-tête Authorization n’est pas un JWT valide. Cela arrive quand le préfixe “Bearer ” est absent ou dupliqué.

**Solution :** Vérifiez le format exact : `Authorization: Bearer eyJhbGciOi...`. Utilisez `authHeader.split(' ')[1]` pour extraire le token.

### 6. CORS : Access-Control-Allow-Origin manquant

**Cause :** Le middleware CORS n’est pas configuré ou l’origine du frontend n’est pas dans la whitelist.

**Solution :** Configurez le middleware `cors` avec l’origine exacte de votre frontend. N’utilisez jamais `origin: '*'` en production avec des cookies ou des credentials.

### 7. EADDRINUSE : Port 3000 déjà utilisé

**Cause :** Un autre processus utilise déjà le port 3000. Cela arrive souvent quand nodemon redémarre pendant qu’une ancienne instance est encore active.

**Solution :** Identifiez le processus avec `lsof -i :3000` (Linux/Mac) ou `netstat -ano | findstr :3000` (Windows), puis terminez-le. Alternativement, changez le port dans le fichier `.env`.

### 8. Timeout de connexion MongoDB au démarrage

**Cause :** Le serveur MongoDB n’est pas démarré ou l’URI de connexion est incorrect. Avec Docker, le conteneur MongoDB peut ne pas être prêt au moment où l’API tente de se connecter.

**Solution :** Vérifiez que MongoDB est en cours d’exécution avec `mongosh --eval "db.version()"`. Avec Docker Compose, utilisez `depends_on` avec un healthcheck comme montré dans l’étape 12.

## Astuces Avancées pour Express.js 5 en Production

Une fois votre API fonctionnelle, ces techniques avancées amélioreront ses performances, sa maintenabilité et sa robustesse en production.

**Mise en cache avec Redis.** Pour les endpoints très sollicités, ajoutez une couche de cache Redis. Un middleware de cache Express peut réduire le temps de réponse de 200 ms à moins de 5 ms pour les données qui changent rarement.

**Pagination par curseur.** La pagination par offset (`skip/limit`) devient lente sur les grandes collections MongoDB car le serveur doit parcourir tous les documents précédents. La pagination par curseur utilise un `_id` de référence et est O(1) quelle que soit la page, idéale pour les collections de plus de 100 000 documents.

