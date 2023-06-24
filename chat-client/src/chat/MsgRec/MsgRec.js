import React from "react";
import "./MsgRec.css"
function MsgRec ({ message }){
  return (
    <div className='msgRec'>
        <img src="https://www.gravatar.com/avatar/1234567890abcdef?s=50&d=identicon&r=PG" />
        <p className='msgRec__txt'>{message}</p>
    </div>
  );
};

export default MsgRec;
