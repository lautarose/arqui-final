import React, { useEffect, useState } from 'react';

var token = getCookie("user");
function getCookie(name) {
  var cookieValue = null;
  if (document.cookie && document.cookie !== '') {
    var cookies = document.cookie.split(';');
    for (var i = 0; i < cookies.length; i++) {
      var cookie = cookies[i].trim();
      if (cookie.substring(0, name.length + 1) === (name + '=')) {
        cookieValue = decodeURIComponent(cookie.substring(name.length + 1));
        break;
      }
    }
  }
  return cookieValue;
}

const Publics = () => {

  fetch('http://localhost:8081/user/' + token)
  .then(response => response.json())
  .then(data => {
    var userId = data.userId;
    fetch('http://localhost:8090/items/user/' + userId)
  .then(response => response.json())
  .then(data => {
    // Aquí puedes utilizar los datos de las publicaciones para mostrarlos en la página.
    console.log(data);
    })
     .catch(error => {
      console.error('Error:', error);
    });
  })
  .catch(error => {
    console.error('Error:', error);
  });

  return (
    <div>hola</div>
  );
}

export default Publics;
