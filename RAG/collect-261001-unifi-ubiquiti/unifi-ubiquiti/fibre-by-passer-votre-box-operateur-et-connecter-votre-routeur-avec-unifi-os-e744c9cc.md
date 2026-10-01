---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/fibre-by-passer-votre-box-operateur-et-connecter-votre-routeur-avec-unifi-os-e744c9cc
title: "fibre-by-passer-votre-box-operateur-et-connecter-votre-routeur-avec-unifi-os-e744c9cc"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-unifi-ubiquiti/fibre-by-passer-votre-box-operateur-et-connecter-votre-routeur-avec-unifi-os-e744c9cc.md
source_anchor: ""
source_lines: [1, 78]
sha256: 22a3bb861b3f2c044d90a06989476d96d3cabb6bab84adc2b969518f254169b8
---

# fibre-by-passer-votre-box-operateur-et-connecter-votre-routeur-avec-unifi-os-e744c9cc

Fibre : by-passer votre box opérateur et connecter votre propre routeur avec UniFi OS
Il est tout à fait possible de se passer de votre box opérateur et d'utiliser ses propres équipements réseau. Toutefois, l'opération n'est pas triviale et s'accompagne de certaines limitations importantes : vous perdrez l'accès à certains services comme le décodeur TV ou le service de téléphonie fixe.
Ce tutoriel s'applique exclusivement aux équipements fonctionnant sous UniFi OS (UniFi Dream Machine, UDM Pro, UDM SE, UXG-Pro, ou toute solution UniFi OS), nous l'avons testé et fonctionne parfaitement pour la Livebox d'Orange (modèle Fibre évidemment), mais des utilisateurs confirment que cela fonctionne aussi pour la Freebox Ultra. 
Disclaimer : cette solution est destinée à un public averti mais qui souhaite tester ce mode de connexion. Nous déclinons toute responsabilité pour toute manipulation qui aurait provoqué une rupture de vos services Fibre. Si ce cas se présente, il ne vous sera possible d'obtenir de support de la part de votre opérateur. Néanmoins, vous pourrez à tout moment rebrancher votre box d'origine.
Pourquoi les opérateurs nous imposent leur box  ?
Contrairement à d'autres pays qui ont franchi le pas, les box opérateurs sont toujours obligatoires en France. Orange, par exemple, impose l'utilisation de la Livebox, avant tout pour faciliter la gestion et le support de ses services.
La box centralise ainsi l'accès Internet avec l'authentification sécurisée, la séparation automatique des flux Internet, TV et téléphonie via VLAN's, la gestion intégrée de la téléphonie fixe (VoIP) et le traitement des flux multicast nécessaires au décodeur TV.
Cette approche permet aux opérateurs de proposer une solution clé en main, simple à installer pour les clients et facile à dépanner pour le service technique.
Toutefois, pour les utilisateurs avancés souhaitant reprendre le contrôle de leur réseau domestique, il reste possible de remplacer la Livebox en connaissance de cause.
Se passer du box opérateur : ce qu’il faut 
1 - Les identifiants de connexion. Chez Orange, ils sont du type fti/xxxxxx .
Ces identifiants sont  fournis dans un papier dans la boîte de la Livebox, si vous les avez perdus, le fti est indiqué dans l'appli Orange et moi. Et si vous avez perdu le mot de passe, il est possible de contacter le Service Client pour le renvoyer .
2 - Un ONT compatible, indispensable si vous êtes en XGS-PON 
Pour ce test, nous utilisons par exemple le modèle WAS-110, disponible sur des plateformes en ligne. Pour le GPON, on peut utiliser le LEOX GPON STICK LXT-010S-H. Pour connaître votre statut d'éligibilité, vous pouvez vous référer à notre article sur le sujet.
🔵 Attention : Nous ne recommandons aucun vendeur spécifique, faites vos propres recherches pour choisir le site qui vous inspire confiance. On rappelle aussi que ces tests ont été réalisés avec la Livebox et les éléments d'identifiant peuvent varier en fonction de l'opérateur.
Configurer le WAS-110 comme ONT 
Après l'installation du firmware communautaire 8311 sur votre WAS-110, vous pouvez l'utiliser pour remplacer la Livebox 7 et accéder à Internet via votre propre routeur. 
Voici les étapes détaillées pour configurer votre ONT et obtenir l'authentification requise pour Orange.
1. Installation du firmware communautaire 8311
Avant toute chose, assurez-vous que votre WAS-110 dispose du firmware 8311 installé.
- Si vous avez acheté un modèle déjà flashé, passez directement à l'étape suivante.
- Sinon, suivez ce guide officiel pour flasher le firmware.
2. Configuration de l’ONU pour l'authentification sur le réseau PON
Il est nécessaire de configurer votre WAS-110 pour qu’il puisse s'authentifier correctement sur le réseau PON d'Orange, en imitant les paramètres d'une Livebox 7 officielle (procédé de "masquerading").
Cette étape est indispensable pour être reconnu par l’OLT et établir une connexion PON.
Depuis l'interface web :
- Accédez à : https://192.168.11.1/cgi-bin/luci/admin/8311/config .
- Connectez-vous avec votre mot de passe root.
- Dans l'onglet PON, renseignez :
  - PON Serial Number (ONT ID) : (copié depuis votre Livebox)
  - Hardware Version : SMBSXLB7400
  - Sync Circuit Pack Version : 1
  - MIB File : /etc/mibs/prx300_1U.ini
