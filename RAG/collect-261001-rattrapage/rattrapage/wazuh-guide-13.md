---
id: collect-261001-rattrapage/rattrapage/wazuh-guide-13
title: "Guide Wazuh — SIEM & XDR Open Source en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents"]
source: docs/RAG/collect-261001-rattrapage/wazuh_guide.md
source_anchor: ""
source_lines: [2272, 2504]
sha256: acd72c4d5ac04901c0dfbcfe56f44bcd1accdebf8fe8793b8853f7c8b77ccc95
---

# Guide Wazuh — SIEM & XDR Open Source en production

alert = json.load(open(sys.argv[1]))
rule = alert.get("rule", {})
agent = alert.get("agent", {}).get("name", "?")
text = (
    f"**[Wazuh] Niveau {rule.get('level')}** — {rule.get('description')}\n"
    f"Agent : `{agent}` | Règle : `{rule.get('id')}` | Heure : `{alert.get('timestamp')}`"
)
# Format Mattermost/Slack :
requests.post(WEBHOOK_URL, json={"text": text}, timeout=10)
# Pour Teams : adaptez au format MessageCard / Adaptive Card.
```

> Filtrez par `<level>` ou `<rule_id>` : envoyer **toutes** les alertes sur un canal = le meilleur moyen que personne ne le lise. Niveau ≥ 10 uniquement.

---

## 57. Envoyer des alertes par e-mail

```xml
<!-- /var/ossec/etc/ossec.conf du MANAGER -->
<global>
  <email_notification>yes</email_notification>
  <email_to>soc@exemple.lan</email_to>
  <smtp_server>10.0.0.25</smtp_server>
  <email_from>wazuh@exemple.lan</email_from>
  <email_maxperhour>30</email_maxperhour>   <!-- anti-spam -->
</global>

<alerts>
  <log_alert_level>3</log_alert_level>
  <email_alert_level>10</email_alert_level>  <!-- e-mail seulement si ≥ 10 -->
</alerts>
```

Avec relais SMTP authentifié (ex. Postfix interne) :

```xml
<global>
  <smtp_server>10.0.0.25</smtp_server>
  <!-- Pour de l'auth SMTP, passez par un relais local postfix plutôt que Wazuh -->
</global>
```

> Astuce : faites relayer par votre **Postfix local** (déjà configuré, voir guide Debian/Ubuntu) plutôt que de configurer SMTP+TLS dans Wazuh. `email_maxperhour` évite de noyer la boîte SOC lors d'une tempête d'alertes.

Modèle d'e-mail utile : objet avec niveau + règle + agent. Wazuh le fait nativement ; vérifiez que l'objet est filtrable (règle de tri côté boîte SOC : `[Wazuh]` + niveau).

---

## 58. API Wazuh : automatiser sans le dashboard

L'API (port 55000) permet tout ce que fait le dashboard, en script.

### 58.1 Authentification

```bash
# Token (valide 15 min par défaut)
TOKEN=$(curl -sk -u wazuh-wui:'<MOT_DE_PASSE>' \
  https://localhost:55000/security/user/authenticate \
  | python3 -c "import sys,json; print(json.load(sys.stdin)['data']['token'])")
echo "Token OK : ${TOKEN:0:20}..."
```

### 58.2 Opérations courantes

```bash
# Lister les agents actifs
curl -sk -H "Authorization: Bearer $TOKEN" \
  'https://localhost:55000/agents?status=active&select=id,name,ip&limit=500'

# Redémarrer un groupe d'agents
curl -sk -X PUT -H "Authorization: Bearer $TOKEN" \
  'https://localhost:55000/agents/group/linux-web/restart'

# Lancer une active-response manuelle (bannir une IP)
curl -sk -X PUT -H "Authorization: Bearer $TOKEN" \
  -H 'Content-Type: application/json' \
  'https://localhost:55000/active-response?agents_list=007' \
  -d '{"command": "firewall-drop.sh", "arguments": ["-add", "203.0.113.7"]}'

# Récupérer les règles (pour audit)
curl -sk -H "Authorization: Bearer $TOKEN" \
  'https://localhost:55000/rules?select=id,description,level&limit=100'

# Statistiques du manager (remontées, files)
curl -sk -H "Authorization: Bearer $TOKEN" \
  'https://localhost:55000/manager/stats/analysisd?pretty=true'
```

### 58.3 Créer un utilisateur API dédié (au lieu de wazuh-wui)

```bash
# Créer l'utilisateur + rôle lecture seule (exemple)
curl -sk -X POST -H "Authorization: Bearer $TOKEN" \
  -H 'Content-Type: application/json' \
  https://localhost:55000/security/users \
  -d '{"username": "api-lecture", "password": "<MOT_DE_PASSE_FORT>"}'
