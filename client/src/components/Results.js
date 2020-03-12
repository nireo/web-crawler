import React, { useEffect, useState } from "react";
import axios from "axios";

const Results = ({query}) => {
    const [loaded, setLoaded] = useState(false);
    const [results, setResults]= useState([]);
    useEffect(() => {
        if (loaded === false) {
            axios.get(`/?query=${query}`).then(response => {
                setResults(response.data);
            }).catch(err => {
                console.log(err)
            })
            setLoaded(true);
        }
    }, []);

    if (!query) {
        return null;
    }

    return (
        <div>
        {!loaded && results.length === 0 && (
             <div>loading...</div>
        )}
        {loaded && results.length === 0 && (
             <div>Problem loading results</div>
        )}
        </div>
    );
}
