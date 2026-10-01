---
id: collect-261001-rattrapage/rattrapage/huawei-licences-tac-guide-8
title: "Licences, support TAC et RMA Huawei — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-26", "2026-09-27"]
keywords: ["attention", "distribution", "license"]
source: docs/RAG/collect-261001-rattrapage/huawei_licences_tac_guide.md
source_anchor: ""
source_lines: [1159, 1348]
sha256: a6bf838880edd1aaffa715eed7bfcb22bf7c681f757681049b61bcd8cc1af025
---

# Licences, support TAC et RMA Huawei — Guide ultra-complet

5. Troubleshooting already performed
   - "display license" confirms expiry date 2026-09-26 for IPS/AV/URL features.
   - Rebooted the USG at 08:30 CET - no change.
   - Checked with partner: renewal order in progress, file expected tomorrow.

6. Attachments
   - diagnostic-information output: diag_USG6680_20260927.txt
   - network topology diagram: topo_paris_site.pdf

7. Requested action
   Please confirm whether a temporary license can be issued to cover the gap
   until the renewal file arrives, and provide the procedure.
```

*(Toutes les valeurs ci-dessus sont fictives : remplacez par vos données réelles.)*

## 72. Template rempli — exemple fictif n°2 (matériel)

```text
Subject: [S1] S310-24P - switch not booting, site Lyon warehouse isolated

1. Customer information
   Company: ACME Industries (fictitious)
   Contact name: Marie Martin (fictitious)
   Phone (24/7 reachable): +33 6 11 11 11 11 (fictitious)
   Email: marie.martin@example.com (fictitious)
   Time zone / availability: CET, available 24/7

2. Equipment information
   Product model: S310-24P (fictitious)
   Software version: unknown - device does not boot (last known: V200R021C00)
   ESN: 2102311ABC8H7000456 (fictitious)
   Serial Number: 2102311ABC8H7000456 (fictitious)
   Support contract: Hi-Care Premier 24x7x4, contract #HC-2026-4452 (fictitious)

3. Problem description
   Summary: Switch does not boot since a power outage this morning.
   First occurrence: 2026-09-27 06:05 CET (after mains power restoration)
   Frequency: permanent - stuck at boot, PWR LED red, SYS LED off.
   Error messages (exact text):
   Console shows "BootROM checksum error" then halts. (fictitious message)

4. Business impact
   Impacted users/services: entire Lyon warehouse (80 users, barcode scanners,
   printers) - no network access.
   Workaround in place: no - single switch on this site, no spare on site.
   Requested severity: S1 - full site isolated, critical for morning shipments.

5. Troubleshooting already performed
   - Power-cycled 3 times, tried another power outlet and power cord.
   - Console cable checked on another device - cable OK.
   - No recent configuration change (last change 3 weeks ago).

6. Attachments
   - console boot log: bootlog_S310_lyon.txt
   - photo of LEDs: leds_s310.jpg
   - site topology: topo_lyon_warehouse.pdf

7. Requested action
   Please diagnose and, if hardware failure confirmed, trigger RMA under
   Hi-Care Premier 24x7x4 (4-hour part delivery expected).
