import React from 'react';
import "./GoButton.css"

const GoButton = ({ handleSubmit }) => {

  return (
    <>
        <button className='gobtn' onClick={handleSubmit} >-&gt;</button>
    </>
  );
};

export default GoButton;
