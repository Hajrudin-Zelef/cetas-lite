---
id: collect-261001-rattrapage/rattrapage/win11-guide-11
title: "Windows 11 en entreprise — Guide technique ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["incident", "valuation"]
source: docs/RAG/collect-261001-rattrapage/win11_guide.md
source_anchor: ""
source_lines: [1542, 1730]
sha256: dcd20202fb6921054f8d7288181e0cfb19645e8ee5f68ea3735eea2689e178f7
---

# Windows 11 en entreprise — Guide technique ultra-complet

| Réglage | GPO |
|---|---|
| Désactiver les recommandations du Démarrer | `Menu Démarrer et barre des tâches` → « Supprimer la liste des recommandations » |
| Désactiver les suggestions d'applications | « Désactiver les suggestions dans le menu Démarrer » |
| Désactiver le contenu recommandé dans l'Explorateur | Stratégies de l'Explorateur de fichiers |
| Spotlight / arrière-plans « à la une » | Personnalisation → empêcher le changement |

```powershell
# Regroupé en script (extrait) :
$policies = @{
    "HKLM:\SOFTWARE\Policies\Microsoft\Windows\CloudContent" = @{
        "DisableWindowsSpotlightFeatures" = 1
        "DisableTailoredExperiencesWithDiagnosticData" = 1
    }
    "HKLM:\SOFTWARE\Policies\Microsoft\Windows\Explorer" = @{
        "HideRecommendedSection" = 1
    }
}
foreach ($path in $policies.Keys) {
    New-Item -Path $path -Force | Out-Null
    foreach ($name in $policies[$path].Keys) {
        Set-ItemProperty -Path $path -Name $name -Value $policies[$path][$name] -Type DWord
    }
}
```

---

## 24. GPO Windows 11 : bonnes pratiques, héritage, filtrage, dépannage

### 24.1 Structurer ses OU et ses GPO

```
DC=entreprise,DC=local
 └─ OU=Postes de travail
     ├─ OU=Win11-Bureautique
     ├─ OU=Win11-Nomades
     ├─ OU=Win11-Kiosques
     └─ OU=Serveurs (séparé !)
```

Règles :

1. **Une GPO = un objectif** (« GPO-BitLocker-OS », « GPO-StartLayout », …) : on désactive sans tout casser.
2. Nommer avec préfixe : `GPO-[Cible]-[Fonction]` (ex. `GPO-Win11-Defender-ASR`).
3. Lier au plus près de la cible, éviter le filtrage de sécurité complexe.
4. Toujours tester sur une OU pilote avant de lier au niveau supérieur.

### 24.2 Filtrage WMI (ex. : cibler Windows 11 uniquement)

```wql
-- Filtre WMI "Windows 11 uniquement" :
SELECT * FROM Win32_OperatingSystem WHERE Version LIKE "10.0.22%" OR Version LIKE "10.0.26%"
-- (Windows 11 = version 10.0.22000+, builds 22621/22631 = 22H2/23H2, 26100 = 24H2)
```

> ⚠️ Les filtres WMI ralentissent l'ouverture de session s'ils sont nombreux ou mal écrits. Préférez les OU dédiées quand c'est possible.

### 24.3 Dépannage : la GPO ne s'applique pas (voir aussi cas pratique 10)

```powershell
# 1. Forcer l'actualisation
gpupdate /force

# 2. Voir les GPO appliquées (résumé)
gpresult /r

# 3. Rapport HTML détaillé (à ouvrir dans Edge)
gpresult /h C:\Admin\gpresult.html

# 4. Simuler pour un autre utilisateur / poste (console GPMC)
# GPMC → "Résolution des problèmes de stratégie de groupe" → Assistant

# 5. Vérifier la réplication SYSVOL entre DC
dcdiag /test:frssysvol
repadmin /syncall /AdeP
```

---

## 25. Windows Update for Business : anneaux de déploiement

### 25.1 Le modèle des anneaux (rings)

```
Anneau 0 — Pilotes/IT      : 0 jour de délai   (5-10 postes, l'équipe IT)
Anneau 1 — Testeurs        : 7 jours           (5 % du parc, volontaires)
Anneau 2 — Production large: 14-21 jours       (le reste du parc)
Anneau 3 — Critiques       : 30 jours + pause  (postes sensibles : caisses, supervision)
```

> 💡 Les **mises à jour de qualité** (mensuelles, Patch Tuesday) et les **mises à jour de fonctionnalités** (annuelles) ont des délais séparés. Les MàJ de qualité : délai court (7-14 j). Les MàJ de fonctionnalités : délai long (60-120 j) pour laisser mûrir.

### 25.2 Définir les anneaux (Intune)

Intune → Appareils → Mises à jour Windows → **Anneaux de mise à jour** :

| Paramètre | Anneau Testeurs | Anneau Production |
|---|---|---|
| MàJ qualité — délai | 7 jours | 14 jours |
| MàJ fonctionnalités — délai | 30 jours | 90 jours |
| Canal de maintenance | General Availability | General Availability |
| Pilotes | Autoriser | Bloquer (déployer via constructeur) |
| Désinstaller auto si problème | Oui | Oui |

