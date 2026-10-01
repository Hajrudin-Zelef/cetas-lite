---
id: collect-261001-cisco/cisco/passkeys-fido2-en-entreprise-11-etapes-2026-4
title: "Activer l'action requise WebAuthn Passwordless"
domain: cisco
role: reference
task: reference
actors: ["Google", "Microsoft"]
dates: []
keywords: []
source: docs/RAG/collect-261001-cisco/passkeys-fido2-en-entreprise-11-etapes-2026.md
source_anchor: ""
source_lines: [230, 322]
sha256: 992e69dc5af35c3f056b2965e9f7a0fa04397028ec1897cc3f774b2500fb91bb
---

# Activer l'action requise WebAuthn Passwordless

**6. Les utilisateurs BYOD ne peuvent pas s’enrôler.** Vérifiez la politique MDM appliquée aux appareils non managés : certaines organisations bloquent volontairement l’enrôlement WebAuthn sur les appareils hors inventaire, ce qui est correct pour Tier 1/2 mais doit être assoupli pour Tier 3/4.

**7. Le SIEM ne remonte aucun événement webauthn_register.** Vérifiez que l’export des logs d’authentification de votre IdP est bien configuré vers votre collecteur (Splunk, Elastic, Sentinel) : ces événements ne sont pas toujours activés par défaut.

**8. Une application interne legacy refuse la connexion après désactivation du mot de passe.** Ne désactivez jamais le mot de passe sur un service tant que sa compatibilité WebAuthn n’a pas été confirmée. Utilisez un proxy d’authentification (reverse proxy SSO) en solution transitoire si l’application ne peut pas être modifiée à court terme.

## Astuces avancées pour les grandes organisations

Pour les entreprises de plus de 1 000 employés ou les environnements multi-filiales, quelques ajustements supplémentaires font gagner un temps considérable :

- **Provisionnement en masse via API.** Plutôt que de laisser chaque utilisateur s’enrôler manuellement, générez des demandes d’enrôlement en lot via l’API de votre IdP et distribuez les clés YubiKey pré-associées via votre outil de gestion d’inventaire matériel.
- **Segmentation par filiale ou pays.** Une entité soumise à des exigences réglementaires locales (banque, santé) peut nécessiter une politique Tier 1 plus stricte que le reste du groupe : gérez ces politiques par groupe d’annuaire plutôt que par exception individuelle.
- **Attestation matérielle vérifiée.** Activez la vérification d’attestation FIDO2 pour vous assurer que seules des clés matérielles certifiées (pas des implémentations logicielles) sont acceptées sur les comptes à privilèges.
- **Rotation planifiée des clés de secours.** Traitez les clés YubiKey de secours comme des secrets à rotation périodique (tous les 18 à 24 mois), en particulier après le départ d’un employé ayant eu accès physique au coffre.
- **Intégration avec la gestion des identités des sous-traitants.** Les comptes prestataires externes doivent suivre un cycle de vie distinct, avec expiration automatique des passkeys à la fin du contrat.

## Comparatif des principales clés de sécurité matérielles compatibles

Pour les comptes Tier 1 et Tier 2, le choix de la clé matérielle influence directement le coût et la facilité de déploiement. Voici un comparatif des options les plus utilisées en entreprise en 2026.

| Clé | Connectique | Support CTAP2.3 | Cas d’usage recommandé | 
|---|---|---|---|
| YubiKey 5.8 | USB-C, USB-A, NFC | Oui | Comptes admin, signature d’autorisation matérielle | 
| Google Titan (génération récente) | USB-C, NFC, Bluetooth | Partiel | Comptes Google Workspace, usage grand volume | 
| Feitian ePass FIDO | USB-A, USB-C | Oui | Déploiements à budget contraint, gros volumes | 

Pour une comparaison plus détaillée de ces trois références, notamment sur le nombre de passkeys stockables et le rapport qualité-prix, consultez notre comparatif YubiKey vs Google Titan vs Feitian.

## Le contexte réglementaire européen et français

