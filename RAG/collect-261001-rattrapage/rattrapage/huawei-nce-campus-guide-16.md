---
id: collect-261001-rattrapage/rattrapage/huawei-nce-campus-guide-16
title: "Guide technique ultra-complet — iMaster NCE-Campus (Huawei)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["incident", "license"]
source: docs/RAG/collect-261001-rattrapage/huawei_nce_campus_guide.md
source_anchor: ""
source_lines: [1382, 1500]
sha256: 5aa1bbbea5a651f153a7a3f2dcd67f3768177d75bab7a931d780c966b011b1db
---

# Guide technique ultra-complet — iMaster NCE-Campus (Huawei)

- **Gestion des fichiers** : dépôt central (firmwares, patches, licences) — Maintenance > File Management.
- **Upgrades/downgrades** : politiques et tâches (section 85).
- **Licences des équipements** : dépôt puis activation en masse (Maintenance > Device Maintenance > Activate Device License) — sans conflit avec l'activation directe sur l'équipement.
- **Signature database** : mise à jour des bases de signatures (sécurité) des équipements via le centre de mise à jour.
- **Redémarrage** : redémarrage planifié d'équipements (avec vérification post-redémarrage).
- **Diagnostic** : ping/traceroute depuis un équipement, collecte de logs, captures.

C'est le « couteau suisse » de l'exploitation quotidienne — former l'équipe à ces écrans en priorité.

---

# PARTIE 11 — SÉCURITÉ

## 101. Comptes et rôles — modèle RBAC

Le **RBAC** (Role-Based Access Control) : on n'attribue pas des droits aux personnes, on attribue des **rôles** aux personnes, et des **permissions** aux rôles.

Principes sous NCE :

- **Moindre privilège** : chacun n'a que les droits nécessaires à sa fonction (un technicien de site ne reconfigure pas le cœur).
- **Séparation** : lecture (supervision) vs écriture (configuration) vs administration (comptes, licences).
- **Périmètre** : les rôles se combinent avec les **sites/régions** (un admin local ne voit que ses sites).
- **Traçabilité** : chaque action est loguée avec le compte nominatif (pas de compte générique partagé).

## 102. Rôles prédéfinis et rôles personnalisés

