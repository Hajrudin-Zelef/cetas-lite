---
id: collect-260926-rattrapage/rattrapage/fr-review-supermicro-storage-a-asg-1115s-ne316r-e3-s-server-review-323ab036-2
title: "fr-review-supermicro-storage-a-asg-1115s-ne316r-e3-s-server-review-323ab036"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Samsung"]
dates: []
keywords: ["benchmark", "gpu", "valuation"]
source: docs/RAG/lot-rattrapage/servers-reviews/fr-review-supermicro-storage-a-asg-1115s-ne316r-e3-s-server-review-323ab036.md
source_anchor: ""
source_lines: [64, 126]
sha256: 3b929c75ef253a9fba0fd71696b75da5ffc0ea6181842063777808ab844783cf
---

# fr-review-supermicro-storage-a-asg-1115s-ne316r-e3-s-server-review-323ab036

Le Blackmagic Disk Speed Test est un autre test pour lequel nous disposons uniquement de résultats pour le Supermicro. Ce test exécute un exemple de fichier de 5 Go pour les vitesses de lecture et d'écriture. Avec ce test, nous avons constaté des vitesses de lecture supérieures à 3 Go/s et près de 1.5 Go/s sur le SSD Samsung PM1A9 M.3 de 2 To du Supermicro.
| Test de vitesse du disque Blackmagic (plus c'est élevé, mieux c'est) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 Go DDR5) | 
|---|---|
| Écrire | 1,415.1 Mo / s | 
| Lire | 3,031.0 Mo / s | 
Cinebench R23
Cinebench R23 de Maxon est une référence de rendu de processeur qui utilise tous les cœurs et threads de processeur. Nous l'avons exécuté pour des tests multicœurs et monocœurs. Des scores plus élevés sont meilleurs. Voici les résultats pour toutes les puces EPYC. Alors que l'on constate un écart attendu entre eux pour la partie multicœur, la partie monocœur ne varie que de huit points entre 84c et 96c. Sur le 128c, il était en retard sur les deux autres de 203 points au 96c et 211pts derrière le 84c.
| Cinebench R23 | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 Go DDR5) | TYAN SX TS70AB8056 (96C EPYC 9684X, 512 Go DDR5) | TYAN SX TS70AB8056 (128C EPYC 9754, 512 Go DDR5) | 
|---|---|---|---|
| Processeur (multicœur) (points) | 81,148 | 93,720 | 103,876 | 
| Processeur (monocœur) (points) | 1,309 | 1,301 | 1,098 | 
| Rapport PM | 61.99x | 72.04X | 94.65x | 
Cinebench 2024
Cinebench 2024 de Maxon est une référence de rendu CPU et GPU qui utilise tous les cœurs et threads du processeur. Nous l'avons exécuté pour des tests multicœurs et monocœurs. Comme cette configuration n’a pas de GPU, nous n’avons pas ces chiffres. Des scores plus élevés sont meilleurs. Encore une fois, nous n'avons que des résultats pour l'ASG-1115S-NE316R dans ce test.
| Cinebench 2024 | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 Go DDR5) | 
|---|---|
| Processeur (multicœur) (points) | 4,913 | 
| Processeur (monocœur) (points) | 81 | 
| Rapport PM | 60.47x | 
Évaluation du processeur Geekbench
Geekbench 6 est un benchmark multiplateforme mesurant les performances globales du système. Toutefois, une analyse comparative des performances monocœur et multicœur, ainsi que du benchmark OpenCL, serait intéressante. Un score élevé indique de meilleures performances. Nous n'avons examiné que les résultats du processeur, ce serveur ne disposant pas de carte graphique. À l'instar de nos observations avec Cinebench 2023, les performances monocœur de l'EPYC 9634 (Supermicro) et de l'EPYC 9684X étaient très proches, avec un écart plus marqué par rapport à l'EPYC 9754. Les performances multicœur étaient également très proches entre l'EPYC 9634 et l'EPYC 9684X, avec respectivement 22 868 et 21 329 points, mais là encore, l'écart avec le 9754 (18 683 points) restait significatif.
Vous pouvez trouver des comparaisons avec n'importe quel système dans le navigateur Geekbench.
| Geekbench 6 | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 Go DDR5) | TYAN SX TS70AB8056 (96C EPYC 9684X, 512 Go DDR5) | TYAN SX TS70AB8056 (128C EPYC 9754, 512 Go DDR5) | 
|---|---|---|---|
| CPU Benchmark - Monocœur | 2,055 | 2,093 | 1,738 | 
| Référence CPU - Multi-Core | 22,868 | 21,329 | 18,683 | 
croque-y
y-cruncher est un programme multithread et évolutif qui peut calculer Pi et d'autres constantes mathématiques jusqu'à des milliards de chiffres. Depuis son lancement en 2009, elle est devenue une application d'analyse comparative et de test de résistance populaire auprès des overclockeurs et des passionnés de matériel. En termes de résultats, nous n'avons que des chiffres de 1 milliard et 10 milliards pour le 9684X et le 9754, mais sur le 9634 EPYC, nous avons des résultats de 1 milliard à 50 milliards. Étonnamment, le Supermicro est arrivé en tête avec le 9634 et seulement 7.274 secondes pour 1 milliard et 71.336 secondes pour 10 milliards.
| y-cruncher (Temps de calcul total) (plus le niveau est bas, mieux c'est) | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 Go DDR5) | TYAN SX TS70AB8056 (96C EPYC 9684X, 512 Go DDR5) | TYAN SX TS70AB8056 (128C EPYC 9754, 512 Go DDR5) | 
|---|---|---|---|
| 1 milliard de chiffres (secondes) | 7.274 secondes | 10.296 secondes | 9.568 secondes | 
| 2.5 milliard de chiffres (secondes) | 17.055 secondes | N/D | N/D | 
| 5 milliard de chiffres (secondes) | 34.336 secondes | N/D | N/D | 
| 10 milliard de chiffres (secondes) | 71.336 secondes | 72.377 secondes | 80.171 | 
| 25 milliard de chiffres (secondes) | 196.695 secondes | N/D | N/D | 
| 50 milliard de chiffres (secondes) | 439.435 secondes | N/D | N/D | 
Compression à 7 zips
L'utilitaire populaire 7-Zip dispose d'un test de mémoire intégré qui démontre les performances du processeur. Dans ce test, nous l'exécutons avec une taille de dictionnaire de 128 Mo lorsque cela est possible. Encore une fois, nous n'avons que les résultats de l'ASG-1115S-NE316R pour ce test.
|  | Supermicro ASG-1115S-NE316R (84C EPYC 9634, 384 Go DDR5) | 
|---|---|
| Compression |  | 
| Utilisation actuelle du processeur | 3,574 % | 
| Courant nominal/utilisation | 5.755 GIPS | 
| Courant | 205.670 GIPS | 
| Utilisation résultante du processeur | 3,572 % | 
| Évaluation/utilisation résultante | 5.733 GIPS | 
| Note résultante | 204.758 GIPS | 
| Décompression |  | 
| Utilisation actuelle du processeur | 3,810 % | 
| Courant nominal/utilisation | 5.753 GIPS | 
| Courant | 219.191 GIPS | 
| Utilisation résultante du processeur | 3,803 % | 
| Évaluation/utilisation résultante | 5.779 GIPS | 
| Note résultante | 219.768 GIPS | 
| Note totale |  | 
| Utilisation totale du processeur | 3,687 % | 
| Note totale/utilisation | 5.756 GIPS | 
| Note totale | 219.768 GIPS | 
Conclusion
Le Supermicro ASG-1115S-NE316R est une plate-forme EPYC dense qui offre beaucoup de potentiel dans 1U. Du côté des performances du processeur, nous avons constaté des chiffres intéressants entre les puces à 84, 96 et 128 cœurs, le 128 à 9754 cœurs étant moins performant dans certaines charges de travail et le 84 à 9634 cœurs tirant vers l'avant. Aujourd'hui plus que jamais, il est essentiel de comprendre les charges de travail et de coupler les processeurs de manière appropriée.
En ce qui concerne les serveurs de stockage, le Supermicro Storage A+ ASG-1115S-NE316R se distingue comme une entrée moderne, conçue de manière efficace et bien conçue pour gérer tout ce que nécessite l'application logicielle installée dessus. Les baies E3.S Gen5 se démarquent également, offrant six baies de plus que d'habitude dans un serveur 1U avec baies U.2. D'une importance cruciale, Supermicro n'a pas non plus lésiné sur les E/S arrière, offrant aux utilisateurs quatre options à exploiter sur ce front.
Dans cette revue, nous nous sommes concentrés sur les performances du système en termes de stockage. Dans un article ultérieur, nous approfondirons l’histoire des performances de stockage lorsqu’il est rempli de flash Gen5 hautes performances.
Page produit Supermicro ASG-1115S-NE316R.
