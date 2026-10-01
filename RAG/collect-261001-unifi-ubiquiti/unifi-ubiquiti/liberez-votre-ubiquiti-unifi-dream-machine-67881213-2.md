---
id: collect-261001-unifi-ubiquiti/unifi-ubiquiti/liberez-votre-ubiquiti-unifi-dream-machine-67881213-2
title: "liberez-votre-ubiquiti-unifi-dream-machine-67881213"
domain: unifi-ubiquiti
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-unifi-ubiquiti/liberez-votre-ubiquiti-unifi-dream-machine-67881213.md
source_anchor: ""
source_lines: [57, 73]
sha256: cd4c27e67cc937464866d6f8a828726b0ead5bb762ba0a847918447977f112dc
---

# liberez-votre-ubiquiti-unifi-dream-machine-67881213

Unifi propose par défaut de réaliser des sauvegardes périodiques de la configuration du contrôleur, avec une gestion de la rotation de celles-ci. L'inconvénient est que les sauvegardes sont réalisées en local uniquement, ce qui est n'est pas optimal au cas où votre équipement vient à rendre l'âme.
Pour remédier à ce détail non négligeable, un projet existe pour envoyer les sauvegardes du contrôleur sur un FTP. Nous allons mettre en place cette solution et par la même occasion en profiter pour exporter les données d'AdGuard Home.
Pour commencer, récupérons le script 80-udm-backup-ftp.sh, toujours sur le même dépôt Github. 
Editons ce script avec la commande vi /mnt/data/on_boot.d/80-udm-backup-ftp.sh pour définir le serveur FTP (ainsi que les informations d'identification nécessaires) sur lequel les sauvegardes vont être effectuées.
Il est également possible de configurer la fréquence d'exécution du script, qui doit être en adéquation avec votre fréquence de sauvegarde du contrôleur. Par défaut, le script est configuré pour être exécuté toutes les 30 minutes : CRON_SCHEDULE='30 * * * *'. Pour mieux comprendre comment écrire la planification d'un cron, je vous dirige vers le site crontab guru.
La dernière personnalisation du script que nous pouvons faire est de commenter les lignes indiquant les éléments que vous ne souhaitez pas sauvegarder, en ajoutant un # en début de ligne, comme pour la ligne qui indique le chemin des fichiers d'Unifi Protect.
Nous allons rendre le script exécutable, télécharger l'image podman puis lancer le script avec les commandes suivantes :
Vous venez de regarder le répertoire de destination de vos sauvegardes sur votre serveur FTP mais celui-ci est vide... C'est normal, vous avez seulement configuré le lancement du conteneur qui va pousser les sauvegardes vers votre serveur FTP, il faut maintenant attendre la première exécution automatique de l'export suivant la fréquence que vous avez configuré.
Si vous ne trouvez aucun fichier après la première planification, vous pouvez consulter les logs d'exécution avec la commande suivante : 
Maintenant que la sauvegarde du contrôleur est exportée, nous allons voir comment gérer la rotation des sauvegardes, n'ayant pas besoin de les garder indéfiniment.
J'utilise mon NAS Synology comme serveur FTP, et il est possible de supprimer les fichiers les plus vieux d'un dossier spécifique en mettant en place une tâche planifiée. On va définir une programmation correspondante à la fréquence des sauvegardes du contrôleur, après la sauvegarde et son export en FTP.
Pour ma part, je fais une sauvegarde du contrôleur Unifi une fois par semaine le dimanche à 2h00, j'ai donc créé la tâche sur le Nas pour tourner tous les lundi à 3h.
Maintenant, la partie la plus importante, le script qui va lister et supprimer les fichiers.
Vous devez bien entendu spécifier le chemin de vos sauvegardes et l'utilisateur avec lequel sont créées les sauvegardes, pour limiter la suppression de mauvais fichiers. Vous pouvez recevoir un mail de rapport du script si vous le souhaitez.
Vous pouvez simplement exécuter le script et voir le résultat de celui-ci grâce à l'interface du Synology.
Voilà, ce tutoriel est arrivé à sa fin, on a pu voir ensemble ce qu'était l'Ubiquiti Unifi Dream Machine et comment améliorer les services qu'il peut nous proposer, en hébergeant Adguard Home, et en externalisant les sauvegardes du contrôleur, pour s'assurer de pouvoir retrouver notre configuration réseau en cas de panne matérielle. Je vous conseille au minimum de suivre cette dernière partie.
Si vous avez des questions, n'hésitez pas à laisser un commentaire ou venir échanger avec nous sur notre groupe Telegram.
