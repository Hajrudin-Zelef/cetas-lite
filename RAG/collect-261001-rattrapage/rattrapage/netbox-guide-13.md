---
id: collect-261001-rattrapage/rattrapage/netbox-guide-13
title: "NetBox — Guide complet : source de vérité IPAM / DCIM"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/netbox_guide.md
source_anchor: ""
source_lines: [2521, 2705]
sha256: 9b2b95a1887813151c4c933297a0015dfe597152235495079661f7e4e6243b4c
---

# NetBox — Guide complet : source de vérité IPAM / DCIM

```sql
-- Taille des plus grosses tables (psql, base netbox)
SELECT relname, pg_size_pretty(pg_total_relation_size(relid))
FROM pg_stat_user_tables ORDER BY pg_total_relation_size(relid) DESC LIMIT 10;
```

---

## 79. Erreur n°10 : certificat TLS expiré ou chaîne incomplète

**Symptômes** : les navigateurs affichent « connexion non sécurisée », les scripts API
échouent avec `SSL: CERTIFICATE_VERIFY_FAILED`, l'inventaire Ansible est vide.

```bash
# Vérifier l'expiration (depuis un poste) :
echo | openssl s_client -connect netbox.lan-entreprise.fr:443 -servername \
  netbox.lan-entreprise.fr 2>/dev/null | openssl x509 -noout -dates

# Chaîne complète ? (doit lister les intermédiaires)
openssl s_client -connect netbox.lan-entreprise.fr:443 -showcerts </dev/null 2>/dev/null \
  | grep -c 'BEGIN CERTIFICATE'

# Renouvellement Let's Encrypt :
certbot renew --dry-run   # test
certbot renew              # réel (recharge nginx automatiquement)
```

**Prévention** : sonde Zabbix « expiration certificat » (section 67) avec alerte à
J-21. Pour la PKI interne : notez la date d'expiration dans un champ personnalisé
ou un ticket récurrent — les certificats internes sont la cause n°1 des pannes
« mystères » un vendredi soir.

---

## 80. Erreur n°11 : import CSV partiellement en échec

**Symptômes** : sur 500 lignes importées, 120 rejetées avec des messages d'erreur
parfois cryptiques.

Méthode :

1. **Téléchargez le modèle** : faites d'abord un export CSV d'objets existants pour
   voir les en-têtes et formats exacts (slugs, pas les noms !).
2. Importez par **petits lots** (50 lignes) : l'erreur est plus facile à localiser.
3. Erreurs fréquentes :
   - `site` inexistant → le slug doit matcher EXACTEMENT (`siege-paris11`, pas `Siège`),
   - statut invalide → valeurs anglaises en base (`active`, pas `actif`) selon version,
   - VLAN référencé par VID sans groupe → ambiguïté si plusieurs groupes.
4. Corrigez le CSV, réimportez **uniquement les lignes en échec** (pas tout le fichier,
   sinon doublons → voir erreur n°6).

> 💡 Gardez chaque CSV d'import versionné (git) avec sa date : c'est la preuve de ce
> qui a été injecté, et ça permet de rejouer un import après correction.

---

## 81. Erreur n°12 : « permission denied » sur `media/` après restauration

**Symptômes** : les pièces jointes ne s'affichent plus, l'upload échoue
(`PermissionError: [Errno 13]`), après une restauration ou une montée de version.

```bash
# Le coupable habituel : l'extraction tar en root a écrasé les propriétaires
ls -ld /opt/netbox/netbox/netbox/media
# -> doit appartenir à netbox:netbox

# Réparation :
chown -R netbox:netbox /opt/netbox/netbox/netbox/media
chmod -R u+rwX,go+rX /opt/netbox/netbox/netbox/media
# + durcissement systemd : vérifier ReadWritePaths dans netbox.service (section 14)
systemctl restart netbox netbox-rq
```

**Prévention** : extrayez toujours les archives en tant que `netbox`
(`sudo -iu netbox tar -xzf ...`), ou corrigez les propriétaires juste après.
Ajoutez ce contrôle à votre checklist de restauration (section 66).

---

---

## 82. Bonnes pratiques : convention de nommage (à faire valider en équipe)

Sans convention, NetBox devient un dépotoir en 6 mois. Faites valider ce référentiel
par l'équipe, affichez-le, et ne transigez pas.

