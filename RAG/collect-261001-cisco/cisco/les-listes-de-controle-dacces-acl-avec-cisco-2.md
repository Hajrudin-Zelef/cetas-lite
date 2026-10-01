---
id: collect-261001-cisco/cisco/les-listes-de-controle-dacces-acl-avec-cisco-2
title: "les-listes-de-controle-dacces-acl-avec-cisco"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-cisco/les-listes-de-controle-dacces-acl-avec-cisco.md
source_anchor: ""
source_lines: [81, 192]
sha256: 8589ce40713b28c7fd36dcd98f2863a6c0b2b95c801d6a9e6d39ef6bcae9e874
---

# les-listes-de-controle-dacces-acl-avec-cisco

**Wildcards mask ?** De quoi parle-t-on ? Une petite explication s’impose. En effet, souvent, l’objectif de mettre en œuvre l’ACL c’est de faire correspondre une plage d’adresse IP (ou un sous-réseau complet) à l'ACL plutôt qu’une seule adresse IP. Cela est faisable à l’aide ce que l’on appelle **Wildcards mask,** ou en français **masque générique (ou inversé).** Notez qu’il ne s’agit pas d’un masque de sous-réseau. Par le biais de cette notation, nous avons **la possibilité d’ignorer des parties de l’adresse IP, comme si elles correspondaient déjà.**

Cela signifie que :

- Quand on met 0 : le routeur doit comparer cet octet comme d’habitude. La présence d'un "0" "figera" la partie correspondante de l'adresse IP
- Lorsqu'on met 255 : le routeur ignore cet octet, considérant qu’il correspond déjà.

C’est-à-dire pour la règle « *deny 10.1.2.0 0.0.0.255* » le routeur va refuser tous les paquets où les trois premiers octets de l'adresse IP source correspondent à ce masque générique. D'autre valeurs sont possible pour le wildcard mask, permettant d'affiner le filtrage en fonction des plages IP ou des regroupements de plages. Ainsi, un wildcard mask en 0.0.255.255 correspondra à un masque en /16; un wildcard mask en 0.0.63.255 correspondra à un masque en /18. Pour calculer la valeur, rien de plus simple, il suffit de soustraire la valeur du masque de sous-réseau à 255.

Exemple :

255.255.255.255

-255.255.255.0

= 0.0.0.255

**On retrouve bien notre wildcard mask correspondant à notre /24.**

Deuxièmement, comme évoqué précédemment, il faut autoriser explicitement les réseaux que l’on veut laisser passer, dans notre exemple c’est « **10.1.3.0/24 ».** À défaut, il n’y aura pas de correspondance entre l’IP source contenue dans l’en-tête du paquet et les règles ACL, ce qui veut dire que le routeur va les refuser implicitement et notre réseau en 10.3.0/24 ne pourra pas non plus communiquer  avec le **10.1.2.0/24.**

Router(config)#access-list 1 permit 10.1.3.0 0.0.0.255

Notez qu'ici, on pourrait utiliser le mot clé "any" à la place de la plage réseau et du wildcard mask. Ce mot clé , qui veut dire "n'importe qui" peut être utilisé lorsque la règle s'adresse à tous les réseaux et hôtes. Comme on à déjà interdit en premier lieu notre premier réseau, on peut très bien estimer que le reste est permis.

Troisièmement, on sélection l’interface concernée par la règle :

Router(config)#interface gigabitEthernet 0/2


Router(config-if)# ip access-group 1 out

Le "1" étant le numéro de l’ACL créée à la première étape, et "out" signifie que la règle s’applique sur les paquets sortants de l’interface.

Nous avons vu les ACLs Standard et leur champ d’application, passons maintenant au deuxième type d’ACLs à savoir les « *ACLs étendues* ».

### B. Les ACL étendues

Les ACL étendues présentent plusieurs similitudes par rapport aux ACL Standards décrites dans la section précédente. Tout comme une ACL standard, **on active les ACL étendues sur les interfaces pour les paquets entrants ou sortants**, puis le routeur cherche dans la liste de manière séquentielle.

Les ACL étendues utilisent également la logique de première correspondance, car dès que la première instruction est mise en correspondance, le routeur arrête la recherche dans la liste d'ACL, en effectuant l’action définie.

En comparaison des ACL standards, **les ACLs étendues vont permettre d'analyser une plus grande variété de champs au sein de l'en-tête d'un paquet**. Cela rend les ACL étendues plus puissantes, plus précises, mais aussi un peu plus complexes.

