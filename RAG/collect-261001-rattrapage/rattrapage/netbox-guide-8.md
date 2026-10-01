---
id: collect-261001-rattrapage/rattrapage/netbox-guide-8
title: "NetBox — Guide complet : source de vérité IPAM / DCIM"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-rattrapage/netbox_guide.md
source_anchor: ""
source_lines: [1464, 1666]
sha256: 53bde40df7301350e7515f394d719496310d808aa82ad0ea66dc508a97b803c9
---

# NetBox — Guide complet : source de vérité IPAM / DCIM

> 💡 Règle : un champ personnalisé doit servir à un **usage réel** (filtre, rapport,
> alerte). Un champ « au cas où » que personne ne remplit = du bruit. Revue annuelle
> des champs personnalisés conseillée.

---

## 48. Étiquettes (tags)

Les **tags** (Admin > Étiquettes) sont des marqueurs libres, transverses aux objets :
`a-migrer`, `supervise-zabbix`, `contrat-gold`, `a-verifier`.

Différences avec les champs personnalisés :

| | Tag | Champ personnalisé |
|---|---|---|
| Nature | marqueur libre, multiple | champ typé, défini à l'avance |
| Usage | tri rapide, lots temporaires | donnée structurée durable |
| Exemple | `incident-2026-09` sur 12 équipements | `date_fin_garantie` sur tous les serveurs |

Bonnes pratiques : couleurs sobres, noms en minuscules avec tirets, et surtout
**nettoyage régulier** : un tag temporaire (`a-verifier`) vieux de 2 ans n'est plus
temporaire, c'est de la dette documentaire.

---

## 49. Modèles (templates) d'export et import en masse

- **Import CSV** : chaque liste propose **Importer** : préparez un CSV avec les
  colonnes exactes (téléchargez d'abord un export pour voir le format), importez,
  corrigez les erreurs signalées ligne par ligne. Idéal pour : 200 VLAN, 48 ports
  d'un panneau, un parc d'imprimantes.
- **Export** : CSV natif (section 41) ou via l'API pour les gros volumes.

Exemple de CSV d'import de préfixes (en-têtes exacts requis) :

```csv
prefix,status,vrf,site,vlan,role,tenant,description
10.20.1.0/24,active,,AGENCE-Lyon,120,LAN-Utilisateurs,DSI,"LAN agence Lyon"
10.20.2.0/24,active,,AGENCE-Lyon,121,WiFi,DSI,"WiFi agence Lyon"
```

> ⚠️ L'import est puissant et irréversible en masse : testez TOUJOURS sur 3 lignes
> dans un tenant `Bac-A-Sable` avant d'importer 500 lignes en production.

---

## 50. Tâches planifiées et maintenance interne

NetBox s'auto-entretient partiellement ; à vous de planifier le reste :

| Tâche | Fréquence | Commande / moyen |
|---|---|---|
| Nettoyage des sessions expirées | hebdo | `python netbox/manage.py clearsessions` |
| Purge du changelog (si rétention) | hebdo | automatique via `CHANGELOG_RETENTION` |
| Vérification disque BDD | mensuel | `SELECT pg_size_pretty(pg_database_size('netbox'));` |
| Test de restauration sauvegarde | trimestriel | section 65 |
| Revue des comptes inactifs | trimestriel | Admin > Utilisateurs (filtre dernière connexion) |
| Revue des permissions | semestriel | avec le RSSI / la direction |

Automatisez avec cron (utilisateur `netbox`) :

```bash
# /etc/cron.d/netbox-maintenance
0 3 * * 0 netbox /opt/netbox/venv/bin/python /opt/netbox/netbox/netbox/manage.py clearsessions
```

---

## 51. API REST : principes et authentification par token

Tout ce qui se fait dans l'interface se fait via l'**API REST** : c'est le socle de
l'automatisation. Base URL : `https://netbox.lan-entreprise.fr/api/`.

1. Créez un token : profil (en haut à droite) > **API Tokens** > + Ajouter.
   - Donnez-lui une description (`Script inventaire Ansible — svc-ansible`).
   - **Permissions d'écriture** : décochez si le script ne fait que lire (principe
     du moindre privilège).
   - Date d'expiration : mettez-en une pour les tokens temporaires.
2. Utilisez-le dans l'en-tête HTTP : `Authorization: Token <votre_token>`.

> ⚠️ Un token = un mot de passe. Ne le commitez JAMAIS dans git en clair, ne le
> collez pas dans un ticket. Stockez-le dans un gestionnaire de secrets / variable
> d'environnement (`NETBOX_TOKEN`). En cas de doute : révoquez et régénérez.

L'API est **auto-documentée** : `https://netbox.lan-entreprise.fr/api/docs/`
(Swagger/OpenAPI) — chaque endpoint, chaque champ, avec exemples. C'est la
référence en cas de doute sur ce guide.

---

## 52. API REST : exemples `curl` (lecture)

