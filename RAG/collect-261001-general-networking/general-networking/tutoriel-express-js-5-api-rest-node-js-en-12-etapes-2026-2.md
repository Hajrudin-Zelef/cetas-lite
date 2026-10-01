---
id: collect-261001-general-networking/general-networking/tutoriel-express-js-5-api-rest-node-js-en-12-etapes-2026-2
title: "v22.x.x (minimum v18)"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/tutoriel-express-js-5-api-rest-node-js-en-12-etapes-2026.md
source_anchor: ""
source_lines: [136, 345]
sha256: ecdf8239065c6e87e0217d3f065ad4fd71c6ea92b80cb133bb54445d2ebe78b8
---

# v22.x.x (minimum v18)

```
// src/config/database.js
import mongoose from 'mongoose';
export const connectDB = async () => {
  try {
    const conn = await mongoose.connect(process.env.MONGODB_URI, {
      maxPoolSize: 10,
      serverSelectionTimeoutMS: 5000,
      socketTimeoutMS: 45000
    });
    console.log(`MongoDB connecté : ${conn.connection.host}`);
    mongoose.connection.on('error', (err) => {
      console.error('Erreur MongoDB :', err);
    });
    mongoose.connection.on('disconnected', () => {
      console.warn('MongoDB déconnecté. Tentative de reconnexion...');
    });
  } catch (error) {
    console.error(`Erreur de connexion MongoDB : ${error.message}`);
    process.exit(1);
  }
};
```
**Piège n°3 :** Les options `useNewUrlParser` et `useUnifiedTopology` ne sont plus nécessaires avec les versions récentes de Mongoose. Les inclure génère un avertissement de dépréciation. Supprimez-les si vous migrez depuis un ancien projet.

Le pool de connexions (`maxPoolSize: 10`) est configuré pour gérer jusqu’à 10 connexions simultanées. En production, augmentez cette valeur selon votre charge, typiquement entre 20 et 50 pour une API à trafic moyen.

## Étape 4 : Créer les Modèles de Données Mongoose

Nous allons créer une API de gestion de produits avec authentification. Commençons par les modèles `User` et `Product`. La validation au niveau du schéma Mongoose constitue la première ligne de défense contre les données invalides.

```
// src/models/User.js
import mongoose from 'mongoose';
import bcrypt from 'bcryptjs';
const userSchema = new mongoose.Schema({
  name: {
    type: String,
    required: [true, 'Le nom est requis'],
    trim: true,
    minlength: [2, 'Le nom doit contenir au moins 2 caractères'],
    maxlength: [50, 'Le nom ne peut pas dépasser 50 caractères']
  },
  email: {
    type: String,
    required: [true, 'L\'email est requis'],
    unique: true,
    lowercase: true,
    match: [/^\S+@\S+\.\S+$/, 'Format d\'email invalide']
  },
  password: {
    type: String,
    required: [true, 'Le mot de passe est requis'],
    minlength: [8, 'Le mot de passe doit contenir au moins 8 caractères'],
    select: false
  },
  role: {
    type: String,
    enum: ['user', 'admin'],
    default: 'user'
  }
}, {
  timestamps: true,
  toJSON: { virtuals: true },
  toObject: { virtuals: true }
});
// Hash du mot de passe avant sauvegarde
userSchema.pre('save', async function(next) {
  if (!this.isModified('password')) return next();
  this.password = await bcrypt.hash(this.password, 12);
  next();
});
// Méthode de comparaison du mot de passe
userSchema.methods.comparePassword = async function(candidatePassword) {
  return bcrypt.compare(candidatePassword, this.password);
};
export default mongoose.model('User', userSchema);
```
```
// src/models/Product.js
import mongoose from 'mongoose';
const productSchema = new mongoose.Schema({
  name: {
    type: String,
    required: [true, 'Le nom du produit est requis'],
    trim: true,
    maxlength: [100, 'Le nom ne peut pas dépasser 100 caractères']
  },
  description: {
    type: String,
    required: [true, 'La description est requise'],
    maxlength: [2000, 'La description ne peut pas dépasser 2000 caractères']
  },
  price: {
    type: Number,
    required: [true, 'Le prix est requis'],
    min: [0, 'Le prix ne peut pas être négatif']
  },
  category: {
    type: String,
    required: [true, 'La catégorie est requise'],
    enum: ['electronique', 'vetements', 'alimentation', 'sport', 'maison']
  },
  stock: {
    type: Number,
    required: true,
    min: [0, 'Le stock ne peut pas être négatif'],
    default: 0
  },
  createdBy: {
    type: mongoose.Schema.Types.ObjectId,
    ref: 'User',
    required: true
  }
}, {
  timestamps: true
});
// Index pour les recherches fréquentes
productSchema.index({ category: 1, price: 1 });
productSchema.index({ name: 'text', description: 'text' });
export default mongoose.model('Product', productSchema);
```
Les index composites sur `category` et `price` accélèrent les requêtes de filtrage. L’index texte sur `name` et `description` permet les recherches full-text intégrées à MongoDB sans moteur de recherche externe.

