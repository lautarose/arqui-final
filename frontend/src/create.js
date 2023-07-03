import React, { useState, useEffect } from 'react';
import { Box, Button, Container, TextField, Typography } from '@mui/material';
import Cookies from "universal-cookie";
import swal from "sweetalert2";
const Cookie = new Cookies();

const Create = () => {
  const [token, setToken] = useState("");
  const [title, setTitle] = useState("");
  const [seller, setSeller] = useState("");
  const [price, setPrice] = useState(0);
  const [currency, setCurrency] = useState("");
  const [picture, setPicture] = useState("");
  const [description, setDescription] = useState("");
  const [state, setState] = useState("");
  const [city, setCity] = useState("");
  const [street, setStreet] = useState("");
  const [number, setNumber] = useState(0);

  useEffect(() => {
    loadTokenFromCookie();
  }, []);

  const loadTokenFromCookie = () => {
    const userToken = Cookie.get("user");
    setToken(userToken);
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    const requestOptions = {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`
      },
      body: JSON.stringify([
        {
          title,
          seller,
          price,
          currency,
          picture,
          description,
          state,
          city,
          street,
          number
        }
      ])
    };

    try {
      const response = await fetch('http://localhost:8090/items/load', requestOptions);
      if (response.status === 201) {
        swal.fire({
          text: "Publicación creada exitosamente",
          icon: 'success'
        }).then(() => {
          setTitle("");
          setSeller("");
          setPrice(0);
          setCurrency("");
          setPicture("");
          setDescription("");
          setState("");
          setCity("");
          setStreet("");
          setNumber(0);
        });
      } else {
        swal.fire({
          text: "Error al crear la publicación",
          icon: 'error'
        });
      }
    } catch (error) {
      swal.fire({
        text: "Error en la solicitud",
        icon: 'error'
      });
    }
  };

  return (
    <Container maxWidth="sm">
      <Box
        sx={{
          marginTop: 4,
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
        }}
      >
        <Typography component="h2" variant="h5">
          Crear Publicación
        </Typography>
        <Box component="form" sx={{ mt: 1 }}>

          <TextField label="Título" type="text" value={title} onChange={(event) => setTitle(event.target.value)} variant="outlined" margin="normal" required fullWidth />
          <TextField label="Seller" type="text" value={seller} onChange={(event) => setSeller(event.target.value)} variant="outlined" margin="normal" required fullWidth />
          <TextField label="Price" type="number" value={price} onChange={(event) => setPrice(event.target.value)} variant="outlined" margin="normal" required fullWidth />
          <TextField label="Currency" type="text" value={currency} onChange={(event) => setCurrency(event.target.value)} variant="outlined" margin="normal" required fullWidth />
          <TextField label="Picture (URL)" type="url" value={picture} onChange={(event) => setPicture(event.target.value)} variant="outlined" margin="normal" required fullWidth />
          <TextField label="Description" type="text" value={description} onChange={(event) => setDescription(event.target.value)} variant="outlined" margin="normal" required fullWidth />
          <TextField label="State" type="text" value={state} onChange={(event) => setState(event.target.value)} variant="outlined" margin="normal" required fullWidth />
          <TextField label="City" type="text" value={city} onChange={(event) => setCity(event.target.value)} variant="outlined" margin="normal" required fullWidth />
          <TextField label="Streer" type="text" value={street} onChange={(event) => setStreet(event.target.value)} variant="outlined" margin="normal" required fullWidth />
          <TextField label="Number" type="number" value={number} onChange={(event) => setNumber(event.target.value)} variant="outlined" margin="normal" required fullWidth />
   
          <Button type="button" fullWidth variant="contained" sx={{ mt: 3, mb: 2 }} onClick={handleSubmit}>
            Crear Publicación
          </Button>

        </Box>
      </Box>
    </Container>
  );
};

export default Create;
