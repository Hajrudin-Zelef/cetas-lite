---
id: collect-261001-general-networking/general-networking/tutoriel-mongodb-atlas-2026-app-node-js-en-13-etapes-4
title: "Stockez le mot de passe dans un fichier .env hors du repo Git"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Google"]
dates: []
keywords: ["acquisition", "aws", "benchmarks", "cost", "embedding", "embeddings", "valuation"]
source: docs/RAG/collect-261001-general-networking/tutoriel-mongodb-atlas-2026-app-node-js-en-13-etapes.md
source_anchor: ""
source_lines: [353, 493]
sha256: 6b92c3c9117f9d0f9d0c7294a64c35a672b53cfb7c463cc47205459f31af9564
---

# Stockez le mot de passe dans un fichier .env hors du repo Git

```
// Définition d'index Atlas Search via la console
{
  "mappings": {
    "dynamic": false,
    "fields": {
      "name": { "type": "string", "analyzer": "lucene.french" },
      "tags": { "type": "string" },
      "price_eur": { "type": "number" }
    }
  }
}
```
```
// Recherche fuzzy sur le nom + filtre prix
db.products.aggregate([
  { $search: {
      index: 'products_search',
      compound: {
        must: [{
          text: {
            query: 'clavier mécanique',
            path: 'name',
            fuzzy: { maxEdits: 1 }
          }
        }],
        filter: [{
          range: { path: 'price_eur', gte: 50, lte: 300 }
        }]
      }
  }},
  { $limit: 20 },
  { $project: { name: 1, price_eur: 1, score: { $meta: 'searchScore' } } }
]);
```
L’analyseur `lucene.french` gère le stemming et l’élision propres au français (« l’écran » ↔ « écran »). Les *lexical prefilters* introduits en février 2026 permettent désormais de filtrer les documents par index B-tree avant l’évaluation Lucene, ce qui réduit considérablement la latence sur des corpus de plusieurs millions de documents – voir la documentation Atlas Search.

## Étape 12 – Vector Search avec Voyage AI

Atlas Vector Search transforme votre cluster en vector database pour les applications RAG, recherche sémantique ou recommandation. Depuis l’acquisition de Voyage AI par MongoDB début 2025, Atlas peut *générer côté serveur* les embeddings via les modèles `voyage-3.5` (1 024 dimensions) ou `voyage-3.5-lite` (512 dimensions), supprimant un appel d’API externe à chaque ingestion. Cette intégration native est l’un des arguments commerciaux majeurs d’Atlas en 2026.

```
// Créer un index vectoriel via la console Atlas
{
  "fields": [
    {
      "type": "vector",
      "path": "description_embedding",
      "numDimensions": 1024,
      "similarity": "cosine"
    },
    {
      "type": "filter",
      "path": "tags"
    }
  ]
}
```
```
# scripts/embed_and_search.py
from pymongo import MongoClient
from voyageai import Client as VoyageClient
import os
vo = VoyageClient(api_key=os.environ['VOYAGE_API_KEY'])
client = MongoClient(os.environ['MONGO_URI'])
col = client.catalog.products
# 1. Générer un embedding pour la requête
query = "clavier silencieux pour le télétravail"
q_emb = vo.embed([query], model='voyage-3.5').embeddings[0]
# 2. Recherche kNN avec filtre
results = col.aggregate([
    {
        '$vectorSearch': {
            'index': 'products_vec',
            'path': 'description_embedding',
            'queryVector': q_emb,
            'numCandidates': 200,
            'limit': 10,
            'filter': {'tags': {'$in': ['clavier', 'périphérique']}}
        }
    },
    {
        '$project': {
            'name': 1,
            'price_eur': 1,
            'score': {'$meta': 'vectorSearchScore'}
        }
    }
])
for r in results:
    print(f"{r['score']:.3f} - {r['name']} ({r['price_eur']} €)")
```
Pour une *recherche hybride* (combinaison BM25 + vecteur), utilisez l’opérateur `$rankFusion` sorti en avril 2025. Il harmonise les scores des deux pipelines via la méthode *Reciprocal Rank Fusion*. Sur des corpus de catalogue produit, l’hybride dépasse régulièrement le pur vectoriel de 8 à 15 points de précision@10, selon les benchmarks de la documentation Vector Search.

