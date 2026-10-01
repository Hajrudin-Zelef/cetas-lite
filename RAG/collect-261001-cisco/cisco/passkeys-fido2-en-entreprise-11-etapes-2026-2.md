---
id: collect-261001-cisco/cisco/passkeys-fido2-en-entreprise-11-etapes-2026-2
title: "Activer l'action requise WebAuthn Passwordless"
domain: cisco
role: reference
task: reference
actors: ["Apple", "Google", "Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/passkeys-fido2-en-entreprise-11-etapes-2026.md
source_anchor: ""
source_lines: [50, 148]
sha256: d09c9913d9bfb225538854492b4ee3571166e0f9894fd37485c2613cec193173
---

# Activer l'action requise WebAuthn Passwordless

Si vous utilisez un IdP auto-hébergé comme Keycloak, l’activation de WebAuthn passe par les écrans *Realm Settings → Authentication*. Voici la configuration de base pour activer les passkeys en tant que méthode d’authentification par défaut, avec un OTP en repli.

```
# Activer l'action requise WebAuthn Passwordless
# Realm Settings → Authentication → Required Actions
# Cocher : "Webauthn Register Passwordless" → Enabled + Default Action
# Créer un flux d'authentification dédié
# Authentication → Flows → Créer un flux "Browser - Passkeys"
kcadm.sh create authentication/flows -r monentreprise \
  -s alias="browser-passkeys" \
  -s providerId="basic-flow" \
  -s topLevel=true \
  -s builtIn=false
# Ajouter l'étape WebAuthn Passwordless comme exécution ALTERNATIVE
kcadm.sh create authentication/executions -r monentreprise \
  -s provider="webauthn-register-passwordless" \
  -s parentFlow="browser-passkeys" \
  -s requirement="ALTERNATIVE"
```
Une fois le flux créé, associez-le au navigateur dans *Authentication → Bindings*. Les nouveaux utilisateurs seront invités à enregistrer une passkey dès leur première connexion, avec un QR code d’enrôlement pour les mobiles.

## Étape 4 : Configurer Okta ou Microsoft Entra ID pour les passkeys

Depuis le 1er août 2026, l’authentificateur FIDO2 (WebAuthn) d’Okta a été rebaptisé « Passkey (FIDO2 WebAuthn) » dans les notes de version de l’Okta Identity Engine, avec de nouveaux contrôles d’administration et une expérience utilisateur simplifiée. Concrètement, activer les passkeys dans Okta se fait via l’API Admin :

```
curl -X POST "https://VOTRE_DOMAINE.okta.com/api/v1/policies" \
  -H "Authorization: SSWS ${OKTA_API_TOKEN}" \
  -H "Content-Type: application/json" \
  -d '{
    "type": "ACCESS_POLICY",
    "name": "Politique Passkeys Tier 1",
    "description": "Passkeys obligatoires pour les comptes admin cloud",
    "conditions": {
      "people": {
        "groups": { "include": ["00g1adminsxyz"] }
      }
    },
    "settings": {
      "factors": {
        "webauthn": { "enroll": { "self": "REQUIRED" } }
      }
    }
  }'
```
Pour les organisations sous Microsoft 365, l’équivalent se configure dans Entra ID via *Sécurité → Méthodes d’authentification → Clé d’accès (FIDO2)*, avec la possibilité de restreindre les fournisseurs (AAGUID) autorisés, une option utile si vous imposez des YubiKey spécifiques sur les comptes à privilèges. Microsoft a d’ailleurs annoncé, dans un point d’étape destiné aux MSP publié en juillet 2026, que les passkeys deviendraient la méthode d’authentification par défaut de Microsoft Entra à compter du 1er septembre 2026, reléguant de fait le mot de passe classique au rang d’option secondaire pour les nouveaux comptes. Autre évolution à noter : Google a commencé le 13 juillet 2026 à déployer les clés de sécurité FIDO2 comme second facteur pour la connexion Windows chez tous les clients Google Workspace, un signe que l’authentification matérielle gagne aussi les postes Windows dans les environnements Google.

## Étape 5 : Provisionner les passkeys via le navigateur d’entreprise

Depuis juillet 2026, les navigateurs d’entreprise (Chrome Enterprise, Edge for Business) exposent des API d’administration qui permettent à l’IT d’enregistrer des passkeys corporate directement sur un compte d’annuaire, d’importer des clés exportées sous enveloppe sécurisée par l’éditeur, ou de déclencher des flux de récupération liés à l’identité de l’entreprise. C’est un changement important : jusqu’ici, la gestion des passkeys reposait presque entièrement sur l’utilisateur final.

Voici la politique Chrome Enterprise à pousser via GPO ou votre MDM pour forcer l’enregistrement des passkeys sur les comptes gérés :

```
{
  "PasskeysEnrollmentEnabled": true,
  "PasskeysSyncEnabled": false,
  "WebAuthenticationRemoteProxiedRequestsAllowed": [
    "https://sso.monentreprise.fr"
  ],
  "SecurityKeyPermitAttestation": [
    "https://sso.monentreprise.fr"
  ]
}
```
Le paramètre `PasskeysSyncEnabled: false` est volontaire sur les postes Tier 1 : il empêche la synchronisation cloud automatique et force l’usage d’une clé matérielle, cohérent avec la matrice de risque définie à l’étape 1.

## Étape 6 : Enregistrer une YubiKey 5.8 pas à pas

Pour les comptes Tier 1 et Tier 2, la clé matérielle reste la référence. Le guide officiel Yubico en français détaille la procédure ; voici la version condensée pour un déploiement en entreprise :

1. Connectez-vous au compte de l’utilisateur et accédez aux paramètres de sécurité du service (Okta, Entra ID, Google Workspace).
2. Recherchez l’option « Clé de sécurité » plutôt que « Créer une passkey » générique : Yubico souligne que si vous ne sélectionnez pas explicitement cette option, le service peut créer une passkey synchronisée non protégée par la YubiKey.
3. Insérez la YubiKey 5.8 dans un port USB-C ou approchez-la du lecteur NFC sur mobile.
4. Touchez le capteur doré de la clé lorsque le voyant clignote pour confirmer la présence physique.
5. Définissez un code PIN FIDO2 (6 chiffres minimum) si le service le demande : c’est ce PIN, combiné à la possession physique de la clé, qui constitue l’authentification à deux facteurs.
6. Répétez l’opération avec une seconde clé de secours et stockez-la dans un coffre sécurisé de l’entreprise.

La YubiKey 5.8, lancée le 24 juillet 2026, ajoute l’autorisation matérielle aux passkeys : elle permet désormais de signer des transactions ou des autorisations sensibles directement avec la clé, au-delà de la simple authentification. Le firmware intègre aussi le support de CTAP 2.3, présenté par Yubico comme la base des « passkeys de nouvelle génération », avec des capacités étendues de gestion multi-comptes sur une seule clé.

## Étape 7 : Tester l’interopérabilité multi-navigateurs et multi-appareils

Avant tout déploiement à grande échelle, les recommandations 2026 sont unanimes : testez un pilote d’interopérabilité incluant au moins deux navigateurs, votre IdP, et votre plateforme MDM, sur des appareils BYOD et professionnels. Voici la check-list minimale à valider :

| Scénario testé | Navigateur / OS | Résultat attendu | 
|---|---|---|
| Enrôlement passkey synchronisée | Chrome sur Windows 11 | Passkey stockée dans le compte Google, réutilisable sur Android | 
| Enrôlement passkey synchronisée | Safari sur macOS | Passkey stockée dans le trousseau iCloud, réutilisable sur iPhone | 
| Connexion cross-device via QR code | Edge sur Windows + Android | Scan du QR, confirmation Bluetooth, connexion réussie | 
| Clé matérielle YubiKey | Firefox + tout OS | Détection USB/NFC, PIN demandé, connexion réussie | 
| Poste BYOD non managé | Chrome personnel | Passkey refusée ou repli MFA selon la politique Tier | 
| Perte de connexion réseau pendant l’enrôlement | Tout navigateur | Message d’erreur clair, reprise possible sans doublon de compte | 

Testez également, et c’est souvent oublié, le parcours de révocation : un administrateur doit pouvoir désactiver une passkey compromise en moins de deux minutes, sans devoir contacter le support de l’éditeur du navigateur.

## Étape 8 : Mettre en place le repli et la récupération de compte

C’est le point de friction numéro un des déploiements passkeys en entreprise : que se passe-t-il quand un utilisateur perd son téléphone ou change d’ordinateur sans avoir migré sa passkey ? Trois options doivent être combinées :