Les ACL étendues suivent la même logique que les ACL standards, elles sont identifiables par un numéro, allant de **100 à 199** et de **200 à 2699**.

Un exemple sera plus parlant, nous allons créer une règle qui aura pour but d’**interdire le Ping de PC2 vers le PC3, tout en l'autorisant vers le PC1,** en posant les règles sur les sous-réseaux (tous en /24). Mettons ça en place en reprenant la même topologie :

Voici la configuration à appliquer sur notre routeur Cisco :

Router(config)# access-list 100 deny icmp 10.1.3.0 0.0.0.255 10.1.2.0 0.0.0.255
Router(config)# access-list 100 permit icmp 10.1.1.0 0.0.0.255 10.1.2.0 0.0.0.255
Router(config)#interface gig 0/2
Router(config-if)#ip access-group 100 in

Dans l'exemple ci-dessus, "*icmp*" correspond au protocole, en l'occurrence ici c'est un moyen de bloquer le ping. Ensuite, "10.1.3.0" c'est l'adresse IP d'origine suivie de son masque générique, et "10.1.2.0" l'adresse IP de destination suivie aussi de son masque générique.

**le numéro de port source et port de destination**. Les numéros de port identifient le service qui envoie ou reçoit les données. Quand le mot "*tcp*" ou "*udp*" est inclus dans la règle de l'ACL, cela permet de préciser le port source et le port de destination afin d'avoir une règle plus précise.

Voici un exemple :

Router(config)# access-list 100 permit tcp 10.1.1.0 0.0.0.255 10.1.2.0 0.0.0.255 eq 21

Concernant « **eq »** il s’agit d’un opérateur qui veut dire égal, et le numéro de port 21 utilisé par le serveur FTP. On autorise le réseau 10.1.1.0/24 à communiquer sur le port 21 (pour le FTP) avec le réseau 10.1.2.0/24.

Prenons un autre exemple...

L’exemple suivant se concentre sur la compréhension de la syntaxe de base. Ici, l'objectif est de créer une ACL pour empêcher l'utilisateur David d’utiliser le protocole FTP pour se connecter sur les serveurs. Pour cela, on va le bloquer en entrée sur l’interface **Gi0/0** de **R3**, car c’est l’interface la plus proche de la source, donc l’ACL sera associée à cette interface. Quant à l'utilisateur Florian, on va lui interdire l’accès au **serveur Web de serveur1**. Dans cet exemple, on active l’ACL sur **R2** en entrée de l'interface **Gi0/0**.

Voici la topologie :

Je vous propose que l'on se connecte sur **R2** afin de créer la première règle.

La première chose à faire, c’est interdire David d’accéder à tous les serveurs FTP :

access-list 101 deny tcp host 172.16.3.10 172.16.1.0 0.0.0.255 eq ftp

Au sein de la commande ci-dessus, on précise bien l'hôte de David avec "**host 172.16.3.10**", car cela correspond à la source à bloquer selon notre schéma initial.

La deuxième règle va permettre d'autoriser tout le reste du trafic :

access-list 101 permit ip any any

On applique l'ACL à l'interface concernée :

R2(config)# interface gi0/0 
R2(config-if)#ip access-group 101 in

Interdisons maintenant à Florian d’accéder au serveur Web :

access-list 101 deny tcp host 172.16.2.10 host 172.16.1.100 eq www

Note : vous pouvez remarquer que j'ai utilisé le même numéro d'ACL que pour la première réalisée. Cela n'a pas d'importance, car ce numéro est interne au routeur, et je ne suis pas sur le même routeur.


Encore une fois, on autorise tout le reste du trafic ensuite :

access-list 101 permit ip any any

Après avoir créé les règles d'ACL, il nous reste plus qu’à les appliquer toujours à partir du routeur R3 :

R3(config)# interface gi0/0
R3(config-if)#ip access-group 101 in

Voilà, mission accomplie ! 

### C. Named IP Access Lists

Comme vous le savez, toutes les listes d’accès doivent être identifiées par un nom ou un numéro. Ce type de liste d’accès est plus pratique, car on peut spécifier un nom significatif qui est facile à retenir et associer à une règle.

Les ACL nommées peuvent correspondre aux mêmes champs qu’une ACL standard et étendue, cependant, elles présentent trois grandes différences par rapport aux listes de contrôles d’accès numérotées, voyons cela ensemble :

