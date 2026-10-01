---
id: collect-261001-general-networking/general-networking/fr-review-esxi-to-pve-51dfd1ff-1
title: "fr-review-esxi-to-pve-51dfd1ff"
domain: general-networking
role: reference
task: reference
actors: ["Broadcom"]
dates: []
keywords: ["agents", "arr", "open source"]
source: docs/RAG/collect-261001-general-networking/fr-review-esxi-to-pve-51dfd1ff.md
source_anchor: ""
source_lines: [1, 29]
sha256: 01697865aa37e2e980b4133682212838a75bf620d1ab694090363cf823c38f73
---

# fr-review-esxi-to-pve-51dfd1ff

Suite au rachat et à la restructuration de VMware par Broadcom, de nombreux utilisateurs, entreprises et particuliers, ont cherché des solutions pour migrer en toute transparence de leur plateforme d'hyperviseur vers une autre. La hausse des coûts de licence et de support logiciel de VMware a généré une forte demande pour des hyperviseurs d'entreprise plus abordables et performants. Proxmox VE est devenu une plateforme populaire grâce à son caractère open source. Grâce à l'outil d'importation intégré, nous allons détailler les étapes de migration d'une machine virtuelle depuis un serveur VMware ESXi autonome vers Proxmox Virtual Environment (PVE).
Conditions préalables et configuration de la migration ESXi
Pour garantir des migrations fluides d'ESXi vers PVE, assurez-vous que les deux serveurs (ou clusters de serveurs) répondent aux critères suivants :
- Accès réseau illimité entre les serveurs – Les serveurs PVE doivent pouvoir accéder à l’API de l’hôte ESXi.
- Accessibilité des identifiants utilisateur « root » pour les serveurs ESXi et PVE – Pour éviter les problèmes d’autorisation pour les fonctionnalités ESXi et PVE, les informations d’identification de l’utilisateur « root » sur les deux serveurs pendant le processus de migration sont recommandées.
- Espace de stockage suffisant sur le(s) serveur(s) PVE – Cela peut sembler évident, mais vérifiez que le ou les serveurs PVE cibles que vous utilisez disposent de suffisamment de stockage pour contenir les disques VM du serveur ESXi.
- Disques de machines virtuelles non chiffrés – Assurez-vous que les disques de machine virtuelle prêts pour la migration n’ont pas été chiffrés à l’aide de stratégies de stockage de machine virtuelle sur ESXi ou de méthodes de chiffrement au niveau du système d’exploitation qui stockent les clés de chiffrement dans un module TPM virtuel.
Une fois que vous avez vérifié que les serveurs ESXi et PVE répondent aux exigences ci-dessus, connectez-vous au serveur Proxmox VE vers lequel vous souhaitez déplacer les machines virtuelles. Cliquez sur l'onglet « Centre de données » à gauche de l'interface web, puis accédez à la vue « Stockage ».
Cliquez sur « Ajouter » et sélectionnez « ESXi » dans le menu déroulant des options de stockage.
Renseignez tous les champs du menu contextuel « Ajouter : ESXi » en fonction de la configuration de votre serveur ESXi. Pour plus de clarté, nous vous recommandons d'indiquer le nom d'hôte du serveur ESXi dans le champ « ID » et de cocher la case « Ignorer la vérification du certificat ». Avant de cliquer sur « Ajouter », assurez-vous que le champ « Nœuds » contient tous les serveurs Proxmox VE de votre cluster qui recevront les machines virtuelles migrées.
Le serveur ESXi apparaîtra comme stockage sous tous les nœuds précédemment sélectionnés, sur le côté gauche de l'interface web. Cliquez sur l'icône Cloud pour afficher les machines virtuelles pouvant être importées sous l'onglet « Invités virtuels ».
Migration d'une machine virtuelle
Dans l'interface Web du serveur ESXi, utilisez la fonction « Console » pour vous connecter à la machine virtuelle que vous souhaitez migrer et désinstaller « VMware Tools » s'il est installé sur l'invité. Les procédures de désinstallation spécifiques peuvent varier selon le système d'exploitation de la machine virtuelle ; suivez donc les étapes appropriées pour votre invité et redémarrez-le si nécessaire. Pour cette démonstration, nous utiliserons une machine virtuelle Windows Server 2025.
Après avoir vérifié que « VMware Tools » n'est pas présent sur la machine virtuelle, arrêtez l'invité et revenez à l'interface web de Proxmox VE. Si nécessaire, cliquez sur l'icône de stockage ESXi récemment ajoutée sous le nœud PVE cible, puis sélectionnez le fichier « .vmx » associé à la machine virtuelle à migrer. Cliquez ensuite sur « Importer » pour commencer la configuration de la migration et de la machine virtuelle résultante sur votre ou vos serveurs PVE.
À noter: Le nom de la machine virtuelle devrait apparaissent dans le nom de fichier correspondant.
Configurez la machine virtuelle dans l'onglet « Général » de la fenêtre « Importer un invité ». Pour des performances optimales, nous vous recommandons de définir le champ « Type de processeur » sur « hôte », sauf si vous devez choisir un autre type pour des raisons de compatibilité.
Remarque : Ce guide ne traite pas de la migration de machines virtuelles à l’aide de la fonction « Importation en direct ». Cette fonction ne doit être utilisée que si les serveurs ESXi et PVE sont connectés à un réseau à haut débit d’au moins 10 Gbit/s.
Une fois les paramètres de base configurés, cliquez sur l'onglet « Avancé ». Pour une compatibilité optimale avec les systèmes d'exploitation sans pilotes VirtIO inclus par défaut (généralement les machines virtuelles Windows), décochez la case « Préparer pour VirtIO-SCSI ». Le champ « Contrôleur SCSI » devrait alors revenir à « VMware PVSCSI » et le type de disque de la machine virtuelle à SCSI. Vous pouvez également personnaliser le stockage de destination de la machine virtuelle en utilisant les champs « Stockage » de chaque disque connecté.
Remarque : Si la machine virtuelle que vous migrez est un invité compatible EFI, vérifiez que l’élément « efidisk » est coché dans la liste « Disques ».
Après avoir consulté l’onglet « Configuration résultante », cliquez sur « Importer » pour vérifier les paramètres de la machine virtuelle résultante.
La fenêtre « Visualiseur de tâches » s'affiche une fois les disques de l'invité copiés et se termine par « TACHE OK ». Cela indique que la machine virtuelle est prête à être démarrée sur votre ou vos serveurs Proxmox VE.
Fermez la fenêtre « Visionneuse de tâches » et recherchez la machine virtuelle nouvellement créée sur le côté gauche de l'interface web de Proxmox VE. Cliquez sur le nom de la machine virtuelle, accédez à l'onglet « Console » et cliquez sur « Démarrer maintenant » pour la démarrer.
Félicitations ! Vous avez migré avec succès une machine virtuelle depuis ESXi grâce à l'outil d'importation intégré de Proxmox VE. Cependant, si votre machine virtuelle exécute une version de Windows, nous pouvons prendre quelques mesures supplémentaires pour améliorer considérablement son accessibilité et ses performances, comme décrit ci-dessous.
Optimisation des machines virtuelles Windows pour Proxmox VE
Les machines virtuelles sont plus efficaces lorsqu'elles disposent des bons outils. L'un des moyens les plus simples d'améliorer la vitesse, l'efficacité et la communication de vos machines virtuelles Windows avec l'hyperviseur Proxmox VE est d'installer les pilotes open source VirtIO, l'invité Qemu et les agents SPICE sur chaque invité.
Vous pouvez télécharger les trois outils regroupés dans un seul fichier ISO ici . Si vos machines virtuelles Windows ont accès à Internet, vous pouvez télécharger l'ISO sur la machine virtuelle ou transférer le fichier vers votre ou vos serveurs Proxmox VE et l'associer à leurs lecteurs CD/DVD virtuels.
Pour télécharger l'ISO sur un serveur PVE, accédez au stockage des fichiers ISO dans l'interface Web et cliquez sur « Télécharger ».
Dans le menu « Télécharger » qui apparaît, utilisez l'option « Sélectionner un fichier » pour choisir le fichier ISO VirtIO et cliquez sur l'option bleue « Télécharger » pour commencer à le copier sur le stockage du serveur.
Une fois le fichier ISO copié sur le serveur, une fenêtre « Visualiseur de tâches » s'affiche. Attendez que « TASK OK » s'affiche avant de quitter et de connecter l'ISO à une machine virtuelle.
