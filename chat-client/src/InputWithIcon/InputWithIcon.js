import React, { useState } from "react";
import "./InputWithIcon.css";
function InputWithIcon(props) {

  const { icon, placeholderText, onChange } = props;
  const [text, settext] = useState('');

  const handleInputChange = (e) => {
    settext(e.target.value);
    onChange(e.target.value)
};
  return (
    <div className="input-container">
      <i className={icon}></i>
      <input
        className="input-field"
        type="text"
        placeholder={placeholderText}
        value={text} 
        onChange={handleInputChange} 
      />
    </div>
  );
}

export default InputWithIcon;
