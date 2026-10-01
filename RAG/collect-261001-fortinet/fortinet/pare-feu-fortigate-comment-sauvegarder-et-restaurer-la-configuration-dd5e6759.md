---
id: collect-261001-fortinet/fortinet/pare-feu-fortigate-comment-sauvegarder-et-restaurer-la-configuration-dd5e6759
title: "pare-feu-fortigate-comment-sauvegarder-et-restaurer-la-configuration-dd5e6759"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-fortinet/pare-feu-fortigate-comment-sauvegarder-et-restaurer-la-configuration-dd5e6759.md
source_anchor: ""
source_lines: [1, 54]
sha256: 8a1b1c0212a3878bb8ba6fd4b7bef04d573bfe5c62ec1c948f15820a3cf85ab1
---

# pare-feu-fortigate-comment-sauvegarder-et-restaurer-la-configuration-dd5e6759

Pare-feu FortiGate : comment sauvegarder et restaurer la configuration ?
I. Présentation
Dans ce tutoriel, nous allons voir comment sauvegarder et restaurer la configuration d'un pare-feu FortiGate. Une manipulation basique, mais indispensable !
Plusieurs situations peuvent pousser à sauvegarder ou restaurer la configuration d'un pare-feu FortiGate, comme le transfert de configurations ou la reprise d'activité après un incident. Il est donc important de toujours avoir au moins une ou plusieurs copies de sauvegarde de vos équipements, stockées à des emplacements distincts (la stratégie de sauvegarde 3-2-1 est valable).
Deux méthodes permettent également de sauvegarder ou de restaurer la configuration d'un pare-feu FortiGate : la méthode graphique et la méthode en ligne de commande. On verra ici comment les utiliser.
Pour ce tutoriel, on aura juste une configuration simple pour un début :
- Le hostname du FortiGate est changé en Meister-Gate
- Le fuseau horaire est (GMT+1:00)Africa/Douala
- La nouvelle langue du système est le français, l'anglais étant la langue par défaut.
C'est cette configuration qu'on va aller sauvegarder, puis restaurer sur un nouveau pare-feu.
II. Sauvegarder un pare-feu FortiGate
A. Depuis l'interface graphique Web
Tout d''abord, on se connecte à l'interface d'administration du pare-feu en http(s). On va ensuite dérouler l'onglet de l'utilisateur en cours (ici admin) dans le coin supérieur droit et on place le curseur de la souris sur configuration.
Un menu déroulant s'ouvre et on sélectionne Sauvegarder. L'assistant de sauvegarde s'ouvre ensuite, on poursuit en choisissant la destination (PC local ou Disque USB) et le format du fichier de sauvegarde (FortiOS ou YAML).
Une option permettant de masquer les mots de passe dans le fichier de sauvegarde est également disponible, bien que seules leurs empreintes (ou hash) y soient enregistrées. Les empreintes des mots de passes et autres clés privées seront alors remplacées par la mention FortinetPasswordMask dans le fichier de sauvegarde. Cependant, les fichiers de sauvegarde utilisant cette option sont destinés au support Fortinet et ne servent pas à restaurer la configuration du pare-feu.
La possibilité de sécuriser l'accès au fichier en le soumettant au chiffrement par mot de passe existe aussi.  Il ne reste plus qu'à valider avec OK pour enclencher le processus de sauvegarde.
Le fichier de configuration sauvegardé peut ensuite être visualisé, modifié ou laissé comme tel.
Si l'on visualise le résultat, on constate bien les modifications apportées à notre pare-feu (nom d'hôte, langue, fuseau horaire). Si votre appareil fonctionne sous FortiOS 7.2.4 (ou une version ultérieure), ce fichier permet également de réinitialiser le mot de passe administrateur en cas d'oubli. Pour ce faire, allez dans la section system admin, puis sous admin et allez remplacer la valeur de password par le nouveau mot de passe désiré, puis sauvegardez le tout.
Rappel : Un fichier sauvegardé avec la mention "Masquer les mots de passe" n'est pas destiné à être utilisé pour effectuer une restauration.
B. En ligne de commande
Comme toujours, l'avantage principal avec la procédure en ligne de commande est le gain de temps. De plus, cette méthode offre des options que l'on ne retrouve pas toujours en mode graphique. Ici, on peut directement transférer la sauvegarde de configuration vers des supports tels que la mémoire flash, un serveur FTP, SFTP, TFTP, un stockage USB (formatage ext2 parfois obligatoire) ou vers une plate-forme de gestion centralisée (FortiManager, FortiCloud, etc.).
Il existe également plusieurs modes de sauvegarde :
- Config : ce mode permet de sauvegarder la configuration au format FortiOS et vers toutes les destinations de sauvegardes listées plus haut.
- Full-config : permet de sauvegarder toutes les configurations, y compris celle par défaut, vers toutes les destinations de sauvegardes listées plus haut, excepté la plate-forme de gestion centralisée.
- Yaml-config : permet la sauvegarde des configurations au format YAML et uniquement vers un serveur FTP ou TFTP.
- Obfuscated-config : même objectif que le mode config , mais avec les mots de passe et les secrets masqués.
- Obfuscated-full-config : même objectif que le mode full-config , mais avec les mots de passe et les secrets masqués.
- Obfuscated-yaml-config : même objectif que le mode yaml , mais avec les mots de passe et les secrets masqués.
La commande de sauvegarde est execute backup. Nous devrons ensuite spécifier le mode de sauvegarde, le support de sauvegarde, le nom du fichier de sauvegarde, ainsi que la destination de sauvegarde (qui inclut l'adresse, et potentiellement un identifiant et un mot de passe). 
Voici un exemple pour effectuer une sauvegarde vers un serveur FTP (port 21) dont l'adresse IP est 192.168.1.117. Autres informations importantes, on utilise le compte meister et le mot de passe 12345 pour s'authentifier sur le serveur distant.
Meister-Gate # execute backup config ftp meister_gate.conf 192.168.1.117:21 meister 12345
Please wait...
Connect to ftp server 192.168.1.117:21 ...
Please wait...
Send config file to ftp server OK.
III. Restauration d'un pare-feu FortiGate
A. Depuis l'interface graphique Web
On est toujours sur l'interface d'administration graphique du pare-feu. Rendez-vous toujours dans le coin supérieur droit, puis déroulez l'onglet de l'utilisateur en cours (ici admin), placez le curseur de la souris sur configuration et choisissez restore cette fois-ci.
Sélectionnez l'emplacement où se trouve le fichier de restauration (dans notre cas, Local PC ), puis cliquez sur Upload pour le charger. Si le fichier de restauration est protégé par un mot de passe, vous devrez le saisir dans la section Password pour que l'équipement puisse le lire. Sinon, on peut sauter cette étape. Lancez la restauration en cliquant sur Ok. L'équipement redémarrera à la fin de la restauration la configuration souhaitée.
B. En ligne de commande
On est toujours ici sur notre terminal console ou SSH.
La commande de restauration est execute restore. Nous devrons ensuite spécifier le mode de restauration, le support de restauration, le nom du fichier de restauration, ainsi que la source de restauration (qui inclut l'adresse, et potentiellement un identifiant et un mot de passe).
Le pare-feu nous avertie ensuite que l'opération de restauration effacera sa configuration actuelle et qu'il sera probablement redémarré. Nous lui demanderons de poursuivre en rentrant y pour yes dans le terminal.  
Voici un exemple, mais cette fois-ci, en utilisant un serveur TFTP comme source de restauration :
FGVMEVZGAQKBGUF9 # execute restore config tftp meister_gate 192.168.1.150 
This operation will overwrite the current setting and could possibly reboot the system!
Do you want to continue? (y/n)y
Please wait...
Connect to ftp server 192.168.1.150 ...
Get file from TFTP server OK.
File check OK.
À la fin de la restauration, le pare-feu redémarre automatiquement, et on retrouvera logiquement nos configurations préalablement sauvegardées.
IV. Conclusion
Ce tutoriel touche à sa fin. Nous avons présenté ici les différentes manières de sauvegarder et restaurer la configuration d'un pare-feu FortiGate, incluant les méthodes, modes et outils appropriés. Comme il a été dit plus haut, ces actions peuvent s'avérer utiles et même salutaires en cas de reprise d'activité après incident ou de transfert de configurations. Donc, pensez régulièrement à sauvegarder vos configurations et à tester vos sauvegardes.
