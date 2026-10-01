---
id: collect-261001-rattrapage/rattrapage/php-guide-1
title: "PHP 8 — Le guide complet du sysadmin qui héberge"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "memory"]
source: docs/RAG/collect-261001-rattrapage/php_guide.md
source_anchor: ""
source_lines: [1, 206]
sha256: 7c8224c747ba7b4d5b6cd0da5ef3cd88aeaa91c8ede1bc9f9d2e7efeddebf2c6
---

# PHP 8 — Le guide complet du sysadmin qui héberge

> **Public :** sysadmin / chef de service systèmes & énergies qui déploie, sécurise et maintient des applications PHP en production.
> **Version couverte :** PHP 8.1 → 8.4 sur Debian 12 / Ubuntu 22.04 / 24.04.
> **Objectif :** installer, configurer, sécuriser, superviser et dépanner PHP + PHP-FPM + nginx comme un pro, et écrire du PHP 8 propre.

---

## 1. Pourquoi ce guide

PHP fait tourner ~75 % du web (WordPress, GLPI, Nextcloud, Dolibarr, applications métier…). En tant que sysadmin, tu n'as pas forcément besoin d'écrire des frameworks : tu dois **installer proprement**, **configurer pour la production**, **sécuriser**, **superviser** et **dépanner** — et comprendre le code que tu héberges pour juger de sa qualité.

Ce guide couvre les deux versants :

| Versant | Contenu |
|---|---|
| **Hébergement** | Installation, `php.ini` prod, PHP-FPM + nginx, OPcache, logs, supervision, sauvegardes, mises à jour |
| **Langage** | Syntaxe PHP 8, types, POO, PDO, sécurité du code, formulaires, sessions, uploads |

**Conventions utilisées :**

- Les blocs `bash` = commandes à exécuter sur le serveur.
- Les blocs `php` = code PHP 8 valide (testé mentalement, pas de placeholder).
- ⚠️ = point de vigilance production. 🔒 = sécurité. 📋 = checklist.

---

## 2. Prérequis et environnement de travail

- Un serveur Debian 12 (Bookworm) ou Ubuntu 22.04/24.04 LTS, à jour :
  ```bash
  sudo apt update && sudo apt upgrade -y
  ```
- Accès root / sudo.
- nginx (ou Apache) — ce guide utilise **nginx + PHP-FPM** (recommandé en prod).
- MariaDB/MySQL ou PostgreSQL si l'appli utilise une base (sections PDO).
- Un éditeur : `nano`, `vim`, ou VS Code en SSH distant.

📋 **Checklist avant de commencer :**

- [ ] OS à jour (`apt upgrade`)
- [ ] Nom d'hôte et DNS configurés
- [ ] Pare-feu actif (UFW : 80/443 ouverts, 22 restreint)
- [ ] Sauvegarde existante avant toute modification
- [ ] Environnement de test/staging séparé de la prod si possible

---

## 3. Installation de PHP 8.x sur Debian/Ubuntu

Debian 12 embarque PHP 8.2, Ubuntu 24.04 embarque PHP 8.3. Pour les versions plus récentes (8.3/8.4 sur Debian 12, 8.4 sur Ubuntu 22.04), utilise le dépôt **Sury** (Ondřej Surý, mainteneur officiel des paquets PHP Debian) :

```bash
# 1. Dépendances pour ajouter un dépôt APT
sudo apt install -y lsb-release ca-certificates apt-transport-https software-properties-common gnupg

# 2. Clé GPG du dépôt Sury
sudo curl -fsSL https://packages.sury.org/php/apt.gpg -o /usr/share/keyrings/deb.sury.org-php.gpg

# 3. Ajout du dépôt
echo "deb [signed-by=/usr/share/keyrings/deb.sury.org-php.gpg] https://packages.sury.org/php/ $(lsb_release -sc) main" \
  | sudo tee /etc/apt/sources.list.d/php.list

# 4. Installation de PHP 8.4 + FPM + extensions courantes
sudo apt update
sudo apt install -y php8.4 php8.4-fpm php8.4-cli \
  php8.4-mysql php8.4-pgsql php8.4-sqlite3 \
  php8.4-curl php8.4-xml php8.4-mbstring php8.4-zip \
  php8.4-gd php8.4-intl php8.4-bcmath php8.4-gmp \
  php8.4-redis php8.4-opcache php8.4-readline
```

🔒 **N'installe que les extensions nécessaires.** Chaque extension = surface d'attaque et mémoire en plus. Adapte la liste à l'appli hébergée (ex. : pas de `php8.4-pgsql` si tu n'as que MySQL).

Vérification :

```bash
php -v            # version CLI
php -m            # modules chargés
php --ini         # fichiers ini lus (CLI)
```

> **Bon à savoir :** PHP existe en deux SAPI (interfaces) : **CLI** (ligne de commande, `php script.php`) et **FPM** (FastCGI, pour le web via nginx). Chacune a **son propre `php.ini`**. Ne confonds jamais les deux.

