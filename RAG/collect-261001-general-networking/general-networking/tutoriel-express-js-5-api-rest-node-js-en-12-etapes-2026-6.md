---
id: collect-261001-general-networking/general-networking/tutoriel-express-js-5-api-rest-node-js-en-12-etapes-2026-6
title: "v22.x.x (minimum v18)"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-general-networking/tutoriel-express-js-5-api-rest-node-js-en-12-etapes-2026.md
source_anchor: ""
source_lines: [926, 1028]
sha256: a001e80044c5a579f5c217348312b0a948f29fb616da7249b43ffe6449f6ac14
---

# v22.x.x (minimum v18)

**Logging structuré avec Pino.** Remplacez Morgan par Pino pour un logging JSON structuré en production. Pino est jusqu’à 5 fois plus rapide que Winston et produit des logs directement exploitables par des outils comme Elasticsearch ou Grafana Loki.

**Graceful shutdown.** Implémentez un arrêt propre du serveur pour terminer les requêtes en cours avant de fermer les connexions MongoDB. C’est essentiel pour les déploiements Kubernetes avec des rolling updates.

```
// Ajouter dans src/server.js
const gracefulShutdown = async (signal) => {
  console.log(`${signal} reçu. Arrêt gracieux en cours...`);
  server.close(async () => {
    await mongoose.connection.close();
    console.log('Connexions fermées. Processus terminé.');
    process.exit(0);
  });
  // Forcer l'arrêt après 10 secondes
  setTimeout(() => {
    console.error('Arrêt forcé après timeout');
    process.exit(1);
  }, 10000);
};
process.on('SIGTERM', () => gracefulShutdown('SIGTERM'));
process.on('SIGINT', () => gracefulShutdown('SIGINT'));
```
**Versioning de l’API.** Préfixez vos routes avec un numéro de version (`/api/v1/products`) dès le départ. Cela permet d’introduire des changements breaking dans une nouvelle version (`/api/v2/products`) sans casser les clients existants.

## Structure Complète du Projet Express.js

Voici l’arborescence finale du projet complet tel que construit dans ce tutoriel. Cette structure suit les conventions de l’industrie pour une API REST Express.js maintenable et scalable.

```
express-api-rest/
├── src/
│   ├── config/
│   │   └── database.js          # Connexion MongoDB
│   ├── controllers/
│   │   ├── authController.js    # Inscription, connexion, profil
│   │   └── productController.js # CRUD produits
│   ├── middleware/
│   │   ├── auth.js              # JWT protect & authorize
│   │   ├── errorHandler.js      # Gestion centralisée des erreurs
│   │   └── validators.js        # Validation express-validator
│   ├── models/
│   │   ├── User.js              # Schéma utilisateur + hash bcrypt
│   │   └── Product.js           # Schéma produit + index
│   ├── routes/
│   │   ├── authRoutes.js        # Routes /api/auth
│   │   └── productRoutes.js     # Routes /api/products
│   ├── app.js                   # Configuration Express
│   └── server.js                # Point d'entrée + graceful shutdown
├── tests/
│   └── product.test.js          # Tests Jest + Supertest
├── .env                         # Variables d'environnement
├── .gitignore                   # Fichiers exclus du versioning
├── Dockerfile                   # Image Docker optimisée
├── docker-compose.yml           # Orchestration API + MongoDB
└── package.json                 # Dépendances et scripts
```
Ce projet contient environ 500 lignes de code réparties sur 12 fichiers. Chaque composant a une responsabilité unique : les modèles gèrent la structure des données, les contrôleurs la logique métier, les middleware les préoccupations transversales (authentification, validation, erreurs), et les routes le mapping HTTP.

## Couverture Connexe

### Articles Connexes sur tech-insider.org

Pour approfondir vos compétences en développement backend et DevOps, consultez ces guides complémentaires :

