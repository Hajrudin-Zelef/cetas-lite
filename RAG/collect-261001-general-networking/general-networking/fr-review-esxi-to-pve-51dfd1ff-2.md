---
id: collect-261001-general-networking/general-networking/fr-review-esxi-to-pve-51dfd1ff-2
title: "fr-review-esxi-to-pve-51dfd1ff"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "arr", "gpu"]
source: docs/RAG/collect-261001-general-networking/fr-review-esxi-to-pve-51dfd1ff.md
source_anchor: ""
source_lines: [30, 62]
sha256: 5f17f10bfcefc4a7ddf761bb1385277495345f41f738ea82ca93e27374c1b9e1
---

# fr-review-esxi-to-pve-51dfd1ff

Ensuite, cliquez sur une machine virtuelle Windows à gauche de l'interface web et sélectionnez l'onglet « Matériel ». Sélectionnez le lecteur CD/DVD et cliquez sur « Modifier ».
Dans la fenêtre contextuelle qui apparaît, remplissez les champs « Stockage » et « Image ISO » avec les valeurs appropriées pour sélectionner le fichier ISO VirtIO que vous avez téléchargé et sélectionnez « OK » pour confirmer.
Connectez-vous à l'invité via l'onglet « Console » et ouvrez l'ISO dans l'Explorateur de fichiers. Localisez le fichier « virtio-win-guest-tools.exe », faites un clic droit dessus et choisissez « Exécuter en tant qu'administrateur » dans le menu qui apparaît.
Si une fenêtre de contrôle de compte d'utilisateur s'affiche, confirmez et poursuivez l'installation en conservant toutes les options par défaut présélectionnées.
Une fois l'installation terminée, fermez le menu et revenez à l'onglet « Matériel » de la machine virtuelle dans l'interface web de Proxmox VE. Cliquez sur « Ajouter » dans ce menu et sélectionnez « Disque dur » dans le menu déroulant.
Renseignez les informations dans le menu « Ajouter : Disque dur », en définissant « Bus/Périphérique » comme « Bloc VirtIO » et « Stockage » comme stockage préféré de la VM. Laissez la valeur du champ « Taille du disque (Gio) » à 32 ou réduisez-la en fonction de l'espace de stockage disponible sur votre/vos serveur(s). Une fois terminé, cliquez sur « Ajouter » pour attacher le disque virtuel nouvellement créé.
Après avoir ajouté le disque, revenez à l'onglet « Console » de la machine virtuelle et ouvrez le menu « Gestion des disques ». Faites un clic droit sur le nouveau disque et choisissez « En ligne » dans le menu suivant. De même, une fois le disque en ligne, faites un nouveau clic droit dessus et sélectionnez « Initialiser le disque ».
Choisissez un style de partition MBR ou GPT pour le nouveau disque et cliquez sur « OK ».
Ensuite, faites un clic droit n’importe où sur l’espace non alloué du nouveau disque et sélectionnez « Nouveau volume simple ».
Utilisez l'assistant de volume simplifié récemment introduit, en choisissant une lettre de lecteur et en générant une étiquette de volume que vous n'avez pas l'intention d'utiliser pour les disques virtuels.
Fermez ensuite la fenêtre « Gestion des disques » et arrêtez la machine virtuelle. Une fois éteinte, revenez à l'onglet « Matériel ».
Suivez ensuite scrupuleusement les étapes ci-dessous, en veillant à ne pas détruire accidentellement les données ou la configuration de la VM. Nous allons maintenant modifier plusieurs options de périphériques virtualisés et de VM afin de tirer parti des pilotes VirtIO et des agents invités installés précédemment :
- Accédez à l'onglet Matériel de la machine virtuelle et cliquez sur le disque VirtIO précédemment ajouté. Cliquez sur « Détacher » pour confirmer sa déconnexion.
- Cliquez sur l'option « Disque inutilisé », puis sur « Supprimer » et confirmez la suppression du lecteur.
- Pour chacun des disques restants de la machine virtuelle, utilisez « Détacher » pour le déconnecter de la machine virtuelle.
- Pour chacun des « Disques inutilisés » de la machine virtuelle, cliquez sur « Modifier », remplacez le champ « Bus/Périphérique » par « Bloc VirtIO » et cliquez sur « Ajouter ».
- Sélectionnez le périphérique « SCSI Controller » et remplacez-le par « VirtIO SCSI single » à l'aide de « Modifier » puis de « OK » lorsque vous avez terminé.
- Pour chacun des « Périphériques réseau » de la machine virtuelle, cliquez sur « Modifier », remplacez le champ « Modèle » par « VirtIO (paravirtualisé) » et cliquez sur « OK ».
  - Ne modifiez pas l'adresse MAC et comprenez que la modification du type d'interface réseau signifie que vous devrez peut-être modifier les paramètres réseau à l'intérieur de l'invité si DHCP n'est pas utilisé.
