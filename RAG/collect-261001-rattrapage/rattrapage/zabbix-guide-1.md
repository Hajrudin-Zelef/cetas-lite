---
id: collect-261001-rattrapage/rattrapage/zabbix-guide-1
title: "Guide Zabbix complet — Supervision d'infrastructure en production"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "open source"]
source: docs/RAG/collect-261001-rattrapage/zabbix_guide.md
source_anchor: ""
source_lines: [1, 134]
sha256: ff13079855dee697791ade2c04a9e7152c55e3c5930c78b5819a14be710c1d7a
---

# Guide Zabbix complet — Supervision d'infrastructure en production

> **Public** : Zelef, chef de service systèmes & énergies, sysadmin.
> **Angle** : pratique, production, contexte Afrique de l'Ouest (coupures électriques fréquentes, liens réseau instables, onduleurs au cœur du métier).
> **Versions couvertes** : Zabbix **6.0 LTS**, **6.4**, **7.0 LTS**, **7.2/7.4** (les différences sont signalées par des ⚠️).
> **OS cibles** : Debian 12/13, Ubuntu 22.04 LTS / 24.04 LTS.
> **Avertissement** : les valeurs de dimensionnement sont des ordres de grandeur à valider sur votre parc réel. Testez toujours en préproduction avant d'appliquer en production.

---

## Table des matières

1. Pourquoi Zabbix (et pas une autre solution)
2. Comment utiliser ce guide
3. Conventions utilisées dans ce guide
4. Architecture générale de Zabbix
5. Les composants en détail : Server
6. Les composants en détail : Agent 2 (et Agent 1)
7. Les composants en détail : Proxy
8. Les composants en détail : Frontend web
9. Les composants en détail : base de données (MySQL vs PostgreSQL)
10. Ports réseau et flux — tableau de référence
11. Choisir sa version : LTS vs standard
12. Dimensionnement : CPU, RAM, disque, IOPS
13. Prérequis système et réseau
14. Installation pas à pas : Zabbix Server + PostgreSQL sur Debian/Ubuntu
15. Installation pas à pas : variante MySQL/MariaDB
16. Installation pas à pas : le frontend web (Nginx + PHP-FPM)
17. Installation pas à pas : Zabbix Agent 2 sur Linux
18. Installation pas à pas : Zabbix Agent sur Windows
19. Installation pas à pas : Zabbix Proxy
20. Sécuriser la communication : PSK et certificats TLS
21. Premier démarrage : assistant de configuration du frontend
22. Configuration initiale : réglages globaux essentiels
23. Utilisateurs, groupes, rôles et permissions
24. Audit et traçabilité
25. Concepts clés : hôtes, groupes d'hôtes, interfaces
26. Ajouter son premier hôte pas à pas
27. Templates : principes et bonnes pratiques
28. Templates officiels essentiels (Linux, Windows, SNMP, réseau)
29. Créer et lier un template personnalisé
30. Items : actifs vs passifs, trapper, dépendants
31. Les clés d'items de l'agent : les indispensables
32. Items calculés et agrégés
33. Prétraitement des valeurs (preprocessing)
34. Intervalles, périodes de collecte et flex intervals
35. Triggers : syntaxe et logique
36. Fonctions de trigger : le catalogue indispensable (last, avg, nodata…)
37. Hystérésis et seuils intelligents
38. Sévérités et bonnes pratiques de nommage
39. Dépendances entre triggers (éviter les tempêtes d'alertes)
40. Actions : le moteur de notification
41. Médias d'alerte : email
42. Médias d'alerte : SMS via passerelle (contexte Afrique de l'Ouest)
43. Médias d'alerte : Telegram
44. Webhooks : Slack, Teams, et autres
45. Escalades et acquittements
46. Découverte réseau automatique
47. Low-Level Discovery (LLD) : disques, interfaces, processus
48. Écrire ses propres règles LLD (UserParameter + JSON)
49. Supervision SNMP : principes (v2c vs v3)
50. Superviser un switch via SNMP pas à pas
51. Superviser un onduleur via SNMP (cas métier — le cœur de ce guide)
52. Template onduleur : items, triggers et dashboard dédiés
53. Supervision JMX (Java/Tomcat)
54. Supervision IPMI (températures, ventilateurs, alimentations)
55. Supervision VMware (intégration native)
56. Supervision web : scénarios et checks HTTP
57. Maintenance planifiée (sans alerte parasite)
58. Haute disponibilité native (Zabbix 6.0+)
59. Housekeeping : comprendre la rétention des données
60. Partitionnement de la base (PostgreSQL/MySQL) pour les gros parcs
61. Performance et tuning du server/proxy
62. Usage quotidien : la routine de l'exploitant
63. Dashboards et vues : construire des écrans d'exploitation
64. Inventaire automatique des hôtes
65. Cartes réseau (maps) et écrans de supervision
66. Bonnes pratiques d'alerting : la chasse au bruit
67. Intégration avec Grafana
68. Sécurité et durcissement en production
69. Superviser Zabbix lui-même (qui surveille le surveillant ?)
70. Sauvegarde : base de données, fichiers, stratégie 3-2-1
71. Restauration : procédure testée
72. Mise à jour de version : méthode sans stress
73. Dépannage : méthodologie générale
74. Erreur n°1 à n°12 : les classiques commentés
75. Cas pratique n°1 : supervision d'une baie complète
76. Cas pratique n°2 : supervision d'un onduleur 40 kVA de A à Z
77. Cas pratique n°3 : alerte coupure électrique avec escalade SMS
78. Cas pratique n°4 : supervision d'un lien VSAT/fibre instable
79. Cas pratique n°5 : capacity planning disque avec prédiction
80. Cas pratique n°6 : supervision d'une salle serveur (température, hygrométrie)
81. Cas pratique n°7 : déploiement multi-sites avec proxies
82. Cas pratique n°8 : supervision d'un cluster Proxmox VE
83. Cas pratique n°9 : blackout test — valider toute la chaîne d'alerte
84. Cas pratique n°10 : rapport mensuel de disponibilité pour la direction
85. Pense-bête de poche (cheat sheet imprimable)
86. Glossaire
87. Quiz : 10 questions pour valider (avec réponses)
88. Pour aller plus loin

