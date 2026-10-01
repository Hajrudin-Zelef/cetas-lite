---
id: collect-261001-general-networking/general-networking/sites-default-files-documents-31363-doc-module-01a-lab-on-basic-ospf-fr-pdf-3cfbfed3-4
title: "sites-default-files-documents-31363-doc-module-01a-lab-on-basic-ospf-fr-pdf-3cfbfed3"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2016-08-12"]
keywords: ["distribution"]
source: docs/RAG/collect-261001-general-networking/sites-default-files-documents-31363-doc-module-01a-lab-on-basic-ospf-fr-pdf-3cfbfed3.md
source_anchor: ""
source_lines: [416, 565]
sha256: 95fd8b6b9a37d185c2fb85f46595fdd547c42b4e06ae75496ce9fd98022ac473
---

# sites-default-files-documents-31363-doc-module-01a-lab-on-basic-ospf-fr-pdf-3cfbfed3

19. Éviter Blackhole de trafic au Redémarraage.  Quand un routeur redémarre après avoir été mis 
hors service, OSPF va commencer la distribution des préfixes dè s que les  adjacences sont établies 
avec ses voisins. Dans la prochaine partie du laboratoire de l' atelier, nous présenterons iBGP. 
Donc, si un routeur redémarre, OSPF va démarrer bien avant que le maillage iBGP est rétabli. Ceci 
entraînera l'atterrissage du routeur dans le trajet de transit pour trafic, sans que la table de routage 
soit complétée par BGP. Il n’y aura pas d’information de routag e complète sur le routeur, de sorte 
que tout le trafic de transit (à  partir de client-to-peer ou en  a m o n t ,  o u  v i c e - v e r s a )  s e r o n t  s o i t  
abandonnées, ou résultant rebondissement des paquets entre rout eurs adjacents. Pour éviter ce 
problème, Il nécessaire que le routeur n’annonce sa disponibili té jusqu'à ce que le maillage iBGP 
est en marche. Pour faire ça, nous devons fournir la commande suivante: 
 
Router1(config)#router ospf 41 
Router1(config-router)#max-metric router-lsa on-startup wait-for-bgp 
 
Ceci met en place OSPF de telles sorte que les routes passant p ar ce routeur sera marqué comme 
inaccessible (très haute métrique) jusqu'à ce que iBGP soit en marche. Une fois iBGP est en 
marche, les préfixes distribués par OSPF vont revenir à des val eurs métriques standard, et le 
routeur va passer le trafic de transit comme d'habitude. 
  
20. (Optionnel). Activer le nom DNS et la résolution d'adresse sur les routeurs. Si les instructeurs 
de l'atelier ont mis en place le nameserver dans l'atelier, à c e stade, toutes les équipes routeur 
devrait maintenant permettre  DNS lookups sur leurs routeurs. O SPF est porteur de tous les 
préfixes, y compris le réseau de connexion à Router15, autour d e la salle de classe, de sorte que 
tous les routeurs doivent être en mesure de voir Router15.

Friday, August 12, 2016   
 10 
 
 
Router2(config)#ip domain-lookup 
Router2(config)#ip name-server 192.168.1.4 
Router2(config)#ip domain-name workshop.net 
 
Ces commandes défont ce qui a été configurée à l'étape 3, au dé but de ce module. Assurez-vous 
que vous pouvez effectuer le ping du nameserver avant de faire cela. Si vous ne pouvez pas 
effectuer le ping sur le nameserver, chercher à savoir pourquoi. 
 
Notez que l'équipe qui opère le routeur 6 sera ajouté pour pouv oir ajouter la configuration qui 
permettra au bloc adresse du serveur DNS de se propager à trave rs le réseau de classe. La 
configuration supplémentaire pour Router6 est la suivante: 
 
router ospf 41 
 network 192.168.1.0 0.0.0.255 area 0 
! 
interface FastEthernet0/1 
 ip addr 192.168.1.254 255.255.255.0 
 no shutdown 
! 
 
21. (Optional). Activer OSPF name lookups sur les routeurs. Enchainant l'étape précédente, 
permettre maintenant les name lookups OSPF sur le routeur. 
 
Router2(config)#ip ospf name-lookup 
 
Cette commande permet l'affichage de l'OSPF router-id comme nom s de domaine. Ainsi, plutôt 
que d'afficher la sortie suivante avec les name lookups désactivés: 
 
router2>sh ip ospf neigh 
 
