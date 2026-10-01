---
id: collect-261001-rattrapage/rattrapage/react-guide-11
title: "Guide React — Le manuel complet"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["datacenter"]
source: docs/RAG/collect-261001-rattrapage/react_guide.md
source_anchor: ""
source_lines: [2497, 2729]
sha256: 491a9765d8be61ddb5780368104352d9f34c72e65fe4913f5acc633909c1c722
---

# Guide React — Le manuel complet

```jsx
// components/CarteUps.jsx — carte individuelle mémorisée
import { memo } from 'react';

const CarteUps = memo(function CarteUps({ ups }) {
  const critique = !ups.enLigne || ups.charge > 80;

  return (
    <article className={`carte ${critique ? 'critique' : ''}`}>
      <h2>{ups.nom}</h2>
      <p className="statut">{ups.enLigne ? '🟢 En ligne' : '🔴 Hors ligne'}</p>
      <dl>
        <dt>Charge</dt><dd>{ups.charge} %</dd>
        <dt>Batterie</dt><dd>{ups.batterieV} V</dd>
        <dt>Autonomie est.</dt><dd>{ups.autonomieMin} min</dd>
      </dl>
      {critique && <p className="alerte">⚠️ Intervention requise</p>}
    </article>
  );
});

export default CarteUps;
```

**Ce que ce cas illustre :** polling via hook personnalisé, états loading/error, `useMemo` pour les agrégats, `memo` sur les cartes, `key` stables, rendu conditionnel.

---

## 48. Cas pratique 2 — formulaire de ticket d'intervention

Formulaire complet : un seul objet d'état, validation, envoi POST, retour utilisateur.

```jsx
// pages/NouveauTicket.jsx
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { api } from '../services/api.js';

const INITIAL = {
  titre: '',
  site: '',
  equipement: '',
  priorite: 'normale',
  description: '',
  urgent: false,
};

export default function NouveauTicket() {
  const [form, setForm] = useState(INITIAL);
  const [erreurs, setErreurs] = useState({});
  const [envoi, setEnvoi] = useState(false);
  const [erreurApi, setErreurApi] = useState(null);
  const navigate = useNavigate();

  // Un seul gestionnaire pour tous les champs (name → clé d'état)
  const gererChangement = (e) => {
    const { name, value, type, checked } = e.target;
    setForm((f) => ({ ...f, [name]: type === 'checkbox' ? checked : value }));
  };

  const valider = () => {
    const err = {};
    if (form.titre.trim().length < 5) err.titre = 'Titre trop court (5 caractères min).';
    if (!form.site) err.site = 'Site obligatoire.';
    if (!form.equipement) err.equipement = 'Équipement obligatoire.';
    if (form.description.trim().length < 10) err.description = 'Décrivez le problème (10 caractères min).';
    setErreurs(err);
    return Object.keys(err).length === 0;
  };

  const envoyer = async (e) => {
    e.preventDefault();
    setErreurApi(null);
    if (!valider()) return;
    setEnvoi(true);
    try {
      const ticket = await api.creerTicket(form);
      navigate(`/tickets/${ticket.id}`); // vers le ticket créé
    } catch (err) {
      setErreurApi(err.message);
    } finally {
      setEnvoi(false);
    }
  };

  const Champ = ({ label, ...props }) => (
    <label>
      {label}
      <input {...props} name={props.name} value={form[props.name]} onChange={gererChangement} />
      {erreurs[props.name] && <span className="err">{erreurs[props.name]}</span>}
    </label>
  );

  return (
    <form onSubmit={envoyer} className="form-ticket">
      <h1>Nouveau ticket d'intervention</h1>

      <Champ label="Titre" name="titre" type="text" />
      <Champ label="Site" name="site" type="text" placeholder="Ex. Datacenter Paris-Nord" />
      <Champ label="Équipement" name="equipement" type="text" placeholder="Ex. UPS-01" />

      <label>
        Priorité
        <select name="priorite" value={form.priorite} onChange={gererChangement}>
          <option value="basse">Basse</option>
          <option value="normale">Normale</option>
          <option value="haute">Haute</option>
        </select>
      </label>

      <label>
        Description
        <textarea name="description" value={form.description} onChange={gererChangement} rows={4} />
        {erreurs.description && <span className="err">{erreurs.description}</span>}
      </label>

      <label className="check">
        <input type="checkbox" name="urgent" checked={form.urgent} onChange={gererChangement} />
        Intervention urgente
      </label>

      {erreurApi && <p className="err">❌ Échec de l'envoi : {erreurApi}</p>}

      <button type="submit" disabled={envoi}>
        {envoi ? 'Envoi en cours…' : 'Créer le ticket'}
      </button>
    </form>
  );
}
```

