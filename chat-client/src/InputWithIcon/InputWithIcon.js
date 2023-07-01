
import React, { useState } from 'react';
import "./InputWithIcon.css"
function AuthInput(props) {
  const { icon, placeholderText } = props;
  return (
<div className="input-container">
  <i className={icon}></i>
  <input className="input-field" type="text" placeholder={placeholderText}/>
</div>
  );
}

export default AuthInput;