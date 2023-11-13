//src/components/pages/Home.tsx
import React from 'react';
import {useNavigate} from 'react-router';

const Home = () => {

   const userName = sessionStorage.getItem("UserFName");
   const history = useNavigate();

   const logout = function() {
    sessionStorage.removeItem('authenticated');
    sessionStorage.removeItem("UserFname");
    sessionStorage.removeItem("UserRole");

    //window.location.href = '/login';
    history('/login');
   }

  return (
    <div>
        <h1 style={{textAlign: 'center'}}>Welcome {userName}! 👋</h1>
        <button id="logout" style={{position:'absolute',top:15,right:15}} onClick={logout}>Logout</button>
    </div>
  );
}
export default Home;