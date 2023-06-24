import React from "react";
import { Link } from "react-router-dom";
import "./Chat.css";

const Chat = ({ id, name, lastMessage }) => {

  return (
    <Link className="chat__link" to={id}>
      <div className="chat">
        <div className="chat__imagediv">
          <img
            className="chat__imagediv--img"
            src="https://www.gravatar.com/avatar/1234567890abcdef?s=50&d=identicon&r=PG"
            alt="Profile"
          />
        </div>
        <div className="chat__content">
          <div className="chat__namediv">
            <h4 className="chat__namediv--txt">{name}</h4>
          </div>
          <div className="chat__messagediv">
            <p className="chat__messagediv--txt">{lastMessage}</p>
          </div>
        </div>
      </div>
    </Link>
  );
};

export default Chat;
