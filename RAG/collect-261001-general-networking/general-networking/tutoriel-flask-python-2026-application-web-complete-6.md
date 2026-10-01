---
id: collect-261001-general-networking/general-networking/tutoriel-flask-python-2026-application-web-complete-6
title: "Étape 1.1 : Créer le dossier du projet"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/tutoriel-flask-python-2026-application-web-complete.md
source_anchor: ""
source_lines: [505, 616]
sha256: 9c4ba82b439d4dbcb4a023048012ff901328b5b678265579b1e42fbc4292571f
---

# Étape 1.1 : Créer le dossier du projet

```
# app/routes/auth.py
from flask import Blueprint, request, jsonify
from flask_jwt_extended import (
    create_access_token, create_refresh_token,
    jwt_required, get_jwt_identity, get_jwt
)
from app.extensions import db
from app.models.user import User
from app.schemas.user_schema import UserRegistrationSchema
auth_bp = Blueprint('auth', __name__)
user_registration_schema = UserRegistrationSchema()
# Ensemble en mémoire pour les tokens invalidés (utiliser Redis en production)
tokens_invalides = set()
@auth_bp.route('/inscription', methods=['POST'])
def inscription():
    """Inscription d'un nouvel utilisateur."""
    donnees = request.get_json()
    erreurs = user_registration_schema.validate(donnees)
    if erreurs:
        return jsonify({'erreurs': erreurs}), 422
    if User.query.filter_by(email=donnees['email']).first():
        return jsonify({'erreur': 'Cette adresse email est déjà utilisée.'}), 409
    if User.query.filter_by(username=donnees['username']).first():
        return jsonify({'erreur': "Ce nom d'utilisateur est déjà pris."}), 409
    nouvel_utilisateur = User(
        username=donnees['username'],
        email=donnees['email']
    )
    nouvel_utilisateur.set_password(donnees['password'])
    db.session.add(nouvel_utilisateur)
    db.session.commit()
    access_token = create_access_token(identity=nouvel_utilisateur.id)
    refresh_token = create_refresh_token(identity=nouvel_utilisateur.id)
    return jsonify({
        'message': 'Inscription réussie.',
        'utilisateur': nouvel_utilisateur.to_dict(),
        'access_token': access_token,
        'refresh_token': refresh_token
    }), 201
@auth_bp.route('/connexion', methods=['POST'])
def connexion():
    """Connexion d'un utilisateur existant."""
    donnees = request.get_json()
    utilisateur = User.query.filter_by(email=donnees.get('email')).first()
    if not utilisateur or not utilisateur.check_password(donnees.get('password', '')):
        return jsonify({'erreur': 'Email ou mot de passe incorrect.'}), 401
    if not utilisateur.is_active:
        return jsonify({'erreur': 'Ce compte a été désactivé.'}), 403
    access_token = create_access_token(identity=utilisateur.id)
    refresh_token = create_refresh_token(identity=utilisateur.id)
    return jsonify({
        'access_token': access_token,
        'refresh_token': refresh_token,
        'utilisateur': utilisateur.to_dict()
    }), 200
@auth_bp.route('/deconnexion', methods=['DELETE'])
@jwt_required()
def deconnexion():
    """Invalide le token JWT courant."""
    jti = get_jwt()['jti']
    tokens_invalides.add(jti)
    return jsonify({'message': 'Déconnexion réussie.'}), 200
@auth_bp.route('/rafraichir', methods=['POST'])
@jwt_required(refresh=True)
def rafraichir_token():
    """Génère un nouveau token d'accès à partir du refresh token."""
    current_user_id = get_jwt_identity()
    nouveau_token = create_access_token(identity=current_user_id)
    return jsonify({'access_token': nouveau_token}), 200
```
Le mécanisme de **refresh token** est essentiel pour offrir une bonne expérience utilisateur tout en maintenant un niveau de sécurité élevé. Les access tokens ont une durée de vie courte (1 heure dans notre configuration) pour limiter l'impact d'un vol de token. Les refresh tokens, valables 30 jours, permettent d'obtenir un nouveau access token sans que l'utilisateur ait à se reconnecter. En production, stockez la liste des tokens invalidés dans **Redis** plutôt qu'en mémoire, pour conserver ce mécanisme de déconnexion entre les redémarrages du serveur et le partager entre tous les workers Gunicorn.

Voici un exemple de flux d'authentification complet depuis la ligne de commande avec HTTPie :

```
# Inscription d'un nouvel utilisateur
http POST http://localhost:5000/api/auth/inscription \
  username=marie_dupont \
  [email protected] \
  password=MonMotDePasse1!
# Réponse attendue (HTTP 201 Created) :
# {
#   "message": "Inscription réussie.",
#   "utilisateur": {"id": 1, "username": "marie_dupont", "email": "[email protected]"},
#   "access_token": "eyJhbGciOiJIUzI1NiIs...",
#   "refresh_token": "eyJhbGciOiJIUzI1NiIs..."
# }
# Connexion et récupération du token
http POST http://localhost:5000/api/auth/connexion \
  [email protected] \
  password=MonMotDePasse1!
# Créer une tâche avec le token (remplacer TOKEN par le vrai token)
http POST http://localhost:5000/api/tasks/ \
  "Authorization:Bearer TOKEN" \
  title="Finaliser le rapport mensuel" \
  priority:=3 \
  status=en_cours
# Réponse attendue (HTTP 201 Created) :
# {
#   "id": 1,
#   "title": "Finaliser le rapport mensuel",
#   "status": "en_cours",
#   "priority": 3,
#   "user_id": 1,
#   "created_at": "2026-03-30T14:25:00.000000"
# }
```
## Étape 10 : Tests unitaires et d'intégration

Les tests automatisés sont la garantie de la qualité et de la maintenabilité de votre application Flask. Une suite de tests complète vous permet de refactoriser votre code en toute confiance, de détecter les régressions avant qu'elles n'atteignent la production et de documenter le comportement attendu de votre API. En 2026, un projet professionnel sans tests automatisés est tout simplement inacceptable dans n'importe quel contexte professionnel sérieux.

