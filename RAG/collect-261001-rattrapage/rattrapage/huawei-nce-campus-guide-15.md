---
id: collect-261001-rattrapage/rattrapage/huawei-nce-campus-guide-15
title: "Guide technique ultra-complet — iMaster NCE-Campus (Huawei)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_nce_campus_guide.md
source_anchor: ""
source_lines: [1285, 1381]
sha256: b9ff2d40171b61ad535860b835c94e8696a2bf9512f0666b0b626a3c82fe178b
---

# Guide technique ultra-complet — iMaster NCE-Campus (Huawei)

- **Zabbix** : supervision transverse (dont onduleurs !) — NCE pousse ses alarmes vers Zabbix (ou Zabbix interroge l'API NCE), corrélation réseau/énergie dans un seul écran (ex : « le switch est down **et** l'onduleur est sur batterie » — diagnostic immédiat).
- **GLPI** : création automatique de tickets depuis les alarmes NCE (via API), inventaire enrichi.
- **Scripts** : exports périodiques (inventaire, conformité), checks custom, rapports pour la direction.
- **Supervision du superviseur** : Zabbix surveille NCE lui-même (services, espace disque, certificats) — « qui supervise le superviseur » (section 22).

Architecture conseillée : NCE = source de vérité **réseau campus Huawei**, Zabbix = supervision **transverse**, GLPI = **tickets/inventaire**. Chacun son rôle, intégrés par API.

---

# PARTIE 10 — ASSURANCE ET SUPERVISION

## 92. Assurance et supervision — vue d'ensemble

L'« assurance » (assurance qualité de service) va au-delà de la supervision classique : il ne s'agit pas seulement de savoir si un équipement est up/down, mais si les **utilisateurs ont une bonne expérience**.

Les briques NCE :

- **Monitoring temps réel** : tableaux de bord, topologie, indicateurs (93-94).
- **Analyse** : détection d'anomalies, cause racine, expérience par utilisateur/application (95-96, CampusInsight section 20).
- **Pilotage** : rapports, alertes, SLA (97-99).
- **Maintenance** : actions correctives depuis NCE (100).

Pour une équipe réduite, l'assurance bien réglée = **moins de tickets**, des tickets **mieux qualifiés**, et des **preuves** (rapports) pour la direction.

## 93. Monitoring temps réel — tableaux de bord

Les tableaux de bord NCE affichent typiquement :

- **Santé globale** : équipements par état et par site (normal/alarme/hors ligne).
- **Utilisateurs** : connectés par SSID/site, échecs d'authentification.
- **Trafic** : top parleurs, applications (si identification activée), utilisation des liens.
- **Wi-Fi** : AP par état, clients par AP, interférences, roaming.
- **Alertes** : alarmes en cours par sévérité.

Bonnes pratiques : un **dashboard « tour de contrôle »** affiché en permanence (NOC/accueil technique), des dashboards **par site** pour les responsables locaux, des dashboards **métier** (ex : disponibilité du Wi-Fi invités pour l'hôtellerie). Revoir les dashboards **trimestriellement** (un dashboard que personne ne regarde est un dashboard à supprimer).

## 94. Cartographie réseau (digital map) — topologie

La **carte numérique** (digital map) : topologie **générée automatiquement** (découverte LLDP + remontées contrôleur), affichant équipements, liens, états en temps réel.

Usages : diagnostic visuel (« où est la coupure ? »), vérification d'architecture (le câblage réel correspond-il au plan ?), présentation à la direction, support aux techniciens distants (« cliquez sur l'AP rouge »).

Limites honnêtes : la topologie auto-découverte peut être **incomplète** (équipements non LLDP, liens non remontés) — la confronter au plan de câblage réel, et la **documenter** (un export mensuel archivé = preuve en cas de litige). Les liens LLDP vers des équipements tiers sont informatifs, pas contractuels.

## 95. Détection d'anomalies — ce que fait l'IA/ML

CampusInsight (et dans une moindre mesure NCE de base) applique du machine learning à la télémétrie :

