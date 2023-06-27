import React from 'react';
import './css/home.css'
import { Link } from 'react-router-dom';


const Home = () => {
  return (
    <div className='container'>
        <Link to="/login">
        <button className="login-button">LOGIN</button>
        </Link>   
    </div>
  );
}

export default Home;