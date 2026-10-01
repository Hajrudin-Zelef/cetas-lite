---
id: collect-261001-general-networking/general-networking/fr-review-synology-diskstation-ds2422-review-dcf1c77f-2
title: "fr-review-synology-diskstation-ds2422-review-dcf1c77f"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/fr-review-synology-diskstation-ds2422-review-dcf1c77f.md
source_anchor: ""
source_lines: [63, 90]
sha256: bafc0257e71ab3d22ed1023de701cb5f4e681f1016b37b9d4928eca31648bbc2
---

# fr-review-synology-diskstation-ds2422-review-dcf1c77f

- Écart-type de latence (écart-type de lecture + écriture moyenné ensemble)
Notre analyse de charge de travail synthétique d'entreprise comprend quatre profils basés sur des tâches réelles. Ces profils ont été développés pour faciliter la comparaison avec nos références passées ainsi qu'avec des valeurs largement publiées telles que la vitesse de lecture et d'écriture maximale de 4k et 8k 70/30, qui est couramment utilisée pour les disques d'entreprise.
4K
- 100 % de lecture ou 100 % d'écriture
- 100% 4K
8K70/30
- 70 % de lecture, 30 % d'écriture
- 100% 8K
8K (séquentiel)
- 100 % de lecture ou 100 % d'écriture
- 100% 8K
128K (séquentiel)
- 100 % de lecture ou 100 % d'écriture
Tout d'abord, les charges de travail d'entreprise, où nous mesurons un long échantillon de performances 4K aléatoires avec une activité d'écriture à 100 % et de lecture à 100 %. Le DS2422+ est passé à 81,238 6 en SSD RAID3622 iSCSI. En comparaison, le DS244,115xs+ est passé à 2422 3622 IOPS en lecture dans la même configuration. Comme vous pouvez le voir ci-dessous, le DSXNUMX+ était également loin derrière le DSXNUMXxs+ lorsqu'il a été testé avec des disques durs (bien que prévu).
La prochaine étape est notre référence de latence moyenne de 4K. Ici, le DS2422+ a enregistré des latences plus élevées que le DS3622xs+ dans toutes les configurations sauf la configuration RAID480 SMB de 6 Go. Avec les disques durs Synology de 16 To, les latences de lecture/écriture du DS2422+ étaient en moyenne de 783 ms/109 ms en SMB contre 720 ms/31 ms contre pour le DS3622xs+ ainsi que 58 ms/107 ms (DS2422+) en iSCSI contre 40 ms/31 ms (DS3622xs+).
En latence maximale de 4K, le DS2422+ a en fait montré de meilleures latences de lecture maximales globales par rapport au DS3622xs+ ; cependant, ses latences d'écriture maximales étaient généralement plus élevées.
Lors de notre dernier test 4K, écart type, le DS2422+ s'est plutôt bien comporté et a en fait obtenu un meilleur résultat de lecture que le DS3622xs+ en utilisant une configuration SSD RAID6 SMB, où il n'affichait que 17.01 ms.
Dans notre test de débit séquentiel 100 % 8K, le DS2422+ a atteint 167,090 113,032 IOPS en lecture et 3622 168,305 IOPS en écriture dans notre configuration SSD iSCSI, bien qu'il soit toujours nettement derrière le DS206,522xs+ (qui a atteint XNUMX XNUMX IOPS en écriture et XNUMX XNUMX IOPS en lecture).
Dans notre prochaine série de tests, nous examinons les performances de la charge de travail mixte 70 % en lecture/30 % en écriture, ce qui fait évoluer les charges de travail de deux threads, deux files d'attente à 16 threads, 16 files d'attente. Dans les configurations SSD, le DS3622xs+ et le DS2422+ sont à peu près à égalité jusqu'à environ la file d'attente 4 threads/8. Passant cette barre, le DS3622xs+ a commencé à bien mieux fonctionner, atteignant plus de 100,000 16 IOPS autour de 8 threads/2422 files d'attente. Les performances du DS8+ ont ralenti autour de la barre des 16 threads/XNUMX files d'attente.
Vient ensuite la latence moyenne 8K 70/30. Comme dans le test de latence moyenne 4K, le DS2422+ a pris un peu de retard sur le DS3622xs+ dans toutes les configurations. Les meilleures configurations de disque dur et de SSD pour le DS2422+ étaient toutes deux en iSCSI, qui affichaient une plage de 7.52 ms à 104.17 ms et de 0.39 ms à 4.34 ms, respectivement.
Avec une latence maximale de 8K 70/30, la configuration HDD et SSD supérieure était à nouveau en iSCSI, affichant une plage de 450 ms à 2,367 41.67 ms et de 305.15 ms à 3622 ms, respectivement. Sans surprise, le DSXNUMXxs+ avait des chiffres globaux inférieurs et des performances plus constantes.
Le dernier de nos tests de charge de travail mixte est l'écart type 8K 70/30, où les résultats étaient beaucoup plus proches (ce qui est visuellement évident dans nos graphiques ci-dessous). Dans l'ensemble, le DS2422+ avait une latence plus lente que le DS3622xs+, mais il fonctionnait toujours bien. La meilleure configuration de disque dur DS2422+ était en iSCSI, avec une latence allant de 450.5 ms à 2,367.3 34.04 ms, tandis que la meilleure configuration SSD était en SMB, avec une latence allant de 281.11 ms à XNUMX ms.
La dernière référence synthétique est notre test séquentiel à grand bloc 128K qui montre la vitesse de transfert séquentielle la plus élevée pour un appareil. Alors que les nombres de lecture étaient, sans surprise, très similaires entre le DS2422+ et le DS3622xs+, les écritures étaient beaucoup plus rapides avec le DS3622xs+. En utilisant des SSD, le DS2422+ a atteint 2.28 Go/s en lecture et 1.61 Go/s en écriture en SMB et 1.82 Go/s en lecture et 1.06 Go/s en écriture en iSCSI. Pour la configuration du disque dur, il affiche 2.28 Go/s en lecture et 1.34 Go/s en écriture en SMB et 1.86 Go/s en lecture et 746 Mo/s en écriture en iSCSI.
Conclusion
Dans l'ensemble, le DiskStation DS2422+ est un solide NAS à 12 baies pour les petites et moyennes entreprises. Alimenté par le système d'exploitation DiskStation Manager (DSM) complet et facile à utiliser de Synology, le DS2422+ offre aux utilisateurs une interface utilisateur simplifiée avec une gamme de support de gestion, d'outils multimédias, de solutions de sauvegarde et d'autres fonctionnalités. Il peut également évoluer jusqu'à 24 disques via une seule unité d'extension DX12 à 1222 baies.
Il y a cependant quelques inconvénients. Si les utilisateurs souhaitent tirer parti de la mise en cache SSD M.2 via la carte d'extension PCIe de Synology (en raison de l'absence d'un emplacement M.2), cela limitera le NAS à un seul port 10GbE. De plus, nous avons trouvé l'insistance de Synology à utiliser leurs disques de stockage de marque (qui sont beaucoup plus chers que d'autres disques NAS comme Seagate ou WD) un peu ennuyeux. Il convient de souligner que vous pouvez toujours utiliser vos propres disques avec le DS2422+ malgré leur messagerie.
Côté performances, aucune surprise notable. Nous nous attendions à ce que le DS2422+ soit en retrait par rapport au DS3622xs+ dans la plupart des catégories (en raison de son processeur Xeon 6 cœurs plus performant), et il offre effectivement un gain de performances significatif par rapport au DS2419+ , avec des vitesses de lecture de 2.28 Go/s et d'écriture de 1.61 Go/s en SMB, et de 1.82 Go/s et d'écriture de 1.06 Go/s en iSCSI pour la configuration SSD lors de notre test synthétique séquentiel de gros blocs de 128 Ko. La configuration HDD a quant à elle affiché des vitesses de lecture de 2.28 Go/s et d'écriture de 1.34 Go/s en SMB, et de 1.86 Go/s et d'écriture de 746 Mo/s en iSCSI lors du même test.
Si vous êtes une petite ou moyenne entreprise qui cherche à dépenser 800 $ de moins (par rapport au DS3622xs+) pour une solution NAS haute capacité aux performances solides pour vos sauvegardes ou vos projets multimédias, le DS2422+ peut être un bon choix. Synology vous facilite la vie avec DSM et une bibliothèque d'applications complète. Soyez simplement prêt à payer une prime pour le stockage de marque Synology si vous ne voulez pas d'alarmes de compatibilité dans votre interface utilisateur.
