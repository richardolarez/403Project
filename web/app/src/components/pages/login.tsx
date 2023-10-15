import React from 'react';
import { Form, Input, Button, Row, Col } from 'antd';
import { UserOutlined, LockOutlined } from '@ant-design/icons';

const Login: React.FC = () => {
  const onFinish = (values: { username: string, password: string }) => {
    // Call the /login endpoint to authenticate the employee
    fetch('http://localhost:8080/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ username: values.username, password: values.password })
    })
    .then(response => {
      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }
      return response.json();
    })
    .then(employee => {
      console.log('Authenticated employee:', employee);
      sessionStorage.setItem('authenticated', 'true');
      sessionStorage.setItem('UserFName', employee.FirstName);
      window.location.reload();
    })
    .catch(error => {
      console.error('Failed to authenticate employee:', error);
      // TODO: Handle authentication error
    });
  };

  return (
  <div>  
    
  <Row justify="center" align="middle" style={{ minHeight: '100vh' }}>
      <Col span={8}>
        <img src="/poos_logo_full.png" alt="Logo" style={{width: '100%', marginBottom: '10%'}}/>
          <Form
          name="login"
          initialValues={{ remember: true }}
          onFinish={onFinish}
          style={{ background: 'white', padding: '20px', boxShadow: '0px 0px 10px rgba(0, 0, 0, 0.2)' }}
        >
          <Form.Item
            name="username" label="Username"
            rules={[{ required: true, message: 'Please enter your username!' }]}
          >
            <Input prefix={<UserOutlined />} placeholder="Username" />
          </Form.Item>

          <Form.Item
            name="password" label="Password"
            rules={[{ required: true, message: 'Please enter your password!' }]}
          >
            <Input.Password prefix={<LockOutlined />} placeholder="Password" />
          </Form.Item>

          <Form.Item>
            <Button type="primary" htmlType="submit" style={{ width: '100%' }}>
              Log In
            </Button>
          </Form.Item>
        </Form>
      </Col>
    </Row>
    </div>
  );
};

export default Login;