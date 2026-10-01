---
id: collect-261001-cisco/cisco/mise-a-jour-bios-am5-13-etapes-45-min-2026-4
title: "verification-bios-am5.ps1"
domain: cisco
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: []
keywords: ["amd", "arr", "gpu", "intel"]
source: docs/RAG/collect-261001-cisco/mise-a-jour-bios-am5-13-etapes-45-min-2026.md
source_anchor: ""
source_lines: [186, 270]
sha256: 7e3e2b2678ff88c2ac699401ecbd2e48ee1ed22282a2ba996d17388e54babdbb
---

# verification-bios-am5.ps1

| Symptôme | Cause probable | Solution | 
|---|---|---|
| Écran noir, aucun signal après le flash | RAM ou GPU mal réinsérée, ou reconfiguration nécessaire | Réinitialiser le CMOS via le jumper ou le bouton Clear CMOS, réinsérer RAM et carte graphique | 
| Aucun démarrage, aucun bip, aucune LED | Alimentation coupée pendant un flash précédent | Tenter un nouveau flash via FlashBack si disponible, sinon contacter le support du fabricant | 
| Boucle de redémarrage continue | Profil mémoire incompatible réactivé automatiquement | Réinitialiser le BIOS aux valeurs par défaut, réactiver EXPO/XMP seulement après un démarrage stable | 
| LED FlashBack qui clignote sans jamais s’arrêter | Fichier BIOS mal nommé ou corrompu | Revérifier le nom exact attendu par le fabricant et le hash SHA-256 du fichier | 
| Clé USB non détectée par le port BIOS | Clé formatée en NTFS/exFAT, ou branchée sur le mauvais port | Reformater en FAT32, utiliser exclusivement le port identifié « BIOS » | 
| Message « fichier BIOS invalide » | Fichier destiné à un autre modèle ou une autre révision de carte | Retélécharger le fichier exact pour la référence et la révision identifiées à l’étape 1 | 
| RAM non reconnue à la fréquence annoncée après la mise à jour | Profil EXPO/XMP désactivé par la réinitialisation du flash | Réactiver le profil dans l’onglet mémoire du BIOS | 
| Nouveau CPU toujours non reconnu malgré le flash | Version AGESA insuffisante malgré la mise à jour, ou flash resté incomplet | Vérifier le numéro de version affiché après redémarrage, recommencer le flash si nécessaire | 
| Courbes de ventilateurs remises à un profil bruyant | Réinitialisation complète des réglages personnalisés | Reconfigurer manuellement dans l’onglet dédié (Smart Fan, Fan Xpert…) | 
| Message « CMOS checksum error » au premier redémarrage | Comportement normal après un flash, ou pile CMOS faible | Entrer dans le BIOS, charger les réglages par défaut, remplacer la pile si elle a plus de trois ans | 

Le guide de JustGeek sur la mise à jour du BIOS propose une check-list complémentaire utile si vous préférez repartir de zéro sur un système qui refuse obstinément de redémarrer normalement.

## Script de vérification complet après mise à jour

Pour clore proprement l’opération, voici un script PowerShell qui rassemble en un seul passage tout ce que ce tutoriel vous a appris à vérifier manuellement : version BIOS, modèle de carte mère, processeur reconnu et vitesse réelle de la RAM. Enregistrez-le sous le nom `verification-bios-am5.ps1` et exécutez-le après chaque mise à jour.

```
# verification-bios-am5.ps1
Write-Host "=== Verification post mise a jour BIOS AM5 ===" -ForegroundColor Cyan
$bios = Get-CimInstance Win32_BIOS
Write-Host "`nBIOS actuel :"
Write-Host "  Fabricant   : $($bios.Manufacturer)"
Write-Host "  Version     : $($bios.SMBIOSBIOSVersion)"
Write-Host "  Date        : $($bios.ReleaseDate)"
$board = Get-CimInstance Win32_BaseBoard
Write-Host "`nCarte mere :"
Write-Host "  Modele      : $($board.Product)"
Write-Host "  Revision    : $($board.Version)"
$cpu = Get-CimInstance Win32_Processor
Write-Host "`nProcesseur :"
Write-Host "  Nom         : $($cpu.Name)"
Write-Host "  Coeurs      : $($cpu.NumberOfCores)"
$ram = Get-CimInstance Win32_PhysicalMemory
$vitesseMoyenne = ($ram | Measure-Object -Property Speed -Average).Average
Write-Host "`nMemoire :"
Write-Host "  Modules     : $($ram.Count)"
Write-Host "  Vitesse moy.: $vitesseMoyenne MT/s (comparez au profil EXPO/XMP attendu)"
Write-Host "`n=== Fin de la verification ===" -ForegroundColor Cyan
```
Exemple de sortie sur un système correctement reconfiguré après le flash :

```
=== Verification post mise a jour BIOS AM5 ===
BIOS actuel :
  Fabricant   : American Megatrends International, LLC.
  Version     : 7D78v1C
  Date        : 20260214000000.000000+000
