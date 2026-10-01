---
id: collect-261001-rattrapage/rattrapage/netbox-guide-9
title: "NetBox — Guide complet : source de vérité IPAM / DCIM"
domain: rattrapage
role: reference
task: reference
actors: ["Meta"]
dates: []
keywords: ["arr", "ethernet", "memory"]
source: docs/RAG/collect-261001-rattrapage/netbox_guide.md
source_anchor: ""
source_lines: [1667, 1887]
sha256: 2d5ac0d05bb69a527b7d82b680068b81672a11d9fc619e567b6e6dcf91992c4e
---

# NetBox — Guide complet : source de vérité IPAM / DCIM

> 💡 **TLS avec PKI interne** : si votre certificat est signé par votre CA d'entreprise,
> ajoutez-la au bundle de confiance du système (`/usr/local/share/ca-certificates/` +
> `update-ca-certificates`), ou passez `nb.http_session.verify = "/chemin/ca.pem"`.
> Ne mettez JAMAIS `verify = False` en production (attaque MITM possible).

---

## 55. API REST : créer un préfixe et une VM en Python (exemple complet)

```python
#!/usr/bin/env python3
"""Provisionner un préfixe + une VM via l'API NetBox."""
import os
import pynetbox

nb = pynetbox.api("https://netbox.lan-entreprise.fr", token=os.environ["NETBOX_TOKEN"])

# 1. Créer un préfixe enfant dans le container 10.10.0.0/16
prefix = nb.ipam.prefixes.create({
    "prefix": "10.10.60.0/24",
    "site": "siege-paris11",
    "vlan": {"vid": 60, "name": "SRV-Test"},   # le VLAN doit exister
    "status": "active",
    "role": "Serveurs",
    "tenant": "dsi",
    "description": "VLAN 60 — serveurs de test (provisionné par API)",
})
print("Préfixe créé :", prefix.prefix, "| id:", prefix.id)

# 2. Réserver la première IP utilisable pour la passerelle
gw = nb.ipam.ip_addresses.create({
    "address": "10.10.60.1/24",
    "status": "active",
    "description": "Passerelle VLAN 60",
})
print("IP passerelle :", gw.address)

# 3. Créer une VM dans le cluster Proxmox
vm = nb.virtualization.virtual_machines.create({
    "name": "VM-TEST-01",
    "cluster": "pve-siege",
    "status": "active",
    "role": "Applicatif",
    "tenant": "dsi",
    "vcpus": 2,
    "memory": 4096,
    "disk": 40,
    "platform": "debian-12",
    "description": "VM de test — créée par script",
})
print("VM créée :", vm.name)

# 4. Interface + IP pour la VM
iface = nb.virtualization.interfaces.create({
    "virtual_machine": vm.id, "name": "eth0", "enabled": True,
})
ip = nb.ipam.ip_addresses.create({
    "address": "10.10.60.10/24", "status": "active",
    "assigned_object_type": "virtualization.vminterface",
    "assigned_object_id": iface.id,
    "description": "Interface de production",
})
# Définir l'IP primaire de la VM (pratique : affichée dans les listes)
vm.primary_ip4 = ip.id
vm.save()
print("VM provisionnée avec IP :", ip.address)
```

Ce script est le patron de tous vos provisionnements : **créer dans NetBox D'ABORD,
déployer ensuite** — jamais l'inverse.

---

## 56. API REST : synchroniser un serveur DHCP (exemple d'usage avancé)

Exemple : générer la configuration des réservations DHCP (ISC DHCPd) depuis NetBox.

```python
#!/usr/bin/env python3
"""Génère les host declarations ISC-DHCP depuis les IP 'reserved' d'un préfixe."""
import os
import pynetbox

nb = pynetbox.api("https://netbox.lan-entreprise.fr", token=os.environ["NETBOX_TOKEN"])

print("# Généré depuis NetBox — ne pas éditer à la main")
for ip in nb.ipam.ip_addresses.filter(parent="10.10.20.0/24", status="reserved"):
    # On ne génère que les IP assignées à une interface avec une MAC connue
    obj = ip.assigned_object
    mac = getattr(obj, "mac_address", None)
    if not mac:
        continue
    hostname = (ip.dns_name or ip.description or f"host-{ip.id}").split(".")[0]
    print(f"host {hostname} {{")
    print(f"  hardware ethernet {mac.lower()};")
    print(f"  fixed-address {str(ip.address).split('/')[0]};")
    print("}")
```

Chaîne d'automatisation complète :

