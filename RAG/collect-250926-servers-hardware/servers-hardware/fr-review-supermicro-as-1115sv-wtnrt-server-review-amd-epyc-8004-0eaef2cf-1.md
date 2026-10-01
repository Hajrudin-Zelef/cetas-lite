---
id: collect-250926-servers-hardware/servers-hardware/fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf-1
title: "fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Broadcom"]
dates: []
keywords: ["amd", "gpu"]
source: docs/RAG/clean4/fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf.md
source_anchor: ""
source_lines: [1, 42]
sha256: 2ccecbf9a0aef26559150bd569729b91f39443bb271d9b42da448375ee498af4
---

# fr-review-supermicro-as-1115sv-wtnrt-server-review-amd-epyc-8004-0eaef2cf

Le serveur Supermicro AS-1115SV-WTNRT est un serveur rackable 1U polyvalent et robuste, conçu pour répondre aux exigences des applications les plus gourmandes en ressources telles que la virtualisation, la gestion de bases de données et l'informatique de périphérie. Adapté aux environnements exigeant un débit élevé et une grande fiabilité, ce serveur offre une architecture équilibrée intégrant le processeur AMD EPYC série 8004 , fournissant ainsi une solution évolutive aux défis informatiques modernes.
Supermicro AS-1115SV-WTNRT Caractéristiques principales et spécifications matérielles
Au cœur de l'AS-1115SV-WTNRT se trouve une configuration à socket unique utilisant le processeur AMD EPYC série 8004, capable de gérer jusqu'à 64 cœurs et 128 threads, ce qui permet une puissance de traitement parallèle importante. Conçus pour les centres de données denses et les fournisseurs de services, les processeurs de la série « Siena » offrent une combinaison optimale de nombre élevé de cœurs et d'efficacité énergétique à un prix compétitif, fonctionnant dans une enveloppe de puissance modeste commençant à seulement 70 W et pouvant atteindre 200 W.
Le serveur prend en charge jusqu'à 576 Go de mémoire DDR5 à 4800 2.5 MHz sur six emplacements DIMM, garantissant un transfert de données à grande vitesse et une gestion efficace des ensembles de données grand public. La polyvalence du stockage est un autre attribut important, avec dix baies de disques XNUMX″ remplaçables à chaud qui prennent en charge une combinaison de disques NVMe, SAS et SATA.
Le Supermicro AS-1115SV-WTNRT est conçu avec un facteur de forme compact 1U, mesurant 437 mm x 43 mm x 597 mm, optimisant l'espace sans compromettre la capacité matérielle interne. Les capacités d'extension incluent trois emplacements PCIe 5.0 x16, deux sont des emplacements pleine hauteur et pleine longueur, offrant suffisamment d'espace pour les cartes plus grandes, tandis que le troisième est un emplacement à profil bas, adapté aux cartes d'extension plus petites.
Le serveur dispose de deux blocs d'alimentation redondants de 860 W pour une alimentation continue et la fiabilité du système.
Sous le capot, la disposition interne du serveur optimise le flux d'air et l'efficacité du refroidissement, assurant des performances constantes dans diverses conditions de fonctionnement, y compris les six ventilateurs internes contrarotatifs (FAN-0163L4).
La plus petite carte mère H13SVW-NT est à l'avant-plan, hébergeant un seul processeur AMD EPYC série 8004 dans un socket SP6, entouré de ses six emplacements DIMM DDR5.
Supermicro AS-1115SV-WTNRT La gestion du système
Supermicro a équipé l'AS-1115SV-WTNRT d'une suite robuste d'outils de gestion, notamment SuperCloud Composer et Supermicro Server Manager (SSM), qui facilitent la gestion et la surveillance rationalisées du système. Le BIOS est équipé de la spécification contemporaine UEFI 2.8, offrant une sécurité améliorée, des temps de démarrage plus rapides et la prise en charge de partitions de disque plus grandes.
L'AS-1115SV-WTNRT dispose également d'un système complet de contrôleur de gestion de la carte mère (BMC) qui offre aux administrateurs des capacités de contrôle et de surveillance. Cette interface BMC (accessible via une adresse IP dédiée) fournit des informations critiques sur l'état du système, permet des modifications de configuration complètes et permet des actions de contrôle à distance sans accès physique au matériel. Il s'agit d'un outil essentiel pour gérer efficacement les opérations du serveur, garantir la disponibilité et résoudre rapidement les problèmes.
Le tableau de bord offre un aperçu complet de l'état du serveur et des mesures essentielles. Il présente des indicateurs d'état du système, des adresses IP et des versions de micrologiciel, fournissant un instantané des données opérationnelles critiques. La consommation électrique est affichée graphiquement au fil du temps, détaillant l'utilisation minimale, moyenne et maximale. Cette visualisation permet aux administrateurs d'évaluer en un coup d'œil l'efficacité énergétique et les coûts opérationnels du serveur. Un aperçu de la console à distance est également disponible, facilitant un accès rapide à l'état opérationnel actuel du serveur, ce qui est inestimable pour un dépannage rapide et une surveillance continue.
La section CPU offre des informations détaillées sur les spécificités du processeur du serveur (dans notre cas, le processeur AMD EPYC 8534P 64 cœurs), telles que la vitesse, la puissance thermique de conception (TDP), le nombre de cœurs, le nombre de threads et le fabricant. Cette section est cruciale pour vérifier les paramètres opérationnels du processeur et garantir qu'il fonctionne dans les limites spécifiées.
Dans l'onglet Mise à jour du micrologiciel, les utilisateurs peuvent gérer et mettre à jour le micrologiciel de divers composants du serveur, tels que BMC, BIOS et CPLD. L'interface permet aux utilisateurs de choisir le type de micrologiciel, de conserver les configurations existantes et de garantir que toutes les mises à jour critiques sont appliquées sans affecter les paramètres opérationnels du serveur. Cette fonction est essentielle pour maintenir la sécurité du serveur, car les mises à jour du micrologiciel contiennent souvent des correctifs pour les vulnérabilités et l'amélioration des performances.
L'onglet Mémoire fournit des informations détaillées sur tous les modules de mémoire installés, affichant l'état de santé de chaque module, son type, la capacité du code de correction d'erreur, la vitesse de fonctionnement, la taille et le numéro de série. L'onglet Mémoire permet aux utilisateurs de suivre et de gérer facilement la mémoire physique du serveur, ce qui est essentiel pour diagnostiquer les problèmes liés à la mémoire ou planifier les mises à niveau.
L'onglet Alimentation affiche des statistiques détaillées de consommation d'énergie sur différentes périodes, y compris la dernière heure, le jour et la semaine. Il affiche les tendances historiques et les valeurs maximales, utiles pour la planification des capacités et l'analyse opérationnelle. Les utilisateurs et les administrateurs peuvent également utiliser ces données pour optimiser les paramètres d'alimentation, planifier des modes basse consommation pendant les heures creuses et prendre des décisions éclairées en matière d'utilisation et d'efficacité énergétique.
Spécifications Supermicro AS-1115SV-WTNRT
| Processeur |  | 
| Processeur | Processeur(s) unique(s) – Processeur AMD EPYC série 8004 | 
| Nombre de noyaux | Jusqu'à 64C/128T | 
| Note | Prend en charge les processeurs TDP jusqu'à 225 W (refroidis par air) | 
| GPU |  | 
| Nombre maximal de GPU | Jusqu'à 1 GPU double largeur ou 2 GPU simple largeur | 
| Mémoire système |  | 
| Mémoire | Nombre d'emplacements : 6 emplacements DIMM – Mémoire maximale (1DPC) : jusqu'à 576 Go 4800 5 MT/s ECC DDRXNUMX RDIMM | 
| Appareils embarqués |  | 
| Chipset | Système sur puce | 
| Connectivité réseau | 2 RJ45 10GBASE-T avec Broadcom BCM57416 | 
| Entrée / Sortie |  | 
| LAN | 1 port(s) LAN IPMI dédié(s) RJ45 1 GbE – 2 port(s) LAN RJ45 10 GBASE-T | 
| USB | 2 ports (avant) – 4 ports (arrière) | 
| Vidéo | 1 port(s) VGA | 
| Port série | 1 port(s) COM (arrière) | 
| TPM | 1 TPM intégré/port 80 | 
| BIOS système |  | 
| Type de BIOS | AMI Flash EEPROM 32 Mo SPI | 
| Fonctionnalités du BIOS | ACPI 6.4, SMBIOS 3.5 ou version ultérieure, UEFI 2.8, Plug and Play (PnP), prise en charge du clavier USB | 
| Direction |  | 
| Logiciels |  | 
| Configurations d'alimentation | Alimentation CA Titanium redondante 1U 800/860 W | 
| Sécurité |  | 
| Hardware | Trusted Platform Module (TPM) 2.0, Silicon Root of Trust (RoT) – Conforme NIST 800-193 | 
