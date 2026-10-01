---
id: collect-261001-rattrapage/rattrapage/netbox-guide-10
title: "NetBox — Guide complet : source de vérité IPAM / DCIM"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "arr", "incident"]
source: docs/RAG/collect-261001-rattrapage/netbox_guide.md
source_anchor: ""
source_lines: [1888, 2109]
sha256: 713480bcc4f00d2d6196ef242d1c053f7ff80aa6ee203c5d0b863fa450e54d4b
---

# NetBox — Guide complet : source de vérité IPAM / DCIM

class ManagementIPReport(Report):
    description = "Liste les équipements actifs sans IP de management assignée."

    def test_device_mgmt_ip(self):
        from dcim.models import Device
        bad = []
        for dev in Device.objects.filter(status="active").exclude(
                role__name__in=["Panneau-Brassage"]):
            has_mgmt = dev.interfaces.filter(
                mgmt_only=True, ip_addresses__isnull=False).exists()
            if not has_mgmt:
                bad.append(dev)
        if bad:
            self.log_failure(
                f"{len(bad)} équipements sans IP de management : "
                + ", ".join(d.name for d in bad[:20])
            )
        else:
            self.log_success("Tous les équipements actifs ont une IP de management.")

    def test_rack_utilization(self):
        from dcim.models import Rack
        for rack in Rack.objects.filter(status="active"):
            pct = rack.get_utilization()
            if pct >= 90:
                self.log_warning(f"Baie {rack.name} remplie à {pct} % — prévoir extension.")
        self.log_success("Contrôle du remplissage des baies terminé.")
```

Idées de rapports pour un service systèmes/énergies :

- Bilan de puissance par baie vs calibre onduleur/PDU (section 26.3),
- Garanties expirant dans < 90 jours (champ `date_fin_garantie`, section 47),
- Engagements opérateurs arrivant à échéance (section 27),
- Préfixes > 80 % d'utilisation,
- Câbles sans étiquette.

---

## 60. Jobs, files d'attente et exécution différée

NetBox 4.x exécute les tâches longues via **django-rq** (files `high`, `default`, `low`),
traitées par `netbox-rq.service` (section 14) :

| File | Contenu typique |
|---|---|
| `high` | webhooks |
| `default` | scripts personnalisés, imports CSV volumineux |
| `low` | rapports planifiés, nettoyage |

Surveillance via l'interface : **Admin > Files d'attente RQ** (ou en CLI) :

```bash
# État des files dans Redis
redis-cli -n 0 llen rq:queue:high
redis-cli -n 0 llen rq:queue:default
# Voir les jobs en échec
sudo -iu netbox /opt/netbox/venv/bin/python \
  /opt/netbox/netbox/netbox/manage.py rqinfo
```

> ⚠️ Une file `default` qui grossit sans se vider = `netbox-rq` arrêté ou planté.
> Symptômes côté utilisateurs : scripts qui restent « en cours » indéfiniment,
> webhooks jamais envoyés. Redémarrez le service et surveillez les logs.

---

## 61. Intégration Ansible : inventaire dynamique depuis NetBox

Le plugin d'inventaire `netbox.netbox.nb_inventory` fait de NetBox LA source de
l'inventaire Ansible : plus de fichier `hosts.ini` à maintenir.

```bash
# Sur le contrôleur Ansible
ansible-galaxy collection install netbox.netbox
pip install pynetbox
```

`inventory/netbox.yml` :

```yaml
plugin: netbox.netbox.nb_inventory
api_endpoint: https://netbox.lan-entreprise.fr
token: "{{ lookup('env', 'NETBOX_TOKEN') }}"   # jamais en clair !
validate_certs: true

# Groupes construits depuis les champs NetBox
group_by:
  - device_roles      # ex. group_switch-acces
  - sites             # ex. group_siege-paris11
  - tenants
plumb_graph: true     # expose les relations (voisins câblés)

# Variables d'hôte depuis NetBox
compose:
  ansible_host: primary_ip4.address | default(interfaces[0].ip_addresses[0].address, '')
  ansible_network_os: >-
    {'Cisco IOS': 'ios', 'Arista EOS': 'eos'}.get(platform.name, 'ios')

# On ne veut que les équipements actifs, joignables
filters:
  status: active
  has_primary_ip: true
```

Test :

```bash
export NETBOX_TOKEN="CHANGEZ_MOI_Token_Exemple"
ansible-inventory -i inventory/netbox.yml --graph | head -30
ansible -i inventory/netbox.yml group_switch-acces -m ping
```

> 💡 **Chaîne vertueuse** : un nouveau switch est créé dans NetBox (avec son IP et
> son rôle) → il apparaît automatiquement dans l'inventaire Ansible → le playbook
> de provisionnement le configure. Le provisionnement « NetBox d'abord » (section 55)
> prend ici tout son sens.

---

## 62. Intégration Zabbix : exporter vers la supervision

Deux stratégies complémentaires :

**A. Création automatique des hôtes Zabbix depuis NetBox** (script Python + API Zabbix) :

```python
#!/usr/bin/env python3
"""Crée dans Zabbix les équipements NetBox taggés 'supervise-zabbix'."""
import os
import pynetbox
from pyzabbix import ZabbixAPI

