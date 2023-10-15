// src/App.tsx
import React from 'react';
import { BrowserRouter } from 'react-router-dom';
import ApplicationRoutes from "./config/ApplicationRoutes";
function App() {
  return (
    <BrowserRouter>
    <ApplicationRoutes />
    </BrowserRouter>
  );
}

export default App;
