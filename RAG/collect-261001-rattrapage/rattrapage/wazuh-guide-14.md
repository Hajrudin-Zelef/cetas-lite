---
id: collect-261001-rattrapage/rattrapage/wazuh-guide-14
title: "Guide Wazuh — SIEM & XDR Open Source en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents"]
source: docs/RAG/collect-261001-rattrapage/wazuh_guide.md
source_anchor: ""
source_lines: [2505, 2708]
sha256: 3c545d7ae05e933f2f58b131567815c68acf390294b67b53ae719c86595115ce
---

# Guide Wazuh — SIEM & XDR Open Source en production

| Élément | Chemin | Fréquence | Outil |
|---|---|---|---|
| Conf manager (règles, decoders, ossec.conf) | `/var/ossec/etc/` | Quotidienne | Borg / rsync / git |
| Clés agents, groupes | `/var/ossec/etc/client.keys`, `/var/ossec/queue/` | Quotidienne | Idem |
| Certificats | `/etc/wazuh-indexer/certs/`, etc. | À chaque changement | Coffre chiffré |
| Indexer (données) | Snapshots OpenSearch | Quotidienne | Dépôt S3/NFS (section 63) |
| Base API (utilisateurs) | `/var/ossec/api/configuration/` | Quotidienne | Borg |
| `wazuh-install-files.tar` | Archive initiale | Une fois, hors serveur | Coffre |

### 62.2 Sauvegarde de la conf (exemple Borg)

```bash
# Initialisation (une fois)
sudo borg init --encryption=repokey /backup/wazuh-conf

# Quotidien (cron)
sudo borg create --compression lz4 \
  /backup/wazuh-conf::'{now:%Y-%m-%d}' \
  /var/ossec/etc /var/ossec/api/configuration \
  /var/ossec/integrations /var/ossec/active-response/bin

# Rétention
sudo borg prune --keep-daily=14 --keep-weekly=8 --keep-monthly=12 /backup/wazuh-conf
```

### 62.3 Versionnez vos règles dans git

```bash
# Vos règles/decoders/scripts custom = du code : versionnez !
cd /var/ossec/etc
sudo git init && sudo git add rules/local_rules.xml decoders/local_decoder.xml
sudo git commit -m "soc: règles initiales"
# Puis poussez vers votre Git interne après chaque modification
```

> Une règle qui a « toujours marché » et qui casse un matin : `git log` vous dira qui a changé quoi et quand. Inestimable.

---

## 63. Snapshots de l'indexer : mise en place pas à pas

Les snapshots OpenSearch = sauvegarde cohérente des index, restaurable.

### Étape 1 — Déclarer un dépôt de snapshots

```bash
# Sur chaque nœud indexer, montez le stockage (ex. NFS /mnt/snapshots)
# Puis déclarez le dépôt :
curl -k -u admin:'<MDP>' -X PUT 'https://localhost:9200/_snapshot/depot_quotidien?pretty' \
  -H 'Content-Type: application/json' -d '{
    "type": "fs",
    "settings": { "location": "/mnt/snapshots/wazuh" }
  }'
```

Le chemin doit être dans `path.repo` de `/etc/wazuh-indexer/opensearch.yml` sur **chaque** nœud :

```yaml
path.repo: ["/mnt/snapshots/wazuh"]
```

