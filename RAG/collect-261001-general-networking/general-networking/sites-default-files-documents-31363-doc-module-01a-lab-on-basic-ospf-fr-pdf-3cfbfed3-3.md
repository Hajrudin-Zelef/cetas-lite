---
id: collect-261001-general-networking/general-networking/sites-default-files-documents-31363-doc-module-01a-lab-on-basic-ospf-fr-pdf-3cfbfed3-3
title: "sites-default-files-documents-31363-doc-module-01a-lab-on-basic-ospf-fr-pdf-3cfbfed3"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2016-08-12"]
keywords: ["arr", "ethernet"]
source: docs/RAG/collect-261001-general-networking/sites-default-files-documents-31363-doc-module-01a-lab-on-basic-ospf-fr-pdf-3cfbfed3.md
source_anchor: ""
source_lines: [278, 415]
sha256: de4bc7efc3e7b080f3b6678a60c50e954fb9419d999b6b4253d76cc60618a7d8
---

# sites-default-files-documents-31363-doc-module-01a-lab-on-basic-ospf-fr-pdf-3cfbfed3

REMARQUE: Les instructeurs de laboratoire auront dessiné une grande cart e réseau sur le 
tableau blanc dans le laboratoire de l'atelier. Lorsque les adr esses IP sont attribuées, prière de les 
annoter et d'informer l'instructeur. Tous les liens point à poi nt DOIVENT être annotées  là pour 
que les équipes d'autres routeurs peuvent documenter et compren dre les liens et routage dans les 
modules actuelles et futures. 
 
Q : Quel masque de réseau doit être utilisé sur le lien point-à-point? 
 
A: Sur les interfaces en série, le masque de réseau devrait être / 30  (ou 255.255.255.252 en format 
dotted quad). Il est inutile d'utiliser une autre taille de mas que car il y’a seulement deux hôtes sur 
un tel lien. Une adresse masque 255.255.255.252 signifie 4 adre sses hôte disponibles, dont deux 
sont utilisables (les deux autres représentant les adresses réseau et broadcast). 
 
12. Connexions Ethernet Les liens Ethernet entre les routeurs seront effectués en util isant des câbles 
RJ-45 cross-over -. Ceux-ci relieront directement les ports Ethernet sur les de ux routeurs sans le 
besoin d'un switch Ethernet. Les subnets IP seront de nouveau t iré du plan d'adressage. Ne faites 
pas l'erreur d'attribuer un mas que / 24 à l'adresse de l'interf ace - il y’a seulement deux hôtes sur le 
réseau Ethernet reliant les deux routeurs, donc un masque / 30 doit être tout à fait suffisant. 
 
13. Ping Test n ° 1. Ping tous les subnets connectés physiquement des routeurs vois ins. Si les subnets  
connectés physiquement sont inaccessibles, consulter vos équipe s voisins pour trouver  le 
problème. Ne pas ignorer le problème – ça peut persister. Utili sez les commandes suivantes pour 
dépanner la connexion: 
 
show arp     : Indique le protocole de résolution d'adresse 
show interface <interface> <number> : État de l'interface et la configuration 
show ip interface     : Résumé bref de l'état des  interfaces IP et la configuration 
 
14. Creation des Loopback Interfaces.  Les Interfaces Loopback seront utilisé dans cet atelier pour 
beaucoup de choses. Ceux-ci comprennent la production de route (à être annoncé) et la 
configuration des peerings BGP. Comme indiqué précédemment dans  l'étape 10, nous allons 
utiliser une partie du bloc d'a dresses IP allouées pour les int erfaces de loopback. La plupart des 
ISP ont tendance à mettre de côté un bloc contigu d'adresses po ur l’utilisation par leurs routeur 
loopbacks. Par exemple, si un ISP a eu 20 routeurs, ils auraien t besoin d'un / 27 (ou 32 adresses 
hôte) afin de fournir une adresse loopback pour chaque routeur. Nous avons 14 routeurs dans notre 
laboratoire – pour faire preuve de prudence et permettre la cro issance, nous allons mettre de côté 
un / 27 (nous permet 32 loopbacks) mais en utiliser seulement 1 4 d'entre eux. Les adresses 
loopbacks assignées sont les suivantes: 
 
R1 10.0.15.241/32 
R2 10.0.15.242/32 
R3 10.0.15.243/32 
R4 10.0.15.244/32 
R5 10.0.15.245/32 
R6 10.0.15.246/32

Friday, August 12, 2016   
 8 
 
R7 10.0.15.247/32 
R8 10.0.15.248/32 
R9 10.0.15.249/32 
R10 10.0.15.250/32 
R11 10.0.15.251/32 
R12 10.0.15.252/32 
R13 10.0.15.253/32 
R14 10.0.15.254/32 
 
