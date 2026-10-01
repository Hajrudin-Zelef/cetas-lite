---
id: collect-261001-rattrapage/rattrapage/zabbix-guide-12
title: "Guide Zabbix complet — Supervision d'infrastructure en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents"]
source: docs/RAG/collect-261001-rattrapage/zabbix_guide.md
source_anchor: ""
source_lines: [1979, 2166]
sha256: 9bf027967403aa8b616b0fc19f9663d43b3dd98dfe08f8d47052351d188306e9
---

# Guide Zabbix complet — Supervision d'infrastructure en production

1. Installez Grafana, ajoutez le plugin **Zabbix datasource** (alexanderzobnin-zabbix-app).
2. *Connections → Data sources → Zabbix* : URL `http://192.168.10.50/api_jsonrpc.php`, credentials d'un compte API **lecture seule** dédié.
3. Créez des dashboards : courbes lissées, annotations d'événements, variables `$host`, `$group`.

Cas d'usage typiques :

- **Rapport direction** : disponibilité mensuelle, jolis graphes exportables PDF.
- **Corrélation** : charge onduleur (Zabbix) + température salle (sonde) + incidents (annotations manuelles).
- **Mur d'écrans** : playlists Grafana en rotation.

> 💡 Compte API dédié `grafana-reader` (rôle User, accès Read sur les groupes nécessaires, **pas** d'accès admin). Si le token fuit, l'impact est limité.

## 68. Sécurité et durcissement en production

🖨️ **Checklist durcissement** :

- [ ] **HTTPS** sur le frontend (Let's Encrypt ou CA interne) :
```bash
apt install -y certbot python3-certbot-nginx
certbot --nginx -d zabbix.votre-domaine.local
```
- [ ] **TLS/PSK** sur tous les flux agent/proxy (section 20)
- [ ] **SNMP v3** partout en production (section 49) ; communautés v2c complexes et ACL sur les équipements
- [ ] Compte `Admin` renommé ou désactivé après création des comptes nominatifs
- [ ] **2FA** pour les Super admins (⚠️ natif 6.4+/7.x)
- [ ] LDAP/AD pour l'authentification centralisée
- [ ] Pare-feu sur le server : n'ouvrir que 10051 (agents/proxys), 443 (admins), 162/UDP (traps)
```bash
# Exemple nftables/ufw (à adapter)
ufw allow from 192.168.10.0/24 to any port 10051 proto tcp
ufw allow from 192.168.10.0/24 to any port 443 proto tcp
ufw allow 162/udp
```
- [ ] Base de données : écoute sur localhost ou réseau de gestion uniquement, mot de passe fort, pas de compte `zabbix` superuser
- [ ] `AllowRoot=0` sur les agents (défaut) ; `DenyKey=` pour bloquer les clés dangereuses :
```ini
# zabbix_agent2.conf — bloque l'exécution de commandes arbitraires
DenyKey=system.run[*]
```
- [ ] Mises à jour OS régulières ; Zabbix mineur à jour (section 72)
- [ ] Logs : `LogFileSize` avec rotation, envoi vers syslog centralisé si possible
- [ ] Sauvegarde chiffrée (section 70)

> ⚠️ **Ne jamais** exposer le frontend Zabbix sur internet sans : HTTPS + 2FA + restriction IP (ou VPN). Un Zabbix compromis = cartographie complète de votre SI offerte à l'attaquant.

## 69. Superviser Zabbix lui-même (qui surveille le surveillant ?)

Créez un hôte `zabbix-srv` (interface agent) + liez les templates officiels :

- `Zabbix server health` : queue, caches, pollers busy, housekeeper, valeur par seconde
- `Linux by Zabbix agent` : OS du server
- Base : `PostgreSQL by Zabbix agent 2` ou `MySQL by Zabbix agent 2`

**Triggers vitaux du monitoring** :

```
# File d'attente qui grandit → server sous-dimensionné ou base lente
{zabbix-srv:zabbix[queue].avg(5m)}>50                    → High
# Cache valeurs saturé
{zabbix-srv:zabbix[wcache,values,all].last()}>80         → High
# Pollers saturés
{zabbix-srv:zabbix[process,poller,avg,busy].avg(5m)}>75  → Average
# Housekeeper à la traîne
{zabbix-srv:zabbix[housekeeper,lastrun].fuzzytime(86400)}=0 → Average
# Le service est-il vivant ? (depuis un proxy ou un 2e site)
{zabbix-srv:net.tcp.service[http,,80].last()}=0          → Disaster
```

> 💡 **Règle absolue** : les alertes "Zabbix lui-même en panne" doivent partir par un **canal indépendant** (SMS via la passerelle GSM locale, pas l'email qui dépend du même réseau). Sinon, quand tout tombe, vous n'apprenez rien.

**Supervision croisée** : un 2e Zabbix (ou un simple script cron + `zabbix_sender` depuis un autre site) qui ping le server principal. En version simple :

```bash
# Cron toutes les 5 min sur une machine distante
*/5 * * * * /usr/bin/zabbix_sender -z 192.168.10.50 -s "watchdog-externe" \
  -k zabbix.frontend.http -o $(curl -s -o /dev/null -w "%{http_code}" https://zabbix.votre-domaine.local/)
```

Trigger : `{watchdog:zabbix.frontend.http.last()}<>200`.

## 70. Sauvegarde : base de données, fichiers, stratégie 3-2-1

**Ce qu'il faut sauvegarder** :

| Quoi | Où | Fréquence |
|---|---|---|
| Base Zabbix complète | `pg_dump` / `mysqldump` | Quotidienne |
| `/etc/zabbix/` | Fichiers de config | À chaque changement (+ Git) |
| `/usr/lib/zabbix/alertscripts/` | Scripts d'alerte | À chaque changement |
| `/usr/share/zabbix/` (custom) | Images, cartes personnalisées | Hebdo |
| Templates exportés (YAML) | Git | À chaque changement |

### PostgreSQL

```bash
#!/bin/bash
# /usr/local/bin/backup_zabbix.sh
set -e
DATE=$(date +%F)
DEST=/backup/zabbix
mkdir -p "$DEST"
sudo -u postgres pg_dump -Fc zabbix > "$DEST/zabbix_$DATE.dump"
# Rétention 30 jours
find "$DEST" -name "zabbix_*.dump" -mtime +30 -delete
# Configs
tar czf "$DEST/zabbix_etc_$DATE.tgz" /etc/zabbix /usr/lib/zabbix/alertscripts
find "$DEST" -name "zabbix_etc_*.tgz" -mtime +30 -delete
```

```bash
chmod +x /usr/local/bin/backup_zabbix.sh
# Cron : tous les jours à 02:00
echo "0 2 * * * root /usr/local/bin/backup_zabbix.sh >> /var/log/zabbix_backup.log 2>&1" \
  > /etc/cron.d/zabbix-backup
```

### MySQL/MariaDB

```bash
mysqldump --single-transaction --routines --events -u zabbix -p'mot_de_passe_ici' \
  zabbix | gzip > /backup/zabbix/zabbix_$(date +%F).sql.gz
```

> ⚠️ **3-2-1** : 3 copies, 2 supports différents, 1 hors site. La sauvegarde Zabbix doit partir **hors du server Zabbix** (NAS, autre site via le proxy, stockage objet). Une sauvegarde sur le même disque que la base ne protège de rien.

### Sauvegarde applicative (items, triggers, templates)

```bash
# Export API de la configuration (sans les données)
# À lancer périodiquement, versionné en Git
python3 export_config.py  # via API : template.get, host.get, action.get...
```

## 71. Restauration : procédure testée

> 💡 **Une sauvegarde non testée = pas de sauvegarde.** Testez la restauration **chaque trimestre** sur une VM isolée.

### Restauration PostgreSQL

```bash
# 1. Stopper les services
systemctl stop zabbix-server zabbix-proxy nginx php8.2-fpm
# 2. Recréer la base vide
sudo -u postgres psql -c "DROP DATABASE zabbix;"
sudo -u postgres psql -c "CREATE DATABASE zabbix OWNER zabbix;"
# 3. Restaurer
sudo -u zabbix pg_restore -d zabbix /backup/zabbix/zabbix_2026-09-25.dump
# 4. Restaurer les configs
tar xzf /backup/zabbix/zabbix_etc_2026-09-25.tgz -C /
# 5. Redémarrer
systemctl start zabbix-server nginx php8.2-fpm
# 6. Vérifier
tail -n 30 /var/log/zabbix/zabbix_server.log
```

### Restauration MySQL

```bash
zcat /backup/zabbix/zabbix_2026-09-25.sql.gz | mysql -u zabbix -p zabbix
```

🖨️ **Checklist de test de restauration** :

- [ ] Restauration effectuée sur machine isolée
- [ ] Frontend accessible, login OK
- [ ] Hôtes et derniers événements visibles
- [ ] Un trigger de test se déclenche
- [ ] Durée mesurée : ______ (objectif < 1 h pour un parc moyen)

## 72. Mise à jour de version : méthode sans stress

### Mise à jour mineure (7.0.1 → 7.0.5) : simple

```bash
# 1. Sauvegarde complète (section 70) — NON NÉGOCIABLE
/usr/local/bin/backup_zabbix.sh
# 2. Snapshot VM si possible
# 3. Mise à jour
apt update && apt install --only-upgrade zabbix-server-pgsql zabbix-frontend-php \
  zabbix-agent2 zabbix-sql-scripts
# 4. Redémarrage (le server migre la base automatiquement si besoin)
systemctl restart zabbix-server zabbix-agent2 nginx php8.2-fpm
# 5. Vérification
tail -n 50 /var/log/zabbix/zabbix_server.log | grep -i "database upgrade\|started"
```

### Mise à jour majeure (6.0 LTS → 7.0 LTS) : méthodique

