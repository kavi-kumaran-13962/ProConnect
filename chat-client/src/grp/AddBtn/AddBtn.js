import React from 'react';
import "./AddBtn.css"
import { Link } from 'react-router-dom';

const AddBtn = () => {

  return (
    <Link to={'/createGrp'}>
        <button className='addBtn' >+</button>
    </Link>
  );
};

export default AddBtn;
