---
id: collect-261001-cisco/cisco/pugnere-presentations-administration-reseau-config-cisco-pdf-302726c4-2
title: "show int"
domain: cisco
role: reference
task: reference
actors: []
dates: ["2010-10-06"]
keywords: ["attention", "dram"]
source: docs/RAG/collect-261001-cisco/pugnere-presentations-administration-reseau-config-cisco-pdf-302726c4.md
source_anchor: ""
source_lines: [162, 292]
sha256: 22e9644e532e9deee9d023361df0bcff0218b9f4e9312ca0c7a74c7091373c47
---

# show int

Tutoriel : Configuration d’un commutateur ou d'un routeur CISCO                            2010-10-06/90
protocole de trunking va utiliser le protocole ISL (propriétaire Cisco) : non conseillé
Exemple de configuration de l'interface GigabitEthernet 0/23 dans le vlan 13 et GigabitEthernet 
0/24 dans en mode trunk : 
switch# conf t
switch(config)#interface GigabitEthernet0/23
switch(config-if)# switchport access vlan 13
switch(config-if)# switchport mode access
switch(config-if)# spanning-tree portfast
switch(config-if)#interface GigabitEthernet0/24
switch(config-if)# switchport trunk encapsulation dot1q
switch(config-if)# switchport mode trunk
switch(config-if)# ^Z
Attribuer une adresse IP à une interface : ip address <address> <mask> : 
switch# conf t
switch(config)# int vlan1
switch(config-if)# ip address 192.168.1.20 255.255.255.0
switch(config-if)# ^Z
Active ou désactive une interface (ou un VLAN) : shutdown ou no shutdown : 
switch# conf t
switch(config)# int fastethernet 0/1
switch(config-if)# no shutdown
switch(config-if)#^Z
Active ou désactive l'auto-croisement d'une interface : mdix
switch# conf t
switch(config)# int fastethernet 0/1
switch(config-if)# mdix auto
switch(config-if)# ^Z
Commandes par interfaces (sous-commandes) : Elles s’adressent à une partie du commutateur. À 1 
ou plusieurs interfaces : 
#conf t
Enter configuration commands, one per line.  End with CNTL/Z.
routeur(config)#interface fastethernet 1/0/1
routeur(config-if)# ip address x.y.z.y 255.255.255.0
routeur(config)#interface fastethernet 1/0/2
routeur(config-if)# ip address x.y.z.z 255.255.255.0
routeur(config)#interface range fastethernet 1/0/2 - 24
routeur(config-if)# switchport mode acess
routeur(config-if)# switchport acess vlan 3
routeur(config-if)# ^Z
Accès au système de fichier de la mémoire FLASH
Affichage du contenu de la mémoire flash : 
sw4#dir flash: 
Directory of flash:/ 
    2  -rwx         270   Jan 1 1970 00:01:36 +00:00  env_vars 
    3  -rwx     3036020   Mar 1 1993 00:03:25 +00:00  c2950-i6q4l2-mz.121-
20.EA1a.bin
    4  -rwx        1368   Mar 1 1993 00:14:18 +00:00  config.text 
    5  -rwx        1996   Mar 1 1993 00:31:25 +00:00  vlan.dat 
Denis Pugnère (CNRS / IN2P3 / IPNL) Page 5 / 10

Tutoriel : Configuration d’un commutateur ou d'un routeur CISCO                            2010-10-06/90
    6  -rwx           5   Mar 1 1993 00:14:18 +00:00  private-config.text
    7  -rwx         110   Mar 1 1993 00:01:45 +00:00  info
    8  drwx        2688   Mar 1 1993 00:07:05 +00:00  html 
   90  -rwx         110   Mar 1 1993 00:07:50 +00:00  info.ver 
