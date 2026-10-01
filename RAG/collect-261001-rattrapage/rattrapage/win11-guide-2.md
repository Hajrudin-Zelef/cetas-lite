---
id: collect-261001-rattrapage/rattrapage/win11-guide-2
title: "Windows 11 en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: ["2025-10-14"]
keywords: ["copilot", "memory", "packaging"]
source: docs/RAG/collect-261001-rattrapage/win11_guide.md
source_anchor: ""
source_lines: [13, 72]
sha256: e679d1990b9b6ba01a047563e764f322bcf9d705784b0a8a5cd4e20d0aeb9700
---

# Windows 11 en entreprise — Guide technique ultra-complet

1. [Vue d'ensemble et positionnement de Windows 11](#1-vue-densemble-et-positionnement-de-windows-11)
2. [Exigences matérielles : TPM 2.0, Secure Boot, CPU, RAM, stockage](#2-exigences-matérielles--tpm-20-secure-boot-cpu-ram-stockage)
3. [Vérifier la compatibilité : PC Health Check, Get-TPM, scripts d'inventaire](#3-vérifier-la-compatibilité--pc-health-check-get-tpm-scripts-dinventaire)
4. [Éditions Windows 11 : Pro vs Entreprise, canaux et cycle de vie](#4-éditions-windows-11--pro-vs-entreprise-canaux-et-cycle-de-vie)
5. [Obtenir les médias : ISO, MCT, VLSC, canaux de mise à jour](#5-obtenir-les-médias--iso-mct-vlsc-canaux-de-mise-à-jour)
6. [Installation propre : clé USB bootable, étapes, pilotes](#6-installation-propre--clé-usb-bootable-étapes-pilotes)
7. [Mise à niveau depuis Windows 10 : chemins, outils, pièges](#7-mise-à-niveau-depuis-windows-10--chemins-outils-pièges)
8. [Installation automatisée : unattend.xml, clés de produit, setupconfig](#8-installation-automatisée--unattendxml-clés-de-produit-setupconfig)
9. [Sysprep : généralisation, audit mode, capture d'image](#9-sysprep--généralisation-audit-mode-capture-dimage)
10. [MDT : installation, deployment share, task sequences](#10-mdt--installation-deployment-share-task-sequences)
11. [MDT : Lite Touch, pilotes, règles CustomSettings.ini](#11-mdt--lite-touch-pilotes-règles-customsettingsini)
12. [Autopilot + Intune : principe, profils, inscription](#12-autopilot--intune--principe-profils-inscription)
13. [Autopilot Reset, Fresh Start et scénarios de réaffectation](#13-autopilot-reset-fresh-start-et-scénarios-de-réaffectation)
14. [Jointure au domaine vs Entra join vs hybride](#14-jointure-au-domaine-vs-entra-join-vs-hybride)
15. [BitLocker : fondamentaux, TPM seul vs TPM+PIN, XTS-AES](#15-bitlocker--fondamentaux-tpm-seul-vs-tpmpin-xts-aes)
16. [BitLocker : clés de récupération (AD, fichier, impression, Entra)](#16-bitlocker--clés-de-récupération-ad-fichier-impression-entra)
17. [BitLocker : stratégies GPO, MBAM / BitLocker Management, déploiement silencieux](#17-bitlocker--stratégies-gpo-mbam--bitlocker-management-déploiement-silencieux)
18. [BitLocker : dépannage — écran de récupération, boucles, causes fréquentes](#18-bitlocker--dépannage--écran-de-récupération-boucles-causes-fréquentes)
19. [Profils utilisateurs : locaux, itinérants, obligatoires](#19-profils-utilisateurs--locaux-itinérants-obligatoires)
20. [FSLogix : principe, conteneurs de profils, cas d'usage](#20-fslogix--principe-conteneurs-de-profils-cas-dusage)
21. [Migration de profils : USMT (ScanState / LoadState)](#21-migration-de-profils--usmt-scanstate--loadstate)
22. [GPO Windows 11 : menu Démarrer, barre des tâches, disposition](#22-gpo-windows-11--menu-démarrer-barre-des-tâches-disposition)
23. [GPO Windows 11 : désactiver Copilot, Widgets, recommandations](#23-gpo-windows-11--désactiver-copilot-widgets-recommandations)
24. [GPO Windows 11 : bonnes pratiques, héritage, filtrage, dépannage](#24-gpo-windows-11--bonnes-pratiques-héritage-filtrage-dépannage)
25. [Windows Update for Business : anneaux de déploiement](#25-windows-update-for-business--anneaux-de-déploiement)
26. [WUfB : stratégies (CSP, GPO, Intune), délais, pause, qualité vs fonctionnalités](#26-wufb--stratégies-csp-gpo-intune-délais-pause-qualité-vs-fonctionnalités)
27. [WSUS vs Windows Update for Business : comparatif et coexistence](#27-wsus-vs-windows-update-for-business--comparatif-et-coexistence)
28. [Déploiement d'applications : winget en entreprise](#28-déploiement-dapplications--winget-en-entreprise)
29. [Microsoft Store en entreprise, applications privées, packaging](#29-microsoft-store-en-entreprise-applications-privées-packaging)
30. [Sécurité : Microsoft Defender Antivirus — configuration entreprise](#30-sécurité--microsoft-defender-antivirus--configuration-entreprise)
31. [Sécurité : SmartScreen, protection réseau, contrôle des applications](#31-sécurité--smartscreen-protection-réseau-contrôle-des-applications)
32. [Sécurité : Credential Guard, LSA Protection, HVCI / Memory Integrity](#32-sécurité--credential-guard-lsa-protection-hvci--memory-integrity)
33. [Sécurité : durcissement — baselines Microsoft, ASR, Attack Surface Reduction](#33-sécurité--durcissement--baselines-microsoft-asr-attack-surface-reduction)
34. [LAPS Windows (nouveau) : principe, GPO, récupération](#34-laps-windows-nouveau--principe-gpo-récupération)
35. [Pare-feu Windows Defender : profils, règles, GPO, journalisation](#35-pare-feu-windows-defender--profils-règles-gpo-journalisation)
36. [Réseau : Wi-Fi d'entreprise 802.1X (WPA2/WPA3-Enterprise)](#36-réseau--wi-fi-dentreprise-8021x-wpa2wpa3-enterprise)
37. [Réseau : VPN — Always On VPN (intro), profils, dépannage](#37-réseau--vpn--always-on-vpn-intro-profils-dépannage)
38. [Réseau : dépannage — pile TCP/IP, DNS, proxy, certificats](#38-réseau--dépannage--pile-tcpip-dns-proxy-certificats)
39. [WinRE : environnement de récupération, outils, personnalisation](#39-winre--environnement-de-récupération-outils-personnalisation)
40. [Points de restauration, sauvegardes système, récupération](#40-points-de-restauration-sauvegardes-système-récupération)
41. [Dépannage : méthodologie et boîte à outils de l'admin](#41-dépannage--méthodologie-et-boîte-à-outils-de-ladmin)
42. [Cas pratique 1 : le PC ne démarre plus (écran noir, BOOTMGR, BCD)](#42-cas-pratique-1--le-pc-ne-démarre-plus-écran-noir-bootmgr-bcd)
43. [Cas pratique 2 : boucle de réparation automatique / WinRE en boucle](#43-cas-pratique-2--boucle-de-réparation-automatique--winre-en-boucle)
44. [Cas pratique 3 : BitLocker demande la clé de récupération en boucle](#44-cas-pratique-3--bitlocker-demande-la-clé-de-récupération-en-boucle)
45. [Cas pratique 4 : mise à jour Windows bloquée / échec répété](#45-cas-pratique-4--mise-à-jour-windows-bloquée--échec-répété)
46. [Cas pratique 5 : mise à niveau 23H2 → 24H2 qui échoue](#46-cas-pratique-5--mise-à-niveau-23h2--24h2-qui-échoue)
47. [Cas pratique 6 : pilote défectueux — écran bleu, rollback, signature](#47-cas-pratique-6--pilote-défectueux--écran-bleu-rollback-signature)
48. [Cas pratique 7 : performances dégradées — disque, mémoire, processus](#48-cas-pratique-7--performances-dégradées--disque-mémoire-processus)
49. [Cas pratique 8 : profil utilisateur corrompu / temporaire](#49-cas-pratique-8--profil-utilisateur-corrompu--temporaire)
50. [Cas pratique 9 : impossible de joindre le domaine / relation d'approbation](#50-cas-pratique-9--impossible-de-joindre-le-domaine--relation-dapprobation)
51. [Cas pratique 10 : GPO qui ne s'applique pas](#51-cas-pratique-10--gpo-qui-ne-sapplique-pas)
52. [Cas pratique 11 : Wi-Fi 802.1X — échec d'authentification](#52-cas-pratique-11--wi-fi-8021x--échec-dauthentification)
53. [Cas pratique 12 : VPN qui ne se connecte pas](#53-cas-pratique-12--vpn-qui-ne-se-connecte-pas)
54. [Cas pratique 13 : applications qui ne se lancent plus / Store corrompu](#54-cas-pratique-13--applications-qui-ne-se-lancent-plus--store-corrompu)
55. [Cas pratique 14 : espace disque saturé après mise à niveau](#55-cas-pratique-14--espace-disque-saturé-après-mise-à-niveau)
56. [Cas pratique 15 : Windows Hello / biométrie ne fonctionne plus](#56-cas-pratique-15--windows-hello--biométrie-ne-fonctionne-plus)
57. [Cas pratique 16 : imprimante réseau / partage SMB inaccessible](#57-cas-pratique-16--imprimante-réseau--partage-smb-inaccessible)
58. [Fin de vie de Windows 10 (14/10/2025) : stratégie de migration](#58-fin-de-vie-de-windows-10-14102025--stratégie-de-migration)
59. [Inventaire du parc : PowerShell, GLPI, Intune, rapports](#59-inventaire-du-parc--powershell-glpi-intune-rapports)
60. [Supervision et journaux : Event Viewer, journaux clés, collecte](#60-supervision-et-journaux--event-viewer-journaux-clés-collecte)
