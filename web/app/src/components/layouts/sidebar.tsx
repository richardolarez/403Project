import React from 'react';
import { Menu } from 'antd';
import {
    UserOutlined,
    VideoCameraOutlined,
    UploadOutlined,
    InfoCircleOutlined,
  } from '@ant-design/icons';
import {useNavigate}  from 'react-router';
const SideNav = () => {
    const history = useNavigate();
const handleUserClick = () => {
        history('/list');
    }
const handleVideosClick = () => {
        history('/videos');
    }
const handleFileClick = () => {
        history('/files');
    }
const handleInfoClick = () => {
        history('/info');
}
return (
      <div>
        <div style={{height: "32px", background: "rgba(255, 255, 255, 0.2)", margin: "16px"}}></div>
            <Menu theme="dark" mode="inline" defaultSelectedKeys={['1']}>
                <Menu.Item key="1" onClick={handleUserClick}>
                    <UserOutlined />
                    <span> Users</span>
                </Menu.Item>
                <Menu.Item key="2" onClick={handleVideosClick}>
                    <VideoCameraOutlined />
                    <span> Videos</span>
                </Menu.Item>
                <Menu.Item key="3" onClick={handleFileClick}>
                    <UploadOutlined />
                    <span> Files</span>
                </Menu.Item>
                <Menu.Item key="4" onClick={handleInfoClick}>
                    <InfoCircleOutlined />
                    <span> Info</span>
                </Menu.Item>
            </Menu>
        </div>
  );
}
export default SideNav;