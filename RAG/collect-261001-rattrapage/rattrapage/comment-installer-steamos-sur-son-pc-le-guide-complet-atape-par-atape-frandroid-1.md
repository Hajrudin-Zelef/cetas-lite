---
id: collect-261001-rattrapage/rattrapage/comment-installer-steamos-sur-son-pc-le-guide-complet-atape-par-atape-frandroid-1
title: "comment-installer-steamos-sur-son-pc-le-guide-complet-atape-par-atape-frandroid"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Microsoft", "Nvidia"]
dates: []
keywords: ["amd", "ethernet", "intel", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/comment-installer-steamos-sur-son-pc-le-guide-complet-atape-par-atape-frandroid.md
source_anchor: ""
source_lines: [1, 74]
sha256: 57a7dee485d4ec7d094730f0fc0d30f6776c0248b3ce548e33318a66276a8471
---

# comment-installer-steamos-sur-son-pc-le-guide-complet-atape-par-atape-frandroid

Transformer un PC en console de salon faÃ§on Steam Deck prend dÃ©sormais environ un quart d’heure. L’opÃ©ration ressemble Ã une installation de Linux classique, avec deux ou trois rÃ©glages qui font toute la diffÃ©rence entre une machine qui dÃ©marre du premier coup et un Ã©cran noir frustrant. On dÃ©roule tout, sans rien laisser au hasard. Si vous n’avez pas encore la machine, on vous explique justement comment construire sa Steam Machine en 8 composants, avant de lui installer son systÃ¨me. Si vous hÃ©sitez encore entre la monter ou acheter celle de Valve, notre comparatif Steam Machine contre PC Ã 1Â 000Â â¬ pose le dÃ©cor.

## Quelle configuration pour faire tourner SteamOS ?

Le point dÃ©terminant de la configuration d’un tel PC gamer, c’est la carte graphique. SteamOS s’appuie sur les pilotes ouverts d’AMD et MesaÂ : une Radeon RX 6000 (RDNA 2) ou RX 7000 (RDNA 3) offre le meilleur support, puisque ce sont les architectures du Steam Deck et de la Steam Machine. Les cartes Intel Arc de bureau, type B580, fonctionnent dÃ©sormais aussi, mais via un bricolage communautaire : l’installation officielle ne les gÃ¨re pas encore, et la mise en route documentÃ©e par des testeurs passe par une carte AMD avant d’Ã©changer pour l’Arc. Les GeForce, en revanche, ne sont pas encore gÃ©rÃ©es, on y revient plus bas. CÃ´tÃ© reste de la machine, visez un SSD NVMe dÃ©diÃ© Ã SteamOS, 16Â Go de RAM pour Ãªtre Ã l’aise, et gardez en tÃªte qu’une carte rÃ©cente exige une image rÃ©cente, sous peine de manquer les pilotes Mesa adaptÃ©s.

Pour une base qui colle Ã la Steam Machine sans se ruiner, notre sÃ©lection de composants pour une Steam Machine maison sert de point de dÃ©part tout trouvÃ©.

Un dÃ©tail de vocabulaire pour cadrer les attentes : le support officiel d’Intel introduit avec SteamOS 3.8 vise d’abord les consoles portables, comme la MSI Claw, et non les cartes graphiques Arc de bureau. Faire tourner SteamOS sur un PC fixe Ã carte Arc reste pour l’instant l’affaire de bricoleurs, et Valve n’a annoncÃ© aucune date pour un support natif des Arc dÃ©diÃ©es.

    Pour aller plus loin

            Nos conseils pour construire sa Steam Machine : 8 composants, moins cher et un piÃ¨ge
            

## Ce qu’il vous faut avant de commencer

- Une clÃ© USB de 8 Go minimum, 16 Go pour Ãªtre tranquille. Son contenu sera effacÃ©.
- L’image de rÃ©cupÃ©ration SteamOS, Ã rÃ©cupÃ©rer sur la page de support officielle de Valve.
- Un outil de gravure : Rufus sur Windows, Balena Etcher sur macOS et Linux, ou la commande dd pour les habituÃ©s.
- Un clavier et une souris pour les Ã©crans d’installation, mÃªme si vous finirez Ã la manette.
- Une connexion Ethernet de prÃ©fÃ©rence, plus fiable que le Wi-Fi pendant l’installation.
- Une sauvegarde de vos fichiers : le disque cible sera entiÃ¨rement rÃ©Ã©crit.

Si vous n’avez jamais touchÃ© Ã Linux, rien d’insurmontable : la logique d’une clÃ© USB bootable et d’un disque Ã prÃ©parer est la mÃªme partout. Notre tuto pour installer Linux sur un PC dÃ©taille ces bases, utiles avant de se lancer dans SteamOS.

## Ãtape 1 : crÃ©er la clÃ© USB d’installation

TÃ©lÃ©chargez l’image sur la page d’installation et de rÃ©paration de SteamOS, la seule source officielle.

Souris, clavier et casque gaming : la sÃ©rie G3 de Logitech G combine prÃ©cision, confort et personnalisation RGB. De plus, G HUB vous permet d’adapter chaque rÃ©glage Ã votre style et de profiter pleinement de vos jeux.

Le fichier pÃ¨se plusieurs gigaoctets, une connexion stable aide. Branchez la clÃ©, ouvrez Rufus ou Balena Etcher, sÃ©lectionnez l’image puis la clÃ©, et lancez l’Ã©criture. VÃ©rifiez bien que vous pointez sur la clÃ© USB et pas sur un autre disque, c’est l’erreur classique qui efface les mauvaises donnÃ©es.

La manipulation est identique Ã celle d’une clÃ© d’installation Linux classique, si vous l’avez dÃ©jÃ faite une fois. Une fois l’Ã©criture terminÃ©e, Ã©jectez proprement la clÃ©.

    Pour aller plus loin

            Comment installer Linux sur un PC Windows 10 : le guide complet
            

## Ãtape 2 : prÃ©parer le BIOS

C’est l’Ã©tape Ã ne pas bÃ¢cler. RedÃ©marrez et entrez dans le BIOS, en gÃ©nÃ©ral via la touche Suppr ou F2. DÃ©sactivez le Secure Boot : SteamOS n’est pas signÃ© avec les clÃ©s de Microsoft, il refuserait de dÃ©marrer sinon. Restez en mode UEFI et coupez le CSM si l’option existe. Si vous le pouvez, dÃ©branchez physiquement les autres disques de la machine le temps de l’installation, c’est le moyen le plus sÃ»r de ne pas Ã©craser le mauvais. Enregistrez les changements, puis quittez.

## Ãtape 3 : lancer l’installation

1. Au dÃ©marrage, ouvrez le menu de boot (souvent F12, F11, Ãchap ou Suppr selon la carte mÃ¨re) et choisissez la clÃ© USB.
2. L’installateur charge quelques lignes Linux, puis affiche un bureau KDE. Lancez l’option qui efface le disque et installe SteamOS.
3. Confirmez l’effacement du disque cible, en vÃ©rifiant une derniÃ¨re fois que c’est le bon. Comptez environ un quart d’heure.
4. Ã la fin, retirez la clÃ© avant le redÃ©marrage, pour ne pas relancer l’installateur.
5. SteamOS dÃ©marre sur sa configuration initiale : langue, rÃ©gion, clavier, puis rÃ©seau, et enfin connexion Ã votre compte Steam.

Si l’installateur refuse de dÃ©marrer, revÃ©rifiez que le Secure Boot est bien dÃ©sactivÃ© et changez de port USB, les faÃ§ades avant posant parfois problÃ¨me.

En cas d’Ã©cran noir aprÃ¨s l’installation, contrÃ´lez l’ordre de boot et assurez-vous que le bon SSD est sÃ©lectionnÃ©. Si une carte rÃ©cente n’affiche rien, repartez d’une image SteamOS 3.8 ou plus rÃ©cente.

## Ãtape 4 : configurer le systÃ¨me

Premier rÃ©flexe avant tout : la mise Ã jour. Direction ParamÃ¨tres puis SystÃ¨me, et installez la derniÃ¨re version de SteamOS. La machine dÃ©marre ensuite directement dans l’interface manette du Steam Deck, avec la prÃ©compilation des shaders qui lisse les saccades.

Pour aller plus loin, basculez en mode Bureau via le menu d’alimentation : vous y trouvez un environnement KDE Plasma complet, pour installer des applications, des Ã©mulateurs ou un navigateur.

Quelques rÃ©glages amÃ©liorent le confort de jeu. Dans les paramÃ¨tres de compatibilitÃ©, activez Steam Play pour tous les titres, afin que Proton fasse tourner l’Ã©crasante majoritÃ© des jeux Windows. Ajoutez Decky Loader si vous voulez des extensions au mode jeu, et un lanceur comme Heroic ou Lutris pour brancher vos bibliothÃ¨ques Epic, GOG ou autres. Avant d’acheter ou de lancer un jeu, un coup d’Åil Ã ProtonDB renseigne sur sa compatibilitÃ©, et Are We Anti-Cheat Yet indique si son anti-triche passe sous Linux.

## Ce qu’il faut savoir avant de se lancer

- **Pas de double dÃ©marrage officiel** : l’installation prend tout le disque. Valve dit travailler sur un installateur capable de cohabiter avec Windows, mais ce n’est pas encore lÃ . Pour garder les deux, le plus propre reste un disque sÃ©parÃ© pour chaque systÃ¨me.
- **Les anti-triches au niveau du noyau bloquent** : Valorant, Call of Duty, Battlefield ou les jeux EA Sports refusent de se lancer sous Linux. Toute une catÃ©gorie de jeux compÃ©titifs reste inaccessible.
- **Pas de HDMI-CEC** : vous ne pourrez pas allumer la TV ni la piloter avec la manette comme sur une vraie console.
- **Le Secure Boot reste dÃ©sactivÃ©** : Ã  garder en tÃªte si vous rÃ©cupÃ©rez la machine pour un autre usage plus tard.

## Et si on a une carte Nvidia ?

