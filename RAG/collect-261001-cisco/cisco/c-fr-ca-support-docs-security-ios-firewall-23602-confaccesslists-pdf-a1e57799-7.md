---
id: collect-261001-cisco/cisco/c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799-7
title: "c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799.md
source_anchor: ""
source_lines: [1150, 1214]
sha256: b618112966ad8ddb542c6231a56e00079ffa79bd79a4da301f4d3f533d6b1c53
---

# c-fr-ca-support-docs-security-ios-firewall-23602-confaccesslists-pdf-a1e57799

Remarque : Dans les versions plus récentes de Cisco IOS, le concept des listes de 
contrôle d’accès compilées (turbo) a évolué ou est devenu obsolète sur de nombreuses 
plates-formes, en raison des progrès réalisés en matière de matériel et de logiciels.
Listes de contrôle d'accès basées sur l'heure distribuées
Les listes de contrôle d'accès basées sur l'heure distribuées ont été introduites dans le logiciel 
Cisco IOS Version 12.2.2.T afin d'implémenter les listes de contrôle d'accès basées sur l'heure sur 
des routeurs de la gamme Cisco 7500 compatibles avec VPN. Avant l'introduction de la 
fonctionnalité de liste de contrôle d'accès basée sur l'heure distribuée, les listes de contrôle 
d'accès basées sur l'heure n'étaient pas prises en charge sur des cartes de ligne pour les routeurs 
de la gamme Cisco 7500. Si les listes de contrôle d'accès basées sur l'heure étaient configurées, 
elles se comportaient comme des listes de contrôle d'accès normales. Si une interface d'une carte 
de ligne était configurée avec des listes de contrôle d'accès basées sur l'heure, les paquets 
commutés dans l'interface n'étaient pas distribués commutés via la carte de ligne, mais transférés 
au processeur de routage pour traitement.

La syntaxe des listes de contrôle d’accès temporelles distribuées est la même que celle des listes 
de contrôle d’accès temporelles, avec l’ajout des commandes relatives à l’état des messages IPC 
(Inter Processor Communication) entre le processeur de routage et la carte de ligne.
 
<#root>
debug time-range ipc 
show time-range ipc 
clear time-range ipc
 
 
Remarque : Dans les versions plus récentes de Cisco IOS, le concept de liste de contrôle 
d’accès temporelle distribuée a évolué ou est devenu obsolète sur de nombreuses plates-
formes, en raison des progrès réalisés en matière de matériel et de logiciels.

Listes de contrôle d'accès de réception
Les listes de contrôle d’accès de réception ont été utilisées afin d’accroître la sécurité sur les 
routeurs Cisco 12000 en protégeant le processeur de routage Gigabit (GRP) du routeur contre le 
trafic inutile et potentiellement néfaste. Les listes de contrôle d'accès de réception ont été ajoutées 
comme dérogation spéciale à la limitation de maintenance pour le logiciel Cisco IOS 
Version 12.0.21S2 et intégrées à 12.0(22)S.
Remarque : Dans les versions plus récentes de Cisco IOS, le concept de réception des 
listes de contrôle d’accès a évolué ou est devenu obsolète sur de nombreuses plates-
formes, en raison des progrès réalisés en matière de matériel et de logiciels.
Listes de contrôle d'accès de protection d'infrastructure
Les listes de contrôle d'accès d'infrastructure sont utilisées afin de minimiser le risque et l'efficacité 
d'une attaque d'infrastructure directe par l'autorisation explicite du seul trafic autorisé vers

l'équipement d'infrastructure alors qu'il autorise tout autre trafic de transit. Reportez-vous à 
Protection de votre noyau : Listes de contrôle d'accès de protection d'infrastructure pour plus 
d'informations.
Listes de contrôle d'accès de transit
Les listes de contrôle d'accès de transit sont employées afin d'améliorer la sécurité du réseau 
puisqu'elles autorisent uniquement de manière explicite le trafic nécessaire dans votre ou vos 
réseaux. Reportez-vous à Listes de contrôle d'accès de transit : Filtrage au niveau de votre 
périphérie pour plus d'informations.
Informations connexes
Configurer les adresses IP couramment utilisées par les listes de contrôle d’accès (ACL)•
RFC 1700 •
RFC 1918 •
Access Lists Support Page•
Cisco IOS Firewall•
Assistance et documentation techniques - Cisco Systems•

À propos de cette traduction
Cisco a traduit ce document en traduction automatisée vérifiée par une personne dans le
cadre d’un service mondial permettant à nos utilisateurs d’obtenir le contenu d’assistance
dans leur propre langue.
 
Il convient cependant de noter que même la meilleure traduction automatisée ne sera pas
aussi précise que celle fournie par un traducteur professionnel.
