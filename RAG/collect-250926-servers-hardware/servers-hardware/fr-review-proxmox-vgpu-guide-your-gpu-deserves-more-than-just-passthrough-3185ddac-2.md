---
id: collect-250926-servers-hardware/servers-hardware/fr-review-proxmox-vgpu-guide-your-gpu-deserves-more-than-just-passthrough-3185ddac-2
title: "fr-review-proxmox-vgpu-guide-your-gpu-deserves-more-than-just-passthrough-3185ddac"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel", "Nvidia"]
dates: []
keywords: ["gpu", "benchmark", "benchmarks", "intel", "nvidia"]
source: docs/RAG/clean4/fr-review-proxmox-vgpu-guide-your-gpu-deserves-more-than-just-passthrough-3185ddac.md
source_anchor: ""
source_lines: [36, 65]
sha256: 61a5ff3bf968049bfa257b1d721cb066366f0d7c64cfca79da8686c5cee6a339
---

# fr-review-proxmox-vgpu-guide-your-gpu-deserves-more-than-just-passthrough-3185ddac

Dans le menu déroulant qui apparaît après avoir cliqué sur le bouton « Ajouter », sélectionnez un « Périphérique PCI ».
Dans la fenêtre contextuelle « Ajouter : périphérique PCI », sélectionnez le nom que vous avez attribué au GPU mappé en ressources dans le champ « Périphérique : ».
Cliquez ensuite sur le champ « Type MDev » et observez la liste des options de type de périphérique médiatisé. Vous remarquerez que chaque choix est désigné par un chiffre et une lettre. Le chiffre représente la quantité de VRAM attribuée à l'invité (en gigaoctets), tandis que les lettres « A », « B » et « Q » font référence au cas d'utilisation du vGPU.
- Q – Postes de travail virtuels avec pilotes RTX Enterprise (nécessite une licence RTX vWS)
- B – Bureaux virtuels (nécessite une licence NVIDIA Virtual PC ou une licence RTX vWS)
- A – Solutions applicatives (nécessite une licence NVIDIA Virtual Applications)
Remarque : Vous pouvez en savoir plus sur les différents profils vGPU et leurs licences requises ici.
Pour cette démonstration, nous avons utilisé le profil « NVIDIA L40S-12Q ». Après avoir sélectionné le type de périphérique souhaité, cochez la case « PCI-Express » et cliquez sur le bouton bleu « Ajouter ».
Un vGPU est désormais attribué à la VM, mais il nous reste à installer le pilote invité et un jeton de licence pour lancer l'opération. Vous pouvez maintenant démarrer la machine virtuelle et vous connecter.
Après vous être connecté, transférez le pilote invité NVIDIA vGPU acquis précédemment à partir du portail de licence vers la machine virtuelle, selon le mode de votre choix (SMB, SCP, téléchargement direct, etc.). Vous pouvez également créer et télécharger un fichier CD-ROM virtuel (.iso) contenant le pilote sur le stockage de votre serveur PVE à l'aide d'un logiciel comme ImgBurn pour déployer simultanément de nombreux invités compatibles vGPU.
Exécutez le fichier d’installation du pilote en double-cliquant sur l’exécutable et suivez les instructions du menu d’installation qui s’affiche.
Extrayez le pilote vers l’emplacement par défaut dans le champ « Chemin d’extraction : » et choisissez l’option d’installation « Express » lorsque vous y êtes invité.
Une fois l’installation du pilote terminée, cliquez sur le bouton « FERMER » pour quitter le menu.
Ensuite, nous devrons activer la licence vGPU sur l'invité. Les méthodes d'activation peuvent varier considérablement selon que vous choisissez un serveur de licences dédié (DLS) ou un serveur de licences cloud (CLS), et selon votre système d'exploitation. Suivez les instructions de NVIDIA. Guide de démarrage rapide du système de licence et Guide de l'utilisateur des licences client pour les étapes détaillées sur l'activation des clients pour votre configuration spécifique.
Nous avons utilisé un serveur de licences cloud et reçu un fichier de jeton à placer sur les invités pour l'activation. Déplacez ce fichier sur l'invité et copiez-le dans le dossier « C:\Program Files\NVIDIA Corporation\vGPU Licensing\ClientConfigToken ».
Ensuite, un redémarrage de l'invité est nécessaire pour terminer le processus d'activation.
Après avoir suivi toutes les étapes de configuration hôte et invité détaillées ici, vous devriez être prêt à exécuter des programmes et applications nécessitant un GPU. N'oubliez pas d'activer le protocole RDP (Remote Desktop Protocol) ou d'installer votre logiciel de bureau à distance préféré sur vos invités après le redémarrage pour profiter de l'affichage à distance accéléré par GPU !
Faire tourner le moteur (test du vGPU Proxmox)
Maintenant que nous disposons de serveurs virtuels avec GPU virtuels, testons-les ! Chacune de nos machines virtuelles est configurée avec 8 processeurs virtuels Intel Xeon Platinum 8580 (4 cœurs hyperthreadés), 32 Go de RAM ECC DDR5 à 4800 40 MT/s et le profil vGPU NVIDIA L12S-12Q (station de travail virtuelle) avec XNUMX Go de VRAM. Vous pouvez consulter la configuration matérielle complète des machines virtuelles ci-dessous :
Cinebench 2024
Basé sur le logiciel de modélisation et d'animation Cinema 4D de Maxon, Cinebench 2024 offre un aperçu intéressant et objectif des performances de rendu sur vGPU. Comparons la puissance du L40S en profil « 48Q » (48 Go de VRAM) avec une machine virtuelle et quatre machines virtuelles exécutant le profil « 12Q ».
Bien qu'il soit absurde d'avoir une seule machine virtuelle monopolisant l'intégralité du L40S, les performances sont impressionnantes avec 21,147 2,514 points lors du benchmark GPU en un seul passage. Cependant, la division du GPU en quatre parties illustre l'impact de l'approche de découpage temporel de NVIDIA pour le partage des cœurs CUDA du GPU, avec des scores individuels allant de 2,567 XNUMX à XNUMX XNUMX lorsque le benchmark a été exécuté simultanément sur toutes les machines virtuelles.
En réexécutant le test sur une seule machine virtuelle avec le profil « 12Q », les trois autres étant inactives, le score remonte à 15,133 XNUMX. Il ne s'agit pas exactement d'un retour au score du GPU complet, mais il reste respectable pour un vGPU partitionné.
Benchmarks Blender
Continuons avec quelques benchmarks de rendu supplémentaires avec Blender. Suivant les tendances de Cinebench 2024, la division du GPU en quatre entraîne une baisse considérable des performances globales par rapport à une seule machine virtuelle exécutant la même charge de travail avec le même profil.
Comme le montre le benchmark Monster, avec seulement quatre machines virtuelles partageant la puissance de calcul du GPU, les performances de rendu individuelles peuvent atteindre jusqu'à 8 % de celles d'une seule machine virtuelle avec le même profil. Cependant, nous avons observé qu'une machine virtuelle prenait une avance considérable sur les autres, jusqu'à 2.4 fois le score de la machine la moins performante.
Les benchmarks Junkshop et Classroom racontent des histoires similaires, avec de fortes baisses de performances pour trois des quatre machines virtuelles et un seul invité obtenant un score beaucoup plus élevé que les autres.
Il est intéressant de noter qu'il semble y avoir de brefs moments où le vGPU d'une machine virtuelle est prioritaire et prend une avance significative. Par exemple, lors du benchmark Classroom, notre deuxième machine virtuelle Windows Server 2025 (WIN2025-2) a atteint plus de trois fois les performances de ses homologues, malgré des exécutions simultanées. Bien que nous ne puissions pas déterminer précisément si cela est dû à la planification du logiciel vGPU ou à la nature même du GPU, cela met en évidence certaines anomalies de performances inhérentes à l'approche de NVIDIA, basée uniquement sur le time-slicing, utilisée avec cette carte.
Conclusion
La configuration et la prise en charge du logiciel vGPU de NVIDIA ne sont peut-être pas aussi abouties que sur d'autres plateformes concurrentes. Il s'agit néanmoins d'une fonctionnalité intéressante et précieuse pour les entreprises et les particuliers qui utilisent déjà des systèmes d'environnement virtuel Proxmox. Bien que les performances soient considérablement réduites lors du partage des ressources GPU, de nombreuses entreprises continuent de tirer parti de la technologie vGPU de NVIDIA et ont déterminé que le partage d'un GPU compense cet inconvénient. Cette approche a été adoptée par de nombreux hyperscalers et centres de données à espace restreint, où l'intégration d'un maximum de locataires (en l'occurrence, des machines virtuelles avec vGPU) dans un espace minimal est l'option la plus efficace et la plus rentable.
