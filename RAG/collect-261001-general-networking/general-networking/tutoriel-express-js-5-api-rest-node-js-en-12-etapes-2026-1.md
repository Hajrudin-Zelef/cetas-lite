---
id: collect-261001-general-networking/general-networking/tutoriel-express-js-5-api-rest-node-js-en-12-etapes-2026-1
title: "v22.x.x (minimum v18)"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/tutoriel-express-js-5-api-rest-node-js-en-12-etapes-2026.md
source_anchor: ""
source_lines: [1, 135]
sha256: d82c40064e785af79d4cdbf19136b38638872e3d8db6430ef570b28522e890e9
---

# v22.x.x (minimum v18)

Express.js reste le framework Node.js le plus utilisé au monde pour créer des API REST, avec environ 69 000 étoiles sur GitHub et 82 494 010 téléchargements sur npm pour la seule semaine du 6 juin 2026, selon les statistiques npmcharts. Avec la sortie d’Express 5.2.1 — la version stable la plus récente, publiée le 1er décembre 2025 — et une adoption qui continue de grimper (TechnologyChecker recense 138 238 domaines actifs utilisant Express.js en août 2026, contre 90 377 seulement en juin 2026), le support natif des Promises et une sécurité renforcée contre les attaques ReDoS font entrer le framework dans une nouvelle ère. Ce tutoriel vous guide pas à pas pour construire une API REST complète avec Express.js 5, de l’installation à la mise en production, en 12 étapes concrètes.

Que vous soyez développeur backend débutant ou que vous migriez depuis Express 4, ce guide couvre tout : routage, middleware, validation, authentification JWT, gestion d’erreurs, tests et déploiement. Chaque étape inclut du code fonctionnel, des exemples de sortie et les pièges courants à éviter. Temps estimé : 60 minutes.

## Prérequis et Versions Requises pour Express.js 5

Avant de commencer, assurez-vous de disposer de l’environnement suivant. Express 5 abandonne le support des versions de Node.js antérieures à la v18, ce qui constitue un changement majeur par rapport à Express 4.

| Outil | Version minimale | Version recommandée | Rôle | 
|---|---|---|---|
| Node.js | 18.x | 22.x LTS | Runtime JavaScript | 
| npm | 9.x | 10.x | Gestionnaire de paquets | 
| Express.js | 5.0.0 | 5.2.1 | Framework HTTP | 
| MongoDB | 7.0 | 8.0 | Base de données NoSQL | 
| VS Code ou éditeur | – | Dernière version | Éditeur de code | 
| Postman ou cURL | – | Dernière version | Test des endpoints | 

Express 5.2.1, publié le 1er décembre 2025, reste la version stable la plus récente et bénéficie d’un support de sécurité actif sur la branche 5.x, un statut confirmé à la fois par CompatHub et par la page Wikipedia consacrée à Express.js, qui référencent toutes deux cette même date de sortie. Fait notable, la version 5.2.0 est sortie le même jour en tant que release stable distincte selon CompatHub, avant d’être rapidement dépassée par le correctif 5.2.1, qui corrige un changement cassant introduit dans la v5.2.0 concernant le parser de requêtes étendues. Le support natif des Promises et le routage mis à jour via path-to-regexp v8 sont les deux améliorations majeures de cette version, qui compte désormais environ 23 400 forks sur GitHub selon SoloDevStack — un signe de son adoption large. Selon le calendrier LTS officiel d’Express.js, la branche 5.x restera maintenue au moins jusqu’en avril 2026, tandis que l’équipe prévoit une version 6.x au plus tôt en janvier 2026.

Vérifiez vos versions installées avant de continuer :

```
node --version
# v22.x.x (minimum v18)
npm --version
# 10.x.x
# Vérifier que MongoDB est en cours d'exécution
mongosh --eval "db.version()"
```
**Piège n°1 :** Si vous utilisez Node.js 16 ou antérieur, Express 5 refusera de s’installer. Mettez à jour vers Node.js 18 ou supérieur avec `nvm install 22` avant de continuer.

## Étape 1 : Initialiser le Projet et Installer Express.js

Créez un nouveau répertoire pour votre projet et initialisez-le avec npm. Nous utilisons l’option `--yes` pour générer un `package.json` par défaut, puis nous installerons Express 5 ainsi que les dépendances nécessaires.