## Étape 5 : Construire les Contrôleurs CRUD

Dans Express 5, les contrôleurs asynchrones n’ont plus besoin de blocs try-catch explicites grâce au support natif des Promises. Si une Promise est rejetée, Express la transmet automatiquement au middleware d’erreur. C’est l’un des changements les plus appréciés de cette version.

```
// src/controllers/productController.js
import Product from '../models/Product.js';
// GET /api/products
export const getProducts = async (req, res) => {
  const { page = 1, limit = 10, category, sort, search } = req.query;
  const filter = {};
  if (category) filter.category = category;
  if (search) filter.$text = { $search: search };
  const sortOptions = {};
  if (sort === 'price_asc') sortOptions.price = 1;
  else if (sort === 'price_desc') sortOptions.price = -1;
  else sortOptions.createdAt = -1;
  const skip = (parseInt(page) - 1) * parseInt(limit);
  const [products, total] = await Promise.all([
    Product.find(filter)
      .sort(sortOptions)
      .skip(skip)
      .limit(parseInt(limit))
      .populate('createdBy', 'name email'),
    Product.countDocuments(filter)
  ]);
  res.json({
    success: true,
    data: products,
    pagination: {
      page: parseInt(page),
      limit: parseInt(limit),
      total,
      pages: Math.ceil(total / parseInt(limit))
    }
  });
};
// GET /api/products/:id
export const getProduct = async (req, res) => {
  const product = await Product.findById(req.params.id)
    .populate('createdBy', 'name email');
  if (!product) {
    res.status(404);
    throw new Error('Produit non trouvé');
  }
  res.json({ success: true, data: product });
};
// POST /api/products
export const createProduct = async (req, res) => {
  const product = await Product.create({
    ...req.body,
    createdBy: req.user._id
  });
  res.status(201).json({ success: true, data: product });
};
// PUT /api/products/:id
export const updateProduct = async (req, res) => {
  const product = await Product.findByIdAndUpdate(
    req.params.id,
    req.body,
    { new: true, runValidators: true }
  );
  if (!product) {
    res.status(404);
    throw new Error('Produit non trouvé');
  }
  res.json({ success: true, data: product });
};
// DELETE /api/products/:id
export const deleteProduct = async (req, res) => {
  const product = await Product.findByIdAndDelete(req.params.id);
  if (!product) {
    res.status(404);
    throw new Error('Produit non trouvé');
  }
  res.status(204).send();
};
```
Remarquez l’absence totale de blocs try-catch. C’est la puissance d’Express 5 : toute erreur levée (synchrone ou asynchrone) est automatiquement interceptée par le middleware de gestion d’erreurs que nous configurerons à l’étape 8. La pagination utilise `Promise.all` pour exécuter la requête et le comptage en parallèle, réduisant le temps de réponse.

**Piège n°4 :** Avec Express 4, oublier un try-catch dans un handler asynchrone provoquait un crash silencieux du serveur. Express 5 résout ce problème, mais vous devez vous assurer d’utiliser la v5 et non la v4 pour en bénéficier.

## Étape 6 : Implémenter l’Authentification JWT

L’authentification par JSON Web Token (JWT) est le standard pour les API REST stateless. Nous allons créer un système complet avec inscription, connexion et middleware de protection des routes.

