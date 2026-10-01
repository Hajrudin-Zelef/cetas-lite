---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-11
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["attention", "datacenter"]
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [951, 1051]
sha256: baac5658c7fa5e9f323d1fac6353efa603dfcd963dc5a51db521e1526852025c
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

Un rapport mensuel d'une page avec ces KPI, c'est ce qui transforme « le gars du Wi-Fi » en « notre prestataire infra ». Et ça justifie le contrat de maintenance.

## 81. Quand le client grandit : les signaux d'alerte

Surveille ces signaux — ils annoncent qu'eKit atteint ses limites :
- plus de **30-40 AP** sur un site unique avec besoin de roaming fin ;
- besoin de **802.1X / NAC** (authentification par certificats, contrôle d'accès dynamique) ;
- besoin de **routage dynamique** (OSPF/BGP), de **VPN maillés** complexes, de **SD-WAN** avancé ;
- **plusieurs Gbit/s** de trafic inter-VLAN à router en permanence ;
- exigences de **redondance** (double passerelle, stacking, double alimentation) contractuelles ;
- équipe IT interne qui veut du **CLI/API/automatisation** (Ansible, Terraform) ;
- secteurs réglementés exigeant des **journaux centralisés** et de l'audit fin.
Dès 2-3 signaux : ouvre le dossier « migration Datacom » (sections 83-90), ne bricole pas.

## 82. Ce que la gamme Datacom entreprise apporte (pour situer)

Sans entrer dans le détail (tu as tes propres ressources Datacom), la gamme entreprise Huawei apporte :
- **iMaster NCE** : le vrai contrôleur campusc/datacenter (politiques, automatisation, assurance réseau) ;
- switchs **S600 et supérieurs**, **CloudEngine** : 25/40/100G, stacking, redondance ;
- **HiSecEngine USG6700F** : pare-feux haut de gamme (100+ Gbit/s) ;
- **AirEngine** (AP entreprise) + **WAC** dédiés : roaming fin, 802.1X, IoT ;
- routage dynamique complet (OSPF, BGP, MPLS), VXLAN/EVPN.
Le prix et la complexité vont avec : c'est un autre métier (ou un partenaire à trouver).

## 83. Critères de décision : rester eKit ou passer Datacom — la grille

| Critère | Reste eKit | Passe Datacom |
|---|---|---|
| Utilisateurs / site | < ~200 | > 300-500 |
| AP / site | < 30 | > 50 ou roaming critique |
| Sites | Multi-sites simples | Topologie hub-and-spoke complexe, SD-WAN |
| Authentification | PSK / portail captif | 802.1X, certificats, NAC |
| Routage | Statique, 1-2 uplinks | OSPF/BGP, multi-opérateurs avec bascule fine |
| Redondance | Onduleur + 4G de secours suffisent | Double équipement exigé contractuellement |
| Admin | App/cloud, pas d'expert réseau | Équipe réseau avec compétences CLI/API |
| Budget | Serré, OPEX minimal | Budget d'infrastructure assumé |

**Phrase à dire au client :** « Aujourd'hui eKit couvre votre besoin avec de la marge. Si vous dépassez [seuil chiffré], on basculera telle brique en Datacom — c'est prévu, pas une surprise. »

## 84. Migration eKit → Datacom : anticiper sans tout jeter

Bonne nouvelle : une migration n'est pas un big-bang.
- Les **AP eKit en mode Fit** peuvent théoriquement rejoindre un WAC (à valider par modèle — **à vérifier sur la documentation officielle** : tous les modèles ne basculent pas).
- Le **câblage** (Cat6, baie, chemins) est réutilisable à 100 %.
- Le **plan d'adressage et les VLAN** se transposent tels quels.
- La **passerelle AR eKit** devient un routeur d'agence ou est remplacée par un AR entreprise / USG supérieur.
- Stratégie : migrer **par couche** (d'abord le cœur : passerelle + switchs d'agrégation, puis les AP par zone), en gardant le cloud eKit pour les sites restants en attendant leur tour.

## 85. Le piège du « on mettra du Datacom plus tard » sans plan

Si tu sais dès le départ que le client dépassera eKit dans 18 mois :
- câble en **Cat6A** (pas Cat5e), prévois des **fourreaux** en réserve ;
- prends des switchs avec des **uplinks 10G** (S310-48P4X plutôt que 48P4S) même si inutilisés aujourd'hui ;
- documente un **plan d'adressage extensible** (/23 plutôt que /24) ;
- écris la trajectoire dans la proposition commerciale (« phase 1 eKit, phase 2 Datacom si > X utilisateurs »).
Le surcoût initial est faible ; le coût d'un recâblage dans 18 mois est énorme.

## 86. Coexistence eKit + Datacom : c'est possible

Un groupe peut très bien avoir : le siège en Datacom (NCE + S600) et 15 agences en eKit supervisées via le cloud. C'est même une architecture **recommandée** économiquement : du lourd où c'est complexe, du simple où c'est simple. Points d'attention : interconnexion via VPN IPsec ou lignes privées, adressage cohérent global, et **deux outils de supervision** (NCE + cloud eKit) — à intégrer dans ton NMS central via SNMP quand c'est possible.

## 87. L'argument commercial honnête face aux concurrents

Face à un concurrent qui propose du « tout-cloud avec licence annuelle » :
- eKit : gestion annoncée **sans licence**, matériel souvent moins cher, déploiement rapide par des non-experts.
- Limites à assumer : écosystème fermé, fonctions avancées limitées, dépendance au cloud Huawei.
Face à du Datacom (Huawei ou autre) :
- eKit : 2-3× moins cher à l'achat, déploiement en jours pas en semaines, maintenance sans expert réseau.
- Ne joue jamais eKit contre du Datacom sur un besoin Datacom : tu perdras le client à la première panne grave. Qualifie le besoin d'abord (section 15).

## 88. Garantie et RMA : ce qu'il faut savoir

- Les distributeurs annoncent typiquement **3 ans de garantie** sur les switchs eKit (ex. : S310 chez les distributeurs cités) — **à vérifier** pour chaque modèle auprès de ton distributeur (la durée peut varier : AP, AR, USG).
- **Procédure RMA** : via ton distributeur Gold (pas directement Huawei en général). Garde : facture d'achat, SN, description de panne, photos.
- **Stock tampon** : pour un client multi-sites, négocie 1 AP et 1 switch de secours chez toi (ou chez le client). Un AP en RMA qui met 2 semaines, c'est 2 semaines de zone blanche.
- **Exclusions typiques** : surtension/foudre (d'où la section 72), mauvaise installation, firmware non officiel. Documente tes installations (photos, parafoudre) pour ne jamais être en tort.

## 89. Sauvegarde des configurations : la méthode complète

1. **Cloud** : la config des sites est conservée dans le cloud eKit (principe du pilotage centralisé) — c'est ta sauvegarde n°1, mais pas la seule.
2. **Export local** : exporte la configuration de chaque site après chaque changement majeur (fichier daté : `CLIENT-Site_2026-09-27.cfg` — nom fictif d'exemple). Stockage : dossier du site + copie hors site (ton NAS / cloud pro).
3. **Fiche site** (section 65) : la version papier/lisible par un humain de l'essentiel.
4. **Avant chaque changement** : sauvegarde + note (« avant : … / après : … »). En cas de problème, tu reviens en arrière en minutes au lieu de reconstruire de mémoire.
5. **Test de restauration** : une fois par an, restaure une config sur un équipement de labo (ou simule la procédure). Une sauvegarde jamais testée n'est pas une sauvegarde.
6. **Rétention** : garde les N-3 dernières sauvegardes par site.

## 90. Mises à jour firmware : le détail par famille

| Famille | Mécanisme | Via cloud | Points d'attention |
|---|---|---|---|
| Switchs (S220/S310/S620) | **HOUP** (dépôt de mise à jour en ligne Huawei) : le switch trouve son chemin de MAJ, 1 clic, pré-chargement | Oui (app/SNC) | Coupure brève malgré le pré-chargement ; Perpetual PoE = les AP restent alimentés |
| AP | Mise à jour via cloud/app | Oui | Les clients sont déconnectés pendant la MAJ → heures creuses |
| AR (passerelles) | Via cloud/app | Oui | Coupe Internet du site pendant la MAJ → prévenir, heures creuses |
| USG6000F-S | Via interface locale/cloud + signatures (centre de sécurité Huawei) | À vérifier | MAJ signatures = sans coupure en général ; MAJ firmware = coupure → planifier |

**Règle commune :** lire la release note, sauvegarder, site pilote, vagues, fenêtre annoncée (section 42). Et **jamais de mise à jour à distance sans personne joignable sur site** : si ça tourne mal, il faut une main pour le reset physique.

---

## 91. Dépannage : la méthode avant les cas (à afficher dans l'atelier)

