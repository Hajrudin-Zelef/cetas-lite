---
id: collect-261001-general-networking/general-networking/throttlestop-undervolt-cpu-portable-12-etapes-2026-4
title: "Verifier les erreurs materielles WHEA (PowerShell, admin)"
domain: general-networking
role: reference
task: reference
actors: ["Intel"]
dates: []
keywords: ["agent", "benchmark", "gpu", "intel"]
source: docs/RAG/collect-261001-general-networking/throttlestop-undervolt-cpu-portable-12-etapes-2026.md
source_anchor: ""
source_lines: [204, 264]
sha256: c0ebee463b67b53ba11f8ab55fde71b7ba57ddbfb0b2813c97a22aba34a9851a
---

# Verifier les erreurs materielles WHEA (PowerShell, admin)

Le **System Agent** alimente le contrôleur mémoire et les bus internes. C’est le rail le plus délicat : un undervolt trop appuyé provoque des erreurs mémoire sournoises, parfois sans plantage immédiat, qui corrompent silencieusement des données. La recommandation des experts de l’undervolting est claire : laissez le System Agent à 0 mV, sauf besoin très précis et test mémoire approfondi (MemTest86). Le rapport bénéfice/risque penche nettement en défaveur de ce rail pour l’usage courant. Concentrez vos efforts sur le duo Core + Cache, qui apporte plus de 90 % du gain thermique total.

## Undervolt sur les portables gaming vendus en France : ROG, Legion, MSI

La compatibilité de l’undervolt dépend surtout de la marque et de la génération du portable. Chez **ASUS ROG**, très populaire en France, beaucoup de modèles Strix et Zephyrus laissent le FIVR accessible sur les générations plus anciennes, mais certains BIOS récents l’ont verrouillé ; la communauté du forum des portables ASUS francophone documente les cas au modèle près. Chez **Lenovo Legion**, l’undervolt fonctionne fréquemment, en complément de Legion Vantage pour les profils de ventilation. **MSI** (via MSI Center) et **HP Omen** présentent des situations variables selon la révision de BIOS.

Les **Razer Blade** et une partie des **Dell/Alienware** récents verrouillent souvent l’interface de tension par défaut, héritage direct de la correction Plundervolt. Avant d’acheter ou de vous lancer, la démarche est simple : identifiez la génération exacte de votre processeur Intel, mettez à jour ou notez votre version de BIOS, puis testez si les curseurs FIVR répondent. S’ils sont verrouillés, rabattez-vous sur les limites de puissance (TPL) et le BD PROCHOT, souvent encore accessibles, ainsi que sur la gestion de la ventilation propre au constructeur. Sur les machines déverrouillées, l’undervolt reste l’optimisation gratuite au meilleur rapport gain/effort pour un portable gaming.

## Résultats concrets : avant/après sur un portable gaming

Voici un exemple représentatif des gains obtenus sur un portable gaming équipé d’un Core i7 de 13e génération (TDP 45 W), avec un offset stable de -125 mV sur les rails Core et Cache. Ces valeurs sont fournies à titre d’illustration : vos résultats dépendront de votre puce (loterie du silicium), de votre refroidissement et de la charge testée.

| Mesure (charge CPU soutenue) | Avant undervolt | Après -125 mV | Gain | 
|---|---|---|---|
| Température package max | 100 °C | 85 °C | -15 °C | 
| Fréquence tous cœurs soutenue | 3,12 GHz | 3,60 GHz | +0,48 GHz | 
| Tension CPU sous charge | 1,248 V | 1,123 V | -0,125 V | 
| Drapeau THERMAL (bridage) | Actif | Inactif | Bridage supprimé | 
| Bruit ventilateurs | Maximum | Modéré | Plus silencieux | 
| FPS moyens (scène CPU-limitée) | Référence | +8 à +12 % | Gain net | 

Le schéma est presque toujours le même : la température chute de 10 à 15 °C, le CPU cesse d’atteindre son Tjmax, il maintient une fréquence plus élevée et plus régulière, et les FPS progressent surtout dans les jeux ou scènes limités par le processeur. Les ventilateurs, sollicités moins fort, deviennent plus discrets. C’est ce triple gain – température, performance, silence – qui explique la popularité durable de ThrottleStop sur les forums de portables gaming comme le fil de discussion officiel de TechPowerUp.

## 8 pièges courants de l’undervolting à éviter

La plupart des échecs d’undervolt ne viennent pas du logiciel mais de la méthode. Voici les huit erreurs qui reviennent le plus souvent, et comment les contourner.

