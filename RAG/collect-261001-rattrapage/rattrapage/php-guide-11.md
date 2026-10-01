---
id: collect-261001-rattrapage/rattrapage/php-guide-11
title: "PHP 8 — Le guide complet du sysadmin qui héberge"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["memory"]
source: docs/RAG/collect-261001-rattrapage/php_guide.md
source_anchor: ""
source_lines: [2296, 2497]
sha256: f220c86b6d94ae3e0976b97772ae982d8d5330c50b2d5e4786f55334b0a9c942
---

# PHP 8 — Le guide complet du sysadmin qui héberge

```nginx
# --- nginx : security headers ---
add_header X-Content-Type-Options "nosniff" always;
add_header X-Frame-Options "SAMEORIGIN" always;          # anti-clickjacking
add_header Referrer-Policy "strict-origin-when-cross-origin" always;
add_header Permissions-Policy "camera=(), microphone=(), geolocation=()" always;
# HSTS : une fois le HTTPS validé et stable
add_header Strict-Transport-Security "max-age=31536000; includeSubDomains" always;

# Content-Security-Policy : la plus puissante (à ajuster selon l'appli)
add_header Content-Security-Policy "default-src 'self'; script-src 'self'; style-src 'self' 'unsafe-inline'; img-src 'self' data:; frame-ancestors 'self'" always;
```

Vérification : `curl -sI https://ton-site.fr | grep -i "^x-\|^content-security\|^strict"`.

Côté PHP, forcer le HTTPS :

```php
<?php declare(strict_types=1);
// En tête du bootstrap : redirige tout le HTTP vers HTTPS
if (($_SERVER['HTTPS'] ?? 'off') !== 'on' && ($_SERVER['HTTP_X_FORWARDED_PROTO'] ?? '') !== 'https') {
    header('Location: https://' . $_SERVER['HTTP_HOST'] . $_SERVER['REQUEST_URI'], true, 301);
    exit;
}
```

---

## 55. Checklist sécurité — l'audit en 15 minutes

📋 **À passer sur chaque appli hébergée :**

**Configuration serveur**
- [ ] `display_errors=Off`, `expose_php=Off`, `allow_url_include=Off`
- [ ] PHP et extensions à jour (`apt upgrade`, `composer audit`)
- [ ] `.env` / `config.php` hors racine web, droits `0600`
- [ ] `vendor/`, `.git/` non accessibles par le web (règle nginx `deny all`)
- [ ] HTTPS partout + HSTS, cookies `Secure`

**Code**
- [ ] 100 % des requêtes SQL en **préparées** (`grep -n "query(\"" src/` → vérifier chacune)
- [ ] 100 % des affichages via `e()` / `htmlspecialchars`
- [ ] Token CSRF sur tous les formulaires POST
- [ ] `password_hash` / `password_verify` (aucun md5/sha1)
- [ ] Uploads : whitelist MIME + nom généré + hors racine web
- [ ] Aucun `include` avec variable non whitelistée
- [ ] `random_int`/`random_bytes` pour les tokens (pas `rand`/`uniqid` seul)

**Exploitation**
- [ ] Logs d'erreurs centralisés et surveillés
- [ ] Sauvegardes code + BDD testées (restauration essayée !)
- [ ] fail2ban sur les 403/404 suspects et les échecs login
- [ ] Comptes admin : 2FA si l'appli le permet, mots de passe forts

---

## 56. nginx + PHP-FPM — configuration complète

```nginx
# /etc/nginx/sites-available/appli.conf
server {
    listen 80;
    server_name appli.entreprise.fr;
    return 301 https://$host$request_uri;  # tout en HTTPS
}

server {
    listen 443 ssl;
    server_name appli.entreprise.fr;

    ssl_certificate     /etc/letsencrypt/live/appli.entreprise.fr/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/appli.entreprise.fr/privkey.pem;
    # (certbot --nginx pour générer)

    root /var/www/appli/public;   # 🔒 SEUL public/ est exposé (index.php dedans)
    index index.php;

    # --- Security headers (section 54) ---
    add_header X-Content-Type-Options "nosniff" always;
    add_header X-Frame-Options "SAMEORIGIN" always;
    add_header Referrer-Policy "strict-origin-when-cross-origin" always;

    # --- Bloquer l'accès aux sensibles ---
    location ~ /\. { deny all; }                        # .git, .env
    location ~* \.(env|log|sql|bak|old)$ { deny all; }

    # --- PHP via FPM (socket Unix) ---
    location / {
        try_files $uri $uri/ /index.php?$query_string;
    }

    location ~ \.php$ {
        try_files $uri =404;                           # 🔒 anti path-info (fichier.php/nimporte-quoi)
        include fastcgi_params;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        fastcgi_param HTTPS on;
        fastcgi_pass unix:/run/php/php8.4-fpm-appli.sock;  # socket DÉDIÉE par site (section 57)
        fastcgi_read_timeout 60s;
    }

    # --- Fichiers statiques : cache agressif ---
    location ~* \.(css|js|png|jpg|webp|woff2)$ {
        expires 30d;
        add_header Cache-Control "public, immutable";
    }

    access_log /var/log/nginx/appli-access.log;
    error_log  /var/log/nginx/appli-error.log;
}
```

