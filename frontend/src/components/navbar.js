import React, { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import { AppBar, Toolbar, IconButton, InputBase, Button } from '@mui/material';
import { Search as SearchIcon } from '@mui/icons-material';
import '../css/home.css'
import  Cookies from 'universal-cookie';



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
            <Button  color="inherit" onClick={handleLogout}>
              LOGOUT
            </Button>
        ) : (
          // Elementos que se mostrarán si el usuario no está logueado
          <Link to="/login">
            <Button color="inherit">
              LOGIN
            </Button>
          </Link>
        )}
      </Toolbar>
    </AppBar>
  );
};

export default Navbar;
