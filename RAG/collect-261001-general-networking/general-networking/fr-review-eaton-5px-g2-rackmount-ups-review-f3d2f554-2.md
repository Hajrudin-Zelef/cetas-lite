---
id: collect-261001-general-networking/general-networking/fr-review-eaton-5px-g2-rackmount-ups-review-f3d2f554-2
title: "fr-review-eaton-5px-g2-rackmount-ups-review-f3d2f554"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "attention"]
source: docs/RAG/collect-261001-general-networking/fr-review-eaton-5px-g2-rackmount-ups-review-f3d2f554.md
source_anchor: ""
source_lines: [55, 76]
sha256: 51ba38fcf4c6e607ef4e1d1c1b7de434e91e455b2a36cc3e1ecdeb0ce5e31233
---

# fr-review-eaton-5px-g2-rackmount-ups-review-f3d2f554

Comme nous l'avons indiqué ci-dessus, l'une des améliorations les plus utiles du nouvel onduleur d'Eaton est l'ajout d'une barre d'état LED lumineuse sous l'écran LCD. En bref, vous n'aurez aucun problème à voir la lumière clignotante lorsque vous passerez devant l'onduleur sur votre rack de serveur et cela nous permet de comprendre facilement ce qui se passe avec l'onduleur en un coup d'œil :
- Lumière bleue : Tout va bien et normal
- Lumière jaune : une attention est nécessaire, mais rien de critique pour le moment
- Feu rouge : un problème critique nécessite une attention et une action immédiates
Le déplacement de l'onduleur vers le panneau arrière révèle une gamme de connectivité et de fonctionnalités. À l'extrême gauche se trouve le connecteur d'alimentation, qui est un cordon d'alimentation de 10 pieds avec 5-15P.
Ensuite, il y a la carte réseau M2 en option, un ancien port USB de type B, le port de mise hors tension à distance (RPO) et de marche/arrêt à distance (ROO), le port de relais de sortie, le port de communication série RS-232, un module de batterie externe (EBM ) port de détection. À droite se trouvent les points de vente gérés :
- Quatre prises gérées 5-15R (situées à gauche, noir : groupe principal)
- Deux prises gérées 5-15R (situées en haut à droite, grises : groupe de segment de charge 1)
- Deux prises gérées 5-15R (situées en bas à droite, gris : groupe de segment de charge 2)
Comme nous l'avons déjà mentionné, Eaton affirme que le nouvel onduleur 5PX G2 est nettement plus silencieux en fonctionnement. Notre modèle de test est équipé d'un seul ventilateur, dont le fonctionnement diffère légèrement des modes marche/arrêt. Par exemple, dans une pièce calme, notre ancien 5PX disposait d'une vitesse élevée et d'une vitesse réduite selon qu'il fonctionnait sur batterie ou avec une charge importante. Le nouveau modèle G2, quant à lui, semble privilégier le mode haute vitesse au démarrage, puis réduire progressivement sa vitesse en plusieurs étapes. Le réglage le plus bas du ventilateur que nous avons testé jusqu'à présent est même plus silencieux que notre onduleur de bureau BeQuiet . Même avec le ventilateur en marche, la différence avec l'onduleur 5PX de première génération reste considérable.
Gestion de l'onduleur Eaton 5PX G2
Les nouvelles solutions d'onduleur Eaton 5PX G2 sont gérées par l'Intelligent Power Manager (IPM) de la société. Il s'agit d'un logiciel complet qui surveille et gère les dispositifs d'alimentation (Eaton et tiers) dans un environnement physique et virtuel via une interface modernisée et facile à utiliser. Il propose également une gamme de niveaux de tarification continus et d'abonnement, permettant aux organisations de trouver le forfait qui leur convient le mieux.
À partir du tableau de bord Accueil, vous pouvez voir le diagramme de flux d'énergie, l'état de la prise, les alarmes actives et l'environnement actuel (c'est-à-dire les températures et l'humidité).
Dans la zone Meters, vous pouvez afficher la tension et la fréquence du 5PX, des détails sur sa sortie, ainsi que l'état et la santé de la batterie.
Sous Contrôles, vous pouvez activer un redémarrage ou un arrêt en toute sécurité de l'ensemble de l'onduleur lui-même, ou par le groupe de prises (dans ce cas, 1 ou 2). Cela vous donne une certaine flexibilité sur les appareils que vous devez éteindre immédiatement (au lieu de tous les appareils connectés).
Les paramètres de protection vous permettent de définir les critères d'arrêt de l'alimentation. Choisissez la stratégie d'alimentation et les conditions à remplir pour que cette action soit mise en œuvre.
La section Environnement affiche le nom de l'appareil, la température, l'humidité, les contacts secs (c'est-à-dire lorsque l'alimentation/la tension est fournie par une autre source) et la communication.
Il existe une gamme d'options de configuration et d'autres informations utiles dans le menu Paramètres. Par exemple, dans la section Réseau et protocole, vous pouvez voir toutes les informations détaillées sur le réseau du 5PX.
La section maintenance vous permet de mettre à jour le firmware.
Conclusion
L'onduleur Eaton 5PX G2 est un autre excellent produit du fabricant de solutions de gestion de l'alimentation, qui s'intègre parfaitement aux autres unités Eaton de notre baie . Les utilisateurs disposant d'un écosystème riche en appareils énergivores, tels que des commutateurs PoE, des serveurs de périphérie et des équipements informatiques complets, constateront que la nouvelle génération de la gamme 5PX répond largement à leurs besoins, leur assurant une protection optimale de leurs appareils. Le 5PX G2 utilise la technologie sinusoïdale pure, offre une protection contre les surtensions améliorée par rapport à la génération précédente et est conforme à la norme de performance des onduleurs IEC 62040-2.
Il a une empreinte physique globale plus petite, un meilleur écran LCD et la possibilité d'ajouter plus de batteries par rapport au modèle précédent. L'Eaton 5PX G2 est également nettement plus silencieux et était souvent à peine perceptible lorsqu'il entrait automatiquement dans les modes de vitesse de ventilation inférieure. De plus, nous avons particulièrement apprécié l'ajout du nouveau voyant d'état LED tricolore sur le panneau avant, qui attire facilement notre attention lorsqu'il se passe quelque chose avec l'onduleur Eaton. Dans l'ensemble, l'Eaton 5PX G2 est un excellent ajout à leur portefeuille déjà impressionnant de solutions d'alimentation.
