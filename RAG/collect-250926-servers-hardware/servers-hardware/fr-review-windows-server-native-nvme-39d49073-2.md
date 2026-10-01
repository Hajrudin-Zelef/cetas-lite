---
id: collect-250926-servers-hardware/servers-hardware/fr-review-windows-server-native-nvme-39d49073-2
title: "fr-review-windows-server-native-nvme-39d49073"
domain: servers-hardware
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/clean4/fr-review-windows-server-native-nvme-39d49073.md
source_anchor: ""
source_lines: [35, 59]
sha256: 0ec3b33afbe2f208d031dc47dd4dc6f54b0f2b317e59b562f40d1133c2f386f5
---

# fr-review-windows-server-native-nvme-39d49073

| IOPS | 1,598,959 | 2,636,516 | 1,217,176 | 1,493,637 | 583,192 | 583,638 | 710,978 | 758,252 | 
| Latence moyenne (ms) | 0.169 | 0.104 | 0.239 | 0.207 | 0.809 | 0.812 | 0.613 | 0.608 | 
| Utilisation totale du processeur (%) | 72.67 | 74.22 | 68.44 | 65.11 | 44.89 | 37.11 | 61.56 | 49.56 | 
| Métrique | 4K aléatoires |  | 64K aléatoires |  | Séquentiel 64K |  | Séquentiel 128K |  | 
|---|---|---|---|---|---|---|---|---|
|  | Non-autochtone | Originaire | Non-autochtone | Originaire | Non-autochtone | Originaire | Non-autochtone | Originaire | 
| Écrire |  |  |  |  |  |  |  |  | 
| Bande passante (Gio/s) | 1.803 | 1.756 | 7.654 | 7.655 | 44.67 | 50.087 | 50.477 | 50.079 | 
| IOPS | 472,725 | 460,383 | 125,391 | 125,406 | 731,859 | 820,603 | 413,495 | 410,232 | 
| Latence moyenne (ms) | 0.992 | 1.028 | 3.814 | 3.816 | 0.399 | 0.558 | 1.022 | 1.149 | 
| Utilisation totale du processeur (%) | 26.00 | 20.67 | 12.22 | 9.33 | 70.44 | 57.78 | 58.44 | 47.33 | 
Analyse des résultats
En commençant par des tests de lecture aléatoire de 4K et 64K, nous avons observé des vitesses de lecture nettement supérieures, avec une différence de près de 4 Gio/s entre les architectures de stockage natives et non natives (respectivement) lors du test de lecture aléatoire de 4K, et une amélioration de près de 16.9 Gio/s lors de la lecture aléatoire de 64K. Nous avons également constaté une augmentation significative des opérations de lecture séquentielles de 128K, nos tests montrant un gain de bande passante d'environ 5.8 Gio/s.
Étonnamment, nous n'avons constaté aucune augmentation significative de la bande passante lors de nos tests d'écriture aléatoire ou séquentielle, la seule différence notable étant une hausse d'environ 5.4 Gio/s pour les écritures séquentielles de 64 Ko. La plupart de nos résultats se situaient à moins de 100 Mio/s les uns des autres, ce qui suggère que les performances de la nouvelle architecture de stockage sont au moins équivalentes à celles de l'ancienne, même lorsqu'elles ne s'améliorent pas.
Le débit étant généralement corrélé à la latence, nous avons également observé des baisses importantes de la latence moyenne de lecture aléatoire pour les tests 4K et 64K. Nous avons constaté une diminution de 38.46 % pour les lectures aléatoires non natives 4K, passant de 0.169 milliseconde à 0.104 milliseconde. Les tests de lecture aléatoire 64K ont montré une diminution plus faible, d'environ 13.39 %. La latence n'a pas varié de façon significative pour les opérations de lecture séquentielle, mais les opérations d'écriture aléatoires et séquentielles ont globalement augmenté, malgré un débit similaire ou supérieur.
Outre l'augmentation des vitesses de lecture aléatoire, nos tests FIO ont révélé une autre tendance intéressante : une diminution substantielle de l'utilisation totale du processeur pour les opérations de lecture et d'écriture séquentielles à 64 Ko et 128 Ko. Les tests d'écriture séquentielle ont affiché les différences les plus marquées, avec une baisse moyenne de 12.66 % de l'utilisation du processeur à 64 Ko et une baisse presque identique de 11.11 % à 128 Ko. Le test de lecture séquentielle à 128 Ko a également enregistré une baisse d'utilisation de 12 %, contre seulement 7.78 % pour la lecture séquentielle à 64 Ko. Il convient de noter qu'avec un processeur suffisamment rapide, il est possible que les deux piles atteignent le plein potentiel d'un périphérique de stockage ; par conséquent, le débit pourrait ne pas augmenter, mais l'utilisation des ressources du processeur diminuerait.
Points clés à retenir
Bien que nombre de nos résultats aient présenté une variabilité inter-exécution après l'activation de la nouvelle architecture de stockage, nous avons pu corroborer plusieurs affirmations de Microsoft, notamment une bande passante de lecture accrue avec une latence réduite et une diminution générale de l'utilisation du processeur. Ce changement étant assez radical pour leur architecture de stockage Windows Server, vieille de plusieurs décennies, Microsoft activera d'abord la technologie NVMe native par défaut dans Windows Server vNext. Heureusement, cette fonctionnalité peut être activée sur Windows Server 2025 via une simple modification du registre ou une stratégie de groupe, permettant ainsi aux administrateurs de serveurs les plus audacieux de profiter de cette nouvelle architecture dès aujourd'hui (après avoir pris connaissance des risques liés à son déploiement).
Nous avons hâte de voir la technologie NVMe native activée par défaut sur la plateforme Windows Server, et nous espérons que les fabricants de SSD NVMe, de cartes RAID et de contrôleurs HBA adhéreront à cette technologie, car ils devraient être en mesure de faire passer les améliorations de Microsoft au niveau supérieur !
Références
Hands, J., Worley, D. et Lakhveer Kaur. (s.d.). Espaces de noms NVMe. Consulté le 30 décembre 2025 sur NVM Express : https://nvmexpress.org/resource/nvme-namespaces/
Lee, S. (15 septembre 2025). SNIA SDC 2025 – Stockage multi-files d'attente sous Windows. San Tomas, CA, États-Unis : Storage Networking Industry Association. Consulté le 29 décembre 2025 sur https://www.youtube.com/watch?v=dR-DWrmCba0&t
Lee, S. (16 septembre 2025). Stockage multi-files sous Windows : une nouvelle architecture pour le matériel de stockage haute performance. Consulté le 29 décembre 2025 sur le site de la conférence SNIA Developer : https://www.snia.org/sites/default/files/2025-10/SNIA-SDC25-Lee-Storage-Multi-Queue-On-Windows.pdf
Shekar, Y. (15 décembre 2025). Annonce de la prise en charge native de NVMe dans Windows Server 2025 : une nouvelle ère de performances de stockage s’ouvre. (Microsoft) Consulté le 29 décembre 2025 sur le site Windows Server News and Best Practices : https://techcommunity.microsoft.com/blog/windowsservernewsandbestpractices/announcing-native-nvme-in-windows-server-2025-ushering-in-a-new-era-of-storage-p/4477353
Association de l'industrie des réseaux de stockage (SNA). (s.d.). Qu'est-ce que le SCSI ? Consulté le 30 décembre 2025 sur le site de la SNA : https://www.snia.org/education/what-is-scsi
