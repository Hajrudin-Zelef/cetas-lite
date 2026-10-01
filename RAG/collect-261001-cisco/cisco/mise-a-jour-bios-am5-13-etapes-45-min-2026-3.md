---
id: collect-261001-cisco/cisco/mise-a-jour-bios-am5-13-etapes-45-min-2026-3
title: "verification-bios-am5.ps1"
domain: cisco
role: reference
task: reference
actors: ["AMD"]
dates: []
keywords: ["amd"]
source: docs/RAG/collect-261001-cisco/mise-a-jour-bios-am5-13-etapes-45-min-2026.md
source_anchor: ""
source_lines: [123, 185]
sha256: 8ecc81715eeaaade206f8eb1d5b33a3a76e918dcd45d9e672144eb362e4adc04
---

# verification-bios-am5.ps1

**Étape 10**, entrez dans le BIOS en appuyant sur Suppr ou F2 au démarrage, puis basculez en mode avancé si votre carte affiche par défaut une interface simplifiée, ce qui est le cas de l’EZ Mode chez ASUS et Gigabyte.

**Étape 11**, lancez l’utilitaire de flash intégré : M-Flash chez MSI, EZ Flash chez ASUS, Q-Flash chez Gigabyte, Instant Flash chez ASRock. Chacun se trouve dans un onglet Outils ou Tool du menu principal. Sélectionnez votre clé USB, puis le fichier BIOS copié à l’étape 6.

**Étape 12**, validez et patientez. L’écran reste noir ou affiche une barre de progression selon les fabricants, pendant 3 à 8 minutes. Le PC redémarre automatiquement une ou deux fois à la fin de l’opération, ce comportement est normal et ne doit surtout pas être interrompu.

**Étape 13**, reconfigurez votre BIOS. C’est l’étape la plus souvent oubliée : une mise à jour réinitialise la quasi-totalité des réglages à leurs valeurs par défaut. Il faut donc systématiquement :

- réactiver le profil mémoire EXPO ou XMP, sans quoi la RAM retombe à sa fréquence JEDEC de base, généralement autour de 4 800 à 5 200 MT/s au lieu des 6 000 MT/s ou plus annoncés sur la boîte
- vérifier l’ordre de démarrage (Boot Order), en particulier si vous utilisez plusieurs disques
- réactiver le Resizable BAR si vous l’aviez configuré pour votre carte graphique
- reconfigurer vos courbes de ventilateurs personnalisées, remises à un profil silencieux par défaut

Un redémarrage complet suivi d’une vérification dans l’utilitaire HWiNFO, que nous détaillons dans notre tutoriel dédié à la surveillance PC avec HWiNFO, permet de confirmer que les fréquences RAM et les températures sont revenues à la normale après la mise à jour.

## Comprendre les versions AGESA : quel firmware pour quel Ryzen

AGESA, pour AMD Generic Encapsulated Software Architecture, est le socle logiciel qu’AMD fournit à tous les fabricants de cartes mères, à charge pour chacun de l’intégrer dans son propre BIOS. C’est ce numéro de version, bien plus que le numéro de BIOS propre au fabricant, qui détermine réellement la compatibilité CPU et RAM d’une carte mère AM5.

| Révision AGESA | Apport principal | CPU ou usage concerné | 
|---|---|---|
| ComboAM5 PI 1.1.7.0 | Version minimale pour démarrer (POST) | Ryzen 9000 non-X3D | 
| ComboAM5 PI 1.1.7.0 Patch A | Compatibilité de démarrage élargie sur X670, B650 et A620 | Ryzen 9000 (selon Gigabyte) | 
| ComboAM5 PI 1.2.0.2 | Amélioration rapportée de la compatibilité mémoire (non confirmée par une source officielle) | Kits DDR5 haute fréquence | 
| ComboAM5 PI 1.2.0.3 | Correctifs de sécurité : dépassement mémoire PeCoffLoader et vérification de signature microcode | Tous CPU AM5 | 
| ComboAM5 PI 1.3.0.1 (2025) | Ajout de la prise en charge | Ryzen 9000 séries X3D | 

