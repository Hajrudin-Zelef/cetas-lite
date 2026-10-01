---
id: collect-261001-general-networking/general-networking/c-fr-ca-support-docs-wireless-mobility-wireless-lan-wlan-68097-accesspt-html-898b14d3-2
title: "c-fr-ca-support-docs-wireless-mobility-wireless-lan-wlan-68097-accesspt-html-898b14d3"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-general-networking/c-fr-ca-support-docs-wireless-mobility-wireless-lan-wlan-68097-accesspt-html-898b14d3.md
source_anchor: ""
source_lines: [60, 123]
sha256: 81746eace807de39951981e730cda56b554fb3173a8f9ae71cdf5327ba04bdb2
---

# c-fr-ca-support-docs-wireless-mobility-wireless-lan-wlan-68097-accesspt-html-898b14d3

    access-list access-list-number [dynamic dynamic-name [timeout minutes]] {deny | permit} protocol source source-wildcard destination destination-wildcard [precedence precedence] [tos tos] [log | log-input] [time-range time-range-name]
Dans la version du logiciel Cisco IOS 12.3(7)JA, les listes de contrôle d'accès étendues peuvent utiliser des numéros dans la plage 100 à 199. Les listes de contrôle d'accès étendues peuvent également utiliser des numéros dans la plage de 2000 à 2699. C'est la plage étendue des listes de contrôle d'accès étendues.
Remarque : Le mot clé log à la fin des entrées de liste de contrôle d'accès individuelles indique :
le numéro et le nom de l'ACL ;
si le paquet a été autorisé ou refusé ;
les informations spécifiques au port.
Les listes de contrôle d'accès étendues peuvent également utiliser des noms au lieu des numéros. C'est la syntaxe pour créer des NACL étendues :
 
    ip access-list extended name {deny | permit} protocol source source-wildcard destination destination-wildcard [precedence precedence] [tos tos] [log | log-input] [time-range time-range-name]
Cet exemple de configuration utilise des NACL étendues. La condition requise est que la NACL étendue doit permettre l'accès Telnet aux clients. Vous devez restreindre tous les autres protocoles sur le réseau WLAN. En outre, les clients utilisent DHCP afin d'obtenir l'adresse IP. Vous devez créer une liste de contrôle d'accès étendue qui :
permet le trafic DHCP et Telnet ;
refuse tous les autres types de trafic.
Une fois que cette liste de contrôle d'accès étendue est appliquée à l'interface radio, les clients s'associent à l'AP et obtiennent une adresse IP du serveur DHCP. Les clients peuvent également utiliser Telnet. Tous les autres types de trafic sont refusés.
Effectuez ces étapes afin de créer une liste de contrôle d'accès étendue sur l'AP :
Exécutez ces commandes afin de créer la liste de contrôle d'accès étendue :
 
      AP<config>#ip access-list extended Allow_DHCP_Telnet !--- Create an extended ACL Allow_DHCP_Telnet. AP<config-extd-nacl>#permit tcp any any eq telnet !--- Allow Telnet traffic. AP<config-extd-nacl>#permit udp any any eq bootpc !--- Allow DHCP traffic. AP<config-extd-nacl>#permit udp any any eq bootps !--- Allow DHCP traffic. AP<config-extd-nacl>#deny ip any any !--- Deny all other traffic types. AP<config-extd-nacl>#exit !--- Return to global configuration mode.
Exécutez ces commandes afin d'appliquer l'ACL à l'interface radio :
 
      AP<config>#interface Dot11Radio 0 AP<config-if>#ip access-group Allow_DHCP_Telnet in !--- Apply the extended ACL Allow_DHCP_Telnet !--- to the radio0 interface.
Vous pouvez utiliser des filtres basés sur l'adresse MAC afin de filtrer des périphériques clients basés sur l'adresse MAC encodée. Quand un client se voit refuser l'accès par un filtre basé sur l'adresse MAC, il ne peut pas s'associer à l'AP. Les filtres d'adresse MAC permettent ou rejettent le transfert des paquets de monodiffusion et de multidiffusion qui sont envoyés depuis ou adressés à des adresses MAC spécifiques.
Voici la syntaxe de commande pour créer une ACL basée sur l'adresse MAC sur l'AP :
Remarque : Cette commande a été encapsulée sur deux lignes en raison de considérations d'espace.
 
    access-list access-list-number {permit | deny} 48-bit-hardware-address 48-bit-hardware-address-mask
