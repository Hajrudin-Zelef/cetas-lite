---
id: collect-261001-cisco/cisco/fr-review-amd-epyc-4005-review-am5-economics-with-enterprise-focus-5f50297a-2
title: "fr-review-amd-epyc-4005-review-am5-economics-with-enterprise-focus-5f50297a"
domain: cisco
role: reference
task: reference
actors: ["AMD", "Nvidia", "Samsung"]
dates: []
keywords: ["amd", "apache", "benchmark", "gpu", "nvidia", "open source"]
source: docs/RAG/collect-261001-cisco/fr-review-amd-epyc-4005-review-am5-economics-with-enterprise-focus-5f50297a.md
source_anchor: ""
source_lines: [48, 103]
sha256: 7cd78d11e1b6391f6f26585c19b29802eddd18cab34fe47c665bc3255fede49e
---

# fr-review-amd-epyc-4005-review-am5-economics-with-enterprise-focus-5f50297a

| Gestion de serveur | (1) Port de gestion de serveur dédié 1000Base-T (Realtek RTL8211FD-CG) ASPEED AST2600 avec micrologiciel basé sur AMI MegaRAC prenant en charge IPMI 2.0 et API DMTF Redfish | 
| Environnement | Température de fonctionnement du système : 0°C ~ 35°C Température hors fonctionnement : -20°C ~ 70°C Humidité relative hors fonctionnement : 5 % ~ 85 % (sans condensation) | 
| Alimentation | (1+1) Bloc d'alimentation redondant 450 W 80+ Platinum | 
| Refroidissement | (1) Module de refroidissement liquide en boucle fermée pour processeur max. 170 W (7) 4028 ventilateurs système | 
| Certifications | CE, FCC (Classe A) | 
| Contenu de l'emballage | (1) Barebone 1U S1102-02 (1) Guide rapide (1) Kit de glissières (1) Module de refroidissement liquide en boucle fermée (assemblé à l'intérieur du châssis) (4) Câbles d'alimentation | 
| Liste des pieces | (1) Guide rapide : G52-S3362X2-Q13 (1) Kit de glissières : E21-S336020-C27 (1) Module de refroidissement liquide en boucle fermée : E34-F000091-AQ0 (4) Câbles d'alimentation : (2)K33-3001002-I45 pour l'Amérique du Nord, (2)K33-3001005-I45 pour l'UE | 
| Pièces en option | Module TPM2.0 : TPM20-IR Cordon d'alimentation C13 250 V/10 A (CN) : K33-3001029-I45 Cordon d'alimentation C13 125 V/13 A (JP) : K33-3001346-I45 | 
Performances de l'AMD EPYC 4005
Dans ce test, nous comparerons le processeur AMD EPYC 4564P à 16 cœurs avec 64 Mo de cache (génération 4004) au processeur AMD EPYC 4585PX à 16 cœurs avec 128 Mo de cache (génération 4005). Les tests ont été effectués sur le serveur MSI S1102-2, comme indiqué précédemment. Voici les composants du serveur de test MSI S1102-02 :
- AMD EPYC 4585PX ou AMD EPYC 4564P
- GPU NVIDIA L4
- 4 disques durs Seagate Exos M 30 To
- SSD de démarrage Solidigm P44 Pro M.2
- 4 x Samsung 16 Go DDR5-4800 ECC UDIMM
Y-Cruncher
y-cruncher est un programme multithread et évolutif capable de calculer Pi et d'autres constantes mathématiques jusqu'à des milliers de milliards de chiffres. Depuis son lancement en 2009, il est devenu une application de benchmarking et de test de résistance populaire auprès des overclockeurs et des passionnés de matériel informatique.
Lors des tests standard (1B à 10B), l'EPYC 4585PX est systématiquement plus rapide que l'EPYC 4564P, réduisant le temps de calcul de 4.5 à 10.2 %. Par exemple, lors du test à 1 milliard de chiffres, il affiche 18.802 s contre 20.933 s pour le 4564P.
Le test BBP montre des gains bien plus importants, jusqu'à 40 %, ce qui souligne la sensibilité de ces charges de travail à la latence mémoire et à la capacité L3. Le cas d'un milliard de BBP le montre clairement : 1 s (0.387 4585 PX) contre 0.630 s (4564 38.6 P), soit une réduction de XNUMX %. Des avantages similaires sont observés pour des tailles de BBP plus élevées.
Ce modèle s'aligne sur le cache L4585 plus grand du 3PX (128 Mo contre 64 Mo sur le 4564P) et sur les améliorations architecturales
| y-cruncher Temps de calcul total (Plus bas, c'est mieux) | AMD EPYC 4585PX 16 cœurs | AMD EPYC 4564P 16 cœurs | 
| 1 milliard | 18.802 secondes | 20.933 secondes | 
| 2.5 milliard | 54.085 secondes | 58.108 secondes | 
| 5 milliard | 121.711 secondes | 128.428 secondes | 
| 10 milliard | 269.096 secondes | 281.677 secondes | 
| 1 milliard de BBP | 0.387 secondes | 0.630 secondes | 
| 10 milliard de BBP | 4.450 secondes | 7.367 secondes | 
| 100 milliard de BBP | 50.311 secondes | 84.051 secondes | 
Mixeur 4.0
Blender 4.0 est une application de modélisation 3D open source. Ce benchmark a été réalisé à l'aide de l'utilitaire Blender Benchmark CLI. Le score est mesuré en échantillons par minute, les valeurs les plus élevées étant les meilleures.
Bien que les deux processeurs aient obtenu de bons résultats pour leur catégorie, le cache accru de l'EPYC 4585PX lui permet d'être encore plus rapide. Par exemple, lors du test Monster, le 4585PX a atteint 351.90 échantillons par minute, tandis que le 4564P a obtenu 306.57 échantillons par minute. Dans Junkshop, les scores étaient de 226.59 pour le 4585PX et de 195.26 pour le 4564P. De même, toujours dans Classroom, avec respectivement 171.87 et 150.27.
| Échantillons CPU par minute de Blender 4.0 (plus c'est élevé, mieux c'est) | AMD EPYC 4585PX 16 cœurs | AMD EPYC 4564P 16 cœurs | 
| Monster | 351.90 | 306.57 | 
| Brocanteur | 226.59 | 195.26 | 
| Salle de classe | 171.87 | 150.27 | 
Points de repère Phoronix
Nous avons utilisé la suite de tests Phoronix pour automatiser les installations, exécuter les charges de travail et collecter les résultats sur cinq plateformes clés : STREAM, 7-Zip, la compilation du noyau Linux, le serveur HTTP Apache et OpenSSL. Ci-dessous, l'EPYC 4585PX est comparé directement à l'EPYC 4564P de génération précédente (tous deux à 16 cœurs).
Bande passante mémoire STREAM : le 4585PX accuse un léger retard, avec 38,472 40,106.9 Mo/s contre 4.1 1,634.9 Mo/s (−4564 %, −XNUMX XNUMX Mo/s). STREAM est très sensible aux horloges/topologies mémoire et aux choix du compilateur ; cela indique une faible marge de bande passante pour le XNUMXP dans notre configuration.
7-Zip (compression + décompression) : le 4585PX délivre 162,951 176,484 MIPS contre 7.7 13,533 MIPS (−7 %, −4564 XNUMX MIPS). XNUMX-Zip s'appuie sur le débit entier et le comportement du cache/mémoire ; ici, le XNUMXP est légèrement en avance.
Compilation du noyau Linux (allmodconfig) : 4585PX s'exécute en 621.967 s contre 747.610 s (-125.643 s, soit 16.8 % de plus). Ce temps réduit reflète un meilleur débit de construction parallèle grâce aux améliorations architecturales.
Requêtes Apache/sec : 4585PX sert 181,764.45 132,744.77 R/s contre 36.9 49,019.68 R/s (+XNUMX %, +XNUMX XNUMX R/s), indiquant un avantage considérable dans le débit HTTP.
Vérification OpenSSL : le 4585PX atteint 400,939,420,057 224,487,686,510 78.6 176.45 vérifications/s contre 4585 3 XNUMX XNUMX vérifications/s (+XNUMX %, +XNUMX milliards de vérifications/s). Les charges de travail cryptographiques bénéficient des gains architecturaux du XNUMXPX et de son LXNUMX plus étendu.
Globalement : le 4585PX affiche des avances substantielles dans les tests de construction, Web et crypto (gains de 17 à 79 %), tandis que le 4564P se démarque sur la bande passante mémoire brute et 7-Zip dans notre configuration.
| Points de repère Phoronix | AMD EPYC 4585PX 16 cœurs | AMD EPYC 4564P 16 cœurs | 
| Discussions | 38,472.0 Mo / s | 40,106.9 Mo / s | 
| 7-ZIP | 162,951 XNUMX MIP/s | 176,484 XNUMX MIP/s | 
| Compilation du noyau (allmod) | 621.967 XNUMX secondes | 747.610 XNUMX secondes | 
| Apache (requêtes par seconde) | 181,764.45 XNUMX R/s | 132,744.77 XNUMX R/s | 
| OpenSSL | 400,939,420,057 XNUMX XNUMX XNUMX Vérifications | 224,487,686,510 XNUMX XNUMX XNUMX Vérifications | 
Performances FIO du disque dur
Le serveur MSI S1102-02 prend en charge quatre disques durs 3.5 pouces en façade, offrant ainsi plusieurs options de stockage. Bien que le stockage le plus rapide soit assuré par les emplacements PCIe 4.0 M.2 intégrés, il est possible d'installer des disques durs jusqu'à 30 To en façade. En utilisant des disques durs SATA Seagate Exos M de 30 To , nous avons mesuré les performances séquentielles et aléatoires avec FIO. Nous avons mesuré une bande passante séquentielle maximale de 1.3 Go/s en lecture et 1.1 Go/s en écriture, avec des IOPS aléatoires 4K de 1 720 IOPS en lecture et 1 990 IOPS en écriture. Ces résultats ont été obtenus en configuration JBOD ; le RAID influencera donc les performances finales.
| Charge de travail | Bande passante la plus élevée | IOPS les plus élevées | Latence la plus faible | 
|---|---|---|---|
| Écriture aléatoire (4K) | 7.8 Mo / s | 1,990 | 2.106 ms | 
| Lecture aléatoire (4K) | 6.7 Mo / s | 1,720 | 6.861 ms | 
| Écriture séquentielle (128 Ko) | 1,102.5 Mo / s | 8,820 | 0.486 ms | 
