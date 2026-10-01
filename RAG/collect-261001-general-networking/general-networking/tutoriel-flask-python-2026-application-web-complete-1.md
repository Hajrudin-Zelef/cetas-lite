---
id: collect-261001-general-networking/general-networking/tutoriel-flask-python-2026-application-web-complete-1
title: "Étape 1.1 : Créer le dossier du projet"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "mai"]
source: docs/RAG/collect-261001-general-networking/tutoriel-flask-python-2026-application-web-complete.md
source_anchor: ""
source_lines: [1, 77]
sha256: bdbb82ba62aecb0a41e22103b31aeab0ef0a29cdf2079490d4fb71ba39c2f0e4
---

# Étape 1.1 : Créer le dossier du projet

**Flask Python** s’impose en 2026 comme l’un des frameworks web les plus polyvalents et appréciés de l’écosystème Python, avec environ 71 400 étoiles sur GitHub et près de 40 millions de téléchargements mensuels sur PyPI selon le comparatif Uvik d’août 2026. Dans ce tutoriel Flask complet, nous allons construire de zéro une **API REST de gestion de tâches** entièrement fonctionnelle, depuis l’installation jusqu’au déploiement en production. Que vous soyez développeur débutant ou expérimenté, ce guide vous accompagnera pas à pas pour maîtriser Flask et créer une application web robuste, sécurisée et prête pour la production.

Publiée le 19 février 2026 selon le dépôt GitHub de Pallets, **Flask 3.1.3** devient la dernière version stable après une lignée de correctifs entamée avec **Flask 3.1.1** le 13 mai 2025 — le tout apportant des corrections de sécurité importantes et des améliorations de performances ; le projet recense désormais 2 743 838 dépôts dépendants sur GitHub. Flask reste dans le **top 3 des frameworks Python** aux côtés de Django et FastAPI, et le comparatif Sencha d’août 2026 continue de citer la version 3.1 comme référence pour les microservices, le déploiement de modèles d’apprentissage automatique et la création d’API légères. Sa position s’effrite toutefois face à la concurrence : l’enquête JetBrains/PSF relayée en juin 2026 montre que Flask reste utilisé par 34 % des développeurs web Python (part stable sur un an), tandis que FastAPI est passé de 14 % à 38 % d’adoption entre 2021 et 2025 selon le State of Python 2025 de JetBrains, dépassant désormais Flask ; le scorecard Uvik 2026 attribue d’ailleurs à Flask une note de 2,5/5, le qualifiant de framework « spécialiste », plutôt réservé aux projets legacy ou pédagogiques. Une version 4.0 est anticipée pour fin 2026, avec la prise en charge native de l’asynchrone et une meilleure sécurité de typage.

À la fin de ce tutoriel, vous disposerez d’une application de gestion de tâches complète avec authentification JWT, base de données relationnelle, validation des données, gestion des erreurs, suite de tests et déploiement conteneurisé avec Docker et Gunicorn.

## Prérequis et versions requises

Avant de plonger dans la création de notre application web avec **flask python**, il est indispensable de vérifier que votre environnement de développement répond aux exigences minimales. Flask 3.1.3 impose des dépendances précises qu’il convient de respecter scrupuleusement pour éviter les incompatibilités — d’autant que, selon le tableau de statut EOL Risk mis à jour en août 2026, le support de sécurité de **Flask 3.1.1** (sorti le 13 mai 2025) s’est arrêté dès le 19 août 2025, à peine 0,27 an après sa sortie, et que celui de **Flask 3.0.0** (sorti le 30 septembre 2023) était déjà clos depuis le 18 janvier 2024, soit près de 2,6 ans avant août 2026 : rester sur une ancienne version expose donc rapidement votre projet à des failles non corrigées.

Depuis la version 3.1.0, Flask a abandonné la compatibilité avec **Python 3.8**. Vous devez obligatoirement utiliser **Python 3.9 ou supérieur**. En 2026, il est fortement recommandé d’utiliser Python 3.12 ou 3.13 pour bénéficier des dernières optimisations de performances et des correctifs de sécurité. La documentation officielle disponible sur docs.python.org détaille toutes les nouveautés de chaque version.

Ces exigences de version remontent à **Flask 3.1.0**, sorti le 13 novembre 2024, qui a relevé les versions minimales requises pour les bibliothèques du cœur de Flask ; les dépendances principales de Flask 3.1.3 sont donc les suivantes :

