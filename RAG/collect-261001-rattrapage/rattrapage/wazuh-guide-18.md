---
id: collect-261001-rattrapage/rattrapage/wazuh-guide-18
title: "Guide Wazuh — SIEM & XDR Open Source en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents"]
source: docs/RAG/collect-261001-rattrapage/wazuh_guide.md
source_anchor: ""
source_lines: [3257, 3390]
sha256: 26f9a5328fe432d681feecf9fc15ada0d31a160c875fa79854826ae39a457eaa
---

# ── Règles / decoders ─────────────────────────────────────
sudo /var/ossec/bin/wazuh-logtest           # tester une ligne de log
sudo /var/ossec/bin/wazuh-analysisd -t      # tester la conf (avec -t)

# ── Cluster ───────────────────────────────────────────────
sudo /var/ossec/bin/cluster_control -l      # nœuds
sudo /var/ossec/bin/cluster_control -s      # statut synchro

# ── Indexer (remplacer <MDP>) ─────────────────────────────
curl -k -u admin:'<MDP>' 'https://localhost:9200/_cluster/health?pretty'
curl -k -u admin:'<MDP>' 'https://localhost:9200/_cat/indices/wazuh-alerts-*?v&s=index'

# ── API (remplacer <MDP>, TOKEN exporté en variable) ──────
TOKEN=$(curl -sk -u wazuh-wui:'<MDP>' https://localhost:55000/security/user/authenticate \
  | python3 -c "import sys,json; print(json.load(sys.stdin)['data']['token'])")
curl -sk -H "Authorization: Bearer $TOKEN" 'https://localhost:55000/agents/summary/status?pretty=true'

