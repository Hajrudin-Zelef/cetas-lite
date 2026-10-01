---
id: collect-261001-general-networking/general-networking/memtest86-tester-sa-ram-en-13-etapes-2026-3
title: "memtest86-tester-sa-ram-en-13-etapes-2026"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "mai", "memory", "open source"]
source: docs/RAG/collect-261001-general-networking/memtest86-tester-sa-ram-en-13-etapes-2026.md
source_anchor: ""
source_lines: [103, 181]
sha256: d53b22fb8ef260e63f237e301c6d32f93c592b52e3cc5810a6b305b696c1b4ee
---

# memtest86-tester-sa-ram-en-13-etapes-2026

**Étape 9 : choisir le nombre de passes et les tests à exécuter.** Par défaut, MemTest86 enchaîne l’ensemble de ses algorithmes de test sur un nombre de passes défini. Pour un premier diagnostic, laissez la configuration par défaut et lancez le test avec la touche F10. Pour une validation plus poussée après avoir trouvé des erreurs, le menu de configuration permet de désélectionner certains tests et de vous concentrer sur ceux qui ont déclenché une alerte, afin de confirmer le problème plus rapidement lors des passes suivantes.

Voici, à titre d’illustration, la structure typique de cet écran de configuration :

```
MemTest86 v11.7 - Configuration
------------------------------------------------
Langue du clavier      : Français
RAM détectée            : 32768 Mo (32 Go)
Tests sélectionnés      : Tous (par défaut)
Nombre de passes        : Continu jusqu'à interruption
CPU / Threads utilisés  : Tous les cœurs disponibles
[F1] Aide   [F2] Config avancée   [F10] Démarrer le test
```
## Étape 10 : lire et interpréter les résultats du test

Pendant le test, l’écran principal affiche en permanence le pourcentage de progression de la passe en cours, le nombre total de passes effectuées, le temps écoulé et, surtout, un compteur d’erreurs. Tant que ce compteur reste à zéro, la RAM se comporte correctement sur les motifs testés jusque-là. Voici un exemple simplifié de ce à quoi ressemble un rapport propre après plusieurs passes :

```
Passe : 4 / Continu       Temps écoulé : 03:12:47
Test en cours : #8 [Random Number Sequence]
Progression : 100 %
Erreurs détectées : 0
Statut : AUCUNE ERREUR - RAM stable sur 4 passes complètes
```
Quand une erreur survient, MemTest86 l’affiche en temps réel dans la partie basse de l’écran, avec le numéro du test concerné, l’adresse mémoire fautive et l’écart entre la valeur attendue et la valeur lue. Un exemple simplifié de ce type de rapport :

```
Passe : 1 / Continu       Temps écoulé : 00:47:12
Test en cours : #5 [Block Move]
Erreurs détectées : 3
------------------------------------------------
Erreur 1 : Adresse 0x3A2F1000 - Attendu FFFFFFFF, Lu FFFFFFFB
Erreur 2 : Adresse 0x3A2F1004 - Attendu FFFFFFFF, Lu FFFFFFFB
Erreur 3 : Adresse 0x7C0E8A20 - Attendu 00000000, Lu 00000004
```
Une erreur isolée sur une seule adresse, trouvée une seule fois sur plusieurs passes, mérite d’être surveillée mais ne condamne pas forcément le module. Des erreurs répétées sur la même plage d’adresses, ou qui se multiplient au fil des passes, indiquent en revanche un problème matériel réel qui justifie de passer aux étapes suivantes.

## Étapes 11 à 13 : isoler, corriger et revérifier une RAM défectueuse

**Étape 11 : isoler le module fautif par dichotomie.** Sur une configuration à plusieurs barrettes, retirez tous les modules sauf un et relancez le test dans le même slot. Répétez l’opération barrette par barrette. Si une seule provoque des erreurs quel que soit le slot utilisé, le module est probablement en cause. Si les erreurs apparaissent uniquement dans un slot précis, quel que soit le module testé, le problème vient plutôt du socket ou de la carte mère elle-même, pas de la RAM.

