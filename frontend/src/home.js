import React, { useState, useEffect } from 'react';
import './css/home.css';
import { Card, CardContent, CardMedia, Typography, Grid } from "@mui/material";
import { Link } from 'react-router-dom';
import SearchIcon from '@mui/icons-material/Search';
import Navbar from './components/navbar';

function truncateDescription(description, limit) {
  const words = description.split(" ");
  if (words.length > limit) {
    return words.slice(0, limit).join(" ") + " ...";
  }
  return description;
}

const Home = () => {
  const url = "http://localhost:8080/search/";
  
  const [products, setProducts] = useState([]);

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
       <Navbar />
      <Grid container spacing={2}>
        {products.map((product) => (
          <Grid item xs={12} sm={6} md={4} key={product.id} className='hover'>
            <Card className='product-card'>
              <CardMedia
                component="img"
                height="200"
                image={product.picture[0]}
                alt={product.title[0]}
              />
              <CardContent>
                <Typography gutterBottom variant="h5" component="div">
                  {product.title[0]}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Price: {product.price[0]} {product.currency[0]}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  {truncateDescription(product.description[0], 15)}
                </Typography>
              </CardContent>
              <Link to={`/product/${product.id}`} className="product-link">
                View Details
              </Link>
            </Card>
          </Grid>
        ))}
      </Grid>
    </div>
  );
};

export default Home;

