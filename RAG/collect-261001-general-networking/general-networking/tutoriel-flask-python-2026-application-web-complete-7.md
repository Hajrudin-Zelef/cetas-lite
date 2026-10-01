---
id: collect-261001-general-networking/general-networking/tutoriel-flask-python-2026-application-web-complete-7
title: "Étape 1.1 : Créer le dossier du projet"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/tutoriel-flask-python-2026-application-web-complete.md
source_anchor: ""
source_lines: [617, 775]
sha256: 7f0b1ab27c2ec4d5d36948cb26030feac87c2334a10370931a90a174251c69df
---

# Étape 1.1 : Créer le dossier du projet

```
# tests/conftest.py
import pytest
from app import create_app
from app.extensions import db as _db
@pytest.fixture(scope='session')
def app():
    """Crée une instance de l'application pour les tests."""
    app = create_app('testing')
    with app.app_context():
        _db.create_all()
        yield app
        _db.drop_all()
@pytest.fixture(scope='function')
def client(app):
    """Retourne un client de test Flask."""
    return app.test_client()
@pytest.fixture(scope='function')
def db(app):
    """Fournit une session de base de données propre pour chaque test."""
    connection = _db.engine.connect()
    transaction = connection.begin()
    yield _db
    transaction.rollback()
    connection.close()
@pytest.fixture
def utilisateur_test(db):
    """Crée un utilisateur de test dans la base de données."""
    from app.models.user import User
    utilisateur = User(username='testeur', email='[email protected]')
    utilisateur.set_password('MotDePasse123!')
    db.session.add(utilisateur)
    db.session.commit()
    return utilisateur
@pytest.fixture
def token_auth(client, utilisateur_test):
    """Retourne un token JWT valide pour l'utilisateur de test."""
    reponse = client.post('/api/auth/connexion', json={
        'email': '[email protected]',
        'password': 'MotDePasse123!'
    })
    return reponse.get_json()['access_token']
# tests/test_tasks.py
def test_creer_tache(client, token_auth):
    """Teste la création d'une nouvelle tâche."""
    reponse = client.post(
        '/api/tasks/',
        json={'title': 'Ma première tâche', 'priority': 2},
        headers={'Authorization': f'Bearer {token_auth}'}
    )
    assert reponse.status_code == 201
    donnees = reponse.get_json()
    assert donnees['title'] == 'Ma première tâche'
    assert donnees['status'] == 'en_attente'
    assert donnees['priority'] == 2
def test_creer_tache_sans_authentification(client):
    """Teste que la création sans token échoue avec 401."""
    reponse = client.post('/api/tasks/', json={'title': 'Tâche non autorisée'})
    assert reponse.status_code == 401
def test_creer_tache_titre_manquant(client, token_auth):
    """Teste la validation : le titre est obligatoire."""
    reponse = client.post(
        '/api/tasks/',
        json={'description': 'Sans titre'},
        headers={'Authorization': f'Bearer {token_auth}'}
    )
    assert reponse.status_code == 422
def test_lister_taches_avec_pagination(client, token_auth):
    """Teste la récupération paginée des tâches."""
    reponse = client.get(
        '/api/tasks/?page=1&per_page=5',
        headers={'Authorization': f'Bearer {token_auth}'}
    )
    assert reponse.status_code == 200
    donnees = reponse.get_json()
    assert 'taches' in donnees
    assert 'total' in donnees
    assert 'pages' in donnees
```
Le pattern **conftest.py** de pytest est la manière standard de définir des fixtures partagées entre tous les fichiers de test. L'utilisation de transactions de base de données dans la fixture **db** garantit que chaque test démarre avec une base propre, sans données résiduelles des tests précédents. Cette isolation est fondamentale pour des tests déterministes et reproductibles, surtout dans un environnement d'intégration continue.

