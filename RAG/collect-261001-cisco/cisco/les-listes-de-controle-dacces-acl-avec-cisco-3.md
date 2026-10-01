---
id: collect-261001-cisco/cisco/les-listes-de-controle-dacces-acl-avec-cisco-3
title: "les-listes-de-controle-dacces-acl-avec-cisco"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/les-listes-de-controle-dacces-acl-avec-cisco.md
source_anchor: ""
source_lines: [193, 283]
sha256: cff95e88ebdf2e5b51f159fa93f196e013978f28e0b19d331706e4ba2c36389d
---

# les-listes-de-controle-dacces-acl-avec-cisco

- L’utilisation de noms au lieu de chiffres pour identifier la liste ACL facilite la mémorisation et l'identification de l'ACL.
- Utilisation de sous-commandes ACL, et non de commandes globales, pour définir l'action et les paramètres correspondants.
- Utilisation des fonctionnalités d'édition de l'ACL qui permettent de supprimer des lignes individuelles de l'ACL et d'insérer de nouvelles instructions à une liste d'accès nommée.

On va reproduire le même schéma :

Cette fois-ci, on va changer un peu les règles du jeu : le but sera d’interdire David d’accéder au serveur Web (Serveur1) via le protocole http. On aura besoin donc d’une ACL qui contiendra plusieurs informations : IP source, IP destination et un protocole bien précis. Pas le choix, il faudrait utiliser une ACL étendue.

Profitons de cet exemple pour utiliser une ACL nommée et poursuivre notre découverte des ACLs.

Sur le R1, tout d’abord, on crée l’ACLs avec le nom "David-HTTP" :

`R2(config)# ip access-list extended David-HTTP`
Puis, on se retrouve dans cette ACL, donc on ajoute notre règle :

`R2(config-ext-nacl)#deny tcp host 172.16.3.10 host 172.16.1.100 eq www`
Vous constatez bien que le prompt a changé, car on se trouve dans l’invite de configuration de l’ACL nommée. On pose une interdiction via la directive "deny" depuis l’adresse source "172.16.3.10" vers le serveur web avec l'adresse IP "172.16.1.100". **« www »** c’est pour faire référence au service web, mais on peut bien évidement mettre le port 80 au lieu de www.

En dernier lieu, on autorise toutes les autres connexions, car souvenez-vous, la directive implicite (deny any ) se trouve à la fin de chaque ACL !

R1(config-ext-nacl)#permit ip any any

Il ne reste plus qu’à associer l’ACL à l’interface G0/0 de R2 comme vous savez le faire maintenant. Vous pouvez retrouver d'autres exemples sur le site de Cisco.

## V. Placement des ACL

Bon, nous savons maintenant qu'il y a deux types d'ACL : standards et étendues. On sait aussi qu'il y a deux possibilités pour les placer : en entrée ou en sortie d'une interface. Mais alors, où placer tel type d'ACL? Sur quel routeur? Quelle interface?

Dans tous les cas, il existe une règle simple et immuable : **une ACL standard se placera toujours au plus proche de la destination, une ACL étendue se placera toujours au plus proche de la source.**

Pourquoi? Tout simplement, car une ACL standard, comme vous avez pu le constater, est très restrictive (blocage de tout le trafic d'un réseau ou d'un hôte), un placement au plus proche de la source risque de limiter fortement les possibilités de celles-ci. En revanche, l'ACL étendue filtrant au niveau de la couche 4 (ports TCP/UDP), son placement au plus proche de la source permet d'éviter de faire transiter des paquets qu'on ne souhaite pas voir sur le réseau, ainsi, on économise du temps de calcul.

Un petit exemple, en reprenant le schéma de la partie II :

Admettons que le PC1 soit la source, à savoir que c'est à partir de lui que je vais déterminer mes règles. Je veux interdire à ce PC1 d'accéder au réseau du serveur S1, je vais donc placer l'ACL standard sur l'interface **Gi 0/3** de **R2 en sortie** **.** En revanche, si je veux interdire à PC1 d'accéder en SSH au serveur S1, je vais placer mon ACL étendue sur l'interface **Gi 1/0** de **R1 en entrée.**

## VI. Conclusion

Nous avons vu ensemble **comment les ACLs peuvent être utiles pour contrôler le réseau afin de filtrer certains flux**. Les ACL permettent un filtrage au niveau du réseau et peuvent être complétées avec le filtrage applicatif (Proxy) qui va venir **au niveau de la couche Application du modèle OSI.** Le filtrage de paquet via **ACLs opère niveaux 3 et 4 du modèle OSI**, ce qui permet de faire déjà beaucoup de choses, mais peu ne pas suffire dans certains cas.

*Merci à Florian Duchemin pour sa relecture technique avant publication.*

Bonjour,

Déjà super tutoriel, merci pour le temps passé dessus.

J’ai cependant l’impression qu’une erreur s’est glissée dans le texte, à ce niveau

« Je vous propose que l’on se connecte sur R2 afin de créer la première règle. »

Sauf erreur de ma part, je pense qu’il s’agit en réalité du R3 dont on parle.

Merci pour le boulot que vous faites, le site est une vraie bible.

Merci ça m’a un peu embrouillé aussi

tout est mélangé notamment sur les ACL standard

dans celle ci on indique la source du paquet IP, or dans les exemples de conf c’est la destination.

cela apporte trop de confusions pour un lecteur non assidu,

vous auriez peut etre du vous relire

A propos des ACL étendues,

il y’aurait une erreur de frappe:

<>

=> [100;199]et[2000;2699] au lieu de [100;199]et[200;2699].

bonjour,

excellent tutoriel.

il y a une erreur sur l’acl etendue.

vous avez inversé la config entre le routeur R2 et R3 david florient

Merci beaucoup pour vos tutos , cependant y’a beaucoup d’erreur je pense , par exemple sur les acls étendus , y’a confusion entre l’ip source et l’ip de destination à plusieurs endroits, ça embrouille

Bonjour, je pense qu’il y a une petite confusion au niveau de la configuration de l’ACL standard.

« nous allons, sur l’interface Gi0/2, interdire les paquets provenant du réseau « 10.1.1.0/24 » »

« Router(config)#access-list 1 deny 10.1.2.0 0.0.0.255 »

cette configuration est appliquée à Gi0/2 en sortie.

N’est-ce pas plutôt « »access-list 1 deny 10.1.1.0 0.0.0.255″ » appliquée à Gi0/2 en sortie ?
