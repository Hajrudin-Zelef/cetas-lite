---
id: collect-260926-mikrotik/mikrotik/configuration-du-vpn-wireguard-pour-les-routeurs-mikrotik
title: "Configuration du VPN WireGuard pour les routeurs MikroTik"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/lot-mikrotik/forum/wireguard/configuration-du-vpn-wireguard-pour-les-routeurs-mikrotik.md
source_anchor: ""
source_lines: [1, 122]
sha256: d4c0f4fa9238821260177d3c8c9ba0e0497e1bbbd42410de0e2f686f1ab9121c
---

# Configuration du VPN WireGuard pour les routeurs MikroTik

pour les utilisateurs de VPN Unlimited

## Comment installer VPN Unlimited sur différentes plateformes et créer un identifiant KeepSolid


**Guides d'installation détaillés pour différentes plateformes :**


**Comment créer un nouvel identifiant KeepSolid ID**

1. Une fois l'application VPN Unlimited installée, vous verrez apparaître la ligne **« Créer un identifiant KeepSolid »** ; cliquez dessus pour lancer la procédure.
2. Saisissez votre adresse e-mail dans le champ prévu à cet effet.
3. Créez un mot de passe et saisissez-le dans le champ prévu à cet effet.
4. Saisissez à nouveau votre mot de passe dans le champ prévu à cet effet pour le confirmer.
5. Cochez la case à côté de la mention **« En continuant, vous acceptez nos…** ».
6. Appuyez sur le bouton **« S'inscrire »** pour terminer l'inscription.

Lors de votre première connexion à un serveur VPN, il vous sera demandé d'autoriser l'ajout de configurations VPN et le téléchargement d'un profil VPN. Pour plus de détails, veuillez vous reporter aux manuels fournis ci-dessus concernant l'installation de VPN Unlimited sur différentes plateformes.


**Important !** Veuillez noter que vous devrez configurer vous-même votre appareil à l’aide des paramètres générés, à vos propres risques.

## I. Créer une configuration VPN pour WireGuard


Avant de pouvoir configurer WireGuard® sur votre routeur MikroTik, vous devez générer une configuration VPN dans votre espace utilisateur KeepSolid.

Pour ce faire, suivez les quelques étapes simples décrites dans le tutoriel « Comment créer manuellement des configurations VPN ».

## II. Accéder à la console Web MikroTik ou à WinBox :


1. **Console Web :** ouvrez un navigateur Web et saisissez 192.168.88.1 dans la barre d'adresse. Connectez-vous à l'aide du nom d'utilisateur et du mot de passe de votre routeur.
2. **WinBox :** téléchargez et ouvrez l'outil WinBox depuis le site Web de MikroTik. Connectez-vous à l'aide de l'adresse IP, du nom d'utilisateur et du mot de passe de votre routeur.


## III. Préparation de la configuration WireGuard :


1. Ouvrez le fichier .conf fourni par VPN Unlimited à l'aide d'un éditeur de texte (tel que le Bloc-notes ou TextEdit). Laissez-le ouvert pour pouvoir vous y référer, car vous devrez copier des valeurs à partir de ce fichier.


## IV. Créer une interface WireGuard :


1. Dans le menu du routeur, accédez à **WireGuard** >**WireGuard** et cliquez sur «**Ajouter un nouvel élément »** .

2. Définissez le nom sur « **VPN-Unlimited** ».

3. Pour la clé privée, copiez-collez la clé figurant dans la section « Interface » de votre fichier .conf.


## V. Configurer l'adresse IP :


1. Accédez à **IP** >**Adresse** .
2. Cliquez **sur****« Ajouter »** et saisissez l'adresse IP et le réseau. Utilisez l'adresse indiquée dans la section Interface du fichier .conf. Pour le réseau, remplacez le dernier octet de votre adresse IP par zéro, afin d'obtenir X.X.X.0.

## VI. Ajouter un pair WireGuard :


1. Cliquez sur l'onglet **WireGuard** >**Peers** . Cliquez sur**« Ajouter un nouveau »** , saisissez les paramètres du fichier .conf, puis cliquez sur**Appliquer** >**OK** .

## VII. Configurer les routes pour le point d'extrémité WireGuard :


1. Accédez à **IP** >**Routes** et cliquez sur**« Ajouter »** .
2. Dans le champ « Adresse de destination », utilisez l'adresse du point de terminaison indiquée dans votre fichier .conf, en ajoutant un masque /32.
3. Définissez la passerelle sur votre passerelle par défaut, que vous trouverez sous IP > Client DHCP.
4. Ajoutez un commentaire tel que « **wgserver »** et marquez la route comme**NON activée** .

## VIII. Configurer la route par défaut pour le trafic VPN :


1. Dans **IP** >**Routes** , cliquez à nouveau sur**« Ajouter »** .
2. Définissez l'adresse de destination sur 0.0.0.0/0 et la passerelle sur %VPN-Unlimited.
3. Cochez la case « **Activé »** pour cette route.

**Remarque :** il se peut que votre connexion Internet cesse de fonctionner à ce stade, jusqu’à ce que le VPN soit entièrement configuré.

## IX. Configurer un client DHCP pour WireGuard :


1. Accédez à **IP** >**Client DHCP** .
2. Modifiez la configuration de votre client DHCP pour y inclure le script fourni.

**Script :**

:local route [/ip route find comment="wgserver"]

:if ($bound=1) do={

/ip route set $route gateway=$"gateway-address" disabled=no

} else={

/ip route set $route disabled=yes

}


3. Ce script permet de gérer l'acheminement de la connexion VPN en fonction de son état.

## X. Créer une liste d'interfaces pour le VPN :

1. Accédez à **Interfaces** >**Liste des interfaces** , puis cliquez sur**« Ajouter »** .
2. Sélectionnez « WAN » comme liste et « VPN-Unlimited » comme interface.

## XI. Configuration du DNS :


1. Accédez à **IP** >**DNS** .
2. Saisissez les serveurs DNS indiqués dans la section Interface de votre fichier .conf.
3. Cochez la case **« Autoriser les requêtes à distance »** pour activer la résolution DNS via le VPN.

Et voilà ! Vous avez configuré avec succès le VPN WireGuard sur votre routeur MikroTik. Vous pouvez désormais profiter des avantages du service VPN Unlimited, ainsi que de la vitesse et de la sécurité offertes par le protocole WireGuard.

Si vous avez besoin d'aide, n'hésitez pas à contacter notre service client à l'adresse [email protected].

Optimisez votre expérience sur le Web, protégez vos données sensibles et défendez-vous contre les menaces en ligne et les pirates informatiques.
