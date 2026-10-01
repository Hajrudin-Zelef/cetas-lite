---
id: collect-250926-servers-hardware/servers-hardware/fr-review-supermicro-jumpstart-review-a-week-with-an-nvidia-hgx-b200-378de64c-2
title: "fr-review-supermicro-jumpstart-review-a-week-with-an-nvidia-hgx-b200-378de64c"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["diffusion", "gpu"]
source: docs/RAG/clean4/fr-review-supermicro-jumpstart-review-a-week-with-an-nvidia-hgx-b200-378de64c.md
source_anchor: ""
source_lines: [15, 41]
sha256: 40f12e031159fba73f843988043b0c0ffdced096b790ccb41374ceb11bc9d90b
---

# fr-review-supermicro-jumpstart-review-a-week-with-an-nvidia-hgx-b200-378de64c

Traditionnellement, lorsqu'un GPU traite des données stockées sur un disque NVMe, ces données doivent d'abord transiter par le CPU et la mémoire système avant d'atteindre le GPU. Ce processus engendre des goulots d'étranglement, le CPU jouant le rôle d'intermédiaire, ce qui augmente la latence et consomme des ressources système précieuses. GPUDirect Storage élimine cette inefficacité en permettant au GPU d'accéder directement aux données depuis le périphérique de stockage via le bus PCIe. Ce chemin direct réduit la surcharge liée aux transferts de données, permettant ainsi des transferts plus rapides et plus efficaces.
Les charges de travail d'IA, notamment celles impliquant l'apprentissage profond, sont très gourmandes en données. L'entraînement de grands réseaux neuronaux nécessite le traitement de téraoctets de données, et tout retard dans le transfert de données peut entraîner une sous-utilisation des GPU et des temps d'entraînement plus longs. GPUDirect Storage relève ce défi en garantissant que les données sont transmises au GPU le plus rapidement possible, minimisant ainsi les temps d'inactivité et maximisant l'efficacité de calcul.
En outre, GDS est particulièrement utile pour les charges de travail impliquant la diffusion de grands ensembles de données, comme le traitement vidéo, le traitement du langage naturel ou l'inférence en temps réel. En réduisant la dépendance au processeur, GDS accélère le déplacement des données et libère les ressources du processeur pour d'autres tâches, améliorant ainsi encore les performances globales du système.
Au-delà de la bande passante brute, GPUDirect avec NVMe-oF (TCP/RDMA) offre également des E/S à très faible latence. Ainsi, les GPU ne manquent jamais de données, ce qui en fait le système idéal pour l'inférence d'IA en temps réel, les pipelines d'analyse et la relecture vidéo.
Débit de lecture séquentielle GDSIO
Lors de nos tests de lecture séquentielle GDSIO sur la plateforme B200, la charge de travail initiale était modeste, avec un seul thread atteignant environ 14 à 15 Gio/s, selon la taille des blocs. Dès que nous avons augmenté le nombre de threads et la taille des blocs, les performances se sont rapidement améliorées. Le passage à 2 et 4 threads a permis d'atteindre un débit de 20 à 36 Gio/s, démontrant ainsi l'excellente montée en puissance du système grâce au parallélisme.
L'accélération réelle s'est produite à partir de huit threads, la plupart des charges de travail se stabilisant alors autour de 30 Gio/s. Les blocs de grande taille ont été les plus performants, avec des résultats constamment excellents pour les tailles de 5 Mo et 10 Mo, quel que soit le nombre de threads.
Le débit a finalement atteint un maximum d'environ 43 Gio/s avec une taille de bloc de 10 Mo à 256 threads, ce qui représente le taux de lecture séquentiel soutenu le plus élevé que nous ayons observé dans ce test.
Latence de lecture séquentielle GDSIO
En termes de latence, la charge de travail a démarré avec une excellente réactivité : les lectures mono-thread se situaient entre 0.06 et 0.1 ms pour les petits blocs. À mesure que le nombre de threads augmentait, la latence diminuait progressivement, restant inférieure à 1 ms jusqu’à 8 threads pour la plupart des charges de travail.
Au-delà de 16 threads, les tailles de blocs plus importantes ont commencé à engendrer des latences de plusieurs millisecondes, le chemin de stockage étant de plus en plus saturé. La latence maximale a été observée à la fin du test, avec une taille de bloc de 10 Mo et 256 threads, atteignant un pic d'un peu plus de 1.2 seconde (environ 1 200 ms), ce qui correspond à une charge extrême conçue pour saturer le système.
Débit d'écriture séquentielle GDSIO
Pour les écritures séquentielles, les performances étaient beaucoup plus stables qu'en lecture. La charge de travail se stabilisait rapidement, la plupart des combinaisons de threads et de tailles de blocs se situant autour de 6.3 à 6.5 Gio/s. Cela indique que le débit d'écriture atteint un plafond constant très tôt, probablement lié au support de stockage et à la mise en mémoire tampon plutôt qu'aux limitations du GPU ou du PCIe.
Augmenter le nombre de threads n'a pas eu d'impact significatif, le débit restant quasiment inchangé entre 2 et 128 threads. Le seul résultat notable est survenu en fin de test : avec une taille de bloc de 10 Mo et 256 threads, le débit a atteint un pic de 18.2 Gio/s, démontrant un léger avantage lorsque le système exploite pleinement les files d'attente profondes et l'agrégation d'écritures.
Latence séquentielle d'écriture GDSIO
La latence d'écriture était initialement relativement faible, se situant entre 0.15 et 0.5 ms pour les charges de travail monothread avec des blocs de petite taille. Avec l'augmentation du nombre de threads, la latence a crû beaucoup plus rapidement qu'en lecture, atteignant 1 à 4 ms pour quatre threads et 4 à 9 ms pour huit threads.
Dès que nous avons atteint 32 threads avec des blocs de plus grande taille, la latence a fortement augmenté, les blocs de 5 et 10 Mo passant dans la plage de 170 à 350 ms. Le cas le plus extrême était celui des blocs de 10 Mo avec 256 threads, qui a culminé à un peu moins de 3 secondes (environ 2 900 ms), illustrant clairement la rapidité avec laquelle le chemin d'écriture est saturé sous une charge parallèle importante.
Débit de lecture aléatoire GDSIO
Pour les lectures aléatoires, la charge de travail a rapidement augmenté. Les performances en mono-thread se situaient entre 11 et 31 Gio/s environ, selon la taille des blocs, les blocs plus volumineux bénéficiant immédiatement d'une bande passante plus élevée. Avec deux puis quatre threads, le débit a atteint 20 à 36 Gio/s, démontrant une excellente scalabilité dès les premières utilisations.
À partir de 8 threads, le système s'est stabilisé autour de 30 Gio/s, un résultat très similaire à celui observé lors du test de lecture séquentielle. Le meilleur résultat a été obtenu avec une taille de bloc de 10 Mo et 256 threads, atteignant un pic d'environ 42.7 Gio/s.
Latence de lecture aléatoire GDSIO
La latence de lecture aléatoire était initialement très faible, de l'ordre de 0.15 à 0.4 ms pour les charges de travail mono-thread avec des blocs de petite taille. À mesure que le nombre de threads augmentait, la latence s'élevait progressivement, restant inférieure à 1 ms jusqu'à 4 threads et atteignant environ 1 à 3 ms à 8 threads.
Dès que nous sommes passés à 32 threads et à des blocs plus volumineux, la latence a augmenté plus fortement, les transferts de 5 et 10 Mo atteignant des valeurs comprises entre 30 et 55 ms. Le cas le plus extrême a été observé avec 256 threads et une taille de bloc de 10 Mo, où la latence a culminé à un peu plus de 1.1 seconde (environ 1180 ms).
Débit d'écriture aléatoire GDSIO
Les performances en écriture aléatoire étaient très homogènes, la plupart des tailles de blocs et des nombres de threads se situant autour de 5.8 à 6.1 Gio/s. La charge de travail a atteint ce niveau presque instantanément et n'a que très peu progressé avec l'ajout de threads supplémentaires, ce qui indique que le chemin d'écriture atteint rapidement sa limite.
La seule exception notable est apparue à la fin du test, où une taille de bloc de 10M à 256 threads a brièvement atteint 12.5Gio/s, bénéficiant probablement d'une mise en file d'attente profonde et d'une agrégation d'écritures sous une charge parallèle importante.
Latence d'écriture aléatoire GDSIO
