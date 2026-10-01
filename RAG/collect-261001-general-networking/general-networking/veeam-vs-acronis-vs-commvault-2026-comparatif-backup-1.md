---
id: collect-261001-general-networking/general-networking/veeam-vs-acronis-vs-commvault-2026-comparatif-backup-1
title: "veeam-vs-acronis-vs-commvault-2026-comparatif-backup"
domain: general-networking
role: reference
task: reference
actors: ["AWS", "Google", "Microsoft"]
dates: []
keywords: ["agent", "agents", "aws", "cyber", "incident"]
source: docs/RAG/collect-261001-general-networking/veeam-vs-acronis-vs-commvault-2026-comparatif-backup.md
source_anchor: ""
source_lines: [1, 28]
sha256: 38e6a8a8d273a26c8e42dffcabc5d9d00fbec38cc1be61361b16a50dcf81ad16
---

# veeam-vs-acronis-vs-commvault-2026-comparatif-backup

Le 5 août 2026, le CERT-FR a publié l’avis CERTFR-2026-AVI-0968 sur de multiples vulnérabilités touchant Veeam Service Provider Console et Veeam ONE. Trois semaines plus tard, une entrée distincte de la base européenne EUVD (EUVD-2026-11595) a confirmé une faille d’élévation de privilèges notée 8,8 sur 10 dans Veeam Backup & Replication. Pour les DSI et RSSI français, la question n’est plus seulement de savoir quel logiciel de sauvegarde restaure le plus vite après un rançongiciel. Elle est aussi de savoir quel éditeur corrige le plus vite ses propres failles, et lequel protège vraiment les données pendant l’attaque. Ce comparatif met face à face Veeam Backup & Replication, Acronis Cyber Protect et Commvault Cloud Platform sur les critères qui comptent en 2026 : architecture, immuabilité, prix par charge de travail, positionnement Gartner et conformité NIS2.

## Pourquoi ce comparatif Veeam, Acronis et Commvault s’impose en août 2026

Trois signaux convergent au même moment. D’abord la sécurité des outils de sauvegarde eux-mêmes est devenue un sujet à part entière. L’avis CERT-FR du 5 août 2026 liste des impacts sérieux sur les produits Veeam concernés : atteinte à la confidentialité, contournement de la politique de sécurité, déni de service à distance, exécution de code arbitraire et injection SQL. Un serveur de sauvegarde compromis donne accès non seulement aux données de production, mais aussi aux copies de restauration, ce qui en fait une cible stratégique pour les groupes de rançongiciel. Le site a déjà documenté ce risque dans son article sur le rançongiciel INC, qui a visé spécifiquement des infrastructures Veeam chez plus de 830 victimes recensées.

Ensuite, le calendrier réglementaire européen pousse les entreprises à revoir leur stratégie de résilience. La directive NIS2 impose des obligations de continuité d’activité aux opérateurs essentiels et importants, et le Cyber Resilience Act ajoute des exigences de transparence sur les vulnérabilités logicielles. Un éditeur de sauvegarde qui publie un correctif en quelques jours n’a pas le même profil de risque qu’un éditeur silencieux pendant des semaines. Enfin, les trois plateformes ont toutes changé de visage depuis un an : Veeam est passé en version 13.1, Acronis a réorienté sa plateforme Cyber Protect Cloud vers les prestataires de services managés, et Commvault a rebaptisé son offre Commvault Cloud Platform, propulsée par Metallic AI. Comparer ces trois éditeurs aujourd’hui donne une photographie beaucoup plus fidèle qu’un comparatif écrit il y a dix-huit mois.

## Veeam Backup & Replication : présentation, architecture et version 13.1

Veeam Backup & Replication reste le nom le plus cité dès qu’on évoque la sauvegarde d’infrastructures virtualisées en entreprise. La dernière build généralement disponible fin août 2026 est la version 13.1.1.18, publiée le 13 août 2026, avec un flux de maintenance parallèle en version 13.0.3.63 sorti le 25 août 2026. La branche précédente, la 12.3.2.4854, avait été livrée les 8 et 9 juin 2026 et reste largement déployée dans les parcs qui n’ont pas encore migré vers la v13.

L’architecture Veeam couvre un périmètre technique très large. Les hyperviseurs pris en charge incluent VMware vSphere de la version 6.x à 8.0, Microsoft Hyper-V, Nutanix AHV, Proxmox VE, Scale Computing HyperCore, HPE Morpheus VM Essentials, oVirt, XCP-ng, Citrix XenServer et Sangfor aSV. Côté cloud public, Veeam sauvegarde et restaure des instances AWS EC2, RDS, DynamoDB, Redshift, EFS et FSx, des machines virtuelles et bases Azure (SQL, Cosmos DB), ainsi que des ressources Google Cloud (VM, SQL, Spanner). La plateforme couvre aussi Microsoft Entra ID, Kubernetes, les partages NAS et, selon des revues indépendantes, Microsoft 365 et Salesforce. C’est aujourd’hui la matrice de compatibilité la plus large des trois éditeurs comparés ici.

## Acronis Cyber Protect Cloud : présentation et positionnement MSP

Acronis a construit sa stratégie autour d’une idée simple : fusionner sauvegarde et cybersécurité dans une seule console, plutôt que de superposer deux outils distincts. Cyber Protect Cloud embarque un moteur anti-malware et une détection comportementale de type EDR (Endpoint Detection and Response) directement sur l’agent de sauvegarde. La dernière branche documentée en juin 2026 correspond à la version 26.06, avec des agents Windows, macOS et Linux en build 26.6.42752. Le rythme de publication d’Acronis suit un schéma mensuel AA.MM, contrairement au cycle de versions majeures de Veeam ou Commvault.

Le positionnement d’Acronis a changé de manière significative depuis 2024. L’éditeur ne figure plus dans le Magic Quadrant Gartner dédié à la sauvegarde d’entreprise, Gartner justifiant cette sortie par une orientation de plus en plus marquée vers les prestataires de services managés (MSP), les postes de travail et les charges de travail en périphérie de réseau plutôt que vers la sauvegarde d’entreprise au sens large. Cela ne rend pas Acronis moins pertinent, cela signale simplement un public cible différent : PME, indépendants et surtout MSP qui gèrent des dizaines de clients depuis une console multi-tenant unique, avec une tarification par poste très lisible.

## Commvault Cloud Platform : architecture et Metallic AI

Commvault a achevé sa transition de marque en 2025-2026 : l’ancienne offre Commvault Complete Backup & Recovery et la solution SaaS Metallic sont désormais réunies sous le nom Commvault Cloud Platform, propulsée par Metallic AI. Les documents accessibles publiquement ne donnent pas de numéro de version unique pour cette plateforme unifiée, contrairement à Veeam ou Acronis, ce qui reflète une approche produit davantage centrée sur des mises à jour continues en mode cloud que sur des jalons de version classiques.

Sur le plan technique, Commvault protège VMware, Hyper-V, Azure Stack, Nutanix AHV et les charges conteneurisées, avec une couverture documentée de plus de 30 services AWS. Sa signature reste la fonctionnalité Air Gap Protect, qui isole les sauvegardes dans un domaine de sécurité air-gappé et immuable chiffré en AES-256. Commvault met aussi en avant des mécanismes d’orchestration de reprise comme Cloud Rewind et la récupération orchestrée d’Active Directory, deux fonctions pensées pour les grands comptes qui doivent redémarrer une chaîne d’applications entière après un incident, pas seulement une machine virtuelle isolée.

## Tableau comparatif : Veeam vs Acronis vs Commvault en 2026

