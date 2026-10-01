---
id: collect-261001-fortinet/fortinet/yasminekechid2-fortigate-en-production-partie-2-haute-disponibilit-c3-a9-ha-a426-cc1022fe-1
title: "Vérifier la version firmware"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/yasminekechid2-fortigate-en-production-partie-2-haute-disponibilit-c3-a9-ha-a426-cc1022fe.md
source_anchor: ""
source_lines: [1, 117]
sha256: 6727d2fd5aac86c43eb81a24445086f66e305966ea3a762d8223aba9831f866a
---

# Vérifier la version firmware

FortiGate en Production — Partie 2 : Haute Disponibilité (HA)
1. Introduction
La continuité de service est un impératif pour toute infrastructure réseau d’entreprise. Une interruption du pare-feu, même brève, peut paralyser l’activité, bloquer l’accès aux applications métier et impacter la productivité. La Haute Disponibilité (High Availability — HA) répond à cette problématique en éliminant le point de défaillance unique.
Ce deuxième article de la série FortiGate présente la mise en œuvre d’un cluster HA actif-passif, permettant un basculement automatique et transparent en cas de défaillance du pare-feu principal. Cette configuration garantit une disponibilité optimale de l’infrastructure réseau.
2. Concept de la Haute Disponibilité FortiGate
Principe de Fonctionnement : La Haute Disponibilité FortiGate repose sur un cluster composé de deux pare-feux fonctionnant en mode actif-passif :
Firewall Primary (actif) : Traite l’intégralité du trafic réseau. C’est l’équipement opérationnel qui applique les politiques de sécurité et route les paquets.
Firewall Secondary (passif) : Surveille en permanence l’état du Primary via des mécanismes de heartbeat. Il maintient une configuration synchronisée et reste prêt à prendre le relais instantanément.
En cas de défaillance du Primary (panne matérielle, redémarrage, perte de connectivité), le Secondary détecte l’anomalie et assume automatiquement le rôle de Primary. Ce basculement (failover) s’effectue en quelques secondes, de manière totalement transparente pour les utilisateurs.
Avantages de l’Architecture HA
- Continuité de service : Élimination du SPOF (Single Point of Failure)
- Basculement transparent : Les sessions actives sont maintenues
- Maintenance sans interruption : Possibilité de mettre à jour ou redémarrer un firewall sans coupure
- Synchronisation automatique : La configuration est répliquée du Primary vers le Secondary
3. Architecture du Cluster HA
L’infrastructure HA déployée dans ce laboratoire comprend deux FortiGate interconnectés selon une topologie redondante :
Composants de l’Architecture
Deux pare-feux FortiGate : Configurés en cluster HA actif-passif
Interfaces de production :
- Port1 : Interface LAN/Management (192.168.0.254)
- Port3 : Interface WAN/NAT (adresse publique)
Interfaces de heartbeat :
- Port4 : Heartbeat primaire (192.168.11.254 sur F1, 192.168.11.253 sur F2)
- Port5 : Heartbeat secondaire (192.168.12.254 sur F1, 192.168.12.253 sur F2)
Interface de management dédiée :
- Port8 : Management HA (10.0.11.251 pour Primary, 10.0.11.252 pour Secondary)
Rôle des Interfaces Heartbeat
Les interfaces heartbeat sont le cœur du mécanisme HA. Elles servent à :
Vérification d’état : Le Primary envoie des messages réguliers (heartbeats) au Secondary pour signaler qu’il est opérationnel.
Détection de panne : Si le Secondary ne reçoit plus de heartbeat pendant un délai défini, il considère le Primary comme défaillant et déclenche le failover.
Synchronisation : Les changements de configuration sont transmis via ces interfaces pour maintenir les deux firewalls identiques.
Redondance du heartbeat : L’utilisation de deux interfaces heartbeat distinctes (port4 et port5) élimine le risque de faux positif. Le Secondary ne bascule que si les deux liens heartbeat sont rompus simultanément, évitant ainsi un split-brain.
Prérequis de Configuration
Avant d’activer le cluster HA, les deux FortiGate doivent être configurés de manière strictement identique :
Configuration Identique Obligatoire
- Interfaces réseau : Mêmes ports physiques activés
- Adresses IP : Adresses identiques sur les interfaces de production
- Politiques de sécurité : Règles de firewall identiques
- Routes : Table de routage identique
Cette homogénéité est critique car, une fois le cluster formé, la configuration du Primary sera automatiquement synchronisée vers le Secondary. Toute divergence initiale peut entraîner des dysfonctionnements.
Vérifications Préalables
Avant de procéder à la configuration HA :
# Vérifier la version firmware
get system status
# Tester la connectivité entre les firewalls
execute ping 192.168.11.253 # Depuis F1 vers F2 sur port4
execute ping 192.168.12.253 # Depuis F1 vers F2 sur port5# Vérifier les interfaces actives
get system interface physical
4. Configuration du Cluster HA
Configuration du Firewall Primary (F1)
- Navigation : System → HA
Paramètres du cluster :
Mode : Active-Passive
Device Priority : 200
Group Name : HA
Password : [mot de passe sécurisé]
Configuration des heartbeat interfaces :
Interface Adresse IP Rôle port4 192.168.11.254/24 Heartbeat primaire port5 192.168.12.254/24 Heartbeat secondaire
Rôle de la Priorité
La priorité détermine quel firewall devient Primary lors de la formation du cluster :
- Priorité 200 (F1) : Devient Primary au démarrage
- Priorité 100 (F2) : Devient Secondary au démarrage
En cas de failover, le firewall avec la plus grande HA uptime (temps depuis le dernier démarrage) conserve le rôle Primary, même si l’autre firewall a une priorité supérieure. Ce comportement évite les basculements intempestifs.
Configuration du Firewall Secondary (F2)
Sur le Firewall 2, appliquez une configuration identique avec une priorité inférieure :
Mode : Active-Passive
Device Priority : 100
Group Name : HA
Password : [même mot de passe que F1]
Interfaces heartbeat :
Interface Adresse IP Rôle port4 192.168.11.253/24 Heartbeat primaire port5 192.168.12.253/24 Heartbeat secondaire
Formation du Cluster
Après application de la configuration sur F2 :
- La connexion GUI peut être temporairement interrompue (redémarrage des services)
- Le cluster HA se forme automatiquement en 30–60 secondes
- F1 assume le rôle Primary (priorité 200)
- F2 assume le rôle Secondary (priorité 100)
- La synchronisation de configuration démarre
Vérification via CLI :
# Sur F1
get system ha status
# Affiche : HA Health Status: OK, Master, vcluster 1# Sur F2
get system ha status
# Affiche : HA Health Status: OK, Slave, vcluster 1
Interface de Management Dédiée
Une amélioration majeure des versions récentes de FortiOS est la possibilité de configurer une interface de management dédiée pour chaque membre du cluster HA.
Problématique Historique
Avant cette fonctionnalité :
- Seul le firewall Primary était accessible via son IP de management
- Le Secondary n’était joignable qu’en console série
- Toute intervention sur le Secondary nécessitait une connexion physique
Solution : HA Management Interface
L’interface de management dédiée (typiquement port8) permet :
- Accès permanent : Chaque firewall possède sa propre IP de management, accessible quelle que soit son rôle (Primary/Secondary)
- Gestion simplifiée : Administration simultanée des deux firewalls via GUI ou SSH
- Isolation du trafic : L’interface de management est exclue de la table de routage, évitant toute interférence avec le trafic utilisateur
- Maintenance facilitée : Diagnostic et dépannage du Secondary sans impacter le trafic de production
Configuration du Management sur F1
Navigation : System → HA → Enable HA Reserved Management Interface
Get Yasmine kechid’s stories in your inbox
Join Medium for free to get updates from this writer.
Configuration de port8 :
Interface : port8
IP Address : 10.0.11.251/24
Allowaccess : ping, https, ssh, snmp
Configuration du Management sur F2
Avant de configurer F2, vérifiez que la synchronisation HA est active et que F1 est bien Primary.
Configuration de port8 sur F2 :
Interface : port8
IP Address : 10.0.11.252/24
Allowaccess : ping, https, ssh, snmp
Test de connectivité :
Depuis un poste du réseau de management :
- Accès GUI F1 : https://10.0.11.251
- Accès GUI F2 : https://10.0.11.252
Importance de cette Fonctionnalité
En environnement de production, cette interface dédiée est essentielle pour :
