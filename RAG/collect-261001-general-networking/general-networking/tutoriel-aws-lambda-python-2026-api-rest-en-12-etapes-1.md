---
id: collect-261001-general-networking/general-networking/tutoriel-aws-lambda-python-2026-api-rest-en-12-etapes-1
title: "Installer AWS CLI v2 sur Linux/macOS"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Lambda"]
dates: ["2012-10-17", "2026-04-10"]
keywords: ["aws", "memory"]
source: docs/RAG/collect-261001-general-networking/tutoriel-aws-lambda-python-2026-api-rest-en-12-etapes.md
source_anchor: ""
source_lines: [1, 125]
sha256: 3d74a14702aadbfee88dcb8d69d3249e58a124817fc41d029bb55d4d54b13750
---

# Installer AWS CLI v2 sur Linux/macOS

AWS Lambda reste en 2026 la plateforme serverless la plus adoptée au monde, avec plus de des milliards d’invocations mensuelles sur l’infrastructure AWS. Depuis les annonces majeures du re:Invent 2025 – Lambda SnapStart, Lambda Managed Instances et l’expansion de SnapStart à Python – le service a franchi un cap décisif pour les développeurs qui cherchent à déployer des applications sans gérer de serveurs. Ce tutoriel vous guide pas à pas, de la création de votre compte AWS jusqu’au déploiement d’une API REST complète avec AWS Lambda, API Gateway et DynamoDB, le tout en Python 3.13.

**Last updated: April 10, 2026**

Que vous soyez développeur backend cherchant à réduire vos coûts d’infrastructure ou architecte cloud souhaitant migrer vers le serverless, ce guide couvre les 12 étapes essentielles avec des blocs de code fonctionnels, des exemples de sortie réels et les pièges les plus courants à éviter. À la fin de ce tutoriel, vous aurez une API REST CRUD entièrement fonctionnelle, déployée en production, capable de gérer des milliers de requêtes par seconde pour moins de 1 € par mois.

## Prérequis et Versions Requises pour AWS Lambda en 2026

Avant de commencer ce tutoriel AWS Lambda, assurez-vous de disposer des outils suivants avec les versions minimales indiquées. L’écosystème serverless AWS a évolué considérablement en 2025-2026, et utiliser des versions obsolètes peut entraîner des incompatibilités avec les nouvelles fonctionnalités comme les Durable Functions ou SnapStart pour Python.

| Outil | Version minimale | Version recommandée | Rôle | 
|---|---|---|---|
| Python | 3.11 | 3.13 | Runtime Lambda et développement local | 
| AWS CLI | 2.15 | 2.22 | Gestion des ressources AWS en ligne de commande | 
| AWS SAM CLI | 1.100 | 1.131 | Déploiement et test local des fonctions Lambda | 
| Docker Desktop | 24.0 | 27.5 | Émulation locale de l’environnement Lambda | 
| Node.js (optionnel) | 20 LTS | 22 LTS | Alternative runtime pour Lambda | 
| Git | 2.40 | 2.47 | Gestion de version du code source | 
| Compte AWS | Free Tier actif | Free Tier actif | Hébergement et exécution des fonctions | 
| pip | 23.0 | 24.3 | Gestionnaire de paquets Python | 

Le Free Tier AWS Lambda offre 1 million de requêtes gratuites et 400 000 Go-secondes de calcul par mois. Pour ce tutoriel, vous ne dépasserez pas ces limites, ce qui signifie un coût nul pendant la phase d’apprentissage. Notez que depuis décembre 2025, AWS facture désormais séparément la phase INIT (initialisation) des fonctions Lambda, ce qui rend l’optimisation des cold starts encore plus importante.

## Étape 1 : Configurer Votre Environnement AWS et les Permissions IAM

La première étape de tout projet AWS Lambda consiste à configurer correctement les permissions IAM. Une erreur de configuration IAM est la cause numéro un des échecs de déploiement Lambda chez les développeurs débutants. Commencez par installer AWS CLI et configurer vos identifiants.

