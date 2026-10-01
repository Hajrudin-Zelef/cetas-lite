---
id: collect-261001-rattrapage/rattrapage/netbox-guide-3
title: "NetBox — Guide complet : source de vérité IPAM / DCIM"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/netbox_guide.md
source_anchor: ""
source_lines: [412, 672]
sha256: 77af5d48380f410dfcd8702bcea6f8761aaedff848c533c118d4c146ba66b96f
---

# NetBox — Guide complet : source de vérité IPAM / DCIM

```bash
cd /opt/netbox/netbox/netbox/
source /opt/netbox/venv/bin/activate
python ../manage.py sendtestemail exploitation@lan-entreprise.fr
```

---

## 12. Initialisation : migrations, statiques, superutilisateur

Toujours en utilisateur `netbox`, virtualenv activé, dans `/opt/netbox/netbox` :

```bash
cd /opt/netbox/netbox
source /opt/netbox/venv/bin/activate

# 1. Générer les SECRET_KEY-dependent artefacts + créer les tables
python netbox/manage.py migrate

# 2. Collecter les fichiers statiques (CSS/JS/images de l'interface)
python netbox/manage.py collectstatic --no-input

# 3. Créer le compte administrateur initial
python netbox/manage.py createsuperuser
#   Username: admin
#   Email: exploitation@lan-entreprise.fr
#   Password: <robuste, gestionnaire de mots de passe>

# 4. (Optionnel) charger des jeux de données de démonstration — UNIQUEMENT EN LAB
# python netbox/manage.py loaddata ...

# 5. Vérifier que tout est sain
python netbox/manage.py check
```

> ⚠️ `createsuperuser` n'est à lancer qu'une fois. Si vous perdez le mot de passe admin,
> régénérez-le avec `python netbox/manage.py changepassword admin`.

Ordre impératif : **migrate AVANT le premier démarrage de gunicorn**, sinon l'application
démarre sur une base vide et renvoie des erreurs 500.

---

## 13. Gunicorn : le serveur d'application

NetBox 4.x s'exécute via gunicorn (serveur WSGI). Fichier conseillé :
`/opt/netbox/gunicorn.py` (propriétaire `netbox:netbox`, 640) :

```python
# /opt/netbox/gunicorn.py
bind = '127.0.0.1:8001'      # écoute locale uniquement, nginx fait le frontal
workers = 5                  # règle : (2 x vCPU) + 1 ; ici pour 2 vCPU
worker_class = 'gthread'
threads = 3
timeout = 120                # les exports CSV/API volumineux prennent du temps
max_requests = 1000          # recycle les workers (anti-fuites mémoire)
max_requests_jitter = 100
preload_app = True
user = 'netbox'
group = 'netbox'
loglevel = 'info'
accesslog = '/var/log/netbox/gunicorn-access.log'
errorlog = '/var/log/netbox/gunicorn-error.log'
```

Créez le répertoire de logs :

```bash
mkdir -p /var/log/netbox
chown netbox:netbox /var/log/netbox
chmod 750 /var/log/netbox
```

Test manuel avant de créer le service systemd (doit afficher des logs sans erreur,
`Ctrl+C` pour arrêter) :

```bash
sudo -iu netbox
cd /opt/netbox/netbox
/opt/netbox/venv/bin/gunicorn --config /opt/netbox/gunicorn.py netbox.wsgi
```

---

## 14. Services systemd : `netbox` et `netbox-rq`

Deux services sont nécessaires : l'application web et les travailleurs de tâches
(webhooks, scripts, exports).

`/etc/systemd/system/netbox.service` :

```ini
[Unit]
Description=NetBox WSGI service (gunicorn)
After=network.target postgresql.service redis-server.service
Wants=postgresql.service redis-server.service

[Service]
Type=simple
User=netbox
Group=netbox
WorkingDirectory=/opt/netbox/netbox
ExecStart=/opt/netbox/venv/bin/gunicorn --config /opt/netbox/gunicorn.py netbox.wsgi
ExecReload=/bin/kill -s HUP $MAINPID
Restart=on-failure
RestartSec=5
# Sécurité systemd (durcissement)
NoNewPrivileges=true
PrivateTmp=true
ProtectSystem=strict
ReadWritePaths=/opt/netbox/netbox/netbox/media /var/log/netbox

[Install]
WantedBy=multi-user.target
```

