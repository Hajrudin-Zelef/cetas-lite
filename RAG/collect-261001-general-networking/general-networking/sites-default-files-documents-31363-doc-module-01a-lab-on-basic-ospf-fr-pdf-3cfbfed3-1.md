---
id: collect-261001-general-networking/general-networking/sites-default-files-documents-31363-doc-module-01a-lab-on-basic-ospf-fr-pdf-3cfbfed3-1
title: "sites-default-files-documents-31363-doc-module-01a-lab-on-basic-ospf-fr-pdf-3cfbfed3"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2016-08-12"]
keywords: []
source: docs/RAG/collect-261001-general-networking/sites-default-files-documents-31363-doc-module-01a-lab-on-basic-ospf-fr-pdf-3cfbfed3.md
source_anchor: ""
source_lines: [1, 143]
sha256: b08bab2e40311c7aa922a59d386f5dad7ae785af96a5240ae5b6b074554c7118
---

# sites-default-files-documents-31363-doc-module-01a-lab-on-basic-ospf-fr-pdf-3cfbfed3

Atelier Labo ISP 
  1   
  
Module 1a – Topologie de base et OSPF 
 
Objectif: Créer une interconnexion physique de base avec une zone OSPF. S’assurez que tous les 
routeurs, les interfaces, les câbles et les connexions fonctionnent correctement. 
 
Pré-requis: Connaissance de routeur Cisco CLI,  expérience pratique antérieure. 
 
Ci-dessous la topologie couramment utilisée pour la première série de séance de labo.  
 
 
Figure 1 – Configuration de base Labo ISP

Friday, August 12, 2016   
 2 
 
Remarques de Labo 
 
Cet atelier est destiné à être exécuté sur un serveur Dynamips avec les topologies labo appropriées 
mises en place. Les routeurs dans l'environnement Dynamips util isent un IOS "Service provider". Les 
configurations et les principes de configurations décrits ci-de ssous fonctionneront sur les  versions 
Cisco IOS 12.4 et les plus récentes. Les versions antérieures C isco IOS ne sont pas supportées mais 
fonctionnent principalement avec les notes ci-dessous, mais von t manquer quelques-unes des 
fonctionnalités couvertes. 
 
Le but de ce module est de construire l’atelier de laboratoire et de présenter les principes de base de la 
construction et de la configuration d'un réseau. Un point impor tant à retenir, et celui qui sera souligné 
à maintes reprises tout au long de cet atelier, c'est qu'il y’a  une séquence distincte à la construction 
d'un réseau opérationnel: 
 
 Après que la conception physique établie, les liens entre le matériel devraient être raccordés et 
vérifiés. 
 
 Ensuite, les routeurs devraient avoir la  configuration de base  installée, et une sécurité 
élémentaire et suffisante doit être mis en place. 
 
 Ensuite la connectivité IP de base  doit être testées et éprouv ées. Ça consiste à attribuer des 
adresses IP sur tous les liens qui doivent être utilisés, et à tester les liens pour les dispositifs 
voisins.  
 
 Une fois qu’un routeur peut voir son voisin il est logique de c ommencer la configuration des 
protocoles de routage. Et commencer par IGP (OSPF est choisi pour cet atelier). La construction 
de BGP ne sert à rien si l'IGP choisi (dans ce cas, OSPF) ne fo nctionne pas correctement. BGP 
s'appuie sur le protocole OSPF pour trouver ses voisins et next  hops, et un OSPF mal ou non-
configuré se traduira par beaucoup de temps perdu à essayer de déboguer les problèmes de 
routage. 
 
 Une fois que l'IGP fonctionne correctement, la configuration BG P  peut être commencée, d'abord 
BGP interne, puis BGP externe. 
 
 N'oubliez pas d’effectuer RTFM. Qu’est ce que RTFM?  Il est essentiel que les ingénieurs réseau 
ISP utilisent pleinement toutes les ressources d'information. L a source n ° 1 est la documentation. 
Lire F#$% Manual (RTFM) est la phase traditionnelle utilisée pour informer les ingénieurs que la 
réponse est dans la documentation. 
 
 Enfin, documentatez, prenez des notes. La documentation est souvent négligée ou oubliée. C’est 
un processus continu dans cet atelier. Si l'instructeur vous de mande de documenter quelque chose, 
que ce soit sur le tableau blanc, ou à la fin de cette brochure , il est dans votre intérêt de le faire. Il 
ne peut jamais y avoir trop de documentation, et au moment de l a conception du réseau et de la 
construction, la documentation peut faire épargner beaucoup de frustration dans le futur.

Atelier Labo ISP 
  3   
  
 
 
