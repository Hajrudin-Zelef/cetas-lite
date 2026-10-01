---
id: collect-261001-rattrapage/rattrapage/onduleurs-ups-guide-6
title: "Onduleurs / UPS 10–120 kVA — Guide ultra-complet (exploitation & maintenance)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["agent", "arr"]
source: docs/RAG/collect-261001-rattrapage/onduleurs_ups_guide.md
source_anchor: ""
source_lines: [998, 1207]
sha256: fe3ac7d79bec7e79aefb97de9662a5cd360aabc47705736ba5517888cc85530e
---

# Onduleurs / UPS 10–120 kVA — Guide ultra-complet (exploitation & maintenance)

- Si l'autonomie s'épuise (pas de groupe, coupure longue), les
  serveurs doivent s'arrêter **proprement** avant la fin.
- **PowerChute** (Schneider/APC) ou équivalent : agent sur chaque
  serveur, scénario « à X min d'autonomie restante → shutdown ».
- Ordre d'arrêt : applicatifs → bases de données → hyperviseurs.
- **Tester 2× par an** (un scénario jamais testé ne marche pas).
- Documenter : qui s'arrête, dans quel ordre, qui redémarre
  (séquence inverse au retour).

---

## 43. Climatisation — dimensionnement

```
Puissance frigo ≈ pertes UPS + pertes charges + apports
Pertes UPS = P_charge × (1 - η)
Ex. : 30 kW à 96 % → 1,2 kW de pertes (chaleur)
      + serveurs 30 kW → ~31 kW à évacuer
```
- **Redondance N+1** pour 40 kVA et +.
- Consigne : **22 °C** (±2). Alarme à 28 °C, critique à 32 °C.
- Maintenance clim : filtres mensuels, contrôle annuel (c'est la
  clim qui protège vos batteries !).
- En cas de panne clim : procédure d'urgence (ouvrir, ventiler,
  délester) — à écrire **avant**.

---

## 44. Groupes électrogènes — dimensionnement détaillé

```
P_groupe ≥ 1,5 × S_UPS (minimum)
P_groupe ≥ 2 × S_UPS (recommandé, charges non linéaires)
Ex. : UPS 60 kVA → groupe 90–120 kVA
```
- **ATS** (inverseur automatique) : bascule réseau/groupe < 10 s.
- Réglages : slew rate (rampe de fréquence) compatible UPS,
  tension et fréquence stables avant couplage.
- **Test mensuel en charge** (30 min à > 30 % de charge) :
  un groupe qui tourne à vide s'encrasse.
- Carburant : stock pour **48 h** minimum, rotation tous les 6 mois,
  contrat de livraison d'urgence.
- Échappement, bruit (capot insonorisé), voisinage : à traiter à
  l'installation.

---

## 45. Onduleurs modulaires (hot-swap)

- Principe : modules de puissance (ex. 25–50 kW) enfichables à chaud
  dans un châssis.
- Avantages : redondance interne N+1, évolutivité (on ajoute des
  modules), MTTR faible (échange en 10 min).
- Marques : Socomec MODULYS, Huawei UPS5000, Schneider Galaxy VS
  (partiellement).
- Maintenance : **mêmes règles** (batteries, condos, ventilos) +
  vérification des connecteurs de modules.

---

## 46. Qualité du réseau — savoir mesurer

