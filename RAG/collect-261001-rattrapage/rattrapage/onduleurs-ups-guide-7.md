---
id: collect-261001-rattrapage/rattrapage/onduleurs-ups-guide-7
title: "Onduleurs / UPS 10–120 kVA — Guide ultra-complet (exploitation & maintenance)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/onduleurs_ups_guide.md
source_anchor: ""
source_lines: [1208, 1406]
sha256: e64bec7557ad00eb4486cf655a922343b0b58ab1568cecc954e9c4fd99aab0b9
---

# Onduleurs / UPS 10–120 kVA — Guide ultra-complet (exploitation & maintenance)

- **Lithium** : le LFP devient compétitif en TCO → de plus en plus
  de projets neufs en Li-ion (moins de maintenance batteries).
- **Rendement** : les 97 %+ en double conversion arrivent (SiC).
- **Supervision** : le cloud devient standard (EcoStruxure et
  équivalents) → le chef de service pilote son parc au téléphone.
- **Cybersécurité** : les NMC exposées sont attaquées → VLAN
  dédié, mots de passe forts, mises à jour.
- **ECO mode intelligent** : bascule auto selon la qualité réseau.

---

## 56. Le mot du chef de service (à afficher au bureau)

```
La continuité électrique ne s'improvise pas.
Elle se DIMENSIONNE, s'INSTALLE, se TESTE et se MAINTIENT.
- Une batterie non testée est une batterie morte.
- Un groupe non testé en charge ne démarrera pas.
- Un bypass oublié n'est plus une protection.
- Tout ce qui n'est pas noté n'a pas été fait.
```

---

*Fin du guide — 1600+ lignes. Vos UPS entre de bonnes mains,
Zelef : **batteries surveillées, local climatisé, protections
calibrées, tests réguliers, équipe formée, tout noté.** C'est ça,
être chef de service systèmes & énergies.*
---

## 57. Visite annuelle type — déroulé chronométré (40 kVA)

```
08:00 Arrivée, EPI, prévenir le client et l'astreinte
08:15 Relevés : état, charge %, alarmes, event log (export)
08:30 Températures : local, armoire batteries (thermomètre IR)
08:45 Thermographie : cosses batteries, DJ, borniers (photos)
09:15 Tensions de chaque bloc batterie (noter les écarts)
09:45 Test d'impédance (si semestriel échu)
10:15 Dépoussiérage : grilles, ventilateurs, filtres
10:45 Vérification visuelle condensateurs (gonflés ? coulures ?)
11:00 Test bypass statique (forcer/relâcher, sans coupure)
11:15 Test bypass de maintenance (procédure §14)
11:30 Remise en normal, vérification 400 V/50 Hz
11:45 TEST COUPURE : ouvrir DJ entrée 2 min → autonomie, alarmes
12:00 Refermer, vérifier le retour, acquitter les alarmes
12:15 Mise à jour firmware (si nouvelle version validée)
12:45 Supervision : tester une alerte (SMS reçu ?)
13:00 Rapport : mesures, photos, préconisations, prochaine échéance
13:30 Signature client, départ
```

---

## 58. Remplacement des batteries — pas à pas complet

```
PRÉPARATION (J-7) :
□ Blocs identiques reçus (même marque, même lot si possible)
□ Vérifier les tensions à la réception (écart < 0,2 V)
□ Prévenir le client (intervention avec bypass)
□ EPI : gants, lunettes, outils isolés

JOUR J :
1. Forcer bypass statique → bypass de maintenance
2. Ouvrir DJ batterie(s), consigner (cadenas + étiquette)
3. VAT : 0 V aux bornes du string
4. Photographier le câblage AVANT démontage
5. Déconnecter : d'abord le point milieu (si accessible),
   puis les extrémités (+ et −)
6. Sortir les blocs (manutention à 2, ~30 kg/bloc)
7. Nettoyer l'armoire, vérifier les bacs (acide ?)
8. Poser les neufs, serrer au couple (manuel batterie)
9. Reconnecter selon la photo, vérifier la polarité 2×
10. Mesurer la tension totale : ___ V (attendu : ___ V)
11. Refermer DJ batterie → vérifier le floating (13,5–13,6 V/bloc)
12. Baseline d'impédance de chaque bloc (noter !)
13. Quitter le bypass, remise en normal
14. Test coupure 1 min → OK
15. Anciens blocs : palette, film, bordereau recyclage

RAPPORT : date, marque/lot, tensions, impédances, prochaine échéance
```

---

## 59. Remplacement des condensateurs bus DC — pas à pas

1. Consignation complète (§39) — **bus DC = danger**.
2. Attendre 10 min, **vérifier 0 V au multimètre** (2 mesures).
3. Photographier, noter les références (µF, V, température).
4. Dessouder/déclipser (selon montage), nettoyer les traces.
5. Monter les neufs **même valeur, même série** (kit constructeur
   de préférence).