`/etc/systemd/system/netbox-rq.service` :

```ini
[Unit]
Description=NetBox Request Queue worker (django-rq)
After=network.target postgresql.service redis-server.service netbox.service

[Service]
Type=simple
User=netbox
Group=netbox
WorkingDirectory=/opt/netbox/netbox
ExecStart=/opt/netbox/venv/bin/python netbox/manage.py rqworker high default low
Restart=on-failure
RestartSec=5

[Install]
WantedBy=multi-user.target
```

Activation :

```bash
systemctl daemon-reload
systemctl enable --now netbox netbox-rq
systemctl status netbox netbox-rq --no-pager
# Vérifier l'écoute locale :
ss -ltnp | grep 8001
```

> 💡 `rqworker high default low` : trois files de priorité. Les webhooks partent en
> `high`, les scripts/rapports en `default`. Si `netbox-rq` est arrêté, les webhooks
> s'accumulent dans Redis sans être envoyés — pensez-y en dépannage (section 70).

---

## 15. Nginx en reverse proxy (HTTP puis HTTPS)

`/etc/nginx/sites-available/netbox` :

```nginx
upstream netbox {
    server 127.0.0.1:8001;
}

server {
    listen 80;
    server_name netbox.lan-entreprise.fr;

    # Redirection systématique vers HTTPS (après Let's Encrypt, section 16)
    # return 301 https://$host$request_uri;

    client_max_body_size 10M;   # imports CSV, images d'équipements

    location /static/ {
        alias /opt/netbox/netbox/netbox/static/;
    }

    location / {
        proxy_pass http://netbox;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_connect_timeout 30s;
        proxy_read_timeout 120s;
    }
}
```

```bash
ln -s /etc/nginx/sites-available/netbox /etc/nginx/sites-enabled/
nginx -t && systemctl reload nginx
```

Vérifiez : `curl -I http://netbox.lan-entreprise.fr/` doit répondre `200 OK`
(ou `302` vers `/login/` si `LOGIN_REQUIRED = True`).

> ⚠️ Servez TOUJOURS les fichiers `/static/` via nginx, jamais via gunicorn :
> c'est plus rapide et ça évite de saturer les workers Python.

---

## 16. HTTPS avec Let's Encrypt (ou PKI interne)

En production, NetBox contient des données sensibles : **HTTPS obligatoire**.

**Option A — Let's Encrypt (serveur exposé sur Internet) :**

```bash
apt install -y certbot python3-certbot-nginx
certbot --nginx -d netbox.lan-entreprise.fr
# Renouvellement automatique déjà planifié par le paquet ; vérifiez :
certbot renew --dry-run
```

**Option B — PKI interne d'entreprise (cas le plus courant pour Zelef) :**

1. Générez une CSR ou demandez un certificat à votre AD CS / PKI interne pour
   `netbox.lan-entreprise.fr`.
2. Déposez le certificat + la chaîne dans `/etc/nginx/ssl/` (600, `root:root`).

```nginx
server {
    listen 443 ssl;
    server_name netbox.lan-entreprise.fr;

    ssl_certificate     /etc/nginx/ssl/netbox.lan-entreprise.fr.crt;
    ssl_certificate_key /etc/nginx/ssl/netbox.lan-entreprise.fr.key;
    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_prefer_server_ciphers on;
    ssl_session_cache shared:SSL:10m;

    # En-têtes de sécurité
    add_header Strict-Transport-Security "max-age=31536000; includeSubDomains" always;
    add_header X-Content-Type-Options nosniff always;
    add_header X-Frame-Options SAMEORIGIN always;

    # ... reprendre les blocs location /static/ et / de la section 15 ...
}

server {
    listen 80;
    server_name netbox.lan-entreprise.fr;
    return 301 https://$host$request_uri;
}
```

Et dans `configuration.py`, quand HTTPS est actif :

```python
SECURE_SSL_REDIRECT = False   # c'est nginx qui redirige, pas Django
SESSION_COOKIE_SECURE = True
CSRF_COOKIE_SECURE = True
```

> ⚠️ Après activation HTTPS, videz les cookies/session de votre navigateur si la
> connexion tourne en boucle : un cookie `Secure` posé en HTTP est ignoré.

---

## 17. Premier lancement : la check-list des 15 premières minutes

Connectez-vous à `https://netbox.lan-entreprise.fr` avec le superutilisateur.

