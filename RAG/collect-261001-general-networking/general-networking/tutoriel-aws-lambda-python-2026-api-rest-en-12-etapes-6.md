---
id: collect-261001-general-networking/general-networking/tutoriel-aws-lambda-python-2026-api-rest-en-12-etapes-6
title: "Installer AWS CLI v2 sur Linux/macOS"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Lambda"]
dates: []
keywords: ["aws", "attention", "graviton", "parameters"]
source: docs/RAG/collect-261001-general-networking/tutoriel-aws-lambda-python-2026-api-rest-en-12-etapes.md
source_anchor: ""
source_lines: [574, 658]
sha256: 881d6bc50423dfbc9c0149ae6827eef01d06ac69d28158922ae67392a2799921
---

# Installer AWS CLI v2 sur Linux/macOS

```
# .github/workflows/deploy.yml
name: Deploy Lambda API
on:
  push:
    branches: [main]
jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-python@v5
        with:
          python-version: '3.13'
      - uses: aws-actions/setup-sam@v2
        with:
          use-installer: true
      - uses: aws-actions/configure-aws-credentials@v4
        with:
          aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
          aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
          aws-region: eu-west-3
      - run: sam build
      - run: sam deploy --no-confirm-changeset --no-fail-on-empty-changeset
# Variables d'environnement par stage dans template.yaml
Parameters:
  Stage:
    Type: String
    Default: dev
    AllowedValues: [dev, staging, prod]
Globals:
  Function:
    Environment:
      Variables:
        STAGE: !Ref Stage
        LOG_LEVEL: !If [IsProd, "WARNING", "INFO"]
Conditions:
  IsProd: !Equals [!Ref Stage, prod]
```
Pour les secrets, ne stockez jamais de clés API ou de mots de passe dans les variables d’environnement Lambda en clair. Utilisez AWS Secrets Manager avec la couche Lambda dédiée qui met en cache les secrets pendant la durée de vie du conteneur. Le coût de Secrets Manager est de 0,40 $ par secret par mois, un investissement minimal pour la sécurité. En production, activez également le chiffrement au repos avec une clé KMS dédiée et configurez les VPC endpoints si vos fonctions accèdent à des ressources dans un VPC privé.

## Tarification AWS Lambda 2026 : Comprendre les Coûts Réels

La tarification d’AWS Lambda reste l’un des modèles les plus avantageux du cloud computing, mais la nouvelle facturation séparée de la phase INIT introduite en 2025 mérite attention. Voici le détail des coûts pour la région eu-west-3 (Paris) en avril 2026, incluant les architectures x86_64 et arm64.

| Composant | x86_64 | arm64 (Graviton) | Free Tier mensuel | 
|---|---|---|---|
| Requêtes | 0,20 $ / million | 0,20 $ / million | 1 million | 
| Durée (Go-seconde) | 0,0000166667 $ | 0,0000133334 $ | 400 000 Go-s | 
| Provisioned Concurrency | 0,0000097222 $/Go-s | 0,0000077778 $/Go-s | – | 
| Stockage éphémère (>512 Mo) | 0,0000000309 $/Go-s | 0,0000000309 $/Go-s | 512 Mo inclus | 
| Lambda URL (requêtes) | 0,20 $ / million | 0,20 $ / million | 1 million | 

Pour une API recevant 100 000 requêtes par jour avec une durée moyenne de 200 ms et 256 Mo de mémoire, le coût mensuel sur arm64 est d’environ **4,80 €**, soit 95 % moins cher qu’un serveur EC2 t3.micro (environ 10 € par mois). Le Free Tier couvre largement les besoins de développement et de petites applications. La tarification arm64 offre une économie de 20 % sur la durée par rapport à x86_64, ce qui s’accumule rapidement pour les applications à fort trafic. Consultez la page de tarification officielle AWS Lambda pour les prix actualisés.

## Guide de Dépannage : 8 Erreurs Fréquentes et Leurs Solutions

Les erreurs AWS Lambda peuvent provenir de multiples sources : permissions IAM, configuration API Gateway, limites de service, problèmes de dépendances ou erreurs de code. Voici les 8 problèmes les plus fréquents rencontrés par les développeurs, avec leurs diagnostics et solutions testés en 2026.

