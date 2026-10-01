---
id: collect-261001-cisco/cisco/pugnere-presentations-administration-reseau-config-cisco-pdf-302726c4-3
title: "show int"
domain: cisco
role: reference
task: reference
actors: []
dates: ["2010-10-06"]
keywords: []
source: docs/RAG/collect-261001-cisco/pugnere-presentations-administration-reseau-config-cisco-pdf-302726c4.md
source_anchor: ""
source_lines: [293, 404]
sha256: 4a453773094b5905b477f4c4d82fe52f4db6f0087a78dcca3d1c2140759f72f4
---

# show int

Tutoriel : Configuration d’un commutateur ou d'un routeur CISCO                            2010-10-06/90
Les  VLAN
Les VLANs peuvent être crées localement sur chaque commutateur (appelé : mode VTP  
transparent) ou alors être propagés à l'aide d'un protocole de propagation des VLAN (appelé : mode 
VTP client/serveur).
Configuration des VLAN en mode transparent : 
#conf t
Enter configuration commands, one per line.  End with CNTL/Z.
switch(config)# vtp mode transparent
switch(config-if)# ^Z
Création d'un VLAN localement : 
#conf t
Enter configuration commands, one per line.  End with CNTL/Z.
switch(config)# vlan 3
switch(config)# name vlan-trois
switch(config)# ^Z
Affectation de l'interface fastEthernet1/0/1 au VLAN 3 : 
#conf t
Enter configuration commands, one per line.  End with CNTL/Z.
routeur(config)#interface fastethernet 1/0/1
routeur(config-if)# switchport mode acess
routeur(config-if)# switchport acess vlan 3
routeur(config-if)# ^Z
Activation du VLAN 3 et affectation d'une adresse IP à ce switch dans ce VLAN : 
#conf t
Enter configuration commands, one per line.  End with CNTL/Z.
switch(config)#interface Vlan3
switch(config-if)# no shutdown
switch(config-if)# ip address x.y.z.z 255.255.255.0
switch(config-if)# ^Z
Cisco dispose d'un protocole d'échange de la liste des VLAN configurés dans un commutateur, c'est 
VTP (Vlan Trunking Protocol). C'est un mode client serveur. Un commutateur doit être configuré en 
mode serveur, les autres commutateurs du réseau local doivent être configurés en mode client : 
sur le commutateur « serveur VTP » : 
vlan database
  vtp domain <domaine>
  vtp server
  vlan 2 
  name vlan-deux
  exit
sur le commutateur « client VTP » : 
vlan database
  vtp domain <domaine>
  vtp client
  exit
On peut visualiser si un client a bien reçu des VLAN par le protocole VTP : 
Denis Pugnère (CNRS / IN2P3 / IPNL) Page 8 / 10

Tutoriel : Configuration d’un commutateur ou d'un routeur CISCO                            2010-10-06/90
# sh vtp status
Une fois que le commutateur Client a reçu la liste des VLAN, la commande suivante affiche la liste 
des VLAN reçus ou configurés localement : 
# show vlan brief
Routage
Commandes pour ajouter une route statique :
Route statique : ip route reseau masque <passerelle>
Route par défaut : ip route 0.0.0.0 0.0.0.0 <passerelle>
routeur#conf t
Enter configuration commands, one per line.  End with CNTL/Z.
routeur(config)# ip route 0.0.0.0 0.0.0.0 192.168.1.13
routeur(config)# ip route 192.168.85.4 255.255.255.252 192.168.84.5
routeur(config)# ip route 192.168.85.16 255.255.255.240 192.168.84.5
routeur(config)# ip route 192.168.85.32 255.255.255.240 192.168.84.4
routeur(config)# ^Z
Le commutateur de niveau 2 ne peut pas faire de routage (statique ou dynamique). Pour pouvoir 
réaliser du routage entre VLAN ou entre réseaux IP différents, le commutateur doit supporter le 
routage (donc de niveau 3 de la couche ISO).
Les commutateurs de la série 2xxx (2950, 2960...) sont de niveau 2 (sans routage).
Les commutateurs de la série 3xxx (3750...) sont de niveau 3 (avec routage). 
En fonction des licences disponibles sur le commutateur, il sera capable (ou non) de supporter les 
protocoles de routage dynamique interne (RIP, OSPF...) ou externe (BGP...).
Par exemple, cette commande peut uniquement être prise en compte sur les commutateurs de niveau 
3 : 
routeur#conf t
Enter configuration commands, one per line.  End with CNTL/Z.
routeur(config)# ip routing
routeur(config)# ^Z
Filtrage, access-lists
Les access-lists cisco sont répertoriées par numéro : 
routeur(config)#access-list ?
  <1-99>       IP standard access list
  <100-199>    IP extended access list
  <1100-1199>  Extended 48-bit MAC address access list
  <200-299>    Protocol type-code access list
  <700-799>    48-bit MAC address access list
• IP standard access list : Ne permet d'utiliser que les adresses source pour identifier les 
Denis Pugnère (CNRS / IN2P3 / IPNL) Page 9 / 10

Tutoriel : Configuration d’un commutateur ou d'un routeur CISCO                            2010-10-06/90
paquets
• IP extended access list : Permet d'identifier un paquet par les adresses IP, protocoles et ports 
source et destination
• Protocol type-code access list : Filtrage sur le protocole
• 48-bit MAC address access list : Filtrage en fonction de l'adresse MAC
Filtres sur ICMP
 Filtrer icmp est toujours problématique. C’est un protocole à la fois très utile mais aussi 
potentiellement dangereux. De nombreuse attaques sont basées dessus (« smurf » et autres 
joyeusetées). Parmis ses différentes fonctions on peut trouver : fragmentation (type 3, message du 
type "destination unreachable"), PATH MTU Discovery (type 4), drop des paquets, traceroute, ping, 
controle de flux.
Pour ceux qui veulent filtrer le protocol icmp , voici au moins ce qu’il faut autoriser :
access-list 102 permit icmp any any unreachable parameter-problem source-quench
time-exceeded ttl-exceeded packet-too-big administratively-prohibited
access-list 102 deny icmp any any
Autrement, pour ceux qui autorisent ping , on peut limiter les excès éventuels en ajoutant ceci sur 
l'interface d'entrée :
rate-limit input access-group 2000 744000 10000 10000 conform-action transmit 
exceedaction drop
 avec :
access-list 2000 permit icmp any any echo-reply
access-list 2000 permit icmp any any echo
Denis Pugnère (CNRS / IN2P3 / IPNL) Page 10 / 10