- (Recommandé uniquement pour les machines virtuelles Windows compatibles UEFI) Sélectionnez l’option « Machine », cliquez sur « Modifier » et cochez la case « Avancé ».
- (Recommandé uniquement pour les machines virtuelles Windows compatibles UEFI) Modifiez le champ « Machine » en « Q35 », modifiez le champ « Version » en le type le plus récent et modifiez le champ « vIOMMU » en « VirtIO ».
- Accédez à l'onglet « Options » de la machine virtuelle, cliquez sur l'option « Agent invité QEMU » et cliquez sur « Modifier ».
- Cochez le champ « Utiliser l'agent invité QEMU », assurez-vous que « Par défaut (VirtIO) » ou « VirtIO » est sélectionné, puis cliquez sur « OK ».
- Sélectionnez l’option « Ordre de démarrage » et cliquez sur « Modifier ».
- Faites glisser, déposez et vérifiez les options de démarrage pour vous assurer que les options de démarrage souhaitées sont sélectionnées pour la machine virtuelle, puis cliquez sur « OK ».
  - Prenez un instant pour examiner les modifications apportées à la machine virtuelle. Tous les périphériques virtuels « Disque dur », « Périphérique réseau » et « Contrôleur SCSI » doivent être configurés avec leurs options VirtIO respectives, et le type de machine virtuelle doit être défini sur « Q35 » si la machine virtuelle est compatible UEFI.
- Enfin, accédez à l’onglet « Console » de la machine virtuelle et cliquez sur « Démarrer maintenant » pour la mettre sous tension.
Remarque : Si vous rencontrez des problèmes de démarrage sur des machines virtuelles avec des configurations BIOS héritées ou un contrôleur SCSI LSI émulé, accédez à l’onglet « Matériel » de la machine virtuelle et utilisez « Détacher » pour déconnecter tous les disques virtuels. Utilisez « Modifier » pour accéder au menu et modifier le type de bus/périphérique en « SATA ». Fermez le menu en cliquant sur le bouton bleu « Ajouter » une fois la modification effectuée, puis accédez à l’onglet « Options » de la machine virtuelle, sélectionnez l’option « Ordre de démarrage » et cliquez sur « Modifier ». Réorganisez l’ordre de démarrage afin que le disque SATA de démarrage soit en premier, puis cochez la case « Activé » pour terminer.
Avant:
Après:
Conclusion
La migration de VMware ESXi vers Proxmox VE est un processus simple grâce à l'outil d'importation intégré de Proxmox. Vous pouvez transférer vos charges de travail avec un minimum de temps d'arrêt et de perturbations en suivant les étapes de préparation appropriées : vérification de la compatibilité, nettoyage de VMware Tools et configuration des paramètres d'importation. Une fois vos machines virtuelles exécutées correctement sur Proxmox, validez leurs performances, installez les pilotes VirtIO le cas échéant et effectuez de nouvelles sauvegardes dans votre environnement mis à jour. Après avoir vérifié que tout fonctionne correctement et de manière fiable, il est recommandé de décommissionner ou d'archiver les machines virtuelles ESXi d'origine afin de libérer des ressources et d'éviter les dérives de configuration entre les plateformes.
Ce processus offre à de nombreux utilisateurs une transition relativement simple et rapide, permettant de maîtriser la hausse des coûts de licences VMware tout en conservant une plateforme robuste et adaptée aux besoins essentiels de virtualisation en entreprise. À ce propos, si vous souhaitez partager des GPU dans Proxmox VE, consultez notre guide sur les vGPU Proxmox.
