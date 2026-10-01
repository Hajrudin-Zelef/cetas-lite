---
id: collect-261001-cisco/cisco/c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799-5
title: "c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799.md
source_anchor: ""
source_lines: [744, 987]
sha256: 031219094d10223d250fb98e192cd0e21c98f8843dc9f7935ef643767a0bef6c
---

# c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799

ACL étendus
Les listes de contrôle d’accès étendues ont été introduites dans le logiciel Cisco IOS version 8.3. 
Les listes de contrôle d’accès étendues contrôlent le trafic en comparant les adresses source et 
de destination des paquets IP aux adresses configurées dans la liste de contrôle d’accès.
Voici le format de la syntaxe de commande des listes de contrôle d'accès étendues. Les lignes 
sont enveloppées ici pour des considérations d'espace.
IP
 
<#root>
access-list 
access-list-number 
     [dynamic 
dynamic-name
 [timeout 
minutes
]] 
     {deny|permit} 
protocol source source-wildcard destination destination-wildcard
 [precedence

precedence
] 
     [tos 
tos
] [log|log-input] [time-range 
time-range-name
]
 
ICMP
 
<#root>
access-list 
access-list-number 
     [dynamic 
dynamic-name
 [timeout 
minutes
]] 
     {deny|permit} icmp 
source source-wildcard destination destination-wildcard 
     [icmp-type [icmp-code] |icmp-message] 
     [precedence 
precedence
] [tos 
tos
] [log|log-input] 
     [time-range 
time-range-name
]
 
TCP
 
<#root>
access-list 
access-list-number 
     [dynamic

dynamic-name
 [timeout 
minutes
]] 
     {deny|permit} tcp 
source source-wildcard
 [operator [
port
]] 
 
destination destination-wildcard 
[operator [
port
]] 
     [established] [precedence 
precedence
] [tos 
tos
] 
     [log|log-input] [time-range 
time-range-name
]
 
UDP
 
<#root>
access-list 
access-list-number 
     [dynamic 
dynamic-name
 [timeout 
minutes
]] 
     {deny|permit} udp 
source source-wildcard 
[operator [
port

]] 
 
destination destination-wildcard
 [operator [
port
]] 
     [precedence precedence] [tos 
tos
] [log|log-input] 
     [time-range 
time-range-name
]
 
Dans toutes les versions du logiciel, le numéro de liste d’accès peut être compris entre 100 et 199. 
Dans le logiciel Cisco IOS Version 12.0.1, les listes de contrôle d’accès étendues commencent à 
utiliser des numéros supplémentaires (2000 à 2699). Ces nombres supplémentaires sont 
désignés sous le nom de listes de contrôle d'accès IP développées. Le logiciel Cisco IOS Version 
11.2 a ajouté la capacité d'utiliser le nom de liste dans les listes de contrôle d'accès étendues.
La valeur 0.0.0.0/255.255.255.255 peut être spécifiée comme any. Une fois que la liste de 
contrôle d'accès est définie, elle doit être appliquée à l'interface (entrante ou sortante). Dans les 
premières versions de logiciel, « out » était la valeur par défaut si aucun mot clé « out » ni « in » 
n'était spécifié. La direction a dû être spécifiée dans les versions de logiciel ultérieures.
 
<#root>
interface 
<interface-name> 
 ip access-group 
 
 
 
 {in|out}
 
 
Cette liste de contrôle d’accès étendue est utilisée pour autoriser le trafic sur le réseau 10.1.1.x 
(interne) et pour recevoir des réponses ping de l’extérieur, tout en empêchant les requêtes ping 
non sollicitées provenant de l’extérieur, ce qui autorise tout autre trafic.
 
<#root>
interface Ethernet0/1

ip address 172.16.1.2 255.255.255.0 
 ip access-group 101 in 
! 
access-list 101 deny icmp any 10.1.1.0 0.0.0.255 echo 
access-list 101 permit ip any 10.1.1.0 0.0.0.255
 
