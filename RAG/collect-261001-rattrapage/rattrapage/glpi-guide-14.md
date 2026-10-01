---
id: collect-261001-rattrapage/rattrapage/glpi-guide-14
title: "Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "apache", "incident"]
source: docs/RAG/collect-261001-rattrapage/glpi_guide.md
source_anchor: ""
source_lines: [2293, 2438]
sha256: d5239feb096141688bba5c2390ccefe172049ea0ef84d1e4b6a0c521e3649d0b
---

# Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV

```
┌─ GLPI — PENSE-BÊTE ─────────────────────────────────────┐
│                                                          │
│ CRÉER UN TICKET : Assistance > Tickets > +                │
│   → catégorie, copieur (élément associé !), compteur,    │
│     code erreur, description factuelle                    │
│                                                          │
│ CYCLE : Nouveau → En cours → Résolu → Clos               │
│   En attente = TOUJOURS un motif + date de relance       │
│                                                          │
│ UN SUIVI À CHAQUE ÉTAPE (public si info client)          │
│ PHOTO du compteur / du défaut en pièce jointe            │
│                                                          │
│ SORTIE DE PIÈCE : depuis le ticket (stock décrémenté)    │
│                                                          │
│ EN ATTENTE > 7 JOURS → relancer / escalader              │
│ SLA : le compteur tourne en heures ouvrées               │
│                                                          │
│ BASE DE CONNAISSANCES : chercher AVANT d'appeler         │
│   un collègue / de se déplacer                          │
│                                                          │
│ MOTS DE PASSE : jamais dans GLPI → coffre dédié         │
│                                                          │
│ EN PANNE ? → prévenir le superviseur + noter l'heure    │
│                                                          │
└──────────────────────────────────────────────────────────┘
```

---

## 67. Glossaire

| Terme | Définition |
|---|---|
| **API REST** | Interface permettant à des scripts d'échanger avec GLPI (tickets, parc...) |
| **CMDB** | Base de données de configuration : l'inventaire et ses liaisons (le « Parc ») |
| **Collecteur** | Boîte mail lue par GLPI qui transforme les courriels en tickets |
| **Cron** | Tâche planifiée système qui fait tourner les actions automatiques de GLPI |
| **Déclarant (post-only)** | Utilisateur qui crée/suit ses tickets, sans voir le reste |
| **Entité** | Découpage organisationnel (société, agence, client) |
| **FusionInventory** | Plugin + agent d'inventaire automatique (matériel, réseau, SNMP) |
| **Gabarit** | Modèle de ticket définissant champs visibles/obligatoires par catégorie |
| **Impact** | Étendue réelle d'un incident (1-mineur → 4-critique) |
| **ITIL** | Référentiel de bonnes pratiques ITSM |
| **ITSM** | Gestion des services informatiques |
| **Priorité** | Calculée depuis urgence × impact (matrice paramétrable) |
| **Profil** | Ensemble d'habilitations (technicien, superviseur, admin...) |
| **SLA** | Engagement de niveau de service (délais de prise en compte/résolution) |
| **SNMP** | Protocole d'interrogation des équipements réseau (compteurs, toner...) |
| **Suivi** | Message horodaté dans un ticket (public ou privé) |
| **Ticket récurrent** | Ticket généré automatiquement (maintenance préventive) |
| **TTO / TTR** | Délai de prise en compte / de résolution |
| **Urgence** | Criticité ressentie par le demandeur |
| **2FA / TOTP** | Double authentification par code temporaire |

---

## 68. Quiz — 10 questions + réponses

**Q1. Que se passe-t-il si le cron système de GLPI ne tourne pas ?**
R : Aucun courriel ne part (file `queuednotification` bloquée), les SLA ne sont
pas calculés, le collecteur mail ne relève rien, les tickets récurrents ne sont
pas créés. C'est la cause n°1 des « GLPI ne fait rien tout seul ».