| Objet | Format | Exemple |
|---|---|---|
| Site | `TYPE-Ville` (majuscules) | `SIEGE-Paris11`, `AGENCE-Lyon`, `DC-Gravelines` |
| Baie | `SITE-Fonction-NN` | `SIEGE-Baie-A01`, `SIEGE-Brassage-B02` |
| Équipement réseau | `SITE-Rôle-Baie-NN` | `SIEGE-SW-A01-01`, `SIEGE-RTR-01` |
| Serveur physique | `SITE-SRV-Usage-NN` | `SIEGE-SRV-HYP-01` |
| VM | `VM-Usage-NN` | `VM-NETBOX-01`, `VM-FICHIERS-02` |
| Onduleur | `SITE-UPS-Baie` | `SIEGE-UPS-A01` |
| PDU | `SITE-PDU-Baie-Voie` | `SIEGE-PDU-A01-A` |
| Circuit | `CT-Fournisseur-AAAA-NNN` | `CT-OBS-2024-001` |
| VLAN | `Usage` (lisible) + VID | `SRV-Production (110)` |
| Tenant | `MAJUSCULES-tirets` | `DSI`, `CLIENT-DUPONT` |

Règles transverses :

- **Pas d'accents, pas d'espaces, pas de caractères spéciaux** dans les noms techniques
  (les descriptions, elles, peuvent être en français courant).
- **Zéros non significatifs** pour le tri : `SW-01`, `SW-02` … `SW-12` (pas `SW-1`).
- **Immuabilité** : on ne renomme pas un équipement « parce que » — le nom est repris
  dans les scripts, la supervision, les tickets. Renommage = procédure (changer aussi
  DNS, Zabbix, Ansible).

---

## 83. Bonnes pratiques : tenancy systématique

Chaque objet NetBox peut porter un tenant : **utilisez-le toujours**. Le tenancy
répond à trois questions vitales :

1. **« À qui appartient cet équipement ? »** → refacturation, fin de contrat client.
2. **« Que dois-je éteindre si le client X part ? »** → filtre par tenant = liste exacte.
3. **« Qu'est-ce qui est mutualisé ? »** → tenant `INFRA-MUTUALISEE` pour le cœur,
   les onduleurs, les baies partagées.

Matrice de rattachement conseillée :

| Objet | Tenant |
|---|---|
| Baie partagée, onduleur, cœur de réseau | `INFRA-MUTUALISEE` |
| Serveurs du SI interne | `DSI` |
| Équipements dédiés à un client hébergé | `CLIENT-<Nom>` |
| Vidéosurveillance | `SECURITE` |
| Objets de test / bac à sable | `BAC-A-SABLE` |

> 💡 Le tenant `BAC-A-SABLE` est votre soupape : les nouveaux utilisateurs, les tests
> d'import CSV et les scripts en développement s'y déroulent. On y fait des bêtises
> sans conséquence, puis on nettoie (filtre par tenant → suppression en masse).

---

## 84. Règle d'or : « documenter AVANT de câbler »

C'est LA règle qui fait la différence entre un NetBox utile et un NetBox décoratif :

> **Aucune intervention physique (brassage, ajout/retrait d'équipement, déplacement)
> sans fiche NetBox créée ou mise à jour AVANT de toucher au matériel.**

Déclinaison opérationnelle :

1. **Préparation** : créez l'équipement, sa position en baie, ses interfaces, les
   câbles prévus (statut `Planifié` si votre workflow le prévoit).
2. **Intervention** : le technicien travaille AVEC la fiche NetBox ouverte
   (ou le QR code de la baie, plugin section 63).
3. **Clôture** : vérification (positions, étiquettes de câbles, photos), passage des
   statuts à `Actif`, clôture du ticket avec lien vers la fiche.

Pour l'imposer : intégrez-la aux **procédures d'intervention** et aux **revues d'équipe**
(« la baie B02 a été recâblée : les fiches sont à jour ? »). Un chef de service qui
montre l'exemple sur ses propres interventions fait plus que dix notes de service.

---

## 85. Hygiène des données : revues périodiques

NetBox se périme comme tout document : planifiez des revues.

| Revue | Fréquence | Action |
|---|---|---|
| Objets `Planifié`/`Réservé` vieux de > 90 j | mensuelle | activer ou supprimer |
| Équipements `En stock` depuis > 6 mois | trimestrielle | réaffecter ou sortir |
| IP `Réservées` jamais assignées | trimestrielle | libérer |
| Tags temporaires (`a-verifier`…) | mensuelle | traiter puis supprimer le tag |
| Comptes utilisateurs inactifs | trimestrielle | désactiver |
| Champs personnalisés inutilisés | annuelle | supprimer |
| Sites/baies `Déprécié` | annuelle | archiver (export) puis supprimer |

Confiez chaque revue à un **référent** (tournante dans l'équipe) : la qualité des
données est une responsabilité partagée, pas « le problème de l'admin NetBox ».

---

## 86. Cas pratique n°1 : documenter un site complet (fil rouge)

**Contexte** : nouveau site `AGENCE-Nantes` (15 personnes). À documenter : 1 baie 42U,
1 onduleur 3 kVA, 2 PDU, 2 switchs d'accès + 1 routeur, 2 serveurs, 1 NAS, câblage,
adressage `10.30.0.0/16`, VLAN.

### Étape 1 — Arborescence géographique et tenants

