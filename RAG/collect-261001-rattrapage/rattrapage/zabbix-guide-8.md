---
id: collect-261001-rattrapage/rattrapage/zabbix-guide-8
title: "Guide Zabbix complet — Supervision d'infrastructure en production"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/zabbix_guide.md
source_anchor: ""
source_lines: [1224, 1432]
sha256: 613017cef310394a338862e9ea639f54f7d54994735a1935f148b02af36b8a17
---

# Guide Zabbix complet — Supervision d'infrastructure en production

- *Default message* : personnalisez le sujet/corps (voir 41).
- *Pause* entre étapes, *durée d'escalade*.
- *Conditions d'acquittement* : ne pas escalader si acquitté.

> 💡 **Une action par criticité/métier**, pas une action par trigger. 3-5 actions bien pensées valent mieux que 50 actions bricolées.

## 41. Médias d'alerte : email

*Administration → Media types → Email* :

| Paramètre | Valeur type |
|---|---|
| SMTP server | `smtp.votre-domaine.local` ou relais interne |
| SMTP port | `587` (STARTTLS) ou `25` |
| Connection security | STARTTLS / SSL/TLS |
| Authentication | Username + password (compte dédié `zabbix@votre-domaine.local`) |

**Testez** avec *Test* (envoie un vrai email). Modèles de message (*Message templates*) :

```
Sujet : [{TRIGGER.SEVERITY}] {HOST.NAME} : {TRIGGER.NAME} ({EVENT.STATUS})
Corps :
  Alerte : {TRIGGER.NAME}
  Hôte : {HOST.NAME} ({HOST.IP})
  Sévérité : {TRIGGER.SEVERITY}
  Heure : {EVENT.DATE} {EVENT.TIME}
  Valeur : {ITEM.VALUE} ({ITEM.NAME})
  Détails : {TRIGGER.DESCRIPTION}
  Acquitter : {EVENT.ACKNOWLEDGE.URL}  (⚠️ 7.x : macro disponible)
```

> 💡 Utilisez un **relais SMTP interne** (Postfix) plutôt que le SMTP du FAI : en cas de coupure internet, les emails internes continuent de partir… si votre relais est local. Pour l'externe, le SMS reste roi (section suivante).

## 42. Médias d'alerte : SMS via passerelle (contexte Afrique de l'Ouest)

L'email ne suffit pas en astreinte. Options classées par fiabilité terrain :

### Option A — Passerelle GSM locale (recommandée)

Un modem/routeur GSM (clé 4G, routeur Teltonika…) + outil d'envoi :

```bash
# Exemple avec gammu + clé USB Huawei
apt install -y gammu
# /etc/gammurc
[gammu]
port = /dev/ttyUSB0
connection = at115200

# Envoi
echo "ZABBIX [High] ups-salle-01 : sur batterie depuis 5 min" \
  | gammu sendsms TEXT +2250700000000
```

Média Zabbix de type **Script** : *Administration → Media types → Create → Type: Script*, script `/usr/lib/zabbix/alertscripts/sendsms.sh` :

```bash
#!/bin/bash
# $1 = numéro, $2 = sujet, $3 = message
TO="$1"; MSG="$2 - $3"
echo "$MSG" | gammu sendsms TEXT "$TO"
```

```bash
chmod +x /usr/lib/zabbix/alertscripts/sendsms.sh
chown zabbix:zabbix /usr/lib/zabbix/alertscripts/sendsms.sh
```

> 💡 Gardez la clé 4G d'alerte sur un **opérateur différent** de votre lien principal, avec du crédit en permanence. Testez l'envoi **chaque mois** (section 62).

### Option B — API d'un agrégateur SMS

Script utilisant `curl` vers l'API de votre fournisseur (Orange/MTN/Wave API, etc.) :

```bash
#!/bin/bash
TO="$1"; SUBJECT="$2"; BODY="$3"
curl -s -X POST "https://api.fournisseur-sms.local/send" \
  -H "Authorization: Bearer CLE_API_ICI" \
  -d "to=$TO" -d "message=$(echo "$SUBJECT $BODY" | cut -c1-160)"
```

> ⚠️ Dépend d'internet : prévoyez l'option A en secours pour les alertes Disaster.

### Option C — SMS via email-to-SMS de l'opérateur

Si votre opérateur offre une passerelle `numero@sms.operateur.xx`, réutilisez le média Email avec une adresse par destinataire.

## 43. Médias d'alerte : Telegram

Gratuit, instantané, avec accusé de lecture — parfait pour une équipe technique.

1. Créez un bot via **@BotFather** → récupérez le **token** (`123456:ABC-DEF...`).
2. Créez un groupe d'astreinte, ajoutez le bot, récupérez le **chat_id** (`https://api.telegram.org/bot<TOKEN>/getUpdates`).
3. *Administration → Media types → Telegram* (média natif depuis 5.x) :

| Paramètre | Valeur |
|---|---|
| Bot token | `123456:ABC-DEF...` (le vôtre) |