**Q2. Quelle est la différence entre les statuts « Résolu » et « Clos » ?**
R : « Résolu » = le technicien propose une solution, en attente de validation du
déclarant. « Clos » = le déclarant a validé (ou clôture automatique après le délai
paramétré). On ne clôt jamais sans validation.

**Q3. Pourquoi le statut « En attente » exige-t-il toujours un motif et une
date de relance ?**
R : Parce qu'il suspend le compteur SLA : sans motif, les statistiques de délais
sont faussées et le ticket peut rester bloqué indéfiniment sans que personne ne
le relance.

**Q4. Citez 4 champs indispensables d'une fiche copieur.**
R : Marque + modèle, numéro de série, compteurs (N&B/couleur + date de relevé),
contrat de maintenance lié. (Plus : localisation client, consommables compatibles.)

**Q5. Comment la priorité d'un ticket est-elle déterminée ?**
R : Par la matrice urgence × impact (paramétrable). L'urgence est fixée par le
demandeur, l'impact validé par la hotline ; la priorité affichée reflète la
réalité opérationnelle.

**Q6. À quoi servent les tickets récurrents ? Donnez deux exemples métier.**
R : À générer automatiquement les tickets de maintenance planifiée. Ex. : relevé
mensuel des compteurs copieurs ; visite préventive trimestrielle.

**Q7. Pourquoi faut-il lier chaque sortie de cartouche à un ticket ?**
R : Pour décrémenter le stock en temps réel, imputer le coût au bon client/contrat,
et calculer la consommation (donc le coût réel) par client — base des
renégociations de contrats.

**Q8. Quelles sont les 3 sauvegardes indispensables de GLPI et leur fréquence ?**
R : Le dump de la base (`mysqldump`, quotidien), les fichiers
(`/var/lib/glpi/files`, quotidien), la configuration (à chaque changement).
Le tout selon la règle 3-2-1, avec test de restauration trimestriel.

**Q9. Un plugin provoque une page blanche après installation. Que faire ?**
R : Consulter les logs (`files/_log/php-errors.log`, log Apache), puis désactiver
le plugin en renommant son dossier (`xxx` → `xxx.off`). Ne jamais installer un
plugin sans vérifier sa compatibilité avec la version exacte de GLPI.

**Q10. Pourquoi ne doit-on jamais noter un mot de passe dans un ticket ?**
R : Les tickets sont visibles par de nombreuses personnes, exportables,
sauvegardés et journalisés : c'est une fuite de credential garantie. Les secrets
vont dans un coffre dédié (KeePass/Vault), jamais dans GLPI.

---

## 69. Pour aller plus loin

### Documentation officielle

- Documentation GLPI 10 (utilisateur + installation) — à consulter pour chaque
  version exacte : les écrans évoluent entre sous-versions.
- Dépôt GitHub `glpi-project/glpi` : notes de version (lisez-les **avant**
  chaque mise à jour), matrice de compatibilité PHP/plugins.

### Approfondir l'ITSM

- Référentiel **ITIL 4** (gestion des incidents, des problèmes, des changements).
- **RGPD** : registre des traitements pour les données clients hébergées dans GLPI.

### Côté technique

- **MariaDB** : réplication pour la haute disponibilité ; `mysqltuner` pour
  le dimensionnement.
- **Supervision** : intégrer les sondes de la section 46 à votre Zabbix/Nagios.
- **API** : automatiser les relevés de compteurs (section 53) puis la
  facturation au coût/copie.
- **SSO** : SAML/OIDC si votre infrastructure l'exige (au-delà du LDAP).

### Côté métier

- Croiser ce guide avec : le **guide copieurs** (codes, maintenance),
  le **guide onduleurs** (énergies), le **guide Proxmox** (hébergement),
  le **guide Debian/Ubuntu** (serveur) — tous transposables en articles
  de la base de connaissances (section 65).
- Rejoindre la **communauté GLPI** (forum, salon de discussion) : les retours
  d'expérience SAV y sont nombreux.

---

## 70. Checklist de mise en production

