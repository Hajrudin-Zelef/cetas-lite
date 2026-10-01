---
id: collect-261001-rattrapage/rattrapage/zabbix-guide-13
title: "Guide Zabbix complet — Supervision d'infrastructure en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "arr", "valuation"]
source: docs/RAG/collect-261001-rattrapage/zabbix_guide.md
source_anchor: ""
source_lines: [2167, 2329]
sha256: 6e7332c9d3fc04bb067fc9c1e878b8a95bc28c31d2cd2b112c0818b9259f2a91
---

# Guide Zabbix complet — Supervision d'infrastructure en production

1. **Lisez les release notes** et la section "Upgrade notes" de la doc officielle (breaking changes : PHP, base, templates).
2. **Maquette** : clonez la prod (VM + dump base), faites la montée en version dessus, validez 1 semaine.
3. **Sauvegarde** + snapshot.
4. Changez le dépôt (`zabbix-release` 7.0), `apt update`, `apt install` des paquets.
5. Le server **7.0** migre automatiquement le schéma 6.0 → 7.0 au démarrage (peut prendre 10-60 min sur grosse base — **ne redémarrez pas pendant la migration**).
6. Mettez à jour les **proxies** (un proxy 6.0 ne parle pas à un server 7.0 — même majeure requise).
7. Les **agents** peuvent rester en version N-1 temporairement (compatibilité ascendante).
8. Videz le cache navigateur, vérifiez les dashboards custom.

> ⚠️ **Ordre impératif** : server d'abord, puis proxies, puis agents. Jamais l'inverse. Et **jamais** de saut de 2 majeures (5.0 → 7.0) : passez par les intermédiaires ou réinstallez proprement.

## 73. Dépannage : méthodologie générale

Quand quelque chose ne marche pas, suivez la chaîne **collecte → stockage → évaluation → alerte → visualisation** (section 4) :

```
1. L'item a-t-il des "Latest data" ?  (Monitoring → Latest data)
   NON → problème de COLLECTE : testez avec zabbix_get / snmpwalk depuis le server
   OUI → continuez
2. Le trigger s'évalue-t-il ? (statut Unknown ? → pas assez d'historique / fonction mal paramétrée)
3. L'action matche-t-elle ? (Reports → Audit / Action log : l'action s'est-elle déclenchée ?)
4. Le média fonctionne-t-il ? (Administration → Media types → Test)
5. Le dashboard affiche-t-il ? (problème de droits ? de filtre temporel ?)
```

**Les 3 logs à connaître** :

```bash
tail -f /var/log/zabbix/zabbix_server.log    # le server
tail -f /var/log/zabbix/zabbix_agent2.log    # l'agent (côté supervisé)
tail -f /var/log/zabbix/zabbix_proxy.log     # le proxy
```

Augmentez temporairement la verbosité : `LogLevel=4` (debug) dans la conf + `systemctl restart`, **puis remettez 3** après diagnostic (le debug remplit le disque).

## 74. Erreur n°1 à n°12 : les classiques commentés

### Erreur n°1 — Le server ne démarre pas : connexion base impossible

```
[Z3001] connection to database 'zabbix' failed: [1045] Access denied
```

**Causes** : mauvais `DBPassword`, utilisateur inexistant, base non créée, PostgreSQL n'écoute pas.
**Correctifs** :

```bash
# Tester la connexion manuellement
sudo -u zabbix psql -h localhost -U zabbix zabbix -c "SELECT 1;"
# Vérifier pg_hba.conf : méthode md5/scram pour l'utilisateur zabbix
grep -v "^#" /etc/postgresql/15/main/pg_hba.conf | grep -v "^$"
```

### Erreur n°2 — Hôte en gris/rouge : "Get value from agent failed"

```
ZBX_NOTSUPPORTED / cannot connect to [[192.168.10.61]:10050]: connection refused
```

**Checklist** :
1. L'agent tourne-t-il ? `systemctl status zabbix-agent2` sur le supervisé.
2. Le pare-feu ? `ss -tlnp | grep 10050` ; test depuis le server : `zabbix_get -s 192.168.10.61 -k agent.ping`.
3. `Server=` dans `zabbix_agent2.conf` contient-il l'IP du server ? (sinon l'agent refuse la connexion — silencieux côté agent, timeout côté server).
4. En mode actif : `Hostname=` **exactement** égal au nom dans le frontend (casse, espaces).

### Erreur n°3 — Import du schéma MySQL : charset incorrect

```
ERROR 1366 (HY000): Incorrect string value...
```

**Cause** : base créée sans `utf8mb4`. **Correctif** : recréer la base avec `CHARACTER SET utf8mb4 COLLATE utf8mb4_bin` (section 15). Il n'y a pas de raccourci propre : refaites l'import.

### Erreur n°4 — "No data" partout après ajout d'un proxy

