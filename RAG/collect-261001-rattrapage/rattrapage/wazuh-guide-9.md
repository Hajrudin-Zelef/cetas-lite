---
id: collect-261001-rattrapage/rattrapage/wazuh-guide-9
title: "Guide Wazuh — SIEM & XDR Open Source en production"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agent", "agents"]
source: docs/RAG/collect-261001-rattrapage/wazuh_guide.md
source_anchor: ""
source_lines: [1387, 1598]
sha256: 3cdac434c831dec35df49345015c859c75fa66198757d13c8a3509bb56ccb4ce
---

# Guide Wazuh — SIEM & XDR Open Source en production

| Attribut | Surveille |
|---|---|
| `check_md5sum="yes"` | Empreinte MD5 |
| `check_sha1sum="yes"` | Empreinte SHA1 |
| `check_sha256sum="yes"` | Empreinte SHA256 |
| `check_size="yes"` | Taille |
| `check_owner="yes"` | Propriétaire |
| `check_group="yes"` | Groupe |
| `check_perm="yes"` | Permissions |
| `check_mtime="yes"` | Date de modification |
| `check_inode="yes"` | Inode |
| `report_changes="yes"` | **Inclut le diff du contenu dans l'alerte** (précieux !) |

> `report_changes="yes"` sur des fichiers texte de config = vous voyez **exactement** quelle ligne a changé. Sur des binaires, inutile (et volumineux) : limitez-le aux fichiers texte.

---

## 35. Configurer syscheck (FIM) sur Windows

```xml
<agent_config os="Windows">
  <syscheck>
    <frequency>7200</frequency>
    <scan_on_start>yes</scan_on_start>

    <directories check_all="yes" realtime="yes">C:\Windows\System32</directories>
    <directories check_all="yes" realtime="yes">C:\Windows\SysWOW64</directories>
    <directories check_all="yes">C:\Program Files</directories>
    <directories check_all="yes">C:\Program Files (x86)</directories>

    <!-- Registre : clés sensibles -->
    <windows_registry arch="both">HKEY_LOCAL_MACHINE\Software\Microsoft\Windows\CurrentVersion\Run</windows_registry>
    <windows_registry arch="both">HKEY_LOCAL_MACHINE\System\CurrentControlSet\Services</windows_registry>
    <windows_registry>HKEY_LOCAL_MACHINE\Software\Microsoft\Windows NT\CurrentVersion\Winlogon</windows_registry>

    <!-- Exclusions : bruit Windows -->
    <ignore>C:\Windows\System32\winevt\Logs</ignore>
    <ignore>C:\Windows\Temp</ignore>
    <ignore type="sregex">\.log$</ignore>
    <ignore type="sregex">\.tmp$</ignore>
  </syscheck>
</agent_config>
```

Le FIM registre est une pépite : 90 % des persistances malware Windows passent par `Run` ou les services. Une alerte « clé de registre ajoutée dans Run » un samedi à 3h = investigation immédiate.

---

## 36. FIM temps réel vs who-data : bien choisir

