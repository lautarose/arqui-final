import React from 'react';
import { useParams } from 'react-router-dom';
import products from './utils/items.js';

const ProductDetail = () => {
  const { id } = useParams();

  // Buscar el producto correspondiente al id en el archivo items.js
  const product = products.find((product) => product.id === parseInt(id));

  if (!product) {
    return <div>No se encontró el producto.</div>;
  }

  return (
    <div>
      <h2>{product.name}</h2>
      <p>{product.description}</p>
    </div>
  );
};

export default ProductDetail;
