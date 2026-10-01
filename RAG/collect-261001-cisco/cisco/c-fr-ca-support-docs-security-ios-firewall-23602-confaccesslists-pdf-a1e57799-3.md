---
id: collect-261001-cisco/cisco/c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799-3
title: "c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-cisco/c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799.md
source_anchor: ""
source_lines: [302, 551]
sha256: 0f3962fb30069aa8cf3af624782b372d7922c9f698409a6657fddfac96918290
---

# c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799

sur le routeur peut être comparé au trafic sur l'autoroute. Si vous étiez un officier de police en 
Pennsylvanie et que vous vouliez arrêter un camion qui va du Maryland à New York, la source du 
camion est le Maryland, et la destination du camion est New York. Le barrage routier pourrait être 
appliqué à la frontière entre la Pennsylvanie et New York (sortie) ou à la frontière entre le 
Maryland et la Pennsylvanie (entrée).
Quand vous faites référence à un routeur, ces termes ont les significations suivantes.
Externe : trafic qui est déjà passé par le routeur et quitte l'interface. La source est l'endroit où 
il était, de l'autre côté du routeur, et la destination est l'endroit où il va.
•
Interne : trafic qui arrive sur l'interface, puis passe par le routeur. La source est l'endroit où il 
était et la destination est l'endroit où il va, de l'autre côté du routeur.
•
Entrant : si la liste d'accès est entrante, quand le routeur reçoit un paquet, le logiciel 
Cisco IOS recherche une correspondance dans les instructions de critères de la liste 
d'accès. Si le paquet est autorisé, le logiciel continue de le traiter. Si le paquet est refusé, le 
logiciel ignore le paquet.
•
Sortant : si la liste d'accès est sortante, une fois que le logiciel a reçu un paquet et l'a routé 
vers l'interface sortante, il recherche une correspondance dans les instructions de critères de 
la liste d'accès. Si le paquet est autorisé, le logiciel le transmet. Si le paquet est refusé, le 
logiciel ignore le paquet.
•
La liste de contrôle d'accès interne a une source sur un segment de l'interface à laquelle elle est 
appliquée et une destination hors de tout autre interface. La liste de contrôle d'accès externe a 
une source sur un segment d'une interface autre que celle à laquelle elle est appliquée et une 
destination hors de l'interface à laquelle elle est appliquée.
Modifier les listes de contrôle d'accès
Lorsque vous modifiez une liste de contrôle d'accès, vous devez être particulièrement vigilant. Par 
exemple, si vous avez l'intention de supprimer une ligne spécifique d'une liste de contrôle d'accès 
numérotée qui existe comme illustré ici, toute la liste est supprimée.
 
<#root>
 
!--- The access-list 101 denies icmp from any to any network  
!--- but permits IP traffic from any to any network. 
 
  router#
configure terminal
  Enter configuration commands, one per line.  End with CNTL/Z. 
  router(config)#
access-list 101 deny icmp any any
  router(config)#
access-list 101 permit ip any any

router(config)#
^Z
 
  router#
show access-list
  Extended IP access list 101 
      deny icmp any any 
      permit ip any any 
  router# 
  *Mar  9 00:43:12.784: %SYS-5-CONFIG_I: Configured from console by console 
 
  router#
configure terminal
  Enter configuration commands, one per line.  End with CNTL/Z. 
  router(config)#
no access-list 101 deny icmp any any
  router(config)#
^Z
 
  router#
show access-list
  router# 
  *Mar  9 00:43:29.832: %SYS-5-CONFIG_I: Configured from console by console
 
Copiez la configuration du routeur sur un serveur TFTP ou un éditeur de texte, tel que le Bloc-
notes, afin de modifier des listes de contrôle d'accès numérotées. Apportez ensuite toutes les 
modifications nécessaires et recopiez la configuration sur le routeur.
Vous pouvez également procéder ainsi.
 
<#root>
router#
configure terminal
  Enter configuration commands, one per line. 
  router(config)#
ip access-list extended test
 
!--- Permits IP traffic from 10.2.2.2 host machine to 10.3.3.3 host machine. 
 
  router(config-ext-nacl)#
permit ip host 10.2.2.2 host 10.3.3.3
 
!--- Permits www traffic from 10.1.1.1 host machine to 10.5.5.5 host machine.

router(config-ext-nacl)#
permit tcp host 10.1.1.1 host 10.5.5.5 eq www
 
!--- Permits icmp traffic from any to any network. 
 
  router(config-ext-nacl)#
permit icmp any any
 
!--- Permits dns traffic from 10.6.6.6 host machine to 10.10.10.0 network. 
 
  router(config-ext-nacl)#
