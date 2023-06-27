import React, { useState } from "react";
import "./Nav.css";
import { Link } from "react-router-dom";

export default function Nav() {
  const [activeIndex, setActiveIndex] = useState(null);

  const handleItemClick = (index) => {
    setActiveIndex(index);
  };

  return (
    <div>
      <nav>
        <ul className="nav__list">
          <li
            className={`nav__li ${activeIndex === 0 ? 'active' : ''}`}
            onClick={() => handleItemClick(0)}
          >
            <Link className="nav__icon" to="profile">
              <i className={`nav__icon--icon fa fa-user`} />
            </Link>
          </li>
          <li
            className={`nav__li ${activeIndex === 1 ? 'active' : ''}`}
            onClick={() => handleItemClick(1)}
          >
            <Link className="nav__icon" to="/chats">
              <i className={`nav__icon--icon far fa-comment-alt`} />
            </Link>
          </li>
          <li
            className={`nav__li ${activeIndex === 2 ? 'active' : ''}`}
            onClick={() => handleItemClick(2)}
          >
            <Link className="nav__icon" to="/grpchats">
              <i className={`nav__icon--icon fa fa-users`} />
            </Link>
          </li>
        </ul>
      </nav>
    </div>
  );
}