Structure type (à adapter à l'organisation — les intitulés exacts dépendent de la version) :

| Rôle | Périmètre typique | Droits typiques |
|---|---|---|
| Super-admin | Tout | Tous (réservé, 1-2 personnes) |
| Admin réseau | Tous sites | Config, templates, maintenance |
| Opérateur NOC | Tous sites | Lecture, acquittement d'alarmes, diagnostic |
| Admin site | Son site | Config locale, onboarding, maintenance |
| Lecteur / Audit | Selon besoin | Lecture seule, rapports |
| Compte API | Selon besoin | API uniquement, droits minimaux |

Créer des **rôles personnalisés** quand les prédéfinis ne collent pas (ex : « technicien astreinte » = lecture + redémarrage AP, rien d'autre). Revoir les rôles **annuellement** (les gens changent de poste, les droits restent — c'est comme ça que naissent les incidents).

## 103. Authentification des administrateurs — locale, RADIUS, LDAP/AD

- **Locale** : comptes dans NCE — simple, mais à gérer (rotation des mots de passe, départs).
- **RADIUS** : centralise l'authentification des admins (même RADIUS que le 802.1X ou dédié).
- **LDAP/AD** : rattachement à l'annuaire d'entreprise — **recommandé** (un départ = une désactivation AD = plus d'accès NCE ; groupes AD → rôles NCE).
- **Hybride recommandé** : AD pour le quotidien + **compte local de secours** (break-glass) pour les situations où l'AD est injoignable (documenté, mot de passe sous enveloppe, usage audité).
- **MFA** : si NCE le supporte dans la version, l'activer pour les comptes privilégiés (**à vérifier**).

## 104. Journal d'audit — traçabilité des actions

Le journal d'audit enregistre **qui a fait quoi, quand, sur quoi** : connexions, modifications de configuration, déploiements de templates, upgrades, changements de comptes, acquittements d'alarmes.

Exigences :

- **Intégrité** : les logs d'audit ne doivent pas être modifiables par les administrés (droits séparés, voire export vers un SIEM).
- **Rétention** : conforme aux obligations (souvent 1 an minimum — **à vérifier** localement).
- **Revue** : revue périodique (au moins trimestrielle) des actions sensibles (qui a touché aux comptes ? aux templates ?).
- **Alertes** : alerter sur les actions critiques (création de compte admin, modification de template de production, désactivation de l'audit — cette dernière ne devrait jamais arriver).

## 105. Durcissement du contrôleur — checklist

- [ ] Mots de passe par défaut changés, politique forte appliquée
- [ ] Comptes inutiles désactivés/supprimés
- [ ] Services/ports non nécessaires fermés sur le serveur
- [ ] Accès au portail restreint (VPN, IP allowlist, pas d'exposition Internet directe)
- [ ] API NBI protégée (TLS, allowlist, compte dédié)
- [ ] Certificats valides (pas d'auto-signé en production si PKI disponible)
- [ ] NTP sécurisé (authentification NTP si possible)
- [ ] Logs d'audit exportés vers SIEM/syslog sécurisé
- [ ] Correctifs NCE appliqués (politique de patch)
- [ ] Sauvegardes chiffrées et testées
- [ ] Revue des rôles et des comptes (trimestrielle)
- [ ] Test d'intrusion ou audit de configuration annuel

## 106. Certificats — gestion et renouvellement

Inventaire (section 44) + processus :

1. **Registre** : liste de tous les certificats (portail, API, EAP, portail captif) avec émetteur, expiration, responsable.
2. **Alertes** : J-90, J-60, J-30, J-7 (plusieurs niveaux — un seul rappel se perd).
3. **Procédure** : renouvellement documenté pas à pas (génération CSR, validation, déploiement, vérification des services, rollback si échec).
4. **Test** : renouveler une fois en maquette avant la première expiration en production.
5. **EAP-TLS** : le renouvellement des certificats clients est le plus délicat (parc de postes) — automatiser via la PKI/GPO.

Un certificat expiré = portail inaccessible, EAP en échec, API rejetée (voir cas pratique 15). C'est l'incident le plus bête et le plus fréquent.

## 107. Sauvegardes NCE — stratégie

- **Quoi** : bases de données (inventaire, configs, alarmes), fichiers (firmwares, templates, licences, rapports), configuration système.
- **Fréquence** : quotidienne (incrémentale/différentielle selon l'outil) + hebdomadaire complète — à adapter au rythme des changements.
- **Où** : **hors du serveur NCE** (stockage dédié, idéalement hors site), chiffré.
- **Rétention** : ex : 30 jours glissants + 12 mensuelles (à valider avec les obligations).
- **Automatisation** : planifiée dans NCE (si la fonction existe) ou par script — **jamais manuelle** (l'humain oublie).
- **Supervision** : alerter sur **échec** de sauvegarde (une sauvegarde qui échoue en silence = pas de sauvegarde).

## 108. Restauration — procédure et tests

- **Procédure écrite** : pas à pas, avec les prérequis (version NCE identique, licence, réseau), les commandes, les vérifications.
- **Test annuel** (minimum) : restaurer en maquette, vérifier (inventaire, un onboarding test, une alarme, l'API). **Chronométrer** : le RTO (temps de restauration) doit être connu, pas deviné.
- **Scénarios** : restauration partielle (une base corrompue) vs totale (serveur perdu) — les deux doivent être documentés.
- **Après incident réel** : revue (qu'est-ce qui a manqué ?), mise à jour de la procédure.

Voir cas pratique 16 (restauration après sinistre).

## 109. Sécurité des flux southbound

- **Chiffrement** : NETCONF sur SSH/TLS, SNMPv3 (authPriv), HTTPS partout — bannir SNMP v1/v2c et HTTP en production.
- **Segmentation** : VLAN de management dédié, ACL sur les équipements (seul NCE — et les admins — joignent les interfaces de management).
- **Authentification mutuelle** : certificats/clés pour les canaux sensibles (télémétrie, authentification).
- **Surveillance** : détecter les tentatives d'accès au management depuis des zones non autorisées (alarme SIEM).
- **ZTP sécurisé** : liste blanche d'ESN (un équipement inconnu ne doit pas s'enregistrer tout seul en production).

---

# PARTIE 12 — MIGRATION ESIGHT → NCE

## 110. Stratégie de migration eSight → NCE — vue d'ensemble

Il n'existe pas de « bouton migrer » : passer d'eSight à NCE-Campus, c'est **reconstruire** la gestion du périmètre campus Huawei sous un nouveau paradigme. La stratégie recommandée :

