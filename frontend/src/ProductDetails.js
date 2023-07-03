import { useParams } from 'react-router-dom';
import React, { useState, useEffect } from 'react';
import { Card, CardContent, CardMedia, Typography, Grid, TextField, Button } from '@mui/material';
import Navbar from './components/navbar';
import { CommentList } from './components/comment';
import Cookies from 'universal-cookie';

const ProductDetails = () => {
  const { id } = useParams();
  const url = `http://localhost:8080/search/${id}`;
  const commentsUrl = `http://localhost:8100/comments/load`;
  const [products, setProducts] = useState([]);
  const [comments, setComments] = useState([]);
  const [newComment, setNewComment] = useState('');
  const [token, setToken] = useState('');
  const cookie = new Cookies();

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
      const response = await fetch(`http://localhost:8100/comments/${id}`);
      const responseJSON = await response.json();
      setComments(responseJSON);
    } catch (error) {
      console.error('Error fetching comments:', error);
    }
  };

  const handleCommentSubmit = async (event) => {
    event.preventDefault();

    const commentData = {
      item_id: id,
      message: newComment,
    };

    try {
      const response = await fetch(commentsUrl, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify(commentData),
      });

      if (response.ok) {
        // Comentario enviado exitosamente, volver a cargar los comentarios
        setNewComment('');
        fetchComments();
      } else {
        console.error('Error submitting comment:', response.status);
      }
    } catch (error) {
      console.error('Error submitting comment:', error);
    }
  };

  useEffect(() => {
    fetchApi();
    fetchComments();
  }, []);

  useEffect(() => {
    const userToken = cookie.get('user');
    setToken(userToken);
  }, []);

  return (
    <div className="container">
      <Navbar />
      <Grid container spacing={2}>
        {products.map((product) => (
          <Grid item xs={12} key={product.id}>
            <Card>
              <CardMedia component="img" height="400" image={product.picture[0]} alt={product.title[0]} />
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

      <form onSubmit={handleCommentSubmit}>
        <TextField
          label="Nuevo comentario"
          value={newComment}
          onChange={(event) => setNewComment(event.target.value)}
          variant="outlined"
          margin="normal"
          fullWidth
        />
        <Button type="submit" variant="contained" color="primary">
          Enviar comentario
        </Button>
      </form>

      <CommentList comments={comments} />
    </div>
  );
};

export default ProductDetails;
