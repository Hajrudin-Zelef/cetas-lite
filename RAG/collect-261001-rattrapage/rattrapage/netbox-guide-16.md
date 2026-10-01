---
id: collect-261001-rattrapage/rattrapage/netbox-guide-16
title: "NetBox — Guide complet : source de vérité IPAM / DCIM"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-26"]
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/netbox_guide.md
source_anchor: ""
source_lines: [3020, 3081]
sha256: 5679d77568c75a83845d957467d2543227cb09efc2f046388fd606c986c3218b
---

# NetBox — Guide complet : source de vérité IPAM / DCIM

**R5.** Une **permission** avec **contrainte** JSON `{"site": ["AGENCE-Lyon"]}` sur les
modèles `dcim.cable` (ajout/modification), assignée à son **groupe** (jamais en
direct à l'utilisateur), puis testée avec un compte du groupe (voir section 43).

**R6.** Dans le dossier `media/` (`/opt/netbox/netbox/netbox/media/`). La sauvegarde
complète = `pg_dump` de PostgreSQL **+** archive de `media/` **+** `configuration.py`
(voir sections 64–65). Sans `media/`, les fiches perdent leurs photos et documents.

**R7.** Le plugin d'inventaire dynamique `netbox.netbox.nb_inventory` : l'inventaire
Ansible est généré à la volée depuis NetBox, donc tout équipement créé dans NetBox
(avec IP et rôle) devient automatiquement configurable par Ansible — zéro fichier
`hosts.ini` à maintenir (voir section 61).

**R8.** Grâce à la modélisation de la chaîne électrique (section 26) : depuis la fiche
de l'onduleur, l'onglet des alimentations connectées liste tous les équipements
branchés via les PDU. On trie par criticité (champ personnalisé), on repère les
mono-alimentés (tag), et on planifie arrêts propres et créneaux sans surprise.

**R9.** (1) Le service `netbox-rq` tourne-t-il ? Les webhooks partent via la file
`high` — s'il est arrêté, les événements s'accumulent dans Redis sans être envoyés.
(2) La file est-elle vide ? (`redis-cli llen rq:queue:high`, Admin > Files d'attente RQ.)
Ensuite seulement : URL, secret/signature, logs du récepteur (voir sections 57, 60).

**R10.** Trois au choix parmi : `configuration.py` en 600, `LOGIN_REQUIRED = True`,
`SECRET_KEY` unique, HTTPS + cookies `Secure` + HSTS, comptes nominatifs + permissions
(plus de compte partagé), LDAP/SSO, PostgreSQL/Redis en écoute locale, pare-feu
(443 uniquement), `DEBUG = False`, sauvegardes testées, bannière d'environnement
(voir section 46).

---

## 95. Pour aller plus loin

| Ressource | Adresse / référence |
|---|---|
| Documentation officielle NetBox | `https://docs.netbox.dev` (installation, API, release notes 4.x) |
| Dépôt GitHub | `https://github.com/netbox-community/netbox` (releases, issues) |
| Démo publique | `https://demo.netbox.dev` (bac à sable pour s'entraîner) |
| Collection Ansible | `netbox.netbox` sur Ansible Galaxy |
| Client Python | `pynetbox` (PyPI) — `https://pynetbox.readthedocs.io` |
| Communauté | Slack NetBox Community, forum GitHub Discussions |

**Feuille de route suggérée pour votre équipe (6 mois) :**

- **Mois 1** : installation + durcissement + convention de nommage validée ; 1 site pilote documenté à 100 %.
- **Mois 2** : migration de l'inventaire existant (section 88) ; inventaire Ansible dynamique en production.
- **Mois 3** : modélisation électrique complète (onduleurs/PDU) + premier rapport « bilan de puissance ».
- **Mois 4** : export Zabbix automatisé ; webhooks vers l'outil de ticketing.
- **Mois 5** : scripts et rapports d'exploitation (garanties, engagements opérateurs, préfixes saturés).
- **Mois 6** : revue qualité des données, audit de sécurité express (section 90), formation des nouveaux arrivants sur ce guide.

> 💡 **Dernier conseil de chef de service** : NetBox ne vaut que par la discipline
> collective qu'on met autour. Nommez un **référent NetBox** (tournant, 6 mois),
> faites-en un point fixe de votre réunion d'équipe mensuelle (revues section 85),
> et mesurez : « % d'équipements avec IP de management », « % de baies documentées ».
> Ce qui est mesuré progresse. Bon courage — votre infrastructure mérite mieux que
> des tableurs. 🚀

---

*Guide rédigé le 2026-09-26 — NetBox 4.x, Debian 12/13, Ubuntu 22.04/24.04 LTS.
Vérifiez les release notes officielles avant toute installation ou montée de version.*
