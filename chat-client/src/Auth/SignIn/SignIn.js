import React from 'react';
import InputWithIcon from "../../InputWithIcon/InputWithIcon"
import GoButton from '../GoButton/GoButton';
import "./SignIn.css"
import { Link } from 'react-router-dom';

const SignIn = () => {

    const [username, setUsername] = React.useState('');
    const [password, setPassword] = React.useState('');
  
    const handleSubmit = async () => {
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
        "token":"abcd",
        "status":200
      }
      // Check if the response was successful.
      if (response.status === 200) {
        // Get the JWT token from the response.
        const token = await response.token;
  
        document.cookie = 'token=' + token
  
        // Redirect the user to the home page.
        window.location.href = '/';
      } else {
        // Display an error message.
        alert(response.statusText);
      }
    };
  
    return (
      <div className='signInContainer'>
        <div className='signInContainer__headingContainer'>
          <h1 className='signInContainer__headingContainer--txt'>Welcome Back !</h1>
        </div>
        <div className='signInContainer__inputContainer'>
          <div className='signInContainer__inputContainer--inputDiv'>
            <InputWithIcon
              placeholderText="Enter username"
              value={username}
              onChange={name => setUsername(name)}
              icon="fa fa-user icon"
            />
          </div>
          <div className='signInContainer__inputContainer--inputDiv'>
            <InputWithIcon
              placeholderText="Enter password"
              value={password}
              onChange={pwd => setPassword(pwd)}
              icon="fa fa-lock icon"
            />
          </div>
        </div>
        <div className='signInContainer__reg'>
          <h2 className= 'signInContainer__reg--txt'>Sign In</h2>
          <GoButton handleSubmit={handleSubmit}/>
        </div>
        <div className='signInContainer__login'>
          <p className='signInContainer__login--txt'>Don’t  have an account? <Link to={'/SignUp'} className='signInContainer__login--link'>Register</Link></p>
        </div>
      </div>
    );
  
  };
  
  export default SignIn;
  