```
# Créer et initialiser le projet
mkdir express-api-rest && cd express-api-rest
npm init --yes
# Installer Express 5 et les dépendances
npm install express@5 mongoose dotenv cors helmet jsonwebtoken bcryptjs
npm install --save-dev nodemon jest supertest
# Structure du projet
mkdir -p src/{routes,controllers,models,middleware,config,utils}
touch src/app.js src/server.js .env .gitignore
```
Express 5 est désormais la version par défaut sur npm depuis la sortie de la v5.1.0 le 31 mars 2025, qui a marqué le début de la phase ACTIVE du support de la branche 5.x selon le blog officiel d’Express.js — une phase toujours en cours à l’été 2026, le calendrier LTS d’Express.js ne prévoyant aucun changement de phase avant avril 2026. Vous n’avez plus besoin de spécifier `@next` comme c’était le cas pendant la phase de développement. Le fichier `package.json` doit inclure `"type": "module"` pour activer les imports ES modules.

Ajoutez ces lignes à votre `package.json` :

```
{
  "type": "module",
  "scripts": {
    "start": "node src/server.js",
    "dev": "nodemon src/server.js",
    "test": "node --experimental-vm-modules node_modules/.bin/jest"
  }
}
```
Configurez le fichier `.env` avec les variables d’environnement :

```
PORT=3000
MONGODB_URI=mongodb://localhost:27017/express-api
JWT_SECRET=votre_cle_secrete_ici_changez_en_production
JWT_EXPIRES_IN=24h
NODE_ENV=development
```
**Piège n°2 :** N’oubliez pas d’ajouter `.env` et `node_modules/` à votre `.gitignore`. Exposer vos clés secrètes JWT dans un dépôt Git est une faille de sécurité critique.

## Étape 2 : Configurer le Serveur Express.js 5 de Base

La configuration du serveur Express.js 5 diffère de la v4 sur plusieurs points. Le support natif des Promises élimine le besoin d’envelopper chaque handler dans un try-catch. Les middleware qui retournent des Promises rejetées sont automatiquement interceptés par le routeur.

```
// src/app.js
import express from 'express';
import cors from 'cors';
import helmet from 'helmet';
import dotenv from 'dotenv';
dotenv.config();
const app = express();
// Middleware de sécurité
app.use(helmet());
app.use(cors({
  origin: process.env.CORS_ORIGIN || 'http://localhost:3000',
  methods: ['GET', 'POST', 'PUT', 'DELETE', 'PATCH'],
  allowedHeaders: ['Content-Type', 'Authorization']
}));
// Parser JSON et URL-encoded
app.use(express.json({ limit: '10mb' }));
app.use(express.urlencoded({ extended: true }));
// Route de santé
app.get('/api/health', (req, res) => {
  res.json({
    status: 'ok',
    timestamp: new Date().toISOString(),
    environment: process.env.NODE_ENV
  });
});
export default app;
```
```
// src/server.js
import app from './app.js';
import { connectDB } from './config/database.js';
const PORT = process.env.PORT || 3000;
const startServer = async () => {
  await connectDB();
  app.listen(PORT, () => {
    console.log(`Serveur Express.js démarré sur le port ${PORT}`);
    console.log(`Environnement : ${process.env.NODE_ENV}`);
    console.log(`Santé : http://localhost:${PORT}/api/health`);
  });
};
startServer();
```
Le middleware `helmet` ajoute automatiquement 11 en-têtes HTTP de sécurité, dont `Content-Security-Policy`, `X-Content-Type-Options` et `Strict-Transport-Security`. C’est une pratique recommandée pour toute API Express.js en production.

Lancez le serveur avec `npm run dev` et testez la route de santé :

```
curl http://localhost:3000/api/health
# Sortie attendue :
{
  "status": "ok",
  "timestamp": "2026-04-10T10:30:00.000Z",
  "environment": "development"
}
```
## Étape 3 : Connexion à MongoDB avec Mongoose

MongoDB est la base de données la plus couramment associée à Express.js dans la stack MERN (MongoDB, Express, React, Node). Mongoose fournit une couche d’abstraction ODM (Object Document Mapping) qui simplifie la validation et les requêtes.

