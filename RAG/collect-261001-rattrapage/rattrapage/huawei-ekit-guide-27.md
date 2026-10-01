---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-27
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [2358, 2467]
sha256: 87bc04b3a1e4451bc417fd5c62d250a8884fd42f59c186357fe1c6f8339ca00f
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

- [ ] Smartphone chargé + batterie externe + câble
- [ ] PC portable + câble console (si le modèle a un port console) + adaptateur USB-série
- [ ] Testeur de câble + pince à sertir + connecteurs + étiquettes + marqueur
- [ ] Jarretières RJ45 (assortiment) + 2 jarretières fibre + injecteur PoE af/at
- [ ] Tournevis, pince coupante, colliers, scotch d'électricien, lampe frontale
- [ ] AP361 + S220 de secours (selon criticité du site)
- [ ] Multimètre (vérifier une prise avant de brancher la baie)
- [ ] Fiche site vierge + PV de recette vierge + ce guide (sur le téléphone)
- [ ] EPI : gants, lunettes (fibre), chaussures de sécurité (chantier/entrepôt)
- [ ] De l'eau et de la patience (les deux s'épuisent vite sur site)

## 220. Compte-rendu de visite technique avant-vente (modèle)

```
CR VISITE TECHNIQUE - [Client prospect] - Date : __________
Participants : __________
Locaux : surface ____ m2, etages ____, murs (placo/beton/verre) ____
Existant : cablage (etat) ____, equipements ____, operateur/debit ____
Besoins exprimes : ____ utilisateurs, Wi-Fi ____, cameras ____, tel IP ____,
                   invites ____, videosurveillance ____, multi-sites ____
Contraintes : budget ____, delai ____, securite ____, electrique ____
Photos : [jointes]   Plan : [croquis joint]
Hypotheses de chiffrage : __________________________________________
Prochaine etape : devis sous ____ jours / 2e visite ____
```
Un CR écrit après chaque visite = un devis juste + une preuve en cas de « vous n'aviez pas vu que… ».

## 221. Erreurs classiques n°26 à 30 (fin de la série)

**N°26 — Le firmware « dernier cri » installé sans lire la note.** Symptôme : régression surprise. Solution : la note de version se lit **avant**, pas après (42). Le « latest » n'est pas toujours le « greatest » en production.
**N°27 — Le test « ça marche sur mon téléphone ».** Symptôme : le client se plaint alors que « ça marchait ». Cause : testé sur 1 seul terminal, près de l'AP. Solution : tester sur 2-3 terminaux différents, aux endroits où les utilisateurs travaillent vraiment (74).
**N°28 — L'AP oublié dans le cloud après déménagement.** Symptôme : alertes fantômes, inventaire faux. Cause : équipement déplacé physiquement mais pas dans l'app. Solution : tout mouvement physique = mouvement dans le cloud le jour même (32).
**N°29 — Le mot de passe du coffre perdu.** Symptôme : personne ne peut administrer. Cause : 1 seul détenteur. Solution : 2 détenteurs + procédure de récupération écrite et scellée (102). Teste la récupération 1×/an.
**N°30 — Le « petit site » sans documentation.** Symptôme : 2 ans plus tard, personne ne sait ce qui est installé. Cause : « c'était petit, pas besoin de fiche ». Solution : **il n'y a pas de petit site** — fiche site (65) même pour 1 AP. Le temps « gagné » se paie ×10 au premier dépannage.

## 222. Mini-lexique fibre (pour parler MiniFTTO sans rougir)

- **OLT** : tête optique (côté baie) — elle pilote les terminaux.
- **ONT/ONU** : terminal optique (côté chambre/bureau) — ici intégré au F700D.
- **Splitter** : diviseur optique passif (1:16, 1:32…) — pas d'alimentation, pas de panne électronique.
- **dB / budget optique** : chaque splitter et soudure « mange » du signal — le budget total doit rester dans les limites de l'OLT.
- **Soudure (fusion)** : épissure des fibres — nécessite soudeuse + technicien formé (sous-traite si besoin, 212).
- **Réflectométrie (OTDR)** : mesure qui localise les défauts sur la fibre — à exiger à la réception d'un câblage fibre.
- **Pigtail / jarretière** : brin de raccordement — ne jamais le pincer ni le courber à angle vif (rayon de courbure !).
En MiniFTTO, la fibre est simple **quand elle est bien posée** — et un cauchemar quand elle est pincée dans une porte. Protège les cheminements.