## Étape 13 – Mise en production et observabilité

Pour passer en production, migrez du M0 vers un tier dédié **M10 minimum** (≈58 $/mois) qui apporte : nœuds dédiés, sauvegardes Cloud Backup, point-in-time recovery, alertes avancées et possibilité de scaling vertical. Activez **Continuous Cloud Backup** avec rétention de 7 jours minimum, et configurez l’export des métriques vers votre stack d’observabilité via **Atlas Datadog Integration** ou Prometheus (push gateway).

```
# Provisionner un cluster M10 via Atlas CLI
atlas auth login
atlas clusters create catalog-prod \
  --provider AWS \
  --region EU_WEST_3 \
  --tier M10 \
  --diskSizeGB 20 \
  --backup \
  --mdbVersion 8.0
# Activer le tag d'environnement (clé pour Cost Explorer)
atlas clusters update catalog-prod \
  --tag environment=production \
  --tag team=backend
```
Côté sécurité, activez **Encryption at Rest with Customer Key Management (CMK)** via AWS KMS, Azure Key Vault ou Google Cloud KMS pour les charges réglementées (santé, finance). MongoDB Atlas a publié en janvier 2026 un nouveau type d’alerte qui détecte les pertes d’accès au KMS empêchant le démarrage d’un nœud – un point critique souvent oublié.

Configurez enfin une politique *Auto-Scaling* qui ajuste le tier en fonction de la charge CPU et IOPS. Sur M10 → M30 par exemple, l’auto-scaling déclenche un upscale en moins de 90 secondes sans interruption de service grâce au *rolling restart* du replica set.

## Pièges courants et antipatterns à éviter

Voici les sept écueils les plus fréquemment rencontrés en production sur MongoDB Atlas, observés en France depuis 2023 par les équipes DevOps que nous avons interrogées.

- **Ouvrir le cluster à 0.0.0.0/0 :** faille critique. Restreignez toujours par CIDR.
- **Stocker la chaîne de connexion en dur :** utilisez Vault, AWS Secrets Manager ou GitHub Actions Secrets.
- **Modéliser comme une base relationnelle :** embarquer plutôt que joindre, sauf cardinalité élevée.
- **Oublier les index composés :** un index single-field ne couvre pas un`find + sort` .
- **Lancer des aggregations sans `$match` initial :** filtrez le plus tôt possible dans le pipeline.
- **Activer le tier M0 en production :** 100 ops/s, pas de backup. Migrez en M10 dès la mise en ligne.
- **Négliger les retries :** activez`retryWrites=true` dans la chaîne de connexion (par défaut depuis le driver 4).

## Astuces avancées pour aller plus loin

Une fois le tutoriel maîtrisé, plusieurs fonctionnalités méritent l’investissement temps. La **Queryable Encryption**, GA depuis MongoDB 7.0, permet de chiffrer côté client des champs sensibles (numéros de carte, IBAN) tout en autorisant des requêtes égalité et range sur la donnée chiffrée. C’est une pierre angulaire des architectures conformes au RGPD pour les données de santé.

Les **Time Series Collections** (GA depuis MongoDB 5.0) compressent jusqu’à 90 % les données chronologiques (métriques IoT, logs), avec un facteur de gain disque de 3 à 10× selon la cardinalité. Très utile pour stocker des données capteurs sans migrer vers InfluxDB ou TimescaleDB.

Les **Change Streams** exposent les modifications de la collection en temps réel via une API push. Couplé à un consommateur Kafka ou à des serverless functions Atlas (Atlas Functions), c’est la base des architectures Event-Driven sur MongoDB sans dépendre de Kafka Connect ou Debezium.

Enfin, **Atlas Data Federation** permet d’exécuter des requêtes MongoDB sur des données stockées en S3 ou Azure Blob (Parquet, JSON, CSV). Idéal pour archiver les vieux documents en S3 Glacier tout en conservant la possibilité d’y faire de l’analytics ad-hoc à coût réduit.

## Dépannage : 8 erreurs courantes et leur résolution

