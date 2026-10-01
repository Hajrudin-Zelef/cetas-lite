---
id: collect-261001-cisco/cisco/fr-review-amd-epyc-4005-review-am5-economics-with-enterprise-focus-5f50297a-3
title: "fr-review-amd-epyc-4005-review-am5-economics-with-enterprise-focus-5f50297a"
domain: cisco
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd", "apache"]
source: docs/RAG/collect-261001-cisco/fr-review-amd-epyc-4005-review-am5-economics-with-enterprise-focus-5f50297a.md
source_anchor: ""
source_lines: [104, 108]
sha256: 717a88e74ce0af3a7ecbbbf1ef2d53099a3b757f520deadfbbd88516bd9da3dc
---

# fr-review-amd-epyc-4005-review-am5-economics-with-enterprise-focus-5f50297a

| Lecture séquentielle (128 Ko) | 1,322.3 Mo / s | 10,578 | 0.459 ms | 
Conclusion
L'EPYC 4005 remplit sa mission. Il simplifie l'AM5 pour les constructeurs de systèmes et les hébergeurs, tout en offrant une marge de manœuvre précieuse grâce à Zen 5, à la DDR5-5600 validée, à l'AVX-512 et à une couverture SKU plus large. Cette famille se positionne parfaitement entre Ryzen sur AM5 et les plateformes SP5 et SP6 plus imposantes, offrant aux fournisseurs de services une solution d'entreprise sans le coût des gros sockets. La carte SKU couvre les modèles de déploiement courants. L'EPYC 4545P propose un processeur 16 cœurs à 65 W pour un hébergement dense et des licences Windows Server propres. L'EPYC 4565P reste la solution standard à 170 W, tandis que l'EPYC 4585PX ajoute le V-Cache 3D pour les services sensibles à la latence qui s'appuient sur un cache de dernier niveau plus important. Une validation mémoire jusqu'à 192 Go sur quatre UDIMM et des E/S stables avec jusqu'à 28 voies PCIe Gen5 complètent la plateforme.
Nos données montrent où cela compte. Le 16PX à 4585 cœurs remporte des victoires nettes dans les tâches sensibles au calcul et au cache. y-cruncher exécute des tâches 4.5 à 10.2 % plus rapidement sur 1 à 10 milliards de données, tandis que les exécutions BBP réduisent le temps jusqu'à 38.6 %. Blender en bénéficie également, avec des gains importants sur Monster, Junkshop et Classroom. Dans la suite Phoronix, le 4585PX réduit le temps de compilation du noyau Linux de 16.8 %, traite 36.9 % de requêtes Apache en plus par seconde et enregistre une augmentation de 78.6 % de la vérification OpenSSL. STREAM et 7-Zip se rapprochent légèrement du précédent 4564P dans notre configuration.
L'EPYC 4005 est une mise à jour modérée mais significative d'AMD. Elle préserve l'aspect économique et la simplicité qui ont facilité l'adoption de la version 4004, tout en améliorant les performances et la flexibilité pratiques dans un large éventail de cas d'utilisation hébergés et en périphérie. Si vous développez des serveurs AM5 monosocket à grande échelle, cette solution par défaut est idéale pour les nouveaux déploiements et constitue une mise à niveau judicieuse lorsque le cache ou l'efficacité sont des enjeux, notamment lorsqu'elle est associée au serveur MSI S1012-02 doté de la boucle liquide interne.