```bash
sudo ln -s /etc/nginx/sites-available/appli.conf /etc/nginx/sites-enabled/
sudo nginx -t && sudo systemctl reload nginx
```

🔒 **Points critiques :** `root` sur `public/` uniquement, `try_files $uri =404` dans le bloc PHP (sinon exécution de fichiers uploadés via path-info), socket par site, `.env`/`.git` bloqués.

---

## 57. Pools PHP-FPM par site — isolation

**Un pool = un site = un utilisateur système.** Si un site est compromis, l'attaquant n'accède pas aux fichiers des autres.

```bash
# Créer l'utilisateur dédié (sans shell, sans home inscriptible)
sudo useradd --system --no-create-home --shell /usr/sbin/nologin appli
sudo mkdir -p /var/www/appli
sudo chown -R appli:www-data /var/www/appli
sudo chmod -R 750 /var/www/appli
```

```ini
; /etc/php/8.4/fpm/pool.d/appli.conf
[appli]
user = appli
group = www-data

listen = /run/php/php8.4-fpm-appli.sock
listen.owner = www-data
listen.group = www-data
listen.mode = 0660

; --- Gestion des processus : dynamic = le bon compromis ---
pm = dynamic
pm.max_children = 20        ; plafond de workers (RAM : 20 x 60 Mo ≈ 1,2 Go)
pm.start_servers = 4
pm.min_spare_servers = 2
pm.max_spare_servers = 6
pm.max_requests = 500       ; recycle chaque worker après 500 requêtes (anti-fuite mémoire)

; --- Logs dédiés ---
php_admin_value[error_log] = /var/log/php-fpm-appli-error.log
php_admin_flag[log_errors] = on

; --- Durcissement par pool (php_admin_value = non surchargeable par l'appli) ---
php_admin_value[open_basedir] = /var/www/appli:/tmp
php_admin_value[upload_tmp_dir] = /var/www/appli/tmp
php_admin_value[session.save_path] = /var/www/appli/sessions
php_admin_value[disable_functions] = exec,passthru,shell_exec,system,proc_open,popen,pcntl_exec
; ⚠️ disable_functions : à ajuster — certaines applis (backups, conversions) ont besoin d'exec.

; --- Status / ping (supervision, section 61) ---
pm.status_path = /fpm-status-appli
ping.path = /fpm-ping-appli
ping.response = pong
```

```bash
sudo systemctl reload php8.4-fpm
```

**Dimensionnement `pm.max_children` :** `RAM_disponible_pour_PHP / mémoire_moyenne_par_worker`. Mesure la mémoire réelle : `ps --sort=-rss -o rss,command -C php-fpm8.4 | awk '{sum+=$1} END {print sum/NR/1024 " Mo/worker"}'`.

| Profil | pm | Quand |
|---|---|---|
| `ondemand` | workers créés à la demande | Faible trafic, économie RAM (démarrage à froid) |
| `dynamic` | min/max spare | **Standard** : la plupart des applis |
| `static` | N workers fixes | Fort trafic constant, latence minimale |

---

## 58. OPcache — le turbo gratuit

OPcache compile chaque script PHP **une fois** et garde le bytecode en mémoire partagée. Gain typique : **2 à 5x** sur le temps de réponse. Activé par défaut, mais à régler :

```ini
; /etc/php/8.4/fpm/conf.d/10-opcache.ini (ou dans php.ini)
opcache.enable = 1
opcache.enable_cli = 0              ; CLI : inutile (processus éphémères), 1 pour les longs workers
opcache.memory_consumption = 256    ; Mo de SHM (128-512 selon taille du code)
opcache.interned_strings_buffer = 16
opcache.max_accelerated_files = 20000 ; > nombre de fichiers PHP (compter : find /var/www -name "*.php" | wc -l)
opcache.validate_timestamps = 0     ; 🔒 PROD : 0 = ne jamais revérifier les fichiers (perf max)
opcache.revalidate_freq = 0         ; ignoré si validate_timestamps=0
opcache.save_comments = 1           ; requis par certaines libs (annotations)
opcache.jit = 1255                  ; PHP 8 : JIT tracing, 1255 = recommandé
opcache.jit_buffer_size = 128M
```

⚠️ **`validate_timestamps=0` en prod** = après chaque déploiement, il faut **vider l'OPcache** (sinon l'ancien code reste servi !) :

