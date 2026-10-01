---
id: collect-261001-general-networking/general-networking/double-authentification-mfa-13-etapes-90-min-2026-3
title: "deploy-mfa-hardening.ps1"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/double-authentification-mfa-13-etapes-90-min-2026.md
source_anchor: ""
source_lines: [101, 219]
sha256: ae6a8601e39af50d5bcef8bbc7946bbc344bc886d1ad4ced6e9529a16225e78a
---

# deploy-mfa-hardening.ps1

```
$params = @{
    displayName = "Exiger le MFA pour tous les utilisateurs"
    state       = "enabledForReportingButNotEnforced"
    conditions  = @{
        users        = @{ includeUsers = @("All"); excludeUsers = @("
```
") }
        applications = @{ includeApplications = @("All") }
    }
    grantControls = @{
        operator        = "OR"
        builtInControls = @("mfa")
    }
}
New-MgIdentityConditionalAccessPolicy -BodyParameter $params Le paramètre `state` défini sur `enabledForReportingButNotEnforced` journalise l’impact de la politique sans bloquer personne. Cela permet de vérifier pendant quelques jours qu’aucun compte légitime n’est pénalisé avant de basculer sur `enabled`.

**Étape 7 : ajoutez une méthode résistante au phishing avec une clé de sécurité FIDO2.** Activez la méthode FIDO2 au niveau de la politique d’authentification du tenant, puis distribuez des clés physiques aux comptes à privilèges.

```
$body = @{
    "@odata.type" = "#microsoft.graph.fido2AuthenticationMethodConfiguration"
    state = "enabled"
    isSelfServiceRegistrationAllowed = $true
    isAttestationEnforced = $true
    keyRestrictions = @{
        isEnforced = $false
        enforcementType = "allow"
        aaGuids = @()
    }
} | ConvertTo-Json -Depth 5
Invoke-MgGraphRequest -Method PATCH `
    -Uri "https://graph.microsoft.com/v1.0/policies/authenticationMethodsPolicy/authenticationMethodConfigurations/Fido2" `
    -Body $body -ContentType "application/json"
```
Le paramètre `isAttestationEnforced` réglé sur `$true` garantit que seules des clés certifiées FIDO2, et non des implémentations logicielles non vérifiées, peuvent être enregistrées. Un détail qui compte pour les comptes à haut risque.

**Étape 8 : activez Windows Hello for Business sur les postes Windows 11.** Pour les postes joints à Entra ID, activez Windows Hello for Business via Intune ou une stratégie de groupe. Les utilisateurs se connectent alors avec leur empreinte, leur visage ou un code PIN local, chaque facteur étant lié cryptographiquement à la puce TPM de l’appareil. Cette méthode couvre le même niveau de résistance au phishing que FIDO2, sans matériel supplémentaire à acheter.

**Étape 9 : désactivez progressivement le SMS et l’appel vocal.** Une fois qu’une majorité d’utilisateurs a basculé vers Authenticator ou une clé FIDO2, restreignez les méthodes les plus faibles.

```
$body = @{
    "@odata.type" = "#microsoft.graph.smsAuthenticationMethodConfiguration"
    state = "disabled"
} | ConvertTo-Json
Invoke-MgGraphRequest -Method PATCH `
    -Uri "https://graph.microsoft.com/v1.0/policies/authenticationMethodsPolicy/authenticationMethodConfigurations/Sms" `
    -Body $body -ContentType "application/json"
```
Ne réalisez cette étape qu’après avoir vérifié, via le rapport d’inscription généré à l’étape 2, qu’aucun utilisateur actif ne dépend plus exclusivement du SMS.

### Étapes 10 à 13 : verrouiller, automatiser et vérifier

**Étape 10 : bloquez les protocoles d’authentification hérités.** Les protocoles comme POP, IMAP basique ou SMTP AUTH ne savent pas transmettre de demande MFA : ils constituent donc une porte dérobée qui contourne entièrement toute politique mise en place plus haut. Créez une politique d’accès conditionnel dédiée qui bloque l’authentification héritée pour tous les utilisateurs, sauf exception documentée pour un système legacy identifié.

**Étape 11 : activez l’accès conditionnel basé sur le risque.** Avec une licence Entra ID P2, activez Identity Protection pour évaluer en temps réel le niveau de risque de chaque connexion (adresse IP inhabituelle, déplacement impossible, jeton potentiellement volé) et exiger automatiquement une réauthentification MFA, ou bloquer la connexion quand le risque dépasse un seuil moyen.

