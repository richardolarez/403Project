//src/components/pages/Home.tsx
import React from 'react';
import {Button, Modal, Input} from 'antd';
import {useEffect, useState} from 'react';
import axios from 'axios';


const Home = () => {

const [isModalVisible, setIsModalVisible] = React.useState(false);
const [isModal2Visible, setIsModal2Visible] = React.useState(false);
const [isModal3Visible, setIsModal3Visible] = React.useState(false);
const [newPassword, setNewPassword] = React.useState('');
const [currentUsername, setCurrentUsername] = React.useState('');
const [currentPassword, setCurrentPassword] = React.useState('');
const [resolveModalPromise, setResolveModalPromise] = React.useState<(() => void) | null>(null);
const [userRole, setUserRole] = useState('');
const [lowQuantityDrugs, setLowQuantityDrugs] = useState<MedicineItem[]>([]);
const [expiredDrugs, setExpiredDrugs] = useState<MedicineItem[]>([]);

interface MedicineItem {
  ID: number;
  Drug: string;
  Doses: number;
  Strength: string;
  Price: number;
  ExpirationDate: string;
}

useEffect(() => {
  // Load user role from session storage
  const storedUserRole = sessionStorage.getItem('UserRole');
  if (storedUserRole) {
    setUserRole(storedUserRole);
  }

}, []);

useEffect(() => {
  // Check for the user role and perform actions accordingly
  if (userRole === 'Manager') {
    axios.post<MedicineItem[]>(`http://localhost:8080/medicines`)
      .then(res => {
        console.log(res);

        // Filter items with quantity less than 20
        const lowQuantityDrugs = res.data
          .filter((item: MedicineItem) => item.Doses < 20)
          .map((item: MedicineItem) => item);

        setLowQuantityDrugs(lowQuantityDrugs);

        // Show the modal if there are low quantity drugs
        if (lowQuantityDrugs.length > 0) {
          showModal2();
        }

        // Filter items with expiration date older than 12 months
        const twelveMonthsAgo = new Date();
          twelveMonthsAgo.setMonth(twelveMonthsAgo.getMonth() - 12);
          
          
          const expiredDrugs = res.data
            .filter((item: MedicineItem) => {
              const expirationDate = new Date(item.ExpirationDate);
              console.log(expirationDate)
              return expirationDate < twelveMonthsAgo;
            })
            .map((item: MedicineItem) => item);

          setExpiredDrugs(expiredDrugs);
      })
      .catch(error => {
        console.log("Error in retrieving medicines:", error);
      });
  } else {
    console.log("Not a manager");
  }
}, [userRole]);

const showModal2 = () => {
  setIsModal2Visible(true);
};

const handleCancel2 = () => {
  setIsModal2Visible(false);
  if(expiredDrugs.length > 0){
    setIsModal3Visible(true)
  }
};

const handleCancel3 = () => {
  setIsModal3Visible(false)
}

const handleOk = async () => {
  if (newPassword && currentUsername) {
    try {
      await updatePassword(currentUsername, currentPassword, newPassword);
      if (resolveModalPromise) {
        resolveModalPromise(); 
      }
      setIsModalVisible(false);
      // Reset the states after successful password update
      setCurrentUsername('');
      setCurrentPassword('');
      setNewPassword('');
    } catch (error) {
      console.error('Error updating password:', error);
    }
  }
};

const userName = sessionStorage.getItem("UserFName");

const showModal = () => {
  setIsModalVisible(true);
  return new Promise<void>(resolve => {
    setResolveModalPromise(() => resolve);
  });
};

const handleCancel = () => {
  setIsModalVisible(false);
  setNewPassword('');
};

const handlePasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
  setNewPassword(e.target.value);
  
};

const handleCurrentPasswordChange = (e: React.ChangeEvent<HTMLInputElement>) => {
  setCurrentPassword(e.target.value);
};

const handleCurrentUsernameChange = (e: React.ChangeEvent<HTMLInputElement>) => {
  setCurrentUsername(e.target.value);
};

const updatePassword = async (username: string, oldPassword: string, newPassword: string) => {
  //Data package
  const data = {
    username: username,
    oldPassword: oldPassword,
    newPassword: newPassword
  };

  // API call to update pass
  try {
    const response = await fetch('http://localhost:8080/updatePassword', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }
  } catch (error) {
    console.error('Error:', error);
    throw error; 
  }
};

  return (
    <div>
      <Modal
        title="Low Quantity Drugs!!11!"
        open={isModal2Visible}
        onCancel={handleCancel2}
        footer={null}
      >
          <p>Low quantity drugs:</p>
          <ul>
          {lowQuantityDrugs.map(drug => (
            <li key={drug.ID}>
            <p>Drug: {drug.Drug}</p>
            <p>Doses: {drug.Doses}</p>
            <p>Strength: {drug.Strength}</p>
            <p>Price: {drug.Price}</p>
            <p>Expiration Date: {drug.ExpirationDate}</p>
          </li>
          ))}
        </ul>
      </Modal>
      <Modal
        title="Expired Drugs!!11!"
        open={isModal3Visible}
        onCancel={handleCancel3}
        footer={null}
      >
          <p>expired drugs:</p>
          <ul>
          {expiredDrugs.map(drug => (
            <li key={drug.ID}>
              <p>Drug: {drug.Drug}</p>
              <p>Doses: {drug.Doses}</p>
              <p>Strength: {drug.Strength}</p>
              <p>Price: {drug.Price}</p>
              <p>Expiration Date: {drug.ExpirationDate}</p>
            </li>
          ))}
        </ul>
      </Modal>
    <Modal
    title="Change Password"
    open={isModalVisible}
    onOk={handleOk}
    onCancel={handleCancel}
    >
      <Input
      value={currentUsername} 
      onChange={handleCurrentUsernameChange} 
      placeholder="Enter Username"
    />
    <Input.Password 
      value={currentPassword} 
      onChange={handleCurrentPasswordChange} 
      placeholder="Enter current password" 
    />
    <Input.Password 
      value={newPassword} 
      onChange={handlePasswordChange} 
      placeholder="Enter new password" 
    />
  </Modal> 
        <h1 style={{textAlign: 'center'}}>Welcome {userName}! 👋</h1>
        <Button onClick={showModal} block>Change Password</Button>
    </div>
  );
}
export default Home;