```

> Un script de supervision (Zabbix, Prometheus exporter) doit utiliser un compte **dédié, en lecture seule**, jamais `wazuh-wui`.

---

## 59. Superviser Wazuh lui-même : le manager en bonne santé

Qui surveille le surveillant ? **Vous**, avec ces indicateurs.

### 59.1 Files d'attente (le pouls du manager)

```bash
# Taille des files analysisd / remoted
sudo /var/ossec/bin/wazuh-analysisd -t   # test de conf (ne pas confondre)
ls -la /var/ossec/queue/
sudo cat /var/ossec/var/run/wazuh-analysisd.state 2>/dev/null | head -30
```

Via API :

```bash
curl -sk -H "Authorization: Bearer $TOKEN" \
  'https://localhost:55000/manager/stats/analysisd?pretty=true'
# Surveillez : events_processed, events_dropped, queue_size
```

Seuils d'alerte (à mettre dans Zabbix/Prometheus) :

| Métrique | Warning | Critical |
|---|---|---|
| `queue_size` analysisd | > 1000 | > 10000 |
| `events_dropped` (croissant) | > 0 sur 1 h | > 100 sur 1 h |
| Agents `disconnected` | > 5 % du parc | > 15 % |
| Espace `/var/ossec` | > 75 % | > 90 % |

### 59.2 EPS réel : mesurez votre charge

```bash
# Compter les alertes des dernières 24h (ordre de grandeur)
sudo wc -l /var/ossec/logs/alerts/alerts.json
# Mieux : via l'API stats, events_processed sur 1h / 3600 = EPS moyen
```

### 59.3 Logs à surveiller

| Log | Ce qu'il dit |
|---|---|
| `/var/ossec/logs/ossec.log` | Manager : erreurs de règles, authd, cluster |
| `/var/ossec/logs/active-responses.log` | Chaque réponse active exécutée |
| `/var/ossec/logs/alerts/alerts.log` | Alertes texte (debug) |
| `/var/log/filebeat/filebeat` | Envoi vers l'indexer |

> Mettez **Wazuh lui-même sous supervision** : un check « le manager tourne + Filebeat envoie + l'indexer est vert » dans votre Zabbix, avec alerte SMS. Un SIEM en panne silencieuse = pire que pas de SIEM.

---

## 60. Superviser l'indexer (OpenSearch)

### 60.1 Santé du cluster

```bash
curl -k -u admin:'<MDP>' 'https://localhost:9200/_cluster/health?pretty'
```

| Statut | Signification | Action |
|---|---|---|
| `green` | Tout va bien | Rien |
| `yellow` | Répliques non assignées (**normal en mono-nœud**) | Rien (ou passer en multi-nœuds) |
| `red` | Des *primaries* manquent = **perte de données possible** | Section 69, immédiat |

### 60.2 Espace disque et watermark

OpenSearch bloque l'écriture à 95 % d'usage disque (flood stage) :

```bash
curl -k -u admin:'<MDP>' 'https://localhost:9200/_cat/allocation?v'
# Surveillez disk.used_percent par nœud
```

Seuils : alerte à **75 %**, action à **85 %** (suppression/rotation d'index, voir section 72).

### 60.3 Rotation des index (ISM)

Wazuh crée un index par jour (`wazuh-alerts-4.x-YYYY.MM.DD`). Sans rotation, le disque se remplit. Via le dashboard (*Index Management → ISM policies*) ou API :

```json
// Politique : supprimer les index d'alertes après 90 jours
{
  "policy": {
    "description": "Rétention alertes 90 jours",
    "default_state": "hot",
    "states": [
      { "name": "hot", "transitions": [{ "state_name": "delete", "conditions": { "min_index_age": "90d" } }] },
      { "name": "delete", "actions": [{ "delete": {} }] }
    ]
  }
}
```

> ⚠️ Avant de supprimer : vérifiez vos **obligations légales** de conservation (souvent 1 an pour les logs de sécurité en entreprise). 90 jours en ligne + archives froides (snapshots, section 63) = bon compromis.

---

## 61. Superviser le dashboard

Le dashboard est stateless : s'il tombe, les données sont sauves, mais le SOC est aveugle.

```bash
# Le service tourne ?
sudo systemctl is-active wazuh-dashboard
# Il répond ?
curl -k -s -o /dev/null -w '%{http_code}\n' https://localhost:443
# Attendu : 302 (redirection vers /app/login) ou 200
```

Points de vigilance :

- **Mot de passe `kibanaserver`** : si changé côté indexer sans mise à jour dans `/etc/wazuh-dashboard/opensearch_dashboards.yml`, le dashboard affiche des erreurs 500.
- **Espace disque** : le dashboard écrit peu, mais vérifiez quand même.
- **Sessions** : en cas de comportement étrange après mise à jour, videz le cache navigateur + redémarrez le service.

```yaml
# /etc/wazuh-dashboard/opensearch_dashboards.yml (extrait)
opensearch.hosts: ["https://127.0.0.1:9200"]
opensearch.username: kibanaserver
opensearch.password: "<MOT_DE_PASSE>"
server.ssl.enabled: true
```

---

## 62. Sauvegarde : stratégie complète

### 62.1 Quoi sauvegarder

