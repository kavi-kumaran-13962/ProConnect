import React from 'react';
import AuthInput from "../AuthInput/AuthInput"
import "./SignUp.css"
function SignUp () {
    return (
        <div className='signUpContainer'>
            <div className='signUpContainer__headingContainer'>
                <h1 className='signUpContainer__headingContainer--txt'>Create an  account</h1>
            </div>
            <div className='signUpContainer__inputContainer'>
                <div className='signUpContainer__inputContainer--inputDiv'>
                    <AuthInput placeholderText="Enter username" icon="fa fa-user icon"/>
                </div>
                <div className='signUpContainer__inputContainer--inputDiv'>
                    <AuthInput placeholderText="Enter password" icon="fa fa-lock icon"/>
                </div>
                <div className='signUpContainer__inputContainer--inputDiv'>
                    <AuthInput placeholderText="confirm password" icon="fa fa-lock icon"/>
                </div>
            </div>
            <div className='signUpContainer__info'>
                <p className='signUpContainer__info--txt'>By clicking register you are creating a new account</p>
            </div>
            <div className='signUpContainer__reg'>
                <h2 className='signUpContainer__reg--txt' >Register</h2>
                <button className='signUpContainer__reg--btn'>-&gt;</button>
            </div>
            <div className='signUpContainer__login'>
                <p className='signUpContainer__login--txt'>Do you have an account? <a className='signUpContainer__login--link'>Login here</a></p>
            </div>
        </div>
    );

};

export default SignUp;