En France, l’Agence nationale de la sécurité des systèmes d’information (ANSSI) recommande depuis plusieurs années le renforcement de l’authentification comme mesure prioritaire face à la recrudescence des attaques par hameçonnage ciblé. Les passkeys s’inscrivent directement dans cette logique, et la directive NIS2, transposée progressivement dans le droit français, impose aux opérateurs de services essentiels et importants de justifier la robustesse de leurs mécanismes d’authentification.

Sur le volet protection des données, la CNIL rappelle que les mécanismes d’authentification biométrique associés aux passkeys (empreinte, reconnaissance faciale) doivent rester locaux à l’appareil et ne jamais transiter vers les serveurs de l’entreprise : la norme WebAuthn respecte nativement ce principe, puisque seule la clé publique quitte l’appareil de l’utilisateur.

L’ENISA, dans son panorama des menaces 2025, structure ses recommandations de déploiement de passkeys autour de cinq piliers : intégration transparente, repli sécurisé, autonomisation de l’utilisateur, vigilance opérationnelle continue, et sélection rigoureuse des fournisseurs. Cette grille de lecture recoupe presque exactement les dix étapes détaillées dans ce tutoriel.

## Passkeys et sécurisation des API : le lien avec OAuth2

Les passkeys sécurisent l’authentification humaine, mais elles ne remplacent pas les mécanismes d’autorisation applicative comme OAuth2 et JWT, utilisés pour sécuriser les échanges entre services. Dans une architecture moderne, les deux se combinent : la passkey authentifie l’utilisateur auprès de l’IdP, qui émet ensuite un jeton OAuth2 pour autoriser les appels API en son nom. Pour les équipes qui doivent aussi durcir leurs API internes, notre guide sur comment sécuriser une API REST avec JWT et OAuth2 détaille cette seconde brique, complémentaire à ce tutoriel.

## Vérifier l’exposition de vos comptes avant le déploiement

Avant de lancer un projet de passkeys, il est utile de vérifier si des identifiants de votre organisation ont déjà fuité dans des bases de données compromises : ces comptes doivent être prioritaires dans votre migration Tier 1. Notre tutoriel Have I Been Pwned : vérifier une fuite explique comment automatiser cette vérification à l’échelle de tout un domaine d’entreprise.

## Projet complet : script d’audit de conformité passkeys

Pour clore ce tutoriel avec un livrable réutilisable, voici un script Python simplifié qui interroge l’API de votre IdP pour générer un rapport de conformité passkeys par département, exploitable directement dans votre reporting mensuel de sécurité.

```
import requests
import csv
OKTA_DOMAIN = "https://votre-domaine.okta.com"
API_TOKEN = "VOTRE_TOKEN_API"
HEADERS = {"Authorization": f"SSWS {API_TOKEN}"}
def get_users():
    resp = requests.get(f"{OKTA_DOMAIN}/api/v1/users", headers=HEADERS)
    resp.raise_for_status()
    return resp.json()
def get_factors(user_id):
    resp = requests.get(
        f"{OKTA_DOMAIN}/api/v1/users/{user_id}/factors", headers=HEADERS
    )
    resp.raise_for_status()
    return resp.json()
def build_report():
    rows = []
    for user in get_users():
        user_id = user["id"]
        email = user["profile"]["email"]
        department = user["profile"].get("department", "N/A")
        factors = get_factors(user_id)
        has_passkey = any(f["factorType"] == "webauthn" for f in factors)
        rows.append({
            "email": email,
            "departement": department,
            "passkey_active": has_passkey,
        })
    return rows
def export_csv(rows, filename="rapport_passkeys.csv"):
    with open(filename, "w", newline="", encoding="utf-8") as f:
        writer = csv.DictWriter(f, fieldnames=["email", "departement", "passkey_active"])
        writer.writeheader()
        writer.writerows(rows)
if __name__ == "__main__":
    data = build_report()
    export_csv(data)
    taux = sum(1 for r in data if r["passkey_active"]) / len(data) * 100
    print(f"Taux d'adoption passkeys : {taux:.1f}%")
```
Adaptez les appels API à votre IdP (Entra ID utilise Microsoft Graph, Keycloak son Admin REST API) : la logique reste identique — lister les utilisateurs, vérifier la présence d’un facteur WebAuthn, exporter un taux d’adoption exploitable par la direction.

## Analyse coûts-bénéfices d’un déploiement passkeys

