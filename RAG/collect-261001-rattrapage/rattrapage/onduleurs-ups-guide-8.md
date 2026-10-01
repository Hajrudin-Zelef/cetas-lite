---
id: collect-261001-rattrapage/rattrapage/onduleurs-ups-guide-8
title: "Onduleurs / UPS 10–120 kVA — Guide ultra-complet (exploitation & maintenance)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/onduleurs_ups_guide.md
source_anchor: ""
source_lines: [1407, 1602]
sha256: f44747597075b9c83d480c85ee9352e8fd220a556706a170ad65b3991e618c9c
---

# Onduleurs / UPS 10–120 kVA — Guide ultra-complet (exploitation & maintenance)

> Ordres de grandeur (400 V tri, cuivre, à valider par note de calcul).

### 65.1 Protections et câbles

| UPS | I approx. | DJ entrée | DJ sortie | Câble AC | DJ batterie (384 V) |
|---|---|---|---|---|---|
| 10 kVA | 16 A | 25 A | 25 A | 4 mm² | 32 A DC |
| 20 kVA | 32 A | 40 A | 40 A | 10 mm² | 63 A DC |
| 40 kVA | 64 A | 80 A | 80 A | 25 mm² | 125 A DC |
| 60 kVA | 96 A | 125 A | 125 A | 35 mm² | 160 A DC |
| 120 kVA | 192 A | 250 A | 250 A | 95 mm² | 250 A DC |

### 65.2 Batteries types (15 min d'autonomie)

| UPS | Charge type | Bus DC | Config indicative |
|---|---|---|---|
| 10 kVA | 7 kW | 384 V | 32 blocs 12 V / 9 Ah (1 string) |
| 20 kVA | 14 kW | 384 V | 32 blocs 12 V / 18 Ah (1 string) |
| 40 kVA | 29 kW | 384–480 V | 32–40 blocs / 24 Ah (1–2 strings) |
| 60 kVA | 43 kW | 480 V | 40 blocs / 40 Ah (2 strings) |
| 120 kVA | 86 kW | 480 V | 40 blocs / 65 Ah (2–3 strings) |

### 65.3 Groupe et climatisation

| UPS | Groupe conseillé | Clim à prévoir |
|---|---|---|
| 10 kVA | 15–20 kVA | 3–5 kW frigo |
| 20 kVA | 30–40 kVA | 5–8 kW frigo |
| 40 kVA | 60–80 kVA | 10–15 kW frigo (N+1) |
| 60 kVA | 90–120 kVA | 15–20 kW frigo (N+1) |
| 120 kVA | 180–250 kVA | 30–40 kW frigo (N+1) |

---

## 66. Procédure d'arrêt d'urgence

```
1. Bouton d'arrêt d'urgence (coup de poing) si danger immédiat
   → coupe TOUT (UPS + bypass)
2. Sinon, arrêt contrôlé :
   a. Prévenir les utilisateurs (arrêt des serveurs)
   b. Arrêt programmé des serveurs (scénario §42)
   c. Forcer le bypass → bypass de maintenance
   d. Ouvrir DJ sortie → DJ entrée → DJ batterie
3. Ne JAMAIS rouvrir sans diagnostic de la cause
4. Consignation avant toute intervention (§39)
```
- Le bouton d'arrêt d'urgence doit être **signalé et accessible**,
  testé annuellement.

---

## 67. Maintenance du groupe électrogène (le jumeau de l'UPS)

```
HEBDOMADAIRE : niveau huile, niveau liquide de refroidissement,
               niveau carburant, voyants, fuites
MENSUEL : démarrage EN CHARGE 30 min (> 30 % de charge),
          vérifier tension/fréquence/stabilité
TRIMESTRIEL : filtres (visuel), courroies, batterie de démarrage
              (tension + chargeur), test ATS complet
ANNUEL : vidange + filtres (huile, gasoil, air),
         liquide de refroidissement, test 2 h en charge,
         contrôle des sécurités (pression huile, température)
```
- Batterie de démarrage du groupe : la panne n°1 des groupes →
  chargeur d'entretien + remplacement tous les 3 ans.
- Carburant : rotation tous les 6 mois, additif anti-bactérien
  (le gasoil stocké se contamine en climat chaud).

---

## 68. Maintenance de la climatisation (la gardienne des batteries)

```
MENSUEL : filtres (nettoyage/remplacement), température de consigne
TRIMESTRIEL : condenseur (dépoussiérage), écoulement condensats,
              test de bascule N+1 (arrêter une clim → l'autre prend tout)
ANNUEL : contrôle frigoriste (pression, fluide), nettoyage complet
```
- Alarme haute température **supervisée** (SMS) : une clim en panne
  un vendredi soir = batteries à 40 °C tout le week-end.