Dans la version du logiciel Cisco IOS 12.3(7)JA, les listes de contrôle d'accès d'adresse MAC peuvent utiliser des numéros dans une plage de 700 à 799 comme numéro d'ACL. Elles peuvent également utiliser des numéros dans une plage de 1100 à 1199.
Cet exemple montre comment configurer un filtre basé sur l'adresse MAC par la CLI, afin de filtrer le client avec une adresse MAC 0040.96a5.b5d4 :
Passez en mode de configuration globale sur la CLI de l'AP :
Créez une ACL d'adresse MAC numéro 700.
Cette ACL ne permet pas au client 0040.96a5.b5d4 de s'associer à l'AP.
 
      access-list 700 deny 0040.96a5.b5d4 0000.0000.0000 !--- This ACL denies all traffic to and from !--- the client with MAC address 0040.96a5.b5d4.
Exécutez cette commande afin d'appliquer cette ACL basée sur l'adresse MAC à l'interface radio :
 
      dot11 association mac-list 700 !--- Apply the MAC-based ACL.
Après avoir configuré ce filtre sur l'AP, le client avec cette adresse MAC, qui a été précédemment associée à l'AP, est dissocié. La console de l'AP envoie ce message :
 
    AccessPoint# *Mar 1 01:42:36.743: %DOT11-6-DISASSOC: Interface Dot11Radio0, Deauthenticating Station 0040.96a5.b5d4
Les listes de contrôle d'accès basées sur l'heure peuvent être activées ou désactivées pour une période de temps spécifique. Cette fonctionnalité offre la robustesse et la souplesse permettant de définir des stratégies de contrôle d'accès qui autorisent ou refusent certains types de trafic.
Cet exemple montre comment configurer une ACL basée sur le temps par la CLI, où la connexion Telnet est autorisée depuis le réseau vers l'extérieur en semaine et pendant les heures ouvrables :
Remarque : Une liste de contrôle d'accès basée sur le temps peut être définie soit sur le port Fast Ethernet, soit sur le port Radio du point d'accès Aironet, en fonction de vos besoins. Elle n'est jamais appliquée sur le Bridge Group Virtual Interface (BVI).
Créez une plage horaire. Pour cela, exécutez cette commande en mode de configuration globale :
 
      AP<config>#time-range Test !--- Create a time-range with name Test. AP(config-time-range)# periodic weekdays 7:00 to 19:00 !--- Allows access to users during weekdays from 7:00 to 19:00 hrs.
Créez une ACL 101 :
 
      AP<config># ip access-list extended 101 AP<config-ext-nacl>#permit tcp 10.1.1.0 0.0.0.255 172.16.1.0 0.0.0.255 eq telnet time-range Test !--- This ACL permits Telnet traffic to and from !--- the network for the specified time-range Test.
Cette ACL autorise une session Telnet à l'AP en semaine.
Exécutez cette commande afin d'appliquer cette ACL basée sur le temps à l'interface Ethernet :
 
      interface Ethernet0/0 ip address 10.1.1.1 255.255.255.0 ip access-group 101 in !--- Apply the time-based ACL.
Aucune procédure de vérification n'est disponible pour cette configuration.
Utilisez cette section pour dépanner votre configuration.
Effectuez ces étapes afin de supprimer une ACL d'une interface :
Passez en mode de configuration de l'interface.
Entrez no devant la commande ip access-group, comme le montre cet exemple :
 
      interface interface no ip access-group {access-list-name | access-list-number} {in | out}
Vous pouvez également utiliser la commande show access-list name | numéro afin de dépanner votre configuration. La commande show ip access-list fournit un nombre de paquets qui indique l'entrée de la liste de contrôle d'accès consultée.
Évitez d'utiliser à la fois la CLI et les interfaces de navigateur Web pour configurer le périphérique sans fil. Si vous configurez le périphérique sans fil avec la CLI, l'interface de navigateur Web peut afficher une traduction inexacte de la configuration. Cependant, cette inexactitude ne signifie pas nécessairement que le périphérique sans fil est mal configuré. Par exemple, si vous configurez les ACL avec la CLI, l'interface de navigateur Web peut afficher ce message :
Si vous voyez ce message, utilisez la CLI afin de supprimer les ACL et utilisez l'interface de navigateur Web pour les reconfigurer.
| Révision | Date de publication | Commentaires | 
|---|---|---|
| 1.0 |                                                                                               17-Jul-2006                                                                                       | Première publication |
