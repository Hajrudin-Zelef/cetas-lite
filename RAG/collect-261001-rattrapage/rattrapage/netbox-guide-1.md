---
id: collect-261001-rattrapage/rattrapage/netbox-guide-1
title: "NetBox — Guide complet : source de vérité IPAM / DCIM"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "datacenter", "open source"]
source: docs/RAG/collect-261001-rattrapage/netbox_guide.md
source_anchor: ""
source_lines: [1, 155]
sha256: 65274c02aab607fc812643942ea59ccf4d883bba426f60f98f786d8fe25e6d84
---

# NetBox — Guide complet : source de vérité IPAM / DCIM

> **Public :** chefs de service systèmes, administrateurs réseau & systèmes, responsables d'infrastructure.
> **Objectif :** installer, configurer, exploiter et maintenir NetBox comme référentiel unique
> (source de vérité) de l'infrastructure : adressage IP (IPAM), baies et équipements (DCIM),
> câblage, alimentation, circuits opérateurs, virtualisation.
> **Versions couvertes :** NetBox 4.x (4.0 → 4.2+), Debian 12/13, Ubuntu 22.04/24.04 LTS,
> PostgreSQL 15/16, Redis 7.
> **Avertissement :** les commandes et chemins correspondent à NetBox 4.x. Les versions 2.x/3.x
> diffèrent sur plusieurs points (Python 3.8+, arborescence, plugins). Vérifiez toujours la
> [documentation officielle](https://docs.netbox.dev) pour votre version exacte.

---

## Sommaire

- Concepts : source de vérité, IPAM, DCIM
- Installation pas à pas (Debian/Ubuntu : PostgreSQL + Redis + NetBox + gunicorn + nginx + systemd)
- Configuration (`configuration.py`)
- Premier lancement et organisation (tenants, régions, sites)
- DCIM : baies, types d'équipements, équipements, interfaces, câblage, alimentation (PDU/onduleurs), circuits
- IPAM : VRF, préfixes, plages IP, adresses IP, VLAN
- Virtualisation, utilisateurs, permissions, sécurité, HTTPS
- API REST (curl + Python), webhooks, scripts et rapports personnalisés
- Intégrations : Ansible (inventaire dynamique), Zabbix (export)
- Sauvegarde (pg_dump + media), restauration, montée de version, supervision
- Dépannage : cas concrets et 10+ erreurs classiques
- Bonnes pratiques (nommage, tenancy, « documenter avant de câbler »)
- Cas pratiques commentés (site complet : baie, onduleur, switchs, serveurs)
- Pense-bête de poche, glossaire, quiz, pour aller plus loin

**Convention de ce guide :**

| Marqueur | Signification |
|---|---|
| `netbox` (shell) | commande exécutée en tant qu'utilisateur `netbox` |
| `root` / `sudo` | commande avec privilèges administrateur |
| `CHANGEZ_MOI_...` | valeur d'exemple à remplacer — jamais utilisée telle quelle en production |
| ⚠️ | point de vigilance (sécurité, perte de données) |

---

## 1. NetBox, c'est quoi ? La « source de vérité »

NetBox est une application web open source (licence Apache 2.0, éditée par NetBox Labs)
qui centralise la documentation d'une infrastructure réseau et datacenter. On l'appelle
**source de vérité** (*source of truth*) : c'est l'endroit unique où l'on consigne
« ce qui existe vraiment » — chaque adresse IP, chaque équipement, chaque câble,
chaque baie.

Sans source de vérité, l'information vit dans des tableurs éparpillés, des têtes
d'administrateurs et des post-it : c'est la garantie des conflits d'IP, des baies
« pleines » qui ne le sont pas, et des interventions à l'aveugle. Avec NetBox :

- **Un seul référentiel** consultable par toute l'équipe (et par les outils, via l'API).
- **Cohérence garantie** : NetBox refuse les doublons (deux fois la même IP, deux fois
  le même nom d'équipement dans un site).
- **Traçabilité** : chaque création/modification/suppression est journalisée (qui, quand, quoi).
- **Automatisation** : Ansible, Zabbix, scripts Python lisent NetBox au lieu de fichiers statiques.

> 💡 **Pour un chef de service systèmes** : NetBox devient le socle de vos procédures
> d'exploitation. « C'est documenté dans NetBox » remplace « demande à Karim, il sait ».
> C'est aussi un formidable outil de management : l'état du parc est visible, mesurable,
> auditable.

---

## 2. Les deux grands modules : IPAM et DCIM

| Module | Sigle | Contenu | Exemples d'objets |
|---|---|---|---|
| **IPAM** | IP Address Management | Tout l'adressage logique | VRF, préfixes, plages IP, adresses IP, VLAN, ASN |
| **DCIM** | Data Center Infrastructure Management | Tout le physique | Régions, sites, baies, équipements, interfaces, câbles, alimentations, circuits |

Les deux modules sont **liés** : une adresse IP (IPAM) est assignée à une interface
d'un équipement (DCIM), elle-même installée dans une baie d'un site. Cette liaison
logique ↔ physique est la vraie force de NetBox : depuis une IP, on remonte au
serveur, à sa baie, à son site, à son PDU et à son onduleur.

Autres modules utiles au quotidien :

- **Circuits** : liaisons opérateurs/FAI (type de circuit, débit, identifiants).
- **Virtualisation** : clusters, machines virtuelles, interfaces virtuelles.
- **Tenancy** : tenants (clients, entités, départements) pour cloisonner la documentation.
- **Secrets** : ⚠️ le module natif a été retiré ; utilisez un gestionnaire dédié (Vault).

---

## 3. Architecture technique de NetBox 4.x

```
                  ┌─────────────┐
                  │   Nginx     │  :443 HTTPS (reverse proxy)
                  └──────┬──────┘
                         │ proxy HTTP local
                  ┌──────▼──────┐
                  │  Gunicorn   │  serveurs WSGI Python (netbox.service)
                  └──────┬──────┘
                         │
            ┌────────────┼────────────┐
            ▼            ▼            ▼
     ┌────────────┐ ┌────────┐ ┌────────────┐
     │ PostgreSQL │ │ Redis  │ │  Fichiers  │
     │  (données) │ │(cache,  │ │  media/    │
     │            │ │ files)   │ │(scripts,   │
     └────────────┘ └────────┘ │ rapports)  │
                               └────────────┘
```

- **PostgreSQL** : base de données relationnelle, TOUTE la donnée métier y vit.
  NetBox 4.x exige PostgreSQL 12 minimum (15/16 recommandés).
- **Redis** : cache applicatif + file d'attente des tâches (cacheops, django-rq).
- **Gunicorn** : serveur d'application Python qui exécute le code Django de NetBox.
- **Nginx** : reverse proxy en frontal (TLS, en-têtes, fichiers statiques).
- **systemd** : supervise `netbox.service` (gunicorn) et `netbox-rq.service` (travailleurs de tâches).

> ⚠️ **Sauvegarde = PostgreSQL + dossier `media/`**. Redis est un cache : il se reconstruit
> seul, inutile de le sauvegarder.

---

## 4. Prérequis matériels et dimensionnement

| Taille d'infrastructure | vCPU | RAM | Disque | Remarques |
|---|---|---|---|---|
| Lab / PME (< 500 équipements) | 2 | 4 Go | 40 Go | Une seule VM suffit |
| Moyenne (500–5 000 équipements) | 4 | 8 Go | 100 Go | PostgreSQL sur disque rapide |
| Grande (> 5 000, API intensive) | 8 | 16 Go | 200 Go+ | Séparer PostgreSQL si besoin |

- **OS** : Debian 12/13 ou Ubuntu 22.04/24.04 LTS, installation minimale, à jour.
- **Réseau** : IP fixe, nom DNS interne (ex. `netbox.lan-entreprise.fr`), accès HTTPS.
- **Python** : 3.10 minimum pour NetBox 4.x (3.12 recommandé). Ne JAMAIS utiliser le Python
  système pour autre chose que ce qu'il faut : NetBox tourne dans un **virtualenv** dédié.

Checklist avant de commencer :

- [ ] Serveur/VM provisionné, OS à jour (`apt update && apt full-upgrade`)
- [ ] Nom d'hôte (FQDN) configuré, résolution DNS directe et inverse OK
- [ ] Accès SSH par clé, compte `sudo` fonctionnel
- [ ] Flux réseau : 443 entrant (utilisateurs), 22 (admin) ; sortant : dépôts APT, PyPI si besoin
- [ ] Sauvegarde de la VM planifiée (snapshot initial avant l'installation)

---

## 5. Préparation du serveur Debian/Ubuntu

Toutes les commandes ci-dessous sont exécutées en `root` (ou via `sudo`).

```bash
# Mise à jour du système
apt update && apt full-upgrade -y

# Paquets de base
apt install -y sudo curl wget git vim htop net-tools dnsutils

