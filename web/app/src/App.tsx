import React, { useState } from 'react';
import './Login.css';

function Login() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    // Call the /login endpoint to authenticate the employee
    const response = await fetch('http://localhost:8080/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ username, password })
    });

    if (response.ok) {
      const employee = await response.json();
      console.log('Authenticated employee:', employee);
      // TODO: Handle authenticated employee
    } else {
      console.error('Failed to authenticate employee:', response.statusText);
      // TODO: Handle authentication error
    }
  };

  return (
    <div className="Login">
      <form onSubmit={handleSubmit}>
        <label>
          Username:
          <input type="text" value={username} onChange={event => setUsername(event.target.value)} />
        </label>
        <label>
          Password:
          <input type="password" value={password} onChange={event => setPassword(event.target.value)} />
        </label>
        <button type="submit">Login</button>
      </form>
    </div>
  );
}

export default Login;