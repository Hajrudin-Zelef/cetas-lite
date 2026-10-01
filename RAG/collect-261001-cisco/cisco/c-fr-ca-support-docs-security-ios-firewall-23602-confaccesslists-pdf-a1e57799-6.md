---
id: collect-261001-cisco/cisco/c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799-6
title: "c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799.md
source_anchor: ""
source_lines: [988, 1149]
sha256: 5c3c6b89447640ccd0aaad683284c8c4e6c974bb76c72cf6d1580cbec17a32ff
---

# c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799

Listes de contrôle d'accès basées sur l'heure utilisant des plages temporelles
Les listes de contrôle d'accès basées sur l'heure ont été introduites dans le logiciel Cisco IOS 
Version 12.0.1.T. Tout en étant semblables aux listes de contrôle d'accès étendues dans leur 
fonctionnement, elles autorisent un contrôle d'accès basé sur l'heure. Une plage temporelle est 
créée qui définit des heures spécifiques de la journée et de la semaine afin d'implémenter des 
listes de contrôle d'accès basées sur l'heure. La plage temporelle est identifiée par un nom et 
référencée par une fonction. Par conséquent, les restrictions horaires sont imposées à la fonction 
elle-même. La plage temporelle repose sur l'horloge système du routeur. L'horloge du routeur peut

être utilisée, mais la fonctionnalité fonctionne de façon optimale avec la synchronisation du 
Protocole d'Heure Réseau (NTP).
Voici des commandes de liste de contrôle d'accès basée sur l'heure.
 
<#root>
 
!--- Defines a named time range. 
 
time-range time-range-name 
 
!--- Defines the periodic times. 
 
periodic days-of-the-week hh:mm to [days-of-the-week] hh:mm  
 
 
!--- Or, defines the absolute times. 
 
absolute [start time date] [end time date] 
 
!--- The time range used in the actual ACL. 
 
ip access-list name|number 
 
 
        time-rangename_of_time-range 
 
Dans cet exemple, une connexion Telnet est autorisée de l'intérieur vers l'extérieur du réseau les 
lundi, mercredi et vendredi pendant les heures d'ouverture :
 
<#root>
interface Ethernet0/0 
 ip address 10.1.1.1 255.255.255.0 
 ip access-group 101 in 
! 
access-list 101 permit tcp 10.1.1.0 0.0.0.255 172.16.1.0 0.0.0.255 eq telnet time-range EVERYOTHERDAY 
! 
time-range EVERYOTHERDAY 
 periodic Monday Wednesday Friday 8:00 to 17:00
 
Entrées de liste de contrôle d'accès IP commentées
Les entrées de listes de contrôle d'accès IP commentées ont été introduites dans le logiciel 
Cisco IOS Version 12.0.2.T Les commentaires facilitent la compréhension des listes de contrôle

d'accès et peuvent être utilisés pour les listes de contrôle d'accès IP standard ou étendues.
Voici la syntaxe de commande pour les listes de contrôle d'accès nommées IP commentées.
 
<#root>
ip access-list {standard|extended} <access-list-name> 
 remark remark
 
Voici la syntaxe de commande pour les listes de contrôle d'accès IP numérotées commentées.
 
<#root>
access-list <access-list-number> remark remark
 
Voici un exemple de commentaires dans une liste de contrôle d’accès numérotée.
 
<#root>
interface Ethernet0/0 
 ip address 10.1.1.1 255.255.255.0 
 ip access-group 101 in 
! 
access-list 101 remark permit_telnet 
access-list 101 permit tcp host 10.1.1.2 host 172.16.1.1 eq telnet 
 
Contrôle d'accès basé sur contexte
Le contrôle d'accès basé sur contexte (CBAC) a été introduit dans le logiciel Cisco IOS Version 
12.0.5.T et nécessite l'ensemble des fonctionnalités du pare-feu Cisco IOS. Le contrôle CBAC 
examine le trafic qui transite par le pare-feu afin de détecter et de gérer les informations d'état 
pour les sessions TCP et UDP. Ces informations d'état sont utilisées afin de créer des ouvertures 
temporaires dans les listes d'accès du pare-feu. Configurez ip inspectlists dans la direction du flux 
d'initiation du trafic afin d'autoriser le trafic de retour et les connexions de données 
supplémentaires pour les sessions autorisées, sessions qui proviennent du réseau interne 
protégé, afin de le faire.
Voici la syntaxe pour le contrôle CBAC.
 
