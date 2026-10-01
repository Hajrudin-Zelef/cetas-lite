---
id: collect-261001-huawei/huawei/fr-review-artesca-veeam-unified-software-appliance-74a26655-3
title: "fr-review-artesca-veeam-unified-software-appliance-74a26655"
domain: huawei
role: reference
task: reference
actors: ["Microsoft", "Oracle"]
dates: []
keywords: []
source: docs/RAG/collect-261001-huawei/fr-review-artesca-veeam-unified-software-appliance-74a26655.md
source_anchor: ""
source_lines: [33, 53]
sha256: dc4dbb0bad7bc0dd7088f2daea6bd28276a793e19cbf6bf17ad9205d324a6501
---

# fr-review-artesca-veeam-unified-software-appliance-74a26655

L'interface de l'appliance logicielle unifiée ARTESCA+ Veeam est constituée de l'instance Veeam intégrée. Fonctionnant sur KubeVirt, cette machine virtuelle Windows Server, équipée de Veeam Backup & Replication, prend en charge la sauvegarde de machines virtuelles issues de divers hyperviseurs, notamment Proxmox VE, Microsoft Hyper-V, VMware vSphere (ESXi/vCenter), Nutanix AHV, Oracle Linux Virtualization Manager et bien d'autres. Face à la migration de nombreux clients vers des solutions de virtualisation autres que celles basées sur VMware, Veeam prévoit d'étendre prochainement la prise en charge à d'autres hyperviseurs populaires, tels que XCP-NG, HPE Morpheus VM Essentials, Citrix XenServer et Red Hat OpenShift Virtualization. Consultez les dernières actualités concernant Veeam Data Platform v13.
Les avantages d'une appliance logicielle unifiée
L'intégration de Veeam améliore les performances, l'efficacité et la sécurité de l'appliance logicielle Scality. Le transfert de données entre Veeam et le référentiel de stockage S3, assuré par ARTESCA+, s'effectue en interne via un réseau virtualisé. Cette approche élimine la nécessité de dupliquer le trafic réseau externe entre les hyperviseurs, les sources de données, Veeam et le référentiel de stockage. Elle permet également de protéger le compartiment S3 contre les attaques, réduisant ainsi la surface d'attaque de la solution.
Pour configurer le compartiment S3 utilisé par la machine virtuelle Veeam intégrée, connectez-vous à l'interface web d'ARTESCA+ et accédez à l'onglet « Comptes » de la page « Gestion des données ». Cliquez ensuite sur le bouton « Démarrer l'assistant Veeam VBR », ce qui affichera un assistant permettant de configurer un compte de stockage et un compartiment.
Dans notre laboratoire, l'assistant Veeam intégré a permis de déployer un référentiel fonctionnel et immuable en quelques minutes. Il a créé le compte ARTESCA, activé le versionnage et le verrouillage d'objet sur le compartiment, configuré le reporting de capacité SOSAPI et nous a fourni le point de terminaison, la région et les clés nécessaires à Veeam. Grâce à la communication interne entre Veeam et le service S3, aucune ressource n'a été exposée sur le réseau et aucune configuration DNS n'a été requise, évitant ainsi les allers-retours habituels entre les équipes de stockage et de sauvegarde. Le processus est simple pour les informaticiens généralistes et, d'après notre expérience, il s'est avéré extrêmement fiable tout en conservant des options de contrôle avancées.
Une fois la configuration initiale du logiciel terminée, l'instance Windows Server exécutant Veeam Backup & Replication est accessible via le protocole RDP (Remote Desktop Protocol) ou VNC. Veeam doit ensuite être connecté au compartiment de stockage S3 d'ARTESCA+ à l'aide de l'onglet « Référentiels de sauvegarde » de la page « Infrastructure de sauvegarde ». Cliquez avec le bouton droit sur l'onglet, puis sélectionnez « Ajouter un référentiel de sauvegarde » pour ajouter le compartiment S3 créé précédemment.
Sélectionnez « Stockage d'objets » dans la liste des types de référentiels de sauvegarde, puis choisissez « Compatible S3 » sur les deux pages suivantes. Un assistant s'affichera alors, vous invitant à saisir les informations de connexion du compartiment S3 que vous avez créé. Le « Point de service » se trouve dans l'onglet « Services de données » de la page « Gestion des données » de l'interface web d'ARTESCA+ (dans notre cas, il s'agissait de « s3.artesca-plus-veeam.local »). Indiquez la région (disponible dans l'onglet « Emplacements ») et les identifiants, sélectionnez le compartiment et le dossier créés précédemment, puis configurez le « Serveur de montage » pour qu'il s'agisse du serveur Veeam intégré à l'appliance logicielle unifiée ARTESCA+ Veeam. Vous pouvez également définir des paramètres supplémentaires, tels qu'une limite d'objets et activer ou non la restauration quasi instantanée avec le service vPower NFS.
Veeam est maintenant prêt à se connecter à une source de données ou à un hyperviseur pour démarrer les sauvegardes.
ARTESCA+ Veeam et Proxmox VE
Pour tester l'appliance logicielle unifiée ARTESCA+ Veeam, nous avons connecté l'unité de test à notre serveur Proxmox VE (PVE). La sauvegarde des machines virtuelles s'effectue en trois étapes simples :
- Connexion du cluster PVE à Veeam
- Ajout d'une machine virtuelle de travail
- Création d'une tâche de sauvegarde
Nous vous présenterons les étapes de démarrage avec Veeam, mais un guide plus détaillé est disponible ici : Guide de l’utilisateur du plug-in Veeam pour Proxmox VE 3. Pour connecter Veeam à PVE via la console Backup & Replication, accédez à la vue « Inventaire », puis sélectionnez « Ajouter un serveur » après avoir cliqué avec le bouton droit sur l’élément de liste « Infrastructure virtuelle ».
Lorsque le menu « Ajouter un serveur » apparaît, suivez les instructions de l’assistant pour connecter une instance Proxmox VE. Il vous sera demandé l’adresse IP ou le nom d’hôte du serveur ainsi que les identifiants ; assurez-vous donc d’avoir ces informations à portée de main pour votre cluster.
Après avoir connecté le PVE, une machine virtuelle de travail est nécessaire pour faciliter les transferts de données entre le serveur Proxmox VE et Veeam. Dans la vue « Infrastructure de sauvegarde », cliquez avec le bouton droit sur « Proxies de sauvegarde », puis sélectionnez « Ajouter un proxy » dans le menu déroulant.
Ajoutez un processus worker à l'aide de l'assistant « Proxmox VE worker » du menu contextuel, puis sélectionnez l'hôte, le processeur, la mémoire et la configuration réseau souhaités. Nommez-le (nous avons utilisé « artesca-worker ») et terminez l'assistant.
Enfin, nous devons créer une tâche de sauvegarde pour démarrer la sauvegarde des machines virtuelles sur Proxmox VE. Dans la vue « Accueil » de Veeam Backup & Replication, cliquez sur « Tâches », puis sur « Sauvegarde ». Faites un clic droit sur « Sauvegarde », survolez « Sauvegarde » dans le menu déroulant et sélectionnez « Machine virtuelle ».
Un menu s'affichera alors, vous permettant de sélectionner les paramètres de la tâche de sauvegarde. Veillez à inclure toutes les machines virtuelles à sauvegarder en utilisant le bouton « Ajouter » à l'étape « Machines virtuelles ». Sélectionnez également le compartiment de stockage ARTESCA+ S3 comme référentiel de sauvegarde. Une fois ces étapes terminées, la tâche de sauvegarde s'exécutera selon votre planification.
Une fois que vous avez configuré Veeam pour sauvegarder et stocker vos hyperviseurs et vos sources de données, l'appliance logicielle unifiée ARTESCA+ Veeam fonctionnera de manière fiable en arrière-plan, protégeant les informations de votre organisation et vous assurant une tranquillité d'esprit.
L'avenir de Scality et Veeam