1. **Partir trop fort d’emblée.** Commencer à -150 mV masque le vrai point de rupture. Débutez à -80 mV et descendez par -10 mV.
2. **Déséquilibrer Core et Cache.** Des offsets différents sur ces deux rails liés provoquent des plantages difficiles à diagnostiquer. Gardez-les identiques.
3. **Oublier « Save voltages immediately ».** Sans cette option, l’undervolt disparaît après une veille et le CPU repart à pleine tension.
4. **Tester uniquement au TS Bench.** Un offset stable au benchmark peut planter en jeu. Validez toujours par une vraie session de jeu prolongée.
5. **Ignorer les erreurs WHEA.** Un système qui ne plante pas mais accumule des corrections WHEA est déjà instable : remontez la tension.
6. **Ne pas lancer en administrateur.** Sans privilèges élevés, certains réglages ne s’appliquent pas et le profil ne se recharge pas au démarrage.
7. **Confondre undervolt verrouillé et undervolt inefficace.** Si la température ne bouge pas, vérifiez d’abord le verrouillage Plundervolt avant de creuser l’offset.
8. **Négliger le contexte thermique.** Un undervolt ne remplace pas une pâte thermique fatiguée ou des grilles d’aération obstruées. Nettoyez d’abord la machine.

## Dépannage : 8 problèmes fréquents avec ThrottleStop

Quand ThrottleStop ne se comporte pas comme prévu, la cause figure presque toujours dans ce tableau. Il couvre les huit situations les plus signalées par les utilisateurs francophones.

| Symptôme | Cause probable | Solution | 
|---|---|---|
| Curseurs FIVR grisés | Undervolt verrouillé (Plundervolt / BIOS) | Chercher un BIOS déverrouillé ; sinon miser sur TPL | 
| La température ne baisse pas | Offset non appliqué ou verrouillé | Vérifier « IA Offset » dans HWiNFO | 
| Écran bleu au démarrage | Offset trop agressif | Redémarrer (le profil non sauvé s’annule) et remonter de 20 mV | 
| Undervolt perdu après veille | « Save voltages immediately » non coché | Réactiver l’option dans FIVR | 
| Bridage à basse température | BD PROCHOT déclenché par un composant | Décocher BD PROCHOT dans Options | 
| Profil non appliqué au boot | Lancement sans droits administrateur | Créer une tâche planifiée /RL HIGHEST | 
| Erreurs WHEA dans l’Observateur | Instabilité silencieuse | Remonter la tension de 10 à 20 mV | 
| PL1/PL2 impossibles à modifier | Verrous MMIO / BIOS constructeur | Vérifier les options BIOS ; certains portables bloquent le TPL | 

Si votre portable reste chaud malgré un undervolt stable, le problème est peut-être ailleurs : pilotes graphiques encrassés, résidus de tension GPU, ou processus en arrière-plan. Un nettoyage propre des pilotes avec notre tutoriel DDU élimine une variable et facilite le diagnostic. Pensez aussi à vérifier que le mode d’alimentation de Windows est réglé sur « Performances optimales » pendant vos tests, puis sur un mode équilibré au quotidien.

## Astuces avancées : profils par jeu, autonomie et Game Mode

ThrottleStop gère quatre profils indépendants, et c’est là que réside sa vraie puissance. Créez un profil « Gaming » agressif (undervolt profond, PL1 relevé, Speed Shift EPP à 0) et un profil « Batterie » économe (undervolt conservateur, PL1 abaissé, EPP à 128 ou plus). ThrottleStop peut basculer automatiquement de l’un à l’autre selon que le portable est branché ou sur batterie, grâce à l’onglet Options et à la case « Battery » / « AC ». Vous obtenez ainsi un portable qui tient plus longtemps en déplacement et qui donne tout branché sur secteur.

Pour les joueurs, la combinaison la plus efficace consiste à associer un undervolt CPU via ThrottleStop à un undervolt GPU. Sur les consoles portables et PC compacts, cette approche double le gain thermique : c’est d’ailleurs l’un des atouts des puces récentes comme le Ryzen Z2 Extreme, pensées pour l’efficacité énergétique. Enfin, gardez un œil sur la tarification des consoles portables gaming : sur ces appareils au refroidissement contraint, un undervolt bien réglé fait souvent la différence entre 45 et 60 FPS.

## Conclusion : un portable plus frais, plus rapide et plus silencieux

