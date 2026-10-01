---
id: collect-250926-servers-hardware/servers-hardware/fr-review-vdi-acceleration-for-all-intel-data-center-gpu-flex-series-170-review-4b06c1ee-2
title: "fr-review-vdi-acceleration-for-all-intel-data-center-gpu-flex-series-170-review-4b06c1ee"
domain: servers-hardware
role: reference
task: reference
actors: ["Apple", "Google", "Intel", "Samsung"]
dates: []
keywords: ["gpu", "intel", "benchmarks"]
source: docs/RAG/clean4/fr-review-vdi-acceleration-for-all-intel-data-center-gpu-flex-series-170-review-4b06c1ee.md
source_anchor: ""
source_lines: [53, 109]
sha256: ffc5a9de3dd41cb6c0a4fe6043a71016660d852d25e61bce2d59214fa43051a3
---

# fr-review-vdi-acceleration-for-all-intel-data-center-gpu-flex-series-170-review-4b06c1ee

Cette approche rentable de l'administration des vGPU améliore l'évolutivité des déploiements de serveurs VDI. Les organisations peuvent provisionner et ajuster dynamiquement les ressources vGPU dans les environnements virtuels sans se soucier des coûts de licence supplémentaires. Cette flexibilité est cruciale pour s'adapter à l'évolution des demandes de charge de travail et prendre en charge un nombre croissant d'utilisateurs, en particulier dans les environnements à haute densité typiques des segments des travailleurs du savoir.
D'un point de vue commercial, les GPU Intel Flex Series offrent des avantages financiers substantiels. En supprimant les complexités et les coûts liés aux licences, Intel permet aux organisations de rationaliser leur infrastructure VDI. Cette simplification accélère les temps de déploiement et réduit le besoin d'allocations budgétaires importantes pour les capacités GPU, faisant de la série Intel Flex une option intéressante pour les entreprises cherchant à optimiser leurs investissements VDI tout en maintenant des performances et une fiabilité élevées.
Pratique : Intel Flex Series 170 avec Supermicro SuperBlade
Dans notre laboratoire, nous avons testé la carte graphique Intel Flex Series 170 avec VMware sur notre système Supermicro SuperBlade X13 GPU Blade. L'installation des cartes graphiques dédiées Intel sur VMware ESXi a été simplifiée pour optimiser l'expérience utilisateur et les performances du système. Il a suffi de télécharger le pilote via SCP sur l'hôte à l'aide d'un fichier ZIP, d'activer l'accès SSH et de lancer l'installation. Après un redémarrage rapide de l'hôte, la Flex Series 170 est apparue dans la liste du matériel et les options SR-IOV (de 0 à 31) étaient disponibles pour la configuration.
Les tests ont été réalisés à l'aide du système Supermicro SuperBlade , indispensable pour valider les performances du GPU Flex Series 170 en tant qu'accélérateur VDI. Le système SuperBlade est conçu pour optimiser la densité de calcul et l'efficacité tout en minimisant la consommation d'énergie. Il constitue ainsi une plateforme idéale pour tester les cas d'utilisation à haute densité destinés aux travailleurs du savoir, tels que ceux proposés par les GPU Intel Flex Series.
| Partie | Configuration SuperMicro SuperBlade | 
|---|---|
| Processeur | 1xIntel Xeon 8562Y+ | 
| Mémoire | 256GB DDR5 | 
| par chaîne | 2x lecteur M.3840 Samsung 2G | 
| GPU | Intel Flex série 170 | 
Le système Supermicro SuperBlade X13 constitue une plateforme idéale pour optimiser la densité de déploiement. Grâce à la possibilité d'intégrer 10 nœuds dans un seul châssis 8U, il est envisageable d'héberger jusqu'à 320 VDI accélérés dans un châssis compact et facile à gérer. À l'inverse, en optant pour davantage de cœurs et de VRAM par carte Flex Series 170, on peut déployer 80 à 160 VDI pour les utilisateurs avertis. De plus, la possibilité d'intégrer jusqu'à 10 GPU dans ce système offre une grande flexibilité, et son réseau interne ultra-rapide permet de mettre en œuvre des solutions de basculement innovantes. Consultez notre test complet pour en savoir plus sur ce serveur SuperBlade, véritable couteau suisse du système.
Performances du GPU Intel Data Center Flex Series 170
Il est important de considérer ici que les tests ont été sélectionnés en fonction de ce qui produirait officiellement un résultat complet, une limitation des tests dans des environnements virtuels sans ajustements ni hacks particuliers. Nous avons sélectionné quelques tranches SR-IOV puis quelques tranches non standard pour voir à quoi cela ressemblerait.
Rendu 3D
Nous nous sommes tournés vers les benchmarks 3D Mark Wildlife pour montrer les capacités de jeu en nuage et de rendu 3D du GPU Flex Series.
3DMark Wild Life propose un outil d'analyse comparative multiplateforme compatible avec les systèmes Windows, Android et Apple iOS. Cet outil évalue et compare les performances graphiques de divers appareils, notamment les ordinateurs portables, les tablettes et les smartphones. Wild Life utilise l'API graphique Vulkan pour les appareils Windows et Android, alors qu'il utilise Metal pour les appareils iOS. Étant donné que ce test fonctionne sur des graphiques intégrés selon différents scores, il peut illustrer la puissance d'un score graphique pur de l'Intel Flex Series 170.
| Test/Tranche SR-IOV | 2GB | 4GB | 7GB | 14GB | 
|---|---|---|---|---|
| Marque 3D de la faune | 29,062 | 42,466 | 49,671 | 45,908 | 
| Marque 3D de la faune extrême | 9,023 | 14,948 | 17,661 | 16,959 | 
LuxMark
Vient ensuite LuxMark, un utilitaire d’analyse comparative des GPU OpenCL. Le Flex Series 170 s'est vraiment montré flexible lors de ce test, affichant des chiffres et une mise à l'échelle impressionnants.
| Test/Tranche SR-IOV | 2GB | 4GB | 7GB | 14GB | 
|---|---|---|---|---|
| Salle Luxmark | 2,961 | 4,382 | 11,002 | 11,202 | 
| Nourriture Luxmark | N/D | 1,316 | 4,502 | 4,525 | 
PCMark 10 Express
Vous vous tournez vers PCMark 10, équipé d'une vaste suite de tests qui reflètent avec précision le large éventail de tâches rencontrées sur le lieu de travail. Cet outil d'analyse comparative comprend diverses évaluations de performances, des options de tests personnalisés, un profil d'autonomie de la batterie et de nouveaux tests de stockage, ce qui en fait une solution globale pour évaluer les performances des PC de bureau modernes.
| Test/Tranche SR-IOV | 2GB | 4GB | 7GB | 14GB | 
|---|---|---|---|---|
| PCMark 10 Express dans son ensemble | 5,111 | 5,146 | 5,311 | 5,218 | 
| Essentials JAMAIS | 10,269 | 10,318 | 10,734 | 10,364 | 
| Score de démarrage des applications | 17,833 | 17,664 | 19,034 | 17,340 | 
| Score de vidéoconférence | 7,798 | 7,933 | 8,095 | 7,980 | 
| Score de navigation Web | 7,789 | 7,839 | 8,028 | 8,046 | 
| Productivité | 6,952 | 7,004 | 7,181 | 7,180 | 
| Score des feuilles de calcul | 6,924 | 6,953 | 7,186 | 7,184 | 
| Écriture de la partition | 6,981 | 7,057 | 7,178 | 7,177 | 
En regardant nos résultats ici, même s’ils ne sont pas aussi spectaculaires, nous pouvons constater une nette mise à l’échelle de diverses tâches qui tirent parti de l’accélération.
Pour faciliter la comparaison, j'ai compilé les résultats de tous les tests dans un seul tableau.
| Test/Tranche SR-IOV | 2GB | 4GB | 7GB | 14GB | 
|---|---|---|---|---|
| Marque 3D Vie Sauvage | 29,062 | 42,466 | 49,671 | 45,908 | 
| Marque 3D Wild Life Extreme | 9,023 | 14,948 | 17,661 | 16,959 | 
| Salle Luxmark | 2,961 | 4,382 | 11,002 | 11,202 | 
| Nourriture Luxmark | N/D | 1,316 | 4,502 | 4,525 | 
| PCMark 10 Express dans son ensemble | 5,111 | 5,146 | 5,311 | 5,218 | 
| Essentials JAMAIS | 10,269 | 10,318 | 10,734 | 10,364 | 
| Score de démarrage des applications | 17,833 | 17,664 | 19,034 | 17,340 | 
| Score de vidéoconférence | 7,798 | 7,933 | 8,095 | 7,980 | 
| Score de navigation Web | 7,789 | 7,839 | 8,028 | 8,046 | 
| Productivité | 6,952 | 7,004 | 7,181 | 7,180 | 
| Score des feuilles de calcul | 6,924 | 6,953 | 7,186 | 7,184 | 
| Écriture de la partition | 6,981 | 7,057 | 7,178 | 7,177 | 
Réflexions de clôture
En analysant toutes ces données et en tenant compte du facteur humain, j'étais constamment stupéfait, voire impressionné, par la simplicité et la puissance de ces cartes. Après avoir alloué 1/32e du GPU à la machine virtuelle et installé le pilote Windows d'Intel , des fonctionnalités comme le Bureau à distance intégré fonctionnaient mieux. L'utilisation de Google Earth dans Chrome, une fonctionnalité phare d'Intel, a démontré que même avec 512 Mo de VRAM et un seul cœur Xe, l'expérience VDI est étonnamment meilleure. Ayant utilisé d'autres produits VDI par le passé, une fois le pilote installé, l'expérience était tout simplement agréable.