1. Le technicien réserve l'IP dans NetBox (statut `Réservé`, MAC renseignée).
2. Ce script tourne chaque nuit (cron) et régénère `/etc/dhcp/dhcpd.conf`.
3. Le service DHCP recharge la conf. **Zéro saisie en double**, zéro dérive.

> ⚠️ Versionnez le fichier généré (git) : en cas de génération corrompue, le
> `dhcpd -t` (test de syntaxe) avant rechargement est obligatoire dans le script cron.

---

## 57. Webhooks : réagir aux événements NetBox

Les **webhooks** envoient un POST HTTP à chaque création/modification/suppression
d'objet : le pont entre NetBox et vos outils (Slack/Teams, GLPI, scripts maison).

**Admin > Webhooks > + Ajouter** :

| Champ | Exemple |
|---|---|
| Nom | `notify-teams-device-change` |
| URL | `https://logic-app-entreprise.azure.com/...` |
| Méthode HTTP | `POST` |
| Content type | `application/json` |
| Secret | `CHANGEZ_MOI_Webhook_Exemple` (signature HMAC) |
| Événements | mise à jour d'équipement |

Exemple de récepteur minimal (Python/Flask) qui journalise les changements :

```python
from flask import Flask, request
import hmac, hashlib

app = Flask(__name__)
SECRET = b"CHANGEZ_MOI_Webhook_Exemple"

@app.post("/netbox-hook")
def hook():
    sig = request.headers.get("X-Hook-Signature", "")
    expected = hmac.new(SECRET, request.data, hashlib.sha512).hexdigest()
    if not hmac.compare_digest(sig, expected):
        return "signature invalide", 403
    event = request.json
    print(f"[{event['event']}] {event['model']} -> {event['data']['name']}")
    return "ok", 200
```

Cas d'usage : notifier l'équipe réseau sur Teams à chaque ajout d'équipement,
déclencher un inventaire Ansible à chaque modification d'IP, alimenter un CMDB externe.

> ⚠️ Les webhooks sont traités par `netbox-rq` (file `high`) : si le service est
> arrêté, les événements s'accumulent dans Redis. Surveillez la profondeur des files
> (section 67).

---

## 58. Scripts personnalisés (custom scripts)

Les **scripts** s'exécutent DANS NetBox (Admin > Scripts) : formulaires web qui
appellent l'API interne. Parfait pour les actions récurrentes d'exploitation.

Emplacement : `/opt/netbox/netbox/netbox/scripts/` (ou un répertoire configuré via
`SCRIPTS_ROOT`). Exemple : script « Réserver les N prochaines IP d'un préfixe » :

```python
from netbox.scripts import Script, StringVar, IntegerVar
import ipaddress

class ReserveIPs(Script):
    class Meta:
        name = "Réserver des IP"
        description = "Réserve les N prochaines IP libres d'un préfixe."

    prefix = StringVar(description="Préfixe CIDR (ex. 10.10.20.0/24)")
    count = IntegerVar(description="Nombre d'IP à réserver", default_value=5)
    description = StringVar(description="Description commune", required=False)

    def run(self, data, commit):
        net = ipaddress.ip_network(data["prefix"])
        # Adresses déjà connues de NetBox dans ce préfixe
        used = {str(ip.address.ip) for ip in
                self.get_ips_in_prefix(data["prefix"])}
        reserved = []
        for host in net.hosts():
            if str(host) in used:
                continue
            if commit:
                # création réelle via l'ORM
                from ipam.models import IPAddress
                IPAddress.objects.create(
                    address=f"{host}/{net.prefixlen}",
                    status="reserved",
                    description=data.get("description") or "Réservée par script",
                )
            reserved.append(str(host))
            if len(reserved) >= data["count"]:
                break
        self.log_success(f"{len(reserved)} IP réservées : {', '.join(reserved)}")

    def get_ips_in_prefix(self, prefix):
        from ipam.models import IPAddress
        return IPAddress.objects.filter(
            address__net_contained=str(ipaddress.ip_network(prefix)))
```

> 💡 Les scripts tournent avec les droits de l'utilisateur qui les lance et
> s'exécutent d'abord en **mode test** (`commit=False`) : on voit ce qui VA se passer
> avant de l'appliquer. Idéal pour déléguer des actions aux techniciens sans leur
> donner un accès API brut.

---

## 59. Rapports personnalisés (custom reports)

Les **rapports** vérifient la cohérence des données et produisent des tableaux de bord
d'exploitation. Emplacement : `/opt/netbox/netbox/netbox/reports/` (`REPORTS_ROOT`).

Exemple : rapport « Équipements sans adresse IP de management » (classique !) :

```python
from netbox.reports import Report

