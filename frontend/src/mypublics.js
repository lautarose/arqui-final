import React, { useEffect, useState } from 'react';
import Cookies from "universal-cookie";
const Cookie = new Cookies();

const Publics = () => {
  const [token, setToken] = useState("");
  const [userId, setUserId] = useState("");

  useEffect(() => {
    loadTokenFromCookie();
  }, []);

  const loadTokenFromCookie = () => {
    const userToken = Cookie.get('user');
    setToken(userToken);
  };

  useEffect(() => {
    fetch('http://localhost:8081/user', {
      headers: {
        'Authorization': 'Bearer ' + token
      }
    })
      .then(response => response.json())
      .then(data => {
        setUserId(data.id);
      })
      .catch(error => {
        console.error('Error:', error);
      });
  }, [token]);

    return (
      <div>
        <h1>User ID: {userId}</h1>
      </div>
    );
}

export default Publics;
