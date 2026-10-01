---
id: collect-261001-general-networking/general-networking/fr-review-truenas-scale-apps-improve-nas-flexibility-143e02cf-2
title: "fr-review-truenas-scale-apps-improve-nas-flexibility-143e02cf"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-general-networking/fr-review-truenas-scale-apps-improve-nas-flexibility-143e02cf.md
source_anchor: ""
source_lines: [43, 68]
sha256: 698457f900ab772413bcbce848a82805a31794d6ae8fe701ee8475f2b80427ff
---

# fr-review-truenas-scale-apps-improve-nas-flexibility-143e02cf

SCALE vous permet de déployer un conteneur Docker en sélectionnant « Lancer une image Docker » dans la section Applications et en le configurant directement via l'interface graphique. Cette méthode est cependant moins intuitive que l'utilisation du dépôt d'applications et nécessite de parcourir le dépôt Docker Hub ( https://hub.docker.com/search?q= ) pour trouver une image. Elle est plutôt destinée aux utilisateurs intermédiaires ayant une certaine expérience des conteneurs (Docker ou Kubernetes).
Entrer le nom du référentiel Docker et le configurer à partir du menu déroulant de l'interface graphique était agréable, mais si c'est votre première fois avec Docker, l'interface graphique ne vous tient pas la main et vous devrez consulter la base de connaissances.
La suppression des conteneurs inutiles est simple ; sélectionnez le conteneur à supprimer et cliquez sur supprimer.
Performances du iX Mini R
Alors que nous étions principalement intéressés par la vérification du catalogue et de l'intégration des applications TrueNAS SCALE, nous voulions pousser un peu le Mini R du côté des performances pour voir comment il se gère. Pour cet examen, notre unité était équipée de 4 disques SSD iX SATA de 1.9 To et de 8 disques durs WD Red Plus de 10 To. Notre unité comprenait également l'empreinte RAM améliorée de 64 Go.
Nous avons abordé les performances du flash à l'intérieur de notre unité d'examen iX Mini R. Nous avons configuré deux pools de stockage RAIDZ2, chacun avec la compression activée mais sans la déduplication activée. Nous avons examiné à la fois les performances externes sur la connexion 10GbE et les performances internes lorsque le stockage a été présenté à une machine virtuelle Server 2022 exécutée sur le système.
Notre processus de référence de stockage partagé et de disque dur d'entreprise préconditionne chaque disque dans un état stable avec la même charge de travail avec laquelle l'appareil sera testé sous une lourde charge de 16 threads avec une file d'attente exceptionnelle de 16 par thread, puis testé à intervalles définis dans plusieurs threads. / profils de profondeur de file d'attente pour afficher les performances en cas d'utilisation légère et intensive. Étant donné que les solutions NAS atteignent très rapidement leur niveau de performance nominal, nous ne représentons graphiquement que les principales sections de chaque test.
Tests de préconditionnement et d'état stable primaire :
- Débit (agrégat IOPS lecture + écriture)
- Latence moyenne (latence de lecture + écriture moyennée ensemble)
- Latence maximale (latence maximale de lecture ou d'écriture)
- Écart-type de latence (écart-type de lecture + écriture moyenné ensemble)
Notre analyse de charge de travail synthétique d'entreprise comprend quatre profils basés sur des tâches réelles. Ces profils ont été développés pour faciliter la comparaison avec nos références passées ainsi qu'avec des valeurs largement publiées telles que la vitesse de lecture et d'écriture maximale de 4k et 8k 70/30, qui est couramment utilisée pour les disques d'entreprise.
- 4K
  - 100 % de lecture ou 100 % d'écriture
- 8K (séquentiel)
  - 100 % de lecture ou 100 % d'écriture
- 128K (séquentiel)
  - 100 % de lecture ou 100 % d'écriture
En commençant par des performances de lecture aléatoire 4K sur 10 GbE sur nos quatre SSD en RAIDZ2, nous avons mesuré 3,545 1,017 IOPS en lecture et XNUMX XNUMX IOPS en écriture.
Ensuite, nous sommes passés à notre test de lecture et d'écriture séquentielle de 8 16, où nous avons mesuré un peu plus de 14.9 XNUMX IOPS en lecture et XNUMX XNUMX IOPS en écriture.
Enfin, lors de notre test de transfert séquentiel de 128k, nous avons mesuré 972 Mo/s en lecture et 738 Mo/s en écriture sur le fil.
Pour mesurer les performances du stockage interne à la box, sans la surcharge d'Ethernet, nous avons utilisé CrystalDiskMark à l'intérieur de la VM fonctionnant sur l'hyperviseur intégré. Ici, nous avons mesuré 811 Mo/s en lecture et 425 Mo/s en écriture avec une charge de travail de transfert séquentiel de 1 M.
Conclusion
Il y a beaucoup à explorer dans le catalogue d'applications TrueNAS SCALE, et pour la plupart, il est facile de travailler avec, en particulier pour ceux qui connaissent TrueNAS SCALE. Il y a d'autres endroits, cependant, où il est facile de se perdre dans les mauvaises herbes. La communauté voudra certainement continuer à créer de la documentation et des guides pour les utilisateurs qui souhaitent explorer en dehors des trains d'applications d'entreprise et officiels.
C'est également formidable de voir du matériel optimisé comme l'iX Mini R. Avec le passage à Linux et la puissance de ZFS, et la capacité d'héberger des conteneurs et des machines virtuelles, il s'agit d'un appareil capable pour les PME avec la polyvalence d'expansion à l'avenir. La croissance combinée continue de la plate-forme logicielle, avec davantage de prise en charge du matériel et des services de données d'entreprise et la variation croissante des plates-formes matérielles, sont toutes de bons pas en avant pour TrueNAS en général et SCALE en particulier, car iX continue de pousser plus loin la pile informatique de l'entreprise.
