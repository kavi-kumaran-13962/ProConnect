import React, { useState } from 'react';
import "./Profile.css"
const Profile = () => {
    return (
        <div className="profile">
            <div className="profile__imagediv">
                <img className='profile__imagediv--img' src="https://www.gravatar.com/avatar/1234567890abcdef?s=50&d=identicon&r=PG" />
            </div>
            <div className='profile__content'>
                <div className="profile__namediv">
                    <h3 className='profile__namediv--txt'>John Doe</h3>
                </div>
            </div>
        </div>
    );
};

export default Profile;
