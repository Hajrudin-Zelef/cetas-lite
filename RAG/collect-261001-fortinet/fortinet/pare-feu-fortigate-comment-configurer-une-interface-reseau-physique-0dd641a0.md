---
id: collect-261001-fortinet/fortinet/pare-feu-fortigate-comment-configurer-une-interface-reseau-physique-0dd641a0
title: "pare-feu-fortigate-comment-configurer-une-interface-reseau-physique-0dd641a0"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["attention", "attribution", "ethernet"]
source: docs/RAG/collect-261001-fortinet/pare-feu-fortigate-comment-configurer-une-interface-reseau-physique-0dd641a0.md
source_anchor: ""
source_lines: [1, 71]
sha256: 3200b8bf781ffd485dfcb96a266edd78a1f1b255d18daca66b4358653079b6a8
---

# pare-feu-fortigate-comment-configurer-une-interface-reseau-physique-0dd641a0

Pare-feu FortiGate : comment configurer une interface réseau physique ?
I. Présentation
Dans ce tutoriel, nous allons voir comment configurer une interface réseau physique sur un pare-feu FortiGate. Nous évoquerons aussi les différentes options de configuration à votre disposition.
Les interfaces réseaux physiques d'un pare-feu lui permettent de se connecter et de communiquer avec d'autres appareils physiques du réseau, comme des routeurs, des commutateurs ou des ordinateurs. Sur un pare-feu FortiGate, ces interfaces se font appeler port.
Selon le modèle du pare-feu, on peut y retrouver des ports Ethernet (RJ-45) ou optiques (SPF, SPF+, QSPF+, etc.). L'objet de cet article est de montrer comment configurer une interface réseau physique sur un pare-feu FortiGate.
II. Options de configuration
Il existe plusieurs options de configuration applicables sur des interfaces physiques d'un pare-feu FortiGate. Toutefois, les plus essentielles et les plus basiques pour établir une communication quelconque sont les suivantes :
- L'alias qui est un nom alternatif utilisé pour désigner le port en question. Il ne remplace cependant pas le nom du port physique.
- Le type d'interface, ici interface physique.
- Le rôle permet de délimiter le champ d'action physique et logique de l'interface. En effet, il indique si l'interface sera utilisée pour le réseau local (LAN), les réseaux externes (WAN), la DMZ, ou si son usage n'est pas encore défini.
- Le mode d'adressage sert à définir la méthode d'attribution de l'adresse IP, du masque de sous-réseau, de la passerelle, etc. à l'interface réseau. Les options disponibles sont la configuration manuelle, le DHCP, l'IPAM, le PPPoE ou le Capture One-Arm
- L'accès administratif indique les protocoles et services réseau que l'on souhaite laisser transiter au niveau l'interface à configurer (http, icmp, ssh, telnet, etc.).
Il existe également d'autres options de configuration plus avancées telles que la gestion de la bande passante, la détection d'équipements, le déploiement d'un serveur DHCP, etc. Nous évoquerons ces dernières dans d'autres tutoriels.
III. Exemple de configuration
A. Via l'interface graphique web
Pour configurer un port physique sur un pare-feu FortiGate depuis son interface d'administration graphique, on procède comme suit :
- On se rend sur le panneau latéral droit, on déroule le menu-flèche Réseau et on clique surInterfaces .
- On choisit ensuite le port à configurer et on lui assigne les paramètres souhaités.
Bien sûr, cette configuration n'est pas figée. N'hésitez pas à l'ajuster ou à la modifier pour qu'elle corresponde parfaitement à vos exigences fonctionnelles et sécuritaires.
Ici, nous allons juste effectuer une configuration basique. Le port1 sera manuellement configuré avec l'adresse IP 192.168.1.192 et le masque de sous-réseau 255.255.255.0. Les protocoles autorisés sur l'interface sont HTTP(S), HTTP, SSH et PING (ICMP). Ceci a un impact sur l'administration du pare-feu, donc attention en fonction de l'usage de l'interface (LAN, WAN, etc...).
Ce qui donne en image :
Si vous y avez prêté attention, l'option de configuration en IPv6 ne semble pas être disponible ici. Eh bien, c'est tout simplement parce qu'elle est désactivée dans les fonctionnalités du pare-feu. Pour la réactiver, on déroule le menu Système et on clique sur Fonctionnalités. On va ensuite cliquer sur la bascule IPv6 et on valide le tout.
Et, si on retourne dans le menu de configuration d'interfaces, on devrait normalement trouver toutes les options de configuration IPv6 disponibles.
B. En ligne de commande
Nous pouvons aussi configurer le pare-feu en utilisant la ligne de commande de plusieurs façons :
- Via un terminal SSH distant.
- En nous connectant directement au port console du pare-feu.
- En accédant à la console disponible via l'interface graphique web du pare-feu.
La configuration en ligne de commande fait gagner du temps, en particulier quand on doit configurer plusieurs interfaces simultanément. On peut, par exemple, copier la configuration d'un port, la coller et modifier les détails (numéro de port, index SNMP, etc.) dans un éditeur de texte sur un PC pour configurer plusieurs ports rapidement. Il suffit ensuite de copier-coller la nouvelle configuration dans le terminal et le tour est joué.
L'autre avantage lié à l'utilisation de la CLI est la possibilité de disposer de plus d'options de configuration qu'en mode graphique.
Allons-y voir comment configurer une interface réseau en ligne de commande.
Nous sommes donc ici sur un terminal en ligne de commande. D'abord, on rentre en mode configuration d'interfaces avec la commande config system interface. Puis, on choisit de configurer le port3 avec la commande edit port3. 
La commande set suivie d'un ou de plusieurs arguments permet d'appliquer des configurations à l'interface choisie. Ici nous allons juste demander à notre interface de se soumettre à la configuration DHCP de notre LAN : la commande est donc set mode dhcp. On va enfin activer les protocoles http, https et icmp (ping) avec la commande set allowaccess https http ping. On sauvegarde le tout avec la commande end.
Ce qui donne finalement cet ensemble de commandes :
Meister-Gate # config system interface
Meister-Gate (interface) # edit port3
Meister-Gate (port3) # set mode dhcp
Meister-Gate (port3) # set allowaccess https http ping
Meister-Gate (port3) # end
En image, la même chose :
On peut aller vérifier l'application de la configuration avec la commande show system interface port3.
Meister-Gate # show system interface port3
config system interface
    edit "port3"
        set vdom "root"
        set mode dhcp
        set allowaccess ping https http
        set type physical
        set snmp-index 3
    next
