---
id: collect-261001-fortinet/fortinet/vmware-workstation-pro-comment-creer-un-pare-feu-fortigate-vm-7cfe91bd-1
title: "vmware-workstation-pro-comment-creer-un-pare-feu-fortigate-vm-7cfe91bd"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "valuation"]
source: docs/RAG/collect-261001-fortinet/vmware-workstation-pro-comment-creer-un-pare-feu-fortigate-vm-7cfe91bd.md
source_anchor: ""
source_lines: [1, 70]
sha256: 8bd7da8c20557fd5835d747e495020f88fbf0b9296bed824ae987ec12ec8264c
---

# vmware-workstation-pro-comment-creer-un-pare-feu-fortigate-vm-7cfe91bd

VMware Workstation Pro : comment créer un pare-feu FortiGate-VM ?
I. Présentation
Dans ce tutoriel, nous allons voir comment importer l'image d'un pare-feu FortiGate dans VMware Workstation Pro pour faire des labs.
Ceci est une alternative à FortiGate-VM sous GNS3. L'installation est simple : il suffit de télécharger l'image du pare-feu sur le site officiel de FortiGate et de l'importer dans VMware Workstation pour la faire fonctionner.
Il est également possible de l'administrer directement depuis son hôte, ou de l'isoler dans un VMnet pour l'administrer depuis une autre machine virtuelle au sein du même VMnet. Ce qui est encore plus intéressant, c'est qu'on peut récupérer la configuration du pare-feu virtuel pour l'appliquer à un pare-feu physique, et ça fonctionne sans problème.
Que vous soyez professionnel, étudiant ou débutant, cette option constitue également une opportunité intéressante pour vous entraîner à l'administration des pare-feu FortiGate.
Pour rappel, voici comment créer un pare-feu FortiGate avec GNS3 :
II. Importation de FortiGate VM dans VMware Workstation Pro
A. Prérequis
Avant de débuter, il est recommandé d'avoir les éléments ci-dessous à sa disposition :
- Un PC dont le processeur prend en charge la virtualisation
- L'hyperviseur de niveau 2 VMware Workstation Pro installé sur le PC (pour rappel, il est gratuit)
Si vous débutez avec cet hyperviseur, la lecture de ces deux articles est recommandée :
Vous devrez aussi télécharger l'image de FortiGate-VM. Il est fortement recommandé de le faire directement depuis le site officiel de Fortinet afin d'éviter toute image potentiellement compromise. Sachez qu'un compte FortiCloud est nécessaire pour ce téléchargement ; la bonne nouvelle, c'est que sa création est gratuite.
Une fois le compte créé, rendez-vous à l'adresse suivante : support.fortinet.com/Download/VMImages.aspx. Choisissez FortiGate comme produit et VMware ESXi comme plate-forme. Sélectionnez la version de l'image à du pare-feu à émuler (fichier zippé) et cliquez sur Download pour télécharger et l'enregistrer.
Il est important de noter que l'utilisation des produits Fortinet, y compris les images FortiGate-VM, est soumise à des licences.
Les licences d'évaluation FortiGate-VM, bien que limitées en termes de ressources matérielles (CPU, RAM, interfaces réseau) et sans support FortiCloud, sont généralement suffisantes pour se familiariser avec l'administration des pare-feu FortiGate. De plus, ces licences sont gratuites et sont attribuées automatiquement dès la création de votre compte FortiCloud.
Notez qu'il n'est possible d'obtenir qu'une seule licence d'évaluation par compte, et celle-ci est limitée à une seule (VM). Autrement dit, si vous voulez utiliser une licence d'évaluation déjà attribuée à VM sur une autre VM, vous devez d'abord la décommissionner depuis votre compte FortiCloud.
Enfin, l'activation d'une licence d'évaluation sur une FortiGate-VM requiert une connexion internet. Vous devrez alors simplement entrer vos identifiants FortiCare via son interface d'administration.
Sans plus tarder, on va procéder à l'importation et la configuration de notre FortiGate-VM dans VMware Worstation Pro.
B. Importer l'image OVF
Suivez les étapes décrites ci-dessous pour effectuer cette tâche.
1. D'abord, on va décompresser l'image de FortiGate-VM précédemment téléchargée, dans un dossier. On peut directement aller double-cliquer sur le fichier FortiGate-VM64.ovf pour lancer l'importation de la VM dans VMware Workstation Pro ou alors aller l'ajouter depuis l'hyperviseur.
2. Il faut ensuite accepter les conditions de licence d'utilisateur final, puis donner un nom quelconque à la VM (ici FortiGate VM) et cliquer sur Import pour finaliser son importation dans VMware Workstation Pro.
3. On peut ensuite aller modifier la VM à sa guise. Ici, on va juste s'assurer que cette dernière soit connectée à notre réseau physique pour pouvoir l'administrer directement depuis notre hôte. Vous pouvez aussi l'ajouter un réseau a isolé, un réseau NAT, ou tout autre. Si vous la connectez en NAT, vous pouvez créer une règle de redirection de ports dans le "Virtual Network Editor" de VMware Workstation pour y accéder malgré tout depuis votre PC physique.
Vous remarquerez sur l'image ci-dessous, que l'image FortiGate correspond à une VM avec une multitude d'interfaces réseau : idéal pour simuler un réseau complet.
Rappel 1 : Il est important de donner un accès internet à votre FortiGate-VM le temps de l'activation de sa licence, et ce, même si vous comptez le placer ultérieurement dans un environnement réseau isolé.
4. On peut déjà démarrer la VM et la laisser s'installer jusqu'à la fin. Le processus dure quelques minutes environ. La VM redémarrera ensuite.
5. On initie une première connexion au pare-feu à l'aide de la console intégré à la VM. Le login par défaut est admin et le mot de passe par défaut est vide (aucun caractère à saisir ici, juste valider avec la touche Entrée du clavier). Après la première connexion, il est obligatoire de remplacer le mot de passe vide par un mot de passe plus robuste. Une fois fait, on peut poursuivre avec la configuration.
FortiFirewall-VM64 login: admin
Password:
You are forced to change your password. Please input a new password.
New Password:
Confirm Password:
Welcome!
En image :
C. Administrer FortiGate-VM
L'installation initiale de FortiGate-VM est terminée. Vous pouvez maintenant poursuivre sa configuration, soit en utilisant l'interface en ligne de commande (via la console intégrée ou SSH), soit l'interface graphique accessible depuis un navigateur web.
Si vous optez pour l'interface graphique ou une connexion SSH distante, il est nécessaire de connaître l'adresse IP d'un des ports physiques du pare-feu ou de la configurer soi-même, pour y accéder. Vous devez aussi vous assurer que les protocoles relatifs à ces services soient autorisés sur ce port physique.
Comme la VM est sur un réseau physique avec un serveur DHCP, le port relié au LAN physique devrait être automatiquement configuré avec une adresse IP. Pour connaitre l'adresse attribuée à notre port, il suffit de rentrer la commande get system interface physical dans le terminal intégré à la VM.
L'adresse IP de notre port est donc 192.168.1.187. Les services ping https ssh http y sont aussi autorisés par défaut.
 FortiFirewall-VM64 # get system interface physical
