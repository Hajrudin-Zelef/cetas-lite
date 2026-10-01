---
id: collect-261001-rattrapage/rattrapage/zabbix-guide-10
title: "Guide Zabbix complet — Supervision d'infrastructure en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-10-01", "2026-11-01"]
keywords: ["agent", "memory"]
source: docs/RAG/collect-261001-rattrapage/zabbix_guide.md
source_anchor: ""
source_lines: [1636, 1813]
sha256: 1942723df35dd2e089537e4a523420b26bc854628bb7c08a6b628d8e9da8290c
---

# 🟡 AVERAGE — Température batterie > 40 °C (vieillissement accéléré)
{ups:battery.temp.avg(10m)}>{$UPS_TEMP_CRIT}

# 🟡 AVERAGE — Pas de données onduleur depuis 3 min
{ups:agent.ping.nodata(3m)}=1   (via icmpping sur l'hôte)

# 🔵 INFO — Retour sur secteur (recovery du trigger batterie)
  Recovery expression : {ups:upsOutputSource.last()}=3
```

> 💡 **L'autonomie restante** (`upsEstimatedMinutesRemaining`) est l'item le plus précieux en coupure : affichez-la en grand sur le dashboard d'astreinte (section 63) et utilisez-la dans l'escalade SMS ("il reste ~22 min").

## 53. Supervision JMX (Java/Tomcat)

Pour les applications Java (Tomcat, Kafka, Elasticsearch…) : activez JMX côté applicatif, puis items **JMX agent** dans Zabbix (interface JMX, port 12345 par défaut).

```bash
# Exemple : Tomcat avec JMX distant
CATALINA_OPTS="$CATALINA_OPTS -Dcom.sun.management.jmxremote \
  -Dcom.sun.management.jmxremote.port=12345 \
  -Dcom.sun.management.jmxremote.authenticate=false \
  -Dcom.sun.management.jmxremote.ssl=false \
  -Djava.rmi.server.hostname=192.168.10.70"
```

Item JMX : clé `jmx["java.lang:type=Memory",HeapMemoryUsage.used]`, type *JMX agent*.

> ⚠️ JMX sans auth ni SSL = à réserver au réseau de gestion isolé. En production, activez l'authentification JMX.

## 54. Supervision IPMI (températures, ventilateurs, alimentations)

Pour le **hardware** des serveurs (Dell iDRAC, HPE iLO, Supermicro…) : interface **IPMI** sur l'hôte, items de type *IPMI agent*.

Prérequis server : `apt install -y libopenipmi0` (souvent déjà dépendance).

Clés IPMI : le nom du capteur tel que remonté (`System Temp`, `FAN 1`, `PS1 Status`) — découvrez-les via la LLD IPMI du template `IPMI` officiel, ou :

```bash
ipmitool -I lanplus -H 192.168.10.61 -U ADMIN -P 'mot_de_passe_ici' sensor list
```

Items types :

```
Capteur : System Temp → trigger si > 75 °C
Capteur : PS1 Status / PS2 Status → trigger si <> "OK" (redondance d'alim !)
Capteur : FAN 1..N → trigger si vitesse = 0
```

> 💡 Pour vous, la **redondance d'alimentation** (`PS1/PS2 Status`) est un item à mettre en HIGH : une alim redondante en panne = plus de redondance, et la prochaine panne = coupure.

## 55. Supervision VMware (intégration native)

Zabbix supervise **vCenter/ESXi** sans agent : *Configuration → Hosts → Create host*, template `VMware Hypervisor` ou `VMware Guest`, macros :

```
{$VMWARE.URL} = https://vcenter.votre-domaine.local/sdk
{$VMWARE.USERNAME} = zabbix-mon@votre-domaine.local
{$VMWARE.PASSWORD} = mot_de_passe_ici
```

La LLD découvre hyperviseurs, VMs, datastores. Surveillez : CPU ready, ballooning mémoire, espace datastore, snapshots oubliés.

> ⚠️ Créez un compte vSphere **lecture seule** dédié. Intervalle de collecte VMware : 5 min minimum (l'API n'aime pas le polling agressif).

## 56. Supervision web : scénarios et checks HTTP

*Configuration → Hosts → [hôte] → Web scenarios*. Exemple : vérifier le portail interne toutes les 5 min :

```
Nom : Portail intranet
Étapes :
  1. Page d'accueil : URL https://intranet.local/, code attendu 200, contient "Bienvenue"
  2. Login : POST https://intranet.local/login, variables user/pass, contient "Déconnexion"
```

Triggers auto : `web.test.fail[Portail intranet]` > 0, `web.test.time[...]` > 5 s.

L'**HTTP agent** (item) permet d'interroger une API REST avec JSONPath en prétraitement — idéal pour les sondes modernes.

## 57. Maintenance planifiée (sans alerte parasite)

*Configuration → Maintenance → Create* : pendant une maintenance, les triggers ne génèrent **pas d'actions** (mais les données continuent d'être collectées — les graphes restent complets).

Cas types :

| Maintenance | Période | Avec collecte de données |
|---|---|---|
| MCO onduleur (bypass) | Ponctuelle, 4 h | Non (pauses) — les tests faussent les mesures |
| Mise à jour serveurs | Hebdo, dimanche 02:00-06:00 | Oui |
| Travaux électriques site | Ponctuelle | Non |

> 💡 **Toujours** poser une maintenance avant une intervention : une alerte "onduleur sur batterie" pendant VOTRE bypass = bruit + SMS inutiles + perte de crédibilité du système d'alerte.

Astuce : maintenance **par tag** (⚠️ 6.x+) — taggez les hôtes `maintenance:elec` et ciblez la maintenance sur le tag.

## 58. Haute disponibilité (HA natif Zabbix 6+)

Depuis la 6.0, le **cluster HA natif** : plusieurs nœuds server en actif/passif automatique, sans outil tiers.

```ini
# /etc/zabbix/zabbix_server.conf — sur chaque nœud
HANodeName=zabbix-node-1      # unique par nœud !
NodeAddress=192.168.10.50:10051
```

1. Installez `zabbix-server-pgsql` sur 2 nœuds (même version !), **même base partagée**.
2. `HANodeName` différent par nœud, `NodeAddress` = IP:port d'écoute de chaque nœud.
3. Démarrez les deux : l'un devient **active**, l'autre **standby** (bascule auto en ~10-30 s).

Vérification : *Administration → General → HA cluster* (⚠️ 6.4+/7.x : *Administration → High availability*).

```
zabbix-node-1   active    last access 5s ago
zabbix-node-2   standby   last access 8s ago
```

> ⚠️ Le HA Zabbix ne couvre **que le server** : la base (PostgreSQL Patroni/réplication) et le frontend (2 instances + reverse proxy) doivent être redondés séparément. Et les **proxies** : déployez-en 2 par site critique avec bascule manuelle des hôtes, ou acceptez le buffer offline.

## 59. Housekeeping : comprendre la rétention des données

Le **housekeeper** supprime périodiquement : history > X jours, trends > Y jours, events, sessions, audit.

Tables concernées : `history`, `history_uint`, `history_str`, `history_text`, `history_log`, `trends`, `trends_uint`, `events`, `alerts`, `auditlog`.

Réglages (*Administration → General → Housekeeping*) — rappelez-vous la section 22 :

```
History storage period : 90d
Trend storage period   : 730d
```

Surveillez le housekeeper lui-même : si `zabbix[housekeeper,...]` (clé interne, section 69) montre qu'il ne suit plus, la base grossit indéfiniment.

> 💡 Alternative moderne : **désactiver le housekeeping interne** et laisser **TimescaleDB** gérer la rétention par policy (compression + drop des chunks). Plus performant sur gros volumes (section 60).

## 60. Partitionnement de la base (PostgreSQL/MySQL) pour les gros parcs

### Avec TimescaleDB (recommandé)

Le script `timescaledb/schema.sql` (section 14.4) a déjà converti `history*` et `trends*` en hypertables. Ajoutez les policies :

```sql
-- Compression après 7 jours, suppression après 90 jours (history)
SELECT add_compression_policy('history', INTERVAL '7 days');
SELECT add_retention_policy('history', INTERVAL '90 days');
SELECT add_compression_policy('history_uint', INTERVAL '7 days');
SELECT add_retention_policy('history_uint', INTERVAL '90 days');
-- Trends : compression après 30 j, rétention 2 ans
SELECT add_compression_policy('trends', INTERVAL '30 days');
SELECT add_retention_policy('trends', INTERVAL '730 days');
```

Vérifiez le gain :

```sql
SELECT hypertable_name,
       pg_size_pretty(hypertable_size(hypertable_name)) AS taille
FROM timescaledb_information.hypertables;
```

### Avec MySQL/MariaDB (partitionnement manuel)

Zabbix fournit un exemple : `/usr/share/doc/zabbix-server-mysql*/create.sql` (script de partitionnement par jour). Principe :

```sql
-- Exemple de partition mensuelle sur history_uint (à adapter)
ALTER TABLE history_uint PARTITION BY RANGE (clock)
(PARTITION p202609 VALUES LESS THAN (UNIX_TIMESTAMP('2026-10-01')),
 PARTITION p202610 VALUES LESS THAN (UNIX_TIMESTAMP('2026-11-01')));
```

> ⚠️ Le partitionnement MySQL manuel exige une **procédure stockée + event scheduler** pour créer/supprimer les partitions. C'est une raison de plus de choisir PostgreSQL + TimescaleDB pour un nouveau déploiement.

## 61. Performance et tuning du server/proxy

### Caches (dans `zabbix_server.conf`)

