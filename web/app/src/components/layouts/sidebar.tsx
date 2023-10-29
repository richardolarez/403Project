import React from 'react';
import { Menu} from 'antd';
import {
    HomeOutlined,
    UserOutlined,
    VideoCameraOutlined,
    UploadOutlined,
    DollarOutlined,
    InfoCircleOutlined,
    FrownOutlined,
    DatabaseOutlined,
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
    const handleFileClick = () => {
        history('/files');
    }
    const handleCheckoutClick = () => {
        history('/checkout');
    }
    const handleInfoClick = () => {
        history('/info');
    }
const handleCustClick = () => {
    history('/customers');
}
    const handleInventoryClick = () => {
        history('/inventory');
    }

    return (
        <div>
            <div style={{height: "32px", margin: "16px", display: 'flex', alignItems: 'center', justifyContent: 'center',}}>
                <img
                    src="/poos_logo.png" // Update the image URL
                    alt="Logo"
                    style={{ marginTop: '5%', height: '100%', marginBottom: '5%',}}
                />
            </div>
            <Menu theme="dark" mode="inline" defaultSelectedKeys={['1']}>
                <Menu.Item key="1" onClick={handleHomeClick}>
                    <HomeOutlined/>
                    <span> Home</span>
                </Menu.Item>
                <Menu.Item key="2" onClick={handleUserClick}>
                    <UserOutlined />
                    <span> Users</span>
                </Menu.Item>
                <Menu.Item key="3" onClick={handleCustClick}>
                    <FrownOutlined />
                    <span> Customers</span>
                </Menu.Item>
                <Menu.Item key="4" onClick={handleCheckoutClick}>
                    <DollarOutlined />
                    <span> Checkout</span>
                </Menu.Item>
                
                <Menu.Item key="5" onClick={handleInventoryClick}>
                    <DatabaseOutlined />
                    <span> Inventory</span>
                </Menu.Item>
                <Menu.Item style={{
                    position: 'absolute',
                    bottom: 0,
                    zIndex: 1,
                    transition: 'all 0.2s',
                    }}key="5" onClick={handleInfoClick}>
                    <InfoCircleOutlined />
                    <span> Info</span>
                </Menu.Item>
            </Menu>
        </div>
    );
}
export default SideNav;