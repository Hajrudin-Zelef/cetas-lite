---
id: collect-261001-general-networking/general-networking/tutoriel-mongodb-atlas-2026-app-node-js-en-13-etapes-3
title: "Stockez le mot de passe dans un fichier .env hors du repo Git"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2026-04-01"]
keywords: ["apache", "exploit"]
source: docs/RAG/collect-261001-general-networking/tutoriel-mongodb-atlas-2026-app-node-js-en-13-etapes.md
source_anchor: ""
source_lines: [218, 352]
sha256: 9f56588e45713d3492f477b8ba2a2bae2e800f56142bcb05b7459c44e692e78d
---

# Stockez le mot de passe dans un fichier .env hors du repo Git

```
python3 -m venv .venv
source .venv/bin/activate
pip install pymongo==4.17.0 python-dotenv==1.0.1 faker==30.3.0
```
```
# scripts/seed.py
import os
from pymongo import MongoClient, ASCENDING
from pymongo.server_api import ServerApi
from faker import Faker
from urllib.parse import quote_plus
from dotenv import load_dotenv
load_dotenv()
fake = Faker('fr_FR')
uri = (
    f"mongodb+srv://{os.environ['MONGO_USER']}:"
    f"{quote_plus(os.environ['MONGO_PASSWORD'])}@"
    f"{os.environ['MONGO_CLUSTER']}/?retryWrites=true&w=majority"
)
client = MongoClient(uri, server_api=ServerApi('1'))
db = client[os.environ['MONGO_DB']]
products = []
for i in range(5000):
    products.append({
        'sku': f'TI-{i:06d}',
        'name': fake.bs().capitalize(),
        'price_eur': round(fake.random.uniform(5, 999), 2),
        'stock': fake.random.randint(0, 500),
        'tags': fake.words(nb=3),
        'created_at': fake.date_time_this_year(),
    })
result = db.products.insert_many(products, ordered=False)
print(f'Inséré {len(result.inserted_ids)} documents')
client.close()
```
Sur un cluster M0, attendez-vous à environ 60 secondes pour 5 000 documents (le quota I/O étant partagé). Sur M10, le même volume passe sous les 4 secondes. Pour des jobs d’ingestion supérieurs au million de documents, basculez vers `BulkWrite` avec batches de 1 000 – au-delà, MongoDB fragmente automatiquement la requête côté serveur.

## Étape 8 – Conception du schéma et validation JSON

MongoDB n’impose aucun schéma par défaut, mais la **JSON Schema validation** côté serveur est la meilleure pratique en production depuis MongoDB 3.6 et reste fortement recommandée en 8.0. Elle garantit qu’une régression côté application n’introduit pas de documents corrompus dans la collection. Définissez le schéma au moment de la création de la collection.

```
// À exécuter une fois dans mongosh
db.createCollection('products', {
  validator: {
    $jsonSchema: {
      bsonType: 'object',
      required: ['sku', 'name', 'price_eur'],
      properties: {
        sku: {
          bsonType: 'string',
          pattern: '^TI-[0-9]{3,8}$',
          description: 'SKU au format TI-XXX'
        },
        name: { bsonType: 'string', minLength: 2, maxLength: 200 },
        price_eur: { bsonType: 'double', minimum: 0 },
        stock: { bsonType: 'int', minimum: 0 },
        tags: { bsonType: 'array', items: { bsonType: 'string' } },
        created_at: { bsonType: 'date' }
      }
    }
  },
  validationLevel: 'strict',
  validationAction: 'error'
});
```
**Antipattern à éviter :** ne modélisez pas vos données comme dans une base relationnelle. MongoDB est conçu pour le pattern *embedded documents*. Au lieu de stocker une table `orders` et une table `order_items` liées par clé étrangère, embarquez les `items` dans le document `order`. La règle empirique : si vous lisez toujours les sous-objets en même temps que le parent, embarquez. Si la cardinalité dépasse plusieurs centaines, référencez par `_id`.

## Étape 9 – Index, performance et profilage

