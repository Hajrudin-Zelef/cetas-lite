---
id: collect-261001-rattrapage/rattrapage/huawei-licences-tac-guide-16
title: "Licences, support TAC et RMA Huawei — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: ["2026-09-27", "2026-10-03"]
keywords: ["incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_licences_tac_guide.md
source_anchor: ""
source_lines: [2412, 2500]
sha256: 4af0320aa9778203eaf0d3f6ad0675b8a79bbb1f61a67145a7b39af25abd7d4a
---

# Licences, support TAC et RMA Huawei — Guide ultra-complet

**Q2. Puis-je utiliser le même `.dat` sur deux équipements identiques ?**
R : Non. Un fichier = un ESN. Il faut un fichier par équipement (section 31).

**Q3. La licence a expiré hier, tout est-il coupé ?**
R : Les fonctions sous licence se désactivent (ou cessent de se mettre à jour)
selon le produit ; le reste continue. Renouvelez et activez au plus vite
(sections 33, 36).

**Q4. Combien de temps pour un transfert de licence après RMA ?**
R : Comptez 24 à 72h ouvrées via le partenaire. Anticipez dès la réception
de la pièce de remplacement (section 32).

**Q5. Le TAC est-il vraiment joignable 24x7 ?**
R : Oui pour l'assistance technique (canaux hotline/email/portail), pour les
clients couverts. Les délais de *réponse* dépendent de la sévérité (section 63).

**Q6. J'ai un contrat Basic : aurai-je une pièce en 4h ?**
R : Non. Le 4h est réservé au palier Premier (P1/P2). En Basic, la pièce est
*expédiée* le jour ouvré suivant (section 83).

**Q7. Mon partenaire eKit ne répond pas, puis-je appeler Huawei ?**
R : Vous pouvez essayer, mais le circuit officiel eKit passe par le
partenaire. Si le vôtre est défaillant, c'est un problème **contractuel** à
traiter (mise en demeure, changement de partenaire), pas un problème TAC
(sections 111, 118).

**Q8. Faut-il une licence par AP361/AP761 ?**
R : En général non : c'est le contrôleur (ou la plateforme de gestion) qui
porte la licence de capacité « nombre d'AP » (section 12).

**Q9. Que faire si l'étiquette S/N est illisible ?**
R : Utilisez le portail (équipements enregistrés, section 49), le bon de
livraison d'origine, ou `display` en CLI si l'équipement démarre (cas n°104).

**Q10. Puis-je renvoyer le défectueux avant de recevoir le remplacement ?**
R : En Advance Replacement, non : gardez-le jusqu'au test du remplacement
(sections 97–98). En RFR, c'est le principe même (renvoi d'abord).

**Q11. Un firmware est-il couvert par la garantie standard ?**
R : Les mises à jour logicielles sont incluses ~1 an en garantie standard ;
au-delà, il faut un contrat logiciel actif (sections 51, 82).

**Q12. Comment prouver à ma direction qu'il faut du Hi-Care Premier ?**
R : Chiffrez le coût d'une indisponibilité (cas n°112 : 4 jours sans réseau)
vs le surcoût annuel du palier. Un seul incident évité rentabilise souvent
le contrat (sections 84–85).

## 142. Registre des incidents : modèle de suivi annuel

Tenez un registre unique de tous les incidents significatifs (S1/S2 et RMA).
Il sert au REX, à l'audit et à la négociation des contrats.

```csv
date;ticket;severite;equipement;site;symptome;cause_racine;duree_indispo;rma_numero;cout_estime;rex_fait;actions
2026-09-27;SR-2026-0912;S2;USG6000;Paris;Licences UTM expirees;Renouvellement non suivi;48h;non;1500 EUR;oui;Script J-90 mis en place
2026-10-03;SR-2026-1044;S1;S310;Lyon;Switch ne boote plus;Panne alimentation;26h;RMA-88412;800 EUR;oui;Spare froid achete
```

*(Valeurs fictives : exemple de format uniquement.)*

Colonnes minimales : date, n° ticket, sévérité, équipement, symptôme, cause
racine, durée d'indisponibilité, n° RMA le cas échéant, REX fait (oui/non).

## 143. Antisèche : les 20 points à retenir (une page)

1. La licence **autorise**, le support **assiste**, la garantie **remplace**.
2. Un `.dat` = **un ESN** = un équipement. Usage unique.
3. Toujours copier l'ESN depuis `display esn`, jamais le retaper.
4. Exiger la **Proof of Entitlement** à la commande, pas après.
5. Les licences UTM des USG sont des **abonnements** : calendrier J-90/J-60/J-30.
6. En HA : **une licence par membre**, mêmes fonctions, mêmes dates.
7. Archiver chaque `.dat` + tenir le **registre de licences**.
8. Ne jamais activer pendant un basculement actif/standby.
9. S1 = 30 min, S2 = 60 min, S3 = 2h, S4 = NBD (réponse TAC).
10. S1 seulement si l'activité est **réellement** bloquée ; rester joignable 24x7.
11. Ticket = modèle + version + ESN/S/N + contrat + 5W + logs + contact.
12. `display diagnostic-information` **pendant** le problème, avant tout reboot.
13. Le logbuffer est **volatil** : l'exporter avant de rebooter.
14. NBD-S = pièce *expédiée* J+1 ; NBD = pièce *reçue* J+1 (si RMA avant 15h).
15. RMA = **n° RMA** d'abord ; les délais courent à partir de sa génération.
16. Tester le remplacement **avant** de renvoyer le défectueux.
17. Après RMA : restaurer la config **et** transférer la licence (nouvel ESN).
18. eKit = **partenaire d'abord**, avec des SLA écrits dans le contrat.
19. Un spare froid sur étagère vaut mieux qu'un NBD un jour férié.
20. Ce qui n'est pas écrit n'existe pas : registre, calendrier, REX, dossier parc.

---

*Fin du guide. Bon courage — et que vos licences n'expirent jamais un vendredi soir.*
