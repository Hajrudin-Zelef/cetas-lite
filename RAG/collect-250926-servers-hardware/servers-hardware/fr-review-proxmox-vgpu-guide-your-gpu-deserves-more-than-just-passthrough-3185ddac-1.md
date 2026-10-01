---
id: collect-250926-servers-hardware/servers-hardware/fr-review-proxmox-vgpu-guide-your-gpu-deserves-more-than-just-passthrough-3185ddac-1
title: "fr-review-proxmox-vgpu-guide-your-gpu-deserves-more-than-just-passthrough-3185ddac"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["gpu", "amd", "arr", "attribution", "nvidia", "valuation"]
source: docs/RAG/clean4/fr-review-proxmox-vgpu-guide-your-gpu-deserves-more-than-just-passthrough-3185ddac.md
source_anchor: ""
source_lines: [1, 35]
sha256: abbf5e99364593cb4dfef0a3dfeb9a0230c3b3b8e449f1de95828573c3533b4a
---

# fr-review-proxmox-vgpu-guide-your-gpu-deserves-more-than-just-passthrough-3185ddac

Proxmox Server Solutions a récemment annoncé la prise en charge du vGPU NVIDIA pour son hyperviseur d'environnement virtuel (VE), ouvrant ainsi la voie à de nouvelles possibilités d'accélération GPU pour le rendu, la VDI, l'IA et d'autres applications. Rejoignez-nous pour découvrir les nouvelles fonctionnalités de l'hyperviseur et évaluer ses performances.
Quel est le problème avec la prise en charge du vGPU ?
Si vous ne connaissez pas le logiciel vGPU de NVIDIA, sachez qu'il permet de partitionner et de répartir les ressources de calcul et de mémoire d'un GPU entre plusieurs machines virtuelles (VM). Dans cette configuration, les VM utilisent alternativement la puissance de traitement du GPU et allouent une partie de la mémoire de la carte à leurs besoins. Cette approche permet au GPU de maintenir une utilisation maximale, même lorsque les charges de travail des VM invitées varient en intensité.
Avant de commencer… (Prérequis)
Avant de déployer un environnement de virtualisation GPU dans Proxmox VE (PVE), vous aurez besoin de quelques éléments. Comme de nombreux outils pour centres de données, le logiciel vGPU de NVIDIA n'est pas gratuit. Vous devrez acheter la version d'évaluation ou vous y inscrire , et créer un compte NVIDIA Enterprise. Ensuite, vous devrez vous procurer une carte graphique compatible vGPU et télécharger les pilotes correspondants depuis le portail de licences . Téléchargez les versions les plus récentes des pilotes « Linux KVM ». Au minimum, téléchargez les pilotes suivants :
- Logiciel vGPU : 18.0
- Pilote hôte : 570.124.03
- Pilote invité Linux : 570.124.06
- Pilote invité Windows : 572.60
Nous avons utilisé le NVIDIA L40S et un Dell PowerEdge R760 pour ce projet.
De plus, vous aurez besoin d'un serveur de licences dédié (DLS) ou d'un serveur de licences cloud (CLS) pour activer la fonctionnalité vGPU sur vos machines virtuelles. Vous trouverez un guide de démarrage rapide en cliquant ici.
Ensuite, vérifiez que les paramètres UEFI (BIOS) de votre serveur activent les fonctionnalités de virtualisation appropriées. Recherchez les options VT-d ou AMD-v , SR-IOV , le décodage 4G , la barre redimensionnable et l'interprétation de l'identifiant de routage alternatif (ARI) , et assurez-vous qu'elles sont toutes activées.
Remarque : Il se peut que vous ne trouviez pas toutes ces fonctionnalités dans le menu UEFI, car certaines peuvent ne pas être accessibles à l'utilisateur.
Enfin, vérifiez que vous utilisez une version appropriée de Proxmox VE. La fonctionnalité vGPU nécessite au minimum la version 8.3.4 de pve-manager, avec le noyau 6.18.12-8-pve ou une version ultérieure. Vous pouvez vérifier les versions logicielles de votre nœud PVE en accédant à l'onglet « Résumé » du serveur souhaité, comme indiqué ci-dessous :
L'hôte avec le plus (configuration vGPU de l'hôte Proxmox)
Maintenant que tout est prêt, il est temps de configurer le serveur Proxmox VE. Dans l'interface web de votre serveur Proxmox VE, cliquez sur le nom du serveur à gauche de l'écran et sélectionnez l'onglet « Shell ». Saisissez la commande suivante dans la fenêtre de console qui apparaît, puis appuyez sur Entrée :
apt install pve-nvidia-vgpu-helper
Cela garantira l'installation de l'outil de configuration vGPU sur votre serveur, le préparant ainsi à la prise en charge de Proxmox vGPU. Une fois l'installation du script terminée ou la présence du script confirmée, exécutez une nouvelle commande pour exécuter l'outil.
configuration de pve-nvidia-vgpu-helper
Répondez « Y » à toutes les questions et continuez jusqu'à ce que la console réapparaisse et que le script soit terminé. Effectuez un redémarrage rapide du serveur en accédant à l'onglet « Résumé » et en cliquant sur le bouton « Redémarrer », ou saisissez la commande de redémarrage dans la console de l'onglet « Shell » et appuyez sur Entrée.
Ensuite, il faut charger le pilote hôte vGPU de NVIDIA sur le serveur. Une fois le serveur redémarré, utilisez un outil de transfert SSH ou SCP tel que WinSCP pour copier le pilote hôte sur le nœud.
Remarque : Si vous avez téléchargé tous les pilotes ensemble sous forme de dossier compressé (.zip), vous devrez peut-être d’abord extraire son contenu et choisir le fichier « .run » dans le dossier « Host_Drivers ».
Placez le fichier dans le répertoire « /home » du serveur et préparez-vous à exécuter le programme d'installation avec les commandes suivantes.
cd /home chown root NVIDIA-Linux-x86_64-570.124.03-vgpu-kvm.run chmod +X NVIDIA-Linux-x86_64-570.124.03-vgpu-kvm.run ./NVIDIA-Linux-x86_64-570.124.03-vgpu-kvm.run --dkms
Remarque : Remplacez « NVIDIA-Linux-x86_64-570.124.03-vgpu-kvm.run » par le nom du pilote que vous avez téléchargé. Vous pouvez utiliser la commande « ls » pour afficher le nom du fichier une fois qu’il aura été placé dans le répertoire « /home ».
Maintenant que le pilote est installé sur le serveur, il ne nous reste plus qu'à configurer le côté hôte de notre vGPU Proxmox ! Avant de pouvoir marquer le GPU comme périphérique pouvant être réparti entre des machines virtuelles, nous devons activer la virtualisation d'E/S à racine unique (SR-IOV). NVIDIA définit cette fonctionnalité comme « …une technologie permettant à un périphérique PCIe physique de se présenter plusieurs fois sur le bus PCIe. Cette technologie permet de créer plusieurs instances virtuelles du périphérique avec des ressources distinctes. » SR-IOV étant une technologie essentielle au fonctionnement de base des GPU virtuels modernes, configurez-la pour qu'elle s'active au démarrage avec la commande suivante :
systemctl enable --now pve-nvidia-sriov@ALL.service
Enfin, nous pouvons cartographier les ressources du GPU afin de les répartir efficacement entre les machines virtuelles. Sur l'interface web du serveur Proxmox VE, cliquez sur « Centre de données » en haut à gauche et faites défiler la page vers le bas pour sélectionner l'onglet « Mappages de ressources ».
Cliquez sur le bouton « Ajouter » sous la section « Périphériques PCI » de la page et remplissez le champ « Nom : » dans la fenêtre suivante avec le nom qui décrit le GPU que vous mappez.
Cochez ensuite la case « Utiliser avec les périphériques médiatisés » et assurez-vous que le menu déroulant « Mappage sur le nœud » contient le serveur sur lequel le GPU est mappé. Faites défiler la liste des périphériques et vérifiez que tous les identifiants de périphérique indiquent « NVIDIA Corporation » dans la colonne « Fournisseur ». Si c'est le cas, cochez la case en haut à gauche du tableau pour sélectionner tous les périphériques ; sinon, sélectionnez uniquement les périphériques dont le fournisseur est « NVIDIA Corporation ».
Remarque : Si plusieurs GPU sont installés dans votre système, vous pouvez utiliser la commande « lspci » dans l’onglet « Shell » du serveur souhaité pour déterminer les identifiants correspondant à chaque carte.
Une fois le périphérique sélectionné, cliquez sur le bouton « Créer » en bas à droite de la fenêtre contextuelle pour confirmer votre sélection. Votre GPU NVIDIA est maintenant prêt à être découpé en vGPU pour les invités de votre serveur Proxmox VE !
Servir les invités (attribution de vGPU aux machines virtuelles)
Tout est en place pour commencer à attribuer et utiliser des vGPU sur nos machines virtuelles. Commencez par créer une nouvelle machine virtuelle, comme d'habitude, ou utilisez une machine virtuelle existante. Pour notre démonstration, nous utiliserons une machine virtuelle Windows Server 2025.
Dans l'interface Web du serveur Proxmox VE, arrêtez la machine virtuelle à l'aide de la méthode qui vous convient (console noVNC, menu d'alimentation de l'invité, etc.) et cliquez sur l'onglet « Matériel » de l'invité.
