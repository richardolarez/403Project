import React from 'react';
import { Route, Routes, Navigate} from "react-router-dom";
import List from "../components/pages/list";
import Form from "../components/pages/form";
import File from "../components/pages/files";
import Videos from "../components/pages/videos";
import Login from '../components/pages/login';
import Home from '../components/pages/home';
import Checkout from '../components/pages/checkout';
import MainLayout from '../components/layouts/MainLayout'; // Import the main application layout
import LoginLayout from '../components/layouts/LoginLayout'; // Import the login page layout

const ApplicationRoutes: React.FC = () => {
  const isUserAuthenticated = sessionStorage.getItem("authenticated"); // Replace with your authentication logic

  return (
    <Routes>
      <Route
        path="/login"
        element={
          isUserAuthenticated ? <Navigate to="/home" /> : <LoginLayout><Login /></LoginLayout>
        }
      />
      <Route
        path="*"
        element={
          isUserAuthenticated ? (
            <MainLayout>
              <Routes>
                <Route path="/Home" element={<Home />} />
                <Route path="/list" element={<List />} />
                <Route path="/form" element={<Form />} />
                <Route path="/files" element={<File />} />
                <Route path="/videos" element={<Videos />} />
                <Route path="/checkout" element={<Checkout />} />
                <Route path="/info" element={<Videos />} />
              </Routes>
            </MainLayout>
          ) : (
            <Navigate to="/login" replace />
          )
        }
      />
    </Routes>
  );
};

export default ApplicationRoutes;