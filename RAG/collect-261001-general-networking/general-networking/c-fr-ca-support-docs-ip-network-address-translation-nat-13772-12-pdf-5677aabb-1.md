---
id: collect-261001-general-networking/general-networking/c-fr-ca-support-docs-ip-network-address-translation-nat-13772-12-pdf-5677aabb-1
title: "c-fr-ca-support-docs-ip-network-address-translation-nat-13772-12-pdf-5677aabb"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-general-networking/c-fr-ca-support-docs-ip-network-address-translation-nat-13772-12-pdf-5677aabb.md
source_anchor: ""
source_lines: [1, 156]
sha256: 0de87693f5482a34aedcf5eb9cfecb82fb514efbf3d83ca25cd7e3363993dd6d
---

# c-fr-ca-support-docs-ip-network-address-translation-nat-13772-12-pdf-5677aabb

Configurer la traduction d’adresse réseau
Table des matières
Introduction
Conditions préalables
Exigences
Composants utilisés
Conventions
Étapes de démarrage rapide pour configurer et déployer la NAT
Définir la NAT sur les interfaces intérieure et extérieure
Exemples
1. Autoriser les utilisateurs internes à accéder à Internet
Configurer la NAT pour permettre aux utilisateurs internes d’accéder à Internet
Configurer la NAT pour permettre aux utilisateurs internes d’accéder à Internet en cas de 
surcharge
2. Autoriser Internet à accéder aux périphériques internes
Configurer la NAT pour permettre à Internet d’accéder aux périphériques internes
3. Rediriger le trafic TCP vers un autre port TCP ou une autre adresse
Configurer la NAT pour rediriger le trafic TCP vers un autre port TCP ou une autre adresse
4. Utiliser la NAT pour une transition réseau
Configurer la NAT pour une utilisation pendant une transition réseau
5. Utilisez la NAT pour les réseaux qui se chevauchent
Différence entre les mappages un-à-un et plusieurs-à-plusieurs
Vérifiez l'opération de NAT
Conclusion
Informations connexes
Introduction
Ce document décrit comment configurer la traduction d’adresses réseau (NAT) sur un routeur 
Cisco.
Conditions préalables
Exigences
Ce document exige une connaissance de base des termes utilisés en relation avec NAT. 
Composants utilisés
Les informations contenues dans ce document sont basées sur les versions de matériel et de 
logiciel suivantes :

Routeurs de la gamme Cisco 2500•
Cisco IOS® Version du logiciel 12.2(10b)•
The information in this document was created from the devices in a specific lab environment. All of 
the devices used in this document started with a cleared (default) configuration. Si votre réseau 
est en ligne, assurez-vous de bien comprendre l’incidence possible des commandes.
Conventions
Pour plus d'informations sur les conventions utilisées dans ce document, reportez-vous à 
Conventions relatives aux conseils techniques Cisco.
Étapes de démarrage rapide pour configurer et déployer la NAT
Remarque : Dans ce document, quand Internet ou un périphérique d'Internet est mentionné, 
cela renvoie à un périphérique sur n'importe quel réseau externe.
Quand vous configurez NAT, il est parfois difficile de savoir où commencer, particulièrement si 
vous débutez avec NAT. Ces étapes vous guident pour définir ce que vous voulez que NAT fasse 
et comment le configurer :
Définissez les interfaces internes et externes de NAT.
Les utilisateurs existent-ils outre plusieurs interfaces ?•
Existe-t-il plusieurs interfaces pour Internet?•
1. 
Définissez ce que vous souhaitez réaliser avec la NAT.
Voulez-vous autoriser les utilisateurs internes à accéder à Internet ?•
Voulez-vous autoriser Internet à accéder aux périphériques internes (comme un 
serveur de messagerie ou un serveur Web)?
•
Voulez-vous rediriger le trafic TCP vers un autre port TCP ou une autre adresse ?•
Voulez-vous utiliser la fonction NAT pendant une transition réseau (par exemple, vous 
avez modifié l'adresse IP d'un serveur et, jusqu'à ce que vous puissiez mettre à jour 
tous les clients, vous voulez que les clients non mis à jour puissent accéder au serveur 
avec l'adresse IP d'origine et autoriser les clients mis à jour à accéder au serveur avec 
la nouvelle adresse) ?
•
Voulez-vous autoriser les réseaux qui se chevauchent à communiquer ?•
2. 
Configurez la NAT afin d’accomplir ce que vous avez défini précédemment. En fonction de 
ce que vous avez défini à l’étape 2, vous devez déterminer laquelle des prochaines 
fonctionnalités utiliser :
3.

