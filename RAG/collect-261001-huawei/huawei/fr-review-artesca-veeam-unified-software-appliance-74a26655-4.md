---
id: collect-261001-huawei/huawei/fr-review-artesca-veeam-unified-software-appliance-74a26655-4
title: "fr-review-artesca-veeam-unified-software-appliance-74a26655"
domain: huawei
role: reference
task: reference
actors: []
dates: []
keywords: ["valuation"]
source: docs/RAG/collect-261001-huawei/fr-review-artesca-veeam-unified-software-appliance-74a26655.md
source_anchor: ""
source_lines: [54, 62]
sha256: c07e942d4a1fbcedcbd6799ca7b430ad9cb4c6d8d79dd0a9a94e983bde11598c
---

# fr-review-artesca-veeam-unified-software-appliance-74a26655

Scality a présenté en avant-première la prochaine version majeure d'ARTESCA+ et a promis une plus grande simplicité, une meilleure intégration et une automatisation accrue pour la configuration de la machine virtuelle Veeam embarquée. L'entreprise prévoit également d'ajouter la prise en charge du déploiement de l'appliance logicielle Veeam (basée sur Linux) et de nombreuses améliorations des fonctionnalités multi-nœuds et de haute disponibilité. Cette mise à jour réduira le temps de déploiement total de la solution et supprimera plusieurs étapes de configuration décrites dans cet article. Par ailleurs, bien que ce document ait été rédigé pour Veeam B&R V12, les organisations utilisant la version V13 constateront une amélioration des performances et de l'efficacité globales du stockage objet, et bénéficieront des dernières fonctionnalités de protection et de sécurité de la plateforme.
Conclusion
ARTESCA+ Veeam remplit parfaitement les promesses de Scality. Dans notre laboratoire, cette appliance logicielle unifiée a éliminé les incertitudes liées à la sauvegarde et au stockage en centralisant le chemin S3, en activant par défaut le versionnage et le verrouillage des objets, et en guidant l'installation grâce à un assistant simple et intuitif, utilisable par un informaticien généraliste. Résultat : un référentiel fonctionnel et immuable en quelques minutes, avec une maintenance minimale.
La sécurité est axée sur la praticité plutôt que sur la performance. Des points d'accès internes uniquement, l'absence de DNS externe, l'isolation des identifiants, la gestion des identités et des accès (IAM) et l'authentification multifacteur (MFA) optionnelle réduisent l'exposition aux risques tout en préservant la simplicité d'administration au quotidien. La supervision via Grafana offre une visibilité complète sur les services, les nœuds et les disques, facilitant ainsi la vérification du bon fonctionnement des sauvegardes et la détection des problèmes avant qu'ils n'entraînent des interruptions de service.
D'un point de vue commercial, la valeur ajoutée réside dans la rapidité de mise en œuvre et la réduction des coûts opérationnels. Les équipes bénéficient de Veeam et d'une cible S3 au sein d'une infrastructure unique sur x86 standard, ce qui leur permet de conserver une grande flexibilité dans le choix du matériel et d'éviter toute dépendance vis-à-vis d'un fournisseur. Scality positionne cette appliance unifiée pour une capacité utilisable d'environ 20 à 440 To, ce qui correspond aux besoins des succursales, des sites périphériques et des PME qui ont dépassé les capacités d'une plateforme scale-up mais ne souhaitent pas déployer une plateforme objet distincte.
Il convient de tenir compte de quelques précautions. L'appliance unifiée est un nœud unique. Pour une résilience accrue ou une capacité supérieure, déployez ARTESCA en cluster multi-nœuds et conservez Veeam comme plan de contrôle. L'exposition externe à S3 est possible, mais il s'agit d'une option à activer qui doit être documentée et gérée avec des identifiants distincts.
Globalement, il s'agit d'une approche mature, s'appuyant sur la longue expérience de Scality et une équipe possédant des décennies d'expérience dans le stockage d'entreprise, qui associe des valeurs par défaut judicieuses, une forte immuabilité et des opérations simples.
Démonstration de la plateforme
Ce rapport est sponsorisé par Scality. Tous les points de vue et opinions exprimés dans ce rapport sont basés sur notre évaluation impartiale du ou des produits étudiés.
