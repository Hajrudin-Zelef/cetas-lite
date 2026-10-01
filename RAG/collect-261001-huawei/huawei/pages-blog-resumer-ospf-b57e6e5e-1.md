---
id: collect-261001-huawei/huawei/pages-blog-resumer-ospf-b57e6e5e-1
title: "pages-blog-resumer-ospf-b57e6e5e"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["cost", "ethernet"]
source: docs/RAG/collect-261001-huawei/pages-blog-resumer-ospf-b57e6e5e.md
source_anchor: ""
source_lines: [1, 85]
sha256: 934dddabfe86c0173c18ea7733d68c94b8e4481661f2e6c9e557757202da4e69
---

# pages-blog-resumer-ospf-b57e6e5e

Table of Contents
OSPF (Ouvrir le chemin le plus court en premier)
Le protocole de routage OSPF fait partie des protocoles à état de liaison.
C'est un standard ouvert, ce qui signifie qu'il peut aussi tourner sur des équipements autres que Cisco.
Contrairement au protocole EIGRP qui lui est propriétaire Cisco, donc il ne fonctionne que sur des routeurs Cisco.
Sur cette topologie réseau, le routeur R1  n'est pas configuré en ce protocole de routage. 
Nous allons donc le configurer :
La commande « router ospf » permet d'activer le protocole.
Le numéro « 1 » est le numéro du processus de ce protocole de routage.
Il n'a pas nécessairement d'importance, on peut mettre ce que l'on veut ! Ici, j'ai choisi le chiffre « 1 ».
Avant de lui indiquer les réseaux, on va lui donner une identité avec la commande « router-id ».
Ici sur l'appeler 1.1.1.1.
On n'est pas obligé de le configurer, car |par défaut, le routeur prendra l'adresse IP la plus élevée de ces interfaces de bouclage.
Et s'il n'y en a pas, il prendra alors l'adresse IP la plus élevée de ces interfaces !
Et ensuite il nous reste plus qu'à lui déclarer les réseaux qui lui sont directement connectés avec la commande « Network ».
Avec le protocole OSPF, on utilise des masques inversés ! Il s'agit de l'inverse du masque de sous-réseau :
Par exemple, si on soustrait 4 fois 255, par le masque de sous-réseau,
Cela nous donne, le masque inversé !
La formule est donc 4 fois 255 moins le masque de sous-réseau, et cela nous donne le masque inversé !
Ce protocole de routage est très gourmand en CPU, c'est pourquoi il utilise le masque inversé, car dans notre exemple, il va se concentrer uniquement sur les 3 premiers octets, il va donc ignorer le dernier octet.
Le masque inversé s'utilise dans OSPF, EIGRP et les access-list !
On va maintenant revenir sur notre topologie !
Dans la suite de la commande, nous avons Area 0
OSPF utilise des airs afin de découper un groupe de routeur en plusieurs petits groupes de travail.
Ça permet de réduire la charge CPU des routeurs, car l'algorithme, qui calcule les routes lors d'un changement, s'exécute que dans le même air OSPF.
Un routeur qui a des interfaces dans plus d'une zone, par exemple notre routeur 4 au milieu, se fait appeler ABR.
Et un routeur qui connecte un réseau OSPF à d'autres domaines de routage, comme notre routeur 7, s'appelle ASBR.
Les autres routeurs sont simplement des routeurs internes à un seul air ospf.
C'est-à-dire que l'ensemble de leur interface appartient à la même air OSPF.
En gros, il faut toujours commencer par l’air 0 en premiers, et ne pas faire des airs ospf de plus de 15 routeurs !
Nous venons donc de déclarer l’ensemble des réseaux qui sont directement connectés à notre routeur 1.
Il peut maintenant découvrir les routeurs faisant partie de l’air 0.
Par contre, on voit que le réseau 192.168.1.0 n’est pas un réseau de routeur, mais un réseau d’utilisateur.
Pour éviter que le routeur1 cherche d’autres routeurs sur un réseau d’utilisateur, on va exécuter la commande « passive-interface » sur son interface Fast Ethernet 0/3 !
Désormais que notre routeur 1 est bien configuré en ce protocole de routage, on va voir ce qui se passe avec des commandes « show » !
La commande « show IP ospf neighbor » permet d’afficher les voisins de ce protocole de routage.
Si 2 routeurs appartiennent au même réseau et qu’ils sont dans la même air OSPF, alors ils vont s’envoyer des paquets de types « Hello » pour se présenter.
Un paquet « Hello » permet :
- De découvrir le réseau
- De se construire de nouvelles relations de voisinage
- Et de s’assurer de leurs disponibilités.
Maintenant que notre routeur R1 à découvert ses voisins, il doit les stocker quelque part.
C’est donc bien, c'est dans la LSDB qu’il stocke toutes les informations de ses voisins !
Et les infos qu’il reçoit de ses voisins s’appellent des LSA !
Pour montrer une image, c’est comme si les LSA’s étaient des livres
Et la LSDB, une bibliothèque !
OSPF met toutes les informations qu’il reçoit dans une Bibliothèque, et les informations reçues s’appellent des LSA.
Et la commande « show IP ospf database » permet ainsi d'afficher le contenu de la bibliothèque et la couverture de nos livres !
La dernière commande « show » que l’on va voir est « show IP route » !
Ici, on peut voir dans la table de routage que le routeur 1 a découvert 3 réseaux :
- Le réseau 2.2.2.0
- Le 5.5.5.0
- Et le 6 6 6 0 !
La lettre O nous informe que ces réseaux ont été appris via le protocole de routage.
Et la dernière chose qui nous reste à voir sur ce protocole de routage est le calcul des coûts de liaison !
Maintenant que notre routeur est bien configuré en ce protocole de routage, Il va utiliser les informations présentes dans sa bibliothèque, la LSDB afin de trouver le meilleur chemin vers ses destinataires en utilisant son algorithme.
Pour cela, il va définir une métrique pour chacune de ses liaisons avec la formule 100 / Débit !
Le chiffre 100 est la bande passante de référence.
Plus le résultat de cette métrique sera faible et plus le chemin sera privilégié !
Ce qui donne pour une liaison Fast Ethernet de 100méga, un coût de 1.
Et pour une interface Ethernet, un coût de 10.
Voici, les différents coûts rapportés sur les interfaces !
Les liaisons les plus lentes sont celles en Ethernet entre le routeur 1et 4, et celle entre le 4 et 7 !
Alors à votre avis, si le routeur 1 souhaite joindre le réseau 192.168.2.0, par où va-t-il passer ?
Par le routeur 2, 4 ou 5 ?
Eh bien, dans cet exemple, les flux passeront aussi bien par le routeur 2 que par le routeur 5, car ils ont des coûts égaux.
Le routeur fera donc du load-balancing.
Pour joindre le réseau à l'autre bout, le coût sera de 4 par R2 et R5, et de 21 en passant par le routeur 4.
Là est la différence avec le protocole de routage RIP, où lui, il serait passé par les deux routeurs du milieu, car le chemin est plus court, mais pas le plus rapide !
Il est possible de modifier le coût d'une interface, soi-même, avec la commande « IP ospf cost ».
Avec les nouvelles liaisons plus rapides qui sont apparues, comme le gigabit ou bien les 10 gigabits, il est indispensable de modifier la bande passante de référence avec la commande : « auto-cost ».
Retrouvez tous nos cours pour réussir votre CCNA sur la chaîne YouTube de Formip.
Pour mieux comprendre les spécificités et les avantages de chaque version du protocole Internet, découvrez notre article détaillé sur les différences entre IPv4 et IPv6.
FAQs
Quelle est la principale différence entre IPv4 et IPv6 ?
IPv4 utilise des adresses sur 32 bits, tandis qu’IPv6 repose sur des adresses sur 128 bits, offrant ainsi un nombre presque illimité d’adresses IP disponibles.
Pourquoi passer d’IPv4 à IPv6 ?
Le passage à IPv6 permet de pallier le manque d’adresses IPv4, d’améliorer la sécurité du réseau et de simplifier la configuration automatique des appareils connectés.
IPv4 et IPv6 peuvent-ils coexister sur un même réseau ?
Oui, il est tout à fait possible de faire coexister IPv4 et IPv6 grâce à des mécanismes de transition comme le dual stack, la traduction (NAT64) ou le tunneling.
Quels sont les avantages d’IPv6 par rapport à IPv4 ?
IPv6 offre une configuration automatique, une sécurité intégrée (IPsec), une meilleure efficacité du routage et la suppression du besoin de NAT (Network Address Translation).
Comment activer IPv6 sur un routeur Cisco ?
Sur un routeur Cisco, il suffit d’activer le routage IPv6 avec la commande globale ipv6 unicast-routing, puis de configurer les interfaces avec des adresses IPv6.
L’IPv6 est-il plus rapide qu’IPv4 ?
