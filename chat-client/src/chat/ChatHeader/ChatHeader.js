import React, { useState } from 'react';
import "./ChatHeader.css"
function ChatHeader ({name}) {
    return (
        <div className="chatHeader">
            <div className="chatHeader__imagediv">
                <img className='chatHeader__imagediv--img' src="https://www.gravatar.com/avatar/1234567890abcdef?s=50&d=identicon&r=PG" />
            </div>
            <div className='chatHeader__name'>
                <h3 className='chatHeader__name--txt'>{name}</h3>
            </div>
        </div>
    );
};

export default ChatHeader;