**Étape 12 : appliquer une correction.** Plusieurs pistes, par ordre de simplicité : réenclencher physiquement les barrettes après avoir nettoyé les contacts dorés avec une gomme douce, désactiver le profil XMP/EXPO pour revenir à la fréquence JEDEC certifiée et voir si les erreurs disparaissent, ou mettre à jour le BIOS si le constructeur a publié un correctif de microcode mémoire (voir notre tutoriel de mise à jour BIOS sur plateforme AM5 pour la procédure complète). Si aucune de ces pistes ne résout le problème, la barrette est probablement défectueuse et il faut engager une procédure de retour auprès du fabricant, comme le service support de Crucial pour cette marque.

**Étape 13 : revérifier après chaque changement.** Après toute intervention, quelle qu’elle soit, relancez MemTest86 pour au moins deux passes complètes avant de conclure. Une seule passe propre après un changement de configuration ne suffit pas à garantir la stabilité : c’est justement le genre de faux positif que ce tutoriel cherche à éviter.

## Identifier votre configuration RAM en ligne de commande

Avant de démonter quoi que ce soit, il est utile de savoir précisément quels modules sont installés, dans quels slots, et à quelle vitesse ils tournent réellement. Sous Windows, PowerShell donne cette information sans ouvrir le boîtier :

`Get-CimInstance Win32_PhysicalMemory | Select-Object BankLabel, Manufacturer, PartNumber, Capacity, Speed, ConfiguredClockSpeed`
Cette commande liste chaque barrette avec son emplacement, son fabricant, sa référence exacte, sa capacité et sa vitesse configurée. La référence (PartNumber) est justement ce que le support technique du fabricant demandera en cas de RMA. Sur un système plus ancien ou en invite de commande classique, l’équivalent existe aussi :

`wmic memorychip get BankLabel, Manufacturer, PartNumber, Capacity, Speed`
Notez les résultats avant de commencer les tests d’isolement de l’étape 11 : cela évite de confondre les slots une fois les barrettes retirées et remises dans un ordre différent.

## Retrouver les erreurs mémoire déjà enregistrées par Windows

Avant même de créer une clé USB, Windows a peut-être déjà enregistré des indices utiles. Chaque plantage lié à la mémoire laisse une trace dans l’Observateur d’événements, sous forme de code d’arrêt (bug check) associé au fichier minidump généré lors du crash. Ouvrez l’Observateur d’événements, puis Journaux Windows > Système, et filtrez sur la source « BugCheck » pour retrouver l’historique de ces incidents.

Les codes suivants pointent le plus souvent vers un problème mémoire plutôt que vers un pilote ou une application isolée :

- **0x0000001A (MEMORY_MANAGEMENT).** Le gestionnaire de mémoire Windows a détecté une incohérence grave dans la gestion des pages mémoire.
- **0x0000000A (IRQL_NOT_LESS_OR_EQUAL).** Un pilote a tenté d’accéder à une zone mémoire à un niveau de priorité incorrect, parfois révélateur d’une adresse physique corrompue.
- **0x00000050 (PAGE_FAULT_IN_NONPAGED_AREA).** Le système a tenté de lire une donnée absente d’une zone mémoire qui ne devrait jamais être vidée sur le disque.
- **0x0000007F (UNEXPECTED_KERNEL_MODE_TRAP).** Erreur matérielle de bas niveau, parfois liée au CPU mais fréquemment corrélée à une instabilité mémoire sous charge.

Retrouver plusieurs occurrences de ces codes sur une période courte constitue un signal fort pour lancer MemTest86 sans attendre. À l’inverse, l’absence totale de ces codes, combinée à des plantages qui touchent toujours le même logiciel, oriente plutôt vers une cause logicielle et rend un test mémoire moins prioritaire.

## MemTest86 vs MemTest86+ vs Windows Memory Diagnostic : quel outil choisir

MemTest86 n’est pas le seul outil disponible pour tester la mémoire vive. Deux alternatives reviennent régulièrement dans les recommandations, chacune avec un usage différent.

### MemTest86 (PassMark) : la référence UEFI moderne

C’est l’outil couvert dans ce tutoriel. Développement actif — la version 11.7 (Build 1000), publiée en mai 2026, reste à ce jour la dernière release UEFI-only répertoriée sur le site de PassMark —, interface UEFI native, support multi-thread pour accélérer les passes sur les configurations à nombreux cœurs, et éditions payantes pour l’automatisation en entreprise. C’est le choix par défaut pour un diagnostic fiable sur du matériel récent.

### MemTest86+ : l’alternative open source

