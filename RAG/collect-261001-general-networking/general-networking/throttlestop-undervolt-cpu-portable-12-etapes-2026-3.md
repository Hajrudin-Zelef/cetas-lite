---
id: collect-261001-general-networking/general-networking/throttlestop-undervolt-cpu-portable-12-etapes-2026-3
title: "Verifier les erreurs materielles WHEA (PowerShell, admin)"
domain: general-networking
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["agent", "attention", "gpu", "intel"]
source: docs/RAG/collect-261001-general-networking/throttlestop-undervolt-cpu-portable-12-etapes-2026.md
source_anchor: ""
source_lines: [120, 203]
sha256: b0af049c25e0e786fca950b5dc08c705b67f0c217cb87d0162817c255231ce9c
---

# Verifier les erreurs materielles WHEA (PowerShell, admin)

Cliquez sur **Apply** pour appliquer l’offset en direct, puis sur **OK – Save voltages immediately**. Cette option est capitale : elle réapplique automatiquement l’undervolt à chaque réveil de veille. Sans elle, de nombreux portables « oublient » l’offset après une mise en veille, et le CPU repart à pleine tension sans que vous vous en aperceviez. Le tableau ci-dessous propose des paliers d’undervolt indicatifs selon le profil de silicium.

| Palier | Offset Core + Cache | Profil de puce | Gain thermique typique | 
|---|---|---|---|
| Découverte | -50 mV | Toutes, ultra-sûr | -3 à -6 °C | 
| Prudent | -80 mV | Point de départ recommandé | -6 à -10 °C | 
| Standard | -100 à -125 mV | Bon silicium (fréquent) | -10 à -15 °C | 
| Agressif | -130 à -160 mV | Très bon silicium | -15 à -20 °C | 
| Extrême | > -180 mV | Loterie du silicium, rare | > -20 °C, souvent instable | 

### Étape 8 – Vérifier que l’offset est bien appliqué

Ouvrez HWiNFO en mode capteurs et cherchez la ligne « IA Offset » ou « Core VID ». La valeur doit refléter votre offset (-80 mV). Si elle reste à 0, l’undervolt n’est pas actif : soit il est verrouillé, soit vous avez oublié de cocher *Unlock Adjustable Voltage* sur les deux rails. Relancez ensuite un TS Bench court pour confirmer la baisse de température :

```
=== HWiNFO – verification de l'offset ===
CPU Core Voltage (SVID)   : 1,168 V (etait 1,248 V)
IA Offset                 : -80,1 mV  <-- offset applique
CPU Package Temperature   : 88 C (etait 100 C)
Core Clocks (moyenne)     : 3,58 GHz (etait 3,12 GHz)
Thermal Throttling        : NON
=> undervolt actif et stable
```
## Étapes 9 à 12 : tester la stabilité et sauvegarder le profil

### Étape 9 – Descendre par paliers de -10 mV

Votre -80 mV tient ? Descendez à -90 mV, appliquez, testez. Puis -100 mV, et ainsi de suite, par pas de -10 mV. À chaque palier, lancez un TS Bench 960M complet. L’objectif est de trouver le point de rupture, puis de remonter d’un cran pour la marge de sécurité. La règle d’or : un undervolt stable au TS Bench mais qui plante en jeu est un undervolt *trop agressif*. Gardez toujours 10 à 20 mV de marge sous votre plus bas offset stable.

### Étape 10 – Valider la stabilité en profondeur

Le TS Bench est un bon filtre rapide, mais la vraie validation demande une charge longue et variée. Enchaînez un Cinebench R23 en boucle 30 minutes, puis une vraie session de jeu de 1 à 2 heures. Surveillez trois signaux d’instabilité : les redémarrages ou BSOD (offset trop bas), les erreurs de calcul dans le TS Bench (supérieures à 0), et les avertissements WHEA dans l’Observateur d’événements de Windows, qui trahissent des corrections d’erreurs matérielles silencieuses.

```
# Verifier les erreurs materielles WHEA (PowerShell, admin)
Get-WinEvent -LogName System |
  Where-Object { $_.Id -eq 18 -or $_.Id -eq 19 } |
  Select-Object -First 10 TimeCreated, Id, LevelDisplayName
# Aucun evenement WHEA-Logger => undervolt stable
# Des evenements Id 18/19 apparaissent => remontez de 10 mV
```
### Étape 11 – Régler Speed Shift EPP et désactiver le BD PROCHOT

