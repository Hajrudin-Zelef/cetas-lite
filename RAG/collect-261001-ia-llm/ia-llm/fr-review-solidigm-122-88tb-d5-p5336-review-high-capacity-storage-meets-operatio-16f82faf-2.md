---
id: collect-261001-ia-llm/ia-llm/fr-review-solidigm-122-88tb-d5-p5336-review-high-capacity-storage-meets-operatio-16f82faf-2
title: "fr-review-solidigm-122-88tb-d5-p5336-review-high-capacity-storage-meets-operatio-16f82faf"
domain: ia-llm
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["diffusion", "distribution", "intel", "nand"]
source: docs/RAG/collect-261001-ia-llm/fr-review-solidigm-122-88tb-d5-p5336-review-high-capacity-storage-meets-operatio-16f82faf.md
source_anchor: ""
source_lines: [21, 62]
sha256: 683b13c0138f0f62c14220c7661071ebb9916d7ec201bc9ce0dbdebaed39b15e
---

# fr-review-solidigm-122-88tb-d5-p5336-review-high-capacity-storage-meets-operatio-16f82faf

| Lecture séquentielle | 7000MB / s | 
| Écriture séquentielle | 3000MB / s | 
| Lecture aléatoire (IOPS) | 900,000 4 (256K, QDXNUMX) | 
| Écriture aléatoire (IOPS) | 19,000 16 (256K, QDXNUMX) | 
| Latence (lecture/écriture) | Lecture : 110 μs (4 Ko) / Écriture : 40 μs (32 Ko) | 
| Latence séquentielle (typ.) | Lecture : 8 μs (4 Ko) / Écriture : 21 μs (32 Ko) | 
| Alimentation (actif/inactif) | Actif : 24 W / Inactif : 5 W | 
| Endurance | 0.6 DWPD (32 134.3 RW) / XNUMX PBW | 
| MTBF | 2 millions d'heures | 
| UBER | <1 secteur par 10 bits lus | 
| Température de fonctionnement | 0 ° C à 70 ° C | 
| Vibration / choc | 2.17 GRMS (en fonctionnement), 1,000 XNUMX G (choc) | 
| Garantie | 5 ans | 
| Poids | 166.4 g ± 10 g | 
Test de performance
Plateforme de test de conduite
Nous utilisons un serveur Dell PowerEdge R760 exécutant Ubuntu 22.04.02 LTS comme plateforme de test pour toutes les charges de travail présentées dans cet article. Équipé d'un boîtier JBOF Serial Cables Gen5 , il offre une large compatibilité avec les SSD U.2, E1.S, E3.S et M.2. La configuration de notre système est décrite ci-dessous :
- 2 x Intel Xeon Gold 6430 (32 cœurs, 2.1 GHz)
- 16 x 64GB DDR5-4400
- Disque SSD Dell BOSS de 480 Go
- Câbles série Gen5 JBOF
Comparaison des lecteurs
- Solidigm P5336 122.88 To (Gen4 | 2.5″ | U.2)
- Solidigm P5336 61.44 To (Gen4 | 2.5″ | U.2)
- Micron 6550 ION 61.44 To (Gen5 | E3.S)
Comme indiqué en introduction, le marché des disques durs d'entreprise haute capacité est complexe, avec différents formats, types de NAND et rapports qualité-prix à prendre en compte. Pour cette analyse, nous avons sélectionné un petit groupe de SSD à comparer au Solidigm P122.88 de 5336 To, dont le Solidigm P61.44 de 5336 To et le Micron 61.44 de 6550 To.
Le Micron 6550 est unique car il est basé sur les technologies Gen5 et TLC, et il est l'un des rares à être produit à ce niveau de capacité. Le lecteur Micron bénéficiera de vitesses d'E/S plus élevées.
Lors de l'analyse des performances, il est essentiel de comprendre cette structuration. Lors du déploiement, ces disques ne sont peut-être pas en concurrence directe, mais leurs capacités se chevauchent. Pour fournir une référence en termes d'évolutivité, nous avons inclus le disque Micron dans cette analyse.
Performances du CDN
Afin de simuler une charge de travail CDN réaliste à contenu mixte, les SSD ont été soumis à une séquence d'analyse comparative en plusieurs phases conçue pour reproduire les schémas d'E/S des serveurs Edge à contenu important. La procédure de test couvre différentes tailles de blocs, grandes et petites, réparties sur des opérations aléatoires et séquentielles, avec différents niveaux de concurrence.
Avant les principaux tests de performance, chaque SSD a effectué un remplissage complet du périphérique avec une passe d'écriture séquentielle à 100 % utilisant des blocs de 1 Mo. Ce processus utilisait des E/S synchrones et une profondeur de file d'attente de quatre, permettant quatre tâches simultanées. Cette phase garantit que le disque entre dans un état stable représentatif d'une utilisation réelle. Après le remplissage séquentiel, une deuxième phase de saturation d'écriture aléatoire de trois heures a été exécutée selon une distribution bssplit pondérée (taille de bloc/pourcentage), privilégiant fortement les transferts de 128 Ko (98.51 %), avec des contributions mineures des blocs inférieurs à 128 Ko jusqu'à 8 Ko. Cette étape émule les schémas d'écriture fragmentés et irréguliers souvent observés dans les environnements de cache distribué.
La suite de tests principale s'est concentrée sur des opérations de lecture et d'écriture aléatoires à grande échelle afin de mesurer le comportement du lecteur sous des profondeurs de file d'attente et des tâches simultanées variables. Chaque test a duré cinq minutes (300 secondes) et a été suivi d'une période d'inactivité de trois minutes, permettant aux mécanismes de récupération internes de stabiliser les indicateurs de performance.
- Exécuté selon une distribution de taille de bloc fixe privilégiant 128 Ko (98.51 %), les 1.49 % restants étant composés de transferts de plus petite taille, compris entre 64 Ko et 8 Ko. Chaque configuration variait entre 1, 2 et 4 tâches simultanées, avec des profondeurs de file d'attente de 1, 2, 4, 8, 16 et 32, afin de profiler l'évolutivité du débit et la latence dans des conditions d'écriture en périphérie classiques.
- Un profil de taille de bloc fortement hétérogène, imitant la récupération de contenu CDN, a été utilisé, commençant par une composante dominante de 128 Ko (83.21 %) suivie d'une longue traîne de plus de 30 blocs plus petits, allant de 4 Ko à 124 Ko, chacun avec une représentation fréquentielle fractionnaire. Cette distribution reflète les divers schémas de requêtes rencontrés lors de la récupération de segments vidéo, de l'accès aux vignettes et de la recherche de métadonnées. Ces tests ont également été effectués sur l'ensemble de la matrice des nombres de tâches et de la profondeur des files d'attente.
Cette combinaison de tests de préconditionnement, de saturation et d'accès aléatoire de taille mixte est conçue pour révéler comment les SSD gèrent les environnements de type CDN soutenus, en mettant l'accent sur la réactivité et l'efficacité dans les scénarios à forte bande passante et hautement parallélisés.
Charge de travail CDN Lecture 1
Lors de ce test de lecture monothread simulant un trafic de diffusion de contenu léger, les modèles Solidigm P5336 122.88 To et Solidigm P5336 61.44 To présentent des caractéristiques d'évolutivité constantes. Le modèle 122.88 To atteint 7,109 32 Mo/s à QD61.44, légèrement supérieur aux 7,002 6550 Mo/s du modèle 61.44 To. Cette évolutivité quasi identique suggère que le modèle Solidigm de plus grande capacité conserve la même efficacité sous une faible pression de lecture sans dégradation des performances. En revanche, le Micron 12,288 XNUMX To affiche une évolutivité beaucoup plus dynamique, atteignant XNUMX XNUMX Mo/s.
Charge de travail CDN Lecture 2
Avec deux threads appliqués, les Solidigm P5336 122.88 To et P5336 61.44 To offrent des performances quasiment identiques, passant de 840 Mo/s au premier jour à environ 1 7,467 Mo/s et 7,469 32 Mo/s respectivement au troisième jour. Les deux disques affichent des gains constants jusqu'au seizième jour, après quoi le débit stagne, indiquant un point de saturation dans leur architecture actuelle. Pour les applications avec un parallélisme modéré, cela constitue une base fiable pour une évolutivité prévisible. Le Micron 16, en revanche, affiche une plage d'évolutivité globale plus élevée, commençant à 6550 1,384 Mo/s et se poursuivant jusqu'à 13,312 32 Mo/s au troisième jour, reflétant les avantages de sa mémoire NAND TLC et de son interface Gen5.
Charge de travail CDN Lecture 4
Ce scénario de lecture à forte demande sollicite davantage les disques dont la simultanéité est accrue. Les disques Solidigm P5336 122.88 To et P5336 61.44 To affichent une évolutivité constante, atteignant environ 7,466 7,469 à 16 32 Mo/s à QD6550 et conservant une bande passante stable jusqu'à QD13,107. Les résultats entre les deux capacités restent pratiquement identiques, ce qui confirme la cohérence du comportement du contrôleur Solidigm sur toute sa gamme haute capacité. En comparaison, le Micron 16 a atteint XNUMX XNUMX Mo/s à QDXNUMX et a maintenu cette bande passante jusqu'à la fin du test.
Écriture de la charge de travail CDN 1
