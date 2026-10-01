---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/installer-opnsense-pare-feu-open-source-en-14-etapes-2026-3
title: "Sous Linux, décompression de l'image téléchargée"
domain: opnsense-pfsense
role: reference
task: reference
actors: ["Google"]
dates: []
keywords: []
source: docs/RAG/collect-261001-opnsense-pfsense/installer-opnsense-pare-feu-open-source-en-14-etapes-2026.md
source_anchor: ""
source_lines: [132, 198]
sha256: 241e53bba84847bed28286a8f07a45d2d7c4bf29fa9da81f604faca058c043e7
---

# Sous Linux, décompression de l'image téléchargée

## Étape 10 : configurer un VPN avec WireGuard

Après avoir installé le plugin os-wireguard, une nouvelle section apparaît sous VPN → WireGuard. Commencez par créer une “Local” instance : un nom, un port d’écoute (le port UDP 51820 par défaut) et une plage d’adresses IP virtuelles pour le tunnel, par exemple `10.10.10.0/24`.

```
# Génération d'une paire de clés WireGuard côté client (Linux/macOS)
wg genkey | tee client_private.key | wg pubkey > client_public.key
# Exemple de configuration client WireGuard à importer sur l'appareil distant
[Interface]
PrivateKey = <client_private.key>
Address = 10.10.10.2/24
DNS = 192.168.1.1
[Peer]
PublicKey = <clé_publique_du_serveur_OPNsense>
Endpoint = votre-ip-publique-ou-ddns:51820
AllowedIPs = 192.168.1.0/24, 10.10.10.0/24
PersistentKeepalive = 25
```
Ajoutez chaque appareil distant comme “Peer” dans l’interface OPNsense, en collant sa clé publique. N’oubliez pas de créer une règle de pare-feu sur l’interface WAN autorisant le trafic UDP entrant sur le port 51820, sans quoi le tunnel ne s’établira jamais, et une règle sur l’interface WireGuard elle-même pour autoriser le trafic du tunnel vers le LAN.

## Étape 11 : activer la détection d’intrusion avec Suricata

Le plugin os-ids intègre le moteur Suricata directement dans OPNsense. Rendez-vous dans Services → Intrusion Detection → Administration. Deux modes sont disponibles : IDS (détection passive, journalisation uniquement) et IPS (prévention active, blocage en ligne du trafic malveillant). Le mode IPS exige davantage de ressources CPU car chaque paquet est inspecté avant d’être transmis.

Sous l’onglet “Download”, activez au moins un jeu de règles gratuit comme “ET Open” (Emerging Threats Open), puis lancez une mise à jour manuelle des règles. Sous “Policy”, associez les interfaces à surveiller (typiquement WAN, et LAN si vous redoutez une compromission interne). Activez ensuite le service et laissez-le tourner 24 à 48 heures en mode IDS avant de basculer en mode IPS, le temps d’identifier les faux positifs propres à votre trafic.

- Commencez toujours en mode détection (IDS) avant de passer en blocage (IPS)
- Désactivez les catégories de règles non pertinentes (par exemple les règles orientées jeux vidéo si votre réseau est un environnement professionnel strict) pour réduire la charge CPU
- Surveillez l’onglet “Alerts” quotidiennement la première semaine pour ajuster les seuils

## Étape 12 : mettre en place la haute disponibilité avec CARP

Pour les environnements qui ne tolèrent pas de coupure, OPNsense prend en charge le protocole CARP (Common Address Redundancy Protocol) permettant de faire fonctionner deux pare-feux en cluster actif/passif. Le nœud principal (MASTER) porte une adresse IP virtuelle partagée ; en cas de panne, le nœud secondaire (BACKUP) reprend automatiquement cette adresse.

La configuration se fait sous Interfaces → Virtual IPs pour créer les adresses CARP, puis sous System → High Availability → Settings pour synchroniser la configuration (XMLRPC Sync) et l’état des connexions (pfsync) entre les deux nœuds via un lien réseau dédié. Cette synchronisation garantit qu’une bascule ne coupe pas les connexions déjà établies, un point critique pour les connexions VPN de longue durée.