Carte mere :
  Modele      : MAG B650 TOMAHAWK WIFI
  Revision    : 1.0
Processeur :
  Nom         : AMD Ryzen 9 9800X3D
  Coeurs      : 8
Memoire :
  Modules     : 2
  Vitesse moy.: 6000 MT/s (comparez au profil EXPO/XMP attendu)
=== Fin de la verification ===
```
Si la vitesse mémoire affichée reste proche de 4 800 ou 5 200 MT/s alors que votre kit est annoncé à 6 000 MT/s ou plus, c’est le signe que le profil EXPO ou XMP n’a pas été réactivé après le flash, exactement le scénario décrit à l’étape 13. Ce script devient ainsi votre projet de référence, réutilisable à chaque future mise à jour, y compris lors de l’installation d’un prochain processeur.

## Astuces avancées pour les utilisateurs expérimentés

Une fois la procédure de base maîtrisée, quelques réglages supplémentaires méritent d’être connus.

**Le double BIOS.** Certaines cartes haut de gamme embarquent deux puces de firmware distinctes, avec un interrupteur physique pour basculer de l’une à l’autre. En cas de flash raté sur la puce principale, basculer sur la seconde permet souvent de redémarrer normalement puis de reprendre la mise à jour dans de meilleures conditions.

**Le downgrade réfléchi.** Revenir à une version antérieure reste possible sur la plupart des cartes via la même procédure de flash, en sélectionnant simplement un fichier plus ancien. Cette option a du sens si une version récente a introduit une instabilité mémoire chez vous spécifiquement, mais elle redevient risquée si un CPU récent, non supporté par l’ancienne révision AGESA, est déjà installé.

**Les canaux bêta.** Plusieurs fabricants publient des BIOS bêta en avance sur leur page de support ou sur des forums officiels, en général pour corriger un bug précis avant la sortie d’une version stable. Réservez ces versions aux cas où vous rencontrez déjà le bug documenté, jamais en usage préventif sur une machine de production.

**Le flash sans carte graphique dédiée.** Sur un processeur doté de graphismes intégrés, la méthode classique depuis le BIOS fonctionne sans GPU externe. Pour un CPU sans partie graphique, seule la méthode FlashBack, qui ne nécessite aucun affichage, reste utilisable en cas de panne de la carte graphique principale.

**L’automatisation du contrôle.** Le script de vérification présenté plus haut peut être programmé via le Planificateur de tâches Windows pour s’exécuter à chaque démarrage pendant une semaine suivant une mise à jour, une façon simple de détecter une régression de fréquence RAM ou un capteur qui serait passé inaperçu.

## BIOS AM5 vs Intel LGA1851 : ce qui diffère

La philosophie de mise à jour n’est pas identique d’une plateforme à l’autre, et cela influence directement la fréquence à laquelle un utilisateur doit s’en soucier.

AMD a fait de la longévité du socket AM5 un argument commercial explicite, avec plusieurs générations de Ryzen prévues sur la même base physique. Cette approche multiplie mécaniquement les occasions de mise à jour BIOS, puisque chaque nouvelle génération de processeurs doit être intégrée au firmware de cartes mères parfois vendues deux ou trois ans plus tôt. C’est tout l’enjeu détaillé dans notre article sur Zen 6 et le maintien de l’AM5 jusqu’en 2029.

Intel, avec le socket LGA1851, a historiquement suivi un cycle de renouvellement plus court, un nouveau socket accompagnant plus fréquemment chaque nouvelle génération de processeurs. Les utilisateurs Intel se retrouvent donc moins souvent dans la situation « nouveau CPU sur ancienne carte mère », mais changent plus régulièrement de carte mère dans son ensemble plutôt que de se contenter d’un flash BIOS.

