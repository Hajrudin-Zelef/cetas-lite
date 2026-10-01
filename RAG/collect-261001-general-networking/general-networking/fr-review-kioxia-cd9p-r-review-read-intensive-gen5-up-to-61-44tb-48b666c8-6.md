---
id: collect-261001-general-networking/general-networking/fr-review-kioxia-cd9p-r-review-read-intensive-gen5-up-to-61-44tb-48b666c8-6
title: "fr-review-kioxia-cd9p-r-review-read-intensive-gen5-up-to-61-44tb-48b666c8"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "gpu", "llama"]
source: docs/RAG/collect-261001-general-networking/fr-review-kioxia-cd9p-r-review-read-intensive-gen5-up-to-61-44tb-48b666c8.md
source_anchor: ""
source_lines: [145, 152]
sha256: b90b992b1722a3d77f3b3018e57c90e773ae38a486e7b31f98837960afef6488
---

# fr-review-kioxia-cd9p-r-review-read-intensive-gen5-up-to-61-44tb-48b666c8

Dans le segment 1 Mo, deux disques ont présenté des pics de latence élevés à 1 Mo/64 Mbps, cohérents avec la dégradation de leur débit dans cette configuration. Les Micron 9550 Max (12.8 To) et Micron 9550 Pro (7.68 To) ont chacun atteint environ 27 000 à 28 000 µs à 1 Mo/64 Mbps, puis leur latence a encore augmenté à 1 Mo/128 Mbps. Le CD9P-R a atteint environ 14 500 µs à 1 Mo/64 Mbps, reflétant son débit d'écriture plus stable dans cette configuration, et a terminé le test à environ 29 100 µs à 1 Mo/128 Mbps. Le Pascari X200P (7.68 To) a affiché la latence d'écriture la plus faible à 1 Mo/128 Mbps, à environ 25 000 µs. Les valeurs les plus élevées à 1M/128 ont été observées pour le Micron 9550 Pro à environ 44 900 µs et le Solidigm PS1010 à environ 40 700 µs, le Micron 7600 Max suivant à environ 36 000 µs.
Conclusion
Le KIOXIA CD9P-R E3.S 7.68 To remplit parfaitement son rôle de disque dédié aux lectures intensives, et les données de test le confirment du début à la fin. Les lectures séquentielles se sont classées parmi les meilleures du groupe avec 14 235,9 Mo/s lors de notre test 128 Ko, égalant ainsi le Pascari X200P. La véritable force du disque s'est révélée à faible profondeur de file d'attente : une latence de lecture aléatoire 4 Ko d'environ 30 µs à QD1, soit près de la moitié des 60 à 90 µs enregistrés par la plupart des concurrents. Cette faible latence en lecture s'est également confirmée lors des tests GPU Direct Storage, où le CD9P-R a montré une excellente progression avec le nombre de threads et a atteint le débit de lecture 1 Mo le plus élevé du groupe, soit environ 6.2 Gio/s.
Les compromis sont tout aussi évidents et s'inscrivent dans la conception 1 DWPD du disque, loin de la contredire. Les performances en écriture se situent en milieu de classement, avec un débit d'écriture séquentielle de 6 912,4 Mo/s pour les requêtes 128K, le plaçant en dernière position. Le débit d'écriture séquentielle GDSIO atteint un pic de 4.9 Gio/s, le classant cinquième parmi les huit disques. La gamme Micron 9550 domine ces charges de travail d'écriture, tant en termes de débit que de latence. Ces résultats ne constituent en aucun cas un reproche au CD9P-R ; il s'agit d'un SSD conçu pour la lecture intensive, et les utilisateurs ayant besoin de performances d'écriture soutenues devraient plutôt se tourner vers le CD9P-V, plus polyvalent.
La constance sous charge est l'autre aspect important. Lors d'un test de pointage DLIO avec le profil LLAMA 3.1 405B, le CD9P-R s'est stabilisé autour de 572 secondes par passe, restant dans la plage de 553 à 590 secondes qui caractérise les solutions Gen5 courantes. Il a ainsi évité les fluctuations qui ont poussé le Pascari X200P jusqu'à 674.5 secondes. La plateforme offre également une évolutivité bien supérieure à celle de notre échantillon : la famille CD9P-R atteint 30.72 To en E3.S et 61.44 To au format 2.5 pouces sur la même architecture, ce qui permet à une seule qualification de couvrir l'ensemble des besoins, des nœuds de calcul haute performance aux niveaux de lecture haute densité.
Pour les parcs de stockage cloud, les pipelines de données IA, la distribution de contenu et les serveurs de lecture virtualisés avec des modèles d'accès dominés par la lecture, le CD9P-R est un disque que nous recommandons sans hésiter. Il combine un débit de lecture Gen5 de pointe avec la latence de lecture basse résolution la plus faible de notre comparatif et une capacité de stockage importante, tout en conservant une stabilité optimale même sous charge soutenue.
Page produit – KIOXIA CD9P-R 7.68 To
Classement : Le Kioxia CD9P-R occupe la première place du classement des meilleurs SSD d'entreprise pour les applications de lecture intensive.