Exercice en labo 
 
 
1. Les routeurs et les participants de l’atelier. Cet atelier est aménagé de telle sorte qu'un groupe 
de deux élèves puissent opérer un seul routeur. 14 routeurs imp liquent généralement au moins 28 
participants. Pour les ateliers avec un plus grand nombre de pa rticipants, ils doivent configurer un 
routeur unique, par groupes de t rois. Les instructeurs de l'ate lier vont partager les routeurs parmi 
les participants de l'atelier. Dans les notes suivantes, une «é quipe routeur» désigne le groupe 
assigné à un routeur particulier. 
 
2. Routeur Hostname.  Chaque routeur sera nommé en fonction de l'emplacement des tab les, 
Router1, Router2, Router3, etc Documentation et labo font égale ment référence à  Router1 comme 
R1. Au prompt du routeur, tout d'abord passer en mode enable, p uis entrez “config terminal”, ou 
simplement "config": 
 
Routeur> enable 
Router# config terminal 
Entrez les commandes de configuration, une par ligne.  Terminez avec CNTL/Z. 
Router(config)# hostname Router1 
Router1(config)# 
 
3. Désactiver la recherche de noms de domaine.  Les Routeurs Cisco tenterons toujours de 
rechercher le DNS pour un nom ou une adresse spécifiée dans la ligne de commande. Vous pouvez 
voir cela lorsque vous faites une  trace sur un routeur sans serveur DNS ou un serveur DNS avec 
aucune entrée in-addr.arpa pour les adresses IP. Nous allons désactiver pour le moment ce lookup 
pour le labo afin d'accélérer les traceroutes. 
 
Router1 (config)# no ip domain-lookup 
 
4. Désactiver  la résolution de noms en ligne de commande (Command -line Name Resolution). 
Le routeur par défaut tente  d'utiliser les différents transpor ts qu'il supporte pour résoudre les 
commandes dans la ligne de commande lors de modes normaux et de  configuration. Si les 
commandes saisies ne font pas partie de Cisco IOS, le routeur t entera d'utiliser ses autres 
transports supportés pour interpr éter la signification de ce no m.  Par exemple, si la commande 
saisie est une adresse IP, le routeur tentera automatiquement d e se connecter à cette destination 
distante. Cette fonctionnalité n'est pas souhaitable sur un routeur d’un ISP, car cela signifie que des 
erreurs typographiques peuvent entraîner des connexions étant tenté à des systèmes distants, ou les 
temps morts pendant que le routeur tente d'utiliser le DNS pour traduire le nom, et ainsi de suite.. 
 
Router1 (config)# line con 0 
Router1 (config-line)# transport preferred none 
Router1 (config-line)# line vty 0 4 
Router1 (config-line)# transport preferred none 
 
5. Désactiver le routage source.  Sauf si vous croyez vraiment qu'il est nécessaire de l’activer , le 
routage source doit être désactivée.  Cette option, activée par  défaut, permet au routeur de traiter 
les paquets avec la source options header de routage.  Cette fo nction est un risque de sécurité bien 
connue car elle permet aux sites distants d'envoyer des paquets  avec une adresse source différente 
à travers le réseau (ce qui était utile pour le dépannage des r éseaux de différents endroits sur 
Internet, mais ces dernières années il a été largement abusé des activités mécréant sur l'Internet).

Friday, August 12, 2016   
 4 
 
 
Router1 (config)# no ip source-route 
 
6. Les noms d'utilisateurs et mots de passe.  Tous les noms d'utilisateur du routeur doivent être 
isplab et tous les mots de passe doivent être lab-PW. S'il vous plaît ,  n e  p a s   c h a n g e r  l e  n o m  
d'utilisateur ou mot de passe, ou laisser le mot de passe non configuré (accès aux ports vty n'est pas 
possible si aucun mot de passe n’est activé). Il est essentiel pour le fonctionnement en douceur 
d’un laboratoire que tous les participants ont accès à tous les routeurs. 
 
  Router1 (config)# username isplab secret lab-PW 
  Router1 (config)# enable secret lab-PW 
  Router1 (config)# service password-encryption 
 
Le directive service password-encryption indique au routeur de crypter les mots de passe stockés 
dans la configuration du routeur (en dehors de >enable secret qui est déjà encrypté). 
 
Remarque A: Il peut etre tentant d'avoir simplement un nom d'utilisateur cisco et mot de passe 
cisco comme une solution au problème nom d'utilisateur / mot de passe. En aucun cas un opérateur 
fournisseur de services ne doit jamais utiliser des mots de passe faciles à deviner sur leur réseau en 
ligne 1opérationnel. 
 
