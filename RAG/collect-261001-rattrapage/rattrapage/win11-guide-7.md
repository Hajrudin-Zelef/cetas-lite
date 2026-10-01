---
id: collect-261001-rattrapage/rattrapage/win11-guide-7
title: "Windows 11 en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/win11_guide.md
source_anchor: ""
source_lines: [793, 1004]
sha256: b17a23625b969212cd92770c35b1f6c8b2fb40450081b33365007efe6b4bd687
---

# Windows 11 en entreprise — Guide technique ultra-complet

[Default]
OSInstall=Y
SkipCapture=YES
SkipAdminPassword=YES
SkipProductKey=YES
SkipComputerBackup=YES
SkipBitLocker=NO
SkipLocaleSelection=YES
SkipTimeZone=YES
SkipApplications=NO
SkipSummary=NO
SkipFinalSummary=YES
KeyboardLocale=fr-FR
UserLocale=fr-FR
UILanguage=fr-FR
TimeZoneName=Romance Standard Time
; Romance Standard Time = fuseau Paris (UTC+1, gère l'heure d'été)
JoinDomain=entreprise.local
DomainAdmin=svc-join-domain
DomainAdminDomain=ENTREPRISE
; Le mot de passe du compte de jointure va dans Bootstrap.ini (chiffré si possible)
MachineObjectOU=OU=Postes Win11,OU=Ordinateurs,DC=entreprise,DC=local
SLShare=\\SRV-MDT\Logs$
EventService=http://SRV-MDT:9800
```

`Bootstrap.ini` (`D:\DeploymentShare\Control\Bootstrap.ini`) :

```ini
[Settings]
Priority=Default

[Default]
DeployRoot=\\SRV-MDT\DeploymentShare$
UserDomain=ENTREPRISE
UserID=svc-mdt
UserPassword=MotDePasseFictif_A_Remplacer
SkipBDDWelcome=YES
KeyboardLocale=fr-FR
```

> 🔐 `svc-mdt` et `svc-join-domain` : comptes de service dédiés, mots de passe complexes, **jamais** de compte admin du domaine. Le compte de jointure doit avoir uniquement le droit « Joindre des ordinateurs au domaine » délégué sur l'OU cible.

### 11.4 Applications dans MDT

Applications → New Application → « Application with source files » ou « Application bundle » :

```powershell
# Exemple : installation silencieuse 7-Zip (commande dans l'application MDT)
# Working directory : .\Applications\7-Zip
# Command line : 7z2408-x64.exe /S

# Exemple : winget en task sequence (section 28)
# Command line : powershell.exe -ExecutionPolicy Bypass -File Install-Apps.ps1
```

### 11.5 Journaux MDT (dépannage)

| Emplacement | Contenu |
|---|---|
| `C:\MININT\SMSOSD\OSDLOGS\BDD.log` | Journal agrégé (le plus utile) |
| `C:\MININT\SMSOSD\OSDLOGS\SMSTS.log` | Task sequence engine |
| `\\SRV-MDT\Logs$` (SLShare) | Copie réseau en temps réel |
| `X:\MININT\...` (WinPE) | Pendant la phase WinPE |

---

## 12. Autopilot + Intune : principe, profils, inscription

### 12.1 Principe (ce qui change par rapport à MDT)

| MDT / image | Autopilot |
|---|---|
| Image maître à maintenir | **Aucune image** : l'OEM livre Windows 11, Autopilot transforme |
| Réseau local, PXE | Cloud, l'utilisateur déballe et allume |
| Jointure AD classique | Entra join (+ hybride possible) |
| Applications via task sequence | Intune (applications, stratégies, conformité) |

Flux Autopilot :

```
1. Le fournisseur enregistre le hardware hash dans Autopilot (ou import CSV manuel)
2. L'appareil est affecté à un profil Autopilot (groupe Entra)
3. L'utilisateur allume → OOBE → connexion Entra ID
4. Intune pousse : stratégies, certificats, applications, conformité
5. L'utilisateur arrive sur un poste configuré (User-Driven ou Self-Deploying)
```

### 12.2 Enregistrer un appareil (import manuel)

```powershell
# Sur le poste (en OOBE : Shift+F10 pour ouvrir une invite)
# Récupérer le script officiel :
Install-Script -Name Get-WindowsAutopilotInfo -Force
# Générer le CSV :
Get-WindowsAutopilotInfo.ps1 -OutputFile C:\Temp\autopilot.csv
# Importer le CSV dans Intune → Appareils → Inscription → Windows Autopilot
```

Format CSV attendu : `Device Serial Number, Windows Product ID, Hardware Hash, Group Tag, Assigned User`.

### 12.3 Profils de déploiement Autopilot (réglages clés)

Dans Intune → Appareils → Inscription → Profils de déploiement :

| Paramètre | Recommandé entreprise |
|---|---|
| Mode de déploiement | Piloté par l'utilisateur (standard) |
| Joindre à | Microsoft Entra ID (pur cloud) ou Hybride |
| Type de compte | Standard (jamais admin local pour l'utilisateur) |
| Langue / clavier | fr-FR |
| Nom d'appareil | Modèle : `PC-%SERIAL%` ou `ENT-%RAND:6%` |
| Masquer les écrans de confidentialité | Oui |
| Masquer le changement de compte | Oui |

### 12.4 Page d'état d'inscription (ESP)

L'ESP bloque l'accès au bureau tant que les éléments critiques ne sont pas installés :

- Applications « bloquantes » (ex. : Defender config, VPN, Office) ;
- Stratégies de sécurité ;
- Certificats.

> 💡 Ne mettez en bloquant **que** le strict nécessaire : chaque application bloquante allonge l'OOBE et chaque échec bloque l'utilisateur. Le reste s'installe en arrière-plan après l'ouverture de session.

### 12.5 Groupes dynamiques Entra (affectation automatique)

```powershell
# Exemple de règle d'appartenance dynamique pour les appareils Autopilot :
# (device.devicePhysicalIDs -any (_ -contains "[ZTDId]"))
# Tous les appareils Autopilot → groupe "GRP-Autopilot-All"

