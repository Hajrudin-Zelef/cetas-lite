---
id: collect-261001-ia-llm/ia-llm/fr-review-hpe-alletra-storage-mp-b10000-and-nist-csf-2-0-a-full-stack-cyber-resi-99bd300c-5
title: "fr-review-hpe-alletra-storage-mp-b10000-and-nist-csf-2-0-a-full-stack-cyber-resi-99bd300c"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["incident", "valuation"]
source: docs/RAG/collect-261001-ia-llm/fr-review-hpe-alletra-storage-mp-b10000-and-nist-csf-2-0-a-full-stack-cyber-resi-99bd300c.md
source_anchor: ""
source_lines: [64, 79]
sha256: 1e7f2d61b494a656b6a80a6ce7827d2836040ebfaee2e1467afd98966420de5d
---

# fr-review-hpe-alletra-storage-mp-b10000-and-nist-csf-2-0-a-full-stack-cyber-resi-99bd300c

Niveau 2 : Protection native VME pour la restauration quotidienne au niveau de la machine virtuelle. Le flux de snapshots HPE Morpheus VM Essentials assure la reprise opérationnelle au niveau de la machine virtuelle. Ce niveau gère les opérations de restauration courantes et sert de solution de repli en cas d’indisponibilité du niveau 1.
Niveau 3 : Instantanés Virtual Lock immuables B10000 pour la restauration au niveau du volume. Ce niveau est crucial en cas d’attaque par ransomware. Les instantanés Virtual Lock sont protégés contre la suppression et la modification, et l’opération de promotion est native à B10000, sans dépendance à un éditeur de logiciels tiers.
Niveau 4 : Restauration StoreOnce Catalyst pour la conservation à long terme et la récupération inter-baies. Les niveaux 1 à 3 protègent contre les incidents durant quelques heures ou quelques jours. Le niveau 4 protège contre les incidents durant quelques semaines ou quelques mois, notamment lorsque l’attaquant a établi une persistance bien avant le déclenchement du chiffrement et lorsque le seul point de restauration valide est antérieur à la période de conservation de tout instantané de stockage principal. Les stockages Catalyst sont dédupliqués, immuables et peuvent être répliqués sur plusieurs systèmes StoreOnce ou détachés vers le stockage objet Cloud Bank. Catalyst Copy et Cloud Bank Detach constituent ensemble le niveau 1 du modèle 3-2-1-1. C’est également à ce niveau que la distinction entre infrastructure physique et VSA, mentionnée dans la présentation de l’architecture, prend toute son importance. En tant qu’option de récupération de dernier recours, le niveau 4 doit pouvoir résister à une compromission de l’hyperviseur. C’est pourquoi il est recommandé de déployer StoreOnce sur une appliance physique dédiée plutôt que sur une VSA en production.
Évaluation finale
Les deux questions essentielles que se pose un expert en sécurité concernant une plateforme de stockage sont : l’infrastructure reste-t-elle sous le contrôle de l’organisation et les données restent-elles protégées ? Après analyse de l’architecture au laboratoire de Fort Collins, il apparaît que la B10000 apporte des réponses plus complètes à ces deux questions qu’une simple baie de stockage, car la réponse ne se limite pas à cette dernière.
La plateforme enregistre chaque action administrative, expose tous les paramètres pertinents via son API et transmet les données d'audit et de télémétrie de sécurité au SIEM déjà utilisé par l'organisation. Le client reste responsable de la rédaction du plan de gouvernance, mais les outils nécessaires à sa mise en œuvre, à sa documentation et à sa validation par un auditeur sont disponibles et accessibles, et non dissimulés derrière un portail fournisseur. Grâce aux snapshots de verrouillage virtuel immuables, même un compte administrateur compromis ne peut pas détruire discrètement les points de restauration ; toute tentative en ce sens est consignée et peut faire l'objet d'une requête.
En matière de protection, la boucle de détection et de restauration au cœur de l'architecture est l'élément qui a démontré la plus grande fiabilité. Le B10000 a détecté la charge de travail de chiffrement simulée en quatre à cinq minutes au niveau des blocs et a capturé un instantané forensique dès sa détection, sans attendre de système externe. Le volume affecté a ensuite été restauré sans problème, aussi bien par une promotion d'instantané Virtual Lock planifiée que par une restauration StoreOnce Catalyst. C'est cette séquence qui est cruciale lors d'un incident, et elle s'est déroulée sans aucun produit tiers dans la chaîne de dépendances.
Ce qui distingue cette architecture, ce n'est pas tant que le B10000 soit nettement plus sécurisé que les baies concurrentes prises individuellement. C'est plutôt que HPE est l'un des rares fournisseurs à proposer le stockage, la virtualisation, la réplication, la sauvegarde et l'observabilité au sein d'un système intégré, et à tester ce système dans son ensemble face à des ransomwares réels dans son propre laboratoire, au lieu de valider chaque composant séparément. Pour une entreprise propriétaire de son infrastructure mais soumise à une équipe de sécurité distincte et à un organisme de réglementation externe, cette coordination fait toute la différence entre un ensemble de produits performants et une stratégie de résilience capable de fonctionner sous pression.
L'architecture évolue rapidement. HPE maintient un rythme de publication régulier pour l'ensemble de la pile technologique, avec une intégration plus poussée entre VM Essentials et le B10000, une détection des ransomwares étendue à d'autres types de données que les volumes de blocs, et une architecture de référence complète publiée, autant d'éléments prévus dans sa feuille de route à court terme. Cette évolution s'oriente vers une coordination davantage gérée nativement par la plateforme et vers une conception plus claire pour la mise en place du type d'architecture de résilience analysé dans cet article.
Pendant des années, la baie de stockage a été considérée comme l'un des éléments les plus sûrs du bâtiment, principalement parce qu'on ne s'en souciait guère. Cette architecture démontre que la baie devrait être tout le contraire : non pas un boîtier négligé dans un coin, mais un acteur clé dans la détection et la protection contre une attaque, à condition que l'organisation prenne les mesures nécessaires pour transformer ces capacités en un plan de sécurité.
Page produit HPE B10000
Ressources supplémentaires
Vidéos
hpe.com/uk/en/resource-library.video.hpe-alletra-storage-mp-b10000-r6-ransomware-detection-with…
hpe.com/uk/en/resource-library.video.hpe-alletra-storage-mp-b10000-and…
Ce rapport est parrainé par HPE. Tous les points de vue et opinions exprimés dans ce rapport sont basés sur notre vision impartiale du ou des produits à l'étude.
