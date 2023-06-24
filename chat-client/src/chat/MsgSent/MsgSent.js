import React, { useState } from 'react';
import "./MsgSent.css"
function MsgSent ({ message }) {

    return (
        <div className='msgSent'>
            <p className='msgSent__txt'>{message}</p>
        </div>
    );
};

export default MsgSent;
