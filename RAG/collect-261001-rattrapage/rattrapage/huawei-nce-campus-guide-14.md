---
id: collect-261001-rattrapage/rattrapage/huawei-nce-campus-guide-14
title: "Guide technique ultra-complet — iMaster NCE-Campus (Huawei)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["arr", "incident", "sandbox"]
source: docs/RAG/collect-261001-rattrapage/huawei_nce_campus_guide.md
source_anchor: ""
source_lines: [1190, 1284]
sha256: 9068b5832b86da2a9466afa26155a1c0cb65357dfdff893aa4565657cc86ff7b
---

# Guide technique ultra-complet — iMaster NCE-Campus (Huawei)

1. **Partir du réel** : prendre la configuration d'un équipement qui fonctionne bien en production, la nettoyer (retirer le spécifique), en faire le template v1.
2. **Identifier les variables** : tout ce qui change par site/équipement devient une variable (section 63). Le reste est fixe (= standard d'entreprise).
3. **Tester en maquette** : appliquer le template sur un équipement de lab, vérifier chaque service (management, 802.1X, SSID, VPN...).
4. **Versionner** : v1.0, v1.1... avec changelog (qui, quoi, pourquoi). Stocker les templates hors NCE aussi (Git) — traçabilité et revue.
5. **Revue par un pair** : un deuxième œil avant chaque version majeure (comme du code).
6. **Déployer par vagues** : site pilote → 10 % → 50 % → 100 %, avec critères de passage (pas d'incident bloquant pendant X jours).
7. **Rollback prêt** : version précédente conservée et testée (section 87).

Anti-patterns : le template « copié de la doc sans test », la variable non documentée, le déploiement « big bang » un vendredi soir.

## 84. Déploiement en masse — vagues et fenêtres de maintenance

Un changement qui touche 100 équipements ne se fait pas en une fois :

- **Vague 0 — maquette** : validation fonctionnelle complète.
- **Vague 1 — pilote** : 1 site non critique (ou quelques équipements), fenêtre de maintenance dédiée, astreinte renforcée.
- **Vague 2 — 10-20 %** : sites variés (pour couvrir les cas particuliers).
- **Vague 3 — généralisation** : le reste, par groupes de sites, avec pause d'analyse entre groupes.
- **Critères de passage** : 0 incident bloquant, conformité 100 %, supervision nominale pendant 48-72 h.

Fenêtres de maintenance : planifiées, **communiquées** (utilisateurs, métiers, astreinte), avec heure de début **et de fin**, et procédure d'arrêt d'urgence (rollback). Un déploiement sans fenêtre annoncée = un incident de communication garanti.

## 85. Mise à jour firmware en masse — procédure

La documentation NCE (Monitoring and O&M V300R020C10) décrit la mécanique : l'administrateur définit une **politique de mise à jour** (manuelle ou automatique) ; à réception de la tâche, l'équipement **télécharge** le package depuis l'adresse indiquée (serveur de fichiers) et s'upgrade.

Procédure recommandée :

1. **Choisir la version cible** : version validée par Huawei pour vos modèles + correctifs de sécurité — jamais « la dernière sortie hier ».
2. **Déposer** le firmware sur le serveur de fichiers NCE (vérifier l'intégrité).
3. **Tester en maquette** : upgrade + downgrade sur chaque modèle, avec test de non-régression des services.
4. **Politique** : manuelle pour la production (l'automatique est réservée aux parcs très homogènes et aux fenêtres maîtrisées).
5. **Vagues** : comme section 84 (pilote → généralisation).
6. **Vérification** : version effective, conformité du template après upgrade (un firmware peut changer des défauts !), supervision nominale.
7. **Fenêtre** : les upgrades redémarrent les équipements — planifier hors production.

## 86. Politique de mise à jour : manuelle vs automatique

| | Manuelle | Automatique |
|---|---|---|
| Déclenchement | L'admin lance, par vague | Planifiée (ex : nuit, week-end) |
| Contrôle | Total | Délégué à la politique |
| Risque | Faible (si vagues respectées) | Moyen (un bug touche tout le monde la même nuit) |
| Charge admin | Élevée | Faible |
| Recommandé pour | Production, sites critiques | Maquette, sites homogènes non critiques |

Règle : **manuelle en production**, automatique éventuellement pour les AP d'un site non critique après plusieurs cycles manuels réussis. Et toujours : **ne jamais** mettre à jour le firmware de tout le parc la veille d'un événement critique (audit, visite officielle...).

## 87. Rollback — revenir en arrière après un déploiement raté

Le rollback, c'est la capacité à **revenir à l'état antérieur** quand un changement tourne mal. Trois niveaux :

1. **Rollback de template** : réappliquer la version précédente du template (conservée et versionnée — section 83). Le plus courant.
2. **Rollback de firmware** : downgrader vers la version précédente (testé en maquette au préalable — un downgrade non testé peut être pire que le bug).
3. **Rollback de contrôleur** : snapshot VM + backup avant upgrade de NCE (section 48).

Conditions d'un rollback qui marche : **détecter vite** (supervision + critères d'échec définis à l'avance : « si > 5 % d'équipements en écart ou 1 site critique impacté, on rollback »), **décider vite** (qui a l'autorité ? — écrit dans la procédure), **exécuter vite** (procédure testée, pas improvisée). Faire un **exercice de rollback** en maquette au moins une fois par an — le jour où on en a besoin, il est trop tard pour apprendre.

## 88. API northbound RESTful — principes

L'API **northbound** (NBI) expose NCE-Campus aux applications : portails, ITSM, supervision tierce, scripts. C'est une API **RESTful** documentée par Huawei (plus de 500 API selon le constructeur — l'écosystème développeur Huawei propose sandbox et outils graphiques).

Principes :

- Ressources (sites, équipements, alarmes, topologie...) manipulées en HTTP (GET/POST/PUT/DELETE), réponses JSON.
- **Ne jamais** exposer l'API sur Internet sans protection (VPN, IP allowlist, TLS).
- Versionner ses scripts (l'API évolue avec les versions NCE — figer la version cible).
- Journaliser les appels sensibles (qui a fait quoi via l'API — traçabilité).

## 89. Authentification API : tokens (POST /controller/v2/tokens)

D'après l'écosystème documenté (NBI V300R024C00) :

- Authentification par **token** : `POST /controller/v2/tokens` avec les identifiants d'un compte du **groupe d'utilisateurs NBI tiers** (third-party NBI User Group) → le contrôleur renvoie un token.
- Le token se transmet dans l'en-tête **`X-ACCESS-TOKEN`** des appels suivants.
- Expiration **glissante** (sliding-expiry) : le token est rafraîchi par l'usage ; prévoir une **reconnexion transparente** sur 401 (non autorisé).
- Les tokens sont **liés à l'IP** (documenté) : un script qui change d'IP source doit se ré-authentifier.

Bonnes pratiques : compte API **dédié** (pas le compte admin humain), droits minimaux (lecture seule si suffisant), secret stocké en coffre (jamais en dur dans un script), rotation périodique.

## 90. Cas d'usage API : topologie, métriques, alarmes

Exemples concrets (logique observée dans l'écosystème, endpoints à valider sur la version cible) :

- **Topologie** : récupérer les nœuds (AP/LSW/AR/FW) et les liens (avec ports d'extrémité) pour afficher la carte dans un outil tiers ou vérifier la cohérence avec l'inventaire.
- **Métriques** : CPU/mémoire/statut par équipement, utilisation des liens — pour alimenter Zabbix/Grafana ou des rapports custom.
- **Alarmes** : lire les alarmes courantes (scroll API, sévérités 1-4 : critique/haute/moyenne/basse) pour créer des **tickets automatiques** dans l'ITSM (GLPI...).
- **Inventaire** : synchroniser l'inventaire NCE avec la CMDB.

C'est par l'API que NCE s'intègre au SI existant au lieu de rester un silo — prévoir ce chantier dans le projet (voir cas pratique 23).

## 91. Intégration avec des outils tiers (Zabbix, GLPI, scripts)

Scénario type pour Zelef (qui connaît Zabbix) :

