---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-23
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter", "distribution", "gpu"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [3316, 3486]
sha256: 369f66282c7fdcac68874df77130e08025902693670e3e1276f1f9176a1798c1
---

# RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE

**Whitebox (Edgecore, Celestica, FS)** : TH5 à −30-50 % vs marques (⚠️),
SONiC ou ONIE+EOS. Verdict : le moins cher au port — à condition d'avoir
l'équipe qui opère SONiC (ou une distribution durcie). Pas pour une
première IA.

**FS.com (optiques et câbles)** : le fournisseur « compatible » de
référence pour optiques tierces codées — 3-5× moins cher que l'OEM (⚠️).
Verdict : valider 1 référence par switch/NIC, puis standardiser. C'est ici
que se gagnent les 30 % du budget optique.

## 160. CAS CHIFFRÉ — refresh 100G → 400G d'un DC existant (500 serveurs)

**Hypothèses** : 500 serveurs 2×25G, leafs 48×25G+8×100G, on passe les
uplinks et les spines en 400G, serveurs inchangés (breakout).

| Étape | Action | Coût (⚠️) |
|---|---|---|
| 1. Fibre | Vérifier/ajouter OS2 inter-rangées | 30-80 k€ |
| 2. Spines | 2× 32×400G (remplacent 2× 32×100G) | 80-160 k€ |
| 3. Uplinks | 44 leafs × 4×100G → 2×400G : 88 optiques 400G-DR4 | 53-132 k€ |
| 4. Breakout | 88 câbles 400G→4×100G côté leaf existant | 18-44 k€ |
| 5. Main d'œuvre | 2 techs × 2 semaines + tests | 30-50 k€ |
| **Total** | | **~211-466 k€** |

