import React, { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import { AppBar, Toolbar, IconButton, InputBase, Button } from '@mui/material';
import { Search as SearchIcon } from '@mui/icons-material';
import '../css/home.css'
import Cookies from 'universal-cookie';
import { red, blue, grey, green } from '@mui/material/colors';


const Navbar = () => {
  const cookies = new Cookies();
  const userToken = cookies.get('user');
  const navigate = useNavigate();
  const [searchTerm, setSearchTerm] = useState('');
  const handleLogout = () => {
    cookies.remove('user'); // Eliminar la cookie 'user'
    window.location.replace('/'); // Redirigir al home
  };

  const handleSearch = () => {
    if (searchTerm === '') {
      navigate('/');
    } else {
      const formattedSearchTerm = searchTerm.replace(' ', '%20');
      navigate(`/search/${formattedSearchTerm}`);
    }
  };

  const handleInputChange = (e) => {
    setSearchTerm(e.target.value);
  };

  return (

    <AppBar position="static" color="primary">
      <Toolbar>
        <IconButton color="inherit" onClick={handleSearch}>
          <SearchIcon />
        </IconButton>
        <InputBase
          placeholder="Search..."
          value={searchTerm}
          onChange={handleInputChange}
        />
        {userToken ? (
          // Elementos que se mostrarán si el usuario está logueado
          <div>
            <Button onClick={handleLogout} sx={{ color: grey[200], backgroundColor: 'red', border: `2px solid ${blue[500]}`, borderRadius: '4px', padding: '8px 16px', transition: 'transform 0.3s', '&:hover': { color: red[500], borderColor: red[500], transform: 'scale(1.02)' }, position: 'absolute', top: '10px', right: '10px', justifyContent: 'flex-end', alignItems: 'flex-start' }}>
              LOGOUT
            </Button>
            <Link to="/create" >
              <Button sx={{ left: '250px', color: grey[300], transform: 'scale(1)', transition: 'transform 0.3s', '&:hover': { color: blue[500], transform: 'scale(1.1)' } }}>
                CREAR PUBLICACIÓN
              </Button>
            </Link>
            <Link to="/mypublics">
              <Button sx={{ left: '250px',color: grey[300], transform: 'scale(1)', transition: 'transform 0.3s', '&:hover': { color: blue[500], transform: 'scale(1.1)' } }}>
                MIS PUBLICACIONES
              </Button>
            </Link>
          </div>
        ) : (
          // Elementos que se mostrarán si el usuario no está logueado
          <>
          <Link to="/login">
            <Button sx={{  display: 'flex', justifyContent: 'center', alignItems: 'center',width: '200px',color: grey[200], backgroundColor: 'green', border: `2px solid ${blue[500]}`, borderRadius: '4px', padding: '8px 16px', transition: 'transform 0.3s', '&:hover': { backgroundColor: 'green',color: grey[300], borderColor: green[500], transform: 'scale(1.02)' }, position: 'absolute', top: '10px', right: '700px'}}>
              LOGIN
            </Button>
          </Link>
          <Link to="/signup">
            <Button sx={{  display: 'flex', justifyContent: 'center', alignItems: 'center',width: '200px',color: blue[500], backgroundColor: blue[200], border: `2px solid ${blue[500]}`, borderRadius: '4px', padding: '8px 16px', transition: 'transform 0.3s', '&:hover': { backgroundColor: blue[500],color: grey[300], borderColor: grey[300], transform: 'scale(1.02)' }, position: 'absolute', top: '10px', right: '400px'}}>
              SIGNUP
            </Button>
          </Link>
          </>
        )}
      </Toolbar>
    </AppBar>
  );
};

export default Navbar;
