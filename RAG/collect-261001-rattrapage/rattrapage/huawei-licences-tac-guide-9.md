---
id: collect-261001-rattrapage/rattrapage/huawei-licences-tac-guide-9
title: "Licences, support TAC et RMA Huawei — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: ["2028-03-30", "2028-06-30"]
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_licences_tac_guide.md
source_anchor: ""
source_lines: [1349, 1503]
sha256: 5ef6aa2a34e400918d24c9ec2b23198299e7b4b503a2192a93ead12d3f4cad84
---

# Licences, support TAC et RMA Huawei — Guide ultra-complet

## 81. Garantie : les types chez Huawei

Deux grands régimes, cumulables dans le temps :

| Régime | Nature | Durée typique |
|---|---|---|
| **Garantie standard** | Incluse à l'achat, définie par la politique de garantie Huawei par gamme | 1 an à vie limitée selon produits (détails section 82) |
| **Hi-Care** | Contrat de support payant (« maintenance », extension de garantie) | 1 à 3 ans, renouvelable |

Logique d'achat : la garantie standard couvre le matériel ; **le TAC complet
et les délais de remplacement courts exigent Hi-Care** (voir section 84).

## 82. La garantie standard : contenu par gamme

D'après la politique de garantie Huawei Enterprise (documents publics) :

| Gamme | Garantie standard constatée* |
|---|---|
| Certains switches (S5710HI, S5720HI, S6700, S7700...) | Garantie à vie limitée (10 ans max.), pièces expédiées à J+10 ouvrés |
| Certains switches (S2700, S3700, S5700...) | Garantie à vie limitée (10 ans max.), pièces expédiées le jour ouvré suivant |
| Routeurs AR, IoT, sécurité, certains switches/AP WLAN | 1 an, avec variantes (retour pour réparation, etc.) |
| Mises à jour logicielles | 1 an inclus en général |

*\* Politique détaillée par référence dans la « Warranty List » officielle —
**à vérifier sur le portail officiel** pour chaque modèle de votre parc
(AP361, AP761, S310, AR720, USG6000). Les conditions varient selon les
régions.*

Modes de remplacement sous garantie standard :
- **AR (Advance Replacement)** : la pièce de remplacement est expédiée avant
  (ou sans attendre) le retour de la défectueuse.
- **RFR (Return For Repair)** : vous renvoyez d'abord, réparation sous
  ~30 jours ouvrés après réception au centre de pièces.

## 83. Hi-Care : les paliers

Hi-Care est le portefeuille de support « branded » Huawei pour l'entreprise.
Paliers constatés dans la documentation officielle :

| Palier | Fenêtre | Pièces de rechange |
|---|---|---|
| **Basic** | 9x5 (9h-18h, lun-ven hors fériés) | Expédiées le jour ouvré suivant (**NBD-S** = NBD Ship) |
| **Standard** | 9x5 | **Reçues** le jour ouvré suivant (**NBD**) si RMA avant 15h00 locales |
| **Premier** | 24x7 pour P1/P2 ; 9x5 pour P3/P4 | **4 heures** (P1/P2) après génération du n° RMA ; NBD (P3/P4) |
| **Onsite Standard** | 9x5 | NBD + intervention sur site 9x5 |
| **Onsite Premier** | 24x7 | 4h + intervention sur site 24x7 |

Tous les paliers incluent : **TAC 24x7**, support en ligne (self-help),
mises à jour logicielles système.

## 84. Tableau comparatif : que choisir pour quel équipement ?

| Équipement (exemple) | Enjeu | Palier recommandé | Justification |
|---|---|---|---|
| USG6000 site principal | Critique, pas de redondance | **Premier 24x7x4** | 4h de remplacement, TAC prioritaire |
| USG6000 site secondaire (HA) | Important, redondé | Standard 9x5xNBD | Le second membre tient la charge |
| S310 bâtiment | Moyen, redondance partielle | Standard 9x5xNBD | NBD acceptable si maillage |
| AP361 / AP761 | Faible unitaire | Basic ou garantie standard | Un AP en panne = gêne locale ; stock de spare (section 112) |
| AR720 agence distante | Moyen | Standard 9x5xNBD | Agence isolée : NBD suffisant si 4G de secours |
| eSight (serveur) | Supervision | Standard | Pas de coupure métier directe |

