---
id: collect-261001-rattrapage/rattrapage/10-idees-de-projets-docker-du-debutant-au-confirme-3
title: "Stage 1: Build"
domain: rattrapage
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-rattrapage/10-idees-de-projets-docker-du-debutant-au-confirme.md
source_anchor: ""
source_lines: [425, 499]
sha256: ba026b16fa9d89cc1cc67935d8fcd17ecb780b2a9f7e9e6d521b3f601856b9ff
---

# Stage 1: Build

@app.post("/predict/")
def predict(data: list):
   return {"prediction": model.predict(data)}
```
- Écrivez le fichier Docker : Créez un fichier Docker qui définit l'environnement pour FastAPI :

```
FROM python:3.9-slim
WORKDIR /app
COPY requirements.txt .
RUN pip install -r requirements.txt
COPY . .
CMD ["uvicorn", "app:app", "--host", "0.0.0.0", "--port", "8000"]
```
- Construisez l'image :

`docker build -t fastapi-app .`
- Exécutez le conteneur :

`docker run -p 8000:8000 fastapi-app`
## Conseils pour travailler sur des projets Docker

Pendant que vous travaillez sur ces projets, gardez les conseils suivants à l'esprit :

- Commencez modestement : Commencez par des projets légèrement difficiles, puis passez à des tâches plus complexes. Il est essentiel d'acquérir de l'assurance pour les tâches les plus simples.
- Consignez vos progrès : Conservez un cursus détaillé de vos projets afin de suivre votre apprentissage et de vous en servir comme référence pour vos projets futurs.
- Rejoignez les communautés Docker : Participez à des forums en ligne et à des rencontres locales pour partager vos expériences, poser des questions et apprendre des autres.
- Expérimentez et personnalisez : N'ayez pas peur de modifier les projets, d'essayer différentes approches et d'explorer les nouvelles fonctionnalités de Docker.
- Continuez à apprendre : Continuez à développer vos connaissances sur Docker en explorant des sujets et des outils avancés tels que Kubernetes, Docker Swarm ou l'architecture microservices.

## Conclusion

La maîtrise de Docker ne se limite pas à l'apprentissage des commandes et des configurations. Il s'agit de comprendre comment Docker s'intègre dans le développement d'applications modernes, les flux de travail de la science des données et la gestion de l'infrastructure.

Les projets présentés dans ce guide vous donnent quelques idées pour acquérir les compétences de base et l'expérience pratique nécessaires pour exceller dans des scénarios réels.

À ce stade, je vous suggère de consolider vos connaissances en suivant ces cours :

## Devenez ingénieur en données

## FAQ

### Quelles sont les meilleures pratiques pour écrire des Dockerfiles efficaces ?

**Les meilleures pratiques pour écrire des Dockerfiles efficaces comprennent la minimisation du nombre de couches en combinant des commandes, l'utilisation de constructions en plusieurs étapes pour réduire la taille de l'image, la sélection d'images de base légères, la mise en cache des dépendances et l'évitement d'inclure des fichiers inutiles dans l'image finale.**

### Qu'est-ce qu'une construction en plusieurs étapes dans Docker ?

**Une construction en plusieurs étapes est une méthode permettant d'optimiser les images Docker en séparant les environnements de construction et d'exécution. Cela permet d'obtenir des images plus petites et plus sûres.**

### Comment réduire la taille d'une image Docker ?

**Utilisez des images de base minimales, gérez efficacement les dépendances et utilisez des constructions en plusieurs étapes pour réduire la taille des images et améliorer les performances.**

### Comment résoudre les erreurs courantes lors de la création d'images Docker ?

**Les erreurs les plus courantes lors de la création d'images Docker sont les problèmes de permission, la syntaxe incorrecte du fichier Docker et l'échec de l'installation des dépendances. Pour résoudre le problème, vérifiez les journaux de construction de Docker, assurez-vous que vous utilisez l'image de base correcte et confirmez que les chemins d'accès ou les autorisations de fichiers sont définis correctement. Des outils tels que `docker build --no-cache` peuvent vous aider à identifier les problèmes de mise en cache.**

### Puis-je utiliser Docker avec Kubernetes pour ces projets ?

**Oui, une fois que vous êtes à l'aise avec Docker, Kubernetes peut être l'étape suivante. Kubernetes permet de gérer les applications conteneurisées à grande échelle. Vous pouvez déployer vos projets Docker sur un cluster Kubernetes pour gérer plusieurs instances, gérer la mise à l'échelle et automatiser les déploiements.**

### Quelles sont les meilleures pratiques pour gérer les volumes Docker et les données persistantes ?

**Lorsque vous travaillez avec des volumes Docker, il est important d'utiliser des volumes nommés pour garantir la persistance des données lors des redémarrages de conteneurs. Sauvegardez régulièrement vos volumes et surveillez les goulets d'étranglement dus aux E/S sur disque. Évitez de stocker directement des données sensibles dans des conteneurs ; utilisez plutôt des solutions de stockage sécurisées ou des bases de données externes.**

### À quoi sert la directive ENTRYPOINT dans un fichier Docker ?

**La directive `ENTRYPOINT` d'un fichier Docker spécifie la commande qui sera toujours exécutée au démarrage du conteneur. Il permet de traiter le conteneur comme un exécutable, où des arguments peuvent être passés pendant l'exécution, ce qui améliore la flexibilité.**

### Quelle est la différence entre CMD et ENTRYPOINT dans un fichier Docker ?

**Les adresses `CMD` et `ENTRYPOINT` indiquent les commandes à exécuter au démarrage d'un conteneur. Cependant, `CMD` fournit des arguments par défaut qui peuvent être remplacés, tandis que `ENTRYPOINT` définit la commande qui s'exécute toujours. `ENTRYPOINT` est utile pour créer des conteneurs qui agissent comme des exécutables, tandis que `CMD` est plus flexible pour spécifier des commandes par défaut.**

Architecte de solutions cloud certifié AWS, DevOps, ingénieur cloud avec une compréhension approfondie de l'architecture et des concepts de haute disponibilité. Je possède des connaissances en matière d'ingénierie du cloud et de DevOps et je sais utiliser les ressources open-source pour exécuter des applications d'entreprise. Je construis des applications basées sur le cloud en utilisant AWS, AWS CDK, AWS SAM, CloudFormation, Serverless Framework, Terraform et Django.
