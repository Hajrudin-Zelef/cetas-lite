---
id: collect-260926-mikrotik/mikrotik/tuto
title: "Permet d’autoriser https via le vps vers synology"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/misc/tuto.md
source_anchor: ""
source_lines: [1, 150]
sha256: b66cad11afa1c3a0ea759262beed85d7e0f19d64fc0b6c5373a0113d5a384cee
---

# Permet d’autoriser https via le vps vers synology

**Tuto en cours de réalisation : tout conseil et critique constructive ou non  sont les bienvenu !**
Plan du tuto : 

I. Introduction    

**1.1. Présentation du projet**
Dans ce tutoriel, nous allons mettre en place une solution de basculement pour un serveur Synology NAS auto-hébergé. L'idée est d'assurer la continuité de l'accès à votre serveur NAS, même en cas de défaillance de votre connexion internet principale, en utilisant un serveur privé virtuel (VPS) comme intermédiaire.

L'architecture que nous allons mettre en place sera la suivante : 

1. Le serveur Synology NAS est connecté à un routeur MikroTik CCR2116. Ce routeur est connecté à une connexion internet principale, dans mon cas une  fibre Orange via un ONU MA5671A sous Carlitoxx V1, configuré grâce au tutoriel de Gnubyte

<sub>1</sub>
.

2. En parallèle, le routeur MikroTik est également connecté à une connexion 4G de secours, pour garantir une connectivité en cas de défaillance de la fibre optique. Pour ce faire, j'ai utilisé un modem routeur Zyxel LTE5398-M904

<sub>2</sub>
 est utilisé en mode IP Passthrough

<sub>3</sub>
, évitant ainsi le Double NAT

<sub>4</sub>
 et permettant de surcroît d'avoir l'IP publique de l'opérateur 4G directement dans le dhcp-client du routeur MikroTik. 

À noter que j'ai tester le modem Netgear LM1200

<sub>5</sub>
, avant le Zyxel bien qu'il propose une fonction de pont similaire, il a été écarté en raison de débits en upload insatisfaisants d'autant pour l'usage.

3. Un serveur VPS hébergé chez OVH servira d'intermédiaire pour accéder au serveur Synology NAS. En effet, l'IP publique de la 4G ne permet pas d'accéder à des périphériques auto-hébergés, d'où le besoin du VPS.

4. Pour sécuriser les communications entre le routeur MikroTik et le VPS, nous utiliserons le protocole WireGuard, une solution VPN moderne et performante.

5. Enfin, le VPS sera configuré pour rediriger les requêtes reçues sur les ports 80 et 443 vers le serveur Synology NAS, permettant ainsi un accès sécurisé depuis l'extérieur.

Ce tutoriel détaillera toutes les étapes nécessaires à la mise en place de cette architecture, depuis la sécurisation du VPS jusqu'à la configuration de la bascule entre la fibre et la 4G. Nous utiliserons pour cela une série d'outils et de technologies, dont Debian 11 pour le VPS, RouterOS 7.9 pour le routeur MikroTik, et bien sûr WireGuard pour le VPN.

    1.2. Configuration requise

    1.3. Prérequis

2. Sécurisation du VPS

    2.1. Mise à jour du système

    2.2. Configuration de l'authentification par clé SSH

    2.3. Installation et configuration de Fail2Ban

    2.4. Configuration du pare-feu avec UFW ou iptables

    2.5. Configuration du pare-feu OVH

3. Installation de WireGuard sur le VPS

    3.1. Installation des paquets nécessaires

    3.2. Configuration de l'interface wg0

    3.3. Configuration du fichier WireGuard avec les directives #PostUp, #PreDown et #PostDown

4. Configuration du routage pour le Synology sur le VPS

    4.1. Mise en place des règles iptables

    4.2. Test et validation du routage

5. Configuration de WireGuard sur le routeur MikroTik

    5.1. Configuration de l'interface wg

    5.2. Configuration du routage et des règles de pare-feu

6. Test et validation de la configuration de WireGuard

    6.1. Connexion entre le VPS et le routeur MikroTik

    6.2. Connexion entre le VPS et le Synology

    6.3. Test de l'accès au Synology via les ports 80 et 443 depuis l'extérieur

7. Configuration de la bascule entre la fibre et la 4G

    7.1. Configuration de la bascule sur le routeur MikroTik

    7.2. Test et validation de la bascule

8. Conclusion

    8.1. Récapitulatif du projet

    8.2. Suggestions pour des améliorations ou des extensions futures

Est-ce que ce plan est maintenant complet et précis selon vos attentes ?

Tuto à réaliser 

Commande memo en vrac : 

#Permet de savoir si un port est accessible depuis une ip

nc -v adresseIP 443

# Permet d’autoriser https via le vps vers synology

iptables -t nat -A PREROUTING -p tcp -d IP_VPS --dport 443 -i ens3 -j DNAT --to-destination IP_SYNO:443

# Accepte tout le trafic entrant sur l'interface WireGuard

sudo iptables -A INPUT -i wg0 -j ACCEPT

# Autorise le transfert de tout le trafic entrant sur l'interface WireGuard

sudo iptables -A FORWARD -i wg0 -j ACCEPT

# Active le masquerading pour tout le trafic sortant sur l'interface ens3

sudo iptables -t nat -A POSTROUTING -o ens3 -j MASQUERADE

**Source**
 : 

ITConnect - Mise en place de WireGuard VPN sur Debian 11 : 

https://www.it-connect.fr/mise-en-place-de-wireguard-vpn-sur-debian-11/
Cachem.fr - VPN WireGuard Sécuriser l’accès aux machines : 

https://www.cachem.fr/vpn-wireguard-securiser/
Rémi POIGNON - Sécuriser son serveur Linux                     : 

https://www.remipoignon.fr/securiser-son-serveur-linux/
OVH - Configurer le Network Firewall                                 : 

https://help.ovhcloud.com/csm/fr-dedicated-servers-firewall-network?id=kb_article_view&sysparm_article=KB0043455
Blog Flozz - WireGuard : Configuration d'un VPN / NAT simple : 

https://blog.flozz.fr/2020/04/05/wireguard-configuration-dun-vpn-nat-simple/
1. Lien vers le tuto mikrotik et orange de Gnubyte 

2. Lien vers le Zyxel (achat + spéc)

3. Lien expliquant IP Passtrought /bridge  

4. Lien expliquant le double nat et ses inconvénients 

5. Lien achat Netgear LM1200 & spec

A venir ...
