---
id: collect-261001-rattrapage/rattrapage/zabbix-guide-2
title: "Guide Zabbix complet — Supervision d'infrastructure en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "valuation"]
source: docs/RAG/collect-261001-rattrapage/zabbix_guide.md
source_anchor: ""
source_lines: [135, 256]
sha256: f89d23f0ba3ca829cba45670d21256ff8a72d5e86cee25615ec5d6df0c8e8ae9
---

# Guide Zabbix complet — Supervision d'infrastructure en production

| Convention | Signification |
|---|---|
| `code en chasse fixe` | Commande, fichier, valeur à saisir |
| ⚠️ | Dépend de la version (6.x vs 7.x) ou point de vigilance |
| 💡 | Conseil de terrain / retour d'expérience |
| 🖨️ | Checklist imprimable |
| `mot_de_passe_ici` | Mot de passe fictif à remplacer |
| `#` en début de ligne shell | Commande exécutée en root |
| `$` en début de ligne shell | Commande exécutée en utilisateur normal |

Les exemples utilisent ces noms fictifs (adaptez à votre plan de nommage) :

| Rôle | Nom d'hôte | IP |
|---|---|---|
| Serveur Zabbix | `zabbix-srv` | `192.168.10.50` |
| Proxy site distant | `zabbix-proxy-sud` | `192.168.20.50` |
| Serveur supervisé | `srv-fichiers-01` | `192.168.10.61` |
| Onduleur 40 kVA | `ups-salle-01` | `192.168.10.90` |
| Switch cœur | `sw-coeur-01` | `192.168.10.2` |

## 4. Architecture générale de Zabbix

```
                                    ┌─────────────────────────┐
                                    │      FRONTEND WEB       │
                                    │  (Nginx + PHP-FPM)      │
                                    └────────────┬────────────┘
                                                 │ SQL (lecture/écriture)
┌──────────┐   ┌──────────┐   ┌──────────────────▼──────────────────┐   ┌──────────┐
│ Agent 2  │   │  Proxy   │   │          ZABBIX SERVER              │   │  SNMP    │
│(actif/   │──▶│(site     │──▶│  - collecte (pollers, trappers)     │◀──│(switch,  │
│ passif)  │   │ distant) │   │  - calcul triggers                  │   │ onduleur)│
└──────────┘   └──────────┘   │  - actions / alertes                │   └──────────┘
                              │  - housekeeper                      │
                              └────────────┬──────────────────┘
                                           │ SQL
                              ┌────────────▼────────────┐
                              │   BASE DE DONNÉES       │
                              │ PostgreSQL ou MySQL     │
                              │ (history, trends, events)│
                              └─────────────────────────┘
```

**Flux de données simplifié** :

1. Le **server** (ou le **proxy**) collecte les métriques : il interroge (mode passif), reçoit (mode actif/trapper), ou scrute (SNMP, IPMI, JMX, HTTP).
2. Les valeurs sont stockées dans la **base** : `history` (données brutes, courte rétention) et `trends` (agrégats horaires, longue rétention).
3. Le server **évalue les triggers** à chaque nouvelle valeur.
4. Si un trigger passe à PROBLEM, le server déclenche les **actions** (email, SMS, Telegram…).
5. Le **frontend** lit la base pour afficher dashboards, graphes, événements.

> 💡 Retenez ce mantra : **collecte → stockage → évaluation → alerte → visualisation**. 90 % des problèmes de dépannage se situent sur l'une de ces 5 étapes.

## 5. Les composants en détail : Server

Le **Zabbix Server** est le chef d'orchestre. C'est un démon (`zabbix_server`) écrit en C, multi-processus :