Par exemple, l’équipe routeur 1 attribuera l’adresse et le masque suivant au loopback sur le routeur 
1: 
 
Router1(config)#interface loopback 0 
Router1(config-if)#ip address 10.0.15.241 255.255.255.255 
 
      Q: Pourquoi utilisons-nous des masques / 32 pour l'adresse d'interface loopback? 
 
A: Il n'y’a pas de  réseau physique attachée au loopback, de sorte qu'il ne peut y  a v o i r  q u ' u n  
dispositif. Donc, nous avons seulement besoin d'affecter un mas que / 32 - c'est un gaspillage 
d'espace d'adressage d'utiliser autre chose. 
 
 
Checkpoint # 1:  appelez l’assistant de laboratoire pour vér ifier la connectivité. Montrez que vous 
pouvez faire un ping et Telnet aux routeurs adjacents. 
 
 
15. OSPF avec une zone dans le même AS - activer le processus OSPF  Chaque équipe routeur doit 
activer le protocole OSPF sur le routeur. L'identificateur de p rocessus OSPF doit être 41 ( v o i r  
exemple). (L'identificateur de processus OSPF est juste un nomb re pour identifier ce processus 
OSPF sur le routeur. Il n'est pas transmis entre les routeurs.)  
 
Router1(config)#router ospf 41 
 
La configuration IOS par défaut doit être modifiée de sorte que toutes les interfaces soient marqués 
comme passif pour OSPF par défaut. Cela supprime les mises à jo ur de routage  sur  toutes les 
interfaces du routeur et arrête le routeur de former involontai rement les contiguïtés OSPF sur les 
interfaces externes, et évite les problèmes potentiels que cela peut apporter2. 
 
Router1(config-router)#passive-interface default 
 
Toutes les interfaces sur lesquelles les contiguïtés OSPF doivent être formés doivent être marquées 
avec la sous-commande no passive-interface. 
 
Router1(config-router)#no passive-interface fastethernet 0/0 
Router1(config-router)#no passive-interface fastethernet 0/1 
Router1(config-router)#no passive-interface serial 1/0 
 
16. Activation du protocole OSPF sur chaque interface.  Maintenant que le processus OSPF est 
configuré, chaque équipe doit activer le protocole OSPF sur les  interfaces du chaque routeur selon 
les besoins. Contrairement aux versions précédentes de l'IOS, I OS 12.4 et versions ultérieures 
permettent également à l’OSPF d’être exécuté sur un lien (plutô t que sur un sub-net). Plutôt que 
                                                            
2 C'est une erreur courante dans de nombreuses configurations ISP d’avoir l'IGP actif sur toutes les interfaces du routeur. Il 
ya eu de nombreux accidents documentés où un client IGP a établi une connexion avec l’IGP du ISP, ce qui entraîne une 
pollution croisée des informations de routage, et le chaos de trafic qui en résulte. Le fait de désactiver cette fonctionnalité 
en marquant toutes les interfaces passives par défaut permet d'éviter les oublis ou les erreurs dans le futur.

Atelier Labo ISP 
  9   
  
d'utiliser la déclaration «network» la plus ancienne,  nous act ivons maintenant le protocole OSPF 
sur chaque interface qui va former une contiguïté: 
 
Router1(config)#interface serial 1/0 
Router1(config-if)#ip ospf 41 area 0 
! 
Router1(config-if)#interface fastethernet 0/0 
Router1(config-if)#ip ospf 41 area 0 
! 
Router1(config-if)#interface fastethernet 0/1 
Router1(config-if)#ip ospf 41 area 0 
. 
17. Annoncer la loopback en / 32.  L'interface de loopback nécessite également que OSPF lui soit 
a c t i v é .  M ê m e  s ' i l  n ' y  a  p a s  d e  c o n t i g u ï t é  q u i  d o i t  ê t r e  f o r m é  (parce qu'il n'y a pas de voisin 
physique et l'interface est mar quée comme passif par défaut à l 'étape précédente), nous devons 
déclarer OSPF sur l'interface loopback afin que l'adresse IP utilisée pour le loopback est placé dans 
le RIB OSPF. 
 
Router1(config)#interface loopback 0 
Router1(config-if)#ip ospf 41 area 0 
  
18. Contiguïtés OSPF. Chaque équipe doit activer la logging des changements de contig uïté OSPF. 
(Remarque: A de partir IOS 12.4, log-voisin-changes est activé par défa ut lorsque OSPF est 
initialement configuré). Il en est ainsi pour qu’une notification soit générée à chaque fois que l'état 
d'un voisins OSPF change, et est utile pour le débogage: 
 
Router2(config)#router ospf 41 
Router2(config-router)#log-adjacency-changes 
 