NAT statique•
NAT dynamique•
Overloading•
Toute combinaison de ces fonctionnalités.•
Vérifiez l'opération de NAT .4. 
Chacun de ces exemples de NAT vous guide tout au long des étapes 1 à 3 des étapes de 
démarrage rapide de l’image précédente. Ces exemples décrivent quelques scénarios communs 
dans lesquels Cisco vous recommande de déployer NAT.
Définir la NAT sur les interfaces intérieure et extérieure
La première étape du déploiement de la NAT consiste à définir les interfaces internes et externes 
de la NAT. Vous trouverez plus facile de définir votre réseau interne comme interne et le réseau 
externe comme externe. Cependant, les termes internes et externes sont également sujets à 
arbitrage. La figure en montre un exemple.

Topologie NAT
Exemples
1. Autoriser les utilisateurs internes à accéder à Internet
Est-il possible que vous souhaitiez autoriser les utilisateurs internes à accéder à Internet, mais 
que vous ne disposiez pas d'adresses valides suffisantes pour accueillir tout le monde ? Si toutes 
les communications avec les appareils Internet proviennent d’appareils internes, vous avez besoin 
d’une adresse unique valide ou d’un ensemble d’adresses valides.
Cette image montre un diagramme de réseau simple avec les interfaces de routeur définies 
comme internes et externes.
Adresses valides disponibles
Dans cet exemple, vous voulez que la fonction NAT autorise certains périphériques internes (les 
31 premiers de chaque sous-réseau) à établir une communication avec des périphériques 
externes et à traduire leur adresse non valide en une adresse ou un pool d'adresses valide. Le 
groupe a été défini tel que la plage d'adresses 172.16.10.1 par 172.16.10.63.
Vous pouvez maintenant configurer la NAT. Afin de réaliser ce qui est défini dans l’image 
précédente, utilisez la NAT dynamique. Avec NAT dynamique, la table de traduction dans le 
routeur est initialement vide et devient peuplée une fois que le trafic qui doit être traduit passe par 
le routeur. Par opposition à la NAT statique, où une traduction est configurée de manière statique

et placée dans la table de traduction sans qu’il soit nécessaire de générer du trafic.
Dans cet exemple, vous pouvez configurer la NAT pour traduire chacun des périphériques 
internes en une adresse valide unique, ou pour traduire chacun des périphériques internes vers la 
même adresse valide. Cette deuxième méthode est connue sous le nom de . overloading Un exemple 
de la façon de configurer chaque méthode est donné ici.
Configurer la NAT pour permettre aux utilisateurs internes d’accéder à Internet
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
 
 
ip nat pool no-overload 172.16.10.1 172.16.10.63 prefix 24 
 
 
!--- Defines a NAT pool named no-overload with a range of addresses  
!--- 172.16.10.1 - 172.16.10.63. 
 
 
ip nat inside source list 7 pool no-overload 
 
 
!--- Indicates that any packets received on the inside interface that  
!--- are permitted by access-list 7 has  
!--- the source address translated to an address out of the  
!--- NAT pool "no-overload". 
 
 
access-list 7 permit 10.10.10.0 0.0.0.31 
access-list 7 permit 10.10.20.0 0.0.0.31 
 
!--- Access-list 7 permits packets with source addresses ranging from  
!--- 10.10.10.0 through 10.10.10.31 and 10.10.20.0 through 10.10.20.31.

