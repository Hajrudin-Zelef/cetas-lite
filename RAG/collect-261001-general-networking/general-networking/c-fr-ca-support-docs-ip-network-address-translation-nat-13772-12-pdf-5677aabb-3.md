---
id: collect-261001-general-networking/general-networking/c-fr-ca-support-docs-ip-network-address-translation-nat-13772-12-pdf-5677aabb-3
title: "c-fr-ca-support-docs-ip-network-address-translation-nat-13772-12-pdf-5677aabb"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-general-networking/c-fr-ca-support-docs-ip-network-address-translation-nat-13772-12-pdf-5677aabb.md
source_anchor: ""
source_lines: [311, 438]
sha256: 301f129ff98b46bb389fd1629feba16eedc503afd5d14829d7ab99f9b0598fa2
---

# c-fr-ca-support-docs-ip-network-address-translation-nat-13772-12-pdf-5677aabb

interface ethernet 1 
 ip address 172.16.50.1 255.255.255.0 
 ip nat inside

!--- Defines Ethernet 1 with an IP address and as a NAT inside interface. 
 
 
interface serial 0 
 ip address 10.200.200.5 255.255.255.252 
 
!--- Defines serial 0 with an IP address. This interface is not  
!--- participating in NAT. 
 
 
ip nat inside source static 172.16.50.8 172.16.10.8 
 
!--- States that any packet received on the inside interface with a  
!--- source IP address of 172.16.50.8 is translated to 172.16.10.8. 
 
 
Remarque : La commande NAT source interne dans cet exemple implique également que 
les paquets reçus sur l'interface externe avec une adresse de destination de 172.16.10.8 ont 
l'adresse de destination traduite en 172.16.50.8.
5. Utilisez la NAT pour les réseaux qui se chevauchent
Des réseaux peuvent en venir à se chevaucher lorsque vous attribuez des adresses IP à des 
appareils internes qui sont déjà utilisées par d’autres appareils sur Internet. Cela peut également 
se produire lorsque deux entreprises, qui utilisent toutes deux les adresses IP RFC 1918 dans 
leurs réseaux, fusionnent. Ces deux réseaux doivent communiquer, de préférence sans que tous 
leurs périphériques soient réadressés.
Différence entre le mappage un-à-un et plusieurs à plusieurs
A configuration de NAT statique crée un mappage linéaire et traduit une adresse spécifique en 
une autre adresse. Ce type de configuration crée une entrée permanente dans la table NAT tant 
que la configuration est présente et permet aussi bien à des hôtes internes qu'externes d'initier 
une connexion. Ceci est en grande partie utile pour les hôtes qui fournissent des services 
d'application comme le courrier, Web, FTP et ainsi de suite. Exemple :
 
<#root>
Router(config)#
ip nat inside source static 10.3.2.11 10.41.10.12
Router(config)#
ip nat inside source static 10.3.2.12 10.41.10.13
 
NAT dynamique est utile quand moins d'adresses sont disponibles que le nombre réel d'hôtes à

traduire. Il crée une entrée dans la table NAT quand l'hôte lance une connexion et établit un 
mappage linéaire entre les adresses. Cependant, le mappage peut varier et dépend des adresses 
enregistrées disponibles dans le pool au moment de la transmission. NAT dynamique permet à 
des sessions d'être lancées seulement de l'intérieur ou de l'extérieur de réseaux externes pour 
lesquels elle est configurée. Des entrées de NAT dynamique sont supprimées de la table de 
traduction si l'hôte ne communique pas pendant une période spécifique qui est configurable. 
L'adresse est alors retournée au pool à l'usage d'un autre hôte.
Par exemple, complétez ces étapes de la configuration détaillée :
Créez un groupe d'adresses.
 
<#root>
Router(config)#
ip nat pool MYPOOLEXAMPLE 10.41.10.1 10.41.10.41 netmask 255.255.255.0
 
1. 
Créez une liste d'accès pour les réseaux internes qui doivent être mappés.
 
<#root>
Router(config)#
access-list 100 permit ip 10.3.2.0 0.0.0.255 any
 
2. 
Associez la liste d’accès 100 qui sélectionne le réseau interne 10.3.2.0 0.0.0.255 pour le 
déploiement de la fonction NAT pour le pool MYPOOLEXAMPLE, puis surchargez les 
adresses.
 
<#root>
Router(config)#
ip nat inside source list 100 pool MYPOOLEXAMPLE overload
 
3. 
Vérifiez l'opération de NAT
Une fois que vous avez configuré la NAT, vérifiez qu’elle fonctionne comme prévu. Vous pouvez 
faire ceci de diverses façons : avec un analyseur de réseau, des commandes show ou des 
commandes debug. Pour obtenir un exemple détaillé de vérification NAT, reportez-vous 
àVérification du fonctionnement NAT et NAT de base.

Conclusion
Les exemples de ce document montrent les étapes de démarrage rapide qui peuvent vous aider à 
configurer et à déployer la NAT.
Ces étapes de démarrage rapide comprennent :
Définissez les interfaces internes et externes de NAT.1. 
Que voulez-vous réaliser avec NAT?2. 
Configurez la NAT afin d’accomplir ce que vous avez défini à l’étape 2.3. 
Vérifiez l'opération de NAT .4. 
Dans chacun des exemples précédents, diverses formes de la commande ip nat Insideont été 
utilisées. Vous pouvez également utiliser la commande ip nat Outsidepour réaliser les mêmes 
objectifs, mais gardez à l’esprit l’ordre des opérations de la NAT. Pour des exemples de 
configuration qui utilisent les commandes ip nat outsidecommandes, référez-vous à Configurer la 
commande IP NAT Outside Source List.
Les exemples précédents ont également montré ces actions :
Commande Action
ip nat inside source
Traduit la source des paquets IP se déplaçant de l'intérieur vers 
l'extérieur.
•
Traduit la destination des paquets IP se déplaçant de l'extérieur vers 
l'intérieur.
•
ip nat outside 
source
Traduit la source des paquets IP se déplaçant de l'extérieur vers 
l'intérieur.
•
Traduit la destination des paquets IP se déplaçant de l'intérieur vers 
l'extérieur.
•
Informations connexes
NAT : Définitions locales et globales.•
Page de support NAT•
Page d'assistance pour les protocoles de routage IP•
Page de support pour le routage IP•
Services d’adressage IP•
Ordre des opérations NAT•
Questions fréquentes au sujet de Cisco IOS NAT•
Assistance et documentation techniques - Cisco Systems•

À propos de cette traduction
Cisco a traduit ce document en traduction automatisée vérifiée par une personne dans le
cadre d’un service mondial permettant à nos utilisateurs d’obtenir le contenu d’assistance
dans leur propre langue.
 
Il convient cependant de noter que même la meilleure traduction automatisée ne sera pas
aussi précise que celle fournie par un traducteur professionnel.