- **Lignes de base** : le système apprend le comportement « normal » (trafic, latence, taux d'erreur par heure/jour) et signale les **écarts**.
- **Prédiction** : dégradation progressive détectée avant la panne (ex : erreurs CRC croissantes sur un port → SFP/câble à changer).
- **Optimisation radio** : ajustement des canaux/puissances (section 75).

Honnêteté : le ML a besoin de **données** (plusieurs semaines d'historique) et produit des **faux positifs** — ne jamais automatiser d'action corrective sur une simple anomalie ML sans validation humaine (au début du moins). Évaluer en pilote : compter les vrais positifs pendant 3 mois avant de s'y fier.

## 96. Analyse de cause racine (root cause analysis)

Quand un incident survient, NCE aide à remonter à la **cause** plutôt que de traiter les symptômes :

- **Corrélation** : une alarme « AP hors ligne » + une alarme « port PoE down sur le S310 » + une alarme « onduleur sur batterie » (via supervision transverse) = probablement une coupure électrique locale, pas un problème Wi-Fi.
- **Base de connaissances** : les pannes connues et leurs solutions (Huawei fournit une base ; l'enrichir avec **vos** cas — chaque incident résolu devient une fiche).
- **Chronologie** : timeline des événements avant l'incident (changement de config ? mise à jour ?).

Méthode : face à une alarme, toujours chercher **ce qui a changé** (config, firmware, environnement) avant de chercher ce qui est cassé. 80 % des incidents réseau suivent un changement.

## 97. Rapports — types et planification

Types de rapports utiles :

- **Disponibilité** : % uptime par site/équipement (pour les SLA internes).
- **Capacité** : utilisation des liens/ports, croissance (pour anticiper les investissements).
- **Sécurité** : échecs d'authentification, équipements non conformes, ports ouverts.
- **Wi-Fi** : expérience par SSID/site, roaming, interférences.
- **Conformité** : écarts de configuration détectés.
- **Direction** : synthèse mensuelle sur une page (état, incidents majeurs, projets).

Planification : rapports **automatiques** (hebdo/mensuel) envoyés par mail aux destinataires — un rapport manuel est un rapport qui ne sortira pas. Archiver les rapports (traçabilité, comparaisons annuelles).

## 98. Alertes — configuration mail/SMS

- **Canaux** : e-mail (via relais SMTP configuré section 40), SMS (via passerelle SMS — prévoir le contrat opérateur), éventuellement webhook vers l'ITSM.
- **Sévérités** : critique (panne de site, contrôleur), haute (équipement down, boucle), moyenne (dégradation, seuil), basse (informatif). Chaque sévérité = un **canal** et un **délai** (le SMS à 3 h du matin, c'est pour le critique uniquement).
- **Anti-bruit** : seuils, temporisations (pas d'alerte sur un flap de 30 s), fenêtres de maintenance (pas d'alerte pendant les travaux planifiés), **escalade** (si pas d'acquittement en X min → niveau supérieur).
- **Test** : provoquer une alarme test **mensuellement** (débrancher un port de maquette) — une alerte qui n'a jamais été testée est une alerte qui ne marche pas.

## 99. SLA management — NQA et mesures actives

La documentation NCE prévoit des **tâches SLA** pour les équipements gérés en SNMP, basées sur le protocole **NQA** (Network Quality Analysis) des équipements :

- **Principe** : l'équipement envoie des sondes actives (ping, jitter, HTTP...) vers des cibles et mesure : latence, gigue, perte, disponibilité.
- **Usages** : vérifier un lien WAN de branche (l'AR720 sonde le siège), valider un SLA opérateur (preuves chiffrées), détecter une dégradation **avant** les plaintes.
- **Configuration** : depuis NCE, planifier les sondes (cibles, fréquence, seuils), collecter les résultats, alerter sur dépassement.
- **Limites** : les sondes consomment (peu) de bande passante ; les résultats dépendent de la charge de l'équipement sondeur — interpréter avec recul.

## 100. Maintenance des équipements depuis NCE

NCE centralise les opérations de maintenance (documentation Monitoring and O&M) :