**Règle budgétaire :** le coût Hi-Care est en général de l'ordre d'un
pourcentage annuel du prix du matériel (**tarifs à vérifier auprès du
partenaire**). Budgétez-le dès l'achat, pas après la première panne.

## 85. Les délais pièces : comprendre NBD-S, NBD et la règle des 15h00

- **NBD-S (Next Business Day - Ship)** : la pièce **part** de chez Huawei le
  jour ouvré suivant. Ajoutez le transit : réception J+2 ou J+3 en pratique.
- **NBD (Next Business Day)** : la pièce **arrive** le jour ouvré suivant,
  **à condition que le n° RMA soit généré avant 15h00 locales**. Après 15h00 :
  le RMA est enregistré le jour ouvré suivant, la pièce arrive le surlendemain.
- **4h (Premier, P1/P2)** : la pièce arrive **dans les 4 heures** suivant la
  génération du n° RMA, 24x7.

**Conséquence opérationnelle :** un ticket S1 ouvert à 16h30 avec un contrat
Standard = pièce le surlendemain, pas le lendemain. D'où l'importance du
palier pour les sites critiques (cas vécu n°112).

## 86. Début et fin de service : la règle des 90 jours

D'après la documentation Hi-Care, si aucune date de début ne figure au bon
de commande :

- Service vendu **avec** le produit : démarrage le **90e jour après
  l'expédition** du produit par Huawei.
- Si Huawei fournit aussi l'**installation / mise en service** : démarrage
  à la **date d'acceptation client**.
- **Renouvellement** : démarrage le lendemain de la fin du contrat précédent
  (pas de trou de couverture si renouvelé à temps).

> **Vérifiez vos contrats** : une installation faite au jour 30 avec un
> démarrage de service au jour 90 = 60 jours de « garantie standard seule ».
> Négociez la date de début sur le bon de commande.

## 87. Ce qui est couvert / non couvert

**Couvert en général :**
- Panne matérielle non provoquée (composant défectueux).
- Bugs logiciels (correctifs via TAC).
- Remplacement à l'identique ou équivalent.

**Non couvert en général :**
- Dommages physiques (choc, eau, surtension sans parafoudre...).
- Modifications non autorisées, firmware tiers.
- Consommables (ventilateurs usés ? selon contrat — **à vérifier**).
- Perte de données de configuration (votre responsabilité : sauvegardes !).
- Interventions dues à une mauvaise utilisation documentée.

**En cas de doute :** le TAC qualifie la panne ; s'il conclut à un dommage
exclu, le RMA est refusé et un devis de réparation est proposé (cas vécu n°107).

## 88. Vérifier la garantie : procédure annuelle

1. Exporter la liste des S/N depuis le portail (section 50).
2. Pour chaque équipement : noter fin de garantie + palier Hi-Care.
3. Marquer en rouge les fins < 6 mois.
4. Chiffrer les renouvellements avec le partenaire.
5. Présenter au budget N+1.

Modèle de ligne de suivi :
```
S310-BatA | S/N 21023... (fictif) | Garantie std + Hi-Care Standard
| Fin : 2028-06-30 | Statut : OK | Action : renouveler avant 2028-03-30
```

## 89. Renouveler / étendre la garantie et le Hi-Care

- Le renouvellement se commande **avant** l'expiration (idéalement J-90,
  comme les licences).
- Un contrat expiré depuis longtemps peut nécessiter une **remise en
  conformité** (inspection, voire refus) — **conditions à vérifier auprès
  du partenaire**.
- À chaque renouvellement : exiger l'avenant écrit avec les dates exactes
  de début/fin et le palier.

## 90. Garantie et achats : les clauses à faire figurer au bon de commande

- [ ] Références exactes des contrats Hi-Care (palier, durée, sites couverts).
- [ ] Date de début de service explicite.
- [ ] Liste des pièces de rechange critiques couvertes.
- [ ] Engagement de délai de génération du n° RMA.
- [ ] Contact support dédié du partenaire (nom, téléphone).
- [ ] Clause de transfert de licence en cas de RMA (section 32).

---

# G. PROCÉDURE RMA

## 91. RMA : définition et vocabulaire

**RMA** (Return Material Authorization) = l'autorisation formelle, avec un
**numéro RMA**, d'échanger un matériel défectueux. Sans numéro RMA, aucun
échange n'est possible : c'est la clé de voûte de tout le processus.

