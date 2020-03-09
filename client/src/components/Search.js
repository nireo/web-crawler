import React, { useState } from 'react';
import Container from '@material-ui/core/Container';

export const Search = () => {
  const [search, setSearch] = useState('');

  return (
    <Container maxWidth="md">
      <div className="search-bar">
        <input
          onChange={({ target }) => setSearch(target.value)}
          value={search}
          style={{
            fontSize: '16px',
            width: '100%',
            border: 'none',
            marginLeft: '1rem',
            marginRight: '1rem'
          }}
        />
      </div>
    </Container>
  );
};
