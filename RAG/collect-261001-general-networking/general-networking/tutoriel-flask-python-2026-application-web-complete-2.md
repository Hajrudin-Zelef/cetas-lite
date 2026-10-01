---
id: collect-261001-general-networking/general-networking/tutoriel-flask-python-2026-application-web-complete-2
title: "Étape 1.1 : Créer le dossier du projet"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-general-networking/tutoriel-flask-python-2026-application-web-complete.md
source_anchor: ""
source_lines: [78, 198]
sha256: 054e5b7298545902c93edcb898d2120a511280a21d7c0e79925c12b17dc2ecd3
---

# Étape 1.1 : Créer le dossier du projet

Une architecture bien pensée est la clé d’un projet Flask maintenable sur le long terme. Contrairement à Django qui impose une structure rigide, Flask vous laisse une grande liberté d’organisation. Cette liberté est à la fois une force et un piège pour les débutants. La structure que nous allons adopter suit les bonnes pratiques de l’industrie et s’appuie sur le **pattern Application Factory**, qui permet de créer plusieurs instances de l’application (développement, test, production) avec des configurations différentes.

```
flask-gestion-taches/
├── app/
│   ├── __init__.py          # Application Factory
│   ├── config.py            # Configurations par environnement
│   ├── extensions.py        # Instances des extensions Flask
│   ├── models/
│   │   ├── __init__.py
│   │   ├── user.py          # Modèle utilisateur
│   │   └── task.py          # Modèle tâche
│   ├── routes/
│   │   ├── __init__.py
│   │   ├── auth.py          # Routes d'authentification
│   │   └── tasks.py         # Routes CRUD des tâches
│   ├── schemas/
│   │   ├── __init__.py
│   │   ├── user_schema.py   # Schémas de validation utilisateur
│   │   └── task_schema.py   # Schémas de validation tâche
│   └── utils/
│       ├── __init__.py
│       └── errors.py        # Gestionnaires d'erreurs
├── tests/
│   ├── __init__.py
│   ├── conftest.py          # Configuration des tests
│   ├── test_auth.py         # Tests d'authentification
│   └── test_tasks.py        # Tests CRUD
├── migrations/              # Migrations Alembic (généré automatiquement)
├── .env                     # Variables d'environnement (ne pas committer)
├── .gitignore
├── requirements.txt
├── Dockerfile
├── docker-compose.yml
└── wsgi.py                  # Point d'entrée WSGI pour Gunicorn
```
Cette structure modulaire présente plusieurs avantages décisifs. Premièrement, elle sépare clairement les responsabilités : les modèles gèrent les données, les routes gèrent les requêtes HTTP, les schémas valident les entrées et sorties. Deuxièmement, elle facilite les tests unitaires car chaque composant peut être testé isolément. Troisièmement, elle prépare l’application à évoluer vers une architecture de microservices si nécessaire.

Le fichier **extensions.py** joue un rôle central dans cette architecture. Il contient les instances de toutes les extensions Flask (SQLAlchemy, Migrate, Login, JWT) créées sans application associée. Ce pattern évite les importations circulaires qui sont l’un des problèmes les plus fréquents dans les projets Flask de taille moyenne. L’initialisation de ces extensions avec l’objet application se fait ensuite dans la factory, ce qui permet une configuration dynamique selon l’environnement.

Pour créer cette structure automatiquement, exécutez les commandes suivantes dans votre terminal. La création manuelle des dossiers et fichiers **__init__.py** vides est une étape souvent négligée par les débutants, mais elle est indispensable pour que Python reconnaisse ces répertoires comme des paquets importables. Utilisez la commande **touch app/models/__init__.py** pour créer des fichiers vides rapidement sur Linux et macOS.

## Étape 3 : Création de l’application Flask de base

Le cœur de notre projet repose sur le pattern **Application Factory**. Cette approche consiste à encapsuler la création de l’instance Flask dans une fonction, généralement appelée **create_app()**. Cela offre une flexibilité maximale pour configurer l’application différemment selon l’environnement d’exécution et simplifie grandement l’écriture des tests automatisés.

