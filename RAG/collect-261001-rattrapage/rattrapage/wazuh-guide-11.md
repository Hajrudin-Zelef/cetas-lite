---
id: collect-261001-rattrapage/rattrapage/wazuh-guide-11
title: "Guide Wazuh — SIEM & XDR Open Source en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "backlog", "incident"]
source: docs/RAG/collect-261001-rattrapage/wazuh_guide.md
source_anchor: ""
source_lines: [1822, 2048]
sha256: c195d78ee9794444fba696a6fac4dad055ea630f7d222833cb3f1e73e323e50b
---

# Guide Wazuh — SIEM & XDR Open Source en production

```bash
#!/bin/bash
# /var/ossec/active-response/bin/disable-account.sh
# Usage : appelé par Wazuh avec les variables d'environnement (alert JSON sur stdin)
read -r ALERT
USER=$(echo "$ALERT" | grep -o '"dstuser":"[^"]*"' | head -1 | cut -d'"' -f4)
ACTION=$1   # add | delete

if [ "$ACTION" = "add" ] && [ -n "$USER" ] && [ "$USER" != "root" ]; then
  /usr/sbin/usermod -L "$USER"
  logger -t wazuh-ar "Compte $USER verrouillé par active-response"
fi
exit 0
```

```bash
sudo chmod 750 /var/ossec/active-response/bin/disable-account.sh
sudo chown root:wazuh /var/ossec/active-response/bin/disable-account.sh
```

Déclaration côté manager :

```xml
<command>
  <name>disable-account</name>
  <executable>disable-account.sh</executable>
  <timeout_allowed>no</timeout_allowed>
  <expect>dstuser</expect>
</command>

<active-response>
  <command>disable-account</command>
  <location>local</location>
  <rules_id>100400</rules_id>   <!-- votre règle "5 échecs sur compte sensible" -->
</active-response>
```

> Déployez le script via les **fichiers partagés du groupe** (`/var/ossec/etc/shared/<groupe>/` ne pousse que la conf...). Pour les scripts, utilisez votre outil de conf (Ansible) ou copiez-les manuellement : **un active-response qui pointe vers un script absent échoue silencieusement** côté agent.

Autres idées de scripts : couper l'interface réseau d'un poste suspect (`quarantine.sh`), redémarrer un service, prendre un snapshot mémoire (forensique), appeler l'API de votre EDR.

---

## 47. Correspondance MITRE ATT&CK

Wazuh mappe nativement ses règles sur la matrice **MITRE ATT&CK** (tactiques, techniques, IDs comme `T1110` *Brute Force*).

Où le voir : Dashboard → *Threat Intelligence → MITRE ATT&CK*. Chaque alerte enrichie affiche tactique (ex. *Credential Access*) et technique.

Ajouter MITRE à vos règles (vu section 29) :

```xml
<rule id="100100" level="10">
  ...
  <mitre>
    <id>T1048</id>
    <id>T1020</id>
  </mitre>
  <group>exfiltration,</group>
</rule>
```

Usage SOC concret :

1. **Revue hebdo par tactique** : « cette semaine, 80 % de nos alertes niveau ≥ 10 sont du *Credential Access* » → priorisez le durcissement des mots de passe / MFA.
2. **Couverture** : identifiez les tactiques **sans aucune règle** chez vous (ex. *Lateral Movement*) → écrivez les règles manquantes.
3. **Reporting direction** : la heatmap MITRE parle mieux aux non-techniques qu'une liste de règles.

---

## 48. Gestion des alertes : cycle de vie

```
Détection → Qualification → Investigation → Réponse → Clôture → Retour d'expérience
```

### 48.1 Qualification (triage)

Pour chaque alerte niveau ≥ 7, répondez à 4 questions :

1. **C'est quoi ?** (règle, machine, utilisateur, IP)
2. **C'est attendu ?** (maintenance, batch, admin légitime)
3. **C'est grave ?** (niveau, MITRE, criticité de l'actif)
4. **J'agis quand ?** (immédiat / aujourd'hui / backlog tuning)

### 48.2 Statuts de traitement (convention d'équipe)

| Statut | Signification |
|---|---|
| `Nouveau` | Non qualifié |
| `En cours` | Un analyste investigue |
| `Faux positif` | Bruit → **crée un ticket de tuning** (section 50) |
| `Bénin expliqué` | Légitime mais à surveiller |
| `Incident` | Attaque réelle → procédure incident |
| `Clos` | Traité, documenté |

Wazuh ne gère pas nativement ce workflow : utilisez **TheHive** (section 54) ou, à défaut, un tableau partagé. L'important : **aucune alerte ≥ 10 ne reste « Nouveau » plus de 24 h**.

### 48.3 Investigation : les 3 pivots

Depuis une alerte, pivotez toujours sur :
- **l'agent** : que s'est-il passé sur cette machine ± 1 h ?
- **l'IP source** : a-t-elle touché d'autres machines ? (recherche `srcip` sur 7 jours)
- **l'utilisateur** : ses autres connexions sont-elles normales ? (géographie, horaires)

