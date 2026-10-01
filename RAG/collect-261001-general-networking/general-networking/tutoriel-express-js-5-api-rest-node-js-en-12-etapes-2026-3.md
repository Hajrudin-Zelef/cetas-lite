---
id: collect-261001-general-networking/general-networking/tutoriel-express-js-5-api-rest-node-js-en-12-etapes-2026-3
title: "v22.x.x (minimum v18)"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/tutoriel-express-js-5-api-rest-node-js-en-12-etapes-2026.md
source_anchor: ""
source_lines: [346, 560]
sha256: 864d027f1244f82cd09bcdb996b249212f3a3531f433c46a3cec4181edfcbb37
---

# v22.x.x (minimum v18)

```
// src/controllers/authController.js
import jwt from 'jsonwebtoken';
import User from '../models/User.js';
const generateToken = (userId) => {
  return jwt.sign({ id: userId }, process.env.JWT_SECRET, {
    expiresIn: process.env.JWT_EXPIRES_IN
  });
};
// POST /api/auth/register
export const register = async (req, res) => {
  const { name, email, password } = req.body;
  const existingUser = await User.findOne({ email });
  if (existingUser) {
    res.status(400);
    throw new Error('Un utilisateur avec cet email existe déjà');
  }
  const user = await User.create({ name, email, password });
  const token = generateToken(user._id);
  res.status(201).json({
    success: true,
    data: {
      user: { id: user._id, name: user.name, email: user.email, role: user.role },
      token
    }
  });
};
// POST /api/auth/login
export const login = async (req, res) => {
  const { email, password } = req.body;
  if (!email || !password) {
    res.status(400);
    throw new Error('Email et mot de passe requis');
  }
  const user = await User.findOne({ email }).select('+password');
  if (!user || !(await user.comparePassword(password))) {
    res.status(401);
    throw new Error('Email ou mot de passe incorrect');
  }
  const token = generateToken(user._id);
  res.json({
    success: true,
    data: {
      user: { id: user._id, name: user.name, email: user.email, role: user.role },
      token
    }
  });
};
// GET /api/auth/me
export const getMe = async (req, res) => {
  res.json({ success: true, data: req.user });
};
```
```
// src/middleware/auth.js
import jwt from 'jsonwebtoken';
import User from '../models/User.js';
export const protect = async (req, res, next) => {
  const authHeader = req.headers.authorization;
  if (!authHeader || !authHeader.startsWith('Bearer ')) {
    res.status(401);
    throw new Error('Accès non autorisé. Token manquant.');
  }
  const token = authHeader.split(' ')[1];
  const decoded = jwt.verify(token, process.env.JWT_SECRET);
  const user = await User.findById(decoded.id);
  if (!user) {
    res.status(401);
    throw new Error('Utilisateur associé au token introuvable');
  }
  req.user = user;
  next();
};
export const authorize = (...roles) => {
  return (req, res, next) => {
    if (!roles.includes(req.user.role)) {
      res.status(403);
      throw new Error('Accès interdit pour ce rôle');
    }
    next();
  };
};
```
Le middleware `protect` vérifie la présence et la validité du token JWT dans l’en-tête `Authorization`. Le middleware `authorize` fournit un contrôle d’accès basé sur les rôles (RBAC). Ces deux middleware peuvent être chaînés sur n’importe quelle route Express.js.

**Piège n°5 :** Ne stockez jamais les tokens JWT dans le `localStorage` côté client — c’est vulnérable aux attaques XSS. Préférez les cookies HttpOnly avec les flags `Secure` et `SameSite` pour les applications en production.

## Étape 7 : Définir les Routes de l’API REST

Express.js utilise un système de routage modulaire qui permet d’organiser les endpoints par ressource. Chaque fichier de routes est autonome et peut être monté sur un préfixe différent dans l’application principale.

```
// src/routes/authRoutes.js
import { Router } from 'express';
import { register, login, getMe } from '../controllers/authController.js';
import { protect } from '../middleware/auth.js';
const router = Router();
router.post('/register', register);
router.post('/login', login);
router.get('/me', protect, getMe);
export default router;
```
```
// src/routes/productRoutes.js
import { Router } from 'express';
import {
  getProducts,
  getProduct,
  createProduct,
  updateProduct,
  deleteProduct
} from '../controllers/productController.js';
import { protect, authorize } from '../middleware/auth.js';
const router = Router();
router.route('/')
  .get(getProducts)
  .post(protect, createProduct);
router.route('/:id')
  .get(getProduct)
  .put(protect, updateProduct)
  .delete(protect, authorize('admin'), deleteProduct);
export default router;
```
Montez les routes dans votre application principale en ajoutant ces lignes dans `src/app.js` :

