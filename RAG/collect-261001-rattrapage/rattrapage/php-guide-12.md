---
id: collect-261001-rattrapage/rattrapage/php-guide-12
title: "PHP 8 — Le guide complet du sysadmin qui héberge"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "memory"]
source: docs/RAG/collect-261001-rattrapage/php_guide.md
source_anchor: ""
source_lines: [2498, 2714]
sha256: 31a34d6e9e5a311a086841d4df6b6f22c4b7685700535dd93e354f706201f032
---

# PHP 8 — Le guide complet du sysadmin qui héberge

```php
<?php // deploy-reset-opcache.php — protégé, usage ponctuel
// opcache_reset() ne vide que le cache du... : attention, en FPM chaque pool a son cache.
// La méthode fiable : reload gracieux de PHP-FPM après déploiement :
// sudo systemctl reload php8.4-fpm
```

**Procédure de déploiement sûre :**

```bash
# 1. Déployer le code
# 2. Vider l'opcache de chaque pool :
sudo systemctl reload php8.4-fpm   # recharge workers + vide OPcache
# 3. Vérifier :
curl -s https://appli.entreprise.fr/ | head -5
```

**Monitoring OPcache** (via `opcache_get_status()` dans un script admin protégé, ou `php -r`) :

```php
$st = opcache_get_status();
echo "Hit rate : " . round($st['opcache_statistics']['opcache_hit_rate'], 2) . " %\n";
echo "Mémoire utilisée : " . round($st['memory_usage']['used_memory'] / 1024 / 1024) . " Mo\n";
echo "Clés : {$st['opcache_statistics']['num_cached_keys']} / {$st['opcache_statistics']['max_cached_keys']}\n";
// Hit rate < 95 % ou clés saturées → augmenter memory_consumption / max_accelerated_files
```

---

## 59. Réglages performance FPM — le calcul

Formule de base pour `pm.max_children` :

```
max_children = (RAM_totale - RAM_OS - RAM_BDD - marge) / RAM_par_worker_PHP
```

Exemple concret — serveur 8 Go, MariaDB 2 Go, OS 1 Go, marge 1 Go → 4 Go pour PHP ; worker moyen 60 Mo → `max_children ≈ 65` répartis sur les pools.

```bash
# Mesurer la mémoire réelle d'un worker (en Mo)
ps -o rss= -C php-fpm8.4 | awk '{s+=$1; n++} END {printf "%.0f Mo/worker (%d workers)\n", s/n/1024, n}'

# Voir les workers par pool
ps aux | grep "php-fpm: pool" | awk '{print $11, $12, $6/1024 " Mo"}' | head -20
```

Autres leviers :

| Levier | Effet |
|---|---|
| `pm.max_requests = 500` | Recycle les workers → contient les fuites mémoire lentes |
| OPcache `validate_timestamps=0` | -5 à -15 % de latence |
| Socket Unix vs TCP | Socket Unix = plus rapide en local (pas de stack TCP) |
| `realpath_cache_size = 4096k` + `realpath_cache_ttl = 600` (php.ini) | Moins de `stat()` disque |
| nginx `fastcgi_cache` | Cache pleine page pour le contenu public (sections non connectées) |
| Redis : sessions + cache appli | Décharge la BDD |

---

## 60. Supervision : les logs — où, quoi, comment

| Log | Chemin | Contenu |
|---|---|---|
| Erreurs PHP (pool) | `/var/log/php-fpm-appli-error.log` | Fatal, warnings, `error_log()` |
| Erreurs FPM globales | `/var/log/php8.4-fpm.log` | Démarrage pools, workers tués |
| Slow log | `/var/log/php-fpm-appli-slow.log` | Requêtes > N secondes (voir ci-dessous) |
| Accès nginx | `/var/log/nginx/appli-access.log` | Chaque requête HTTP |
| Erreurs nginx | `/var/log/nginx/appli-error.log` | Timeouts FPM, 502... |

```ini
; pool.d/appli.conf — slow log : détecte les requêtes qui rament
request_slowlog_timeout = 5s
slowlog = /var/log/php-fpm-appli-slow.log
; Toute requête > 5 s = stack trace PHP dans le slow log → identifie la fonction fautive
```

```bash
# Surveillance temps réel
tail -f /var/log/php-fpm-appli-error.log

# Top des erreurs du jour
grep "$(date +%Y-%m-%d)" /var/log/php-fpm-appli-error.log | cut -d']' -f2- | sort | uniq -c | sort -rn | head

# Corréler nginx 502 ↔ FPM : même fenêtre temporelle
grep " 502 " /var/log/nginx/appli-access.log | tail -5
grep -E "WARNING|ERROR" /var/log/php8.4-fpm.log | tail -5

# Rotation : logrotate est préconfiguré pour php-fpm sur Debian — vérifier :
cat /etc/logrotate.d/php8.4-fpm
```

