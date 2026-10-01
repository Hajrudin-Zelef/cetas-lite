---
id: collect-261001-general-networking/general-networking/test-de-l-iphone-18-pro-le-meilleur-iphone-est-aussi-le-plus-prudent-frandroid-3
title: "test-de-l-iphone-18-pro-le-meilleur-iphone-est-aussi-le-plus-prudent-frandroid"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "DeepSeek", "Google", "Xiaomi"]
dates: []
keywords: ["agent", "deepseek", "gemini", "gpu", "transcription"]
source: docs/RAG/collect-261001-general-networking/test-de-l-iphone-18-pro-le-meilleur-iphone-est-aussi-le-plus-prudent-frandroid.md
source_anchor: ""
source_lines: [103, 142]
sha256: 34796ce65cd9d62f75c8a2b5e106ecebb10bcd9c341ba7aea2c4ccdaed4b1e67
---

# test-de-l-iphone-18-pro-le-meilleur-iphone-est-aussi-le-plus-prudent-frandroid

Face Ã Android, la comparaison devient une histoire d’Ã©cosystÃ¨me avant d’Ãªtre une histoire de silicium. Dans la mÃªme application PocketPal, le Xiaomi 17 Ultra et son Snapdragon 8 Elite Gen 5 plafonnent Ã 15 tokens par seconde, le Pixel 11 Pro Fold Ã 2,5 : sur Android, l’app tourne sur CPU, et activer le NPU Hexagon du Xiaomi divise la vitesse par trois faute de format compatible. Pour une mesure honnÃªte, nous sommes passÃ©s par Google AI Edge Gallery, l’app de Google, avec le modÃ¨le de Google, Gemma 4 E4B, sur GPU partout : le 18 Pro dÃ©code entre 25 et 34 tokens par seconde selon les runs, l’Air et le Xiaomi tournent Ã 19, le Pixel Ã 7,5. Le Tensor G6 est excellent avec Gemini Nano dans les applications Google, et moyen dÃ¨s qu’on sort de l’Ã©cosystÃ¨me Google, l’Air et le Xiaomi font jeu Ã©gal, le 18 Pro est seul devant.

Et sur iPhone, une application tierce obtient l’accÃ©lÃ©ration matÃ©rielle sans rien configurer. Sur Android, le mÃªme modÃ¨le peut aller six fois plus ou moins vite selon l’application, le rÃ©glage (CPU, GPU, etc.) et le smartphone.

TroisiÃ¨me niveau, les fonctions natives, et lÃ le discours se retourne. Nous avons pris la mÃªme photo, prise avec l’Air, et lui avons fait subir les mÃªmes retouches sur les deux iPhone : effacer un panneau immobilier sur un balcon avec Nettoyer, enlever la plaque d’Ã©gout, redresser la perspective avec la nouvelle RÃ©orientation, agrandir le cadre.

Le rÃ©sultat est identique au pixel prÃ¨s, ce qui est logique puisque c’est le mÃªme modÃ¨le Apple Intelligence des deux cÃ´tÃ©s, et le 18 Pro va un peu plus vite, de l’ordre de quelques secondes sur une rÃ©orientation, rien qu’on remarque sans chronomÃ¨tre.

Ãa s’explique assez simplement, en fait : Apple calibre ses fonctions pour l’iPhone le moins puissant qui y a droit. Au passage, la RÃ©orientation laisse des coins vides, que le smartphone remplit en inventant. Sur notre photo, le bÃ¢timent de gauche a gagnÃ© une fenÃªtre qui n’existe pas. Elle est bien alignÃ©e, bien Ã©clairÃ©e, et totalement fausse.

On a continuÃ© les tests, et on ne vous les livre pas tous ici. Mais on se demandait Ã quoi ressemble l’intelligence d’un modÃ¨le de 4 milliards de paramÃ¨tres, celui qui tient dans un smartphone ? Nous avons soumis Ã Gemma 4 E4B, sur les quatre appareils, un petit problÃ¨me de physique Ã partir de nos propres mesures de recharge, six valeurs Ã calculer (le test de charge est plus bas). Les quatre smartphones ont donnÃ© des rÃ©ponses diffÃ©rentes, avec le mÃªme modÃ¨le et le mÃªme prompt, parce que l’arithmÃ©tique des GPU n’accumule pas dans le mÃªme ordre.

Aucun n’a dÃ©passÃ© 4 sur 6, tous ont rÃ©ussi le rendement, trois sur quatre ont ratÃ© une soustraction. Et quand, par erreur, nous avons laissÃ© le corrigÃ© dans le prompt, le Pixel a calculÃ© 18,25 W, a vu que la rÃ©ponse attendue Ã©tait 18,2, et Ã©crit qu’il fallait Â« s’en tenir Ã cette valeur Â».

Comme l’ensemble des autres, la puce du 18 Pro est plus rapide, mais elle n’est pas plus intelligente. C’est mÃªme le fil rouge de cette section : plus on s’approche des usages qu’Apple a prÃ©vus, moins la puce compte, plus on s’en Ã©loigne, plus elle Ã©crase tout. Le Neural Engine 32 cÅurs est un investissement pour les applications de demain, et pour l’instant, on est encore loin de l’exploiter Ã fond.