7741440 bytes total (1606144 bytes free) 
sw4#
Suppression de la base des VLAN (ATTENTION : efface la liste des VLAN reçus ou enregistrés) :  
#delete flash:vlan.dat 
Delete filename [vlan.dat]? 
Delete flash:vlan.dat? [confirm]
Affichage d'un fichier contenu sur la flash : 
#more flash:info.ver 
image_name: c2950-i6q4l2-mz.121-20.EA1a.bin 
image_file_size: 3041280 
image_min_dram: 16 
tar_file_size_k: 3133
Suppression de la configuration
Pour effacer la configuration de mémoire non volatile (FLASH) : ATTENTION : Cela efface la 
configuration qui est chargée au démarrage, si on re-démarre le commutateur après cette 
commande, il perd toute sa configuration.
# erase startup-config
Supprimer le fichier (vlan.dat) contenant la liste des VLAN enregistrés localement sur le switch : 
# delete flash:vlan.dat
Rédémarrage du routeur (reboot)
# reload 
Une autre méthode consiste à ré-initialiser le switch à sa configuration par défaut d'usine, pour 
cela : 
– couper l'alimentation du switch
– tout en maintenant appuyé le bouton MODE, brancher l'alimentation du switch,
– le switch démarre en mode « maintenance »
Taper :
Switch: flash_init
Le switch donne alors accès au contenu de la flash, il est alors possible de modifier le contenu de la 
flash, le fichier contenant la configuration courante s'appelle config.text, il est possible de le 
renommer ou de le supprimer :
Switch: rename flash:config.text flash:config.ancien
Switch: delete flash:config.text
Denis Pugnère (CNRS / IN2P3 / IPNL) Page 6 / 10

Tutoriel : Configuration d’un commutateur ou d'un routeur CISCO                            2010-10-06/90
Idem pour supprimer la liste des VLAN enregistrés localement sur le switch (fichier vlan.dat) il est 
possible de le renommer ou de le supprimer :
Switch: delete flash:vlan.dat
Puis redémarrer le switch
Switch: reset
Sauvegarde de la configuration du routeur sur un serveur
On peut sauvegarder la configuration du routeur sur un serveur du réseau via TFTP, RCP ou FTP ou 
HTTP... (en fonction de ce que supporte le commutateur). La commande copy demande 2 
paramètres : le premier est la source et le second est la destination.
La référence à un fichier sur un serveur TFTP se fait par la syntaxe suivante (spécifier l'adresse IP 
du serveur et le nom du fichier) : 
tftp://serveur/fichier
La référence à un fichier sur un serveur FTP se fait par la syntaxe suivante (spécifier l'utilisateur à la 
place de 'user', le mot de passe de l'utilisateur dans 'password', puis l'adresse IP du serveur et le nom 
du fichier) : 
ftp://user:password@serveur/fichier
Il faut d’abord mettre en place un serveur TFTP. On suppose que /tftpboot est le répertoire de 
chargement du serveur tftp : 
# copy system:/running-config tftp://192.168.1.14/test
Address or name of remote host [192.168.1.14]? 
Destination filename [test]? 
!!!!!!!!
35761 bytes copied in 1.157 secs (30908 bytes/sec)
#
La configuration est sauvegardée dans /tftpboot/test du serveur tftp.
Chargement de la configuration à partir d’un serveur TFTP
De même on peut charger la configuration via un serveur TFTP ou ftp. Cette méthode présente 
l’avantage de pouvoir écrire tranquillement sa configuration via un éditeur de texte et la charger 
quand on veut.
# copy tftp://192.168.1.14/test system:/running-config
# copy ftp://user:password@192.168.1.14/test system:/running-config
# copy system:/running-config system:/startup-config
# copy ftp://user:password@192.168.1.14/test system:/startup-config
On peut donc charger un fichier de configuration et soit le rendre actif immédiatement (destination 
= system:/running-config) ou au prochain démarrage (destination = system:/startup-config)
Commandes de tests et de visualisation de l’état du routeur : 
 
Denis Pugnère (CNRS / IN2P3 / IPNL) Page 7 / 10

