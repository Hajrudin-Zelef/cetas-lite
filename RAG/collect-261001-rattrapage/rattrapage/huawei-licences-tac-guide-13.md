---
id: collect-261001-rattrapage/rattrapage/huawei-licences-tac-guide-13
title: "Licences, support TAC et RMA Huawei — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["distribution", "license"]
source: docs/RAG/collect-261001-rattrapage/huawei_licences_tac_guide.md
source_anchor: ""
source_lines: [1964, 2082]
sha256: 2254e7007883f426e51bc9f9d90fbf8616518e401f3692d5389991c354cfcbf1
---

# Licences, support TAC et RMA Huawei — Guide ultra-complet

| Sujet | Spécificité eKit |
|---|---|
| **Garantie matérielle** | Assurée via le circuit distribution ; enregistrez quand même les S/N sur le portail Huawei. Durées **à vérifier sur le portail officiel** par modèle. |
| **RMA** | Déclenché via le partenaire (qui gère la logistique avec le distributeur). Gardez les mêmes réflexes : photos, n° de dossier, test du remplacement avant retour (partie G). |
| **Licences** | Mêmes principes (ESN, `.dat`, portail). Les AP eKit n'ont en général pas de licence unitaire (section 12). |
| **Firmwares** | Via le partenaire et/ou le portail selon droits ; l'application eKit peut proposer les mises à jour directement. |
| **TAC direct** | Non : le partenaire est le TAC de fait (niveau 1). S'il est compétent et contractuellement engagé, c'est même plus rapide. |
| **Documentation** | Portail Huawei + ressources eKit dédiées (guides d'installation simplifiés, très bien faits pour les PME). |

**Recommandation de chef de service :** traitez votre partenaire eKit comme
une extension de votre équipe — avec des SLA écrits. Un bon partenaire eKit
vaut mieux qu'un mauvais accès direct ; un partenaire sans SLA écrite est un
risque (cas n°111).

---

# K. RÉFÉRENCES

## 120. Glossaire

| Terme | Définition |
|---|---|
| **AR (Advance Replacement)** | La pièce de remplacement est expédiée avant le retour de la défectueuse. |
| **Batch Import** | Génération de fichiers de licence en masse sur le portail (template + ESN). |
| **Commissioning license** | Licence temporaire (~30 jours) pour mise en service et tests. |
| **Co-terming** | Alignement des dates d'expiration de plusieurs contrats/licences. |
| **.dat** | Fichier de licence chiffré Huawei, lié à un ESN. |
| **Diagnostic-information** | Commande agrégeant version, config, états et logs pour le TAC. |
| **DOA (Dead On Arrival)** | Matériel défectueux dès réception. |
| **EID (Entitlement ID)** | Identifiant de droit d'utilisation d'une licence. |
| **ESDP** | Plateforme de distribution électronique des logiciels/licences Huawei. |
| **ESN** | Equipment Serial Number : identifiant logique lié aux licences. |
| **GSC** | Global Service Center : centre de services Huawei (ex. : Europe). |
| **HA** | Haute disponibilité (actif/standby). |
| **HedEx** | Lecteur de documentation électronique Huawei (fichiers .hdx). |
| **Hi-Care** | Portefeuille de contrats de support Huawei pour l'entreprise. |
| **LAC** | License Authorization Code (selon gammes). |
| **Logbuffer** | Journal en mémoire vive (volatil au reboot). |
| **Logfile** | Journal persistant en flash. |
| **NBD** | Next Business Day : pièce reçue le jour ouvré suivant (si RMA avant 15h). |
| **NBD-S** | NBD Ship : pièce expédiée le jour ouvré suivant. |
| **OPEX** | Dépenses d'exploitation (dont renouvellements de licences). |
| **P1–P4 / S1–S4** | Niveaux de priorité/sévérité des tickets TAC. |
| **PN** | Part Number : référence commerciale. |
| **Proof of Entitlement** | Document prouvant le droit : EID + mot de passe d'activation. |
| **RFR (Return For Repair)** | Retour d'abord, réparation sous ~30 jours ouvrés. |
| **RMA** | Return Material Authorization : autorisation + n° d'échange. |
| **RTU** | Right To Use : licence d'activation d'une capacité (ex. : ports 100G). |
| **S/N** | Numéro de série physique du matériel. |
| **SLA** | Engagement de niveau de service (délais contractuels). |
| **Spare froid** | Équipement de secours stocké, pré-configuré, non sous tension. |
| **TAC** | Technical Assistance Center : support technique Huawei. |
| **Temporary license** | Licence temporaire (~60 jours) : démo, avant-vente, secours. |
| **UTM** | Unified Threat Management : lames IPS/AV/URL/antispam des USG. |
| **VRP** | Versatile Routing Platform : OS des équipements réseau Huawei. |
| **WAC** | Contrôleur WLAN (gestion centralisée des AP). |

## 121. Quiz — 10 questions

1. Un fichier `.dat` généré pour l'ESN d'un USG peut-il être activé sur un
   autre USG du même modèle ?
2. Que signifie le message `The ESN of the license file does not match with
   the device` et que faire ?
3. Citez les trois informations minimales pour générer un fichier de licence
   sur le portail.
4. Un ticket S1 est ouvert à 16h30 avec un contrat Hi-Care Standard 9x5xNBD.
   Quand la pièce de remplacement arrivera-t-elle au plus tôt ?
5. Quelle est la différence entre NBD-S et NBD ?
6. En S1, que se passe-t-il si l'ingénieur TAC ne parvient pas à vous joindre
   dans l'heure ?
7. Pourquoi ne faut-il jamais activer une licence pendant un basculement
   actif/standby ?
8. Un AP361 eKit tombe en panne : quel est le bon premier réflexe côté support ?
9. Citez 4 des 10 informations obligatoires d'un ticket TAC.
10. Après un RMA d'USG6000, quelle vérification licence ne faut-il jamais oublier ?

## 122. Réponses du quiz

1. **Non.** Le `.dat` est cryptographiquement lié à un ESN précis ; sur un
   autre équipement, l'activation échoue (ESN mismatch). Il faut régénérer un
   fichier pour le nouvel ESN (transfert, section 32).
2. L'ESN saisi sur le portail ne correspond pas à l'équipement. Vérifier
   l'ESN réel (`display esn`, éventuellement après reboot), puis **régénérer
   un fichier avec le bon ESN** (section 30).
3. L'**Entitlement ID** (ou mot de passe d'activation de la preuve de droit),
   l'**ESN exact** de l'équipement cible, un **compte** sur le portail de
   licences (sections 18–20).
