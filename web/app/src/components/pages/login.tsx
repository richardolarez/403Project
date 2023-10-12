import React from 'react';
import { Form, Input, Button, Row, Col } from 'antd';
import { UserOutlined, LockOutlined } from '@ant-design/icons';

const Login: React.FC = () => {
  const onFinish = (values: any) => {
    console.log('Received values:', values);
    // You can add your authentication logic here
  };

  return (
    <Row justify="center" align="middle" style={{ minHeight: '100vh' }}>
      <Col span={8}>
        <Form
          name="login"
          initialValues={{ remember: true }}
          onFinish={onFinish}
          style={{ background: 'white', padding: '20px', boxShadow: '0px 0px 10px rgba(0, 0, 0, 0.2)' }}
        >
          <Form.Item
            name="username"
            rules={[{ required: true, message: 'Please enter your username!' }]}
          >
            <Input prefix={<UserOutlined />} placeholder="Username" />
          </Form.Item>

          <Form.Item
            name="password"
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
  );
};

export default Login;