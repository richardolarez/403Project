import React from 'react';
import { Route, Routes, Navigate} from "react-router-dom";
import Login from '../components/pages/login';
import Home from '../components/pages/home';
import List from "../components/pages/list";
import Prescriptions from '../components/pages/prescriptions';
import Form from "../components/pages/form";
import Checkout from '../components/pages/checkout';
import InventoryList from '../components/pages/InventoryList'; // Import the inventory page
import MainLayout from '../components/layouts/MainLayout'; // Import the main application layout
import LoginLayout from '../components/layouts/LoginLayout'; // Import the login page layout
import Customers from '../components/pages/customers';
import CustForm from '../components/pages/custForm';
import Inventory from '../components/pages/inventory';
import PaymentPage from '../components/pages/payment';
import Reports from '../components/pages/reports';
import Info from '../components/pages/info';

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
                <Route path="/prescriptions" element={<Prescriptions />} />
                <Route path="/form" element={<Form />} />
                <Route path="/checkout" element={<Checkout />} />
                <Route path="/customers" element={<Customers />} />
                <Route path="/custForm" element={<CustForm />} />
                <Route path="/inventoryList" element={<InventoryList />} />
                <Route path="/inventory" element={<Inventory />} />
                <Route path='/payment' element={<PaymentPage />} />
                <Route path='/reports' element={<Reports />} />
                <Route path="/info" element={<Info />} />
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