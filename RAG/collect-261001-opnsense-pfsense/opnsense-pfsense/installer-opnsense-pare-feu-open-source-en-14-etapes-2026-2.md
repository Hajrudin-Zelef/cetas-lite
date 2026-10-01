---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/installer-opnsense-pare-feu-open-source-en-14-etapes-2026-2
title: "Sous Linux, décompression de l'image téléchargée"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-opnsense-pfsense/installer-opnsense-pare-feu-open-source-en-14-etapes-2026.md
source_anchor: ""
source_lines: [50, 131]
sha256: 068ede23569c9a3a5874d0658aa54257987d365d5ca473623093ce0af33bd10c
---

# Sous Linux, décompression de l'image téléchargée

Sous macOS ou Linux, balenaEtcher offre une interface identique en trois clics : sélection de l’image, sélection du disque cible, écriture. Etcher a l’avantage de vérifier automatiquement l’écriture après la gravure, ce qui élimine une source d’échec fréquente à l’étape suivante.

## Étape 3 : démarrer l’installateur et effectuer l’installation

Branchez la clé USB sur la machine cible, allumez-la et entrez dans le menu de démarrage (Boot Menu), généralement accessible avec F11, F12 ou Échap selon le fabricant de la carte mère. Sélectionnez la clé USB comme périphérique de démarrage.

OPNsense démarre dans un environnement live et affiche un menu console. Connectez-vous avec l’identifiant `installer` ou lancez directement le programme d’installation depuis le menu. L’assistant vous demande ensuite :

- Le clavier à utiliser (choisissez “French” pour un clavier AZERTY)
- Le type d’installation : “Install (ZFS)” est recommandé pour bénéficier des snapshots et de la protection contre la corruption de données, “Install (UFS)” reste plus léger pour du matériel très modeste
- Le disque de destination
- Le mot de passe du compte `root` , à choisir robuste puisqu’il protège l’accès administrateur complet du pare-feu

Une fois l’installation terminée, l’assistant propose de redémarrer. Retirez la clé USB avant le redémarrage pour éviter de rebooter sur l’installateur en boucle, une erreur d’inattention très courante.

## Étape 4 : assigner les interfaces réseau WAN et LAN

Au premier démarrage, la console affiche le menu principal d’OPNsense avec une série d’options numérotées. Choisissez l’option “Assign interfaces” (assigner les interfaces). Le système liste les interfaces détectées avec leur nom de pilote FreeBSD, par exemple `igb0`, `igb1` ou `em0`, `em1`.

C’est l’étape la plus sensible de toute l’installation. Le système vous demande de désigner physiquement l’interface WAN (connectée à votre box ou modem) puis l’interface LAN (connectée à votre commutateur interne). Le moyen le plus fiable de ne pas se tromper consiste à débrancher tous les câbles réseau sauf un, observer quelle interface devient active dans la console, puis répéter l’opération pour chaque câble.

```
# Dans le menu console OPNsense
# Option 1 : Assign interfaces
# Le système demande :
Enter the WAN interface name or 'a' for auto-detection: igb0
Enter the LAN interface name or 'a' for auto-detection: igb1
Enter the Optional interface 1 name or 'a' for auto-detection (or nothing if finished): [Entrée]
Do you want to proceed? [y/n]: y
```
## Étape 5 : configurer l’adresse IP du LAN et accéder à l’interface web

Toujours dans le menu console, sélectionnez l’option de configuration de l’interface LAN. Attribuez une adresse IPv4 statique, par exemple `192.168.1.1/24`, et activez le serveur DHCP si vous voulez que les postes du réseau interne reçoivent automatiquement une adresse IP dès cette étape.

Connectez ensuite un ordinateur au port LAN du pare-feu (directement ou via un commutateur), laissez-le obtenir une adresse par DHCP, puis ouvrez un navigateur et accédez à `https://192.168.1.1`. Le certificat étant auto-signé, votre navigateur affiche un avertissement de sécurité : c’est normal à ce stade, acceptez l’exception. Connectez-vous avec l’identifiant `root` et le mot de passe défini lors de l’installation.

## Étape 6 : suivre l’assistant de configuration initiale

