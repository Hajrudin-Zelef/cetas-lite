---
id: collect-261001-rattrapage/rattrapage/huawei-usg6000-guide-14
title: "Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["license", "memory"]
source: docs/RAG/collect-261001-rattrapage/huawei_usg6000_guide.md
source_anchor: ""
source_lines: [2062, 2213]
sha256: 556217d27afabcdcccd78cfa16afb465f2a09f119fbe7b2a66af192e939326d4
---

# Guide ULTRA-COMPLET — Huawei USG6000 (Firewall UTM)

Les packs de signatures (AV/IPS/URL) se téléchargent depuis le centre de mise à jour. En cas de restauration :
1. Restaurer la config (section 116).
2. Vérifier les versions de signatures (`display av update info`).
3. Si obsolètes : relancer une MAJ (section 76) — la config référence les **profils**, les signatures se rechargent.

## 121. Inventaire : ce qu'il faut archiver pour chaque USG

📋 Fiche par équipement (dans ton GLPI/gestion de parc) :
- [ ] Modèle, numéro de série (ESN), version logicielle.
- [ ] Plan d'adressage (interfaces, zones, VLAN).
- [ ] Schéma réseau (où il est branché, quoi autour).
- [ ] Config sauvegardée (3 générations).
- [ ] Licences : type, expiration.
- [ ] Mots de passe dans le **coffre** (jamais dans la fiche en clair).
- [ ] Contacts support + contrat de maintenance.
- [ ] Historique des interventions.

## 122. Planning annuel type (récapitulatif, détail section 126)

| Fréquence | Action |
|-----------|--------|
| Quotidien | Coup d'œil supervision (CPU, sessions, VPN) — automatisé |
| Hebdo | Rapport UTM, état HA, échecs d'authentification |
| Mensuel | MAJ signatures (si pas auto), revue des logs d'attaque, backup config |
| Trimestriel | Revue des politiques (règles inutilisées), revue des exceptions UTM, test de restauration |
| Semestriel | Test de bascule HA, revue des licences (J-60) |
| Annuel | Upgrade firmware (si besoin), audit de durcissement, test PRA |

## 123. Upgrade en production : la checklist du jour J

📋 Jour J :
- [ ] Fenêtre de maintenance **validée** et **communiquée**.
- [ ] Sauvegarde config **vérifiée** (le fichier s'ouvre, il est complet).
- [ ] Firmware **vérifié** (checksum OK, bonne version pour le modèle).
- [ ] Console **branchée** et testée.
- [ ] Collègue **prévenu** (binôme en cas de pépin).
- [ ] Plan de rollback **écrit** (pas « dans la tête »).
- [ ] Après reboot : trafic, VPN, HA, licences, supervision — **tout vert** avant de clore.

## 124. Ce qu'il ne faut JAMAIS faire avec un firmware

- ❌ Upgrader **sans sauvegarde**.
- ❌ Upgrader **les deux nœuds HA en même temps**.
- ❌ Couper l'alimentation pendant l'écriture flash.
- ❌ Installer un firmware d'un **autre modèle** (même proche).
- ❌ Upgrader un vendredi 17h sans astreinte.
- ❌ Supprimer l'ancien firmware avant **une semaine** de validation.

---
---

# BLOC L — MAINTENANCE PRÉVENTIVE

## 125. Philosophie : le firewall qu'on oublie est celui qui tombe

Un USG bien configuré tourne des mois sans broncher — c'est exactement pour ça qu'on l'oublie, et qu'on découvre la panne (disque plein, licence expirée, ventilateur mort) **le jour où on a besoin des logs**. La maintenance préventive, c'est **rendre l'oubli impossible** : checks automatisés + rituels calendaires.

## 126. Plan de maintenance détaillé

| Fréquence | Tâches | Commandes/outils |
|-----------|--------|------------------|
| **Quotidien** (auto) | CPU, mémoire, sessions, état VPN/HA, traps | SNMP + supervision |
| **Hebdomadaire** | Rapport UTM, échecs auth, état standby HA | Web Monitor, `display hrp state` |
| **Mensuel** | Backup config, MAJ signatures (si manuel), espace disque, revue logs d'attaque | Sections 115, 76, `display disk` |
| **Trimestriel** | Règles inutilisées, exceptions UTM, test restauration (maquette), poussière/ventilation | `display security-policy`, section 116 |
| **Semestriel** | Test bascule HA, licences J-60, firmware (évaluer) | Section 101, `display license` |
| **Annuel** | Audit durcissement, test PRA complet, remplacement préventif (ventilos 4 ans, disques 3-5 ans) | Section 133+ |

## 127. Checklist mensuelle (imprimable)

📋 Maintenance mensuelle USG6000 — Site : _______ Date : _______
- [ ] `display cpu` / `display memory` — valeurs normales (< 70 % en pointe) ?
- [ ] `display firewall session table` — pas d'anomalie (sessions fantômes) ?
- [ ] `display version` — toujours la version attendue ?
- [ ] Signatures AV/IPS/URL à jour (< 7 jours) ?
- [ ] `display license` — aucune expiration < 90 jours ?
- [ ] HA : `display hrp state verbose` — standby synchronisé ?
- [ ] VPN : `display ike sa` — tunnels attendus montés ?
- [ ] Disque : espace libre > 20 % ?
- [ ] Backup config effectué et **vérifié** ?
- [ ] Logs : pas d'attaque massive non traitée ?
- [ ] Baie : voyants verts, pas de bruit anormal, température OK ?

## 128. Nettoyage des politiques : la règle des 90 jours

```huawei
[USG] display security-policy rule all   # repérer les règles à 0 hit
```

- Toute règle avec **0 hit depuis 90 jours** = candidate à la **désactivation** (pas suppression immédiate : désactiver 30 jours, puis supprimer si personne ne pleure).
- 📋 Documenter chaque suppression (date, raison, demandeur). Une règle « inutile » qui servait au PRA annuel, ça fait désordre de la réécrire en urgence.

## 129. Santé matérielle : ventilateurs, alimentations, disques

```huawei
[USG] display device            # état des modules
[USG] display fan               # ventilateurs (selon version)
[USG] display power             # alimentations (selon version)
[USG] display disk              # disques et espace (selon version)
[USG] display temperature       # températures (selon version)
[USG] display alarm              # alarmes actives
```

Seuils d'alerte terrain : ventilateur en défaut = **remplacement sous 48h** (un seul ventilo en moins = surchauffe en été) ; disque en erreur = **remplacement + resync RAID1** ; alim redondante en défaut = **commander le spare immédiatement** (tu n'as plus de redondance).

