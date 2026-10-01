---
id: collect-261001-rattrapage/rattrapage/huawei-licences-tac-guide-5
title: "Licences, support TAC et RMA Huawei — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-26", "2026-09-27", "2027-09-26"]
keywords: ["license"]
source: docs/RAG/collect-261001-rattrapage/huawei_licences_tac_guide.md
source_anchor: ""
source_lines: [619, 786]
sha256: edc1fe0a7d34ff03ba295e16781fd4f238dd85e2b3e53146bfee66bb098e225a
---

# Licences, support TAC et RMA Huawei — Guide ultra-complet

Cas d'usage : équipement mis au rebut, licence transférée, erreur d'activation.
Méthodes :

- **Côté équipement :** selon les modèles, `license revoke` / `undo license
  active` ou suppression via l'interface web (commande exacte **à vérifier
  sur le portail officiel** pour votre version).
- **Côté portail :** marquer le droit comme révoqué/retourné pour libérer
  l'Entitlement si le contrat le permet.

> Une licence révoquée côté équipement mais pas côté portail peut bloquer une
> réémission ultérieure : faites les deux.

## 36. Licence expirée : ce qui se passe techniquement

Comportement typique (varie selon produit et version) :

| Produit | À l'expiration |
|---|---|
| USG6000 (UTM) | Les lames IPS/AV/URL cessent de se mettre à jour puis se désactivent ; le firewall de base continue. Alarmes/logs « license expired ». |
| Contrôleur WLAN | Les AP en dépassement de capacité ne montent plus ; les AP existants peuvent rester (selon version) ou tomber. |
| eSight | Modules sous licence désactivés ; supervision réduite. |

**Ce n'est jamais silencieux** : des alarmes sont générées. Le problème, c'est
qu'elles arrivent souvent un vendredi soir et que personne ne les regarde
avant lundi (cas vécu n°101).

## 37. Période de grâce : mythe ou réalité ?

Certains produits appliquent une courte période de grâce après expiration
(quelques jours) pendant laquelle les fonctions restent actives avec des
avertissements. **Ce comportement varie selon les produits et les versions —
à vérifier sur le portail officiel** pour chacun de vos équipements.

**Règle de pilotage :** ne jamais compter sur une période de grâce. Pilotez
au J-90 comme si l'expiration était couperet.

## 38. L'inventaire des licences : construire le registre

Le registre de licences est **le** document administratif du chef de service.
Une ligne par licence, les colonnes suivantes au minimum :

| Colonne | Exemple fictif |
|---|---|
| Équipement (nom) | USG6000-Paris |
| Modèle | USG6680 *(fictif)* |
| ESN | 2102311HSL8H6000123 *(fictif)* |
| Fonction / capacité | IPS+AV+URL Filtering |
| Part Number | L-USG-IPS-1Y *(fictif)* |
| Entitlement ID | EID-2026-8F3K2P *(fictif)* |
| Date d'achat | 2026-09-26 |
| Date d'activation | 2026-09-27 |
| Date d'expiration | 2027-09-26 |
| Type | Commercial 1 an |
| Fichier .dat | `licences/2026/USG6000-Paris_21023..._IPS.dat` |
| Fournisseur | [À COMPLÉTER] |
| Statut | Active / Expirée / Transférée / Révoquée |
| Notes | Renouvellement 3 ans envisagé |

