---
id: collect-261001-rattrapage/rattrapage/zabbix-guide-4
title: "Guide Zabbix complet — Supervision d'infrastructure en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/zabbix_guide.md
source_anchor: ""
source_lines: [416, 682]
sha256: 78cd7faa4f46e54302e37fd65133dbe39407613e1ee9f1071b5b959c5cf2b310
---

# Guide Zabbix complet — Supervision d'infrastructure en production

```bash
apt install -y zabbix-server-pgsql zabbix-sql-scripts
# Import du schéma (long : plusieurs minutes sur petite VM)
zcat /usr/share/zabbix-sql-scripts/postgresql/server.sql.gz \
  | sudo -u zabbix psql zabbix
# Avec TimescaleDB : activez la compression des hypertables
zcat /usr/share/zabbix-sql-scripts/postgresql/timescaledb/schema.sql.gz \
  | sudo -u zabbix psql zabbix
```

⚠️ **Zabbix 7.x** : le script `timescaledb/schema.sql` crée les hypertables et les policies de compression. Ne l'oubliez pas, sinon vous perdez le bénéfice TimescaleDB.

### 14.5. Configurer `zabbix_server.conf`

```bash
cp /etc/zabbix/zabbix_server.conf /etc/zabbix/zabbix_server.conf.orig
```

Éditez `/etc/zabbix/zabbix_server.conf` (extrait des paramètres essentiels) :

```ini
# --- Base de données ---
DBHost=localhost
DBName=zabbix
DBUser=zabbix
DBPassword=mot_de_passe_ici
DBPort=5432

# --- Caches (à ajuster selon RAM, voir section 61) ---
CacheSize=256M
HistoryCacheSize=128M
HistoryIndexCacheSize=64M
TrendCacheSize=128M
ValueCacheSize=256M

# --- Pollers (point de départ raisonnable) ---
StartPollers=20
StartPollersUnreachable=5
StartTrappers=10
StartPreprocessors=10
StartHistorySyncers=8
StartEscalators=5
StartAlerters=5
StartDiscoverers=3

# --- Housekeeping ---
HousekeepingFrequency=1
MaxHousekeeperDelete=5000

# --- Divers ---
LogFile=/var/log/zabbix/zabbix_server.log
LogFileSize=50
PidFile=/run/zabbix/zabbix_server.pid
SocketDir=/run/zabbix
Timeout=10
FpingLocation=/usr/sbin/fping
```

### 14.6. Démarrer et activer le service

```bash
systemctl enable --now zabbix-server
systemctl status zabbix-server --no-pager
tail -n 50 /var/log/zabbix/zabbix_server.log
```

Le log doit se terminer par quelque chose comme `Zabbix Server started. Zabbix 7.0.x`. Si ce n'est pas le cas → section 74 (dépannage).

## 15. Installation pas à pas : variante MySQL/MariaDB

Si votre équipe maîtrise mieux MySQL, voici la variante (Debian 12) :

```bash
apt install -y mariadb-server zabbix-server-mysql zabbix-sql-scripts
systemctl enable --now mariadb
mysql_secure_installation   # répondez aux questions (root, suppression anonyme...)
```

Créez la base avec le bon charset (**obligatoire** : `utf8mb4`) :

```sql
CREATE DATABASE zabbix CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
CREATE USER 'zabbix'@'localhost' IDENTIFIED BY 'mot_de_passe_ici';
GRANT ALL PRIVILEGES ON zabbix.* TO 'zabbix'@'localhost';
SET GLOBAL log_bin_trust_function_creators = 1;
```

Importez le schéma :

```bash
zcat /usr/share/zabbix-sql-scripts/mysql/server.sql.gz | mysql --default-character-set=utf8mb4 -u zabbix -p zabbix
# Puis remettez la sécurité :
mysql -e "SET GLOBAL log_bin_trust_function_creators = 0;"
```

Dans `/etc/zabbix/zabbix_server.conf` : `DBHost=localhost`, `DBName=zabbix`, `DBUser=zabbix`, `DBPassword=mot_de_passe_ici`.

