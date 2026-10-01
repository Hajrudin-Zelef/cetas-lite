---
id: collect-261001-rattrapage/rattrapage/zabbix-guide-11
title: "Guide Zabbix complet — Supervision d'infrastructure en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/zabbix_guide.md
source_anchor: ""
source_lines: [1814, 1978]
sha256: 4a0ac82d3ce72751b9e6450185950700492a0c733a2b6ba0f64f6897ea9d579b
---

# Guide Zabbix complet — Supervision d'infrastructure en production

| Paramètre | Défaut | Parc moyen | Rôle |
|---|---|---|---|
| `CacheSize` | 32M | 256M–1G | Config (hosts, items, triggers) |
| `HistoryCacheSize` | 16M | 128–512M | Valeurs en attente d'écriture |
| `TrendCacheSize` | 4M | 128M | Agrégats |
| `ValueCacheSize` | 8M | 256M–1G | ⚠️ Le plus critique pour les triggers |
| `HistoryIndexCacheSize` | — | 64M | Index (7.x) |

Symptômes d'un cache trop petit : `zabbix[wcache,...]` > 80 %, triggers en retard, "value cache is fully used" dans les logs.

### Pollers

Règle : augmentez `StartPollers` tant que la file (`zabbix[queue]`, *Administration → Queue*) ne se résorbe pas. Repères :

```
< 1 000 items SNMP : StartPollers=20 suffit
5 000+ items       : 40-60
```

Ne dépassez pas ~2× le nombre de cœurs sans mesurer : trop de pollers = contention base.

### Base de données

```ini
# postgresql.conf (serveur dédié 16 Go RAM, à ajuster)
shared_buffers = 4GB
effective_cache_size = 12GB
work_mem = 64MB
maintenance_work_mem = 512MB
max_connections = 200
```

```bash
# Surveillez les requêtes lentes
# postgresql.conf :
log_min_duration_statement = 1000   # log > 1 s
```

### Clés internes de monitoring (section 69)

```
zabbix[queue]              # items en attente (> 10 = sous-dimensionné)
zabbix[wcache,values,all]  # % cache valeurs
zabbix[process,poller,avg,busy]  # % occupation pollers
```

> 💡 **Tunez par la mesure** : dashboard "Zabbix server health" (template `Zabbix server health` officiel) affiché en permanence pendant 1 semaine avant de toucher aux paramètres.

## 62. Usage quotidien : la routine de l'exploitant

🖨️ **Routine quotidienne (15 min)** :

- [ ] *Monitoring → Problems* : 0 problème non acquitté en High/Disaster ?
- [ ] *Monitoring → Dashboard* "Santé Zabbix" : queue < 10, caches < 80 % ?
- [ ] *Reports → Availability* : un hôte < 99 % sur 24 h ?
- [ ] Vérifier les acquittements de la veille (commentaires renseignés ?)
- [ ] *Administration → Queue* : aucun pic anormal ?

🖨️ **Routine hebdomadaire (1 h)** :

- [ ] *Monitoring → Discovery* : nouveaux équipements à intégrer ?
- [ ] Items **Not supported** (*Configuration → Hosts → Items*, filtre) : corriger ou désactiver
- [ ] Test d'envoi SMS mensuel (section 42) — noter le résultat
- [ ] Relecture d'un trigger bruyant : ajuster seuil/hystérésis (section 37)
- [ ] Vérifier l'espace disque de la base (`df -h`, taille des tables)
- [ ] Contrôler la sauvegarde (restauration test trimestrielle, section 71)

🖨️ **Routine mensuelle** :

- [ ] Rapport de disponibilité pour la direction (section 84)
- [ ] Revue des maintenances planifiées du mois suivant
- [ ] Mise à jour mineure Zabbix si disponible (section 72)
- [ ] Test de bout en bout de la chaîne d'alerte (section 83)

## 63. Dashboards et vues : construire des écrans d'exploitation

*Monitoring → Dashboards → Create dashboard*. Widgets essentiels :

| Widget | Usage |
|---|---|
| Problems | Les problèmes en cours (filtre par sévérité/groupe) |
| Graph / Item value | Courbes et valeurs temps réel |
| Top hosts | Top 10 CPU/mémoire/disque |
| Geomap | Sites distants sur carte |
| URL / Clock | Contexte |

### Dashboard "Énergie" (votre écran prioritaire)