<#root>
ip inspect name inspection-name protocol [timeoutseconds]

Voici un exemple de l'utilisation du contrôle CBAC pour examiner le trafic sortant. La liste de 
contrôle d'accès étendue 111 bloque normalement le trafic entrant autre que le trafic ICMP sans 
les ouvertures du contrôle CBAC pour le trafic entrant.
 
<#root>
ip inspect name myfw ftp timeout 3600 
ip inspect name myfw http timeout 3600 
ip inspect name myfw tcp timeout 3600 
ip inspect name myfw udp timeout 3600 
ip inspect name myfw tftp timeout 3600 
! 
interface Ethernet0/1 
 ip address 172.16.1.2 255.255.255.0 
 ip access-group 111 in 
 ip inspect myfw out 
! 
access-list 111 deny icmp any 10.1.1.0 0.0.0.255 echo 
access-list 111 permit icmp any 10.1.1.0 0.0.0.255
 
 
Proxy d'authentification
Le proxy d'authentification a été introduit dans le logiciel Cisco IOS Version 12.0.5.T. Ce proxy 
nécessite l'ensemble de fonctionnalités du pare-feu Cisco IOS. Le proxy d'authentification est 
employé pour authentifier les utilisateurs entrants ou sortants, ou les deux. Les utilisateurs qui 
sont normalement bloqués par une liste de contrôle d'accès peuvent amener un navigateur à 
passer par le pare-feu et s'authentifier sur un serveur TACACS+ ou RADIUS. Le serveur transmet 
des entrées de la liste de contrôle d'accès supplémentaires au routeur afin de permettre aux 
utilisateurs de passer après authentification.
Le proxy d'authentification est semblable à la fonctionnalité de verrou et clé (listes de contrôle 
d'accès dynamiques). Les différences sont les suivantes :
La fonctionnalité de verrou et clé est activée par une connexion Telnet au routeur. Le proxy 
d'authentification est activé par HTTP via le routeur.
•
Le proxy d'authentification doit utiliser un serveur externe.•
Le proxy d'authentification peut gérer l'ajout de plusieurs listes dynamiques. La fonctionnalité 
de verrou et clé ne peut en ajouter qu'une.
•
Le proxy d'authentification a un délai d'attente absolu, mais pas inactif. La fonctionnalité de 
verrou et clé a les deux.
•
Pour obtenir des exemples de proxy d'authentification, reportez-vous au Manuel de configuration 
logicielle intégrée sécurisée Cisco.

Listes de contrôle d'accès turbo
Les listes de contrôle d'accès turbo ont été introduites dans le logiciel Cisco IOS Version 12.1.5.T 
et figurent uniquement sur les plates-formes 7200, 7500 et autres plates-formes haut de gamme. 
La fonctionnalité de liste de contrôle d'accès turbo est conçue pour traiter les listes de contrôle 
d'accès plus efficacement afin d'améliorer les performances du routeur.
Utilisez la commande access-list compiled pour les listes de contrôle d'accès turbo. Voici un 
exemple d'une liste de contrôle d'accès compilée.
 
<#root>
access-list 101 permit tcp host 10.1.1.2 host 172.16.1.1 eq telnet 
access-list 101 permit tcp host 10.1.1.2 host 172.16.1.1 eq ftp 
access-list 101 permit udp host 10.1.1.2 host 172.16.1.1 eq syslog 
access-list 101 permit udp host 10.1.1.2 host 172.16.1.1 eq tftp 
access-list 101 permit udp host 10.1.1.2 host 172.16.1.1 eq ntp
 
Une fois que la liste de contrôle d'accès standard ou étendue est définie, utilisez la commande 
global configuration pour la compilation.
 
<#root>
 
!--- Tells the router to compile. 
 
access-list compiled 
! 
interface Ethernet0/1 
 ip address 172.16.1.2 255.255.255.0 
 
!--- Applies to the interface. 
 
 ip access-group 101 in