```bash
export NETBOX_URL="https://netbox.lan-entreprise.fr"
export NETBOX_TOKEN="CHANGEZ_MOI_Token_Exemple_ABCDEF1234567890"

# Lister les sites (paginé : 50 objets par page par défaut)
curl -s -H "Authorization: Token $NETBOX_TOKEN" \
     -H "Accept: application/json" \
     "$NETBOX_URL/api/dcim/sites/" | python3 -m json.tool | head -40

# Chercher un préfixe précis
curl -s -H "Authorization: Token $NETBOX_TOKEN" \
     "$NETBOX_URL/api/ipam/prefixes/?prefix=10.10.20.0/24" | python3 -m json.tool

# Filtrer : équipements actifs du site Siège
curl -s -H "Authorization: Token $NETBOX_TOKEN" \
     "$NETBOX_URL/api/dcim/devices/?site=siege-paris11&status=active" \
     | python3 -c "import json,sys; d=json.load(sys.stdin); print(d['count'], 'équipements')"

# Suivre la pagination (exemple : récupérer TOUTES les adresses IP)
curl -s -H "Authorization: Token $NETBOX_TOKEN" \
     "$NETBOX_URL/api/ipam/ip-addresses/?limit=1000" -o /tmp/ips.json
```

Points d'API les plus utilisés :

| Endpoint | Contenu |
|---|---|
| `/api/dcim/sites/` | sites |
| `/api/dcim/racks/` | baies |
| `/api/dcim/devices/` | équipements |
| `/api/dcim/interfaces/` | interfaces |
| `/api/dcim/cables/` | câbles |
| `/api/ipam/vrfs/` | VRF |
| `/api/ipam/prefixes/` | préfixes |
| `/api/ipam/ip-addresses/` | adresses IP |
| `/api/ipam/vlans/` | VLAN |
| `/api/virtualization/virtual-machines/` | VM |
| `/api/circuits/circuits/` | circuits |
| `/api/tenancy/tenants/` | tenants |

---

## 53. API REST : écriture avec `curl` (création)

```bash
# Créer un VLAN (notez le Content-Type JSON)
curl -s -X POST -H "Authorization: Token $NETBOX_TOKEN" \
     -H "Content-Type: application/json" \
     -d '{"name": "SRV-Production", "vid": 110, "site": "siege-paris11",
          "status": "active", "description": "VLAN serveurs prod"}' \
     "$NETBOX_URL/api/ipam/vlans/" | python3 -m json.tool

# Créer une adresse IP (le masque fait partie de la valeur)
curl -s -X POST -H "Authorization: Token $NETBOX_TOKEN" \
     -H "Content-Type: application/json" \
     -d '{"address": "10.10.10.99/24", "status": "reserved",
          "description": "Réservée — futur cluster"}' \
     "$NETBOX_URL/api/ipam/ip-addresses/" | python3 -m json.tool

# Modifier (PATCH) la description d'un préfixe existant (id=12)
curl -s -X PATCH -H "Authorization: Token $NETBOX_TOKEN" \
     -H "Content-Type: application/json" \
     -d '{"description": "LAN utilisateurs — bâtiment A (maj sept 2026)"}' \
     "$NETBOX_URL/api/ipam/prefixes/12/" | python3 -m json.tool

# Supprimer (DELETE) — irréversible, à manier avec précaution
# curl -s -X DELETE -H "Authorization: Token $NETBOX_TOKEN" \
#      "$NETBOX_URL/api/ipam/ip-addresses/456/"
```

> ⚠️ L'API respecte les permissions du compte propriétaire du token : un token
> « lecture seule » renverra `403 Forbidden` en écriture. C'est une sécurité, pas un bug.

---

## 54. API REST : Python avec `pynetbox` (installation et bases)

`pynetbox` est le client Python officiel : bien plus agréable que `curl` pour scripter.

```bash
# Dans un venv dédié à vos scripts (pas celui de NetBox !)
python3 -m venv ~/venv-scripts && source ~/venv-scripts/bin/activate
pip install pynetbox
```

```python
#!/usr/bin/env python3
"""Connexion de base à NetBox et lecture d'objets."""
import os
import pynetbox

nb = pynetbox.api(
    "https://netbox.lan-entreprise.fr",
    token=os.environ["NETBOX_TOKEN"],
)
nb.http_session.verify = True   # vérification TLS (certificat interne : voir encadré)

# Lister les sites actifs
for site in nb.dcim.sites.filter(status="active"):
    print(f"{site.name:20s} {site.region.name if site.region else '-'}")

# Chercher un équipement par nom
sw = nb.dcim.devices.get(name="SIEGE-SW-A01-01")
print(sw.name, "->", sw.device_type.model, "| baie:", sw.rack.name, "| U:", sw.position)

# Toutes les IP d'un préfixe, avec leur statut
prefix = nb.ipam.prefixes.get(prefix="10.10.20.0/24")
for ip in nb.ipam.ip_addresses.filter(parent=str(prefix.prefix)):
    print(f"{ip.address}  [{ip.status.value}]  {ip.description or ''}")
```

