---
id: collect-261001-rattrapage/rattrapage/glpi-guide-5
title: "Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/glpi_guide.md
source_anchor: ""
source_lines: [515, 720]
sha256: b6ba9450717864113590dd9652393a17a8211edb5b942cadb08a834ce1f140fa
---

# Guide GLPI ultra-complet — Ticketing, gestion de parc et pilotage SAV

| Action | Fréquence conseillée | Rôle |
|---|---|---|
| `queuednotification` | 1 min | Envoi des courriels en file |
| `mailcollector` | 5 min | Relève IMAP → tickets |
| `sla` | 15 min | Calculs et escalades SLA |
| `purgeticket` | 1 jour | Clôture/purge selon règles |
| `alertnotclosed` | 1 jour | Alerte tickets non clos |
| `createinquest` | 1 heure | Envoi des sondages |
| `sendpasswordexpires` | 1 jour | Expiration mots de passe |

> Chacune a un **journal d'exécution** (dernière exécution, durée, erreurs).
> En dépannage, c'est le premier écran à ouvrir (section 55).

### 9.4 Supervision du cron (exemple Nagios/Zabbix)

```bash
#!/bin/bash
# check_glpi_cron.sh — alerte si le cron GLPI n'a pas tourné depuis 10 min
STAMP_FILE="/var/lib/glpi/files/_cron/last_run"
# alternative : interroger la BDD
AGE=$(mysql -N -s -e "SELECT TIMESTAMPDIFF(MINUTE, MAX(lastrun), NOW())
  FROM glpidb.glpi_crontasks;" 2>/dev/null)
[ "${AGE:-999}" -gt 10 ] && echo "CRITICAL: cron GLPI inactif depuis ${AGE} min" && exit 2
echo "OK: cron GLPI actif (dernière exécution il y a ${AGE} min)"
```

---

# PARTIE II — ORGANISATION : ENTITÉS, PROFILS, UTILISATEURS

## 10. Entités : l'arborescence multi-sites

L'entité est le découpage organisationnel. Tout objet GLPI (ticket, matériel,
contrat...) appartient à une entité. Les entités forment un **arbre**.

### Exemple pour une société de maintenance multi-sites

```
Entité racine : ACME Maintenance
├── Agence Nord
│   ├── Atelier
│   └── Clients rattachés (sous-entités par client important)
├── Agence Sud
└── Siège (services internes : informatique, comptabilité)
```

Ou, variante « par client » si chaque client est une entité :

```
ACME Maintenance
├── Client — Mairie de Villeurbanne
├── Client — Clinique Saint-Roch
└── Interne
```

### Bonnes pratiques

- **Ne multipliez pas les entités** : 5 à 15 suffisent pour une PME. Trop d'entités
  = administration cauchemardesque.
- Les objets « mutualisés » (contrats-cadres, modèles de documents) se mettent en
  **récursif** depuis l'entité racine pour être visibles partout.
- Un technicien itinérant reçoit son profil en **récursif** sur l'entité racine :
  il voit tous les sites.
- Les clients externes (portail déclarant) sont cantonnés à **leur** entité :
  ils ne voient que leurs tickets et leurs matériels.

### Création

`Administration > Entités` → `+` : nom, entité parente, adresse, responsable,
délégations (qui peut administrer cette entité). Pensez à renseigner l'adresse :
elle sert sur les documents et les notifications.

---

## 11. Profils et habilitations : qui fait quoi