Bref, avec sa puce A20 Pro, l’iPhone 18 Pro reste le smartphone le plus douÃ© pour l’IA embarquÃ©e, alors mÃªme que Siri AI n’est pas disponible en France. C’est sans doute lÃ que le bÃ¢t blesse.

### Ã quoi sert toute cette puissance, alors ?

Il faut poser la question franchement, parce que je me la pose : qui, dans la vraie vie, va faire tourner un modÃ¨le de langage dans une application tierce sur son smartphone ? Presque personne.

Nos tokens par seconde sont la seule faÃ§on de mesurer le Neural Engine, mais au quotidien, il travaille surtout ailleurs. Ce qui s’en sert, c’est tout le reste, et tout le reste est invisible. Chaque photo passe par des rÃ©seaux de neurones avant mÃªme d’apparaÃ®tre Ã l’Ã©cran, pour la fusion des expositions, le dÃ©bruitage, le dÃ©tourage, la balance des blancs. Le clavier en utilise pour la correction, Photos pour la recherche, Dictaphone pour la transcription, l’Appareil photo pour reconnaÃ®tre un texte ou un objet dans le viseur, Face ID pour vous ouvrir le tÃ©lÃ©phone, et Siri pour Ã peu prÃ¨s tout ce qu’il fait encore chez nous.

Ces fonctions tournent en local, sans rÃ©seau, des centaines de fois par jour, et c’est prÃ©cisÃ©ment pour elles qu’Apple a doublÃ© son Neural Engine. Le problÃ¨me, c’est qu’aucune ne se laisse mesurer. On ne chronomÃ¨tre pas le dÃ©bruitage d’une photo, on ne compare pas la correction du clavier entre deux puces, on ne sait pas si un iPhone 18 Pro trouve Â« la maison bleue Â» dans Photos plus vite qu’un Air parce que la recherche est faite avant qu’on tape. Le gain existe probablement, rÃ©parti en millisecondes sur mille gestes, et il est impossible Ã dÃ©montrer. C’est la limite de ce test, et de tous les tests : on mesure la puissance lÃ oÃ¹ elle se voit, et elle sert surtout lÃ oÃ¹ elle ne se voit pas. Apple le sait, et c’est sans doute pour Ã§a que sa communication parle de Â« 2x Â» plutÃ´t que de vous montrer oÃ¹.

Et pour demain, mÃªme si ce n’est pas l’objet de ce test, je pense Ã un systÃ¨me d’exploitation qui ne lance plus des apps mais les fabrique, un agent local qui comprend le besoin et construit l’outil, sans rien envoyer Ã personne. Et Ã§a n’a rien de science-fiction, puisque c’est ce que les modÃ¨les frontiÃ¨re font dÃ©jÃ dans le cloud.

iPhone 18 Pro Ã 99â¬ + forfait 250Go Ã 29,99â¬/mois pendant 12 mois. Profitez de 50â¬ de remise immÃ©diate avec le code NEW50 et jusqu’Ã 150â¬ de bonus reprise pour votre ancien mobile. Livraison offerte.

Ce qui bloque, c’est de les faire tenir dans une poche, et lÃ , franchement, c’est la physique qui dÃ©cide, quelle que soit l’ingÃ©niositÃ© qu’on y met. Un modÃ¨le de classe frontiÃ¨re compte des centaines de milliards de paramÃ¨tres ; mÃªme quantifiÃ©, DeepSeek en version 671 milliards demande un Mac Studio Ã 512 Go de mÃ©moire unifiÃ©e et plus de 800 Go/s de bande passante pour tourner.

Un iPhone 18 Pro a 12 Go et une fraction de cette bande passante, et l’on a vu ce que cette bande passante fait aux tokens par seconde. Le contexte est l’autre mur : un agent de codage qui charge un dÃ©pÃ´t entier dÃ©passe largement le cache mÃ©moire qu’une carte de 16 Go peut contenir, alors pour une puce de smartphone, n’en parlons pas. Reste alors ce qu’on a testÃ©, les modÃ¨les de 4 milliards de paramÃ¨tres, et ils sont bluffants pour leur taille, mais leur taille est le sujet.

Ma propre expÃ©rience est plus modeste et plus parlante : quatre calculs de physique de niveau lycÃ©e, et Gemma 4 E4B en rate deux sur six (l’arithmÃ©tique de tÃªte est une faiblesse de tous les modÃ¨les de langage, mais bon…).

La bonne nouvelle, c’est que la frontiÃ¨re descend : un modÃ¨le de 3 Ã 14 milliards fait aujourd’hui ce qu’un 70 milliards faisait il y a douze Ã dix-huit mois. Le rÃªve viendra donc, par le bas, et l’iPhone qui aura la mÃ©moire et la bande passante pour l’accueillir n’est pas encore sorti.

## iOS : du bon et de l’absent

