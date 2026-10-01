---
id: collect-261001-rattrapage/rattrapage/wazuh-guide-12
title: "Guide Wazuh — SIEM & XDR Open Source en production"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: ["2026-09-26"]
keywords: ["open source", "agent", "intel"]
source: docs/RAG/collect-261001-rattrapage/wazuh_guide.md
source_anchor: ""
source_lines: [2049, 2271]
sha256: 3df854cbf93ac91c447a951c0339c8591dcf371c8d72d1a871c9e283bc0ac1f9
---

# Guide Wazuh — SIEM & XDR Open Source en production

## 52. Collecter les logs d'équipements réseau (syslog)

Switchs, pare-feu, **onduleurs**, bornes Wi-Fi : la plupart parlent syslog. Deux options :

### Option A — Syslog direct vers le manager

Sur le manager, activez la réception syslog (port 514) :

```xml
<!-- /var/ossec/etc/ossec.conf du MANAGER -->
<remote>
  <connection>syslog</connection>
  <port>514</port>
  <protocol>udp</protocol>
  <allowed-ips>10.0.0.0/8</allowed-ips>
  <local_ip>10.0.0.21</local_ip>
</remote>
```

```bash
# Autoriser le port
sudo ufw allow 514/udp comment 'Syslog equipements'
sudo systemctl restart wazuh-manager
```

Côté équipement : pointez son syslog vers `10.0.0.21:514`. Chaque équipement apparaît comme un « agent » syslog dans le dashboard.

### Option B — Via un agent relais (recommandé)

Un agent Linux collecte en local (rsyslog) puis envoie au manager via le canal chiffré 1514 :

```bash
# Sur l'agent relais : rsyslog écoute les équipements
# /etc/rsyslog.d/10-equipements.conf
module(load="imudp")
input(type="imudp" port="514" ruleset="tofile")
template(name="EqFmt" type="string" string="%timegenerated% %hostname% %msg%\n")
ruleset(name="tofile") {
  action(type="omfile" file="/var/log/equipements.log" template="EqFmt")
}
```

```xml
<!-- agent.conf : Wazuh lit ce fichier -->
<localfile>
  <log_format>syslog</log_format>
  <location>/var/log/equipements.log</location>
</localfile>
```

> L'option B est **chiffrée et authentifiée** (clé d'agent), contrairement au syslog UDP brut. En environnement sensible, préférez-la. Pensez aussi à vos **onduleurs** : leurs logs (coupures, batteries) méritent d'être centralisés — croisez-les avec NUT (voir guide onduleurs) pour une vision énergie + sécurité.

---

## 53. Surveiller un pare-feu / un switch avec Wazuh

Exemple concret : un pare-feu qui logue en syslog les connexions refusées.

### 53.1 Decoder pour logs de pare-feu générique

```
2026-09-26 22:30:01 fw-dmz kernel: DROP IN=eth0 OUT= SRC=203.0.113.7 DST=10.0.1.10 PROTO=TCP SPT=52341 DPT=22
```

```xml
<!-- /var/ossec/etc/decoders/local_decoder.xml -->
<decoder name="fw-drop">
  <prematch>kernel: DROP</prematch>
  <regex>kernel: DROP .* SRC=(\S+) DST=(\S+) .* PROTO=(\S+) .* SPT=(\d+) DPT=(\d+)</regex>
  <order>srcip, dstip, protocol, srcport, dstport</order>
</decoder>
```

```xml
<!-- Règle : scan de ports détecté par le pare-feu -->
<rule id="100700" level="10" frequency="20" timeframe="60">
  <decoded_as>fw-drop</decoded_as>
  <description>Scan probable : 20 paquets droppés en 60s depuis $(srcip)</description>
  <mitre><id>T1595</id></mitre>
</rule>
```

### 53.2 Ce que Wazuh sait déjà décoder

Le ruleset officiel inclut des decoders pour : Cisco IOS/ASA, pfSense/OPNsense, iptables, Windows Firewall, Juniper, Fortinet (partiel). Avant d'écrire un decoder, cherchez :

```bash
ls /var/ossec/ruleset/decoders/ | grep -iE 'cisco|forti|pfsense|iptables'
sudo /var/ossec/bin/wazuh-logtest   # collez une ligne réelle, voyez si ça matche déjà
```

> Pour vos équipements MikroTik / Huawei / Cisco du lab RAG : la plupart émettent du syslog standard. Un decoder `prematch` sur le hostname + 2–3 règles suffisent souvent pour les événements critiques (login, changement de config, lien down).

---

## 54. Intégration TheHive : créer des cases depuis les alertes

**TheHive** = plateforme de gestion d'incidents (open source). L'idée : toute alerte Wazuh niveau ≥ 12 crée automatiquement une *case* dans TheHive pour suivi.

### 54.1 Principe

```
Wazuh (alerte) → intégration custom (script) → API TheHive → Case créée
```