**Gain** : uplinks ×4 (de 400G à 1,6T par leaf en 400G), oversubscription
3:1 → 0,75:1 (surprovisionné — on pourra densifier les serveurs).
**Leçon** : le refresh se fait **par le haut** (spines/uplinks d'abord),
les serveurs suivent. Et la fibre d'abord (§113).

## 161. Topologie multi-tenant cloud — l'overlay en détail

```
  Tenant A (VNI 10001)          Tenant B (VNI 10002)
  ┌──────────────┐              ┌──────────────┐
  │ 10.1.0.0/16  │              │ 10.1.0.0/16  │  ← mêmes IP possibles !
  │ (isolé)      │              │ (isolé)      │
  └──────┬───────┘              └──────┬───────┘
         │ VTEP Leaf-1                │ VTEP Leaf-2
  ───────┴──────── underlay BGP ───────┴────────
```

- **Isolation** : VNI distincts, pas de fuite inter-VNI (vérifié par test).
- **Passerelle** : anycast par VNI (1 IP gateway par tenant, même MAC).
- **Sortie** : border leafs avec NAT/firewall par tenant.
- **Quotas** : rate-limit par VNI sur les border leafs (un tenant ne noie
  pas les autres).
- **Supervision** : compteurs par VNI (pas seulement par port).

**Dimensionnement** : 16M de VNI possibles, mais la table MAC/FIB du leaf
est la vraie limite (136K MAC sur un Arista 7060X, ⚠️) — ~1000 VMs/leaf
en pratique avec marge.

## 162. Checklists de mise en service — jour J

**J-7** : [ ] spares reçus et testés [ ] FW validés au lab [ ] golden
config gelée [ ] plan d'adressage dans NetBox [ ] équipe briefée
[ ] rollback préparé [ ] fenêtre de maintenance déclarée.

**Jour J** : [ ] 1 baie à la fois [ ] câblage → étiquetage → photo
[ ] 1er lien : BER + DOM avant de continuer [ ] BGP/BFD up
[ ] tests §107 par palier (pas tout à la fin) [ ] supervision branchée
**avant** la charge.

**J+1** : [ ] burn-in 72 h lancé [ ] alertes vérifiées (en faire sonner
une pour de faux) [ ] doc à jour [ ] REX à chaud (15 min, ce qui a
coincé).

**Règle** : on ne met en service **que ce qu'on a testé**, et on ne teste
**que ce qu'on a documenté**. Dans cet ordre.

## 163. 10 erreurs de câblage illustrées

**E1. Tx vers Tx** (deux modules émettent l'un vers l'autre) :
```
  [Module A Tx] ──fibre── [Module B Tx]   ❌ pas de link
  [Module A Tx] ──fibre── [Module B Rx]   ✅ (croiser !)
```

**E2. MPO mâle-mâle** : broches tordues, à jeter des 2 côtés.

**E3. OM4 sur port DR (monomode)** : le laser 1310 nm dans du multimode =
pertes énormes, lien instable.

**E4. OS2 sur port SR (850 nm)** : le VCSEL dans du monomode = ça peut
« marcher » sur 2 m puis mourir — trompeur.

**E5. Jarretière de 30 m enroulée en boule** : micro-courbures + 2 dB.

**E6. DAC plié à 90°** : le Twinax n'aime pas — faux contacts intermittents.

**E7. MPO-12 sur port MPO-16** : 4 fibres dans le vide, lanes manquantes.

**E8. Breakout branché sur port non-breakout** : 1 lane up, 7 down.

**E9. Fibre sous le rail de la baie** : écrasée à la fermeture — BER qui
monte avec la température (dilatation).

**E10. Deux trunks méthode B bout-à-bout** : B+B = droit au lieu de croisé
→ Tx vers Tx (retour à E1). **Un seul** élément croisé par lien.

## 164. Conversions dB/dBm/W — l'antisèche

| Conversion | Formule / valeur |
|---|---|
| dBm → mW | P(mW) = 10^(dBm/10) |
| 0 dBm | 1 mW |
| −3 dBm | 0,5 mW |
| −10 dBm | 0,1 mW |
| −20 dBm | 0,01 mW |
| 3 dB | ×2 (ou ÷2) |
| 10 dB | ×10 (ou ÷10) |
| Budget 4 dB | = on peut perdre 60 % de la puissance |
| Marge 3 dB | = la moitié de la puissance en réserve |

**Exemple** : Tx −2 dBm (0,63 mW), perte 4 dB → Rx −6 dBm (0,25 mW).
Si la sensibilité est −8 dBm : marge 2 dB → **limite** (voulu : ≥ 3 dB).

---

*Fin du guide — 164 sections, 40+40 termes de glossaire, 45 pièges,
10 questions de quiz, 30 FAQ. Vérification : `wc -l` — objectif ≥ 4000 lignes.*

# PARTIE P — DERNIERS COMPLÉMENTS

## 165. Étude de cas — datacenter IA 10 MW : le réseau de bout en bout

**Hypothèses** : 10 MW IT, dont 8 MW GPU (512 nœuds × 8 B200, ⚠️),
1 MW réseau, 1 MW stockage+divers.

**Réseau backend** (1:1, 800G, rail-optimized) :
- 4096 NIC 800G (8/GPU × 512 nœuds).
- 2 tiers : 64 leafs + 128 spines 64×800G ? Non — 4096/64 = 64 leafs
  (64 down + 64 up... en 2×400G). Prenons des 64×800G : 64 leafs,
  64 spines → **128 switchs**, 8192 ports 800G.

**Bilan énergie** (⚠️) :

| Poste | Calcul | kW |
|---|---|---|
| GPU (8 MW) | 512 × 15,6 kW | 8000 |
| NIC 800G | 4096 × 60 W | 246 |
| Switchs backend | 128 × 1,5 kW | 192 |
| Optiques backend | ~6000 × 15 W | 90 |
| Front-end (3:1, 100G) | — | 60 |
| Stockage NVMe-oF | — | 150 |
| Divers (OOB, mgmt) | — | 20 |
| **Total IT** | | **~8758** |
| Au compteur (PUE 1,35) | 8758 × 1,35 | **~11 800** |

**Le réseau (528 kW IT) = 6 % de l'IT** — mais 100 % du risque : sans lui,
les 8 MW de GPU sont des radiateurs à 40 000 € pièce.
**Refroidissement** : ~3,2 MW thermiques à évacuer (PUE 1,35 → 35 % =
pertes+clim). Le lot réseau seul (528 kW) = **~150 tonnes** frigorifiques.

## 166. Tester les optiques à la réception — procédure

1. **Visuel** : emballage, bouchons présents, étiquette (référence, SN).
2. **Inspection** : microscope — PASS/FAIL IEC 61300-3-35 (§25).
3. **Codage** : enficher dans le switch/NIC cible → `show inventory` :
   la référence lue = la référence commandée (détecte les contrefaçons).
4. **Boucle** : loopback optique (atténuateur si besoin) → BER 0 sur
   1e12 bits, puissances Tx/Rx dans la spec.
5. **Échantillon** : 100 % des 10 premiers, puis 10 % par lot (min 5).
   Un lot avec > 2 % de rejet = lot refusé.

**Outillage** : 1 microscope, 2 atténuateurs variables, 1 jeu de
loopbacks LC/MPO, 1 power-meter. ~3 000-5 000 € (⚠️) — amorti au premier
lot défectueux détecté **avant** montage en baie.

## 167. Distances fibre — table complète par norme (⚠️ usuels)

