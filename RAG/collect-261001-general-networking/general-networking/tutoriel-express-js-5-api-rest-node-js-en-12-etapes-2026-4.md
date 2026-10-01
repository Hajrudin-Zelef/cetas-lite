---
id: collect-261001-general-networking/general-networking/tutoriel-express-js-5-api-rest-node-js-en-12-etapes-2026-4
title: "v22.x.x (minimum v18)"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/tutoriel-express-js-5-api-rest-node-js-en-12-etapes-2026.md
source_anchor: ""
source_lines: [561, 762]
sha256: 2b537cf051d887b85f1eb3d60545ef5c108d59eef9cd5e29594722058f5d5a7f
---

# v22.x.x (minimum v18)

`npm install express-validator````
// src/middleware/validators.js
import { body, validationResult } from 'express-validator';
export const handleValidation = (req, res, next) => {
  const errors = validationResult(req);
  if (!errors.isEmpty()) {
    res.status(400);
    throw new Error(
      errors.array().map(e => e.msg).join(', ')
    );
  }
  next();
};
export const validateRegister = [
  body('name')
    .trim()
    .notEmpty().withMessage('Le nom est requis')
    .isLength({ min: 2, max: 50 }).withMessage('Le nom doit contenir entre 2 et 50 caractères'),
  body('email')
    .trim()
    .notEmpty().withMessage('L\'email est requis')
    .isEmail().withMessage('Format d\'email invalide')
    .normalizeEmail(),
  body('password')
    .notEmpty().withMessage('Le mot de passe est requis')
    .isLength({ min: 8 }).withMessage('Le mot de passe doit contenir au moins 8 caractères')
    .matches(/\d/).withMessage('Le mot de passe doit contenir au moins un chiffre')
    .matches(/[A-Z]/).withMessage('Le mot de passe doit contenir au moins une majuscule'),
  handleValidation
];
export const validateProduct = [
  body('name')
    .trim()
    .notEmpty().withMessage('Le nom du produit est requis')
    .isLength({ max: 100 }).withMessage('Le nom ne peut pas dépasser 100 caractères'),
  body('description')
    .trim()
    .notEmpty().withMessage('La description est requise'),
  body('price')
    .isFloat({ min: 0 }).withMessage('Le prix doit être un nombre positif'),
  body('category')
    .isIn(['electronique', 'vetements', 'alimentation', 'sport', 'maison'])
    .withMessage('Catégorie invalide'),
  body('stock')
    .optional()
    .isInt({ min: 0 }).withMessage('Le stock doit être un entier positif'),
  handleValidation
];
```
Intégrez les validateurs dans les routes :

```
// Modifier src/routes/authRoutes.js
import { validateRegister } from '../middleware/validators.js';
router.post('/register', validateRegister, register);
// Modifier src/routes/productRoutes.js
import { validateProduct } from '../middleware/validators.js';
router.route('/')
  .get(getProducts)
  .post(protect, validateProduct, createProduct);
```
**Piège n°6 :** La validation doit toujours intervenir AVANT le contrôleur dans la chaîne de middleware. Placer la validation après le contrôleur est inutile car les données auront déjà été traitées.

## Étape 10 : Ajouter le Rate Limiting et la Sécurité

En mars 2026, trois vulnérabilités de sécurité ont été corrigées dans path-to-regexp, la bibliothèque de routage utilisée par Express.js, dont CVE-2026-4867 (critique) liée à une croissance exponentielle des regex. Le rate limiting et les mesures de sécurité avancées sont essentiels pour protéger votre API.

`npm install express-rate-limit hpp express-mongo-sanitize````
// Ajouter dans src/app.js
import rateLimit from 'express-rate-limit';
import mongoSanitize from 'express-mongo-sanitize';
import hpp from 'hpp';
// Rate limiting global
const limiter = rateLimit({
  windowMs: 15 * 60 * 1000, // 15 minutes
  max: 100, // 100 requêtes par IP
  standardHeaders: true,
  legacyHeaders: false,
  message: {
    success: false,
    message: 'Trop de requêtes. Réessayez dans 15 minutes.'
  }
});
// Rate limiting strict pour l'authentification
const authLimiter = rateLimit({
  windowMs: 15 * 60 * 1000,
  max: 5, // 5 tentatives de connexion par IP
  message: {
    success: false,
    message: 'Trop de tentatives de connexion. Réessayez dans 15 minutes.'
  }
});
app.use(limiter);
app.use('/api/auth/login', authLimiter);
app.use(mongoSanitize()); // Prévient les injections NoSQL
app.use(hpp()); // Prévient la pollution des paramètres HTTP
```
Le rate limiting à deux niveaux protège l’API contre les attaques par force brute. Le limiteur global autorise 100 requêtes par fenêtre de 15 minutes, tandis que le limiteur d’authentification restreint à 5 tentatives. `express-mongo-sanitize` supprime les opérateurs MongoDB (`$gt`, `$ne`) des entrées utilisateur pour prévenir les injections NoSQL.