## 223. Les logiques de l'app à connaître par cœur (sans les nommer)

L'app évolue, mais les logiques restent. Retiens les **principes**, pas les libellés exacts (qui changent à chaque version) :
1. Tout part du **site** : pas de site = pas d'équipement gérable.
2. L'**onboarding** précède la **configuration** : un équipement non onboardé ne reçoit rien.
3. La config se pense en **modèle** (VLAN, SSID) puis s'**applique** aux équipements — pas l'inverse.
4. Les **alertes** sont rattachées aux équipements **et** aux sites : lis toujours les deux niveaux.
5. Une **mise à jour** est une opération à part, jamais mélangée à un changement de config.
6. Ce que tu ne trouves pas dans l'app se fait (si possible) en **local** — l'app n'est pas le seul chemin.
Quand l'app change d'interface après une MAJ : respire, les 6 principes n'ont pas bougé — cherche où ils se cachent dans la nouvelle interface.

## 224. Tableau — qui fait quoi dans ton équipe (RACI simplifié)

| Tâche | Chef de service (toi) | Technicien senior | Technicien junior | Client |
|---|---|---|---|---|
| Avant-vente / devis | **R**esponsable | Consulté | — | Informé |
| Déploiement site simple | Approuve | **R** | Exécute | Informé |
| Déploiement site complexe | **R** | Exécute | Assiste | Informé |
| Astreinte niveau 1 | Escalade | **R** (rotation) | Assiste | Appelle |
| MAJ firmware | Valide | **R** | — | Informé |
| Audit annuel | **R** | Exécute | Assiste | Reçoit rapport |
| Accès cloud (création/retrait) | **R** | Demande | — | — |
| Facturation / litige | **R** | — | — | — |

**R** = responsable final. Une seule personne « R » par tâche — sinon personne ne l'est vraiment.

## 225. Naviguer dans ce guide : la table de décision rapide

```
JE DOIS...                                    -> VA VOIR...
Choisir du materiel pour un devis             -> §4-8, fiches §116-130, budgets §148
Convaincre un client                          -> §13, §87, FAQ §164, CR visite §220
Preparer un deploiement                       -> §15 (check-list), §64-69, §70 (ordre)
Onboarder des equipements                     -> §19-21, pense-bete §107
Configurer VLAN/SSID/PoE                      -> §64-69, procedures §152-154
Tester et recetter                            -> §74, PV §160, cablage §187
Depanner                                      -> §91 (methode), §92-100, §156, §221, §108
Superviser au quotidien                        -> §43 (rituel), §76-80
Securiser                                     -> §102-104, audit §155, durcir §199
Faire evoluer / migrer                        -> §61-63 (interop), §83-86 (Datacom)
Vendre de la maintenance                      -> §101, §178-180, §202
Former un technicien                          -> §158, maquette §210, valise §219
Gerder le cap sur 12 mois                     -> §215
```
**Fin du guide — 225 sections. Bon déploiement, chef.**

---

## 226. Ton premier client eKit : la check-list de zéro à facturé

- [ ] Maquette montée et maîtrisée (210) — tu as onboardé 3 équipements les yeux fermés
- [ ] Distributeur Gold identifié, tarifs et délais RMA connus par écrit (184, 201)
- [ ] Visite technique faite + CR écrit (220)
- [ ] Devis signé avec hypothèses écrites (159, 203)
- [ ] Matériel reçu, pré-staging au bureau (19)
- [ ] Câblage réceptionné (187), baie et onduleur prêts
- [ ] Déploiement selon l'ordre (70) + check-list (74)
- [ ] PV de recette signé (160), fiche site remise (65)
- [ ] Contrat de maintenance proposé (101, 178) — même s'il dit non, propose
- [ ] Inspection cloud à J+2/J+7 faite, cas notés, guide annoté
Le premier client est le plus dur et le plus précieux : c'est lui qui fera tes 5 suivants par recommandation.

## 227. Les 5 phrases à bannir devant un client

