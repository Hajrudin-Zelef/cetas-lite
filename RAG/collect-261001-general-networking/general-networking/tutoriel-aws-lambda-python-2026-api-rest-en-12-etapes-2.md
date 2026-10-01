---
id: collect-261001-general-networking/general-networking/tutoriel-aws-lambda-python-2026-api-rest-en-12-etapes-2
title: "Installer AWS CLI v2 sur Linux/macOS"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Lambda"]
dates: ["2010-09-09", "2016-10-31"]
keywords: ["aws", "graviton"]
source: docs/RAG/collect-261001-general-networking/tutoriel-aws-lambda-python-2026-api-rest-en-12-etapes.md
source_anchor: ""
source_lines: [126, 209]
sha256: ccbdbdb88be56c52183e78b61fa103146a1e73da433e2eac27f854af6a7fe831
---

# Installer AWS CLI v2 sur Linux/macOS

SAM CLI nécessite Docker pour émuler l’environnement Lambda. Si Docker n’est pas démarré, vous obtiendrez l’erreur **“Error: Running AWS SAM projects locally requires Docker”**. Lancez Docker Desktop avant d’exécuter **sam local invoke**. Sur Linux, assurez-vous que votre utilisateur fait partie du groupe docker avec **sudo usermod -aG docker $USER**.

## Étape 3 : Configurer le Template SAM avec API Gateway

Le fichier **template.yaml** est le cœur de votre infrastructure serverless. Il définit vos fonctions Lambda, les événements déclencheurs (API Gateway, S3, DynamoDB Streams), les permissions IAM et les ressources associées. Modifiez-le pour configurer une API REST complète avec les endpoints CRUD.

```
AWSTemplateFormatVersion: '2010-09-09'
Transform: AWS::Serverless-2016-10-31
Description: API REST CRUD Serverless avec Lambda et DynamoDB
Globals:
  Function:
    Timeout: 30
    MemorySize: 256
    Runtime: python3.13
    Architectures:
      - arm64
    Environment:
      Variables:
        TABLE_NAME: !Ref ProductsTable
        REGION: eu-west-3
Resources:
  # Fonction Lambda pour les opérations CRUD
  CrudFunction:
    Type: AWS::Serverless::Function
    Properties:
      CodeUri: crud_api/
      Handler: app.lambda_handler
      Policies:
        - DynamoDBCrudPolicy:
            TableName: !Ref ProductsTable
      Events:
        GetProducts:
          Type: Api
          Properties:
            Path: /products
            Method: get
        GetProduct:
          Type: Api
          Properties:
            Path: /products/{id}
            Method: get
        CreateProduct:
          Type: Api
          Properties:
            Path: /products
            Method: post
        UpdateProduct:
          Type: Api
          Properties:
            Path: /products/{id}
            Method: put
        DeleteProduct:
          Type: Api
          Properties:
            Path: /products/{id}
            Method: delete
  # Table DynamoDB
  ProductsTable:
    Type: AWS::DynamoDB::Table
    Properties:
      TableName: Products
      AttributeDefinitions:
        - AttributeName: id
          AttributeType: S
      KeySchema:
        - AttributeName: id
          KeyType: HASH
      BillingMode: PAY_PER_REQUEST
Outputs:
  ApiUrl:
    Description: URL de l'API Gateway
    Value: !Sub "https://${ServerlessRestApi}.execute-api.${AWS::Region}.amazonaws.com/Prod/"
```
Notez l’architecture **arm64** dans la section Globals. Depuis 2025, les fonctions Lambda exécutées sur processeurs AWS Graviton offrent des cold starts 13 à 24 % plus rapides et un coût réduit de 20 % par rapport à l’architecture x86_64. C’est un choix recommandé par AWS pour toutes les nouvelles fonctions Python. Le mode de facturation **PAY_PER_REQUEST** pour DynamoDB est idéal pour les charges de travail variables, car vous ne payez que les lectures et écritures réelles.

### Piège n°3 : Choisir x86_64 par défaut

Beaucoup de tutoriels utilisent encore l’architecture x86_64 par habitude. En 2026, AWS Graviton (arm64) offre de meilleures performances pour les fonctions Python et Node.js. Vérifiez simplement que vos dépendances natives (comme numpy ou pandas) sont compilées pour arm64. La commande **pip install –platform manylinux2014_aarch64** permet de télécharger les bonnes versions.

## Étape 4 : Implémenter les Opérations CRUD Complètes

Créez le répertoire **crud_api/** et implémentez le handler principal avec toutes les opérations CRUD. Cette fonction unique gère les 5 méthodes HTTP grâce au routage interne basé sur le champ **httpMethod** de l’événement API Gateway. Cette approche “fat Lambda” est plus simple pour les petits projets, mais vous pouvez aussi créer une fonction par route pour les applications plus grandes.

