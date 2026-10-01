---
id: collect-261001-general-networking/general-networking/fr-review-dapustor-r6060-122tb-review-read-heavy-gen5-qlc-at-scale-98545c1d-6
title: "fr-review-dapustor-r6060-122tb-review-read-heavy-gen5-qlc-at-scale-98545c1d"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["attention"]
source: docs/RAG/collect-261001-general-networking/fr-review-dapustor-r6060-122tb-review-read-heavy-gen5-qlc-at-scale-98545c1d.md
source_anchor: ""
source_lines: [130, 133]
sha256: 7c308a5f76cfbb056a583e82e4baa2de11702e90930d75711476a2dcde12463c
---

# fr-review-dapustor-r6060-122tb-review-read-heavy-gen5-qlc-at-scale-98545c1d

DapuStor propose le R6060 dans une large gamme de formats, avec des options U.2, E3.L et E1.L allant de 15.36 To à 122.88 To, ainsi qu'une version haut de gamme de 245 To. Cette diversité offre une grande flexibilité aux intégrateurs, mais souligne également l'importance d'une vérification préalable de la compatibilité. La variante E3.L 2T de 122.88 To que nous avons testée se situe en dehors des déploiements U.2 et E3.S les plus courants dans les centres de données. De manière générale, les SSD EDSFF nécessitent toujours une attention particulière concernant les emplacements de disques, l'épaisseur, la longueur et l'allocation des lignes PCIe avant toute commande. Un disque de 2 To ne s'insère pas dans un emplacement de 1 To, et un disque E3.L ne s'installe pas dans un emplacement E3.S. Bien que cela soit courant pour l'intégration de SSD EDSFF aujourd'hui, il est important de le signaler pour toute personne intégrant ces variantes plus denses pour la première fois dans un déploiement.
Pour les niveaux de stockage pour lesquels le R6060 est conçu, il a bien fonctionné et offre à DapuStor une option QLC Gen5 haute capacité crédible.
Page produit – DapuStor R6060 122 To
Classement : Le DapuStor R6060 122 To occupe la première place du classement des meilleurs SSD d'entreprise en tant que meilleure alternative haute capacité.
