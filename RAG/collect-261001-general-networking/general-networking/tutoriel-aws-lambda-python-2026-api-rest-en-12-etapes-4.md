---
id: collect-261001-general-networking/general-networking/tutoriel-aws-lambda-python-2026-api-rest-en-12-etapes-4
title: "Installer AWS CLI v2 sur Linux/macOS"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Lambda"]
dates: []
keywords: ["aws", "graviton"]
source: docs/RAG/collect-261001-general-networking/tutoriel-aws-lambda-python-2026-api-rest-en-12-etapes.md
source_anchor: ""
source_lines: [350, 462]
sha256: 6720e35cf4ab20800b5807a203fd2de6610f67dbfe562be0cfdd93ae8c7fae21
---

# Installer AWS CLI v2 sur Linux/macOS

```
# Variable pour l'URL de base
API_URL="https://abc123def4.execute-api.eu-west-3.amazonaws.com/Prod"
# 1. Créer un produit (POST)
curl -X POST "$API_URL/products" \
  -H "Content-Type: application/json" \
  -d '{"name":"Clavier Mécanique","price":89.99,"description":"Clavier RGB switches Cherry MX","category":"périphériques"}'
# Réponse :
# {"id":"f47ac10b-58cc-4372-a567-0e02b2c3d479","name":"Clavier Mécanique","price":89.99,"description":"Clavier RGB switches Cherry MX","category":"périphériques","created_at":"2026-04-08T14:23:01.234567","updated_at":"2026-04-08T14:23:01.234567"}
# 2. Lister tous les produits (GET)
curl "$API_URL/products"
# Réponse :
# {"products":[{"id":"f47ac10b-...","name":"Clavier Mécanique","price":89.99,...}],"count":1}
# 3. Récupérer un produit (GET by ID)
curl "$API_URL/products/f47ac10b-58cc-4372-a567-0e02b2c3d479"
# 4. Mettre à jour un produit (PUT)
curl -X PUT "$API_URL/products/f47ac10b-58cc-4372-a567-0e02b2c3d479" \
  -H "Content-Type: application/json" \
  -d '{"name":"Clavier Mécanique Pro","price":119.99,"description":"Clavier RGB switches Cherry MX Brown"}'
# 5. Supprimer un produit (DELETE)
curl -X DELETE "$API_URL/products/f47ac10b-58cc-4372-a567-0e02b2c3d479"
# Réponse : {"message":"Produit f47ac10b-58cc-4372-a567-0e02b2c3d479 supprimé"}
```
Si une requête retourne une erreur 502 (Bad Gateway), cela indique généralement un problème dans votre code Lambda. Consultez les logs CloudWatch avec **sam logs –name CrudFunction –stack-name lambda-api-crud –tail** pour identifier l’erreur. Les erreurs 403 indiquent un problème de permissions IAM, tandis que les erreurs 404 signifient que la route n’est pas configurée dans API Gateway.

## Étape 7 : Optimiser les Cold Starts et les Performances Lambda

Les cold starts restent le défi principal du serverless en 2026, malgré les améliorations significatives d’AWS. Un cold start survient lorsque Lambda doit initialiser un nouveau conteneur d’exécution pour votre fonction. Avec Python 3.13 sur arm64, le cold start moyen est de 250 à 450 ms, contre 400 à 800 ms sur x86_64 avec Python 3.11. Voici les techniques d’optimisation les plus efficaces.

| Technique d’optimisation | Réduction cold start | Complexité | Coût additionnel | 
|---|---|---|---|
| Architecture arm64 (Graviton) | 13-24 % | Faible | -20 % (économie) | 
| SnapStart pour Python | Jusqu’à 4,3x | Moyenne | Aucun | 
| Provisioned Concurrency | ~100 % (élimine) | Faible | 0,015 $/Go/h | 
| Réduire la taille du package | 10-30 % | Moyenne | Aucun | 
| Lambda Layers | Variable | Moyenne | Aucun | 
| Initialisation lazy | 5-15 % | Faible | Aucun | 

L’expansion de **SnapStart à Python**, annoncée fin 2024 et stabilisée en 2025, est l’amélioration la plus significative. SnapStart prend un instantané de votre conteneur après l’initialisation et le restaure lors des invocations suivantes, réduisant les cold starts d’un facteur 4,3x. Pour l’activer, ajoutez la configuration suivante dans votre template SAM :

