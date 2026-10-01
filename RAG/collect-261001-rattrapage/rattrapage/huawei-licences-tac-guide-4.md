---
id: collect-261001-rattrapage/rattrapage/huawei-licences-tac-guide-4
title: "Licences, support TAC et RMA Huawei — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: ["2026-09-27", "2027-06-28", "2027-07-28", "2027-08-17", "2027-09-16", "2027-09-26", "2027-11-15"]
keywords: ["license"]
source: docs/RAG/collect-261001-rattrapage/huawei_licences_tac_guide.md
source_anchor: ""
source_lines: [457, 618]
sha256: fa918e57afbd3b6d90a0630aec3277b6cad165855c2b89edfced0ee091ff9d3b
---

# Licences, support TAC et RMA Huawei — Guide ultra-complet

1. Connectez-vous à l'interface web d'administration.
2. Allez dans **System > License Management**.
3. Cliquez **Upload / Activate**, sélectionnez le `.dat`.
4. Validez : le système affiche le résultat et le récapitulatif de la licence.
5. **Rechargez la page** et vérifiez que la licence apparaît avec le bon
   statut et la bonne date d'expiration.

L'activation web et l'activation CLI sont équivalentes ; choisissez selon
vos habitudes et vos droits d'accès.

## 26. Étape 6 — Vérifier : `display license`

```
<USG6000> display license
MainBoard:
  ESN code          : 2102311HSL8H6000123
  License state     : Demo / Normal / Expired
  License file name : licence.dat
  ...
  Feature           : IPS
  Expire date       : 2027-09-26
```

Ce que vous devez voir : **le bon ESN**, le bon fichier, chaque fonction
attendue avec sa date d'expiration. Tout écart = retour à l'étape concernée.

## 27. Lecture commentée d'un `display license`

```
<USG6000> display license
MainBoard:
  ESN code          : 2102311HSL8H6000123   ← doit matcher display esn
  License state     : Normal                ← « Normal » = OK
                                                « Demo » = licence d'essai/temporaire
                                                « Expired » = expirée → renouveler
  License file name : licence.dat           ← le fichier activé
  Created time      : 2026-09-27            ← date d'activation
  Feature           : IPS                   ← une ligne par fonction
  Expire date       : 2027-09-26            ← SURVEILLER : mettre une alerte J-90
  Feature           : AV
  Expire date       : 2027-09-26
  Feature           : URL Filtering
  Expire date       : 2027-09-26
```

**Automatisation conseillée :** un script qui se connecte en SSH chaque nuit,
lance `display license` sur les USG, parse les dates d'expiration et alerte
à J-90 / J-60 / J-30. C'est le garde-fou qui évite le cas vécu n°101.

## 28. Cas particulier USG6000 : activer la licence UTM complète

Sur les USG récents (HiSecEngine), la procédure suit le même schéma, avec
un guide dédié (« License Usage Guide » du modèle). Particularités :

- Les licences UTM sont gérées comme un **lot** (IPS+AV+URL) mais peuvent
  avoir des dates distinctes si achetées séparément.
- Après activation, forcez une **mise à jour des signatures** pour valider
  que la licence est bien prise en compte :
  `update signature` (commande exacte **à vérifier sur le portail officiel**
  pour votre version VRP).
- En HA actif/standby : activez sur les **deux** membres, chacun avec son
  propre fichier (deux ESN = deux fichiers).

## 29. Cas particulier eSight : charger la licence

1. Connectez-vous à la console d'administration eSight.
2. Rubrique **System > License Management** (intitulé exact selon version).
3. Importez le fichier de licence (format spécifique eSight, pas forcément
   `.dat` — **à vérifier sur le portail officiel**).
4. Vérifiez les modules déverrouillés et le compteur de nœuds.
5. Redémarrez les services concernés si la documentation l'exige.

## 30. Erreurs classiques d'activation : tableau de dépannage

Tableau établi d'après la documentation officielle de dépannage des licences
(HiSecEngine USG6000F License Usage Guide) :