```
# app/extensions.py
from flask_sqlalchemy import SQLAlchemy
from flask_migrate import Migrate
from flask_login import LoginManager
from flask_jwt_extended import JWTManager
db = SQLAlchemy()
migrate = Migrate()
login_manager = LoginManager()
jwt = JWTManager()
# app/config.py
import os
from datetime import timedelta
class Config:
    SECRET_KEY = os.environ.get('SECRET_KEY', 'dev-secret-key-changez-moi')
    SQLALCHEMY_TRACK_MODIFICATIONS = False
    JWT_ACCESS_TOKEN_EXPIRES = timedelta(hours=1)
    JWT_REFRESH_TOKEN_EXPIRES = timedelta(days=30)
class DevelopmentConfig(Config):
    DEBUG = True
    SQLALCHEMY_DATABASE_URI = os.environ.get('DATABASE_URL', 'sqlite:///dev.db')
    SQLALCHEMY_ECHO = True
class TestingConfig(Config):
    TESTING = True
    SQLALCHEMY_DATABASE_URI = 'sqlite:///:memory:'
    JWT_ACCESS_TOKEN_EXPIRES = timedelta(minutes=5)
class ProductionConfig(Config):
    DEBUG = False
    SQLALCHEMY_DATABASE_URI = os.environ.get('DATABASE_URL')
config = {
    'development': DevelopmentConfig,
    'testing': TestingConfig,
    'production': ProductionConfig,
    'default': DevelopmentConfig
}
# app/__init__.py
from flask import Flask
from .config import config
from .extensions import db, migrate, login_manager, jwt
def create_app(config_name='default'):
    app = Flask(__name__)
    app.config.from_object(config[config_name])
    # Initialisation des extensions
    db.init_app(app)
    migrate.init_app(app, db)
    login_manager.init_app(app)
    jwt.init_app(app)
    # Enregistrement des blueprints
    from .routes.auth import auth_bp
    from .routes.tasks import tasks_bp
    app.register_blueprint(auth_bp, url_prefix='/api/auth')
    app.register_blueprint(tasks_bp, url_prefix='/api/tasks')
    # Enregistrement des gestionnaires d'erreurs
    from .utils.errors import register_error_handlers
    register_error_handlers(app)
    return app
# wsgi.py
import os
from dotenv import load_dotenv
from app import create_app
load_dotenv()
app = create_app(os.environ.get('FLASK_ENV', 'production'))
if __name__ == '__main__':
    app.run()
```
L’utilisation des **Blueprints** Flask est une autre bonne pratique fondamentale. Les Blueprints permettent de diviser votre application en composants modulaires réutilisables. Chaque blueprint encapsule un ensemble de routes, de templates et de fichiers statiques liés à une fonctionnalité spécifique. Dans notre cas, nous avons deux blueprints : **auth_bp** pour l’authentification et **tasks_bp** pour la gestion des tâches.

Testez votre application de base en lançant **flask run** depuis la racine du projet. Flask doit démarrer sans erreur et être accessible à l’adresse **http://127.0.0.1:5000**. Si vous obtenez une erreur “Could not import app”, vérifiez que votre variable d’environnement **FLASK_APP=app** est bien définie dans votre fichier **.env** et que python-dotenv est correctement installé dans votre environnement virtuel actif.

La séparation des configurations selon l’environnement (**DevelopmentConfig**, **TestingConfig**, **ProductionConfig**) est une pratique indispensable. En développement, **SQLALCHEMY_ECHO=True** affiche toutes les requêtes SQL dans la console, ce qui facilite le débogage. En production, cette option doit être désactivée pour éviter des fuites d’informations sensibles dans les journaux.

## Étape 4 : Configuration de la base de données avec Flask-SQLAlchemy

Flask-SQLAlchemy est l’extension ORM (Object-Relational Mapper) de référence pour les projets Flask Python. Elle encapsule SQLAlchemy, la bibliothèque ORM la plus puissante de l’écosystème Python, en l’intégrant harmonieusement avec le contexte d’application Flask. Grâce à Flask-SQLAlchemy, vous interagissez avec votre base de données en utilisant des objets Python plutôt que d’écrire du SQL brut, ce qui améliore la lisibilité du code et la sécurité contre les injections SQL.

