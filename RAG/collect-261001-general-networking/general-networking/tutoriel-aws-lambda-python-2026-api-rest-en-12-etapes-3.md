---
id: collect-261001-general-networking/general-networking/tutoriel-aws-lambda-python-2026-api-rest-en-12-etapes-3
title: "Installer AWS CLI v2 sur Linux/macOS"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Lambda"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-general-networking/tutoriel-aws-lambda-python-2026-api-rest-en-12-etapes.md
source_anchor: ""
source_lines: [210, 349]
sha256: 0a01cf7867d4533eaa6bbfaf19ec44e9e37de8ca5701eea6a096623897659f0b
---

# Installer AWS CLI v2 sur Linux/macOS

```
import json
import os
import uuid
import boto3
from datetime import datetime
from decimal import Decimal
# Initialisation en dehors du handler pour réutiliser la connexion
dynamodb = boto3.resource('dynamodb', region_name=os.environ.get('REGION', 'eu-west-3'))
table = dynamodb.Table(os.environ.get('TABLE_NAME', 'Products'))
class DecimalEncoder(json.JSONEncoder):
    """Encoder personnalisé pour les types Decimal de DynamoDB."""
    def default(self, obj):
        if isinstance(obj, Decimal):
            return float(obj)
        return super().default(obj)
def build_response(status_code, body):
    """Construire une réponse HTTP standardisée."""
    return {
        "statusCode": status_code,
        "headers": {
            "Content-Type": "application/json",
            "Access-Control-Allow-Origin": "*",
            "Access-Control-Allow-Methods": "GET,POST,PUT,DELETE,OPTIONS"
        },
        "body": json.dumps(body, cls=DecimalEncoder)
    }
def lambda_handler(event, context):
    """Router les requêtes vers les opérations CRUD."""
    http_method = event.get('httpMethod', '')
    path = event.get('path', '')
    path_params = event.get('pathParameters') or {}
    try:
        if http_method == 'GET' and 'id' in path_params:
            return get_product(path_params['id'])
        elif http_method == 'GET':
            return list_products()
        elif http_method == 'POST':
            body = json.loads(event.get('body', '{}'))
            return create_product(body)
        elif http_method == 'PUT' and 'id' in path_params:
            body = json.loads(event.get('body', '{}'))
            return update_product(path_params['id'], body)
        elif http_method == 'DELETE' and 'id' in path_params:
            return delete_product(path_params['id'])
        else:
            return build_response(405, {"error": "Méthode non autorisée"})
    except Exception as e:
        return build_response(500, {"error": str(e)})
def list_products():
    """Lister tous les produits."""
    result = table.scan()
    return build_response(200, {
        "products": result.get('Items', []),
        "count": result.get('Count', 0)
    })
def get_product(product_id):
    """Récupérer un produit par son ID."""
    result = table.get_item(Key={'id': product_id})
    item = result.get('Item')
    if not item:
        return build_response(404, {"error": "Produit non trouvé"})
    return build_response(200, item)
def create_product(data):
    """Créer un nouveau produit."""
    if not data.get('name') or not data.get('price'):
        return build_response(400, {"error": "Les champs 'name' et 'price' sont requis"})
    product = {
        'id': str(uuid.uuid4()),
        'name': data['name'],
        'price': Decimal(str(data['price'])),
        'description': data.get('description', ''),
        'category': data.get('category', 'general'),
        'created_at': datetime.utcnow().isoformat(),
        'updated_at': datetime.utcnow().isoformat()
    }
    table.put_item(Item=product)
    return build_response(201, product)
def update_product(product_id, data):
    """Mettre à jour un produit existant."""
    result = table.update_item(
        Key={'id': product_id},
        UpdateExpression="SET #n = :name, price = :price, description = :desc, updated_at = :updated",
        ExpressionAttributeNames={'#n': 'name'},
        ExpressionAttributeValues={
            ':name': data.get('name', ''),
            ':price': Decimal(str(data.get('price', 0))),
            ':desc': data.get('description', ''),
            ':updated': datetime.utcnow().isoformat()
        },
        ReturnValues="ALL_NEW"
    )
    return build_response(200, result.get('Attributes', {}))
def delete_product(product_id):
    """Supprimer un produit."""
    table.delete_item(Key={'id': product_id})
    return build_response(200, {"message": f"Produit {product_id} supprimé"})
```
Remarquez que le client DynamoDB est initialisé en dehors du handler. C’est une optimisation fondamentale : AWS Lambda réutilise le conteneur d’exécution entre les invocations successives (warm starts), ce qui signifie que la connexion DynamoDB est établie une seule fois puis réutilisée. Cette technique réduit la latence des invocations chaudes de 15 à 40 ms. Le **DecimalEncoder** est nécessaire car DynamoDB stocke les nombres sous forme de Decimal, incompatible avec le sérialiseur JSON standard de Python.

## Étape 5 : Déployer Votre API Serverless sur AWS

Le déploiement avec SAM CLI est un processus en deux étapes : le build (compilation et empaquetage) puis le deploy (envoi vers AWS). La première exécution de **sam deploy –guided** crée un fichier **samconfig.toml** qui mémorise vos paramètres pour les déploiements suivants.

```
# Build du projet (compilation des dépendances)
sam build
# Sortie attendue :
# Building codeuri: /home/user/lambda-api-crud/crud_api runtime: python3.13
# Running PythonPipBuilder:ResolveDependencies
# Running PythonPipBuilder:CopySource
# Build Succeeded
# Déployer en mode guidé (première fois)
sam deploy --guided
# Réponses recommandées :
# Stack Name: lambda-api-crud
# AWS Region: eu-west-3
# Confirm changes before deploy: y
# Allow SAM CLI IAM role creation: y
# Disable rollback: n
# Save arguments to configuration file: y
# Sortie de déploiement :
# CloudFormation outputs from deployed stack
# -------------------------------------------
# Key: ApiUrl
# Value: https://abc123def4.execute-api.eu-west-3.amazonaws.com/Prod/
# -------------------------------------------
# Déploiements suivants (utilise samconfig.toml)
sam deploy
```
Après le déploiement, SAM affiche l’URL de votre API Gateway. Conservez cette URL, elle constitue le point d’entrée de votre API REST. Le déploiement crée automatiquement un stack CloudFormation contenant la fonction Lambda, la table DynamoDB, l’API Gateway, le rôle IAM et les permissions associées. En cas d’erreur, CloudFormation effectue un rollback automatique vers l’état précédent.

### Piège n°4 : Déployer sans tester le build

Exécutez toujours **sam build** avant **sam deploy**. Un déploiement sans build préalable envoie le code source brut sans les dépendances compilées. Si votre **requirements.txt** contient des bibliothèques comme **boto3** ou **requests**, elles ne seront pas incluses dans le package de déploiement et votre fonction échouera avec une erreur **ModuleNotFoundError**.

## Étape 6 : Tester l’API REST avec curl et Valider les Réponses

Une fois l’API déployée, testez chaque endpoint CRUD avec curl. Remplacez l’URL par celle retournée par votre déploiement SAM. Ces tests vérifient que l’intégration entre API Gateway, Lambda et DynamoDB fonctionne correctement de bout en bout.

