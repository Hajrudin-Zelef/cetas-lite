---
id: collect-261001-rattrapage/rattrapage/huawei-nce-campus-guide-8
title: "Guide technique ultra-complet — iMaster NCE-Campus (Huawei)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["distribution", "license"]
source: docs/RAG/collect-261001-rattrapage/huawei_nce_campus_guide.md
source_anchor: ""
source_lines: [584, 670]
sha256: 78b1e5b8f3c5427ad692ab61048688cbdbca00d849c8b2dd34b111cf45123e19
---

# Guide technique ultra-complet — iMaster NCE-Campus (Huawei)

**Aucun prix public fiable** n'existe pour les licences NCE-Campus : les tarifs dépendent du partenaire, du volume, de la durée de souscription, du pays, des remises projet. Tout chiffre trouvé sur un forum est à considérer comme **non vérifié**.

Ce que l'on peut dire honnêtement :

- Le **ticket d'entrée** (licences + serveurs/VM + intégration + formation) est **sans commune mesure** avec eSight ou eKit : prévoir un **budget projet**, pas un simple achat de licences.
- Les **services d'intégration** (dimensionnement, installation, templates, migration, formation) représentent souvent **autant que les licences** sur un premier déploiement — les oublier est l'erreur budgétaire n°1.
- Le **renouvellement annuel** des souscriptions doit être inscrit au budget récurrent (OPEX), avec une clause de révision.
- **À vérifier auprès du partenaire** : devis détaillé avec (a) licences plateforme, (b) device-days par an, (c) CampusInsight, (d) services d'intégration, (e) formation, (f) support.

## 34. Checklist avant achat de licences

- [ ] Inventaire exact des équipements à gérer (modèle, quantité, version logicielle) — S310, AP361, AP761, AR720, USG6000...
- [ ] Matrice de compatibilité NCE × modèles × versions validée avec le partenaire
- [ ] Croissance prévisionnelle du parc sur la durée de souscription (+ marge)
- [ ] Choix de la durée (1 an / 3 ans) et de la cotermination
- [ ] CampusInsight : oui/non, quels packages
- [ ] Mode MSP ou entreprise — trancher
- [ ] Devis comparé sur TCO 3-5 ans vs statu quo (eSight/eKit)
- [ ] Services d'intégration chiffrés séparément
- [ ] Formation de l'équipe chiffrée et planifiée
- [ ] Procédure de renouvellement et alerte de fin de souscription définies
- [ ] Clause contractuelle : que se passe-t-il en fin de souscription non renouvelée ? (accès lecture ? blocage ? — **à vérifier** et à écrire)

---

# PARTIE 5 — INSTALLATION

## 35. Architecture matérielle requise — serveur / VM

NCE-Campus s'installe sur des **serveurs x86** (physiques ou VMs). Les exigences exactes (vCPU, RAM, disque, IOPS) dépendent de la **version** et du **dimensionnement** (nombre d'équipements gérés, télémétrie, CampusInsight ou non) — **toujours** utiliser le guide d'installation et l'outil de dimensionnement officiels de la version cible.

Principes généraux d'un chef de service :

- **Ne jamais sous-dimensionner la RAM** : bases de données + analyse temps réel = mémoire. Le premier symptôme d'un NCE sous-dimensionné est une interface lente et des traitements d'alarmes en retard (voir cas pratique 17).
- **Stockage** : disques rapides (SSD/NVMe) en RAID redondant, volumétrie calculée avec la rétention (section 19).
- **Réseau** : plusieurs interfaces (management, service, éventuellement backup), redondance.
- **Virtualisation** : si VM, **réservations** de CPU/RAM (pas de sur-allocation sauvage), datastore performant, anti-affinité entre les nœuds HA.
- **HA** : deux machines sur des hôtes/racks/sites distincts (section 22).

## 36. Système d'exploitation supporté

