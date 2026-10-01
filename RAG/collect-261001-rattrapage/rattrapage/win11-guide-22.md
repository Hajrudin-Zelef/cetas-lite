---
id: collect-261001-rattrapage/rattrapage/win11-guide-22
title: "Windows 11 en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Intel", "Microsoft"]
dates: []
keywords: ["energy", "ethernet", "gpu", "incident", "intel"]
source: docs/RAG/collect-261001-rattrapage/win11_guide.md
source_anchor: ""
source_lines: [3704, 3860]
sha256: 14e1c60d47c8bd1c05086cd77aed46c22765d15f16791cd8cd06711804041eb3
---

# Windows 11 en entreprise — Guide technique ultra-complet

| Stratégie | Effet |
|---|---|
| Configure favorites / ManagedFavorites | Favoris d'entreprise imposés |
| HomepageLocation / RestoreOnStartup | Page d'accueil intranet |
| PasswordManagerEnabled | Désactiver le gestionnaire de mots de passe intégré (si coffre d'entreprise) |
| SmartScreenEnabled | Forcer SmartScreen |
| IE Mode (InternetExplorerIntegration) | Compatibilité applications legacy |

```powershell
# Exemple : forcer la page d'accueil intranet via registre
$edge = "HKLM:\SOFTWARE\Policies\Microsoft\Edge"
New-Item -Path $edge -Force | Out-Null
Set-ItemProperty -Path $edge -Name "HomepageLocation" -Value "https://intranet.entreprise.local" -Type String
Set-ItemProperty -Path $edge -Name "HomepageIsNewTabPage" -Value $false -Type DWord
```

### 70.3 Profils et synchronisation

- Synchronisation du profil Edge via le compte Entra (favoris, mots de passe) : pratique, à encadrer (données hors UE ? → vérifier la résidence des données du tenant).
- Pour les postes partagés : **désactiver la synchronisation** et effacer les données à la fermeture (stratégie `ClearBrowsingDataOnExit`).

---

## 71. Sauvegarde et PRA poste de travail

### 71.1 Ce qu'on sauvegarde (et ce qu'on ne sauvegarde pas)

| Données | Solution | RPO |
|---|---|---|
| Fichiers utilisateurs | OneDrive KFM (Bureau/Documents/Images) | Continu |
| Profils spécifiques (métier) | Dossier redirigé ou script de copie vers partage | Quotidien |
| Configuration poste | **Rien** → redéploiement MDT/Autopilot | N/A |
| Clés BitLocker | AD / Entra (automatique) | À l'activation |

### 71.2 PRA « poste de travail » (procédure type)

```
INCIDENT : poste volé / HS / ransomware localisé
1. Isoler : désactiver le compte AD/Entra, révoquer les sessions, effacer à distance (Intune → Effacer)
2. Clés BitLocker : vérifier qu'elles sont en AD/Entra (récupération si besoin forensique)
3. Remplacer : déployer un poste neuf (MDT/Autopilot) — objectif < 4 h
4. Restaurer : l'utilisateur se connecte → OneDrive resynchronise → apps via Intune/winget
5. Post-mortem : cause, ticket GLPI, ajustement des protections
```

### 71.3 Effacement à distance (Intune)

Intune → Appareils → sélectionner → **Effacer** (avec ou sans conservation de l'inscription) ou **Retirer** (retire la gestion, conserve les données — pour un départ propre).

```powershell
# Effacement local sécurisé d'un disque avant recyclage (hors BitLocker)
cipher /w:C:
# Ou, plus radical : le chiffrement BitLocker + destruction de la clé = données irrécupérables
# (méthode "crypto-erase" : chiffrer puis supprimer les protecteurs)
```

---

## 72. Gestion de l'énergie : stratégies, veille moderne, batteries

### 72.1 Plans d'alimentation en entreprise

```powershell
# Voir les plans disponibles / actif
powercfg /list
# Activer un plan (GUID du plan "Performances élevées" en exemple)
powercfg /setactive 8c5e7fda-e8bf-4a96-9a85-a6e23a8c635c

# Imposer par GPO : Configuration ordinateur\Modèles d'administration\Système\Gestion de l'alimentation
# → "Spécifier un plan d'alimentation actif personnalisé"

# Empêcher la mise en veille pendant une tâche critique
powercfg /requestsoverride PROCESS monprocessus.exe Display System
```

### 72.2 Veille moderne (Modern Standby / S0)

Windows 11 privilégie la **veille moderne** (S0 low-power idle) sur les machines récentes : réveil instantané, MàJ en arrière-plan — mais **consommation résiduelle** et parfois réveils intempestifs.

```powershell
# Quel mode de veille supporte la machine ?
powercfg /a
# Voir les réveils intempestifs :
powercfg /waketimers
powercfg /lastwake

# Désactiver un périphérique qui réveille :
powercfg /devicedisablewake "Intel(R) Ethernet Connection"
```

### 72.3 Batteries (parc de portables)

```powershell
# Rapport de batterie (autonomie, usure)
powercfg /batteryreport /output C:\Admin\battery-report.html
# Indicateur clé : "Design Capacity" vs "Full Charge Capacity"
# → en dessous de 70-80 % de la capacité d'origine : batterie à remplacer

# Rapport d'énergie (diagnostic 60 s)
powercfg /energy /output C:\Admin\energy-report.html
```

> 💡 **Lien avec le métier énergie** : un parc de 200 portables en veille moderne consomme en permanence quelques watts chacun. Une GPO « mise en veille prolongée après 2 h » + extinction nocturne via tâche planifiée = économies mesurables. Voir le guide onduleurs pour le dimensionnement des protections électriques associées.

---

(Bloc 5/6 — sections 61 à 72)

---

## 73. Erreurs classiques (20) : ce qu'il ne faut pas faire

### Erreur n°1 — Migrer sans inventaire de compatibilité
Déployer Windows 11 « à l'aveugle » puis découvrir 40 % de postes sans TPM 2.0. **Toujours** commencer par le script de la section 3.2.

### Erreur n°2 — Utiliser le contournement de registre TPM/CPU en production
`AllowUpgradesWithUnsupportedTPMOrCPU` = pas de support, pas de MàJ garanties. En entreprise : les postes non compatibles se **remplacent**, ils ne se « bricolent » pas.

### Erreur n°3 — Activer BitLocker sans sauvegarder les clés
Chiffrement sans GPO de récupération = le premier écran bleu de récupération devient une perte de données. **Règle** : pas de clé en AD/Entra = pas de chiffrement.

### Erreur n°4 — Mettre à jour le BIOS sans suspendre BitLocker
Le classique du lundi matin : 15 appels « écran bleu BitLocker » après une campagne de MàJ BIOS. Procédure écrite + `Suspend-BitLocker -RebootCount 1` systématique.

### Erreur n°5 — Donner les droits d'administrateur local aux utilisateurs
« Pour qu'il puisse installer ses logiciels » = la porte ouverte aux ransomwares et aux configurations exotiques. Solutions : winget/Intune pour les logiciels validés, LAPS pour le dépannage.

### Erreur n°6 — Une seule GPO « fourre-tout »
Une GPO de 150 paramètres impossible à déboguer. **Une GPO = un objectif**, nommée `GPO-[Cible]-[Fonction]`.

### Erreur n°7 — Oublier les droits « Lire + Appliquer » sur le filtrage de sécurité
Depuis 2016, filtrer une GPO sur un groupe sans donner « Lire » aux « Ordinateurs authentifiés » = GPO silencieusement ignorée (voir cas pratique 10).

### Erreur n°8 — Tester les MàJ directement en production
Pas d'anneaux WUfB = la KB fautive touche 100 % du parc le même jour. Anneau pilote obligatoire (section 25).

### Erreur n°9 — Laisser les pilotes Windows Update actifs en production
Un pilote GPU poussé par WU un vendredi soir = 30 postes en BSOD le lundi. Pilotes via le constructeur, validés (section 26.3).

### Erreur n°10 — Sysprepper une machine jointe au domaine ou déjà utilisée
Image avec SID dupliqués, relations d'approbation cassées en série. Sysprep = **toujours** depuis une machine de référence propre, jamais jointe (section 9).

### Erreur n°11 — Supprimer un profil en effaçant juste `C:\Users\...`
Le registre garde la trace → profil temporaire au prochain logon. Utiliser `Win32_UserProfile` (sections 19.3, 49).

### Erreur n°12 — Stocker des mots de passe en clair dans unattend.xml / scripts
Le XML traîne sur un partage, dans un dépôt Git… Utiliser LAPS + comptes de service dédiés (sections 8.2, 34).

### Erreur n°13 — Injecter « tous les pilotes » dans MDT
Conflits, BSOD, task sequences de 3 heures. **Un dossier par modèle + profil de sélection** (section 11.2).

### Erreur n°14 — Négliger la période de rollback de 10 jours
`/ResetBase` ou suppression de `Windows.old` le jour J = impossible de revenir en arrière si un problème métier apparaît à J+5 (section 55).

### Erreur n°15 — Ne pas documenter les exceptions (DRA, exclusions Defender, GPO)
« Pourquoi ce poste n'est pas chiffré ? » — sans registre des exceptions, impossible de répondre à un audit. Tout écart = ticket + date de révision.

