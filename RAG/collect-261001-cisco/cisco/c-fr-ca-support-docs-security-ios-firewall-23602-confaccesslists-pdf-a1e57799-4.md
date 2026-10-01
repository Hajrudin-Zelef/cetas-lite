---
id: collect-261001-cisco/cisco/c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799-4
title: "c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "attention"]
source: docs/RAG/collect-261001-cisco/c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799.md
source_anchor: ""
source_lines: [552, 743]
sha256: 409fd57c872e536b890f0232c13df2acc012a968e7a241d4ab1dbadce297ebbb
---

# c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799

internetrouter(config)#
ip access-list standard 2

internetrouter(config-std-nacl)#
25 per 172.16.1.7
internetrouter(config-std-nacl)#
15 per 172.16.1.16
 
internetrouter#
show access-lists
Standard IP access list 2 
 
15 permit 172.16.1.16
    30 permit 172.16.1.11 
    20 permit 172.16.1.10 
 
25 permit 172.16.1.7
    10 permit 172.16.1.2
 
La principale différence dans une liste d’accès standard est que Cisco IOS ajoute une entrée dans 
l’ordre descendant de l’adresse IP, et non sur un numéro d’ordre.
Cet exemple illustre les différentes entrées, par exemple comment autoriser une adresse IP 
(192.168.100.0) ou les réseaux (10.10.10.0).
 
<#root>
internetrouter#
show access-lists
Standard IP access list 19 
    10 permit 192.168.100.0 
    15 permit 10.10.10.0, wildcard bits 0.0.0.255 
    19 permit 10.101.110.0, wildcard bits 0.0.0.255 
    25 deny any
 
Ajoutez l'entrée dans la liste d'accès 2 afin d'autoriser l'adresse IP 172.22.1.1 :
 
<#root>
internetrouter(config)#
ip access-list standard 2
internetrouter(config-std-nacl)#
18 permit 172.22.1.1

Cette entrée est ajoutée en haut de la liste afin de donner la priorité à l'adresse IP spécifique 
plutôt qu'au réseau.
 
<#root>
internetrouter#
show access-lists
Standard IP access list 19 
    10 permit 192.168.100.0 
 
18 permit 172.22.1.1
    15 permit 10.10.10.0, wildcard bits 0.0.0.255 
    19 permit 10.101.110.0, wildcard bits 0.0.0.255 
    25 deny   any
 
Remarque : Les listes de contrôle d'accès précédentes ne sont pas prises en charge dans 
un dispositif de sécurité, tel que le pare-feu ASA/PIX.
Instructions pour modifier les listes de contrôle d'accès lorsqu'elles sont appliquées à des crypto-
cartes.
Si vous ajoutez à une configuration de liste d'accès actuelle, il n'est pas nécessaire de 
supprimer la crypto-carte. En cas d'ajout direct sans suppression de la carte de chiffrement, 
cette opération est prise en charge et acceptable.
•
Si vous devez modifier ou supprimer une entrée de liste d'accès d'une liste d'accès actuelle, 
vous devez supprimer la carte de chiffrement de l'interface. Après avoir supprimé la carte de 
chiffrement, apportez toutes les modifications nécessaires à la liste d'accès et ajoutez à 
nouveau la carte de chiffrement. Si vous apportez des modifications telles que la 
suppression de la liste d'accès sans suppression de la carte de chiffrement, cette opération 
n'est pas prise en charge et peut avoir comme conséquence un comportement imprévisible.
•
Dépannage
Comment supprimer une liste de contrôle d'accès d'une interface ?
Passez en mode de configuration et entrez no devant la commande access-group, comme illustré 
dans cet exemple, afin de supprimer une liste de contrôle d'accès d'une interface.
 
<#root>
interface <interface-name> 
 no ip access-group <acl> {in|out}