== [onboard]
        ==[port1]
                mode: dhcp
                ip: 192.168.1.187 255.255.255.0
                ipv6: ::/0
                status: up
                speed: 10000Mbps (Duplex: full)
                FEC: none
                FEC_cap: none
        ==[port2]
                mode: static
                ip: 0.0.0.0 0.0.0.0
                ipv6: ::/0
                status: up
                speed: 10000Mbps (Duplex: full)
                FEC: none
                FEC_cap: none
        ==[port3]
                mode: static
                ip: 0.0.0.0 0.0.0.0
                ipv6: ::/0
                status: up
                speed: 10000Mbps (Duplex: full)
                FEC: none
                FEC_cap: none
En image :
Ensuite, il suffit d'entrer cette adresse IP dans la barre d'adresse de votre navigateur. Vous serez alors invité à saisir le nom d'utilisateur et le mot de passe que vous avez configuré.
Une fois la connexion établie, une licence est requise pour utiliser la machine virtuelle. Si vous avez une licence FortiGate, il suffit de la renseigner après avoir cliqué sur Full license pour activer le système. Sinon, vous pouvez directement lier la VM à votre licence d'évaluation en cliquant sur le bouton Evaluation license, puis en saisissant les identifiants de votre compte FortiCloud et en cliquant sur OK. La FortiGate-VM redémarrera ensuite.