# Par profil (Group Tag) :
# (device.devicePhysicalIDs -any (_ -eq "[OrderID]:Compta"))
```

---

## 13. Autopilot Reset, Fresh Start et scénarios de réaffectation

### 13.1 Comparatif des méthodes de réinitialisation

| Méthode | Conserve l'inscription Entra/Autopilot | Conserve les données | Usage |
|---|---|---|---|
| Autopilot Reset | ✅ (réinscription auto) | ❌ | Réaffectation à un nouvel utilisateur |
| Réinitialiser ce PC (paramètres) | ❌ | Au choix | Poste hors gestion |
| Fresh Start (Intune) | ❌ (réinscription) | ❌ | Repartir de zéro, supprime les apps préinstallées |
| Reimage MDT | ❌ | ❌ | Retour au socle MDT |

### 13.2 Déclencher un Autopilot Reset

```powershell
# Méthode 1 : depuis Intune (Appareils → ... → Autopilot Reset)
# Méthode 2 : en local (admin)
systemreset.exe --factoryreset
# puis choisir "Autopilot Reset" si proposé

# Méthode 3 : à distance via Intune (action groupée sur N appareils)
```

Processus détaillé :

```
1. L'appareil reçoit la commande (Intune ou locale)
2. Redémarrage → suppression des applications, paramètres, données utilisateurs
3. L'inscription Entra ID et l'enregistrement Autopilot sont CONSERVÉS
4. OOBE Autopilot → nouvel utilisateur se connecte
5. Intune réapplique profils + applications
Durée typique : 15 à 30 minutes.
```

### 13.3 Scénario de réaffectation (procédure type)

1. Sauvegarder les données de l'ancien utilisateur (OneDrive Known Folder Move = automatique).
2. Révoquer les sessions (Entra → Révoquer les sessions) si départ sensible.
3. Autopilot Reset via Intune.
4. Affecter le nouvel utilisateur dans le profil Autopilot.
5. Le nouvel utilisateur allume et se connecte.

---

## 14. Jointure au domaine vs Entra join vs hybride

### 14.1 Tableau de décision

| Critère | AD joint (classique) | Entra join (pur cloud) | Hybride (Hybrid Entra join) |
|---|---|---|---|
| Prérequis réseau | Ligne vers les DC | Internet | Les deux |
| Gestion | GPO | Intune | GPO + Intune (co-management) |
| SSO applications cloud | Via AD FS / Seamless SSO | Natif (PRT) | Natif après sync |
| Ressources locales (partages, impression) | Natif (Kerberos) | Via VPN/Cloud (Kerberos Cloud Trust) | Natif |
| Idéal pour | Parc sédentaire existant | Nomades, nouveaux parcs, filiales sans DC | Transition, contraintes locales fortes |

**Recommandation 2026** : pour un nouveau déploiement, visez **Entra join + Intune**. L'hybride est un état transitoire : il cumule les complexités des deux mondes (double inscription, délais de sync, dépannage plus difficile).

### 14.2 Jointure au domaine (classique)

```powershell
# Joindre le domaine en PowerShell
Add-Computer -DomainName "entreprise.local" -OUPath "OU=Postes Win11,DC=entreprise,DC=local" -Credential (Get-Credential) -Restart

# Vérifier
Get-CimInstance Win32_ComputerSystem | Select-Object PartOfDomain, Domain

# Quitter le domaine (avant réaffectation)
Remove-Computer -UnjoinDomainCredential (Get-Credential) -Restart
```

### 14.3 Entra join (manuel, hors Autopilot)

Paramètres → Comptes → Accès professionnel ou scolaire → Se connecter → « Joindre cet appareil à Microsoft Entra ID ».