> ⚠️ Oublier `utf8mb4` provoque des erreurs d'import et des problèmes d'affichage (accents, emojis dans les noms d'hôtes). C'est l'erreur classique n°3 (section 74).

## 16. Installation pas à pas : le frontend web (Nginx + PHP-FPM)

```bash
apt install -y zabbix-frontend-php zabbix-nginx-conf php8.2-fpm php8.2-pgsql \
  php8.2-bcmath php8.2-mbstring php8.2-xml php8.2-gd php8.2-ldap
# (adaptez php8.2-* à la version PHP de votre OS : php8.1-* sur Ubuntu 22.04)
```

⚠️ **Versions PHP** : Zabbix 7.0 exige **PHP 8.0 minimum**. Debian 12 fournit PHP 8.2, Ubuntu 22.04 PHP 8.1, Ubuntu 24.04 PHP 8.3 — tous compatibles 7.0.

Configuration Nginx fournie : `/etc/zabbix/nginx.conf`. Copiez-la ou incluez-la :

```bash
cp /etc/zabbix/nginx.conf /etc/nginx/conf.d/zabbix.conf
# Éditez : server_name, listen 80/443, root /usr/share/zabbix
systemctl enable --now nginx php8.2-fpm
```

Extrait type de `/etc/nginx/conf.d/zabbix.conf` :

```nginx
server {
    listen 80;
    server_name zabbix.votre-domaine.local;
    root /usr/share/zabbix;
    index index.php;

    location = /favicon.ico { log_not_found off; }
    location / {
        try_files $uri $uri/ =404;
    }
    location /assets {
        access_log off;
        expires 10d;
    }
    location ~ /\.ht {
        deny all;
    }
    location ~ /(api\/|conf[^\.]|include|locale) {
        deny all;
        return 404;
    }
    location ~ [^/]\.php(/|$) {
        fastcgi_pass unix:/var/run/php/php8.2-fpm.sock;
        fastcgi_split_path_info ^(.+\.php)(/.+)$;
        fastcgi_index index.php;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        include fastcgi_params;
    }
}
```

> 💡 **Passez en HTTPS** dès que possible (section 68) : les identifiants transitent sinon en clair.

## 17. Installation pas à pas : Zabbix Agent 2 sur Linux

```bash
# Dépôt déjà ajouté (section 14.1)
apt install -y zabbix-agent2 zabbix-agent2-plugin-*
```

`/etc/zabbix/zabbix_agent2.conf` — mode **actif** (recommandé derrière NAT/pare-feu) :

```ini
PidFile=/run/zabbix/zabbix_agent2.pid
LogFile=/var/log/zabbix/zabbix_agent2.log
LogFileSize=20

# Le server/proxy à qui envoyer les données actives
ServerActive=192.168.10.50:10051
# Autoriser aussi les checks passifs depuis le server (optionnel)
Server=192.168.10.50
# DOIT correspondre exactement au "Host name" dans le frontend
Hostname=srv-fichiers-01

Include=/etc/zabbix/zabbix_agent2.d/*.conf
```

```bash
systemctl enable --now zabbix-agent2
systemctl status zabbix-agent2 --no-pager
```

Vérification depuis le server :

```bash
# Test d'un check passif
zabbix_get -s 192.168.10.61 -k agent.hostname
zabbix_get -s 192.168.10.61 -k system.cpu.num
```

## 18. Installation pas à pas : Zabbix Agent sur Windows

1. Téléchargez l'archive `zabbix_agent2-7.0.x-windows-amd64-openssl.zip` depuis zabbix.com/download.
2. Décompressez dans `C:\Program Files\Zabbix Agent 2\`.
3. Éditez `conf\zabbix_agent2.conf` :

```ini
ServerActive=192.168.10.50:10051
Server=192.168.10.50
Hostname=SRV-WIN-01
```

> ⚠️ Le `Hostname` doit correspondre **exactement** (casse comprise) au nom d'hôte déclaré dans le frontend, sinon les données actives sont rejetées (`"host not found"` dans le log).

4. Installez le service (PowerShell en administrateur) :

```powershell
cd "C:\Program Files\Zabbix Agent 2"
.\zabbix_agent2.exe --config conf\zabbix_agent2.conf --install
Start-Service "Zabbix Agent 2"
Get-Service "Zabbix Agent 2"
```

5. Ouvrez le pare-feu Windows pour le 10050/TCP entrant (si checks passifs) :

```powershell
New-NetFirewallRule -DisplayName "Zabbix Agent" -Direction Inbound `
  -Protocol TCP -LocalPort 10050 -Action Allow
```

## 19. Installation pas à pas : Zabbix Proxy

Sur la machine du site distant (Debian/Ubuntu, dépôt ajouté) :

```bash
# Proxy actif + SQLite (simple) ou PostgreSQL (recommandé si > 2000 items)
apt install -y zabbix-proxy-sqlite3
# ou : apt install -y zabbix-proxy-pgsql
```

`/etc/zabbix/zabbix_proxy.conf` (proxy **actif**, SQLite) :

```ini
ProxyMode=0
Server=192.168.10.50
Hostname=zabbix-proxy-sud
LogFile=/var/log/zabbix/zabbix_proxy.log
PidFile=/run/zabbix/zabbix_proxy.pid
DBName=/var/lib/zabbix/zabbix_proxy.db
DBUser=zabbix

# Buffer offline : encaisse 24 h de coupure de liaison
ProxyOfflineBuffer=24h
HeartbeatFrequency=60
ConfigFrequency=300
DataSenderFrequency=5

# Dimensionnement collecte locale
StartPollers=15
StartTrappers=5
CacheSize=128M
Timeout=10
```

```bash
mkdir -p /var/lib/zabbix && chown zabbix:zabbix /var/lib/zabbix
systemctl enable --now zabbix-proxy
```

Puis **créez le proxy dans le frontend** : *Administration → Proxies → Create proxy*, nom **exact** `zabbix-proxy-sud`, mode *Active*. Sans cette déclaration côté server, le proxy reste en `no data`.

> 💡 `ProxyOfflineBuffer=24h` : avec une liaison 4G capricieuse, c'est la différence entre "trou dans les graphes" et "données complètes au retour du lien".

## 20. Sécuriser la communication : PSK et certificats TLS

Par défaut, les échanges server ↔ agent/proxy sont **en clair**. En production : chiffrez.