end
En image :
Pour retrouver la configuration IP obtenue par DHCP, on saisit la commande get system interface physical.
Meister-Gate # get system interface physical
== [onboard]
        ....
        ==[port3]
                mode: dhcp
                ip: 192.168.1.104 255.255.255.0
                ipv6: ::/0
                status: up
                speed: 10000Mbps (Duplex: full)
                FEC: none
                FEC_cap: none
Dans la configuration globale des interfaces, nous retrouvons bien les paramètres de notre port3 :
IV. Conclusion
L'objectif principal de la configuration des interfaces physiques d'un pare-feu est d'établir sa communication avec les autres équipements physiques du réseau. La mise en place de services avancés comme les VLANs ou les tunnels VPN requiert la création et la configuration d'une interface réseau virtuelle. Nous aborderons ce point dans nos prochains tutoriels.
Ce qui est essentiel d'un point de vue de la sécurité, c'est de contrôler les protocoles d'administration activés sur les différentes interfaces. Autrement dit, n'activez pas l'interface HTTPS ni le SSH sur l'interface exposée sur Internet.
Ce tutoriel est maintenant terminé. On attend vos retours ! Laissez un commentaire ou rejoignez la discussion sur Discord.
Envie de vous exercer ? Créez un pare-feu virtuel avec VMware Workstation et FortiGate VM !
Bravo 🤗🤗
