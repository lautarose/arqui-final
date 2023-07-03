import React from 'react';
import { Typography, Card, CardContent, Button } from '@mui/material';
import { format } from 'date-fns';
import Cookies from 'universal-cookie';

const Comment = ({ comment, onDeleteComment }) => {
  const { comment_id, user_id, message, created_at } = comment;

  const formattedDate = (date) => {
    return format(new Date(date), 'dd/MM/yyyy HH:mm');
  };

  const handleDelete = async () => {
    const cookies = new Cookies();
    const token = cookies.get('user');

    try {
      const response = await fetch(`http://localhost:8100/comments/${comment_id}`, {
        method: 'DELETE',
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (response.ok) {
        onDeleteComment(comment_id);
        window.location.reload(); // Actualizar la página
      } else {
        console.error('Error deleting comment:', response.status);
      }
    } catch (error) {
      console.error('Error deleting comment:', error);
    }
  };

  return (
    <Card>
      <CardContent>
        <Typography variant="subtitle1">Usuario: {user_id}</Typography>
        <Typography variant="body1">{message}</Typography>
        <Typography variant="caption">Publicado el {formattedDate(created_at)}</Typography>
        <Button variant="outlined" color="error" onClick={handleDelete}>
          Eliminar
        </Button>
      </CardContent>
    </Card>
  );
};

export default Comment;


export function CommentList({ comments }) {
    return (
      <div>
        {comments ? (
          comments.map((comment) => (
            <Comment key={comment.comment_id} comment={comment} />
          ))
        ) : (
          <p>No hay comentarios</p>
        )}
      </div>
    );
  }
  