---
id: collect-261001-fortinet/fortinet/pare-feu-fortigate-comment-configurer-une-interface-vlan-58463c2f-1
title: "Création de l'interface VLAN et configurations de base"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["attribution"]
source: docs/RAG/collect-261001-fortinet/pare-feu-fortigate-comment-configurer-une-interface-vlan-58463c2f.md
source_anchor: ""
source_lines: [1, 51]
sha256: eda49ff445509e852de9326254e953473cddff2a2eff8b4a49395c72ea3beb89
---

# Création de l'interface VLAN et configurations de base

Pare-feu FortiGate : comment configurer une interface VLAN ?
I. Présentation
Dans ce tutoriel, nous allons voir comment configurer une interface VLAN sur un pare-feu FortiGate. Nous évoquerons aussi les différentes options de configuration à notre disposition.
Il est également possible de configurer une ou plusieurs interfaces (ou sous-interfaces) VLAN sur un pare-feu FortiGate. À titre de rappel, un VLAN ou réseau local virtuel, permet de diviser un réseau physique en plusieurs sous-réseaux logiques. Cette segmentation peut être mise en place pour des raisons de fonctionnalité, de gestion ou de sécurité.
Les sous-réseaux virtuels ainsi créés communiqueront entre eux par le biais d'interfaces virtuelles, toutes connectées à une sorte de commutateur virtuel.
Une interface VLAN sous FortiOS se configure de façon très similaire à une interface réseau physique, avec quelques légères variations. On y retrouve souvent les mêmes paramètres de base comme l'adressage IP, la gestion des accès administratifs, la définition du rôle ou encore l'activation d'un service DHCP.
Si ce n'est pas déjà fait, jetez un coup d'œil au tutoriel ci-dessous. Il vous montre comment configurer une interface réseau physique sur un pare-feu FortiGate :
Découvrons ensemble les étapes pour configurer une interface VLAN sur un pare-feu FortiGate.
Remarque 1 : le menu de configuration peut varier en fonction de la version de FortiOS, cependant les options de configuration restent quasiment les mêmes. Ici, nous travaillerons sous FortiOS 7.6.3 qui est l'une des versions les plus récentes.
II. Options de configuration
Il existe plusieurs options de configuration disponibles pour une interface VLAN. Toutefois, les plus importantes sont les suivantes :
- Le nom de l'interface.
- Le type d'interface (ici VLAN).
- Le protocole du VLAN (802.1q ou 802.1qd).
- Le choix de l'interface physique sur laquelle nous voulons créer notre interface virtuelle. Une interface physique est capable de supporter environ 4096 VLANs.
- L'ID du VLAN qui permet d'identifier le VLAN de manière unique sur le réseau. Après la création de l'interface, vous ne pouvez pas modifier ou changer cet identifiant, ni le nom de l'interface.
- L'ID de virtual routing and forwarding.
- Le rôle (LAN, WAN, DMZ, indéfini).
Comme options toutes aussi importantes, mais pas indispensables, nous avons :
- Le mode d'adressage qui sert à définir la méthode d'attribution de l'adresse IP, du masque de sous-réseau, de la passerelle, etc. à l'interface réseau. Les options disponibles sont la configuration manuelle, le DHCP, l'IPAM, le PPPoE ou le Capture One-Arm.
- L'accès administratif indique les protocoles et services réseau que l'on souhaite laisser transiter au niveau l'interface à configurer (http, icmp, ssh, telnet, etc.).
La configuration d'une interface VLAN peut également se faire via l'interface graphique Web ou via la ligne de commande.
III. Exemple de configuration
A. Via l'interface graphique Web
Pour configurer une interface VLAN sur un pare-feu FortiGate depuis une interface graphique Web, on procède comme suit :
- On se rend sur le panneau latéral droit, on déroule le menu-flèche Réseau et on clique surInterfaces .
- On va ensuite dérouler le menu-flèche Nouveau et on clique surInterface .
La configuration présentée ici est juste à titre illustratif ! N'hésitez pas à l'ajuster ou à la modifier pour qu'elle corresponde parfaitement à vos exigences fonctionnelles et sécuritaires.
Voici donc un exemple de configuration :
- Nom : Mon_VLAN
- Type : VLAN
- Protocole VLAN : 802.1q
- Interface : port2
- ID VLAN : 100
- ID de VRF : 0
- Rôle : LAN
- Mode d'adressage : manuel (192.168.100.1/24 )
- Accès administratif : https, ping, ssh en IPv4
- Serveur DHCP : activé, avec une plage d'adresses pour les clients allant de 192.168.100.2 à192.168.100.254 , l'IP de l'interface comme passerelle par défaut et le DNS du système comme serveur DNS. Le bail DHCP qui est de 604800 secondes (7 jours), est également le bail par défaut.
- Détection d'équipements pour le filtrage d'adresses MAC.
En images :
Remarque 2 : D'autres options telles que la gestion de la bande passante, la configuration d'un portail captif, l'ajout de serveur TFTP, etc. sont également disponibles pour la configuration d'une interface VLAN. Nous aborderons ces points dans nos prochains tutoriels.
Une fois la configuration terminée, nous pouvons vérifier que la nouvelle interface VLAN ait bien été créée. Pour cela, on retourne dans le menu de configuration d'interfaces, on localise le port physique abritant le VLAN et on va cliquer sur le bouton + à gauche de ce dernier pour dérouler. 
B. En ligne de commande
La configuration en ligne de commande peut se faire via un terminal SSH distant, en se connectant directement au port console du pare-feu ou en accédant à la console disponible via l'interface graphique web. Cette option est généralement utilisée par les administrateurs plus expérimentés.
Comme toujours, la configuration en ligne de commande fait gagner du temps, en particulier quand on doit configurer plusieurs interfaces simultanément. On peut, par exemple, copier la configuration d'un port, la coller et modifier les détails (Nom du VLAN, ID de VLAN, interface physique, index SNMP, etc.) dans un éditeur de texte sur un PC pour configurer plusieurs ports rapidement. Il suffit ensuite de copier-coller la nouvelle configuration dans le terminal et le tour est joué.
De plus, la ligne de commande offre des options de configuration qui ne sont pas toujours disponibles en mode graphique. Sachez par ailleurs que, une fois dans le mode de configuration d'interfaces (config system interface > edit "nom de l'interface"), l'ordre des commandes n'a plus d'importance. Gardez aussi à l'esprit que la configuration des options clés, évoquées ci-dessus, reste indispensable et utilisez toujours la commande end pour valider la configuration ! 
Voici un exemple de configuration en ligne de commande :
# Création de l'interface VLAN et configurations de base
Meister-Gate # config system interface
    
