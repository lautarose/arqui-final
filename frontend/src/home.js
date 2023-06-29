import React from 'react';
import './css/home.css';
import products from "./utils/items.js";
import { Card, CardContent, CardMedia, Typography, Grid } from "@mui/material";
import { Link } from 'react-router-dom';
import SearchIcon from '@mui/icons-material/Search';

function truncateDescription(description, limit) {
  const words = description.split(" ");
  if (words.length > limit) {
    return words.slice(0, limit).join(" ") + " ...";
  }
  return description;
}

const Home = () => {
  return (
    <div className='container'>
      <div className='barra'>
        <img src='https://cdn.freebiesupply.com/logos/large/2x/jl-logo-black-and-white.png' className='loguito'/>
        <Link to="/login">
        <button className="login-button">LOGIN</button>
        </Link>  
        <div className='busqueda'>
          <input type="text" placeholder="Buscar" class="search-input"/>
          <button className='lupa'>
            <SearchIcon />
          </button>
        </div>
      </div>
        <Grid container spacing={2} >
      {products.map((product) => (
        <Grid item key={product.id} xs={12} sm={6} md={4} >
          <Card sx={{ height: "100%" }} className="product-card">
            <CardMedia 
              component="img"
              height="200"
              image={product.picture}
              alt={product.id}
            />
            <CardContent>
              <Typography variant="h6" component="div">
                {product.title}
              </Typography>
              <Typography variant="subtitle1" color="textSecondary">
                {product.seller}
              </Typography>
              <Typography variant="body2" color="textSecondary">
                Precio: {product.price} {product.currency}
              </Typography>
              <Typography variant="body2" component="p">
                Descripción:{truncateDescription(product.description, 15)}
              </Typography>
              <Typography variant="body2" color="textSecondary">
                Estado: {product.state}
              </Typography>
              <Typography variant="body2" color="textSecondary">
                Ciudad: {product.city}
              </Typography>
              <Typography variant="body2" color="textSecondary">
                Calle: {product.street} {product.number}
              </Typography>
            </CardContent>
          </Card>
        </Grid>
      ))}
    </Grid>
    </div>
  );
}

export default Home;