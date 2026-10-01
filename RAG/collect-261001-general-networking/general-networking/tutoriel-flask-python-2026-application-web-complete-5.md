---
id: collect-261001-general-networking/general-networking/tutoriel-flask-python-2026-application-web-complete-5
title: "Étape 1.1 : Créer le dossier du projet"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/tutoriel-flask-python-2026-application-web-complete.md
source_anchor: ""
source_lines: [367, 504]
sha256: 709dc54228cff7af92c306d576911f8d681a1552729664fd79e0bab76a57989b
---

# Étape 1.1 : Créer le dossier du projet

```
# app/schemas/task_schema.py
from datetime import datetime
from marshmallow import Schema, fields, validate, validates, ValidationError
class TaskSchema(Schema):
    id = fields.Int(dump_only=True)
    title = fields.Str(
        required=True,
        validate=validate.Length(min=1, max=200,
            error="Le titre doit contenir entre 1 et 200 caractères.")
    )
    description = fields.Str(
        allow_none=True,
        validate=validate.Length(max=2000,
            error="La description ne peut pas dépasser 2000 caractères.")
    )
    status = fields.Str(
        validate=validate.OneOf(
            ['en_attente', 'en_cours', 'termine', 'annule'],
            error="Statut invalide. Valeurs acceptées : en_attente, en_cours, termine, annule."
        ),
        missing='en_attente'
    )
    priority = fields.Int(
        validate=validate.Range(min=1, max=3,
            error="La priorité doit être 1 (faible), 2 (moyenne) ou 3 (haute).")
    )
    due_date = fields.DateTime(
        allow_none=True,
        format='iso',
        error_messages={'invalid': 'Format de date invalide. Utilisez le format ISO 8601.'}
    )
    user_id = fields.Int(dump_only=True)
    created_at = fields.DateTime(dump_only=True)
    updated_at = fields.DateTime(dump_only=True)
    @validates('due_date')
    def valider_date_echeance(self, valeur):
        if valeur and valeur < datetime.now():
            raise ValidationError(
                "La date d'échéance ne peut pas être dans le passé."
            )
# app/schemas/user_schema.py
from marshmallow import Schema, fields, validate, validates, ValidationError
import re
class UserRegistrationSchema(Schema):
    username = fields.Str(
        required=True,
        validate=validate.Length(min=3, max=80,
            error="Le nom d'utilisateur doit contenir entre 3 et 80 caractères.")
    )
    email = fields.Email(
        required=True,
        error_messages={'invalid': 'Adresse email invalide.'}
    )
    password = fields.Str(
        required=True,
        load_only=True,
        validate=validate.Length(min=8,
            error="Le mot de passe doit contenir au moins 8 caractères.")
    )
    @validates('password')
    def valider_complexite_mot_de_passe(self, valeur):
        if not re.search(r'[A-Z]', valeur):
            raise ValidationError("Le mot de passe doit contenir au moins une majuscule.")
        if not re.search(r'[0-9]', valeur):
            raise ValidationError("Le mot de passe doit contenir au moins un chiffre.")
```
Les schémas Marshmallow offrent un niveau de contrôle très fin sur la validation. Les champs marqués **dump_only=True** ne sont jamais pris en compte lors de la désérialisation (chargement de données), ce qui empêche les utilisateurs malveillants de modifier des champs critiques comme l'**id**, le **user_id** ou les horodatages. À l'inverse, les champs **load_only=True** comme les mots de passe ne sont jamais inclus dans la sérialisation, donc ne seront jamais retournés dans les réponses de l'API.

La méthode **@validates** permet d'ajouter des validations métier personnalisées qui vont au-delà des simples contraintes de format. Dans notre exemple, nous vérifions que la date d'échéance n'est pas dans le passé et que le mot de passe respecte des critères de complexité minimaux. Ces validations sont exécutées automatiquement lors de l'appel à **schema.validate()** ou **schema.load()**, et les erreurs sont agrégées et retournées en une seule réponse structurée.

## Étape 8 : Gestion des erreurs et middleware

Une API REST professionnelle doit retourner des réponses d'erreur cohérentes et informatives. Flask permet de définir des gestionnaires d'erreurs globaux qui interceptent les exceptions et les convertissent en réponses JSON structurées. Cette approche centralise la gestion des erreurs et garantit que votre API retourne toujours le même format de réponse, qu'il s'agisse d'une erreur 404, 422 ou 500, ce qui simplifie considérablement l'intégration côté client.

```
# app/utils/errors.py
from flask import jsonify
from sqlalchemy.exc import IntegrityError
def register_error_handlers(app):
    @app.errorhandler(400)
    def mauvaise_requete(erreur):
        return jsonify({
            'erreur': 'Mauvaise requête',
            'message': str(erreur.description),
            'code': 400
        }), 400
    @app.errorhandler(401)
    def non_autorise(erreur):
        return jsonify({
            'erreur': 'Non autorisé',
            'message': 'Authentification requise pour accéder à cette ressource.',
            'code': 401
        }), 401
    @app.errorhandler(403)
    def interdit(erreur):
        return jsonify({
            'erreur': 'Accès refusé',
            'message': "Vous n'avez pas les permissions nécessaires.",
            'code': 403
        }), 403
    @app.errorhandler(404)
    def non_trouve(erreur):
        return jsonify({
            'erreur': 'Ressource introuvable',
            'message': "La ressource demandée n'existe pas ou a été supprimée.",
            'code': 404
        }), 404
    @app.errorhandler(422)
    def entite_non_traitable(erreur):
        return jsonify({
            'erreur': 'Données invalides',
            'message': 'Les données fournies ne respectent pas le format attendu.',
            'code': 422
        }), 422
    @app.errorhandler(500)
    def erreur_interne(erreur):
        return jsonify({
            'erreur': 'Erreur interne du serveur',
            'message': 'Une erreur inattendue s\'est produite. Veuillez réessayer.',
            'code': 500
        }), 500
    @app.errorhandler(IntegrityError)
    def erreur_integrite(erreur):
        return jsonify({
            'erreur': 'Conflit de données',
            'message': "Cette ressource existe déjà ou viole une contrainte d'unicité.",
            'code': 409
        }), 409
```
Au-delà de la gestion des erreurs, les hooks **before_request** et **after_request** de Flask permettent d'implémenter des middlewares qui s'exécutent avant et après chaque requête. Ces hooks sont idéaux pour la journalisation des requêtes, la vérification des en-têtes CORS, la mise à jour des statistiques d'utilisation ou l'injection d'informations de contexte dans **g** (l'objet de contexte de requête Flask).

La gestion de l'**IntegrityError** SQLAlchemy est particulièrement importante. Cette exception est levée lorsqu'une contrainte de base de données est violée, par exemple lors de la tentative de création d'un utilisateur avec un email déjà existant. Sans ce gestionnaire, Flask retournerait une erreur 500 générique peu informative au lieu d'un 409 Conflict explicite. En production, pensez également à utiliser **db.session.rollback()** dans vos gestionnaires d'erreurs pour s'assurer que les transactions échouées sont correctement annulées.

## Étape 9 : Authentification avec Flask-Login et JWT

L'authentification est l'une des fonctionnalités les plus critiques de toute application web. Notre application utilise **Flask-JWT-Extended** pour l'authentification sans état des endpoints API (REST), ce qui est parfaitement adapté aux applications monopage (SPA) et aux clients mobiles. Les tokens JWT (JSON Web Tokens) permettent une authentification stateless, scalable et sécurisée, sans nécessiter de stockage de session côté serveur.

