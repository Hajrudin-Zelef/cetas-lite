---
id: collect-261001-fortinet/fortinet/pare-feu-fortigate-comment-restaurer-la-configuration-dusine-41a0af6f
title: "pare-feu-fortigate-comment-restaurer-la-configuration-dusine-41a0af6f"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/pare-feu-fortigate-comment-restaurer-la-configuration-dusine-41a0af6f.md
source_anchor: ""
source_lines: [1, 22]
sha256: a4d62ec7688d04cd15c95bd12493480d2334eaecd9cb11a3a21f919460ca6a08
---

# pare-feu-fortigate-comment-restaurer-la-configuration-dusine-41a0af6f

Pare-feu FortiGate : comment restaurer la configuration d’usine ?
I. Présentation
Dans ce tutoriel, nous allons voir comment restaurer la configuration d'usine d'un pare-feu FortiGate.
La fonction de restauration des paramètres d'usine d'un équipement permet de supprimer les configurations apportées par l'administrateur et de rétablir les paramètres définis par le fabricant de ce dernier. Cette option est particulièrement intéressante dans le cas où on acquière un équipement en seconde main ou qu'on souhaite simplement le céder à un tiers. Ceci peut également correspondre à une action de remédiation suite à la compromission d'un équipement.
Nous allons voir comment implémenter concrètement cette action sur un pare-feu FortiGate.
Remarque : si vous envisagez de remettre en service ce pare-feu, il peut s'avérer judicieux de procéder à une sauvegarde de la configuration avant la réinitialisation.
II. Procédure
D'abord, on se connecte au pare-feu à l'aide d'un PC et d'un câble console. On lance ensuite le logiciel PuTTY. Puis, on initie la connexion en ligne de commande. On se connecte à l'aide d'un des comptes administrateurs du pare-feu. En cas d'oubli ou de non-connaissance du mot de passe administrateur, veuillez vous référer au tutoriel ci-dessous pour la procédure à suivre :
La commande permettant de lancer la réinitialisation du pare-feu est execute factoryreset. On saisit la commande et on valide. Le système nous averti alors que toutes les données et les configurations présentes dans l'équipement seront effacées. Il nous demande ensuite si l'on souhaite continuer : on valide avec la lettre y pour "yes" et la réinitialisation peut commencer. Elle dure quelques minutes... Patientez.
Meister-Gate # execute factoryreset
This operation will reset the system to factory default!
Do you want to continue? (y/n)y
System is reseting to factory default...
Le pare-feu redémarre une fois la réinitialisation effectuée.
D'abord, on peut remarquer ici que l'équipement a retrouvé son hostname par défaut, c'est-à-dire FortiGate-VM64-KVM (initialement changé en Meister-Gate). Le login à utiliser pour se connecter au pare-feu est désormais le login par défaut (admin). Le mot de passe à insérer ici est vide (juste valider avec la touche Entrée du clavier). Après la première connexion, le pare-feu nous invite à créer un mot de passe plus robuste.
Tout ceci démontre clairement que l'équipement a bien été restauré aux paramètres d'usine.
Remarque 1 :
Il existe également une deuxième commande appelée execute factoryreset2. Cette dernière, par contre, permet de conserver certains paramètres après la réinitialisation, à l'exemple des routes statiques permettant de garder l'accès à distance sur l'équipement.
Remarque 2 :
Une autre méthode consistant à réinitialiser l'équipement via l'interface web existe également. Le chemin de navigation peut varier en fonction de l'équipement à sa disposition (System > Maintenance > Factory Reset), mais le résultat sera le même qu'en ligne de commande. La méthode en ligne de commande reste par ailleurs la plus rapide.
III. Conclusion
Nous venons de voir ici comment réinitialiser un pare-feu FortiGate. Vous avez désormais la possibilité de le reconfigurer de zéro pour une nouvelle utilisation, ou de faire plaisir à quelqu'un en lui offrant !