nb = pynetbox.api("https://netbox.lan-entreprise.fr", token=os.environ["NETBOX_TOKEN"])
zx = ZabbixAPI("https://zabbix.lan-entreprise.fr")
zx.login(os.environ["ZABBIX_USER"], os.environ["ZABBIX_PASSWORD"])

for dev in nb.dcim.devices.filter(tag="supervise-zabbix", status="active"):
    if not dev.primary_ip4:
        print(f"SKIP {dev.name} : pas d'IP primaire")
        continue
    ip = str(dev.primary_ip4.address).split("/")[0]
    # mapping rôle NetBox -> modèle Zabbix
    templates = {"Switch-Acces": "Cisco IOS SNMP",
                 "Serveur": "Linux by Zabbix agent",
                 "Onduleur": "UPS SNMP"}.get(dev.device_role.name, "ICMP Ping")
    print(f"CREATE {dev.name} ({ip}) template={templates}")
    # zx.host.create(...) — à adapter à votre version de Zabbix
```

**B. Zabbix → NetBox** : remonter les alertes vers un champ personnalisé
(`alarme_en_cours`) via webhook Zabbix → API NetBox : la fiche équipement affiche
l'état de supervision en temps réel.

> ⚠️ Pour les onduleurs (votre métier !) : supervisez via SNMP les variables
> `upsBatteryStatus`, `upsEstimatedChargeRemaining`, `upsSecondsOnBattery` et
> faites pointer la fiche NetBox de l'onduleur vers son hôte Zabbix (champ
> personnalisé `zabbix_hostid`). En incident énergie, on part de NetBox (chaîne
> électrique, section 26) et on bascule sur Zabbix (métriques temps réel).

---

## 63. Plugins utiles (écosystème NetBox 4.x)

Les plugins s'installent via pip dans le venv et se déclarent dans `configuration.py`
(`PLUGINS = [...]`). Sélection éprouvée :

| Plugin | Usage | Remarque |
|---|---|---|
| `netbox-topology-views` | cartes topologiques du réseau | très parlant en réunion |
| `netbox-bgp` | sessions BGP, communities | si vous faites du BGP |
| `netbox-dns` | zones DNS gérées dans NetBox | alternative au DNS dispersé |
| `netbox-inventory` | suivi des assets (garanties, mouvements) | lien avec l'inventaire comptable |
| `netbox-qrcode` | QR codes sur les fiches (baies, équipements) | flashez la baie → sa fiche ! |
| `netbox-attachments` | pièces jointes enrichies | selon version |

Exemple d'activation :

```python
PLUGINS = [
    "netbox_topology_views",
    "netbox_qrcode",
]
```

```bash
sudo -iu netbox /opt/netbox/venv/bin/pip install netbox-topology-views netbox-qrcode
sudo -iu netbox /opt/netbox/venv/bin/python /opt/netbox/netbox/netbox/manage.py migrate
systemctl restart netbox netbox-rq
```

> ⚠️ **Règle de prudence** : chaque plugin est du code tiers exécuté avec les droits
> de NetBox. N'installez que des plugins maintenus, compatibles avec votre version
> 4.x (vérifiez la matrice de compatibilité sur leur dépôt), et testez-les en LAB
> avant la production. Un plugin abandonné peut bloquer une montée de version.

---

## 64. Sauvegarde : base PostgreSQL (`pg_dump`)

La donnée métier vit dans PostgreSQL : la sauvegarde est non négociable, quotidienne,
et **testée**.

```bash
#!/bin/bash
# /opt/netbox/backup-netbox-db.sh  (propriétaire root, 700)
set -euo pipefail
BACKUP_DIR="/var/backups/netbox"
RETENTION=30
mkdir -p "$BACKUP_DIR"
DATE=$(date +%Y%m%d-%H%M%S)

# Dump complet au format custom (restauration sélective possible)
sudo -u postgres pg_dump -Fc netbox > "$BACKUP_DIR/netbox-$DATE.dump"
# + dump SQL texte (lisible, dépannage)
sudo -u postgres pg_dump netbox > "$BACKUP_DIR/netbox-$DATE.sql"

gzip -f "$BACKUP_DIR/netbox-$DATE.sql"
find "$BACKUP_DIR" -name 'netbox-*' -mtime +$RETENTION -delete
echo "OK $DATE : $(du -h "$BACKUP_DIR/netbox-$DATE.dump" | cut -f1)"
```

