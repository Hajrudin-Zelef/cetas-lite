---
id: collect-261001-general-networking/general-networking/tutoriel-aws-lambda-python-2026-api-rest-en-12-etapes-7
title: "Installer AWS CLI v2 sur Linux/macOS"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Lambda"]
dates: []
keywords: ["aws", "compute"]
source: docs/RAG/collect-261001-general-networking/tutoriel-aws-lambda-python-2026-api-rest-en-12-etapes.md
source_anchor: ""
source_lines: [659, 738]
sha256: cbe7f3a53704a028805b653d9e6998376c20d8b168228e55b73dda5ef48c9ef6
---

# Installer AWS CLI v2 sur Linux/macOS

**Déploiements canary et alias** – Utilisez les alias Lambda avec des déploiements progressifs (canary) pour tester les nouvelles versions en production avec un pourcentage limité du trafic. SAM supporte nativement les **DeploymentPreference** de type Canary10Percent5Minutes (10 % du trafic pendant 5 minutes, puis bascule complète si aucune alarme ne se déclenche). Cette approche réduit drastiquement le risque des mises en production.

## Projet Complet : Structure Finale de l’API REST Serverless

Voici la structure finale de votre projet après avoir suivi les 12 étapes de ce tutoriel. Cette architecture respecte les bonnes pratiques AWS Well-Architected Framework pour le serverless et peut servir de base pour des applications de production.

```
# Structure finale du projet
lambda-api-crud/
├── .github/
│   └── workflows/
│       └── deploy.yml          # Pipeline CI/CD GitHub Actions
├── crud_api/
│   ├── __init__.py
│   ├── app.py                  # Handler principal avec Powertools
│   └── requirements.txt        # Dépendances de la fonction
├── layers/
│   └── common/
│       └── python/             # Dépendances partagées (Layer)
├── tests/
│   ├── __init__.py
│   ├── unit/
│   │   └── test_handler.py     # Tests unitaires
│   └── integration/
│       └── test_api.py         # Tests d'intégration
├── events/
│   ├── create_product.json     # Événement test POST
│   ├── get_products.json       # Événement test GET
│   └── delete_product.json     # Événement test DELETE
├── samconfig.toml              # Configuration SAM (générée)
├── template.yaml               # Template SAM (infrastructure)
└── README.md
```
Cette API REST serverless peut gérer entre 1 000 et 10 000 requêtes par seconde sans aucune modification, grâce au scaling automatique d’AWS Lambda et de DynamoDB en mode PAY_PER_REQUEST. Le coût pour 1 million de requêtes mensuelles avec une durée moyenne de 150 ms est inférieur à **3 €** sur arm64. Pour aller plus loin, consultez le blog AWS Compute qui publie régulièrement des études de cas et des optimisations pour Lambda.

## Nouvelles Fonctionnalités AWS Lambda 2025-2026 : Ce Qui a Changé

L’année 2025-2026 a été marquée par des annonces majeures pour AWS Lambda. Au re:Invent 2025 (décembre), AWS a dévoilé trois fonctionnalités qui redéfinissent le serverless : les Durable Functions, les Managed Instances et le provisioned mode pour SQS. Voici un résumé des changements les plus importants qui impactent directement les développeurs.

Les **Lambda SnapStart** permettent des exécutions multi-étapes allant jusqu’à un an, avec gestion automatique de l’état et reprise après échec. Cette fonctionnalité remplace de nombreux cas d’usage de Step Functions pour les workflows simples à moyens. Les **Lambda Managed Instances** fusionnent le modèle serverless avec les instances EC2, offrant des performances prévisibles pour les charges de travail intensives. Le **provisioned mode pour SQS** élimine les pics de throttling lors du traitement de files d’attente à fort débit.

Côté runtimes, AWS supporte désormais Python 3.13 et 3.14, Node.js 22 et 24, Java 21 avec SnapStart natif, et .NET 10 avec Native AOT. L’expansion de SnapStart à Python est l’amélioration de performance la plus significative de 2025, avec une réduction des cold starts allant jusqu’à 4,3x. La facturation séparée de la phase INIT encourage les développeurs à optimiser l’initialisation de leurs fonctions, un changement qui a poussé 34 % des utilisateurs Lambda à réduire leurs imports inutiles selon les données internes AWS.