4. Le n° RMA généré après 15h00 un jour ouvré → enregistré le jour ouvré
   suivant → pièce reçue le **surlendemain ouvré** (section 85). (Si la panne
   est un jeudi avant un week-end prolongé : voir cas n°112.)
5. **NBD-S** : la pièce est *expédiée* le jour ouvré suivant (réception J+2/J+3).
   **NBD** : la pièce est *reçue* le jour ouvré suivant, si le RMA est généré
   avant 15h00 (section 85).
6. La sévérité est **temporairement abaissée** jusqu'au rétablissement du
   contact (section 64). D'où l'astreinte joignable 24x7 en S1.
7. Risque d'erreur *Failed to dispatch the license file* : le fichier peut
   échouer à se répliquer vers la MPU standby, surtout si son disque est
   plein ou si le basculement survient pendant l'activation (section 30).
8. Contacter le **partenaire** (circuit eKit : le partenaire est le point
   d'entrée unique, section 118), pas la hotline Huawei Enterprise directe.
9. Quatre parmi : modèle exact, version logicielle (VRP), ESN/S/N, n° de
   contrat de support, description du problème, impact métier, logs
   (diagnostic-information), topologie, contact joignable, actions déjà
   tentées (section 69).
10. Vérifier que la licence a été **transférée vers le nouvel ESN** :
    `display license` sur le nouvel équipement, fonctions + dates OK
    (sections 32 et 99, cas n°110).

## 123. Pour aller plus loin

**Documentation officielle (à consulter sur https://support.huawei.com) :**
- *HiSecEngine USG6000F License Usage Guide* — la référence licences des USG
  (activation, erreurs, dépannage).
- *Huawei Hi-Care Service Description* — paliers, SLA, responsabilités TAC/client.
- *Enterprise Product Warranty Policy* (par région) — durées et modalités
  par gamme de produits.
- *HiCare Support Services User Guide* — classification des sévérités S1–S4.
- Release Notes de votre version VRP exacte — prérequis et incompatibilités.

