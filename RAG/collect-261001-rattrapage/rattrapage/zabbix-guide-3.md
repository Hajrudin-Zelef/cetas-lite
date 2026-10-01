---
id: collect-261001-rattrapage/rattrapage/zabbix-guide-3
title: "Guide Zabbix complet — Supervision d'infrastructure en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "apache"]
source: docs/RAG/collect-261001-rattrapage/zabbix_guide.md
source_anchor: ""
source_lines: [257, 415]
sha256: cf7884c38edfe275ecd67456da2bd03c39e3172777548bc4ac3ff8d94fe9c1a6
---

# Guide Zabbix complet — Supervision d'infrastructure en production

> 💡 **Cas d'usage typique en Afrique de l'Ouest** : un proxy par site distant (agence, site technique) avec `ProxyMode=0` (actif). Le proxy encaisse les coupures de liaison sans perdre l'historique local.

Paramètres clés : `ProxyMode=`, `Server=`, `Hostname=` (doit matcher le nom du proxy créé dans le frontend), `DBName/DBUser/DBPassword`, `ProxyOfflineBuffer=` (défaut 1h — montez à 24h+ sur lien instable), `HeartbeatFrequency=`.

## 8. Les composants en détail : Frontend web

Interface PHP (`zabbix-web`) : dashboards, configuration, graphes, cartes, inventaire, audit. Elle **ne collecte rien** : elle lit/écrit uniquement en base.

- Serveur web : **Nginx (recommandé)** ou Apache + PHP-FPM.
- ⚠️ **Zabbix 7.x exige PHP 8.0+** ; Zabbix 6.0 LTS accepte PHP 7.2.5+. Vérifiez la matrice officielle pour votre version exacte.
- Le frontend peut être installé **sur une machine dédiée** (bonne pratique : séparer web et server en production).
- Fichier de config : `/etc/zabbix/web/zabbix.conf.php` (généré par l'assistant d'installation).

## 9. Les composants en détail : base de données (MySQL vs PostgreSQL)

| Critère | PostgreSQL (+ TimescaleDB) | MySQL / MariaDB |
|---|---|---|
| Recommandation Zabbix | ✅ **Recommandé officiellement** | Supporté |
| Gros volumes | Excellent (TimescaleDB : compression, retention policies) | Correct avec partitionnement manuel |
| Partitionnement | Natif via TimescaleDB | Manuel (section 60) |
| Simplicité | Un peu plus exigeant | Plus familier pour beaucoup d'équipes |
| HA base | Patroni / streaming replication | Galera / réplication |

> 💡 **Choix recommandé pour un nouveau déploiement** : **PostgreSQL 15/16 + TimescaleDB**. La compression des historiques divise l'espace disque par 5 à 10 sur les métriques numériques — décisif quand on supervise des onduleurs et des switchs avec des pas de 30 s.

Tailles indicatives (à valider) :

| Parc | history (90 j) | trends (2 ans) | Total indicatif |
|---|---|---|---|
| 100 hôtes / 5 000 items | ~30 Go | ~5 Go | ~40 Go |
| 500 hôtes / 30 000 items | ~200 Go | ~30 Go | ~250 Go |
| 1 000 hôtes / 80 000 items | ~500 Go | ~80 Go | ~650 Go |

## 10. Ports réseau et flux — tableau de référence

🖨️ **À afficher en salle d'exploitation / à fournir à l'équipe réseau pour les ouvertures de flux.**

| Source → Destination | Port | Protocole | Usage |
|---|---|---|---|
| Server → Agent | 10050 | TCP | Checks passifs |
| Agent → Server/Proxy | 10051 | TCP | Checks actifs, envoi proxy |
| Proxy → Server | 10051 | TCP | Transfert données (proxy actif) |
| Server → Proxy | 10050 | TCP | Proxy passif (rare) |
| Server → équipement | 161 | UDP | SNMP poll |
| Équipement → Server | 162 | UDP | SNMP traps |
| Server → équipement | 623 | UDP | IPMI |
| Server → équipement | 12345 | TCP | JMX (variable) |
| Navigateur → Frontend | 80/443 | TCP | Interface web |
| Frontend → Base | 5432 / 3306 | TCP | PostgreSQL / MySQL |
| Server → Base | 5432 / 3306 | TCP | Écriture données |
| Server → SMTP | 25/587 | TCP | Alertes email |
| Server → API Telegram | 443 | TCP | Alertes Telegram |

> ⚠️ N'ouvrez que les flux nécessaires. Un agent en **mode 100 % actif** n'a besoin d'aucun flux entrant : seul le 10051/TCP sortant vers le server/proxy.

## 11. Choisir sa version : LTS vs standard

| Version | Type | Fin de support (indicative — vérifiez le site officiel) |
|---|---|---|
| 6.0 | LTS | Déjà/ancien — ne plus déployer en 2026 |
| 6.4 | Standard | Support court — éviter en production |
| **7.0** | **LTS** | ✅ **Choix recommandé en 2026** |
| 7.2 / 7.4 | Standard | Nouveautés, support court |