4. Dans chaque utilisateur : *Media → Add → Type: Telegram*, *Send to* : le chat_id (ou `@pseudo`).
5. Message template adapté (court !) :

```
[{TRIGGER.SEVERITY}] {HOST.NAME}
{TRIGGER.NAME}
{EVENT.DATE} {EVENT.TIME} — {ITEM.VALUE}
```

> 💡 Un groupe Telegram "Astreinte Zabbix" où le bot poste : toute l'équipe voit qui a acquitté quoi. Zéro coût, historique intégré.

## 44. Webhooks : Slack, Teams, et autres

*Administration → Media types* : des webhooks pré-packagés existent (Slack, Teams, PagerDuty, Jira…). Principe : un script JavaScript embarqué appelle l'API distante.

Exemple minimal (webhook générique, à adapter) :

```javascript
// Paramètres du média : URL, Token
var params = JSON.parse(value);
var req = new HttpRequest();
req.addHeader('Content-Type: application/json');
req.addHeader('Authorization: Bearer ' + params.Token);
var body = JSON.stringify({
    text: params.Subject + "\n" + params.Message
});
var resp = req.post(params.URL, body);
if (req.getStatus() != 200) {
    throw 'Echec webhook: ' + resp;
}
return 'OK';
```

> ⚠️ Les webhooks sortants exigent internet : ne faites **jamais** reposer une alerte Disaster uniquement sur un webhook.

## 45. Escalades et acquittements

### Escalade type (à configurer dans l'action, section 40)

```
Étape 1 — immédiate : Email + Telegram → groupe "Exploitation"
Étape 2 — après 15 min sans acquittement : SMS → groupe "Exploitation"
Étape 3 — après 45 min sans acquittement : SMS + appel → "Chef de service"
Récupération : Email + Telegram → tous ("RÉSOLU")
```

### Acquittements

Depuis *Monitoring → Problems* : *Acknowledge* → message ("Pris en charge, intervention à 14h"), éventuellement **changer la sévérité** ou **mettre en maintenance**. L'acquittement **stoppe l'escalade** (si l'action est configurée ainsi) mais ne résout pas le problème.

> 💡 Règle d'équipe : **tout PROBLEM High/Disaster est acquitté sous 15 min** avec un commentaire. Un problème non acquitté qui escalade, c'est le système qui fonctionne.

## 46. Découverte réseau automatique

*Configuration → Discovery → Create discovery rule*. Exemple : balayer `192.168.10.1-254` :

| Check | Port | Règle |
|---|---|---|
| ICMP ping | — | `icmpping` |
| SNMPv2 agent | 161 | `snmpget` OID `1.3.6.1.2.1.1.1.0` (sysDescr) |
| Zabbix agent | 10050 | `system.hostname` |

Puis *Configuration → Actions → Discovery actions* : "si SNMP détecté → ajouter au groupe `Réseau/Auto`, lier le template `Generic SNMP`".

> ⚠️ La découverte, c'est bien ; la **découverte + action auto** non relue, c'est 300 imprimantes dans votre inventaire. Limitez les plages, relisez la file *Monitoring → Discovery*, et **désactivez** la règle après le déploiement initial.

## 47. Low-Level Discovery (LLD) : disques, interfaces, processus

La LLD découvre **à l'intérieur** d'un hôte : partitions, interfaces réseau, disques, processus… et crée items/triggers/graphes par entité.

### Exemple : LLD des systèmes de fichiers (déjà dans le template Linux)

Règle `vfs.fs.discovery` → prototypes d'items :

```
Item prototype : vfs.fs.size[{#FSNAME},pused]
Trigger prototype : "{#FSNAME} plein à > {$VFS_FS_PUSED_CRIT}%"
Graph prototype : espace disque
```

**Macros LLD** `{#FSNAME}`, `{#IFNAME}`… : remplacées par entité découverte.

### Filtres LLD (indispensables)

Sans filtre, vous supervisez `/snap/*`, `docker0`, `veth*`… :

```
Filtre : {#FSNAME} matches ^/(|var|home|opt)$      (regex)
     ET  {#FSTYPE} matches ^(ext4|xfs)$
```

> 💡 **Toujours** filtrer les LLD réseau (`{#IFNAME} not matches ^(lo|docker|veth)`) et disque. Une LLD non filtrée = des centaines d'items parasites et des alertes absurdes.

## 48. Écrire ses propres règles LLD (UserParameter + JSON)

Une règle LLD attend un **JSON** : `{"data":[{"{#MACRO1}":"val1",...}, ...]}`.

Exemple : découvrir les onduleurs USB branchés (ou tout parc maison) :

```bash
# /etc/zabbix/zabbix_agent2.d/llb_custom.conf
UserParameter=custom.ups.discovery,python3 /usr/local/bin/ups_discovery.py
```

```python
#!/usr/bin/env python3
"""Découverte des onduleurs : sort un JSON LLD."""
import json, subprocess