Que faire quand trop de trafic est refusé ?
Si trop de trafic est refusé, étudiez la logique de votre liste ou essayez de définir et d'appliquer une 
autre liste plus importante. La commande show ip access-lists fournit un nombre de paquets qui 
indique l'entrée de la liste de contrôle d'accès consultée. Le mot clé log à la fin des entrées 
individuelles de la liste de contrôle d'accès indique le numéro de la liste de contrôle d'accès et si le 
paquet a été autorisé ou refusé, en plus des informations spécifiques au port.
Remarque : Le mot-clé log-input existe dans le logiciel Cisco IOS version 11.2 et ultérieure. 
L’utilisation de ce mot clé consigne des informations supplémentaires sur le trafic 
correspondant à la liste de contrôle d’accès, telles que l’interface d’entrée et l’adresse MAC 
source, le cas échéant.
Comment déboguer au niveau du paquet qui utilise un routeur Cisco ?
Cette procédure explique le processus de débogage. Avant de commencer, assurez-vous 
qu'aucune liste de contrôle d'accès n'est actuellement appliquée, qu'il existe une liste de contrôle 
d'accès et que la commutation rapide n'est pas désactivée.
Remarque : Faites très attention quand vous déboguez un système avec un trafic intense. 
Utilisez une liste de contrôle d'accès afin de déboguer un trafic spécifique. Assurez-vous du 
processus et du flux de trafic.
Utilisez la commande access-list afin de capturer les données souhaitées.
Dans cet exemple, la capture de données est définie pour l'adresse de destination 10.2.6.6 
ou l'adresse source 10.2.6.6.
 
<#root>
access-list 101 permit ip any host 10.2.6.6 
access-list 101 permit ip host 10.2.6.6 any
 
1. 
Désactivez la commutation rapide sur les interfaces impliquées. Vous voyez seulement le 
premier paquet si la commutation rapide n'est pas désactivée.
 
<#root>
configure terminal  
 interface 
 
 
          
 
 
 
  no ip route-cache 
2.

Utilisez la commande terminal monitor en mode enable afin d'afficher la sortie de la 
commande debug et les messages d'erreur système pour la session et le terminal actuels.
3. 
Utilisez la commande debug ip packet 101 ou debug ip packet 101 detail afin de commencer 
le processus de débogage.
4. 
Exécutez la commande no debug all en mode enable et la commande interface configuration 
afin d'arrêter le processus de débogage.
5. 
Redémarrez la mise en cache.
 
<#root>
configure terminal  
 interface 
 
 
          
 
 
 
  ip route-cache 
 
 
 
 
6. 
Types de listes de contrôle d'accès IP
Cette section du document décrit les types de listes de contrôle d'accès.
Diagramme du réseau

ACL standards
Les listes de contrôle d'accès standard sont le plus ancien type de liste de contrôle d'accès. Elles 
remontent à la version 8.3 du logiciel Cisco IOS. Les listes de contrôle d’accès standard contrôlent 
le trafic en comparant l’adresse source des paquets IP aux adresses configurées dans la liste de 
contrôle d’accès.
Voici le format de la syntaxe de commande d'une liste de contrôle d'accès standard.
 
<#root>
access-list <access-list-number> {permit|deny} {host|source source-wildcard|any}
 
Dans toutes les versions du logiciel, le numéro de liste d’accès peut être compris entre 1 et 99. 
Dans le logiciel Cisco IOS Version 12.0.1, les listes de contrôle d’accès standard commencent à 
utiliser des numéros supplémentaires (1300 à 1999). Ces nombres supplémentaires sont 
désignés sous le nom de listes de contrôle d'accès IP développées. Le logiciel Cisco IOS Version 
11.2 a ajouté la capacité d'utiliser le nom de liste dans les listes de contrôle d'accès standard.
Un paramètre source/source-wildcard de 0.0.0.0/255.255.255.255 peut être spécifié comme any. 
Le générique peut être omis s'il n'est composé que de zéros. Par conséquent, l’hôte 10.1.1.2 
0.0.0.0 est identique à l’hôte 10.1.1.2.
Une fois que la liste de contrôle d'accès est définie, elle doit être appliquée à l'interface (entrante 
ou sortante). Dans les premières versions de logiciel, « out » était la valeur par défaut si aucun 
mot clé « out » ni « in » n'était spécifié. La direction a dû être spécifiée dans les versions de 
logiciel ultérieures.
 
<#root>
interface <interface-name>
 ip access-group

{in|out} 
 
 
 
 
Voici un exemple de l'utilisation d'une liste de contrôle d'accès standard afin de bloquer tout le 
trafic, à l'exception de celui provenant de la source 10.1.1.x.
 
<#root>
interface Ethernet0/0 
 ip address 10.1.1.1 255.255.255.0 
 ip access-group 1 in 
! 
access-list 1 permit 10.1.1.0 0.0.0.255 
 
