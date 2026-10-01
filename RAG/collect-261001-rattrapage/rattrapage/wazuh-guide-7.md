---
id: collect-261001-rattrapage/rattrapage/wazuh-guide-7
title: "Guide Wazuh — SIEM & XDR Open Source en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-26"]
keywords: ["agent", "agents", "apache"]
source: docs/RAG/collect-261001-rattrapage/wazuh_guide.md
source_anchor: ""
source_lines: [910, 1161]
sha256: 4d386d20bdc9cd1ea4715f129567f50beb511c0f107e9ad39009c204b166d299
---

# Guide Wazuh — SIEM & XDR Open Source en production

La clé d'agent chiffre et authentifie le canal agent↔manager. Elle vit dans `/var/ossec/etc/client.keys` (agent) et dans la base du manager.

```bash
# Lister les clés (manager)
sudo /var/ossec/bin/manage_agents -l

# Extraire la clé d'un agent (pour réimport)
sudo /var/ossec/bin/manage_agents -e 007

# Supprimer un agent proprement (puis purger)
sudo /var/ossec/bin/manage_agents -r 007
```

### 21.1 Rotation des clés

Il n'y a pas de rotation automatique native : procédez par vague (groupe par groupe) :

```bash
# 1. Côté manager : supprimer l'ancienne inscription
sudo /var/ossec/bin/manage_agents -r 007
# 2. Côté agent : forcer une réinscription
sudo rm -f /var/ossec/etc/client.keys
sudo /var/ossec/bin/agent-auth -m 10.0.0.21 -P 'NouveauMotDePasse' -G linux-serveurs
sudo systemctl restart wazuh-agent
```

Planifiez une rotation annuelle et **après tout départ d'administrateur** ayant eu accès aux clés.

---

## 22. Groupes d'agents : l'arme de la configuration centralisée

Un groupe = un ensemble d'agents qui partagent des fichiers de configuration poussés par le manager.

```bash
# Créer un groupe (manager)
sudo /var/ossec/bin/agent_groups -a -g linux-web -q

# Assigner un agent à un groupe
sudo /var/ossec/bin/agent_groups -a -g linux-web -i 007 -q

# Lister
sudo /var/ossec/bin/agent_groups -l
```

Groupes conseillés pour démarrer :

| Groupe | Contenu |
|---|---|
| `linux-serveurs` | FIM `/etc`, audit SSH, inventaire |
| `linux-web` | + logs nginx/apache, FIM webroot |
| `windows-postes` | Event logs, FIM dossiers sensibles |
| `windows-serveurs` | + audit avancé, RDP |
| `dmz` | Politique restrictive, active response agressive |

