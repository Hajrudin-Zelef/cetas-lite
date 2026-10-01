---
id: collect-261001-general-networking/general-networking/tutoriel-flask-python-2026-application-web-complete-4
title: "Étape 1.1 : Créer le dossier du projet"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/tutoriel-flask-python-2026-application-web-complete.md
source_anchor: ""
source_lines: [277, 366]
sha256: 58b386d6bcec729b5b89f2d1ac57c81b340876222ae595ef9d7d09bfa0708942
---

# Étape 1.1 : Créer le dossier du projet

```
# app/routes/tasks.py
from flask import Blueprint, request, jsonify
from flask_jwt_extended import jwt_required, get_jwt_identity
from app.extensions import db
from app.models.task import Task
from app.schemas.task_schema import TaskSchema
tasks_bp = Blueprint('tasks', __name__)
task_schema = TaskSchema()
tasks_schema = TaskSchema(many=True)
@tasks_bp.route('/', methods=['GET'])
@jwt_required()
def get_tasks():
    """Récupère toutes les tâches de l'utilisateur connecté avec pagination."""
    current_user_id = get_jwt_identity()
    page = request.args.get('page', 1, type=int)
    per_page = request.args.get('per_page', 10, type=int)
    status_filter = request.args.get('status')
    query = Task.query.filter_by(user_id=current_user_id)
    if status_filter:
        query = query.filter_by(status=status_filter)
    tasks_paginated = query.order_by(Task.priority.desc(), Task.created_at.desc())\
                           .paginate(page=page, per_page=per_page, error_out=False)
    return jsonify({
        'taches': tasks_schema.dump(tasks_paginated.items),
        'total': tasks_paginated.total,
        'pages': tasks_paginated.pages,
        'page_courante': page
    }), 200
@tasks_bp.route('/', methods=['POST'])
@jwt_required()
def create_task():
    """Crée une nouvelle tâche pour l'utilisateur connecté."""
    current_user_id = get_jwt_identity()
    donnees = request.get_json()
    erreurs = task_schema.validate(donnees)
    if erreurs:
        return jsonify({'erreurs': erreurs}), 422
    nouvelle_tache = Task(
        title=donnees['title'],
        description=donnees.get('description'),
        status=donnees.get('status', 'en_attente'),
        priority=donnees.get('priority', 1),
        due_date=donnees.get('due_date'),
        user_id=current_user_id
    )
    db.session.add(nouvelle_tache)
    db.session.commit()
    return jsonify(task_schema.dump(nouvelle_tache)), 201
@tasks_bp.route('/
```
', methods=['GET'])
@jwt_required()
def get_task(task_id):
    """Récupère une tâche spécifique par son identifiant."""
    current_user_id = get_jwt_identity()
    tache = Task.query.filter_by(id=task_id, user_id=current_user_id).first_or_404()
    return jsonify(task_schema.dump(tache)), 200
@tasks_bp.route('/', methods=['PUT'])
@jwt_required()
def update_task(task_id):
    """Met à jour une tâche existante."""
    current_user_id = get_jwt_identity()
    tache = Task.query.filter_by(id=task_id, user_id=current_user_id).first_or_404()
    donnees = request.get_json()
    erreurs = task_schema.validate(donnees, partial=True)
    if erreurs:
        return jsonify({'erreurs': erreurs}), 422
    for cle, valeur in donnees.items():
        setattr(tache, cle, valeur)
    db.session.commit()
    return jsonify(task_schema.dump(tache)), 200
@tasks_bp.route('/', methods=['DELETE'])
@jwt_required()
def delete_task(task_id):
    """Supprime une tâche par son identifiant."""
    current_user_id = get_jwt_identity()
    tache = Task.query.filter_by(id=task_id, user_id=current_user_id).first_or_404()
    db.session.delete(tache)
    db.session.commit()
    return jsonify({'message': 'Tâche supprimée avec succès'}), 200   Remarquez l’utilisation systématique du décorateur **@jwt_required()** sur toutes les routes protégées. Ce décorateur de Flask-JWT-Extended vérifie automatiquement la présence et la validité du token JWT dans l’en-tête **Authorization: Bearer {token}** de chaque requête. Si le token est absent, expiré ou invalide, Flask retourne automatiquement une réponse 401 Unauthorized sans exécuter la fonction de vue.

La pagination via **paginate()** est une fonctionnalité essentielle pour toute API REST professionnelle. Sans pagination, une requête GET sur une table contenant des milliers d’enregistrements pourrait saturer la mémoire de votre serveur et ralentir considérablement les temps de réponse. Notre implémentation retourne le nombre total d’éléments, le nombre de pages et la page courante, permettant au client de naviguer efficacement dans les résultats.

La méthode **first_or_404()** de SQLAlchemy est une convention élégante qui retourne l’objet trouvé ou lève automatiquement une exception HTTP 404 si aucun résultat ne correspond. Couplée à notre gestionnaire d’erreurs global, elle garantit des réponses d’erreur cohérentes et informatives sans code répétitif dans chaque vue.

## Étape 7 : Validation des données avec Marshmallow

La validation des données entrantes est une responsabilité critique de toute API REST. **Marshmallow** est la bibliothèque de sérialisation et de validation de référence pour les projets **créer application web flask**. Elle permet de définir des schémas qui décrivent la structure attendue des données, de valider automatiquement les requêtes entrantes et de sérialiser les objets SQLAlchemy en dictionnaires JSON prêts à être retournés dans les réponses.