| Perturbation | Effet sur l'UPS | Action |
|---|---|---|
| Creux de tension | Passage sur batterie | Enregistreur, seuils |
| Surtension | Bypass / protection | Parafoudre, signalement |
| Harmoniques (THDv) | Échauffement | Filtres, UPS IGBT |
| Variation de fréquence | Refus du réseau (batterie) | Groupe à régler |
| Coupures brèves | 0 ms (VFI) → invisible | Rien (c'est le job) |
| Déséquilibre phases | Alarme | Rééquilibrer les départs |

- **Enregistreur de qualité réseau** (Fluke 1750 ou équivalent) :
  l'outil qui prouve que « c'est le réseau, pas l'UPS ».

---

## 47. Audit énergétique d'un site — trame

```
□ Relevé des compteurs (12 mois) : kWh, pics
□ Mesure en charge (pince wattmètre) : par départ
□ Facteur de puissance global (pénalités ? → batterie de condensateurs)
□ Qualité réseau (enregistreur 1 semaine)
□ Climatisation : puissance, âge, redondance
□ Groupe : puissance, tests, carburant
□ UPS existants : âge, charge %, batteries, alarmes
□ Recommandations chiffrées : ROI de chaque action
```
- L'audit est une **prestation vendable** qui amène les contrats.

---

## 48. Gestion des incidents majeurs — plan de crise

```
1. DÉTECTION : alarme supervision → astreinte (15 min max)
2. QUALIFICATION : UPS ? réseau ? groupe ? (arbre de décision §25)
3. COMMUNICATION : client prévenu (délai, impact estimé)
4. ACTION : procédure (bypass, groupe, délestage…)
5. SUIVI : point toutes les 30 min jusqu'à résolution
6. RETEX : rapport 48 h — cause racine, actions correctives
```
- **Exercice annuel** : coupure simulée avec le client (comme un
  exercice incendie).

---

## 49. Dossier d'exploitation — le classeur du site

Chaque site doit avoir (papier + numérique) :
- Schéma unifilaire **à jour**.
- Note de calcul, rapports de mise en service.
- Fiches techniques (UPS, batteries, groupe, clim).
- Procédures : consignation, bypass, arrêt d'urgence, crise.
- Historique : rapports de visite, mesures, remplacements.
- Contacts : astreinte, constructeur, fournisseur pièces.
- **Sans dossier d'exploitation, pas de maintenance sérieuse.**

---

## 50. Lexique étendu

| Terme | Signification |
|---|---|
| ATS | Inverseur automatique de source |
| TGBT | Tableau général basse tension |
| THDv / THDi | Distorsion harmonique tension/courant |
| Slew rate | Vitesse de variation de la fréquence |
| Floating / Boost | Régimes de charge batterie |
| Baseline | Mesure de référence (impédance) |
| NMC | Carte réseau UPS |
| Trap SNMP | Alarme poussée par l'équipement |
| DOD | Profondeur de décharge |
| SOC | État de charge |
| EODV | Tension de fin de décharge |
| Load bank | Banc de charge |
| RETEX | Retour d'expérience |
---

## 51. Fiche réflexe astreinte (à garder sur le téléphone)

```
ALARME UPS — CONDUITE À TENIR
1. Lire le code EXACT + l'heure
2. État ? ☐ Normal ☐ Sur batterie ☐ Bypass ☐ Défaut
3. Charge ___ % | Batterie ___ V | Autonomie ___ min
4. Cause probable (§18) : _______________
5. Si sur batterie : autonomie restante ? groupe démarré ?
6. Si défaut : bypass possible ? charge protégée ?
7. Appeler : client ___ / senior ___ / constructeur ___
8. Noter TOUT (heure, actions) → fiche d'intervention
NE JAMAIS : forcer un redémarrage sans diagnostic,
            laisser en bypass sans prévenir,
            intervenir seul sur TGBT/HTA.
```

---

## 52. Erreurs classiques des débutants (et comment les éviter)

1. **Oublier le disjoncteur batterie ouvert** → l'UPS ne tient pas
   la coupure. → Checklist §13.
2. **Serrer les cosses batteries à la main** → échauffement, faux
   contact. → Clé dynamométrique + thermographie.
3. **Mélanger blocs neufs et vieux** → le string est limité par le
   plus faible. → Remplacement complet du string.
4. **Négliger la clim** → batteries mortes en 2 ans. → 20–25 °C,
   clim N+1.
5. **Firmware jamais mis à jour** → bugs, failles. → Annuel.
6. **Pas de test de coupure** → on découvre la panne le jour J. →
   Test à chaque visite annuelle.
7. **Mot de passe NMC perdu** → supervision morte. → Coffre.
8. **Bypass de maintenance oublié** → plus de protection. → Alarme
   supervisée + checklist.
9. **Intervenir sans consignation** → **danger de mort**. → §24,
   toujours.
10. **Ne rien noter** → pas de traçabilité, pas d'amélioration. →
    Fiche systématique.

---

## 53. Budget maintenance type (ordre de grandeur annuel)

| Poste | % du prix de l'UPS |
|---|---|
| Visites préventives (MO) | 3–5 % |
| Pièces d'usure (ventilos, filtres) | 1–2 % |
| Provision batteries (amorti /5 ans) | 8–12 % |
| Supervision / astreinte | 2–3 % |
| Banc de charge (location annuelle) | 1 % |
| **Total** | **15–23 %** |

- C'est ce chiffrage qui justifie vos contrats : montrez-le au
  client, il comprend ce qu'il paie.

---

## 54. Normes et références (à connaître)

- **CEI 62040-1/2/3** : sécurité / CEM / performances des UPS.
- **CEI 60364 / NF C 15-100** : installations électriques BT.
- **NF C 18-510** : habilitation électrique (référence).
- **EN 50272-2** : sécurité des batteries stationnaires.
- **Eurobat** : classification des durées de vie batteries.
- Conservez les **déclarations de conformité** de chaque UPS.

---

## 55. Ce qui change (2024-2026)

