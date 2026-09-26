---
id: collect-260926-mikrotik/mikrotik/tuto-remplacement-bbox-fibre-par-mikrotik-ipv4-ipv6-et-tv-replay-2
title: "Remplacer le numéro de série par celui de la BBox"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/tuto-remplacement-bbox-fibre-par-mikrotik-ipv4-ipv6-et-tv-replay.md
source_anchor: ""
source_lines: [168, 217]
sha256: 195a53e8ee7ef5b522c789238ec01ac31be62e1df98f5215375e744221cfeffc
---

# Remplacer le numéro de série par celui de la BBox

```
/routing igmp-proxy
set query-interval=1m quick-leave=yes
/routing igmp-proxy interface
add alternative-subnets=89.86.96.0/24,89.86.97.0/24,176.165.8.0/24,193.251.97.0/24 interface=BouyguesNet upstream=yes
add interface=LAN
#Ajouter l'option DNS pour les serveurs
/ip/dhcp-server/option
add code=6 name=dns-bboxtv value="'194.158.122.10''194.158.122.15'"
#Créer un bail dhcp spécifique pour la BBox TV
/ip dhcp-server lease
add address=192.168.100.10 comment="BBox TV" dhcp-option=dns-bboxtv mac-address=D0:5A:xx:xx:xx:xx server=LAN
/ip firewall mangle
add action=set-priority chain=output new-priority=5 out-interface=BouyguesNet passthrough=no protocol=igmp
```
Forcer IGMPv2 (workaround)
- Relier le port WAN du routeur à un switch

- Relier la première interface du second igmpproxy au LAN (pour l'installation  bon fonctionnement d'IGMP Proxy)

- Relier la seconde interface du second igmpproxy au même switch que SFP GPON

- Le switch doit être correctement configuré pour autoriser le VLAN100 sur les 2 ports, ne pas activer l'IGMP Snooping sur le VLAN100

**Installation du conteneur LXC sous proxmox**
Voici une très rapide explication de ce que j'ai fait.

*Se connecter en ssh au serveur proxmox et créer le conteneur*```
#Télécharger le script de création du conteneur LXC
wget https://github.com/tteck/Proxmox/raw/main/ct/debian.sh
#Modifier les droits pour l'exécution
chmod +x ./debian.sh
```
*Exécuter le script et configurer le conteneur avec l'interface connectée au LAN**Créer une seconde interface eth1 pour le conteneur lxc via CLI ou WebUI**Se connecter au conteneur lxc*```
#Lancer la console du conteneur
lxc-console XXX 
#Saisir les identifiants renseignés lors de l'installation du contenur
#Installer iptables + iptables-persistent + tcpdump + igmpproxy
apt-get update;apt-get install iptables iptables-persistent tcpdump igmpproxy
#Bloquer le traffic par défaut pour les chain INPUT, OUTPUT, FORWARD - parce qu'on a besoin de rien d'autre
iptables -P INPUT DROP
iptables -P OUTPUT DROP
iptables -P FORWARD DROP
#Autoriser 12 paquets IGMP par heure vers l'adresse 224.0.0.1/32 et l'interface de sortie eth1
iptables -A OUTPUT -d 224.0.0.1/32 -o eth1 -p igmp -m limit --limit 12/hour -j ACCEPT
#Sauvegarder les règles iptables
netfilter-persistent save
#Quitter la console avec ctrl-a puis q
```
Bon courage à tous.
