import React, { useState } from 'react';
import { TextField, Button } from '@mui/material';
import Swal from 'sweetalert2';

const SignupForm = () => {
  const [name, setName] = useState('');
  const [lastName, setLastName] = useState('');
  const [email, setEmail] = useState('');
  const [userName, setUserName] = useState('');
  const [password, setPassword] = useState('');

  const handleFormSubmit = (e) => {
    e.preventDefault();

    const userData = {
      name,
      last_name: lastName,
      email,
      user_name: userName,
      password,
    };

    // Aquí puedes hacer la llamada POST a tu backend
    // utilizando fetch, axios u otra librería de tu elección

    fetch('http://localhost:8081/user/signup', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(userData),
    })
      .then((response) => {
        if (response.ok) {
          Swal.fire({
            title: '¡Registro exitoso!',
            text: 'Usuario registrado correctamente.',
            icon: 'success'
          }).then(() => {
            // Redirigir al home después de hacer clic en el botón "OK"
            window.location.replace("/");
          });
        } else {
          // Error en el registro
          Swal.fire('Error', 'No se pudo registrar el usuario.', 'error');
        }
      })
      .catch((error) => {
        console.error('Error:', error);
        Swal.fire('Error', 'Hubo un problema al realizar el registro.', 'error');
      });
  };

  return (
    <form onSubmit={handleFormSubmit}>
      <TextField
        label="Nombre"
        value={name}
        onChange={(e) => setName(e.target.value)}
        required
      />
      <TextField
        label="Apellido"
        value={lastName}
        onChange={(e) => setLastName(e.target.value)}
        required
      />
      <TextField
        label="Email"
        type="email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        required
      />
      <TextField
        label="Nombre de usuario"
        value={userName}
        onChange={(e) => setUserName(e.target.value)}
        required
      />
      <TextField
        label="Contraseña"
        type="password"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
        required
      />
      <Button type="submit" variant="contained" color="primary">
        Registrarse
      </Button>
    </form>
  );
};

export default SignupForm;