| SAPI | Binaire / service | Fichier ini (Debian) |
|---|---|---|
| CLI | `php` | `/etc/php/8.4/cli/php.ini` |
| FPM | `php8.4-fpm` (service) | `/etc/php/8.4/fpm/php.ini` |

---

## 4. Installation et premier démarrage de PHP-FPM

PHP-FPM (FastCGI Process Manager) est le SAPI recommandé pour servir PHP via nginx : pools de workers, gestion fine des processus, status temps réel.

```bash
# Le paquet php8.4-fpm installe et démarre le service
sudo systemctl enable --now php8.4-fpm
sudo systemctl status php8.4-fpm

# Socket Unix par défaut (à utiliser dans nginx) :
ls -l /run/php/php8.4-fpm.sock
```

Test de fonctionnement avec nginx (config complète section 56). Test rapide en CLI :

```bash
# Page d'info PHP en CLI (équivalent de phpinfo() web, sans exposer le web)
php -r 'phpinfo();' | head -40
```

⚠️ **Ne mets JAMAIS `phpinfo()` sur un site en production** : ça expose versions, chemins, variables d'environnement. En dev local uniquement.

---

## 5. Configuration php.ini pour la production

Deux fichiers à régler : `/etc/php/8.4/fpm/php.ini` (web) et `/etc/php/8.4/cli/php.ini` (CLI/cron). Après modification : `sudo systemctl reload php8.4-fpm`.

### 5.1 Réglages de sécurité (FPM — les plus importants)

```ini
; === EXPOSITION ===
expose_php = Off              ; ne pas envoyer "X-Powered-By: PHP/8.4" dans les en-têtes

; === ERREURS : jamais affichées au visiteur en prod ===
display_errors = Off
display_startup_errors = Off
log_errors = On
error_log = /var/log/php8.4-fpm-error.log
error_reporting = E_ALL       ; on loggue TOUT, on affiche RIEN

; === UPLOADS ===
file_uploads = On
upload_max_filesize = 10M     ; selon besoin métier (jamais plus que nécessaire)
max_file_uploads = 5

; === LIMITES RESSOURCES ===
max_execution_time = 30       ; 30 s max par requête web (CLI : 0 = illimité)
max_input_time = 30
memory_limit = 256M           ; par worker ; x N workers = RAM totale à prévoir
post_max_size = 12M           ; > upload_max_filesize (marge pour les champs du form)
max_input_vars = 2000         ; protection contre les attaques par saturation

; === SESSIONS (voir section 40) ===
session.cookie_httponly = 1
session.cookie_secure = 1     ; HTTPS uniquement (si le site est en HTTPS, ce qui est obligatoire)
session.cookie_samesite = "Lax"
session.use_strict_mode = 1
session.use_only_cookies = 1

; === DIVERS ===
allow_url_fopen = Off         ; bloque file_get_contents("http://...") sauf besoin explicite
allow_url_include = Off        ; TOUJOURS Off (inclusion de code distant = faille critique)
default_charset = "UTF-8"
date.timezone = "Europe/Paris" ; ou la TZ du serveur ; évite les warnings DateTime
```

### 5.2 Différences CLI

```ini
; /etc/php/8.4/cli/php.ini
max_execution_time = 0   ; pas de limite pour les scripts longs (imports, crons)
memory_limit = 512M      ; les batchs consomment plus
log_errors = On
```

### 5.3 Vérifier la config effective

```bash
php -i | grep -E "^(expose_php|display_errors|memory_limit|upload_max_filesize)"
# ou, ciblé :
php -r 'echo ini_get("memory_limit"), PHP_EOL;'
```

📋 **Checklist php.ini prod :** `expose_php=Off`, `display_errors=Off`, `log_errors=On`, `allow_url_include=Off`, cookies de session durcis, `memory_limit` cohérent avec la RAM, `upload_max_filesize` au besoin strict.

---

## 6. PHP en ligne de commande (CLI)

Le CLI sert pour les crons, les scripts d'import, Composer, les tests.

```bash
php -v                          # version
php -l script.php                # vérifie la SYNTAXE sans exécuter (lint) — indispensable avant déploiement
php script.php arg1 arg2         # exécution avec arguments ($argv)
php -r 'echo "hello\n";'         # one-liner
php -a                           # shell interactif (pratique pour tester une fonction)
```

Récupérer les arguments :

```php
<?php
// script.php appelé via : php script.php import --force
var_dump($argv);   // [0 => 'script.php', 1 => 'import', 2 => '--force']
var_dump($argc);   // 3
```

🔒 **Permissions :** un script PHP CLI lancé par cron doit avoir les droits minimaux (propriétaire dédié, jamais root sauf nécessité absolue). Les scripts web (`/var/www/...`) ne doivent **jamais** être exécutables en écriture par l'utilisateur FPM.

---

## 7. Syntaxe de base