```
// Ajouter dans src/app.js après les middleware existants
import authRoutes from './routes/authRoutes.js';
import productRoutes from './routes/productRoutes.js';
app.use('/api/auth', authRoutes);
app.use('/api/products', productRoutes);
```
La méthode `router.route()` permet de chaîner les handlers HTTP sur un même chemin, ce qui réduit la duplication et améliore la lisibilité. Les routes publiques (GET) ne nécessitent pas d’authentification, tandis que la suppression requiert le rôle admin.

| Méthode | Endpoint | Auth requise | Rôle | Description | 
|---|---|---|---|---|
| POST | /api/auth/register | Non | – | Inscription utilisateur | 
| POST | /api/auth/login | Non | – | Connexion et obtention du token | 
| GET | /api/auth/me | Oui | Tous | Profil de l’utilisateur connecté | 
| GET | /api/products | Non | – | Liste des produits avec pagination | 
| GET | /api/products/:id | Non | – | Détail d’un produit | 
| POST | /api/products | Oui | Tous | Créer un produit | 
| PUT | /api/products/:id | Oui | Tous | Mettre à jour un produit | 
| DELETE | /api/products/:id | Oui | Admin | Supprimer un produit | 

## Étape 8 : Gestion Centralisée des Erreurs Express.js 5

Express 5 révolutionne la gestion des erreurs avec un flux prévisible pour les erreurs synchrones et asynchrones. Le middleware d’erreur centralisé capture toutes les exceptions et fournit des réponses JSON standardisées, évitant les fuites d’informations sensibles.

```
// src/middleware/errorHandler.js
export const notFound = (req, res, next) => {
  res.status(404);
  throw new Error(`Route non trouvée : ${req.originalUrl}`);
};
export const errorHandler = (err, req, res, next) => {
  let statusCode = res.statusCode === 200 ? 500 : res.statusCode;
  let message = err.message;
  // Erreur de cast Mongoose (ID invalide)
  if (err.name === 'CastError' && err.kind === 'ObjectId') {
    statusCode = 400;
    message = 'Identifiant de ressource invalide';
  }
  // Erreur de duplication MongoDB
  if (err.code === 11000) {
    statusCode = 400;
    const field = Object.keys(err.keyValue)[0];
    message = `La valeur du champ '${field}' existe déjà`;
  }
  // Erreur de validation Mongoose
  if (err.name === 'ValidationError') {
    statusCode = 400;
    message = Object.values(err.errors).map(e => e.message).join(', ');
  }
  // Erreur JWT
  if (err.name === 'JsonWebTokenError') {
    statusCode = 401;
    message = 'Token invalide';
  }
  if (err.name === 'TokenExpiredError') {
    statusCode = 401;
    message = 'Token expiré';
  }
  res.status(statusCode).json({
    success: false,
    message,
    ...(process.env.NODE_ENV === 'development' && { stack: err.stack })
  });
};
```
Ajoutez les middleware d’erreur en dernière position dans `src/app.js` :

```
// En bas de src/app.js, APRÈS les routes
import { notFound, errorHandler } from './middleware/errorHandler.js';
app.use(notFound);
app.use(errorHandler);
```
En mode développement, la stack trace complète est incluse dans la réponse pour faciliter le débogage. En production, seul le message d’erreur est retourné. Le middleware gère automatiquement cinq types d’erreurs courants : CastError MongoDB, duplication de clé unique, validation Mongoose, et les deux erreurs JWT.

Exemple de réponse d’erreur :

```
# Requête avec un ID invalide
curl http://localhost:3000/api/products/abc123
# Réponse :
{
  "success": false,
  "message": "Identifiant de ressource invalide"
}
```
## Étape 9 : Ajouter la Validation des Données

La validation côté serveur est indispensable pour toute API REST en production. Nous utilisons express-validator pour définir des règles de validation déclaratives sur chaque route. Installez d’abord le paquet :

