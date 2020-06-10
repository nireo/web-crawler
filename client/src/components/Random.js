import React, { useState, useEffect, useCallback } from 'react';
import axios from 'axios';

export const Random = () => {
  const [results, setResults] = useState(null);
  const loadData = useCallback(async () => {
    const data = await axios.get('https://localhost:3001/random');
    setResults(data.data);
  }, []);

  useEffect(() => {
    if (results === null) {
      loadData();
    }
  }, []);

  console.log(results);

  return (
    <div>
      <h2>Random results</h2>
      <p>Does this not work?</p>
    </div>
  );
};
