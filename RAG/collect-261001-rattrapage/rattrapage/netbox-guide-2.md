---
id: collect-261001-rattrapage/rattrapage/netbox-guide-2
title: "NetBox — Guide complet : source de vérité IPAM / DCIM"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/netbox_guide.md
source_anchor: ""
source_lines: [156, 411]
sha256: 5357e05b69258a65128ace44ec0299940f150310050764d56c0870353237e372
---

# Dépendances système de NetBox 4.x (Debian 12/13, Ubuntu 22.04/24.04)
apt install -y python3 python3-pip python3-venv python3-dev \
    build-essential libxml2-dev libxslt1-dev libffi-dev libpq-dev \
    libssl-dev zlib1g-dev

# Fuseau horaire (important pour les journaux et les tâches planifiées)
timedatectl set-timezone Europe/Paris
timedatectl status
```

Créez l'utilisateur système dédié (NetBox ne tourne JAMAIS en root) :

```bash
# Groupe + utilisateur sans shell interactif inutile, avec home
adduser --system --group --home /opt/netbox netbox
# Sur Ubuntu, équivalent :
# adduser --system --disabled-password --group --home /opt/netbox netbox
```

> 💡 **Bonne pratique** : un utilisateur système par application. Les fichiers de NetBox
> appartiennent à `netbox:netbox`, jamais à `root`.

---

## 6. Installation et configuration de PostgreSQL

```bash
# Installation (Debian 12 → PostgreSQL 15, Debian 13/Ubuntu 24.04 → 16)
apt install -y postgresql postgresql-contrib

# Démarrage + activation au boot
systemctl enable --now postgresql

# Vérification
pg_lsclusters
psql --version
```

Créez la base et l'utilisateur NetBox. ⚠️ Choisissez un mot de passe robuste et
stockez-le dans votre gestionnaire de mots de passe (ne le notez jamais dans ce guide) :

```bash
sudo -u postgres psql <<'EOF'
CREATE DATABASE netbox;
CREATE USER netbox WITH PASSWORD 'CHANGEZ_MOI_PgSQL_Exemple_2026!';
ALTER DATABASE netbox OWNER TO netbox;
-- Extension utile pour certaines recherches (optionnel)
\c netbox
CREATE EXTENSION IF NOT EXISTS pg_trgm;
EOF
```

Testez la connexion avec l'utilisateur applicatif :

```bash
PGPASSWORD='CHANGEZ_MOI_PgSQL_Exemple_2026!' \
  psql -h localhost -U netbox -d netbox -c 'SELECT version();'
```

Réglages PostgreSQL recommandés pour NetBox (`/etc/postgresql/<version>/main/postgresql.conf`) :

| Paramètre | Valeur conseillée | Pourquoi |
|---|---|---|
| `max_connections` | `200` | gunicorn ouvre plusieurs connexions |
| `shared_buffers` | `25 % de la RAM` | cache principal de PostgreSQL |
| `work_mem` | `8 MB` | tris/hachages des requêtes |
| `listen_addresses` | `localhost` | pas d'écoute réseau inutile si tout est local |

```bash
systemctl restart postgresql
```

---

## 7. Installation et configuration de Redis

NetBox 4.x utilise Redis pour le cache applicatif et les files de tâches (django-rq).
Redis 7 est recommandé.

```bash
apt install -y redis-server

# Sécurisation de base : écoute locale uniquement
sed -i 's/^bind .*/bind 127.0.0.1 ::1/' /etc/redis/redis.conf
# (Optionnel mais recommandé) mot de passe Redis :
# sed -i 's/^# requirepass .*/requirepass CHANGEZ_MOI_Redis_Exemple_2026!/' /etc/redis/redis.conf

systemctl enable --now redis-server
redis-cli ping   # doit répondre : PONG
```

> ⚠️ Si vous activez `requirepass`, répercutez-le dans `configuration.py`
> (`REDIS['password']`), sinon NetBox ne pourra pas se connecter au cache.

Test de persistance (Redis doit survivre à un redémarrage) :

```bash
redis-cli set test_netbox ok
systemctl restart redis-server
redis-cli get test_netbox   # doit répondre : ok
redis-cli del test_netbox
```

---

## 8. Téléchargement et installation de NetBox

Travaillez désormais en tant qu'utilisateur `netbox`. Repérez la dernière version 4.x
sur la [page des releases GitHub](https://github.com/netbox-community/netbox/releases).

```bash
# Passer sous l'utilisateur netbox
sudo -iu netbox
cd /opt/netbox

# Télécharger l'archive de la version voulue (exemple : 4.2.x)
wget https://github.com/netbox-community/netbox/archive/refs/tags/v4.2.3.tar.gz \
     -O netbox.tar.gz
tar -xzf netbox.tar.gz
ln -s netbox-4.2.3 netbox
cd netbox

# Environnement virtuel Python
python3 -m venv /opt/netbox/venv
source /opt/netbox/venv/bin/activate

