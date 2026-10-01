---
id: collect-261001-general-networking/general-networking/tutoriel-mongodb-atlas-2026-app-node-js-en-13-etapes-2
title: "Stockez le mot de passe dans un fichier .env hors du repo Git"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["open source"]
source: docs/RAG/collect-261001-general-networking/tutoriel-mongodb-atlas-2026-app-node-js-en-13-etapes.md
source_anchor: ""
source_lines: [60, 217]
sha256: 20785c2007f9b100abd81631f478d69330fbb2b1367439758557454b83057df2
---

# Stockez le mot de passe dans un fichier .env hors du repo Git

Passez ensuite dans **Security > Database Access** et créez un utilisateur applicatif. Cliquez « Add New Database User », choisissez l’authentification *SCRAM* (par défaut), nommez l’utilisateur `catalog_app`, et générez un mot de passe fort via le bouton « Autogenerate » – copiez-le immédiatement, il ne sera plus affiché. Affectez le rôle *readWrite* sur la base `catalog` uniquement. Évitez systématiquement le rôle *atlasAdmin* pour des comptes applicatifs en production.

```
# Stockez le mot de passe dans un fichier .env hors du repo Git
echo "MONGO_USER=catalog_app" >> .env
echo "MONGO_PASSWORD=VotreMotDePasseFort_2026" >> .env
echo "MONGO_CLUSTER=catalog-cluster.abc12.mongodb.net" >> .env
echo "MONGO_DB=catalog" >> .env
# Ajoutez .env à .gitignore avant le moindre commit
echo ".env" >> .gitignore
```
**Piège fréquent :** les caractères spéciaux du mot de passe (`@`, `:`, `/`, `?`) doivent être *URL-encodés* dans la chaîne de connexion. Préférez un mot de passe long alphanumérique pour éviter ce casse-tête. Les rotations de mot de passe via Atlas CLI `atlas dbusers update` propagent le changement en moins de 30 secondes sur tous les nœuds réplica.

## Étape 4 – Se connecter avec MongoDB Compass

Téléchargez **MongoDB Compass 1.46** depuis la page Tools de Atlas (Linux .deb, macOS .dmg, Windows .exe). Compass reste l’outil GUI de référence pour explorer collections, profiler les requêtes lentes et tester les pipelines d’agrégation visuellement. Sur la console Atlas, cliquez sur « Connect » à côté de votre cluster, puis « Compass » – Atlas génère une URI au format SRV.

```
# Format de la chaîne de connexion (à coller dans Compass)
mongodb+srv://catalog_app:<password>@catalog-cluster.abc12.mongodb.net/?retryWrites=true&w=majority&appName=catalog-app
# Avec mot de passe URL-encodé en CLI
mongosh "mongodb+srv://catalog-cluster.abc12.mongodb.net/" \
  --apiVersion 1 \
  --username catalog_app
```
Une fois connecté, créez la base `catalog` et la collection `products` via le bouton « Create Database ». Atlas n’instancie réellement la base que lors de la première écriture, ce qui est conforme au comportement standard de MongoDB. Insérez un document de test pour valider la chaîne :

```
// Dans Compass, onglet "Insert Document"
{
  "sku": "TI-2026-001",
  "name": "Clavier mécanique Keychron Q1",
  "price_eur": 199.00,
  "stock": 42,
  "tags": ["periphérique", "mécanique", "QMK"],
  "created_at": new Date()
}
```
Si l’insertion échoue avec `MongoServerSelectionError`, retournez vérifier votre IP dans Network Access – c’est la cause numéro un des échecs de connexion en environnement résidentiel français, où les opérateurs (Orange, Free, SFR, Bouygues) renouvellent l’IP publique toutes les 24 à 72 heures.

## Étape 5 – Initialiser le projet Node.js et installer le driver

Créez le squelette de l’application. Le tutoriel utilise Express 5 (LTS depuis octobre 2024), Node.js 22 LTS et le driver natif `mongodb 7.2`. On évite Mongoose pour ce premier projet : rester sur le driver natif vous expose au comportement réel de l’API et facilite l’apprentissage des opérateurs MongoDB.

```
mkdir catalog-api && cd catalog-api
npm init -y
npm install express@5 mongodb@7 dotenv@16
npm install --save-dev nodemon@3 typescript@5
# Structure recommandée
mkdir -p src/routes src/db tests
touch src/index.js src/db/client.js src/routes/products.js
```
Configurez le client MongoDB de manière singleton pour éviter d’ouvrir un pool à chaque requête HTTP. Le driver gère un pool interne par défaut de **100 connexions**, ce qui couvre largement les besoins d’un service web standard. Plus de détails dans la documentation officielle du driver Node.js.

