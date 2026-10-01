---
id: collect-261001-general-networking/general-networking/fr-review-after-the-pi-record-serving-a-130-tb-dataset-with-backblaze-b2-c0672672-2
title: "fr-review-after-the-pi-record-serving-a-130-tb-dataset-with-backblaze-b2-c0672672"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Google"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-general-networking/fr-review-after-the-pi-record-serving-a-130-tb-dataset-with-backblaze-b2-c0672672.md
source_anchor: ""
source_lines: [17, 35]
sha256: ca266532510a729422525eaff1eb707160da516075a66d1215a9a4aa8d0ac2bd
---

# fr-review-after-the-pi-record-serving-a-130-tb-dataset-with-backblaze-b2-c0672672

Le tableau de bord réseau UniFi du laboratoire confirme les caractéristiques de chargement durant la période de transfert. L'interface WAN indique que Backblaze est l'application dominante en termes de volume de trafic, avec un débit montant de 2.27 Gbit/s et une consommation mensuelle de données WAN de 90.9 To enregistrée sur la période. La connexion est restée stable tout au long du transfert, sans perte de paquets significative.
Débit de transfert par minute
Le graphique ci-dessous illustre le débit de transfert mesuré par minute sur une période représentative du chargement. Les barres atteignent régulièrement entre 15 et 16 gigaoctets par minute, ce qui correspond à un débit de ligne soutenu d'environ 2 Gbit/s. Les brèves interruptions visibles sur le graphique correspondent à des pauses de validation de somme de contrôle effectuées périodiquement pendant le transfert afin de garantir l'intégrité des données avant de poursuivre.
Progrès cumulatif des transferts
Le graphique cumulatif des octets transférés illustre la progression linéaire et constante des données, passant de 0 à environ 100 To sur cette période. Cette stabilité, observée sur l'ensemble de la période, témoigne de la fluidité du transfert, sans interruption majeure ni baisse de débit.
Disposition finale des seaux
Backblaze a créé le bucket pi-314-trillion, qui contient l'ensemble des 628 fichiers et dont la taille confirmée est de 132 210,5 Go. Ce bucket est configuré comme privé, conserve toutes les versions des fichiers et est accessible via le point de terminaison compatible S3 à l'adresse s3.us-west-004.backblazeb2.com. Le stockage objet simplifie la gestion d'un ensemble de données de cette ampleur. Chaque fichier est adressable individuellement, la liste complète peut être récupérée par programmation et il n'y a aucune hiérarchie de système de fichiers ni limite de volume à contourner.
Accéder à l'ensemble de données
Des chercheurs nous demandent d'accéder à ces données ; d'ailleurs, un projet est déjà en cours. Michael Kleber est ingénieur logiciel principal chez Google, mais il étudie les décimales de pi depuis l'obtention de son doctorat en mathématiques en 1999. Les mathématiciens considèrent pi comme un nombre normal , il est donc raisonnable de se demander : « Parmi les 10^d séquences de d chiffres, laquelle met le plus de temps à apparaître dans pi, et combien de chiffres faut-il ? » Kleber a mené la recherche jusqu'à d = 7, et lorsque Fabrice Bellard a calculé 2.7 billions de décimales de pi en 2009, Kleber l'a encouragé à étendre ses recherches jusqu'à d = 11, la limite du possible à l'époque. « Avec 314 billions de chiffres aléatoires, il y a environ 79 % de chances de trouver toutes les séquences de longueur 13 », explique Kleber, « alors j'espère que nous aurons de la chance ! » Nous nous attendons à ce que cela se produise beaucoup plus souvent maintenant que Backblaze a rendu les données accessibles à tous.
L'ensemble de données PI est hébergé sur Backblaze B2 et peut être téléchargé par toute personne souhaitant l'utiliser. L'accès se fait via un lien de demande permettant d'obtenir des identifiants ou les instructions de téléchargement pour récupérer les fichiers. Backblaze hébergera l'ensemble de données afin d'en garantir la disponibilité pour la recherche et la vérification pendant toute la durée de l'hébergement.
Options de téléchargement
Les utilisateurs peuvent récupérer des fichiers individuels ou l'intégralité des 130 To de données, selon leurs besoins. Le bucket est structuré de manière à ce que chaque objet puisse être accédé et téléchargé directement, sans avoir à récupérer l'ensemble des données. Pour ceux qui souhaitent tout récupérer, une synchronisation complète du bucket peut être effectuée à l'aide des outils décrits ci-dessous. Cette option nécessite 135 To d'espace libre.
Outils recommandés
- Rclone est l'outil recommandé pour accéder à l'ensemble de données. Il s'intègre parfaitement à Backblaze B2 et permet aux utilisateurs d'adapter le processus de téléchargement à leurs besoins en bande passante.
- API compatible S3Ainsi, tout outil de téléchargement compatible avec S3 peut être utilisé pour récupérer les données. La seule condition est que cet outil permette de modifier l'URL du point de terminaison S3 par défaut afin de pointer vers le point de terminaison B2 plutôt que vers AWS.
Un exemple concret d'infrastructure hybride
Notre calcul de Pi à 314 billions de décimales illustre parfaitement ce à quoi ressemble une infrastructure hybride en pratique. Le calcul a été entièrement exécuté sur un seul serveur Dell PowerEdge R7725 du laboratoire StorageReview. Cependant, le système devant être réaffecté à d'autres projets et tâches une fois le calcul terminé, conserver en permanence 130 To de résultats au sein du laboratoire n'était pas une solution viable.
Le désir de mettre les données à la disposition de ceux qui souhaitent les utiliser, que ce soit pour des travaux scientifiques rigoureux ou par simple curiosité, est bien réel. Cependant, héberger un ensemble de données de cette taille au sein du laboratoire représente rapidement une charge importante pour les opérations. La bande passante est saturée, l'infrastructure est constamment sollicitée et les tâches quotidiennes du laboratoire sont perturbées par chaque requête de téléchargement.
Une solution comme Backblaze B2 élimine tous ces obstacles. Les données résident dans une infrastructure cloud conçue spécifiquement pour les volumes de données de cette envergure, avec un débit évolutif, de multiples points de redondance pour garantir leur intégrité, et la sécurité ainsi que l'expertise opérationnelle propres aux plateformes de stockage d'entreprise. Les ressources de calcul étaient hébergées sur site car le projet l'exigeait. Le stockage est dans le cloud car c'est tout simplement la solution la plus adaptée pour la suite.
