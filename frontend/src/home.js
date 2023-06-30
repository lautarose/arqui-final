import React, { useState, useEffect } from 'react';
import './css/home.css';
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
  const url = "http://localhost:8090/items/649f4ffd766a191ba2a3ece4";
  
  const [products, setProducts] = useState({});

  const fetchApi = async () => {
    try {
      const response = await fetch(url);
      const responseJSON = await response.json();
      setProducts(responseJSON);
    } catch (error) {
      console.error('Error fetching products:', error);
    }
  };
  

  useEffect(() => {
    fetchApi();
  }, []);

  return ( 
      <div className='container'>
        <Typography variant="h5" component="div">
          {products.title}
        </Typography>
        <Typography variant="body2" color="text.secondary">
          Price: {products.price}
        </Typography>
      </div>
  );
};

export default Home;
