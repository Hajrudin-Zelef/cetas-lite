---
id: collect-261001-general-networking/general-networking/fr-review-how-it-works-proxmox-import-wizard-vmware-migration-tool-8000e879-1
title: "fr-review-how-it-works-proxmox-import-wizard-vmware-migration-tool-8000e879"
domain: general-networking
role: reference
task: reference
actors: ["Broadcom"]
dates: []
keywords: ["arr", "attention", "datacenter"]
source: docs/RAG/collect-261001-general-networking/fr-review-how-it-works-proxmox-import-wizard-vmware-migration-tool-8000e879.md
source_anchor: ""
source_lines: [1, 42]
sha256: 3e3f56d4872708e04d4f474c6905aed6ccefbde873c72c4a70a04db47a20153d
---

# fr-review-how-it-works-proxmox-import-wizard-vmware-migration-tool-8000e879

Depuis le rachat récent de VMware par Broadcom, de nombreux clients de petite taille s'inquiètent, à juste titre, des hausses de prix et envisagent par conséquent des alternatives à VMware . Proxmox espère séduire ces utilisateurs, qui n'ont pas forcément besoin de la suite VMware complète, grâce à un nouvel outil : l'assistant d'importation Proxmox.
La facilité de migration d’un hyperviseur à un autre est un problème majeur, ou peut-être de manière plus appropriée, le manque de facilité de migration. Avec l'arrivée récente de l'assistant Proxmox dans GA, nous voulions voir à quel point il est facile de déplacer une VM de VMware vers Proxmox. Pour faciliter la conversation, nous avons créé un guide de migration étape par étape pour aider nos lecteurs à mieux comprendre le nouvel assistant d'importation Proxmox. Bien que la plupart des étapes et des captures d'écran se concentrent sur Windows, certaines seront interchangeables avec Linux et d'autres systèmes d'exploitation.
Assistant d'importation Proxmox – Préparation et considérations
Avant de commencer, vous devez effectuer quelques étapes au niveau du système d'exploitation invité pour préparer les machines virtuelles à l'exportation. Vous devez supprimer les outils spécifiques à l'hyperviseur de votre VM invitée, tels que les outils VMware dans notre instance. Ceux-ci doivent être supprimés avant la migration, sinon ils peuvent causer des problèmes lors du changement d'hyperviseur. Ils peuvent également être difficiles à supprimer après le transfert.
Une autre tâche consiste à supprimer toute configuration réseau statique avant la migration si votre VM exécute Windows. Après la migration, la VM recevra une nouvelle carte réseau et vous devrez configurer les paramètres statiques sur cette carte. Si vous ne parvenez pas à supprimer la configuration statique, Windows affichera un avertissement lorsque vous tenterez de définir les mêmes paramètres statiques sur la nouvelle carte réseau, même si la précédente a été supprimée.
Pour les réservations DHCP, vous devez soit basculer la réservation sur le MAC de la nouvelle carte réseau sur la VM, soit définir manuellement le MAC sur la carte réseau dans la VM.
Si vous utilisez un périphérique TPM virtuel, il n'est actuellement pas possible de migrer l'état vTPM de VMware vers Proxmox. Si vous utilisez le chiffrement complet du disque avec les clés stockées dans vTPM, envisagez de le désactiver. Après la désactivation, assurez-vous de disposer des clés de déchiffrement manuel au cas où.
En termes de stockage, l'importation d'une VM avec des disques sauvegardés par le stockage VMware vSAN ne fonctionne pas. Proxmox répertorie le déplacement des vDisks de la VM vers un SSD ou un disque dur local du serveur hyperviseur comme solution de contournement. Il est également recommandé de ne pas importer simultanément plus de quatre disques de VM.
Avant de migrer, assurez-vous de mettre hors tension la VM source.
Importer des sources
À l'heure actuelle, VMware ESXi est la seule source d'importation prise en charge, mais il est prévu d'ajouter la prise en charge de l'importation des fichiers OVA/OVF VMware . L'importation d'ESXi a été testée des versions 6.5 à 8.0.
Étapes de migration
Proxmox propose trois méthodes d'importation de machines virtuelles : manuelle, automatique et à chaud . La fonction d'importation à chaud démarre la machine virtuelle pendant le processus d'importation afin de réduire l'interruption de service. En cas d'échec de l'importation à chaud, toutes les données écrites depuis le début de l'importation seront perdues ; il est donc recommandé d'effectuer un test préalable sur une machine virtuelle dédiée.
Bien qu'il existe trois méthodes d'importation, ce guide étape par étape se concentre sur la méthode d'importation automatique. Nous pensons que ce sera la méthode la plus couramment utilisée, à moins que les machines virtuelles sources nécessitent une attention particulière lors du processus d'importation.
Importation automatique complète de la VM
L'importateur de VM intégré intègre le système de plugin de stockage dans l'interface Web. Cette méthode importe la VM complète, avec la plupart de sa configuration mappée sur le modèle de configuration de Proxmox, réduisant ainsi les temps d'arrêt.
- Assurez-vous que votre Proxmox VE est égal ou supérieur à la version 8 et dispose des dernières mises à jour du système.
- Accédez à Datacenter → Stockage → Ajouter et sélectionnez ESXi comme stockage source d'importation.
- Entrez le domaine ou l'adresse IP de l'hôte ESXi et les informations d'identification d'un compte administrateur. Si votre instance ESXi dispose d'un certificat auto-signé, vous pouvez ignorer la vérification du certificat ou ajouter l'autorité de certification à votre magasin de confiance système dans Proxmox.
- Dans le menu de gauche, sélectionnez la nouvelle icône de stockage pour votre hôte et assurez-vous que vous pouvez voir toutes vos machines virtuelles invitées.
- Sélectionnez la VM que vous souhaitez importer, puis appuyez sur le bouton « importer » en haut à gauche.
- Sélectionnez au minimum votre stockage cible et le pont réseau pour la VM. Si nécessaire, utilisez l'onglet Avancé pour une configuration plus détaillée.
- Examinez la configuration résultante pour votre VM pour vous assurer que tout semble correct.
- Assurez-vous que votre VM source est prête à être importée et éteignez-la.
- Démarrez l'import côté Proxmox. En fonction de la taille du disque et des capacités de votre réseau, cette étape peut prendre un certain temps. (N'importez pas plus de 4 disques VM à la fois)
- Démarrez la machine virtuelle et inspectez le système d'exploitation pour voir si des modifications post-migration sont nécessaires.
Si votre VM fonctionne correctement, vous avez terminé votre première importation. Cependant, vous n’avez pas encore terminé. Il y a quelques points à régler après la migration pour obtenir la meilleure expérience avec Windows.
Post-migration (pilotes VirtIO)
- Assurez-vous de vérifier vos paramètres réseau après l'importation. Le nom de la carte réseau a probablement changé.
- (Principalement pour les machines virtuelles Windows) Installez les pilotes manquants.
- Pour les machines virtuelles Windows, téléchargez et joignez le VirtIO ISO. Une fois le pilote installé, le disque de démarrage doit être basculé vers VirtIO, qui est entièrement documenté dans cet article du wiki Proxmox.
  - L'utilisation d'un périphérique VirtIO SCSI fonctionne mieux que le SCSI émulé.