| Message affiché | Cause probable | Conduite à tenir |
|---|---|---|
| `The ESN of the license file does not match with the device` | ESN saisi erroné sur le portail | 1. Redémarrer l'équipement et revérifier l'ESN (`display esn`). 2. Si l'ESN du fichier reste différent : **régénérer un fichier avec le bon ESN**. |
| `The product type of the license file does not match with the device` | Fichier généré pour un autre modèle | Régénérer un fichier pour le bon type de produit. |
| `The license file has expired` | Fichier généré depuis trop longtemps ou licence à durée dépassée | Obtenir un fichier non expiré / renouveler la licence. |
| `Failed to dispatch the license file` | Échec de transmission : disque standby plein, fichier absent, basculement actif/standby pendant l'activation | Libérer de l'espace sur la MPU standby ; vérifier la présence du fichier ; ne pas basculer pendant l'activation, puis réactiver. |
| `The license file has been activated` | Fichier déjà activé | Rien à faire : c'est normal, pas une erreur. |
| `Failed to activate the license` | Fichier corrompu ou problème interne | Vérifier l'intégrité (re-télécharger, re-transférer) ; si persistant : contacter le support technique. |

**Règle d'or :** devant une erreur, toujours vérifier dans l'ordre :
1) l'ESN (`display esn`), 2) le type de produit, 3) l'intégrité du fichier,
4) l'espace disque, 5) l'absence de basculement HA en cours.

---

# C. TRANSFÉRER, RENOUVELER, RÉVOQUER UNE LICENCE

## 31. Principe fondamental : la licence suit l'ESN, pas le propriétaire

Une licence Huawei est liée à **un ESN précis**. Elle ne se « déplace » pas
toute seule :

- Changer d'équipement = nouvel ESN = **nouveau fichier** à générer.
- Il n'existe pas de « copier-coller » du `.dat` d'un équipement vers un autre :
  l'activation échouera (ESN mismatch, section 30).
- Le transfert est une **procédure administrative** : on demande à Huawei (via
  le partenaire ou le portail) de réémettre la licence pour le nouvel ESN,
  en justifiant (RMA, remplacement, erreur d'ESN initiale).

## 32. Cas du RMA : transférer la licence vers l'équipement de remplacement

C'est le scénario le plus courant. Procédure type :

```
1. Le TAC valide le RMA et vous expédie l'équipement de remplacement.
2. À réception, relevez le NOUVEL ESN : display esn  →  ESN_NEW
3. Rassemblez : ancien ESN (ESN_OLD), numéro de RMA, preuve d'achat initiale.
4. Demandez le transfert :
   a) via votre partenaire (canal recommandé : il connaît votre dossier), ou
   b) via le portail de licences (fonction de réémission / rehost — intitulé
      exact à vérifier sur le portail officiel).
5. Générez le nouveau .dat pour ESN_NEW, activez-le (partie B).
6. Vérifiez : display license sur le nouvel équipement.
7. Mettez à jour le registre de licences (section 38) : l'ancien ESN passe
   au statut « retourné / révoqué ».
```

**Délais :** le transfert prend en général 24 à 72h ouvrées via le partenaire.
**Ne renvoyez pas l'ancien équipement avant d'avoir le nouveau fichier** si
vous devez encore l'utiliser (période de chevauchement à négocier).

## 33. Renouveler une licence UTM (le workflow annuel)

Les licences IPS/AV/URL des USG6000 sont des abonnements. Le renouvellement
annuel suit ce cycle :

```
J-90  : alerte automatique (script display license, section 27)
        → informer le N+1 / les achats du budget à prévoir
J-60  : demander un devis au partenaire (licence 1 an ou 3 ans)
J-30  : passer commande ; exiger la Proof of Entitlement
J-15  : générer le .dat (nouvel ESN ? non : même équipement, même ESN)
J-7   : activer le nouveau fichier (l'ancien reste valide jusqu'à sa fin)
J     : vérifier display license → nouvelle date d'expiration
```

**3 ans vs 1 an :** le 3 ans coûte moins cher par an en général (**tarifs à
vérifier auprès du partenaire**) et divise par trois la charge administrative.
Recommandé pour les équipements critiques.

## 34. Le calendrier de renouvellement : l'outil du chef de service

Tenez un calendrier partagé (ou un simple tableur) avec, pour chaque licence
à durée limitée :

| Équipement | Fonction | Expiration | Alerte J-90 | Alerte J-60 | Statut | Responsable |
|---|---|---|---|---|---|---|
| USG6000-PARIS | IPS/AV/URL | 2027-09-26 | 2027-06-28 | 2027-07-28 | À commander | Zelef |
| USG6000-LYON | IPS/AV/URL | 2027-11-15 | 2027-08-17 | 2027-09-16 | Devis demandé | [À COMPLÉTER] |

Revue **mensuelle** en réunion d'équipe : 5 minutes suffisent à éviter 90 %
des expirations surprises.

## 35. Révoquer / désactiver une licence

