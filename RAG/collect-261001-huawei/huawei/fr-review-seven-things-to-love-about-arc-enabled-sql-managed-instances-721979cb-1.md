---
id: collect-261001-huawei/huawei/fr-review-seven-things-to-love-about-arc-enabled-sql-managed-instances-721979cb-1
title: "fr-review-seven-things-to-love-about-arc-enabled-sql-managed-instances-721979cb"
domain: huawei
role: reference
task: reference
actors: ["Intel", "Microsoft"]
dates: []
keywords: ["intel"]
source: docs/RAG/collect-261001-huawei/fr-review-seven-things-to-love-about-arc-enabled-sql-managed-instances-721979cb.md
source_anchor: ""
source_lines: [1, 26]
sha256: 22d145fdcf273c5ba9976488ecba0700453e6c49ab9faf8677705c14f03b79d0
---

# fr-review-seven-things-to-love-about-arc-enabled-sql-managed-instances-721979cb

Suite à notre récent article sur les services managés compatibles avec Azure Arc , nous avons poursuivi notre exploration de la puissance d'Azure Arc et d'Azure Stack HCI avec DataON, partenaire de Microsoft et d'Intel. Nous avons rapidement constaté leurs avantages et un cas d'usage s'est particulièrement distingué : l'instance SQL Server managée compatible avec Azure Arc. Cette instance est une plateforme en tant que service (PaaS) qui utilise la dernière version de SQL Server (Enterprise Edition), mise à jour et sauvegardée automatiquement. De plus, pour les applications critiques, elle intègre une haute disponibilité.
En explorant SQL Managed Instance compatible avec Azure Arc, nous avons découvert plusieurs fonctionnalités uniques, intéressantes ou puissantes. Ces éléments sont développés ci-dessous.
Tirer parti de la puissance de l'infrastructure hyperconvergée
Très tôt, les entreprises ont découvert la puissance du cloud public Azure et les services qu'il pouvait fournir. Cependant, pour certaines charges de travail, il est nécessaire de les conserver sur site pour des raisons de conformité. Azure Stack HCI répond aux exigences réglementaires en utilisant la puissance et les services offerts par Azure (y compris l'instance SQL gérée par Arc), permettant à ces charges de travail de s'exécuter sur le matériel de l'entreprise à l'emplacement de son choix.
DataON, l'une des entreprises avec lesquelles nous travaillons en partenariat, a été l'un des premiers à adopter ces technologies et nous a aidés à mieux les comprendre.
Avec Azure Arc, les clients peuvent afficher et gérer leurs applications et leurs bases de données de manière cohérente avec un ensemble d'outils et une interface familiers, quel que soit l'endroit où ces services s'exécutent, du local au multicloud en passant par la périphérie.
Désormais, chaque nœud de cluster Azure Stack HCI est activé pour Arc lors de l'inscription d'un cluster auprès d'Azure. Cela signifie que toutes ces puissantes fonctionnalités de gestion Azure sont disponibles pour vos nœuds Azure Stack HCI.
Adopter la sécurité basée sur le matériel
Au début de son développement, Microsoft a donné la priorité à la sécurité lors de la création d'Azure Stack HCI, Arc et SQL Managed Instance compatible avec Arc. Microsoft et Intel ont collaboré pour fournir une solution de sécurité complète avec Azure Stack HCI, couvrant l'ensemble de l'infrastructure informatique. Ils ont également intégré Azure Arc pour étendre la sécurité basée sur Azure aux environnements hybrides et multi-cloud. La sécurité et les extensions intégrées d'Intel renforcent encore cette solution, assurant une protection complète du silicium au cloud.
Les mesures de sécurité d'Intel garantissent que les appareils et les données sont fiables, tout en fournissant une accélération de la charge de travail et du chiffrement. Cela permet une protection des données sécurisée et isolée du matériel et une fiabilité des logiciels afin de se prémunir contre les cybermenaces.
La plate-forme d'Azure intègre des outils et des contrôles de sécurité facilement accessibles et conviviaux. Les contrôles natifs de DevOps et de Security Center peuvent être personnalisés pour protéger et superviser toutes les ressources cloud et tous les niveaux d'architecture. Microsoft a développé Azure en utilisant les principes de confiance zéro standard de l'industrie, qui impliquent une vérification explicite et l'hypothèse qu'une violation s'est produite.
La sécurité commence au niveau matériel. L'utilisation d'un serveur Secured-core et d'un tableau de bord, disponible via Azure Stack HCI, permet la vérification et l'audit du matériel pour s'assurer que le serveur répond aux exigences de Secured-core.
L'engagement avec DataON (un partenaire Intel Platinum) garantit que la base matérielle pour un déploiement sur site d'Azure Stack HCI utilise les derniers serveurs basés sur Intel pour répondre aux exigences du serveur Secured-core. TPM2.0, Démarrage sécurisé, Virtualization Based Security (VBS), Hypervisor-protected Code Integrity, Pre-boot DMA protection et DRTM protection sont quelques fonctionnalités de sécurité fournies par les serveurs Intel et vérifiées par Azure Stack HCI.
Exploiter la puissance de Kubernetes
L'instance SQL gérée par Arc exploite Kubernetes (K8s) pour héberger l'instance SQL et fournir des fonctionnalités de gestion supplémentaires pour ces instances SQL. K8s est une technologie éprouvée (elle existe depuis environ une décennie) dans le centre de données, et en l'utilisant, Microsoft capitalise sur ses caractéristiques et fonctions et sur son écosystème puissant et riche.
SQL Managed Instance compatible avec Arc masque la complexité de l'exécution de conteneurs via des tableaux de bord et des assistants tout en permettant aux autres de travailler directement avec les K8.
Tarification SQL transparente et instantanée
Les coûts de licence pour votre instance SQL gérée par Arc sont calculés et affichés au fur et à mesure que l'instance est configurée, révélant le coût de la base de données avant le déploiement. Cela permet également aux clients d'effectuer des calculs de simulation et de peser les compromis au moment de décider quoi déployer. Par exemple, vous pouvez déterminer si vous voulez une, deux ou trois répliques pour la haute disponibilité ou tout autre attribut que SQL Managed Instance compatible avec Arc peut fournir. Ces informations sur les coûts évitent les surprises à la fin du mois et permettent aux secteurs d'activité de configurer leurs instances en fonction de leurs budgets.
En prime, si vous possédez déjà une licence SQL Server, vous pouvez utiliser l'avantage Azure Hybrid pour économiser sur les coûts de licence.
Simplification de la création de nouvelles bases de données
Comme Azure Arc est piloté par des stratégies, un administrateur ou même l'utilisateur final d'une base de données peut créer une nouvelle instance gérée SQL à l'aide de l'interface Web Azure. Azure Stack HCI regroupe tout le calcul et le stockage des serveurs sous son contrôle. Ainsi, la création d'une nouvelle base de données implique de sélectionner les attributs nécessaires, mais de ne pas avoir à décider quels composants individuels et discrets sont utilisés pour l'hébergement.
En seulement quelques minutes de déploiement, une instance gérée SQL hautement disponible compatible Arc avec des fonctionnalités intégrées telles que les sauvegardes automatisées, la surveillance, la haute disponibilité, la reprise après sinistre, etc., sera prête à l'emploi.
Pour utiliser la base de données, SQL Managed Instance compatible avec Arc fournit une liste de chaînes de connexion pour les langages de programmation courants. C'est un petit changement, mais cela peut épargner beaucoup de frustration aux programmeurs qui cherchent à s'y connecter.
La migration des bases de données existantes est un jeu d'enfant
À l'aide du service de migration de données Azure entièrement automatisé de Microsoft, le déplacement d'une base de données vers Azure Stack HCI en tant qu'instance SQL gérée par Arc est un jeu d'enfant. Même pour les professionnels qualifiés et expérimentés, la migration vers une base de données peut être une perspective anxiogène. Microsoft a créé un assistant pour guider les utilisateurs tout au long du processus, en supprimant le stress de le faire soi-même ou les frais de sous-traitance.
Surveillance intégrée