Requête dashboard type (KQL) :

```
rule.level >= 10 and agent.name : "srv-web-01"
```

---

## 49. Niveaux de sévérité : que faire pour chaque niveau

| Niveau | Canal | Délai de traitement | Exemple d'action |
|---|---|---|---|
| 0–2 | Aucun (log seul) | — | Rien |
| 3–6 | Dashboard (revue quotidienne) | 24 h | Qualifier, tuner si bruit |
| 7–9 | Dashboard + e-mail groupé | 4 h ouvrées | Investiguer |
| 10–12 | **Alerte push** (e-mail prioritaire, webhook) | 1 h | Investiguer + active response auto |
| 13–16 | **Réveil d'astreinte** | Immédiat | Procédure incident |

> Formalisez ce tableau dans votre **procédure SOC** (même une page) : en pleine nuit, on ne réfléchit pas, on exécute.

---

## 50. Tuning : réduire les faux positifs

Le tuning est **le** travail continu du SOC. Sans lui, Wazuh devient une machine à bruit que plus personne ne lit.

### 50.1 Mesurer avant de tuner

Dashboard → *Discover* : comptez les alertes par `rule.id` sur 7 jours. Les 10 règles les plus bruyantes = vos cibles.

```bash
# Top 10 des règles bruyantes via l'API (7 derniers jours, exemple)
curl -sk -H "Authorization: Bearer $TOKEN" \
  'https://localhost:55000/alerts?pretty=true&limit=1' | head -5
# Pour l'agrégation fine, utilisez le dashboard (visualisation par rule.id)
```

### 50.2 Les 5 techniques de tuning

**1. Exclure des chemins/IP bruit (global)**

```xml
<global>
  <white_list>10.0.0.5</white_list>  <!-- sonde de supervision qui scanne -->
</global>
```

**2. Baisser le niveau d'une règle bruyante** (voir `soc_tuning.xml`, section 32.3).

**3. Restreindre une règle à un contexte** avec `<if_sid>` + conditions :

```xml
<!-- N'alerter sur les échecs sudo QUE pour les comptes sensibles -->
<rule id="100500" level="10">
  <if_sid>5402</if_sid>
  <list field="dstuser" lookup="address_match_key">etc/lists/comptes-sensibles.txt</list>
  <description>Échec sudo répété sur compte sensible : $(dstuser)</description>
</rule>
```

**4. Ignorer via `rule_exclude`** un lot entier non pertinent chez vous :

```xml
<ruleset>
  <rule_exclude>0400-imapd_rules.xml</rule_exclude>  <!-- pas de serveur IMAP -->
</ruleset>
```

**5. Ajuster les seuils de fréquence** (ex. brute force) :

```xml
<!-- Par défaut la règle 5763 se base sur 8 échecs ; resserrez ou élargissez -->
<rule id="5763" level="10" frequency="12" timeframe="120" overwrite="yes">
  <if_matched_sid>5716</if_matched_sid>
  <description>sshd: brute force attack (seuil SOC : 12/120s).</description>
</rule>
```

### 50.3 Rituel de tuning

- **Quotidien (15 min)** : qualifier les nouvelles alertes, taguer les faux positifs.
- **Hebdomadaire (1 h)** : top 10 bruit → 1–2 actions de tuning.
- **Mensuel** : revoir les règles `overwrite`, vérifier qu'elles sont toujours pertinentes.

> Un bon objectif : **moins de 50 alertes/jour/analyste à qualifier** après 3 mois de tuning. Au-delà, le SOC sature et rate les vraies attaques.

---

## 51. Supervision réseau via agents (état des interfaces, ports)

Le module `syscollector` remonte aussi l'état réseau ; le module `command` (wodle) exécute des commandes périodiques dont la sortie est analysée comme un log.

Exemple : surveiller les **ports en écoute** d'un serveur (détection d'un backdoor qui ouvre un port) :

```xml
<!-- agent.conf du groupe linux-serveurs -->
<agent_config>
  <wodle name="command">
    <disabled>no</disabled>
    <command>ss -tlnp</command>
    <interval>1h</interval>
    <ignore_output>no</ignore_output>
    <run_on_start>yes</run_on_start>
    <timeout>0</timeout>
  </wodle>
</agent_config>
```

```xml
<!-- Règle : alerter si un port inhabituel écoute -->
<rule id="100600" level="10">
  <if_sid>530</if_sid>   <!-- règle générique wodle command -->
  <regex>:(4444|31337|1337|6667)\s</regex>
  <description>Port suspect en écoute : possible backdoor ($(full_log))</description>
  <mitre><id>T1205</id></mitre>
</rule>
```

Autres commandes utiles en wodle : `ip addr` (changement d'IP), `arp -a`, `iptables -L -n` (règle pare-feu modifiée), `df -h` (disque plein, voir section 72).

---

