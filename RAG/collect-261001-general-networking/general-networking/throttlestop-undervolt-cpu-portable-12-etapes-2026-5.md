---
id: collect-261001-general-networking/general-networking/throttlestop-undervolt-cpu-portable-12-etapes-2026-5
title: "Verifier les erreurs materielles WHEA (PowerShell, admin)"
domain: general-networking
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["amd", "arr", "benchmark", "gpu", "intel"]
source: docs/RAG/collect-261001-general-networking/throttlestop-undervolt-cpu-portable-12-etapes-2026.md
source_anchor: ""
source_lines: [265, 305]
sha256: 96d948bb565ee41b32ab18f93cea8dff8b79664fcb52c64428750c8bb53d0255
---

# Verifier les erreurs materielles WHEA (PowerShell, admin)

L’undervolting avec ThrottleStop est l’une des rares optimisations qui ne demande aucun compromis : vous baissez la température, vous gagnez en performance soutenue et vous réduisez le bruit, le tout gratuitement et de façon totalement réversible. La méthode se résume à quatre réflexes : partir prudemment à -80 mV sur les rails Core et Cache, descendre par paliers de -10 mV, valider chaque palier par un test de stabilité réel, puis sauvegarder un profil qui se relance au démarrage. Si votre machine verrouille l’undervolt à cause de la correction Plundervolt, les limites de puissance (TPL) et le BD PROCHOT restent des leviers efficaces.

Comptez une trentaine de minutes pour un premier réglage, un peu plus si vous cherchez le dernier millivolt stable. Le jeu en vaut la chandelle : sur un portable gaming qui touchait son Tjmax à 100 °C, il n’est pas rare de repasser sous les 85 °C tout en gagnant plusieurs centaines de mégahertz soutenus. Associez ce réglage CPU à un undervolt GPU et à une surveillance HWiNFO, et vous obtiendrez la machine la plus efficace possible sans ouvrir le châssis. Vos ventilateurs, vos FPS et vos oreilles vous remercieront.

### Couverture associée

## FAQ : ThrottleStop et l’undervolting

### ThrottleStop peut-il endommager mon processeur ?

Non, tant que vous ne touchez qu’à l’Offset Voltage négatif. Undervolter réduit la tension : le pire scénario est un plantage ou un écran bleu, sans dommage matériel. Un simple redémarrage annule tout réglage non sauvegardé. Le seul vrai risque serait d’augmenter la tension au-delà des valeurs d’usine, ce que ce tutoriel ne fait jamais.

### ThrottleStop fonctionne-t-il sur un processeur AMD Ryzen ?

Non. ThrottleStop est conçu exclusivement pour les processeurs Intel. Pour undervolter un CPU AMD Ryzen, utilisez AMD Ryzen Master et son Curve Optimizer, qui offrent une logique d’offset comparable. La partie GPU, en revanche, se règle avec MSI Afterburner quelle que soit la marque du processeur.

### Pourquoi les curseurs de tension sont-ils grisés ?

Parce que l’undervolt est verrouillé au niveau du microcode Intel. Depuis la correction de la faille Plundervolt (CVE-2019-11157), de nombreux BIOS figent l’interface de tension pour empêcher l’attaque. Sur ces machines, aucun logiciel ne peut contourner le verrou ; seule une version de BIOS déverrouillée peut rouvrir l’accès à l’undervolting.

### Quel offset d’undervolt choisir pour commencer ?

Commencez à -80 mV sur les rails CPU Core et CPU Cache simultanément. C’est une valeur que la grande majorité des puces encaissent. Descendez ensuite par paliers de -10 mV en testant la stabilité à chaque étape, puis remontez de 10 à 20 mV par rapport à votre plus bas offset stable pour garder une marge de sécurité.

### Combien de degrés puis-je réellement gagner ?

Cela dépend de votre puce et de votre refroidissement, mais un undervolt de -100 à -125 mV fait généralement chuter la température de 10 à 15 °C sous charge, parfois jusqu’à 20 °C sur les portables qui chauffaient le plus. Le gain le plus visible est souvent l’arrêt du bridage thermique, qui redonne des fréquences soutenues plus élevées.

### Faut-il lancer ThrottleStop à chaque démarrage ?

Oui, l’undervolt n’est actif que lorsque ThrottleStop tourne. Le mieux est de créer une tâche planifiée avec privilèges administrateur (option /RL HIGHEST) qui le lance à l’ouverture de session, sans invite UAC. Le profil sauvegardé s’applique alors automatiquement dès l’arrivée sur le bureau.

### Comment savoir si mon undervolt est stable ?

Trois vérifications : aucun redémarrage ni écran bleu, zéro erreur de calcul dans le TS Bench, et aucun événement WHEA (Id 18/19) dans l’Observateur d’événements de Windows après une charge prolongée. Validez toujours par une vraie session de jeu de 1 à 2 heures, plus révélatrice qu’un simple benchmark.

### ThrottleStop est-il vraiment gratuit ?

Oui, ThrottleStop est entièrement gratuit et distribué sans publicité par TechPowerUp. Il pèse 1,7 Mo dans sa version 9.7.3 du 2 avril 2025, ne nécessite aucune installation et fonctionne sous Windows 7 à 11 ; le paquet TechPowerUp.ThrottleStop est même référencé sur winstall pour une installation via WinGet, avec une fiche mise à jour le 11 juin 2026. C’est l’un des utilitaires d’optimisation matérielle les plus téléchargés au monde – la version espagnole d’Uptodown recensait déjà 208 466 téléchargements pour la seule archive 9.7 (1,71 Mo, SHA256 d83d…800d) au compteur de septembre 2026 –, développé et maintenu bénévolement par Kevin Glynn (UncleWebb).

*Avertissement : l’undervolting reste une opération à vos risques. Procédez par petits paliers, sauvegardez vos données importantes et testez soigneusement la stabilité avant d’utiliser votre portable en production.*
