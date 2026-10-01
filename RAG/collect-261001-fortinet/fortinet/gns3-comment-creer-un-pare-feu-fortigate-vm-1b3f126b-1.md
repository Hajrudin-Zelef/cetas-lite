---
id: collect-261001-fortinet/fortinet/gns3-comment-creer-un-pare-feu-fortigate-vm-1b3f126b-1
title: "gns3-comment-creer-un-pare-feu-fortigate-vm-1b3f126b"
domain: fortinet
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/gns3-comment-creer-un-pare-feu-fortigate-vm-1b3f126b.md
source_anchor: ""
source_lines: [1, 71]
sha256: ec0f3b3df68a513d2012755581c84d041be2641f0c2979d76051850f00ea6707
---

# gns3-comment-creer-un-pare-feu-fortigate-vm-1b3f126b

GNS3 : comment créer un pare-feu FortiGate VM ?
I. Présentation
Dans ce tutoriel, nous allons voir comment importer et configurer FortiGate VM dans le simulateur réseau GNS3. Une configuration idéale pour vous exercer !
Vous aimeriez découvrir l'administration des pare-feux FortiGate, mais vous n'avez pas d'équipement physique à disposition ? Une solution à ce problème consiste à importer une appliance FortiGate dans le simulateur réseau GNS3.
Il s'agit ici d'une émulation de FortiOS (système d'exploitation des équipements Fortinet) exécutée sur une machine virtuelle. Cette émulation (souvent appelée FortiGate VM) fonctionnera comme si nous étions sur un pare-feu réel. Ce qui nous permettra d'effectuer des simulations et autres configurations (routing, NAT, règles firewall, etc.) comme si on le faisait sur une infrastructure réelle.
Les professionnels de la sécurité informatique, experts réseaux, étudiants et débutants utilisent largement cette méthode pour les simulations. Le but de ce tutoriel est donc de présenter comment importer cette appliance FortiGate dans GNS3.
II. Importation de FortiGate VM dans GNS3
A. Prérequis
Avant de débuter, il est recommandé d'avoir les éléments ci-dessous à sa disposition :
- Un PC dont le processeur prend en charge la virtualisation
- Le simulateur réseau GNS3 installé sur le PC (la version la plus récente et la plus stable est la 2.2.53)
- L'hyperviseur de niveau 2 VMware Workstation installé sur le PC.
- L'image de GNS3 VMware Workstation (elle se télécharge en même temps que l'installation de GNS3)
- Un terminal web (un PC virtuel importé dans GNS3) comme client pour le pare-feu.
- Une image FortiGate téléchargée sur GNS3 Marketplace ou directement sur le site officiel de Fortinet.
Remarque : Pour la sécurité de votre système, téléchargez toujours les images FortiGate depuis les sources officielles. Les images provenant d'autres plateformes pourraient être compromises et l'endommager.
Voici d'ailleurs les liens de GNS3 Marketplace et Fortinet pour le téléchargement des appliances :
Si vous débutez avec GNS3, retrouvez nos précédents tutoriels à ce sujet :
- Introduction à GNS3 : installation étape par étape pour bien débuter
- GNS3 : comment importer une image Cisco et créer son premier switch ?
- GNS3 : comment créer et importer une machine virtuelle VMware ?
- Débuter avec VMware Workstation Pro
B. Procédure
1. Premièrement, on va créer un réseau virtuel VMnet0 dans VMWare, qu'on lie ensuite à notre réseau physique. On introduit la GNS3 VM dans ce réseau virtuel. On la démarre et on se rassure qu'elle est a bien obtenu la configuration réseau désirée. Cette configuration est très importante pour la suite et surtout pour la communication avec GNS3.
2. Puis, on démarre GNS3. On crée un nouveau projet (nommé à sa guise), et on va cliquer sur "New template". Une fenêtre s'ouvre et on choisit l'option "Install an appliance from the GNS3 server (recommended)" et on clique sur "Next".
3. On s'assure que les appliances disponibles dans le cache local de GNS3 sont bien celles disponibles depuis GNS3 Marketplace, en cliquant sur "Update from online registry".
Cette option n'est bien évidemment pas nécessaire si on choisit de télécharger ses images sur des sites tiers. On rappelle encore ici que ce choix peut causer des dommages irréversibles dans le cas où vous tombez sur des images vérolées.
4. On sélectionne ensuite le menu déroulant "Firewalls". Puis, on sélectionne "FortiGate" et on clique sur "Install".
5. Ensuite, on clique directement sur le bouton "Next".
6. On choisit le binaire qui fera tourner notre appliance sur la VM. On peut laisser celui sélectionné pour nous par GNS3. Donc, on clique sur "Next" une nouvelle fois.
7. On va ensuite choisir la version de l'image qu'on veut importer. On clique sur "Create new version" puis sur "Next".
Si vous utilisez une image provenant d'une source externe, cochez "Allow custom files". Cela empêchera GNS3 de comparer son hash attendu avec celui de votre fichier, et donc d'interrompre l'importation. Cela reste cependant risqué pour votre PC (sauf si vous travaillez en environnement bac à sable).
8. On va récupérer l'image téléchargée et on continue. Ici, vous devez sélectionner le fichier correspondant à votre image FortiGate VM.
9. À la question de savoir si on souhaite installer cette version de FortiGate, on répond "yes".
10. On valide l'importation en cliquant sur le bouton "Finish".
11. On importe ensuite le nouveau pare-feu dans notre projet et on démarre tous les équipements. Vous obtenez alors une représentation similaire à celle présentée ci-dessous.
12. On lance une première connexion au pare-feu à l'aide de la console. Le login par défaut est admin et le mot de passe par défaut est vide (aucun caractère à saisir ici, juste valider avec la touche Entrée du clavier). Après la première connexion, il est obligatoire de remplacer le mot de passe vide par un mot de passe plus robuste. Une fois fait, on peut poursuivre avec la configuration.
FortiFirewall-VM64-KVM login: admin
Password:
You are forced to change your password. Please input a new password.
New Password:
Confirm Password:
Welcome!
Ce qui donne en image :
13. On passe ensuite en mode configuration d'interfaces avec la commande config system interface. Le port à configurer ici est le port 2 du pare-feu (port communicant avec le LAN) avec la commande edit. On lui attribue l'adresse IP 192.168.1.150 en /24, on active l'écoute sur les protocoles https, http, ssh et icmp (ping) avec la commande set allowaccess et on valide le tout avec la commande end.
FortiFirewall-VM64-KVM # config system interface
FortiFirewall-VM64-KVM (interface) # edit port2
FortiFirewall-VM64-KVM (port2) # set ip 192.168.1.150 255.255.255.0
FortiFirewall-VM64-KVM (port2) # set allowaccess https http ssh ping
FortiFirewall-VM64-KVM (port2) # end
Ce qui donne en image :
14. On va sur notre client (terminal web), et on lui attribue une adresse IP tirée de la même plage que celle du pare-feu.
root@webterm-1:~# ifconfig eth0 inet 192.168.1.151 netmask 255.255.255.0
root@webterm-1:~# ifconfig eth0 up
La configuration étant terminée, on va tester tout ça !
C. Se connecter à l'interface web du FortiGate
Tout d'abord, on commence par un test de la connectivité entre le terminal Web et le pare-feu. Dans ce cas, la commande ping est notre alliée.
root@webterm-1:~# ping 192.168.1.150
PING 192.168.1.150 (192.168.1.150) 56(84) bytes of data.
64 bytes from 192.168.1.150: icmp_seq=1 ttl=255 time=2.45 ms
64 bytes from 192.168.1.150: icmp_seq=2 ttl=255 time=0.817 ms
64 bytes from 192.168.1.150: icmp_seq=3 ttl=255 time=0.974 ms
64 bytes from 192.168.1.150: icmp_seq=4 ttl=255 time=0.953 ms
64 bytes from 192.168.1.150: icmp_seq=5 ttl=255 time=2.17 ms
64 bytes from 192.168.1.150: icmp_seq=6 ttl=255 time=0.850 ms
64 bytes from 192.168.1.150: icmp_seq=7 ttl=255 time=0.883 ms
64 bytes from 192.168.1.150: icmp_seq=8 ttl=255 time=0.905 ms
64 bytes from 192.168.1.150: icmp_seq=9 ttl=255 time=1.24 ms
Ensuite, on rentre l'adresse du pare-feu la barre d'adresse du navigateur du terminal. La page de connexion apparait et nous rentrons nos identifiants précédemment créés. Pour rappel, il s'agit du compte admin et du mot de passe personnalisé que vous avez choisi.
À la question de savoir si on souhaite ajouter un Hostname ou configurer le Tableau de bord, on peut cliquer sur "Later". Une fois fait, on se retrouve directement sur l'interface d'administration du pare-feu et on peut débuter avec les configurations de services.
III. Conclusion
