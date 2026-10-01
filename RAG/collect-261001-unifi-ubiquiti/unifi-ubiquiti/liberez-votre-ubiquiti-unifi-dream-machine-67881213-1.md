---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/liberez-votre-ubiquiti-unifi-dream-machine-67881213-1
title: "liberez-votre-ubiquiti-unifi-dream-machine-67881213"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/liberez-votre-ubiquiti-unifi-dream-machine-67881213.md
source_anchor: ""
source_lines: [1, 56]
sha256: 0a877b7404fba596652155e79451acf94eb9c8485f98adf45839ff91d4d93e92
---

# liberez-votre-ubiquiti-unifi-dream-machine-67881213

Libérez votre Ubiquiti Unifi Dream Machine
L'Unifi Dream Machine est un équipement réseau tout-en-un (routeur, switch, point d'accès wifi) puissant et sur lequel on peut ajouter d'autres services.
Pour une fois, nous n'allons pas parler de Docker, mais de réseau, avec Ubiquiti !
Ubiquiti, si vous ne connaissez pas, est un fabricant de matériel proposant sous diverses marques (UniFi, EdgeMax, UISP, AirMax, AirFiber, GigaBeam et UFiber pour les citer) des équipements réseaux filaires et sans-fil pour les entreprises et les particuliers. Unifi est sa marque la plus répandue et présentant la plus large gamme.
En complément de cette activité matérielle, Ubiquiti développe ses propres logiciels de gestion, que ce soit Unifi Network Controller pour gérer leurs équipements réseaux, Unifi Protect pour la vidéosurveillance, Unifi Access pour les solutions d'accès à des bâtiments, ou bien le dernier en date, en bêta actuellement, Unifi Talk, la solution VoIP maison.
Nous sommes plusieurs sur notre communauté Telegram à être équipés en matériel Ubiquiti Unifi, notamment vos auteurs, Hexamus, Trashoune et Guillaume. Parmi la gamme variée de produits au catalogue, il en existe un, sorti récemment, qui cumule les fonctionnalités de plusieurs autres, l'Unifi Dream Machine.
L'Unifi Dream Machine (ou UDM) est un appareil tout-en-un incluant un point d'accès wifi, un switch Gigabit 4 ports, une passerelle de sécurité et le contrôleur Unifi Cloud Key intégré.
Une version Pro existe, rackable pour être intégrée dans une baie 19", avec un switch Gigabit 8 ports et 2 ports fibre SFP+, ainsi que tous les contrôleurs Unifi existants à ce jour, et la possibilité de lui ajouter un disque dur pour les enregistrements avec Unifi Protect.
|  | Unifi Dream Machine | Unifi Dream Machine Pro | 
|---|---|---|
| Ports Gigabit RJ45 | 1 WAN + 4 LAN | 1 WAN + 8 LAN | 
| Ports 10G SFP+ | Non | 1 WAN + 1 LAN | 
| Point d'accès Wifi | Oui | Non | 
| Unifi Network Controller | Oui | Oui | 
| Security Gateway | Oui | Oui | 
| Unifi Protect | Non | Oui | 
| Unifi Access | Non | Oui | 
| Unifi Talk | Non | Oui | 
| Processeur | Quad-Core 1.7 GHz | Quad-Core 1.7 GHz | 
| Mémoire | 2 GB DDR3 | 4 GB DDR4 | 
| Stockage | 16 GB | 16 GB | 
| Emplacement Disque dur | Non | Oui | 
| Débit maximal avec IDS/IPS | 850 Mbps | 3.5 Gbps | 
| Prix officiel | 299€ | 383€ | 
| Achat | Lien | Lien | 
Après cette rapide présentation d'Ubiquiti et plus précisément des modèles Unifi Dream Machine et sa version Pro, passons au vif du sujet, à propos de ces 2 matériels justement. Propulsés par le système Unifi OS, il est possible de procéder à quelques améliorations, notamment pouvoir y installer des applications, parmi lesquelles Adguard, PiHole, Let's Encrypt, Wireguard...
Toutes les explications nécessaires à la mise en place de ces différentes modifications sont détaillées dans le dépôt Github suivant, et nous allons vous éclaircir un peu cette procédure :
Prérequis
Pour nous permettre de faire les installations sur notre UDM, nous avons besoin de nous connecter en SSH en tant que root. Pour cela, dans l'interface UniFi, rendez-vous dans Settings > Advanced, puis activez le SSH et définissez un mot de passe. Vous pourrez ensuite utiliser Putty ou un autre terminal.
Première étape : On-Boot-Script
A chaque démarrage ou mise à jour du firmware, l'UDM se réinitialise (pas d'inquiétude, les réglages effectués dans les différents contrôleurs sont bien sauvegardés).
C'est là qu'intervient "On-Boot-Script", qu'on va mettre en place pour disposer de la mise en cache de tous les packages d'installation de style Debian et pour qu'ils soient exécutés au démarrage, un peu comme le fait init.d. Cela nous permettra de conserver et d'exécuter des scripts et des personnalisations pendant le processus de démarrage de l'UDM.
Commençons par nous connecter en SSH à l'UDM, avec le compte root et le mot de passe que vous avez choisi. Il faut ensuite basculer à l'intérieur de l'OS Unifi :
Vient ensuite le téléchargement du paquet et son installation :
Voilà, c'est tout, votre Unifi Dream Machine est maintenant libérée et délivrée pour accueillir de nouveaux services !
Ajout de AdGuard en natif sur l'UDM
Voilà enfin venu le moment d'aborder ce qui a été à l'initiative de cet article, Adguard ! Suite au tutoriel de Guillaume sur AdGuard et ayant vu qu'il était possible de l'installer sur l'UDM directement, je me suis dit que cela serait bien que je mette ça en place.
Mais comment va-t-on faire tourner AdGuard sur l'UDM ? L'installer directement sur Unifi OS n'est pas forcément la meilleure solution... Ah, le dépôt Github parle de conteneurs !
En effet, j'ai dit en introduction qu'on n'allait pas parler de Docker ! Mais c'est vrai, parce qu'on va parler de conteneurs podman ! Pour faire vite, podman est un moteur de conteneurs, similaire à Docker, mais qui tourne sans droits root grâce au lancement de chaque conteneur dans son propre démon.
Bon, fermons cette parenthèse, et mettons nous au travail !
Il faut commencer par créer un nouveau réseau dans le contrôleur Unifi, et sans DHCP. Dans notre exemple, nous allons utiliser le VLAN 5 et le sous-réseau 10.0.5.1/24 (plus simple car c'est celui qui est défini dans les scripts que nous allons utiliser), mais vous pouvez choisir le VLAN et l'IP que vous voulez.
On retourne ensuite sur le terminal SSH de notre UDM, où il faut télécharger le fichier 10-dns.sh dans le répertoire /mnt/data/on_boot.d.
Vous devrez modifier dans les premières lignes du fichier le VLAN et les adresses IP pour suivre votre configuration, ainsi que la ligne 28 où il faut renseigner le nom du conteneur, dans notre cas adguardhome. Pour éditer le fichier, la commande est la suivante : vi /mnt/data/on_boot.d/10-dns.sh.
Petite aide rapide pour utiliser vi :
- Utilisez la touche i pour passer en mode édition.
- Enregistrez vos modifications en quittant le mode édition en appuyant sur Esc
- Enfin, sauvegardez en tapant :wq puisEnter
Une fois le fichier téléchargé et adapté, nous devons le rendre exécutable et l'exécuter une première fois.
Continuons avec la création de l'adresse mac dédiée au conteneur podman adguardhome sur le vlan. Pour cela, téléchargez le fichier 20-dns.conflist dans le répertoire /mnt/data/podman/cni. Comme pour le précédent fichier, on utilise la commande curl. On utilise également à nouveau l'éditeur vi pour personnaliser l'adresse IP du conteneur podman adguardhome. L'adresse mac devra être du format suivant : 84:DD:B7:19:DA:08, vous pouvez laisser libre court à votre imagination (chiffres de 0 à 9 et lettres de A à F sont acceptées) ou utiliser un générateur d'adresse mac.
Nous allons maintenant créer les répertoires où AdguardHome va stocker ses données.
Nous pouvons enfin lancer le conteneur adguardhome :
Pour vérifier que vous avez tout bien suivi, il vous suffit de vous rendre avec votre navigateur sur l'adresse de conteneur que vous avez défini http://10.0.5.3:3000. Bravo, vous êtes sur l'interface d'initialisation d'Adguard Home et je vous renvoie vers l'article écrit par Guillaume pour sa configuration.
Dernière étape, il vous faut renseigner l'adresse IP d'AdguardHome en tant que DNS pour votre réseau.
Vous pouvez vérifier dans podman le conteneur adguardhome et par la même occasion constater qu'unifi-os est aussi lui-même un conteneur podman.
Je dois avouer que j'ai rencontré quelques difficultés à comprendre les explications sur le github avant d'arriver à mon but, et qu'en plus j'ai rajouté des fautes de frappe, donc cela n'aidait pas.
Externalisation des sauvegardes du contrôleur Unifi
