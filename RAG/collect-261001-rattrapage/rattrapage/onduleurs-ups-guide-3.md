---
id: collect-261001-rattrapage/rattrapage/onduleurs-ups-guide-3
title: "Onduleurs / UPS 10–120 kVA — Guide ultra-complet (exploitation & maintenance)"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/onduleurs_ups_guide.md
source_anchor: ""
source_lines: [369, 579]
sha256: 427f2c73e70060f4faaefaab229164b231d8246f1b0b18f74826f43272605da2
---

# Onduleurs / UPS 10–120 kVA — Guide ultra-complet (exploitation & maintenance)

- **Parallèle** : 2 à 6 UPS identiques (même modèle, même firmware)
  sur le même bus de sortie.
- **N+1** : N UPS portent la charge, 1 en réserve. Ex. : besoin
  60 kVA → 2× 60 kVA en N+1 (ou 3× 40 kVA).
- Avantages : maintenance **sans bypass** (on isole un UPS, les
  autres portent la charge), tolérance de panne.
- Exigences : câbles de **même longueur/impédance**, mise en service
  par procédure constructeur, supervision commune.
- ⚠️ Le parallèle ne pardonne pas l'approximation : une erreur de
  câblage = court-circuit entre onduleurs.

---

## 17. Monitoring et supervision

### 17.1 Carte réseau (NMC)
- **Network Management Card** (Schneider AP9640/41 ou équivalent) :
  web + SNMP + Modbus TCP.
- À chaque installation : IP fixe, **mot de passe changé**, SNMP
  configuré, alertes email/SMS.

### 17.2 Protocoles
- **SNMP** : supervision (Zabbix, PRTG, Centreon) — OID standard
  UPS-MIB (RFC 1628).
- **Modbus RTU/TCP** : intégration GTB/BMS du bâtiment.
- **Contacts secs** : report d'alarmes simple (défaut, sur batterie,
  bypass) vers une alarme technique.

### 17.3 EcoStruxure IT (Schneider) / équivalents
- Supervision cloud : alertes, historique, recommandations.
- Pour un chef de service : **un tableau de bord unique** pour tout
  le parc = vous voyez les batteries qui faiblissent **avant**
  la panne.

### 17.4 Ce qu'il faut superviser (minimum)
- État (normal / sur batterie / bypass / défaut).
- Charge % par phase.
- Tension batterie + température.
- Alarmes actives.
- Autonomie restante estimée.

---

## 18. Alarmes et défauts — diagnostic

> Les libellés exacts varient selon les marques. Notez le **code
> complet** affiché + l'historique (event log).

### 18.1 Famille BATTERIES (la plus fréquente)
- **Batterie déconnectée** : disjoncteur ouvert, fusible, câble.
- **Tension batterie basse** : fin d'autonomie ou batterie HS.
- **Test batterie échoué** : impédance trop haute → prévoir
  remplacement (§20).
