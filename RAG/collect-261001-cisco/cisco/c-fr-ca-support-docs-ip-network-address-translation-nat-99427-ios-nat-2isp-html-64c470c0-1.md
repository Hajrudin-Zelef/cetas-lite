---
id: collect-261001-cisco/cisco/c-fr-ca-support-docs-ip-network-address-translation-nat-99427-ios-nat-2isp-html-64c470c0-1
title: "c-fr-ca-support-docs-ip-network-address-translation-nat-99427-ios-nat-2isp-html-64c470c0"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["full-duplex"]
source: docs/RAG/collect-261001-cisco/c-fr-ca-support-docs-ip-network-address-translation-nat-99427-ios-nat-2isp-html-64c470c0.md
source_anchor: ""
source_lines: [1, 25]
sha256: 45618904ffdfc95e8ae926a80a6fa2c89e040f41e899da560eaf4296ab6fe001
---

# c-fr-ca-support-docs-ip-network-address-translation-nat-99427-ios-nat-2isp-html-64c470c0

Ce document décrit la configuration d'un routeur Cisco IOS® pour connecter un réseau à Internet avec traduction d'adresses réseau via deux connexions ISP. La traduction d’adresses de réseau (NAT) du logiciel Cisco IOS peut distribuer les connexions TCP et les sessions UDP suivantes sur plusieurs connexions réseau si des routes à coût égal vers une destination donnée sont disponibles. Si l’une des connexions devient inutilisable, le suivi d’objets, un composant du routage de périphérie optimisé (OER), peut être utilisé pour désactiver la route jusqu’à ce que la connexion soit à nouveau disponible, ce qui garantit la disponibilité du réseau en dépit de l’instabilité ou du manque de fiabilité d’une connexion Internet.
Ce document suppose que vous avez des connexions LAN et WAN fonctionnelles ; il ne fournit pas d'arrière-plan de configuration ou de dépannage pour établir la connectivité initiale.
Ce document ne décrit pas un moyen de différencier les routes, il n'y a donc aucun moyen de préférer une connexion plus souhaitable à une connexion moins souhaitable.
Ce document décrit la configuration d'OER pour activer ou désactiver l'une ou l'autre route Internet en fonction de l'accessibilité des serveurs DNS du FAI. Vous devez identifier des hôtes spécifiques qui peuvent être accessibles via une seule des connexions ISP et qui ne peuvent pas être disponibles si cette connexion ISP n'est pas disponible.
Cette configuration a été développée avec un routeur Cisco 1811 avec le logiciel Advanced IP Services 12.4(15)T. Si une autre version du logiciel est utilisée, certaines fonctionnalités ne sont peut-être pas disponibles ou les commandes de configuration peuvent différer de celles présentées dans ce document. Des configurations similaires sont disponibles sur toutes les plates-formes de routeur Cisco IOS, bien que la configuration de l’interface varie probablement d’une plate-forme à l’autre.
The information in this document was created from the devices in a specific lab environment. All of the devices used in this document started with a cleared (default) configuration. If your network is live, make sure that you understand the potential impact of any command.
Pour plus d'informations sur les conventions utilisées dans ce document, reportez-vous à Conventions relatives aux conseils techniques Cisco.
Vous devez potentiellement ajouter un routage basé sur des stratégies pour un trafic spécifique pour vous assurer qu'il utilise toujours une connexion ISP. Les clients VPN IPSec, les combinés VoIP et tout autre trafic qui n'utilise qu'une seule des options de connexion du FAI pour préférer la même adresse IP, une vitesse supérieure ou une latence inférieure à la connexion sont des exemples de trafic nécessitant ce comportement.
Remarque : Utilisez l’outil de recherche de commandes (clients inscrits seulement) pour en savoir plus sur les commandes figurant dans le présent document.
Ce document utilise la configuration réseau suivante :
Cet exemple de configuration, illustré dans le schéma de réseau, décrit un routeur d’accès qui utilise une connexion IP configurée par DHCP à un FAI (comme illustré par FastEthernet 0) et une connexion PPPoE sur l’autre connexion FAI. Les types de connexion n'ont aucun impact particulier sur la configuration, sauf si le suivi d'objet et le routage OER et/ou basé sur des stratégies doivent être utilisés avec une connexion Internet attribuée par DHCP. Dans ces cas, il peut être très difficile de définir un routeur de tronçon suivant pour le routage de stratégie ou OER.
| Exemple de configuration de routeur | 
|---|
| track timer interface 5 ! ! Configure timers on route tracking ! track 123 rtr 1 reachability  delay down 15 up 10 ! track 345 rtr 2 reachability  delay down 15 up 10 ! ! Use “ip dhcp client route track [number]”     ! to monitor route on DHCP interfaces ! Define ISP-facing interfaces with “ip nat outside” ! interface FastEthernet0  ip address dhcp  ip dhcp client route track 345  ip nat outside  ip virtual-reassembly ! interface FastEthernet1  no ip address  pppoe enable  no cdp enable ! interface FastEthernet2  no cdp enable ! interface FastEthernet3  no cdp enable ! interface FastEthernet4  no cdp enable ! interface FastEthernet5  no cdp enable ! interface FastEthernet6  no cdp enable ! interface FastEthernet7  no cdp enable ! interface FastEthernet8  no cdp enable ! interface FastEthernet9  no cdp enable ! ! Define LAN-facing interfaces with “ip nat inside” ! interface Vlan1  description LAN Interface  ip address 192.168.108.1 255.255.255.0  ip nat inside  ip virtual-reassembly  ip tcp adjust-mss 1452 ! ! Define ISP-facing interfaces with “ip nat outside” ! Interface Dialer 0  description PPPoX dialer  ip address negotiated  ip nat outside  ip virtual-reassembly  ip tcp adjust-mss ! ip route 0.0.0.0 0.0.0.0 dialer 0 track 123 ! ! Configure NAT overload (PAT) to use route-maps ! ip nat inside source route-map fixed-nat    interface Dialer0 overload ip nat inside source route-map dhcp-nat     interface FastEthernet0 overload ! ! Configure an OER tracking entry     ! to monitor the first ISP connection ! ip sla 1  icmp-echo 172.16.108.1 source-interface Dialer0  timeout 1000  threshold 40  frequency 3 ! ! Configure a second OER tracking entry     ! to monitor the second ISP connection ! ip sla 2  icmp-echo 172.16.106.1 source-interface FastEthernet0  timeout 1000  threshold 40  frequency 3 ! ! Set the SLA schedule and duration ! ip sla schedule 1 life forever start-time now ip sla schedule 2 life forever start-time now ! ! Define ACLs for traffic that     ! will be NATed to the ISP connections ! access-list 110 permit ip 192.168.108.0 0.0.0.255 any ! ! Route-maps associate NAT ACLs with NAT     ! outside on the ISP-facing interfaces ! route-map fixed-nat permit 10  match ip address 110  match interface Dialer0 ! route-map dhcp-nat permit 10  match ip address 110  match interface FastEthernet0   | 
Avec le suivi de route attribué par DHCP :
| Exemple de configuration du suivi de route attribué par DHCP (facultatif) | 
|---|
| interface FastEthernet0  description Internet Intf  ip dhcp client route track 123  ip address dhcp  ip nat outside  ip virtual-reassembly  speed 100  full-duplex  no cdp enable  | 
Référez-vous à cette section pour vous assurer du bon fonctionnement de votre configuration.
L'Outil Interpréteur de sortie (clients enregistrés uniquement) (OIT) prend en charge certaines commandes show. Utilisez l'OIT pour afficher une analyse de la sortie de la commande show .
show ip nat translation - Affiche l'activité NAT entre les hôtes internes NAT et les hôtes NAT extérieurs. Cette commande permet de vérifier que les hôtes internes sont traduits en adresses externes NAT.
 
      Router# sh ip nat tra Pro Inside global Inside local Outside local Outside global tcp 172.16.108.44:54486 192.168.108.3:54486 172.16.104.10:22 172.16.104.10:22 tcp 172.16.106.42:49620 192.168.108.3:49620 172.16.102.11:80 172.16.102.11:80 tcp 172.16.108.44:1623 192.168.108.4:1623 172.16.102.11:445 172.16.102.11:445 Router#
show ip route - Vérifie que plusieurs itinéraires vers Internet sont disponibles.
 
