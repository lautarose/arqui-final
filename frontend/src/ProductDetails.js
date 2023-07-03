import { useParams } from 'react-router-dom';
import React, { useState, useEffect } from 'react';
import { Card, CardContent, CardMedia, Typography, Grid } from "@mui/material";
import { Link } from 'react-router-dom';
import Navbar from './components/navbar';
import { CommentList } from './components/comment';

const ProductDetails = () => {
  const { id } = useParams();
  const url = "http://localhost:8080/search/" + id;
  const [products, setProducts] = useState([]);
  const [comments, setComments] = useState([]);

  const fetchApi = async () => {
    try {
      const response = await fetch(url);
      const responseJSON = await response.json();
      setProducts(responseJSON);
    } catch (error) {
      console.error('Error fetching products:', error);
    }
  };

  const fetchComments = async () => {
    try {
      const commentsUrl = `http://localhost:8100/comments/${id}`;
      const response = await fetch(commentsUrl);
      const responseJSON = await response.json();
      setComments(responseJSON);
    } catch (error) {
      console.error('Error fetching comments:', error);
    }
  };

  useEffect(() => {
    fetchApi();
    fetchComments();
  }, []);

  return (
    <div className='container'>
      <Navbar />
      <Grid container spacing={2}>
        {products.map((product) => (
          <Grid item xs={12} key={product.id}>
            <Card>
              <CardMedia
                component="img"
                height="400"
                image={product.picture[0]}
                alt={product.title[0]}
              />
              <CardContent>
                <Typography gutterBottom variant="h5" component="div">
                  {product.title[0]}
                </Typography>
                <Typography variant="h6" color="text.secondary">
                  Price: {product.price[0]} {product.currency[0]}
                </Typography>
                <Typography variant="body2" color="text.secondary">
                  Description: {product.description[0]}
                </Typography>
              </CardContent>
            </Card>
          </Grid>
        ))}
      </Grid>

      <CommentList comments={comments} />
    </div>
  );
};

export default ProductDetails;