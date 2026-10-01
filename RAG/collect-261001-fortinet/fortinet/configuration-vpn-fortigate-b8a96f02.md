---
id: collect-261001-fortinet/fortinet/configuration-vpn-fortigate-b8a96f02
title: "configuration-vpn-fortigate-b8a96f02"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/configuration-vpn-fortigate-b8a96f02.md
source_anchor: ""
source_lines: [1, 60]
sha256: 452f409198b6ba89e63991b257e08da137808efe50a6fc67bb80fbdba6cc3b30
---

# configuration-vpn-fortigate-b8a96f02

Sommaire
Toggle
Introdcution
« Dans cet article, nous allons explorer en détail le processus de configuration d’un VPN SSL sur un pare-feu Fortigate. Nous aborderons les étapes nécessaires pour créer un tunnel sécurisé entre les utilisateurs distants et le réseau interne, en utilisant le protocole SSL pour garantir la confidentialité des communications. Nous verrons également comment configurer les paramètres de sécurité et configurer les politiques de pare-feu appropriées pour contrôler l’accès aux ressources internes.
Activer le module VPN SSL
Pour commencer, nous allons donc aller dans « System », « Feature Visibility » et activer l’option « SSL-VPN » et valider avec le bouton « Apply ».
Configuration du VPN
Nous allons dans « VPN » puis « SSL-VPN Settings »
Nous devons mettre le « SSL-VPN Status » sur « Enable ».
Sélectionner l’interface d’écoute qui servira a se connecter au VPN dans le champ « Listen on Interface »
Choisir le port de vous souhaiter utiliser dans le champ « Listen on Interface »
Dans notre exemple nous utiliserons donc le certificat par défaut du Fortigate
Activer la redirection de HTTP vers le SSL en valident l’option « Reduction HTTP to SSL-VPN »
Dans « Address Range » vous pouvez garder le plan d’adresse du tunnel VPN ou le personnaliser avec « Specify custom IP ranges »
Vous avez la possibilité de renseigner vos DNS avec l’option « DNS Server »
Avec le bouton « Create New » nous allons créé un nouvelle utilisateur.
Il faut cliqué sur le + pour ajouter l’utilisateur, la liste s’affiche a droit. Maintenant il faut faire « Create ».
Choisir « User » un nouveau menu d’affiche.
Dans le wizard ci-dessous nous allons donc pouvoir créé l’utilisateur.
Sélectionner « Local User » et faire « Next ».
Renseigner le champ « Username » et « Password » et faire « Next ».
Nous avons la possibilité d’activé l’authentification à deux facteurs. Pour cela il faut activé l’option « Two-factor Authentification » choisir « FortiToken », sélectionner un « Token » dans le menu déroulant et renseigner l’adresse mail associé et faire « Next »
Vous avez par defaut 2 licences pour le FortiToken.
Activer l’utilisateur avec « Enabled » puis faire « Submit ».
Sélectionner l’utilisateur dans la liste et dans le menu déroulant « Portal » indiquer « tunnel-acces » puis faire « OK ».
Nous sommes de nouveau sur la page « SSL-VPN Setting.
Nous avons bien l’utilisateur et le profil que nous avons choisi. Faire « Apply pour appliqué les paramètres.
Configuration du portail
Ici nous avons la possibilité d’autoriser certaine action dans « Tunnel Mode Client Option ».
Il est également possible de vérifier la conformité du poste à l’aide des option « Host Check » qui va vérifier que le firewall et/ou l’antivirus de la machine est activé ».
Nous avons également la possibilité de faire des restrictions sur le type d’OS dans l’option « Restrict to Specific OS Version ».
Valider la configuration avec le bouton « OK ».
Il faut maintenant ouvrir les flux entrants et sortant pour les clients VPN.
Dans le menu « Policy & Objects » puis « Firewall Policy » et faire « Create new ».
Nous allons créer une règle pour autoriser les connexions VPN depuis l’extérieur vers le réseau interne « LAN ».
Il faut renseigner les champs suivants:
- Name
- Incoming interface
- Outgoing interface
- adresse source et utilisateur autoriser
- les services (ports) de destination
- Déactivation du NAT
- « Log allowed traffic » pour activer tous les logs
- « Enable this policy » pour activé la régler
Il est possible de mettre des profils de sécurité. nous verrons cela dans une autre article.
Vous pouvez valider avec « OK ».
Nous allons maintenant créer une règle pour autoriser l’acces internet a notre utilisateur depuis le VPN.
Pour cela nous devons créer la règle ci-dessous:
- Name
- Incoming interface
- Outgoing interface
- Source
- Destination
- Service
- Action
- NAT
- Log allowed traffic
Nous pouvons valider avec le bouton « OK ».
Il ne reste plus qu’a installer le Forticlient sur vos machines et le configurer avec les informations de connexion.
Lors de la connexion il faudra renseigner le token, celui-ci vous sera envoyé par mail. Il faudra l’enregistrer dans l’application « FortiToken ».