## 130. Dépoussiérage et environnement

- **Trimestriel** : dépoussiérage des façades/filtres (air comprimé sec, boîtier **éteint** si ouverture nécessaire — sinon par l'extérieur).
- **Température de baie** : 20–25 °C idéal, alarme au-delà de 35 °C.
- ⚠️ **Ne jamais** poser de carton/dossier devant les grilles d'aération. Vu en vrai. Plusieurs fois.

## 131. Revue des logs d'attaque : le rituel hebdo

1. Ouvrir le rapport hebdo (section 111) ou les logs UTM.
2. Chercher : **nouveaux top talkers**, **pics d'attaques**, **pays inhabituels**, **tentatives sur les comptes admin**.
3. Toute anomalie → **ticket/enquête** (pas « on verra la semaine prochaine »).
4. 📋 Tenir un **journal des incidents** (même les petits) — c'est lui qui justifie les budgets sécu.

## 132. Plan de Reprise d'Activité (PRA) : le firewall dedans

- Le PRA inclut : **config sauvegardée hors site**, **firmware de référence** archivé, **licences** (fichiers .dat), **procédure de reconstruction** (ce guide, sections 11–18), **boîtier de secours** (si contrat) ou **délai fournisseur**.
- 📋 **Tester 1 fois par an** : reconstruire un USG de maquette depuis la sauvegarde. Un backup jamais restauré = un espoir, pas un plan.

---
---

# BLOC M — DURCISSEMENT

## 133. Durcir l'accès management

```huawei
system-view
# 1. HTTPS uniquement (désactiver HTTP) :
[USG] web-manager enable port 8443            # port non standard (optionnel)
# Désactiver le HTTP clair via la config des services de management par interface
# 2. SSH uniquement (désactiver Telnet) :
[USG] undo telnet server enable
[USG] ssh server enable
[USG] ssh server port 2222                    # port non standard (optionnel, à documenter !)
# 3. Restreindre les sources d'administration :
[USG] acl 2000
[USG-acl-basic-2000] rule permit source 192.168.10.0 0.0.0.255   # VLAN admin uniquement
[USG-acl-basic-2000] quit
[USG] ssh server acl 2000
# 4. Comptes nominatifs (pas de compte générique partagé) - voir section 13
# 5. Timeout de session :
[USG] idle-timeout 10                         # minutes (console/web selon version)
save
```

📋 Règles d'or : **pas de Telnet**, **pas de HTTP**, **pas d'admin depuis Internet** (ou via VPN uniquement), **comptes nominatifs**, **double authentification** si possible.