Neighbor ID     Pri   State           Dead Time   Address       Interface 
10.0.15.241       1   FULL/BDR        00:00:36    10.0.15.1     FastEthernet0/0 
10.0.15.244       1   FULL/  -        00:00:32    10.0.15.18    Serial1/0 
10.0.15.254       1   FULL/DR         00:00:38    10.0.15.26    FastEthernet0/1 
 
      le routeur affichera les informations suivantes: 
 
   router2#sh ip ospf neigh 
 
   Neighbor ID     Pri   State           Dead Time   Address       Interface 
   router1.worksho   1   FULL/BDR        00:00:33    10.0.15.1     FastEthernet0/0 
   router4.worksho   1   FULL/  -        00:00:39    10.0.15.18    Serial1/0 
   router14.worksh   1   FULL/DR         00:00:35    10.0.15.26    FastEthernet0/1 
 
     ce qui est beaucoup plus informatif. 
 
22. Ping Test #2. Ping toutes les interfaces loopback dans la salle de classe. C ela permettra d'assurer 
que l'IGP OSPF est connecté End-to-End. Si vous rencontrez des problèmes utilisez les 
commandes suivantes pour les résoudre :  
 
show ip route   : : Voir s'il ya un itinéraire pour la destination prévue 
show ip ospf   : Voir informations générales OSPF

Atelier Labo ISP 
  1 1   
  
show ip ospf interface : Vérifier si le protocole OSPF est activé sur toutes les 
interfaces destinées 
show ip ospf neighbor : Voir la liste des voisins OSPF que le routeur voit 
 
Checkpoint #2:  appelez l’assistant de laboratoire pour vérifier la connectivité. Sauvegarder la 
configuration telle qu'elle est sur le routeur - utiliser une feuille de calcul distincte ou l'espace de 
travail à la fin de ce module. Vous aurez besoin de  cette configuration à plusieurs reprises tout au 
long de l'atelier. 
 
 
23. Traceroute vers tous les routeurs.  Une fois que vous pouvez effectuer le ping sur tous les 
routeurs, essayez de suivre des routes vers tous les routeurs u tilisant  trace xxxx commande. Par 
exemple, équipe routeur 1 entre: 
 
Router1# trace 10.0.15.252 
  
pour suivre une route vers le routeur R12. Si le temps limite est écoulé sur chaque hop en raison de 
destinations inaccessibles, il est possible d'interrompre le traceroute  à  l ' a i d e  d e  l a  s é q u e n c e  
d'interruption Cisco CTRL-^. 
 
 Q. Pourquoi certains chemins traces montrent plusieurs adresses IP par hop? 
  
A. S'il ya plus d'un des chemins de coût égal, OSPF va effectuer l e "load share" trafic entre ces 
chemins. 
 
Router1>trace router12 
 
Tapez la séquence d'échappement pour annuler. 
Tracing the route to router12.workshop.net (10.0.15.224) 
 
  1 fe0-0.router2.workshop.net (10.0.15.2) 4 msec 
    fe0-1.router13.workshop.net (10.0.15.6) 0 msec 
    fe0-0.router2.workshop.net (10.0.15.2) 0 msec 
  2 fe0-0.router14.workshop.net (10.0.15.54) 4 msec 
    fe0-1.router14.workshop.net (10.0.15.26) 4 msec 
    fe0-0.router14.workshop.net (10.0.15.54) 0 msec 
  3 ser0-0.router12.workshop.net (10.0.15.69) 4 msec *  4 msec 
Router1> 
 
24. Autres caractéristiques dans OSPF. Revoir la documentation ou la ligne de commande d'aide en 
en tapant ?  pour voir d'autres commandes show et d’autres fonctions de configuration OSPF . 
 
25. Configuration avancée.  Ces équipes routeur qui ont achevé ce module doivent se référe r au 
module 11 de l'atelier avancée BGP. Les étapes set-up ont été é tendues pour inclure toutes les 
exigences de base d'un routeur qui est utilisé dans un backbone  ISP. En attendant que le module 
soit  conclut, il serait maintenant un bon moment pour réviser le module avancé et incorporer les 
ajouts à la configuration utilisée ici.

Friday, August 12, 2016   
 12 
 
 
 
Questions de révision 
 
1. Quel protocole IP que Ping et Traceroute utilisent? 
 
2. Ping sur l'adresse IP du routeur de votre voisin (par exemple 1 0.0.15.2). Notez le temps qu'il a 
fallu pour le ping à compléter la tache. Maintenant, un autre p ing sur l'adresse IP de votre routeur 
sur le même segment (par exemple 10.0.15.1). Notez le temps qu' il a fallu pour compléter une 
table de ping. Quels sont les résultats? Pourquoi y a-t-il une différence? 
 
3. Quel IOS show command(s) affiche la table de forwarding du routeur? 
 
4. Quel IOS show command(s) affiche la base de données OSPF du routeur?
