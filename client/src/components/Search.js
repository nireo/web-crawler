import React, { useState } from 'react';
import Container from '@material-ui/core/Container';
import Button from '@material-ui/core/Button';
import { useLocation } from 'react-router-dom';

const useQuery = () => {
  return new URLSearchParams(useLocation().search);
};

export const Search = () => {
  const [search, setSearch] = useState('');
  let query = useQuery();

  const handleSearch = (event) => {
    event.preventDefault();
    if (search === '') {
      return;
    }

    query.set('query', search);
  };

  return (
    <Container maxWidth="md">
      <form onSubmit={handleSearch}>
        <div className="search-bar">
          <input
            onChange={({ target }) => setSearch(target.value)}
            value={search}
            style={{
              fontSize: '16px',
              border: 'none',
              marginLeft: '1rem',
              marginRight: '1rem',
              width: '100%',
            }}
            className="search-input"
          />
        </div>
        <Button type="submit" style={{ marginTop: '2rem' }} variant="contained">
          Search
        </Button>
      </form>
    </Container>
  );
};
