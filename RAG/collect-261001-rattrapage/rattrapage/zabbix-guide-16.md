---
id: collect-261001-rattrapage/rattrapage/zabbix-guide-16
title: "Guide Zabbix complet — Supervision d'infrastructure en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/zabbix_guide.md
source_anchor: ""
source_lines: [2613, 2625]
sha256: b56bee914e5c8670b49a91ef0fbb030c3392b46f65e823c46332111fac098070
---

# Guide Zabbix complet — Supervision d'infrastructure en production

**Feuille de route conseillée (6 mois)** :
1. Mois 1 : maquette + 10 hôtes pilotes, chaîne d'alerte email/Telegram validée.
2. Mois 2 : templates onduleur/switch, dépendances, premier site avec proxy.
3. Mois 3 : SMS opérationnel, dashboards énergie/réseau/direction, blackout test n°1.
4. Mois 4 : durcissement (TLS, SNMPv3, 2FA), sauvegarde testée (restauration réelle).
5. Mois 5 : Grafana, rapports mensuels automatisés, LLD maison pour le parc spécifique.
6. Mois 6 : revue du bruit, HA si besoin, blackout test n°2, documentation d'exploitation à jour.

> 💡 **Dernier conseil** : la supervision est un organisme vivant. Un Zabbix qu'on n'entretient pas (seuils jamais relus, templates jamais mis à jour, sauvegardes jamais testées) devient en 2 ans un générateur de bruit que plus personne ne regarde. La routine de la section 62 est aussi importante que l'installation.

---

*Fin du guide — bon monitoring !* 🔍
