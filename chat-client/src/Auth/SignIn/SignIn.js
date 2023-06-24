import React from 'react';
import AuthInput from "../AuthInput/AuthInput"
import GoButton from '../GoButton/GoButton';
import "./SignIn.css"
function SignIn () {
    return (
        <div className='signInContainer'>
            <div className='signInContainer__headingContainer'>
                <h1 className='signInContainer__headingContainer--txt'>Welcome Back !</h1>
            </div>
            <div className='signInContainer__inputContainer'>
                <div className='signInContainer__inputContainer--inputDiv'>
                    <AuthInput placeholderText="Enter username" icon="fa fa-user icon"/>
                </div>
                <div className='signInContainer__inputContainer--inputDiv'>
                    <AuthInput placeholderText="Enter password" icon="fa fa-lock icon"/>
                </div>
            </div>
            <div className='signInContainer__reg'>
                <h2 className= 'signInContainer__reg--txt'>Sign In</h2>
                <GoButton/>
            </div>
            <div className='signInContainer__login'>
                <p className='signInContainer__login--txt'>Don’t  have an account? <a className='signInContainer__login--link'>Register</a></p>
            </div>
        </div>
    );

};

export default SignIn;
