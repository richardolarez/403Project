//src/components/pages/Home.tsx
import React from 'react';
const Home = () => {
   const userName = sessionStorage.getItem("UserFName");
  return (
    <div>
        <h1>Hewwo UwU {userName}</h1> 
    </div>
  );
}
export default Home;