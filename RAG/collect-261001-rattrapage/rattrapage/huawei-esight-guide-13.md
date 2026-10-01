---
id: collect-261001-rattrapage/rattrapage/huawei-esight-guide-13
title: "Huawei eSight — Guide ultra-complet d'exploitation terrain"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_esight_guide.md
source_anchor: ""
source_lines: [1970, 2111]
sha256: 5b181a28ab8f160c4b332d954fce411f08788336e927b6e278d2ddfd8f7573c0
---

# Huawei eSight — Guide ultra-complet d'exploitation terrain

- [ ] Environmental Health Check : tout au vert ?
- [ ] Test réel des notifications (SMS + mail) : reçus ?
- [ ] Revue des règles de masquage : encore justifiés ? dates de fin ?
- [ ] Alarmes les plus fréquentes du mois : top 10 extrait, actions ?
- [ ] Comptes : mouvements (arrivées/départs) pris en compte ?
- [ ] Espace disque et croissance de la base : tendance OK ?
- [ ] Rapports programmés : bien générés et envoyés ?
- [ ] Documentation : changements du mois consignés ?

## 104. Mises à jour et patchs : la méthode sans stress

1. **Lire les notes de version** : prérequis, incompatibilités,
   changements de comportement (surtout sur les API northbound et les
   MIB).
2. **Vérifier la compatibilité** avec vos modules et votre version de base.
3. **Maquette d'abord** : rejouer la mise à jour sur un clone/lab
   avant la production.
4. **Fenêtre de maintenance** : informer (la supervision sera
   indisponible — prévoir une surveillance manuelle minimale pendant
   l'opération).
5. **Snapshot + sauvegarde** juste avant (section 70-71).
6. **Appliquer**, vérifier les services, rejouer les tests (login,
   découverte test, alarme test, notification test).
7. **Garder le snapshot** quelques jours avant de le consolider.
8. Documenter : version avant/après, date, incidents rencontrés.

**Fréquence** : patchs de sécurité sans tarder (un NMS exposé avec des
accès SNMP/SSH sur tout le parc est une cible de choix) ; montées de
version fonctionnelles 1-2 fois par an, jamais « parce qu'une nouvelle
version existe ».

## 105. Superviser eSight lui-même

Paradoxe classique : l'outil qui supervise tout n'est supervisé par
personne. Corrigez ça :

- **Supervision externe** : un outil tiers (ou un second eSight léger,
  ou même un simple script) qui ping eSight, teste le port HTTPS et
  vérifie la fraîcheur des données (dernière alarme reçue il y a
  moins de X minutes).
- **Alerte sur les sauvegardes** : échec de backup = alerte immédiate.
- **Espace disque, CPU, RAM** du serveur eSight dans votre supervision
  système habituelle.
- **Canal d'alerte indépendant** : si eSight tombe, c'est un autre
  système qui prévient l'astreinte (sinon : silence radio pendant
  la panne — le pire scénario).

## 106. Le Fault Information Collection Tool : préparer le support

Quand un problème dépasse vos compétences, le support Huawei aura besoin
d'un **package de diagnostic**. eSight fournit le **Fault Information
Collection Tool** (icône eSight Console → Tools) :

1. Le lancer, renseigner les paramètres demandés, cliquer Collect.
2. Récupérer le package généré (noter le **mot de passe de décompression**
   affiché — sans lui, le support ne peut pas l'ouvrir).
3. Transmettre au support avec : version eSight exacte, description du
   problème, heure de survenue, ce qui a déjà été tenté.

**Réflexe** : générer ce package **avant** de bidouiller quand un
problème sérieux survient — les logs d'origine valent de l'or pour
le diagnostic.

---

# 17. POUR ALLER PLUS LOIN

## 107. Documentation officielle : où chercher

- **Portail support Huawei** (support.huawei.com) : guides d'installation,
  d'exploitation et notes de version **de votre version exacte** —
  c'est LA référence, ce guide ne la remplace pas.
- **Brochure eSight 23.1** (mars 2024) : vision produit et périmètre
  des composants.
- **Datasheet eSight** : dimensionnement et éditions.
- **Aide en ligne** intégrée à la console eSight : souvent plus à jour
  que les PDF pour les chemins de menus.
- Mention obligatoire : toute procédure sensible (cluster, restauration,
  upgrade) se fait **avec la doc de votre version sous les yeux**.

## 108. Formation et certification

- Parcours officiels Huawei (HCIA/HCIP) : utiles pour les fondamentaux
  réseau, moins pour eSight spécifiquement — la formation eSight se fait
  surtout via les **partenaires** (formations produit).
- **Formation interne** : le meilleur investissement. Programme type
  pour un nouvel exploitant : 1/2 journée « lire eSight » (topo, alarmes,
  acquittement), 1 journée « administrer » (découverte, seuils, rapports),
  compagnonnage d'un mois avec un senior.
- **Exercices** : alarmes simulées en labo, tempête simulée, restauration
  à blanc — on apprend eSight en le cassant (en labo).

## 109. Communauté et retours d'expérience

- Échanges entre exploitants (forums, groupes d'utilisateurs) : les
  problèmes eSight sont souvent les mêmes partout (découverte SNMP,
  base qui grossit, storms) — les solutions aussi.
- Capitaliser en interne : **journal d'expérience** des alarmes
  (section 40), runbooks (section 98), et ce guide annoté de vos
  spécificités locales.
- Rester attentif à la trajectoire **iMaster NCE** : suivre les
  annonces Huawei via votre partenaire pour anticiper la migration
  (section 5).

---

# 18. PENSE-BÊTE DE POCHE

## 110. Mémo protocoles et ports

| Protocole | Port | Sens | À retenir |
|---|---|---|---|
| SNMP poll | UDP 161 | eSight → équipement | Timeouts ? Vérifier ACL + routage |
| SNMP traps | UDP 162 | Équipement → eSight | Pas de traps = pas d'alarmes temps réel |
| Syslog | UDP 514 | Équipement → eSight | Complète les traps |
| NetStream | UDP (configuré) | Équipement → eSight | NTA, volumineux |
| SSH | TCP 22 | eSight → équipement | Backup de configs |
| Telnet | TCP 23 | eSight → équipement | À éviter si possible (clair) |
| HTTPS console | TCP (selon version) | Navigateur → eSight | **À vérifier sur la documentation officielle** |
| SMTP | TCP 25/587 | eSight → relais | Notif mail |
| ICMP | — | eSight → équipement | Test de joignabilité |

## 111. Mémo cycle de vie des alarmes

**Courante → (acquittement) → Acquittée → (clear) → Historique.**
Règles : on acquitte avec un commentaire/ticket ; on ne masque qu'avec
une date de fin ; le storm se traite par la cause racine, pas en
acquittant tout.

## 112. Mémo sauvegardes

- **eSight lui-même** : base quotidienne + système hebdo → stockage externe.
- **Configs équipements** : quotidienne + à chaque changement → via eSight.
- **Tester la restauration** 1×/an. Une sauvegarde non testée n'existe pas.
- **Licences** archivées dans le PRA.

---

# 19. LES 20 PIÈGES CLASSIQUES

## 113. Les 20 erreurs à ne pas commettre

