import React from 'react';
import { Search } from './components/Search';
import { BrowserRouter as Router } from 'react-router-dom';

function App() {
  return (
    <Router>
      <h1>gosear</h1>
      <Search />
    </Router>
  );
}

export default App;