(Redémarrez l'indexer après modification.)

### Étape 2 — Politique de snapshot (SLM)

```bash
# Snapshot quotidien à 2h, rétention 30 jours
curl -k -u admin:'<MDP>' -X PUT 'https://localhost:9200/_slm/policy/wazuh-quotidien?pretty' \
  -H 'Content-Type: application/json' -d '{
    "schedule": "0 0 2 * * ?",
    "name": "<wazuh-snap-{now/d}>",
    "repository": "depot_quotidien",
    "config": { "indices": "wazuh-alerts-*" },
    "retention": { "expire_after": "30d", "min_count": 7, "max_count": 30 }
  }'
```

### Étape 3 — Vérifier

```bash
curl -k -u admin:'<MDP>' 'https://localhost:9200/_snapshot/depot_quotidien/_all?pretty' | grep -E 'snapshot|state'
curl -k -u admin:'<MDP>' 'https://localhost:9200/_slm/policy/wazuh-quotidien?pretty'
```

> Testez la **restauration** au moins une fois par semestre sur un environnement isolé (section 64). Une sauvegarde jamais testée = pas de sauvegarde.

---

## 64. Restaurer après sinistre : procédure

### Scénario A — Perte du manager (données indexer intactes)

1. Réinstallez un manager (même version !).
2. Restaurez `/var/ossec/etc/` depuis Borg (section 62).
3. Restaurez les certificats si le FQDN est identique.
4. Redémarrez, vérifiez les agents (ils se reconnectent seuls).

### Scénario B — Perte de l'indexer (manager intact)

```bash
# 1. Réinstallez l'indexer (même version)
# 2. Re-déclarez le dépôt de snapshots (section 63, étape 1)
# 3. Listez les snapshots disponibles
curl -k -u admin:'<MDP>' 'https://localhost:9200/_snapshot/depot_quotidien/_all?pretty' | grep '"snapshot"'
# 4. Restaurez (fermez d'abord les index existants si besoin)
curl -k -u admin:'<MDP>' -X POST \
  'https://localhost:9200/_snapshot/depot_quotidien/wazuh-snap-2026.09.25/_restore?pretty' \
  -H 'Content-Type: application/json' -d '{
    "indices": "wazuh-alerts-*",
    "ignore_unavailable": true
  }'
# 5. Vérifiez la santé : _cluster/health doit redevenir green/yellow
```

### Scénario C — Tout est perdu (rebuild complet)

1. Rebuild all-in-one (section 9).
2. Restaurez la conf manager (Borg + git).
3. Restaurez les snapshots indexer.
4. Les agents se réinscrivent ? **Non** : leurs clés sont dans `client.keys` restauré → ils se reconnectent automatiquement. (Si `client.keys` perdu : réinscription massive, prévoyez le mot de passe d'inscription.)

> Tenez cette procédure **imprimée** dans votre classeur d'astreinte. Le jour du sinistre, le dashboard sera peut-être inaccessible : le papier ne plante jamais.

---

## 65. Mise à jour du manager et des composants centraux

### 65.1 Règle d'or

**Jamais de mise à jour sans : sauvegarde (section 62) + lecture du changelog + fenêtre de maintenance.** Les montées de version mineure 4.x sont généralement fluides, mais les majeures (ex. 4.x → 5.x un jour) peuvent changer les formats.

### 65.2 Procédure (paquets Debian/Ubuntu)

```bash
# 1. Sauvegarde complète (conf + snapshot indexer)
# 2. Lire le changelog : https://documentation.wazuh.com/.../release-notes/
# 3. Défiger, mettre à jour dans l'ordre : indexer → manager → dashboard
sudo apt-mark unhold wazuh-indexer wazuh-manager wazuh-dashboard filebeat

sudo apt update
sudo apt install -y wazuh-indexer          # 1. indexer d'abord
sudo systemctl restart wazuh-indexer
# vérifier : curl .../_cluster/health → green/yellow

sudo apt install -y wazuh-manager          # 2. manager
sudo systemctl restart wazuh-manager

sudo apt install -y wazuh-dashboard filebeat   # 3. dashboard + filebeat
sudo systemctl restart wazuh-dashboard filebeat

sudo apt-mark hold wazuh-manager wazuh-indexer wazuh-dashboard filebeat

# 4. Vérifications post-update (section 10) + 1 agent test + 1 alerte test
```

### 65.3 Pièges fréquents

- **Fichiers de conf écrasés** : les paquets peuvent proposer de remplacer `ossec.conf`. Gardez toujours vos versions (Borg/git, section 62) et fusionnez à la main.
- **Ruleset officiel mis à jour** : vos `overwrite` (section 32) peuvent ne plus s'appliquer proprement si la règle officielle a changé → relancez `wazuh-logtest` sur vos échantillons.
- **Incompatibilité de version manager/agents** : un manager 4.12 avec des agents 4.7 fonctionne, mais l'inverse **non**. Mettez toujours le manager à jour **avant** les agents.

---

## 66. Mise à jour des agents (WPK et manuelle)

### 66.1 Via WPK (paquets Wazuh, depuis le manager)

Le manager peut pousser les mises à jour via l'API :

```bash
# 1. Vérifier les versions des agents
curl -sk -H "Authorization: Bearer $TOKEN" \
  'https://localhost:55000/agents?select=id,name,version&limit=500' | grep -o '"version":"[^"]*"'

# 2. Lancer la mise à jour d'un groupe (par vagues !)
curl -sk -X PUT -H "Authorization: Bearer $TOKEN" \
  'https://localhost:55000/agents/upgrade?agents_list=007,008,009'

# 3. Suivre la progression
curl -sk -H "Authorization: Bearer $TOKEN" \
  'https://localhost:55000/agents/upgrade_result?agents_list=007'
```

**Par vagues** : d'abord 2–3 agents pilotes, puis un groupe, puis le reste. Jamais tout le parc d'un coup un vendredi soir.

### 66.2 Manuelle (Ansible / script)

```bash
# Linux
sudo apt-mark unhold wazuh-agent
sudo apt update && sudo apt install -y wazuh-agent
sudo apt-mark hold wazuh-agent
sudo systemctl restart wazuh-agent
```

```powershell
# Windows : poussez le nouveau MSI en silencieux (GPO / Intune / script)
msiexec /i wazuh-agent-4.12-1.msi /q REINSTALL=ALL REINSTALLMODE=vomus
```

---

## 67. Dépannage : méthodologie générale