Un **profil** = un jeu d'habilitations. GLPI 10 fournit des profils prédéfinis,
à dupliquer et ajuster (ne modifiez jamais les profils natifs directement :
dupliquez-les d'abord).

### Matrice des profils types pour un SAV

| Capacité | Post-only (déclarant) | Tech | Superviseur | Admin |
|---|---|---|---|---|
| Créer un ticket | ✅ | ✅ | ✅ | ✅ |
| Voir ses propres tickets | ✅ | ✅ | ✅ | ✅ |
| Voir tous les tickets de l'entité | ❌ | ✅ | ✅ | ✅ |
| S'assigner / assigner un ticket | ❌ | ✅ (soi) | ✅ | ✅ |
| Changer statut, priorité, SLA | ❌ | ✅ | ✅ | ✅ |
| Résoudre / clore | ❌ | ✅ | ✅ | ✅ |
| Gérer le parc (créer/modifier) | ❌ | lecture | ✅ | ✅ |
| Gérer contrats/budgets | ❌ | ❌ | ✅ | ✅ |
| Base de connaissances (publier) | ❌ | ✅ | ✅ | ✅ |
| Configuration / utilisateurs | ❌ | ❌ | partiel | ✅ |

### Profils à créer pour le métier

1. **Technicien SAV** (dupliqué de Tech) : + gestion des imprimantes/copieurs,
   + réservation de pièces, − suppression définitive.
2. **Superviseur / Chef de service** : tout le Tech + contrats, budgets, rapports,
   + réassignation inter-équipes, + validation des solutions.
3. **Déclarant client** (dupliqué de Post-only) : création de tickets avec gabarit
   « Intervention copieur », vision de ses matériels (ses copieurs !).
4. **Magasinier** : gestion du stock (cartouches, pièces détachées) sans accès tickets.

### Principe du moindre privilège

> Donnez à chacun le minimum nécessaire. Un technicien n'a pas besoin de voir les
> budgets ; un déclarant client ne doit jamais voir les tickets d'un autre client.
> Revoyez les profils **une fois par an** (départs, changements de poste).

---

## 12. Utilisateurs : création, import LDAP/AD, règles

### Création manuelle

`Administration > Utilisateurs` → `+` : identifiant, nom, courriel, entité(s) +
profil(s). Activez « Actif ». Le courriel est **indispensable** (notifications).

### Import depuis Active Directory / LDAP (recommandé)

`Configuration > Authentification > Annuaires LDAP` → `+` :

| Champ | Exemple |
|---|---|
| Nom | `AD-Entreprise` |
| Serveur | `ad.entreprise.lan` |
| Port | `389` (ou `636` LDAPS) |
| Base DN | `DC=entreprise,DC=lan` |
| DN de connexion | `CN=svc_glpi,OU=Services,DC=entreprise,DC=lan` |
| Filtre utilisateurs | `(&(objectClass=user)(!(userAccountControl:514)))` |
| Champ identifiant | `sAMAccountName` |
| Champ courriel | `mail` |

Puis `Administration > Utilisateurs` → « Import LDAP » : import de masse ou
**import à la première connexion** (l'utilisateur est créé automatiquement quand
il se connecte avec ses identifiants AD — zéro administration).

> **Sécurité :** préférez LDAPS (636) avec certificat valide. Le compte de
> connexion LDAP doit avoir le minimum de droits (lecture seule sur l'annuaire).

### Règles d'affectation automatique

`Administration > Règles > Règles d'affectation d'habilitations` : ex.
« si membre du groupe AD `SAV-Techniciens` → profil Technicien SAV récursif sur
l'entité racine ». L'arrivée d'un nouveau technicien = l'ajouter au groupe AD,
rien à faire dans GLPI.

---

## 13. Groupes : équipes, escalades, attrib
utions

Les **groupes** servent à assigner des tickets à une équipe plutôt qu'à une personne,
et à gérer les escalades.

### Groupes types d'un SAV

| Groupe | Rôle |
|---|---|
| `SAV — Hotline N1` | Qualification, diagnostic à distance |
| `SAV — Techniciens terrain` | Interventions sur site |
| `SAV — Atelier` | Réparations en atelier |
| `SAV — Superviseurs` | Escalade, arbitrage |
| `Informatique interne` | Parc interne de l'entreprise |
| `Astreinte` | Hors heures ouvrées |

`Assistance > Groupes` (ou `Administration > Groupes`) : créez le groupe,
ajoutez-y les utilisateurs, définissez s'il peut être **assigné** à un ticket
et/ou **notifié**.

### Bonnes pratiques

- Assignez d'abord au **groupe**, le superviseur redistribue aux techniciens :
  évite les tickets « orphelins » quand quelqu'un est absent.
- Un ticket non pris en compte sous X minutes (SLA) → **escalade automatique**
  vers le groupe `Superviseurs` (section 21).
- Le groupe `Astreinte` ne reçoit que les tickets d'urgence maximale hors horaires.

---

## 14. Notifications par courriel : configuration complète

### 14.1 Configuration SMTP

`Configuration > Notifications > Configuration des notifications` :

| Paramètre | Valeur type |
|---|---|
| Méthode d'envoi | SMTP |
| Hôte SMTP | `smtp.entreprise.lan` ou relais du FAI |
| Port | `587` (STARTTLS) ou `465` (SSL) |
| Chiffrement | STARTTLS |
| Authentification | Oui — compte dédié `glpi@entreprise.lan` |
| Expéditeur | `glpi@entreprise.lan` / nom « SAV — GLPI » |

> Utilisez un **compte dédié** à GLPI, jamais une boîte personnelle.
> Testez avec le bouton « Envoyer un courriel de test » (section 8.5).

### 14.2 Modèles de notification

`Configuration > Notifications > Modèles de notifications` : chaque événement
(nouveau ticket, nouveau suivi, résolution...) possède un modèle en français,
modifiable. Variables utiles :

