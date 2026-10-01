---
id: collect-261001-cisco/cisco/passkeys-fido2-en-entreprise-11-etapes-2026-5
title: "Activer l'action requise WebAuthn Passwordless"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/passkeys-fido2-en-entreprise-11-etapes-2026.md
source_anchor: ""
source_lines: [323, 417]
sha256: 2536ebe96be4e5e68e049bfc2ccd7dbd2d2e3ca3d2b442eb107593cc0a794855
---

# Activer l'action requise WebAuthn Passwordless

Avant de présenter le projet à la direction, il est utile de chiffrer précisément l’investissement face aux gains attendus. Le coût principal reste le matériel (clés YubiKey ou équivalent) et le temps d’intégration IT, mais les économies indirectes sur le support et sur la réduction du risque de compromission sont généralement sous-estimées dans les premiers arbitrages budgétaires.

| Poste | Coût / gain estimé | Détail | 
|---|---|---|
| Clés matérielles (300 comptes Tier 1/2, 2 clés chacun) | 15 000 € – 36 000 € | Achat unique, amortissable sur 5 à 7 ans selon la durée de vie de la clé | 
| Intégration IdP et configuration | 40 à 120 heures d’ingénierie | Variable selon la complexité de l’annuaire existant et le nombre d’applications à raccorder | 
| Formation support IT | 2 à 3 sessions d’une demi-journée | Concentré sur les scénarios de perte, de récupération et de révocation | 
| Réduction des tickets « mot de passe oublié » | -60 % à -80 % en moyenne | Constaté sur les comptes migrés, d’après les retours de pilotes d’entreprise en 2026 | 
| Réduction du risque de compromission par hameçonnage | Quasi nulle sur les comptes passkey | Le hameçonnage classique par formulaire de connexion ne fonctionne plus sans le secret partagé | 

Le calcul du retour sur investissement dépend surtout de la taille de votre service desk. Pour une organisation où les réinitialisations de mot de passe représentent une part significative des tickets, la bascule vers les passkeys réduit mécaniquement la charge du support, souvent avec un retour sur investissement observé en moins de 18 mois d’après les retours de terrain publiés en 2026 sur les déploiements en entreprise. Pour donner un ordre de grandeur de la montée en charge à anticiper côté infrastructure, le rapport *Passkey Power 20* de Dashlane, publié fin 2025 et relayé par Help Net Security, recensait déjà 1,3 million d’authentifications par passkey chaque mois dès octobre 2025 sur les seuls services grand public suivis par l’étude : un volume qui donne une idée du trafic que votre IdP doit pouvoir absorber une fois le mot de passe désactivé à grande échelle.

## Feuille de route détaillée sur 12 mois

Pour donner une vision d’ensemble avant de vous lancer, voici la trame de planification type utilisée par les équipes IAM ayant mené un déploiement complet en 2026, du cadrage initial jusqu’à la désactivation du mot de passe sur les derniers services.

| Phase | Durée | Livrables clés | 
|---|---|---|
| Cadrage et matrice de risque | Semaines 1-2 | Inventaire des services, classification Tier 1-4, budget matériel validé | 
| Intégration technique IdP | Semaines 3-6 | Configuration WebAuthn, politiques d’accès, restriction AAGUID sur Tier 1 | 
| Pilote restreint | Semaines 6-10 | 20 à 50 utilisateurs, mesure du taux d’échec et des tickets support | 
| Vague 1 : comptes Tier 1 | Mois 3 | Comptes admin, VPN, finance migrés à 100 % | 
| Vague 2 : comptes Tier 2 | Mois 4-6 | Messagerie, CRM, outils RH migrés par groupe d’annuaire | 
| Vague 3 : comptes Tier 3 | Mois 6-9 | Intranet et outils collaboratifs, adoption majoritairement volontaire | 
| Vague 4 : BYOD et prestataires | Mois 9-12 | Politique dédiée, expiration automatique en fin de contrat | 
| Désactivation du mot de passe | Mois 12+ | Service par service, après validation du taux d’adoption et du plan de repli | 