Un index manquant est la première cause de lenteur sur MongoDB. La règle de base : tout champ utilisé dans `find`, `sort` ou comme jointure dans `$lookup` doit être indexé. Atlas inclut un **Performance Advisor** qui analyse les requêtes des dernières 24 heures et propose les index manquants – accessible dans l’onglet Performance Advisor du cluster.

```
// Créer un index composé sur SKU (unique) et prix descendant
db.products.createIndex({ sku: 1 }, { unique: true });
db.products.createIndex({ price_eur: -1, stock: 1 });
// Index TTL pour expirer les documents après 30 jours
db.cache.createIndex({ created_at: 1 }, { expireAfterSeconds: 2592000 });
// Profiler une requête
db.products.find({ price_eur: { $lt: 50 } })
  .sort({ price_eur: 1 })
  .explain('executionStats');
```
Le retour de `explain('executionStats')` doit afficher `"stage": "IXSCAN"` pour confirmer que l’index est exploité. Si vous voyez `"COLLSCAN"`, votre requête lit la collection en entier : ajoutez un index ciblé. Le ratio `nReturned / totalDocsExamined` doit s’approcher de 1 sur des collections en production.

| Type d’index | Cas d’usage | Coût en écriture | 
|---|---|---|
| Single field | Filtre simple sur un champ | Faible | 
| Compound | Filtre + tri sur plusieurs champs | Modéré | 
| Multikey | Tableaux (ex: tags) | Modéré | 
| Text | Recherche full-text basique | Élevé | 
| Geo (2dsphere) | Coordonnées GPS | Modéré | 
| TTL | Expiration auto de documents | Faible | 
| Wildcard | Schémas dynamiques | Élevé | 
| Hashed | Sharding | Modéré | 

**Piège fréquent :** les index *text* ne supportent qu’un seul index text par collection. Si vous avez besoin de recherche multi-champ pondérée, basculez vers Atlas Search (étape 11), qui s’appuie sur Lucene et offre une bien meilleure pertinence.

## Étape 10 – Pipelines d’agrégation

L’agrégation MongoDB est l’équivalent des CTE et des fonctions analytiques de SQL. Elle s’écrit comme une succession d’*étages* (`$match`, `$group`, `$lookup`, `$project`) appliqués séquentiellement sur le flux de documents. C’est l’outil principal pour générer des rapports analytics depuis MongoDB sans transit par un data warehouse externe.

```
// Top 5 catégories par chiffre d'affaires sur le mois courant
db.orders.aggregate([
  { $match: {
      status: 'paid',
      created_at: { $gte: new Date('2026-04-01') }
  }},
  { $unwind: '$items' },
  { $lookup: {
      from: 'products',
      localField: 'items.product_id',
      foreignField: '_id',
      as: 'product'
  }},
  { $unwind: '$product' },
  { $group: {
      _id: { $arrayElemAt: ['$product.tags', 0] },
      total_eur: { $sum: { $multiply: ['$items.qty', '$product.price_eur'] } },
      orders: { $sum: 1 }
  }},
  { $sort: { total_eur: -1 } },
  { $limit: 5 }
]);
```
Le pipeline ci-dessus joint `orders` à `products` via `$lookup`, calcule le CA par tag principal et retourne le Top 5. Sur M10, comptez 80–150 ms pour 100 000 commandes filtrées si les index `created_at` et `items.product_id` sont en place. **Limite mémoire :** chaque étage limite à 100 Mo de RAM par défaut. Au-delà, ajoutez `{ allowDiskUse: true }` en option globale du `aggregate`.

L’opérateur `$densify` introduit en MongoDB 5.3 reste indispensable pour combler les trous d’une série temporelle (ventes par jour avec des jours sans vente). Pensez aussi à `$setWindowFields` pour les calculs de moyenne mobile, équivalent de `OVER (PARTITION BY...)` en SQL.

## Étape 11 – Atlas Search avec Lucene

Atlas Search est intégré à chaque cluster (y compris M0). C’est un moteur Apache Lucene managé qui s’exécute sur les mêmes nœuds que vos données – pas besoin d’Elasticsearch externe. Vous créez l’index via la console ou l’API Atlas, puis vous l’interrogez via l’opérateur `$search` de l’agrégation.