# ── Logs à connaître par cœur ─────────────────────────────
sudo tail -f /var/ossec/logs/ossec.log                 # manager
sudo tail -f /var/ossec/logs/alerts/alerts.json        # alertes JSON
sudo tail -f /var/ossec/logs/active-responses.log      # réponses actives
sudo tail -f /var/ossec/logs/ossec.log                 # agent (sur l'agent)
```

---

## 83. Ports et flux réseau : tableau de référence

| Port | Proto | Source → Destination | Service | Exposer |
|---|---|---|---|---|
| 1514 | UDP (TCP opt.) | Agents → Manager | Remontée événements | Parc supervisé |
| 1515 | TCP | Agents → Manager | Inscription (authd) | Parc (ou via provisioning) |
| 1516 | TCP | Manager ↔ Manager | Cluster | Inter-managers uniquement |
| 443 | TCP | Admins → Dashboard | IHM web | Admins / VPN |
| 9200 | TCP | Manager/Dashboard → Indexer | OpenSearch API | Inter-composants + admins |
| 9300 | TCP | Indexer ↔ Indexer | Cluster OpenSearch | Inter-indexers uniquement |
| 55000 | TCP | Admins/scripts → Manager | API Wazuh | Admins / supervision |
| 514 | UDP | Équipements → Manager/relais | Syslog | Équipements réseau |
| 22 | TCP | Admins → tous | SSH admin | Admins / bastion |

> Principe : **1514/1515** vers le parc, **tout le reste** réservé aux admins et aux interconnexions. Quand on hésite, on restreint.

---

## 84. Fichiers importants : où est quoi

### Manager (`/var/ossec/`)

| Fichier | Rôle |
|---|---|
| `/var/ossec/etc/ossec.conf` | Configuration principale du manager |
| `/var/ossec/etc/rules/local_rules.xml` | **Vos** règles |
| `/var/ossec/etc/rules/soc_tuning.xml` | Vos surcharges (suggéré) |
| `/var/ossec/etc/decoders/local_decoder.xml` | **Vos** decoders |
| `/var/ossec/etc/lists/` | Vos CDB lists |
| `/var/ossec/etc/shared/<groupe>/agent.conf` | Conf poussée aux agents |
| `/var/ossec/etc/client.keys` | Clés des agents |
| `/var/ossec/ruleset/` | Ruleset officiel (**ne pas modifier**) |
| `/var/ossec/logs/ossec.log` | Log principal |
| `/var/ossec/logs/alerts/alerts.json` | Alertes (source de Filebeat) |
| `/var/ossec/logs/active-responses.log` | Réponses actives |
| `/var/ossec/active-response/bin/` | Scripts de réponse |
| `/var/ossec/integrations/` | Scripts d'intégration (TheHive, webhooks) |
| `/var/ossec/api/configuration/api.yaml` | Configuration de l'API |
| `/var/ossec/queue/` | Files d'attente, DB agents |

### Indexer / Dashboard

| Fichier | Rôle |
|---|---|
| `/etc/wazuh-indexer/opensearch.yml` | Configuration de l'indexer |
| `/etc/wazuh-indexer/certs/` | Certificats TLS |
| `/var/lib/wazuh-indexer/` | **Données** (à mettre sur volume dédié) |
| `/etc/wazuh-dashboard/opensearch_dashboards.yml` | Configuration du dashboard |
| `/usr/share/wazuh-dashboard/data/wazuh/config/wazuh.yml` | Connexion dashboard → API manager |

### Agent Linux (`/var/ossec/`)

| Fichier | Rôle |
|---|---|
| `/var/ossec/etc/ossec.conf` | Configuration de l'agent |
| `/var/ossec/etc/client.keys` | Clé d'inscription |
| `/var/ossec/logs/ossec.log` | Log de l'agent |
| `/var/ossec/queue/syscheck/` | Base FIM locale |
| `/var/ossec/queue/sca/` | Résultats SCA |

---

## 85. Glossaire

| Terme | Définition |
|---|---|
| **Active Response** | Exécution automatique d'un script (blocage IP, etc.) sur déclenchement d'une règle. |
| **Agent** | Programme installé sur la machine surveillée ; collecte et transmet. |
| **Alerte** | Événement ayant déclenché une règle de niveau > seuil ; stockée dans `alerts.json`. |
| **API** | Interface REST du manager (port 55000) pour automatiser. |
| **authd** | Service d'inscription des agents (port 1515). |
| **CDB list** | Liste de valeurs (IP, utilisateurs...) utilisable dans les règles. |
| **Cluster** | Ensemble de managers synchronisés (1 master + N workers). |
| **CVE** | Référence publique d'une vulnérabilité (*Common Vulnerabilities and Exposures*). |
| **CVSS** | Score de gravité d'un CVE (0.0 à 10.0). |
| **Dashboard** | Interface web (OpenSearch Dashboards) de visualisation. |
| **Decoder** | Règle de parsing : transforme un log brut en champs structurés. |
| **EPS** | *Events Per Second* : débit d'événements, unité de dimensionnement. |
| **Filebeat** | Agent qui envoie `alerts.json` du manager vers l'indexer. |
| **FIM** | *File Integrity Monitoring* : surveillance des modifications de fichiers (syscheck). |
| **Indexer** | Composant de stockage/recherche basé sur OpenSearch (port 9200). |
| **ISM** | *Index State Management* : politiques de cycle de vie des index (rétention). |
| **MITRE ATT&CK** | Référentiel de tactiques et techniques d'attaque ; Wazuh mappe ses règles dessus. |
| **Manager** | Cœur de Wazuh : analyse (decoders, règles), FIM, SCA, réponses actives. |
| **Niveau (level)** | Sévérité d'une règle, de 0 (ignoré) à 16 (attaque en cours). |
| **Règle** | Condition sur des champs décodés → déclenche une alerte. |
| **SCA** | *Security Configuration Assessment* : audit de conformité (CIS...). |
| **SIEM** | *Security Information and Event Management* : centralisation et corrélation des logs. |
| **SLM** | *Snapshot Lifecycle Management* : snapshots automatiques de l'indexer. |
| **SOAR** | *Security Orchestration, Automation and Response* (ex. Shuffle). |
| **Syscheck** | Nom historique du module FIM dans Wazuh/OSSEC. |
| **Syscollector** | Module d'inventaire (logiciels, matériel, réseau, processus). |
| **Tuning** | Ajustement continu des règles pour réduire les faux positifs. |
| **Vulnerability Detector** | Module croisant l'inventaire logiciel avec les bases CVE. |
| **Wodle** | Module « wodle » : exécution périodique de commandes ou tâches (ex. `command`, `sca`). |
| **WPK** | Paquet Wazuh pour la mise à jour distante des agents. |
| **XDR** | *Extended Detection and Response* : détection + réponse sur les endpoints. |

---

## 86. Quiz : 10 questions + réponses

**Q1. Quels sont les 3 composants centraux de Wazuh et leurs ports par défaut ?**
> R : Le **manager** (1514/UDP agents, 1515/TCP inscription, 55000/TCP API), l'**indexer** (9200/TCP, basé sur OpenSearch) et le **dashboard** (443/TCP, basé sur OpenSearch Dashboards). Filebeat fait le lien manager → indexer.

**Q2. Pourquoi ne faut-il jamais modifier les fichiers de `/var/ossec/ruleset/` ?**
> R : C'est le ruleset officiel, **écrasé à chaque mise à jour**. Le custom va dans `/var/ossec/etc/rules/local_rules.xml` et `/var/ossec/etc/decoders/local_decoder.xml`.

