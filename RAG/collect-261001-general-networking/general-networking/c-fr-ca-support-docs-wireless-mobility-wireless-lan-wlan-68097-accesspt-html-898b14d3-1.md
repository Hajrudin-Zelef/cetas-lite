---
id: collect-261001-general-networking/general-networking/c-fr-ca-support-docs-wireless-mobility-wireless-lan-wlan-68097-accesspt-html-898b14d3-1
title: "c-fr-ca-support-docs-wireless-mobility-wireless-lan-wlan-68097-accesspt-html-898b14d3"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-general-networking/c-fr-ca-support-docs-wireless-mobility-wireless-lan-wlan-68097-accesspt-html-898b14d3.md
source_anchor: ""
source_lines: [1, 59]
sha256: b879f7ced1c33df7a165c56c0ecaeb800750684cfc90a0925c4d8e35aa499f24
---

# c-fr-ca-support-docs-wireless-mobility-wireless-lan-wlan-68097-accesspt-html-898b14d3

Ce document explique comment configurer les filtres basés sur des listes de contrôle d'accès (ACL) sur des points d'accès Cisco (AP) Aironet à l'aide de l'interface de ligne de commande (CLI).
Cisco vous recommande de prendre connaissance des rubriques suivantes :
Configuration d'une connexion sans fil à l'aide d'un AP Aironet et d'un adaptateur client Aironet 802.11 a/b/g
ACL
Les informations contenues dans ce document sont basées sur les versions de matériel et de logiciel suivantes :
Point d'accès (AP) de la gamme Aironet 1200 qui exécute le logiciel Cisco IOS® Version 12.3(7)JA1
Adaptateur client Aironet 802.11a/b/g
Aironet Desktop Utility (ADU), version 2.5
The information in this document was created from the devices in a specific lab environment. All of the devices used in this document started with a cleared (default) configuration. If your network is live, make sure that you understand the potential impact of any command.
Pour plus d'informations sur les conventions utilisées dans ce document, reportez-vous à Conventions relatives aux conseils techniques Cisco.
Vous pouvez utiliser des filtres sur l'application pour effectuer ces tâches :
Restreindre l'accès au réseau sans fil LAN (WLAN)
Fournir une couche supplémentaire de sécurité sans fil
Vous pouvez utiliser différents types de filtres pour filtrer le trafic en fonction :
de protocoles spécifiques ;
de l'adresse MAC du périphérique client ;
de l'adresse IP du périphérique client.
Vous pouvez également permettre à des filtres de restreindre le trafic depuis des utilisateurs sur le réseau local câblé. Les filtres d'adresse IP et d'adresse MAC permettent ou rejettent le transfert des paquets de monodiffusion et de multidiffusion qui sont envoyés vers ou depuis des adresses IP ou MAC spécifiques.
Les filtres basés sur des protocoles fournissent une façon plus précise de restreindre l'accès aux protocoles spécifiques par les interfaces Ethernet et radios de l'AP. Vous pouvez utiliser l'une ou l'autre de ces méthodes pour configurer les filtres sur les AP :
GUI Web
CLI
Ce document explique comment utiliser les filtres de liste de contrôle d'accès par la CLI. Pour obtenir des informations sur la façon de configurer des filtres par l'interface graphique, consultez Configuration des filtres.
Vous pouvez utiliser la CLI pour configurer ces types de filtres basés sur ACL sur l'AP :
Filtres qui utilisent des listes de contrôle d'accès standard
Filtres qui utilisent des listes de contrôle d'accès étendues
Filtres qui utilisent des listes de contrôle d'accès d'adresse MAC
Remarque : Le nombre d'entrées autorisées sur une liste de contrôle d'accès est limité par le CPU du point d'accès. S'il faut ajouter un grand nombre d'entrées de routage à une ACL, par exemple en filtrant une liste d'adresses MAC pour les clients, utilisez un commutateur dans le réseau qui peut effectuer la tâche.
Cette section vous fournit des informations pour configurer les fonctionnalités décrites dans ce document.
Utilisez l'outil Command Lookup Tool (clients enregistrésseulement) pour trouver plus d'informations sur les commandes utilisées dans ce document.
Toutes les configurations dans ce document supposent qu'une connexion sans fil est déjà établie. Ce document se focalise seulement sur la façon d'utiliser la CLI afin de configurer des filtres. Si vous n'avez pas une connexion sans fil de base, consultez la section Exemple de configuration de connexion LAN sans fil de base.
Vous pouvez utiliser des listes de contrôle d'accès standard pour autoriser ou rejeter l'entrée de périphériques clients dans le réseau WLAN basé sur l'adresse IP du client. Les listes de contrôle d'accès standard comparent l'adresse source des paquets IP aux adresses qui sont configurées dans la liste de contrôle d'accès afin de contrôler le trafic. Ce type de liste de contrôle d'accès peut être désigné comme étant basé sur l'adresse IP source.
La syntaxe des commandes applicables aux listes de contrôle d'accès standard est la suivante : access-list access-list-number {permit | deny} {adresse IP hôte | source-ip source-wildcard | any}.
Dans la version 12.3(7)JA de Cisco IOS®, le numéro de la liste de contrôle d'accès peut être n'importe quel numéro de 1 à 99. Les listes de contrôle d'accès standard peuvent également utiliser une plage étendue de 1300 à 1999. Ces numéros supplémentaires sont des listes de contrôle d'accès IP étendues.
Quand une liste de contrôle d'accès standard est configurée pour refuser l'accès à un client, le client s'associe toujours à l'AP. Cependant, il n'y a aucune communication de données entre l'AP et le client.
Cet exemple montre une liste de contrôle d'accès standard qui est configurée pour filtrer l'adresse IP 10.0.0.2 du client depuis l'interface sans fil (interface radio0). L'adresse IP de l'AP est 10.0.0.1.
Après cela, le client avec l'adresse IP 10.0.0.2 ne peut pas envoyer ou recevoir de données par le réseau WLAN même s'il est associé à l'AP.
Effectuez ces étapes afin de créer une liste de contrôle d'accès standard par la CLI :
Connectez-vous à l'AP par la CLI.
Utilisez le port de console ou Telnet afin d'accéder à l'ACL par l'interface Ethernet ou l'interface sans fil.
Passez en mode de configuration globale sur l'AP :
 
      AP#configure terminal