---

## 1. Pourquoi Zabbix (et pas une autre solution)

Zabbix est une plateforme de supervision **open source**, mature (depuis 2001), sans coût de licence, capable de superviser **plusieurs dizaines de milliers d'équipements** avec un seul serveur correctement dimensionné. Pour un service systèmes & énergies en Afrique de l'Ouest, ses atouts sont concrets :

- **Zéro coût de licence** : budget préservé pour le matériel (sondes, serveurs, onduleurs).
- **SNMP natif et puissant** : switchs, routeurs, onduleurs, sondes environnementales parlent SNMP ; Zabbix les interroge sans agent.
- **Proxy** : supervise des sites distants même avec un lien intermittent (le proxy stocke localement et envoie quand le lien revient — vital avec des liaisons VSAT ou 4G instables).
- **LLD (Low-Level Discovery)** : découvre automatiquement disques, interfaces réseau, partitions — fini la configuration manuelle machine par machine.
- **Alerting flexible** : email, SMS via passerelle GSM locale, Telegram — adapté aux réalités du terrain où l'email n'est pas toujours lu à temps.
- **Pas de dépendance cloud obligatoire** : tout tourne on-premise, vos données restent chez vous.

Comparaison rapide (ordre de grandeur, à valider selon votre contexte) :

| Solution | Licence | Agent | SNMP | Proxy/offline | Courbe d'apprentissage |
|---|---|---|---|---|---|
| Zabbix | Gratuite (GPL) | Oui (Agent 2) | Natif | Oui (proxy) | Moyenne |
| Nagios / Centreon | Gratuite / payante | NRPE/NSClient | Via plugins | Partielle | Moyenne à forte |
| Prometheus | Gratuite | Exporters | Via snmp_exporter | Non natif | Forte (PromQL) |
| PRTG | Payante au capteur | Oui | Oui | Sondes distantes | Faible |

> 💡 **Conseil de terrain** : si votre équipe est petite (1 à 3 personnes), Zabbix offre le meilleur ratio puissance/simplicité d'exploitation au quotidien.

## 2. Comment utiliser ce guide

- **Débutant Zabbix** : lisez dans l'ordre, de la section 4 à la section 47, en montant une maquette en parallèle (une VM suffit).
- **Déjà utilisateur** : piochez dans les sections 48+ (SNMP, onduleur, HA, tuning, dépannage).
- **En astreinte** : imprimez la section 85 (pense-bête) et gardez la section 74 (erreurs classiques) sous la main.
- Tous les exemples de configuration sont **testés syntaxiquement** pour Zabbix 6.x/7.x. Les mots de passe sont fictifs (`mot_de_passe_ici`) : remplacez-les.

## 3. Conventions utilisées dans ce guide

