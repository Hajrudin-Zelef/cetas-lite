---
id: collect-261001-general-networking/general-networking/fr-review-unlocking-the-power-of-nas-exploring-qnaps-top-applications-7c4f1527-1
title: "fr-review-unlocking-the-power-of-nas-exploring-qnaps-top-applications-7c4f1527"
domain: general-networking
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-general-networking/fr-review-unlocking-the-power-of-nas-exploring-qnaps-top-applications-7c4f1527.md
source_anchor: ""
source_lines: [1, 27]
sha256: dabe032954e037d85b6d27473cd20e73b9943b91e02edc5717a2b6bdb4cecf7e
---

# fr-review-unlocking-the-power-of-nas-exploring-qnaps-top-applications-7c4f1527

QNAP a une longue histoire et un catalogue impressionnant de périphériques de stockage en réseau prenant en charge le stockage personnel, les petites et les grandes entreprises. Ils proposent des solutions de stockage suffisamment petites pour sécuriser les albums photo et les collections musicales, ainsi que des monstres de stockage de données massifs à l'échelle du pétaoctet pour sauvegarder des centres de données entiers. Les appareils de stockage en réseau (NAS) QNAP sont populaires dans un certain nombre de cas d'utilisation allant des PME aux amateurs. Le système d'exploitation comprend plusieurs applications pour aider les administrateurs à configurer et à étendre l'utilisation des appareils.
Plus tôt cette année, nous avons testé un NAS surpuissant, le QNAP TVS-h874 . Aujourd'hui, nous allons utiliser cette machine pour passer en revue certaines des applications les plus populaires et utiles de QNAP disponibles sur son App Center.
Le QNAP App Center est similaire aux autres magasins d'applications et centres de téléchargement, offrant une large sélection d'applications typiques pour les appareils QNAP. Cependant, ce n’est pas le seul endroit pour acquérir des candidatures. Nous ne couvrirons aucun emplacement de téléchargement tiers dans cet article, à moins qu'ils ne soient fournis directement et approuvés par QNAP.
Applications QNAP natives
La plupart des appareils NAS sont préchargés avec une petite sélection d'applications pour aider les utilisateurs à faire fonctionner leurs appareils et à offrir des services supplémentaires qu'ils n'exécutent peut-être pas déjà dans leur environnement.
MonQNAPcloud
L'une des applications les plus connues de l'écosphère est « myQNAPcloud ». Cette application principale est fournie aux propriétaires en tant que solution de simplification de connexion et de partage à distance. Au lieu d'avoir à parcourir des configurations réseau difficiles pour activer un service DNS dynamique, myQNAPcloud fournit aux propriétaires leur propre « SmartURL » unique et facile à retenir.
https://q.link.to/ <customlinkname>
Grâce à cette application, les administrateurs peuvent créer des dossiers et des partages de fichiers hautement granulaires et sécurisés. Ils ont deux possibilités : SmartShare, qui est étonnamment similaire aux autorisations, mots de passe, etc. LDAP de Windows, et les services de publication, où une page Web est créée qui permet à d'autres utilisateurs disposant d'un accès approprié d'afficher et de modifier des fichiers.
QuObjects
Cette application est un service de stockage d'objets dans le cloud doté de fonctionnalités similaires aux compartiments S3 d'AWS. Elle est conçue pour offrir aux utilisateurs de NAS QNAP une solution simple et évolutive pour stocker et gérer de grandes quantités de données non structurées, telles que des photos, des vidéos et des fichiers, dans un environnement de type cloud.
Ce service est avantageux pour les utilisateurs qui ont besoin d'un moyen flexible et efficace de gérer des ensembles de données volumineux. Avec QuObjects, les utilisateurs peuvent profiter des avantages du stockage cloud, tels que l'accessibilité à distance, une évolutivité facile et une protection robuste des données, tout en conservant le contrôle et la sécurité d'un système NAS local. Il est particulièrement utile pour les entreprises ou les particuliers qui ont besoin d'une solution rentable de stockage de données combinant les avantages des systèmes locaux et basés sur le cloud.
Échelle de queue
Tailscale pour NAS QNAP transforme votre stockage réseau en un hub sécurisé et facile d'accès. En téléchargeant Tailscale depuis le QNAP App Center, vous pouvez accéder à votre NAS de n'importe où sans ouvrir de ports dans votre pare-feu, pour plus de simplicité et de sécurité. Vous pouvez ainsi partager votre NAS QNAP avec des utilisateurs Tailscale désignés et bénéficier d'un contrôle d'accès sélectif.
De plus, vous pouvez mettre en œuvre des listes de contrôle d'accès (ACL) pour restreindre l'accès, garantissant ainsi que seules les personnes autorisées peuvent accéder aux données sensibles. De plus, votre NAS QNAP peut fonctionner comme un routeur de sous-réseau, fournissant un accès externe à votre réseau local. Cette fonctionnalité remplace efficacement le besoin d'un serveur VPN autonome traditionnel, rationalisant ainsi la configuration de votre réseau.
Enfin, dans des scénarios tels que l'utilisation d'un café Internet, votre NAS QNAP peut servir de nœud de sortie, vous offrant un accès Internet sécurisé même à partir d'emplacements non fiables. Cette solution complète offre un mélange d'accessibilité, de sécurité et de polyvalence pour vos besoins NAS et réseau.
AMIZ Cloud et Centre d'Organisation
L’un des défis auxquels sont confrontés de nombreux administrateurs informatiques est la gestion de plusieurs appareils et services au sein de dizaines d’interfaces, de portails et de sites Web. Je peux attester que le fait d'avoir moins de portails de gestion combinant des appareils et des services dans un seul (ou du moins moins) volets constitue une amélioration considérable du flux de travail et de la productivité.
QNAP a relevé ce défi avec son AMIZ Cloud Center . Cette application permet aux administrateurs de centraliser tous leurs périphériques QNAP sur un portail de gestion unique. Elle offre les outils nécessaires pour assurer le suivi et la maintenance du NAS, surveiller l'utilisation des ressources et consulter les journaux. Elle permet également de gérer les mises à jour système et applicatives et, grâce à l'installation de l'application Security Counselor, de surveiller et de résoudre les problèmes de sécurité en temps réel.
AMIZ Cloud and Organization Center est l'une des applications indispensables pour tout administrateur QNAP responsable de plusieurs appareils et services cloud.
Virtualisation et conteneurisation
Station de virtualisation
QNAP Virtualization Station est une solution complète pour les NAS QNAP. Lancée en 2014, elle permet d'exécuter plusieurs machines virtuelles (VM) utilisant des images natives Windows ou Linux. Chaque VM dispose de matériel virtualisé indépendant, incluant interfaces réseau, disques et cartes graphiques.
Une fois l'application installée sur le NAS, l'utilisateur est redirigé vers la Place de marché des machines virtuelles, où il peut rechercher des applications préconfigurées sous forme d'images de machines virtuelles correspondant à ses besoins. Parmi ces images d'applications figurent pfSense, Zabbix, AWS File Gateway et bien d'autres.
Bien entendu, la possibilité d’installer votre propre VM est toujours disponible. QNAP simplifie le processus de création de VM et est étonnamment similaire à celui de vSphere de VMware.
Station de conteneurs
QNAP propose également Container Station comme application d'abstraction numérique . Elle intègre LXD et Docker, technologies de virtualisation légères de Kata, permettant d'exécuter plusieurs systèmes Linux isolés sur un NAS QNAP et de télécharger des applications depuis le registre intégré Docker Hub/LXD Image Server.
