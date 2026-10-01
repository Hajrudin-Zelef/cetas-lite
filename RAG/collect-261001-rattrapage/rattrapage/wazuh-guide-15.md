---
id: collect-261001-rattrapage/rattrapage/wazuh-guide-15
title: "Guide Wazuh — SIEM & XDR Open Source en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents"]
source: docs/RAG/collect-261001-rattrapage/wazuh_guide.md
source_anchor: ""
source_lines: [2709, 2894]
sha256: e451a6115c655c26442aa65ea4a6c02dfeb3cd365f4886ca701f7465edb2f247
---

# Guide Wazuh — SIEM & XDR Open Source en production

```
1. DÉFINIR : quel est le symptôme exact ? (pas d'alerte ? agent rouge ? dashboard vide ?)
2. LOCALISER : à quelle étape de la chaîne ça casse ?
   Agent → (1514) → Manager (analysisd) → alerts.json → Filebeat → (9200) → Indexer → Dashboard
3. VÉRIFIER : chaque maillon avec UNE commande (tableau ci-dessous)
4. CORRIGER : une seule chose à la fois, puis re-tester
5. DOCUMENTER : notez la cause dans votre wiki (section 74)
```

| Maillon | Commande de test |
|---|---|
| Agent → Manager | `agent_control -i <ID>` : `active` ? |
| Manager analyse | `tail /var/ossec/logs/ossec.log` : erreurs ? |
| alerts.json s'écrit | `tail -5 /var/ossec/logs/alerts/alerts.json` |
| Filebeat envoie | `systemctl status filebeat`, `/var/log/filebeat/` |
| Indexer reçoit | `curl .../_cat/indices/wazuh-alerts-*?v` : docs.count augmente ? |
| Dashboard affiche | Requête *Discover* sur les 15 dernières minutes |

> 80 % des pannes se résolvent à l'étape 2 : **la chaîne est linéaire, suivez-la dans l'ordre** au lieu de tout redémarrer en vrac.

---

## 68. Dépannage : l'agent ne remonte pas

### Arbre de décision

```
Agent "disconnected" ou "never_connected" ?
│
├─► 1. L'agent tourne-t-il ?
│     Linux   : systemctl is-active wazuh-agent
│     Windows : Get-Service wazuh-agent
│     Non → démarrez-le, lisez ossec.log
│
├─► 2. L'agent est-il inscrit ? (clé présente ?)
│     Linux   : ls -la /var/ossec/etc/client.keys
│     Absent → (ré)inscription : agent-auth (section 20)
│
├─► 3. Le réseau passe-t-il ? (depuis l'AGENT)
│     Test UDP 1514 : nc -u -v -z <manager> 1514
│     Test TCP 1515 : nc -v -z <manager> 1515
│     Échec → pare-feu local, pare-feu réseau, routage, NAT
│
├─► 4. Le manager voit-il l'agent ?
│     /var/ossec/bin/agent_control -i <ID>
│     "Active" côté manager mais rien ne remonte → vérifiez l'heure (NTP !)
│
└─► 5. L'horloge est-elle synchro ? (± 5 min max)
      timedatectl status   (agent ET manager)
      Décalage → les messages sont rejetés silencieusement
```

### Cas vécu : l'agent derrière un NAT

Si plusieurs agents partagent la même IP publique, utilisez `use_source_ip=no` côté `authd` et des noms d'agents uniques. Sinon le manager confond les agents.

### Cas vécu : l'agent Windows « actif » mais sans logs

Souvent le service tourne mais l'inscription a échoué silencieusement : supprimez `C:\Program Files (x86)\ossec-agent\client.keys`, réinscrivez via l'interface (*Manage → Authenticate*), redémarrez le service.

---

## 69. Dépannage : indexer rouge / cluster unhealthy

```bash
# 1. Diagnostic express
curl -k -u admin:'<MDP>' 'https://localhost:9200/_cluster/health?pretty'
curl -k -u admin:'<MDP>' 'https://localhost:9200/_cat/shards?v' | grep -v STARTED
curl -k -u admin:'<MDP>' 'https://localhost:9200/_cat/allocation?v'
```

| Symptôme | Cause probable | Action |
|---|---|---|
| `status: red` | Shard primaire non assigné | Identifier le shard (`_cat/shards`), vérifier le nœud qui le portait |
| `unassigned_shards > 0` durable | Nœud manquant / disque plein | Libérer du disque (section 72), redémarrer le nœud |
| `number_of_nodes` < attendu | Nœud indexer down | `systemctl status wazuh-indexer`, logs `/var/log/wazuh-indexer/` |
| Erreurs `flood stage` | Disque > 95 % | **Urgence** : supprimer des vieux index (section 72) |

### Sortir du flood stage (disque plein côté indexer)

