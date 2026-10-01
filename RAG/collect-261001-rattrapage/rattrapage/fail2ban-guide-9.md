---
id: collect-261001-rattrapage/rattrapage/fail2ban-guide-9
title: "Guide fail2ban — Le bouclier anti-brute-force"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "arr"]
source: docs/RAG/collect-261001-rattrapage/fail2ban_guide.md
source_anchor: ""
source_lines: [1966, 2192]
sha256: 68e612fba840fcccad5ed8617ac5b116ffb968efd0097c7bca6412c84dc7d9cc
---

# Guide fail2ban — Le bouclier anti-brute-force

## 33. AbuseIPDB : remonter les attaquants

AbuseIPDB est une base communautaire d'IP malveillantes. Deux usages : **consulter** (est-ce un attaquant connu ?) et **contribuer** (remonter vos attaquants).

### Consulter (enquête manuelle)

```bash
# Via l'API (clé gratuite : 1000 requêtes/jour)
curl -s -G https://api.abuseipdb.com/api/v2/check \
  -d ipAddress=203.0.113.45 \
  -d maxAgeInDays=90 \
  -H "Key: VOTRE_CLE_API" \
  -H "Accept: application/json" | python3 -m json.tool
```

Réponse (extraits) : `abuseConfidenceScore`, `totalReports`, `lastReportedAt`, pays, FAI.

### Remonter automatiquement (action)

Créez `/etc/fail2ban/action.d/abuseipdb.conf` :

```ini
[Definition]
actionban = curl -s --fail -X POST https://api.abuseipdb.com/api/v2/report \
              -H "Key: <abuseipdb_key>" -H "Accept: application/json" \
              --data-urlencode "ip=<ip>" \
              --data-urlencode "categories=18,22" \
              --data-urlencode "comment=fail2ban [<name>] : <failures> echecs" \
              >/dev/null
actionunban =

[Init]
abuseipdb_key = VOTRE_CLE_API
```

```ini
[sshd]
action = %(action_)s
         abuseipdb[name=sshd, abuseipdb_key=VOTRE_CLE_API]
```

Catégories AbuseIPDB utiles : `18` = brute-force, `22` = SSH, `21` = web app attack, `14` = port scan.

### Bonnes pratiques de remontée