**Étape 12 : générez des Temporary Access Pass pour l’inscription à distance.** Pour les nouveaux employés ou les utilisateurs qui perdent l’accès à leurs méthodes existantes, générez un Temporary Access Pass à durée de vie limitée depuis le portail Entra ID. Ce code temporaire permet d’enregistrer une nouvelle méthode MFA sans jamais transmettre de mot de passe par téléphone ou par e-mail, un vecteur d’ingénierie sociale classique.

**Étape 13 : testez le parcours utilisateur et surveillez les journaux de connexion.** Connectez-vous avec un compte test pour valider chaque méthode de bout en bout, puis mettez en place une surveillance continue des journaux de connexion Entra ID, idéalement remontés dans Microsoft Sentinel ou un autre SIEM.

```
SigninLogs
| where ResultType == 0
| where AuthenticationRequirement == "multiFactorAuthentication"
| where TimeGenerated > ago(1d)
| summarize Connexions = count() by UserPrincipalName, AppDisplayName
| where Connexions > 20
| order by Connexions desc
```
Cette requête repère les comptes qui accumulent un nombre anormal de connexions MFA réussies en une journée, un signal parfois révélateur d’un jeton volé rejoué en boucle par un script automatisé.

## Script complet : automatiser le déploiement avec Microsoft Graph PowerShell

Pour éviter de reproduire chaque commande une à une sur plusieurs tenants, voici un script qui regroupe l’audit initial, la création de la politique d’accès conditionnel et l’activation de FIDO2 en une seule exécution. Remplacez l’identifiant du compte break-glass par le vôtre, et testez-le systématiquement sur un tenant de démonstration avant la production.

```
# deploy-mfa-hardening.ps1
# Script complet : audit, politique d'accès conditionnel et FIDO2
param(
    [string]$BreakGlassAccountId = "
```
"
)
Connect-MgGraph -Scopes "Reports.Read.All","Policy.ReadWrite.ConditionalAccess","Policy.ReadWrite.AuthenticationMethod"
Write-Host "Étape 1/3 : audit des méthodes MFA enregistrées..."
Get-MgReportAuthenticationMethodUserRegistrationDetail -All |
    Select-Object UserPrincipalName, IsMfaRegistered, IsPasswordlessCapable |
    Export-Csv -Path .\audit-mfa-2026.csv -NoTypeInformation -Encoding UTF8
Write-Host "Rapport enregistré dans audit-mfa-2026.csv"
Write-Host "Étape 2/3 : création de la politique d'accès conditionnel (mode audit)..."
$caParams = @{
    displayName = "Exiger le MFA pour tous les utilisateurs"
    state       = "enabledForReportingButNotEnforced"
    conditions  = @{
        users        = @{ includeUsers = @("All"); excludeUsers = @($BreakGlassAccountId) }
        applications = @{ includeApplications = @("All") }
    }
    grantControls = @{ operator = "OR"; builtInControls = @("mfa") }
}
New-MgIdentityConditionalAccessPolicy -BodyParameter $caParams
Write-Host "Politique créée en mode rapport uniquement."
Write-Host "Étape 3/3 : activation de FIDO2 comme méthode résistante au phishing..."
$fido2Body = @{
    "@odata.type" = "#microsoft.graph.fido2AuthenticationMethodConfiguration"
    state = "enabled"
    isSelfServiceRegistrationAllowed = $true
    isAttestationEnforced = $true
} | ConvertTo-Json -Depth 5
Invoke-MgGraphRequest -Method PATCH `
    -Uri "https://graph.microsoft.com/v1.0/policies/authenticationMethodsPolicy/authenticationMethodConfigurations/Fido2" `
    -Body $fido2Body -ContentType "application/json"
Write-Host "Déploiement terminé. Vérifiez le portail Entra ID avant de passer en mode enforced." Enregistrez ce script sous le nom `deploy-mfa-hardening.ps1` et exécutez-le avec un compte disposant des rôles Administrateur d’accès conditionnel et Administrateur de méthode d’authentification. Le script journalise chaque action dans la console, ce qui facilite l’audit après coup.

## Résultats attendus : exemples de sortie et vérifications

Une fois le script exécuté, la console affiche une progression semblable à celle-ci :