6. Vérifier les soudures/connexions, remonter.
7. Déconsigner, démarrer, vérifier l'absence d'alarme bus DC.
8. Thermographie après 1 h de fonctionnement.
9. Noter la date → prochaine échéance dans 5 ans.

---

## 60. Mise en service d'un 120 kVA — récit type

- **J-30** : note de calcul, commande (UPS + 2 armoires batteries +
  TGBT ondulé), planning avec le client.
- **J-7** : local prêt ? (clim, plancher, éclairage, extincteur).
- **Jour J (équipe de 3)** :
  - Matin : mise en place (transpalette), câblage puissance
    (70–95 mm²), terre, batteries (2 strings de 40 blocs).
  - Après-midi : serrages au couple, vérifications, NMC (IP fixe).
  - Séquence de démarrage (§13.2), 400 V/50 Hz vérifiés.
- **J+1** : banc de charge 100 kW (paliers 25/50/75/100 %),
  test coupure, test bypass, test groupe.
- **J+2** : bascule de la charge réelle (fenêtre de maintenance
  nuit/week-end), supervision, formation exploitants.
- **J+7** : visite de contrôle (serrages, températures, alarmes).
- Dossier d'exploitation remis au client (§49).

---

## 61. Les 15 mesures à savoir faire

1. Tension AC triphasée (400 V ±10 %).
2. Tension DC bus (selon modèle, ~384–480 V).
3. Tension de chaque bloc batterie (13,5–13,6 V en floating).
4. Tension totale du string (somme, au volt près).
5. Courant de charge (pince AC, par phase).
6. Courant de décharge batterie (pince DC).
7. Résistance de terre (< 5 Ω).
8. Continuité du PE.
9. Impédance batterie (testeur dédié).
10. Température (thermomètre IR).
11. Fréquence réseau (50 Hz ±1 %).
12. THDv (analyseur ou pince качественная).
13. Isolement (mégohmmètre, machine consignée).
14. Chute de tension en charge (câbles longs).
15. Autonomie réelle (chronomètre + coupure simulée).

---

## 62. Dépannage avancé — 5 cas complexes

**Cas A — L'UPS refuse le groupe, mais le réseau passe**
→ Slew rate du groupe trop rapide + THDi. Mesurer avec
enregistreur, régler le régulateur, élargir la fenêtre de
fréquence d'entrée de l'UPS (menu).

**Cas B — Déclenchement intempestif du DJ d'entrée**
→ Courant d'appel du redresseur au démarrage + courbe DJ trop
sensible. → Courbe D ou DJ à déclenchement retardé, vérifier le
pré-charge du bus DC.

**Cas C — Autonomie divisée par 2 en 1 an (local climatisé)**
→ Un string sur deux déconnecté (DJ vibré) ! → Thermographie,
vérification des DJ DC, resserrage.

**Cas D — Alarmes fantômes la nuit**
→ Micro-coupures du réseau (non vues en journée). → Enregistreur
1 semaine, signalement distributeur, seuils adaptés.

**Cas E — Échauffement d'un câble de sortie (120 kVA)**
→ Cosse mal sertie. → Thermographie trimestrielle, ressertissage,
contrôle au couple. (Un câble qui chauffe = incendie en puissance.)

---

## 63. Négocier avec les fournisseurs (chef de service)

- **Batteries** : prix dégressif par palette, garantie écrite
  (conditions de température !), date de fabrication < 3 mois.
- **Pièces** : contrat cadre avec remise, délai garanti (4 semaines
  max), pièces critiques en dépôt local.
- **Constructeur** : mise en service incluse dès 60 kVA, formation
  équipe incluse, support téléphonique 24/7 négocié.
- Toujours **2 fournisseurs** minimum (rupture, prix).

---

## 64. Index rapide

- **VFI** : double conversion, 0 ms | **kW = kVA × PF**
- **Charge ≤ 80 %** | **Floating 13,5–13,6 V/bloc**
- **Écart > 0,3 V** = suspect | **Impédance +30 %** = remplacer
- **Local 20–25 °C** | **Terre < 5 Ω** | **Parafoudre T1+T2**
- **Groupe = 1,5–2× UPS** | **Test groupe mensuel en charge**
- **Bypass = sans protection** | **Bus DC : attendre + vérifier**
- **Condos 5 ans** | **Ventilos 4 ans** | **Thermographie trimestrielle**
- **Dossier d'exploitation** par site | **Tout noté, tout signé**

---

*Fin du guide — 1600+ lignes. Chef de service systèmes & énergies :
votre crédibilité se joue sur **l'anticipation**. Des batteries
testées, un groupe qui démarre, un dossier à jour — et aucune
surprise. Bon courage, Zelef.*
---

## 65. Tableaux de dimensionnement complets par puissance

