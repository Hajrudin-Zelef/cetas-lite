---
id: collect-261001-cisco/cisco/opnsense-mise-en-place-dun-pare-feu-0c72dee5-2
title: "OPNsense : mise en place d’un pare-feu"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-cisco/opnsense-mise-en-place-dun-pare-feu-0c72dee5.md
source_anchor: ""
source_lines: [5, 89]
sha256: 7804a721a341af08a36ad8935d144604bd55cc77e0dad08ca5e35677c5fe8a7d
---

# OPNsense : mise en place d’un pare-feu

Dans ce tutoriel, je vais vous expliquer comment mettre un place OPNsense qui est une distribution qui permet de mettre en place un pare-feu périmétrique.
OPNsense est fork basé sur PFsense qui est lui même un fork de m0n0wall basé sur FreeBSD.
OPNsense intègre les fonctionnalités standard des firewall opensource :
OPNsense peut embarqué par un Greffon (nom des plugins) Crowdsec, qui va vous permettre de protéger toutes les publications sur Internet avec la base d’adresse IP communautaire.
Pour installer OPNsense, vous avez besoin d’un ordinateur (physique ou virtuel) avec deux cartes réseaux minimum :
Dans ce tutoriel, mon interface WAN sera en DHCP
Concernant les ressources CPU / RAM et disques, cela va dépendre principalement du nombre de périphériques et de fonctionnalités que vous allez activer.
Pour l’espace disque prévoir à minima 20Go.
Ici, je vais utilisé un ordinateur virtuel sous VMWare.
Commencer par télécharger l’ISO OPNsense : https://opnsense.org/download/
Particularité en OPNsense, l’installation se déroule en deux temps :
Commencer par démarrer sur l’ISO :
A un moment du démarrage, on vous demande d’appuyer sur une touche pour configurer les interfaces réseaux (LAN et WAN) :
Ensuite on vous demande si vous souhaitez configurer un LAG et des VLAN, dans mon cas je répond Non (N) au deux questions :
On arrive maintenant à la configuration des Interfaces réseaux. commencer par entrer le nom de la carte réseau WAN (em1) et le nom de l’interface LAN (em0) puis entrer rien pour l’interface optionnelle.
Confirmer l’application de la configuration des interfaces (Y) :
Une fois la configuration terminée, on récupère le prompt sur login :
Sur la capture, on peut voir que les adresses IP attribués aux interfaces.
En fonction de votre environnement, les adresses IP peuvent vous convenir, notamment si vous avez un service DHCP sur les deux interfaces.
Dans mon cas, l’interface LAN à une adresse IP en 192.168.1.X/24, ce qui me pose problème car cela est sur la même plage IP que ma box et je ne peux donc pas accéder à l’administration.
Je vais donc vous montrer comment changer l’adresse IP.
Commencer par vous identifiant en avec le compte root et le mot de passe opnsense.
Par défaut, on arrive sur un menu qui nous permet de faire plusieurs opérations « facilement ».
Pour modifier une adresse IP, entrer l’option 2.
Entrer le numéro de l’interface que vous souhaitez configurer.
Plusieurs « questions » sont demandés pour configurer l’interface réseau :
Ensuite vous pouvez configurer l’IPv6 si nécessaire.
Patienter pendant l’application de la configuration, une fois terminer, on vous indiquer l’url de l’interface Web de OPNsense.
Avant de faire l’installation définitive sur le disque, on va suivre l’assistant de configuration sur l’interface Web.
Depuis un navigateur Internet, aller sur l’interface Web du firewall.
Connecter vous avec le compte root et mot de passe opnsense.
Patienter pendant le chargement de l’assistant…
Au lancement de l’assistant, cliquer sur le bouton Next 1.
Configurer les informations de l’hôte : nom, domaine et lange de l’interface et serveur DNS 1 puis valider en cliquant sur Next 2.
Indiquer le fuseau horaire 1 et cliquer sur Next 2.
Si nécessaire modification la configuration de l’interface WAN et cliquer sur Suivant.
Faire de même pour l’interface LAN et cliquer sur Suivant.
On fini par le changement du mot de passe du compte root et cliquer sur Suivant.
Pour appliquer la configuration, cliquer sur Recharger 1.
Une fois le rechargement de la configuration effectué, une page indiquant que OPNsense est configuré s’affiche.
Sur la tableau de bord (Dashboard) du firewall, un message d’alerte s’affiche indiquant que nous sommes en LiveCD.
On va maintenant rendre persistant notre firewall.
Retourner sur la console en Shell et cette fois ci, identifier vous avec le compte installer et le même mot de passe que root.
Commencer par configurer le disposition du clavier.
Choisir ensuite la tâche Install (UFS).
Si un message indiquant que les prérequis ne sont pas respecter (CPU / RAM) valider en cliquant sur Proceed anyway.
Sélectionner le disque dur où sera installé OPNsense.
Confirmer la sélection du disque.
Patienter pendant le clonage du système du le disque …
Valider en sélectionner Complete Install.
Le système redémarre …
Penser à retirer l’ISO
De retour sur l’interface Web, on peut voir que le message d’alerte indiquant que nous sommes en liveCD à disparu.
Dans cette partie, je ne vais pas rentrer dans le détail complet de l’utilisation du firewall OPNsense (on verra ça dans d’autres tutoriels).
Ici on va faire un tour « rapide » pour le rendre fonctionnel avec la mise en place du service DHCP et la mise en place d’un proxy transparent.
Par défaut, le firewall est configuré de la manière suivante :
Les règles du pare-feu sont accessible dans le menu : Pare-feu -> Règle -> (INTERFACE)
Règles LAN :
Règles WAN :
Comme on peut le voir, il n’y a pas de règle WAN, car de façon implicite, tout le trafic venant d’internet vers le LAN est bloqué.
Dans mon environnement de LAB, je n’ai pas de serveur DHCP, on va donc configuré ce service.
Depuis le menu de navigation, aller sur Services 1 et dérouler DHCPv4 2 et cliquer sur LAN 3.
Cocher la case Activer le serveur DHCP sur l’interface LAN 1, puis configurer la plage DHCP 2 et indiquer le ou les serveurs DNS 3 et la Passerelle 4.
Aller en bas de la page et cliquer sur le bouton Sauvegarder 1.
Le serveur DHCP est configuré sur OPNsense.
Maintenant, on va pouvoir tester le pare-feu, pour cela j’ai ajouté un ordinateur virtuel dans la partie LAN.
J’ai oublié la capture de la configuration IP.
Pour vérifier le fonctionnement du DHCP, il suffit d’aller voir les baux : Services / DHCPv4 / Baux
Sur la capture, on voit l’adresse qui a été attribué à mon ordinateur.
On peut vérifier le trafic entre le LAN et le WAN directement depuis le firewall, Windows étant « bavard », on va pouvoir voir le trafic, depuis le menu aller : Rapports / Trafic.
Sur la capture, on peut voir le trafic sortant.
Dans ce partie, je vais expliquer comment configurer le Proxy en HTTP et HTTPS en mode transparent.
Le mode transparent contrairement au mode implicite ne nécessite pas de d’indiquer au ordinateur client les paramètres du proxy.
Le pare-feu va intercepter la requête HTTP/HTTPS pour ensuite l’envoyer au service Proxy (SQUID).
Le mode transparent nécessite le déploiement d’un certificat SSL pour le trafic HTTPS, car du passage par le proxy, le trafic est chiffré avec le certificat du firewall.
Depuis le menu, aller sur : Système / Gestion des Certificats / Autorités, sur la page de gestion de CA, cliquer sur bouton + 1 pour en créer une nouvelle.
Remplir tous les champs obligatoire 1 et cliquer sur le bouton Sauvegarder 2.
L’autorité est créée.
On va maintenant passer à la configuration du Proxy.
Pour être franc, ce OPNsense ce n’est pas une partie de plaisir par rapport à d’autre solution …
On va commencer par activer le Proxy Web, depuis le menu aller sur Services / Proxy Web / Administration.
Sur l’onglet Réglages Proxy généraux, cocher la case Activer le proxy 1.
Pour accéder à plus d’option, activé le mode avancé 1.
Aller ensuite sur l’onglet Forward Proxy 1, cocher les cases suivantes : Activer le proxy HTTP Transparent 2 et Activer l’inspection SSL 3. Sélectionner l’autorité de certificat 4 que l’on a créé précédemment puis valider les paramètres en cliquant sur Appliquer 5.
Sur les autres Firewall que j’ai l’habitude de manipuler, normalement la configuration est terminée mais pas sur OPNsense, il faut maintenant créer deux règles pour faire passer le trafic dans le proxy.
