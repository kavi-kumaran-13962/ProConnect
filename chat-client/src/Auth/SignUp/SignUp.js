import React from 'react';
import InputWithIcon from "../../InputWithIcon/InputWithIcon"
import "./SignUp.css"
import { Link } from 'react-router-dom';
import GoButton from '../GoButton/GoButton';
function SignUp () {

    const [username, setUsername] = React.useState('');
    const [password, setPassword] = React.useState('');
    const [confirmpassword, setconfirmpassword] = React.useState('');

    const handleSubmit = async () => {
        if (password !== confirmpassword) {
            alert("Password and confirm password are not equal");
            return
        }
        const data = {
          username,
          password
        };
    
      // Make an API call to sign in the user and get the JWT token.
      //   const response = await fetch('/api/signin', {
      //     method: 'POST',
      //     body: JSON.stringify(data)
      //   });
      console.log(username,password)
        const response = {
          "status":200
        }
        // Check if the response was successful.
        if (response.status === 200) {

          // Redirect the user to the home page.
          window.location.href = '/SignIn';
        } else {
          // Display an error message.
          alert(response.statusText);
        }
      };
      
    return (
        <div className='signUpContainer'>
            <div className='signUpContainer__headingContainer'>
                <h1 className='signUpContainer__headingContainer--txt'>Create an  account</h1>
            </div>
            <div className='signUpContainer__inputContainer'>
                <div className='signUpContainer__inputContainer--inputDiv'>
                    <InputWithIcon onChange={name => setUsername(name)} placeholderText="Enter username" icon="fa fa-user icon"/>
                </div>
                <div className='signUpContainer__inputContainer--inputDiv'>
                    <InputWithIcon onChange={pwd => setPassword(pwd)} placeholderText="Enter password" icon="fa fa-lock icon"/>
                </div>
                <div className='signUpContainer__inputContainer--inputDiv'>
                    <InputWithIcon onChange={pwd => setconfirmpassword(pwd)} placeholderText="confirm password" icon="fa fa-lock icon"/>
                </div>
            </div>
            <div className='signUpContainer__info'>
                <p className='signUpContainer__info--txt'>By clicking register you are creating a new account</p>
            </div>
            <div className='signUpContainer__reg'>
                <h2 className='signUpContainer__reg--txt' >Register</h2>
                <GoButton handleSubmit={handleSubmit}/>
            </div>
            
            <div className='signUpContainer__login'>
                <p className='signUpContainer__login--txt'>Do you have an account? <Link to={'/SignIn'} className='signUpContainer__login--link'>Login here</Link></p>
            </div>
        </div>
    );

};

export default SignUp;
