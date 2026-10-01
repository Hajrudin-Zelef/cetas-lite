---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/remplacer-bbox-remplacer-bbox-par-unifi-dream-machine-pro-internet-tel-ok-12-6da00aad
title: "remplacer-bbox-remplacer-bbox-par-unifi-dream-machine-pro-internet-tel-ok-12-6da00aad"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/remplacer-bbox-remplacer-bbox-par-unifi-dream-machine-pro-internet-tel-ok-12-6da00aad.md
source_anchor: ""
source_lines: [1, 15]
sha256: ab3856e1cb0fb1e7037f0adb330896b057f965c8e30116ca10c3ac22946009f1
---

# remplacer-bbox-remplacer-bbox-par-unifi-dream-machine-pro-internet-tel-ok-12-6da00aad

Je suis également équipé de matériel Unifi : un AP et un Switch , mon contrôleur est hébergé sur un NAS Qnap et j'ai un compte Ubiquiti. Mes questions avec L'UDM:
Le contrôleur est hébergé sur l'UDM et la config initiale se fait avec un mobile en Bluetooth. Dans mon cas avec le contrôleur hébergé sur mon Qnap je peux avoir accès au contrôleur via le nom de domaine du Nas :
xxxx.myqnapcloud.com:38193 (38193 étant le port par défaut) aussi bien en local ou d'un PC distant avec un navigateur que d'un mobile avec l'appli Unifi Network et ce même en 4G .
Comment cela se passe avec l'UDM pour avoir un accès externe au contrôleur?
Hello
Le controleur de l'UDM pro va prendre naturellement le relai de ton contrôleur hébergé sur ton NAS. 
Il y a une méthode pour faire un backup et un restore mais le mieux est de repartir de zéro (d'autant plus que tu n'as pas une configuration très complexe a priori). 
J'ai par ailleurs, la même configuration avec un NAS QNAP et voici ce que j'ai fait :
-Achat d'un nom de domaine chez ionos avec certificat ssl wildcard * inclus pour 1,2 euros TTC la première année (promo encore valable quand j'écris ces lignes). 
-Ajout du certificat sur mon NAS
-UDM Pro qui héberge le contrôleur et ajout du certificat sur mon UDM Pro. 
-Hébergement sur mon NAS de Bitwarden server pour gérer mes mots de passe, via docker. Cette solution intègre nginx, un reverse proxy, sécurisée par ssl avec le certificat. 
-Parametrage de nginx pour rediriger mes pages de NAS vers nas.monsite.fr, de mon udm pro vers ubnt.monsite.fr et de mon serveur bitwarden vers bitwarden.monsite.fr. 
Je peux dans la configuration de nginx interdire l'accès vers l'udm pro uniquement aux ressources sur mon réseau. 
Je trouve cette solution sécurisée et plus propre.