**Points à retenir :** état objet unique + handler générique via `name`, validation avant envoi, bouton désactivé pendant l'envoi (anti double-clic), redirection après succès.

---

## 49. Cas pratique 3 — liste d'équipements filtrable

Recherche + filtre par statut + tri, avec état local bien placé.

```jsx
// pages/Parc.jsx
import { useState, useMemo } from 'react';
import { usePolling } from '../hooks/usePolling.js';
import { Link } from 'react-router-dom';

export default function Parc() {
  const { data: equipements, chargement, erreur } = usePolling('/api/equipements', 15000);
  const [recherche, setRecherche] = useState('');
  const [statut, setStatut] = useState('tous'); // tous | ligne | alarme | hs
  const [tri, setTri] = useState('nom');

  const visibles = useMemo(() => {
    if (!equipements) return [];
    const r = recherche.toLowerCase();
    const filtres = equipements.filter((e) => {
      const okRecherche = e.nom.toLowerCase().includes(r) || e.site.toLowerCase().includes(r);
      const okStatut = statut === 'tous' ||
        (statut === 'ligne' && e.enLigne && e.charge <= 80) ||
        (statut === 'alarme' && e.enLigne && e.charge > 80) ||
        (statut === 'hs' && !e.enLigne);
      return okRecherche && okStatut;
    });
    return [...filtres].sort((a, b) =>
      tri === 'charge' ? b.charge - a.charge : a.nom.localeCompare(b.nom)
    );
  }, [equipements, recherche, statut, tri]);

  if (chargement) return <p>⏳ Chargement du parc…</p>;
  if (erreur) return <p className="erreur">❌ {erreur}</p>;

  return (
    <main>
      <h1>Parc d'équipements</h1>

      <div className="filtres">
        <input
          type="search"
          placeholder="Rechercher (nom, site)…"
          value={recherche}
          onChange={(e) => setRecherche(e.target.value)}
          aria-label="Rechercher un équipement"
        />
        <select value={statut} onChange={(e) => setStatut(e.target.value)} aria-label="Filtrer par statut">
          <option value="tous">Tous statuts</option>
          <option value="ligne">En ligne</option>
          <option value="alarme">En alarme</option>
          <option value="hs">Hors service</option>
        </select>
        <select value={tri} onChange={(e) => setTri(e.target.value)} aria-label="Trier par">
          <option value="nom">Tri : nom</option>
          <option value="charge">Tri : charge ↓</option>
        </select>
      </div>

      <p>{visibles.length} équipement(s) affiché(s)</p>

      {visibles.length === 0 ? (
        <p>Aucun équipement ne correspond aux critères.</p>
      ) : (
        <table>
          <thead>
            <tr><th>Nom</th><th>Site</th><th>Charge</th><th>Statut</th></tr>
          </thead>
          <tbody>
            {visibles.map((e) => (
              <tr key={e.id}>
                <td><Link to={`/equipements/${e.id}`}>{e.nom}</Link></td>
                <td>{e.site}</td>
                <td>{e.charge} %</td>
                <td>{e.enLigne ? (e.charge > 80 ? '⚠️ Alarme' : '🟢 En ligne') : '🔴 HS'}</td>
              </tr>
            ))}
          </tbody>
        </table>
      )}
    </main>
  );
}
```

**Ce que ce cas illustre :** `useMemo` pour filtrer/trier (recalcul uniquement si les critères changent), état des filtres **local** au composant (pas besoin de le remonter), `key` sur les lignes, message "aucun résultat".

---

## 50. Checklist — mettre une app en production

### Avant le build
- [ ] `StrictMode` conservé (aucun warning en console dev)
- [ ] ESLint : zéro erreur (les warnings `react-hooks/*` corrigés)
- [ ] `console.log` de debug supprimés
- [ ] Variables `VITE_*` documentées (`.env.example` fourni)
- [ ] Aucun secret / mot de passe dans le code ou les `.env` front