**Règle d'or** : en production, déployez une **LTS** (7.0 en 2026). Les versions standard servent à tester les nouveautés en labo.

⚠️ Ce guide couvre 6.x et 7.x ; les différences notables sont signalées. Si vous êtes encore en 6.0 LTS, planifiez la migration (section 72).

## 12. Dimensionnement : CPU, RAM, disque, IOPS

Formule de base : comptez en **valeurs par seconde (vps)** = nombre d'items / intervalle moyen.

```
Exemple : 10 000 items à 60 s  →  ~167 vps
          30 000 items à 60 s  →  ~500 vps
```

| Charge | vCPU server | RAM server | Disque base | IOPS |
|---|---|---|---|---|
| < 100 vps (petit parc) | 2 | 4 Go | 100 Go SSD | 500 |
| 100–500 vps (moyen) | 4 | 8–16 Go | 300 Go SSD/NVMe | 2 000 |
| 500–2 000 vps (gros) | 8 | 32 Go | 1 To NVMe | 5 000+ |
| > 2 000 vps | Cluster HA + base dédiée + TimescaleDB | | | |

> 💡 Le **disque de la base** est presque toujours le goulot : privilégiez du **SSD/NVMe local** plutôt qu'un SAN lent. `CacheSize`, `HistoryCacheSize`, `TrendCacheSize` (section 61) absorbent les pics mais ne remplacent pas des IOPS corrects.

## 13. Prérequis système et réseau

🖨️ **Checklist pré-installation :**

- [ ] Debian 12/13 ou Ubuntu 22.04/24.04 à jour (`apt update && apt full-upgrade`)
- [ ] Nom d'hôte FQDN résolu (`hostname -f` répond)
- [ ] NTP/chrony synchronisé (`chronyc tracking` : `Leap status : Normal`)
- [ ] Partition `/var/lib/postgresql` (ou `/var/lib/mysql`) dimensionnée
- [ ] Flux réseau ouverts (section 10)
- [ ] Accès internet pour les dépôts (ou miroir local)
- [ ] Sauvegarde existante du système (snapshot VM avant l'installation)

```bash
# Vérifications de base
hostname -f
timedatectl status | grep -i "synchronized\|NTP"
df -h /var/lib/postgresql
free -h
```

> ⚠️ **Le temps est critique** : un décalage NTP fausse les timestamps des données et casse les triggers `nodata()`. Synchronisez **tous** les supervisés et le server sur la même source NTP.

## 14. Installation pas à pas : Zabbix Server + PostgreSQL sur Debian/Ubuntu

> Procédure pour **Zabbix 7.0 LTS** sur **Debian 12** (adaptable à Ubuntu 22.04/24.04 — seuls les noms de dépôts changent). ⚠️ Vérifiez toujours la [page de téléchargement officielle](https://www.zabbix.com/download) : les URL de dépôts ci-dessous suivent le schéma officiel.

### 14.1. Ajouter le dépôt officiel Zabbix

```bash
# En root
wget https://repo.zabbix.com/zabbix/7.0/debian/pool/main/z/zabbix-release/zabbix-release_latest_7.0+debian12_all.deb
dpkg -i zabbix-release_latest_7.0+debian12_all.deb
apt update
```

Pour Ubuntu 22.04, remplacez `debian` par `ubuntu` et `debian12` par `ubuntu22.04` dans l'URL (schéma officiel).

### 14.2. Installer PostgreSQL et TimescaleDB

```bash
apt install -y postgresql-15 postgresql-client-15
# TimescaleDB (fortement recommandé) — dépôt officiel Timescale :
apt install -y gnupg postgresql-common apt-transport-https lsb-release wget
/usr/share/postgresql-common/pgdg/apt.postgresql.org.sh
echo "deb https://packagecloud.io/timescale/timescaledb/debian/ $(lsb_release -c -s) main" \
  > /etc/apt/sources.list.d/timescaledb.list
wget --quiet -O - https://packagecloud.io/timescale/timescaledb/gpgkey | apt-key add -
apt update
apt install -y timescaledb-2-postgresql-15
```

Activez l'extension dans `postgresql.conf` :

```bash
echo "shared_preload_libraries = 'timescaledb'" >> /etc/postgresql/15/main/postgresql.conf
systemctl restart postgresql
```

### 14.3. Créer la base et l'utilisateur Zabbix

```bash
sudo -u postgres psql <<'EOF'
CREATE USER zabbix WITH PASSWORD 'mot_de_passe_ici';
CREATE DATABASE zabbix OWNER zabbix;
\c zabbix
CREATE EXTENSION IF NOT EXISTS timescaledb;
EOF
```

> 💡 Notez ce mot de passe dans votre coffre (KeePass, Vault…) — il sera réutilisé dans `zabbix_server.conf`.

### 14.4. Installer le server et importer le schéma

