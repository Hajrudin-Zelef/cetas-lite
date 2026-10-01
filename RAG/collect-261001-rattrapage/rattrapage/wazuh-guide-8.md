---
id: collect-261001-rattrapage/rattrapage/wazuh-guide-8
title: "Guide Wazuh — SIEM & XDR Open Source en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "attention"]
source: docs/RAG/collect-261001-rattrapage/wazuh_guide.md
source_anchor: ""
source_lines: [1162, 1386]
sha256: 92d19cb5acb1840ee5927c381d08eef8e3da250665dc9bd3adde65ece152b718
---

# Guide Wazuh — SIEM & XDR Open Source en production

```bash
sudo /var/ossec/bin/wazuh-logtest
# Collez la ligne de log, observez :
# **Phase 1: Completed decoding.
#        decoder: 'monapp-base'
#        srcip: '10.0.5.23'
#        dstuser: 'jdupont'
#        action: 'export'
```

Si le decoder ne matche pas : vérifiez les échappements (`\\S` dans XML), l'ordre des groupes, et que `<prematch>` est bien présent dans la ligne.

---

## 29. Écrire sa première règle personnalisée

On reprend le decoder `monapp-base` : alerter quand quelqu'un exporte `clients.csv` (donnée sensible).

```xml
<!-- /var/ossec/etc/rules/local_rules.xml -->
<group name="monapp,">
  <rule id="100100" level="10">
    <decoded_as>monapp-base</decoded_as>
    <field name="action">export</field>
    <field name="filename">clients\.csv</field>
    <description>Export sensible clients.csv par $(dstuser) depuis $(srcip)</description>
    <mitre>
      <id>T1048</id>   <!-- Exfiltration Over Alternative Protocol (exemple) -->
    </mitre>
  </rule>
</group>
```

Points d'attention :

- **IDs 100000–199999** réservés aux règles locales (jamais de collision avec le ruleset officiel).
- `$(dstuser)` : référence un champ décodé dans la description.
- Testez avec `wazuh-logtest` : la phase 2 doit afficher `**Rule 100100 matched.`
- Rechargez les règles : `sudo systemctl reload wazuh-manager` (ou redémarrage complet si doute).

---

## 30. Niveaux de règles (0 à 16) : la grammaire de la sévérité

