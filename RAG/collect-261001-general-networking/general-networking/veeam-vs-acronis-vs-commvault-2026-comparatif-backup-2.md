---
id: collect-261001-general-networking/general-networking/veeam-vs-acronis-vs-commvault-2026-comparatif-backup-2
title: "veeam-vs-acronis-vs-commvault-2026-comparatif-backup"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "CISA", "Google"]
dates: ["2026-03-31", "2026-08-13"]
keywords: ["agent", "agents", "aws", "cyber", "license", "sandbox"]
source: docs/RAG/collect-261001-general-networking/veeam-vs-acronis-vs-commvault-2026-comparatif-backup.md
source_anchor: ""
source_lines: [29, 63]
sha256: 475551d6b4b2f2c96d806130338d8b3b2f4aa66f85e6a1a8caf65176a912c688
---

# veeam-vs-acronis-vs-commvault-2026-comparatif-backup

| Critère | Veeam Backup & Replication | Acronis Cyber Protect Cloud | Commvault Cloud Platform | 
|---|---|---|---|
| Dernière version connue (août 2026) | 13.1.1.18 (13/08/2026) | 26.06, agent 26.6.42752 | Non communiquée (mise à jour continue) | 
| Modèle de licence | Veeam Universal License (par charge de travail) | Par poste / par charge de travail | Sur devis, par capacité ou abonnement | 
| Hyperviseurs supportés | VMware, Hyper-V, Nutanix AHV, Proxmox VE, XCP-ng, XenServer, oVirt, Sangfor aSV | Agents Windows/macOS/Linux, VM via agent | VMware, Hyper-V, Azure Stack, Nutanix AHV | 
| Cloud public | AWS, Azure, Google Cloud (matrice large) | Couverture cloud via agents, matrice partielle | 30+ services AWS, Azure | 
| Immuabilité | Object Lock S3/Azure Blob/Wasabi, dépôts Linux durcis | Non détaillée publiquement | Air Gap Protect, AES-256 | 
| Air gap | Veeam Data Cloud Vault, copies hors ligne | Non documenté explicitement | Domaine air-gappé natif (Air Gap Protect) | 
| Anti-malware intégré | Secure Restore (scan avant restauration) | EDR intégré à l’agent de sauvegarde | Détection via Metallic AI | 
| Restauration vérifiée en sandbox | SureBackup | Non mis en avant | Cloud Rewind (orchestration) | 
| Authentification renforcée | RBAC, SSO SAML, validation à quatre yeux | Console multi-tenant, MFA | RBAC entreprise | 
| Cible principale | ETI, grands comptes, environnements hybrides | PME, indépendants, MSP | Grands comptes, environnements hybrides complexes | 
| Statut Magic Quadrant Gartner 2026 | Leader (en tête de l’exécution) | Hors périmètre depuis 2024 | Leader | 
| Note Gartner Peer Insights | 4,8/5 (124 avis, au 31/03/2026) | Avis positifs, moyenne agrégée non publiée | Non publiée dans les sources consultées | 

## Modèles de déploiement : on-premise, cloud ou hybride

Le choix d’un éditeur de sauvegarde ne se limite pas aux fonctionnalités : le modèle de déploiement pèse tout autant sur le coût total et sur la surface d’attaque exposée. Veeam Backup & Replication reste conçu en priorité pour un déploiement on-premise, avec un serveur de sauvegarde installé sur l’infrastructure du client, ce qui donne un contrôle total sur les données mais reporte aussi sur l’équipe IT la responsabilité du durcissement de ce serveur, exactement le composant visé par l’avis CERT-FR du 5 août 2026. Veeam propose toutefois des briques cloud complémentaires, comme Veeam Data Cloud Vault pour le stockage immuable hors site, et un modèle SaaS via des partenaires hébergeurs, illustré par l’offre référencée sur le catalogue G-Cloud britannique.