Dans la fenêtre principale, cochez **Speed Shift EPP** et réglez la valeur entre 0 et 255 (0 = performances maximales, 255 = économie maximale ; 80 est un bon compromis en jeu). Ouvrez ensuite **Options** et repérez la case **BD PROCHOT** : si elle est cochée, un composant comme le chargeur ou le GPU peut forcer le CPU à se brider même à basse température. La décocher règle bien des bridages inexpliqués – à condition de surveiller les températures ensuite.

### Étape 12 – Sauvegarder le profil et lancer ThrottleStop au démarrage

Nommez votre profil (par exemple « Gaming ») et cliquez sur **Save** pour l’écrire dans le fichier ThrottleStop.ini. Pour que l’undervolt s’applique à chaque démarrage, ne mettez pas ThrottleStop dans le dossier Démarrage classique : créez plutôt une tâche planifiée qui le lance avec les privilèges administrateur, sinon Windows demandera une confirmation à chaque ouverture. La commande suivante crée cette tâche en une ligne :

```
:: Lancer ThrottleStop au demarrage, sans invite UAC (cmd admin)
schtasks /Create /TN "ThrottleStop" ^
  /TR "C:\ThrottleStop\ThrottleStop.exe" ^
  /SC ONLOGON /RL HIGHEST /F
:: Verifier la tache
schtasks /Query /TN "ThrottleStop"
:: => l'undervolt s'applique des l'ouverture de session
```
## Régler les limites de puissance (TPL) et Speed Shift EPP

L’undervolt réduit la température ; les limites de puissance décident de la performance soutenue. Ouvrez la fenêtre **TPL** (Turbo Power Limits). Vous y trouverez PL1 (Long Power) et PL2 (Short Power), ainsi que le Turbo Time Limit. Sur un portable dont le refroidissement le permet, relever PL1 au niveau de PL2 peut débloquer une performance soutenue nettement supérieure – mais uniquement *après* avoir sécurisé l’undervolt, sinon vous ne ferez que chauffer davantage.

```
=== Fenetre TPL – exemple portable 45 W bien refroidi ===
[x] MMIO Lock desactive (si accessible)
Long Power PL1  (Turbo Boost Long)  : 55 W
Short Power PL2 (Turbo Boost Short) : 65 W
Turbo Time Limit (Tau)              : 56 s
[x] Speed Shift        active
Clamp                               : decoche
=> a valider avec un test thermique de 30 min
```
Attention : augmenter les limites de puissance ne sert à rien si votre système de refroidissement sature déjà. Sur un ultraportable fin, mieux vaut souvent *baisser* PL1 pour un fonctionnement silencieux et frais. La combinaison gagnante sur la majorité des portables gaming reste : undervolt agressif mais stable + PL1 relevé modérément + BD PROCHOT décoché. Testez chaque variable isolément pour comprendre son effet réel sur votre machine.

## BD PROCHOT, C-States et Speed Shift : aller plus loin

Trois réglages avancés méritent votre attention une fois l’undervolt maîtrisé. Le **BD PROCHOT** (Bi-Directional PROCHOT) est un signal qui permet à un autre composant – GPU, VRM, chargeur – d’ordonner au CPU de se brider. Sur certains portables, il se déclenche à des températures GPU pourtant raisonnables et plombe les performances CPU. Le décocher dans Options est souvent le gain le plus spectaculaire, mais surveillez alors la température des autres composants avec HWiNFO.

Les **C-States** gèrent les états de veille des cœurs au repos. Les laisser activés améliore l’autonomie et réduit la température au ralenti, sans nuire aux performances en charge. Enfin, **Speed Shift Technology** (SST) permet au CPU de changer de fréquence en quelques millisecondes plutôt qu’en dizaines de millisecondes : cochez la case Speed Shift et laissez ThrottleStop l’activer au démarrage. Pour la partie graphique et l’affichage des FPS, notre guide MSI Afterburner complète idéalement ce réglage CPU.

## Undervolter l’iGPU Intel et le System Agent (optionnel)

Une fois le CPU Core et le CPU Cache stabilisés, deux rails supplémentaires apparaissent dans la fenêtre FIVR : l’**Intel GPU** (iGPU) et le **System Agent** (SA). L’iGPU alimente le circuit graphique intégré ; l’undervolter apporte un léger gain thermique dans les tâches qui sollicitent le GPU embarqué (lecture vidéo, jeux légers, affichage sur écran externe). Un offset prudent de -30 à -50 mV sur ce rail est généralement sûr, à condition de le tester séparément avec une charge graphique dédiée. Sur un portable gaming équipé d’une carte dédiée, le gain reste marginal puisque l’iGPU est souvent au repos en jeu.