| Bibliothèque | Version minimale | Rôle | Version recommandée 2026 | 
|---|---|---|---|
| Python | 3.9+ | Langage de base | 3.12.x ou 3.13.x | 
| Werkzeug | >= 3.1 | Utilitaires WSGI | 3.1.x | 
| Jinja2 | >= 3.1 | Moteur de templates (support async) | 3.1.x | 
| ItsDangerous | >= 2.2 | Sérialisation sécurisée | 2.2.x | 
| Blinker | >= 1.9 | Système de signaux | 1.9.x | 
| Click | >= 8.1 | Interface ligne de commande | 8.1.x | 

Sur le plan des connaissances préalables, ce tutoriel suppose une maîtrise élémentaire de Python (fonctions, classes, décorateurs), une compréhension basique du protocole HTTP (méthodes GET, POST, PUT, DELETE) et des notions sur les bases de données relationnelles et SQL. Si vous souhaitez comparer Flask avec d’autres approches, notre Tutoriel FastAPI Python 2026 vous fournira une perspective complémentaire très utile.

Concernant les outils de développement, vous aurez besoin de : un éditeur de code (VS Code, PyCharm ou Vim), Git pour la gestion de versions, Docker Desktop si vous souhaitez suivre la section déploiement, et un client HTTP comme HTTPie ou Postman pour tester vos endpoints. Une connexion Internet stable est également nécessaire pour l’installation des dépendances via pip.

## Étape 1 : Installation et configuration de l’environnement

La première étape de tout projet **flask python** sérieux consiste à créer un environnement virtuel isolé. Cette pratique fondamentale évite les conflits de dépendances entre projets et garantit la reproductibilité de votre environnement de développement. En 2026, l’utilisation de **venv** (intégré à Python) ou de **uv** (l’outil nouvelle génération ultra-rapide) est vivement recommandée.

```
# Étape 1.1 : Créer le dossier du projet
mkdir flask-gestion-taches
cd flask-gestion-taches
# Étape 1.2 : Créer l'environnement virtuel (Python 3.12+)
python3 -m venv venv
# Étape 1.3 : Activer l'environnement virtuel
# Sur Linux/macOS :
source venv/bin/activate
# Sur Windows :
venv\Scripts\activate
# Étape 1.4 : Mettre à jour pip
pip install --upgrade pip
# Étape 1.5 : Installer Flask et les dépendances principales
pip install Flask==3.1.3
pip install Flask-SQLAlchemy==3.1.1
pip install Flask-Migrate==4.0.7
pip install Flask-Login==0.6.3
pip install Flask-JWT-Extended==4.7.1
pip install marshmallow==3.22.0
pip install flask-marshmallow==1.2.1
pip install marshmallow-sqlalchemy==1.3.0
pip install gunicorn==23.0.0
pip install python-dotenv==1.0.1
# Étape 1.6 : Geler les dépendances
pip freeze > requirements.txt
```
Une fois Flask installé, vérifiez l’installation en lançant la commande **flask –version** dans votre terminal. Vous devriez voir apparaître “Flask 3.1.3” ainsi que les versions de Python et Werkzeug. Si la commande n’est pas reconnue, assurez-vous que votre environnement virtuel est bien activé — le préfixe “(venv)” doit apparaître dans votre invite de commande.

Créez ensuite un fichier **.env** à la racine du projet pour stocker vos variables d’environnement sensibles. Cette pratique de sécurité est absolument indispensable : ne committez jamais vos clés secrètes directement dans votre code source. Ajoutez immédiatement **.env** à votre fichier **.gitignore**.

```
# Contenu du fichier .env
FLASK_APP=app
FLASK_ENV=development
FLASK_DEBUG=1
SECRET_KEY=votre-cle-secrete-tres-longue-et-aleatoire-ici
JWT_SECRET_KEY=votre-cle-jwt-differente-et-tout-aussi-secrete
DATABASE_URL=sqlite:///gestion_taches.db
# En production, remplacer par :
# DATABASE_URL=postgresql://utilisateur:motdepasse@localhost:5432/nom_base
```
La variable **SECRET_KEY** est utilisée par Flask pour signer les cookies de session et les tokens CSRF. Elle doit être une chaîne aléatoire d’au moins 32 caractères. En Python, vous pouvez en générer une avec la commande : **python -c “import secrets; print(secrets.token_hex(32))”**. Ne réutilisez jamais la même clé entre vos environnements de développement et de production.

## Étape 2 : Structure du projet Flask