Wazuh appelle un script d'intégration via `<integration>` dans `ossec.conf` :

```xml
<!-- /var/ossec/etc/ossec.conf du MANAGER -->
<integration>
  <name>custom-thehive.py</name>
  <rule_id>100100,100300,100700</rule_id>
  <alert_format>json</alert_format>
</integration>
<!-- ou : <level>12</level> pour tout ce qui est ≥ 12 -->
```

### 54.2 Script d'intégration (squelette)

```python
#!/usr/bin/env python3
# /var/ossec/integrations/custom-thehive.py
import sys, json, requests

THEHIVE_URL = "https://thehive.exemple.lan:9000"
API_KEY = "VOTRE_CLE_API_THEHIVE"   # via variable d'environnement en prod !

def main():
    alert = json.load(sys.stdin)
    rule = alert.get("rule", {})
    case = {
        "title": f"[Wazuh] {rule.get('description', 'Alerte')}",
        "description": json.dumps(alert, indent=2)[:5000],
        "severity": 3 if int(rule.get("level", 0)) >= 12 else 2,
        "tags": ["wazuh", f"rule:{rule.get('id')}", f"agent:{alert.get('agent', {}).get('name')}"],
        "tlp": 2,
        "pap": 2,
    }
    r = requests.post(f"{THEHIVE_URL}/api/v1/case", json=case,
                      headers={"Authorization": f"Bearer {API_KEY}"},
                      timeout=15, verify=False)  # verify=False seulement avec autosigné interne
    r.raise_for_status()
    print(f"Case créée : {r.json().get('id')}")

if __name__ == "__main__":
    # Wazuh appelle : custom-thehive.py <fichier_alerte_json>
    sys.stdin = open(sys.argv[1])
    main()
```

```bash
sudo chmod 750 /var/ossec/integrations/custom-thehive.py
sudo chown root:wazuh /var/ossec/integrations/custom-thehive.py
```

> ⚠️ Ne laissez jamais une clé API en dur dans un script versionné : utilisez une variable d'environnement ou un fichier de secrets (droits 600).

### 54.3 Bonnes pratiques

- Ne créez des cases que pour les niveaux **≥ 12** (sinon TheHive devient une poubelle).
- Dédupliquez : une case par (règle, agent, jour), pas une par alerte.
- Enrichissez avec le lien direct vers l'alerte dans le dashboard Wazuh.

---

## 55. Intégration Shuffle (SOAR)

**Shuffle** = plateforme SOAR open source (automatisation de réponse). Cas d'usage : à la réception d'une alerte, Shuffle peut automatiquement : interroger un threat intel, isoler la machine (via API Wazuh active-response), créer un ticket, notifier.

### 55.1 Webhook Wazuh → Shuffle

Côté Wazuh, utilisez l'intégrateur `slack`-like générique ou un script custom (même mécanisme que section 54) qui POSTe vers le webhook d'entrée Shuffle :

```python
# extrait : envoi vers Shuffle
SHUFFLE_WEBHOOK = "https://shuffle.exemple.lan:3001/api/v1/hooks/<ID_DU_WORKFLOW>"
payload = {
    "rule_id": rule.get("id"),
    "level": rule.get("level"),
    "agent": alert.get("agent", {}).get("name"),
    "srcip": alert.get("data", {}).get("srcip"),
    "description": rule.get("description"),
}
requests.post(SHUFFLE_WEBHOOK, json=payload, timeout=10)
```

### 55.2 Workflow Shuffle type « force brute SSH »

1. **Trigger** : webhook reçoit l'alerte (règle 5763).
2. **Enrichissement** : appel API AbuseIPDB / VirusTotal sur la `srcip`.
3. **Décision** : si score malveillant > seuil → étape 4, sinon log seul.
4. **Action** : appel API Wazuh pour lancer `firewall-drop` (voir section 58), + création d'un ticket.
5. **Notification** : résumé sur le canal SOC.

> Commencez **en mode « dry-run »** (Shuffle notifie sans bloquer) pendant 2 semaines. Passez en blocage auto quand le taux de faux positifs est < 5 %.

---

## 56. Webhooks génériques : Slack, Teams, Mattermost

### 56.1 Intégrateur Slack natif

```xml
<!-- /var/ossec/etc/ossec.conf du MANAGER -->
<integration>
  <name>slack</name>
  <hook_url>https://hooks.slack.com/services/T000/B000/XXXXXXXX</hook_url>
  <rule_id>5763</rule_id>
  <alert_format>json</alert_format>
</integration>
```

### 56.2 Webhook générique (Teams, Mattermost, n8n...)

Pas d'intégrateur natif → script custom (mécanisme `<integration>` + script dans `/var/ossec/integrations/`) :

```python
#!/usr/bin/env python3
# /var/ossec/integrations/custom-webhook.py
import sys, json, requests

WEBHOOK_URL = "https://mattermost.exemple.lan/hooks/xxx"  # ou Teams, n8n...

