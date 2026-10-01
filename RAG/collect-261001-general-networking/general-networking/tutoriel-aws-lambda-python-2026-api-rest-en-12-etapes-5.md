---
id: collect-261001-general-networking/general-networking/tutoriel-aws-lambda-python-2026-api-rest-en-12-etapes-5
title: "Installer AWS CLI v2 sur Linux/macOS"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Lambda"]
dates: []
keywords: ["aws", "memory"]
source: docs/RAG/collect-261001-general-networking/tutoriel-aws-lambda-python-2026-api-rest-en-12-etapes.md
source_anchor: ""
source_lines: [463, 573]
sha256: 2804391c66f39331089d386415bda8d4fb4c926b15695a6e29e08a22a94bc03f
---

# Installer AWS CLI v2 sur Linux/macOS

```
# Créer une Layer pour les dépendances communes
mkdir -p layers/common/python
pip install aws-lambda-powertools boto3 requests \
  -t layers/common/python \
  --platform manylinux2014_aarch64 \
  --only-binary=:all:
# Ajouter la Layer dans template.yaml
Resources:
  CommonLayer:
    Type: AWS::Serverless::LayerVersion
    Properties:
      LayerName: common-dependencies
      Description: Dépendances partagées (powertools, boto3)
      ContentUri: layers/common/
      CompatibleRuntimes:
        - python3.13
      CompatibleArchitectures:
        - arm64
    Metadata:
      BuildMethod: python3.13
  CrudFunction:
    Type: AWS::Serverless::Function
    Properties:
      Layers:
        - !Ref CommonLayer
      # ... reste de la configuration
```
L’avantage principal des Layers est la séparation des préoccupations : vous pouvez mettre à jour vos dépendances indépendamment de votre code métier. Cela accélère les déploiements de 40 à 60 %, car seul le code modifié est téléchargé. De plus, les Layers bénéficient d’un cache au niveau du conteneur Lambda, ce qui signifie que les dépendances ne sont chargées en mémoire qu’une seule fois, même si plusieurs fonctions les partagent.

### Piège n°5 : Inclure boto3 dans le package de déploiement

AWS Lambda inclut déjà **boto3** dans l’environnement d’exécution. Ajouter boto3 dans votre requirements.txt augmente inutilement la taille du package de 30 Mo. Cependant, si vous avez besoin d’une version spécifique plus récente que celle fournie par AWS, incluez-la explicitement dans une Layer. Vérifiez la version intégrée avec **python -c “import boto3; print(boto3.__version__)”** dans votre handler.

## Étape 10 : Implémenter les Lambda SnapStart pour les Workflows Complexes

Annoncées au re:Invent 2025, les **Lambda SnapStart** représentent l’évolution la plus importante d’AWS Lambda depuis son lancement. Elles permettent de créer des workflows multi-étapes qui s’exécutent sur des périodes allant de quelques secondes à un an, avec gestion automatique de l’état, des points de contrôle et de la reprise en cas d’échec. Contrairement au timeout standard de 15 minutes, les Durable Functions orchestrent plusieurs invocations Lambda de manière transparente.

```
# Exemple de Durable Function pour un pipeline de traitement
import json
from aws_lambda_powertools import Logger
logger = Logger()
def lambda_handler(event, context):
    """Workflow durable : traitement de commande en 4 étapes."""
    # Étape 1 : Validation de la commande
    order = validate_order(event['order_data'])
    logger.info("Commande validée", extra={"order_id": order['id']})
    # Étape 2 : Vérification du stock (appel externe)
    stock_status = check_inventory(order['items'])
    if not stock_status['available']:
        return {"status": "failed", "reason": "Stock insuffisant"}
    # Étape 3 : Traitement du paiement
    payment = process_payment(order['payment_info'])
    # Étape 4 : Confirmation et notification
    confirmation = send_confirmation(order, payment)
    return {
        "status": "completed",
        "order_id": order['id'],
        "confirmation_number": confirmation['number']
    }
# Configuration SAM pour Durable Functions
# template.yaml
#  OrderWorkflow:
#    Type: AWS::Serverless::Function
#    Properties:
#      FunctionName: order-workflow
#      Handler: workflow.lambda_handler
#      Runtime: python3.13
#      Timeout: 900
#      DurableFunction:
#        Enabled: true
#        MaxDuration: 86400  # 24 heures max
```
Les Durable Functions sont particulièrement adaptées aux workflows d’intelligence artificielle, où une chaîne de traitement peut inclure l’appel à plusieurs modèles, la transformation de données et l’agrégation de résultats. Depuis leur lancement, AWS rapporte que les Durable Functions sont utilisées dans 23 % des nouvelles applications serverless d’entreprise, principalement pour les pipelines de données et les workflows d’approbation. La gestion automatique de l’état élimine le besoin de bases de données intermédiaires ou de files d’attente complexes comme SQS + Step Functions.

## Étape 11 : Surveiller et Déboguer avec CloudWatch et X-Ray

Le monitoring est crucial pour toute application serverless en production. AWS Lambda s’intègre nativement avec CloudWatch pour les métriques, les logs et les alarmes, ainsi qu’avec X-Ray pour le tracing distribué. Configurez des alarmes CloudWatch pour être notifié en cas de taux d’erreur élevé, de durée d’exécution anormale ou de throttling (limitation du débit).

```
# Consulter les logs en temps réel
sam logs --name CrudFunction --stack-name lambda-api-crud --tail
# Sortie typique :
# 2026-04-08T14:30:01.234 START RequestId: abc-123
# 2026-04-08T14:30:01.256 {"level":"INFO","service":"products-api","message":"Listing all products","correlation_id":"abc-123"}
# 2026-04-08T14:30:01.312 END RequestId: abc-123
# 2026-04-08T14:30:01.312 REPORT Duration: 78.45 ms Billed Duration: 79 ms Memory: 256 MB Max Memory Used: 89 MB
# Créer une alarme CloudWatch pour le taux d'erreur
aws cloudwatch put-metric-alarm \
  --alarm-name "Lambda-CrudFunction-Errors" \
  --metric-name Errors \
  --namespace AWS/Lambda \
  --statistic Sum \
  --period 300 \
  --threshold 5 \
  --comparison-operator GreaterThanThreshold \
  --dimensions Name=FunctionName,Value=lambda-api-crud-CrudFunction \
  --evaluation-periods 1 \
  --alarm-actions arn:aws:sns:eu-west-3:123456789012:alerts
# Activer X-Ray dans template.yaml
Globals:
  Function:
    Tracing: Active
```
Les métriques clés à surveiller pour une fonction AWS Lambda sont la **durée d’exécution** (Duration), les **invocations** (Invocations), les **erreurs** (Errors), le **throttling** (Throttles) et les **invocations concurrentes** (ConcurrentExecutions). En 2026, AWS a ajouté la métrique **InitDuration** séparée qui mesure spécifiquement le temps de la phase d’initialisation, facilitant l’optimisation des cold starts. La limite de concurrence par défaut est de 1 000 invocations simultanées par région, extensible sur demande.

## Étape 12 : Déployer en Production avec CI/CD et Variables d’Environnement

Pour un déploiement de production robuste, intégrez votre pipeline Lambda dans un workflow CI/CD. AWS SAM s’intègre naturellement avec GitHub Actions, GitLab CI et AWS CodePipeline. Configurez des variables d’environnement séparées pour chaque stage (dev, staging, prod) et utilisez AWS Systems Manager Parameter Store ou Secrets Manager pour les données sensibles comme les clés API.

