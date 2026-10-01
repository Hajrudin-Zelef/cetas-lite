---
id: collect-261001-rattrapage/rattrapage/php-guide-13
title: "PHP 8 — Le guide complet du sysadmin qui héberge"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "attention", "memory"]
source: docs/RAG/collect-261001-rattrapage/php_guide.md
source_anchor: ""
source_lines: [2715, 2877]
sha256: ebddedcba7dc23568c92af7a6b171ac36b1269e9b72ff444d4b28d94c90e74f0
---

# 4. Configs (nginx, pools FPM, php.ini)
tar -czf "$DEST/appli-configs-$DATE.tar.gz" \
  /etc/nginx/sites-available/appli.conf \
  /etc/php/8.4/fpm/pool.d/appli.conf

# 5. Rétention : 7 jours glissants
find "$DEST" -type f -mtime +7 -delete

# 6. Test d'intégrité du dump (corrompu = inutile)
gzip -t "$DEST/appli-bdd-$DATE.sql.gz" && echo "dump OK"
```

```bash
# Cron : tous les jours à 2h
0 2 * * * /usr/local/bin/backup-appli.sh >> /var/log/backup-appli.log 2>&1
```

**Restauration BDD (à TESTER au moins une fois par trimestre sur un serveur de test) :**

```bash
zcat /srv/backups/appli/appli-bdd-20261005-0200.sql.gz | mysql -u root -p appli_restauree
```

📋 **Checklist sauvegarde :** BDD (`--single-transaction`), uploads, configs, commit git tracé, chiffrement si offsite, **test de restauration réel** (un backup non testé = pas de backup), alerte si le cron échoue (log + monitoring du fichier le plus récent).

---

## 64. Dépannage — méthode générale

```
1. REPRODUIRE : quelle URL/action exacte ? quel utilisateur ? quelle heure ?
2. LIRE LES LOGS (dans cet ordre) :
   a. /var/log/nginx/appli-error.log      → 502 ? timeout ? (le symptôme HTTP)
   b. /var/log/php-fpm-appli-error.log    → fatal/warning PHP (la cause code)
   c. /var/log/php8.4-fpm.log             → pool saturé ? workers tués ?
   d. slow log                            → quelle fonction rame ?
   e. logs MariaDB / appli                → cause BDD/métier
3. ISOLER : staging ? un seul pool ? après un déploiement ? (git log)
4. CORRIGER → tester → déployer → SURVEILLER les logs 15 min
```

```bash
# Kit de survie : statut des services
systemctl is-active php8.4-fpm nginx mariadb
# File d'attente FPM (si > 0 durable → saturation)
curl -s "http://127.0.0.1/fpm-status-appli?json" | grep -E "listen queue|active processes|max children"
# Processus qui consomment
ps --sort=-rss -o pid,rss,command -C php-fpm8.4 | head
# Espace disque (1re cause des "ça marchait hier")
df -h / /var /tmp
# Inodes (un disque plein à 30 % mais 100 % d'inodes = même effet)
df -i /var
```

---

## 65. Dépannage — cas concret n°1 : erreur 502 Bad Gateway

**Symptômes :** nginx répond 502 sur les pages PHP, le HTML statique marche.

| Cause | Diagnostic | Remède |
|---|---|---|
| FPM arrêté | `systemctl is-active php8.4-fpm` → `inactive` | `systemctl start php8.4-fpm`, chercher pourquoi (OOM ? `dmesg \| grep -i oom`) |
| Mauvais socket | `connect() to unix:/run/php/...sock failed (2: No such file)` dans nginx error log | Aligner `fastcgi_pass` nginx ↔ `listen` du pool |
| Permissions socket | `(13: Permission denied)` | `listen.owner/group/mode` : www-data doit lire/écrire |
| Pool saturé | `max children reached` dans php8.4-fpm.log | Augmenter `max_children` (si RAM) ou traiter la lenteur (slow log) |
| Fatal PHP au démarrage du worker | Erreur dans php-fpm-appli-error.log | Corriger le code (souvent après déploiement → rollback) |

```bash
# Vérifier que le socket existe et qui peut l'utiliser
ls -l /run/php/php8.4-fpm-appli.sock
# Tester FPM directement (bypass nginx)
SCRIPT_NAME=/fpm-ping-appli SCRIPT_FILENAME=/fpm-ping-appli \
  cgi-fcgi -bind -connect /run/php/php8.4-fpm-appli.sock  # doit répondre "pong"
```

---

## 66. Dépannage — cas concret n°2 : page blanche (WSOD)

**Symptômes :** HTTP 200 mais page vide, ou 500 sans message (prod : `display_errors=Off`, c'est normal).

```bash
# 1. Le log PHP du pool dit tout :
tail -30 /var/log/php-fpm-appli-error.log
# Exemples typiques :
# "PHP Fatal error: Uncaught Error: Class 'App\\Sonde' not found" → autoload (composer dump-autoload)
# "PHP Parse error: syntax error, unexpected '=>' " → code PHP 8 sur... non, ici c'est l'inverse :
#    code écrit pour PHP 8.4 déployé sur 8.1 → vérifier php -v vs code
# "Allowed memory size of 268435456 bytes exhausted" → memory_limit trop bas OU boucle infinie