## Étape 13 : mises à jour, sauvegarde et restauration de la configuration

Toute la configuration d’OPNsense (règles de pare-feu, VPN, DHCP, plugins installés) est stockée dans un fichier unique, `config.xml`. Sous System → Configuration → Backups, téléchargez régulièrement ce fichier sur un support externe, hors du pare-feu lui-même. En cas de panne matérielle, une réinstallation suivie d’un import de ce fichier restaure l’intégralité de la configuration en quelques minutes.

```
# Vérification manuelle des mises à jour disponibles en ligne de commande
opnsense-update -c
# Application des mises à jour (équivalent à System > Firmware > Updates dans l'UI)
opnsense-update -bkr
```
Sous System → Firmware → Updates, OPNsense vérifie automatiquement la disponibilité de nouvelles versions et affiche un changelog avant application. Compte tenu du rythme de sortie soutenu du projet (trois versions mineures pour la seule branche 26.7 entre juillet et août 2026, un rythme déjà observé sur la branche 25.1 dont la dernière mineure, la 25.1.12, était sortie le 21 juillet 2025 selon endoflife.date), planifiez une vérification hebdomadaire au minimum, et systématiquement après toute annonce de faille de sécurité sur le forum officiel.

## Étape 14 : sécuriser l’accès administrateur avec la double authentification

Un pare-feu constitue une cible de choix : compromettre le compte root d’OPNsense revient à prendre le contrôle de tout le filtrage réseau d’une organisation. Avant de mettre l’équipement en production, il est fortement recommandé d’activer la double authentification (2FA) sur l’interface web, en plus d’un mot de passe robuste.

Sous System → Access → Servers, ajoutez un serveur d’authentification local avec un second facteur de type TOTP (Time-based One-Time Password), compatible avec des applications comme Google Authenticator, Aegis ou une YubiKey configurée en mode OTP. Une fois le serveur créé, associez-le au compte administrateur sous System → Access → Users, puis scannez le QR code généré avec votre application d’authentification.

- Créez un compte administrateur secondaire nommé (par exemple `admin-nom` ) plutôt que de n’utiliser que le compte`root` générique, pour garder une traçabilité des actions dans les journaux
- Restreignez l’accès à l’interface web à des adresses IP sources précises sous Firewall → Rules → LAN, en particulier si l’administration s’effectue aussi via VPN
- Changez le port HTTPS par défaut (443) de l’interface de gestion sous System → Settings → Administration si celle-ci reste exposée sur une interface autre que le LAN de confiance
- Désactivez l’accès administrateur en HTTP non chiffré, qui reste parfois actif par erreur après une mise à niveau

Ce durcissement prend une dizaine de minutes et élimine l’essentiel du risque lié au vol ou à la réutilisation d’un mot de passe, un vecteur d’attaque encore responsable d’une part importante des compromissions d’équipements réseau signalées par le CERT-FR ces dernières années.

## Multi-WAN et répartition de charge entre deux connexions Internet

Pour les petites entreprises qui ne peuvent pas se permettre une coupure Internet totale, OPNsense permet de connecter deux liens WAN distincts (par exemple une fibre et un lien 4G/5G de secours) et de répartir automatiquement la charge ou de basculer sur le second lien en cas de panne du premier.

La configuration se fait sous System → Gateways, où chaque connexion WAN reçoit une passerelle distincte surveillée par un test de latence (dpinger) vers une cible fiable, par exemple un résolveur DNS public. Sous System → Gateways → Group, créez ensuite un groupe de passerelles définissant l’ordre de priorité (tier 1, tier 2) : en fonctionnement normal, le trafic emprunte la passerelle de priorité la plus haute, et bascule automatiquement sur la passerelle de secours si les tests de latence ou de perte de paquets dépassent les seuils configurés.

