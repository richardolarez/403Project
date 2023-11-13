import React, { useState, useEffect } from 'react';
import { Layout } from 'antd';
import SideNav from './sidebar';
import {
  MenuUnfoldOutlined,
  MenuFoldOutlined
} from '@ant-design/icons';

const { Header, Sider, Content } = Layout;

interface LayoutProps {
  children: React.ReactNode;
}

function logout() {
  
  console.log('what the fuck is going on here');
  
  sessionStorage.removeItem('authenticated');
  sessionStorage.removeItem("UserFname");
  sessionStorage.removeItem("UserRole");

  window.location.href = '/login';
}

const MainLayout: React.FC<LayoutProps> = ({ children }) => {
  const [collapse, setCollapse] = useState(false);

  useEffect(() => {
    window.innerWidth <= 760 ? setCollapse(true) : setCollapse(false);
  }, []);

  const handleToggle = (event: any) => {
    event.preventDefault();
    collapse ? setCollapse(false) : setCollapse(true);
  };

  return (
    <Layout>
      <button id="logout" style={{position:'absolute',top:15,right:15}} onClick={logout}>Logout</button>
      <Sider trigger={null} collapsible collapsed={collapse}>
        <SideNav />
      </Sider>
      <Layout>
        <Header className="siteLayoutBackground" style={{ padding: 0, background: "#001529" }}>
          {React.createElement(collapse ? MenuUnfoldOutlined : MenuFoldOutlined, {
            className: 'trigger',
            onClick: handleToggle,
            style: { color: "#fff"}
          })}
          <span style={{ color: 'white', marginLeft: '20px', fontFamily: 'Arial, sans-serif', fontWeight: 'bold' }}>Pharmacy Order & Operations System</span>
        </Header>
        <Content style={{ margin: '24px 16px', padding: 24, minHeight: "calc(100vh - 114px)", background: "#fff" }}>
          {children}
        </Content>
      </Layout>
    </Layout>
  );
};

export default MainLayout;
