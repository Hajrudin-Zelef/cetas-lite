---
id: collect-261001-rattrapage/rattrapage/netbox-guide-7
title: "NetBox — Guide complet : source de vérité IPAM / DCIM"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/netbox_guide.md
source_anchor: ""
source_lines: [1263, 1463]
sha256: b22883e8a2b14018262026692bcf7f98b4905b5277ab715af71cefb43674eb5f
---

# NetBox — Guide complet : source de vérité IPAM / DCIM

> 💡 En audit de sécurité, l'export des services par site donne la cartographie
> des surfaces d'exposition sans scanner le réseau. Pensez à le maintenir à jour
> lors de chaque mise en production.

---

## 40. Recherche globale, filtres et vues personnalisées

- **Recherche globale** (barre en haut) : tapez `10.10.10.15`, `SIEGE-SW`, `FOC1234`
  (n° de série) : NetBox cherche dans tous les objets. C'est le réflexe n°1.
- **Filtres** : chaque liste (préfixes, équipements…) a un panneau de filtres
  (statut, site, tenant, rôle…). Combinez-les : « équipements Actif + site Siège +
  rôle Switch-Acces ».
- **Colonnes** : personnalisez les colonnes affichées par liste (roue crantée).
- **Signets** : enregistrez vos filtres fréquents comme signets d'équipe
  (« Baies pleines à +80 % », « IP sans DNS »).

Astuce : les filtres sont encodés dans l'URL : copiez l'URL d'une vue filtrée dans
votre documentation d'exploitation ou vos tickets — le destinataire voit exactement
la même sélection.

---

## 41. Exports CSV natifs

Chaque liste propose **Exporter > CSV** : parfait pour les inventaires ponctuels,
les audits, les tableaux de reporting.

Limites à connaître :

- L'export suit les filtres actifs : filtrez d'abord, exportez ensuite.
- Gros volumes (> 50 000 lignes) : préférez l'API paginée (section 54).
- L'export CSV n'inclut pas les pièces jointes ni les champs personnalisés complexes :
  pour un vrai reporting, utilisez les rapports personnalisés (section 59).

---

## 42. Comptes utilisateurs et groupes

**Admin > Utilisateurs** : créez UN compte nominatif par personne. Jamais de compte
partagé « admin » ou « tech » : le journal d'audit (section 45) perd tout son sens
si on ne sait pas QUI a modifié quoi.

**Groupes** (Admin > Groupes) : modélisez vos rôles d'équipe :

| Groupe | Périmètre type |
|---|---|
| `Admins-Infra` | tous droits |
| `Exploitants-Resau` | lecture/écriture sur DCIM+IPAM, pas d'admin |
| `Techniciens-Terrain` | lecture globale + modification des câbles et baies |
| `Lecture-Seule` | consultation (management, auditeurs, prestataires) |
| `Stagiaires` | lecture seule + bac à sable (un tenant dédié) |

> 💡 Le compte `admin` initial : réservez-le aux urgences (mot de passe sous scellés
> dans le coffre). Au quotidien, chacun utilise son compte nominatif, y compris vous.

---

## 43. Permissions : le modèle objet de NetBox

NetBox applique des permissions fines : **qui peut faire quoi sur quels objets**.
Admin > Permissions > + Ajouter :

1. **Actions** : voir, ajouter, modifier, supprimer (par modèle : préfixe, équipement…).
2. **Contraintes** (optionnel mais puissant) : limite l'action à un sous-ensemble,
   exprimé en JSON. Exemple : les techniciens terrain ne modifient que les objets
   du site de Lyon :

```json
{
  "site": ["AGENCE-Lyon"]
}
```

3. Assignez la permission au **groupe** (jamais directement à l'utilisateur :
   ingérable à terme).

Exemple complet : groupe `Techniciens-Terrain` :

| Permission | Contrainte |
|---|---|
| Voir : tout (ou presque) | — |
| Modifier : `dcim.cable`, `dcim.device` (champs position/statut) | `{"site": ["AGENCE-Lyon"]}` |
| Ajouter : `dcim.cable` | `{"site": ["AGENCE-Lyon"]}` |

> ⚠️ Testez TOUJOURS une permission avec un compte du groupe concerné avant de la
> déployer : une contrainte mal écrite bloque silencieusement (l'utilisateur voit
> des listes vides et ouvre un ticket « NetBox est cassé »).

---

## 44. Authentification externe : LDAP / SSO (OIDC)

En entreprise, on ne gère pas les mots de passe dans chaque appli : on délègue à
l'annuaire. NetBox 4.x supporte LDAP (via `django-auth-ldap`) et l'OIDC
(fournisseurs SSO) via configuration.

**LDAP** — ajoutez dans `configuration.py` (paquet `django-auth-ldap` à installer
dans le venv) :