| Mode | Fonctionnement | Avantages | Inconvénients |
|---|---|---|---|
| **Planifié** (`frequency`) | Scan périodique complet | Simple, peu gourmand | Détection différée (jusqu'à `frequency`) |
| **Temps réel** (`realtime="yes"`) | Notification immédiate (inotify/Win32 API) | Détection instantanée | Plus gourmand sur dossiers très actifs |
| **Who-data** (`whodata="yes"`) | Auditd (Linux) : **qui** a modifié | Identifie l'utilisateur/processus | Linux seul, auditd requis, plus lourd |

Recommandation :

```xml
<!-- Binaires système : temps réel, on veut savoir TOUT DE SUITE -->
<directories check_all="yes" realtime="yes">/usr/bin,/usr/sbin,/bin,/sbin</directories>

<!-- Fichiers de conf : who-data pour savoir QUI a changé quoi -->
<directories check_all="yes" whodata="yes" report_changes="yes">/etc/ssh,/etc/sudoers.d</directories>

<!-- Gros dossiers applicatifs : planifié pour ne pas saturer -->
<directories check_all="yes">/opt/monapp</directories>
```

Prérequis who-data sur Linux : `auditd` installé et démarré sur l'agent.

```bash
sudo apt install -y auditd
sudo systemctl enable --now auditd
```

> ⚠️ **Ne mettez jamais `realtime` sur `/var/log`** : chaque écriture de log déclencherait un scan → boucle infernale et explosion d'EPS.

---

## 37. Exploiter les alertes FIM au quotidien

Une alerte FIM typique (niveau 7 par défaut, règles 550-557) :

```json
{
  "rule": { "id": "550", "level": 7, "description": "Integrity checksum changed." },
  "syscheck": {
    "path": "/etc/passwd",
    "mode": "whodata",
    "user_name": "root",
    "process_name": "/usr/sbin/useradd",
    "changed_attributes": ["mtime", "sha256_after", "size_after"],
    "diff": "+jdupont:x:1001:1001::/home/jdupont:/bin/bash\n"
  }
}
```

Lecture SOC :

1. **Quel fichier ?** `/etc/passwd` = sensible.
2. **Qui ?** `user_name: root`, `process_name: /usr/sbin/useradd` = ajout d'utilisateur légitime probable.
3. **Quoi ?** le `diff` montre la ligne ajoutée.
4. **Contexte** : y a-t-il un ticket de création de compte ? Si non → investigation.

Règles de triage FIM :

| Signal | Interprétation | Action |
|---|---|---|
| Fichier binaire modifié + `realtime`, heure ouvrée, package manager actif | Mise à jour logicielle | Vérifier les logs apt/yum, classer |
| Fichier web `.php` modifié la nuit | Compromission probable | Isoler, analyser (cas pratique 3) |
| `/etc/shadow` modifié + who-data `unknown` | Suspect | Investigation immédiate |
| Dizaines de fichiers modifiés d'un coup | Mise à jour système OU ransomware | Corréler avec l'activité (cas pratique 2) |

**Montez le niveau** des FIM critiques via `local_rules.xml` :

```xml
<rule id="100300" level="12">
  <if_sid>550</if_sid>
  <field name="file">/etc/shadow</field>
  <description>FIM CRITIQUE : /etc/shadow modifié !</description>
</rule>
```

---

## 38. Détection de vulnérabilités : inventaire logiciel

Le module **Vulnerability Detector** (manager) croise l'inventaire logiciel remonté par les agents (module `syscollector`) avec les flux CVE (NVD, Canonical, Microsoft, Red Hat...).

Ce que ça donne : pour chaque agent, la liste des CVE le concernant, avec score CVSS, sévérité et références. C'est votre **tableau de bord patch management**.

Prérequis côté agent : `syscollector` activé (inventaire paquets/OS). Il l'est par défaut sur les agents 4.x récents ; vérifiez :

```bash
# Sur l'agent : forcer un inventaire immédiat (au lieu d'attendre l'intervalle)
sudo /var/ossec/bin/wazuh-control restart   # syscollector tourne au démarrage
sudo grep -i syscollector /var/ossec/logs/ossec.log | tail -5
```

---

## 39. Configurer le module vulnerability detector

Sur le **manager**, `/var/ossec/etc/ossec.conf` :

```xml
<vulnerability-detector>
  <enabled>yes</enabled>
  <interval>1d</interval>              <!-- rescan quotidien -->
  <min_full_scan_interval>1d</min_full_scan_interval>
  <run_on_start>yes</run_on_start>

  <!-- Fournisseurs : laissez ceux de vos OS -->
  <provider name="canonical">
    <enabled>yes</enabled>
    <os>jammy</os>      <!-- Ubuntu 22.04 -->
    <os>focal</os>      <!-- Ubuntu 20.04 -->
    <update_interval>1h</update_interval>
  </provider>
  <provider name="debian">
    <enabled>yes</enabled>
    <os>bookworm</os>   <!-- Debian 12 -->
    <os>bullseye</os>   <!-- Debian 11 -->
    <update_interval>1h</update_interval>
  </provider>
  <provider name="redhat">
    <enabled>yes</enabled>
    <os>9</os>
    <update_interval>1h</update_interval>
  </provider>
  <provider name="nvd">
    <enabled>yes</enabled>
    <update_interval>1h</update_interval>
  </provider>
  <provider name="msu">
    <enabled>yes</enabled>
    <update_interval>1h</update_interval>
  </provider>
</vulnerability-detector>

<!-- Indexer : stocker l'inventaire -->
<indexer>
  <enabled>yes</enabled>
  <hosts>
    <host>https://127.0.0.1:9200</host>
  </hosts>
  <username>admin</username>
  <password><MOT_DE_PASSE></password>
  <ssl>
    <certificate_authorities>/etc/filebeat/certs/root-ca.pem</certificate_authorities>
    <certificate>/etc/filebeat/certs/filebeat.pem</certificate>
    <key>/etc/filebeat/certs/filebeat-key.pem</key>
  </ssl>
</indexer>
```

> ⚠️ **Ne mettez jamais de mot de passe en clair dans un exemple partagé.** En production, le bloc `<indexer>` est généré par l'assistant avec les bons secrets ; ne l'éditez qu'en connaissance de cause.

Vérification :

```bash
sudo grep -i "vulnerability" /var/ossec/logs/ossec.log | tail -10
# Attendu : téléchargement des bases CVE, puis "Vulnerability scan finished"
```

Le premier scan télécharge plusieurs centaines de Mo de bases CVE : **prévoyez le temps et l'espace disque** (`/var/ossec/queue/vulnerabilities/`).

---

## 40. Lire et prioriser les alertes de vulnérabilités