Acronis Cyber Protect Cloud a été pensé dès l’origine comme une plateforme cloud multi-tenant, hébergée par Acronis ou par un partenaire MSP, avec des agents légers installés sur chaque poste ou serveur protégé. Ce choix réduit la charge d’administration pour une PME qui ne veut pas gérer d’infrastructure de sauvegarde dédiée, mais implique aussi de confier la garde des données à un tiers, un point à vérifier au regard des exigences de souveraineté des données propres à certains secteurs réglementés en France. Commvault Cloud Platform, de son côté, propose un déploiement hybride natif : les workloads peuvent rester on-premise, être répliqués vers le cloud public, ou vivre nativement dans un environnement cloud, avec une seule console de gestion pour orchestrer l’ensemble. Cette flexibilité explique en partie pourquoi Commvault reste privilégié par les grands comptes qui opèrent des environnements hérités aux côtés de charges de travail cloud-natives récentes.

Pour une entreprise française soumise à NIS2 ou évoluant dans un secteur sensible, la question de la localisation des données mérite d’être posée explicitement à chaque éditeur avant signature. Un contrat de sauvegarde cloud qui héberge les données hors de l’Union européenne peut compliquer une démonstration de conformité, même si la solution technique elle-même est irréprochable.

## Le bulletin CERT-FR du 5 août 2026 : ce qu’il faut vraiment savoir

L’avis CERTFR-2026-AVI-0968 vise deux composants précis de l’écosystème Veeam : Veeam Service Provider Console dans les versions antérieures à 9.3.0.35057, et Veeam ONE dans les versions antérieures à 13.1.0.7034. Le CERT-FR recommande d’appliquer sans délai les correctifs listés dans le bulletin de sécurité de l’éditeur. Pour les hébergeurs et MSP français qui exploitent Service Provider Console afin de superviser plusieurs clients depuis une seule interface, une console compromise ouvre potentiellement l’accès à l’ensemble des locataires gérés, ce qui explique le niveau de vigilance demandé par l’agence.

En parallèle, la base européenne des vulnérabilités EUVD référence sous l’identifiant EUVD-2026-11595 une élévation de privilèges locale dans Veeam Backup & Replication, avec un score CVSS 3.1 de 8,8. Les versions concernées sont la branche 13 antérieure à 13.0.1 et la branche 12 antérieure à 12.3.2, ainsi que des plages similaires pour Veeam Backup and Recovery. Le score EPSS, qui estime la probabilité d’exploitation dans les 30 jours, se situe autour de 0,22 %, et la vulnérabilité n’apparaît ni dans le catalogue KEV de la CISA ni dans son équivalent européen au moment de la dernière mise à jour de la fiche. Aucune activité d’exploitation n’a par ailleurs été détectée par les honeypots Shadowserver. Le détail technique de cette faille et de son correctif figure dans la note KB4743 publiée par Veeam, qui documente les vulnérabilités corrigées dans la version 12.3.2.

Un point de prudence s’impose sur un troisième signalement. Des rapports de veille sur les menaces publiés fin août 2026 évoquent une éventuelle faille d’exécution de code à distance, référencée de manière informelle CVE-2026-44963, qui reposerait sur une désérialisation non sécurisée dans Veeam Backup & Replication 12.x. À la date du 29 août 2026, cette information n’a été confirmée ni par une analyse technique publique, ni par une reconnaissance officielle de l’éditeur. Elle doit donc être traitée comme une rumeur à surveiller, pas comme un fait acquis. Pour approfondir la méthode de défense en profondeur recommandée face à ce type de scénario, l’article du site sur la règle de sauvegarde 3-2-1-1-0 anti-ransomware détaille comment limiter l’impact d’une compromission du serveur de sauvegarde lui-même.

## Immuabilité, air gap et détection IA : la vraie bataille anti-ransomware

Sur le papier, les trois éditeurs promettent une protection contre le rançongiciel. Dans le détail, les mécanismes ne se valent pas. Veeam s’appuie sur l’immuabilité au niveau du stockage objet : verrouillage Object Lock sur AWS S3, Azure Blob ou Wasabi, et dépôts Linux durcis avec attributs étendus non modifiables. Veeam Data Cloud Vault ajoute une couche de stockage immuable et chiffré préconfigurée, pensée pour les équipes qui ne veulent pas gérer elles-mêmes l’infrastructure de stockage cible. Avant toute restauration, la fonction Secure Restore analyse le point de restauration avec un moteur anti-malware, et SureBackup permet de tester automatiquement, dans un environnement isolé, qu’une machine restaurée démarre réellement et fonctionne.

