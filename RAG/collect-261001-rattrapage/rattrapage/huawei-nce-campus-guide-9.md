---
id: collect-261001-rattrapage/rattrapage/huawei-nce-campus-guide-9
title: "Guide technique ultra-complet — iMaster NCE-Campus (Huawei)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["license"]
source: docs/RAG/collect-261001-rattrapage/huawei_nce_campus_guide.md
source_anchor: ""
source_lines: [671, 776]
sha256: 3ac0588d7d466deb092c8fd4e4d70928a2aaf2cd07f3eb58d21c35699bd5a173
---

# Guide technique ultra-complet — iMaster NCE-Campus (Huawei)

1. **Préparer la VM/le serveur** : créer la VM selon les specs, monter l'ISO/OVA, configurer le réseau (IP statique de management).
2. **Installer le système** : lancer l'installation, choisir le mode (standalone / HA — décider **avant**), définir les mots de passe initiaux (politique forte dès le départ).
3. **Premier démarrage** : vérifier les services (tous les composants manager/controller/analyzer démarrés), consulter les logs d'initialisation.
4. **Accéder au portail** : https://<IP-NCE> — accepter/vérifier le certificat.
5. **Assistant de configuration initiale** : fuseau horaire, NTP, DNS, nom d'hôte, paramètres réseau complémentaires.
6. **Importer la licence** : Maintenance > File Management (dépôt du fichier), puis activation (Device Maintenance > Activate Device License — logique documentée dans le guide Monitoring and O&M). Vérifier le compteur de device-days.
7. **Configurer les services tiers** : relais SMTP (alertes), serveur de fichiers (firmwares), centre de mise à jour / signature database (System > System Management > Third-Party Service).
8. **Créer l'organisation** : tenant, régions, sites (section 43).
9. **Sauvegarde initiale** : snapshot VM + backup applicatif (section 45/107).
10. **Test de bout en bout** : onboarder **un** équipement pilote (un S310 de maquette) et vérifier supervision + déploiement de template avant d'industrialiser.

## 41. Premier login — comptes par défaut et politique

- Utiliser les **comptes initiaux** documentés dans le guide d'installation (généralement un compte `admin` système).
- **Dès le premier login** : changer tous les mots de passe par défaut, appliquer la politique de complexité (longueur, rotation), créer les comptes nominatifs des administrateurs (jamais de compte partagé en production).
- Activer le **verrouillage après échecs** et la **déconnexion automatique** d'inactivité.
- Tracer : vérifier que les connexions alimentent le journal d'audit (section 104).
- Si annuaire d'entreprise : planifier le raccordement LDAP/AD (section 103) — mais garder **un compte local de secours** (break-glass) avec mot de passe sous enveloppe scellée.

## 42. Configuration initiale — assistant de démarrage

L'assistant post-installation couvre typiquement :

- **Système** : nom d'hôte, fuseau horaire, langue du portail.
- **Temps** : serveurs NTP (redondants), vérification de la dérive.
- **Réseau** : DNS, passerelle, éventuellement proxy.
- **Sécurité** : politique de mot de passe, certificats initiaux.
- **Organisation** : création du premier tenant/site (détaillée section 43).
- **Services** : SMTP, serveur de fichiers, mise à jour.

Conseil : faire cette configuration **en maquette d'abord**, documenter chaque écran (captures), puis rejouer en production — c'est le premier « template » de votre exploitation NCE.

## 43. Création de l'organisation : tenant, régions, sites

La hiérarchie logique type dans NCE-Campus :

```
Tenant (entreprise / organisation)
└── Région (géographique : ex. « Siège », « Ouest », « Est »)
    └── Site (bâtiment / agence : ex. « Siège — Bâtiment A »)
        └── Équipements (switches, AP, AR, USG, WAC)
```

- **Tenant** : entité de gestion (mono-tenant en entreprise classique ; multi-tenant en MSP).
- **Région** : regroupement géographique/administratif pour les rapports et les droits.
- **Site** : l'unité opérationnelle — c'est au niveau site que s'appliquent les templates, les politiques d'accès, les fenêtres de maintenance, les composants d'authentification.

Bonnes pratiques : nommer les sites de façon **stable et explicite** (code site + ville + bâtiment), ne jamais réutiliser un nom de site supprimé sans vérifier les restes en base, documenter la convention de nommage dans le dossier d'exploitation.