```
// src/db/client.js
import { MongoClient, ServerApiVersion } from 'mongodb';
import 'dotenv/config';
const uri = `mongodb+srv://${process.env.MONGO_USER}:${encodeURIComponent(process.env.MONGO_PASSWORD)}@${process.env.MONGO_CLUSTER}/?retryWrites=true&w=majority&appName=catalog-api`;
const client = new MongoClient(uri, {
  serverApi: { version: ServerApiVersion.v1, strict: true, deprecationErrors: true },
  maxPoolSize: 50,
  minPoolSize: 5,
  serverSelectionTimeoutMS: 5000,
});
let db;
export async function getDb() {
  if (!db) {
    await client.connect();
    db = client.db(process.env.MONGO_DB);
    console.log('[mongo] Connecté au cluster Atlas');
  }
  return db;
}
export async function closeDb() {
  await client.close();
}
```
**Astuce :** activez l’API Stable v1 (`strict: true`). Elle gèle le contrat de l’API serveur indépendamment des évolutions du moteur, ce qui réduit les risques de régression lors d’une mise à niveau majeure de MongoDB. Le coût est négligeable et l’effet bénéfique sur la stabilité d’une production de longue durée.

## Étape 6 – CRUD complet avec Express

Implémentez les quatre verbes REST sur la collection `products`. On utilise les méthodes natives `insertOne`, `find`, `updateOne` et `deleteOne`. Notez l’usage systématique de `ObjectId` pour caster l’identifiant venant de l’URL et la validation minimale par *JSON Schema* que vous appliquerez à l’étape 8.

```
// src/routes/products.js
import { Router } from 'express';
import { ObjectId } from 'mongodb';
import { getDb } from '../db/client.js';
const router = Router();
router.post('/', async (req, res) => {
  const db = await getDb();
  const result = await db.collection('products').insertOne({
    ...req.body,
    created_at: new Date(),
  });
  res.status(201).json({ id: result.insertedId });
});
router.get('/', async (req, res) => {
  const db = await getDb();
  const items = await db.collection('products')
    .find({})
    .sort({ created_at: -1 })
    .limit(50)
    .toArray();
  res.json(items);
});
router.get('/:id', async (req, res) => {
  const db = await getDb();
  const item = await db.collection('products')
    .findOne({ _id: new ObjectId(req.params.id) });
  if (!item) return res.status(404).json({ error: 'introuvable' });
  res.json(item);
});
router.patch('/:id', async (req, res) => {
  const db = await getDb();
  const result = await db.collection('products').updateOne(
    { _id: new ObjectId(req.params.id) },
    { $set: { ...req.body, updated_at: new Date() } },
  );
  res.json({ modified: result.modifiedCount });
});
router.delete('/:id', async (req, res) => {
  const db = await getDb();
  const result = await db.collection('products')
    .deleteOne({ _id: new ObjectId(req.params.id) });
  res.json({ deleted: result.deletedCount });
});
export default router;
```
Le point d’entrée Express est court et lit le fichier `.env` via dotenv. Lancez l’application avec `nodemon src/index.js` et testez avec curl ou Bruno (alternative open source à Postman très populaire en France depuis 2024).

```
// src/index.js
import express from 'express';
import productsRouter from './routes/products.js';
import { getDb, closeDb } from './db/client.js';
const app = express();
app.use(express.json());
app.use('/products', productsRouter);
const port = process.env.PORT || 3000;
app.listen(port, async () => {
  await getDb();
  console.log(`[api] http://localhost:${port}`);
});
process.on('SIGTERM', async () => {
  await closeDb();
  process.exit(0);
});
```
Sortie attendue après `curl -X POST localhost:3000/products -H 'content-type: application/json' -d '{"sku":"TI-001","name":"Souris Logitech MX","price_eur":99}'` :

`{ "id": "6620b1f3a7a14e2c9a5f1234" }`
## Étape 7 – Connexion avec PyMongo et opérations bulk

Pour les jobs d’ingestion massive ou les pipelines de data engineering, Python avec PyMongo reste plus ergonomique que Node.js. Installez **PyMongo 4.17.0** dans un environnement virtuel et utilisez `insert_many` avec `ordered=False` pour paralléliser les insertions. La page PyPI de PyMongo liste les versions disponibles et leurs notes de release.