### 25.3 Définir les anneaux (GPO — AD joint)

`Configuration ordinateur\Modèles d'administration\Composants Windows\Windows Update\Gérer les mises à jour proposées par Windows Update` :

- « Sélectionner quand les mises à jour qualité sont reçues » : délai en jours.
- « Sélectionner quand les builds d'évaluation et les mises à jour de fonctionnalités sont reçus » : délai en jours.
- « Mettre en pause les mises à jour » : en cas d'incident majeur.

```powershell
# Équivalent registre (à n'utiliser qu'en scripté, préférez la GPO) :
$wu = "HKLM:\SOFTWARE\Policies\Microsoft\Windows\WindowsUpdate\AU"
New-Item -Path $wu -Force | Out-Null
Set-ItemProperty -Path $wu -Name "DeferFeatureUpdates" -Value 1 -Type DWord
Set-ItemProperty -Path $wu -Name "DeferFeatureUpdatesPeriodInDays" -Value 90 -Type DWord
Set-ItemProperty -Path $wu -Name "DeferQualityUpdates" -Value 1 -Type DWord
Set-ItemProperty -Path $wu -Name "DeferQualityUpdatesPeriodInDays" -Value 14 -Type DWord
```

---

## 26. WUfB : stratégies (CSP, GPO, Intune), délais, pause, qualité vs fonctionnalités

### 26.1 Les trois types de mises à jour

| Type | Contenu | Fréquence | Délai conseillé |
|---|---|---|---|
| Qualité (quality) | Correctifs sécurité + bugs cumulatifs | Mensuel (2e mardi) | 7-14 j |
| Fonctionnalités (feature) | Nouvelle version annuelle (24H2…) | 1/an | 60-120 j |
| Pilotes / firmware | Via Windows Update (optionnel) | Continu | Bloquer en prod, canal constructeur |

### 26.2 Mettre en pause (gestion d'incident)

```powershell
# Pause via registre (équivalent GPO "Mettre en pause") :
$wu = "HKLM:\SOFTWARE\Microsoft\WindowsUpdate\UX\Settings"
Set-ItemProperty -Path $wu -Name "PauseFeatureUpdatesStartTime" -Value (Get-Date -Format "yyyy-MM-dd") -Type String
Set-ItemProperty -Path $wu -Name "PauseQualityUpdatesStartTime"  -Value (Get-Date -Format "yyyy-MM-dd") -Type String
# Reprendre : supprimer ces valeurs ou via Paramètres → Windows Update.
```

Procédure d'incident « MàJ qui casse » :

```
1. Mettre en pause l'anneau Production (GPO/Intune) — 7 jours
2. Identifier la KB fautive (journaux, retours testeurs)
3. Tester le correctif / la désinstallation sur l'anneau pilote
4. Si nécessaire : wusa /uninstall /kb:XXXXXXX (section 45)
5. Lever la pause par anneau, en cascade
```

### 26.3 Pilotes via Windows Update : oui ou non ?

- **En entreprise : NON** par défaut. Les pilotes doivent venir du constructeur (Dell Command Update, Lenovo System Update, HP Image Assistant) validés par l'IT.
- GPO : « Ne pas inclure les pilotes avec les mises à jour Windows » → Activé.
- Exception : périphériques exotiques sans canal constructeur.

### 26.4 Heures d'activité et redémarrages

```powershell
# Forcer les heures d'activité 8h-19h (les MàJ ne redémarrent pas dedans)
$wu = "HKLM:\SOFTWARE\Policies\Microsoft\Windows\WindowsUpdate"
Set-ItemProperty -Path $wu -Name "SetActiveHours" -Value 1 -Type DWord
Set-ItemProperty -Path $wu -Name "ActiveHoursStart" -Value 8 -Type DWord
Set-ItemProperty -Path $wu -Name "ActiveHoursEnd" -Value 19 -Type DWord
# Notification de redémarrage :
Set-ItemProperty -Path $wu -Name "AutoRestartRequiredNotificationDismissal" -Value 2 -Type DWord
```

---

## 27. WSUS vs Windows Update for Business : comparatif et coexistence

### 27.1 Comparatif

| Critère | WSUS | WUfB |
|---|---|---|
| Infrastructure | Serveur(s) WSUS, stockage local (To) | Aucune (cloud Microsoft) |
| Bande passante | Optimisée en LAN (1 téléchargement) | Delivery Optimization (P2P LAN) |
| Contrôle fin (approuver KB par KB) | ✅ | ❌ (délais uniquement) |
| Postes nomades / télétravail | Compliqué (VPN) | ✅ natif |
| Coût | Licences serveur + stockage + admin | Inclus (Intune/GPO) |
| Rapports | Console WSUS (limitée) | Update Compliance / Intune |
| Avenir | En maintenance, **déprécié à terme** | **Stratégie Microsoft** |

