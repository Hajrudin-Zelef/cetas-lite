---
id: collect-261001-general-networking/general-networking/fr-review-how-to-get-started-with-truenas-scale-75339f23-1
title: "fr-review-how-to-get-started-with-truenas-scale-75339f23"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["open source"]
source: docs/RAG/collect-261001-general-networking/fr-review-how-to-get-started-with-truenas-scale-75339f23.md
source_anchor: ""
source_lines: [1, 25]
sha256: 1190a14d50b9b15348c2a1e710b21f2b2c15912a6c9562d43c2a64b3ac3b7e41
---

# fr-review-how-to-get-started-with-truenas-scale-75339f23

TrueNAS a gagné du terrain dans les communautés d'auto-hébergement et de homelab pour plusieurs raisons. L'une des principales raisons est qu'il est gratuit et Open Source. Une autre raison importante est que, dès le départ, il est assez facile à configurer et possède une interface attrayante. Pour les personnes qui se lancent dans ce passe-temps ou qui souhaitent travailler sur quelque chose, les interfaces CLI peuvent être intimidantes et la sélection des paramètres corrects peut être déroutante. Ainsi, pour vous aider à démarrer, nous avons décidé de créer un guide de configuration rapide pour montrer comment configurer les partages SMB et iSCSI avec les ACL.
TrueNAS est compatible avec de nombreux matériels, mais il est essentiel de vérifier la compatibilité du matériel choisi. Consultez le guide matériel de TrueNAS Scale pour plus de détails. Côté stockage, il est primordial de disposer de plusieurs disques de même capacité pour créer un pool. Bien qu'il soit possible de combiner différents disques, cela risque d'engendrer plus de problèmes qu'autre chose. Hormis ce point, TrueNAS est généralement assez peu exigeant.
Installation
L'installation globale de TrueNAS Scale est très simple. La majeure partie de la configuration réelle est effectuée via l'interface Web après l'installation. Une fois que vous avez démarré votre programme d'installation TrueNAS, vous appuyez sur le chargeur de démarrage GRUB avec la possibilité de démarrer l'installation, qui passera automatiquement au menu de configuration de la console.
Dans le menu initial, sélectionnez l'option Installer/Mise à niveau et cliquez sur OK.
L'écran suivant varie en fonction de la configuration réelle. Sélectionnez le lecteur pour l'installation TrueNAS. Utilisez les touches fléchées pour faire défiler pour mettre en surbrillance le lecteur préféré et appuyez sur espace. Notre capture d'écran est un peu différente car ces disques étaient auparavant utilisés dans une configuration TrueNAS et attribués à des pools. En règle générale, vous ne verrez que les noms, adresses et capacités des lecteurs disponibles. Sélectionnez le lecteur de démarrage souhaité et cliquez sur OK.
Il est important de créer un compte d'utilisateur administratif et TrueNAS propose deux bonnes options. La sélection de l'option 1 créera un mot de passe administrateur dans la console. Le choix de l'option 3 nécessite que l'utilisateur crée le mot de passe administrateur lors de sa première connexion à l'interface Web. L’une ou l’autre option est acceptable ici et dépend de vos préférences personnelles. L'option 2 n'est pas une bonne option ! Après avoir fait votre choix, cliquez sur OK.
Le programme d'installation suivra toutes ses étapes et, une fois terminé, le message d'installation réussie vous sera présenté. Retirez votre support d'installation et cliquez sur OK pour passer à l'écran suivant.
Toutes nos félicitations! Si vous voyez cet écran, cela signifie que TrueNAS a été installé. A partir de là, vous ne verrez pas grand-chose sur la console locale. Accédez à une machine connectée au réseau et utilisez l’interface Web. En règle générale, à moins que les choses tournent mal, il ne sera pas nécessaire de revoir cette interface et vous pourrez exécuter votre boîtier TrueNAS complètement sans tête.
Configuration de l'échelle TrueNAS
Accédez à l'interface Web à l'adresse IP répertoriée dans la console. Une fois que vous avez saisi l'adresse IP sur la machine en réseau, l'écran de connexion s'affiche. Vous devrez soit configurer le mot de passe administrateur, soit entrer admin et le mot de passe que vous avez précédemment configuré. Ensuite, vous rencontrerez l’écran d’atterrissage.
TrueNAS est installé, mais il nécessite une configuration supplémentaire pour fonctionner correctement en tant que NAS. Pour commencer, accédez à l’onglet Stockage dans la barre latérale de gauche. Tout d’abord, nous devons créer un pool, alors cliquons sur Créer un pool.
Création d'un pool à l'échelle TrueNAS
La création d'un pool de lecteurs est la première chose à faire pour rendre cet espace utilisable. En règle générale, vous souhaitez plusieurs disques de la même capacité pour éviter les maux de tête. Nous disposons de six SSD NVMe de 1.6 To et de neuf SSD SAS de 480 Go, nous allons donc les configurer en deux pools distincts. Par souci de simplicité, nous les divisons en pool NVMe et pool SAS et les plaçons dans ZFS RAIDZ1, en les répartissant avec un seul disque de rechange. Pour les configurations comportant plus de six disques ou plus de 2 To chacun, il est probablement judicieux d'utiliser RAIDZ2 pour le deuxième disque de rechange. Cela permet de réduire les risques de perte de données si vous rencontrez un autre échec alors que vous essayez de reconstruire le pool à partir du premier échec. ZFS RAIDZ est un moyen rapide et simple de regrouper et de mettre en ligne des SSD NVMe.
C'est tout pour créer votre pool, à moins que vous ne souhaitiez explorer d'autres options telles que la mise en cache et les pièces de rechange. Bien que cela ne soit pas essentiel pour les configurations NAS standard, vous pouvez enregistrer votre travail et accéder à la page de révision.
Création d'utilisateur
Créons maintenant des comptes d'utilisateurs afin que le compte administrateur ne soit pas utilisé pour accéder au stockage. Cela nous permet de configurer des autorisations pour d'autres utilisateurs et de leur accorder un stockage partagé ou séparé. Vous pouvez également configurer des autorisations par groupes pour simplifier l'accès des utilisateurs au même stockage.
Accédez à l'onglet « Identifiants » à gauche et sélectionnez « Utilisateurs locaux ». Cliquez sur « Ajouter un utilisateur » en haut à droite et saisissez directement les informations de l'utilisateur. Ce guide ne traite pas des paramètres du répertoire personnel , bien qu'ils offrent des options utiles. L'authentification via LDAP et Active Directory dépasse le cadre de ce guide et ne sera donc pas abordée ici. Toutefois, ces options permettent de s'authentifier via des comptes existants sans que les utilisateurs aient besoin de créer un compte TrueNAS dédié.
Pour plus d'informations sur les partages de fichiers à domicile, vous pouvez consulter la documentation TrueNAS ici.
Ensembles de données et partages
Une fois les comptes d'utilisateurs et les groupes créés, vous pouvez configurer un stockage accessible aux utilisateurs et utiliser votre NAS. Vous pouvez choisir entre iSCSI et SMB pour cela, mais SMB est votre meilleur choix pour le partage avec plusieurs utilisateurs. Nous couvrirons la configuration des deux ici.
Configuration PME
Pour configurer SMB, nous devons commencer par créer un ensemble de données. Pour ce faire, accédez à l'onglet Ensembles de données sur le côté gauche, cliquez sur le pool dont vous souhaitez être le parent, puis appuyez sur le bouton Ajouter un ensemble de données en haut à droite. Ici, vous pouvez lui donner un nom, modifier les paramètres de compression et de cryptage et définir le type de partage.
Puisque nous configurons un partage SMB, nous voulons nous assurer que l'ensemble de données est correctement configuré. Utilisez simplement le menu déroulant pour sélectionner SMB dans le champ Type de partage. Nous pouvons économiser pour permettre au partage SMB d’occuper la totalité du pool, mais pour diviser le pool, nous devons passer aux options avancées.
Maintenant que nous sommes dans les options avancées, nous pouvons modifier la taille du jeu de données en entrant simplement la capacité souhaitée dans le champ quota. À moins que vous n'ayez besoin de modifier d'autres options, nous avons créé l'ensemble de données et pouvons continuer.
