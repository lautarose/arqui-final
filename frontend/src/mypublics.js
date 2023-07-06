import React, { useEffect, useState } from 'react';
import Cookies from "universal-cookie";
import { Card, CardContent, Typography, IconButton } from '@mui/material';
import DeleteIcon from '@mui/icons-material/Delete';

const Cookie = new Cookies();

const Publics = () => {
  const [token, setToken] = useState("");
  const [userId, setUserId] = useState("");
  const [userName, setUserName] = useState("");
  const [publications, setPublications] = useState([]);
  const [publicationDetails, setPublicationDetails] = useState([]);

  useEffect(() => {
    loadTokenFromCookie();
  }, []);

  const loadTokenFromCookie = () => {
    const userToken = Cookie.get('user');
    setToken(userToken);
  };

  const fetchPublications = () => {
    fetch(`http://localhost:8090/items/user/${userId}`)
      .then(response => response.json())
      .then(data => {
        setPublications(data);
      })
      .catch(error => {
        console.error('Error:', error);
      });
  };

  const deletePublication = (publicationId) => {
    fetch(`http://localhost:8090/items/${publicationId}`, {
      method: 'DELETE',
      headers: {
        'Authorization': 'Bearer ' + token
      }
    })
      .then(response => {
        if (response.ok) {
          fetchPublications();
          window.location.replace("/mypublics")
        } else {
          console.error('Error deleting publication:', response.statusText);
        }
      })
      .catch(error => {
        console.error('Error:', error);
      });
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
        setUserName(data.name);
      })
      .catch(error => {
        console.error('Error:', error);
      });
  }, [token]);

  useEffect(() => {
    fetchPublications();
  }, [userId, token]);

  useEffect(() => {
    const fetchPublicationDetails = async () => {
      const publicationDetailsPromises = publications.map(publicationId =>
        fetch(`http://localhost:8090/items/${publicationId}`)
          .then(response => response.json())
      );

      const publicationDetailsData = await Promise.all(publicationDetailsPromises);
      setPublicationDetails(publicationDetailsData);
    };

    if (publications.length > 0) {
      fetchPublicationDetails();
    }
  }, [publications]);

  return (
    <div>
      <div style={{ display: 'flex', justifyContent: 'center' }}>
        <h1>Publicaciones de {userName}</h1>
      </div>

      <div>
        {publicationDetails.map(publication => (
          <Card key={publication.id} style={{ margin: '10px' }}>
            <CardContent>
              <Typography variant="h5" component="div">
                {publication.title}
              </Typography>
              <Typography color="text.secondary">
                Seller: {publication.seller}
              </Typography>
              <Typography color="text.secondary">
                Price: {publication.price} {publication.currency}
              </Typography>
              <IconButton onClick={() => deletePublication(publication.id)}>
                <DeleteIcon />
              </IconButton>
            </CardContent>
          </Card>
        ))}
      </div>
    </div>
  );
}

export default Publics;
