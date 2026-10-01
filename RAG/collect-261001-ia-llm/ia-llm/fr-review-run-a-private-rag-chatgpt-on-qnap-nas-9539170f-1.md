---
id: collect-261001-ia-llm/ia-llm/fr-review-run-a-private-rag-chatgpt-on-qnap-nas-9539170f-1
title: "fr-review-run-a-private-rag-chatgpt-on-qnap-nas-9539170f"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Nvidia", "OpenAI"]
dates: []
keywords: ["chatgpt", "amd", "exploit", "gpu", "nvidia"]
source: docs/RAG/collect-261001-ia-llm/fr-review-run-a-private-rag-chatgpt-on-qnap-nas-9539170f.md
source_anchor: ""
source_lines: [1, 31]
sha256: 59b1931b347d3267135432c13de784168b4611bb0b74310972d3c96324c06ef7
---

# fr-review-run-a-private-rag-chatgpt-on-qnap-nas-9539170f

QNAP est réputé pour la conception de son matériel, et notamment pour sa capacité à intégrer une puissance, une évolutivité et une flexibilité supérieures à celles de ses concurrents. Nous avons récemment testé le TS-h1290FX , un NAS NVMe de 12 pouces équipé d'un processeur AMD EPYC 7302P (16 cœurs/32 threads), de 256 Go de mémoire vive, d'une connectivité 25 GbE intégrée et de nombreux ports PCI. Avec une telle puissance et autant d'applications préinstallées , que se passe-t-il si nous ajoutons une carte graphique et voyons jusqu'où nous pouvons pousser ce NAS pour exécuter des applications d'intelligence artificielle, comme un serveur ChatGPT privé ?
Potentiel de stockage NAS pour l’IA
Le QNAP TS-h1290FX a beaucoup à offrir aux entreprises souhaitant se lancer dans l'IA. Ce NAS présente un avantage unique : il prend en charge un GPU interne et offre un potentiel de stockage massif. Les grands modèles d'IA nécessitent une quantité importante de données, qui doivent être stockées et accessibles efficacement. Cela peut s'avérer complexe pour les plateformes de stockage utilisant des disques durs, mais le TS-h1290FX, grâce à sa compatibilité U.2 NVMe, relève tous les défis.
Quand on pense aux NAS de grande capacité, on imagine souvent des plateformes avec disques durs 3.5 pouces compatibles avec des disques allant jusqu'à 24 To. Cela paraît énorme, mais c'est peu de chose comparé aux SSD QLC U.2. QNAP a récemment ajouté la compatibilité avec la gamme Solidigm P5336 , qui atteint une capacité incroyable de 61.44 To par disque. Pour un modèle à 12 baies comme le TS-h1290FX, les utilisateurs bénéficient de 737 To de stockage brut avant réduction des données. En matière de NAS compacts de bureau, rares sont les systèmes qui peuvent rivaliser.
Alors que les entreprises adoptent rapidement l’IA, disposer d’un système capable de fournir une capacité de stockage pour les flux de travail d’IA et d’exécuter des modèles constitue un énorme avantage. L'exploit impressionnant, cependant, est que ce NAS QNAP peut exécuter ces flux de travail d'IA tout en s'acquittant de ses tâches principales de partage du stockage dans l'environnement PME.
Il faut également dire que l’IA n’est pas une chose monolithique. Différents projets d'IA nécessitent différents types de stockage pour les prendre en charge. Bien que nous nous concentrions ici sur l'unité de bureau, QNAP propose de nombreux autres systèmes NAS prenant en charge le flash et la mise en réseau à haute vitesse, des éléments essentiels pour répondre à un besoin d'IA plus ambitieux que ce que nous avons couvert ici.
Comment QNAP prend-il en charge les GPU ?
QNAP prend en charge les GPU dans bon nombre de leurs systèmes NAS. Ils proposent également quelques applications prenant également en charge les GPU. Pour cet article, nous examinons principalement le GPU à travers le prisme de Virtualization Station. Virtualization Station est un hyperviseur pour le NAS QNAP, qui permet aux utilisateurs de créer une variété de machines virtuelles. Virtualization Station dispose également d'un ensemble de fonctionnalités approfondies qui prennent en charge les sauvegardes de VM, les instantanés, les clones et, plus important encore, le relais GPU dans le contexte de cet article.
À l'intérieur de notre unité de test, le QNAP TS-h1290FX est équipé d'une carte serveur typique avec plusieurs emplacements PCIe disponibles pour l'extension. QNAP fournit également les câbles d'alimentation GPU nécessaires à l'intérieur du châssis, donc aucune affaire amusante n'est requise pour les cartes qui nécessitent plus que l'alimentation du slot PCIe. Nous avons trouvé que le NVIDIA RTX A4000 à emplacement unique s'adaptait parfaitement avec suffisamment d'espace pour le refroidissement. Sur cette plateforme, un GPU avec un refroidisseur actif est préféré. Votre choix de GPU sera déterminé par la charge de travail et par ce que le NAS peut physiquement prendre en charge et refroidir.
Configuration de QNAP pour l'IA
Configurer une machine virtuelle (VM) avec transfert direct du GPU sur un NAS QNAP nécessite plusieurs étapes. Cela requiert un NAS QNAP compatible avec la virtualisation et doté des capacités matérielles nécessaires. Vous trouverez ci-dessous un guide expliquant comment configurer un NAS QNAP avec transfert direct du GPU.
1. Vérifier la compatibilité matérielle
Assurez-vous que votre NAS QNAP prend en charge Virtualization Station, qui est l'application de virtualisation de QNAP.
- Confirmez que le NAS dispose d'un emplacement PCIe disponible pour un GPU et que le GPU prend en charge le relais. Les listes de compatibilité sont souvent disponibles sur le site Web de QNAP. Bien que la liste de compatibilité actuelle ne prenne pas officiellement en charge le NVIDIA A4000, nous n'avons eu aucun problème avec les fonctionnalités.
2. Installez le GPU
- Éteignez le NAS et débranchez-le de l'alimentation. Ouvrez le boîtier et insérez le GPU dans un emplacement PCIe disponible. Connectez tous les câbles d’alimentation nécessaires au GPU. Fermez le boîtier, rebranchez l'alimentation et allumez le NAS.
3. Mettez à jour votre micrologiciel et logiciel QNAP
Assurez-vous que votre NAS QNAP exécute la dernière version de QTS (le système d'exploitation de QNAP). Nous avons utilisé Virtualization Station 4, qui est une version bêta ouverte de QNAP, pour offrir une meilleure prise en charge et de meilleures performances pour le travail GPU. Virtualization Station 4 est un package à installation automatique, contrairement à d'autres qui sont installés directement via QNAP App Center.
4. Installez le système d'exploitation sur la VM
Après avoir installé la Virtualization Station de QNAP sur votre NAS, vous pouvez accéder à l'interface de gestion pour déployer votre machine virtuelle (VM). Lorsque vous cliquez sur « Créer », une fenêtre d'invite apparaîtra pour vous permettre de fournir le nom de la VM et de sélectionner l'emplacement sur le NAS où la VM s'exécutera. Dans la plupart des cas, vous devrez peut-être apporter quelques ajustements mineurs aux informations sur le système d'exploitation et la version.
Ensuite, ajustez les ressources et le type de compatibilité du processeur que la VM verra au niveau du système d'exploitation invité. Dans notre cas, nous avons donné à notre VM 64 Go de mémoire et 8 processeurs. Nous avons sélectionné le type de processeur passthrough pour le modèle et modifié le BIOS en UEFI.
Pour démarrer et installer le système d'exploitation, vous devez télécharger et monter un fichier ISO en tant que lecteur de CD/DVD virtuel. Une fois le processus d'installation terminé, activez RDP pour la gestion avant de passer à l'étape suivante. La fonctionnalité de gestion des machines virtuelles QNAP change une fois le relais GPU activé, et RDP simplifie considérablement ce processus. À ce stade, éteignez la VM.
5. Configurer le relais GPU
Dans Virtualization Station :
- Avec la VM existante hors tension, modifiez votre VM.
- Dans le menu des paramètres de la VM, recherchez l'onglet Périphériques physiques. À partir de là, sélectionnez PCIe. Vous verrez un appareil disponible pour le relais. Dans notre cas, il s’agissait du NVIDIA RTX A4000. Appliquez ce changement.
- Si vous devez allouer d'autres ressources à votre VM, telles que des cœurs de processeur, de la RAM et du stockage, c'est le moment de le faire.
- Rallumez la VM.
6. Installez les pilotes GPU dans la VM
Une fois de retour dans la VM en utilisant RDP avec le GPU connecté, téléchargez et installez les pilotes appropriés pour votre GPU dans la VM. Cette étape est cruciale pour que le GPU fonctionne correctement et fournisse les améliorations de performances attendues.
7. Vérifier la fonctionnalité GPU Passthrough