Ces étapes pour VirtIO ne sont nécessaires que pour les machines Windows. Bien qu'il soit possible de fonctionner sans eux, cela aura un impact négatif sur les performances de la VM.
- Téléchargez la dernière version ISO stable du pilote VirtIO à partir de ici.
- Téléchargez l'ISO du pilote sur votre stockage ISO Proxmox.
- Attachez l’ISO du pilote au lecteur de CD-ROM de la VM.
- Ajoutez un disque de 1 Go avec Bus Type SCSI ou VirtIO Block à la VM.
- Si le disque se connecte à chaud, passez à l'étape suivante ; sinon, redémarrez la VM. Si les pilotes ne sont pas installés, le disque apparaîtra dans le gestionnaire de périphériques en tant que contrôleur SCSI avec une erreur.
- Si vous voyez le contrôleur SCSI, cliquez dessus avec le bouton droit, sélectionnez « Mettre à jour le pilote », puis sélectionnez « Parcourir mon ordinateur pour les pilotes ». Localisez le lecteur de CD et cliquez sur OK, puis sur Suivant.
- Vous devriez maintenant recevoir un message indiquant que Windows a mis à jour avec succès vos pilotes.
- Arrêtez la VM.
- Détachez et supprimez le disque de 1 Go que nous avons créé précédemment et l'ISO dans le lecteur de CD-ROM.