NCE-Campus est livré comme une **solution packagée** (appliance logicielle / image) — le système d'exploitation sous-jacent est **imposé par Huawei** (généralement une distribution Linux durcie/validée par l'éditeur, installée avec le package). Concrètement :

- On ne choisit pas « Ubuntu vs RHEL » : on déploie **l'image/la procédure Huawei**.
- Vérifier les prérequis hyperviseur (versions VMware/Hyper-V/KVM supportées) dans le guide d'installation.
- Le durcissement OS est normalement intégré ; le durcissement **applicatif** reste à faire (section 105).

**À vérifier** pour la version cible : image ISO/OVA fournie, versions d'hyperviseur supportées, procédure d'installation (graphique/ligne de commande).

## 37. Prérequis réseau (ports, flux, DNS, NTP)

Avant d'installer, verrouiller ces prérequis — 80 % des échecs d'onboarding viennent d'ici :

- **Connectivité IP** : routage entre NCE et tous les sites/équipements gérés (ou via VPN/MPLS) ; VLAN de management dédié recommandé.
- **DNS** : résolution directe/inverse fonctionnelle pour NCE et les équipements (le ZTP par nom en dépend).
- **NTP** : **même source de temps** partout (NCE, équipements, serveurs d'authentification). Un décalage > quelques minutes casse les certificats, les tokens API et les logs corrélés.
- **Flux southbound** : autoriser NETCONF (TCP/830), SNMP (UDP/161-162), CAPWAP (UDP/5246-5247 si WAC), HTTPS (TCP/443), HTTP/2/télémétrie et TCP de synchro d'authentification entre NCE et les équipements/sites.
- **Flux northbound** : HTTPS vers les admins (portail web), port **18002** pour l'API NBI RESTful (documenté dans l'écosystème — POST /controller/v2/tokens), SMTP vers le relais mail (alertes), éventuellement SMS via passerelle.
- **Proxy Internet** : si NCE doit joindre les services Huawei (licences, mises à jour), prévoir le proxy avec authentification.
- **Pare-feu** : documenter **chaque flux** dans un tableau (source/destination/port/usage) et le faire valider par l'équipe sécurité — c'est un livrable du projet.

## 38. Prérequis d'installation — checklist complète

- [ ] Serveurs/VM conformes au dimensionnement validé (CPU/RAM/disque/réseau)
- [ ] OS/image Huawei disponible (ISO/OVA) + checksum vérifié
- [ ] Hyperviseur en version supportée
- [ ] Adressage IP, DNS, NTP, passerelle validés
- [ ] Flux réseau ouverts et testés (tableau des flux signé)
- [ ] Certificats (si PKI interne) prêts ou procédure d'auto-signé définie
- [ ] Compte de service / accès hyperviseur disponibles
- [ ] Fenêtre de maintenance planifiée et communiquée
- [ ] Sauvegarde de l'existant (si migration) réalisée
- [ ] Accès au support Huawei/partenaire vérifié (compte, contrat)
- [ ] Plan de rollback écrit (que faire si l'install échoue ?)
- [ ] Équipe formée aux concepts (au minimum : architecture, licence, ZTP)

## 39. Obtenir le package d'installation

- Via le **partenaire Huawei** ou le portail de distribution logicielle (ESDP) avec les droits associés au contrat.
- Vérifier : **référence exacte de version** (ex : V300R024C00 — ne pas mélanger les versions entre doc et package), **checksum** du fichier, **notes de version** (prérequis, incompatibilités connues, correctifs).
- Télécharger aussi : guide d'installation, guide de déploiement HA, matrice de compatibilité des équipements, License Usage Guide de la version.
- Stocker le package sur un support **fiable et sauvegardé** (pas sur le poste d'un technicien).

## 40. Installation pas à pas — grandes étapes

(Le détail écran par écran dépend de la version — suivre le guide d'installation officiel. Ci-dessous le déroulé logique.)