permit udp host 10.6.6.6 10.10.10.0 0.0.0.255 eq domain
  router(config-ext-nacl)#^Z 
  1d00h: %SYS-5-CONFIG_I: Configured from console by consoles-l 
 
  router#
show access-list
  Extended IP access list test 
      permit ip host 10.2.2.2 host 10.3.3.3 
      permit tcp host 10.1.1.1 host 10.5.5.5 eq www 
      permit icmp any any 
      permit udp host 10.6.6.6 10.10.10.0 0.0.0.255 eq domain
 
Toutes les suppressions sont retirées de la liste de contrôle d'accès et tous les ajouts sont 
apportés à la fin de la liste de contrôle d'accès.
 
<#root>
router#
configure terminal
   Enter configuration commands, one per line.  End with CNTL/Z. 
   router(config)#
ip access-list extended test
 
!--- ACL entry deleted. 
 
   router(config-ext-nacl)#
no permit icmp any any
 
!--- ACL entry added. 
 
   router(config-ext-nacl)#
permit gre host 10.4.4.4 host 10.8.8.8
   router(config-ext-nacl)#
^Z

1d00h: %SYS-5-CONFIG_I: Configured from console by consoles-l 
 
   router#
show access-list
   Extended IP access list test 
       permit ip host 10.2.2.2 host 10.3.3.3 
       permit tcp host 10.1.1.1 host 10.5.5.5 eq www 
       permit udp host 10.6.6.6 10.10.10.0 0.0.0.255 eq domain 
       permit gre host 10.4.4.4 host 10.8.8.8
 
Vous pouvez également ajouter des lignes de liste de contrôle d'accès à des listes de contrôle 
d'accès standard ou étendues numérotées par numéro de séquence dans Cisco IOS. Voici un 
exemple de la configuration :
Configurez la liste de contrôle d'accès étendue de la façon suivante :
 
<#root>
Router(config)#
access-list 101 permit tcp any any
Router(config)#
access-list 101 permit udp any any
Router(config)#
access-list 101 permit icmp any any
Router(config)#
exit
Router#
 
Émettez la commande show access-list afin d'afficher les entrées ACL. Les numéros de séquence 
tels que 10, 20 et 30 apparaissent également ici.
 
<#root>
Router#
show access-list
Extended IP access list 101 
    10 permit tcp any any 
    20 permit udp any any 
    30 permit icmp any any
 
Ajoutez l'entrée pour la liste d'accès 101 avec le numéro de séquence 5.

Exemple 1 :
 
<#root>
Router#
configure terminal
Enter configuration commands, one per line.  End with CNTL/Z. 
Router(config)#
ip access-list extended 101
Router(config-ext-nacl)#
5 deny tcp any any eq telnet
Router(config-ext-nacl)#
exit
Router(config)#
exit
Router#
 
Dans la sortie de la commande show access-list, la liste de contrôle d'accès numéro d'ordre 5 est 
ajoutée comme première entrée à la liste de contrôle d'accès 101.
 
<#root>
Router#
show access-list
Extended IP access list 101 
 
5 deny tcp any any eq telnet
    10 permit tcp any any 
    20 permit udp any any 
    30 permit icmp any any 
Router#
 
Exemple 2 :
 
<#root>
internetrouter#
show access-lists
Extended IP access list 101 
    10 permit tcp any any

15 permit tcp any host 172.16.2.9 
    20 permit udp host 172.16.1.21 any 
    30 permit udp host 172.16.1.22 any 
 
internetrouter#
configure terminal
Enter configuration commands, one per line.  End with CNTL/Z. 
internetrouter(config)#
ip access-list extended 101
internetrouter(config-ext-nacl)#
18 per tcp any host 172.16.2.11
internetrouter(config-ext-nacl)#
^Z
 
internetrouter#
show access-lists
Extended IP access list 101 
    10 permit tcp any any 
    15 permit tcp any host 172.16.2.9 
    18 permit tcp any host 172.16.2.11 
    20 permit udp host 172.16.1.21 any 
    30 permit udp host 172.16.1.22 any 
internetrouter#
 
De même, vous pouvez configurer la liste d'accès standard de la façon suivante :
 
<#root>
internetrouter(config)#
access-list 2 permit 172.16.1.2
internetrouter(config)#
access-list 2 permit 172.16.1.10
internetrouter(config)#
access-list 2 permit 172.16.1.11
 
internetrouter#
show access-lists
Standard IP access list 2 
    30 permit 172.16.1.11 
    20 permit 172.16.1.10 
    10 permit 172.16.1.2 
 