```
┌─────────────────────────────────────────────────────┐
│ COUPURES EN COURS (Problems, filtre groupe=Onduleurs)│
├──────────────┬──────────────┬───────────────────────┤
│ ups-salle-01 │ ups-salle-02 │ ups-agence-sud        │
│ Source: BAT. │ Source: sect.│ Source: secteur       │
│ Batterie: 68%│ Batterie:100%│ Batterie: 100%        │
│ Autonomie: 24│ Autonomie: -- │ Autonomie: --         │
│ min          │              │                       │
├──────────────┴──────────────┴───────────────────────┤
│ Charge des onduleurs (graphe multi-hôtes, 24h)       │
│ Température salle (graphe)                           │
└─────────────────────────────────────────────────────┘
```

> 💡 **Un dashboard par public** : "Énergie" (vous), "Réseau" (équipe réseau), "Direction" (disponibilité %, SLA, sans jargon). Le dashboard direction ne montre JAMAIS de problèmes techniques bruts.

Partage : *Dashboard → Sharing* (public/privé, par groupe d'utilisateurs). ⚠️ 7.x : les dashboards peuvent être définis dans les templates.

## 64. Inventaire automatique des hôtes

*Configuration → Hosts → [hôte] → Inventory → Automatic*. L'agent remplit : OS, version, serial, modèle… (items `system.hw.*`, `system.sw.os`).

*Inventory → Overview* : tableau filtrable ("tous les serveurs Debian 12", "tous les switchs Cisco"). Export CSV pour la CMDB.

> 💡 Renseignez aussi les champs **manuels** utiles en intervention : contact, localisation ("Salle serveur, armoire A3"), contrat de maintenance. En astreinte à 3h du matin, ça vaut de l'or.

## 65. Cartes réseau (maps) et écrans de supervision

*Monitoring → Maps → Create map* : fond de plan (importez le plan de votre salle), icônes par type, liens colorés selon l'état.

- **Liens** : entre switch et serveurs, avec indicateurs de débit (`{sw:ifInOctets[{#IFINDEX}]}` en label).
- **Éléments** : icône = dernier état (OK vert, PROBLEM rouge clignotant).
- **Map imbriquées** : carte "Sites" → clic → carte "Site Sud" détaillée.

> 💡 Affichez la carte des sites sur un **écran mural** en mode kiosque (URL + `&kiosk=1`… vérifiez le paramètre selon version, ou mode plein écran du navigateur avec refresh auto). C'est votre "tour de contrôle".

## 66. Bonnes pratiques d'alerting : la chasse au bruit

Le bruit tue la confiance : une équipe qui reçoit 50 emails/jour **ne lit plus** les alertes. Règles d'or :

1. **Chaque alerte doit exiger une action.** Sinon : dashboard, pas email.
2. **Seuils + durée** : jamais de `.last()` seul sur une métrique bruitée (section 37).
3. **Hystérésis** systématique sur les seuils critiques.
4. **Dépendances** (section 39) : pas d'alerte serveur si le switch parent est down.
5. **Plages horaires** : `and {TRIGGER.VALUE}=1` combiné à `time()`/`dayofweek()` pour les alertes ouvrées.
6. **Dédupliquez** : un trigger "disque > 90 %" par partition via LLD, pas 12 triggers manuels.
7. **Acquittement obligatoire** avec commentaire (section 45).
8. **Revue mensuelle du bruit** : *Reports → Notifications* : quel trigger a envoyé le plus d'emails ? Corrigez le top 3.

### Exemple : alerte disque "propre"

```
Nom : "Espace disque {#FSNAME} > {$VFS_FS_PUSED_WARN}% sur {HOST.NAME}"
Problème  : {srv:vfs.fs.size[{#FSNAME},pused].avg(15m)}>{$VFS_FS_PUSED_WARN}
Recovery  : {srv:vfs.fs.size[{#FSNAME},pused].avg(15m)}<{$VFS_FS_PUSED_WARN}-10
Sévérité  : Average (Warning si > 95 via 2e trigger)
Action    : Email groupe Exploitation, heures ouvrées uniquement
```

### Ce qu'il ne faut JAMAIS alerter par SMS

- CPU > 80 % (bruit)
- Disque > 85 % (email suffit)
- Reboot planifié (maintenance !)

### Ce qui mérite le SMS

- Onduleur sur batterie (Disaster)
- Site injoignable (High)
- Température salle > 35 °C (High)
- Batterie onduleur < 30 % pendant coupure (Disaster)

## 67. Intégration avec Grafana

Zabbix a de beaux dashboards ; Grafana excelle dans les **vues mixtes** (Zabbix + autres sources) et les rapports.