| Niveau | Signification | Exemples |
|---|---|---|
| 0 | Ignoré (pas d'alerte) | Bruit filtré |
| 1 | Événement basique | Démarrage service |
| 2 | Notification système | — |
| 3 | Événement notable | **Agent déconnecté**, connexion réussie inhabituelle |
| 4 | Erreur système | Échec de service |
| 5 | Erreur utilisateur | Mauvais mot de passe isolé |
| 6 | Attaque de faible pertinence | Scan de ports léger |
| 7 | Mauvaise configuration / « bad word » | Mot-clé suspect |
| 8 | Première occurrence / sévérité notable | Premier succès après échecs |
| 9 | Erreur d'application critique | — |
| 10 | **Attaque confirmée** | Force brute SSH réussie à bloquer |
| 11 | Erreur d'intégrité | Fichier modifié (FIM) |
| 12 | Attaque très pertinente | Exploitation probable |
| 13 | Corrélation / anomalie | Comportement anormal |
| 14 | Corrélation temporelle forte | Plusieurs attaques liées |
| 15 | Compromission probable | **Alerte critique, agir immédiatement** |
| 16 | Attaque active en cours | Rare, réponse immédiate |

Règle d'or pour une petite équipe : **tout ce qui est ≥ 10 doit arriver sur un canal qui réveille** (e-mail prioritaire, webhook, SMS via passerelle). Le reste se traite en revue quotidienne (section 81).

---

## 31. Tester ses règles sans casser la production

`wazuh-logtest` simule decoders + règles sans toucher à la production :

```bash
sudo /var/ossec/bin/wazuh-logtest
```

Session type :

```
Type one log per line.

Sep 26 22:15:01 srv-web sshd[1234]: Failed password for root from 203.0.113.7 port 51234 ssh2

**Phase 1: Completed decoding.
       decoder: 'sshd'
       srcip: '203.0.113.7'
       dstuser: 'root'

**Phase 2: Completed filtering (rules).
       Rule id: '5716'
       Level: '5'
       Description: 'SSHD: authentication failed.'
```

Bonnes pratiques :

1. Testez **chaque** règle avant déploiement.
2. Testez aussi les **non-déclenchements** (une ligne normale ne doit pas matcher).
3. Gardez un fichier de lignes de test : `/root/wazuh-tests/echantillons.log` avec cas positifs et négatifs.

```bash
# Rejouer un fichier de test complet
while IFS= read -r ligne; do
  echo "$ligne" | sudo /var/ossec/bin/wazuh-logtest 2>/dev/null | grep -E "Rule id|Level"
done < /root/wazuh-tests/echantillons.log
```

---

## 32. Organiser ses règles : local_rules, CDB lists

### 32.1 Structure conseillée

```
/var/ossec/etc/rules/
├── local_rules.xml          # vos règles métier (inclus par défaut)
├── soc_tuning.xml           # surcharges : désactivation / baisse de niveau
└── listes-noires.xml        # règles basées sur CDB lists
/var/ossec/etc/lists/
├── bad-ips.txt              # une IP par ligne : 203.0.113.7:
├── comptes-sensibles.txt    # root: / admin: / backup:
└── ...
```

Incluez vos fichiers dans `ossec.conf` du manager :

```xml
<ruleset>
  <rule_dir>ruleset/rules</rule_dir>
  <decoder_dir>ruleset/decoders</decoder_dir>
  <rule_exclude>0215-policy_rules.xml</rule_exclude>  <!-- exemple : exclure un lot -->
</ruleset>
```

Et déclarez les listes :

```xml
<ruleset>
  <list>etc/lists/bad-ips.txt</list>
</ruleset>
```

### 32.2 Règle basée sur CDB list

```
# /var/ossec/etc/lists/bad-ips.txt  (format clé: — le ":" final est obligatoire)
203.0.113.7:
198.51.100.23:
```

```xml
<rule id="100200" level="12">
  <if_sid>5700</if_sid>
  <list field="srcip" lookup="address_match_key">etc/lists/bad-ips.txt</list>
  <description>Connexion depuis une IP inscrite sur liste noire SOC : $(srcip)</description>
</rule>
```

### 32.3 Surcharger une règle officielle (sans la modifier)

```xml
<!-- /var/ossec/etc/rules/soc_tuning.xml -->
<group name="tuning,">
  <!-- Baisser le niveau d'une règle trop bruyante -->
  <rule id="5710" level="3" overwrite="yes">
    <description>SSHD: Attempt to login using a non-existent user (bruit réduit)</description>
  </rule>
  <!-- Désactiver proprement -->
  <!-- <rule id="510" level="0" overwrite="yes" /> = ne JAMAIS faire : préférez l'exclusion ciblée -->
</group>
```

> ⚠️ `overwrite="yes"` sur une règle officielle : documentez **pourquoi** dans un commentaire XML. Dans 6 mois, personne ne s'en souviendra sinon.

---

## 33. FIM — surveillance d'intégrité des fichiers : principes

Le FIM (*File Integrity Monitoring*, module `syscheck`) calcule des empreintes (MD5, SHA1, SHA256) des fichiers surveillés et alerte à **toute modification** : contenu, permissions, propriétaire, taille.

C'est votre détecteur de :
- webshell déposé sur un serveur web,
- binaire système remplacé (rootkit),
- modification de `/etc/passwd`, `/etc/shadow`, des tâches cron,
- changement de configuration non déclaré.

Cycle : **scan initial** (baseline) → scans périodiques → comparaison → alerte si différence. La baseline initiale génère un pic d'alertes : normal, ne paniquez pas (voir section 50 pour le tuning).

---

## 34. Configurer syscheck (FIM) sur Linux

Via `agent.conf` du groupe (recommandé) ou `/var/ossec/etc/ossec.conf` de l'agent :

```xml
<agent_config>
  <syscheck>
    <frequency>7200</frequency>          <!-- scan complet toutes les 2h -->
    <scan_on_start>yes</scan_on_start>

    <!-- Dossiers critiques système -->
    <directories check_all="yes" realtime="yes" report_changes="yes">/etc,/usr/bin,/usr/sbin</directories>
    <directories check_all="yes" realtime="yes" report_changes="yes">/bin,/sbin,/boot</directories>

    <!-- Webroot applicatif -->
    <directories check_all="yes" realtime="yes" report_changes="yes" restrict="\.php$">/var/www/html</directories>

    <!-- Exclusions : bruit inutile -->
    <ignore>/etc/mtab</ignore>
    <ignore>/etc/hosts.deny</ignore>
    <ignore type="sregex">^/etc/.*~$</ignore>
    <ignore>/var/log</ignore>
    <ignore>/tmp</ignore>

    <!-- Ne pas suivre ces montages -->
    <skip_nfs>yes</skip_nfs>
    <skip_dev>yes</skip_dev>
    <skip_proc>yes</skip_proc>
    <skip_sys>yes</skip_sys>
  </syscheck>
</agent_config>
```

Options de `check_*` (au lieu de `check_all`) :

