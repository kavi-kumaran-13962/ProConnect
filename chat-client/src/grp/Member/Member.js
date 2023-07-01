import React from "react";
import "./Member.css";

const Member = ({ id, name, onClick }) => {

  const [isTicked, setIsTicked] = React.useState(false);
  
  const handleClick = () => {
    setIsTicked(!isTicked);
    onClick(id);
  };

  return (
      <div
        className="member"
        id={id}
        onClick={handleClick}
      >
        <div className="member__imagediv">
          <img
            className="member__imagediv--img"
            src="https://www.gravatar.com/avatar/1234567890abcdef?s=50&d=identicon&r=PG"
            alt="Profile"
          />
        </div>
        <div className="member__content">
          <div className="member__namediv">
            <h4 className="member__namediv--txt">{name}</h4>
          </div>
        </div>
        <div
            className="member__tick"
            style={{
              color: isTicked ? "white" : "transparent",
              right: "0"
            }}
          >
            <i className="fas fa-check"></i>
          </div>
      </div>
  );
};

export default Member;