Un agent peut appartenir à **plusieurs groupes** : la configuration est fusionnée (le groupe `default` s'applique toujours en base).

---

## 23. Configuration centralisée : pousser la conf aux agents

Les fichiers partagés vivent dans `/var/ossec/etc/shared/<groupe>/`. Le fichier magique : `agent.conf`.

```bash
sudo mkdir -p /var/ossec/etc/shared/linux-web
sudo nano /var/ossec/etc/shared/linux-web/agent.conf
```

```xml
<agent_config>
  <!-- Surveillance FIM du webroot -->
  <syscheck>
    <frequency>7200</frequency>
    <directories check_all="yes" realtime="yes">/var/www/html</directories>
  </syscheck>

  <!-- Collecte des logs nginx -->
  <localfile>
    <log_format>syslog</log_format>
    <location>/var/log/nginx/access.log</location>
  </localfile>
  <localfile>
    <log_format>syslog</log_format>
    <location>/var/log/nginx/error.log</location>
  </localfile>

  <!-- Inventaire logiciel hebdo -->
  <wodle name="syscollector">
    <disabled>no</disabled>
    <interval>1w</interval>
    <packages>yes</packages>
    <os>yes</os>
  </wodle>
</agent_config>
```

Le manager pousse automatiquement `agent.conf` aux agents du groupe (vérifiez avec `cluster_control` en multi-nœuds). L'agent applique au prochain redémarrage — ou **sans redémarrage** pour certains modules.

> ⚠️ Une erreur XML dans `agent.conf` = **aucun agent du groupe ne l'applique**. Validez toujours avec `/var/ossec/bin/wazuh-logtest` ou `xmllint` avant de pousser.

```bash
xmllint --noout /var/ossec/etc/shared/linux-web/agent.conf && echo "XML OK"
```

---

## 24. Vérifier la santé des agents

### 24.1 En ligne de commande (manager)

```bash
# Vue d'ensemble
sudo /var/ossec/bin/agent_control -l

# Détail d'un agent
sudo /var/ossec/bin/agent_control -i 007

# Agents déconnectés depuis plus de X
sudo /var/ossec/bin/agent_control -l | grep -i disconnected
```

### 24.2 Via l'API

```bash
curl -sk -H "Authorization: Bearer $TOKEN" \
  'https://localhost:55000/agents?status=disconnected&pretty=true'
```

### 24.3 Les 4 états d'un agent

| État | Signification | Action |
|---|---|---|
| `active` | Connecté, envoie des données | Rien |
| `disconnected` | Ne répond plus (seuil : `notify_time` + `time`) | Voir section 68 |
| `pending` | Inscrit, jamais connecté | Vérifier l'inscription / réseau |
| `never_connected` | Inscrit mais aucune connexion historique | Agent jamais démarré ? |

### 24.4 Alerte sur agent silencieux

Wazuh peut alerter quand un agent se déconnecte (règle 501 / niveau 3). Pour une petite équipe, créez une règle surélevée (voir section 29) : un serveur qui cesse d'émettre un vendredi soir mérite un SMS, pas un log.

---

## 25. Désinstaller / réinscrire un agent proprement

```bash
# Linux : désinstallation propre
sudo systemctl stop wazuh-agent
sudo apt remove --purge wazuh-agent -y
sudo rm -rf /var/ossec

# Côté manager : purger l'inscription
sudo /var/ossec/bin/manage_agents -r <ID>

# Réinscription d'une machine réinstallée (évite les doublons) :
# 1. purger l'ancien ID côté manager
# 2. réinstaller l'agent
# 3. agent-auth avec le MÊME nom → nouvel ID, historique propre
```

> Ne réutilisez jamais un ancien `client.keys` sur une machine réinstallée : supprimez et réinscrivez. Les doublons d'ID sont une source classique d'alertes fantômes.

---

## 26. Le moteur de règles : comment Wazuh analyse les logs

Chaîne de traitement d'un événement :

```
Log brut (agent)
   │  1. DECODER : reconnaît le programme/source, extrait les champs
   ▼     (ex. sshd → srcip, dstuser, action)
Champs structurés
   │  2. RÈGLES : les règles sont évaluées par ID croissant ;
   ▼     la première qui matche peut en déclencher d'autres (if_sid)
Alerte (si level > seuil de log)
   │  3. Stockage alerts.json → Filebeat → Indexer → Dashboard
   ▼
Réponse active éventuelle
```

Points clés :

- Les règles sont dans `/var/ossec/ruleset/rules/` (ne **jamais** modifier : écrasé à chaque mise à jour).
- Vos règles vont dans `/var/ossec/etc/rules/local_rules.xml`.
- Une règle ne se déclenche que si le log a été **décodé** (decoder parent OK).
- `wazuh-logtest` (section 31) permet de simuler toute la chaîne.

---

## 27. Decoders : extraire du sens des logs bruts

Un decoder = une expression régulière nommée qui transforme `Failed password for root from 203.0.113.7 port 51234 ssh2` en champs `srcip=203.0.113.7`, `dstuser=root`.

```bash
# Où vivent les decoders
ls /var/ossec/ruleset/decoders/ | head -20
# Vos decoders :
sudo nano /var/ossec/etc/decoders/local_decoder.xml
```

Anatomie d'un decoder :

```xml
<decoder name="mon-app-web">
  <prematch>^monapp: </prematch>          <!-- filtre rapide -->
  <regex>^monapp: user=(\\S+) action=(\\S+) ip=(\\S+)</regex>
  <order>dstuser, action, srcip</order>   <!-- mappe les groupes -->
</decoder>
```

- `<prematch>` : testé en premier, peu coûteux. Mettez-y une chaîne distinctive.
- `<regex>` : groupes capturants dans l'ordre de `<order>`.
- Les champs extraits deviennent utilisables dans les règles (`<field name="srcip">`).

Decoders parents/enfants : un decoder peut affiner un decoder générique via `<parent>`. Exemple : le decoder `sshd` est parent de dizaines de sous-decoders (succès, échec, clé invalide...).

---

## 28. Écrire son premier decoder personnalisé

Cas concret : votre application métier logue dans `/var/log/monapp/app.log` :

```
2026-09-26 22:10:01 [WARN] user=jdupont action=export ip=10.0.5.23 file=clients.csv
```

### Étape 1 — Collecter le log côté agent

```xml
<!-- /var/ossec/etc/shared/mon-groupe/agent.conf -->
<agent_config>
  <localfile>
    <log_format>syslog</log_format>
    <location>/var/log/monapp/app.log</location>
  </localfile>
</agent_config>
```

### Étape 2 — Écrire le decoder (manager)

```xml
<!-- /var/ossec/etc/decoders/local_decoder.xml -->
<decoder name="monapp-base">
  <prematch>user=</prematch>
  <regex>^(\S+ \S+) \[(\w+)\] user=(\S+) action=(\S+) ip=(\S+) file=(\S+)</regex>
  <order>timestamp, loglevel, dstuser, action, srcip, filename</order>
</decoder>
```

### Étape 3 — Tester avec wazuh-logtest