### Couverture associée

Pour approfondir vos connaissances sur les technologies associées, consultez nos autres tutoriels et comparatifs :

## FAQ : Questions Fréquentes sur AWS Lambda

**Quel est le coût réel d’AWS Lambda pour une petite application ?**

Pour une application recevant 50 000 requêtes par mois avec une durée moyenne de 200 ms et 256 Mo de mémoire, le coût est de 0 € grâce au Free Tier (1 million de requêtes et 400 000 Go-secondes inclus). Au-delà du Free Tier, le même volume coûterait environ 0,50 € par mois sur arm64.

**AWS Lambda est-il adapté aux applications à fort trafic ?**

Oui. AWS Lambda gère nativement jusqu’à 1 000 invocations concurrentes par région (extensible sur demande). Des entreprises comme Coca-Cola, Netflix et iRobot utilisent Lambda en production avec des millions d’invocations par heure. Pour les charges de travail prévisibles à très fort trafic, les Lambda Managed Instances (2025) offrent des performances encore plus stables.

**Quelle est la différence entre Lambda et EC2 pour une API REST ?**

Lambda est facturé à l’exécution (pay-per-use), tandis qu’EC2 est facturé à l’heure, même sans trafic. Pour les APIs avec un trafic variable ou faible, Lambda est 5 à 50x moins cher. Pour les charges constantes et prévisibles (plus de 70 % d’utilisation), EC2 peut être plus économique. Les Managed Instances de Lambda (2025) comblent cet écart en combinant les deux modèles.

**Comment réduire les cold starts AWS Lambda en 2026 ?**

Utilisez SnapStart pour Python (réduction 4,3x), passez à l’architecture arm64 (réduction 13-24 %), réduisez la taille du package sous 50 Mo, et activez Provisioned Concurrency pour les endpoints critiques. Évitez les imports inutiles et initialisez les clients SDK en dehors du handler.

**Peut-on utiliser AWS Lambda avec un VPC privé ?**

Oui, depuis les améliorations de 2019 (Hyperplane ENI), les fonctions Lambda dans un VPC n’ont plus de cold starts supplémentaires significatifs. Cependant, les fonctions dans un VPC nécessitent un NAT Gateway (environ 32 € par mois) pour accéder à Internet. Utilisez les VPC endpoints pour les services AWS (DynamoDB, S3, Secrets Manager) afin d’éviter le trafic Internet.

**Quels langages de programmation sont supportés par AWS Lambda en 2026 ?**

AWS Lambda supporte nativement Python 3.13/3.14, Node.js 22/24, Java 21, .NET 8/10, Ruby 3.3, et les runtimes personnalisés via Amazon Linux 2023. Python reste le langage le plus populaire avec 58 % des fonctions Lambda, suivi de Node.js à 31 % et Java à 7 %. Les runtimes personnalisés permettent d’utiliser n’importe quel langage compilé vers Linux, comme Rust ou Go.

**AWS Lambda est-il conforme au RGPD pour les données européennes ?**

Oui, à condition de déployer dans une région européenne (eu-west-1 Irlande, eu-west-3 Paris, eu-central-1 Francfort). AWS propose des Data Processing Agreements (DPA) et les régions européennes sont conformes au RGPD. Activez le chiffrement au repos avec AWS KMS et le chiffrement en transit (TLS 1.3 par défaut sur API Gateway) pour une conformité complète.

**Quelle est la limite de taille maximale d’une fonction Lambda ?**

Le package de déploiement ne peut pas dépasser 50 Mo (compressé .zip) ou 250 Mo (décompressé avec Layers). L’image de conteneur peut atteindre 10 Go. La mémoire configurable va de 128 Mo à 10 240 Mo (10 Go), et le timeout maximum est de 900 secondes (15 minutes) par invocation. Le stockage éphémère /tmp est configurable de 512 Mo à 10 240 Mo.
