---
id: collect-261001-cisco/cisco/3512-unifi-controller-dubiquiti-utilisez-raspberry-pi-comme-cloud-key-3038ed0e-2
title: "3512-unifi-controller-dubiquiti-utilisez-raspberry-pi-comme-cloud-key-3038ed0e"
domain: cisco
role: reference
task: reference
actors: []
dates: ["2019-02-05", "2019-02-06", "2019-02-07"]
keywords: ["attention", "ethernet"]
source: docs/RAG/collect-261001-cisco/3512-unifi-controller-dubiquiti-utilisez-raspberry-pi-comme-cloud-key-3038ed0e.md
source_anchor: ""
source_lines: [45, 110]
sha256: 7a8b0cb8511e12243969bb94f12a8d14e157443078f1eef8650df535ae8b3052
---

# 3512-unifi-controller-dubiquiti-utilisez-raspberry-pi-comme-cloud-key-3038ed0e

On peut alors installer Unifi Controller :
sudo apt install unifi
Une fois la procédure terminée, Unifi Controller sera accessible sur le port 8443 de votre Raspberry Pi. Vous pouvez donc le gérer à travers les applications mobiles ou l'interface web via l'URL suivante :
https://IP_de_votre_Raspberry_PI:8443
Contrairement à une Cloud Key vous n'aurez en effet pas le choix entre son interface de gestion et un accès à UniFi Controller, qui sera la seule possibilité. La mise à jour du système devra donc se faire de manière classique (via APT). La procédure de configuration et l'interface sont identiques à la version classique.
Une démonstration en ligne est accessible par ici (après création d'un compte).
Les étapes post-installation
Vous pouvez effectuer quelques tâches une fois l'installation finie, comme désactiver le service MongoDB puisque UniFi Controller dispose du sien :
sudo systemctl stop mongodb
sudo systemctl disable mongodb
Vous pouvez également indiquer que le système ne doit pas mettre à jour UniFi Controller en même temps que le reste du système afin de pouvoir procéder à cette étape de manière manuelle :
sudo apt-mark hold unifi
Puis pour la mise à jour manuelle :
sudo apt-mark unhold unifi
sudo apt update
sudo apt install --only-upgrade unifi
sudo apt-mark hold unifi
Vous pouvez voir à tout moment la liste des paquets « marqués » de la sorte :
sudo apt-mark showhold
UniFi Controller au-delà du Raspberry Pi
Bien entendu, cela fonctionne pour toute machine utilisant un système dérivé de Debian avec APT comme gestionnaire de paquets, et donc d'autres micro PC sous Armbian par exemple. Vous pouvez également l'utiliser depuis une machine de récupération, un serveur sous la forme d'une machine virtuelle ou d'un système de conteneurs tel que Docker.
UniFi Controller peut être utilisé via une connexion filaire ou sans fil (si celle-ci ne dépend pas d'un point d'accès UniFi qu'elle contrôle) et sur des machines moins performantes qu'un RPi 3B+, même si cela n'est pas forcément conseillé du fait de la puissance et de la mémoire limitées. Nous l'avons néanmoins installé avec succès sur un RPi Zero W.
Commentaires (13)
Le 05/02/2019 à 12h23
Dans mon ancien boulot j’administrais environ 1000 AP et 30 switch sur un contrôleur Unifi déporté (qu’on hébergeais sur un “petit” serveur dédié). C’est du java et ça consomme pas mal de RAM (c’est proportionnel au nombre d’appareils gérés). En plus, il y doit y avoir des fuites mémoires parce qu’il faut forcément le redémarrer au bout d’un moment.
En dehors de ces défauts, le logiciel est vraiment bien et le matériel aussi. Les MAJ sont assez fréquentes et la communauté est (relativement) écoutée.
Si on ne met pas de stats (ou qu’il n’y a pas de trafic), un Rpi 2 peut faire l’affaire, mais devient très vite ingérable de part sa lenteur.
Ubiquiti propose des produits au rapport qualité/perf/prix plus que correct, je recommande à ceux qui veulent mettre les mains dedans ou qui ont envie d’avoir un réseau wifi (ou pas), de qualité et “hautement” paramétrable.
P.S. : je ne travaille pas pour eux, ils font juste du bon boulot ;)
Le 05/02/2019 à 13h10
+1 pour Ubiquiti et Unifi.
Très pratique pour les entreprises.
J’ai aussi réalisé plusieurs interconnexion avec leurs radio (Bullet, AirFiber…) c’est très simple et efficace.
Le 05/02/2019 à 15h41
Je cherche un installateur qui pourrait mettre un place une installation de vidéosurveillance Unifi dans le parking de ma résidence (dans le 94). Hélas, je n’ai pour l’instant pas réussi à en trouver un seul… Si vous avez des noms, je suis preneur !
J’aime cette solution car je l’utilise chez moi à titre personnel et j’aimerais pouvoir utiliser du matériel que je connais et en lequel j’ai confiance.
Le 05/02/2019 à 16h10
Bonjour ! Je viens de t’écrire un MP pour te proposer mes services :) J’espère à très bientôt !
Le 06/02/2019 à 10h59
Vu le nombre de commentaires, je ne suis pas le seul a ne pas comprendre de quoi on parle dans l’article…
Le 06/02/2019 à 11h17
Il n’y a rien de compliqué pourtant, par contre pour un particulier l’intérêt est plutôt faible en fait.
Le 06/02/2019 à 13h03
Bah tout dépend, je préfère voir un particulier utiliser des AP en complément du routeur de sa box que d’utiliser de mauvais répéteurs. Après ubiquiti ou pas, contrôleur ou pas, RPi ou pas, c’est un autre sujet.
Le souci reste qu’on considère souvent que coller un AP lambda avec le même SSID est une bonne solution. Disons que ça fonctionne à peu près, mais ça pose aussi pas mal de soucis (dont on se fout quand on veut juste un signal qui passe sans se poser plus de questions).
Il y a une astuce : lire l’article ;)
Le 07/02/2019 à 16h33
A la construction de ma maison, j’ai prévu dès le départ 3 AP Unifi pour avoir un excellent wifi partout.
Je ne suis pas déçu. J’avoue que par fainéantise, j’ai préféré acheter une cloud key :p. Je viens d’ailleurs de commander la Gen2 Plus car j’ai aussi prévu de mettre quelques caméras chez moi (c’est clair que quand on construit c’est plus facile de faire passer quelques cables Ethernet).
Pour le POE, j’ai pris des TP-LINK 5 ports (dont 4 POE), juste pour le fait qu’ils restent passifs. Je me serais bien laissé tenter par le switch POE Unifi, mais il semble assez bruyant.
Je reste aussi a l’affut de leur plateforme UAG que je trouve intéressante. Le “petit” modèle est maintenant trop petit pour faire de l’introspection. Quand on regarde les nouveaux modèles qui sont sortis côté routeurs, c’est clair que la gamme UAG va être remise au gout du jour.
D’ailleurs David, tu as jeté un coup d’oeil à ce que sait faire la gamme UAG ?
Le 07/02/2019 à 16h38
Tu as pris quoi comme AP par curiosité ? Parce que même si la couverture est bonne, le débit reste assez léger (surtout si on passe en Mesh). Mais je suis en train de tester dans différents cas pour avoir des données chiffrées.
Par contre attention, passif pour le PoE ça veut juste dire que ; 1. ça ne respecte pas la norme, 2. le jus est envoyé tout le temps sans aucune vérification. Cela n’a rien à voir avec une consommation et une chauffe plus ou moins importante. Voir :
https://www.inpact-hardware.com/article/1100/poe-et-802-3atafbt-ce-quil-faut-savoir-alimentation-par-cable-reseau
Après si tu veux dire que le switch n’a pas de ventilation, attention quand même à sa chauffe ;) (mais sur 5 ports pour des AP ça devrait aller).
Pour l’article, je vais commencer par les AP et la gestion distante, la partie routage ce sera plus tard. C’est déjà bien assez long/compliqué comme ça
Le 07/02/2019 à 17h00
Oppss.
Petite correction. C’est de la gamme USG (Unifi Security Gateway) et non UAG dont je voulais parler. La gamme UAG c’est encore autre chose.
David, c’est vrai que leur gamme n’est pas très lisible. Mais quel plaisir d’avoir un système qui marche “tout seul” par la suite. J’en ai même monté chez mes parents ;).
Le 07/02/2019 à 17h24
Pour les AP, c’est les AC Pro. Je suis resté dans des prix “raisonnables”. Après ça monte en flèche ! Pas de mesh chez moi. Je n’ai pas constaté de problèmes de débits, mais je n’ai certainement pas poussé les tests comme toi.
Pour les POE, c’est bien passif au niveau ventilation. J’ai pris celui la TP LINK TL-SG1005P. A 56W maxi (passé cela, c’est ventilateur obligatoire), c’est suffisant pour les 3 AP + CloudKey. Le second attend pour les caméras (3 G3 Pro qu’il faut que je commande). Pour avoir regardé la température, ce qui chauffe le plus, c’est le transformateur externe :p. Le switch en acier dissipe bien mieux la chaleur. C’est chaud au toucher, pas brulant. Mais j’ai laissé de l’espace autour pour la ventilation.
Le 07/02/2019 à 20h43
