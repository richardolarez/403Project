//src/components/pages/Home.tsx
import React from 'react';
import {Button, Modal, Input} from 'antd';
import {useEffect, useState} from 'react';
import axios from 'axios';


const Home = () => {

const [isModalVisible, setIsModalVisible] = React.useState(false);
const [newPassword, setNewPassword] = React.useState('');
//const [oldPassword, setOldPassword] = React.useState('');
const [currentUsername, setCurrentUsername] = React.useState('');
const [currentPassword, setCurrentPassword] = React.useState('');
//const [shouldUpdatePassword, setShouldUpdatePassword] = React.useState(false);
const [resolveModalPromise, setResolveModalPromise] = React.useState<(() => void) | null>(null);



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