Pour lancer la suite de tests, installez pytest et pytest-cov avec **pip install pytest pytest-cov**, puis exécutez **pytest tests/ -v --cov=app --cov-report=html**. Cette commande lance tous les tests avec un rapport de couverture de code détaillé en HTML. Visez une couverture minimale de 80% pour votre code métier. Pour une intégration continue automatisée, notre Tutoriel GitHub Actions CI/CD 2026 vous guidera dans la mise en place d'une pipeline complète.

## Étape 11 : Déploiement avec Docker et Gunicorn

Le déploiement en production d'une application **flask python** nécessite plusieurs étapes importantes. Le serveur de développement intégré à Flask n'est pas conçu pour la production : il est mono-thread, non sécurisé et peu performant. En production, on utilise **Gunicorn** (disponible sur gunicorn.org) comme serveur WSGI devant Flask. **Docker** (docker.com) garantit la portabilité et la reproductibilité de l'environnement d'exécution. Pour une orchestration multi-conteneurs complète, consultez notre Tutoriel Docker Compose 2026.

```
# Dockerfile
FROM python:3.12-slim
# Créer un utilisateur non-root pour la sécurité
RUN groupadd -r flaskapp && useradd -r -g flaskapp flaskapp
# Définir le répertoire de travail
WORKDIR /app
# Copier et installer les dépendances (layer cache Docker)
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
# Copier le code source
COPY . .
# Changer le propriétaire des fichiers
RUN chown -R flaskapp:flaskapp /app
# Passer à l'utilisateur non-root
USER flaskapp
# Exposer le port applicatif
EXPOSE 8000
# Commande de démarrage avec Gunicorn
CMD ["gunicorn", \
     "--bind", "0.0.0.0:8000", \
     "--workers", "4", \
     "--worker-class", "sync", \
     "--worker-connections", "1000", \
     "--timeout", "30", \
     "--keepalive", "5", \
     "--max-requests", "1000", \
     "--max-requests-jitter", "100", \
     "--log-level", "info", \
     "--access-logfile", "-", \
     "--error-logfile", "-", \
     "wsgi:app"]
# docker-compose.yml
version: '3.9'
services:
  web:
    build: .
    ports:
      - "8000:8000"
    environment:
      - FLASK_ENV=production
      - DATABASE_URL=postgresql://flask_user:motdepasse@db:5432/flask_db
      - SECRET_KEY=${SECRET_KEY}
      - JWT_SECRET_KEY=${JWT_SECRET_KEY}
    depends_on:
      db:
        condition: service_healthy
    restart: unless-stopped
  db:
    image: postgres:16-alpine
    volumes:
      - postgres_data:/var/lib/postgresql/data
    environment:
      - POSTGRES_USER=flask_user
      - POSTGRES_PASSWORD=motdepasse
      - POSTGRES_DB=flask_db
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U flask_user -d flask_db"]
      interval: 10s
      timeout: 5s
      retries: 5
volumes:
  postgres_data:
```
Le nombre de workers Gunicorn doit être calibré selon les ressources de votre serveur. La règle de base est **(2 × nombre_de_cœurs_CPU) + 1**. Sur un serveur avec 2 cœurs, vous utiliserez donc 5 workers. L'option **--max-requests** redémarre chaque worker après un certain nombre de requêtes, ce qui prévient les fuites mémoire progressives. L'option **--max-requests-jitter** ajoute une variation aléatoire pour éviter que tous les workers redémarrent simultanément, ce qui provoquerait une brève interruption de service.

Pour construire et lancer l'application en production, exécutez **docker-compose up --build -d**. Vérifiez que les conteneurs sont bien démarrés avec **docker-compose ps** et consultez les journaux avec **docker-compose logs -f web**. N'oubliez pas d'appliquer les migrations de base de données dans le conteneur avec **docker-compose exec web flask db upgrade** avant d'ouvrir l'application au trafic.

## Pièges courants et solutions

Après avoir accompagné des centaines de développeurs sur des projets Flask, nous avons identifié les erreurs les plus fréquentes et leurs solutions. Éviter ces pièges vous fera économiser de nombreuses heures de débogage frustrant.