- Ne remontez que les bans **automatiques** des jails fiables (sshd, nginx-botsearch), pas les bans manuels.
- Ne remontez pas les IP en `ignoreip` (évident, mais l'action ne le sait pas — filtrez en amont si besoin).
- Une remontée par ban suffit : AbuseIPDB déduplique, mais n'abusez pas du quota gratuit.

### Enrichir les alertes mail avec AbuseIPDB

Dans vos mails d'alerte (section 22), ajoutez le lien direct :

```
https://www.abuseipdb.com/check/203.0.113.45
```

Un clic suffit pour voir l'historique de l'IP avant de décider d'un ban permanent.

---

## 34. Monitoring Zabbix

### Ce qu'il faut superviser

| Métrique | Seuil d'alerte | Signification |
|---|---|---|
| fail2ban-server en cours d'exécution | process absent → critique | le service est tombé |
| Nombre de jails actifs | < attendu → warning | un jail ne s'est pas chargé |
| IP bannies par jail | pic brutal → warning | attaque en cours ou faux positifs |
| Total bans / heure | > 100/h → warning | campagne de scan massive |
| Âge de `/var/log/fail2ban.log` | pas modifié depuis 1h → info | peut-être normal (nuit calme) |

### UserParameter Zabbix

Sur l'agent Zabbix (`/etc/zabbix/zabbix_agentd.d/fail2ban.conf`) :

```ini
# Nombre de jails
UserParameter=fail2ban.jails,sudo /usr/bin/fail2ban-client status 2>/dev/null | grep "Number of jail" | awk '{print $NF}'

# IP actuellement bannies pour un jail donné : fail2ban.banned[sshd]
UserParameter=fail2ban.banned[*],sudo /usr/bin/fail2ban-client status "$1" 2>/dev/null | grep "Currently banned" | awk '{print $NF}'

# Total des bans pour un jail : fail2ban.total[sshd]
UserParameter=fail2ban.total[*],sudo /usr/bin/fail2ban-client status "$1" 2>/dev/null | grep "Total banned" | awk '{print $NF}'

# Échecs en cours (compteur non encore banni) : fail2ban.failed[sshd]
UserParameter=fail2ban.failed[*],sudo /usr/bin/fail2ban-client status "$1" 2>/dev/null | grep "Currently failed" | awk '{print $NF}'
```

Droits sudo pour l'utilisateur `zabbix` (via `visudo`) :

```
zabbix ALL=(ALL) NOPASSWD: /usr/bin/fail2ban-client status *
```

> Restreignez au strict nécessaire : `status` seul ne modifie rien. Ne donnez JAMAIS `fail2ban-client set *` ou `reload` à l'agent.

### Template Zabbix (idée de structure)

- Items : `fail2ban.jails`, `fail2ban.banned[sshd]`, `fail2ban.banned[recidive]`, `fail2ban.total[sshd]`.
- Triggers :
  - `{host:fail2ban.jails.last()}<5` → « fail2ban : jail manquant » (severity: average)
  - `{host:fail2ban.banned[recidive].last()}>50` → « fail2ban : nombreux récidivistes » (severity: warning)
  - `nodata(/host/proc.num[fail2ban-server],5m)=1` → « fail2ban-server arrêté » (severity: high)

---

## 35. Monitoring Prometheus

### Exporter les métriques (script simple)

Pas d'exporter officiel maintenu universellement ; le plus fiable en prod est un petit script + node_exporter `textfile`.

`/usr/local/bin/fail2ban_metrics.sh` :

```bash
#!/bin/bash
# Exporte les métriques fail2ban pour node_exporter (textfile)
OUT=/var/lib/node_exporter/textfile_collector/fail2ban.prom
TMP="$OUT.$$"

{
  echo "# HELP fail2ban_jails_total Nombre de jails actifs"
  echo "# TYPE fail2ban_jails_total gauge"
  JAILS=$(fail2ban-client status 2>/dev/null | grep "Number of jail" | awk '{print $NF}')
  echo "fail2ban_jails_total ${JAILS:-0}"

  echo "# HELP fail2ban_banned_current IP actuellement bannies par jail"
  echo "# TYPE fail2ban_banned_current gauge"
  echo "# HELP fail2ban_banned_total Total des bans par jail"
  echo "# TYPE fail2ban_banned_total counter"
  echo "# HELP fail2ban_failed_current Echecs en cours par jail"
  echo "# TYPE fail2ban_failed_current gauge"

  for j in $(fail2ban-client status 2>/dev/null | grep "Jail list" | sed 's/.*Jail list://;s/,//g'); do
    ST=$(fail2ban-client status "$j" 2>/dev/null)
    CUR=$(echo "$ST" | grep "Currently banned" | awk '{print $NF}')
    TOT=$(echo "$ST" | grep "Total banned" | awk '{print $NF}')
    FAIL=$(echo "$ST" | grep "Currently failed" | awk '{print $NF}')
    echo "fail2ban_banned_current{jail=\"$j\"} ${CUR:-0}"
    echo "fail2ban_banned_total{jail=\"$j\"} ${TOT:-0}"
    echo "fail2ban_failed_current{jail=\"$j\"} ${FAIL:-0}"
  done
} > "$TMP" && mv "$TMP" "$OUT"
```

Cron (toutes les 2 minutes) :

```cron
*/2 * * * * root /usr/local/bin/fail2ban_metrics.sh
```

### Alertes Prometheus (exemples)

```yaml
groups:
  - name: fail2ban
    rules:
      - alert: Fail2BanDown
        expr: absent(fail2ban_jails_total) or fail2ban_jails_total == 0
        for: 5m
        labels: { severity: critical }
        annotations:
          summary: "fail2ban ne tourne plus sur {{ $labels.instance }}"

      - alert: Fail2BanJailMissing
        expr: fail2ban_jails_total < 5
        for: 10m
        labels: { severity: warning }
        annotations:
          summary: "Jails manquants sur {{ $labels.instance }} ({{ $value }} actifs)"

      - alert: Fail2BanBanStorm
        expr: increase(fail2ban_banned_total[1h]) > 200
        for: 15m
        labels: { severity: warning }
        annotations:
          summary: "Tempête de bans sur {{ $labels.instance }} : {{ $value }} bans/h"
```

### Dashboard Grafana (panneaux suggérés)

1. `fail2ban_jails_total` — stat unique (vert si == attendu).
2. `fail2ban_banned_current` par jail — bar gauge.
3. `increase(fail2ban_banned_total[1h])` par jail — time series (courbe des attaques).
4. `fail2ban_failed_current{jail="sshd"}` — les attaques en cours non encore bannies.

---

## 36. Alertes mail et webhooks

### Prérequis : un MTA qui envoie

```bash
# Test d'envoi (postfix en satellite, exim, msmtp...)
echo "corps du message" | mail -s "[test] fail2ban" admin@example.com
```

Si le mail n'arrive pas, le problème est le MTA, pas fail2ban. Débuggez d'abord `mail.log`.

### Alerte mail par jail

```ini
[sshd]
enabled = true
action  = %(action_mwl)s[name=sshd, dest=admin@example.com, sender=fail2ban@srv-web.example.com]
```

Contenu du mail (`action_mwl`) : IP bannie, whois, nombre d'échecs, extrait des lignes de log. C'est verbeux : réservez-le aux jails critiques (sshd, recidive), pas à nginx-botsearch (sinon : 200 mails/nuit).

### Réduire le bruit : n'alerter que recidive

```ini
[DEFAULT]
action = %(action_)s          # silencieux par défaut

[recidive]
enabled = true
action  = %(action_mwl)s[name=recidive, dest=admin@example.com, sender=fail2ban@srv-web.example.com]
```

Vous ne recevez un mail que pour les vrais récidivistes (5 bans en 1 jour) : signal utile, bruit minimal.

### Webhook Mattermost/Slack (rappel section 22)

