---
id: collect-261001-cisco/cisco/pugnere-presentations-administration-reseau-config-cisco-pdf-302726c4-1
title: "show int"
domain: cisco
role: reference
task: reference
actors: []
dates: ["2010-10-06"]
keywords: ["attention", "copyright", "ethernet", "memory"]
source: docs/RAG/collect-261001-cisco/pugnere-presentations-administration-reseau-config-cisco-pdf-302726c4.md
source_anchor: ""
source_lines: [1, 161]
sha256: 2138f8a968a475ed4d950acf20d0f05826bd599f7229715a68d895a477d3aaea
---

# show int

Tutoriel : Configuration d’un commutateur ou d'un routeur CISCO                            2010-10-06/90
Configuration d’un commutateur ou d'un 
routeur CISCO
Ce document n'est pas un tutoriel complet de configuration d'un commutateur ou d'un routeur 
Cisco, mais plutôt un recueil de commandes utilisées fréquemment pour commencer à configurer 
un commutateur ou un routeur.  La validité des commandes présentées varient en fonction du type 
de commutateur et de la version de l'IOS du matériel.
Connexion en série
Pour la première configuration d'un commutateur, la seule façon de le configurer est par 
l'intermédiaire d'une liaison Serie RS232 connectée au port console du commutateur.
La plupart des commutateurs ont une connexion série RS232 configurée par défaut en : 
– 9600 bauds,
– 8 bits de données, 
– pas de parité, 
– 1 bit de stop,
– pas de contrôle de flux.
Sous windows, on peut utiliser l'hyper terminal, sous linux on peut utiliser un émulateur comme 
kermit. Voici quelques commandes relatives à kermit : 
set modem type none     ; on utilise une connexion directe (sans modem)
set line /dev/ttyS0     ; Specify device name
set carrier-watch off   ; If DTR CD are not cross-connected
set escape-character ^A ; changer le caractere d'echapement
set flow none           ; If you can't use RTS/CTS
set speed 57600         ; Or other desired speed
set parity none         ; (or "mark" or "space", if necessary)
set stop-bits 1         ; (rarely necessary)
show escape
connect                 ; Enter Connect (terminal) state
Exemple pour se connecter avec le minimum de commandes : 
$ kermit
C-Kermit 8.0.209, 17 Mar 2003, for Red Hat Linux 8.0 
 Copyright (C) 1985, 2003, 
  Trustees of Columbia University in the City of New York. 
Type ? or HELP for help. 
C-Kermit>set modem type none 
C-Kermit>set line /dev/ttyS0 
C-Kermit>set carrier-watch off 
C-Kermit>set escape-character ^A    
C-Kermit>connect 
Dans l'exemple ci-dessus, la sortie de la connexion de kermit se fera par la combinaison de 
touches :  
Ctrl-A
q
Denis Pugnère (CNRS / IN2P3 / IPNL) Page 1 / 10

Tutoriel : Configuration d’un commutateur ou d'un routeur CISCO                            2010-10-06/90
Tant que le commutateur n'est pas configuré, le seul moyen de le configurer est par l'intermédiaire 
du port console (RS232). Au démarrage, il est possible de rentrer dans le setup. À l'aide de plusieurs 
questions, le setup permet de rentrer les paramètres de base et les adresses IP. Si on ne rentre pas 
dans le setup, il est alors possible de le configurer en ligne de commande.
Identification
2 modes d’identification : 
– Mode utilisateur (user) : le mot de passe est demandé lors de la connexion via telnet/ssh ou 
à la console. Dans ce mode, nous avons accès à un sous ensemble des commandes : 
uniquement certaines commandes de visualisation d'informations
– Mode administrateur (privileged ou enable) : C'est dans ce mode que l'on pourra configurer 
le commutateur.
Visualisation de la configuration en mode utilisateur :
Liste des commandes disponibles. Attention elle est différente suivant le mode d’identification. Le ? 
Est aussi utilisable pour connaître les arguments d'une commande : 
> ? 
On peut saisir le début de la commande, elle peut être complétée (complétion) comme le bash en 
utilisant la touche 'TAB'.
Visualisation du modèle de commutateur, de la version de l'IOS et des numéros de série : 
> sh version
 