- Dans l'onglet ISP Fixes, désactivez Fix VLANs.
- Sauvegardez et redémarrez.
3. Accès au WAS-110 si inaccessible directement
Si vous ne parvenez pas à accéder directement au WAS-110 depuis votre réseau :
- Activez SSH sur votre routeur UniFi (UDM/UDM Pro/UDM Pro max/Etc).
- Connectez-vous en SSH au routeur.
- Exécutez les commandes suivantes (remplacez eth9 par le port où est branché le WAS-110) :
iptables -t nat -I POSTROUTING -o eth9 -d 192.168.11.0/24 -j SNAT --to 192.168.11.2
ip addr add dev eth9 local 192.168.11.2/24
Vérification de l'état de la connexion
Après que la configuration du WAS-110 est terminée, vous devez obtenir le statut O5.1 cela signifie que vous êtes désormais authentifié avec succès sur l'OLT d'Orange.
La partie la plus compliquée est alors passée !
Authentification via DHCP Option 90
Petit rappel : souvenez-vous de votre identifiant PPPoE fti/xxxxxxx.
Pour obtenir une adresse IP sur le réseau Orange, il faut transmettre vos identifiants via DHCP Option 90.
Voici comment faire :
- Allez sur cet outil en ligne.
- Entrez votre identifiant PPPoE.
- Entrez votre mot de passe si fourni par Orange.
- Cliquez sur Run pour générer la valeur de l'Option 90.
Sur votre routeur UniFi :
- Type de connexion WAN : DHCP
- VLAN ID : 832
- MAC Address Clone : votre adresse MAC de la Livebox
- DHCP CoS : 6
- DHCP Options :
- Option 60 : sagem
- Option 77 : FSVDSL_livebox.Internet.softathome.Livebox7
- Option 90 : (valeur générée)
Vérification finale
Si tout est correctement configuré, vous recevrez directement votre IP publique Orange.
Si vous recevez une IP 172.xxx.xxx.xxx, cela indique une erreur dans la configuration (VLAN, MAC clone, DHCP Options) — vérifiez et corrigez.
En cas de difficultés supplémentaires, pensez à inspecter les logs DHCP du routeur pour comprendre où le processus bloque.
Remarque importante sur le support
Maintenant que le bypass est opérationnel, nous tenons à nouveau à rappeler que vous ne pourrez plus contacter le support technique de votre opérateur en cas de problème, tant que vous utilisez un autre moyen que leur matériel officiel. Mais si vous êtes arrivé jusqu'ici, c'est que vous êtes bien décidé alors bravo :) . On rappelle aussi qu'en cas de soucis, reconnectez tout simplement votre box d'origine à la fibre.
Utilisateurs avancés : Mikrotik, OpenWRT, OPNsense
Bien que ce guide cible principalement UniFi OS, les étapes de principe s'appliquent aussi à d'autres routeurs comme :
- Mikrotik RouterOS,
- OpenWRT,
- OPNsense/PFSense,
- Autres équipements compatibles DHCP Option 90/60/77 etc.
Ressources supplémentaires :
Pour plus d'informations et d'exemples de configurations sur d'autres matériels, nous vous recommandons vivement de consulter le forum LaFibre.info, qui a inspiré une grande partie de ce tutoriel :
👉 Forum LaFibre.info - Remplacer la Livebox
N'hésitez pas non plus à poster vos commentaires ou, mieux, nous rejoindre sur le Discord Communautaire Carto Telco si vous souhaitez discuter avec d'autres utilisateurs de la solution.