Stockage : tableur partagé + sauvegarde. Accès : chef de service + adjoint
a minima (pas sur le seul PC d'une personne).

## 39. Template de registre prêt à copier (CSV)

```csv
equipement;modele;esn;fonction;part_number;entitlement_id;date_achat;date_activation;date_expiration;type;fichier_dat;fournisseur;statut;notes
USG6000-Paris;USG6680;2102311HSL8H6000123;IPS+AV+URL;L-USG-IPS-1Y;EID-2026-8F3K2P;2026-09-26;2026-09-27;2027-09-26;Commercial 1 an;licences/2026/USG6000-Paris_IPS.dat;Partenaire X;Active;Passer en 3 ans en 2027
S310-BatA;S310-24P;2102311ABC...;Base;Incluse;-;-;-;-;Permanent;-;Partenaire X;Active;Rien à renouveler
```

*(ESN et références ci-dessus sont fictifs : remplacez par vos valeurs.)*

## 40. L'audit annuel des licences

Une fois par an (idéalement 3 mois avant la clôture budgétaire) :

1. [ ] Exporter le registre complet.
2. [ ] Pour chaque équipement critique : `display license` réel vs registre
      (écarts = licences activées sans trace, ou expirations non vues).
3. [ ] Lister les expirations des 12 prochains mois + budget associé.
4. [ ] Identifier les licences inutilisées (fonction payée jamais activée).
5. [ ] Vérifier les transferts en cours (RMA de l'année).
6. [ ] Présenter la synthèse à la direction : budget N+1.

## 41. Licences et haute disponibilité (actif/standby)

- Chaque membre d'un cluster HA a son **propre ESN** et donc besoin de son
  **propre fichier** `.dat`.
- Les deux fichiers doivent couvrir les **mêmes fonctions** avec des dates
  cohérentes, sinon le basculement dégrade le service (IPS actif d'un côté,
  inactif de l'autre).
- À l'achat d'un couple HA : commander **2×** chaque licence.

## 42. Licences et remplacement de carte (châssis modulaires)

Sur châssis (S12700, etc. — pas votre parc actuel, mais à savoir) :
- La licence est liée à l'ESN de la **MPU active**.
- Remplacement de la MPU = nouvel ESN = transfert de licence (section 32).
- Prévoir dans le contrat de maintenance une clause de **transfert express**
  en cas de remplacement de carte.

## 43. Coûts : ordres de grandeur et vigilance

Les tarifs des licences varient selon les gammes, les durées et les remises
partenaires. **Aucun prix n'est donné ici : tout tarif est à vérifier auprès
de votre partenaire.** En revanche, retenez les règles de gestion :

- Budgétez les renouvellements UTM comme des **charges récurrentes**
  (OPEX), pas comme des investissements ponctuels.
- Un devis doit toujours détailler : part number, quantité, durée, dates
  de début/fin de droits.
- Comparez le 1 an vs 3 ans **à service égal** (mêmes fonctions, mêmes dates).

## 44. Négocier les renouvellements avec le partenaire

Leviers d'un chef de service :

1. **Anticipation** (J-90) : commander dans l'urgence coûte toujours plus cher.
2. **Volume** : regrouper les renouvellements de plusieurs sites sur un seul
   bon de commande.
3. **Durée** : le 3 ans fait baisser le coût annuel et la charge de gestion.
4. **Alignement des dates** : faire coïncider les expirations (co-terming)
   pour une seule négociation annuelle.
5. **Concurrence loyale** : faire chiffrer par deux partenaires agréés.

## 45. Pièges classiques du cycle de vie des licences

| Piège | Conséquence | Parade |
|---|---|---|
| Preuve de droit jamais réclamée | Impossible de générer le `.dat` | L'exiger à la commande (checklist 113) |
| ESN recopié à la main | Fichier inutilisable | Copier-coller depuis `display esn` |
| Renouvellement commandé trop tard | Expiration + réactivation en urgence | Alertes J-90/J-60/J-30 |
| Licence activée sur un seul membre du HA | Dégradation au basculement | 2 fichiers, 2 ESN, vérification croisée |
| `.dat` perdu (pas d'archive) | Transfert RMA compliqué | Archivage systématique (section 22) |
| Registre non tenu | Aucune visibilité | Audit annuel (section 40) |
| Licence temporaire oubliée | Expiration à 60 jours en pleine prod | Ne jamais mettre de temporaire en production sans plan de bascule |

---

# D. LE PORTAIL SUPPORT HUAWEI (support.huawei.com)

## 46. Présentation : la tour de contrôle de votre parc

https://support.huawei.com (rubrique Enterprise) est le point d'entrée unique
pour :

- Enregistrer vos équipements (par numéro de série).
- Télécharger **firmwares, patchs et documentation**.
- Gérer les **licences** (lien avec les parties B et C).
- Vérifier les **garanties**.
- Ouvrir et suivre les **tickets TAC**.
- Accéder à la **base de connaissances** et à la communauté.

> Sans compte, l'accès est limité (documentation publique partielle). Avec un
> compte et des équipements enregistrés, les téléchargements logiciels sont
> débloqués **selon vos droits contractuels** (voir cas vécu n°105).

## 47. Créer un compte, pas à pas

1. Allez sur https://support.huawei.com → **Register / S'enregistrer**.
2. Renseignez : nom, prénom, **email professionnel** (évitez les adresses
   personnelles : le compte suit le collaborateur, pas l'entreprise —
   voir section 60), entreprise, pays, téléphone.
3. Choisissez un mot de passe robuste (gestionnaire de mots de passe
   recommandé) et activez la double authentification si proposée.
4. Validez l'email de confirmation.
5. Complétez le profil : rôle (administrateur réseau, chef de service...),
   secteur d'activité.