Cette feuille de route reste indicative : une PME de 50 employés peut compresser ce calendrier à 3 ou 4 mois, tandis qu’un groupe multi-filiales avec des contraintes réglementaires locales (banque, assurance, santé) dépassera souvent les 12 mois annoncés, notamment à cause des cycles de validation de conformité propres à chaque pays européen.

## Passkeys pour les développeurs : intégrer WebAuthn dans une application maison

Si votre entreprise développe des applications internes qui doivent elles-mêmes supporter les passkeys, plutôt que de dépendre uniquement de votre IdP, voici le squelette d’intégration côté serveur avec la bibliothèque `@simplewebauthn/server` pour Node.js, une des implémentations WebAuthn les plus utilisées en 2026.

```
import {
  generateRegistrationOptions,
  verifyRegistrationResponse,
} from '@simplewebauthn/server';
// Étape 1 : générer les options d'enregistrement côté serveur
app.post('/webauthn/register-options', async (req, res) => {
  const user = req.session.user;
  const options = await generateRegistrationOptions({
    rpName: 'MonEntreprise SSO',
    rpID: 'sso.monentreprise.fr',
    userID: Buffer.from(user.id),
    userName: user.email,
    attestationType: 'direct',
    authenticatorSelection: {
      residentKey: 'required',
      userVerification: 'preferred',
    },
  });
  req.session.currentChallenge = options.challenge;
  res.json(options);
});
// Étape 2 : vérifier la réponse envoyée par le navigateur
app.post('/webauthn/register-verify', async (req, res) => {
  const verification = await verifyRegistrationResponse({
    response: req.body,
    expectedChallenge: req.session.currentChallenge,
    expectedOrigin: 'https://sso.monentreprise.fr',
    expectedRPID: 'sso.monentreprise.fr',
  });
  if (verification.verified) {
    await saveCredentialToDatabase(req.session.user.id, verification.registrationInfo);
  }
  res.json({ verified: verification.verified });
});
```
Ce code illustre le principe fondamental de WebAuthn : le serveur génère un défi (challenge) unique, le navigateur le signe avec la clé privée stockée sur l’appareil ou la clé matérielle, et le serveur vérifie la signature sans jamais voir ni stocker la clé privée elle-même. C’est ce mécanisme qui rend le phishing structurellement inefficace contre les passkeys : même un faux site ne peut pas obtenir de signature valide, puisque l’origine (`expectedOrigin`) est vérifiée cryptographiquement.

## Foire aux questions

### Combien de temps prend un déploiement de passkeys en entreprise ?

Comptez généralement entre 6 et 12 mois pour un déploiement complet, de l’intégration technique avec votre fournisseur d’identité jusqu’à l’inscription de l’ensemble des utilisateurs. Le pilote initial dure typiquement 3 à 4 semaines.

### Les passkeys remplacent-elles complètement le mot de passe ?

À terme, oui, sur les services compatibles. En pratique, la plupart des organisations conservent un mécanisme de repli (MFA classique ou mot de passe) pendant toute la phase de transition, et pour les applications legacy qui ne supportent pas encore WebAuthn.

### Passkey synchronisée ou liée à l’appareil : laquelle choisir ?

Pour les comptes à privilèges (admin cloud, VPN, finance), privilégiez une passkey liée à l’appareil sur clé matérielle YubiKey. Pour les comptes standards, une passkey synchronisée offre un meilleur équilibre entre sécurité et confort d’usage. Près de la moitié des entreprises déployées combinent les deux approches selon le niveau de risque du compte.

### Que faire si un employé perd sa clé YubiKey ?

Utilisez la clé de secours pré-enregistrée ou les codes de récupération générés à l’enrôlement. Révoquez immédiatement la clé perdue dans votre IdP pour empêcher toute utilisation frauduleuse si elle est retrouvée par un tiers.

### Les passkeys sont-elles conformes au RGPD ?

Oui. Les données biométriques éventuellement utilisées pour déverrouiller une passkey (empreinte, reconnaissance faciale) restent locales à l’appareil et ne sont jamais transmises à l’entreprise ni au fournisseur de service : seule la clé publique cryptographique est partagée, ce qui limite fortement la surface de données personnelles traitées.

### Peut-on déployer des passkeys sur des appareils BYOD non managés ?

