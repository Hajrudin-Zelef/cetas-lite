---
id: collect-261001-cisco/cisco/c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799-1
title: "c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799.md
source_anchor: ""
source_lines: [1, 147]
sha256: 733d56e6124b63d74be688eeba254096015bc450e0d36e8af15f19751dcaf19b
---

# c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799

Configuration des listes d'accès IP
Table des matières
Introduction
Conditions préalables
Exigences
Composants utilisés
Conventions
Informations générales
Concepts relatifs aux listes de contrôle d'accès
Masques
Récapitulation des listes de contrôle d'accès
Traiter les listes de contrôle d'accès
Définir les ports et les types de messages
Appliquer les listes de contrôle d'accès
Définir les termes interne, externe, entrant, sortant, source et destination
Modifier les listes de contrôle d'accès
Dépannage
Types de listes de contrôle d'accès IP
Diagramme du réseau
ACL standards
ACL étendus
IP
ICMP
TCP
UDP
Verrou et clé (listes de contrôle d'accès dynamiques)
Listes de contrôle d'accès nommées IP
Listes de contrôle d'accès réflexives
Listes de contrôle d'accès basées sur l'heure utilisant des plages temporelles
Entrées de liste de contrôle d'accès IP commentées
Contrôle d'accès basé sur contexte
Proxy d'authentification
Listes de contrôle d'accès turbo
Listes de contrôle d'accès basées sur l'heure distribuées
Listes de contrôle d'accès de réception
Listes de contrôle d'accès de protection d'infrastructure
Listes de contrôle d'accès de transit
Informations connexes
Introduction

Ce document décrit différents types de listes de contrôle d'accès IP (ACL) et comment elles 
peuvent filtrer le trafic réseau.
Conditions préalables
Exigences
Aucune condition préalable spécifique n'est requise pour ce document.
Composants utilisés
Ce document traite de divers types de listes de contrôle d'accès. Certaines de ces versions 
existent depuis la version 8.3 du logiciel Cisco IOS® et d'autres ont été introduites dans des 
versions ultérieures. Ceci est noté dans la discussion de chaque type.
The information in this document was created from the devices in a specific lab environment. All of 
the devices used in this document started with a cleared (default) configuration. Si votre réseau 
est en ligne, assurez-vous de bien comprendre l’incidence possible des commandes.
Conventions
Reportez-vous aux conventions des conseils techniques Cisco pour plus d’information sur les 
conventions utilisées dans ce document.
Informations générales
Ce document décrit comment les listes de contrôle d'accès (ACL) IP peuvent filtrer le trafic sur le 
réseau. Il contient également de brèves descriptions des types de listes de contrôle d'accès IP, de 
la disponibilité des fonctionnalités et un exemple d'utilisation sur un réseau.
Remarque : Le document RFC 1700 contient les numéros attribués des ports réservés.Le 
document RFC 1918 contient l’allocation d’adresses pour les réseaux Internet privés, des 
adresses IP qui ne doivent normalement pas être vues sur Internet.
Remarque : Seuls les utilisateurs Cisco enregistrés peuvent accéder aux informations 
internes.
Remarque : Les listes de contrôle d’accès peuvent également être utilisées pour définir le 
trafic vers les protocoles NAT (Network Address Translate), crypter ou filtrer des protocoles 
non IP tels qu’AppleTalk ou IPX. Une discussion relative à ces fonctions sort du cadre de ce 
document.

Concepts relatifs aux listes de contrôle d'accès
Masques
Les masques sont utilisés dans les listes de contrôle d’accès IP pour définir les parties d’une 
adresse IP qui doivent correspondre et celles qui peuvent varier. Lorsque vous configurez des 
adresses IP sur des interfaces, vous utilisez un masque de sous-réseau, qui commence par des 
valeurs plus grandes (255) à gauche, par exemple, l’adresse IP 10.1.1.129 avec le masque de 
sous-réseau 255.255.255.0. En revanche, les listes de contrôle d’accès utilisent un masque 
générique (également appelé masque inverse), qui fonctionne de la manière opposée. Par 
exemple, le masque générique 0.0.0.255 signifie que les trois premiers octets doivent 
correspondre exactement, tandis que le dernier octet peut varier. En termes binaires, chaque 0 du 
masque générique signifie que ce bit doit correspondre exactement et chaque 1 signifie que ce bit 
peut être ignoré (peu importe). Le tableau suivant illustre ce concept plus en détail.
Exemple de masque
adresse réseau (trafic à traiter) 10.1.1.0
subnet mask (masque de sous-réseau) 255.255.255.0
masque inversé 0.0.0.255
adresse réseau (binaire) 00001010.00000001.00000001.00000000
masque inverse (binaire) 00000000.00000000.00000000.11111111
D’après le masque inverse binaire, vous pouvez voir que les trois premiers octets doivent 
correspondre exactement à l’adresse réseau binaire donnée (00001010.00000001.00000001). Le 
dernier ensemble de numéros est « ne vous en souciez pas » (.11111111). Par conséquent, tout 
le trafic qui commence par 10.1.1. correspond depuis le dernier octet est « ne s’en soucient pas ». 
Par conséquent, les adresses réseau 10.1.1.1 à 10.1.1.255 (10.1.1.x) sont traitées.
Pour déterminer le masque inverse de la liste de contrôle d’accès, vous devez soustraire le 
masque de sous-réseau de 255.255.255.255. Dans cet exemple, le masque inverse est déterminé 
pour l’adresse réseau 172.16.1.0 avec le masque de sous-réseau 255.255.255.0.
255.255.255.255 - 255.255.255.0 (masque de sous-réseau) = 0.0.0.255 (masque inverse)•
Notez les équivalents ACL.
Le caractère générique/source de 0.0.0.0/255.255.255.255 signifie any.•
Le masque générique/source de 10.1.1.2/0.0.0.0 est identique à l’hôte 10.1.1.2.•
Récapitulation des listes de contrôle d'accès
Remarque : les masques de sous-réseau peuvent également être représentés comme 
notation de longueur fixe. Par exemple, 192.168.10.0/24 représente 192.168.10.0 
255.255.255.0.

Cette liste décrit comment récapituler une plage de réseaux en un seul réseau pour l'optimisation 
des listes de contrôle d'accès. Examinez les réseaux ci-dessous.
 
192.168.32.0/24 
192.168.33.0/24 
192.168.34.0/24 
192.168.35.0/24 
192.168.36.0/24 
192.168.37.0/24 
192.168.38.0/24 
192.168.39.0/24
 
Les deux premiers octets et le dernier octet sont les mêmes pour chaque réseau. Le tableau ci-
dessous explique comment les récapituler en un seul réseau.
Le troisième octet des réseaux précédents peut être écrit comme indiqué dans ce tableau, 
correspondant à la position de bit d’octet et à la valeur d’adresse pour chaque bit.
Décimal 128 64 32 16 8 4 2 1
32 0 0 1 0 0 0 0 0
33 0 0 1 0 0 0 0 1
34 0 0 1 0 0 0 1 0
35 0 0 1 0 0 0 1 1
36 0 0 1 0 0 1 0 0
37 0 0 1 0 0 1 0 1
38 0 0 1 0 0 1 1 0
39 0 0 1 0 0 1 1 1
L L L L L D D D
Puisque les cinq premiers bits correspondent, les huit réseaux précédents peuvent être 
récapitulés en un réseau (192.168.32.0/21 ou 192.168.32.0 255.255.248.0). Chacune des huit 
combinaisons possibles des trois bits de poids faible est appropriée pour les plages de réseaux en 
question. Cette commande définit une liste de contrôle d'accès qui autorise ce réseau. Si vous 
soustrayez 255.255.248.0 (masque de sous-réseau) de 255.255.255.255, le résultat est 0.0.7.255.
 
<#root>
access-list acl_permit permit ip 192.168.32.0 0.0.7.255
 
Examinez cet ensemble de réseaux pour plus d'explications.
 
192.168.146.0/24 
192.168.147.0/24

192.168.148.0/24 
192.168.149.0/24
 
