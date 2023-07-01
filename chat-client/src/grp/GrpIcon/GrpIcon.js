import React, { useState } from 'react';
import "./GrpIcon.css"
const GrpIcon = () => {
    return (
        <div className="grpicon">
            <div className="grpicon__imagediv">
                <img className='grpicon__imagediv--img' src="https://www.gravatar.com/avatar/1234567890abcdef?s=50&d=identicon&r=PG" />
            </div>
            <div className='grpicon__content'>
                <div className="grpicon__namediv">
                    <h3 className='grpicon__namediv--txt'>Add Grp Icon</h3>
                </div>
            </div>
        </div>
    );
};

export default GrpIcon;
