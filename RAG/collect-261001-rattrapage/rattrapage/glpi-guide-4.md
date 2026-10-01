---
id: collect-261001-rattrapage/rattrapage/glpi-guide-4
title: "Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "memory"]
source: docs/RAG/collect-261001-rattrapage/glpi_guide.md
source_anchor: ""
source_lines: [278, 514]
sha256: 128d27abb7b55651e93ce37b8dde33ec87f0bf31911fe513f3e0f0fb30d90fbb
---

# Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV

```bash
sudo mysql -u root -p <<'EOF'
CREATE DATABASE glpidb CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER 'glpi'@'localhost' IDENTIFIED BY 'REMPLACER_PAR_UN_MOT_DE_PASSE_FORT';
GRANT ALL PRIVILEGES ON glpidb.* TO 'glpi'@'localhost';
FLUSH PRIVILEGES;
EOF
```

> **Ne jamais** utiliser l'utilisateur `root` de MariaDB dans GLPI.
> Générez le mot de passe avec `openssl rand -base64 24` et conservez-le dans
> un coffre (KeePass, Vault) — jamais en clair dans un ticket ou un courriel.

### Étape 5 — téléchargement et déploiement de GLPI

```bash
cd /tmp
GLPI_VERSION="10.0.18"   # adapter à la dernière 10.0.x disponible
wget "https://github.com/glpi-project/glpi/releases/download/${GLPI_VERSION}/glpi-${GLPI_VERSION}.tgz"
tar xzf "glpi-${GLPI_VERSION}.tgz"
sudo mv glpi /var/www/
sudo chown -R www-data:www-data /var/www/glpi
sudo find /var/www/glpi -type d -exec chmod 755 {} \;
sudo find /var/www/glpi -type f -exec chmod 644 {} \;
```

### Étape 6 — virtual host Apache

Créer `/etc/apache2/sites-available/glpi.conf` (contenu complet en annexe B),
puis :

```bash
sudo a2enmod rewrite headers ssl
sudo a2ensite glpi
sudo a2dissite 000-default
sudo apache2ctl configtest
sudo systemctl reload apache2
```

### Étape 7 — paramètres PHP (fichier `/etc/php/8.2/apache2/php.ini` ou `8.3`)

```ini
memory_limit = 256M
max_execution_time = 120
max_input_vars = 5000
upload_max_filesize = 32M
post_max_size = 40M
session.cookie_httponly = 1
date.timezone = Europe/Paris   ; adapter
opcache.enable = 1
```

Puis `sudo systemctl restart apache2`.

### Étape 8 — lancer l'assistant web

Ouvrez `https://glpi.entreprise.lan/` (ou `http://<ip>/glpi`) et suivez la
[section 7](#7-assistant-dinstallation-web--écran-par-écran).

---

## 6. Installation pas à pas — variante nginx + PHP-FPM

Si vous préférez nginx (performances, reverse proxy existant) :

```bash
sudo apt install -y nginx php-fpm php-mysql php-curl php-gd php-intl \
  php-ldap php-apcu php-xml php-mbstring php-zip php-bz2 mariadb-server
```

### Pool PHP-FPM dédié : `/etc/php/8.2/fpm/pool.d/glpi.conf`

```ini
[glpi]
user = www-data
group = www-data
listen = /run/php/glpi.sock
listen.owner = www-data
listen.group = www-data
pm = dynamic
pm.max_children = 20
pm.start_servers = 4
pm.min_spare_servers = 2
pm.max_spare_servers = 6
php_admin_value[memory_limit] = 256M
php_admin_value[upload_max_filesize] = 32M
php_admin_value[post_max_size] = 40M
```

### Serveur nginx : `/etc/nginx/sites-available/glpi` (extrait, complet en annexe C)

```nginx
server {
    listen 443 ssl;
    server_name glpi.entreprise.lan;
    root /var/www/glpi/public;

    ssl_certificate     /etc/ssl/certs/glpi.crt;
    ssl_certificate_key /etc/ssl/private/glpi.key;

    location / {
        try_files $uri /index.php$is_args$args;
    }
    location ~ ^/index\.php(/|$) {
        fastcgi_pass unix:/run/php/glpi.sock;
        include fastcgi_params;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
    }
    # Bloquer l'accès direct aux dossiers sensibles
    location ~ ^/(config|files|inc|install|tests)/ { deny all; }
}
```

> **Note GLPI 10 :** le DocumentRoot recommandé est `/var/www/glpi/public`
> (le routeur frontal `public/index.php`). Ne pointez plus la racine sur
> `/var/www/glpi` directement comme en 9.x.

```bash
sudo ln -s /etc/nginx/sites-available/glpi /etc/nginx/sites-enabled/
sudo nginx -t && sudo systemctl reload nginx
sudo systemctl restart php8.2-fpm
```

---

## 7. Assistant d'installation web : écran par écran

1. **Sélection de la langue** → Français.
2. **Licence** → accepter (GPL v3+).
3. **Installation / Mise à jour** → « Installation ».
4. **Vérification de l'environnement** : tous les voyants doivent être verts.
   - Orange « recommandé » (ex. `exif`, `sodium`) : acceptable mais à corriger.
   - Rouge « obligatoire » : bloquant — installez l'extension manquante puis rechargez.
5. **Connexion à la base** : serveur `localhost`, utilisateur `glpi`, mot de passe
   choisi à l'étape 4, base `glpidb` (l'assistant peut la créer si l'utilisateur a
   les droits ; ici elle existe déjà).
