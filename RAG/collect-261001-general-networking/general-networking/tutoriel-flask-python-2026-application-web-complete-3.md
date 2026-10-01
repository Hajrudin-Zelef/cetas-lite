---
id: collect-261001-general-networking/general-networking/tutoriel-flask-python-2026-application-web-complete-3
title: "Étape 1.1 : Créer le dossier du projet"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-general-networking/tutoriel-flask-python-2026-application-web-complete.md
source_anchor: ""
source_lines: [199, 276]
sha256: a1d05d502f46cbf4b990d4f5512d94c010b4213e68babe12eb463ff377ea4cfa
---

# Étape 1.1 : Créer le dossier du projet

Notre application supporte SQLite pour le développement (aucune installation requise) et PostgreSQL pour la production. Ce choix est délibéré : SQLite est parfait pour les tests locaux et les petits projets, tandis que PostgreSQL offre des performances et une robustesse adaptées à la production. Pour basculer entre les deux, il suffit de modifier la variable **DATABASE_URL** dans votre fichier **.env**. Le pilote PostgreSQL pour Python s’installe avec **pip install psycopg2-binary**.

Flask-Migrate, basé sur Alembic, gère les migrations de base de données. Les migrations permettent de faire évoluer le schéma de votre base de données de manière contrôlée et versionnée, sans perdre les données existantes. C’est un outil indispensable en production. Les commandes essentielles à retenir sont : **flask db init** (initialisation du répertoire migrations, à exécuter une seule fois), **flask db migrate -m “description”** (génération d’une nouvelle migration) et **flask db upgrade** (application des migrations en attente sur la base de données).

La configuration du pool de connexions SQLAlchemy est critique en production. Par défaut, SQLAlchemy maintient un pool de 5 connexions actives avec un débordement maximum de 10. Pour une application sous charge, ajustez ces paramètres dans votre configuration de production : **SQLALCHEMY_POOL_SIZE=20**, **SQLALCHEMY_MAX_OVERFLOW=40** et **SQLALCHEMY_POOL_TIMEOUT=30**. Une mauvaise configuration du pool est l’une des causes les plus fréquentes d’erreurs de performance en production.

## Étape 5 : Création des modèles de données

Les modèles de données définissent la structure de votre base de données à travers des classes Python. Chaque classe hérite de **db.Model** et chaque attribut de classe représente une colonne de la table correspondante. Flask-SQLAlchemy traduit automatiquement ces définitions en instructions SQL **CREATE TABLE**. Une conception soigneuse des modèles dès le début évite de nombreuses migrations correctives ultérieures.

```
# app/models/user.py
from datetime import datetime, timezone
from flask_login import UserMixin
from werkzeug.security import generate_password_hash, check_password_hash
from app.extensions import db
class User(UserMixin, db.Model):
    __tablename__ = 'users'
    id = db.Column(db.Integer, primary_key=True)
    username = db.Column(db.String(80), unique=True, nullable=False, index=True)
    email = db.Column(db.String(120), unique=True, nullable=False, index=True)
    password_hash = db.Column(db.String(256), nullable=False)
    is_active = db.Column(db.Boolean, default=True, nullable=False)
    created_at = db.Column(db.DateTime, default=lambda: datetime.now(timezone.utc))
    updated_at = db.Column(db.DateTime, default=lambda: datetime.now(timezone.utc),
                           onupdate=lambda: datetime.now(timezone.utc))
    # Relation One-to-Many avec Task
    tasks = db.relationship('Task', backref='owner', lazy='dynamic',
                            cascade='all, delete-orphan')
    def set_password(self, password):
        self.password_hash = generate_password_hash(password)
    def check_password(self, password):
        return check_password_hash(self.password_hash, password)
    def to_dict(self):
        return {
            'id': self.id,
            'username': self.username,
            'email': self.email,
            'created_at': self.created_at.isoformat()
        }
# app/models/task.py
from datetime import datetime, timezone
from app.extensions import db
class Task(db.Model):
    __tablename__ = 'tasks'
    id = db.Column(db.Integer, primary_key=True)
    title = db.Column(db.String(200), nullable=False)
    description = db.Column(db.Text, nullable=True)
    status = db.Column(db.String(20), default='en_attente', nullable=False)
    priority = db.Column(db.Integer, default=1, nullable=False)  # 1=faible, 2=moyenne, 3=haute
    due_date = db.Column(db.DateTime, nullable=True)
    created_at = db.Column(db.DateTime, default=lambda: datetime.now(timezone.utc))
    updated_at = db.Column(db.DateTime, default=lambda: datetime.now(timezone.utc),
                           onupdate=lambda: datetime.now(timezone.utc))
    user_id = db.Column(db.Integer, db.ForeignKey('users.id'), nullable=False, index=True)
    def to_dict(self):
        return {
            'id': self.id,
            'title': self.title,
            'description': self.description,
            'status': self.status,
            'priority': self.priority,
            'due_date': self.due_date.isoformat() if self.due_date else None,
            'created_at': self.created_at.isoformat(),
            'updated_at': self.updated_at.isoformat(),
            'user_id': self.user_id
        }
```
Plusieurs points méritent une attention particulière dans ces modèles. L’utilisation de **UserMixin** de Flask-Login fournit automatiquement les méthodes **is_authenticated**, **is_active**, **is_anonymous** et **get_id()**, nécessaires au bon fonctionnement de l’authentification par session. Le hachage des mots de passe via **werkzeug.security** garantit que les mots de passe ne sont jamais stockés en clair dans la base de données.

La relation **backref=’owner’** sur le modèle User crée automatiquement un attribut **owner** sur le modèle Task, permettant d’accéder à l’utilisateur propriétaire d’une tâche via **tache.owner**. L’option **cascade=’all, delete-orphan’** assure que la suppression d’un utilisateur supprime automatiquement toutes ses tâches associées, évitant les enregistrements orphelins. L’ajout d’index sur les colonnes **username**, **email** et **user_id** est essentiel pour maintenir de bonnes performances de recherche même avec des millions d’enregistrements.

Après avoir défini vos modèles, initialisez les migrations et créez la base de données avec les commandes suivantes : **flask db init**, puis **flask db migrate -m “modeles initiaux”** et enfin **flask db upgrade**. Vérifiez que le fichier **gestion_taches.db** a bien été créé dans votre répertoire courant en listant les fichiers avec **ls -la**.

## Étape 6 : Implémentation des routes CRUD

L’implémentation des routes CRUD (Create, Read, Update, Delete) constitue le cœur fonctionnel de notre **flask api rest python**. Chaque opération correspond à une méthode HTTP spécifique : POST pour créer, GET pour lire, PUT/PATCH pour mettre à jour et DELETE pour supprimer. Flask rend cette correspondance naturelle et intuitive grâce à son système de routage basé sur des décorateurs Python.