| Processus | Rôle | Paramètre de dimensionnement |
|---|---|---|
| `poller` | Interroge les agents en mode passif, SNMP, etc. | `StartPollers=` |
| `unreachable poller` | Réessaie les hôtes injoignables | `StartPollersUnreachable=` |
| `trapper` | Reçoit les données envoyées (actif, trapper, proxy) | `StartTrappers=` |
| `preprocessing` | Applique le prétraitement des items | `StartPreprocessors=` |
| `history syncer` | Écrit les valeurs en base | `StartHistorySyncers=` / `HistoryStorageDateIndex` |
| `escalator` | Gère les escalades d'actions | `StartEscalators=` |
| `alert manager` | Dispatche vers les médias | `StartAlerters=` |
| `housekeeper` | Supprime les vieilles données | `HousekeepingFrequency=` |
| `discoverer` | Découverte réseau | `StartDiscoverers=` |
| `http poller` | Checks web/HTTP | `StartHTTPPollers=` |
| `snmp trapper` | Reçoit les SNMP traps | `StartSNMPTrapper=` (désactivé par défaut) |

⚠️ **Zabbix 7.x** : le server sait écrire l'historique dans un ** TimescaleDB** ? Non — restons précis : Zabbix supporte officiellement **PostgreSQL + TimescaleDB** (recommandé pour les gros volumes, compression native) et **MySQL/MariaDB**. Le support TimescaleDB existe depuis la 5.0.

Points clés à retenir :

- Le server est **le seul** à évaluer les triggers et à envoyer les alertes (le proxy ne fait que collecter et relayer).
- Le server peut fonctionner en **cluster HA natif** depuis la 6.0 (section 58).
- Fichier de config : `/etc/zabbix/zabbix_server.conf`.

## 6. Les composants en détail : Agent 2 (et Agent 1)

Deux agents coexistent :

| Critère | Zabbix Agent (1) | Zabbix Agent 2 |
|---|---|---|
| Langage | C | Go |
| Plugins | Limités | **Modulaires** (natif : PostgreSQL, MySQL, Redis, Docker, systemd, MQTT, SNMP…) |
| Checks actifs/passifs | Oui | Oui |
| Exécution persistante | Non | Oui (connexions maintenues) |
| Recommandé pour | Parc existant | ✅ **Tout nouveau déploiement** |

> 💡 **Déployez Agent 2 partout** sur les nouveaux serveurs. L'agent 1 reste supporté mais ne reçoit plus les nouveautés.

**Modes de fonctionnement** :

- **Passif** : le server/proxy **demande** la valeur à l'agent (port 10050/TCP sur l'agent). Simple, mais le server porte la charge de planification.
- **Actif** : l'agent **envoie** les valeurs au server/proxy (port 10051/TCP côté server). L'agent récupère sa liste d'items à surveiller, puis pousse. Idéal derrière un NAT ou un pare-feu (seul le flux sortant est nécessaire), et recommandé pour les gros parcs.

```
PASSIF :  server ──(demande :10050)──▶ agent ──(réponse)──▶ server
ACTIF  :  agent ──(récupère liste :10051)──▶ server
          agent ──(envoie valeurs :10051)──▶ server
```

Fichier de config : `/etc/zabbix/zabbix_agent2.conf`. Clés utiles : `Server=` (qui peut interroger), `ServerActive=` (à qui envoyer en actif), `Hostname=` (doit matcher le nom dans le frontend !).

## 7. Les composants en détail : Proxy

Le **proxy** (`zabbix_proxy`) collecte pour le compte du server sur un site distant ou un segment isolé :

- **Stocke localement** (SQLite par défaut, MySQL/PostgreSQL possible) puis **transfère** au server. Si le lien tombe, **aucune donnée n'est perdue** (dans la limite du disque et de `ProxyOfflineBuffer`).
- **Modes** : actif (le proxy se connecte au server — à privilégier derrière un NAT) ou passif (le server se connecte au proxy).
- Le proxy **ne calcule pas les triggers** et **n'envoie pas d'alertes** : il relaie les données brutes.

```
Site distant (liaison 4G instable)
┌─────────────┐      ┌──────────────┐              ┌──────────────┐
│  Agents     │─────▶│    PROXY     │────(quand le──▶│    SERVER    │
│  SNMP       │      │ (buffer local│   lien est up) │  (central)   │
└─────────────┘      └──────────────┘              └──────────────┘
```