**1. “Task timed out after X seconds”** – Votre fonction dépasse le timeout configuré (par défaut 3 secondes). Augmentez le timeout dans template.yaml (maximum 900 secondes) ou optimisez votre code. Les causes courantes : appels API externes lents, scans DynamoDB sur de grandes tables, ou connexions réseau bloquées dans un VPC sans NAT Gateway.

**2. “ModuleNotFoundError: No module named ‘xxx'”** – La dépendance n’est pas incluse dans le package. Vérifiez votre requirements.txt, exécutez **sam build** avant le déploiement, et confirmez que l’architecture correspond (arm64 vs x86_64). Pour les bibliothèques avec des extensions C, utilisez **–platform manylinux2014_aarch64**.

**3. “AccessDeniedException” sur DynamoDB** – Le rôle IAM de votre fonction n’a pas les permissions suffisantes. Vérifiez que la politique DynamoDBCrudPolicy est correctement attachée et que le nom de la table correspond. Utilisez **aws iam simulate-principal-policy** pour tester les permissions.

**4. Erreur 502 Bad Gateway d’API Gateway** – Le format de la réponse Lambda est incorrect. API Gateway attend un objet JSON avec les champs **statusCode** (entier), **headers** (objet) et **body** (chaîne JSON). Vérifiez que body est bien une chaîne (**json.dumps()**) et non un objet Python.

**5. “Rate Exceeded” / Throttling** – Vous avez atteint la limite de concurrence (1 000 par défaut par région). Demandez une augmentation via AWS Service Quotas ou implémentez une file d’attente SQS en amont. Le provisioned mode pour SQS, lancé au re:Invent 2025, élimine les pics de throttling pour les sources d’événements SQS.

**6. Cold starts excessifs (>2 secondes)** – Activez SnapStart pour Python, passez à arm64, réduisez la taille du package sous 50 Mo et utilisez Provisioned Concurrency pour les endpoints critiques. Évitez d’importer des bibliothèques lourdes comme pandas (150 Mo) si elles ne sont pas nécessaires.

**7. “CORS header ‘Access-Control-Allow-Origin’ missing”** – API Gateway ne renvoie pas les headers CORS. Ajoutez les headers dans la réponse de votre fonction Lambda ET configurez la section CORS dans API Gateway. Pour SAM, ajoutez **Cors: “‘*'”** sous la propriété Api dans Globals.

**8. “Unable to import module ‘app’: No module named ‘app'”** – Le chemin du handler est incorrect. Vérifiez que le champ Handler dans template.yaml correspond à la structure de votre répertoire : **dossier.fichier.fonction** (par exemple **crud_api/app.lambda_handler** → Handler: **app.lambda_handler** avec CodeUri: **crud_api/**).

## Astuces Avancées pour AWS Lambda en Production

Au-delà des bases, plusieurs techniques avancées distinguent les déploiements Lambda professionnels des projets d’apprentissage. Ces optimisations sont utilisées par les équipes DevOps gérant des milliers de fonctions Lambda en production.

**Lambda URLs vs API Gateway** – Pour les microservices internes qui n’ont pas besoin d’authentification avancée ou de transformation de requête, les Lambda URLs (lancées en 2022, améliorées en 2025) offrent un endpoint HTTP direct sans la latence additionnelle d’API Gateway (environ 30 ms). Le coût est identique, mais la simplicité d’architecture est un avantage. Consultez la documentation officielle des runtimes Lambda pour les compatibilités.

**Lambda Managed Instances** – Annoncées en novembre 2025, les Managed Instances permettent d’exécuter des fonctions Lambda sur des instances EC2 dédiées (y compris Graviton4) tout en conservant le modèle de programmation serverless. AWS gère le cycle de vie, le patching et le scaling. Cette fonctionnalité est disponible dans 5 régions, dont eu-west-1 (Irlande), et est idéale pour les charges de travail avec des exigences de performance prévisibles ou des besoins de conformité spécifiques.

**Power Tuning** – L’outil AWS Lambda Power Tuning teste automatiquement votre fonction avec différentes configurations de mémoire (128 Mo à 10 240 Mo) et identifie le rapport coût/performance optimal. Augmenter la mémoire augmente proportionnellement le CPU alloué, ce qui peut réduire la durée d’exécution suffisamment pour compenser le surcoût. Pour les fonctions CPU-bound, passer de 256 Mo à 1024 Mo peut réduire la durée de 60 % tout en n’augmentant le coût que de 10 %.