**Cause** : le proxy n'est pas déclaré côté server, ou `Hostname` différent entre `zabbix_proxy.conf` et *Administration → Proxies*.
**Correctif** : les 3 noms doivent matcher : `Hostname=` du proxy, nom du proxy dans le frontend, et les hôtes doivent avoir *Monitored by* = ce proxy. Log proxy : `cannot send list of active checks` = le server ne connaît pas ce proxy.

### Erreur n°5 — Items SNMP "Not supported"

**Causes** : mauvais OID, communauté erronée, pare-feu UDP/161, équipement sans SNMP activé.
**Diagnostic** :

```bash
# Depuis le server Zabbix lui-même :
snmpwalk -v2c -c communaute_lecture_ici 192.168.10.90 1.3.6.1.2.1.33.1.2.3.1.1.2.1
# Timeout → réseau/communauté ; "No Such Object" → mauvais OID
```

> 💡 90 % des problèmes SNMP se résolvent en CLI avant même d'ouvrir le frontend.

### Erreur n°6 — Trigger en "Unknown"

**Causes** : pas encore de données (nouvel item), fonction sur période trop longue (`avg(1h)` avec 5 min d'historique), division par zéro dans un calculé.
**Correctif** : attendez la première collecte ; *Monitoring → Latest data* pour vérifier que l'item reçoit des valeurs.

### Erreur n°7 — Tempête d'emails après une coupure

**Cause** : pas de dépendances (section 39), pas de maintenance (section 57).
**Correctif immédiat** : *Configuration → Maintenance* d'urgence sur le site ; **correctif durable** : arbre de dépendances + action "site down" unique au lieu de N actions par hôte.

### Erreur n°8 — "Value cache is fully used"

```
value cache is fully used
```

**Cause** : `ValueCacheSize` trop petit (section 61). **Correctif** : doublez `ValueCacheSize`, redémarrez le server, surveillez `zabbix[wcache,values,all]`.

### Erreur n°9 — Le frontend affiche "Database error"

**Causes** : base arrêtée, `zabbix.conf.php` avec mauvais mot de passe, base pleine (disque).
**Diagnostic** :

```bash
df -h /var/lib/postgresql
systemctl status postgresql
# Tester les identifiants du frontend :
grep DB_ /etc/zabbix/web/zabbix.conf.php
```

### Erreur n°10 — Les graphes ont des "trous"

**Causes** : collecte plus lente que l'intervalle (pollers saturés), lien instable (normal — le proxy bufferise, section 7), NTP désynchronisé, housekeeper agressif.
**Diagnostic** : *Administration → Queue* (âge des items en retard), `zabbix[queue]`, vérifier chrony des deux côtés.

### Erreur n°11 — L'alerte SMS ne part pas

**Checklist** :
1. *Administration → Media types → [SMS] → Test* : fonctionne ?
2. Le script a-t-il les bons droits ? (`zabbix:zabbix`, +x)
3. La clé 4G a-t-elle du crédit / du réseau ? (`gammu networkinfo`)
4. L'action a-t-elle vraiment déclenché une opération SMS ? (*Reports → Action log*)
5. Le numéro du destinataire est-il au format international ? (`+225...`)

### Erreur n°12 — Après mise à jour : "Database version mismatch"

```
Database version does not match: expected X, got Y
```

**Cause** : le server a été mis à jour mais pas les scripts SQL, ou la migration a été interrompue.
**Correctif** : installez `zabbix-sql-scripts` de la **même version** que le server, relancez le server et **laissez la migration se terminer** (surveillez le log — ne redémarrez pas). En cas d'interruption brutale : restaurez la sauvegarde pré-upgrade (section 71) et recommencez.

## 75. Cas pratique n°1 : supervision d'une baie complète

**Objectif** : 1 switch + 4 serveurs + 1 onduleur dans une armoire, alertes propres.

1. Groupes : `Baie-A3/Réseau`, `Baie-A3/Serveurs`, `Baie-A3/Onduleurs`.
2. Hôtes + templates : switch (SNMP), serveurs (Agent 2), onduleur (template section 52).
3. Dépendances : triggers serveurs/onduleur "injoignable" → dépendent du trigger "switch injoignable".
4. Maintenance récurrente : créneau MCO mensuel.
5. Dashboard "Baie A3" : map de l'armoire + problèmes filtrés.
6. Action : High+ du groupe `Baie-A3/*` → Email + Telegram, escalade SMS à 15 min.

**Résultat attendu** : une coupure du switch = **1 seule alerte** (le switch), pas 6.

## 76. Cas pratique n°2 : supervision d'un onduleur 40 kVA de A à Z

**Contexte** : onduleur `ups-salle-01` (192.168.10.90), 40 kVA, salle serveur principale.

**Étape 1 — Validation SNMP** (depuis le server) :