6. **Initialisation** : création des tables (quelques minutes).
7. **Collecte de données (télémétrie)** : répondez selon votre politique interne.
   En entreprise, on choisit généralement « Ne pas envoyer ».
8. **Comptes par défaut créés** — ⚠️ **à sécuriser immédiatement** (section 8) :
   - `glpi` / `glpi` (super-admin)
   - `tech` / `tech` (technicien)
   - `normal` / `normal` (post-only)
   - `post-only` / `postonly` (déclarant)

### Checklist immédiate après l'assistant

- [ ] Changer les 4 mots de passe par défaut (ou désactiver les comptes inutiles).
- [ ] Supprimer le dossier `/var/www/glpi/install` (ou le renommer `install.bak`).
- [ ] Vérifier `Administration > Journaux` : aucune erreur d'installation.
- [ ] Configurer le cron système (section 9) **avant** toute utilisation réelle.

---

## 8. Configuration post-installation : vérifications critiques

### 8.1 Supprimer / sécuriser l'installeur

```bash
sudo mv /var/www/glpi/install /var/www/glpi/install.bak_$(date +%F)
sudo chown -R www-data:www-data /var/www/glpi
```

### 8.2 Comptes par défaut

`Administration > Utilisateurs` : pour chaque compte par défaut,
définissez un mot de passe fort (20+ caractères) ou désactivez-le
(« Actif : Non »). Ne gardez actif que ce dont vous avez besoin.

### 8.3 Identité de l'application

`Configuration > Générale` :

| Paramètre | Valeur conseillée |
|---|---|
| Nom | `GLPI — SAV Entreprise` |
| URL de l'application | `https://glpi.entreprise.lan` (utilisée dans les courriels !) |
| Fuseau horaire | celui du site principal |
| Langue par défaut | Français |

> **Piège classique :** une URL d'application fausse = des liens morts dans tous
> les courriels de notification. À vérifier en premier quand « les liens ne marchent pas ».

### 8.4 Dossiers de données hors racine web (recommandé)

```bash
sudo mkdir -p /var/lib/glpi/{config,files}
sudo cp -a /var/www/glpi/config/* /var/lib/glpi/config/
sudo cp -a /var/www/glpi/files/* /var/lib/glpi/files/
sudo chown -R www-data:www-data /var/lib/glpi
```

Dans le virtual host Apache (ou `.htaccess`), définir :

```apache
SetEnv GLPI_CONFIG_DIR /var/lib/glpi/config
SetEnv GLPI_VAR_DIR /var/lib/glpi/files
```

(nginx : `fastcgi_param GLPI_CONFIG_DIR /var/lib/glpi/config;` etc.)

### 8.5 Test d'envoi de courriel

`Configuration > Notifications > Configuration des notifications` → renseigner le
SMTP, puis « Envoyer un courriel de test ». Si rien n'arrive : section 56.

---

## 9. Tâches planifiées (cron) : le cœur battant de GLPI

Sans cron, GLPI n'envoie **aucun** courriel, ne calcule **aucun** SLA, ne purge rien
et ne relève pas les boîtes mail. C'est la cause n°1 des « GLPI ne fait rien tout seul ».

### 9.1 Cron système (mode CLI, recommandé)

```bash
# /etc/cron.d/glpi — toutes les minutes, en www-data
* * * * * www-data /usr/bin/php /var/www/glpi/front/cron.php
```

Vérification :

```bash
sudo -u www-data php /var/www/glpi/front/cron.php
# sortie attendue : exécution silencieuse, code 0
```

### 9.2 Mode pseudo-cron (si pas d'accès au cron système)

`Configuration > Actions automatiques` → « Exécuter en mode CLI » : Non.
GLPI exécute alors les tâches lors des visites de pages. **Déconseillé en production**
(irrégulier, dépend du trafic).

### 9.3 Actions automatiques essentielles à vérifier

`Configuration > Actions automatiques` :

