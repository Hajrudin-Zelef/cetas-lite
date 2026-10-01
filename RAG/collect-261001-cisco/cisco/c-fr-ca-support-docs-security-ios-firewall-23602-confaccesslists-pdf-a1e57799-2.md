---
id: collect-261001-cisco/cisco/c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799-2
title: "c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799.md
source_anchor: ""
source_lines: [148, 301]
sha256: ac9e17c440383a29685dcde615fcd64c0890d6bd56321a58a28ca39273b7acf2
---

# c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799

Les deux premiers octets et le dernier octet sont les mêmes pour chaque réseau. Le tableau ci-
dessous explique comment les récapituler.
Le troisième octet des réseaux précédents peut être écrit comme indiqué dans ce tableau, 
correspondant à la position de bit d’octet et à la valeur d’adresse pour chaque bit.
Décimal 128 64 32 16 8 4 2 1
146 1 0 0 1 0 0 1 0
147 1 0 0 1 0 0 1 1
148 1 0 0 1 0 1 0 0
149 1 0 0 1 0 1 0 1
L L L L L ? ? ?
À la différence de l'exemple précédent, vous ne pouvez pas récapituler ces réseaux en un seul 
réseau. S'ils sont récapitulés en un seul réseau, ils deviennent 192.168.144.0/21 parce que cinq 
bits sont semblables dans le troisième octet. Ce réseau résumé 192.168.144.0/21 couvre une 
plage de réseaux allant de 192.168.144.0 à 192.168.151.0. Parmi ceux-ci, 192.168.144.0, 
192.168.145.0, 192.168.150.0 et 192.168.151.0 ne figurent pas dans la liste de quatre réseaux 
donnée. Afin de couvrir les réseaux spécifiques en question, vous avez besoin d'un minimum de 
deux réseaux récapitulés. Les quatre réseaux donnés peuvent être récapitulés dans ces deux 
réseaux :
Pour les réseaux 192.168.146.x et 192.168.147.x, tous les bits correspondent à l’exception 
du dernier, qui est une erreur. Il peut être écrit comme 192.168.146.0/23 (ou 192.168.146.0 
255.255.254.0).
•
Pour les réseaux 192.168.148.x et 192.168.149.x, tous les bits correspondent à l’exception 
du dernier, qui est une erreur. Il peut être écrit comme 192.168.148.0/23 (ou 192.168.148.0 
255.255.254.0).
•
Ce résultat définit une liste de contrôle d’accès récapitulée pour les réseaux précédents.
 
<#root>
 
!--- This command is used to allow access access for devices with IP  
!--- addresses in the range from 192.168.146.0 to 192.168.147.254. 
 
access-list 10 permit 192.168.146.0 0.0.1.255  
 
 
<#root>

!--- This command is used to allow access access for devices with IP  
!--- addresses in the range from 192.168.148.0 to 192.168.149.254 
 
access-list 10 permit 192.168.148.0 0.0.1.255
 
Traiter les listes de contrôle d'accès
Le trafic qui entre dans le routeur est comparé aux entrées de la liste de contrôle d'accès basées 
sur l'ordre dans lequel les entrées arrivent dans le routeur. De nouvelles instructions sont ajoutées 
à la fin de la liste. Le routeur continue de chercher jusqu'à ce qu'il trouve une correspondance. Si 
aucune correspondance n'est trouvée quand le routeur atteint la fin de la liste, le trafic est refusé. 
Pour cette raison, vous devez avoir les entrées fréquemment consultées en haut de la liste. 
Chaque liste de contrôle d’accès a un refus implicite à la fin pour tout trafic qui n’est pas 
explicitement autorisé. Une liste de contrôle d’accès à entrée unique avec une seule entrée de 
refus peut refuser tout le trafic. Vous devez avoir au moins une instruction d'autorisation dans une 
liste de contrôle d'accès ou tout le trafic est bloqué. Ces deux listes de contrôle d'accès (101 et 
102) ont le même effet.
 
<#root>
 
!--- This command is used to permit IP traffic from 10.1.1.0  
!--- network to 172.16.1.0 network. All packets with a source  
!--- address not in this range will be rejected. 
 
access-list 101 permit ip 10.1.1.0 0.0.0.255 172.16.1.0 0.0.0.255
 
 
<#root>
 