Remarque : Certaines applications, telles que la gestion de réseau, nécessitent des 
messages Ping pour une fonction de conservation de connexion active. Si c'est le cas, vous 
pouvez limiter les requêtes ping entrantes qui sont bloquées ou être plus granulaires dans 
les adresses IP autorisées/refusées.
Verrou et clé (listes de contrôle d'accès dynamiques)
Le verrouillage et la clé, également appelés listes de contrôle d’accès dynamiques, ont été 
introduits dans le logiciel Cisco IOS Version 11.1. Cette fonctionnalité dépend de Telnet, de 
l’authentification (locale ou distante) et des listes de contrôle d’accès étendues.
La configuration de la fonctionnalité de verrou et clé démarre avec l'application d'une liste de 
contrôle d'accès étendue pour bloquer le trafic par le routeur. Les utilisateurs qui veulent traverser 
le routeur sont bloqués par la liste de contrôle d'accès étendue jusqu'à ce qu'ils établissent une 
connexion Telnet au routeur et soient authentifiés. La connexion Telnet est alors abandonnée et 
une liste de contrôle d’accès dynamique à entrée unique est ajoutée à la liste de contrôle d’accès 
étendue existante. Le trafic est ainsi autorisé pendant une période particulière ; des délais 
d'attente inactifs et absolus sont possibles.

Remarque : Dans les versions plus récentes de Cisco IOS, le concept de « listes de 
contrôle d’accès verrouillées et à clé » a évolué ou est devenu obsolète sur de 
nombreuses plates-formes, en raison des avancées matérielles et logicielles.
Listes de contrôle d'accès nommées IP
Les listes de contrôle d’accès nommées IP ont été introduites dans le logiciel Cisco IOS Version 
11.2. Cela permet d’attribuer des noms aux listes de contrôle d’accès standard et étendues plutôt 
que des numéros.
Voici le format de la syntaxe de commande pour les listes de contrôle d'accès nommées IP.
 
<#root>
ip access-list {extended|standard} name

Voici un exemple TCP :
 
<#root>
{permit|deny} tcp source source-wildcard [operator [port]] 
destination destination-wildcard [operator [port]] [established] 
[precedence precedence] [tos tos] [log] [time-range time-range-name]
 
Voici un exemple de l'utilisation d'une liste de contrôle d'accès nommée afin de bloquer tout le 
trafic, à l'exception de la connexion Telnet entre l'hôte 10.1.1.2 et l'hôte 172.16.1.1.
 
<#root>
interface Ethernet0/0 
 ip address 10.1.1.1 255.255.255.0 
 ip access-group in_to_out in 
! 
ip access-list extended in_to_out 
 permit tcp host 10.1.1.2 host 172.16.1.1 eq telnet 
 
Listes de contrôle d'accès réflexives
Les listes de contrôle d’accès réflexives ont été introduites dans la version 11.3 du logiciel Cisco 
IOS. Elles permettent de filtrer les paquets IP en fonction des informations de session de couche 
supérieure. Elles sont généralement employées pour autoriser le trafic sortant et pour limiter le 
trafic entrant en réponse aux sessions initialisées à l'intérieur du routeur.
Les listes de contrôle d'accès réflexives peuvent être définies seulement avec les listes de 
contrôle d'accès nommées IP étendues. Elles ne peuvent pas être définies avec des listes de 
contrôle d'accès nommées IP standard ou numérotées, ni avec d'autres listes de contrôle d'accès 
de protocole. Les listes de contrôle d'accès réflexives peuvent être utilisées en même temps que 
d'autres listes de contrôle d'accès étendues standard et statiques.
Voici la syntaxe pour différentes commandes de liste de contrôle d'accès réflexive.
 
<#root>
interface 
<interface-name>
 
 ip access-group 
 
 
 
 {in|out} 
!

ip access-list extended 
<name>
 permit 
protocol
 any any reflect 
name [timeoutseconds] 
! 
ip access-list extended 
<name>
 evaluate 
<name> 
 
Ceci est un exemple de l'autorisation du trafic ICMP sortant et entrant, alors qu'il n'autorise que le 
trafic TCP qui a démarré de l'intérieur, l'autre trafic est refusé.
 
<#root>
ip reflexive-list timeout 120 
! 
interface Ethernet0/1 
 ip address 172.16.1.2 255.255.255.0 
 ip access-group inboundfilters in 
 ip access-group outboundfilters out 
! 
ip access-list extended inboundfilters 
 permit icmp 172.16.1.0 0.0.0.255 10.1.1.0 0.0.0.255 
 evaluate tcptraffic 
 
!--- This ties the reflexive ACL part of the outboundfilters ACL,  
!--- called tcptraffic, to the inboundfilters ACL. 
 
ip access-list extended outboundfilters 
 permit icmp 10.1.1.0 0.0.0.255 172.16.1.0 0.0.0.255 
 permit tcp 10.1.1.0 0.0.0.255 172.16.1.0 0.0.0.255 reflect tcptraffic
 