---

## 69. Tableau de bord mensuel du service (modèle)

```
PARC UPS — ___/___
Nb UPS suivis : ___ | Disponibilité : ___ % | Incidents : ___
MTTR : ___ h | Alarmes traitées : ___ | En retard : ___
BATTERIES : âge moyen ___ ans | tests OK ___/___ | remplacements prévus : ___
GROUPES : tests mensuels OK ___/___ | carburant mini : ___
CLIM : alarmes ___ | bascules N+1 testées : ___
CONTRATS : renouvellements à 90 j : ___
BUDGET : dépensé ___ % | investissements à prévoir : ___
RISQUES : ___
ACTIONS : ___
```

---

## 70. Fiche site — modèle (une par site)

```
SITE : ___  Responsable : ___  Tél : ___
UPS : marque ___ modèle ___ n° série ___  Puissance ___ kVA
Mise en service : ___  Charge : ___ %  Firmware : ___
BATTERIES : marque ___ ___ blocs ___ V ___ Ah ___ strings
Mise en service : ___  Prochain remplacement estimé : ___
GROUPE : marque ___ puissance ___ kVA  Dernier test en charge : ___
CLIM : marque ___ puissance ___  N+1 : oui/non
SUPERVISION : NMC ___ (IP ___)  SNMP : oui/non  Alertes SMS : oui/non
CONTRAT : niveau ___  Échéance : ___
DOSSIER D'EXPLOITATION : ☐ à jour le ___
RISQUES NOTÉS : ___
```

---

*1600+ lignes atteintes. Chef de service : **votre parc est aussi
bon que votre pire checklist.** Appliquez-les toutes, et dormez
tranquille.*
---

## 71. Les 10 règles d'or (à afficher dans chaque local UPS)

```
1. NE JAMAIS intervenir sans consignation complète.
2. Le bus DC reste chargé après coupure : attendre + vérifier.
3. En bypass, la charge n'est PLUS protégée.
4. Une batterie non testée est une batterie morte.
5. Local batteries : 20–25 °C, ventilé, pas de flamme.
6. Outils isolés, pas de bijoux, une seule main sur le DC.
7. Groupe : test mensuel EN CHARGE, pas à vide.
8. Tout écart de tension batterie > 0,3 V = suspect.
9. Thermographie trimestrielle : le chaud annonce la panne.
10. Tout noté, tout signé, tout classé au dossier d'exploitation.
```

---

## 72. Fiche de mise en service — modèle à signer

```
MISE EN SERVICE UPS
Date : ___  Site : ___  Techniciens : ___
UPS : ___ (n° série : ___)  Puissance : ___ kVA
□ Local conforme (clim, ventilation, extincteur, accès)
□ Câblage vérifié (serrage, repérage, terre = ___ Ω)
□ Protections conformes à la note de calcul
□ Batteries : ___ blocs, tension totale = ___ V, polarité OK
□ Séquence de démarrage OK, 400 V / 50 Hz en sortie
□ Test coupure : ___ min d'autonomie mesurée
□ Test bypass statique + maintenance : OK
□ NMC configurée (IP : ___), alertes testées : OK
□ Firmware : ___
□ Dossier d'exploitation remis : oui
□ Formation exploitants : oui (___ personnes)
Réserves : ___
Signatures : installateur ___ / client ___
```

---

*Fin du guide — 1600+ lignes. À toi, chef : **dimensionne juste,
installe propre, teste tout, note tout.** Ton app sera redoutable
sur les UPS avec ce contenu.*
---

## 73. Annexe — conversions et formules utiles

```
P (kW) = S (kVA) × cos φ
I (A) = S (VA) / (√3 × 400)            [triphasé 400 V]
I (A) = P (W) / 230                    [monophasé 230 V]
Chute de tension : ΔU = √3 × I × L × (R cos φ + X sin φ)
Autonomie : Ah = P×t / (Vbus × ηond × ηbat)
Pertes UPS : P_thermique = P_charge × (1 − η)
Arrhenius : vie batterie / 2 tous les +10 °C
√3 ≈ 1,732
```

---

*1600+ lignes. Guide terminé.*

---

## 74. Pour aller plus loin (prochaines étapes possibles)

- Fiches détaillées par modèle Easy UPS (3S / 3M) avec menus pas à pas.
- Dictionnaire d'alarmes par marque (Schneider, Eaton, Vertiv).
- Modèle de contrat de maintenance UPS (clauses, SLA, bordereau de prix).
- Formation « habilitation électrique » : plan de cours complet.
- Guide groupes électrogènes approfondi (moteur, alternateur, ATS).