🔒 **Logs = données sensibles potentielles** (e-mails, traces SQL) : droits `640`, propriétaire `appli:adm`, jamais exposés par le web, rotation active, rétention 30-90 jours selon besoin.

---

## 61. Supervision : status FPM, ping, métriques

```ini
; pool.d/appli.conf (déjà vu section 57)
pm.status_path = /fpm-status-appli
ping.path = /fpm-ping-appli
```

```nginx
# nginx : exposer le status UNIQUEMENT en local (jamais au public !)
location ~ ^/(fpm-status-appli|fpm-ping-appli)$ {
    allow 127.0.0.1;
    allow 10.0.0.0/8;      # ton réseau de supervision
    deny all;
    include fastcgi_params;
    fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
    fastcgi_pass unix:/run/php/php8.4-fpm-appli.sock;
}
```

```bash
# Ping : le pool répond ?
curl -s http://127.0.0.1/fpm-ping-appli
# pong

# Status : état des workers (format texte ou ?json)
curl -s "http://127.0.0.1/fpm-status-appli?json" | python3 -m json.tool
```

Exemple de sortie (`?json`) :

```json
{
    "pool": "appli",
    "process manager": "dynamic",
    "start time": 1727222400,
    "accepted conn": 15230,
    "listen queue": 0,
    "max listen queue": 4,
    "idle processes": 3,
    "active processes": 2,
    "total processes": 5,
    "max active processes": 18,
    "max children reached": 0
}
```

**Alertes à poser (Zabbix/Prometheus/node_exporter ou script cron) :**

| Métrique | Seuil d'alerte |
|---|---|
| `max children reached` = 1 | ⚠️ pool saturé → augmenter `max_children` ou chercher la lenteur |
| `listen queue` > 0 durable | requêtes en attente → sous-dimensionné |
| `slow requests` croissant | code ou BDD qui rame |
| ping ≠ `pong` | pool HS → redémarrer + alerter |

Script de check minimal (cron toutes les 5 min) :

```bash
#!/bin/bash
# /usr/local/bin/check-fpm.sh
for pool in appli intranet; do
  rep=$(curl -sm 5 "http://127.0.0.1/fpm-ping-$pool")
  [ "$rep" = "pong" ] || echo "ALERTE : pool FPM $pool ne répond pas" | mail -s "FPM $pool HS" astreinte@entreprise.fr
done
```

---

## 62. Durcissement : open_basedir et disable_functions

Deux remparts **par pool** (section 57), non contournables par l'appli (`php_admin_value`) :

```ini
; --- open_basedir : PHP ne peut ouvrir des fichiers QUE dans ces dossiers ---
php_admin_value[open_basedir] = /var/www/appli:/tmp:/usr/share/php
; Effet : file_get_contents('/etc/passwd') → Warning "open_basedir restriction in effect"
; ⚠️ Inclure tous les chemins légitimes (vendor, sessions, tmp, uploads) sinon l'appli casse.
; Tester après activation : parcourir les fonctions critiques de l'appli.

; --- disable_functions : désactive les exécutions shell ---
php_admin_value[disable_functions] = exec,passthru,shell_exec,system,proc_open,popen,pcntl_exec,show_source,phpinfo
; ⚠️ Certaines applis en ont besoin (ex. : génération PDF via binaire, backups).
; Alternative : les autoriser mais auditer chaque appel dans le code.
```

**Test d'efficacité** (après activation, en CLI avec la conf du pool ou via une page de test **temporaire et protégée**) :

```php
<?php
var_dump(@file_get_contents('/etc/passwd')); // doit échouer (open_basedir)
var_dump(function_exists('exec') && @exec('id')); // exec désactivée
```

🔒 Ces deux réglages ne remplacent **pas** la correction du code (injections, uploads) : ce sont des **filets**, pas des murs. Un attaquant qui uploade un shell reste bloqué hors de `/etc`, mais le shell tourne quand même dans le pool.


---

## 63. Sauvegarde — code + base de données

On ne sauvegarde pas « le serveur », on sauvegarde de quoi **reconstruire** : code versionné + dump BDD + fichiers utilisateurs + configs.

```bash
#!/bin/bash
# /usr/local/bin/backup-appli.sh — à lancer en cron (utilisateur dédié "backup")
set -euo pipefail
DATE=$(date +%Y%m%d-%H%M)
DEST=/srv/backups/appli
mkdir -p "$DEST"

# 1. Base de données (MariaDB) — transactionnel, sans verrou global
mysqldump --single-transaction --quick --routines --triggers \
  -u backup -p"$DB_BACKUP_PASS" appli | gzip > "$DEST/appli-bdd-$DATE.sql.gz"

# 2. Fichiers utilisateurs (uploads) — hors git par nature
tar -czf "$DEST/appli-uploads-$DATE.tar.gz" -C /var/www/appli uploads/

# 3. Code : le git fait foi — on archive le commit déployé (traçabilité)
git -C /var/www/appli rev-parse HEAD > "$DEST/appli-commit-$DATE.txt"

