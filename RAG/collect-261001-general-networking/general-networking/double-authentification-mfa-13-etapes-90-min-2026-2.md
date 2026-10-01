---
id: collect-261001-general-networking/general-networking/double-authentification-mfa-13-etapes-90-min-2026-2
title: "deploy-mfa-hardening.ps1"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Microsoft"]
dates: []
keywords: ["distribution"]
source: docs/RAG/collect-261001-general-networking/double-authentification-mfa-13-etapes-90-min-2026.md
source_anchor: ""
source_lines: [33, 100]
sha256: 579b284f005905088eed7a8d4758f609689a7daec108f5c21e6e280b11de8f27
---

# deploy-mfa-hardening.ps1

Le tutoriel qui suit permet justement de construire cette transition en douceur, sans bloquer vos utilisateurs du jour au lendemain. Si vous cherchez une alternative entièrement sans mot de passe, notre tutoriel dédié aux passkeys FIDO2 détaille la configuration côté utilisateur final.

## Panorama des méthodes d’authentification : ce qui fonctionne vraiment en 2026

Avant de choisir votre configuration cible, il vaut mieux comparer objectivement les méthodes disponibles sur Microsoft Entra ID. Le tableau ci-dessous résume leur niveau de résistance au phishing, leur complexité de déploiement et leur coût réel.

| Méthode | Résistance au phishing | Facilité de déploiement | Coût | Recommandation 2026 | 
|---|---|---|---|---|
| SMS / appel vocal (OTP) | Faible (interceptable, SIM swap) | Très simple | Frais opérateur | À éliminer en priorité | 
| Application TOTP (Authenticator) | Moyenne (vulnérable à l’AiTM) | Simple | Gratuit | Acceptable en transition | 
| Notification push + number matching | Moyenne à bonne | Simple | Gratuit | Bon compromis pour une PME | 
| Clé de sécurité FIDO2 / passkey | Élevée (liée au domaine) | Modérée (matériel à distribuer) | 25 à 60 € par clé | Recommandé pour comptes sensibles | 
| Windows Hello for Business | Élevée | Modérée (nécessite un join Entra ID) | Inclus dans la licence | Recommandé sur parc Windows 11 | 
| Certificat / carte à puce | Élevée | Complexe (PKI à gérer) | Variable | Secteur public et réglementé | 

Dans la majorité des PME, la combinaison la plus réaliste consiste à démarrer avec Microsoft Authenticator et le number matching pour l’ensemble des utilisateurs, puis à réserver les clés FIDO2 aux comptes à privilèges où le coût du matériel se justifie face au risque. C’est exactement la logique suivie dans les 13 étapes détaillées plus loin.

## Prérequis : comptes, licences et versions nécessaires

Réunissez les éléments suivants avant de commencer. Un déploiement mal préparé reste la première cause d’échec d’un projet MFA, bien avant la complexité technique elle-même.

| Élément | Version ou niveau minimal | Rôle dans le tutoriel | 
|---|---|---|
| Abonnement Microsoft 365 | Business Premium, E3/E5, ou Entra ID P1/P2 | Accès conditionnel et Identity Protection | 
| Système d’exploitation admin | Windows 11 ou macOS 14 et supérieur | Poste d’administration | 
| PowerShell | 7.4 LTS ou supérieur | Exécution des scripts | 
| Module Microsoft.Graph | Dernière version (PowerShell Gallery) | Automatisation via Microsoft Graph API | 
| Microsoft Authenticator | Dernière version (iOS et Android) | Méthode MFA principale | 
| Clé de sécurité FIDO2 (optionnel) | Certifiée FIDO2, ex. série YubiKey 5 | MFA résistant au phishing | 
| Rôle d’administrateur | Administrateur d’accès conditionnel ou de sécurité | Création des politiques | 

Si vous gérez moins de 300 utilisateurs et n’avez pas de licence Entra ID P1/P2, les valeurs de sécurité par défaut (Security Defaults) de Microsoft 365 constituent un point de départ correct : elles imposent un MFA basique gratuitement, sans la granularité de l’accès conditionnel décrite ici. Les étapes suivantes supposent une licence Entra ID P1 au minimum. Vérifiez également que votre compte dispose du rôle Administrateur d’accès conditionnel avant de passer à l’étape 1 : un rôle insuffisant provoque l’erreur la plus fréquente rencontrée en support, détaillée dans la section dépannage plus bas.

