---
id: collect-261001-opnsense-pfsense/opnsense-pfsense/comment-creer-et-configurer-une-dmz-sur-un-firewall-opnsense-4b5a25f1-2
title: "Comment créer et configurer une DMZ sur un firewall OPNsense ?"
domain: opnsense-pfsense
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-opnsense-pfsense/comment-creer-et-configurer-une-dmz-sur-un-firewall-opnsense-4b5a25f1.md
source_anchor: ""
source_lines: [5, 51]
sha256: ddda280c7914006c029dd758f17b60dc70838916ca0afdf99970672e124b7b82
---

# Comment créer et configurer une DMZ sur un firewall OPNsense ?

Dans ce tutoriel, nous allons apprendre à créer et à configurer une DMZ sur un pare-feu OPNsense. Ceci sera utile pour isoler les services de votre organisation qui ont pour vocation à être exposés sur Internet.
Nous partons du principe que le firewall est déjà en place et qu'il dispose de deux interfaces : "LAN" et "WAN". Nous allons lui ajouter une troisième interface, destinée à accueillir notre DMZ au sein de laquelle se situera un serveur Web. La DMZ sera associée au réseau "10.10.10.0/24" et elle hébergera un serveur Web associé à l'adresse IP "10.10.10.12/24".
Si vous avez besoin d'aide pour la mise en route d'OPNsense et des précisions sur le rôle d'une DMZ, voici quelques liens utiles :
Sur le firewall OPNsense, l'interface réseau destinée à assurer la connectivité de la DMZ peut prendre la forme d'une interface réseau physique (cas d'un boitier physique), une interface réseau virtuelle (cas d'une VM) ou une interface de type VLAN (interface logique).
Pour ma part, OPNsense est mis en place sur une machine virtuelle Hyper-V, donc une nouvelle carte réseau virtuelle devra être ajoutée à la VM. Avant cela, je vais créer un nouveau commutateur virtuel de type "Externe", rattaché à une carte physique de mon hyperviseur. Ceci est utile, car je souhaite connecter à ma DMZ, une machine qui n'est pas hébergée sur cet hyperviseur. Sinon, vous pouvez créer un commutateur virtuel "Privé" (et il faudra rattacher à ce commutateur virtuel la ou les VM à connecter en DMZ).
Ensuite, vous devez éditer les propriétés de la VM "OPNsense" pour lui ajouter une carte réseau. Dans les propriétés, vous devez cliquer sur "Ajouter un matériel" et choisir "Carte réseau" avant de valider. Cette opération peut être effectuée à chaud, ce qui évite les interruptions éventuelles de la production.
Cette nouvelle carte réseau doit être associée au commutateur virtuel "DMZ" nouvellement créé. Voici un exemple :
La suite des opérations se déroulera sur l'interface d'administration d'OPNsense.
Connectez-vous à l'interface d'administration d'OPNsense et naviguez dans les menus de cette façon :
1 - Cliquez sur "Interfaces" puis "Assignations".
2 - Sous "Assign a new interface", vous devriez voir une carte disponible. Ici, elle est nommée "hn2". Ajoutez la description "DMZ".
3 - Cliquez sur "Ajouter" pour valider.
Remarque : le nom donné à l'interface n'a aucun impact, nous pourrions indiquer n'importe quel nom... L'essentiel étant de pouvoir s'y retrouver.
Voilà, désormais, il y a bien 3 interfaces assignées sur notre firewall OPNsense. Pour finir, cliquez sur "Sauvegarder" pour valider la modification définitivement. D'un point de vue d'OPNsense, cette interface porte l'identifiant "opt1".
Pour le moment, l'interface n'est pas active bien qu'elle soit créée. Pour la suite de la configuration, voici les étapes à effectuer :
1 - Sous "Interfaces", sélectionnez l'interface "DMZ".
2 - Sous "Configuration de base", cochez l'option nommée "Activer l'interface".
3 - Ajouter une nouvelle fois "DMZ" en tant que description, bien que ce soit facultatif.
Toujours sur cette même page, il est important d'effectuer la configuration réseau de cette interface. Au niveau de "Type de configuration IPv4", choisissez "Adresse IPv4 statique".
Puis, descendez dans la page afin de définir l'adresse IPv4 et le masque de sous-réseau. Par exemple : 10.10.10.1/24. Puisqu'il s'agit d'une DMZ, le masque de sous-réseau pourrait être plus restrictif parce que généralement le nombre d'hôtes connectés est faible. Configurez éventuellement d'autres paramètres avant de sauvegarder.
À partir de maintenant, le firewall OPNsense dispose d'une interface réseau "DMZ" prête à l'emploi.
Notre firewall dispose de trois interfaces associées à trois réseaux distincts, avec des fonctions différentes : LAN, WAN et DMZ. Désormais, nous devons nous intéresser aux flux réseau autorisés ou bloqués entre la DMZ et les autres réseaux.
De part sa fonction, la zone réseau "DMZ" doit être isolée des autres réseaux et les accès entrants ou sortants doivent être limités au strict nécessaire. Autrement dit, il ne doit pas être possible de communiquer entre la DMZ et le LAN, à quelques exceptions près. Qu'est-ce qu'une exception ? Il s'agit d'un flux spécifique que l'on va autoriser.
Dans notre cas, le LAN doit pouvoir communiquer avec la DMZ uniquement sur le port 80, correspondant au HTTP, pour que l'on soit en mesure de se connecter à notre serveur Web.
La DMZ ne doit pas pouvoir communiquer avec les autres réseaux locaux, et vers Internet, les flux doivent être limités (HTTP/HTTPS). Nous pouvons utiliser le firewall OPNsense en tant que résolveur DNS, ce qui évite d'ouvrir un flux direct vers un serveur DNS externe.
En résumé, nous devons créer les règles suivantes.
| Action | Interface | Source | Source Port | Destination | Dest. Port | Protocole | 
|---|---|---|---|---|---|---|
| Autoriser | DMZ | DMZ Net | Any | DMZ Address | 53/UDP | DNS | 
| Action | Interface | Source | Source Port | Destination | Dest. Port | Protocole | 
|---|---|---|---|---|---|---|
| Autoriser | DMZ | DMZ Net | Any | Any sauf LAN Net | 80/TCP | HTTP | 
| Autoriser | DMZ | DMZ Net | Any | Any sauf LAN Net | 443/TCP | HTTPS | 
| Action | Interface | Source | Source Port | Destination | Dest. Port | Protocole | 
|---|---|---|---|---|---|---|
| Autoriser | LAN | LAN Net | Any | 10.10.10.12/32 | 80/TCP | HTTP | 
Pour que les hôtes connectés au LAN puissent accéder à Internet, il convient aussi de créer les règles n°1 et n°2 sur l'interface LAN, en adaptant les sources et les destinations. Ici, nous nous concentrons sur la configuration de la DMZ.
Pour créer une règle sur l'interface DMZ du firewall OPNsense, il faut se rendre à l'emplacement suivant : Pare-feu > Règles > DMZ. Ensuite, il est nécessaire de créer une nouvelle règle en appuyant sur le bouton "+".
Vous devez savoir que :
Ensuite, vous devez créer les règles en vous appuyant sur les tableaux ci-dessus. Veillez aussi à ajouter une description pour indiquer l'utilité de chaque règle.
Lorsque c'est indiqué en tant que destination "Any sauf LAN Net", cela signifie que l'on autorise le flux sauf si la destination se situe dans le LAN. Ceci permet de couper les communications de la DMZ vers le LAN et inversement.
Dans l'assistant de création de règles d'OPNsense, vous devez cocher l'option "Destination / Inverser" et indiquer "LAN Net" comme "Destination" pour avoir ce comportement (tout ce qui n'est pas du "LAN Net" sera autorisé).
Ici, il n'y a qu'un seul réseau local, donc c'est simple. S'il y a plusieurs réseaux, plutôt que de créer plusieurs règles, vous pouvez créer un groupe de réseaux via la fonction d'Alias (Aliases). Il suffira de sélectionner l'Alias par son nom lors de la création de la règle. Cette fonction se situe dans : Pare-feu > Alias.
Remarque : une autre façon de procéder, c'est de déclarer les réseaux privés des différentes classes (10.0.0.0/8 - 172.16.0.0/12 - 192.168.0.0/16), pour être certain d'isoler votre DMZ. Si vous gérez une liste manuelle, ceci vous oblige à la maintenir.
La création de règles est une action répétitive, bien que vous puissiez dupliquer les règles grâce au bouton mis en évidence sur l'image ci-dessous. Ceci vous permettra de gagner du temps (mais attention aux erreurs !).
Pour vous aiguiller, voici les règles que vous devez avoir sur les interfaces "LAN" et "DMZ". Dans ma configuration, le NAT sortant est en mode automatique, donc OPNsense a créé des règles de lui-même pour les réseaux correspondants à mon LAN et à ma DMZ.
Vous pouvez consulter la documentation officielle en complément :
