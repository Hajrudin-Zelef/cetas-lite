---
id: collect-261001-general-networking/general-networking/fr-review-budget-truenas-core-system-showdown-84d7db31-1
title: "fr-review-budget-truenas-core-system-showdown-84d7db31"
domain: general-networking
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "attention", "ethernet"]
source: docs/RAG/collect-261001-general-networking/fr-review-budget-truenas-core-system-showdown-84d7db31.md
source_anchor: ""
source_lines: [1, 36]
sha256: dabab8bb625a8565e93bf547645387aec1e16e31e87020049450db9c08de7ece
---

# fr-review-budget-truenas-core-system-showdown-84d7db31

Il y a quelques mois, nous avons lancé notre concours de montage de NAS TrueNAS à 800 $ . En résumé, Brian, Ben et Kevin ont utilisé 800 $ pour construire leur propre système NAS en utilisant TrueNAS Core comme système d'exploitation. Grâce à Western Digital, ils n'ont pas eu à se soucier du stockage : WD leur a fourni un large choix de SSD et de disques durs . Nous voulions voir ce que notre équipe était capable de faire avec chaque système et comparer leurs performances. Ils ont dépensé l'argent, se sont lancés des défis et ont fait tester leurs systèmes. Découvrons maintenant leurs résultats.
Pour ceux qui n'ont pas suivi, nous avons réalisé une vidéo sur notre petit concours, que vous pouvez trouver intégrée ici ou sur notre page YouTube :
Systèmes TrueNAS CORE économiques
Le PC de Ben, le stagiaire, était sans doute le plus fait maison de tous. Il s'est rendu à la Cincinnati Computer Cooperative et chez MicroCenter, a acheté tous les composants séparément et les a assemblés lui-même. Son PC comprenait une alimentation OCZ GSX600, une carte mère ASRock B550, 64 Go de RAM G.Skill Ripjaws V (2 x 32 Go) DDR4-3600, un processeur Ryzen 5 3600, un boîtier Chelsio 111-00603+A0 et un boîtier Lian Li Liancool 205. Avec les quelques euros qui lui restaient, il a ajouté des bandes LED à sa configuration.
Kevin a utilisé le serveur HPE MicroServer Gen10 Plus équipé d'un processeur Xeon et de mémoire ECC. Il y a également ajouté une carte Mellanox ConnectX-5 100 GbE pour optimiser les performances par rapport aux autres configurations et simplifier la configuration réseau. Alors que les autres serveurs utilisent une carte réseau double port, Kevin n'a besoin de configurer qu'une seule interface 100 GbE.
La configuration de Brian se situe entre les deux autres. Il a commencé avec une carte mère Supermicro M11SDV-8CT-LN4F, équipée d'un processeur AMD EPYC 3201 SoC et de quatre ports Ethernet 1 GbE, ce qui a considérablement pesé sur son budget. Pour la mémoire vive, Brian a opté pour deux modules SK hynix PC4-2400T-RD1-11 DDR4 ECC de 8 Go. Il a également installé une alimentation Thermaltake de 500 W et une carte réseau 10 GbE. Le tout a été intégré dans un boîtier Fractal Design Node 304. Bien que Brian ait trouvé la carte réseau 10 GbE à un prix avantageux, celle-ci s'est avérée incompatible avec le logiciel TrueNAS, l'obligeant à utiliser une carte réseau Emulex de laboratoire. La mémoire vive d'occasion, provenant de Chine, a également posé problème et a dû être remplacée.
Système TrueNAS CORE économique - Performances
Passons maintenant à la véritable raison de notre présence ici : lequel des trois est le meilleur ? Outre nos trois configurations DIY, nous offrons également un TrueNAS Mini . La configuration iXsystems utilise RAIDZ2 puisqu'elle est livrée avec 5 disques durs. La plateforme iXsystems TrueNAS Mini X+ offre le meilleur compromis entre taille du châssis et compatibilité avec les disques. Elle prend en charge cinq disques durs 3.5" et dispose même de deux baies 2.5" pour SSD. Alors pourquoi ne pas la tester comme référence ? Tout simplement parce que le Mini X+ est optimisé pour une résilience maximale des données, et non pour les performances. Les trois autres ont été optimisés pour être les plus rapides de ce comparatif, ce qui comporte toutefois certains risques. Si iXsystems voulait surpasser nos concurrents, ils pourraient tout simplement les écraser avec une configuration bien plus performante.
Un mot rapide sur les configurations RAID : TrueNAS en prend en charge plusieurs selon la version. Comme nous avons utilisé des versions radicalement différentes, il y aura différentes configurations RAID. Les versions de Ben et Kevin utilisent RAIDZ sur quatre SSD, et la version de Brian utilise Mirror sur quatre disques durs.
Nous n'avons examiné que le protocole de partage de fichiers SMB pour cette confrontation. Un élément intéressant à mentionner est l'importance accordée à la configuration de la carte mère et du châssis. La plate-forme de bureau de Ben qui a sans doute l'air la plus cool, n'a que deux baies de lecteur de 3.5 pouces et est également de loin le plus grand boîtier.
Le boîtier de Brian prend en charge jusqu'à six baies de lecteur 3.5″ avec une attention particulière au refroidissement, mais sa carte mère ne dispose que de quatre ports SATA intégrés. Le microserveur HPE de Kevin construit en tant que version de base a quatre baies et quatre ports, mais c'est exactement ainsi que la plate-forme est conçue.
Le stockage diffère légèrement selon les modèles. La configuration de Brian comportait quatre disques durs WD Red de 10 To, mais malheureusement, le port M.2 NVMe ne fonctionnait pas correctement. Les configurations de Ben et Kevin, quant à elles, utilisaient quatre SSD WD Red de 4 To.
Il est important de noter dans la section des performances que la configuration RAID a un rôle énorme dans la façon dont les performances sont mesurées, au-delà de la sélection du lecteur lui-même. RAIDZ aura moins de surcharge que RAIDZ2, et Mirror aura encore moins de surcharge que RAIDZ. Cela dit, la configuration RAID doit tenir compte de l'application finale ultime, de la capacité dont vous avez besoin et de la résistance aux pannes que vous souhaitez pour votre construction. En fin de compte, ces résultats ne visent pas à montrer quel NAS est le plus rapide, mais plutôt comment les configurations TrueNAS fonctionnent sur des versions similaires, certaines utilisant les mêmes disques, dans différentes configurations RAID.
Analyse synthétique de la charge de travail d'entreprise
Notre processus de référence de stockage partagé et de disque dur d'entreprise préconditionne chaque disque dans un état stable avec la même charge de travail avec laquelle l'appareil sera testé sous une lourde charge de 16 threads avec une file d'attente exceptionnelle de 16 par thread, puis testé à intervalles définis dans plusieurs profils de profondeur de thread/file d'attente pour montrer les performances en cas d'utilisation légère et intensive. Étant donné que les solutions NAS atteignent très rapidement leur niveau de performance nominal, nous ne représentons graphiquement que les principales sections de chaque test.
Tests de préconditionnement et d'état stable primaire :
- Débit (agrégat IOPS lecture + écriture)
- Latence moyenne (latence de lecture + écriture moyennée ensemble)
- Latence maximale (latence maximale de lecture ou d'écriture)
- Écart-type de latence (écart-type de lecture + écriture moyenné ensemble)
Notre analyse de charge de travail synthétique d'entreprise comprend quatre profils basés sur des tâches réelles. Ces profils ont été développés pour faciliter la comparaison avec nos références passées ainsi qu'avec des valeurs largement publiées telles que la vitesse de lecture et d'écriture maximale de 4k et 8k 70/30, qui est couramment utilisée pour les disques d'entreprise.
- 4K
- 
  - 100 % de lecture ou 100 % d'écriture
  - 100% 4K
- 8K 70/30
  - 70 % de lecture, 30 % d'écriture
  - 100% 8K
- 8K (séquentiel)
  - 100 % de lecture ou 100 % d'écriture
  - 100% 8K
- 128K (séquentiel)
  - 100 % de lecture ou 100 % d'écriture
  - 100% 128K
Tout d'abord, notre test de débit de lecture/écriture 4K. Pour la lecture, le plus performant était Ben's avec 14,865 11,476 IOPS. Kevin est arrivé deuxième avec 595 3,868. Brian a terminé troisième avec 2,517 IOPS. Pour l'écriture, Kevin a pris la première place avec 923 XNUMX IOPS. Ben a décroché la deuxième place avec XNUMX XNUMX IOPS. Brian est resté troisième avec XNUMX IOPS.
Une grande partie de cela se résume au type de RAID déployé, bien que, avec le microserveur de Kevin par rapport à la construction DIY de Ben, la différence d'IOPS joue sur la vitesse du processeur dans chaque version.