Le point important à retenir : deux cartes mères de marques différentes portant des numéros de BIOS totalement distincts peuvent embarquer la même révision AGESA, et donc offrir exactement le même niveau de compatibilité. Inversement, deux versions BIOS consécutives chez un même fabricant peuvent conserver la même révision AGESA si la mise à jour ne portait que sur un réglage mineur sans rapport avec la compatibilité CPU ou RAM.

Ces informations évoluent vite, et certains apports listés ci-dessus proviennent de guides techniques tiers plutôt que de bulletins officiels AMD. La meilleure pratique reste donc de toujours vérifier la liste de compatibilité officielle de votre référence exacte via le site chipset AM5 d’AMD, plutôt que de se fier uniquement à un tableau générique, aussi à jour soit-il.

## Tableau comparatif des méthodes de flash par marque

Avant de choisir votre méthode, ce tableau récapitule les différences pratiques entre le flash sans CPU et le flash classique depuis le BIOS, marque par marque.

| Marque / méthode | CPU requis ? | Emplacement | Format du fichier | Durée approximative | 
|---|---|---|---|---|
| ASUS BIOS FlashBack | Non | Port USB dédié « BIOS » + bouton arrière | Nom précis selon le modèle, à la racine de la clé | Quelques minutes, LED clignotante | 
| MSI Flash BIOS Button | Non | Port USB dédié + bouton Flash BIOS | Fichier à la racine, FAT32 | 5 à 10 minutes | 
| Gigabyte Q-Flash Plus | Non | Port USB dédié + bouton Q-Flash | Renommé « GIGABYTE.bin » | Quelques minutes, LED clignotante | 
| ASRock BIOS Flashback | Non | Port USB dédié + bouton (variable selon modèle) | Voir manuel du modèle exact | Variable selon modèle | 
| Méthode classique (M-Flash / EZ Flash / Q-Flash / Instant Flash) | Oui, CPU fonctionnel requis | Depuis le menu BIOS/UEFI | Fichier à la racine, FAT32 | 3 à 8 minutes | 

Pour un premier montage ou un changement de processeur qui empêche déjà le démarrage, la méthode sans CPU reste la seule option viable. Pour une mise à jour de routine sur un système qui fonctionne déjà, la méthode classique depuis le BIOS est généralement plus rapide, car elle évite d’avoir à démonter le radiateur du processeur ou à débrancher les câbles d’alimentation.

## Les 7 erreurs les plus fréquentes lors d’une mise à jour BIOS

La plupart des cartes mères rendues inutilisables après un flash ne le sont pas à cause d’un défaut matériel, mais d’une de ces sept erreurs, toutes évitables.

1. **Couper l’alimentation pendant le flash.** Même une coupure de quelques secondes peut laisser le firmware dans un état incomplet et rendre la carte muette au démarrage suivant.
2. **Flasher le mauvais fichier.** Confondre une révision de carte, une variante WiFi et non-WiFi, ou un modèle proche dans la même gamme reste la deuxième cause la plus fréquente de blocage.
3. **Utiliser une clé USB mal formatée.** Une clé en NTFS ou exFAT n’est tout simplement pas lue par la plupart des ports BIOS dédiés, sans message d’erreur explicite.
4. **Télécharger depuis une source non officielle.** Un fichier modifié ou corrompu récupéré sur un forum peut sembler fonctionner au moment de la copie, puis échouer en plein flash.
5. **Ignorer une pile CMOS fatiguée.** Sur une carte mère de plus de deux ou trois ans, une pile faible augmente sensiblement le risque de coupure inopinée pendant l’opération.
6. **Oublier de reconfigurer après le flash.** Repartir sur une RAM au ralenti pendant des semaines parce que le profil EXPO n’a pas été réactivé reste étonnamment courant.
7. **Flasher une version plus ancienne alors qu’un CPU récent est déjà installé.** Un downgrade peut faire disparaître la prise en charge du processeur en place, provoquant un blocage qui ne peut alors être résolu que via la méthode sans CPU si la carte la propose.

Le guide de Grosbill sur la mise à jour du BIOS sans plantage recense des précautions similaires côté PC portable, où la marge d’erreur est encore plus faible en l’absence de fonction FlashBack.

## Dépannage : 10 problèmes courants et leurs solutions

Si malgré ces précautions quelque chose se passe mal, voici les symptômes les plus signalés et la marche à suivre pour chacun.

