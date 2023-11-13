//src/components/pages/Home.tsx
import React from 'react';

const Home = () => {

   const userName = sessionStorage.getItem("UserFName");

  return (
    <div>
        <h1 style={{textAlign: 'center'}}>Welcome {userName}! 👋</h1>
    </div>
  );
}
export default Home;