| Couche de sécurité | Paquet | Protection | Impact performance | 
|---|---|---|---|
| En-têtes HTTP | helmet | 11 en-têtes de sécurité (CSP, HSTS, etc.) | Négligeable | 
| Rate limiting | express-rate-limit | Brute force, DDoS basique | Faible (stockage mémoire) | 
| Injection NoSQL | express-mongo-sanitize | Opérateurs MongoDB malveillants | Négligeable | 
| Pollution HTTP | hpp | Paramètres de requête dupliqués | Négligeable | 
| CORS | cors | Requêtes cross-origin non autorisées | Négligeable | 
| JWT | jsonwebtoken | Authentification stateless | Faible (vérification crypto) | 

## Étape 11 : Écrire les Tests avec Jest et Supertest

Les tests automatisés garantissent que votre API fonctionne correctement et permettent de détecter les régressions. Jest combiné à Supertest offre un framework de test complet pour les API Express.js.

```
// tests/product.test.js
import request from 'supertest';
import mongoose from 'mongoose';
import app from '../src/app.js';
import User from '../src/models/User.js';
import Product from '../src/models/Product.js';
let token;
let productId;
beforeAll(async () => {
  await mongoose.connect(process.env.MONGODB_URI_TEST ||
    'mongodb://localhost:27017/express-api-test');
  // Créer un utilisateur de test
  await User.deleteMany({});
  await Product.deleteMany({});
  const res = await request(app)
    .post('/api/auth/register')
    .send({
      name: 'Test User',
      email: '[email protected]',
      password: 'Password123'
    });
  token = res.body.data.token;
});
afterAll(async () => {
  await mongoose.connection.dropDatabase();
  await mongoose.connection.close();
});
describe('POST /api/products', () => {
  it('doit créer un produit avec un token valide', async () => {
    const res = await request(app)
      .post('/api/products')
      .set('Authorization', `Bearer ${token}`)
      .send({
        name: 'Laptop Pro',
        description: 'Un ordinateur portable performant',
        price: 1299.99,
        category: 'electronique',
        stock: 50
      });
    expect(res.status).toBe(201);
    expect(res.body.success).toBe(true);
    expect(res.body.data.name).toBe('Laptop Pro');
    productId = res.body.data._id;
  });
  it('doit refuser la création sans authentification', async () => {
    const res = await request(app)
      .post('/api/products')
      .send({
        name: 'Produit Test',
        description: 'Description test',
        price: 10,
        category: 'electronique'
      });
    expect(res.status).toBe(401);
  });
  it('doit refuser un produit avec un prix négatif', async () => {
    const res = await request(app)
      .post('/api/products')
      .set('Authorization', `Bearer ${token}`)
      .send({
        name: 'Produit',
        description: 'Description',
        price: -5,
        category: 'electronique'
      });
    expect(res.status).toBe(400);
  });
});
describe('GET /api/products', () => {
  it('doit retourner la liste paginée des produits', async () => {
    const res = await request(app)
      .get('/api/products?page=1&limit=10');
    expect(res.status).toBe(200);
    expect(res.body.success).toBe(true);
    expect(res.body.pagination).toBeDefined();
    expect(Array.isArray(res.body.data)).toBe(true);
  });
  it('doit filtrer par catégorie', async () => {
    const res = await request(app)
      .get('/api/products?category=electronique');
    expect(res.status).toBe(200);
    res.body.data.forEach(p => {
      expect(p.category).toBe('electronique');
    });
  });
});
```
Lancez les tests avec `npm test`. Sortie attendue :

