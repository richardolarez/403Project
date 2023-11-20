import React, { useState } from 'react';
import { Layout, Button } from 'antd';
import SideNav from './sidebar';
import {
  MenuUnfoldOutlined,
  MenuFoldOutlined
} from '@ant-design/icons';
import {useNavigate} from 'react-router';

const { Header, Sider, Content } = Layout;

interface LayoutProps {
  children: React.ReactNode;
}

const MainLayout: React.FC<LayoutProps> = ({ children }) => {
  const [collapse, setCollapse] = useState(false);

  const handleToggle = (event: any) => {
    event.preventDefault();
    setCollapse(!collapse);
  };

  const bigLogout = () => {
    console.log('what the heck is going on here');
    sessionStorage.removeItem('authenticated');
    sessionStorage.removeItem("UserFName");
    sessionStorage.removeItem("UserRole");
    window.location.reload();
  };

  return (
    <Layout>
      <Button id="bigLogout" style={{ position: 'absolute', top: 15, right: 15 }} onClick={bigLogout}>Logout</Button>
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
