---
id: collect-261001-general-networking/general-networking/passkeys-fini-le-mot-de-passe-en-13-etapes-2026-3
title: "Installe une autorité de certification locale de confiance"
domain: general-networking
role: reference
task: reference
actors: ["Apple", "Google", "Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/passkeys-fini-le-mot-de-passe-en-13-etapes-2026.md
source_anchor: ""
source_lines: [87, 184]
sha256: a42975f126bd5f6a7e14c8d813917b2d37a3ae7cdc9f933a169e8c4d901adf08
---

# Installe une autorité de certification locale de confiance

Pour ces quatre comptes, répétez la même vérification finale : déconnectez-vous complètement, puis reconnectez-vous en choisissant explicitement l’option passkey plutôt que le mot de passe, pour confirmer que le geste fonctionne réellement de bout en bout et pas seulement au moment de la création. Profitez-en aussi pour créer un deuxième passkey depuis un appareil différent lorsque le service le permet, ce qui vous évite de dépendre d’un unique téléphone pour l’accès à des comptes que vous utilisez quotidiennement.

## Étape 5 : choisir et synchroniser votre gestionnaire de passkeys

Un point bloque souvent les nouveaux utilisateurs : où vivent exactement ces passkeys, et que se passe-t-il en cas de changement d’appareil. Deux familles existent. Les passkeys liés à l’appareil (device-bound) restent physiquement sur une clé de sécurité ou une puce sécurisée et ne se copient jamais. Les passkeys synchronisables se répliquent via un écosystème cloud chiffré, que ce soit celui de Google, d’Apple, de Microsoft ou d’un gestionnaire tiers comme Bitwarden.

Pour un usage personnel multiplateforme, un gestionnaire tiers évite de rester enfermé dans un seul écosystème. Si vous hébergez déjà vos identifiants avec Vaultwarden, la version auto-hébergée de Bitwarden, vous pouvez y stocker vos passkeys au même endroit que vos mots de passe restants, avec le même niveau de contrôle sur vos données. C’est l’option la plus cohérente pour quelqu’un qui a déjà investi dans une infrastructure d’identifiants souveraine.

Pour un déploiement en entreprise, la question se pose différemment. Un fournisseur d’identité comme Okta, Entra ID ou Ping Identity centralise la politique d’authentification et impose ou non les passkeys selon le profil de risque de chaque utilisateur. Nous reviendrons sur cette intégration plus loin, une fois le projet de démonstration construit.

Un débat revient régulièrement dans les équipes sécurité au sujet de la portabilité : que se passe-t-il si un utilisateur quitte l’écosystème Apple pour Android, ou change de gestionnaire de mots de passe ? Le standard WebAuthn ne définit pas de mécanisme universel d’export d’un passkey d’un écosystème vers un autre, chaque fournisseur gère sa propre synchronisation en interne. Dans la pratique, cela signifie que l’utilisateur doit recréer un nouveau passkey dans le nouvel écosystème plutôt que de migrer l’ancien. Ce n’est pas idéal, mais ce n’est pas non plus bloquant tant qu’une méthode de secours reste disponible pendant la transition, un point sur lequel nous revenons dans la section consacrée aux erreurs à éviter.

## Étapes 6 à 8 : construire le backend WebAuthn avec Node.js

Place à la partie développeur. L’objectif : un projet minimal mais complet, qui gère l’inscription et la connexion par passkey de bout en bout. Nous utilisons Express pour le serveur et la bibliothèque SimpleWebAuthn, qui encapsule la complexité de l’encodage CBOR et de la vérification des signatures WebAuthn derrière une API JavaScript lisible.

L’architecture générale tient en quatre échanges entre le navigateur et le serveur. Pour l’inscription : le navigateur demande des options d’enregistrement, le serveur répond avec un challenge et des paramètres, le navigateur transmet ces paramètres à l’authentificateur local puis renvoie l’attestation obtenue, et le serveur vérifie cette attestation avant de sauvegarder le nouveau credential. Pour la connexion, le schéma se répète à l’identique en remplaçant l’attestation par une assertion. Écrire cette logique à la main serait long et risqué sur le plan cryptographique, raison pour laquelle nous nous appuyons sur SimpleWebAuthn plutôt que de réimplémenter l’encodage CBOR ou la vérification de signature nous-mêmes.

### Étape 6 : initialiser le projet

```
mkdir passkey-demo && cd passkey-demo
npm init -y
npm install express express-session @simplewebauthn/server
npm install --save-dev nodemon
```
Créez ensuite un module de stockage minimal. Pour ce tutoriel, un stockage en mémoire suffit à démontrer le flux complet. En production, remplacez-le par une vraie base de données (PostgreSQL, MySQL ou SQLite), en conservant la même structure de champs.

### Étape 7 : le modèle de données

```
// db.js : stockage en mémoire pour la démo
// En production, remplacez cette Map par une vraie table utilisateurs/credentials
const users = new Map();
function getUser(username) {
  if (!users.has(username)) {
    users.set(username, {
      id: username,
      devices: [],           // liste des passkeys enregistrés
      currentChallenge: undefined,
    });
  }
  return users.get(username);
}
module.exports = { getUser };
```
### Étape 8 : générer les options d’enregistrement

Le serveur doit d’abord générer un défi cryptographique aléatoire (challenge) et des paramètres d’inscription, que le navigateur transmettra ensuite à l’authentificateur de l’utilisateur.

```
// server.js
const express = require('express');
const session = require('express-session');
const {
  generateRegistrationOptions,
  verifyRegistrationResponse,
  generateAuthenticationOptions,
  verifyAuthenticationResponse,
} = require('@simplewebauthn/server');
const { getUser } = require('./db');
const app = express();
app.use(express.json());
app.use(session({
  secret: process.env.SESSION_SECRET || 'changez-moi-en-production',
  resave: false,
  saveUninitialized: false,
}));
const rpName = 'Mon Application';
const rpID = 'localhost';           // votre domaine réel en production
const origin = `https://${rpID}:3000`;
app.post('/register/options', async (req, res) => {
  const { username } = req.body;
  const user = getUser(username);
  const options = await generateRegistrationOptions({
    rpName,
    rpID,
    userName: username,
    attestationType: 'none',
    excludeCredentials: user.devices.map(d => ({
      id: d.credentialID,
      transports: d.transports,
    })),
    authenticatorSelection: {
      residentKey: 'required',
      userVerification: 'preferred',
    },
  });
  user.currentChallenge = options.challenge;
  res.json(options);
});
```
Notez le paramètre **attestationType: ‘none’**. Il évite de demander une preuve d’origine matérielle de l’authentificateur, une information rarement utile hors contextes réglementés et qui soulève des questions de confidentialité si elle est collectée sans raison. Le paramètre **residentKey: ‘required’** force la création d’une clé découvrable, ce qui permettra plus tard la connexion sans même saisir de nom d’utilisateur.

## Étapes 9 et 10 : la cérémonie d’enregistrement côté client et sa vérification

Côté navigateur, la bibliothèque **@simplewebauthn/browser** encapsule l’appel à `navigator.credentials.create()`. C’est cet appel qui déclenche la fenêtre native de votre système (empreinte, Face ID, code PIN).