- **Remplacer batterie** : fin de vie estimée atteinte.
- **Surchauffe batterie** : local trop chaud → agir vite (risque
  d'emballement thermique).

### 18.2 Famille RÉSEAU / ENTRÉE
- **Réseau hors tolérance** : tension/fréquence hors plage → l'UPS
  passe sur batterie (vérifier le réseau, pas l'UPS !).
- **Séquence de phases incorrecte** : câblage d'entrée à corriger.
- **Neutre manquant** : **dangereux** → vérifier immédiatement.

### 18.3 Famille ONDULEUR / BYPASS
- **Surcharge** : > 100 % → délester ou augmenter la puissance.
- **Défaut onduleur** : bascule sur bypass statique → diagnostic
  (surchauffe ? court-circuit en sortie ?).
- **Bypass hors tolérance** : le bypass ne peut pas prendre le
  relais → situation critique, intervenir.
- **Court-circuit en sortie** : l'UPS protège, puis bypass → trouver
  le défaut aval.

### 18.4 Famille DIVERS
- **Défaut ventilateur** : remplacer (surchauffe sinon).
- **Surchauffe** : clim du local, dépoussiérage, ventilateurs.
- **Bus DC en surtension** : carte à diagnostiquer.
- **Défaut carte** : noter le code, appeler le support avec l'event log.

### 18.5 Méthode générale
```
1. Noter code + heure + circonstances (orage ? travaux ?)
2. Consulter l'event log (historique des 50–200 derniers événements)
3. Classer : batterie / réseau / onduleur / environnement
4. Traiter la cause, pas le symptôme
5. Tester (coupure simulée) après intervention
```

---

## 19. Maintenance préventive — le programme total

### 19.1 Visite mensuelle (30 min)
```
□ Relevé : état, charge %, alarmes actives
□ Event log : nouveaux événements ?
□ Température local + armoire batteries
□ Ventilation : grilles propres, flux d'air
□ Voyants : tout au vert ?
□ Supervision : alertes reçues ?
```

### 19.2 Visite trimestrielle (2 h)
```
□ Tout le mensuel +
□ Tension de chaque bloc batterie (écart > 0,3 V = suspect)
□ Serrage des connexions batteries (au couple, machine consignée)
□ Dépoussiérage (aspirateur, jamais d'air comprimé sur les cartes)
□ Thermographie IR (points chauds) — §22
□ Vérifier les seuils d'alarmes et la supervision
```

### 19.3 Visite semestrielle (demi-journée)
```
□ Tout le trimestriel +
□ Test d'impédance batteries — §20
□ Test des ventilateurs (écoute, débit)
□ Test bypass statique + bypass de maintenance
□ Vérifier les condensateurs (visuel : gonflés ? coulures ?)
□ Nettoyage filtres à air
```

### 19.4 Visite annuelle (journée complète)
```
□ Tout le semestriel +
□ Test de décharge batterie (autonomie réelle) ou banc de charge — §23
□ Vérification des protections (déclenchement, sélectivité)
□ Mesure de terre
□ Mise à jour firmware (avec procédure de secours)
□ Rapport annuel complet + plan de remplacement (batteries ? ventilos ? condos ?)
```

---

## 20. Batteries — test et remplacement

### 20.1 Test d'impédance (la méthode pro)
- Appareil : testeur d'impédance/conductance (Hioki, Megger BITE…).
- **Baseline à la mise en service** : mesure de chaque bloc neuf.
- Suivi : **+20–30 %** par rapport à la baseline = bloc à
  remplacer. **+50 %** = remplacement urgent.
- Fréquence : semestrielle (trimestrielle après 3 ans).

### 20.2 Test de tension
- Floating : chaque bloc à 13,5–13,6 V (à 20 °C).
- Écart > 0,3 V entre blocs = bloc faible → test d'impédance ciblé.

### 20.3 Test de décharge (autonomie réelle)
- Couper l'entrée, chronométrer jusqu'à l'alarme « batterie basse ».
- ⚠️ Ne jamais décharger à fond (arrêter à 1,75 V/élément).
- Si l'autonomie < 80 % de la nominale → remplacement.

### 20.4 Remplacement — procédure
1. Commander des blocs **identiques** (même marque, même lot si
   possible).
2. Consignation : disjoncteur batterie **ouvert**, vérifier
   l'absence de tension.
3. Remplacer **tout le string** (jamais un seul bloc neuf parmi
   des vieux).
4. EPI : gants, lunettes, outils isolés, pas de bijoux.
5. Vérifier la tension totale avant de refermer le disjoncteur.
6. Nouvelle **baseline d'impédance**.
7. Anciens blocs : **recyclage** (filière plomb, §32).

---

## 21. Condensateurs et ventilateurs — les oubliés

### 21.1 Condensateurs du bus DC
- Vie : **5–7 ans** (moins en chaud). Signes : gonflement,
  coulures, évent ouvert.
- Un condensateur qui lâche = **explosion possible** + carte HS.
- **Remplacement préventif à 5 ans** sur les sites critiques
  (kit constructeur, machine consignée, attendre la décharge du
  bus — 5–10 min — et **vérifier au multimètre**).

### 21.2 Ventilateurs
- Vie : **3–5 ans** en continu. Signes : bruit, débit faible,
  alarme « fan fault ».
- Remplacement préventif à 4 ans (pièce peu chère, panne coûteuse).
- ⚠️ Sens de montage : noter avant démontage (U-autocollant).

---

## 22. Thermographie infrarouge

- Caméra IR (ou smartphone + module) : **l'outil du chef de service**.
- Points à scanner : connexions batteries, disjoncteurs, cosses,
  borniers, transformateurs.
- **Écart > 10 °C** entre phases ou par rapport à l'ambiance =
  anomalie (serrage, surcharge, contact).
- Fréquence : trimestrielle. Photos **avant/après** dans le rapport.
---

## 23. Test sur banc de charge (load bank)

- **Pourquoi** : tester l'UPS à 100 % sans risquer la charge réelle.
- **Quand** : mise en service 60 kVA+, après grosse intervention,
  annuellement sur sites critiques.
- **Procédure** :
  1. Basculer la charge réelle sur bypass/groupe (ou tester à vide
     de charge avec le banc seul).
  2. Brancher le banc sur la sortie UPS.
  3. Monter par paliers (25/50/75/100 %), 15 min par palier.
  4. Surveiller : températures, tensions, alarmes, autonomie batterie.
  5. Redescendre, rebasculer la charge, rapport.
- Location de bancs : prévoir dans le budget maintenance.

---

## 24. Sécurité — consignation, DC, batteries, arc flash