Dès la première connexion, OPNsense lance un assistant de configuration en plusieurs écrans. Il vous demande le nom d’hôte, le domaine, les serveurs DNS à utiliser (ou de laisser le résolveur Unbound intégré s’en charger), le fuseau horaire et un serveur NTP. Vient ensuite le type de connexion WAN : IP statique, DHCP (le plus courant chez les box grand public en mode bridge) ou PPPoE (fréquent chez certains FAI xDSL en France).

L’assistant propose enfin de changer le mot de passe root et de recharger la configuration. À l’issue de ces écrans, le pare-feu applique déjà une politique par défaut sensée : le trafic sortant du LAN vers le WAN est autorisé, le trafic entrant depuis le WAN vers le LAN est bloqué.

## Étape 7 : installer les extensions (plugins) essentielles

OPNsense repose sur une architecture modulaire : l’installation de base reste volontairement légère et les fonctionnalités avancées s’ajoutent via des plugins téléchargés depuis le dépôt officiel. Rendez-vous dans le menu Système → Firmware → Plugins.

| Plugin | Fonction | Recommandé pour | 
|---|---|---|
| os-wireguard | VPN WireGuard, site-à-site ou nomade | Tous les usages, performance élevée | 
| os-ids (Suricata) | Détection et prévention d’intrusion réseau | Entreprises, réseaux exposés | 
| os-acme-client | Certificats TLS automatiques via Let’s Encrypt | Interfaces web exposées publiquement | 
| os-theme-cicada / os-theme-rebellion | Thèmes d’interface alternatifs | Confort visuel, aucun impact fonctionnel | 
| os-netflow / Insight | Statistiques de trafic détaillées par hôte et protocole | Supervision de bande passante | 
| os-nut | Gestion d’onduleur (UPS) réseau | Salles serveurs, arrêt propre sur coupure secteur | 

Pour installer un plugin, tapez son nom dans le champ de recherche, cliquez sur l’icône “+” en face du résultat, puis patientez quelques secondes pendant le téléchargement et l’installation. Certains plugins comme os-wireguard ajoutent immédiatement un nouveau menu dans l’interface, sans nécessiter de redémarrage du pare-feu.

## Étape 8 : configurer les règles de pare-feu et le NAT

La version 26.7 généralise le nouveau moteur de règles MVC (Modèle-Vue-Contrôleur), qui restructure l’affichage des règles sans changer leur logique. Rendez-vous dans Firewall → Rules → LAN pour consulter la règle par défaut autorisant tout le trafic sortant du réseau local.

Pour ouvrir un service hébergé en interne vers l’extérieur (un serveur web, un accès VPN), il faut créer une redirection de port sous Firewall → NAT → Port Forward. OPNsense génère alors automatiquement la règle de pare-feu correspondante sur l’interface WAN, sauf si vous décochez l’option “Filter rule association”.

```
# Exemple : redirection du port 443 (HTTPS) vers un serveur web interne
# Firewall > NAT > Port Forward > Add
Interface: WAN
Protocol: TCP
Destination: WAN address
Destination port range: HTTPS (443)
Redirect target IP: 192.168.1.50
Redirect target port: HTTPS (443)
Description: Reverse proxy interne - serveur web
```
Pour le NAT sortant, le mode par défaut “Automatic outbound NAT” convient à la majorité des installations. Si vous avez besoin d’un contrôle plus fin (par exemple pour forcer certaines machines à sortir par une adresse IP publique dédiée sur une connexion multi-WAN), basculez sur le mode “Hybrid” ou “Manual” dans Firewall → NAT → Outbound.

## Étape 9 : mettre en place le DHCP et le résolveur DNS Unbound

Sous Services → DHCPv4 → LAN, activez le serveur DHCP et définissez une plage d’adresses, par exemple de `192.168.1.100` à `192.168.1.200`, en laissant les 99 premières adresses disponibles pour des attributions statiques (imprimantes, serveurs, équipements réseau). Renseignez la passerelle et le serveur DNS avec l’adresse du pare-feu lui-même (`192.168.1.1`), qui fait office de résolveur.

OPNsense embarque Unbound comme résolveur DNS récursif par défaut, accessible sous Services → Unbound DNS → General. Le mode récursif interroge directement les serveurs racine DNS sans dépendre d’un fournisseur tiers, ce qui améliore la confidentialité. Cochez “Register DHCP leases in the DNS Resolver” pour que les machines du réseau local soient résolubles par leur nom d’hôte.

