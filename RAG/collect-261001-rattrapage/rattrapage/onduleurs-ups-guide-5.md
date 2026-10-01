---
id: collect-261001-rattrapage/rattrapage/onduleurs-ups-guide-5
title: "Onduleurs / UPS 10–120 kVA — Guide ultra-complet (exploitation & maintenance)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/onduleurs_ups_guide.md
source_anchor: ""
source_lines: [781, 997]
sha256: 33d631a9ddab428841248304b1e7d1f1611d59963edf01045f4daab13ab5ed7f
---

# Onduleurs / UPS 10–120 kVA — Guide ultra-complet (exploitation & maintenance)

| Terme | Signification |
|---|---|
| UPS / onduleur | Alimentation sans interruption |
| VFI / VI / VFD | Topologies (IEC 62040-3) |
| Double conversion | AC→DC→AC en permanence |
| PF | Facteur de puissance (kW/kVA) |
| THDi | Distorsion harmonique du courant |
| Bus DC | Liaison continue interne (~384–480 V) |
| Floating | Charge d'entretien des batteries |
| Boost | Charge accélérée/égalisation |
| Bypass statique | Bascule électronique auto sur réseau |
| Bypass maintenance | Bascule manuelle pour intervention |
| N+1 | Redondance : 1 UPS de plus que nécessaire |
| NMC | Carte réseau de supervision |
| SNMP / Modbus | Protocoles de supervision |
| Autonomie | Temps sur batteries |
| Load bank | Banc de charge (test) |
| Arrhenius | Loi : +10 °C = vie batterie / 2 |
| GTR | Garantie de temps de rétablissement |
| MTTR | Temps moyen de réparation |

---

## 34. Pense-bête de poche

```
VFI = 0 ms | kW = kVA × PF | Charge cible ≤ 80 %
Floating 13,5–13,6 V/bloc | Écart > 0,3 V = suspect
Impédance +30 % = remplacer | Batteries : local 20–25 °C
Bus DC chargé après coupure : attendre + vérifier !
Groupe = 1,5–2× l'UPS | Test groupe mensuel EN CHARGE
Bypass = plus de protection | Thermographie trimestrielle
Condos 5 ans | Ventilos 4 ans | Firmware annuel
```

---

## 35. Fiches par puissance — l'essentiel en une page

### 35.1 10 kVA (~9 kW)
- Usage : petite salle serveurs, agence.
- Protection : 20–25 A, câble ~4 mm², terre < 5 Ω.
- Batteries : 32 blocs 12 V (384 V), autonomie 10–15 min standard.
- Vigilance : disjoncteur amont souvent sous-dimensionné, local
  non climatisé.

### 35.2 20 kVA (~18 kW)
- Usage : salle serveurs PME.
- Protection : 40 A, câble 6–10 mm².
- Batteries : 32–40 blocs, 1–2 strings.
- Vigilance : prévoir la croissance (souvent saturé en 2 ans).

### 35.3 40 kVA (~36 kW)
- Usage : data center PME, industrie.
- Protection : 80 A, câble 16–25 mm², note de calcul exigée.
- Batteries : 40 blocs (480 V) ou 32 (384 V), 1–2 strings.
- Vigilance : local dédié + clim, parafoudre, supervision NMC.

### 35.4 60 kVA (~54 kW)
- Usage : data center moyen, hôpital.
- Protection : 125 A, câble 35 mm², étude de sélectivité.
- Batteries : armoire(s) dédiée(s), 2 strings conseillés.
- Vigilance : mise en service avec banc de charge, groupe
  90–120 kVA, clim N+1.

### 35.5 120 kVA (~108 kW)
- Usage : gros site, N+1 de 60 kVA.
- Protection : 250 A, câble 70–95 mm², TGBT dédié.
- Batteries : plusieurs armoires, 2+ strings, local batteries
  ventilé.
- Vigilance : plancher (1–2 t), étude complète, contrat Gold,
  équipe formée, stock pièces.

---

## 36. Questions pièges (validez vos connaissances)

1. Pourquoi 0 ms en double conversion ? → La charge est toujours
   sur l'onduleur, pas de bascule.
2. Un 40 kVA PF 0,9 alimente combien de kW ? → 36 kW.
3. Charge à 95 % en permanence : problème ? → Oui, pas de marge
   (pics, croissance) → viser 80 %.
4. Local batteries à 35 °C : conséquence ? → Vie / 2 (Arrhenius).
5. Pourquoi 2 strings identiques ? → Sinon déséquilibre de
   décharge, le string faible limite l'ensemble.
6. On remplace un seul bloc gonflé ? → Non, tout le string.
7. Le bus DC après consignation ? → Reste chargé : attendre +
   vérifier.
8. Groupe de même puissance que l'UPS : OK ? → Non, 1,5–2×.
9. En bypass de maintenance, la charge est protégée ? → Non.
10. Parafoudre : où ? → Tête d'installation, Type 1+2.

---

*Fin du guide. La maintenance des UPS tient en une phrase :
**batteries surveillées, local climatisé, protections calibrées,
tests réguliers — et tout noté.** Avec ça, votre parc ne vous
fera jamais de surprise un vendredi à 18h.*
---