```bash
# 1. Libérer de la place EN DEHORS d'OpenSearch d'abord (logs, tmp)
# 2. Lever le bloc en lecture seule :
curl -k -u admin:'<MDP>' -X PUT 'https://localhost:9200/_all/_settings' \
  -H 'Content-Type: application/json' \
  -d '{"index.blocks.read_only_allow_delete": null}'
# 3. Supprimer les vieux index :
curl -k -u admin:'<MDP>' -X DELETE 'https://localhost:9200/wazuh-alerts-4.x-2026.06.*'
# 4. Vérifier : _cluster/health → yellow/green
```

> Un indexer `red` = **priorité absolue**, devant toute analyse d'alerte : sans stockage, le SIEM est aveugle ET perd des données.

---

## 70. Dépannage : pas d'alertes / alertes en retard

```
Le dashboard est vide mais les agents sont verts ?
│
├─► 1. alerts.json se remplit-il ? (MANAGER)
│     tail -5 /var/ossec/logs/alerts/alerts.json
│     Non → analysisd ne produit pas : regardez /var/ossec/logs/ossec.log
│           (erreur de règle ? testez avec wazuh-logtest, section 31)
│
├─► 2. Filebeat tourne-t-il et envoie-t-il ?
│     systemctl is-active filebeat
│     tail -20 /var/log/filebeat/filebeat
│     Erreurs TLS → certificats (section 16) ; erreurs 401 → mot de passe
│
├─► 3. L'indexer reçoit-il ? (doc count augmente ?)
│     watch -n 5 'curl -sk -u admin:... "https://localhost:9200/_cat/indices/wazuh-alerts-*?v&s=index" | tail -3'
│
├─► 4. Le dashboard interroge-t-il le bon index ?
│     Discover → vérifiez le pattern d'index et la plage de temps !
│     (90 % des "dashboard vide" = plage de temps sur "aujourd'hui" alors
│      que les données datent d'hier, ou pattern d'index cassé après update)
│
└─► 5. Retard seulement ? → file d'attente analysisd saturée (section 59)
      = manager sous-dimensionné : augmentez CPU/RAM ou passez en distribué
```

---

## 71. Dépannage : dashboard inaccessible

| Symptôme | Vérification | Solution |
|---|---|---|
| `Connection refused` sur 443 | `systemctl status wazuh-dashboard` | Redémarrer ; voir logs `/var/log/wazuh-dashboard/` |
| Erreur 500 / « Cannot connect to OpenSearch » | `curl -k https://localhost:9200` | Indexer down ? (section 69) ; mot de passe `kibanaserver` désynchronisé ? |
| Page blanche après mise à jour | Cache navigateur | Ctrl+Shift+R ; vider le cache ; tester en navigation privée |
| Certificat invalide | `openssl s_client` (section 16) | Certificat expiré → régénérer |
| Login rejeté en boucle | Utilisateur verrouillé ? | Vérifier via API ou réinitialiser avec `securityadmin.sh` |

Réinitialiser le mot de passe admin (indexer) en dernier recours :

```bash
# Générer le hash du nouveau mot de passe puis l'appliquer via securityadmin.sh
# (procédure détaillée dans la doc officielle "password management")
# Pensez ensuite à mettre à jour wazuh.yml côté dashboard (utilisateur wazuh-wui).
```

---

## 72. Dépannage : disque plein

### Où chercher (par ordre de probabilité)

```bash
df -h
du -sh /var/lib/wazuh-indexer /var/ossec/logs /var/log 2>/dev/null
# Index les plus gros :
curl -k -u admin:'<MDP>' 'https://localhost:9200/_cat/indices?v&s=store.size:desc' | head -15
```

### Actions immédiates

```bash
# 1. Purger les vieux logs manager (gardez 30 jours en local, le reste en archives)
sudo find /var/ossec/logs/alerts -name 'alerts.log-*' -mtime +30 -delete
sudo journalctl --vacuum-time=14d

# 2. Supprimer les vieux index (en respectant la rétention légale !)
curl -k -u admin:'<MDP>' -X DELETE 'https://localhost:9200/wazuh-alerts-4.x-2026.05.*'

# 3. Vérifier la rotation ISM (section 60) : est-elle vraiment active ?
```

### Prévention durable

- Alertes à 75 % / 85 % (section 59).
- Rétention ISM alignée sur vos obligations (souvent 1 an → snapshots froids, pas tout en ligne).
- **Séparez les volumes** : données indexer ≠ système ≠ logs.

---

## 73. 10 erreurs classiques (et comment les éviter)

### Erreur n°1 — Modifier le ruleset officiel au lieu de local_rules.xml
**Symptôme :** vos règles disparaissent à la mise à jour.
**Solution :** tout le custom va dans `/var/ossec/etc/rules/`, `/var/ossec/etc/decoders/`. Le dossier `/var/ossec/ruleset/` est en lecture seule mentale.

### Erreur n°2 — Oublier la allowlist avant d'activer l'active response
**Symptôme :** vous vous bannissez vous-même (ou votre supervision) en testant.
**Solution :** `<white_list>` complète (section 45) **avant** le premier `<active-response>`. Testez depuis une IP de lab non allowlistée.

