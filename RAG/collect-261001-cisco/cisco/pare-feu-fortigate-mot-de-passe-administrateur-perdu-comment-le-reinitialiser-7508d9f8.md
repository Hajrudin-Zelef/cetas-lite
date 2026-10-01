---
id: collect-261001-cisco/cisco/pare-feu-fortigate-mot-de-passe-administrateur-perdu-comment-le-reinitialiser-7508d9f8
title: "pare-feu-fortigate-mot-de-passe-administrateur-perdu-comment-le-reinitialiser-7508d9f8"
domain: cisco
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-cisco/pare-feu-fortigate-mot-de-passe-administrateur-perdu-comment-le-reinitialiser-7508d9f8.md
source_anchor: ""
source_lines: [1, 47]
sha256: 35ab8378ca3dbd1f7f66fc8a67d88d6795ab34e58ba05c0b8e3a428f85a7984c
---

# pare-feu-fortigate-mot-de-passe-administrateur-perdu-comment-le-reinitialiser-7508d9f8

Pare-feu FortiGate : mot de passe administrateur perdu, comment le réinitialiser ?
I. Présentation
Dans ce tutoriel, nous allons voir comment réinitialiser le mot de passe administrateur d'un pare-feu FortiGate. Les équipements fonctionnant avec FortiOS 7.2.4 ou une version ultérieure ne sont pas compatibles avec la méthode qu'on présentera ici. Nous reviendrons sur ce point dans la suite de cet article.
Comme tout dispositif de sécurité, l'accès à l'administration des pare-feu FortiGate requiert une authentification préalable. Actuellement, ces équipements utilisent encore la méthode d'authentification par mot de passe.
Il arrive parfois que l'on oublie ce mot de passe ou qu'on l'ignore complètement (notamment en cas de reprise d'activité). L'accès à l'interface d'administration de l'équipement s'en trouve alors rendu difficile, voire impossible.
L'option la plus judicieuse est de procéder à une réinitialisation totale du mot de passe administrateur. De plus, cette action maintient la configuration actuelle de l'équipement à la fin, contrairement à une réinitialisation d'usine via le bouton reset.
Comment faire pour remettre à zéro le mot de passe admin d'un pare-feu FortiGate en gardant la configuration telle qu'elle est ? C'est ce que nous allons voir dans la suite de ce tutoriel.
Si vous souhaitez vous exercer dans la configuration d'un pare-feu FortiGate, vous pouvez utiliser GNS3. A ce propos, suivez ce tutoriel :
II. Réinitialiser le mot de passe
A. Prérequis
Pour réinitialiser le mot de passe d'un pare-feu FortiGate, on aura besoin du pare-feu lui-même, ainsi que :
- Un PC
- Un câble console
- Le logiciel PuTTY
- Le numéro de série du pare-feu
Comme on peut le deviner ici, la réinitialisation du mot de passe administrateur se fait en ligne de commande. Le compte qu'on utilisera pour se connecter au pare-feu s'appelle maintainer. Il s'agit en fait ici d'un super-utilisateur en quelque sorte. Son mot de passe est aléatoire, c'est-à-dire qu'il est constitué au début des caractères bcpb suivi du numéro de série de l'équipement (souvent inscrit sur le pare-feu), chaque équipement ayant bien-sûr un numéro de série propre à lui.
Voyons ce que cela donne en pratique.
B. Procédure
1. D'abord, on connecte le PC au pare-feu à l'aide d'un câble console et on démarre le pare-feu. Ensuite, on lance le logiciel PuTTY pour avoir accès à l'interface en ligne de commande de l'équipement depuis le PC. Comme cela a été dit plus haut, le login est maintainer et le mot de passe est bcpbFGVMEVPLDHK7E671 (bcpb + numéro de série du pare-feu).
Meister-Gate login: maintainer
Password: bcpbFGVMEVPLDHK7E671
Welcome!
Remarque : L'établissement de la connexion avec l'utilisateur de secours (maintainer) peut nécessiter plusieurs tentatives ou plusieurs redémarrages du pare-feu en cas d'échec initial.
2. L'étape suivante consiste à attribuer un nouveau mot de passe à l'administrateur. On rentre d'abord la commande config system admin. Ensuite, on choisit de modifier le compte-cible avec la commande edit admin. On modifie le mot de passe avec la commande set password <insérer un mot de passe robuste> et on valide le tout avec la commande end. On peut éventuellement se déconnecter avec la commande exit.
Ce qui donne :
Meister-Gate # config system admin
Meister-Gate (admin)# edit admin
Meister-Gate (admin)# set password 123456
Meister-Gate (admin)# end
Meister-Gate # exit
En image :
C. Connexion
On peut initier une nouvelle connexion depuis le terminal à l'aide du compte admin et le mot de passe nouvellement créé. Ce qui, logiquement, devrait être un succès !
Meister-Gate login: admin
Password: 123456
Welcome !
Vous pouvez essayer de vous connecter via l'interface web si l'équipement avait déjà une configuration réseau avant la réinitialisation du mot de passe. Sinon, vous devrez configurer une adresse IP LAN, un masque de sous-réseau et activer le service http(s) en mode console. Vous devez enfin connecter l'équipement à votre PC par câble Ethernet, puis accéder à l'interface web en utilisant l'adresse IP.
On pourrait encore aller plus loin en allant confirmer tout ceci dans les logs du système (Accueil > Journaux et Rapports > Évènements  > Évènements Système). On peut, par exemple, aller vérifier le log de modification du mot de passe admin avec l'utilisateur maintainer. On peut aussi aller voir celui relatif au succès de la nouvelle tentative de connexion avec le compte admin et son nouveau mot de passe.
III. Conclusion
Ce tutoriel touche à sa fin. Ainsi, vous savez maintenant comment réinitialiser le mot de passe administrateur d'un pare-feu FortiGate si vous l'avez oublié ou si vous l'ignorez tout simplement.
Comme il a été mentionné plus haut, cette méthode n'est compatible qu'avec les pare-feu FortiOS antérieurs à la version 7.2.4, car l'utilisateur maintainer a été supprimé à partir de cette version. 
À partir de la version 7.2.4, pour reprendre le contrôle de l'équipement, la réinitialisation d'usine (via le bouton de réinitialisation) est la seule option. Cependant, vous pouvez restaurer vos paramètres précédents si vous avez sauvegardé le fichier de configuration avant de perdre le mot de passe. Il suffit alors d'aller modifier ce fichier en allant changer le mot de passe dans sa section admin, de venir restaurer notre configuration dans le pare-feu et le tour est joué. Nous aborderons d'ailleurs ces points dans un prochain tutoriel.
Bonjour, il serait peut-être intéressant de préciser que cette méthode fonctionne uniquement pour les FortiGate avec une version inférieure à la 7.2.4.
Le compte maintainer n’existe plus sur les versions plus récentes, et il faut maintenant passer par une réinitialisation du FortiGate avec réinjection du fichier de config actuel modifié.
En tout cas, c’est bien ce qui est indiqué par Fortinet.
Bonjour, nous y tiendrons compte !
Merci ! 🙂