## 37. Schémas unifilaires types

### 37.1 Schéma 10–20 kVA (simple)
```
Réseau ──[DJ 25/40A]──┬──► ENTRÉE UPS ──► SORTIE UPS ──[DJ]──► Charge
                      └──► BYPASS UPS ──► (interne)
Terre ──► PE ──► châssis UPS + armoire batteries
Batteries ──[DJ DC]──► UPS
```

### 37.2 Schéma 40–120 kVA (avec TGBT ondulé)
```
Réseau ──[DJ général]──[Parafoudre T1+T2]──┬──► ENTRÉE UPS ──►┐
                                           └──► BYPASS UPS ──►┤
                                                              ▼
                                              TGBT ONDULÉ ──[DJ départs]──► Charges
                                              (sectionneur général + report alarme)
Batteries ──[DJ DC par string]──► UPS
Groupe ──[Inverseur de source]──► (amont UPS)
```

### 37.3 Schéma N+1 (2× 60 kVA)
```
Réseau ──┬──[DJ]──► UPS1 ──┐
         └──[DJ]──► UPS2 ──┴──► Bus commun ──[DJ]──► Charge
Bypass externe de maintenance (contourne les 2 UPS)
```
- Câbles de sortie **strictement identiques** (longueur, section).
- Un seul bypass de maintenance externe pour l'ensemble.

---

## 38. Note de calcul électrique — modèle

> Document exigé pour 40 kVA et +. Modèle de trame :

```
1. OBJET : alimentation ondulée site ___, UPS ___ kVA
2. HYPOTHÈSES : 400 V tri, cos φ ___, température ___°C,
   mode de pose ___, longueur ___ m
3. BILAN DE PUISSANCE
   Charge 1 : ___ kW / cos φ ___ → ___ kVA
   ...
   TOTAL : ___ kW / ___ kVA → UPS retenu : ___ kVA (charge ___%)
4. COURANTS : I = S/(√3×U) = ___ A
5. CÂBLES : section ___ mm², chute de tension ___ % (< 5 %)
6. PROTECTIONS : DJ ___ A courbe ___, pouvoir de coupure ___ kA
7. SÉLECTIVITÉ : étude amont/aval (tableau)
8. TERRE : < 5 Ω, section PE ___ mm²
9. BATTERIES : ___ blocs ___ V ___ Ah, ___ strings,
   autonomie ___ min à ___ kW
10. GROUPE : ___ kVA (1,5–2× UPS)
```

---

## 39. Consignation détaillée — pas à pas

```
AVANT :
□ Prévenir le client / l'astreinte (la charge passe en bypass)
□ EPI : gants isolants, lunettes, écran facial si TGBT
□ Outillage isolé 1000 V, VAT vérifié

CONSIGNATION :
1. Forcer le bypass statique (menu UPS)
2. Basculer le bypass de maintenance → charge sur réseau direct
3. Ouvrir : DJ sortie → DJ entrée → DJ bypass → DJ batterie(s)
4. Cadenas + étiquette sur chaque organe
5. ATTENDRE 10 min (décharge bus DC)
6. VAT : vérifier absence de tension AC (entrée/sortie)
   ET DC (bus, batteries)
7. Si intervention batteries : déconnecter un pôle du string

DÉCONSIGNATION (ordre inverse) :
1. Reconnecter batteries, vérifier tension totale
2. Retirer cadenas/étiquettes
3. Fermer : DJ batterie → DJ bypass → DJ entrée
4. Démarrer l'onduleur, vérifier 400 V/50 Hz
5. Fermer DJ sortie → repasser en mode normal
6. Quitter le bypass de maintenance
7. TEST : coupure simulée 30 s → OK
```

---

## 40. Mise en parallèle — procédure détaillée

1. Vérifier : **mêmes modèles, mêmes firmwares**, cartes parallèles
   installées.
2. Câbles de sortie **identiques** (même longueur au cm près).
3. Adressage : UPS1 = maître, UPS2 = esclave (menu).
4. Démarrage séquentiel selon la procédure constructeur.
5. Test : couper UPS1 → UPS2 prend 100 % sans coupure.
6. Supervision : les 2 UPS dans le même tableau de bord.
7. ⚠️ Ne jamais paralléliser des UPS de marques/modèles
   différents.

---

## 41. Supervision SNMP — mise en pratique

- **UPS-MIB (RFC 1628)** : OID standard — état, charge, batterie,
  autonomie, température.
- **Zabbix/PRTG/Centreon** : template « UPS » + seuils :
  - Charge > 80 % → warning ; > 90 % → critique.
  - Sur batterie → alerte immédiate (SMS).
  - Batterie basse → critique.
  - Température > 30 °C → warning.
- **Trap SNMP** : l'UPS pousse les alarmes (plus rapide que le poll).
- Testez les alertes **après chaque installation** (coupure simulée
  → le SMS arrive-t-il ?).

---

## 42. Arrêt programmé des serveurs (graceful shutdown)