```
# Dans template.yaml, sous Properties de votre fonction
CrudFunction:
  Type: AWS::Serverless::Function
  Properties:
    SnapStart:
      ApplyOn: PublishedVersions
    AutoPublishAlias: live
    # ... reste de la configuration
```
Une autre optimisation essentielle consiste à réduire la taille de votre package de déploiement. Chaque mégaoctet supplémentaire ajoute environ 10 ms au cold start. Excluez les fichiers inutiles (tests, documentation, fichiers __pycache__) et utilisez des Lambda Layers pour les dépendances partagées entre plusieurs fonctions. Depuis 2025, AWS recommande également d’utiliser le SDK AWS v3 avec les paramètres **keepAlive: true**, **maxAttempts: 2** et un **timeout de 3000 ms** pour le client DynamoDB afin de réduire la latence réseau.

## Étape 8 : Ajouter la Validation, la Gestion d’Erreurs et les Logs Structurés

Une API de production nécessite une validation robuste des entrées, une gestion d’erreurs cohérente et des logs structurés pour le débogage. AWS Lambda intègre nativement CloudWatch Logs, mais les logs non structurés sont difficiles à rechercher et analyser. Implémentez un système de logging structuré avec le module **aws-lambda-powertools**, la bibliothèque officielle d’AWS pour les bonnes pratiques Lambda.

```
# requirements.txt
boto3>=1.35.0
aws-lambda-powertools>=3.5.0
# app.py - Version améliorée avec Powertools
import json
import os
import uuid
from decimal import Decimal
from datetime import datetime
from aws_lambda_powertools import Logger, Tracer, Metrics
from aws_lambda_powertools.event_handler import APIGatewayRestResolver
from aws_lambda_powertools.logging import correlation_paths
from aws_lambda_powertools.utilities.validation import validate
import boto3
logger = Logger(service="products-api")
tracer = Tracer(service="products-api")
metrics = Metrics(namespace="ProductsAPI")
app = APIGatewayRestResolver()
dynamodb = boto3.resource('dynamodb')
table = dynamodb.Table(os.environ.get('TABLE_NAME', 'Products'))
@app.get("/products")
@tracer.capture_method
def list_products():
    logger.info("Listing all products")
    result = table.scan()
    metrics.add_metric(name="ProductsListed", unit="Count", value=1)
    return {"products": result.get('Items', []), "count": result.get('Count', 0)}
@app.post("/products")
@tracer.capture_method
def create_product():
    data = app.current_event.json_body
    if not data.get('name') or not data.get('price'):
        raise ValueError("Les champs 'name' et 'price' sont requis")
    product = {
        'id': str(uuid.uuid4()),
        'name': data['name'],
        'price': Decimal(str(data['price'])),
        'description': data.get('description', ''),
        'created_at': datetime.utcnow().isoformat()
    }
    table.put_item(Item=product)
    logger.info("Product created", extra={"product_id": product['id']})
    metrics.add_metric(name="ProductCreated", unit="Count", value=1)
    return product, 201
@logger.inject_lambda_context(correlation_id_path=correlation_paths.API_GATEWAY_REST)
@tracer.capture_lambda_handler
@metrics.log_metrics(capture_cold_start_metric=True)
def lambda_handler(event, context):
    return app.resolve(event, context)
```
La bibliothèque **aws-lambda-powertools** (version 3.5 en avril 2026) fournit trois utilitaires essentiels. Le **Logger** produit des logs JSON structurés avec corrélation automatique des requêtes. Le **Tracer** intègre AWS X-Ray pour le tracing distribué. Les **Metrics** envoient des métriques personnalisées vers CloudWatch, incluant automatiquement une métrique de cold start. Cette approche permet de diagnostiquer les problèmes de performance en quelques secondes plutôt qu’en heures.

## Étape 9 : Configurer les Lambda Layers pour les Dépendances Partagées

Les Lambda Layers permettent de séparer les dépendances du code métier, réduisant la taille de vos packages de déploiement et accélérant les cycles de développement. Une Layer peut contenir des bibliothèques Python, des binaires ou des fichiers de configuration partagés entre plusieurs fonctions. Depuis 2026, chaque fonction Lambda peut référencer jusqu’à 5 Layers, avec une taille totale décompressée maximale de 250 Mo.

