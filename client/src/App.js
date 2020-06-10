import React from 'react';
import { Search } from './components/Search';
import { BrowserRouter as Router, Route, Link } from 'react-router-dom';
import { Random } from './components/Random';

function App() {
  return (
    <Router>
      <h1>gosear</h1>
      <Link to="/random">random</Link>
      <Route exact path="random" render={() => <Random />} />
      <Route exact path="/" render={() => <Search />} />
    </Router>
  );
}

export default App;