# Dépendances Python (peut prendre plusieurs minutes)
pip install --upgrade pip
pip install -r requirements.txt
```

Vérifiez que l'installation s'est bien passée :

```bash
python -m netbox --version 2>/dev/null || python netbox/manage.py version
# ou simplement :
/opt/netbox/venv/bin/python -c "import django; print(django.get_version())"
```

> 💡 **Astuce de mise à jour future** : le lien symbolique `/opt/netbox/netbox` pointe
> vers la version active. Pour monter de version, on télécharge la nouvelle archive,
> on bascule le lien, on rejoue les migrations. (Détails : section 66.)

---

## 9. Le fichier `configuration.py` : principes

`configuration.py` est LE fichier de configuration de NetBox. Modèle fourni :
`netbox/netbox/configuration_example.py`.

```bash
# Toujours en utilisateur netbox, dans /opt/netbox/netbox/netbox/
sudo -iu netbox
cd /opt/netbox/netbox/netbox
cp configuration_example.py configuration.py
chmod 600 configuration.py
```

Règles d'or :

1. **Permissions 600** (`netbox:netbox`) : ce fichier contient des secrets (clé Django,
   mots de passe). Personne d'autre ne doit pouvoir le lire.
2. **Versionnez un modèle**, pas le fichier réel : gardez `configuration.py` hors de git,
   ou chiffrez-le (Ansible Vault, SOPS).
3. **Après chaque modification** : `systemctl restart netbox netbox-rq`.
4. **Validez la syntaxe Python** avant de redémarrer :
   `python -m py_compile configuration.py && echo OK`.

Structure minimale d'un `configuration.py` fonctionnel (détaillée sections 10 à 13) :

```python
ALLOWED_HOSTS = ['netbox.lan-entreprise.fr', '10.0.0.10']
DATABASE = { ... }      # PostgreSQL
REDIS = { ... }         # Redis : cache + files d'attente
SECRET_KEY = '...'      # >= 50 caractères aléatoires
```

---

## 10. `configuration.py` : base de données et paramètres vitaux

```python
# --- Hôtes autorisés (anti Host header attack) ---
ALLOWED_HOSTS = [
    'netbox.lan-entreprise.fr',
    '10.10.0.15',          # IP du serveur si accès direct
]

# --- Base de données PostgreSQL ---
DATABASE = {
    'NAME': 'netbox',
    'USER': 'netbox',
    'PASSWORD': 'CHANGEZ_MOI_PgSQL_Exemple_2026!',  # identique à la section 6
    'HOST': 'localhost',
    'PORT': '',            # '' = socket/port par défaut
    'CONN_MAX_AGE': 300,   # connexions persistantes (secondes)
}

# --- Clé secrète Django : GÉNÉREZ-LA, ne l'inventez pas ---
# python3 -c "import secrets; print(secrets.token_urlsafe(50))"
SECRET_KEY = 'CHANGEZ_MOI_Cle_Aleatoire_50_Caracteres_Minimum_Exemple='

# --- Redis : cache + files de tâches ---
REDIS = {
    'tasks': {
        'HOST': 'localhost',
        'PORT': 6379,
        'DATABASE': 0,
        # 'PASSWORD': 'CHANGEZ_MOI_Redis_Exemple_2026!',  # si requirepass activé
        'SSL': False,
    },
    'caching': {
        'HOST': 'localhost',
        'PORT': 6379,
        'DATABASE': 1,
        # 'PASSWORD': 'CHANGEZ_MOI_Redis_Exemple_2026!',
        'SSL': False,
    },
}
```

> ⚠️ **Ne réutilisez JAMAIS la `SECRET_KEY` d'exemple** entre deux installations :
> elle sert à signer les sessions et les jetons. Une clé connue = sessions falsifiables.

---

## 11. `configuration.py` : e-mail, fuseau horaire et divers

```python
# --- E-mail (notifications, réinitialisation de mot de passe) ---
EMAIL = {
    'SERVER': 'smtp.lan-entreprise.fr',
    'PORT': 587,
    'USERNAME': 'netbox@lan-entreprise.fr',
    'PASSWORD': 'CHANGEZ_MOI_SMTP_Exemple_2026!',
    'USE_SSL': False,
    'USE_TLS': True,
    'TIMEOUT': 10,
    'FROM_EMAIL': 'netbox@lan-entreprise.fr',
}

# --- Localisation ---
TIME_ZONE = 'Europe/Paris'
DATE_FORMAT = 'd/m/Y'      # format français dans l'interface
SHORT_DATE_FORMAT = 'd/m/Y'
TIME_FORMAT = 'H:i'

# --- Divers utiles en production ---
ADMINS = [('Exploitation', 'exploitation@lan-entreprise.fr')]
LOGIN_REQUIRED = True          # oblige l'authentification pour TOUT (recommandé)
BANNER_TOP = '⚠ PRODUCTION — NetBox de la DSI — toute modification est tracée'
MAINTENANCE_MODE = False
```

Testez l'envoi d'e-mail après le premier démarrage :

