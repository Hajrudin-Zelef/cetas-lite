---
id: collect-260926-mikrotik/mikrotik/routeur-mikrotik-rb5009ug-s-in-sfp-gpon-onu-34-20bi-1
title: "ONT-G"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/routeur-mikrotik-rb5009ug-s-in-sfp-gpon-onu-34-20bi.md
source_anchor: ""
source_lines: [1, 149]
sha256: 7cf2b807c9748021c5df3dea72c7a24f3f97ab37a9560d9cfc0c450956f25d4b
---

# ONT-G

Bonjour, 

Le routeur MikroTik RB5009UG+S+IN peut servir à remplacer le livebox, je propose de créer un thread dédié.

Ce routeur n'est pas aussi avancé que son grand frère le CCR2004, mais est largement plus accessibles pour nos bourses. 

Il est également plus évolué que le HeX S. 

La complexité du CCR2004 fait que la configuration proposée dans le thread original 

https://lafibre.info/remplacer-livebox/guide-de-connexion-fibre-directement-sur-un-routeur-voire-meme-en-2gbps/
 est loin de coller directement à ce routeur, elle m'a à la fois aidé, mais par la suite découragé la première fois que j'ai essayé de configurer mon RB5009UG. 

Je vais ici détailler ce que j'ai fait pour arriver à faire marcher l'ipv4, et on verra où cela nous mènera. 

J'ai comme décrit sur la page du CCR2004, acheté un ONU acheté chez FS.com. 

Réf. de l'ONU : SFP GPON-ONU-34-20BI (

https://www.fs.com/fr/products/183843.html?attribute=46672&id=3208708
)

**PART 1 : Quelques notes et configuration basique**
On se prépare psychologiquement, et on sauvegarde quelque part les élements suivant que l'on trouve notés sur l'arrière de sa livebox, dans les menus admin, ou dans la doc reçu avec la livebox.

- Serial number : SMBSXXXXXXXX

- Hardware version : SMBSSXXXXXXX

- Login/password pour la connexion.

On branche le routeur, cable rj45 sur le port ether8. 

On accède à sa console d'admin à partir de son navigateur, à partir de l'adresse : 192.168.88.1

On change le mot de passe, on met à jour son routeur et on effectue les modifications suivantes : 

On débranche/rebranche son cable réseau.

**PART 2 : Configuration de l'ONU**
1/ sur le GPON après un ssh sur 192.168.1.10

[Conditions : 

- Faut avoir la fibre connectée sur le module 

- Si vous êtes sur votre réseau en 192.168.1.255 , alors vérifier bien que vous n'avez rien en .10]

`ssh -o KexAlgorithms=diffie-hellman-group14-sha1 -oHostKeyAlgorithms=+ssh-dss  192.168.1.10 -l ONTUSER`
Mot de passe 

`7sp!lwUBz1`
2/ on rentre le numéro de série précédement noté 

`set_serial_number SMBSXXXXXXXX`
3/ on rentre le vendor_ID 

`sfp_i2c -i 7 -s “ SMBS”`
4/ pour ma part, j'ai dû mettre le hardware version de l'ONT. 

```
cp /etc/mibs/data_1g_8q.ini /etc/mibs/data_1g_8q.ini.ORG
vi /etc/mibs/data_1g_8q.ini
```
rempalacer la ligne ci-dessous avec le hardware version de l'ONT =>

# ONT-G

```
256 0 HWTC 0000000000000 00000000 2 0 0 0 0 #0
```
PAR

```
256 0 SMBS SMBSSXXXXXXX\0 00000000 2 0 0 0 0 #0
```
5/ un petit reboot

`reboot`
6/ On se reconnecte, puis on tape cette commande dont le résultat devrait faire apparaitre 

**"curr_state=5 previous_state=4"**`onu ploamsg`**PART 3 : Configuration IPv4**
Je désactive les interfaces que je ne compte pas utiliser.

De ether2 à ether7. 

Pour le moment on est branché à ether8, on utilisera ether8.

L'idéal étant d'utiliser ether1 qui est le port en 2,5Gb/s

```
/interface ethernet
set [ find default-name=ether2 ] disabled=yes
set [ find default-name=ether3 ] disabled=yes
set [ find default-name=ether4 ] disabled=yes
set [ find default-name=ether5 ] disabled=yes
set [ find default-name=ether6 ] disabled=yes
set [ find default-name=ether7 ] disabled=yes
```
On met l'adresse mac de sa livebox (adaptez bien cette ligne... n'allez pas la recopier)

```
/interface bridge
add admin-mac=20:XX:XX:XX:XX:XX auto-mac=no name=bridge
```
VLAN 832 pour Internet

```
/interface vlan
add comment="Internet ONT" interface=sfp-sfpplus1 loop-protect-disable-time=0s \
    loop-protect-send-interval=1s name=vlan832-internet vlan-id=832
```
Set des options du client DHCP, adaptez les 140 caractères YYYYYYYY, comme suivant :

1- allez sur 

https://jsfiddle.net/kgersen/3mnsc6wy/
, entrez le login/pass et copiez la sortie simplement.

2- dans votre éditeur préféré : ajouter 0x au début, et supprimez les deux points ":" partout.

```
/ip dhcp-client option
add code=60 name=vendor-class-identifier value=0x736167656d
add code=77 name=userclass value="0x2b46535644534c5f6c697665626f782e496e7465726e65742e736f66746174686f6d652e4c697665626f7833"
add code=90 name=authsend  value=0xYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY
```
```
/ip dhcp-server
add address-pool=dhcp interface=bridge lease-time=1w name=LAN
```
```
/interface bridge filter
add action=set-priority chain=output dst-port=67 ip-protocol=udp log=yes \
    log-prefix="Set CoS6 on DHCP request" mac-protocol=ip new-priority=6 \
    out-interface=vlan832-internet passthrough=yes
```
```
/ip dhcp-client
add dhcp-options=hostname,clientid,authsend,userclass,vendor-class-identifier \
    interface=bridge
```
```
/ip dhcp-server network
add address=192.168.1.0/24 comment=defconf dns-server=1.1.1.1,1.0.0.1 \
    gateway=192.168.1.1 netmask=24
```
Règles firewall, et je suis totalement dépassé ! 

