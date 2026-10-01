---
id: collect-261001-cisco/cisco/3512-unifi-controller-dubiquiti-utilisez-raspberry-pi-comme-cloud-key-3038ed0e-1
title: "3512-unifi-controller-dubiquiti-utilisez-raspberry-pi-comme-cloud-key-3038ed0e"
domain: cisco
role: reference
task: reference
actors: ["Oracle", "Qualcomm"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/3512-unifi-controller-dubiquiti-utilisez-raspberry-pi-comme-cloud-key-3038ed0e.md
source_anchor: ""
source_lines: [1, 44]
sha256: cff036ce3aeaa36060bc5ac056c75b1ac1bbb66671ed62278af0c8a73486b3be
---

# 3512-unifi-controller-dubiquiti-utilisez-raspberry-pi-comme-cloud-key-3038ed0e

UniFi Controller d’Ubiquiti : utilisez un Raspberry Pi comme une Cloud Key
De quoi faire une belle économie !
Le 05 février 2019 à 11h16
UniFi Controller d’Ubiquiti : utilisez un Raspberry Pi comme une Cloud Key
De quoi faire une belle économie !
Hardware
Hardware
9 min
Ubiquiti est une marque bien connue des amateurs de solutions Wi-Fi destinées aux professionnels, à travers sa gamme UniFi. Celle-ci repose sur une gestion logicielle proposée dans des produits clé en main ou à installer sur votre propre machine. Nous avons tenté le coup avec un Raspberry Pi.
Comme nous aurons l’occasion de le voir dans un prochain dossier, les solutions Wi-Fi dites « professionnelles » ont un fonctionnement un peu différent de celles proposées au grand public.
Le routeur n’est en effet pas le centre d’un système sans fil auquel on peut ajouter des satellites et/ou répéteurs. Chaque élément dispose de son propre rôle : routeur, switchs, points d’accès, caméras, etc. La gestion du réseau sans fil, bien que pouvant être individuelle, est le plus souvent pensée pour être centralisée et multisite.
Elle passe en général par un élément connu sous le petit nom de contrôleur, proposé sous forme d’un matériel clé en main à connecter à votre réseau, ou d’un logiciel à installer sur une machine locale ou distante. Un nombre croissant de constructeurs proposent d’ailleurs des services hébergés, Netgear ayant par exemple fait le choix de ne se reposer que sur cette méthode pour la gestion centralisée de ses produits via Insight.
Aujourd’hui, nous allons évoquer le cas d’un concurrent bien connu des bidouilleurs : Ubiquiti.
Ubiquiti, c’est fouillis mais plutôt réussi
L'Américain s'est taillé en quelques années une place de choix dans le secteur des réseaux, notamment sans fil. Son offre est complète, modulaire, efficace et relativement ouverte. Elle bénéficie aussi d'interfaces agréables à utiliser, ce qui est plutôt rare dans le secteur. Le tout avec des produits plus ou moins abordables.
Ubiquiti a néanmoins un défaut, de taille : sa proposition commerciale est peu lisible. Distinguer ses gammes et les noms de produits/services, comprendre son fonctionnement... rien de tout cela n'est vraiment simple de prime abord. Il faut donc fouiller dans la documentation, analyser et parfois tester pour y parvenir.
Reste que la solution de gestion centralisée de la marque, UniFi Controller, est l'une des plus convaincantes du marché. Elle peut être utilisée comme service hébergé, mais c'est assez cher : 299 dollars par an pour 10 appareils ou moins, 498 dollars jusqu'à 20 appareils, 697 dollars jusqu'à 30 puis 199 dollars par tranche de dix ensuite.
Heureusement, elle est également installable sur la machine (locale/distante) de votre choix ou proposée à travers quatre produits clé en main : les Cloud Key, Cloud Key Gen2 (Plus) et UniFi XG Server.
Les contrôleurs matériels
La première (UC-CK) exploite un SoC à quatre cœurs (MediaTek MT7623), 2 Go de DDR, un port réseau Gigabit et une alimentation PoE (passif/802.3af) ou USB Type-C (5V, 1A). Elle est donnée pour une consommation minimum de 5 watts. On regrette surtout qu'elle soit livrée sans solution de fixation pour rack, son format y étant peu adapté.
Elle est néanmoins compacte et légère : 110 grammes, pour des dimensions de 121,9 x 43,4 x 21,7 mm. Elle dispose d'un micro-bouton reset et d'un autre permettant de l'éteindre. Elle est simplement livrée avec un câble réseau et une carte SD de 8 Go qui est utilisée en complément des 16 Go intégrés, contenant le système. Ce dernier est un dérivé de Debian Jessie (8.11) dans sa version actuelle (3.10.20-ubnt-mtk). Cette Cloud Key se vend aux alentours de 100 euros.
La Gen2 (UCK-G2) annoncée plus récemment se veut bien plus complète avec un SoC huit cœurs (Qualcomm APQ8053), 32 Go de stockage, du Bluetooth, un écran en façade et la possibilité d'un montage en rack (via un adaptateur 1U CKG2-RM). Elle est un peu plus volumineuse et lourde : 150 grammes pour 119,75 x 46,8 x 27,1 mm.
Elle est présentée comme la manière de gérer jusqu'à une cinquantaine d'appareils, sans précisions sur les capacités de la version précédente sur ce point.
Une déclinaison Gen2 Plus (UCK-G2-PLUS) est également proposée avec un disque dur (2,5") de 1 To intégré, pouvant être changé. De quoi lui permettre de gérer 20 caméras, ou 15 caméras et jusqu'à 50 appareils selon le constructeur. Le tout pèse 585 grammes pour des dimensions de 134,2 x 131,16 x 27,1 mm.
Le montage peut également être effectué en rack via l'adaptateur CKG2-RM, mais l'alimentation a été revue à la hausse. Ainsi, outre le PoE (12,95 watts) ou l'USB Type-C, un adaptateur Quick Charge 2.0/3.0 (9V, 2A) peut être utilisé. Ce produit doit permettre à Ubiquiti d'évoluer dans la vidéosurveillance, la société venant d'annoncer sa gamme Protect.
Côté tarif, comptez respectivement 220 euros et 260 euros environ. Vu la faible différence de prix, on aurait préféré que le constructeur ne propose que le Gen2 Plus, mais sans stockage par défaut, à un tarif plus abordable.
Enfin, on trouve l'UniFi XG Server (UAS-XG), qui est, comme son nom l'indique, un serveur (sous Ubuntu). Il intègre un Xeon D-1521, 32 Go de mémoire, un SSD de 120 Go en complément de 2x 2 To de stockage plus classique.
Il propose deux ports 10 Gb/s (RJ45), un Gigabit pour la gestion à distance (IPMI), deux USB 2.0, deux USB 3.0 et du Bluetooth BLE 4.0 pour la phase d'installation. Il intègre UniFi Controller/Video mais peut également faire fonctionner les applications de votre choix en parallèle. Il est proposé à pas moins de 2 000 euros pièce.
Installation d'UniFi Controller sur un Raspberry Pi
Les solutions logicielles d'Ubiquiti reposent en bonne partie sur Debian et ses dérivés. Il est donc possible d'installer le logiciel Unifi Controler sur une machine sous Linux, mais pas seulement : des versions macOS et Windows sont disponibles. Deux canaux existent : LTS (support à long terme) et classique. Notez que dans tous les cas, Java est nécessaire.
Notre objectif du jour est de créer notre propre « Cloud Key », mais pour un tarif moindre. De quoi gérer quelques points d'accès au sein d'un petit réseau local. Nous allons utiliser un simple Raspberry Pi. Un dépôt APT est en effet proposé par Ubiquiti, permettant une installation simple (Raspbian étant lui aussi un dérivé de Debian).
Nous partons du principe que votre Raspberry Pi est configuré sous Raspbian Lite, à jour et accessible via SSH ou une interface graphique. Si ce n'est pas le cas, suivez notre guide d'installation :
Pensez également à lui attribuer une IP fixe au niveau de votre routeur afin d'éviter qu'il n'en change, ce qui pourrait poser des soucis pour l'accès à l'interface ou même la communication avec les appareils.
Il faut tout d'abord ajouter le dépôt d'Ubiquiti afin qu'APT puisse l'utiliser, ainsi que les clés GPG qui servent à vérifier la provenance des fichiers. Pour cela, ouvrez un terminal ou une connexion SSH et tapez :
echo 'deb http://www.ui.com/downloads/unifi/debian stable ubiquiti' | sudo tee /etc/apt/sources.list.d/100-ubnt-unifi.list
sudo apt install dirmngr
sudo apt-key adv --keyserver keyserver.ubuntu.com --recv 06E85760C0A52C50
Il faut ensuite mettre à jour le contenu des dépôts, puis s'occuper du cas de Java. En effet, les dépôts de Raspbian proposent la version 9 qui n'est pas encore supportée par Unifi Controller. En attendant, deux choix s'offrent à nous : OpenJDK 8 ou la version d'Oracle. Cette dernière étant la plus performante, nous la choisissons :
sudo apt update
sudo apt install oracle-java8-jdk
Notez que si vous préférez OpenJDK, la procédure à suivre est :
sudo apt update
sudo apt remove openjdk-9-jre-headless
sudo apt install openjdk-8-jre-headless
