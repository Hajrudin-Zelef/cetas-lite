---
id: collect-261001-general-networking/general-networking/c-fr-ca-support-docs-ip-network-address-translation-nat-13772-12-pdf-5677aabb-2
title: "c-fr-ca-support-docs-ip-network-address-translation-nat-13772-12-pdf-5677aabb"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-general-networking/c-fr-ca-support-docs-ip-network-address-translation-nat-13772-12-pdf-5677aabb.md
source_anchor: ""
source_lines: [157, 310]
sha256: 9e3dc38de2f963e35e7464ccf85bf91bbe2d221fbf3948889a52d59af0c00e74
---

# c-fr-ca-support-docs-ip-network-address-translation-nat-13772-12-pdf-5677aabb

Remarque : Cisco recommande fortement de ne pas configurer les listes d'accès 
référencées par les commandes NAT avec permit any. Si vous utilisez permit any dans NAT, 
il consomme trop de ressources de routeur qui peuvent causer des problèmes réseau.
Dans la configuration précédente, notez que seules les 32 premières adresses du sous-réseau 
10.10.10.0 et les 32 premières adresses du sous-réseau 10.10.20.0 sont autorisées par la liste 
d’accès 7. Par conséquent, seules ces adresses source sont traduites. Il peut y avoir d’autres 
appareils avec d’autres adresses sur le réseau interne, mais ceux-ci ne sont pas traduits.
Configurer la NAT pour permettre aux utilisateurs internes d’accéder à Internet en cas de 
surcharge
Routeur NAT
 
interface ethernet 0 
 ip address 10.10.10.1 255.255.255.0 
 ip nat inside 
 
!--- Defines Ethernet 0 with an IP address and as a NAT inside interface. 
 
 
interface ethernet 1 
 ip address 10.10.20.1 255.255.255.0 
 ip nat inside 
 
!--- Defines Ethernet 1 with an IP address and as a NAT inside interface. 
 
 
interface serial 0 
 ip address 172.16.10.64 255.255.255.0 
 ip nat outside 
 
!--- Defines serial 0 with an IP address and as a NAT outside interface. 
 
 
ip nat pool ovrld 172.16.10.1 172.16.10.1 prefix 24 
 
 
!--- Defines a NAT pool named ovrld with a range of a single IP  
!--- address, 172.16.10.1. 
 
 
ip nat inside source list 7 pool ovrld overload 
 
 
!--- Indicates that any packets received on the inside interface that  
!--- are permitted by access-list 7 has the source address  
!--- translated to an address out of the NAT pool named ovrld.  
!--- Translations are overloaded, which allows multiple inside  
!--- devices to be translated to the same valid IP address. 
 
access-list 7 permit 10.10.10.0 0.0.0.31 
access-list 7 permit 10.10.20.0 0.0.0.31 
 
!--- Access-list 7 permits packets with source addresses ranging from  
!--- 10.10.10.0 through 10.10.10.31 and 10.10.20.0 through 10.10.20.31.

Remarquez que dans la deuxième configuration précédente, le pool NAT n’a qu’une seule 
adresse. ovrld Le mot clé overload utilisé dans la commande ip nat inside source list 7 pool ovrld 
overload permet à NAT de traduire plusieurs périphériques internes en l'adresse unique dans le 
groupe.
Une autre variante de cette commande estip nat inside source list 7 interface serial 0 overload, qui 
configure la NAT pour la surcharge sur l'adresse assignée à l'interface Serial 0.
Lorsque  est configuré, le routeur conserve suffisamment d’informations des protocoles de niveau 
supérieur (par exemple, les numéros de port TCP ou UDP) pour traduire l’adresse globale en 
adresse locale correcte.overloading Pour les définitions d'adresse globale et locale, référez-vous à 
NAT : Définitions locales et globales.
2. Autoriser Internet à accéder aux périphériques internes
Vous pouvez avoir besoin de périphériques internes pour échanger des informations avec des 
périphériques sur Internet, où la communication est initiée à partir d’appareils Internet, par 
exemple, pour échanger des courriels. Il est courant que les périphériques Internet envoient des 
courriels à un serveur de messagerie qui réside sur le réseau interne.
Origine des communications
Configurer la NAT pour permettre à Internet d’accéder aux périphériques internes
Dans cet exemple, vous définissez d’abord les interfaces interne et externe de la NAT, comme 
illustré dans le diagramme de réseau précédent.
Ensuite, vous définissez que vous souhaitez que les utilisateurs internes puissent établir des

