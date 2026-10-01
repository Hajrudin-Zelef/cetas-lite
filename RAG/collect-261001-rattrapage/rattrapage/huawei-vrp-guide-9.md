---
id: collect-261001-rattrapage/rattrapage/huawei-vrp-guide-9
title: "VRP — Le système d'exploitation transversal Huawei"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-rattrapage/huawei_vrp_guide.md
source_anchor: ""
source_lines: [1663, 1914]
sha256: 8163ea04b4c9184511af678c17900432ffb64dc11b50ccf4f50e67da2d909593
---

# VRP — Le système d'exploitation transversal Huawei

```vrp
[AR720]snmp-agent
[AR720]snmp-agent sys-info version v3
[AR720]snmp-agent group v3 group-secure privacy
[AR720]snmp-agent usm-user v3 user-secure group-secure
[AR720]snmp-agent usm-user v3 user-secure authentication-mode sha AuthPass123! privacy-mode aes128 PrivPass456!
[AR720]snmp-agent target-host trap-hostname ZABBIX trap address udp-domain 192.168.100.60 params securityname user-secure v3 privacy
```

SNMPv3 = authentification + chiffrement. En 2026, c'est le standard à
viser pour toute supervision en production.

## 93. Accès web (HTTP/HTTPS)

```vrp
[AR720]http server enable
[AR720]http secure-server enable               # HTTPS (recommandé)
[AR720]http server-source -i Vlanif 10
[AR720]aaa
[AR720-aaa]local-user admin service-type http
[AR720-aaa]quit
```

URL : `https://192.168.10.1:8443` (le port exact varie selon modèle —
à vérifier avec `display http server` / `display http secure-server`).
L'interface web est utile pour la visualisation, mais le CLI reste la
référence pour l'exploitation sérieuse et l'automatisation.

## 94. AAA et RADIUS — centraliser l'authentification

```vrp
[AR720]radius-server template RADIUS_SIEGE
[AR720-radius-RADIUS_SIEGE]radius-server shared-key cipher ClePartageeRADIUS123
[AR720-radius-RADIUS_SIEGE]radius-server authentication 192.168.100.70 1812
[AR720-radius-RADIUS_SIEGE]radius-server accounting 192.168.100.70 1813
[AR720-radius-RADIUS_SIEGE]quit
[AR720]aaa
[AR720-aaa]authentication-scheme AUTH_RADIUS
[AR720-aaa-authen-AUTH_RADIUS]authentication-mode radius local
[AR720-aaa-authen-AUTH_RADIUS]quit
[AR720-aaa]domain default
[AR720-aaa-domain-default]authentication-scheme AUTH_RADIUS
[AR720-aaa-domain-default]radius-server RADIUS_SIEGE
[AR720-aaa-domain-default]quit
[AR720-aaa]quit
[AR720]user-interface vty 0 4
[AR720-ui-vty0-4]authentication-mode aaa
[AR720-ui-vty0-4]quit
```

Le `local` en secours (`authentication-mode radius local`) garantit l'accès
si le serveur RADIUS tombe — **ne l'oubliez jamais**, sinon une panne
RADIUS = porte fermée partout.

## 95. Sécuriser les accès distants — durcissement

```vrp
[AR720]acl number 2000
[AR720-acl-basic-2000]rule 5 permit source 192.168.100.0 0.0.0.255
[AR720-acl-basic-2000]rule 10 deny
[AR720-acl-basic-2000]quit
[AR720]user-interface vty 0 4
[AR720-ui-vty0-4]acl 2000 inbound
[AR720-ui-vty0-4]protocol inbound ssh
[AR720-ui-vty0-4]quit
[AR720]undo telnet server enable
[AR720]undo http server enable                  # si le web n'est pas utilisé
[AR720]undo ftp server enable                   # idem FTP
```

Checklist durcissement accès :

```text
[ ] Telnet désactivé (SSH uniquement)
[ ] ACL sur les VTY (plage admin uniquement)
[ ] Mots de passe en irreversible-cipher
[ ] idle-timeout configuré (10 min)
[ ] Banner légal configuré
[ ] Comptes nominatifs (pas de compte générique partagé si possible)
[ ] SNMP v3 ou communauté v2c forte, traps vers supervision
[ ] RADIUS avec secours local
```

## 96. Python + Netmiko sur VRP — prise en main

Netmiko supporte les équipements Huawei via `device_type="huawei"`.

```bash
pip install netmiko
```

```python
from netmiko import ConnectHandler

device = {
    "device_type": "huawei",
    "host": "192.168.10.1",
    "username": "admin",
    "password": "MotDePasseFort123!",   # en prod : variable d'environnement / coffre
    "port": 22,
}

conn = ConnectHandler(**device)
print(conn.send_command("display version"))
print(conn.send_command("display ip interface brief"))
conn.disconnect()
```