Visualisation des switches de la stack (si plusieurs switches sont stackés) : 
> sh switch
Visualisation de l'allocation de la mémoire aux objets de l'IOS : 
> sh memory
Pour surveiller la charge du routeur (table des processus de l'IOS)
> sh processes
Visualisation de la mémoire utilisée et disponible : 
> sh processes memory
Visualisation d'un résumé des statistiques des queues de toutes les interfaces : 
> sh interfaces summary
Visualisation d'un résumé de l'état de toutes les interfaces (connecté/non connecté, vlan, vitesse, 
type) : 
> sh interfaces status
Visualisation du détail d'une interface (ici la GigabitEthernet 1/0/1) : 
Denis Pugnère (CNRS / IN2P3 / IPNL) Page 2 / 10

Tutoriel : Configuration d’un commutateur ou d'un routeur CISCO                            2010-10-06/90
> sh interface Gigabitethernet 1/0/1
Passer en mode privilégié afin de modifier la configuration : 
> enable
Un mot de passe « enable » est éventuellement demandé. Le prompt '>' devient '#' pour indiquer que 
l'on est en mode privilégié.
Visualiser l’état des interfaces :
# show int 
visualiser la table des routes IP : 
# sh ip route 
Visualisation de la table arp : 
# sh ip arp
Compte les trames à destination du routeur (et non toutes celles qui passent) : 
# sh ip traffic 
Interrogation des logs des ACL : 
# show ip accounting access-violations 
Réinitialisation des compteurs de l'accounting : 
# clear ip accounting
Affichage de la table des adresses ethernet (dynamic, static, vlan...) : 
#sh mac address-table dynamic     
          Mac Address Table 
------------------------------------------- 
Vlan    Mac Address       Type        Ports 
----    -----------       --------    ----- 
   1    0024.98ec.6c04    DYNAMIC     Fa0/20 
  10    0024.98ec.6c04    DYNAMIC     Fa0/20 
  12    001b.63b6.6ece    DYNAMIC     Fa0/20 
Total Mac Addresses for this criterion: 3 
Effacer la table d'adresses mac : 
#clear mac address-table dynamic
Affichage de la table ARP : 
#sh arp 
Protocol  Address          Age (min)  Hardware Addr   Type   Interface 
Internet  192.168.10.4            -   0011.bb49.6080  ARPA   Vlan10 
Effacer la table ARP : 
#clear arp-cache
Denis Pugnère (CNRS / IN2P3 / IPNL) Page 3 / 10

Tutoriel : Configuration d’un commutateur ou d'un routeur CISCO                            2010-10-06/90
Modification de la configuration
Pour entrer dans l’éditeur de configuration et la modifier (en mode privilégié) : 
# configure  terminal
Enter configuration commands, one per line.  End with CNTL/Z.
pour sortir de l’éditeur de configuration : utiliser la combinaison de touches CTRL-Z : 
switch(config)# ^Z
switch#
Assigne un mot de passe encrypté à enable : 
# enable secret <password>
pour visualiser la configuration en mémoire non volatile (celle en mémoire RAM, pas celle sur la 
mémoire flash) : 
# write terminal 
# show running-config 
pour visualiser la configuration de démarrage (celle stockée sur la mémoire flash) :
# show startup-config 
Modifier le nom de l'équipement réseau  : 
# hostname <hostname> 
Pour sauvegarder la nouvelle configuration en mémoire non volatile (FLASH) : 
# write memory 
# copy running-config startup-config
 Toutes les commandes peuvent être rentrées sous forme complète ou sous forme abrégée :
switch# write memory
switch# wr mem
switch# conf t
switch(config)# int fa 0/1
switch(config)# interface fastethernet 0/1
switch(config-if)# ^Z
Configuration des interfaces
Le mode de chaque interface physique : 
switchport mode access : dans ce mode, on indique au commutateur que le port n'est pas 
un uplink et que l'on connectera une ou plusieurs machines dans le même VLAN
switchport access vlan N : dans ce mode, on indique au commutateur que le port n'est 
pas un uplink et que l'on connectera une ou plusieurs machines dans le VLAN numéro N
switchport mode trunk : dans ce mode, on indique au commutateur que le port va 
transporter tous les VLAN (sauf prunning)
switchport trunk encap dot1q : dans ce mode, on indique au commutateur que le 
protocole de trunking va utiliser le protocole normalisé 802.1q
switchport trunk encap isl : dans ce mode, on indique au commutateur que le 
Denis Pugnère (CNRS / IN2P3 / IPNL) Page 4 / 10