## 44. Intégration DNS / NTP / certificats

- **DNS** : NCE doit résoudre les noms des équipements/serveurs (et inversement pour les logs). Enregistrements A/PTR propres, TTL raisonnables. Tester avec `nslookup`/`dig` depuis NCE avant l'onboarding.
- **NTP** : hiérarchie claire (source stratum 1/2 → NCE → équipements). Vérifier l'offset (< 1 s idéalement). Un NTP en panne = des certificats qui semblent expirés, des tokens API rejetés, des logs incohérents entre NCE et équipements.
- **Certificats** : inventaire des certificats NCE (portail HTTPS, API, 802.1X/EAP, portail captif). Noter les **dates d'expiration** dans un calendrier avec alerte à J-60/J-30 (voir cas pratique 15). Privilégier la PKI d'entreprise pour l'EAP-TLS ; à défaut, planifier le renouvellement des auto-signés.

## 45. Sauvegarde initiale — avant toute chose

Avant d'onboarder le moindre équipement de production :

1. **Snapshot VM** (si virtualisé) après installation + config initiale validée.
2. **Backup applicatif NCE** (base + fichiers) via la fonction prévue — tester la restauration **en maquette** (un backup non testé n'est pas un backup).
3. Documenter : quoi, où, fréquence, rétention, responsable, procédure de restauration (section 107-108).

## 46. Checklist post-installation

- [ ] Tous les services NCE démarrés et stables (24 h d'observation)
- [ ] Licence importée, compteur device-days cohérent
- [ ] NTP/DNS vérifiés (offset, résolution)
- [ ] Organisation créée (tenant/région/site pilote)
- [ ] SMTP testé (mail d'alerte reçu)
- [ ] Serveur de fichiers joignable (dépôt firmware)
- [ ] Comptes admin nominatifs créés, mots de passe par défaut changés
- [ ] Audit activé et vérifié
- [ ] Sauvegarde initiale réalisée **et restauration testée**
- [ ] Un équipement pilote onboardé et supervisé
- [ ] Documentation d'installation archivée (versions, captures, flux réseau)

## 47. Échecs d'installation fréquents et remèdes

| Symptôme | Cause probable | Remède |
|---|---|---|
| L'ISO ne boote pas | Image corrompue / mauvais mode (UEFI/BIOS) | Vérifier le checksum, re-télécharger, tester l'autre mode |
| Services qui ne démarrent pas | RAM/disque insuffisants | Re-dimensionner la VM, vérifier les prérequis |
| Portail inaccessible | Flux HTTPS bloqué / certificat | Vérifier pare-feu, écoute du service, logs |
| Licence refusée | Fichier corrompu / mauvaise version / ESN | Re-télécharger depuis l'ESDP, vérifier la compatibilité de version |
| NTP qui dérive | Pas de connectivité UDP/123 | Ouvrir le flux, vérifier la hiérarchie |
| Lenteurs dès le départ | Sur-allocation hyperviseur | Réservations CPU/RAM, datastore dédié |

## 48. Mise à jour du contrôleur lui-même

NCE-Campus évolue par versions (V300R020 → R022 → R024...) et correctifs :

- **Politique** : rester à N-1 maximum (ni la toute dernière le jour de sa sortie, ni une version sans support). Lire les notes de version (breaking changes, prérequis).
- **Procédure** : sauvegarde complète avant, snapshot VM, fenêtre de maintenance, test en maquette d'abord, vérification post-upgrade (services, licences, onboarding d'un équipement test, API).
- **Rollback** : snapshot + backup = retour arrière possible (section 86/108).
- **Signature database** : mettre à jour la base de signatures via le centre de mise à jour configuré (System > System Management > Third-Party Service > Signature Database Server), périodiquement ou en temps réel selon la politique.

## 49. Notion de site dans NCE-Campus

Le **site** est l'unité de déploiement : il regroupe les équipements d'un lieu, porte les templates, les politiques (VLAN, SSID, 802.1X), les fenêtres de maintenance et le composant d'authentification éventuel. Bien découper les sites, c'est :

- Isoler les **blast radius** (un template erroné n'impacte qu'un site — voir cas pratique 6).
- Permettre des **fenêtres de maintenance** par site.
- Déléguer des droits par site (RBAC).
- Obtenir des **rapports** par site.