```python
import ldap
from django_auth_ldap.config import LDAPSearch, GroupOfNamesType

AUTH_LDAP_SERVER_URI = "ldaps://ad.lan-entreprise.fr:636"
AUTH_LDAP_BIND_DN = "CN=svc-netbox,OU=Comptes de service,DC=lan-entreprise,DC=fr"
AUTH_LDAP_BIND_PASSWORD = "CHANGEZ_MOI_LDAP_Exemple_2026!"
AUTH_LDAP_USER_SEARCH = LDAPSearch(
    "OU=Utilisateurs,DC=lan-entreprise,DC=fr",
    ldap.SCOPE_SUBTREE, "(sAMAccountName=%(user)s)",
)
AUTH_LDAP_GROUP_SEARCH = LDAPSearch(
    "OU=Groupes,DC=lan-entreprise,DC=fr",
    ldap.SCOPE_SUBTREE, "(objectClass=group)",
)
AUTH_LDAP_GROUP_TYPE = GroupOfNamesType(name_attr="cn")
AUTH_LDAP_USER_ATTR_MAP = {
    "first_name": "givenName",
    "last_name": "sn",
    "email": "mail",
}
# Miroir des groupes AD -> groupes NetBox (pratique !)
AUTH_LDAP_MIRROR_GROUPS = True
```

Avec `AUTH_LDAP_MIRROR_GROUPS = True`, l'appartenance aux groupes AD
(`GG-NetBox-Admins`…) pilote les groupes NetBox : **l'habilitation se gère dans
l'AD**, NetBox suit automatiquement. C'est le montage recommandé.

---

## 45. Journal d'audit (changelog) et traçabilité

Chaque création, modification, suppression est enregistrée : **qui, quand, quoi,
avant/après**. Consultez :

- le journal global : en bas de la page d'accueil ou via l'objet « Journal »,
- le journal d'un objet : onglet **Journal** sur sa fiche.

Cas d'usage :

- « Qui a déplacé le serveur SRV-02 de la baie A01 vers A02 ? » → réponse en 30 s.
- « Quand l'IP 10.10.10.15 a-t-elle été assignée ? » → horodatage exact.
- Conflit d'équipe : l'historique tranche sans débat.

**Rétention** : le changelog grossit avec l'usage. Paramètre dans `configuration.py` :

```python
CHANGELOG_RETENTION = 90   # jours ; 0 = conservation infinie (défaut)
```

> ⚠️ En environnement audité (ISO 27001…), conservez le changelog longtemps et
> incluez la table `extras_objectchange` dans votre politique de sauvegarde :
> c'est elle qui porte la traçabilité (la sauvegarde BDD standard l'inclut déjà).

---

## 46. Sécurité : durcissement de l'installation

Checklist de durcissement (à appliquer dès la mise en production) :

- [ ] `configuration.py` en `600`, propriétaire `netbox:netbox`
- [ ] `LOGIN_REQUIRED = True`, `SECRET_KEY` unique et robuste
- [ ] HTTPS partout, cookies `Secure`, HSTS (section 16)
- [ ] Comptes nominatifs, groupes + permissions (sections 42–43), compte `admin`
    sous scellés
- [ ] LDAP/SSO si disponible (section 44) ; sinon politique de mots de passe forte
- [ ] PostgreSQL : écoute `localhost` uniquement, mot de passe robuste
- [ ] Redis : `bind 127.0.0.1`, `requirepass` si exposition réseau inévitable
- [ ] Pare-feu (nftables/UFW) : n'ouvrir que 443 (et 22 restreint aux admins)
- [ ] `DEBUG = False` (jamais `True` en production : fuite d'informations garantie)
- [ ] Mises à jour OS + NetBox suivies (section 66)
- [ ] Sauvegardes testées (section 65) — une sauvegarde non testée n'existe pas
- [ ] Bannière `BANNER_TOP` pour marquer l'environnement (PROD vs LAB)

```bash
# Exemple UFW minimaliste
ufw default deny incoming
ufw allow from 10.10.0.0/16 to any port 443 proto tcp
ufw allow from 10.10.254.10 to any port 22 proto tcp   # bastion uniquement
ufw enable
```

---

## 47. Champs personnalisés (custom fields)

Les **champs personnalisés** (Admin > Champs personnalisés) étendent les objets
NetBox sans plugin : texte, entier, booléen, date, URL, JSON, sélection.

Exemples utiles pour un service systèmes :

| Objet | Champ | Type | Usage |
|---|---|---|---|
| Équipement | `date_fin_garantie` | date | alertes de renouvellement |
| Équipement | `contrat_maintenance` | texte | n° de contrat |
| Équipement | `criticite` | sélection (`Critique/Majeur/Mineur`) | priorisation |
| Baie | `disjoncteur_amont` | texte | `TGBT-DJ-14 — 32A` (lien énergies !) |
| Préfixe | `responsable` | texte | référent métier |
| Circuit | `date_engagement` | date | fin d'engagement opérateur |