```
# Installer AWS CLI v2 sur Linux/macOS
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install
# Vérifier l'installation
aws --version
# aws-cli/2.22.x Python/3.13.x Linux/6.x.x
# Configurer les identifiants
aws configure
# AWS Access Key ID: AKIAIOSFODNN7EXAMPLE
# AWS Secret Access Key: wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
# Default region name: eu-west-3
# Default output format: json
# Installer SAM CLI
pip install aws-sam-cli
sam --version
# SAM CLI, version 1.131.0
```
Créez ensuite un rôle IAM dédié pour vos fonctions Lambda. Ce rôle doit inclure les permissions minimales nécessaires selon le principe du moindre privilège. Connectez-vous à la console AWS IAM, puis créez un rôle avec la politique de confiance suivante :

```
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}
```
Attachez la politique **AWSLambdaBasicExecutionRole** pour les logs CloudWatch, puis **AmazonDynamoDBFullAccess** pour l’accès à DynamoDB. En production, vous remplacerez cette dernière par une politique personnalisée restreignant l’accès à une table spécifique. La région **eu-west-3** (Paris) est recommandée pour les utilisateurs européens afin de minimiser la latence et respecter les exigences RGPD.

### Piège n°1 : Utiliser les identifiants root

Ne configurez jamais AWS CLI avec les identifiants de votre compte root. Créez un utilisateur IAM dédié avec des permissions programmatiques. En cas de fuite d’identifiants root, un attaquant peut accéder à l’intégralité de votre compte AWS, créer des instances coûteuses et supprimer vos données. Activez systématiquement l’authentification multi-facteurs (MFA) sur le compte root.

## Étape 2 : Créer Votre Première Fonction AWS Lambda avec Python

Initialisez un nouveau projet SAM qui servira de base à votre API REST serverless. SAM (Serverless Application Model) simplifie considérablement le déploiement en générant automatiquement les templates CloudFormation nécessaires. Depuis la version 1.120 de SAM CLI sortie en 2025, le support natif de Python 3.13 est intégré avec les optimisations SnapStart.

```
# Initialiser un projet SAM
sam init --runtime python3.13 --app-template hello-world --name lambda-api-crud
# Structure du projet généré
lambda-api-crud/
├── README.md
├── __init__.py
├── events/
│   └── event.json
├── hello_world/
│   ├── __init__.py
│   ├── app.py
│   └── requirements.txt
├── samconfig.toml
├── template.yaml
└── tests/
    ├── __init__.py
    └── unit/
        └── test_handler.py
```
Ouvrez le fichier **hello_world/app.py** et remplacez le contenu par votre premier handler Lambda. Ce handler minimaliste retourne un JSON avec un code HTTP 200. Le paramètre **event** contient les données de la requête entrante (headers, body, query string), tandis que **context** fournit les métadonnées d’exécution (mémoire disponible, temps restant, nom de la fonction).

```
import json
def lambda_handler(event, context):
    """Handler principal de la fonction Lambda."""
    return {
        "statusCode": 200,
        "headers": {
            "Content-Type": "application/json",
            "Access-Control-Allow-Origin": "*"
        },
        "body": json.dumps({
            "message": "Bonjour depuis AWS Lambda !",
            "runtime": "Python 3.13",
            "region": "eu-west-3"
        })
    }
```
Testez votre fonction localement avant tout déploiement. SAM CLI émule l’environnement Lambda grâce à Docker, ce qui vous permet de détecter les erreurs sans consommer de ressources AWS. Cette étape est cruciale : selon AWS, 67 % des erreurs de déploiement Lambda peuvent être détectées par un test local préalable.

```
# Tester localement avec SAM
sam local invoke HelloWorldFunction --event events/event.json
# Sortie attendue :
# Invoking app.lambda_handler (python3.13)
# START RequestId: 1a2b3c4d-5e6f-7890-abcd-ef1234567890
# END RequestId: 1a2b3c4d-5e6f-7890-abcd-ef1234567890
# REPORT RequestId: 1a2b3c4d Duration: 3.45 ms Billed Duration: 4 ms Memory Size: 128 MB Max Memory Used: 36 MB
# {"statusCode": 200, "headers": {"Content-Type": "application/json", "Access-Control-Allow-Origin": "*"}, "body": "{\"message\": \"Bonjour depuis AWS Lambda !\", \"runtime\": \"Python 3.13\", \"region\": \"eu-west-3\"}"}
```
### Piège n°2 : Oublier Docker pour les tests locaux