```

## 73. Ouvrir le ticket : pas à pas sur le portail

1. Connectez-vous à https://support.huawei.com → **Service Request / TAC**.
2. Sélectionnez le produit et l'équipement enregistré (section 49).
3. Choisissez la **sévérité** (sections 63–68) — soyez factuel.
4. Remplissez le formulaire en collant votre template (section 70).
5. Joignez les pièces : diagnostic-information, topologie, captures.
6. Validez : notez le **numéro de ticket** (format type `SR-XXXX...` —
   format exact **à vérifier sur le portail officiel**).
7. Informez votre équipe et votre astreinte : qui suit le ticket, à quel rythme.

## 74. Suivi du ticket : statuts et SLA

Cycle de vie typique d'un ticket :

```
Ouvert → Accepté par le TAC → En diagnostic → Contournement proposé →
Cause racine identifiée → Résolu → Clôturé (avec votre accord)
```

- **Réponse initiale** : selon la sévérité (30 min / 60 min / 2 h / NBD).
- **Mises à jour** : le TAC doit vous tenir informé ; si 24h sans nouvelles
  sur un S2, relancez (toujours dans le ticket, jamais « en parallèle » par
  un autre canal : un seul fil = pas de perte d'info).
- **Ne fermez jamais un ticket vous-même** tant que la cause racine n'est pas
  comprise ou qu'un contournement stable n'est pas validé en production.

## 75. L'escalade : quand et comment

Escalader = demander un niveau d'attention supérieur. Motifs légitimes :

| Motif | Action |
|---|---|
| SLA de réponse dépassé | Relance formelle dans le ticket + appel hotline en citant le n° de ticket |
| Diagnostic au point mort depuis 48h (S1/S2) | Demander l'intervention d'un expert produit / du R&D (via le TAC) |
| Impact métier aggravé | Requalifier la sévérité à la hausse (S3→S2) avec justification |
| Désaccord sur la cause | Demander une revue par un second ingénieur |

**Comment escalader proprement** (en anglais dans le ticket) :
```text
Escalation request: no progress since [date/time], business impact increasing
([describe]). Requesting senior engineer review and management attention.
Ticket: [number]. Contact 24/7: [name, phone].
```
Toujours factuel, jamais agressif : l'ingénieur en face est votre allié,
pas votre adversaire.

---

# F. COLLECTER LES INFOS POUR LE TAC + GARANTIE

## 76. `display diagnostic-information` : la pièce maîtresse

C'est **la** commande que le TAC vous demandera dans 9 cas sur 10. Elle
agrège en une seule sortie : version, configuration, états des interfaces,
logs, diagnostic matériel.

```
<USG6000> display diagnostic-information
```

- Redirigez vers un fichier : la sortie fait plusieurs milliers de lignes.
- Nommez le fichier : `diag_<modèle>_<site>_AAAAMMJJ.txt`.
- Collectez-la **pendant le problème** si possible (pas 3 jours après
  un reboot qui a tout effacé).

> Si la commande n'existe pas sur votre version, l'équivalent est une
> combinaison de `display version`, `display current-configuration`,
> `display logbuffer`, `display device` — demandez au TAC la liste exacte.

## 77. Exporter les logs : logbuffer et logfile

```
<USG6000> display logbuffer          ← logs en mémoire (volatils !)
<USG6000> display logfile            ← logs persistants (flash)
```

- Le **logbuffer** se vide au reboot : en cas de panne avec redémarrage,
  exportez-le **avant** de rebooter si l'équipement répond encore.
- Filtrez par période : repérez l'heure exacte du début de panne et
  fournissez ±30 minutes autour.
- Exportez en texte brut, pas en capture d'écran (le TAC doit pouvoir
  chercher dedans).

## 78. Les captures et schémas à fournir

| Élément | Format conseillé | Contenu |
|---|---|---|
| Topologie du site | PDF / PNG | Équipements, liens, adresses IP, VLAN |
| Face avant/arrière | Photo nette | LED allumées/éteintes, câblage |
| Étiquette S/N | Photo nette | S/N, modèle, date de fabrication |
| Message d'erreur | Texte copié-collé | Le texte exact, pas une paraphrase |
| Graphique de supervision | PNG | Trafic/CPU avant-pendant-après la panne |

Une photo floue de l'étiquette = un aller-retour de 24h avec le TAC.
Prenez 30 secondes pour la refaire nette.

## 79. Décrire une panne efficacement : la méthode

Structure imposée à votre équipe pour tout compte-rendu de panne :

1. **Quoi** : le symptôme observable (pas l'interprétation).
   - ✅ « Les utilisateurs du VLAN 20 n'obtiennent plus d'IP depuis 8h15. »
   - ❌ « Le DHCP est cassé. » (interprétation, peut-être fausse)
2. **Quand** : date/heure précise + fuseau, première occurrence.
3. **Où** : quel équipement, quel site, quel segment.
4. **Fréquence** : permanent / intermittent (toutes les X minutes) /
   lié à un événement (coupure électrique, mise à jour...).
5. **Changements récents** : config, firmware, câblage, coupure de courant
   dans les 7 derniers jours.
6. **Ce qui a été tenté** et le résultat de chaque action.

## 80. La règle des 5W adaptée au TAC

| W | Question | Exemple de réponse |
|---|---|---|
| What | Que se passe-t-il ? | Plus de distribution DHCP sur VLAN 20 |
| When | Depuis quand ? | 2026-09-27 08:15 CET, permanent |
| Where | Où ? | S310-BatA, Lyon, ports 1-12 |
| Who | Qui est impacté ? | 40 utilisateurs du bâtiment A |
| Why (piste) | Qu'est-ce qui a changé ? | Mise à jour VRP hier soir 22h |

Rédigez ces 5 lignes **avant** d'appeler : l'ingénieur TAC vous les
demandera de toute façon, autant les avoir prêtes.

---

