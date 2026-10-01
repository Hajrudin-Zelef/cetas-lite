---
id: collect-261001-general-networking/general-networking/fr-review-how-to-get-started-with-truenas-scale-75339f23-2
title: "fr-review-how-to-get-started-with-truenas-scale-75339f23"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/fr-review-how-to-get-started-with-truenas-scale-75339f23.md
source_anchor: ""
source_lines: [26, 55]
sha256: f03768f10df9b6e23d641c88437ec8ffade22a7897b465dcb45cd51a571fea4a
---

# fr-review-how-to-get-started-with-truenas-scale-75339f23

Maintenant que ces éléments fondamentaux sont éliminés, nous pouvons configurer le partage SMB réel à utiliser. Accédez à l'onglet Partages. Dans la barre latérale, recherchez les partages SMB et sélectionnez Ajouter.
Nous devons utiliser l'option Chemin pour sélectionner l'ensemble de données créé précédemment et le rendre disponible pour le partage. Ensuite, ajoutez un nom qui s’affichera une fois le service SMB accédé. L'option Objectif peut être configurée pour plusieurs utilisations, mais pour le stockage de fichiers normal, les paramètres de partage par défaut fonctionneront.
Après avoir configuré le partage SMB et cliqué sur Enregistrer, vous verrez la fenêtre contextuelle pour activer le service SMB. Si vous souhaitez que le service démarre automatiquement au démarrage de TrueNAS, sélectionnez le démarrage automatique avant d'activer le service. Si vous changez d'avis plus tard, vous pouvez également activer le démarrage automatique.
À ce stade, vous devriez pouvoir voir TrueNAS dans la section réseau de l'explorateur de fichiers Windows, mais en fonction de vos préférences, il y a quelques étapes supplémentaires. Si vous pouvez fournir à tous les utilisateurs l'accès à ce partage, ainsi que le nom par défaut, alors vous avez terminé. Sinon, le reste est assez simple. Sélectionnez simplement l'icône de partage à côté du commutateur activé sur le partage et modifiez la liste de contrôle d'accès (ACL). Par défaut, l'ACL accorde à chacun une autorisation complète. Cela peut convenir si vous êtes le seul utilisateur de l'instance, mais sinon, vous souhaiterez peut-être être un peu plus précis.
Vous pouvez soit ajouter une règle de refus spécifique pour empêcher un utilisateur d'accéder à un pool, soit supprimer l'entrée « tout le monde @ » et la configurer pour n'autoriser qu'un seul utilisateur, comme nous le montrons dans la capture d'écran ci-dessous. Cela restreint l’accès au pool uniquement à notre utilisateur storagereview. Vous pouvez y ajouter plus d'utilisateurs ou même des groupes entiers. Il est également possible d'avoir des utilisateurs en lecture seule et des utilisateurs spécifiques refusés.
Et enfin, pour les PME, la dernière étape (même si elle n'est pas critique) consiste à changer le nom du service. Pour ce faire, nous ouvrons le menu kebab (ou menu « trois points ») à côté du bouton Ajouter et cliquons sur configurer le service. Ici, nous pouvons simplement changer le nom NetBIOS pour modifier la façon dont le service SMB apparaît sur les appareils du réseau.
Après l'enregistrement, nous devrions voir le nom dans l'onglet réseau de l'Explorateur de fichiers sous Windows.
À partir de là, double-cliquez sur le nom et vous devriez voir une fenêtre d'authentification. Entrez les informations d'identification du compte que vous avez créé et défini dans l'ACL, et vous devriez pouvoir voir les partages autorisés pour cet utilisateur. Toutes nos félicitations! Vous disposez désormais d’un partage SMB simple et fonctionnel sur TrueNAS.
Si vous souhaitez consulter la documentation TrueNAS pour la configuration SMB, vous la trouverez ici.
Configuration iSCSI
Si vous êtes intéressé par la configuration iSCSI, vous pouvez parcourir cette section. Sinon, sautez-le.
iSCSI est une bête légèrement différente de SMB, et elle s'affichera également un peu différemment sous Windows. iSCSI se monte comme un disque local au lieu d'un partage réseau, mais il aime un peu moins être utilisé par plusieurs machines clientes à la fois. SMB est la meilleure option pour le partage avec plusieurs machines clientes. Soyez prêt : la configuration d'iSCSI est assez simple, mais le mappage sous Windows est un peu plus complexe que celui de SMB.
Avant de créer le partage réel, nous devrons créer soit un ensemble de données, soit un Zvol à partager. Nous utiliserons un Zvol. Nous créons un Zvol de 2 To sous notre pool SAS. Ceci est disponible sous l’onglet Ensembles de données.
Ensuite, nous pouvons accéder à l'onglet Partages et appuyer sur le bouton Assistant sur iSCSI. Par souci de simplicité, nous appelons ce partage « iscsi-sas », en sélectionnant le Zvol comme appareil et en changeant la plate-forme de partage en système d'exploitation moderne. Puisqu'il s'agit de notre premier partage iSCSI, nous devons créer une cible et un portail.
Ce sont les paramètres sous le portail que nous utilisons pour iSCSI. L'utilisateur est également configuré ici. Après cela, nous pouvons ignorer l'initiateur et passer à autre chose.
Après l'enregistrement, nous serons invités à démarrer iSCSI comme nous l'avons fait avec SMB.
On pourrait penser que nous avons terminé ici et que nous sommes du côté TrueNAS, mais il est maintenant temps de passer du côté de Windows. La machine utilisée dans cette configuration exécute Windows 11 ; les vues sur d’autres systèmes d’exploitation peuvent être différentes.
Nous commençons par rechercher et ouvrir l'initiateur iSCSI.
Dans l'initiateur iSCSI, entrez votre adresse IP TrueNAS dans la zone cible et cliquez sur Connexion rapide.
Une fenêtre contextuelle pour la connexion rapide avec la cible découverte s'affichera. Si votre statut indique Connecté, vous pouvez appuyer sur Terminé.
Quittez l'initiateur iSCSI. Ensuite, nous devons ouvrir la gestion des disques. Vous devriez soit obtenir une fenêtre contextuelle pour initialiser un disque, soit simplement voir votre espace non alloué.
Une fois le disque initialisé, nous pouvons créer un nouveau volume simple.
Nous avons simplement laissé le volume remplir tout l'espace disque et lui avons donné la lettre E. Nous avons ensuite formaté la partition avec NTFS avec une taille d'unité d'allocation par défaut et défini l'étiquette sur « Truenas-iSCSI-SAS ».
Une fois le formatage et la configuration terminés, nous devrions pouvoir voir le partage iSCSI monté.
Et si nous regardons l'onglet Ce PC dans l'Explorateur de fichiers, nous devrions voir notre partage iSCSI monté !
Si le partage est monté ici, alors tout s'est bien passé et vous pouvez désormais profiter de votre stockage en réseau. Ce processus est un peu complexe pour ceux qui ne sont pas habitués au bricolage, mais dans l'ensemble, il ne va pas trop loin dans les détails.
Pour plus de détails, vous pouvez consulter la page de partage iSCSI dans la documentation TrueNAS ici.
Conclusion
Dans l’ensemble, TrueNAS Scale est relativement simple à configurer une fois que vous avez suivi les étapes dans l’ordre. Alors que de nombreuses plates-formes NAS de vente au détail proposent des assistants pour vous guider à travers les étapes, TrueNAS propose un processus plus granulaire. Certaines sections peuvent prêter à confusion, mais avec quelques conseils, elles peuvent être résolues. Il existe également une bonne marge de personnalisation et d'adaptation supplémentaire pour une utilisation par les utilisateurs les plus avancés.
Un autre avantage de TrueNAS Scale est qu'il peut exécuter des applications et dispose également d'un hyperviseur. L'hyperviseur KVM ne remplacera probablement pas le besoin d'hôtes dédiés exécutant Proxmox ou ESXi, mais il peut fonctionner à la rigueur pour faire tourner une simple VM. Les applications vous permettent d'héberger des éléments comme Immich pour rationaliser le stockage de photos, Plex pour le streaming vidéo ou même Nextcloud pour la collaboration et le partage. Ces applications ajoutent des fonctionnalités à TrueNAS et vous permettent de le personnaliser en fonction de vos besoins.