!--- This command is used to permit IP traffic from 10.1.1.0  
!--- network to 172.16.1.0 network. All packets with a source  
!--- address not in this range will be rejected. 
 
access-list 102 permit ip 10.1.1.0 0.0.0.255 172.16.1.0 0.0.0.255    
access-list 102 deny ip any any 
 
Dans l'exemple suivant, la dernière entrée est suffisante. Vous n'avez pas besoin des trois 
premières entrées car IP inclut TCP, le protocole UDP (User Datagram Protocol) et le protocole 
ICMP (Internet Control Message Protocol).
 
<#root>
 
!--- This command is used to permit Telnet traffic

!--- from machine 10.1.1.2 to machine 172.16.1.1. 
 
access-list 101 permit tcp host 10.1.1.2 host 172.16.1.1 eq telnet
 
 
<#root>
 
!--- This command is used to permit tcp traffic from  
!--- 10.1.1.2 host machine to 172.16.1.1 host machine. 
 
access-list 101 permit tcp host 10.1.1.2 host 172.16.1.1
 
 
<#root>
 
!--- This command is used to permit udp traffic from 
 !--- 10.1.1.2 host machine to 172.16.1.1 host machine. 
 
access-list 101 permit udp host 10.1.1.2 host 172.16.1.1
 
 
<#root>
 
!--- This command is used to permit ip traffic from  
!--- 10.1.1.0 network to 172.16.1.10 network. 
 
access-list 101 permit ip 10.1.1.0 0.0.0.255 172.16.1.0 0.0.0.255
 
Définir les ports et les types de messages
Vous pouvez non seulement définir la source et la destination de la liste de contrôle d’accès, mais 
également définir les ports, les types de message ICMP et d’autres paramètres. Une source utile 
d'informations pour les ports connus est la RFC 1700. Les types de messages ICMP sont décrits 
dans la RFC 792.
Le routeur peut afficher un texte descriptif sur certains des ports connus. Sélectionnez un ?pour 
obtenir de l'aide.
 
<#root>
access-list 102 permit tcp host 10.1.1.1 host 172.16.1.1 eq ?
  bgp          Border Gateway Protocol (179) 
  chargen      Character generator (19) 
  cmd          Remote commands (rcmd, 514)

Pendant la configuration, le routeur convertit également des valeurs numériques en valeurs plus 
conviviales. Il s’agit d’un exemple dans lequel vous tapez le numéro de type de message ICMP et 
le routeur le convertit en nom.
 
<#root>
access-list 102 permit icmp host 10.1.1.1 host 172.16.1.1 14
 
devient
 
<#root>
access-list 102 permit icmp host 10.1.1.1 host 172.16.1.1 timestamp-reply
 
Appliquer les listes de contrôle d'accès
Vous pouvez définir des listes de contrôle d’accès sans les appliquer. Toutefois, les listes de 
contrôle d'accès n'ont aucun effet tant qu'elles ne sont pas appliquées à l'interface du routeur. Il 
est judicieux d'appliquer la liste de contrôle d'accès sur l'interface la plus proche de la source du 
trafic. Comme le montre cet exemple, lorsque vous essayez de bloquer le trafic de la source à la 
destination, vous pouvez appliquer une liste de contrôle d’accès entrante à E0 sur le routeur A au 
lieu d’une liste sortante à E1 sur le routeur C. Une liste de contrôle d’accès a une instruction deny 
ip any implicitement à la fin de toute liste de contrôle d’accès. Si le trafic est lié à une requête 
DHCP et s’il n’est pas explicitement autorisé, le trafic est abandonné car lorsque vous examinez la 
requête DHCP dans IP, l’adresse source est s=0.0.0.0 (Ethernet1/0), d=255.255.255.255, len 604, 
rcvd 2 UDP src=68, dst=67. Notez que l’adresse IP source est 0.0.0.0 et l’adresse de destination 
est 255.255.255.255. Le port source est 68 et destination 67. Par conséquent, vous devez 
autoriser ce type de trafic dans votre liste de contrôle d’accès, sinon le trafic est abandonné en 
raison d’un refus implicite à la fin de l’instruction.
Remarque : Pour que le trafic UDP transite, il doit également être autorisé explicitement par 
la liste de contrôle d’accès.
Définir les termes interne, externe, entrant, sortant, source et destination
Le routeur utilise les termes interne, externe, source et destination comme références. Le trafic