- Tutoriel FastAPI Python 2026 : Créer une API REST Complète — Comparez l’approche Python avec Express.js pour le développement d’API REST.
- Tutoriel Docker pour Débutants — Maîtrisez la conteneurisation pour déployer votre API Express.js en production.
- Tutoriel Docker Compose : Applications Multi-Conteneurs — Orchestrez Express.js et MongoDB avec Docker Compose.
- Tutoriel GraphQL avec Node.js et Apollo Server — Explorez une alternative REST avec GraphQL sur Node.js.
- Tutoriel Prometheus Grafana : Monitoring Complet — Surveillez les performances de votre API Express.js en production.
- Tutoriel GitHub Actions CI/CD — Automatisez les tests et le déploiement de votre API.

## FAQ : Questions Fréquentes sur Express.js

### Express.js est-il toujours pertinent en 2026 ?

Express.js reste le framework Node.js le plus utilisé au monde, avec environ 69 000 étoiles et 23 400 forks sur GitHub ainsi que plus de 102,4 millions de téléchargements hebdomadaires sur npm en février 2026 selon SoloDevStack. Avec la sortie d’Express 5, le framework bénéficie du support natif des Promises, d’une sécurité renforcée contre les attaques ReDoS via path-to-regexp v8, et d’un calendrier LTS officiel. Son écosystème de middleware et sa communauté active en font un choix solide pour les API REST en production.

### Quelle est la différence entre Express.js et Fastify ?

Fastify est généralement plus rapide qu’Express.js en termes de requêtes par seconde grâce à son architecture basée sur les schémas JSON. Cependant, Express.js dispose d’un écosystème de middleware beaucoup plus large et d’une communauté plus vaste. Pour les projets nécessitant des performances brutes, Fastify est préférable. Pour les projets nécessitant flexibilité et compatibilité, Express.js reste le choix pragmatique.

### Dois-je migrer d’Express 4 vers Express 5 ?

Express 4 continue de recevoir des correctifs de sécurité (dernière version 4.22.1), mais Express 5.0.0 a été officiellement publiée en janvier 2025 — dix ans après la sortie d’Express 4 — selon InfoQ, qui a également rapporté à cette occasion l’abandon du support de Node.js antérieur à la v18. La migration vers Express 5 est recommandée pour les nouveaux projets et les projets existants qui bénéficieront du support natif des Promises et de la protection ReDoS. Testez votre code avec Express 5 dans un environnement de staging avant de migrer la production.

### Comment gérer le CORS en production avec Express.js ?

Configurez le middleware `cors` avec une liste blanche d’origines autorisées. N’utilisez jamais `origin: '*'` en production si votre API utilise des cookies ou des credentials. Utilisez une variable d’environnement pour définir les origines autorisées : `CORS_ORIGIN=https://monapp.fr`.

### Quelle base de données utiliser avec Express.js ?

MongoDB avec Mongoose est le choix le plus courant dans l’écosystème Node.js (stack MERN/MEAN). Pour les données relationnelles, PostgreSQL avec Prisma ou Drizzle ORM est recommandé. Le choix dépend de la structure de vos données : MongoDB pour les documents flexibles, PostgreSQL pour les relations complexes et les transactions ACID.

### Comment sécuriser une API Express.js en production ?

Appliquez ces couches de sécurité : helmet pour les en-têtes HTTP, rate limiting pour prévenir les attaques par force brute, express-mongo-sanitize contre les injections NoSQL, validation stricte des entrées avec express-validator, authentification JWT avec des tokens courts, et HTTPS obligatoire. Maintenez Express.js et ses dépendances à jour, notamment path-to-regexp qui a reçu trois correctifs de sécurité en mars 2026.

### Express.js supporte-t-il TypeScript ?

Express.js 5 fonctionne avec TypeScript via les types `@types/express`. Vous pouvez écrire vos contrôleurs et middleware en TypeScript et compiler vers JavaScript avec `tsc` ou utiliser `tsx` pour l’exécution directe en développement. La définition de types personnalisés pour `req.user` via la déclaration de module améliore significativement l’expérience développeur.

### Comment déployer Express.js sur le cloud ?

