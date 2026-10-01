---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-20
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: ["Apple"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [1706, 1779]
sha256: 9be621bacae043425706b92c87c0ae0ab141cf271a07a28e9eb96974d910608a
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

**N°17 — Le « ça marchait avant » sans sauvegarde.** Cause : changement non tracé, pas de sauvegarde. Solution : toujours sauvegarder avant (89) ; pour retrouver : interroger (« qui a touché à quoi ? »), comparer avec la fiche site, et reconstruire proprement.
**N°18 — Le DHCP qui s'épuise.** Symptômes : les nouveaux clients n'ont plus d'IP le matin. Cause : bail trop long + beaucoup de passages (hôtel, boutique). Solution : baux courts sur VLAN invités (4-12 h), élargir le scope.
**N°19 — L'AP « trop puissant ».** Symptômes : clients accrochés à l'AP du bout du couloir, débits faibles. Cause : puissance au max partout. Solution : **baisser** la puissance, ajouter des AP si besoin — la densité se règle en ajoutant des cellules, pas en criant plus fort.
**N°20 — Le câble « qui marchait » en 100 Mbit/s.** Symptômes : AP bridé à ~90 Mbit/s. Cause : paire abîmée, prise mal sertie. Solution : testeur, refaire la prise — un câble à 100 Mbit/s sur un AP Wi-Fi 6, c'est 80 % du potentiel jeté.
**N°21 — L'oubli du VLAN natif.** Symptômes : l'AP ne remonte pas / pas de DHCP. Cause : trunk avec natif incohérent entre switch et AP. Solution : aligner le natif des deux côtés (ou tout tagger explicitement).
**N°22 — La MAJ du vendredi soir.** Symptômes : week-end de panne. Cause : optimisme. Solution : MAJ en début de semaine, heures creuses, site pilote (42). Toujours.
**N°23 — Le SSID « TEST » oublié.** Symptômes : faille béante (souvent sans mot de passe). Cause : SSID de chantier jamais supprimé. Solution : l'audit annuel (155) les chasse ; nommer les SSID temporaires `ZZZ-TEMP-` pour les repérer.
**N°24 — Le switch plein à 100 % de PoE.** Symptômes : après ajout d'une caméra, un AP redémarre. Cause : budget dépassé (68). Solution : calculer avant d'ajouter, configurer les priorités PoE, 2e switch si besoin.
**N°25 — Le client qui « n'a rien touché ».** Symptômes : config modifiée, personne ne sait par qui. Cause : accès partagés / non tracés. Solution : comptes nominatifs (34), registre (106), et la phrase magique : « ce n'est pas grave, on va remettre d'aplomb — mais racontez-moi tout, même les détails bêtes ».

## 157. Plan de continuité (PRA) version PME — le minimum

Un PRA n'est pas réservé aux grands groupes. Pour chaque site critique :
1. **Inventaire vital** : passerelle, switch(s), AP critiques, onduleur — avec SN et contacts (distributeur, opérateur, électricien).
2. **Scénarios** : coupure électrique (→ onduleur + groupe ?), coupure opérateur (→ 4G de secours ?), panne passerelle (→ équipement de secours, 151), incendie/dégât des eaux de la baie (→ sauvegarde hors site, 89).
3. **RTO/RPO** : en PME, vise simple — RTO 4 h ouvrées (équipement de secours + config cloud), RPO 24 h (sauvegarde quotidienne de la config si changements fréquents).
4. **Test annuel** : coupe (simulée) — est-ce que l'onduleur tient ? Est-ce que l'équipe sait quoi faire ? Un PRA non testé = du papier.
5. **Document** : 2 pages max, dans la baie (plastifié) + chez toi. Le jour J, personne ne lit un pavé.

## 158. Former ton équipe à eKit : programme en 5 modules

1. **Fondamentaux** (1/2 j) : écosystème eKit, modèles, Fit/Fat/Cloud, VLAN/SSID/PoE de base.
2. **L'app** (1/2 j) : compte, sites, onboarding Wi-Fi + scan, config d'un site type — **en maquette** (113), pas en théorie.
3. **Déploiement** (1 j) : câblage, ordre de mise en service (70), étiquetage, check-list (74), sur un vrai site avec un senior.
4. **Dépannage** (1/2 j) : la méthode (91), les cas 92-100, le pense-bête (108), gestion de l'astreinte.
5. **Sécurité & admin** (1/2 j) : coffre, accès, sauvegardes, audit (155), relation client.
**Validation :** chaque technicien déploie seul un site « boutique » en maquette avant d'aller seul chez un client. Et ce guide = le manuel de référence de l'équipe.

## 159. Chiffrer une affaire eKit : la méthode de devis

1. **Visite technique** (jamais de devis sans voir le site) : mesures, murs, existant, besoins (check-list 15).
2. **Matériel** : liste (scénarios 46-60) + 10 % de marge (câbles, fixations, imprévus) — prix distributeur à jour.
3. **Main-d'œuvre** : câblage (au point ou au forfait), installation, configuration, tests, formation client — **ne brade jamais** la configuration et les tests, c'est là que se joue la qualité.
4. **Onduleur + électrique** : souvent oubliés — chiffre-les (lien avec ton guide onduleurs).
5. **Récurrent** : contrat de maintenance annuel (plan 101 + KPI 80) — c'est ta marge durable, propose-le systématiquement.
6. **Options** : USG (sécurité), AP supplémentaires (densité), MiniFTTO (si neuf), extension de garantie.
7. **Présentation** : 1 page de synthèse (besoin → solution → prix), puis le détail. Le client achète la **confiance**, pas la liste de courses.

## 160. Modèle de PV de recette (réception de chantier)

```
PROCES-VERBAL DE RECETTE - [Client - Site]              Date : __________
Installation realisee conformement au devis n° __________
[ ] Tous les equipements sont en ligne dans le cloud eKit (capture jointe)
[ ] SSID testes : __________ (debits releves : __________)
[ ] Isolation invites verifiee (pas d'acces LAN)
[ ] Portail captif teste (Android + iPhone)
[ ] Equipements filaires testes (imprimante, caisses, cameras)
[ ] Firmwares notes : __________
[ ] Sauvegarde remise : oui / Fiche site remise : oui / Photos : oui
[ ] Formation client effectuee (nom du forme : __________)
[ ] Reserves eventuelles : __________________________________________
Signature client : __________            Signature prestataire : __________
```
Le PV signé = la fin officielle du chantier et le début de la maintenance (et de la facturation du contrat). Sans PV, « il reste un petit truc » pendant 6 mois.

## 161. Diagnostic visuel : lire les LED (mémo générique)

Les codes exacts varient par modèle (**à vérifier sur la documentation officielle** du modèle), mais la logique est commune :
- **PWR/SYS vert fixe** : alimenté, démarré.
- **SYS clignotant** : démarrage en cours ou activité — attendre la fin du boot (1-3 min) avant de conclure à une panne.
- **LED éteinte** : pas d'alimentation → vérifier prise, PoE, câble.
- **LED port allumée** : lien établi ; **clignotante** : trafic ; **éteinte** : pas de lien → câble, port distant, PoE.
- **Rouge / orange** : alarme ou défaut → noter le motif exact (clignotement rapide/lent) et chercher dans le guide du modèle.
**Règle :** photographie les LED avant de toucher — c'est une preuve de l'état initial, utile pour le RMA et le rapport.

## 162. Câbles et connectique : le mémo de la sacoche

- **Toujours avoir** : 5 jarretières RJ45 Cat6 (0,5/1/2/3/5 m), 2 jarretières fibre (selon connectique du site), 1 injecteur PoE af/at, 1 testeur de câble, étiquettes + marqueur, colliers, tournevis, pince à sertir + 10 connecteurs RJ45.
- **Ne jamais faire** : rallonger un câble PoE avec un raccord « domino », sertir à l'arrache sans tester, laisser un câble en tension sur son connecteur, mélanger fibre monomode et multimode (ou leurs modules SFP).
- **Tester** chaque brin après sertissage : continuité + paires + longueur. Un câble non testé = un câble en panne.

## 163. eKit et la réglementation : ce qu'il faut respecter

