---
id: collect-261001-general-networking/general-networking/fr-review-hp-elite-t755-review-powerful-thin-client-for-demanding-users-7182b627-3
title: "fr-review-hp-elite-t755-review-powerful-thin-client-for-demanding-users-7182b627"
domain: general-networking
role: reference
task: reference
actors: ["Nvidia"]
dates: []
keywords: ["arr", "gpu", "nvidia", "voice"]
source: docs/RAG/collect-261001-general-networking/fr-review-hp-elite-t755-review-powerful-thin-client-for-demanding-users-7182b627.md
source_anchor: ""
source_lines: [79, 109]
sha256: 9e5a3269e404be7f4fadbcd1a0f5fa7efa0ec32c8d6327e606502ab04e79f22f
---

# fr-review-hp-elite-t755-review-powerful-thin-client-for-demanding-users-7182b627

Après la mise sous tension du client, l'écran de démarrage HP est apparu et vous êtes automatiquement connecté à l'aide de l'utilisateur et du mot de passe utilisateur par défaut. Après une connexion réussie, l’écran Window IoT nous a été présenté. L'écran ressemblait à n'importe quel autre système Windows 10. Le menu Démarrer de Windows comprenait des icônes pour Citrix Receiver VMware Horizon Client et d'autres outils Windows standard.
De nombreux outils standard, tels que l'invite de commande, ont été désactivés et ne pouvaient être exécutés qu'en se connectant à l'appareil à l'aide du compte administrateur.
Le t755 est conçu dans un souci de sécurité et, à ce titre, HP Write Manager protège le contenu et réduit l'usure du lecteur flash en redirigeant et en mettant en cache les écritures vers un espace de stockage virtuel dans la RAM. Le cache est vidé lors d'un redémarrage du système et toutes les modifications apportées depuis le dernier démarrage du système sont définitivement perdues. Cela protège l'appareil contre les codes malveillants et les configurations non sécurisées. Comme il s'agissait d'un système de test, nous avons désactivé HP Write Manager lors de nos tests en accédant au Panneau de configuration et en utilisant l'outil de configuration HP Write Manager. Les modifications nécessitaient une connexion en tant qu'administrateur (mot de passe par défaut Admin). Le système a nécessité un redémarrage une fois les modifications apportées.
La configuration de l'appareil était la même que celle des autres systèmes Windows 10, mais seul le compte utilisateur Administrateur peut modifier définitivement l'appareil. Le compte utilisateur a un accès limité à de nombreux outils. Les outils standard tels que l'Explorateur de fichiers ne sont pas disponibles pour les utilisateurs réguliers.
L'appareil dispose de deux lecteurs locaux : C et Z. Le lecteur C : (protégé par le filtre d'écriture HP) est un lecteur flash et héberge le système d'exploitation et les applications. Le lecteur Z: est un lecteur RAM virtuel qui se comporte comme un lecteur physique mais est créé au démarrage du système et détruit à l'arrêt du système.
Nous avons installé ControlUp Edge DX pour surveiller le système pendant les tests. À titre de clause de non-responsabilité, l'un de nos analystes (Tom Fenton) travaille pour ControlUp.
Poste de travail Horizon local
Pour avoir une idée du fonctionnement de l'appareil dans le monde réel, nous l'avons utilisé avec un bureau virtuel Horizon local pour effectuer nos tâches quotidiennes pendant deux semaines.
Comme mentionné précédemment, le t755 était connecté à notre réseau avec un câble Cat 6 via un réseau 1GbE via un switch connecté à un serveur ou un routeur WAN. Le serveur hébergeait notre bureau virtuel VMware Horizon local tandis que le routeur WAN se connectait aux bureaux virtuels basés sur le cloud. Pour créer un environnement contrôlé, le réseau a été surveillé pendant les tests pour garantir qu'aucun autre trafic n'était présent.
Le bureau virtuel que nous avons utilisé fonctionnait sous Windows 10 et disposait de deux processeurs virtuels, de 8 Go de mémoire et de 128 Go de stockage basé sur NVMe.
Nous avons lancé le client Horizon et l'avons configuré pour se connecter à un poste de travail Horizon local. Nous étions liés au bureau virtuel à la résolution native 2K du moniteur.
Le premier test que nous avons effectué consistait à utiliser VLC pour lire une vidéo (1720 × 720 à 24 ips) stockée sur le bureau virtuel. Tout d’abord, nous avons lu la vidéo dans sa résolution native, puis en mode plein écran. En mode natif et plein écran, la vidéo est lue sans aucune perte d'image. Le son est parfaitement diffusé via le haut-parleur intégré de l'appareil, en mode natif et plein écran. Le haut-parleur intégré était suffisamment puissant pour que nous puissions l'entendre, mais vous auriez besoin d'un casque ou de haut-parleurs externes dans un environnement de bureau.
Au cours de ce test, Edge DX a montré qu'environ 4 % du processeur de l'appareil et 1.25 Go de RAM étaient utilisés, tandis que le bureau virtuel montrait que plus de 90 % de son processeur était utilisé. Cela indiquait que la gigue provenait du bureau virtuel et non du client et que le client pouvait gérer une charge exigeante à partir d'un bureau virtuel sans stresser ses ressources.
Nous avons connecté un casque Jabra Voice 150 à l'un des ports USB situés à l'avant de l'appareil ; le bureau virtuel a découvert le casque Jabra et a fonctionné sans problème avec une bonne qualité sonore.
Aucun problème n’a été rencontré au cours de la fenêtre de test de deux semaines. Cela incluait l'utilisation d'applications bureautiques, telles que le navigateur Web Chrome, la lecture de musique en streaming sur Internet, etc.
Test d'un bureau virtuel basé sur le cloud et activé par GPU
Aux débuts du VDI, la base d’utilisateurs était principalement composée de travailleurs qui utilisaient une ou deux applications bureautiques. Nous voyons désormais des entreprises utiliser des applications nécessitant des GPU. Pour voir à quel point cela fonctionnait, nous avons utilisé VMware TestDrive pour nous connecter à un ordinateur de bureau compatible GPU, puis avons exécuté des applications à forte intensité vidéo.
Nous avons joué une vidéo haute définition à partir d'un navigateur Web sur les bureaux virtuels.
Le processeur du client a atteint 2 %, tandis qu'environ 0.45 Mbps de données étaient transférées. Le processeur sur le bureau virtuel était d'environ 37 pour cent. L'audio et la vidéo ont été parfaitement lus et même si le moniteur ne prend en charge que le QHD (2560 1440 x XNUMX XNUMX), la sortie était fantastique sur le moniteur HP.
Nous avons lancé la démo NVIDIA FaceWorks et elle s'est déroulée sans problème. Bien que le bureau virtuel utilise 42 % de la puissance de son processeur et que le GPU atteigne sa capacité maximale, l'appareil a pu gérer la charge et l'utilisation du processeur est restée inférieure à 3 %.
Configuration à deux moniteurs
Il n'est pas rare que les utilisateurs disposent de plusieurs moniteurs. Pour tester dans quelle mesure l'appareil gère une configuration à deux moniteurs, nous avons connecté un moniteur 4K au client via le DisplayPort sur le GPU. L'appareil a immédiatement reconnu le moniteur et nous avons configuré le deuxième affichage en mode Paysage.
Sur le moniteur 4k, nous avons affiché un bureau VMware TestDrive Horizon ; de l'autre, nous avons exécuté notre bureau virtuel Horizon local. Nous avons diffusé différentes vidéos sur chacun des moniteurs en même temps.
Les deux vidéos ont été lues sans gigue et Edge DX a montré que seulement 4 % environ du processeur du client et 2.3 Mo/s de trafic réseau étaient envoyés à l'appareil.
Configuration du moniteur Penta
Nous n'avons jamais testé de client léger avec plus de trois écrans auparavant, mais voyant comment ce client léger peut en prendre en charge jusqu'à six, nous en avons connecté cinq ! Nous avons connecté trois moniteurs à l'aide du DisplayPort natif et utilisé les ports GPU pour rejoindre les deux autres. Nous avons utilisé le moniteur Dell P4317Q pour afficher quatre écrans car il prend en charge jusqu'à quatre moniteurs PIP sur un seul écran.
Nous avons configuré les écrans et en avons pris une capture d'écran.
Après avoir vérifié qu'il prenait en charge cinq écrans 2K, nous sommes revenus à une configuration d'affichage unique.
Gestion des appareils
L'appareil peut être utilisé avec les outils Windows standard ou HP Device Manager (HPDM) pour une administration centralisée et basée sur le serveur de ce client. Cependant, l’utilisation de HPDM dépasse le cadre de cette revue.
Conclusion