communications avec l’extérieur. Les périphériques externes doivent pouvoir communiquer avec 
le serveur de messagerie uniquement à l’intérieur.
La troisième étape est de configurer NAT. Pour réaliser ce que vous avez défini, vous pouvez 
configurer ensemble la NAT statique et dynamique. Pour plus d'informations sur la façon de 
configurer cet exemple, référez-vous à Configurer la NAT statique et dynamique simultanément.
3. Rediriger le trafic TCP vers un autre port TCP ou une autre adresse
Un serveur Web sur le réseau interne est un autre exemple où il peut être nécessaire que des 
périphériques Internet établissent la communication avec des périphériques internes. Dans 
certains cas, le serveur Web interne peut être configuré pour écouter le trafic Web sur un port 
TCP autre que le port 80. Par exemple, le serveur Web interne peut être configuré pour écouter le 
port TCP 8080. Dans ce cas, vous pouvez utiliser NAT pour rediriger le trafic destiné au port TCP 
80 au port TCP 8080.
Port TCP du trafic Web
Après avoir défini les interfaces comme indiqué dans le diagramme de réseau précédent, vous 
pouvez décider que vous souhaitez que la NAT redirige les paquets de l’extérieur destinés à 
172.16.10.8:80 vers 172.16.10.8:8080. Vous pouvez utiliser une commande static nat pour 
traduire le numéro de port TCP pour y parvenir. Un exemple de configuration est présenté ici.
Configurer la NAT pour rediriger le trafic TCP vers un autre port TCP ou une autre adresse
Routeur NAT
 
interface ethernet 0 
 ip address 172.16.10.1 255.255.255.0 
 ip nat inside 
 
!--- Defines Ethernet 0 with an IP address and as a NAT inside interface.

interface serial 0 
 ip address 10.200.200.5 255.255.255.252 
 ip nat outside 
 
!--- Defines serial 0 with an IP address and as a NAT outside interface. 
 
 
ip nat inside source static tcp 172.16.10.8 8080 172.16.10.8 80 
 
!--- Static NAT command that states any packet received in the inside  
!--- interface with a source IP address of 172.16.10.8:8080 is  
!--- translated to 172.16.10.8:80. 
 
 
Remarque : La description de la configuration pour la commande NAT statique indique que 
tout paquet reçu dans l’interface interne avec une adresse source de 172.16.10.8:8080 est 
traduit en 172.16.10.8:80. Cela implique également que tout paquet reçu sur l’interface 
externe avec une adresse de destination de 172.16.10.8:80 a la destination traduite en 
172.16.10.8:8080.
 
show ip nat translations 
Pro Inside global      Inside local       Outside local      Outside global 
tcp 172.16.10.8:80     172.16.10.8:8080   ---                ---
 
4. Utiliser la NAT pour une transition réseau
La NAT est utile lorsque vous devez réadresser des périphériques sur le réseau ou lorsque vous 
remplacez un périphérique par un autre. Par exemple, si tous les périphériques du réseau utilisent 
un serveur particulier et que ce serveur doit être remplacé par un nouveau serveur qui a une 
nouvelle adresse IP, la reconfiguration de tous les périphériques du réseau pour utiliser la 
nouvelle adresse de serveur prendra un certain temps. En attendant, vous pouvez utiliser la NAT 
pour configurer les périphériques avec l’ancienne adresse pour traduire leurs paquets afin de 
communiquer avec le nouveau serveur.

Transition du réseau NAT
Une fois que vous avez défini les interfaces NAT, comme illustré à l’image précédente, vous 
pouvez décider que vous souhaitez que la NAT permette que les paquets de l’extérieur destinés à 
l’ancienne adresse du serveur (172.16.10.8) soient traduits et envoyés à la nouvelle adresse du 
serveur. N'oubliez pas que le nouveau serveur se trouve sur un autre réseau local et que les 
périphériques de ce réseau local ou les périphériques accessibles via ce réseau local 
(périphériques situés à l'intérieur du réseau) doivent être configurés pour utiliser la nouvelle 
adresse IP du serveur, si possible.
Vous pouvez utiliser NAT statique pour effectuer ce dont vous avez besoin. Voici un exemple de 
configuration.
Configurer la NAT pour une utilisation pendant une transition réseau
Routeur NAT
 
interface ethernet 0 
 ip address 172.16.10.1 255.255.255.0 
 ip nat outside 
 
!--- Defines Ethernet 0 with an IP address and as a NAT outside interface. 
 
 