Pour la configuration (mode `system-view` géré automatiquement) :

```python
commands = [
    "sysname S310-ETAGE1",
    "vlan 20",
    "description SERVEURS",
]
output = conn.send_config_set(commands)
print(output)
conn.send_command("save")   # NE PAS OUBLIER
conn.disconnect()
```

> ⚠️ Netmiko ne fait pas le `save` à votre place : terminez toujours par
> `send_command("save")` + lecture de la confirmation.

## 97. Envoi de commandes en masse (multi-équipements)

```python
from netmiko import ConnectHandler
from getpass import getpass

PASSWORD = getpass("Mot de passe admin : ")

devices = [
    {"device_type": "huawei", "host": "192.168.10.1", "username": "admin",
     "password": PASSWORD},   # AR720-SIEGE
    {"device_type": "huawei", "host": "192.168.10.2", "username": "admin",
     "password": PASSWORD},   # S310-ETAGE1
    {"device_type": "huawei", "host": "192.168.10.3", "username": "admin",
     "password": PASSWORD},   # USG6000-SIEGE
]

COMMANDS = [
    "ntp-service unicast-server 192.168.100.1",
    "info-center loghost 192.168.100.60",
]

for dev in devices:
    try:
        conn = ConnectHandler(**dev)
        out = conn.send_config_set(COMMANDS)
        conn.send_command("save")
        print(f"[OK] {dev['host']}")
        with open(f"change_{dev['host']}.log", "w") as f:
            f.write(out)
        conn.disconnect()
    except Exception as e:
        print(f"[ERREUR] {dev['host']} : {e}")
```

Bonnes pratiques : un appareil à la fois en dry-run d'abord (`send_command`
de lecture), journaliser chaque sortie, timeout généreux, et **toujours**
un plan de rollback.

## 98. Sauvegarde automatisée des configurations

Script de sauvegarde quotidienne (à mettre en cron sur le serveur admin) :

```python
from netmiko import ConnectHandler
from datetime import date
from getpass import getpass
import os

PASSWORD = getpass()
BACKUP_DIR = "/srv/backups/vrp"
os.makedirs(BACKUP_DIR, exist_ok=True)

devices = {
    "AR720-SIEGE": "192.168.10.1",
    "S310-ETAGE1": "192.168.10.2",
    "USG6000-SIEGE": "192.168.10.3",
}

today = date.today().isoformat()
for name, host in devices.items():
    conn = ConnectHandler(device_type="huawei", host=host,
                         username="admin", password=PASSWORD)
    conn.send_command("screen-length 0 temporary")
    cfg = conn.send_command("display current-configuration")
    path = f"{BACKUP_DIR}/{name}_{today}.cfg"
    with open(path, "w") as f:
        f.write(cfg)
    print(f"[OK] {name} -> {path}")
    conn.disconnect()
```

```cron
# /etc/cron.d/vrp-backup — sauvegarde quotidienne à 2h00
0 2 * * * backup /usr/bin/python3 /opt/scripts/vrp_backup.py >> /var/log/vrp_backup.log 2>&1
```

Rétention : 30 jours glissants + 12 mensuels (règle 3-2-1 : voir le guide
Debian/Ubuntu de Zelef pour la stratégie de sauvegarde globale).

## 99. Ansible et VRP — principes

Ansible n'a pas de module natif Huawei officiel maintenu comme pour
d'autres vendeurs, mais deux approches fonctionnent :

**Approche 1 : `ansible.netcommon.cli_command` / `cli_config`**
(transport `ansible_network_os=huawei` via la collection communautaire
`community.network` — disponibilité à vérifier selon votre version
d'Ansible) :

```yaml
---
- name: Configurer NTP sur les VRP
  hosts: vrp
  gather_facts: false
  vars:
    ansible_connection: network_cli
    ansible_network_os: huawei
    ansible_user: admin
  tasks:
    - name: Appliquer la config NTP
      ansible.netcommon.cli_config:
        config: |
          ntp-service unicast-server 192.168.100.1
          info-center loghost 192.168.100.60
    - name: Sauvegarder
      ansible.netcommon.cli_command:
        command: save
        prompt: "(y/n)"
        answer: "y"
```

**Approche 2 : template + Netmiko via un module custom** — souvent plus
fiable : générez les configs avec Jinja2 côté Ansible, poussez avec un
script Python Netmiko (section 97).

## 100. Inventaire et faits — `display` utiles aux scripts

```bash
# Faits minimaux à collecter par équipement (pour CMDB / inventaire)
display version              # modèle + version VRP + uptime
display device elabel        # numéro de série
display current-configuration # config complète
display startup              # fichiers de boot
display interface brief      # état des ports
```