Exécutez ces commandes afin de créer la liste de contrôle d'accès standard :
 
      AP<config>#access-list 25 deny host 10.0.0.2 !--- Create a standard ACL 25 to deny access to the !--- client with IP address 10.0.0.2. AP<config>#access-list 25 permit any !--- Allow all other hosts to access the network.
Exécutez ces commandes afin d'appliquer cette ACL à l'interface radio :
 
      AP<config>#interface Dot11Radio 0 AP<config-if>#ip access-group 25 in !--- Apply the standard ACL to the radio interface 0.
Vous pouvez également créer une ACL standard nommée (NACL). La NACL emploie un nom au lieu d'un numéro pour définir l'ACL.
 
    AP#configure terminal AP<config>#ip access-list standard name AP<config>#permit | deny {host ip-address | source-ip [source-wildcard] | any} log
Exécutez ces commandes afin d'utiliser des NACL standard pour refuser l'accès de l'hôte 10.0.0.2 au réseau WLAN :
 
    AP#configure terminal AP<config>#ip access-list standard TEST !--- Create a standard NACL TEST. AP<config-std-nacl>#deny host 10.0.0.2 !--- Disallow the client with IP address 10.0.0.2 !--- access to the network. AP<config-std-nacl>#permit any !--- Allow all other hosts to access the network. AP<config-std-nacl>#exit !--- Exit to global configuration mode. AP<config>#interface Dot11Radio 0 !--- Enter dot11 radio0 interface mode. AP<config-if>#ip access-group TEST in !--- Apply the standard NACL to the radio interface.
Les listes de contrôle d'accès étendues comparent les adresses source et de destination des paquets IP aux adresses configurées dans la liste de contrôle d'accès pour contrôler le trafic. Les listes de contrôle d'accès étendues fournissent également un moyen de filtrer le trafic en fonction de protocoles de routage spécifiques. Ceci fournit un contrôle plus précis pour l'implémentation des filtres sur un réseau WLAN.
Les listes de contrôle d'accès étendues permettent à un client d'accéder à certaines ressources sur le réseau mais pas à toutes. Par exemple, vous pouvez implémenter un filtre qui autorise le trafic DHCP et Telnet au client tandis qu'il restreint tout autre trafic.
Voici la syntaxe de commande des listes de contrôle d'accès étendues :
Remarque : cette commande est encapsulée sur quatre lignes pour des raisons d'espace.
 
