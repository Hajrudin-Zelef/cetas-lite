---
id: collect-261001-rattrapage/rattrapage/windows-server-guide-1
title: "Windows Server en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["datacenter", "distribution"]
source: docs/RAG/collect-261001-rattrapage/windows_server_guide.md
source_anchor: ""
source_lines: [1, 165]
sha256: 340ff06762f40de0fb320b9b23788b6472de84b9f59b6cfe4cf015efaf7f6b94
---

# Windows Server en entreprise — Guide technique ultra-complet

> **Public :** Zelef, chef de service systèmes & énergies — administrateurs, techniciens senior.
> **Versions couvertes :** Windows Server 2019, 2022, 2025 (LTSC). Versions 2025 signalées quand le comportement diffère.
> **Ton :** direct, dense, pratique. Commandes PowerShell testables, tableaux de référence, checklists.
> **Avertissement :** les mots de passe et secrets dans ce guide sont des exemples fictifs (`P@ssw0rd-Exemple-2026!`). Remplace-les systématiquement.

---

## Table des matières

1. Éditions et canaux de distribution
2. Standard vs Datacenter : le tableau qui décide
3. Checklist de choix d'édition (10 critères)
4. Installation pas à pas (Desktop Experience)
5. Installation Server Core pas à pas
6. Post-installation : la checklist des 20 minutes
7. Server Core vs Desktop Experience : différences concrètes
8. sconfig : le couteau suisse de Server Core
9. Administration à distance : RSAT, Windows Admin Center, WinRM
10. Quand choisir Core, quand choisir Desktop
11. Rôles et fonctionnalités : concepts
12. Installation via Server Manager
13. Installation via PowerShell (Install-WindowsFeature)
14. Les 15 rôles courants et quand les installer
15. Hyper-V : architecture et prérequis
16. Commutateurs virtuels : externe, interne, privé
17. VM génération 1 vs génération 2
18. Création d'une VM pas à pas (PowerShell)
19. Mémoire dynamique, processeurs virtuels, disques
20. Checkpoints : standard vs production
21. Réplication Hyper-V (Hyper-V Replica)
22. Cluster de basculement + CSV
23. Bonnes pratiques Hyper-V (20 règles d'or)
24. WSUS : architecture et prérequis
25. Installation et configuration initiale de WSUS
26. Groupes d'ordinateurs et ciblage côté client
27. Approbations : manuelles et règles automatiques
28. Maintenance : Server Cleanup Wizard et scripts
29. Dépannage WSUS (clients qui ne remontent pas, reset base)
30. Partages réseau : SMB, création, bonnes pratiques
31. Permissions NTFS vs permissions de partage : la combinaison
32. Cas pratiques de permissions (tableau de vérité)
33. DFS-N : espaces de noms
34. DFS-R : réplication
35. FSRM : quotas et filtrage de fichiers
36. Déduplication des données
37. Dépannage d'accès aux fichiers (méthode en 7 étapes)
38. Serveur d'impression : rôle et architecture
39. Déploiement d'imprimantes et copieurs via GPO
40. Pilotes Type 3 vs Type 4, Print Management
41. Files d'attente, spooler : redémarrage et migration
42. Dépannage d'impression (lien métier copieurs)
43. RDS : les 4 rôles (Broker, Session Host, Passerelle, Accès Web)
44. Déploiement RDS pas à pas
45. Licences CAL RDS : par utilisateur vs par appareil
46. Dépannage RDS
47. AD CS / PKI : concepts (AC racine, subordonnée, modèles)
48. Installation d'une AC d'entreprise pas à pas
49. Modèles de certificats et auto-enrollment via GPO
50. Renouvellement et révocation (CRL, OCSP)
51. Dépannage PKI
52. Sauvegarde : Windows Server Backup
53. Planifications, bare metal, restauration
54. Stratégie 3-2-1 et rotation des supports
55. Restauration : scénarios pas à pas
56. Veeam et alternatives : quand passer au niveau supérieur
57. Licences : KMS vs ADBA
58. CAL utilisateur vs appareil, inventaire
59. Activation : slmgr, DISM, dépannage
60. Durcissement de base (hardening)
61. Pare-feu Windows Defender : règles et profils
62. Microsoft Defender Antivirus sur serveur
63. LAPS et gestion des mots de passe locaux
64. Audit, journaux et supervision minimale
65. Cas pratiques commentés (1 à 5)
66. Cas pratiques commentés (6 à 10)
67. Cas pratiques commentés (11 à 15)
68. Cas pratiques commentés (16 à 18)
69. Erreurs classiques (1 à 5)
70. Erreurs classiques (6 à 10)
71. Erreurs classiques (11 à 15)
72. Erreurs classiques (16 à 18)
73. Checklist d'exploitation quotidienne
74. Checklist hebdomadaire
75. Checklist mensuelle
76. Checklist trimestrielle / annuelle
77. Checklist de mise en production d'un nouveau serveur
78. Pense-bête de poche : PowerShell
79. Pense-bête de poche : commandes et raccourcis
80. Pense-bête de poche : ports réseau essentiels
81. Glossaire (A–M)
82. Glossaire (N–Z)
83. Quiz : 10 questions
84. Quiz : réponses commentées
85. Pour aller plus loin : 20 ressources et chantiers

---

## 1. Éditions et canaux de distribution

Windows Server existe en deux canaux :

| Canal | Versions | Support | Usage |
|---|---|---|---|
| **LTSC** (Long-Term Servicing Channel) | 2019, 2022, 2025 | 10 ans (5 mainstream + 5 étendu) | Production, charges stables, serveurs d'infrastructure |
| **Canaux annuels** (anciennement SAC) | ex. 23H2, 25H2 | 24 mois | Conteneurs, dev/test, environnements agiles |

**Règle d'or en entreprise : toujours LTSC pour la production.** Les versions annuelles servent aux conteneurs et aux bancs d'essai. Windows Server 2025 est la LTSC actuelle ; 2022 reste le choix le plus éprouvé en 2026 ; 2019 entre en support étendu (correctifs de sécurité uniquement).

Éditions principales (LTSC) :

| Édition | Cible | Différence clé |
|---|---|---|
| **Standard** | PME, peu de virtualisation | 2 VM (OSE) incluses par licence |
| **Datacenter** | Virtualisation intensive | VM illimitées |
| **Essentials** | ≤ 25 utilisateurs / 50 appareils | Rôles limités, pas de virtualisation (2019 = dernière vraie Essentials) |

> Note 2025 : l'édition Essentials n'existe plus en 2025. Les petites structures prennent Standard.

---

## 2. Standard vs Datacenter : le tableau qui décide

| Critère | Standard | Datacenter |
|---|---|---|
| VM (OSE) par licence | **2** | **Illimitées** |
| Conteneurs Hyper-V illimités | Non | Oui |
| Storage Spaces Direct (S2D) | Non | Oui |
| Réplica de stockage (Storage Replica) | Limité (1 partenariat, 1 volume ≤ 2 To) | Illimité |
| Réseau SDN complet (Network Controller) | Non | Oui |
| Machines virtuelles protégées (Shielded VM) | Non (hôte uniquement) | Oui |
| Prix (ordre de grandeur) | ~1 000 € / 16 cœurs | ~6 000 € / 16 cœurs |

**Calcul de rentabilité :** si un hôte Hyper-V fait tourner plus de ~6 VM en production durable, Datacenter est généralement moins cher que d'empiler des licences Standard (chaque licence Standard = 2 VM ; il faut re-licencier par tranches de 2 VM).

```powershell
# Vérifier l'édition installée
Get-ComputerInfo | Select-Object WindowsProductName, WindowsEditionId, WindowsVersion

# Changer d'édition (montée en gamme Standard -> Datacenter, sans réinstallation)
DISM /Online /Set-Edition:ServerDatacenter /ProductKey:XXXXX-XXXXX-XXXXX-XXXXX-XXXXX /AcceptEula
```

> La montée en gamme Standard → Datacenter fonctionne. La descente Datacenter → Standard est **impossible** sans réinstallation.

---

## 3. Checklist de choix d'édition (10 critères)

1. [ ] Nombre de VM prévues par hôte sur 5 ans → > 6 = Datacenter
2. [ ] Storage Spaces Direct prévu ? → Datacenter obligatoire
3. [ ] Réplication de stockage multi-sites ? → Datacenter (au-delà des limites Standard)
4. [ ] SDN / Network Controller ? → Datacenter
5. [ ] Budget licences sur 5 ans (comparer 3× Standard vs 1× Datacenter)
6. [ ] Modèle de licence : 16 cœurs minimum par serveur, par tranches de 2/16 cœurs
7. [ ] CAL : prévoir CAL Utilisateur ou Appareil pour chaque accès (voir §58)
8. [ ] Canal LTSC confirmé (pas de version annuelle en prod)
9. [ ] Compatibilité applicative : l'éditeur supporte-t-il 2025 ? (sinon 2022)
10. [ ] Stratégie de sortie : 2019 → planifier la migration (support étendu = sécurité seule)

---

## 4. Installation pas à pas (Desktop Experience)

