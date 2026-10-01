---
id: collect-250926-servers-hardware/servers-hardware/fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887-3
title: "fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "benchmark", "benchmarks", "gpu", "open source", "valuation"]
source: docs/RAG/clean4/fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887.md
source_anchor: ""
source_lines: [105, 139]
sha256: 5fb82f6ec3910315763863986adf37356f74996b2f0864e0d50430fea29ab92c
---

# fr-review-supermicro-a-server-as-2126hs-tn-review-dual-amd-epyc-9005-turin-7a015887

Le routeur AS-2126HS-TN est livré avec le BMC intégré de Supermicro, une plateforme que nous avons déjà présentée pour plusieurs systèmes A+, et qui demeure un élément essentiel de la gestion quotidienne. Le tableau de bord offre une vue consolidée de l'état et de la configuration du système, permettant une visibilité immédiate sur l'état du firmware, l'inventaire matériel et les données de télémétrie des capteurs. Depuis la page d'accueil, les administrateurs peuvent rapidement vérifier les versions du BMC, du BIOS, du CPLD et de Redfish, confirmer la configuration réseau et valider l'identité de l'hôte avant de procéder à des opérations de dépannage ou de maintenance plus approfondies.
Dans la section « Informations sur les composants », le BMC affiche avec précision les processeurs AMD EPYC installés, notamment les deux processeurs 192 cœurs présents dans le système. Chaque socket est équipé d'un AMD EPYC 9965, doté de 192 cœurs actifs et de 384 threads par processeur, avec un TDP configuré de 500 W. Cette vue permet de vérifier rapidement la configuration des processeurs et les autres caractéristiques des composants directement depuis le BMC, garantissant ainsi leur fonctionnement optimal.
Dans l'onglet Refroidissement des informations sur les composants, le BMC permet de contrôler quatre modes de ventilation : Vitesse standard, Vitesse maximale, Vitesse optimale et Vitesse pour activités E/S intensives. Ces préréglages permettent d'ajuster rapidement le flux d'air en fonction de la charge de travail et des besoins thermiques, sans modifier le BIOS. Sous la sélection du mode, l'interface affiche l'état de chaque ventilateur et sa vitesse de rotation en temps réel, offrant ainsi une vue claire de l'état du refroidissement du système.
Pour l'accès à distance, le BMC de l'AS-2126HS-TN intègre une console KVM accessible via une interface HTML5 ou un plug-in Java. Les paramètres de la console permettent de sélectionner le mode souris de base pour s'adapter aux différents systèmes d'exploitation, et une option de réinitialisation IKVM est disponible en cas de blocage de la session. Ceci offre un accès complet à la console hors bande pour l'installation, le dépannage et la restauration, sans nécessiter de périphériques locaux.
Tests de performance du Supermicro AS-2126HS-TN
Pour les tests, la carte mère Supermicro AS-2126HS-TN n'a pas été configurée avec une architecture de stockage dense ou une alimentation et un câblage optimisés pour les GPU. Par conséquent, notre évaluation s'est concentrée exclusivement sur les performances du processeur, en utilisant les benchmarks Blender, y-cruncher et Phoronix pour caractériser le débit de calcul brut sur la plateforme AMD EPYC biprocesseur.
Le serveur AS-2126HS-TN était équipé de deux processeurs AMD EPYC 9965 à 192 cœurs, offrant un total de 384 cœurs et 768 threads. À titre de comparaison, les résultats sont obtenus avec ceux d'un serveur Dell PowerEdge R7725 configuré avec les mêmes deux processeurs AMD EPYC 9965 et 1.5 To de mémoire système. Les deux systèmes ont été testés avec les mêmes suites de tests et paramètres afin de garantir la cohérence des résultats.
Supermicro Configuration du serveur A+ AS -2126HS-TN
- CPU: 2x AMD EPYC 9965 (192 cœurs)
- RAM: 1.5 To 24 x 64 Go DDR5 6000 MHz
- SSD SSD NVMe Micron 7.68 To pour centres de données
Mixeur 4.5
Blender est une application de modélisation 3D open source. Ce benchmark a été réalisé avec l'utilitaire Blender Benchmark. Le score est mesuré en échantillons par minute, les valeurs les plus élevées indiquant de meilleures performances.
Dans le benchmark CPU SMT de Blender, le serveur Supermicro A+ AS-2126HS-TN, équipé de deux processeurs AMD EPYC 9965, affiche d'excellents résultats, illustrant les gains de performance d'une architecture à 192 cœurs pour les charges de travail de rendu multithreadées. Avec l'utilitaire Blender Benchmark, le système a enregistré 3 070,84 échantillons par minute pour Monster, 2 063,61 pour Junkshop et 1 527,39 pour Classroom.
À titre de comparaison, le Dell PowerEdge R7725, équipé des mêmes deux processeurs EPYC 9965, affiche un débit légèrement supérieur dans ce scénario avec SMT activé, atteignant 3 193,11 échantillons par minute dans Monster, 2 174,63 dans Junkshop et 1 608,79 dans Classroom. L’écart relativement faible entre les deux plateformes suggère un comportement globalement similaire sous des charges de rendu intensives en SMT, les différences étant probablement dues à l’optimisation au niveau de la plateforme plutôt qu’à la puissance de calcul brute.
| SMT du processeur Blender (Échantillons par minute ; plus c'est élevé, mieux c'est) | Serveur Supermicro A+ AS -2126HS-TN (AMD EPYC 9965 192C) | Dell PowerEdge R7725 (double processeur AMD EPYC 9965 192C) | 
|---|---|---|
| Monster | 3,070.84 | 3,193.11 | 
| Brocanteur | 2,063.61 | 2,174.63 | 
| Salle de classe | 1,527.39 | 1,608.79 | 
Avec la technologie SMT désactivée, le système Supermicro affiche une augmentation notable du débit brut dans les trois scènes. Les performances atteignent 4 018,10 échantillons par minute dans Monster, 2 707,10 dans Junkshop et 1 990,51 dans Classroom, démontrant ainsi comment le moteur de rendu CPU de Blender peut tirer parti d'une réduction des conflits entre les threads sur les systèmes à très grand nombre de cœurs.
Dans cette configuration, le Dell PowerEdge R7725 affiche toujours un débit absolu supérieur, avec respectivement 4 304,21, 2 870,34 et 2 079,79 échantillons par minute dans les environnements Monster, Junkshop et Classroom. Bien que le système Dell conserve un avantage en termes de performances maximales, le Supermicro AS-2126HS-TN présente des gains substantiels avec le SMT désactivé, ce qui souligne l'influence significative de la stratégie de multithreading sur les performances des plateformes de cette envergure.
| Processeur Blender sans SMT (Échantillons par minute ; plus c'est élevé, mieux c'est) | Serveur Supermicro A+ AS -2126HS-TN (AMD EPYC 9965192C) | Dell PowerEdge R7725 (double processeur AMD EPYC 9965 192C) | 
|---|---|---|
| Monster | 4,018.10 | 4,304.21 | 
| Brocanteur | 2,707.10 | 2,870.34 | 
| Salle de classe | 1,990.51 | 2,079.79 | 
croque-y
y-cruncher est un programme multithread et évolutif capable de calculer Pi et d'autres constantes mathématiques jusqu'à des milliers de milliards de chiffres. Depuis son lancement en 2009, il est devenu une application de benchmarking et de test de résistance populaire auprès des overclockeurs et des passionnés de matériel informatique.
Dans le benchmark y-cruncher, le Supermicro AS-2126HS-TN démontre une mise à l'échelle constante et prévisible à mesure que la taille des problèmes augmente. Le système effectue le calcul d'un milliard de décimales en 8.092 secondes, puis passe à 15.245 secondes pour 2.5 milliards de décimales et à 24.961 secondes pour 5 milliards de décimales.
À mesure que la charge de travail augmente, les temps de calcul évoluent de manière linéaire, atteignant 44.350 secondes pour 10 milliards de chiffres, 114.107 secondes pour 25 milliards et 246.688 secondes pour 50 milliards. Lors du calcul le plus exigeant, portant sur 100 milliards de chiffres, le système effectue l'opération en 572.800 secondes, démontrant ainsi sa capacité à gérer un nombre élevé de threads sur de longues périodes d'exécution.
Le Dell PowerEdge R7725 exécute les mêmes charges de travail légèrement plus rapidement sur toutes les tailles de test, terminant l'exécution de 100 milliards de chiffres en 481.207 secondes.
| Y-Cruncher (temps de calcul total) | Serveur Supermicro A+ AS -2126HS-TN (AMD EPYC 9965 192C) | Dell PowerEdge R7725 (double processeur AMD EPYC 9965 192C) | 
|---|---|---|
| 1 milliard | 8.092 s | 7.879 secondes | 
