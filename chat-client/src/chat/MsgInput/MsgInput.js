import React, { useState } from 'react';
import "./MsgInput.css"
const MsgInput = () => {
    const [message, setMessage] = useState('');

    const handleInputChange = (e) => {
        setMessage(e.target.value);
    };

    const handleSend = () => {
        // Handle sending the message here
        console.log(message);
        setMessage('');
    };

    return (
        <div className='msgInput'>
            <input 
                type="text" 
                value={message} 
                onChange={handleInputChange} 
                placeholder="Write your message..."     
                className='msgInput__inp'
            />
            <div 
                onClick={handleSend} 
                className='msgInput__sendDiv'
            >
            <i className='msgInput__sendicon fa fa-paper-plane icon'></i>
            </div>
        </div>
    );
};

export default MsgInput;