## Tutoriel : activer la double authentification en 13 étapes

Les treize étapes suivantes font passer un tenant sans politique MFA cohérente à une configuration qui impose une authentification résistante au phishing pour les comptes sensibles, tout en gardant une voie de secours pour ne jamais se retrouver bloqué dehors. Comptez environ 90 minutes pour l’ensemble, hors temps de distribution des clés physiques si vous en commandez.

### Étapes 1 à 4 : préparer le terrain avant le déploiement

**Étape 1 : vérifiez les licences et attribuez les rôles nécessaires.** Confirmez dans le Centre d’administration Microsoft 365 que les utilisateurs concernés disposent d’une licence incluant Entra ID P1 (Business Premium, E3 ou E5 la couvrent déjà). Attribuez ensuite le rôle Administrateur d’accès conditionnel à votre compte de déploiement, plutôt que d’utiliser un compte Administrateur général, pour limiter la surface d’attaque de vos propres identifiants.

**Étape 2 : connectez-vous à Microsoft Graph PowerShell et auditez l’état actuel.** Avant de modifier quoi que ce soit, générez un inventaire complet des méthodes déjà enregistrées par vos utilisateurs. Cela évite de bloquer des comptes qui n’ont encore aucune méthode MFA configurée.

```
Connect-MgGraph -Scopes "Reports.Read.All","Policy.Read.All","UserAuthenticationMethod.Read.All"
Get-MgReportAuthenticationMethodUserRegistrationDetail -All |
    Select-Object UserPrincipalName, IsMfaRegistered, IsPasswordlessCapable, MethodsRegistered |
    Export-Csv -Path .\audit-mfa-2026.csv -NoTypeInformation -Encoding UTF8
```
La commande exporte un fichier CSV listant, pour chaque utilisateur, si le MFA est enregistré, si le compte est compatible sans mot de passe et quelles méthodes sont actives. Voici un extrait représentatif du résultat obtenu sur un tenant de test :

```
UserPrincipalName            IsMfaRegistered  IsPasswordlessCapable  MethodsRegistered
----------------------------  ---------------  ---------------------  --------------------------------
[email protected]           True             False                  {microsoftAuthenticatorPush, sms}
[email protected]            False            False                  {}
[email protected]           True             True                   {fido2SecurityKey}
```
**Étape 3 : créez un compte d’accès d’urgence (break-glass) avant toute modification.** Créez au moins deux comptes cloud-only, exclus de toute politique d’accès conditionnel et de toute exigence MFA, avec un mot de passe long généré aléatoirement et stocké hors ligne, par exemple dans un coffre-fort physique. Ces comptes évitent de se retrouver totalement bloqué hors de son tenant si une politique mal configurée bloque tous les administrateurs en même temps, un scénario plus courant qu’on ne le pense.

**Étape 4 : déployez Microsoft Authenticator auprès des utilisateurs.** Lancez une campagne d’inscription groupée depuis Entra ID (Protection, puis Authentication methods, puis Registration campaign) pour inciter chaque utilisateur à installer Microsoft Authenticator et à enregistrer son compte avant la date limite fixée. Prévoyez une communication interne claire : une bascule surprise vers le MFA obligatoire génère mécaniquement un pic de tickets au support technique.

### Étapes 5 à 9 : basculer vers un MFA qui résiste au phishing

**Étape 5 : activez le number matching et les notifications enrichies.** Dans les paramètres de la méthode Microsoft Authenticator, activez la correspondance de numéro : l’utilisateur doit désormais saisir le chiffre affiché à l’écran de connexion dans l’application, au lieu de simplement appuyer sur « Approuver ». Cette seule modification réduit fortement l’efficacité des attaques par lassitude MFA, car elle empêche l’approbation accidentelle ou automatique.

**Étape 6 : créez une politique d’accès conditionnel exigeant le MFA.** Plutôt que de cliquer manuellement dans le portail, définissez la politique via Microsoft Graph pour pouvoir la reproduire sur plusieurs tenants.

