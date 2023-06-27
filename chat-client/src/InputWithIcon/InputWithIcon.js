
import React, { useState } from 'react';
import "./InputWithIcon.css"
function AuthInput(props) {
  const { icon, placeholderText } = props;
  return (
<div class="input-container">
  <i class={icon}></i>
  <input class="input-field" type="text" placeholder={placeholderText}/>
</div>
  );
}

export default AuthInput;