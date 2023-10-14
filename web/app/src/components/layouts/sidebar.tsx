import React from 'react';
import { Menu } from 'antd';
import {
    HomeOutlined,
    UserOutlined,
    VideoCameraOutlined,
    UploadOutlined,
    DollarOutlined,
    InfoCircleOutlined,
  } from '@ant-design/icons';
import {useNavigate}  from 'react-router';
const SideNav = () => {
    const history = useNavigate();

const handleHomeClick = () => {
        history('/home');
    }
const handleUserClick = () => {
        history('/list');
    }
const handleVideosClick = () => {
        history('/videos');
    }
const handleFileClick = () => {
        history('/files');
    }
const handleCheckoutClick = () => {
        history('/checkout');
    }
const handleInfoClick = () => {
        history('/info');
}
return (
      <div>
<div style={{height: "32px", background: "rgba(255, 255, 255, 0.2)", margin: "16px"}}></div>
            <Menu theme="dark" mode="inline" defaultSelectedKeys={['1']}>
                <Menu.Item key="1" onClick={handleHomeClick}>
                    <HomeOutlined/>
                    <span> Home</span>
                </Menu.Item>
                <Menu.Item key="2" onClick={handleUserClick}>
                    <UserOutlined />
                    <span> Users</span>
                </Menu.Item>
                <Menu.Item key="3" onClick={handleVideosClick}>
                    <VideoCameraOutlined />
                    <span> Videos</span>
                </Menu.Item>
                <Menu.Item key="4" onClick={handleFileClick}>
                    <UploadOutlined />
                    <span> Files</span>
                </Menu.Item>
                <Menu.Item key="5" onClick={handleCheckoutClick}>
                    <DollarOutlined />
                    <span> Checkout</span>
                </Menu.Item>
                <Menu.Item style={{
                    position: 'absolute',
                    bottom: 0,
                    zIndex: 1,
                    transition: 'all 0.2s',
                    }}key="6" onClick={handleInfoClick}>
                    <InfoCircleOutlined />
                    <span> Info</span>
                </Menu.Item>
            </Menu>
        </div>
  );
}
export default SideNav;