# 2. Reproduire en CLI (même code, erreurs visibles) :
sudo -u appli php /var/www/appli/public/index.php
# ⚠️ attention : le CLI n'a pas les mêmes ini ni les mêmes variables ($_SERVER vide)

# 3. Vérifier l'OPcache après déploiement :
# "j'ai déployé mais l'ancien code tourne" → systemctl reload php8.4-fpm (section 58)
```

**Checklist WSOD :** logs pool → `php -l` sur le fichier modifié → droits fichiers (`appli:www-data`, pas root:root 600) → OPcache vidé → `.env` présent → extensions requises (`composer check-platform-reqs`).

---

## 67. Dépannage — cas concret n°3 : lenteurs

```bash
# 1. Slow log : qui rame ?
tail -20 /var/log/php-fpm-appli-slow.log
# "script_filename = /var/www/appli/public/index.php ... mysql_query() ..." → la BDD

# 2. C'est la BDD ? Activer le log des requêtes lentes MariaDB :
# /etc/mysql/mariadb.conf.d/50-server.cnf : slow_query_log=1, long_query_time=2
# Puis : mysqldumpslow /var/log/mysql/slow.log | head -20

# 3. C'est le réseau/appel externe ? (API, LDAP...)
# Mesurer avec curl depuis le serveur :
curl -o /dev/null -s -w "connexion: %{time_connect}s, total: %{time_total}s\n" https://api.externe.fr/health

# 4. Workers occupés par des requêtes longues → file d'attente :
curl -s "http://127.0.0.1/fpm-status-appli?json" | python3 -c "import json,sys; d=json.load(sys.stdin); print(d['listen queue'], d['active processes'])"
```

**Remèdes par cause :**

| Cause | Remède |
|---|---|
| Requête SQL lente | Index manquant (`EXPLAIN`), pagination, cache |
| Appel HTTP externe synchrone | Timeout court (`CURLOPT_TIMEOUT=5`), file de jobs (cron/worker) |
| `max_children` trop bas | Augmenter (si RAM dispo) |
| OPcache mal réglé | `validate_timestamps=0`, mémoire suffisante |
| Session verrouillée | `session_write_close()` tôt (section 40) |

---

## 68. Les 12 erreurs classiques — diagnostic éclair

| # | Message | Cause la plus probable | Fix |
|---|---|---|---|
| 1 | `Headers already sent by ... :12` | Espace/BOM avant `<?php`, ou `echo` avant `header()`/`session_start()` | Supprimer tout output avant les en-têtes ; pas de `?>` final |
| 2 | `Undefined array key "x"` | Accès tableau sans test | `$_GET['x'] ?? défaut`, `isset()` |
| 3 | `Call to undefined function mb_strlen()` | Extension `php8.4-mbstring` non installée | `apt install php8.4-mbstring` + reload FPM |
| 4 | `Class 'PDO' not found` | `php8.4-mysql`/`php8.4-pgsql` absent | Installer le driver PDO adéquat |
| 5 | `SQLSTATE[HY000] [2002] Connection refused` | MariaDB arrêté, mauvais host/port, pare-feu | `systemctl status mariadb`, tester `mysql -h ... -u ...` |
| 6 | `Access denied for user 'appli'@'...'` | Mauvais mot de passe ou host non autorisé | Vérifier `.env` + `SELECT user,host FROM mysql.user` |
| 7 | `Permission denied` (fopen, move_uploaded_file) | Droits fichiers / `open_basedir` | `chown appli:www-data`, vérifier `open_basedir` |
| 8 | `Allowed memory size exhausted` | `memory_limit` trop bas ou fuite/boucle | Augmenter avec mesure, ou corriger le code (fetch ligne à ligne) |
| 9 | `Maximum execution time exceeded` | Script trop long ou blocage (BDD, HTTP) | Timeout côté cause ; `set_time_limit()` en CLI si légitime |
| 10 | `Failed to open stream: No such file` (include) | Mauvais chemin relatif | Chemins absolus via `__DIR__`, autoloader PSR-4 |
| 11 | `Typed property must not be accessed before initialization` | Propriété typée non initialisée au constructeur | Initialiser dans le constructeur ou `?type` + `= null` |
| 12 | `Cannot modify readonly property` | Assignation après construction | `readonly` = constructeur uniquement ; revoir le design |

---

## 69. Mise à jour PHP 7.4 → 8.x — la procédure

```bash
# 1. INVENTAIRE : quelles applis, quelle version PHP actuelle, quelles extensions ?
php -v && php -m | sort

# 2. AUDIT de compatibilité (outil officiel : rector en mode dry-run, ou phpstan)
composer require --dev rector/rector
vendor/bin/rector process src --dry-run --config=rector-74-to-84.php

