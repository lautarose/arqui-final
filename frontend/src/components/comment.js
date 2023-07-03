import React from 'react';
import { Typography, Card, CardContent } from '@mui/material';
import { format } from 'date-fns';

export function Comment({ comment }) {
  const { user_id, message, created_at } = comment;

  const formattedDate = (date) => {
    return format(new Date(date), 'dd/MM/yyyy HH:mm');
  };

  return (
    <Card>
      <CardContent>
        <Typography variant="subtitle1">
          Usuario: {user_id}
        </Typography>
        <Typography variant="body1">{message}</Typography>
        <Typography variant="caption">
          Publicado el {formattedDate(created_at)}
        </Typography>
      </CardContent>
    </Card>
  );
}

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
  