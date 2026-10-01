---
id: collect-250926-servers-hardware/servers-hardware/fr-review-supermicro-hyper-superserver-sys-221h-tn24r-review-e7dd20c1-2
title: "fr-review-supermicro-hyper-superserver-sys-221h-tn24r-review-e7dd20c1"
domain: servers-hardware
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["benchmark", "gpu", "intel", "open source", "valuation"]
source: docs/RAG/clean4/fr-review-supermicro-hyper-superserver-sys-221h-tn24r-review-e7dd20c1.md
source_anchor: ""
source_lines: [84, 141]
sha256: e3378598d42aae8bf59c9ed1580c9bbe07b520b3337e14f2fb983c3c33cddf09
---

# fr-review-supermicro-hyper-superserver-sys-221h-tn24r-review-e7dd20c1

Pour explorer les capacités avancées du SYS-221H-TN24R, nous examinerons l'interface de gestion à distance. Cela offre une fenêtre sur l'état opérationnel du serveur, fournissant des fonctionnalités de surveillance et de gestion en temps réel qui sont cruciales pour garantir des performances optimales et des réponses de maintenance rapides.
Par exemple, la présentation inclut le contrôle UID (Unit Identification) et des informations générales sur le serveur. Il fournit le nom du fabricant, le nom du produit, le numéro de série, l'état de l'alimentation, le nom d'hôte et les informations réseau, y compris les adresses IP et MAC du BMC. Le contrôle UID est généralement utilisé pour identifier visuellement le serveur dans un rack avec un voyant clignotant. De plus, les versions du micrologiciel du BMC et du BIOS sont également affichées pour fournir des informations utiles en un coup d'œil, ce qui en fait un espace utile pour la gestion des stocks et le dépannage.
La section CPU répertorie l'état de santé, le modèle, la vitesse, les watts de puissance thermique de conception (TDP), le nombre de cœurs, le nombre de threads et le fabricant des processeurs installés. Dans ce cas, il présente deux processeurs Intel Xeon Platinum 8562Y+, chacun avec 32 cœurs et 64 threads, fonctionnant à 2800 300 MHz et un TDP de XNUMX watts.
La zone de consommation électrique fournit des données historiques et en temps réel sur la consommation électrique du serveur. Comme le montre l'image ci-dessous, il affiche la consommation d'énergie minimale, moyenne et maximale au cours de la dernière heure, et sous le graphique se trouve un résumé de la consommation d'énergie depuis la mise sous tension du serveur. Cela inclut les valeurs de pointe et leurs heures correspondantes. De plus, il fournit également aux utilisateurs des tendances historiques pour la dernière heure, le dernier jour et la dernière semaine, détaillant l'utilisation moyenne, la consommation maximale et la consommation électrique minimale.
Ces informations sont extrêmement utiles pour ceux qui cherchent à gérer l'efficacité énergétique de leur système ainsi qu'à garantir que le serveur fonctionne dans les limites de sa capacité électrique.
Performances du Supermicro Hyper SuperServer SYS-221H-TN24R
Nous continuons d'exploiter les résultats de performance obtenus lors de notre analyse initiale des processeurs Intel Xeon Scalable de 5e génération. Notre étude porte sur le système SYS-221H-TN24R, équipé de deux processeurs 8562Y+. Cette configuration sera comparée à notre plateforme Supermicro E1.S , à titre de référence, cette dernière étant équipée de deux processeurs 8460H.
Spécifications du Supermicro Hyper SuperServer SYS-221H-TN24R
- 2 processeurs Intel Xeon Platinum 8562Y+
- 512GB DDR5
- Windows Server 2022
Spécifications Supermicro Storage SuperServer SSG-121E-NES24R
- 2 processeurs Intel Xeon Platinum 8460H
- 512GB DDR5
- Windows Server 2022
Les Intel Xeon Platinum 8562Y et 8460H, appartenant respectivement aux 5e et 4e générations de processeurs évolutifs Intel Xeon, présentent des différences notables. Le 8562Y+ fonctionne à une fréquence de base de 2.80 GHz et offre 32 cœurs avec 64 threads, contrastant avec les 8460 cœurs et 40 threads du 64H à une fréquence de base de 2.2 GHz.
Le 8460H se démarque également par un cache nettement plus important de 105 Mo, comparé aux 60 Mo du 8562Y+, ce qui pourrait s'avérer avantageux dans les applications nécessitant un traitement de données poussé. Les deux modèles prennent en charge la mémoire DDR5. Cependant, le 8460H est moins économe en énergie, avec un TDP de 330W contre les 300W du 8562Y+. En termes de prix, le 8460H est nettement plus cher, à plus de 10,000 8562 $, tandis que le 6,000Y+ coûte un peu moins de XNUMX XNUMX $.
Pour fournir une analyse complète des performances, nos tests englobent une variété de tests intensifs, chacun ciblant différents aspects des performances du processeur. Nous utiliserons Blender OptiX pour les performances de rendu, Blackmagic pour évaluer les capacités de traitement vidéo et Geekbench pour l'évaluation globale des performances du système. De plus, Cinebench nous donnera un aperçu des graphiques et de l'efficacité du rendu, un Y-cruncher pour les calculs mathématiques et des tests de compression 7-Zip pour évaluer les vitesses de traitement et de compression des données.
Ces divers points de référence nous aideront à fournir une évaluation complète, couvrant à la fois les tâches informatiques générales et spécialisées, afin de voir comment elles se comportent dans des scénarios du monde réel.
Mixeur OptiX
Le premier est le test Blender, une application de modélisation 3D open source. Ce benchmark a été exécuté à l'aide de l'utilitaire Blender Benchmark. Le score est exprimé en échantillons par minute, le plus élevé étant le meilleur.
Dans la version 4.0 de Blender, voici les scores des deux CPU Intel :
| Processeur Blender 4.0 | 2 x Platine 32Y+ à 8562 cœurs | 2x Platine 40H à 8460 cœurs | 
| Monster | 805.254137 | 671.717058 | 
| Brocanteur | 513.800608 | 424.825295 | 
| Salle de classe | 414.999628 | 347.791937 | 
Test de vitesse Blackmagic RAW
Nous avons également commencé à exécuter le test de vitesse RAW de Blackmagic, qui teste les performances de lecture vidéo. De plus, ce test est davantage un test hybride regroupant à la fois le CPU et le GPU dans un scénario réel pour le décodage RAW.
| Test de vitesse Blackmagic RAW | 2 x Platine 32Y+ à 8562 cœurs | 2x Platine 40H à 8460 cœurs | 
| Processeur 8k | FPS 175 | FPS 145 | 
Geekbench 6
Geekbench 6 est un outil d'évaluation multiplateforme mesurant les performances globales d'un système. Toutefois, il serait intéressant d'analyser les performances monocœur et multicœur, ainsi que les résultats du benchmark OpenCL. Un score élevé indique de meilleures performances. Précisons que nous n'avons examiné que les résultats du processeur, aucun GPU n'étant installé sur les serveurs.
| Banc Geek 6 | 2 x Platine 32Y+ à 8562 cœurs | 2x Platine 40H à 8460 cœurs | 
| Single Core | 2,149 | 1,111 | 
| Multi-Core | 22,494 | 8,589 | 
Vous pouvez trouver des comparaisons avec n'importe quel système dans le navigateur Geekbench.
Cinebench R23
Cinebench R23 de Maxon est une référence de rendu de processeur qui utilise tous les cœurs et threads de processeur. Nous l'avons exécuté pour des tests multicœurs et monocœurs. Des scores plus élevés sont meilleurs.
| Cinebench R23 | 2 x Platine 32Y+ à 8562 cœurs | 2x Platine 40H à 8460 cœurs | 
| Processeur multicœur | 103,848 | 75,720 | 
| CP monocœur | 1,500 | 1,068 | 
| Rapport PM | 69.25x | 70.92x | 
Cinebench 2024
Voici les résultats pour la version 2024 de Cinebench, en regardant le CPU.
| Cinebench R24 | 2 x Platine 32Y+ à 8562 cœurs | 2x Platine 40H à 8460 cœurs | 
| Processeur multicœur | 5,289 | 4,243 | 
| CP monocœur | 82 | 62 | 
| Rapport PM | 62.1 | 68.52 | 
croque-y
y-cruncher est un programme multithread et évolutif qui peut calculer Pi et d'autres constantes mathématiques jusqu'à des milliards de chiffres. Depuis son lancement en 2009, y-cruncher est devenu une application d'analyse comparative et de test de stress populaire auprès des overclockeurs et des passionnés de matériel. Plus vite, c'est mieux dans ce test.
| croque-y | 2 x Platine 32Y+ à 8562 cœurs | 2x Platine 40H à 8460 cœurs | 
| 50 milliard | 363.758 secondes | N/D | 
| 25 milliard | 164.066 secondes | 191.909 secondes | 
| 10 milliard | 57.868 secondes | 67.484 secondes | 
| 5 milliard | 26.934 secondes | 30.813 secondes | 
| 2.5 milliard | 12.036 secondes | 14.059 secondes | 
| 1 milliard | 4.344 secondes | 5.134 secondes | 
Conclusion
